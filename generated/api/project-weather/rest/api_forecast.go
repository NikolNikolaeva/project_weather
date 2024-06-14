// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Weather API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ForecastAPIController binds http requests to an api service and writes the service results to the http response
type ForecastAPIController struct {
	service      ForecastAPIServicer
	errorHandler ErrorHandler
}

// ForecastAPIOption for how the controller is set up.
type ForecastAPIOption func(*ForecastAPIController)

// WithForecastAPIErrorHandler inject ErrorHandler into controller
func WithForecastAPIErrorHandler(h ErrorHandler) ForecastAPIOption {
	return func(c *ForecastAPIController) {
		c.errorHandler = h
	}
}

// NewForecastAPIController creates a default api controller
func NewForecastAPIController(s ForecastAPIServicer, opts ...ForecastAPIOption) *ForecastAPIController {
	controller := &ForecastAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ForecastAPIController
func (c *ForecastAPIController) Routes() Routes {
	return Routes{
		"GetForecastsByCityIdAndPeriod": Route{
			strings.ToUpper("Get"),
			"/api/cities/{id}/forecasts",
			c.GetForecastsByCityIdAndPeriod,
		},
	}
}

// GetForecastsByCityIdAndPeriod -
func (c *ForecastAPIController) GetForecastsByCityIdAndPeriod(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	var periodParam string
	if query.Has("period") {
		param := query.Get("period")

		periodParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "period"}, nil)
		return
	}
	result, err := c.service.GetForecastsByCityIdAndPeriod(r.Context(), idParam, periodParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}