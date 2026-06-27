# Aperture
### Consent-Led Agentic AI Customer Acquisition Copilot for SBI

> Built for SBI Global FinTech Fest 2026 Hackathon

---

## What is Aperture?

Aperture transforms SBI's customer acquisition from reactive to proactive. Instead of waiting for customers to fill forms and walk into branches, Aperture intelligently identifies high-intent prospects inside SBI-owned channels, qualifies them through explainable AI, and guides them through RBI-compliant conversational onboarding.

**The acquisition journey Aperture automates:**

```
Identify → Qualify → Personalise → Convert → Onboard → Explain → Audit
```

---

## The Problem

Banks today are reactive. A customer visits the SBI website, browses a home loan page for 7 minutes, uses the EMI calculator, and then leaves. Nothing happens. No intelligent follow-up. No personalised offer. The lead is lost.

- **High acquisition cost:** ₹400–₹1,200 per lead via traditional channels
- **No intelligent qualification:** Generic scoring, no product fit analysis
- **Poor conversion:** High drop-off at form and document stages
- **No explainability:** AI decisions are black boxes — a compliance risk

---

## The Solution

Aperture's four-stage acquisition engine:

| Stage | What Happens |
|-------|-------------|
| **1. Identify** | Detects high-intent sessions inside YONO, SBI website, and campaigns using first-party session signals |
| **2. Qualify** | Runs explainable 5-dimension scoring to classify prospects as Hot/Warm/Cold |
| **3. Convert** | Serves hyper-personalised conversational offers in 6+ languages |
| **4. Onboard** | Guides through RBI-compliant KYC, document upload, and application in minutes |

---

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│  CHANNEL LAYER                                                  │
│  SBI Website | YONO | QR Campaign | Branch Tablet | RM App     │
└────────────────────────────┬────────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────────┐
│  CONSENT LAYER  — Purpose binding, revocation, audit receipts   │
└────────────────────────────┬────────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────────┐
│  DECISION LAYER (8 AI Agents)                                   │
│  Intent │ Qualify │ Personalise │ Converse │ Document │         │
│  Compliance │ Explainability │ Audit                            │
└────────────────────────────┬────────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────────┐
│  DATA LAYER  — PostgreSQL | Redis | Kafka | Object Store        │
└─────────────────────────────────────────────────────────────────┘
```

Full architecture documentation: [docs/architecture/](docs/architecture/)

---

## Technology Stack

| Layer | Technology |
|-------|-----------|
| Backend API | Go 1.22 + Gin |
| Frontend | Next.js 14 + TypeScript + Tailwind + shadcn/ui |
| AI Agents | Python 3.11 + FastAPI + LangChain |
| LLM | Claude Sonnet (Anthropic) |
| Database | PostgreSQL 16 |
| Cache | Redis 7 |
| Events | Kafka (mock for MVP) |
| Auth | JWT RS256 |
| Containers | Docker + Docker Compose |

---

## Repository Structure

```
aperture/
├── backend/          Go/Gin REST API (Clean Architecture)
├── frontend/         Next.js dashboard application
├── agents/           Python AI agents (8 independent services)
├── shared/           Shared types, constants, schemas
├── docs/             Architecture, API contracts, DB design
├── infra/            PostgreSQL, Redis, Kafka configs
├── docker/           Dockerfiles + docker-compose.yml
├── scripts/          Setup, seed, migration scripts
├── configs/          Environment configurations
├── assets/           Design assets
└── .github/          CI/CD workflows
```

---

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Go 1.22+
- Node.js 20+
- Python 3.11+

### Run the full stack

```bash
# Clone the repository
git clone https://github.com/swaroopthakare/aperture
cd aperture

# Copy environment config
cp .env.example .env
# Add your ANTHROPIC_API_KEY to .env

# Start all services
docker compose up -d

# Seed demo data
./scripts/seed.sh

# Open the app
open http://localhost:3000
```

### Demo credentials

| Role | Email | Password |
|------|-------|---------|
| Customer | arjun@demo.aperture | Demo@1234 |
| Relationship Manager | rm.rajesh@sbi.demo | Demo@1234 |
| Compliance Officer | compliance@sbi.demo | Demo@1234 |
| Operations | ops@sbi.demo | Demo@1234 |
| Admin | admin@aperture | Demo@1234 |

---

## AI Agents

| Agent | Purpose | Port |
|-------|---------|------|
| Intent Agent | Detects high-intent sessions from behavioral signals | 8101 |
| Qualification Agent | 5-dimension explainable product-fit scoring | 8102 |
| Personalisation Agent | Generates personalised multilingual offers | 8103 |
| Conversation Agent | Manages onboarding dialogue and field extraction | 8104 |
| Document Agent | OCR extraction from PAN, salary slips, statements | 8105 |
| Compliance Agent | Policy checks, consent verification, data minimisation | 8106 |
| Explainability Agent | Human-readable explanations for every AI decision | 8107 |
| Audit Agent | Immutable event logging and risk flagging | 8108 |

---

## Compliance

Aperture is built compliance-first:

- **Consent before data:** No data accessed without explicit purpose-bound consent
- **DPDP Act 2023:** Granular consent, revocation, data deletion
- **RBI Digital KYC:** Document extraction follows RBI Master Direction
- **Explainability:** Every AI decision has reason + evidence + policy reference
- **Audit trail:** Immutable logs retained per RBI 10-year requirement
- **Data minimisation:** Only data needed for stated purpose is collected

See [docs/architecture/system-architecture.md](docs/architecture/system-architecture.md) for the full compliance design.

---

## API Documentation

| Module | Contract |
|--------|---------|
| Authentication | [docs/api/authentication.md](docs/api/authentication.md) |
| Customers | [docs/api/customer.md](docs/api/customer.md) |
| Qualification | [docs/api/qualification.md](docs/api/qualification.md) |
| Conversation | [docs/api/chat.md](docs/api/chat.md) |
| Consent | [docs/api/consent.md](docs/api/consent.md) |
| Documents | [docs/api/document.md](docs/api/document.md) |
| Audit & Explain | [docs/api/audit.md](docs/api/audit.md) |

---

## Business Impact

| Metric | Traditional | Aperture |
|--------|------------|---------|
| Acquisition cost per lead | ₹400–₹1,200 | ₹50–₹200 |
| Qualification time | Days (manual) | Seconds (AI) |
| Onboarding completion | ~30% | ~80%+ |
| Decision explainability | None | 100% |
| KYC turnaround | 3–7 days | Minutes |

---

## Documentation

- [Product Requirements](docs/product-requirements.md)
- [User Personas](docs/user-personas.md)
- [User Flows](docs/user-flows.md)
- [Feature List](docs/feature-list.md)
- [System Architecture](docs/architecture/system-architecture.md)
- [Component Diagram](docs/architecture/component-diagram.md)
- [Sequence Diagrams](docs/architecture/sequence-diagrams.md)
- [Technology Decisions](docs/architecture/technology-decisions.md)
- [Deployment Guide](docs/architecture/deployment.md)
- [Database ERD](docs/database/ERD.md)
- [Database Tables](docs/database/tables.md)

---

## Team

**Swaroop Thakare** — AI Engineer, Arealis  
Building AI Agents & Distributed Backend Systems in Go  
PCCOE '26 | Hackathon Winner

---

## License

MIT License — see [LICENSE](LICENSE)
