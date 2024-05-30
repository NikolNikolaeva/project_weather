package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NikolNikolaeva/project_weather/resources/swagger"
	"github.com/stretchr/testify/assert"
)

func Test_NewWeatherDataGetter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer server.Close()
	client := server.Client()

	newWeatherDataGetter := NewWeatherDataGetter(client)
	assert.NotNil(t, newWeatherDataGetter)
}

func Test_WeatherDataGetter_GetData(t *testing.T) {
	testCases := []struct {
		description    string
		serverResponse string
		serverStatus   int
		expectedData   *swagger.WeatherDTO
		expectedError  string
	}{
		{
			description: "Successful data retrieval",
			serverResponse: `{
                "location": {
                    "name": "Sofia",
                    "country": "Bulgaria",
                    "lat": 42.698,
                    "lon": 23.322
                },
                "Current": {
                    "last_updated": "2023-01-01 15:04",
                    "temp_c": 15.0,
                    "condition": {
                        "text": "Sunny"
                    }
                }
            }`,
			serverStatus: http.StatusOK,
			expectedData: &swagger.WeatherDTO{
				Location: swagger.LocationDTO{
					Name:    "Sofia",
					Country: "Bulgaria",
					Lat:     42.698,
					Lon:     23.322,
				},
				Current: swagger.CurrentDTO{
					LastUpdated: "2023-01-01 15:04",
					TempC:       15.0,
					Condition: swagger.ConditionDTO{
						Text: "Sunny",
					},
				},
			},
			expectedError: "",
		},
		{
			description:    "HTTP error",
			serverResponse: ``,
			serverStatus:   http.StatusInternalServerError,
			expectedData:   nil,
			expectedError:  "failed to fetch weather data: status code 500",
		},
		{
			description:    "JSON parsing error",
			serverResponse: `{invalid json}`,
			serverStatus:   http.StatusOK,
			expectedData:   nil,
			expectedError:  "failed to parse weather data: invalid character",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.serverStatus)
				fmt.Fprintln(w, testCase.serverResponse)
			}))
			defer server.Close()

			client := server.Client()
			dataGetter := NewWeatherDataGetter(client)

			data, err := dataGetter.GetData(server.URL)

			if testCase.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), testCase.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedData, data)
			}
		})
	}
}
