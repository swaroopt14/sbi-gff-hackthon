package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CustomerStatus string
type CustomerTier string
type CustomerSource string
type OccupationType string

const (
	StatusProspect   CustomerStatus = "prospect"
	StatusQualified  CustomerStatus = "qualified"
	StatusApplied    CustomerStatus = "applied"
	StatusOnboarding CustomerStatus = "onboarding"
	StatusOnboarded  CustomerStatus = "onboarded"
	StatusRejected   CustomerStatus = "rejected"

	TierHot  CustomerTier = "hot"
	TierWarm CustomerTier = "warm"
	TierCold CustomerTier = "cold"

	SourceWeb      CustomerSource = "web"
	SourceYONO     CustomerSource = "yono"
	SourceCampaign CustomerSource = "campaign"
	SourceBranch   CustomerSource = "branch"
	SourceRM       CustomerSource = "rm_manual"

	OccupationSalaried    OccupationType = "salaried"
	OccupationSelfEmployed OccupationType = "self_employed"
	OccupationBusiness    OccupationType = "business"
	OccupationRetired     OccupationType = "retired"
	OccupationStudent     OccupationType = "student"
)

type Customer struct {
	ID                   uuid.UUID      `db:"id"`
	UserID               uuid.NullUUID  `db:"user_id"`
	Name                 sql.NullString `db:"name"`
	PhoneEncrypted       []byte         `db:"phone_encrypted"`
	EmailEncrypted       []byte         `db:"email_encrypted"`
	Age                  sql.NullInt16  `db:"age"`
	Occupation           sql.NullString `db:"occupation"`
	IncomeMonthly        sql.NullInt32  `db:"income_monthly"`
	Location             sql.NullString `db:"location"`
	State                sql.NullString `db:"state"`
	ExistingSBICustomer  bool           `db:"existing_sbi_customer"`
	SBICIFId             sql.NullString `db:"sbi_cif_id"`
	LanguagePreference   string         `db:"language_preference"`
	Status               CustomerStatus `db:"status"`
	Tier                 sql.NullString `db:"tier"`
	IntentScore          sql.NullInt32  `db:"intent_score"`
	QualificationScore   sql.NullInt32  `db:"qualification_score"`
	Source               CustomerSource `db:"source"`
	CampaignSource       sql.NullString `db:"campaign_source"`
	RMID                 uuid.NullUUID  `db:"rm_id"`
	CreatedAt            time.Time      `db:"created_at"`
	UpdatedAt            time.Time      `db:"updated_at"`
}

type CreateCustomerInput struct {
	Name           string
	PhoneEncrypted []byte
	EmailEncrypted []byte
	Source         CustomerSource
	SessionID      *uuid.UUID
	ProductInterest string
}

type UpdateCustomerInput struct {
	Name          *string
	Age           *int16
	Occupation    *OccupationType
	IncomeMonthly *int32
	Location      *string
	State         *string
	Status        *CustomerStatus
	Tier          *CustomerTier
	IntentScore   *int32
	QualScore     *int32
	RMID          *uuid.UUID
}
