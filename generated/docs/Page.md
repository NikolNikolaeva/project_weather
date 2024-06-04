# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int64** |  | 
**Next** | Pointer to **int64** |  | [optional] 
**Last** | Pointer to **int64** |  | [optional] 
**Count** | **int64** |  | 
**Current** | Pointer to **int64** |  | [optional] 
**Previous** | Pointer to **int64** |  | [optional] 

## Methods

### NewPage

`func NewPage(limit int64, count int64, ) *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Page) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Page) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Page) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *Page) GetNext() int64`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Page) GetNextOk() (*int64, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Page) SetNext(v int64)`

SetNext sets Next field to given value.

### HasNext

`func (o *Page) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetLast

`func (o *Page) GetLast() int64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *Page) GetLastOk() (*int64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *Page) SetLast(v int64)`

SetLast sets Last field to given value.

### HasLast

`func (o *Page) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetCount

`func (o *Page) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Page) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Page) SetCount(v int64)`

SetCount sets Count field to given value.


### GetCurrent

`func (o *Page) GetCurrent() int64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Page) GetCurrentOk() (*int64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Page) SetCurrent(v int64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *Page) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetPrevious

`func (o *Page) GetPrevious() int64`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Page) GetPreviousOk() (*int64, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Page) SetPrevious(v int64)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *Page) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


