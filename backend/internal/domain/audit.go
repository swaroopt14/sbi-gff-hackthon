package domain

import (
	"time"

	"github.com/google/uuid"
)

type AuditEventType string
type AuditActorType string
type RiskLevel string

const (
	EventIntentDetected      AuditEventType = "intent_detected"
	EventConsentGranted      AuditEventType = "consent_granted"
	EventConsentRevoked      AuditEventType = "consent_revoked"
	EventQualificationScored AuditEventType = "qualification_scored"
	EventQualificationOverride AuditEventType = "qualification_override"
	EventDocumentUploaded    AuditEventType = "document_uploaded"
	EventDocumentExtracted   AuditEventType = "document_extracted"
	EventChatStarted         AuditEventType = "chat_started"
	EventChatCompleted       AuditEventType = "chat_completed"
	EventApplicationSubmitted AuditEventType = "application_submitted"
	EventApplicationDecided  AuditEventType = "application_decided"
	EventComplianceCheck     AuditEventType = "compliance_check"
	EventRMOverride          AuditEventType = "rm_override"

	ActorSystem  AuditActorType = "system"
	ActorUser    AuditActorType = "user"
	ActorAgent   AuditActorType = "agent"

	RiskLow      RiskLevel = "low"
	RiskMedium   RiskLevel = "medium"
	RiskHigh     RiskLevel = "high"
	RiskCritical RiskLevel = "critical"
)

type AuditLog struct {
	ID            uuid.UUID      `db:"id"`
	EventType     AuditEventType `db:"event_type"`
	EntityType    string         `db:"entity_type"`
	EntityID      uuid.NullUUID  `db:"entity_id"`
	ActorType     AuditActorType `db:"actor_type"`
	ActorID       uuid.NullUUID  `db:"actor_id"`
	ActorName     string         `db:"actor_name"`
	Description   string         `db:"description"`
	InputSummary  string         `db:"input_summary"`
	OutputSummary string         `db:"output_summary"`
	FullInput     []byte         `db:"full_input"`  // JSONB
	FullOutput    []byte         `db:"full_output"` // JSONB
	PolicyApplied string         `db:"policy_applied"`
	RiskLevel     RiskLevel      `db:"risk_level"`
	Flags         []byte         `db:"flags"` // JSONB
	TraceID       uuid.NullUUID  `db:"trace_id"`
	Timestamp     time.Time      `db:"timestamp"`
}

type LogEventInput struct {
	EventType     AuditEventType
	EntityType    string
	EntityID      *uuid.UUID
	ActorType     AuditActorType
	ActorID       *uuid.UUID
	ActorName     string
	Description   string
	InputSummary  string
	OutputSummary string
	FullInput     interface{}
	FullOutput    interface{}
	PolicyApplied string
	RiskLevel     RiskLevel
	TraceID       *uuid.UUID
}
