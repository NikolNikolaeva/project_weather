# GetAllForecasts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | [**Page**](Page.md) |  | 
**Data** | [**[]Forecast**](Forecast.md) |  | 

## Methods

### NewGetAllForecasts200Response

`func NewGetAllForecasts200Response(page Page, data []Forecast, ) *GetAllForecasts200Response`

NewGetAllForecasts200Response instantiates a new GetAllForecasts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllForecasts200ResponseWithDefaults

`func NewGetAllForecasts200ResponseWithDefaults() *GetAllForecasts200Response`

NewGetAllForecasts200ResponseWithDefaults instantiates a new GetAllForecasts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetAllForecasts200Response) GetPage() Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetAllForecasts200Response) GetPageOk() (*Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetAllForecasts200Response) SetPage(v Page)`

SetPage sets Page field to given value.


### GetData

`func (o *GetAllForecasts200Response) GetData() []Forecast`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllForecasts200Response) GetDataOk() (*[]Forecast, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllForecasts200Response) SetData(v []Forecast)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


