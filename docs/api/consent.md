# API Contract — Consent Management

**Base URL:** `/api/v1/consent`  
**Auth:** Bearer JWT required

---

## POST /api/v1/consent/capture

Capture explicit customer consent.

**Request:**
```json
{
  "customer_id": "uuid",
  "purposes": [
    {
      "purpose_code": "HOME_LOAN_KYC",
      "description": "Verify your identity and income for SBI Home Loan application",
      "data_categories": ["pan", "income", "aadhaar_offline"],
      "retention_days": 90,
      "third_party_sharing": false
    }
  ],
  "channel": "web | mobile | branch | voice",
  "ip_address": "string (hashed)",
  "user_agent": "string",
  "captured_at": "ISO8601"
}
```

**Response 201:**
```json
{
  "consent_id": "uuid",
  "customer_id": "uuid",
  "status": "active",
  "purposes": [
    {
      "purpose_id": "uuid",
      "purpose_code": "HOME_LOAN_KYC",
      "status": "active",
      "granted_at": "ISO8601",
      "expires_at": "ISO8601"
    }
  ],
  "receipt": {
    "receipt_id": "uuid",
    "receipt_url": "/api/v1/consent/{consent_id}/receipt",
    "generated_at": "ISO8601"
  }
}
```

---

## GET /api/v1/consent/:customer_id

List all consents for a customer.

**Response 200:**
```json
{
  "customer_id": "uuid",
  "consents": [
    {
      "consent_id": "uuid",
      "purpose_code": "HOME_LOAN_KYC",
      "description": "Home Loan identity and income verification",
      "status": "active | revoked | expired",
      "data_categories": ["pan", "income"],
      "granted_at": "ISO8601",
      "expires_at": "ISO8601",
      "revoked_at": "ISO8601 | null"
    }
  ]
}
```

---

## DELETE /api/v1/consent/:consent_id

Revoke consent.

**Request:**
```json
{
  "revocation_reason": "no_longer_interested | withdraw_application | other",
  "revoked_by": "customer | rm | compliance"
}
```

**Response 200:**
```json
{
  "consent_id": "uuid",
  "status": "revoked",
  "revoked_at": "ISO8601",
  "data_deletion_scheduled_at": "ISO8601",
  "retained_data": ["application_form"],
  "retained_reason": "RBI regulation requires 10-year retention of submitted applications",
  "receipt_url": "/api/v1/consent/{consent_id}/revocation-receipt"
}
```

---

## GET /api/v1/consent/:consent_id/receipt

Download consent receipt (PDF or JSON).

**Query Params:**
```
format   string   — json (default) | pdf
```

**Response 200 (JSON format):**
```json
{
  "receipt_id": "uuid",
  "consent_id": "uuid",
  "customer_name": "Arjun Sharma",
  "purposes": ["Home Loan identity verification"],
  "data_collected": ["PAN Card", "Salary Slip"],
  "retention_period": "90 days from application date",
  "third_party_sharing": "None",
  "your_rights": [
    "Revoke consent at any time via aperture.sbi.co.in/consents",
    "Request deletion of your data",
    "Access all data we hold about you"
  ],
  "contact": "dpo@sbi.co.in",
  "granted_at": "ISO8601",
  "receipt_generated_at": "ISO8601"
}
```

---

## POST /api/v1/consent/verify

Check if valid consent exists for a specific purpose before data access.

**Request:**
```json
{
  "customer_id": "uuid",
  "purpose_code": "HOME_LOAN_KYC",
  "data_category": "pan"
}
```

**Response 200:**
```json
{
  "authorized": true,
  "consent_id": "uuid",
  "granted_at": "ISO8601",
  "expires_at": "ISO8601"
}
```

**Response 403:**
```json
{
  "authorized": false,
  "reason": "no_consent | consent_revoked | consent_expired | purpose_mismatch"
}
```
