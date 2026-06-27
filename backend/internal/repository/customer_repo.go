package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/swaroopt14/sbi-gff-hackthon/backend/internal/domain"
)

type CustomerRepo struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (r *CustomerRepo) Create(ctx context.Context, in domain.CreateCustomerInput) (*domain.Customer, error) {
	const q = `
		INSERT INTO customers (name, phone_encrypted, email_encrypted, source, language_preference, status)
		VALUES ($1, $2, $3, $4, 'en', 'prospect')
		RETURNING id, name, phone_encrypted, email_encrypted, existing_sbi_customer,
		          language_preference, status, source, created_at, updated_at`

	c := &domain.Customer{}
	err := r.db.QueryRowContext(ctx, q,
		nullString(in.Name),
		in.PhoneEncrypted,
		in.EmailEncrypted,
		in.Source,
	).Scan(
		&c.ID, &c.Name, &c.PhoneEncrypted, &c.EmailEncrypted,
		&c.ExistingSBICustomer, &c.LanguagePreference, &c.Status,
		&c.Source, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if isPQUniqueViolation(err) {
			return nil, domain.ErrConflict
		}
		return nil, fmt.Errorf("create customer: %w", err)
	}
	return c, nil
}

func (r *CustomerRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	const q = `
		SELECT id, user_id, name, phone_encrypted, email_encrypted, age, occupation,
		       income_monthly, location, state, existing_sbi_customer, sbi_cif_id,
		       language_preference, status, tier, intent_score, qualification_score,
		       source, campaign_source, rm_id, created_at, updated_at
		FROM customers
		WHERE id = $1`

	c := &domain.Customer{}
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&c.ID, &c.UserID, &c.Name, &c.PhoneEncrypted, &c.EmailEncrypted,
		&c.Age, &c.Occupation, &c.IncomeMonthly, &c.Location, &c.State,
		&c.ExistingSBICustomer, &c.SBICIFId, &c.LanguagePreference,
		&c.Status, &c.Tier, &c.IntentScore, &c.QualificationScore,
		&c.Source, &c.CampaignSource, &c.RMID, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get customer %s: %w", id, err)
	}
	return c, nil
}

type ListCustomersFilter struct {
	Status    string
	Tier      string
	ProductID string
	Search    string
	Page      int
	Limit     int
	SortBy    string
	SortDir   string
}

type ListCustomersResult struct {
	Customers  []*domain.Customer
	Total      int
}

func (r *CustomerRepo) List(ctx context.Context, f ListCustomersFilter) (*ListCustomersResult, error) {
	if f.Limit <= 0 || f.Limit > 100 {
		f.Limit = 20
	}
	if f.Page <= 0 {
		f.Page = 1
	}
	offset := (f.Page - 1) * f.Limit

	args := []interface{}{}
	where := "WHERE 1=1"
	argIdx := 1

	if f.Status != "" {
		where += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, f.Status)
		argIdx++
	}
	if f.Tier != "" {
		where += fmt.Sprintf(" AND tier = $%d", argIdx)
		args = append(args, f.Tier)
		argIdx++
	}

	orderBy := "created_at DESC"
	if f.SortBy == "intent_score" || f.SortBy == "qualification_score" {
		dir := "DESC"
		if f.SortDir == "asc" {
			dir = "ASC"
		}
		orderBy = fmt.Sprintf("%s %s", f.SortBy, dir)
	}

	countQ := "SELECT COUNT(*) FROM customers " + where
	var total int
	if err := r.db.QueryRowContext(ctx, countQ, args...).Scan(&total); err != nil {
		return nil, fmt.Errorf("count customers: %w", err)
	}

	listQ := fmt.Sprintf(`
		SELECT id, name, status, tier, intent_score, qualification_score,
		       location, source, created_at, updated_at
		FROM customers
		%s
		ORDER BY %s
		LIMIT $%d OFFSET $%d`, where, orderBy, argIdx, argIdx+1)

	args = append(args, f.Limit, offset)
	rows, err := r.db.QueryContext(ctx, listQ, args...)
	if err != nil {
		return nil, fmt.Errorf("list customers: %w", err)
	}
	defer rows.Close()

	var customers []*domain.Customer
	for rows.Next() {
		c := &domain.Customer{}
		if err := rows.Scan(
			&c.ID, &c.Name, &c.Status, &c.Tier,
			&c.IntentScore, &c.QualificationScore,
			&c.Location, &c.Source, &c.CreatedAt, &c.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan customer row: %w", err)
		}
		customers = append(customers, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate customers: %w", err)
	}

	return &ListCustomersResult{Customers: customers, Total: total}, nil
}

func (r *CustomerRepo) UpdateScores(ctx context.Context, id uuid.UUID, intentScore, qualScore *int32, tier *domain.CustomerTier) error {
	const q = `
		UPDATE customers
		SET intent_score = COALESCE($2, intent_score),
		    qualification_score = COALESCE($3, qualification_score),
		    tier = COALESCE($4, tier),
		    updated_at = NOW()
		WHERE id = $1`

	_, err := r.db.ExecContext(ctx, q, id, intentScore, qualScore, tier)
	if err != nil {
		return fmt.Errorf("update customer scores %s: %w", id, err)
	}
	return nil
}

// helpers

func nullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

func isPQUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505"
	}
	return false
}

func mustMarshal(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}
