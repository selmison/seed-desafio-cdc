package domain

import (
	"fmt"

	"github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

var (
	ErrEmailIsNotValidated = fmt.Errorf("%s %w", "e-mail", domain.ErrIsNotValidated)
	ErrDescIsNotValidated  = fmt.Errorf("%s %w", "description", domain.ErrIsNotValidated)
	ErrTimeIsNotValidated  = fmt.Errorf("%s %w", "time", domain.ErrIsNotValidated)
	ErrNameCouldNotBeEmpty = fmt.Errorf("%s %w", "name", domain.ErrCouldNotBeEmpty)
)
