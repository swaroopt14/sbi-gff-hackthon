# Aperture — User Flows

**Version:** 1.0  
**Date:** 2026-06-27

---

## Flow 1 — New Prospect Acquisition (Primary Journey)

```
[Customer on SBI Website]
          |
          | Browses Home Loan page
          | Time on page > 3 minutes
          | Uses EMI calculator
          |
          v
[Intent Detection Agent triggers]
  - Intent Score: 82/100
  - Category: "Active Comparison"
  - Signals: Page depth, EMI interaction, return visit
          |
          v
[Qualification Agent runs silently]
  - Uses declared session data only (no PII yet)
  - Eligibility check against age/location
  - Preliminary product fit: Home Loan ✓
          |
          v
[Personalised Chat Widget appears]
  - "Hi! We noticed you're exploring home loans."
  - "Based on your profile, you may be eligible for up to ₹60L at 8.4%."
  - CTA: "Check My Eligibility in 2 Minutes"
          |
          |---[Customer ignores]---> Session ends, prospect record saved for re-engagement
          |
          | Customer clicks CTA
          |
          v
[Consent Screen — Step 1]
  "To check your eligibility, we need basic details."
  "This will NOT affect your credit score."
  [Allow] [Cancel]
          |
          v
[Conversation Agent — Data Collection]
  Q: What's your name?
  Q: What's your approximate monthly income?
  Q: What's your employment type? (Salaried/Self-employed/Business)
  Q: Which city are you looking to buy property in?
  Q: What's the approximate property value?
          |
          v
[Qualification Agent — Full Scoring]
  Intent Strength: 30/30 (high session depth, EMI usage, campaign click)
  Eligibility:     22/25 (income ₹1.2L, age 28, metro city)
  KYC Readiness:   15/20 (no docs yet, but self-declared PAN available)
  Product Fit:     13/15 (Home Loan match 92%)
  Completion Prob:  9/10 (high engagement, all questions answered)
  TOTAL SCORE:     89/100 → HOT 🔥
          |
          v
[Personalised Offer Shown]
  "Congratulations! You're pre-qualified for:
   SBI Home Loan — ₹60L at 8.40% p.a.
   EMI: ~₹45,300/month (30 years)
   Processing Fee: ₹15,000
   [Show Full Offer] [Apply Now]"
          |
          v
[Consent Screen — Step 2 (Full KYC Consent)]
  "To complete your application, we need to verify:"
  ☑ PAN Card (for identity verification)
  ☑ Income documents (salary slip/bank statement)
  "Your data will be used only for this home loan application."
  "Retention: 90 days if application is not completed."
  [I Agree & Continue] [Not Now]
          |
          v
[Document Upload]
  Upload PAN card image
  Upload last 3 months salary slips
  (Optional) Upload last 6 months bank statements
          |
          v
[Document Agent — Processing]
  PAN extracted: Name ✓ PAN number ✓ DOB ✓ (Confidence: 0.96)
  Salary slip: Employer ✓ Net salary ✓ Date ✓ (Confidence: 0.91)
  Verification: PAN format valid ✓ Income consistent ✓
          |
          v
[Application Pre-filled]
  Name: Arjun Sharma
  PAN: ABCDE1234F
  Income: ₹1,20,000/month
  Employer: Infosys Ltd
  Loan Amount: ₹60,00,000
  Property City: Pune
  [Review & Submit]
          |
          v
[Application Submitted]
  Application ID: APRT-2024-001234
  Status: Under Review
  Estimated Decision: 48 hours
  [Track Application] [Download Receipt]
          |
          v
[Customer Dashboard Active]
  - Application status tracker
  - Document upload status
  - Next steps
  - RM contact details
```

---

## Flow 2 — RM Workflow (Daily Acquisition Management)

```
[RM Opens Dashboard]
          |
          v
[Today's Prioritised Pipeline]
  HOT (5 leads) | WARM (18 leads) | COLD (47 leads)
          |
          v
[Opens Lead Card: Arjun Sharma]
  Intent: 89/100
  Product: Home Loan ₹60L
  Last activity: 2 hours ago
  Conversation Summary: "Has salary of ₹1.2L, looking at Pune property ₹75L"
  Documents: PAN ✓ Salary Slip ✓
  Status: Awaiting RM call confirmation
  
  Next Best Action: "Call now — hot window within 6 hours"
          |
          v
[RM Clicks "Initiate Contact"]
  - System logs: RM initiated contact at [timestamp]
  - RM gets customer phone number (if consent given)
          |
          v
[Post-call: RM logs outcome]
  Outcome: [Interested] [Not Interested] [Call Later] [Escalate]
  Notes: [Free text]
          |
          v
[System updates prospect record]
  - Qualification score adjusted based on RM input
  - Next Best Action recalculated
```

---

## Flow 3 — Compliance Officer Review

```
[Compliance Officer Opens Dashboard]
          |
          v
[Compliance Overview]
  Consent Health: 98.2% (flagged: 12 cases)
  Audit Flags: 3 high-severity
  Policy Violations: 0
  Pending Human Reviews: 8
          |
          v
[Reviews Flagged Case]
  Prospect: Meena Pillai
  Flag: "Qualification score based on inferred income"
  Risk: Medium
  Recommendation: "Request income documents before proceeding"
          |
          v
[Compliance Officer Actions]
  [Approve] [Override with Reason] [Reject & Block] [Escalate]
          |
          v
[All actions logged in audit trail]
  Actor: Vikram Nair (Compliance Officer)
  Action: Override — Request additional documents
  Reason: "Inferred income used without declared income confirmation"
  Timestamp: 2024-01-15 14:32:07 IST
```

---

## Flow 4 — Consent Revocation Flow

```
[Customer logged into Customer Portal]
          |
          v
[My Consents Page]
  Shows all active consents with:
  - Purpose
  - Data collected
  - Retention period
  - Date granted
          |
          v
[Customer clicks "Revoke" on Home Loan consent]
          |
          v
[Revocation Confirmation]
  "Revoking consent will:
   ✗ Stop your home loan application
   ✗ Delete your documents from our system
   ✓ Your data will be deleted within 30 days
   Note: Submitted application forms may be retained as per RBI requirement"
  [Confirm Revocation] [Keep Consent]
          |
          v
[System Actions on Revocation]
  - Consent record marked REVOKED with timestamp
  - Document store flagged for deletion
  - Active conversation terminated
  - RM notified: "Customer revoked consent"
  - Revocation receipt generated and emailed
```

---

## Flow 5 — Document Agent Error Handling

```
[Customer Uploads Document]
          |
          v
[Document Agent Processing]
          |
          |---[Image too blurry]---> "Please upload a clearer image"
          |---[PAN format invalid]-> "PAN number not matching expected format"
          |---[Income mismatch]----> Flag for manual review, continue
          |---[Missing field]------> "Could not read [field], please enter manually"
          |
          v
[Confidence Score per field]
  If ANY field confidence < 0.7:
    → Flag for manual review
    → Allow customer to correct
  If ALL fields confidence > 0.85:
    → Auto-proceed
  If critical field (PAN number) confidence < 0.9:
    → Mandatory manual entry
```

---

## Flow 6 — Multilingual Conversation Flow

```
[Conversation Agent starts in English]
          |
          v
[Language Detection]
  Customer types in Hindi → Switch to Hindi
  Customer types in Tamil → Switch to Tamil
  (or) Customer explicitly selects language from menu
          |
          v
[All subsequent messages in selected language]
  All form labels translated
  All error messages translated
  All offer details translated
          |
          v
[If AI confidence in response drops < 0.7 in regional language]
  → Offer: "Would you like to continue in English?"
  → Or: "Let me connect you with a regional language specialist"
```
