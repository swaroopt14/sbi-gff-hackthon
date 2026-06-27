# Aperture — Deployment Architecture

**Version:** 1.0  
**Date:** 2026-06-27

---

## MVP Deployment (Local / Demo)

```
docker compose up

┌──────────────────────────────────────────────────────────────────┐
│ Docker Compose Network: aperture_default                        │
│                                                                  │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────────┐   │
│  │  frontend   │     │   backend   │     │  agent-gateway  │   │
│  │  :3000      │────▶│   :8080     │────▶│   :8100         │   │
│  │  Next.js    │     │   Go/Gin    │     │   FastAPI proxy │   │
│  └─────────────┘     └──────┬──────┘     └────────┬────────┘   │
│                             │                     │             │
│             ┌───────────────┼─────────────────────┤             │
│             │               │                     │             │
│       ┌─────▼─────┐   ┌─────▼─────┐   ┌──────────▼──────┐    │
│       │ postgres  │   │   redis   │   │  agent services │    │
│       │  :5432    │   │   :6379   │   │  :8101-:8108    │    │
│       └───────────┘   └───────────┘   └─────────────────┘    │
│                                                                  │
│  ┌─────────────┐     ┌─────────────┐                          │
│  │  kafka      │     │  zookeeper  │                          │
│  │  :9092      │     │  :2181      │  (mock mode for MVP)    │
│  └─────────────┘     └─────────────┘                          │
│                                                                  │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │  minio (S3-compatible object store)  :9000              │   │
│  │  For document storage during demo                       │   │
│  └─────────────────────────────────────────────────────────┘   │
└──────────────────────────────────────────────────────────────────┘
```

---

## Service Port Map

| Service | Port | Description |
|---------|------|-------------|
| Frontend | 3000 | Next.js dev server |
| Backend | 8080 | Go/Gin API |
| Agent Gateway | 8100 | Routes to individual agents |
| Intent Agent | 8101 | Intent scoring |
| Qualification Agent | 8102 | Qualification scoring |
| Personalisation Agent | 8103 | Message generation |
| Conversation Agent | 8104 | Chat management |
| Document Agent | 8105 | OCR + extraction |
| Compliance Agent | 8106 | Policy checks |
| Explainability Agent | 8107 | Explanation generation |
| Audit Agent | 8108 | Event logging |
| PostgreSQL | 5432 | Primary database |
| Redis | 6379 | Cache |
| Kafka | 9092 | Events (mock) |
| MinIO | 9000 | Object store |
| MinIO Console | 9001 | MinIO UI |

---

## Production Deployment (Post-Hackathon Target)

```
┌─────────────────────────────────────────────────────────────────────────┐
│                        AWS / Azure Cloud                                │
│                                                                         │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │  CloudFront / Azure CDN                                         │   │
│  │  (Static assets + Next.js edge)                                 │   │
│  └───────────────────────────┬─────────────────────────────────────┘   │
│                              │                                          │
│  ┌───────────────────────────▼─────────────────────────────────────┐   │
│  │  Application Load Balancer (ALB)                                │   │
│  │  SSL termination | Rate limiting | WAF                          │   │
│  └──────────────────┬────────────────────┬───────────────────────┘   │
│                     │                    │                             │
│  ┌──────────────────▼───┐   ┌────────────▼──────────────────────┐    │
│  │  EKS / AKS           │   │  EKS / AKS                        │    │
│  │  Go Backend Pods (n) │   │  Python Agent Pods (n)            │    │
│  │  HPA: CPU+RPS based  │   │  HPA: queue depth based           │    │
│  └──────────────────────┘   └───────────────────────────────────┘    │
│                                                                         │
│  ┌─────────────┐  ┌──────────────┐  ┌──────────────┐  ┌───────────┐  │
│  │  RDS Aurora │  │ ElastiCache  │  │  MSK Kafka   │  │    S3     │  │
│  │  PostgreSQL │  │  Redis       │  │  (managed)   │  │ Documents │  │
│  │  Multi-AZ   │  │  Cluster     │  │              │  │           │  │
│  └─────────────┘  └──────────────┘  └──────────────┘  └───────────┘  │
│                                                                         │
│  ┌─────────────────────────────────────────────────────────────────┐   │
│  │  Observability: Prometheus + Grafana + ELK Stack               │   │
│  └─────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## Environment Strategy

| Environment | Purpose | Infra |
|-------------|---------|-------|
| local | Developer workstation | Docker Compose |
| development | CI/CD integration testing | Docker Compose on CI |
| staging | Pre-production validation | Kubernetes (reduced replicas) |
| production | Live system | Kubernetes (full HA) |

---

## CI/CD Pipeline (GitHub Actions)

```
Push to branch
      │
      ▼
[Lint + Test]
  go test ./...
  pytest agents/
  tsc --noEmit
      │
      ▼
[Build Docker Images]
  backend:sha
  frontend:sha
  agents:sha
      │
      ▼
[Push to Registry]
  ghcr.io/aperture/backend:sha
      │
      ▼
[Deploy to Staging]
  helm upgrade aperture ./infra/helm
      │
      ▼
[Integration Tests]
  API contract tests
  E2E smoke tests
      │
      ▼
[Manual Gate: Production Deploy]
```

---

## Secret Management

| Secret | Storage | Rotation |
|--------|---------|----------|
| ANTHROPIC_API_KEY | HashiCorp Vault / AWS Secrets Manager | 90 days |
| DATABASE_URL | HashiCorp Vault | On rotation |
| JWT_PRIVATE_KEY | HashiCorp Vault | 90 days |
| REDIS_PASSWORD | HashiCorp Vault | 90 days |
| KAFKA credentials | HashiCorp Vault | 90 days |

**MVP:** `.env` files with `.env.example` template. Never committed to git.

---

## Health Check Strategy

Every service exposes:

```
GET /health      → 200 {"status":"ok","version":"1.0.0"}
GET /ready       → 200 {"status":"ready"} or 503 (during startup)
GET /metrics     → Prometheus metrics
```

Docker Compose healthchecks ensure services wait for dependencies before accepting traffic.
