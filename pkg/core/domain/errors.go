package domain

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExists       = errors.New("already exists")
	ErrCouldNotBeEmpty     = errors.New("could not be empty")
	ErrInternalApplication = errors.New("internal application error")
	ErrIsRequired          = errors.New("is required")
	ErrNotFound            = errors.New("not found")
	ErrIsNotValidated      = errors.New("is not validated")
	ErrIdCouldNotBeEmpty   = fmt.Errorf("%s %w", "id", ErrCouldNotBeEmpty)
)
