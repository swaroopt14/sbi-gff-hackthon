# Aperture — Entity Relationship Diagram

**Version:** 1.0  
**Date:** 2026-06-27

---

## Entity Relationship Overview

```
users
  │ 1
  │ ├─── has many ──▶ customers (1:1 for customer role users)
  │ ├─── has many ──▶ audit_logs (as actor)
  │ └─── has many ──▶ rm_overrides (as rm)

customers
  │ 1
  ├─── has many ──▶ intent_sessions (1:N)
  │                    │
  │                    └─── has many ──▶ intent_events (1:N)
  │                    └─── has one  ──▶ intent_scores (1:1 per session)
  │
  ├─── has many ──▶ qualification_scores (1:N, one per product attempt)
  │                    └─── has one ──▶ rm_overrides (0:1)
  │
  ├─── has many ──▶ product_recommendations (1:N)
  │                    └─── belongs to ──▶ products
  │
  ├─── has many ──▶ conversations (1:N)
  │                    └─── has many ──▶ conversation_turns (1:N)
  │
  ├─── has many ──▶ documents (1:N)
  │                    └─── has many ──▶ extracted_fields (1:N)
  │
  ├─── has many ──▶ consents (1:N)
  │                    └─── has many ──▶ consent_purposes (1:N)
  │
  ├─── has many ──▶ applications (1:N)
  │                    ├─── belongs to ──▶ products
  │                    └─── has many ──▶ application_events (1:N)
  │
  ├─── has many ──▶ audit_logs (as entity)
  │
  └─── has many ──▶ explanations (polymorphic)

products
  ├─── has many ──▶ policy_rules (1:N)
  ├─── has many ──▶ product_recommendations (1:N)
  └─── has many ──▶ applications (1:N)
```

---

## Cardinality Matrix

| Parent | Child | Relationship |
|--------|-------|-------------|
| users | customers | 1:1 (customer-role users) |
| customers | intent_sessions | 1:N |
| intent_sessions | intent_events | 1:N |
| intent_sessions | intent_scores | 1:1 |
| customers | qualification_scores | 1:N |
| qualification_scores | rm_overrides | 1:0..1 |
| customers | product_recommendations | 1:N |
| customers | conversations | 1:N |
| conversations | conversation_turns | 1:N |
| customers | documents | 1:N |
| documents | extracted_fields | 1:N |
| customers | consents | 1:N |
| consents | consent_purposes | 1:N |
| customers | applications | 1:N |
| applications | application_events | 1:N |
| products | policy_rules | 1:N |
| products | applications | 1:N |
| products | product_recommendations | 1:N |

---

## Indexes (Performance-Critical)

```sql
-- Customer search
CREATE INDEX idx_customers_status ON customers(status);
CREATE INDEX idx_customers_tier ON customers(tier);
CREATE INDEX idx_customers_phone ON customers(phone);  -- hashed

-- Intent processing
CREATE INDEX idx_intent_sessions_customer ON intent_sessions(customer_id);
CREATE INDEX idx_intent_sessions_created ON intent_sessions(created_at DESC);
CREATE INDEX idx_intent_scores_session ON intent_scores(session_id);

-- Qualification pipeline
CREATE INDEX idx_qual_scores_customer ON qualification_scores(customer_id);
CREATE INDEX idx_qual_scores_product ON qualification_scores(product_id);
CREATE INDEX idx_qual_scores_tier ON qualification_scores(tier);

-- Consent enforcement (hot path)
CREATE INDEX idx_consents_customer ON consents(customer_id);
CREATE INDEX idx_consents_status ON consents(status);
CREATE INDEX idx_consent_purposes_consent ON consent_purposes(consent_id);
CREATE INDEX idx_consent_purposes_code ON consent_purposes(purpose_code);

-- Audit queries
CREATE INDEX idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_logs_timestamp ON audit_logs(timestamp DESC);
CREATE INDEX idx_audit_logs_event_type ON audit_logs(event_type);
CREATE INDEX idx_audit_logs_risk ON audit_logs(risk_level) WHERE risk_level != 'low';

-- Document processing
CREATE INDEX idx_documents_customer ON documents(customer_id);
CREATE INDEX idx_documents_status ON documents(status);

-- Conversation
CREATE INDEX idx_conversations_customer ON conversations(customer_id);
CREATE INDEX idx_conv_turns_conversation ON conversation_turns(conversation_id);
```
