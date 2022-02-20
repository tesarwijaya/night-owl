package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tesarwijaya/night-owl/internal/domain/player/model"
	"github.com/tesarwijaya/night-owl/internal/domain/player/repository"
)

func Test_FindAll(t *testing.T) {
	testCases := []struct {
		Name      string
		Data      *repository.Data
		Expect    []model.PlayerModel
		ExpectErr error
	}{
		{
			Name: "when_data_present",
			Data: &repository.Data{
				Data: map[int64]*model.PlayerModel{
					1: {
						Name: "some-player-name",
					},
				},
			},
			Expect: []model.PlayerModel{{
				Name: "some-player-name",
			}},
		},
		{
			Name:      "when_data_not_found",
			Data:      &repository.Data{},
			ExpectErr: errors.New("data not found"),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			repo := repository.NewPlayerReposity(test.Data)

			actual, err := repo.FindAll(context.Background())
			if test.ExpectErr == nil {
				assert.Equal(t, test.Expect, actual)
				assert.Nil(t, err)
			}

			assert.Equal(t, test.ExpectErr, err)
		})
	}
}

func Test_FindByID(t *testing.T) {
	testCases := []struct {
		Name      string
		Data      *repository.Data
		Param     int64
		Expect    model.PlayerModel
		ExpectErr error
	}{
		{
			Name: "when_data_present",
			Data: &repository.Data{
				Data: map[int64]*model.PlayerModel{
					1: {
						Name: "some-player-name",
					},
				},
			},
			Param: 1,
			Expect: model.PlayerModel{
				Name: "some-player-name",
			},
		},
		{
			Name: "when_data_not_present",
			Data: &repository.Data{
				Data: map[int64]*model.PlayerModel{
					1: {
						Name: "some-player-name",
					},
				},
			},
			Param:     2,
			ExpectErr: errors.New("data not found"),
		},
	}

	for _, test := range testCases {
		repo := repository.NewPlayerReposity(test.Data)

		actual, err := repo.FindByID(context.Background(), test.Param)
		if test.ExpectErr == nil {
			assert.Equal(t, test.Expect, actual)
			assert.Nil(t, err)
		}

		assert.Equal(t, test.ExpectErr, err)
	}
}

func Test_FindByTeamID(t *testing.T) {
	testCases := []struct {
		Name      string
		Param     int64
		Data      *repository.Data
		Expect    []model.PlayerModel
		ExpectErr error
	}{
		{
			Name:  "when_data_present",
			Param: 1,
			Data: &repository.Data{
				Data: map[int64]*model.PlayerModel{
					1: {
						Name:   "some-player-name",
						TeamID: 1,
					},
				},
			},
			Expect: []model.PlayerModel{{
				Name:   "some-player-name",
				TeamID: 1,
			}},
		},
		{
			Name:  "when_data_not_found",
			Param: 2,
			Data: &repository.Data{
				Data: map[int64]*model.PlayerModel{
					1: {
						Name:   "some-player-name",
						TeamID: 1,
					},
				},
			},
			ExpectErr: errors.New("data not found"),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			repo := repository.NewPlayerReposity(test.Data)

			actual, err := repo.FindByTeamID(context.Background(), test.Param)
			if test.ExpectErr == nil {
				assert.Equal(t, test.Expect, actual)
				assert.Nil(t, err)
			}

			assert.Equal(t, test.ExpectErr, err)
		})
	}
}

func Test_Insert(t *testing.T) {
	testCases := []struct {
		Name      string
		Data      *repository.Data
		Param     model.PlayerModel
		Expect    model.PlayerModel
		ExpectErr error
	}{
		{
			Name: "when_successful",
			Data: &repository.Data{
				Data: map[int64]*model.PlayerModel{},
			},
			Param: model.PlayerModel{
				Name: "some-player-name",
			},
		},
	}

	for _, test := range testCases {
		repo := repository.NewPlayerReposity(test.Data)

		err := repo.Insert(context.Background(), test.Param)

		assert.Equal(t, test.ExpectErr, err)
	}
}
