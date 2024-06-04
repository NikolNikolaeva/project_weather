package services

import (
	"encoding/json"
	"fmt"
	"github.com/NikolNikolaeva/project_weather/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	swagger "github.com/weatherapicom/go"
)

func TestWeatherDataGetter_GetCurrentData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &swagger.InlineResponse2001{
			Location: &swagger.Location{
				Name:    "Sofia",
				Country: "Bulgaria",
				Lat:     42.698,
				Lon:     23.322,
			},
			Current: &swagger.Current{
				TempC: 15.0,
				Condition: &swagger.CurrentCondition{
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

	dataGetter := mocks.NewMockWeatherDataGetter(controller)
	data, err := dataGetter.GetData(server.URL)

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "Sofia", data.Location.Name)
	assert.Equal(t, "Bulgaria", data.Location.Country)
	assert.Equal(t, 15.0, data.Current.TempC)
	assert.Equal(t, "Sunny", data.Current.Condition.Text)
}

func TestWeatherDataGetter_GetCurrentData_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	controller := gomock.NewController(t)
	defer controller.Finish()

	dataGetter := mocks.NewMockWeatherDataGetter(controller)
	data, err := dataGetter.GetData(server.URL)

	assert.Error(t, err)
	assert.Nil(t, data)
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

	dataGetter := mocks.NewMockWeatherDataGetter(controller)
	data, err := dataGetter.GetData(server.URL)

	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Contains(t, err.Error(), "failed to parse weather data")
}
