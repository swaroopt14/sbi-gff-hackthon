package domain

import (
	"time"

	"github.com/google/uuid"
)

type IntentEventType string
type IntentCategory string

const (
	EventPageView      IntentEventType = "page_view"
	EventClick         IntentEventType = "click"
	EventDwell         IntentEventType = "dwell"
	EventEMICalc       IntentEventType = "emi_calc"
	EventSearch        IntentEventType = "search"
	EventCampaignClick IntentEventType = "campaign_click"
	EventReturnVisit   IntentEventType = "return_visit"
	EventAbandon       IntentEventType = "abandon"
	EventFormStart     IntentEventType = "form_start"
	EventFormAbandon   IntentEventType = "form_abandon"

	CategoryProductDiscovery IntentCategory = "product_discovery"
	CategoryActiveComparison IntentCategory = "active_comparison"
	CategoryReadyToApply     IntentCategory = "ready_to_apply"
	CategoryDropoffRecovery  IntentCategory = "dropoff_recovery"
)

type IntentSession struct {
	ID             uuid.UUID      `db:"id"`
	CustomerID     uuid.NullUUID  `db:"customer_id"`
	SessionToken   string         `db:"session_token"`
	Channel        string         `db:"channel"`
	CampaignSource string         `db:"campaign_source"`
	DeviceType     string         `db:"device_type"`
	ProductPages   []string       `db:"product_pages"`
	EventCount     int            `db:"event_count"`
	StartedAt      time.Time      `db:"started_at"`
	LastEventAt    time.Time      `db:"last_event_at"`
	CreatedAt      time.Time      `db:"created_at"`
}

type IntentEvent struct {
	ID        uuid.UUID       `db:"id"`
	SessionID uuid.UUID       `db:"session_id"`
	EventType IntentEventType `db:"event_type"`
	Page      string          `db:"page"`
	Product   string          `db:"product"`
	Metadata  []byte          `db:"metadata"` // JSONB
	Timestamp time.Time       `db:"timestamp"`
}

type IntentScore struct {
	ID                uuid.UUID      `db:"id"`
	SessionID         uuid.UUID      `db:"session_id"`
	CustomerID        uuid.NullUUID  `db:"customer_id"`
	Score             int            `db:"score"`
	Tier              CustomerTier   `db:"tier"`
	Category          IntentCategory `db:"category"`
	SignalsDetail     []byte         `db:"signals_detail"` // JSONB
	Explanation       string         `db:"explanation"`
	RecommendedProduct string        `db:"recommended_product"`
	ShouldTriggerChat bool           `db:"should_trigger_chat"`
	AgentVersion      string         `db:"agent_version"`
	ComputedAt        time.Time      `db:"computed_at"`
}

// IntentScoreResult is returned from the Intent Agent
type IntentScoreResult struct {
	SessionID          string         `json:"session_id"`
	Score              int            `json:"intent_score"`
	Tier               CustomerTier   `json:"intent_tier"`
	Category           IntentCategory `json:"intent_category"`
	ShouldTriggerChat  bool           `json:"should_trigger_chat"`
	RecommendedProduct string         `json:"recommended_product"`
	Explanation        string         `json:"explanation"`
	ComputedAt         time.Time      `json:"computed_at"`
}
