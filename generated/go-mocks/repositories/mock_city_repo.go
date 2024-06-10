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

// DeleteCityByID mocks base method.
func (m *MockCityRepo) DeleteCityByID(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCityByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCityByID indicates an expected call of DeleteCityByID.
func (mr *MockCityRepoMockRecorder) DeleteCityByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCityByID", reflect.TypeOf((*MockCityRepo)(nil).DeleteCityByID), arg0)
}

// FindCityByID mocks base method.
func (m *MockCityRepo) FindCityByID(arg0 string) (*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCityByID", arg0)
	ret0, _ := ret[0].(*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCityByID indicates an expected call of FindCityByID.
func (mr *MockCityRepoMockRecorder) FindCityByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCityByID", reflect.TypeOf((*MockCityRepo)(nil).FindCityByID), arg0)
}

// GetAllCity mocks base method.
func (m *MockCityRepo) GetAllCity() ([]*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCity")
	ret0, _ := ret[0].([]*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCity indicates an expected call of GetAllCity.
func (mr *MockCityRepoMockRecorder) GetAllCity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCity", reflect.TypeOf((*MockCityRepo)(nil).GetAllCity))
}

// RegisterCity mocks base method.
func (m *MockCityRepo) RegisterCity(arg0 *model.City) (*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCity", arg0)
	ret0, _ := ret[0].(*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterCity indicates an expected call of RegisterCity.
func (mr *MockCityRepoMockRecorder) RegisterCity(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCity", reflect.TypeOf((*MockCityRepo)(nil).RegisterCity), arg0)
}

// UpdateCityByID mocks base method.
func (m *MockCityRepo) UpdateCityByID(arg0 string, arg1 *model.City) (*model.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCityByID", arg0, arg1)
	ret0, _ := ret[0].(*model.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCityByID indicates an expected call of UpdateCityByID.
func (mr *MockCityRepoMockRecorder) UpdateCityByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCityByID", reflect.TypeOf((*MockCityRepo)(nil).UpdateCityByID), arg0, arg1)
}
