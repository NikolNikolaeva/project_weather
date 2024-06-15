// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Weather API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package modules

type City struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	Country string `json:"country"`

	Longitude string `json:"longitude,omitempty"`

	Latitude string `json:"latitude,omitempty"`

	CreatedAt int64 `json:"createdAt,omitempty"`

	UpdatedAt int64 `json:"updatedAt,omitempty"`
}

// AssertCityRequired checks if the required fields are not zero-ed
func AssertCityRequired(obj City) error {
	elements := map[string]interface{}{
		"name":    obj.Name,
		"country": obj.Country,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCityConstraints checks if the values respects the defined constraints
func AssertCityConstraints(obj City) error {
	return nil
}
