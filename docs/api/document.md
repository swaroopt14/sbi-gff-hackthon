# API Contract — Documents

**Base URL:** `/api/v1/documents`  
**Auth:** Bearer JWT required

---

## POST /api/v1/documents/upload

Upload a document for OCR extraction.

**Content-Type:** `multipart/form-data`

**Form Fields:**
```
customer_id   string (required)
document_type string (required) — pan | aadhaar_offline_xml | salary_slip | bank_statement | form_16
consent_id    string (required) — consent must exist before upload
file          file   (required) — max 10MB, types: jpg, png, pdf, xml
```

**Response 201:**
```json
{
  "document_id": "uuid",
  "customer_id": "uuid",
  "document_type": "pan",
  "status": "processing",
  "uploaded_at": "ISO8601",
  "processing_eta_seconds": 8
}
```

---

## GET /api/v1/documents/:document_id

Get document processing status and extracted fields.

**Response 200 (processing):**
```json
{
  "document_id": "uuid",
  "status": "processing",
  "progress_percent": 65
}
```

**Response 200 (completed):**
```json
{
  "document_id": "uuid",
  "customer_id": "uuid",
  "document_type": "pan",
  "status": "completed | needs_review | failed",
  "uploaded_at": "ISO8601",
  "processed_at": "ISO8601",
  "extracted_fields": [
    {
      "field_name": "name",
      "extracted_value": "ARJUN SHARMA",
      "normalized_value": "Arjun Sharma",
      "confidence": 0.97,
      "needs_correction": false
    },
    {
      "field_name": "pan_number",
      "extracted_value": "ABCDE1234F",
      "normalized_value": "ABCDE1234F",
      "confidence": 0.99,
      "needs_correction": false
    },
    {
      "field_name": "date_of_birth",
      "extracted_value": "12/05/1996",
      "normalized_value": "1996-05-12",
      "confidence": 0.94,
      "needs_correction": false
    },
    {
      "field_name": "father_name",
      "extracted_value": "RAMESH SHARMA",
      "normalized_value": "Ramesh Sharma",
      "confidence": 0.88,
      "needs_correction": false
    }
  ],
  "overall_confidence": 0.95,
  "verification_status": {
    "format_valid": true,
    "checksum_valid": true,
    "cross_match": null
  },
  "flags": [],
  "reviewer_id": null
}
```

**Response 200 (needs_review):**
```json
{
  "document_id": "uuid",
  "status": "needs_review",
  "reason": "Low confidence on PAN number field (0.62)",
  "extracted_fields": [...],
  "manual_review_required_fields": ["pan_number"]
}
```

---

## PATCH /api/v1/documents/:document_id/correct

Customer manually corrects an extracted field.

**Request:**
```json
{
  "corrections": [
    {
      "field_name": "pan_number",
      "corrected_value": "ABCDE1234F",
      "customer_confirmed": true
    }
  ]
}
```

**Response 200:**
```json
{
  "document_id": "uuid",
  "corrected_fields": ["pan_number"],
  "status": "correction_applied",
  "updated_at": "ISO8601"
}
```

---

## GET /api/v1/documents/customers/:customer_id

List all documents for a customer.

**Response 200:**
```json
{
  "customer_id": "uuid",
  "documents": [
    {
      "document_id": "uuid",
      "document_type": "pan",
      "status": "completed",
      "overall_confidence": 0.95,
      "uploaded_at": "ISO8601"
    },
    {
      "document_id": "uuid",
      "document_type": "salary_slip",
      "status": "completed",
      "overall_confidence": 0.91,
      "uploaded_at": "ISO8601"
    }
  ]
}
```

---

## DELETE /api/v1/documents/:document_id

Delete a document (on consent revocation or customer request).

**Auth:** Roles: customer (own), admin

**Response 200:**
```json
{
  "document_id": "uuid",
  "deleted_at": "ISO8601",
  "audit_log_id": "uuid"
}
```
