# Forecast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CityId** | **string** |  | 
**ForecastDate** | **string** |  | 
**Temperature** | **string** |  | 
**Condition** | **string** |  | 
**CreatedAt** | Pointer to **int64** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewForecast

`func NewForecast(cityId string, forecastDate string, temperature string, condition string, ) *Forecast`

NewForecast instantiates a new Forecast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastWithDefaults

`func NewForecastWithDefaults() *Forecast`

NewForecastWithDefaults instantiates a new Forecast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Forecast) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Forecast) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Forecast) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Forecast) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCityId

`func (o *Forecast) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *Forecast) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *Forecast) SetCityId(v string)`

SetCityId sets CityId field to given value.


### GetForecastDate

`func (o *Forecast) GetForecastDate() string`

GetForecastDate returns the ForecastDate field if non-nil, zero value otherwise.

### GetForecastDateOk

`func (o *Forecast) GetForecastDateOk() (*string, bool)`

GetForecastDateOk returns a tuple with the ForecastDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastDate

`func (o *Forecast) SetForecastDate(v string)`

SetForecastDate sets ForecastDate field to given value.


### GetTemperature

`func (o *Forecast) GetTemperature() string`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *Forecast) GetTemperatureOk() (*string, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *Forecast) SetTemperature(v string)`

SetTemperature sets Temperature field to given value.


### GetCondition

`func (o *Forecast) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Forecast) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Forecast) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetCreatedAt

`func (o *Forecast) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Forecast) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Forecast) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Forecast) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Forecast) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Forecast) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Forecast) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Forecast) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


