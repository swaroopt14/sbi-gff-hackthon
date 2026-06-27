# API Contract — Audit & Explainability

**Base URL:** `/api/v1`  
**Auth:** Bearer JWT required  
**Roles:** compliance, admin (full); rm (limited); customer (own record)

---

## GET /api/v1/audit/logs

Query audit logs.

**Query Params:**
```
customer_id   uuid     — Filter by customer
actor_id      uuid     — Filter by actor (user who performed action)
event_type    string   — intent_detected | consent_granted | qualification_scored | document_uploaded | application_submitted | override | revocation | compliance_check
start_date    ISO8601
end_date      ISO8601
risk_level    string   — low | medium | high
page          int      — Default 1
limit         int      — Default 50, max 200
```

**Response 200:**
```json
{
  "data": [
    {
      "log_id": "uuid",
      "event_type": "qualification_scored",
      "entity_type": "customer",
      "entity_id": "uuid",
      "actor_type": "system | user",
      "actor_id": "uuid | null",
      "actor_name": "AI Qualification Agent | Rajesh Kumar (RM)",
      "description": "Qualification score computed: 89/100 (Hot tier)",
      "input_summary": "Income ₹1.2L, Age 28, Intent 82/100",
      "output_summary": "Score 89, Tier: Hot, Product: Home Loan",
      "risk_level": "low",
      "flags": [],
      "timestamp": "ISO8601",
      "trace_id": "uuid"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 50,
    "total": 384
  }
}
```

---

## GET /api/v1/audit/logs/:log_id

Get detailed audit log entry.

**Response 200:**
```json
{
  "log_id": "uuid",
  "event_type": "string",
  "entity_type": "string",
  "entity_id": "uuid",
  "actor_type": "string",
  "actor_id": "uuid",
  "actor_name": "string",
  "description": "string",
  "full_input": {},
  "full_output": {},
  "policy_applied": "string",
  "risk_level": "string",
  "flags": [
    {
      "flag_type": "data_sensitivity | human_review_required | policy_exception",
      "description": "string",
      "resolved": false
    }
  ],
  "related_events": ["uuid"],
  "timestamp": "ISO8601",
  "trace_id": "uuid"
}
```

---

## GET /api/v1/audit/customers/:customer_id/timeline

Get audit timeline for a customer journey (visual timeline data).

**Response 200:**
```json
{
  "customer_id": "uuid",
  "customer_name": "Arjun Sharma",
  "timeline": [
    {
      "log_id": "uuid",
      "event_type": "session_started",
      "icon": "session",
      "title": "High-intent session detected",
      "description": "Browsed Home Loan page for 7 minutes. EMI calculator used.",
      "timestamp": "ISO8601",
      "risk_level": "low",
      "expandable": true
    },
    {
      "log_id": "uuid",
      "event_type": "consent_granted",
      "icon": "consent",
      "title": "Consent granted",
      "description": "Customer consented to Home Loan KYC data use",
      "timestamp": "ISO8601",
      "risk_level": "low",
      "expandable": false
    }
  ]
}
```

---

## POST /api/v1/audit/export

Export audit logs as CSV or PDF.

**Auth:** Roles: compliance, admin

**Request:**
```json
{
  "filters": {
    "customer_id": "uuid (optional)",
    "start_date": "ISO8601",
    "end_date": "ISO8601",
    "event_types": ["consent_granted", "qualification_scored"]
  },
  "format": "csv | pdf",
  "include_full_input_output": false
}
```

**Response 200:**
```json
{
  "export_id": "uuid",
  "download_url": "/api/v1/audit/exports/{export_id}/download",
  "expires_at": "ISO8601",
  "record_count": 384
}
```

---

## GET /api/v1/explain/:entity_type/:entity_id

Get AI explanation for a decision.

**Path Params:**
- `entity_type`: `qualification | intent | recommendation | document_extraction`
- `entity_id`: UUID of the entity

**Query Params:**
```
audience   string   — customer | rm | compliance (controls detail level)
```

**Response 200:**
```json
{
  "explanation_id": "uuid",
  "entity_type": "qualification",
  "entity_id": "uuid",
  "audience": "rm",
  "summary": "Customer qualifies as Hot lead for SBI Home Loan",
  "decision": "Qualified — Score 89/100 (Hot tier)",
  "reasoning": [
    {
      "factor": "Income eligibility",
      "outcome": "PASS",
      "confidence": 0.95,
      "detail": "Declared income ₹1,20,000/month exceeds minimum ₹50,000 threshold for ₹60L loan (ratio 2.0x — recommended 2.5x minimum met)",
      "policy": "HOME_LOAN_MIN_INCOME_RATIO_V2",
      "evidence": "Self-declared via conversation; pending document verification"
    },
    {
      "factor": "Age eligibility",
      "outcome": "PASS",
      "confidence": 1.0,
      "detail": "Age 28 is within 21-65 years eligible range. Loan tenure 30 years ends at age 58.",
      "policy": "HOME_LOAN_AGE_ELIGIBILITY_V1",
      "evidence": "Declared in conversation"
    },
    {
      "factor": "KYC readiness",
      "outcome": "PARTIAL",
      "confidence": 0.75,
      "detail": "PAN and Salary Slip available. Aadhaar not yet uploaded. CKYCR check pending.",
      "policy": "KYC_MINIMUM_DOCUMENTS",
      "evidence": "Document uploads in system"
    }
  ],
  "next_steps": [
    "Upload Aadhaar offline XML for full KYC",
    "RM to confirm income via salary slip verification"
  ],
  "confidence": 0.91,
  "policy_version": "HOME_LOAN_ELIGIBILITY_V2",
  "generated_at": "ISO8601",
  "can_be_overridden": true
}
```
