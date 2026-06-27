# Aperture Backend

Go 1.22 + Gin REST API — Clean Architecture, Domain-Driven Design

## Structure

```
backend/
├── cmd/aperture/        main.go entry point
├── internal/
│   ├── domain/          Entities, value objects, domain errors
│   ├── application/     Use cases, orchestration logic
│   ├── repository/      Data access interfaces + implementations
│   ├── service/         Business services (intent, qualify, consent…)
│   ├── handler/         HTTP handlers (Gin route handlers)
│   ├── middleware/       Auth, RBAC, rate limit, logger, request ID
│   ├── config/          Viper config loading
│   ├── logger/          Zap structured logger setup
│   └── utils/           Shared helpers (crypto, pagination, errors)
├── pkg/
│   ├── llm/             LLM client abstraction (Claude / OpenAI)
│   ├── eventbus/        Kafka / mock event bus interface
│   ├── cache/           Redis client wrapper
│   └── storage/         Object store client (MinIO / S3)
├── api/v1/              OpenAPI/Swagger spec
├── configs/             YAML config files (app.yaml, db.yaml…)
├── migrations/          PostgreSQL migration files (golang-migrate)
├── test/                Integration and API tests
└── docs/                Swagger generated docs
```

## Running locally

```bash
cd backend
cp ../configs/development/.env .env
go run ./cmd/aperture
# Server starts on :8080
```

## Testing

```bash
go test ./...
go test ./... -race -cover
```
