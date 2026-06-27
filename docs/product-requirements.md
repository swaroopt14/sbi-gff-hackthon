# Aperture — Product Requirements Document (PRD)

**Version:** 1.0  
**Date:** 2026-06-27  
**Author:** Aperture Engineering Team  
**Status:** Approved for Development

---

## 1. Executive Summary

Aperture is a consent-led, agentic AI customer acquisition platform built for the State Bank of India. It transforms SBI's customer acquisition model from reactive (waiting for customers) to proactive (identifying, qualifying, personalising, and converting high-intent prospects inside SBI-owned channels).

The platform is governed by RBI compliance requirements and built on explainability-first AI principles.

---

## 2. Problem Statement

### Current State
- Banks wait for customers to visit branches or fill web forms
- No intelligent intent detection across YONO, SBI website, or campaign channels
- Qualification is manual and generic
- No personalisation of product offers
- High customer acquisition cost (₹400–₹1,200 per lead)
- Poor digital onboarding completion (high drop-off)
- No audit trail for AI-assisted decisions

### Root Causes
1. Siloed data across YONO, CBS, CRM, and campaign systems
2. No real-time intent processing pipeline
3. Absence of AI-driven qualification scoring
4. Legacy form-based onboarding with no conversational assistance
5. Compliance treated as an afterthought rather than a foundation

---

## 3. Product Vision

> "Give every SBI customer the experience of having a personal banker available 24/7 — one who knows them, understands their needs, and guides them through the right product in minutes."

---

## 4. Functional Requirements

### 4.1 Intent Detection

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-01 | System SHALL detect high-intent sessions based on page visits, dwell time, click patterns, and product exploration | Must Have |
| FR-02 | System SHALL capture campaign attribution (SMS/email/QR/YONO push) as an intent signal | Must Have |
| FR-03 | System SHALL identify abandoned applications and re-engage prospects | Must Have |
| FR-04 | System SHALL assign an Intent Score (0–100) with weighted signal breakdown | Must Have |
| FR-05 | System SHALL categorise intent: Product Discovery / Active Comparison / Ready to Apply / Drop-off Recovery | Must Have |
| FR-06 | System SHALL not infer sensitive life events (marriage, pregnancy, illness) from browsing data | Must Have |

### 4.2 Qualification

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-07 | System SHALL score prospects on five dimensions: Intent Strength, Eligibility, KYC Readiness, Product Fit, Completion Probability | Must Have |
| FR-08 | System SHALL produce a composite qualification score (0–100) with tier classification: Hot / Warm / Cold | Must Have |
| FR-09 | Qualification SHALL use only customer-declared information, SBI internal profile, and post-consent data | Must Have |
| FR-10 | Every qualification decision SHALL include a human-readable explanation with evidence | Must Have |
| FR-11 | Compliance Officer SHALL be able to override any qualification score with audit trail | Must Have |
| FR-12 | System SHALL support configurable eligibility policy rules per product | Must Have |

### 4.3 Personalisation

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-13 | System SHALL generate personalised offer messages based on prospect profile | Must Have |
| FR-14 | System SHALL support at least 6 languages: English, Hindi, Tamil, Telugu, Kannada, Malayalam | Must Have |
| FR-15 | System SHALL recommend the most suitable SBI product for each prospect | Must Have |
| FR-16 | System SHALL display an interactive EMI/savings calculator during conversation | Must Have |
| FR-17 | Personalisation SHALL be explainable: "This offer was generated because..." | Must Have |

### 4.4 Conversational Onboarding

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-18 | System SHALL provide a conversational AI interface to guide prospects through onboarding | Must Have |
| FR-19 | Conversation Agent SHALL collect: Name, PAN, Income, Occupation, Purpose, Consent | Must Have |
| FR-20 | Conversation SHALL adapt dynamically based on prospect responses | Must Have |
| FR-21 | System SHALL support handoff to human RM when confidence is low or customer requests it | Must Have |
| FR-22 | System SHALL support V-CIP (Video Customer Identification Process) handoff for KYC completion | Should Have |

### 4.5 Document Processing

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-23 | System SHALL accept document uploads: PAN card, Aadhaar (offline XML), Salary slip, Bank statement | Must Have |
| FR-24 | System SHALL extract structured data from documents using OCR/AI | Must Have |
| FR-25 | System SHALL return a confidence score per extracted field | Must Have |
| FR-26 | System SHALL flag missing or unclear fields for manual review | Must Have |
| FR-27 | System SHALL NOT store Aadhaar number in full; SHALL mask after use | Must Have |

### 4.6 Consent Management

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-28 | System SHALL present clear purpose statements before any data collection | Must Have |
| FR-29 | System SHALL capture explicit, granular consent per data category | Must Have |
| FR-30 | System SHALL generate a consent receipt with timestamp, purpose, data categories, and retention period | Must Have |
| FR-31 | System SHALL allow consent revocation at any time | Must Have |
| FR-32 | System SHALL enforce purpose binding — data collected for purpose X cannot be used for purpose Y | Must Have |
| FR-33 | System SHALL delete non-legally-required data upon consent revocation | Must Have |

### 4.7 Explainability

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-34 | Every AI output SHALL include: Reason, Confidence, Evidence, Policy Used, Recommendation, Next Step | Must Have |
| FR-35 | Explainability logs SHALL be visible to: Customer (simplified), RM (detailed), Compliance (full) | Must Have |
| FR-36 | System SHALL maintain a complete decision trail for every prospect journey | Must Have |

### 4.8 Audit

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-37 | System SHALL maintain immutable audit logs for every event | Must Have |
| FR-38 | Audit logs SHALL capture: Event, Actor, Timestamp, Input, Output, Policy Applied | Must Have |
| FR-39 | Compliance Officer SHALL be able to query and export audit logs | Must Have |
| FR-40 | System SHALL flag events that require human review | Must Have |

### 4.9 Dashboards

| ID | Requirement | Priority |
|----|-------------|----------|
| FR-41 | Customer Dashboard: application status, consent records, documents, next steps | Must Have |
| FR-42 | RM Dashboard: prospect pipeline, intent signals, qualification scores, next best actions | Must Have |
| FR-43 | Compliance Dashboard: consent status, audit flags, policy violations, override log | Must Have |
| FR-44 | Operations Dashboard: funnel metrics, conversion rates, drop-off analysis, volume | Must Have |
| FR-45 | Admin Dashboard: user management, agent configuration, policy rules, system health | Should Have |

---

## 5. Non-Functional Requirements

### 5.1 Performance

| ID | Requirement | Target |
|----|-------------|--------|
| NFR-01 | API response time (p95) | < 500ms |
| NFR-02 | Intent detection latency | < 200ms |
| NFR-03 | Document OCR processing | < 10 seconds |
| NFR-04 | Concurrent prospect sessions | 10,000+ |
| NFR-05 | System availability | 99.9% |

### 5.2 Security

| ID | Requirement |
|----|-------------|
| NFR-06 | All data at rest: AES-256 encryption |
| NFR-07 | All data in transit: TLS 1.3 |
| NFR-08 | Authentication: OAuth2 + JWT with short expiry |
| NFR-09 | PII data: field-level encryption in database |
| NFR-10 | Aadhaar number: masked/tokenised after verification |
| NFR-11 | No PII in application logs |
| NFR-12 | Role-based access control (RBAC) with principle of least privilege |

### 5.3 Compliance

| ID | Requirement |
|----|-------------|
| NFR-13 | RBI Master Direction on Digital KYC compliance |
| NFR-14 | RBI Master Direction on Information Technology compliance |
| NFR-15 | DPDP Act 2023 compliance (consent, purpose, minimisation, revocation) |
| NFR-16 | Aadhaar Act 2016 — no full Aadhaar storage, no demographic share |
| NFR-17 | PCI-DSS — no raw financial data stored beyond necessity |
| NFR-18 | Audit logs retained for 10 years per RBI requirement |

### 5.4 Scalability

| ID | Requirement |
|----|-------------|
| NFR-19 | Horizontal scalability — stateless services |
| NFR-20 | Kafka for async event processing |
| NFR-21 | Redis for session caching (TTL managed) |
| NFR-22 | Database read replicas for analytics queries |

### 5.5 Observability

| ID | Requirement |
|----|-------------|
| NFR-23 | Structured JSON logging (no PII) |
| NFR-24 | Distributed tracing with trace IDs |
| NFR-25 | Health check endpoints on all services |
| NFR-26 | Metric dashboards (Prometheus + Grafana) |

---

## 6. MVP Scope vs Future

### MVP (Hackathon Demo)

| Module | Included |
|--------|----------|
| Intent Detection Agent | Yes — session-based scoring |
| Qualification Agent | Yes — 5-dimension scoring |
| Personalisation Agent | Yes — English + Hindi |
| Conversation Agent | Yes — text-based |
| Document Agent | Yes — PAN + Salary Slip OCR |
| Consent Management | Yes — full consent flow |
| Explainability | Yes — all decisions |
| Audit Trail | Yes — all events |
| Customer Dashboard | Yes |
| RM Dashboard | Yes |
| Compliance Dashboard | Yes |
| Operations Dashboard | Yes |
| Admin Dashboard | Partial |

### Future Scope (Post-Hackathon)

| Feature | Phase |
|---------|-------|
| Voice AI (Gnani.ai / Posibl) | Phase 5 |
| 6+ languages full support | Phase 5 |
| Account Aggregator integration | Phase 6 |
| CKYCR live integration | Phase 6 |
| V-CIP live integration | Phase 6 |
| DigiLocker integration | Phase 6 |
| Kubernetes deployment | Phase 7 |
| Multi-bank white-labelling | Phase 8 |

---

## 7. Assumptions

1. SBI internal data (YONO profile, existing relationship) is available via mock API for demo
2. CKYCR, NSDL, and UIDAI integrations are mocked for demo
3. Kafka is used in mock mode for demo (no real broker required for MVP)
4. Account Aggregator consent flow is demonstrated via mock
5. V-CIP handoff is simulated in demo
6. AI agents use Claude API (claude-sonnet-4-6) via Anthropic SDK
7. PAN and Aadhaar extraction uses mock OCR for demo

---

## 8. Risks

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| AI hallucination in qualification | Medium | High | Confidence thresholds, human review triggers |
| Regulatory interpretation ambiguity | Medium | High | Compliance officer review gate at every step |
| Document OCR accuracy on low-quality images | High | Medium | Confidence scoring + manual review fallback |
| Session data latency affecting intent scoring | Low | Medium | Redis cache + async processing |
| Consent fatigue leading to user drop-off | Medium | High | Single-screen consent with clear language |

---

## 9. Acceptance Criteria

| Feature | Acceptance Criterion |
|---------|---------------------|
| Intent Score | Score generated within 200ms with ≥3 signal sources |
| Qualification | Score 0–100 with explanation, tier assigned, RM can override |
| Conversation | Collects all required fields, handles ambiguous input, falls back gracefully |
| Document OCR | Extracts 5+ fields from PAN/salary slip with confidence >0.8 |
| Consent | Consent receipt generated, purpose bound, revocable in UI |
| Audit | Every event logged with actor, timestamp, input, output |
| Explainability | Every AI output has reason, confidence, evidence, policy |
| Dashboards | Load within 2 seconds, all data accurate, role-gated |
