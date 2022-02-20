package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/tesarwijaya/night-owl/internal/domain/player/model"
)

var i int64

type Data struct {
	mu   sync.Mutex
	Data map[int64]*model.PlayerModel
}

func NewPlayerData() *Data {
	return &Data{
		Data: map[int64]*model.PlayerModel{},
	}
}

type PlayerRepository interface {
	FindAll(ctx context.Context) ([]model.PlayerModel, error)
	FindByID(ctx context.Context, id int64) (model.PlayerModel, error)
	FindByTeamID(ctx context.Context, teamID int64) ([]model.PlayerModel, error)
	Insert(ctx context.Context, payload model.PlayerModel) error
}

type PlayerRepositoryImpl struct {
	Data *Data
}

func NewPlayerReposity(data *Data) PlayerRepository {
	return &PlayerRepositoryImpl{
		Data: data,
	}
}

func (r *PlayerRepositoryImpl) FindAll(ctx context.Context) ([]model.PlayerModel, error) {
	var res []model.PlayerModel
	for _, v := range r.Data.Data {
		res = append(res, *v)
	}

	if len(res) == 0 {
		return nil, errors.New("data not found")
	}

	return res, nil
}

func (r *PlayerRepositoryImpl) FindByID(ctx context.Context, id int64) (model.PlayerModel, error) {
	var res *model.PlayerModel
	for k, v := range r.Data.Data {
		if k == id {
			res = v
		}
	}

	if res == nil {
		return model.PlayerModel{}, errors.New("data not found")
	}

	return *res, nil
}

func (r *PlayerRepositoryImpl) FindByTeamID(ctx context.Context, teamID int64) ([]model.PlayerModel, error) {
	var res []model.PlayerModel
	for _, v := range r.Data.Data {
		if v.TeamID == teamID {
			res = append(res, *v)
		}
	}

	if len(res) == 0 {
		return nil, errors.New("data not found")
	}

	return res, nil
}

func (r *PlayerRepositoryImpl) Insert(ctx context.Context, payload model.PlayerModel) error {
	r.Data.mu.Lock()
	defer r.Data.mu.Unlock()

	i++
	r.Data.Data[i] = &payload
	return nil
}
