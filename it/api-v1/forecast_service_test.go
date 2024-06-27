////go:build integration

package api_v1

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/NikolNikolaeva/project_weather/it/testbed"
)

// GET: /cities/{id}/forecasts -> OK (No forecasts)
func Test_givenNoForecastsExist_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// prepare
		city := tb.CreateCity(testbed.NewTestCity())

		// execute
		result, res, err := tb.APIClient.ForecastAPI.GetByCityIdAndPeriod(context.Background(), city.ID).Period("").Execute()

		// assert
		Expect(err).ShouldNot(HaveOccurred())
		Expect(res).To(HaveHTTPStatus(http.StatusOK))
		Expect(result).To(BeNil())
	})
}

//
//// GET: /cities/{id}/forecasts -> OK (1 forecast)
//func Test_givenForecastExists_whenGetByCityIdAndCurrentPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//		city := tb.CreateCity(testbed.NewTestCity())
//
//		//currentData := tb.CreateCurrentWeather(city.Name)
//		//tb.FakeWeatherApi.RespondOn(
//		//	testbed.CreateWeatherDataRequestMatcher(city.Name, "current"),
//		//	testbed.CreateWeatherDataResponder(currentData),
//		//)
//		//fmt.Println(currentData)
//		// execute
//		result, res, err := tb.APIClient.ForecastAPI.GetByCityIdAndPeriod(context.Background(), city.ID).Period("current").Execute()
//
//		fmt.Println(result)
//		fmt.Println("-----------------", err)
//		fmt.Println(res)
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(res).To(HaveHTTPStatus(http.StatusOK))
//		Expect(result).ToNot(BeNil())
//	})
//}

// GET: /cities/{id}/forecasts -> BadRequest (Empty cityId)
func Test_givenAnEmptyCityId_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsBadRequest(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// execute
		result, res, err := tb.APIClient.ForecastAPI.GetByCityIdAndPeriod(context.Background(), "").Period("").Execute()
		// assert
		Expect(err).Should(HaveOccurred())
		Expect(res).To(HaveHTTPStatus(http.StatusInternalServerError))
		Expect(result).To(BeNil())
	})
}

// GET: /cities/{id}/forecasts -> NotFound (Invalid cityId)
func Test_givenAnInvalidCityId_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsNotFound(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// execute
		result, res, err := tb.APIClient.ForecastAPI.GetByCityIdAndPeriod(context.Background(), "another_invalid_id").Period("current").Execute()

		// assert
		Expect(err).Should(HaveOccurred())
		Expect(res).To(HaveHTTPStatus(http.StatusInternalServerError))
		Expect(result).To(BeNil())
	})
}

//// GET: /cities/{id}/forecasts -> OK (Current forecast)
//func Test_givenValidCityIdAndPeriodCurrent_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//		city := tb.CreateCity(testbed.NewTestCity())
//
//		result, res, err := tb.APIClient.ForecastAPI.GetByCityIdAndPeriod(context.Background(), city.ID).Period("current").Execute()
//
//		// assert
//		Expect(res).To(HaveHTTPStatus(http.StatusOK))
//		Expect(result).NotTo(BeNil())
//		Expect(err).ShouldNot(HaveOccurred())
//	})
//}

// GET: /cities/{id}/forecasts -> BadRequest (Invalid period)
func Test_givenInvalidPeriod_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsBadRequest(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// prepare
		city := tb.CreateCity(testbed.NewTestCity())

		// execute
		result, res, err := tb.APIClient.ForecastAPI.GetByCityIdAndPeriod(context.Background(), city.ID).Period("invalid_period").Execute()

		fmt.Println(err)
		// assert
		Expect(err).Should(HaveOccurred())
		Expect(result).To(BeNil())
		Expect(res).To(HaveHTTPStatus(http.StatusNotImplemented))
	})
}

// GET: /cities/{id}/forecasts -> OK (Daily forecasts)
func Test_givenValidCityIdAndPeriodDaily_whenGetByCityIdAndPeriodIsInvoked_thenResponseIsOK(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// prepare
		city := tb.CreateCity(testbed.NewTestCity())
		//currentData := tb.(city.Name)
		//tb.FakeWeatherApi.RespondOn(
		//	testbed.CreateWeatherDataRequestMatcher(city.Name, "current"),
		//	testbed.CreateWeatherDataResponder(currentData),
		//)

		result, res, err := tb.APIClient.ForecastAPI.GetByCityIdAndPeriod(context.Background(), city.ID).Period("daily").Execute()

		// assert
		Expect(err).ShouldNot(HaveOccurred())
		Expect(res).To(HaveHTTPStatus(http.StatusOK))
		Expect(result).NotTo(BeNil())

	})
}
