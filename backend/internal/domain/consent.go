package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ConsentStatus string
type PurposeStatus string

const (
	ConsentActive  ConsentStatus = "active"
	ConsentRevoked ConsentStatus = "revoked"
	ConsentExpired ConsentStatus = "expired"

	PurposeActive   PurposeStatus = "active"
	PurposeRevoked  PurposeStatus = "revoked"
	PurposeFulfilled PurposeStatus = "fulfilled"
	PurposeExpired  PurposeStatus = "expired"
)

type Consent struct {
	ID              uuid.UUID     `db:"id"`
	CustomerID      uuid.UUID     `db:"customer_id"`
	Status          ConsentStatus `db:"status"`
	Channel         string        `db:"channel"`
	IPHash          string        `db:"ip_hash"`
	UserAgentHash   string        `db:"user_agent_hash"`
	ReceiptID       uuid.NullUUID `db:"receipt_id"`
	GrantedAt       time.Time     `db:"granted_at"`
	RevokedAt       sql.NullTime  `db:"revoked_at"`
	RevocationReason sql.NullString `db:"revocation_reason"`
	RevokedBy       sql.NullString `db:"revoked_by"`
}

type ConsentPurpose struct {
	ID                uuid.UUID     `db:"id"`
	ConsentID         uuid.UUID     `db:"consent_id"`
	PurposeCode       string        `db:"purpose_code"`
	Description       string        `db:"description"`
	DataCategories    []string      `db:"data_categories"`
	RetentionDays     int           `db:"retention_days"`
	ThirdPartySharing bool          `db:"third_party_sharing"`
	Status            PurposeStatus `db:"status"`
	ExpiresAt         time.Time     `db:"expires_at"`
	RevokedAt         sql.NullTime  `db:"revoked_at"`
	CreatedAt         time.Time     `db:"created_at"`
}

type ConsentVerification struct {
	Authorized bool
	ConsentID  uuid.UUID
	GrantedAt  time.Time
	ExpiresAt  time.Time
	Reason     string // populated when Authorized == false
}

type CaptureConsentInput struct {
	CustomerID    uuid.UUID
	Channel       string
	IPHash        string
	UserAgentHash string
	Purposes      []PurposeInput
}

type PurposeInput struct {
	PurposeCode       string
	Description       string
	DataCategories    []string
	RetentionDays     int
	ThirdPartySharing bool
}

type RevokeConsentInput struct {
	Reason    string
	RevokedBy string // customer | rm | compliance | system
}
