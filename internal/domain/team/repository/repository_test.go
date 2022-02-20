package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tesarwijaya/night-owl/internal/domain/team/model"
	"github.com/tesarwijaya/night-owl/internal/domain/team/repository"
)

func Test_FindAll(t *testing.T) {
	testCases := []struct {
		Name      string
		Data      *repository.Data
		Expect    []model.TeamModel
		ExpectErr error
	}{
		{
			Name: "when_data_present",
			Data: &repository.Data{
				Data: map[int64]*model.TeamModel{
					1: {
						Name: "some-team-name",
					},
				},
			},
			Expect: []model.TeamModel{{
				Name: "some-team-name",
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
			repo := repository.NewTeamReposity(test.Data)

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
		Expect    model.TeamModel
		ExpectErr error
	}{
		{
			Name: "when_data_present",
			Data: &repository.Data{
				Data: map[int64]*model.TeamModel{
					1: {
						Name: "some-team-name",
					},
				},
			},
			Param: 1,
			Expect: model.TeamModel{
				Name: "some-team-name",
			},
		},
		{
			Name: "when_data_not_present",
			Data: &repository.Data{
				Data: map[int64]*model.TeamModel{
					1: {
						Name: "some-team-name",
					},
				},
			},
			Param:     2,
			ExpectErr: errors.New("data not found"),
		},
	}

	for _, test := range testCases {
		repo := repository.NewTeamReposity(test.Data)

		actual, err := repo.FindByID(context.Background(), test.Param)
		if test.ExpectErr == nil {
			assert.Equal(t, test.Expect, actual)
			assert.Nil(t, err)
		}

		assert.Equal(t, test.ExpectErr, err)
	}
}

func Test_Insert(t *testing.T) {
	testCases := []struct {
		Name      string
		Data      *repository.Data
		Param     model.TeamModel
		Expect    model.TeamModel
		ExpectErr error
	}{
		{
			Name: "when_successful",
			Data: &repository.Data{
				Data: map[int64]*model.TeamModel{},
			},
			Param: model.TeamModel{
				Name: "some-team-name",
			},
		},
	}

	for _, test := range testCases {
		repo := repository.NewTeamReposity(test.Data)

		err := repo.Insert(context.Background(), test.Param)

		assert.Equal(t, test.ExpectErr, err)
	}
}
