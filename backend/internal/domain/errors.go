package domain

import "errors"

var (
	ErrNotFound          = errors.New("record not found")
	ErrConflict          = errors.New("record already exists")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrValidation        = errors.New("validation error")
	ErrConsentMissing    = errors.New("valid consent not found for this purpose")
	ErrConsentRevoked    = errors.New("consent has been revoked")
	ErrConsentExpired    = errors.New("consent has expired")
	ErrPurposeMismatch   = errors.New("data purpose does not match granted consent")
	ErrDocumentNotReady  = errors.New("document processing not complete")
	ErrAgentUnavailable  = errors.New("agent service unavailable")
	ErrRateLimited       = errors.New("rate limit exceeded")
)

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error { return e.Err }

func NewAppError(code, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}
