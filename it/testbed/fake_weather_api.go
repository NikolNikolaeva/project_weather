package testbed

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	clientapi "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/it/internal/gomisc/lang"
	arrays "github.com/NikolNikolaeva/project_weather/it/internal/gomisc/lang/array"
	"github.com/NikolNikolaeva/project_weather/it/internal/gomisc/types/pair"
)

type Responder func(http.ResponseWriter)
type RequestMatcher func(*http.Request) bool

func CreateAnyRequestMatcher() RequestMatcher {
	return func(request *http.Request) bool {
		return true
	}
}

func CreateErrorStatusResponder(status int) Responder {
	return func(writer http.ResponseWriter) {
		writer.WriteHeader(status)
		lang.Must(writer.Write([]byte(http.StatusText(status))))
	}
}

func CreateCityGetCityRequestMatcher(cityId string) RequestMatcher {
	return func(request *http.Request) bool {
		return request.Method == http.MethodGet &&
			request.URL.Path == fmt.Sprintf("/cities/%v", cityId)
	}
}

func CreateForecastsGetForecastRequestMatcher(cityId string) RequestMatcher {
	return func(request *http.Request) bool {
		return request.Method == http.MethodGet &&
			request.URL.Path == fmt.Sprintf("/forecasts/%s", cityId)
	}
}

func CreateCityGetCityResponder(cityName string, country string) Responder {
	return func(writer http.ResponseWriter) {
		writer.WriteHeader(http.StatusOK)
		lang.Must(writer.Write(lang.Must(
			json.Marshal(&clientapi.City{Name: cityName, Country: country}),
		)))
	}
}

type FakeWeatherApi interface {
	Reset() FakeWeatherApi
	RespondOn(matcher RequestMatcher, responder Responder) FakeWeatherApi
}

func NewFakeWeatherApi(address string) (FakeWeatherApi, func()) {
	result := &_FakeWeatherApi{}

	result.server = &http.Server{
		Addr: fmt.Sprintf(":%d", lang.Must(strconv.Atoi(lang.Must(url.Parse(address)).Port()))),
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			arrays.Map(
				arrays.Filter(
					result.responders,
					func(pair pair.Pair[RequestMatcher, Responder]) bool {
						return pair.Left()(request)
					},
				),
				func(_ int, pair pair.Pair[RequestMatcher, Responder]) Responder {
					return pair.Right()
				},
			)[0](writer)
		}),
	}

	go func() {
		_ = result.server.ListenAndServe()
	}()

	return result.Reset(), func() { _ = result.server.Close() }
}

type _FakeWeatherApi struct {
	server     *http.Server
	responders []pair.Pair[RequestMatcher, Responder]
}

func (self *_FakeWeatherApi) Reset() FakeWeatherApi {
	self.responders = []pair.Pair[RequestMatcher, Responder]{
		pair.NewPair[RequestMatcher, Responder](
			CreateCityGetCityRequestMatcher("4fb57332-6423-429e-a4e9-a9fad76e4c07"),
			func(writer http.ResponseWriter) {
				writer.WriteHeader(http.StatusCreated)
				lang.Must(writer.Write([]byte(`{
					"id": "4fb57332-6423-429e-a4e9-a9fad76e4c07",
					"name": "Paris",
					"country": "France",
                    "longitude":"0",
                    "latitude":"0",
				}`)))
			},
		),
		pair.NewPair[RequestMatcher, Responder](
			CreateForecastsGetForecastRequestMatcher("4fb57332-6423-429e-a4e9-a9fad76e4c07"),
			func(writer http.ResponseWriter) {
				writer.WriteHeader(http.StatusOK)
				lang.Must(writer.Write([]byte(`{
"id":"5fb57332-6423-429e-a4e9-a9fad76e4c07",
					"cityId": "4fb57332-6423-429e-a4e9-a9fad76e4c07",
					"temperature":"15.0",
"condition":"Sunny",
"forecastDate": "2006-01-02",
				}`)))
			},
		),
		pair.NewPair(
			CreateAnyRequestMatcher(),
			CreateErrorStatusResponder(http.StatusInternalServerError),
		),
	}

	return self
}

func (self *_FakeWeatherApi) RespondOn(matcher RequestMatcher, responder Responder) FakeWeatherApi {
	self.responders = append(
		[]pair.Pair[RequestMatcher, Responder]{pair.NewPair(matcher, responder)},
		self.responders...,
	)

	return self
}
