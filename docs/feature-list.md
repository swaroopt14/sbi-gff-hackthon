# Aperture — Feature List

**Version:** 1.0  
**Date:** 2026-06-27

---

## MVP Features (Hackathon Demo)

### Module 1: Intent Detection
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 1.1 | Session Tracking | Track page visits, dwell time, click patterns | P0 |
| 1.2 | Intent Scoring | Weighted scoring across 5 signal types | P0 |
| 1.3 | Intent Classification | Hot / Warm / Cold / Drop-off Recovery | P0 |
| 1.4 | Campaign Attribution | Link intent to source campaign | P0 |
| 1.5 | Intent Timeline | Visual timeline of customer's intent signals | P1 |
| 1.6 | Re-engagement Trigger | Alert when a prospect returns after drop-off | P1 |

### Module 2: Qualification
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 2.1 | Multi-dimension Scoring | 5-dimension score with configurable weights | P0 |
| 2.2 | Product Fit Matrix | Match prospect profile to SBI product catalogue | P0 |
| 2.3 | Tier Classification | Hot / Warm / Cold with threshold configuration | P0 |
| 2.4 | Explainability Output | Reason + evidence per dimension | P0 |
| 2.5 | RM Override | Human override with audit log | P0 |
| 2.6 | Policy Rules Engine | Configurable eligibility rules per product | P1 |

### Module 3: Personalisation
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 3.1 | Personalised Message Generation | LLM-generated offer messages | P0 |
| 3.2 | Language Support | English + Hindi for MVP | P0 |
| 3.3 | EMI Calculator | Interactive EMI/savings calculator | P0 |
| 3.4 | Product Recommendation | Top 3 product recommendations with reasons | P0 |
| 3.5 | Offer Card | Visual offer card with key terms | P1 |

### Module 4: Conversation
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 4.1 | Chat Interface | Real-time chat widget | P0 |
| 4.2 | Structured Data Collection | Collect 8 required fields via conversation | P0 |
| 4.3 | Dynamic Flow | Adapt questions based on previous answers | P0 |
| 4.4 | Human Handoff | Transfer to RM with context | P0 |
| 4.5 | Conversation History | Full transcript per session | P0 |
| 4.6 | Typing Indicators | UX for AI response generation | P1 |

### Module 5: Document Processing
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 5.1 | Document Upload | Multi-file upload with preview | P0 |
| 5.2 | PAN Extraction | Extract name, PAN number, DOB, photo | P0 |
| 5.3 | Salary Slip Extraction | Extract employer, salary, date, deductions | P0 |
| 5.4 | Confidence Scoring | Field-level confidence scores | P0 |
| 5.5 | Manual Correction | Customer can correct AI-extracted fields | P0 |
| 5.6 | Document Status | Processing status per document | P1 |

### Module 6: Consent Management
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 6.1 | Consent Capture | Granular consent per data category | P0 |
| 6.2 | Consent Receipt | Generated PDF/screen receipt | P0 |
| 6.3 | Purpose Binding | Enforcement layer preventing data reuse | P0 |
| 6.4 | Consent Dashboard | View all active consents | P0 |
| 6.5 | Revocation | One-click revocation with data deletion | P0 |
| 6.6 | Retention Management | Automatic data deletion after retention period | P1 |

### Module 7: Explainability
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 7.1 | Decision Panel | Per-decision explanation with evidence | P0 |
| 7.2 | Confidence Indicators | Visual confidence score per AI output | P0 |
| 7.3 | Policy Reference | Which policy rule was applied | P0 |
| 7.4 | Customer View | Simplified explanation for customers | P0 |
| 7.5 | RM View | Detailed explanation for RMs | P0 |
| 7.6 | Compliance View | Full technical explanation for compliance | P0 |

### Module 8: Audit
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 8.1 | Event Logging | Immutable log for every system event | P0 |
| 8.2 | Actor Attribution | Who did what, when | P0 |
| 8.3 | Audit Timeline | Visual timeline per prospect | P0 |
| 8.4 | Risk Flags | Automated flagging of suspicious events | P0 |
| 8.5 | Export | CSV/PDF export of audit logs | P1 |
| 8.6 | Search & Filter | Query audit logs by actor, event, date | P1 |

### Module 9: Dashboards
| # | Feature | Description | Priority |
|---|---------|-------------|----------|
| 9.1 | Customer Dashboard | Application status + consent + docs | P0 |
| 9.2 | RM Dashboard | Pipeline + intent + next best action | P0 |
| 9.3 | Compliance Dashboard | Consent health + flags + overrides | P0 |
| 9.4 | Operations Dashboard | Funnel + conversion + drop-off analytics | P0 |
| 9.5 | Admin Dashboard | User mgmt + agent config + system health | P1 |

---

## Future Features (Post-Hackathon)

| # | Feature | Phase |
|---|---------|-------|
| F1 | Voice conversation (multilingual) | Phase 5 |
| F2 | 6+ language support (Tamil, Telugu, Kannada, Malayalam) | Phase 5 |
| F3 | Account Aggregator integration (Saafe/OneMoney) | Phase 6 |
| F4 | CKYCR live integration | Phase 6 |
| F5 | UIDAI Aadhaar e-KYC | Phase 6 |
| F6 | DigiLocker document retrieval | Phase 6 |
| F7 | V-CIP (Video KYC) live integration | Phase 6 |
| F8 | SBI Core Banking API integration | Phase 6 |
| F9 | Mobile app (iOS/Android) | Phase 7 |
| F10 | Real-time Kafka event streaming | Phase 7 |
| F11 | ML model fine-tuning on SBI data | Phase 8 |
| F12 | Multi-bank white-labelling | Phase 9 |

---

## Product Metrics (Success Criteria for Demo)

| Metric | Target |
|--------|--------|
| Intent detection time | < 200ms |
| Qualification score generation | < 500ms |
| Document OCR accuracy | > 90% on clean images |
| Consent capture completion | 100% before data use |
| Conversation completion rate | > 80% in demo |
| Audit event latency | < 100ms write time |
| Dashboard load time | < 2 seconds |
