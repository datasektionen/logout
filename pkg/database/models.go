// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

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
	Kthid      string
	UgKthid    string
	Email      string
	FirstName  string
	FamilyName string
	YearTag    string
	MemberTo   pgtype.Date
	WebauthnID []byte
}
