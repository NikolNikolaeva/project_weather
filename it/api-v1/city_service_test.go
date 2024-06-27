////go:build integration

package api_v1

import (
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/NikolNikolaeva/project_weather/it/testbed"
	"github.com/NikolNikolaeva/project_weather/it/testbed/generated/client"
)

// GET: /api/cities -> OK (No cities)
func Test_givenThatNoCitiesExist_whenGetCitiesIsInvoked_thenResponseIsOK(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// execute
		result, response, err := tb.APIClient.CityAPI.
			GetAll(tb.Context).Execute()

		// assert
		Expect(err).ShouldNot(HaveOccurred())
		Expect(result).To(BeNil())
		Expect(response).To(HaveHTTPStatus(http.StatusOK))
		Expect(result).To(BeNil())
	})
}

// GET: /api/cities -> OK  -> 1 city
func Test_givenThatNoCitiesExist_whenGetAllIsInvoked_thenResponseIsOK(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// prepare
		city := tb.CreateCity(testbed.NewTestCity())

		// execute
		result, response, err := tb.APIClient.CityAPI.
			GetAll(tb.Context).Execute()

		// assert
		Expect(err).ShouldNot(HaveOccurred())
		Expect(response.StatusCode).To(Equal(http.StatusOK))
		Expect(result).ToNot(BeNil())
		Expect(result).To(Equal([]client.City{testbed.AsCityDto(city)}))
	})
}

// DELETE: /api/cities/{id} -> OK
func Test_givenCityExists_whenDeleteByIdIsInvoked_thenResponseIsOK(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// prepare
		city := tb.CreateCity(testbed.NewTestCity())

		// execute
		response, err := tb.APIClient.CityAPI.DeleteById(tb.Context, city.ID).Execute()

		// assert
		Expect(err).ShouldNot(HaveOccurred())
		Expect(response).To(HaveHTTPStatus(http.StatusOK))

	})
}

// DELETE: /api/cities/{non-existent id} -> Gone
func Test_givenAnUnknownCity_whenDeleteCityIsInvoked_thenResponseIsNotFound(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// execute
		response, err := tb.APIClient.CityAPI.
			DeleteById(tb.Context, "some_invalid_id").Execute()

		// assert
		Expect(err).Should(HaveOccurred())
		Expect(response).To(HaveHTTPStatus(http.StatusNotFound))
	})
}

//// DELETE: /api/cities/{empty id} -> BadRequest
//func Test_givenAnEmptyId_whenDeleteCityIsInvoked_thenResponseIsBadRequest(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// execute
//		response, err := tb.APIClient.CityAPI.
//			DeleteById(tb.Context, "").Execute()
//
//		// assert
//		Expect(err).Should(HaveOccurred())
//		//Expect(response.Body).To(Equal("id is required"))
//		Expect(response.StatusCode).To(Equal(http.StatusBadRequest))
//	})
//}   ???

// GET: /api/cities/{non-existent id} -> Gone
func Test_givenThatCitiesExist_whenGetCityIsInvokedWithWrongId_thenResponseIsStatusNotFound(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// execute
		result, response, err := tb.APIClient.CityAPI.
			GetById(tb.Context, "random-id").Execute()
		// assert
		Expect(err).Should(HaveOccurred())
		Expect(response).To(HaveHTTPStatus(http.StatusInternalServerError))
		Expect(result).To(BeNil())
	})
}

// GET: /api/cities/{id} -> OK
func Test_givenThatCitiesExist_whenGetCityIsInvoked_thenResponseIsOK(t *testing.T) {
	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
		defer tb.Reset()

		// prepare
		city := tb.CreateCity(testbed.NewTestCity())

		// execute
		result, response, err := tb.APIClient.CityAPI.
			GetById(tb.Context, city.ID).Execute()
		// assert
		Expect(err).ShouldNot(HaveOccurred())
		Expect(response).To(HaveHTTPStatus(http.StatusOK))
		Expect(result.Name).To(Equal(testbed.AsCityDto(city).Name))
		Expect(result.Country).To(Equal(testbed.AsCityDto(city).Country))
		Expect(result.Id).To(Equal(testbed.AsCityDto(city).Id))
	})
}

//// Get: /api/cities/{id} -> BadRequest
//func Test_givenThatCitiesExist_whenGetCityIsInvoked_thenResponseIsBadRequest(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// execute
//		result, response, err := tb.APIClient.CityAPI.
//			GetById(tb.Context, "").Execute()
//
//		// assert
//		Expect(err).To(BeNil())
//		Expect(response).To(HaveHTTPStatus(http.StatusBadRequest))
//		Expect(*result).To(Equal("id is required"))
//	})
//}

//// POST: /api/v1/tasks -> Created
//func Test_givenThatNoCitiesExist_whenCreateCityIsInvoked_thenResponseIsCreated(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		city := tb.CreateCity(testbed.NewTestCity())
//
//		// execute
//		result, response, err := tb.APIClient.CityAPI.Register(tb.Context).
//			City(*client.NewCity(city.Name, city.Country)).
//			Execute()
//
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(response).To(HaveHTTPStatus(http.StatusCreated))
//		Expect(result).ToNot(BeNil())
//		Expect(result).To(Equal(city))
//		Expect(tb.Dao.City.Count()).To(Equal(int64(1)))
//	})
//}
//
//// PUT: /api/cities/{id} -> OK
//func Test_givenACity_whenUpdateSingleCityIsInvoked_thenResponseIsOK(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//
//		city := tb.CreateCity(testbed.NewTestCity())
//		Long := "123"
//		Lat := "456"
//
//		// execute
//		updated, response, err := tb.APIClient.CityAPI.
//			UpdateByID(tb.Context, city.ID).
//			City(client.City{
//				Name:      city.Name,
//				Country:   city.Country,
//				Latitude:  &Lat,
//				Longitude: &Long,
//			}).Execute()
//
//		// assert
//		Expect(err).ShouldNot(HaveOccurred())
//		Expect(response).To(HaveHTTPStatus(http.StatusOK))
//		Expect(updated.Latitude).To(Equal(&Lat))
//		Expect(updated.Longitude).To(Equal(&Long))
//	})
//}
//
//// PUT: /api/cities/{invalid_id} -> StatusInternalServerError
//func Test_givenACity_whenUpdateSingleCityIsInvoked_thenResponseIsStatusInternalServerError(t *testing.T) {
//	testbed.Prepare(t, 30*time.Second, func(t *testing.T, tb *testbed.Runner) {
//		defer tb.Reset()
//
//		// prepare
//
//		city := tb.CreateCity(testbed.NewTestCity())
//		Long := "123"
//		Lat := "456"
//
//		// execute
//		updated, response, err := tb.APIClient.CityAPI.
//			UpdateByID(tb.Context, "invalid_id").
//			City(client.City{
//				Name:      city.Name,
//				Country:   city.Country,
//				Latitude:  &Lat,
//				Longitude: &Long,
//			}).Execute()
//		fmt.Println("---------------", err)
//		// assert
//		//Expect(err).Should(HaveOccurred())
//		Expect(response).To(HaveHTTPStatus(http.StatusInternalServerError))
//		Expect(updated).To(BeNil())
//	})
//}
