# \PortfolioMarginEndpointsAPI

All URIs are relative to *https://dapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfuturesGetPmAccountInfoV1**](PortfolioMarginEndpointsAPI.md#CfuturesGetPmAccountInfoV1) | **Get** /dapi/v1/pmAccountInfo | Classic Portfolio Margin Account Information (USER_DATA)



## CfuturesGetPmAccountInfoV1

> CfuturesGetPmAccountInfoV1Resp CfuturesGetPmAccountInfoV1(ctx).Asset(asset).RecvWindow(recvWindow).Execute()

Classic Portfolio Margin Account Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginEndpointsAPI.CfuturesGetPmAccountInfoV1(context.Background()).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginEndpointsAPI.CfuturesGetPmAccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPmAccountInfoV1`: CfuturesGetPmAccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginEndpointsAPI.CfuturesGetPmAccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPmAccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetPmAccountInfoV1Resp**](CfuturesGetPmAccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

