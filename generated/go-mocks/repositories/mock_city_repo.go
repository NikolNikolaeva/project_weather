// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NikolNikolaeva/project_weather/repositories (interfaces: CityRepo)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_city_repo.go . CityRepo
//

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gen "gorm.io/gen"

	model "github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

// MockCityRepo is a mock of CityRepo interface.
type MockCityRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCityRepoMockRecorder
}

// MockCityRepoMockRecorder is the mock recorder for MockCityRepo.
type MockCityRepoMockRecorder struct {
	mock *MockCityRepo
}

// NewMockCityRepo creates a new mock instance.
func NewMockCityRepo(ctrl *gomock.Controller) *MockCityRepo {
	mock := &MockCityRepo{ctrl: ctrl}
	mock.recorder = &MockCityRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCityRepo) EXPECT() *MockCityRepoMockRecorder {
	return m.recorder
}

// DeleteByID mocks base method.
func (m *MockCityRepo) DeleteByID(arg0 string) (gen.ResultInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(gen.ResultInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockCityRepoMockRecorder) DeleteByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockCityRepo)(nil).DeleteByID), arg0)
}

// FindByID mocks base method.
func (m *MockCityRepo) FindByID(arg0 string) (*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCityRepoMockRecorder) FindByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCityRepo)(nil).FindByID), arg0)
}

// GetAll mocks base method.
func (m *MockCityRepo) GetAll() ([]*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCityRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCityRepo)(nil).GetAll))
}

// Register mocks base method.
func (m *MockCityRepo) Register(arg0 *model.City) (*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockCityRepoMockRecorder) Register(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockCityRepo)(nil).Register), arg0)
}

// UpdateByID mocks base method.
func (m *MockCityRepo) UpdateByID(arg0 string, arg1 *model.City) (*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", arg0, arg1)
	ret0, _ := ret[0].(*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockCityRepoMockRecorder) UpdateByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockCityRepo)(nil).UpdateByID), arg0, arg1)
}
