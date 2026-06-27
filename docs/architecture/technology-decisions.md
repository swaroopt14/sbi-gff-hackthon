# Aperture — Technology Decisions

**Version:** 1.0  
**Date:** 2026-06-27

---

## Decision Framework

Each technology decision is evaluated on: Performance, Ecosystem, Compliance, Team Familiarity, and Operational Cost.

---

## Backend Language: Go

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| Performance | 5/5 | Handles 10,000+ concurrent requests with low memory |
| Banking fitness | 5/5 | Used by banks (ING, Monzo, N26, SBI FinTech) |
| Type safety | 5/5 | Compile-time safety reduces runtime errors |
| Concurrency | 5/5 | Goroutines ideal for agent orchestration |
| Container size | 5/5 | Small Docker images (~20MB) |
| Ecosystem | 4/5 | Excellent HTTP/DB/Kafka libraries |

**Alternatives Considered:**
- Java (Spring Boot): Mature but heavy; JVM startup overhead
- Node.js: Fast to write but weak type safety for financial logic
- Python: Agent layer only — not suitable for high-performance APIs

**Decision: Go + Gin for backend API. Python + FastAPI for AI agents.**

---

## Frontend Framework: Next.js 14 (App Router)

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| SSR/SSG | 5/5 | Server components reduce client bundle size |
| TypeScript | 5/5 | Native TypeScript support |
| Ecosystem | 5/5 | shadcn/ui, Framer Motion, React Query all first-class |
| Performance | 5/5 | Streaming + Suspense for progressive loading |
| Auth | 4/5 | Middleware for route protection |

**Decision: Next.js 14 with App Router, TypeScript, Tailwind CSS, shadcn/ui**

---

## AI Provider: Anthropic Claude (claude-sonnet-4-6)

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| Reasoning quality | 5/5 | Best in class for structured output + compliance |
| Safety | 5/5 | Constitutional AI — critical for banking context |
| Structured output | 5/5 | Native JSON mode, function calling |
| Context window | 5/5 | 200K tokens for long document processing |
| Explainability | 5/5 | Detailed reasoning traces |
| Cost | 4/5 | Competitive for the quality tier |

**Alternatives Considered:**
- GPT-4o: Excellent but OpenAI TOS may conflict with banking data
- Llama 3 (local): Self-hosted but requires significant infra; quality gap
- Gemini 2.0 Flash: Fast, but less structured output reliability

**Decision: Claude claude-sonnet-4-6 as primary LLM. Abstraction layer allows future swap.**

**Abstraction Design:**
```go
type LLMClient interface {
    Complete(ctx context.Context, req CompletionRequest) (CompletionResponse, error)
    StreamComplete(ctx context.Context, req CompletionRequest) (<-chan string, error)
}
// Implementations: ClaudeClient, OpenAIClient, GeminiClient
```

---

## Database: PostgreSQL 16

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| ACID compliance | 5/5 | Required for financial data |
| JSON support | 5/5 | JSONB for flexible agent outputs |
| Performance | 4/5 | Excellent with proper indexing |
| Audit trails | 5/5 | Native triggers, partitioning for audit logs |
| Compliance | 5/5 | Row-level security for multi-role access |

**Alternatives Considered:**
- MySQL: Less powerful JSON + window functions
- MongoDB: Non-ACID at transaction level; unsuitable for financial data
- CockroachDB: Overkill for MVP

**Decision: PostgreSQL 16 with read replicas for analytics queries.**

---

## Cache: Redis 7

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| Session storage | 5/5 | TTL-managed, atomic operations |
| Rate limiting | 5/5 | Sliding window counter support |
| Pub/Sub | 4/5 | Real-time notifications (intent triggers) |
| Conversation state | 5/5 | Fast read/write for chat turns |

**Decision: Redis 7 with Redis Cluster in production.**

---

## Message Queue: Kafka (Mock for MVP)

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| Intent event streaming | 5/5 | High-throughput, ordered, replayable |
| Audit event pipeline | 5/5 | Append-only, durable |
| Async agent triggers | 5/5 | Decouples intent detection from API path |

**MVP Decision:** Kafka producer/consumer mocked via in-process channels. Same interface, swap to real Kafka with one config change.

**Interface:**
```go
type EventBus interface {
    Publish(ctx context.Context, topic string, event Event) error
    Subscribe(topic string, handler EventHandler) error
}
// MockEventBus: in-process
// KafkaEventBus: real Kafka
```

---

## Agent Framework: LangChain / FastAPI

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| LLM abstraction | 5/5 | Model-agnostic |
| Tool use | 5/5 | Function calling integration |
| Memory | 5/5 | Redis + in-memory memory classes |
| Structured output | 5/5 | Pydantic output parsers |
| HTTP interface | 5/5 | FastAPI for clean REST exposure |

**Decision: Each agent = standalone FastAPI service + LangChain chains. No LangServe for MVP (adds complexity).**

---

## Authentication: JWT (RS256)

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| Stateless | 5/5 | Backend can scale horizontally |
| Role embedding | 5/5 | Roles in JWT claims |
| Security | 5/5 | RS256 (asymmetric) — private key signs, public key verifies |
| Refresh | 4/5 | Refresh token in HttpOnly cookie |

**Token Lifetime:**
- Access token: 15 minutes
- Refresh token: 7 days (rotated on use)

---

## Containerisation: Docker + Docker Compose

| Criterion | Score | Rationale |
|-----------|-------|-----------|
| Dev parity | 5/5 | Every service in container |
| Setup time | 5/5 | `docker compose up` = full stack |
| Postgres seeding | 5/5 | Init scripts for demo data |
| Networking | 5/5 | Internal DNS between services |

**Production upgrade path:** Docker Compose → Kubernetes (same images, different orchestrator)

---

## Summary Decision Matrix

| Component | Choice | Priority |
|-----------|--------|----------|
| Backend API | Go 1.22 + Gin | P0 |
| Frontend | Next.js 14 + TypeScript | P0 |
| UI Library | shadcn/ui + Tailwind | P0 |
| AI Agents | Python 3.11 + FastAPI + LangChain | P0 |
| LLM | Claude claude-sonnet-4-6 | P0 |
| Database | PostgreSQL 16 | P0 |
| Cache | Redis 7 | P0 |
| Message Queue | Kafka (mock) | P0 |
| Auth | JWT RS256 + HttpOnly refresh | P0 |
| Container | Docker + Docker Compose | P0 |
| Animations | Framer Motion | P1 |
| Charts | Recharts | P1 |
| State | Zustand + React Query | P0 |
| Forms | React Hook Form + Zod | P0 |
| Logging | Zap (Go) + structlog (Python) | P0 |
| Validation | go-playground/validator | P0 |
