package catalog

import (
	kitLog "github.com/go-kit/kit/log"
	"gorm.io/gorm"

	"github.com/selmison/seed-desafio-cdc/gen/catalog"
)

type service struct {
	repo   *gorm.DB
	logger kitLog.Logger
}

// NewService returns a new Actors Service.
func NewService(repo *gorm.DB, logger kitLog.Logger) catalog.Service {
	return &service{repo, logger}
}
