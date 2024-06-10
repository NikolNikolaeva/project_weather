// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NikolNikolaeva/project_weather/services (interfaces: WeatherDataGetter)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/services/mock_weather_data_getter.go . WeatherDataGetter
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	weatherapi "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
)

// MockWeatherDataGetter is a mock of WeatherDataGetter interface.
type MockWeatherDataGetter struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherDataGetterMockRecorder
}

// MockWeatherDataGetterMockRecorder is the mock recorder for MockWeatherDataGetter.
type MockWeatherDataGetterMockRecorder struct {
	mock *MockWeatherDataGetter
}

// NewMockWeatherDataGetter creates a new mock instance.
func NewMockWeatherDataGetter(ctrl *gomock.Controller) *MockWeatherDataGetter {
	mock := &MockWeatherDataGetter{ctrl: ctrl}
	mock.recorder = &MockWeatherDataGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherDataGetter) EXPECT() *MockWeatherDataGetterMockRecorder {
	return m.recorder
}

// GetCurrentData mocks base method.
func (m *MockWeatherDataGetter) GetCurrentData(arg0, arg1 string) (*weatherapi.Current, *weatherapi.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentData", arg0, arg1)
	ret0, _ := ret[0].(*weatherapi.Current)
	ret1, _ := ret[1].(*weatherapi.Location)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCurrentData indicates an expected call of GetCurrentData.
func (mr *MockWeatherDataGetterMockRecorder) GetCurrentData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentData", reflect.TypeOf((*MockWeatherDataGetter)(nil).GetCurrentData), arg0, arg1)
}

// GetForecastData mocks base method.
func (m *MockWeatherDataGetter) GetForecastData(arg0 string, arg1 int32, arg2 string) (*weatherapi.Forecast, *weatherapi.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForecastData", arg0, arg1, arg2)
	ret0, _ := ret[0].(*weatherapi.Forecast)
	ret1, _ := ret[1].(*weatherapi.Location)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetForecastData indicates an expected call of GetForecastData.
func (mr *MockWeatherDataGetterMockRecorder) GetForecastData(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForecastData", reflect.TypeOf((*MockWeatherDataGetter)(nil).GetForecastData), arg0, arg1, arg2)
}
