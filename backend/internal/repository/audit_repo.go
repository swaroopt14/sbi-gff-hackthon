package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/swaroopt14/sbi-gff-hackthon/backend/internal/domain"
)

// AuditRepo is append-only — no Update or Delete methods by design.
type AuditRepo struct {
	db *sql.DB
}

func NewAuditRepo(db *sql.DB) *AuditRepo {
	return &AuditRepo{db: db}
}

func (r *AuditRepo) Log(ctx context.Context, in domain.LogEventInput) (*domain.AuditLog, error) {
	fullInput, _ := json.Marshal(in.FullInput)
	fullOutput, _ := json.Marshal(in.FullOutput)

	riskLevel := in.RiskLevel
	if riskLevel == "" {
		riskLevel = domain.RiskLow
	}

	const q = `
		INSERT INTO audit_logs (
			event_type, entity_type, entity_id, actor_type, actor_id, actor_name,
			description, input_summary, output_summary, full_input, full_output,
			policy_applied, risk_level, flags, trace_id, timestamp
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11,
			$12, $13, '[]', $14, NOW()
		)
		RETURNING id, timestamp`

	log := &domain.AuditLog{
		EventType:     in.EventType,
		EntityType:    in.EntityType,
		ActorType:     in.ActorType,
		ActorName:     in.ActorName,
		Description:   in.Description,
		InputSummary:  in.InputSummary,
		OutputSummary: in.OutputSummary,
		PolicyApplied: in.PolicyApplied,
		RiskLevel:     riskLevel,
	}
	if in.EntityID != nil {
		log.EntityID = uuid.NullUUID{UUID: *in.EntityID, Valid: true}
	}
	if in.ActorID != nil {
		log.ActorID = uuid.NullUUID{UUID: *in.ActorID, Valid: true}
	}
	if in.TraceID != nil {
		log.TraceID = uuid.NullUUID{UUID: *in.TraceID, Valid: true}
	}

	err := r.db.QueryRowContext(ctx, q,
		log.EventType, log.EntityType, log.EntityID, log.ActorType, log.ActorID, log.ActorName,
		log.Description, log.InputSummary, log.OutputSummary, fullInput, fullOutput,
		log.PolicyApplied, log.RiskLevel, log.TraceID,
	).Scan(&log.ID, &log.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("audit log: %w", err)
	}
	return log, nil
}

type ListAuditFilter struct {
	CustomerID *uuid.UUID
	EntityType string
	EventType  string
	RiskLevel  string
	StartDate  *time.Time
	EndDate    *time.Time
	Page       int
	Limit      int
}

func (r *AuditRepo) List(ctx context.Context, f ListAuditFilter) ([]*domain.AuditLog, int, error) {
	if f.Limit <= 0 || f.Limit > 200 {
		f.Limit = 50
	}
	if f.Page <= 0 {
		f.Page = 1
	}
	offset := (f.Page - 1) * f.Limit

	args := []interface{}{}
	where := "WHERE 1=1"
	argIdx := 1

	if f.EntityType != "" {
		where += fmt.Sprintf(" AND entity_type = $%d", argIdx)
		args = append(args, f.EntityType)
		argIdx++
	}
	if f.EventType != "" {
		where += fmt.Sprintf(" AND event_type = $%d", argIdx)
		args = append(args, f.EventType)
		argIdx++
	}
	if f.RiskLevel != "" {
		where += fmt.Sprintf(" AND risk_level = $%d", argIdx)
		args = append(args, f.RiskLevel)
		argIdx++
	}
	if f.StartDate != nil {
		where += fmt.Sprintf(" AND timestamp >= $%d", argIdx)
		args = append(args, f.StartDate)
		argIdx++
	}
	if f.EndDate != nil {
		where += fmt.Sprintf(" AND timestamp <= $%d", argIdx)
		args = append(args, f.EndDate)
		argIdx++
	}

	var total int
	countQ := "SELECT COUNT(*) FROM audit_logs " + where
	if err := r.db.QueryRowContext(ctx, countQ, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count audit logs: %w", err)
	}

	listQ := fmt.Sprintf(`
		SELECT id, event_type, entity_type, entity_id, actor_type, actor_id, actor_name,
		       description, input_summary, output_summary, policy_applied,
		       risk_level, flags, trace_id, timestamp
		FROM audit_logs
		%s
		ORDER BY timestamp DESC
		LIMIT $%d OFFSET $%d`, where, argIdx, argIdx+1)

	args = append(args, f.Limit, offset)
	rows, err := r.db.QueryContext(ctx, listQ, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list audit logs: %w", err)
	}
	defer rows.Close()

	var logs []*domain.AuditLog
	for rows.Next() {
		l := &domain.AuditLog{}
		if err := rows.Scan(
			&l.ID, &l.EventType, &l.EntityType, &l.EntityID,
			&l.ActorType, &l.ActorID, &l.ActorName,
			&l.Description, &l.InputSummary, &l.OutputSummary,
			&l.PolicyApplied, &l.RiskLevel, &l.Flags, &l.TraceID, &l.Timestamp,
		); err != nil {
			return nil, 0, fmt.Errorf("scan audit log: %w", err)
		}
		logs = append(logs, l)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("iterate audit logs: %w", err)
	}

	return logs, total, nil
}

func (r *AuditRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.AuditLog, error) {
	const q = `
		SELECT id, event_type, entity_type, entity_id, actor_type, actor_id, actor_name,
		       description, input_summary, output_summary, full_input, full_output,
		       policy_applied, risk_level, flags, trace_id, timestamp
		FROM audit_logs WHERE id = $1`

	l := &domain.AuditLog{}
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&l.ID, &l.EventType, &l.EntityType, &l.EntityID,
		&l.ActorType, &l.ActorID, &l.ActorName,
		&l.Description, &l.InputSummary, &l.OutputSummary,
		&l.FullInput, &l.FullOutput,
		&l.PolicyApplied, &l.RiskLevel, &l.Flags, &l.TraceID, &l.Timestamp,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get audit log %s: %w", id, err)
	}
	return l, nil
}
