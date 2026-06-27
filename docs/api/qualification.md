# API Contract — Qualification & Intent

**Base URL:** `/api/v1`  
**Auth:** Bearer JWT required

---

## POST /api/v1/intent/score

Compute intent score for a session (called by frontend or Kafka consumer).

**Request:**
```json
{
  "session_id": "uuid",
  "customer_id": "uuid (optional — null for anonymous)",
  "events": [
    {
      "type": "page_view | click | dwell | emi_calc | search | campaign_click | return_visit | abandon",
      "page": "string",
      "timestamp": "ISO8601",
      "metadata": {}
    }
  ],
  "campaign_source": "string (optional)"
}
```

**Response 200:**
```json
{
  "session_id": "uuid",
  "intent_score": 82,
  "intent_tier": "hot | warm | cold",
  "intent_category": "product_discovery | active_comparison | ready_to_apply | dropoff_recovery",
  "signals_used": [
    {
      "type": "dwell_time",
      "value": 420,
      "weight": 0.3,
      "contribution": 25
    }
  ],
  "should_trigger_chat": true,
  "recommended_product": "home_loan",
  "explanation": "High dwell time (7 min), EMI calculator used, return visit detected",
  "computed_at": "ISO8601"
}
```

---

## GET /api/v1/intent/sessions/:session_id

Get intent data for a specific session.

**Response 200:**
```json
{
  "session_id": "uuid",
  "customer_id": "uuid | null",
  "started_at": "ISO8601",
  "last_event_at": "ISO8601",
  "event_count": 14,
  "intent_score": 82,
  "intent_tier": "hot",
  "events": []
}
```

---

## POST /api/v1/qualify/score

Run full qualification scoring for a customer.

**Request:**
```json
{
  "customer_id": "uuid",
  "product_id": "uuid",
  "profile": {
    "age": 28,
    "income_monthly": 120000,
    "income_type": "salaried | self_employed | business",
    "location": "Pune",
    "existing_sbi_customer": false,
    "existing_loans": 0,
    "credit_score_available": false
  },
  "session_data": {
    "intent_score": 82,
    "intent_category": "ready_to_apply"
  },
  "documents_available": ["pan", "salary_slip"]
}
```

**Response 200:**
```json
{
  "qualification_id": "uuid",
  "customer_id": "uuid",
  "product_id": "uuid",
  "overall_score": 89,
  "tier": "hot",
  "dimensions": {
    "intent_strength": {
      "score": 30,
      "max": 30,
      "weight": 0.30,
      "evidence": "Session score 82/100, EMI calculator used"
    },
    "eligibility": {
      "score": 22,
      "max": 25,
      "weight": 0.25,
      "evidence": "Age 28 (eligible), Income ₹1.2L (>threshold), Metro city (tier-1)"
    },
    "kyc_readiness": {
      "score": 15,
      "max": 20,
      "weight": 0.20,
      "evidence": "PAN available, Salary slip available, Aadhaar not yet"
    },
    "product_fit": {
      "score": 13,
      "max": 15,
      "weight": 0.15,
      "evidence": "Income supports ₹60L loan. Property in metro. First-time buyer profile."
    },
    "completion_probability": {
      "score": 9,
      "max": 10,
      "weight": 0.10,
      "evidence": "All questions answered, high engagement, prior completion rate >80%"
    }
  },
  "recommended_offer": {
    "product": "SBI Home Loan",
    "amount": 6000000,
    "rate": 8.4,
    "tenure_years": 30,
    "emi_monthly": 45326
  },
  "explanation_id": "uuid",
  "policy_applied": "HOME_LOAN_ELIGIBILITY_V2",
  "human_review_required": false,
  "computed_at": "ISO8601"
}
```

---

## POST /api/v1/qualify/:qualification_id/override

RM or Compliance Officer override of qualification score.

**Auth:** Roles: rm, compliance, admin

**Request:**
```json
{
  "overridden_score": 75,
  "overridden_tier": "warm",
  "reason": "Customer has undisclosed existing loan based on phone conversation",
  "action": "downgrade | upgrade | reject | approve"
}
```

**Response 200:**
```json
{
  "override_id": "uuid",
  "qualification_id": "uuid",
  "original_score": 89,
  "overridden_score": 75,
  "overridden_by": "uuid (user)",
  "reason": "string",
  "logged_at": "ISO8601"
}
```

---

## GET /api/v1/qualify/:qualification_id/explanation

Get full human-readable explanation for a qualification decision.

**Response 200:**
```json
{
  "qualification_id": "uuid",
  "explanation": {
    "summary": "Customer qualifies for SBI Home Loan because income exceeds minimum threshold, age is within eligible range, PAN is available for KYC, and session shows high purchase intent.",
    "factors": [
      {
        "factor": "Income",
        "status": "pass",
        "detail": "Monthly income ₹1,20,000 exceeds minimum ₹50,000 for ₹60L loan",
        "policy_ref": "HOME_LOAN_MIN_INCOME_RATIO"
      },
      {
        "factor": "Age",
        "status": "pass",
        "detail": "Age 28 is within eligible range 21-65 years",
        "policy_ref": "HOME_LOAN_AGE_ELIGIBILITY"
      },
      {
        "factor": "Intent",
        "status": "pass",
        "detail": "Session score 82/100 — highest confidence tier",
        "policy_ref": "INTENT_THRESHOLD_HOT"
      }
    ],
    "confidence": 0.91,
    "generated_at": "ISO8601"
  }
}
```
