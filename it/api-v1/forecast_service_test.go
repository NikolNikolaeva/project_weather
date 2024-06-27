//go:build integration

package api_v1

//
//import (
//	"context"
//	"net/http"
//	"testing"
//	"time"
//
//	. "github.com/onsi/gomega"
//
//	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
//	"github.com/NikolNikolaeva/project_weather/it/testbed"
//	"github.com/NikolNikolaeva/project_weather/services"
//)
//
//// GET: /cities/{id}/forecasts -> OK (No forecasts)
//func Test_givenNoForecastsExist_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//		city := tb.CreateCity(testbed.NewTestCity())
//
//		// create service
//		service := services.NewForecastAPIService(tb.ForecastRepo, tb.Converter, tb.Handler, tb.Config, tb.CityRepo)
//
//		// execute
//		result, err := service.GetByCityIdAndPeriod(context.Background(), city.ID, "")
//
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(result.Code).To(Equal(http.StatusOK))
//		Expect(result.Body).To(Equal([]api.Forecast{}))
//	})
//}
//
//// GET: /cities/{id}/forecasts -> OK (1 forecast)
//func Test_givenForecastExists_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//		city := tb.CreateCity(testbed.NewTestCity())
//		forecast := tb.CreateForecast(testbed.NewTestForecast())
//
//		// create service
//		service := services.NewForecastAPIService(tb.ForecastRepo, tb.Converter, tb.Handler, tb.Config, tb.CityRepo)
//
//		// execute
//		result, err := service.GetByCityIdAndPeriod(context.Background(), city.ID, "")
//
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(result.Code).To(Equal(http.StatusOK))
//		Expect(result.Body).To(Equal([]api.Forecast{*tb.Converter.ConvertModelForecastToApiForecast(forecast)}))
//	})
//}
//
//// GET: /cities/{id}/forecasts -> BadRequest (Empty cityId)
//func Test_givenAnEmptyCityId_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsBadRequest(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// create service
//		service := services.NewForecastAPIService(tb.ForecastRepo, tb.Converter, tb.Handler, tb.Config, tb.CityRepo)
//
//		// execute
//		result, err := service.GetByCityIdAndPeriod(context.Background(), "", "")
//
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(result.Code).To(Equal(http.StatusNotImplemented))
//		Expect(result.Body).To(Equal("City id is required"))
//	})
//}
//
//// GET: /cities/{id}/forecasts -> NotFound (Invalid cityId)
//func Test_givenAnInvalidCityId_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsNotFound(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// create service
//		service := services.NewForecastAPIService(tb.ForecastRepo, tb.Converter, tb.Handler, tb.Config, tb.CityRepo)
//
//		// execute
//		result, err := service.GetByCityIdAndPeriod(context.Background(), "invalidId", "daily")
//
//		// assert
//		Expect(err).Should(HaveOccurred())
//		Expect(result.Code).To(Equal(http.StatusNotImplemented))
//		Expect(result.Body).To(Equal("City not found"))
//	})
//}
//
//// GET: /cities/{id}/forecasts -> OK (Current forecast)
//func Test_givenValidCityIdAndPeriodCurrent_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//		city := tb.CreateCity(testbed.NewTestCity())
//		currentData := tb.CreateCurrentWeather("Sofia")
//		tb.FakeWeatherApi.RespondOn(
//			testbed.CreateWeatherDataRequestMatcher(city.Name),
//			testbed.CreateWeatherDataResponder(currentData),
//		)
//
//		// create service
//		service := services.NewForecastAPIService(tb.ForecastRepo, tb.Converter, tb.Handler, tb.Config, tb.CityRepo)
//
//		// execute
//		result, err := service.GetByCityIdAndPeriod(context.Background(), city.ID, "current")
//
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(result.Code).To(Equal(http.StatusOK))
//		Expect(result.Body).To(Equal(currentData))
//	})
//}
//
//// GET: /cities/{id}/forecasts -> BadRequest (Invalid period)
//func Test_givenInvalidPeriod_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsBadRequest(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//		city := tb.CreateCity(testbed.NewTestCity())
//
//		// create service
//		service := services.NewForecastAPIService(tb.ForecastRepo, tb.Converter, tb.Handler, tb.Config, tb.CityRepo)
//
//		// execute
//		result, err := service.GetByCityIdAndPeriod(context.Background(), city.ID, "invalid-period")
//
//		// assert
//		Expect(err).Should(HaveOccurred())
//		Expect(result.Code).To(Equal(http.StatusNotImplemented))
//		Expect(result.Body).To(Equal("Invalid period. Please specify 'current', 'daily', 'weekly', or 'monthly'."))
//	})
//}
//
//// GET: /cities/{id}/forecasts -> OK (Daily forecasts)
//func Test_givenValidCityIdAndPeriodDaily_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//		city := tb.CreateCity(testbed.NewTestCity())
//		forecasts := tb.CreateForecastWeather(city.ID, 1)
//
//		// create service
//		service := services.NewForecastAPIService(tb.ForecastRepo, tb.Converter, tb.Handler, tb.Config, tb.CityRepo)
//
//		// execute
//		result, err := service.GetByCityIdAndPeriod(context.Background(), city.ID, "daily")
//
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(result.Code).To(Equal(http.StatusOK))
//		Expect(result.Body).To(Equal(forecasts))
//	})
//}
