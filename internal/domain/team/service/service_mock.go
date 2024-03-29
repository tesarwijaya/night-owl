// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/team/service/service.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tesarwijaya/night-owl/internal/domain/team/model"
)

// MockTeamService is a mock of TeamService interface.
type MockTeamService struct {
	ctrl     *gomock.Controller
	recorder *MockTeamServiceMockRecorder
}

// MockTeamServiceMockRecorder is the mock recorder for MockTeamService.
type MockTeamServiceMockRecorder struct {
	mock *MockTeamService
}

// NewMockTeamService creates a new mock instance.
func NewMockTeamService(ctrl *gomock.Controller) *MockTeamService {
	mock := &MockTeamService{ctrl: ctrl}
	mock.recorder = &MockTeamServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamService) EXPECT() *MockTeamServiceMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockTeamService) FindAll(ctx context.Context) ([]model.TeamModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]model.TeamModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTeamServiceMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTeamService)(nil).FindAll), ctx)
}

// FindByID mocks base method.
func (m *MockTeamService) FindByID(ctx context.Context, id int64) (model.TeamModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(model.TeamModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockTeamServiceMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockTeamService)(nil).FindByID), ctx, id)
}

// FindTeamPlayer mocks base method.
func (m *MockTeamService) FindTeamPlayer(ctx context.Context, id int64) (model.TeamPlayerRespModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTeamPlayer", ctx, id)
	ret0, _ := ret[0].(model.TeamPlayerRespModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTeamPlayer indicates an expected call of FindTeamPlayer.
func (mr *MockTeamServiceMockRecorder) FindTeamPlayer(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTeamPlayer", reflect.TypeOf((*MockTeamService)(nil).FindTeamPlayer), ctx, id)
}

// Insert mocks base method.
func (m *MockTeamService) Insert(ctx context.Context, payload model.TeamModel) (model.TeamModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, payload)
	ret0, _ := ret[0].(model.TeamModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTeamServiceMockRecorder) Insert(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTeamService)(nil).Insert), ctx, payload)
}
