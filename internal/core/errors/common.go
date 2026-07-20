package core_errors

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrInvalideArgument = errors.New("invalid argument")
)
