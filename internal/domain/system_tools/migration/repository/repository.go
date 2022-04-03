package repository

import (
	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/dig"
)

type MigrationRepository interface {
	Up() error
	Down() error
}

type MigrationRepositoryImpl struct {
	dig.In
	Migrate *migrate.Migrate
}

func NewMigrationRepository(repo MigrationRepositoryImpl) MigrationRepository {
	return &repo
}

func (r MigrationRepositoryImpl) Up() error {
	return r.Migrate.Up()
}

func (r MigrationRepositoryImpl) Down() error {
	return r.Migrate.Up()
}
