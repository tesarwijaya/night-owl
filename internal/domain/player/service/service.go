package service

import (
	"context"

	"github.com/tesarwijaya/night-owl/internal/domain/player/model"
	"github.com/tesarwijaya/night-owl/internal/domain/player/repository"
	team_repository "github.com/tesarwijaya/night-owl/internal/domain/team/repository"
	"go.uber.org/dig"
)

type PlayerService interface {
	FindAll(ctx context.Context) ([]model.PlayerModel, error)
	FindByID(ctx context.Context, id int64) (model.PlayerModel, error)
	Insert(ctx context.Context, payload model.PlayerModel) (model.PlayerModel, error)
}

type PlayerServiceImpl struct {
	dig.In
	Repo     repository.PlayerRepository
	TeamRepo team_repository.TeamRepository
}

func NewPlayerService(svc PlayerServiceImpl) PlayerService {
	return &svc
}

func (s *PlayerServiceImpl) FindAll(ctx context.Context) ([]model.PlayerModel, error) {
	return s.Repo.FindAll(ctx)
}

func (s *PlayerServiceImpl) FindByID(ctx context.Context, id int64) (model.PlayerModel, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *PlayerServiceImpl) Insert(ctx context.Context, payload model.PlayerModel) (model.PlayerModel, error) {
	_, err := s.TeamRepo.FindByID(ctx, payload.TeamID)
	if err != nil {
		return model.PlayerModel{}, err
	}

	if err := s.Repo.Insert(ctx, payload); err != nil {
		return model.PlayerModel{}, err
	}

	return payload, nil
}
