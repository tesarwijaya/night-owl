package service

import (
	"github.com/tesarwijaya/night-owl/internal/domain/system_tools/migration/repository"
	"go.uber.org/dig"
)

type MigrationService interface {
	Up() error
	Down() error
}

type MigrationServiceImpl struct {
	dig.In
	R repository.MigrationRepository
}

func NewMigrationService(svc MigrationServiceImpl) MigrationService {
	return &svc
}

func (s *MigrationServiceImpl) Up() error {
	return s.R.Up()
}

func (s *MigrationServiceImpl) Down() error {
	return s.R.Down()
}
