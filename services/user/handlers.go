package user

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/datasektionen/logout/pkg/database"
	"github.com/datasektionen/logout/pkg/httputil"
	"github.com/datasektionen/logout/pkg/kthldap"
	"github.com/datasektionen/logout/pkg/pls"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const nextUrlCookie string = "_logout_next-url"

func validNextURL(url string) bool {
	if url == "" {
		return true
	}
	// The "url" must be a path (and possibly params)
	if len(url) > 0 && url[0] != '/' {
		return false
	}
	// If it starts with two slashes there will be an implicit `https:` in front, so then it's not a path
	if len(url) > 1 && url[1] == '/' {
		return false
	}
	return true
}

func (s *service) index(w http.ResponseWriter, r *http.Request) httputil.ToResponse {
	nextURL := r.FormValue("next-url")
	if !validNextURL(nextURL) {
		return httputil.BadRequest("Invalid return url")
	}
	hasCookie := false
	if nextURL == "" {
		c, _ := r.Cookie(nextUrlCookie)
		if c != nil {
			nextURL = c.Value
			hasCookie = true
		}
	}
	if nextURL == "" {
		nextURL = "/account"
	}
	if kthid, err := s.GetLoggedInKTHID(r); err != nil {
		return err
	} else if kthid != "" {
		if hasCookie {
			http.SetCookie(w, &http.Cookie{Name: nextUrlCookie, MaxAge: -1})
		}
		http.Redirect(w, r, nextURL, http.StatusSeeOther)
		return nil
	}
	if nextURL != "" {
		http.SetCookie(w, &http.Cookie{
			Name:     nextUrlCookie,
			Value:    nextURL,
			MaxAge:   int((time.Minute * 10).Seconds()),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
	}
	return index(s.passkey.PasskeyLogin(), s.dev.LoginForm)
}

func (s *service) account(w http.ResponseWriter, r *http.Request) httputil.ToResponse {
	user, err := s.GetLoggedInUser(r)
	if err != nil {
		return err
	}
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}
	passkeySettings, err := s.passkey.PasskeySettings(r.Context(), user.KTHID)
	if err != nil {
		return err
	}
	isAdmin, err := pls.CheckUser(r.Context(), user.KTHID, "admin-read")
	return account(*user, passkeySettings)
}

func (s *service) acceptInvite(w http.ResponseWriter, r *http.Request) httputil.ToResponse {
	idString := r.PathValue("id")
	if idString == "-" {
		idCookie, _ := r.Cookie("invite")
		if idCookie == nil {
			return httputil.BadRequest("No invite id found")
		}
		idString = idCookie.Value
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		return httputil.BadRequest("Invalid uuid")
	}
	inv, err := s.db.GetInvite(r.Context(), id)
	if err == pgx.ErrNoRows {
		return httputil.BadRequest("No such invite")
	} else if err != nil {
		return err
	}
	if time.Now().After(inv.ExpiresAt.Time) {
		return httputil.BadRequest("Invite expired")
	}
	if inv.MaxUses.Valid && inv.CurrentUses >= inv.MaxUses.Int32 {
		return httputil.BadRequest("This invite cannot be used to create more users")
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "invite",
		Value:    id.String(),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	return acceptInvite()
}

func (s *service) FinishInvite(w http.ResponseWriter, r *http.Request, kthid string) (bool, httputil.ToResponse) {
	idCookie, _ := r.Cookie("invite")
	if idCookie == nil {
		return false, nil
	}
	id, err := uuid.Parse(idCookie.Value)
	if err != nil {
		return true, httputil.BadRequest("Invalid uuid")
	}
	inv, err := s.db.GetInvite(r.Context(), id)
	if err == pgx.ErrNoRows {
		return true, httputil.BadRequest("No such invite")
	} else if err != nil {
		return true, err
	}
	if time.Now().After(inv.ExpiresAt.Time) {
		return true, httputil.BadRequest("Invite expired")
	}
	if inv.MaxUses.Valid && inv.CurrentUses >= inv.MaxUses.Int32 {
		return true, httputil.BadRequest("This invite has reached its usage limit")
	}
	person, err := kthldap.Lookup(r.Context(), kthid)
	if err != nil {
		return true, err
	}
	if person == nil {
		slog.Error("Could not find user in ldap", "kthid", kthid, "invite id", id)
		return true, errors.New("Could not find user in ldap")
	}
	tx, err := s.db.Begin(r.Context())
	if err != nil {
		return true, err
	}
	defer tx.Rollback(r.Context())
	if err := tx.CreateUser(r.Context(), database.CreateUserParams{
		Kthid:      kthid,
		UgKthid:    person.UGKTHID,
		Email:      kthid + "@kth.se",
		FirstName:  person.FirstName,
		FamilyName: person.FamilyName,
	}); err != nil {
		return true, err
	}
	if err := tx.IncrementInviteUses(r.Context(), id); err != nil {
		return true, err
	}
	if err := tx.Commit(r.Context()); err != nil {
		return true, err
	}
	http.SetCookie(w, &http.Cookie{Name: "invite", MaxAge: -1})
	slog.Info("User invite link used", "kthid", kthid, "invite-id", inv.ID)
	return true, s.LoginUser(r.Context(), kthid)
}
