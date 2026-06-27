# Aperture — Sequence Diagrams

**Version:** 1.0  
**Date:** 2026-06-27

---

## Sequence 1 — Customer Acquisition (Full Primary Flow)

```
Customer    Frontend    Backend     IntentAgent  QualAgent  PersonAgent  ConvAgent  DocAgent  ConsentSvc  AuditAgent

   │           │           │            │            │           │           │          │          │           │
   │─ Opens ──▶│           │            │            │           │           │          │          │           │
   │  Website  │─ GET / ──▶│            │            │           │           │          │          │           │
   │           │◀─ 200 ───│            │            │           │           │          │          │           │
   │           │           │            │            │           │           │          │          │           │
   │─ Browses ─▶│           │            │            │           │           │          │          │           │
   │  HomeLoan │─ POST ───▶│            │            │           │           │          │          │           │
   │  page     │  /sessions │            │            │           │           │          │          │           │
   │           │           │─ POST ────▶│            │           │           │          │          │           │
   │           │           │  /score    │            │           │           │          │          │           │
   │           │           │            │─ Process ─▶│            │           │          │          │           │
   │           │           │            │  signals   │            │           │          │          │           │
   │           │           │            │◀─ score:82 │            │           │          │          │           │
   │           │           │◀─ Intent ─│            │           │           │          │          │           │
   │           │           │  Score:82  │            │           │           │          │          │           │
   │           │           │─ POST ────────────────▶│           │           │          │          │           │
   │           │           │  /qualify  │            │           │           │          │          │           │
   │           │           │            │            │◀─ Score ─▶│           │          │          │           │
   │           │           │            │            │   89/100  │           │          │          │           │
   │           │           │─ POST ────────────────────────────▶│           │          │          │           │
   │           │           │  /personalise          │            │◀─ Offer ─▶│          │          │           │
   │           │           │◀─ offer message        │            │           │          │          │           │
   │           │─ Show ───▶│           │            │            │           │          │          │           │
   │  Chat Widget          │           │            │            │           │          │          │           │
   │◀──────────│           │           │            │            │           │          │          │           │
   │           │           │           │            │            │           │          │          │           │
   │─ Clicks ─▶│           │           │            │            │           │          │          │           │
   │  Widget   │─ Show ───▶│           │            │            │           │          │          │           │
   │           │  Consent  │           │            │            │           │          │          │           │
   │◀─ Consent─│           │           │            │            │           │          │          │           │
   │  Screen   │           │           │            │            │           │          │          │           │
   │─ Agrees ─▶│           │           │            │            │           │          │          │           │
   │           │─ POST ───▶│           │            │            │           │          │          │           │
   │           │  /consent/capture     │            │            │           │          │          │           │
   │           │           │─────────────────────────────────────────────────────────────────────▶│           │
   │           │           │  Consent persisted + receipt generated          │          │◀─ OK ───│           │
   │           │           │─ POST ─────────────────────────────────────────▶│          │          │           │
   │           │           │  /audit/log            │            │           │          │          │           │
   │           │           │  event: ConsentGranted │            │           │          │◀─ OK ───│           │
   │           │           │           │            │            │           │          │          │           │
   │─ Chat ───▶│           │           │            │            │           │          │          │           │
   │  starts   │─ POST ───▶│           │            │            │           │          │          │           │
   │           │  /chat    │─ POST ─────────────────────────────▶│           │          │          │           │
   │           │           │  /agents/chat          │            │           │          │          │           │
   │           │           │           │            │            │◀─ response│          │          │           │
   │           │◀─ reply ─│           │            │            │           │          │          │           │
   │◀──────────│           │           │            │            │           │          │          │           │
   │           │           │           │            │            │           │          │          │           │
   │─ Uploads ─▶│           │           │            │            │           │          │          │           │
   │  docs     │─ POST ───▶│           │            │            │           │          │          │           │
   │           │  /documents│─ POST ──────────────────────────────────────────────────▶│          │           │
   │           │           │  /agents/document/extract          │           │          │          │           │
   │           │           │           │            │            │           │◀─ fields │          │           │
   │           │◀─ prefill─│           │            │            │           │          │          │           │
   │◀──────────│           │           │            │            │           │          │          │           │
   │           │           │           │            │            │           │          │          │           │
   │─ Submits ─▶│           │           │            │            │           │          │          │           │
   │           │─ POST ───▶│           │            │            │           │          │          │           │
   │           │  /applications        │            │            │           │          │          │           │
   │           │           │─ POST ─────────────────────────────────────────────────────────────────────────▶│
   │           │           │  /audit/log → ApplicationSubmitted │           │          │          │           │
   │           │◀─ app_id ─│           │            │            │           │          │          │           │
   │◀──────────│           │           │            │            │           │          │          │           │
```

---

## Sequence 2 — Intent Detection Flow

```
Browser         Session Service     Kafka          Intent Agent      DB

   │─ Page View ─▶│                    │                │              │
   │              │─ Produce ─────────▶│                │              │
   │              │  intent.event      │                │              │
   │              │  {type:PAGE_VIEW}  │                │              │
   │              │                    │─ Consume ─────▶│              │
   │─ EMI Calc ──▶│                    │                │              │
   │              │─ Produce ─────────▶│                │              │
   │              │  {type:INTERACT}   │                │              │
   │              │                    │─ Consume ─────▶│              │
   │─ +3min dwell▶│                    │                │              │
   │              │─ Produce ─────────▶│                │              │
   │              │  {type:DWELL_HIGH} │                │              │
   │              │                    │─ Consume ─────▶│              │
   │              │                    │                │─ Aggregate ─▶│
   │              │                    │                │  signals     │
   │              │                    │                │─ Compute ───▶│
   │              │                    │                │  score       │
   │              │                    │                │─ Persist ───▶│
   │              │                    │                │  IntentScore │
   │              │                    │                │◀─ OK ───────│
   │              │                    │                │              │
   │              │◀─ Trigger ─────────────────────────│              │
   │              │  (score > threshold)                │              │
   │              │─ notify frontend to show widget     │              │
```

---

## Sequence 3 — Document Processing Flow

```
Customer     Frontend     Backend     DocAgent     ClaudeAPI    DB

   │─ Upload ──▶│           │            │             │          │
   │  PAN.jpg  │─ POST ───▶│            │             │          │
   │           │  /documents│            │             │          │
   │           │            │─ Validate ─│             │          │
   │           │            │  file type │             │          │
   │           │            │─ Store ────│             │          │
   │           │            │  to object │             │          │
   │           │            │─ POST ────▶│             │          │
   │           │            │  /extract  │             │          │
   │           │            │            │─ Pre-process│          │
   │           │            │            │  (resize,   │          │
   │           │            │            │  enhance)   │          │
   │           │            │            │─ POST ─────▶│          │
   │           │            │            │  Claude:    │          │
   │           │            │            │  extract    │          │
   │           │            │            │  PAN fields │          │
   │           │            │            │◀─ {name,   │          │
   │           │            │            │  pan_no,   │          │
   │           │            │            │  dob,      │          │
   │           │            │            │  confidence}│          │
   │           │            │            │─ Validate ─│          │
   │           │            │            │  PAN format │          │
   │           │            │            │─ POST ─────│─────────▶│
   │           │            │◀─ fields ─│  persist   │          │
   │           │◀─ prefill ─│            │             │          │
   │◀──────────│            │            │             │          │
   │  [Review fields, submit]             │             │          │
```

---

## Sequence 4 — Consent Revocation Flow

```
Customer     Frontend      Backend      ConsentSvc     DB        DocStore

   │─ Open ───▶│             │              │            │           │
   │  Consents │─ GET ──────▶│              │            │           │
   │           │  /consents  │─ List ──────▶│            │           │
   │           │             │              │─ Query ───▶│           │
   │           │             │              │◀─ rows ───│           │
   │           │◀─ consents ─│              │            │           │
   │◀──────────│             │              │            │           │
   │─ Revoke ─▶│             │              │            │           │
   │           │─ DELETE ───▶│              │            │           │
   │           │  /consents  │─ Revoke ────▶│            │           │
   │           │  /:id       │              │─ Update ──▶│           │
   │           │             │              │  status=REVOKED        │
   │           │             │              │─ Schedule ─│──────────▶│
   │           │             │              │  doc delete│  (async)  │
   │           │             │              │─ Notify ──▶│           │
   │           │             │              │  RM of revocation      │
   │           │◀─ receipt ─│              │            │           │
   │◀──────────│             │              │            │           │
   │  [Receipt + confirmation email]        │            │           │
```

---

## Sequence 5 — RM Dashboard Load + NBA

```
RM           Frontend       Backend       QualSvc     NBA Engine     DB

   │─ Login ──▶│             │              │              │           │
   │           │─ POST ─────▶│              │              │           │
   │           │  /auth/login│─ Validate ──▶│              │           │
   │           │◀─ JWT ─────│              │              │           │
   │           │             │              │              │           │
   │─ Opens ──▶│             │              │              │           │
   │  Dashboard│─ GET ──────▶│              │              │           │
   │           │  /rm/pipeline              │              │           │
   │           │             │─ Load ──────▶│              │           │
   │           │             │  qualified   │              │           │
   │           │             │  prospects   │              │           │
   │           │             │◀─ list ─────│              │           │
   │           │             │─ Run NBA ─────────────────▶│           │
   │           │             │  per prospect│              │           │
   │           │             │              │              │─ Query ──▶│
   │           │             │              │              │  intent + │
   │           │             │              │              │  qual data│
   │           │             │              │              │◀─ data ──│
   │           │             │              │              │─ Compute ─│
   │           │             │              │              │  next step│
   │           │◀─ pipeline ─│              │              │           │
   │  with NBA hints         │              │              │           │
   │◀──────────│             │              │              │           │
```

---

## Sequence 6 — Compliance Check (Embedded in Application Submit)

```
ApplicationService    ComplianceAgent    ConsentSvc    AuditSvc

        │─ Submit app ──▶│                   │              │
        │                │─ Check consent ──▶│              │
        │                │◀─ exists, valid ─│              │
        │                │─ Check purpose ──▶│              │
        │                │◀─ bound ─────────│              │
        │                │─ Check min data ──│              │
        │                │◀─ minimal ───────│              │
        │                │─ Policy check ────│              │
        │                │◀─ PASS ──────────│              │
        │◀─ APPROVED ───│                   │              │
        │─ Log ───────────────────────────────────────────▶│
        │  ComplianceCheck: PASSED          │              │
        │◀─ OK ────────────────────────────────────────────│
```
