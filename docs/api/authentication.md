# API Contract — Authentication

**Base URL:** `/api/v1/auth`  
**Version:** v1

---

## POST /api/v1/auth/login

Login and receive JWT tokens.

**Request:**
```json
{
  "email": "string (required)",
  "password": "string (required)"
}
```

**Response 200:**
```json
{
  "access_token": "eyJhbGci...",
  "token_type": "Bearer",
  "expires_in": 900,
  "user": {
    "id": "uuid",
    "email": "string",
    "name": "string",
    "role": "customer | rm | compliance | operations | admin",
    "avatar_url": "string | null"
  }
}
```
Refresh token returned in HttpOnly cookie: `aperture_refresh`

**Response 401:** `{"error": "invalid_credentials"}`  
**Response 429:** `{"error": "too_many_attempts", "retry_after": 60}`

---

## POST /api/v1/auth/refresh

Exchange refresh token for new access token.

**Request:** (no body — reads `aperture_refresh` cookie)

**Response 200:**
```json
{
  "access_token": "eyJhbGci...",
  "expires_in": 900
}
```

---

## POST /api/v1/auth/logout

Revoke refresh token.

**Headers:** `Authorization: Bearer {token}`

**Response 200:** `{"message": "logged_out"}`

---

## GET /api/v1/auth/me

Get current user profile.

**Headers:** `Authorization: Bearer {token}`

**Response 200:**
```json
{
  "id": "uuid",
  "email": "string",
  "name": "string",
  "role": "string",
  "permissions": ["string"],
  "created_at": "ISO8601"
}
```

---

## POST /api/v1/auth/change-password

**Headers:** `Authorization: Bearer {token}`

**Request:**
```json
{
  "current_password": "string",
  "new_password": "string (min 12 chars)"
}
```

**Response 200:** `{"message": "password_changed"}`
