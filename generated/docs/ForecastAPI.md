# \ForecastAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForecast**](ForecastAPI.md#CreateForecast) | **Post** /forecasts | 
[**GetAllForecasts**](ForecastAPI.md#GetAllForecasts) | **Get** /forecasts | 
[**UpdateForecast**](ForecastAPI.md#UpdateForecast) | **Put** /forecasts/{id} | 



## CreateForecast

> Forecast CreateForecast(ctx).Forecast(forecast).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	forecast := *openapiclient.NewForecast("922613c8-60ce-42f4-9823-a0ee9df38828", "ForecastDate_example", "20.0", "sunny") // Forecast |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForecastAPI.CreateForecast(context.Background()).Forecast(forecast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForecastAPI.CreateForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateForecast`: Forecast
	fmt.Fprintf(os.Stdout, "Response from `ForecastAPI.CreateForecast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forecast** | [**Forecast**](Forecast.md) |  | 

### Return type

[**Forecast**](Forecast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllForecasts

> GetAllForecasts200Response GetAllForecasts(ctx).Page(page).Limit(limit).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int64(789) // int64 |  (optional) (default to 1)
	limit := int64(789) // int64 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForecastAPI.GetAllForecasts(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForecastAPI.GetAllForecasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllForecasts`: GetAllForecasts200Response
	fmt.Fprintf(os.Stdout, "Response from `ForecastAPI.GetAllForecasts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllForecastsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | [default to 1]
 **limit** | **int64** |  | [default to 20]

### Return type

[**GetAllForecasts200Response**](GetAllForecasts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateForecast

> Forecast UpdateForecast(ctx, id).Forecast(forecast).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | 
	forecast := *openapiclient.NewForecast("922613c8-60ce-42f4-9823-a0ee9df38828", "ForecastDate_example", "20.0", "sunny") // Forecast |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForecastAPI.UpdateForecast(context.Background(), id).Forecast(forecast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForecastAPI.UpdateForecast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateForecast`: Forecast
	fmt.Fprintf(os.Stdout, "Response from `ForecastAPI.UpdateForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forecast** | [**Forecast**](Forecast.md) |  | 

### Return type

[**Forecast**](Forecast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

