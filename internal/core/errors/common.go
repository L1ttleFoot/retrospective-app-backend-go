package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("not_found")
	ErrInvalidArgument = errors.New("invalid_argument")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrInternal        = errors.New("internal_server_error")
)
