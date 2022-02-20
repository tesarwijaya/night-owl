package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/tesarwijaya/night-owl/internal/domain/team/model"
)

var i int64 = 0

type Data struct {
	mu   sync.Mutex
	Data map[int64]*model.TeamModel
}

func NewTeamData() *Data {
	return &Data{
		Data: map[int64]*model.TeamModel{},
	}
}

type TeamRepository interface {
	FindAll(ctx context.Context) ([]model.TeamModel, error)
	FindByID(ctx context.Context, id int64) (model.TeamModel, error)
	Insert(ctx context.Context, payload model.TeamModel) error
}

type TeamRepositoryImpl struct {
	Data *Data
}

func NewTeamReposity(data *Data) TeamRepository {
	return &TeamRepositoryImpl{
		Data: data,
	}
}

func (r *TeamRepositoryImpl) FindAll(ctx context.Context) ([]model.TeamModel, error) {
	var res []model.TeamModel
	for _, v := range r.Data.Data {
		res = append(res, *v)
	}

	if len(res) == 0 {
		return nil, errors.New("data not found")
	}

	return res, nil
}

func (r *TeamRepositoryImpl) FindByID(ctx context.Context, id int64) (model.TeamModel, error) {
	var res *model.TeamModel
	for k, v := range r.Data.Data {
		if k == id {
			res = v
		}
	}

	if res == nil {
		return model.TeamModel{}, errors.New("data not found")
	}

	return *res, nil
}

func (r *TeamRepositoryImpl) Insert(ctx context.Context, payload model.TeamModel) error {
	r.Data.mu.Lock()
	defer r.Data.mu.Unlock()

	i++
	payload.ID = i
	r.Data.Data[i] = &payload
	return nil
}
