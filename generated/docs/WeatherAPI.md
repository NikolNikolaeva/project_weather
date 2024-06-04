# \WeatherAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWeatherByCity**](WeatherAPI.md#GetWeatherByCity) | **Get** /weather/{city}/{period} | 



## GetWeatherByCity

> Forecast GetWeatherByCity(ctx, city, period).Execute()





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
	city := "city_example" // string | 
	period := "period_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WeatherAPI.GetWeatherByCity(context.Background(), city, period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WeatherAPI.GetWeatherByCity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWeatherByCity`: Forecast
	fmt.Fprintf(os.Stdout, "Response from `WeatherAPI.GetWeatherByCity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**city** | **string** |  | 
**period** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeatherByCityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Forecast**](Forecast.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

