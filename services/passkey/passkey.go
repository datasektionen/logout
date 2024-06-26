package passkey

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/datasektionen/logout/pkg/config"
	"github.com/datasektionen/logout/pkg/database"
	"github.com/datasektionen/logout/pkg/httputil"
	"github.com/datasektionen/logout/services/passkey/export"
	user "github.com/datasektionen/logout/services/user/export"
	"github.com/go-webauthn/webauthn/webauthn"
)

type service struct {
	db       *database.Queries
	webauthn *webauthn.WebAuthn
	user     user.Service
}

func NewService(db *database.Queries) (*service, error) {
	wa, err := webauthn.New(&webauthn.Config{
		RPID:          config.Config.Origin.Hostname(),
		RPDisplayName: "Konglig Datasektionen",
		RPOrigins:     []string{config.Config.Origin.String()},
	})
	if err != nil {
		return nil, err
	}

	s := &service{db: db, webauthn: wa}

	http.Handle("POST /login/passkey/begin", httputil.Route(s.beginLoginPasskey))
	http.Handle("POST /login/passkey/finish", httputil.Route(s.finishLoginPasskey))
	http.Handle("POST /passkey/add/begin", httputil.Route(s.beginAddPasskey))
	http.Handle("POST /passkey/add/finish", httputil.Route(s.finishAddPasskey))
	http.Handle("POST /passkey/remove", httputil.Route(s.removePasskey))
	http.Handle("GET /passkey/list", httputil.Route(s.listPasskeys))

	return s, nil
}

func (s *service) Assign(user user.Service) {
	s.user = user
}

func (s *service) listPasskeysForUser(ctx context.Context, kthid string) ([]export.Passkey, error) {
	dbPasskeys, err := s.db.ListPasskeysByUser(ctx, kthid)
	if err != nil {
		return nil, err
	}
	passkeys := make([]export.Passkey, len(dbPasskeys))
	for i, passkey := range dbPasskeys {
		var c webauthn.Credential
		if err := json.Unmarshal([]byte(passkey.Data), &c); err != nil {
			return nil, err
		}
		passkeys[i] = export.Passkey{
			ID:   passkey.ID,
			Name: passkey.Name,
			Cred: c,
		}
	}
	return passkeys, nil
}
