# Aperture — Component Diagram

**Version:** 1.0  
**Date:** 2026-06-27

---

## Core Component Map

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          APERTURE COMPONENTS                               │
└─────────────────────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────────────────┐
│ FRONTEND (Next.js 14 / App Router)                                       │
│                                                                           │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌─────────────┐  │
│  │  Auth Pages  │  │  Customer    │  │  RM Portal   │  │ Compliance  │  │
│  │  /login      │  │  Dashboard   │  │  Dashboard   │  │  Dashboard  │  │
│  │  /register   │  │  /chat       │  │  /pipeline   │  │  /audit     │  │
│  └──────────────┘  │  /consent    │  │  /leads      │  │  /flags     │  │
│                    │  /documents  │  │  /leads/:id  │  └─────────────┘  │
│  ┌──────────────┐  │  /apply      │  └──────────────┘                   │
│  │  Operations  │  └──────────────┘                   ┌─────────────┐  │
│  │  Dashboard   │                                      │    Admin    │  │
│  │  /analytics  │  ┌─────────────────────────────────┐ │  Dashboard  │  │
│  │  /funnel     │  │         Shared Components        │ │  /settings  │  │
│  └──────────────┘  │  <Navbar> <Sidebar> <ChatWidget> │ └─────────────┘  │
│                    │  <IntentTimeline> <ScoreCard>    │                   │
│                    │  <ConsentModal> <DocUpload>       │                   │
│                    │  <ExplainPanel> <AuditTimeline>  │                   │
│                    └─────────────────────────────────┘                   │
└────────────────────────────────────┬─────────────────────────────────────┘
                                     │ HTTP / REST (Next.js API layer)
                                     ▼
┌──────────────────────────────────────────────────────────────────────────┐
│ BACKEND (Go / Gin — Modular Monolith)                                    │
│                                                                           │
│  ┌─────────────────────────────────────────────────────────────────────┐ │
│  │ API Layer (Gin Router)                                              │ │
│  │ /api/v1/auth  /customers  /intent  /qualify  /chat  /documents     │ │
│  │ /consent  /applications  /audit  /explain  /recommend  /admin      │ │
│  └──────────────────────────┬──────────────────────────────────────────┘ │
│                             │                                             │
│  ┌──────────────────────────▼──────────────────────────────────────────┐ │
│  │ Middleware Chain                                                    │ │
│  │ [Auth JWT] → [RBAC] → [RateLimit] → [RequestID] → [Logger]        │ │
│  └──────────────────────────┬──────────────────────────────────────────┘ │
│                             │                                             │
│  ┌──────────────────────────▼──────────────────────────────────────────┐ │
│  │ Service Layer (Business Logic)                                      │ │
│  │                                                                     │ │
│  │  AuthService       CustomerService    IntentService                 │ │
│  │  QualifyService    PersonalService    ConversationService           │ │
│  │  DocumentService   ConsentService     ApplicationService            │ │
│  │  AuditService      ExplainService     RecommendationService         │ │
│  │  AdminService      AgentOrchestrator                                │ │
│  └──────────────────────────┬──────────────────────────────────────────┘ │
│                             │                                             │
│  ┌──────────────────────────▼──────────────────────────────────────────┐ │
│  │ Repository Layer (Data Access)                                      │ │
│  │                                                                     │ │
│  │  UserRepo          CustomerRepo        SessionRepo                  │ │
│  │  IntentRepo        QualificationRepo   ConversationRepo             │ │
│  │  DocumentRepo      ConsentRepo         ApplicationRepo              │ │
│  │  AuditRepo         ExplanationRepo     RecommendationRepo           │ │
│  └──────────────────────────┬──────────────────────────────────────────┘ │
│                             │                                             │
│  ┌──────────────────────────▼──────────────────────────────────────────┐ │
│  │ Infrastructure Layer                                                │ │
│  │                                                                     │ │
│  │  PostgresDB    RedisClient    KafkaProducer    ObjectStore          │ │
│  │  AgentClient   LLMClient      EmailClient      MetricsClient        │ │
│  └─────────────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────────────────────────────────┘
         │                    │                      │
         ▼                    ▼                      ▼
    PostgreSQL            Redis Cache          Agent Services
    (Primary DB)          (Sessions,           (Python / FastAPI)
                           Rate Limits)
```

---

## Agent Components

```
┌────────────────────────────────────────────────────────────────────────────┐
│ AGENT LAYER (Python / FastAPI — each independently deployable)            │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Intent Agent (:8101)                                               │   │
│  │  signal_collector → signal_aggregator → scorer → classifier       │   │
│  │  Tools: session_reader, page_event_parser                          │   │
│  │  Memory: short-term (Redis TTL 30min)                              │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Qualification Agent (:8102)                                        │   │
│  │  profile_reader → policy_engine → scorer → tier_classifier        │   │
│  │  Tools: product_eligibility_checker, income_estimator             │   │
│  │  Memory: none (stateless per request)                              │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Personalisation Agent (:8103)                                      │   │
│  │  profile_reader → product_matcher → message_generator             │   │
│  │  Tools: product_catalogue, language_selector, emi_calculator      │   │
│  │  Memory: customer profile (session-scoped)                         │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Conversation Agent (:8104)                                         │   │
│  │  turn_processor → response_generator → field_extractor            │   │
│  │  Tools: form_filler, language_detector, human_handoff_trigger     │   │
│  │  Memory: conversation history (Redis, per session)                 │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Document Agent (:8105)                                             │   │
│  │  image_preprocessor → ocr_engine → field_extractor → validator   │   │
│  │  Tools: pan_parser, salary_parser, aadhaar_offline_parser        │   │
│  │  Memory: none (stateless)                                          │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Compliance Agent (:8106)                                           │   │
│  │  action_checker → policy_validator → risk_assessor → reporter    │   │
│  │  Tools: consent_verifier, data_minimisation_checker              │   │
│  │  Memory: policy rules (in-memory cache)                            │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Explainability Agent (:8107)                                       │   │
│  │  decision_reader → evidence_collector → explanation_generator    │   │
│  │  Tools: policy_reference, decision_history_reader                 │   │
│  │  Memory: none (reads from audit store)                             │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Audit Agent (:8108)                                                │   │
│  │  event_receiver → enricher → immutable_logger → flag_detector   │   │
│  │  Tools: timeline_builder, risk_flag_engine                        │   │
│  │  Memory: append-only (writes to PostgreSQL audit table)           │   │
│  └────────────────────────────────────────────────────────────────────┘   │
│                                                                            │
│  ┌────────────────────────────────────────────────────────────────────┐   │
│  │ Shared: LLM Client (Claude claude-sonnet-4-6 via Anthropic SDK)    │   │
│  │ Shared: Prompt Library (versioned YAML prompt templates)           │   │
│  │ Shared: Config (agent-level policy config in JSON)                 │   │
│  └────────────────────────────────────────────────────────────────────┘   │
└────────────────────────────────────────────────────────────────────────────┘
```

---

## Database Component

```
┌────────────────────────────────────────────────────────────────────────────┐
│ POSTGRESQL SCHEMA (see database/tables.md for full column definitions)    │
│                                                                            │
│  users ──────────────────── user_id (PK)                                  │
│  customers ───────────────── customer_id (PK) → user_id (FK)              │
│  intent_sessions ─────────── session_id (PK) → customer_id (FK)           │
│  intent_scores ──────────── score_id (PK) → session_id (FK)               │
│  qualification_scores ──── qual_id (PK) → customer_id (FK)               │
│  products ──────────────── product_id (PK)                                │
│  product_recommendations ── rec_id (PK) → customer_id, product_id (FK)   │
│  conversations ──────────── convo_id (PK) → customer_id (FK)              │
│  conversation_turns ─────── turn_id (PK) → convo_id (FK)                  │
│  documents ─────────────── doc_id (PK) → customer_id (FK)                │
│  extracted_fields ────────── field_id (PK) → doc_id (FK)                  │
│  consents ───────────────── consent_id (PK) → customer_id (FK)            │
│  consent_purposes ──────── purpose_id (PK) → consent_id (FK)             │
│  applications ───────────── app_id (PK) → customer_id, product_id (FK)   │
│  application_events ─────── event_id (PK) → app_id (FK)                   │
│  audit_logs ────────────── log_id (PK) → actor_id, entity_id              │
│  explanations ───────────── expl_id (PK) → entity_id, entity_type        │
│  policy_rules ───────────── rule_id (PK) → product_id (FK)               │
│  rm_overrides ───────────── override_id (PK) → qual_id, user_id (FK)     │
└────────────────────────────────────────────────────────────────────────────┘
```

---

## Communication Patterns

| From | To | Protocol | Sync/Async |
|------|----|----------|------------|
| Frontend | Go Backend | HTTPS REST | Sync |
| Go Backend | PostgreSQL | TCP/TLS | Sync |
| Go Backend | Redis | TCP | Sync |
| Go Backend | Kafka | TCP | Async |
| Go Backend | Agent Services | HTTP REST | Sync (with timeout) |
| Agent Services | Claude API | HTTPS | Sync |
| Kafka Consumer | Intent Agent | Internal | Async |
| Document Agent | Object Store | HTTPS | Sync |
