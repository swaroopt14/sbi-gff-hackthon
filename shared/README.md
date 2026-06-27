# Aperture Shared

Shared constants, types, schemas, and utilities used across backend, frontend, and agents.

## Structure

```
shared/
├── constants/    Domain constants (event types, status codes, product codes, policy codes)
├── types/        Shared TypeScript types (frontend) / Go structs (backend reuses domain)
├── models/       Shared data models (JSON schema definitions)
├── utils/        Pure utility functions (no framework dependencies)
└── schemas/      JSON Schema / Zod schemas for API contract validation
```

## Usage

- **Frontend** imports from `@aperture/shared`
- **Backend** Go packages re-declare types inline (Go modules are self-contained)
- **Agents** import shared schemas from `shared/schemas/` for input/output validation
