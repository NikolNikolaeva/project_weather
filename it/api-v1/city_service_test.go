package api_v1

import (
	"github.com/NikolNikolaeva/project_weather/it/testbed"
	"github.com/NikolNikolaeva/project_weather/it/testbed/generated/client"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
	"time"
)

// GET: /cities -> OK  -> 1 city
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
		Expect(result).ToNot(BeNil())
		Expect(response).To(HaveHTTPStatus(http.StatusOK))
		Expect(result).To(Equal([]client.City{testbed.AsCityDto(city)}))
	})
}

// DELETE: /cities/{id} -> OK
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
		Expect(response.Body).To(Equal([]client.City{testbed.AsCityDto(city)}))

	})
}
