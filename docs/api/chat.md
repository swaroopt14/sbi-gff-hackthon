# API Contract — Conversation (Chat)

**Base URL:** `/api/v1/chat`  
**Auth:** Bearer JWT required

---

## POST /api/v1/chat/sessions

Start a new conversation session.

**Request:**
```json
{
  "customer_id": "uuid (optional — null for anonymous)",
  "session_id": "uuid",
  "product_intent": "home_loan | savings | fd | mf | credit_card",
  "language": "en | hi | ta | te | kn | ml",
  "channel": "web | mobile | branch | rm_assist"
}
```

**Response 201:**
```json
{
  "conversation_id": "uuid",
  "customer_id": "uuid | null",
  "language": "en",
  "opening_message": {
    "id": "uuid",
    "role": "assistant",
    "content": "Hi! I noticed you were exploring home loans. I'm here to help you check your eligibility in just a few minutes. Shall we get started?",
    "timestamp": "ISO8601",
    "metadata": {
      "personalized": true,
      "offer_preview": {
        "product": "SBI Home Loan",
        "max_amount": 6000000,
        "rate": 8.4
      }
    }
  }
}
```

---

## POST /api/v1/chat/sessions/:conversation_id/messages

Send a message in an existing conversation.

**Request:**
```json
{
  "content": "My salary is 1.2 lakhs per month",
  "message_type": "text | file_reference | button_click",
  "metadata": {}
}
```

**Response 200:**
```json
{
  "message_id": "uuid",
  "role": "assistant",
  "content": "That's great! ₹1.2 lakhs per month makes you eligible for a significant home loan. What type of employment are you in — salaried or self-employed?",
  "timestamp": "ISO8601",
  "metadata": {
    "extracted_fields": {
      "income_monthly": 120000
    },
    "form_progress": 0.4,
    "fields_collected": ["income"],
    "fields_remaining": ["employment_type", "property_location", "loan_amount", "purpose"]
  },
  "quick_replies": [
    {"label": "Salaried", "value": "salaried"},
    {"label": "Self-employed", "value": "self_employed"},
    {"label": "Business Owner", "value": "business"}
  ],
  "action": null
}
```

**Special response when all fields collected:**
```json
{
  "message_id": "uuid",
  "role": "assistant",
  "content": "Perfect! Based on what you've shared, here's your personalised offer.",
  "timestamp": "ISO8601",
  "action": {
    "type": "show_offer",
    "payload": {
      "qualification_id": "uuid",
      "product": "SBI Home Loan",
      "offered_amount": 6000000,
      "rate": 8.4,
      "emi": 45326,
      "tenure_years": 30,
      "cta": "Apply Now"
    }
  }
}
```

---

## GET /api/v1/chat/sessions/:conversation_id

Get full conversation history.

**Response 200:**
```json
{
  "conversation_id": "uuid",
  "customer_id": "uuid",
  "status": "active | completed | abandoned | transferred",
  "language": "en",
  "product_intent": "home_loan",
  "messages": [
    {
      "id": "uuid",
      "role": "assistant | user",
      "content": "string",
      "timestamp": "ISO8601",
      "metadata": {}
    }
  ],
  "collected_data": {
    "name": "Arjun Sharma",
    "income_monthly": 120000,
    "employment_type": "salaried",
    "location": "Pune",
    "loan_amount": 6000000,
    "purpose": "home_purchase"
  },
  "form_completion": 1.0,
  "started_at": "ISO8601",
  "last_message_at": "ISO8601"
}
```

---

## POST /api/v1/chat/sessions/:conversation_id/transfer

Transfer conversation to RM.

**Request:**
```json
{
  "reason": "customer_request | low_confidence | sensitive_query | escalation",
  "rm_id": "uuid (optional — system assigns if null)"
}
```

**Response 200:**
```json
{
  "transfer_id": "uuid",
  "assigned_rm": {
    "id": "uuid",
    "name": "Rajesh Kumar",
    "phone": "****5678"
  },
  "estimated_wait_minutes": 5,
  "summary_for_rm": "Customer Arjun Sharma (income ₹1.2L, salaried) is interested in a ₹60L home loan in Pune. PAN uploaded. Intent score 82/100."
}
```

---

## DELETE /api/v1/chat/sessions/:conversation_id

End/abandon a conversation session.

**Response 200:**
```json
{
  "conversation_id": "uuid",
  "status": "abandoned",
  "ended_at": "ISO8601"
}
```
