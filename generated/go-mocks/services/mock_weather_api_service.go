// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NikolNikolaeva/project_weather/services (interfaces: WeatherService)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/services/mock_weather_api_service.go . WeatherService
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWeatherService is a mock of WeatherService interface.
type MockWeatherService struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherServiceMockRecorder
}

// MockWeatherServiceMockRecorder is the mock recorder for MockWeatherService.
type MockWeatherServiceMockRecorder struct {
	mock *MockWeatherService
}

// NewMockWeatherService creates a new mock instance.
func NewMockWeatherService(ctrl *gomock.Controller) *MockWeatherService {
	mock := &MockWeatherService{ctrl: ctrl}
	mock.recorder = &MockWeatherServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherService) EXPECT() *MockWeatherServiceMockRecorder {
	return m.recorder
}

// StartFetching mocks base method.
func (m *MockWeatherService) StartFetching() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartFetching")
}

// StartFetching indicates an expected call of StartFetching.
func (mr *MockWeatherServiceMockRecorder) StartFetching() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartFetching", reflect.TypeOf((*MockWeatherService)(nil).StartFetching))
}