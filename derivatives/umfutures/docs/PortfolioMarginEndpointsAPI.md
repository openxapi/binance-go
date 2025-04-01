# \PortfolioMarginEndpointsAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UmfuturesGetPmAccountInfoV1**](PortfolioMarginEndpointsAPI.md#UmfuturesGetPmAccountInfoV1) | **Get** /fapi/v1/pmAccountInfo | Classic Portfolio Margin Account Information (USER_DATA)



## UmfuturesGetPmAccountInfoV1

> UmfuturesGetPmAccountInfoV1Resp UmfuturesGetPmAccountInfoV1(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Classic Portfolio Margin Account Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginEndpointsAPI.UmfuturesGetPmAccountInfoV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginEndpointsAPI.UmfuturesGetPmAccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetPmAccountInfoV1`: UmfuturesGetPmAccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginEndpointsAPI.UmfuturesGetPmAccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetPmAccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UmfuturesGetPmAccountInfoV1Resp**](UmfuturesGetPmAccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

