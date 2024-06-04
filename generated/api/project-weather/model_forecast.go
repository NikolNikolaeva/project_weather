/*
Weather API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Forecast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Forecast{}

// Forecast struct for Forecast
type Forecast struct {
	Id           *string `json:"id,omitempty"`
	CityId       string  `json:"cityId"`
	ForecastDate string  `json:"forecastDate"`
	Temperature  string  `json:"temperature"`
	Condition    string  `json:"condition"`
	CreatedAt    *int64  `json:"createdAt,omitempty"`
	UpdatedAt    *int64  `json:"updatedAt,omitempty"`
}

type _Forecast Forecast

// NewForecast instantiates a new Forecast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecast(cityId string, forecastDate string, temperature string, condition string) *Forecast {
	this := Forecast{}
	this.CityId = cityId
	this.ForecastDate = forecastDate
	this.Temperature = temperature
	this.Condition = condition
	return &this
}

// NewForecastWithDefaults instantiates a new Forecast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastWithDefaults() *Forecast {
	this := Forecast{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Forecast) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forecast) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Forecast) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Forecast) SetId(v string) {
	o.Id = &v
}

// GetCityId returns the CityId field value
func (o *Forecast) GetCityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CityId
}

// GetCityIdOk returns a tuple with the CityId field value
// and a boolean to check if the value has been set.
func (o *Forecast) GetCityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CityId, true
}

// SetCityId sets field value
func (o *Forecast) SetCityId(v string) {
	o.CityId = v
}

// GetForecastDate returns the ForecastDate field value
func (o *Forecast) GetForecastDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForecastDate
}

// GetForecastDateOk returns a tuple with the ForecastDate field value
// and a boolean to check if the value has been set.
func (o *Forecast) GetForecastDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForecastDate, true
}

// SetForecastDate sets field value
func (o *Forecast) SetForecastDate(v string) {
	o.ForecastDate = v
}

// GetTemperature returns the Temperature field value
func (o *Forecast) GetTemperature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value
// and a boolean to check if the value has been set.
func (o *Forecast) GetTemperatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Temperature, true
}

// SetTemperature sets field value
func (o *Forecast) SetTemperature(v string) {
	o.Temperature = v
}

// GetCondition returns the Condition field value
func (o *Forecast) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *Forecast) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *Forecast) SetCondition(v string) {
	o.Condition = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Forecast) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forecast) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Forecast) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Forecast) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Forecast) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Forecast) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Forecast) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Forecast) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o Forecast) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Forecast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["cityId"] = o.CityId
	toSerialize["forecastDate"] = o.ForecastDate
	toSerialize["temperature"] = o.Temperature
	toSerialize["condition"] = o.Condition
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *Forecast) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cityId",
		"forecastDate",
		"temperature",
		"condition",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varForecast := _Forecast{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForecast)

	if err != nil {
		return err
	}

	*o = Forecast(varForecast)

	return err
}

type NullableForecast struct {
	value *Forecast
	isSet bool
}

func (v NullableForecast) Get() *Forecast {
	return v.value
}

func (v *NullableForecast) Set(val *Forecast) {
	v.value = val
	v.isSet = true
}

func (v NullableForecast) IsSet() bool {
	return v.isSet
}

func (v *NullableForecast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecast(val *Forecast) *NullableForecast {
	return &NullableForecast{value: val, isSet: true}
}

func (v NullableForecast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}