package models

import "errors"

var (
	InvalidTokenError    = errors.New("invalid token")
	AccessDeniedError    = errors.New("access denied")
	ValidationTokenError = errors.New("validation token error: %s")
	JwkKetNotFound       = errors.New("jwk key not found")
)
