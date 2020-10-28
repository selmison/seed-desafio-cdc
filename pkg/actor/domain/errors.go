package domain

import (
	"fmt"

	"github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

var (
	ErrEmailIsNotValidated  = fmt.Errorf("%s %w", "e-mail", domain.ErrIsNotValidated)
	ErrDescIsNotValidated   = fmt.Errorf("%s %w", "description", domain.ErrIsNotValidated)
	ErrEmailCouldNotBeEmpty = fmt.Errorf("%s %w", "e-mail", domain.ErrCouldNotBeEmpty)
	ErrNameCouldNotBeEmpty  = fmt.Errorf("%s %w", "name", domain.ErrCouldNotBeEmpty)
)
