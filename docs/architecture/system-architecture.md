# Aperture — System Architecture

**Version:** 1.0  
**Date:** 2026-06-27

---

## 1. Overview

Aperture is a distributed, event-driven, microservice-adjacent monolith (modular monolith for MVP, decomposable to microservices for production). It is composed of:

- A **Go/Gin backend** exposing REST APIs
- A **Next.js frontend** with role-based dashboards
- A **Python-based agent layer** orchestrating 8 AI agents
- A **PostgreSQL** primary data store
- A **Redis** cache for sessions and short-lived state
- A **Kafka** event bus (mocked for MVP)
- A **Claude API** integration for AI reasoning

---

## 2. Six-Layer Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                          APERTURE PLATFORM                                 │
└─────────────────────────────────────────────────────────────────────────────┘

╔═════════════════════════════════════════════════════════════════════════════╗
║  LAYER 1 — CHANNEL LAYER                                                   ║
║  SBI Website | YONO App (mock) | QR Campaign | Branch Tablet | RM App     ║
║  Entry point for all customer interactions                                 ║
╚═════════════════════════════════════════════════════════════════════════════╝
                                      │
                              Next.js Frontend
                                      │
╔═════════════════════════════════════════════════════════════════════════════╗
║  LAYER 2 — API GATEWAY LAYER                                               ║
║  Authentication │ Rate Limiting │ Request Routing │ CORS │ Logging        ║
║  Go/Gin Backend — /api/v1/*                                               ║
╚═════════════════════════════════════════════════════════════════════════════╝
                                      │
╔═════════════════════════════════════════════════════════════════════════════╗
║  LAYER 3 — CONSENT LAYER (enforced before any data access)                ║
║  ConsentService │ PurposeEngine │ RevocationManager │ ConsentReceipt      ║
╚═════════════════════════════════════════════════════════════════════════════╝
                                      │
╔═════════════════════════════════════════════════════════════════════════════╗
║  LAYER 4 — DECISION LAYER (Agentic AI Orchestration)                      ║
║                                                                             ║
║  ┌─────────────┐  ┌──────────────┐  ┌───────────────┐  ┌──────────────┐  ║
║  │   Intent    │  │ Qualification│  │Personalisation│  │ Conversation │  ║
║  │   Agent     │→ │    Agent     │→ │    Agent      │→ │    Agent     │  ║
║  └─────────────┘  └──────────────┘  └───────────────┘  └──────────────┘  ║
║                                                                             ║
║  ┌─────────────┐  ┌──────────────┐  ┌───────────────┐  ┌──────────────┐  ║
║  │  Document   │  │  Compliance  │  │Explainability │  │    Audit     │  ║
║  │   Agent     │  │    Agent     │  │    Agent      │  │    Agent     │  ║
║  └─────────────┘  └──────────────┘  └───────────────┘  └──────────────┘  ║
╚═════════════════════════════════════════════════════════════════════════════╝
                                      │
╔═════════════════════════════════════════════════════════════════════════════╗
║  LAYER 5 — DATA LAYER                                                      ║
║  PostgreSQL (primary) │ Redis (cache/session) │ S3-compatible (documents)  ║
╚═════════════════════════════════════════════════════════════════════════════╝
                                      │
╔═════════════════════════════════════════════════════════════════════════════╗
║  LAYER 6 — GOVERNANCE LAYER                                               ║
║  Audit Trail │ Explainability Logs │ Human Review Queue │ Model Monitor   ║
╚═════════════════════════════════════════════════════════════════════════════╝
```

---

## 3. Service Decomposition

### 3.1 Backend Services (Go — Modular Monolith)

```
backend/
├── auth/           JWT auth, refresh, RBAC
├── customers/      Customer/prospect CRUD, search
├── sessions/       Intent session tracking
├── intent/         Intent score computation
├── qualification/  Qualification scoring engine
├── conversation/   Chat session management
├── documents/      Upload, OCR trigger, storage
├── consent/        Consent capture, enforcement, revocation
├── applications/   Loan application lifecycle
├── audit/          Immutable event logging
├── explainability/ Explanation storage and retrieval
├── recommendations/ NBA engine
└── admin/          User management, config
```

### 3.2 Agent Services (Python — FastAPI)

```
agents/
├── intent-agent/        Session signal → Intent score
├── qualification-agent/ Profile → Qualification score
├── personalisation-agent/ Profile → Offer message
├── conversation-agent/  Chat turn → Response + data extraction
├── document-agent/      Document image → Extracted fields
├── compliance-agent/    Action → Policy check
├── explainability-agent/ Decision → Human explanation
└── audit-agent/         Event → Audit record
```

### 3.3 Frontend (Next.js)

```
frontend/
├── app/
│   ├── (auth)/         Login, register
│   ├── (customer)/     Customer dashboard, chat, consent, docs
│   ├── (rm)/           RM dashboard, pipeline, lead detail
│   ├── (compliance)/   Compliance dashboard, flags, audit
│   ├── (operations)/   Funnel analytics, operations dashboard
│   └── (admin)/        Admin settings, user management
```

---

## 4. Data Flow

```
[Customer Session Start]
        │
        ▼
[Session Service]
  Creates session ID
  Starts event collection
        │
        ▼
[Intent Events streamed to Kafka topic: intent.events]
        │
        ▼
[Intent Agent consumes events]
  Computes score
  Classifies intent
  Persists IntentScore to DB
        │
        ▼
[Trigger: score > threshold]
        │
        ▼
[Personalisation Agent]
  Generates offer message
  Selects product recommendation
        │
        ▼
[Frontend: Chat Widget appears]
        │
        ▼
[Conversation Agent]
  Manages turn-by-turn dialogue
  Extracts structured fields
  Each turn logged to AuditLog
        │
        ▼
[Consent Service]
  Before any PII collection:
    Purpose presented
    Consent captured
    Receipt generated
        │
        ▼
[Document Service → Document Agent]
  File received
  OCR extraction
  Confidence scored
  Fields returned for pre-fill
        │
        ▼
[Qualification Agent — Full Scoring]
  All signals + declared data + document data
  Score computed
  Explanation generated
        │
        ▼
[Application Service]
  Application created with pre-filled data
  Submitted to review queue
        │
        ▼
[Compliance Agent]
  Validates consent is present
  Validates data minimisation
  Policy check passes
        │
        ▼
[Audit Agent]
  Logs: ApplicationSubmitted event
  Logs: All prior events in transaction
        │
        ▼
[Explainability Agent]
  Generates full decision explanation
  Stores in ExplanationStore
  Available to: Customer, RM, Compliance
```

---

## 5. Integration Points

| Integration | Type | MVP | Production |
|------------|------|-----|------------|
| Claude API (Anthropic) | HTTP | Live | Live |
| PostgreSQL | TCP | Live | Live |
| Redis | TCP | Live | Live |
| Kafka | TCP | Mock | Live |
| NSDL PAN Verification | HTTP | Mock | Live |
| CKYCR | HTTP | Mock | Live |
| UIDAI Aadhaar e-KYC | HTTP | Mock | Live |
| Account Aggregator | HTTP | Mock | Live |
| DigiLocker | HTTP | Mock | Live |
| SBI Core Banking | HTTP | Mock | Live |

---

## 6. Security Architecture

```
[Client]
    │
    │ HTTPS (TLS 1.3)
    │
[API Gateway]
    │ JWT Validation (RS256)
    │ Rate Limiting (Redis-backed)
    │ CORS Policy
    │
[Backend Services]
    │ RBAC — Role checked per endpoint
    │ Input validation (Go validator)
    │ No PII in logs (field masking)
    │
[Data Layer]
    │ Encryption at rest (AES-256)
    │ Field-level encryption for PII
    │ Aadhaar masked after use
    │ Connection pool with TLS
    │
[Document Storage]
    └── Object store with signed URLs (TTL 15 min)
```

---

## 7. Scalability Design

```
[Load Balancer]
      │
  ┌───┴────┐
  │   Go   │  (n replicas — stateless)
  │ Backend│
  └───┬────┘
      │
  ┌───┴──────────┐
  │ Redis Cluster │  Session cache, rate limit counters
  └───┬───────────┘
      │
  ┌───┴────────────┐
  │   PostgreSQL   │  Primary + Read Replicas
  │   Cluster      │
  └───┬────────────┘
      │
  ┌───┴──────┐
  │  Kafka   │  Topic-partitioned for parallel processing
  │ Cluster  │
  └──────────┘
```

---

## 8. Agent Orchestration

Each agent is independently callable via HTTP (FastAPI). The Go backend acts as the orchestrator, calling agents in sequence or parallel depending on the pipeline stage.

```
Go Backend (Orchestrator)
    │
    ├──[POST /agents/intent/score]──────── Intent Agent (FastAPI)
    │                                           │
    │                                     Returns: IntentScoreResponse
    │
    ├──[POST /agents/qualify/score]──────── Qualification Agent (FastAPI)
    │                                           │
    │                                     Returns: QualificationResponse
    │
    ├──[POST /agents/personalise]─────────── Personalisation Agent (FastAPI)
    │
    ├──[POST /agents/chat]─────────────────── Conversation Agent (FastAPI)
    │
    ├──[POST /agents/document/extract]──────── Document Agent (FastAPI)
    │
    ├──[POST /agents/compliance/check]────── Compliance Agent (FastAPI)
    │
    ├──[POST /agents/explain]───────────── Explainability Agent (FastAPI)
    │
    └──[POST /agents/audit/log]────────────── Audit Agent (FastAPI)
```

---

## 9. Technology Decisions Summary

See [technology-decisions.md](technology-decisions.md) for rationale.

| Component | Technology | Reason |
|-----------|-----------|--------|
| Backend | Go + Gin | Performance, concurrency, type safety |
| Frontend | Next.js + TypeScript | SSR, app router, ecosystem |
| UI | Tailwind + shadcn/ui | Rapid consistent UI |
| Agents | Python + FastAPI | LangChain/LangGraph ecosystem |
| AI Provider | Claude (claude-sonnet-4-6) | Best reasoning, compliance focus |
| Database | PostgreSQL | ACID, complex queries, JSON support |
| Cache | Redis | Session TTL, rate limiting |
| Events | Kafka (mock) | Async intent processing |
| Auth | JWT (RS256) | Stateless, role-embeddable |
| Container | Docker + Docker Compose | Local dev parity |
