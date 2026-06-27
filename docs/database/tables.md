# Aperture — Database Table Definitions

**Version:** 1.0  
**Date:** 2026-06-27  
**Database:** PostgreSQL 16

---

## Table: users

Primary identity store for all platform users (customers, RMs, compliance, admin).

```sql
CREATE TABLE users (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email           VARCHAR(255) UNIQUE NOT NULL,
  phone_hash      VARCHAR(64),              -- SHA-256 of phone number
  password_hash   VARCHAR(255) NOT NULL,
  name            VARCHAR(255) NOT NULL,
  role            VARCHAR(50) NOT NULL      -- customer, rm, compliance, operations, admin
                    CHECK (role IN ('customer','rm','compliance','operations','admin')),
  status          VARCHAR(50) NOT NULL DEFAULT 'active'
                    CHECK (status IN ('active','suspended','deleted')),
  avatar_url      TEXT,
  last_login_at   TIMESTAMPTZ,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
```

---

## Table: customers

Prospect and customer profile data. Separate from users to allow anonymous prospects.

```sql
CREATE TABLE customers (
  id                    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id               UUID REFERENCES users(id),  -- NULL for anonymous prospects
  name                  VARCHAR(255),
  phone_encrypted       BYTEA,                       -- AES-256 encrypted
  email_encrypted       BYTEA,
  age                   SMALLINT,
  occupation            VARCHAR(50)
                          CHECK (occupation IN ('salaried','self_employed','business','retired','student','other')),
  income_monthly        INTEGER,                     -- in INR
  location              VARCHAR(255),
  state                 VARCHAR(100),
  existing_sbi_customer BOOLEAN DEFAULT FALSE,
  sbi_cif_id            VARCHAR(20),                 -- NULL for non-SBI customers
  language_preference   VARCHAR(10) DEFAULT 'en',
  status                VARCHAR(50) NOT NULL DEFAULT 'prospect'
                          CHECK (status IN ('prospect','qualified','applied','onboarding','onboarded','rejected')),
  tier                  VARCHAR(20)
                          CHECK (tier IN ('hot','warm','cold')),
  intent_score          INTEGER,
  qualification_score   INTEGER,
  source                VARCHAR(50) DEFAULT 'web'
                          CHECK (source IN ('web','yono','campaign','branch','rm_manual','referral')),
  campaign_source       VARCHAR(255),
  rm_id                 UUID REFERENCES users(id),   -- Assigned RM
  created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: intent_sessions

Browser/app session that triggers intent detection.

```sql
CREATE TABLE intent_sessions (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id     UUID REFERENCES customers(id),   -- NULL for anonymous
  session_token   VARCHAR(255) UNIQUE NOT NULL,
  channel         VARCHAR(50) NOT NULL
                    CHECK (channel IN ('web','yono','mobile','branch','campaign')),
  campaign_source VARCHAR(255),
  device_type     VARCHAR(50),
  started_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  last_event_at   TIMESTAMPTZ,
  ended_at        TIMESTAMPTZ,
  event_count     INTEGER DEFAULT 0,
  product_pages   TEXT[],                           -- Products browsed
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: intent_events

Individual behavioral events within a session.

```sql
CREATE TABLE intent_events (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  session_id      UUID NOT NULL REFERENCES intent_sessions(id) ON DELETE CASCADE,
  event_type      VARCHAR(100) NOT NULL
                    CHECK (event_type IN (
                      'page_view','click','scroll','dwell','search',
                      'emi_calc','campaign_click','return_visit','abandon',
                      'form_start','form_abandon','document_hover'
                    )),
  page            VARCHAR(500),
  product         VARCHAR(100),
  metadata        JSONB DEFAULT '{}',
  timestamp       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: intent_scores

Computed intent score per session.

```sql
CREATE TABLE intent_scores (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  session_id      UUID NOT NULL REFERENCES intent_sessions(id),
  customer_id     UUID REFERENCES customers(id),
  score           INTEGER NOT NULL CHECK (score BETWEEN 0 AND 100),
  tier            VARCHAR(20) NOT NULL CHECK (tier IN ('hot','warm','cold')),
  category        VARCHAR(100) NOT NULL,
  signals_detail  JSONB NOT NULL DEFAULT '[]',
  explanation     TEXT,
  recommended_product VARCHAR(100),
  should_trigger_chat BOOLEAN DEFAULT FALSE,
  agent_version   VARCHAR(50),
  computed_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: qualification_scores

Multi-dimension qualification scoring per customer per product.

```sql
CREATE TABLE qualification_scores (
  id                    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id           UUID NOT NULL REFERENCES customers(id),
  product_id            UUID NOT NULL REFERENCES products(id),
  overall_score         INTEGER NOT NULL CHECK (overall_score BETWEEN 0 AND 100),
  tier                  VARCHAR(20) NOT NULL CHECK (tier IN ('hot','warm','cold')),
  
  -- Dimension scores (stored as JSONB for flexibility)
  dimensions            JSONB NOT NULL DEFAULT '{}',
  
  -- Recommended offer
  offered_amount        BIGINT,
  offered_rate          NUMERIC(5,2),
  offered_tenure_years  SMALLINT,
  emi_monthly           INTEGER,
  
  -- Metadata
  explanation_id        UUID,
  policy_version        VARCHAR(100),
  human_review_required BOOLEAN DEFAULT FALSE,
  data_sources          TEXT[],
  agent_version         VARCHAR(50),
  computed_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  
  -- Override (if any)
  overridden            BOOLEAN DEFAULT FALSE,
  override_id           UUID
);
```

---

## Table: rm_overrides

Human override of AI qualification decisions.

```sql
CREATE TABLE rm_overrides (
  id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  qualification_id    UUID NOT NULL REFERENCES qualification_scores(id),
  overridden_by       UUID NOT NULL REFERENCES users(id),
  original_score      INTEGER NOT NULL,
  overridden_score    INTEGER,
  original_tier       VARCHAR(20),
  overridden_tier     VARCHAR(20),
  action              VARCHAR(50) NOT NULL CHECK (action IN ('upgrade','downgrade','approve','reject')),
  reason              TEXT NOT NULL,
  created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: products

SBI product catalogue.

```sql
CREATE TABLE products (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  code            VARCHAR(50) UNIQUE NOT NULL,  -- HOME_LOAN, SAVINGS, FD, MF, CC
  name            VARCHAR(255) NOT NULL,
  category        VARCHAR(100) NOT NULL,
  description     TEXT,
  min_amount      BIGINT,
  max_amount      BIGINT,
  base_rate       NUMERIC(5,2),
  min_age         SMALLINT,
  max_age         SMALLINT,
  min_income      INTEGER,
  features        JSONB DEFAULT '[]',
  is_active       BOOLEAN DEFAULT TRUE,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: product_recommendations

Personalised product recommendations for each customer.

```sql
CREATE TABLE product_recommendations (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id     UUID NOT NULL REFERENCES customers(id),
  product_id      UUID NOT NULL REFERENCES products(id),
  rank            SMALLINT NOT NULL,
  fit_score       INTEGER NOT NULL CHECK (fit_score BETWEEN 0 AND 100),
  offer_message   TEXT,
  offer_amount    BIGINT,
  offer_rate      NUMERIC(5,2),
  reasons         JSONB DEFAULT '[]',
  language        VARCHAR(10) DEFAULT 'en',
  shown_to_customer BOOLEAN DEFAULT FALSE,
  customer_response VARCHAR(50),  -- accepted | rejected | ignored
  agent_version   VARCHAR(50),
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: conversations

Chat conversation session.

```sql
CREATE TABLE conversations (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id     UUID REFERENCES customers(id),
  session_id      UUID REFERENCES intent_sessions(id),
  channel         VARCHAR(50) NOT NULL,
  language        VARCHAR(10) NOT NULL DEFAULT 'en',
  product_intent  VARCHAR(100),
  status          VARCHAR(50) NOT NULL DEFAULT 'active'
                    CHECK (status IN ('active','completed','abandoned','transferred')),
  form_completion NUMERIC(3,2) DEFAULT 0,   -- 0.0 to 1.0
  collected_data  JSONB DEFAULT '{}',       -- Extracted form fields
  transferred_to  UUID REFERENCES users(id), -- RM if transferred
  transfer_reason VARCHAR(255),
  started_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  ended_at        TIMESTAMPTZ,
  last_message_at TIMESTAMPTZ
);
```

---

## Table: conversation_turns

Individual message turns in a conversation.

```sql
CREATE TABLE conversation_turns (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  conversation_id UUID NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
  role            VARCHAR(20) NOT NULL CHECK (role IN ('user','assistant','system')),
  content         TEXT NOT NULL,
  message_type    VARCHAR(50) DEFAULT 'text',
  extracted_fields JSONB DEFAULT '{}',
  quick_replies   JSONB DEFAULT '[]',
  action          JSONB,
  tokens_used     INTEGER,
  agent_version   VARCHAR(50),
  timestamp       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: documents

Uploaded documents metadata.

```sql
CREATE TABLE documents (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id     UUID NOT NULL REFERENCES customers(id),
  consent_id      UUID NOT NULL REFERENCES consents(id),
  document_type   VARCHAR(100) NOT NULL
                    CHECK (document_type IN ('pan','aadhaar_offline_xml','salary_slip','bank_statement','form_16','itr')),
  status          VARCHAR(50) NOT NULL DEFAULT 'processing'
                    CHECK (status IN ('processing','completed','needs_review','failed','deleted')),
  storage_key     VARCHAR(500),             -- Object store key (not public URL)
  file_size_bytes INTEGER,
  mime_type       VARCHAR(100),
  overall_confidence NUMERIC(3,2),
  needs_manual_review BOOLEAN DEFAULT FALSE,
  reviewed_by     UUID REFERENCES users(id),
  uploaded_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  processed_at    TIMESTAMPTZ,
  deleted_at      TIMESTAMPTZ
);
```

---

## Table: extracted_fields

OCR-extracted fields per document.

```sql
CREATE TABLE extracted_fields (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  document_id     UUID NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
  field_name      VARCHAR(100) NOT NULL,
  extracted_value TEXT,
  normalized_value TEXT,
  confidence      NUMERIC(3,2) NOT NULL,
  needs_correction BOOLEAN DEFAULT FALSE,
  customer_corrected BOOLEAN DEFAULT FALSE,
  corrected_value TEXT,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: consents

Customer consent records.

```sql
CREATE TABLE consents (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id     UUID NOT NULL REFERENCES customers(id),
  status          VARCHAR(50) NOT NULL DEFAULT 'active'
                    CHECK (status IN ('active','revoked','expired')),
  channel         VARCHAR(50) NOT NULL,
  ip_hash         VARCHAR(64),
  user_agent_hash VARCHAR(64),
  receipt_id      UUID UNIQUE,
  granted_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  revoked_at      TIMESTAMPTZ,
  revocation_reason VARCHAR(255),
  revoked_by      VARCHAR(50)               -- customer | rm | compliance | system
);
```

---

## Table: consent_purposes

Individual purpose entries within a consent.

```sql
CREATE TABLE consent_purposes (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  consent_id      UUID NOT NULL REFERENCES consents(id) ON DELETE CASCADE,
  purpose_code    VARCHAR(100) NOT NULL,
  description     TEXT NOT NULL,
  data_categories TEXT[] NOT NULL,
  retention_days  INTEGER NOT NULL,
  third_party_sharing BOOLEAN DEFAULT FALSE,
  status          VARCHAR(50) NOT NULL DEFAULT 'active'
                    CHECK (status IN ('active','revoked','fulfilled','expired')),
  expires_at      TIMESTAMPTZ NOT NULL,
  revoked_at      TIMESTAMPTZ,
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: applications

Loan/product application records.

```sql
CREATE TABLE applications (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id     UUID NOT NULL REFERENCES customers(id),
  product_id      UUID NOT NULL REFERENCES products(id),
  qualification_id UUID REFERENCES qualification_scores(id),
  consent_id      UUID NOT NULL REFERENCES consents(id),
  reference_number VARCHAR(100) UNIQUE NOT NULL,  -- APRT-2024-001234
  status          VARCHAR(100) NOT NULL DEFAULT 'submitted'
                    CHECK (status IN ('draft','submitted','under_review','approved','rejected','disbursed','cancelled')),
  applied_amount  BIGINT,
  applied_tenure_years SMALLINT,
  application_data JSONB NOT NULL DEFAULT '{}',   -- Pre-filled form data
  decision        VARCHAR(50),
  decision_reason TEXT,
  decision_at     TIMESTAMPTZ,
  decided_by      UUID REFERENCES users(id),
  submitted_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

---

## Table: audit_logs

Immutable append-only audit trail.

```sql
CREATE TABLE audit_logs (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  event_type      VARCHAR(100) NOT NULL,
  entity_type     VARCHAR(100) NOT NULL,
  entity_id       UUID,
  actor_type      VARCHAR(50) NOT NULL CHECK (actor_type IN ('system','user','agent')),
  actor_id        UUID,
  actor_name      VARCHAR(255),
  description     TEXT NOT NULL,
  input_summary   TEXT,
  output_summary  TEXT,
  full_input      JSONB,                          -- Stored encrypted for sensitive events
  full_output     JSONB,
  policy_applied  VARCHAR(255),
  risk_level      VARCHAR(20) NOT NULL DEFAULT 'low'
                    CHECK (risk_level IN ('low','medium','high','critical')),
  flags           JSONB DEFAULT '[]',
  trace_id        UUID,
  timestamp       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- Audit logs are NEVER updated or deleted (append-only enforced via policy)
```

---

## Table: explanations

AI decision explanations stored per entity.

```sql
CREATE TABLE explanations (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  entity_type     VARCHAR(100) NOT NULL,  -- qualification | intent | recommendation | document
  entity_id       UUID NOT NULL,
  summary         TEXT NOT NULL,
  decision        TEXT NOT NULL,
  reasoning       JSONB NOT NULL DEFAULT '[]',
  confidence      NUMERIC(3,2),
  policy_version  VARCHAR(100),
  next_steps      TEXT[],
  agent_version   VARCHAR(50),
  generated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_explanations_entity ON explanations(entity_type, entity_id);
```

---

## Table: policy_rules

Configurable eligibility and scoring rules per product.

```sql
CREATE TABLE policy_rules (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  product_id      UUID NOT NULL REFERENCES products(id),
  rule_code       VARCHAR(100) NOT NULL,
  rule_name       VARCHAR(255) NOT NULL,
  rule_type       VARCHAR(100) NOT NULL,  -- eligibility | scoring | threshold
  dimension       VARCHAR(100),           -- Which scoring dimension this applies to
  conditions      JSONB NOT NULL,         -- Rule conditions as JSON
  weight          NUMERIC(4,3),           -- Scoring weight (0.0 - 1.0)
  version         VARCHAR(50) NOT NULL DEFAULT 'v1',
  is_active       BOOLEAN DEFAULT TRUE,
  created_by      UUID REFERENCES users(id),
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```
