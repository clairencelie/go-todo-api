package helper

import "errors"

var (
	ErrLoginFailed        = errors.New("invalid username or password")
	ErrNotFound           = errors.New("data not found")
	ErrRowsNotAffected    = errors.New("no rows affected")
	ErrorTokenInvalid     = errors.New("token invalid")
	ErrBearerTokenMissing = errors.New("bearer token missing")
)
