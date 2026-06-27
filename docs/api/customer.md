# API Contract — Customers

**Base URL:** `/api/v1/customers`  
**Auth:** Bearer JWT required for all endpoints  
**Roles:** rm, compliance, operations, admin (customer role: own record only)

---

## GET /api/v1/customers

List/search prospects and customers.

**Query Params:**
```
search       string  — Name, phone, PAN (partial match)
status       string  — prospect | qualified | applied | onboarded
tier         string  — hot | warm | cold
product_id   uuid    — Filter by product interest
page         int     — Default 1
limit        int     — Default 20, max 100
sort_by      string  — intent_score | qual_score | created_at
sort_dir     string  — asc | desc
```

**Response 200:**
```json
{
  "data": [
    {
      "id": "uuid",
      "name": "string",
      "phone": "string (masked: ****1234)",
      "email": "string (masked: a***@gmail.com)",
      "status": "prospect",
      "tier": "hot",
      "intent_score": 82,
      "qualification_score": 89,
      "product_interest": "home_loan",
      "location": "Pune",
      "last_activity": "ISO8601",
      "created_at": "ISO8601"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 248,
    "total_pages": 13
  }
}
```

---

## GET /api/v1/customers/:id

Get full customer/prospect detail.

**Response 200:**
```json
{
  "id": "uuid",
  "name": "string",
  "phone": "string",
  "email": "string",
  "age": 28,
  "occupation": "salaried",
  "income_monthly": 120000,
  "location": "Pune",
  "existing_sbi_customer": false,
  "status": "applied",
  "tier": "hot",
  "intent_score": {
    "score": 82,
    "tier": "hot",
    "signals": [
      {"type": "page_view", "weight": 0.2, "value": "home_loan_page"},
      {"type": "dwell_time", "weight": 0.3, "value": "420s"},
      {"type": "emi_interaction", "weight": 0.3, "value": true},
      {"type": "campaign_click", "weight": 0.2, "value": "sms_oct24"}
    ],
    "computed_at": "ISO8601"
  },
  "qualification_score": {
    "score": 89,
    "tier": "hot",
    "dimensions": {
      "intent_strength": {"score": 30, "max": 30},
      "eligibility": {"score": 22, "max": 25},
      "kyc_readiness": {"score": 15, "max": 20},
      "product_fit": {"score": 13, "max": 15},
      "completion_probability": {"score": 9, "max": 10}
    },
    "recommended_products": [
      {"product_id": "uuid", "name": "SBI Home Loan", "fit_score": 92}
    ],
    "explanation_id": "uuid",
    "computed_at": "ISO8601"
  },
  "documents": [
    {"id": "uuid", "type": "pan", "status": "verified", "uploaded_at": "ISO8601"}
  ],
  "consents": [
    {"id": "uuid", "purpose": "home_loan_application", "status": "active"}
  ],
  "applications": [
    {"id": "uuid", "product": "SBI Home Loan", "status": "under_review"}
  ],
  "next_best_action": {
    "action": "call_customer",
    "reason": "Hot lead — all docs verified, awaiting RM call",
    "priority": "high",
    "deadline": "ISO8601"
  },
  "created_at": "ISO8601",
  "updated_at": "ISO8601"
}
```

---

## POST /api/v1/customers

Create new prospect (from session data or manual RM entry).

**Request:**
```json
{
  "name": "string",
  "phone": "string",
  "email": "string (optional)",
  "source": "web_session | rm_manual | campaign | branch",
  "session_id": "uuid (required if source=web_session)",
  "product_interest": "string"
}
```

**Response 201:**
```json
{
  "id": "uuid",
  "status": "prospect",
  "created_at": "ISO8601"
}
```

---

## PATCH /api/v1/customers/:id

Update customer record (RM or system update).

**Request (partial update):**
```json
{
  "name": "string",
  "income_monthly": 120000,
  "occupation": "salaried",
  "location": "Pune"
}
```

**Response 200:** Full customer object (same as GET /:id)

---

## GET /api/v1/customers/:id/timeline

Get full event timeline for customer.

**Response 200:**
```json
{
  "customer_id": "uuid",
  "events": [
    {
      "id": "uuid",
      "type": "intent_detected",
      "timestamp": "ISO8601",
      "description": "High-intent session detected on Home Loan page",
      "metadata": {"intent_score": 82},
      "actor": "system"
    },
    {
      "id": "uuid",
      "type": "consent_granted",
      "timestamp": "ISO8601",
      "description": "Customer consented to Home Loan application data use",
      "metadata": {"purpose": "home_loan_application"},
      "actor": "customer"
    }
  ]
}
```

---

## GET /api/v1/customers/:id/next-best-action

Get NBA recommendation.

**Response 200:**
```json
{
  "customer_id": "uuid",
  "action": "call_customer",
  "reason": "Hot lead — income verified, EMI calc completed, awaiting RM contact",
  "priority": "high",
  "suggested_script": "Hi, this is [RM name] from SBI. I noticed you were exploring home loans. I wanted to share that you may be eligible for ₹60L at 8.4% p.a. Can I share the details?",
  "product": {
    "id": "uuid",
    "name": "SBI Home Loan",
    "offered_amount": 6000000,
    "offered_rate": 8.4
  },
  "deadline": "ISO8601",
  "generated_at": "ISO8601"
}
```
