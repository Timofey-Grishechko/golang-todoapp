package core_postgres_pool

import "errors"

var (
	ErrNoRows              = errors.New("no rows")
	ErrViolatesForeighnKey = errors.New("violates foreighn key")
	ErrUnknown             = errors.New("unknown")
)
