// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NikolNikolaeva/project_weather/repositories (interfaces: ForecastRepo)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_forecast_repo.go . ForecastRepo
//

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"

	model "github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

// MockForecastRepo is a mock of ForecastRepo interface.
type MockForecastRepo struct {
	ctrl     *gomock.Controller
	recorder *MockForecastRepoMockRecorder
}

// MockForecastRepoMockRecorder is the mock recorder for MockForecastRepo.
type MockForecastRepoMockRecorder struct {
	mock *MockForecastRepo
}

// NewMockForecastRepo creates a new mock instance.
func NewMockForecastRepo(ctrl *gomock.Controller) *MockForecastRepo {
	mock := &MockForecastRepo{ctrl: ctrl}
	mock.recorder = &MockForecastRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockForecastRepo) EXPECT() *MockForecastRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockForecastRepo) Create(arg0 *model.Forecast) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockForecastRepoMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockForecastRepo)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockForecastRepo) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockForecastRepoMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockForecastRepo)(nil).Delete), arg0)
}

// DeleteByCityId mocks base method.
func (m *MockForecastRepo) DeleteByCityId(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByCityId", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByCityId indicates an expected call of DeleteByCityId.
func (mr *MockForecastRepoMockRecorder) DeleteByCityId(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByCityId", reflect.TypeOf((*MockForecastRepo)(nil).DeleteByCityId), arg0)
}

// FindAll mocks base method.
func (m *MockForecastRepo) FindAll() ([]*model.Forecast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*model.Forecast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockForecastRepoMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockForecastRepo)(nil).FindAll))
}

// FindByCityId mocks base method.
func (m *MockForecastRepo) FindByCityId(arg0 string) ([]*model.Forecast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCityId", arg0)
	ret0, _ := ret[0].([]*model.Forecast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCityId indicates an expected call of FindByCityId.
func (mr *MockForecastRepoMockRecorder) FindByCityId(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCityId", reflect.TypeOf((*MockForecastRepo)(nil).FindByCityId), arg0)
}

// FindByCityIdAndDate mocks base method.
func (m *MockForecastRepo) FindByCityIdAndDate(arg0 string, arg1 time.Time) (*model.Forecast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByCityIdAndDate", arg0, arg1)
	ret0, _ := ret[0].(*model.Forecast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByCityIdAndDate indicates an expected call of FindByCityIdAndDate.
func (mr *MockForecastRepoMockRecorder) FindByCityIdAndDate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByCityIdAndDate", reflect.TypeOf((*MockForecastRepo)(nil).FindByCityIdAndDate), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockForecastRepo) FindByID(arg0 string) (*model.Forecast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*model.Forecast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockForecastRepoMockRecorder) FindByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockForecastRepo)(nil).FindByID), arg0)
}

// Update mocks base method.
func (m *MockForecastRepo) Update(arg0 string, arg1 *model.Forecast) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockForecastRepoMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockForecastRepo)(nil).Update), arg0, arg1)
}
