# \MarketDataAPI

All URIs are relative to *https://papi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginGetPingV1**](MarketDataAPI.md#PmarginGetPingV1) | **Get** /papi/v1/ping | Test Connectivity



## PmarginGetPingV1

> map[string]interface{} PmarginGetPingV1(ctx).Execute()

Test Connectivity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.PmarginGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.PmarginGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.PmarginGetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetPingV1Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

