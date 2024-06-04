package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_services "github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"
	"go.uber.org/mock/gomock"

	api "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
	"github.com/stretchr/testify/assert"
)

func TestWeatherDataGetter_GetCurrentData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &api.InlineResponse2001{
			Location: &api.Location{
				Name:    "Sofia",
				Country: "Bulgaria",
				Lat:     42.698,
				Lon:     23.322,
			},
			Current: &api.Current{
				TempC: 15.0,
				Condition: &api.CurrentCondition{
					Text: "Sunny",
				},
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	controller := gomock.NewController(t)
	defer controller.Finish()

	dataGetter := mock_services.NewMockWeatherDataGetter(controller)
	current, location, err := dataGetter.GetCurrentData(server.URL, "")

	assert.NoError(t, err)
	assert.NotNil(t, current)
	assert.NotNil(t, location)
	assert.Equal(t, 15.0, current.TempC)
	assert.Equal(t, "Sofia", location.Name)
	assert.Equal(t, "Bulgaria", location.Country)
	assert.Equal(t, "Sunny", current.Condition.Text)
}

func TestWeatherDataGetter_GetCurrentData_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	controller := gomock.NewController(t)
	defer controller.Finish()

	dataGetter := mock_services.NewMockWeatherDataGetter(controller)
	current, location, err := dataGetter.GetCurrentData(server.URL, "")

	assert.Error(t, err)
	assert.Nil(t, current)
	assert.Nil(t, location)
	assert.Contains(t, err.Error(), "status code 500")
}

func TestWeatherDataGetter_GetCurrentData_JSONError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{invalid json}`)
	}))
	defer server.Close()

	controller := gomock.NewController(t)
	defer controller.Finish()

	dataGetter := mock_services.NewMockWeatherDataGetter(controller)
	current, location, err := dataGetter.GetCurrentData(server.URL, "")

	assert.Error(t, err)
	assert.Nil(t, current)
	assert.Nil(t, location)
	assert.Contains(t, err.Error(), "failed to parse weather data")
}
