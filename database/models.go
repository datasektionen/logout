// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Invite struct {
	ID          uuid.UUID
	Name        string
	CreatedAt   pgtype.Timestamp
	ExpiresAt   pgtype.Timestamp
	MaxUses     pgtype.Int4
	CurrentUses int32
}

type LegacyapiToken struct {
	ID         uuid.UUID
	Kthid      string
	LastUsedAt pgtype.Timestamp
}

type OidcClient struct {
	ID           []byte
	RedirectUris []string
}

type Passkey struct {
	ID    uuid.UUID
	Name  string
	Kthid string
	Data  string
}

type Session struct {
	ID         uuid.UUID
	Kthid      string
	LastUsedAt pgtype.Timestamp
}

type User struct {
	Kthid                   string
	UgKthid                 string
	Email                   string
	FirstName               string
	FamilyName              string
	YearTag                 string
	MemberTo                pgtype.Date
	WebauthnID              []byte
	FirstNameChangeRequest  string
	FamilyNameChangeRequest string
}

type WebauthnSessionDatum struct {
	Kthid string
	Data  []byte
}
