/*
Weather API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetAllForecasts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllForecasts200Response{}

// GetAllForecasts200Response struct for GetAllForecasts200Response
type GetAllForecasts200Response struct {
	Page Page       `json:"page"`
	Data []Forecast `json:"data"`
}

type _GetAllForecasts200Response GetAllForecasts200Response

// NewGetAllForecasts200Response instantiates a new GetAllForecasts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllForecasts200Response(page Page, data []Forecast) *GetAllForecasts200Response {
	this := GetAllForecasts200Response{}
	this.Page = page
	this.Data = data
	return &this
}

// NewGetAllForecasts200ResponseWithDefaults instantiates a new GetAllForecasts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllForecasts200ResponseWithDefaults() *GetAllForecasts200Response {
	this := GetAllForecasts200Response{}
	return &this
}

// GetPage returns the Page field value
func (o *GetAllForecasts200Response) GetPage() Page {
	if o == nil {
		var ret Page
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *GetAllForecasts200Response) GetPageOk() (*Page, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *GetAllForecasts200Response) SetPage(v Page) {
	o.Page = v
}

// GetData returns the Data field value
func (o *GetAllForecasts200Response) GetData() []Forecast {
	if o == nil {
		var ret []Forecast
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAllForecasts200Response) GetDataOk() ([]Forecast, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetAllForecasts200Response) SetData(v []Forecast) {
	o.Data = v
}

func (o GetAllForecasts200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllForecasts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetAllForecasts200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"page",
		"data",
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

	varGetAllForecasts200Response := _GetAllForecasts200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAllForecasts200Response)

	if err != nil {
		return err
	}

	*o = GetAllForecasts200Response(varGetAllForecasts200Response)

	return err
}

type NullableGetAllForecasts200Response struct {
	value *GetAllForecasts200Response
	isSet bool
}

func (v NullableGetAllForecasts200Response) Get() *GetAllForecasts200Response {
	return v.value
}

func (v *NullableGetAllForecasts200Response) Set(val *GetAllForecasts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllForecasts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllForecasts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllForecasts200Response(val *GetAllForecasts200Response) *NullableGetAllForecasts200Response {
	return &NullableGetAllForecasts200Response{value: val, isSet: true}
}

func (v NullableGetAllForecasts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllForecasts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
