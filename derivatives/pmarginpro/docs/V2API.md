# \V2API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginproGetPortfolioAccountV2**](V2API.md#PmarginproGetPortfolioAccountV2) | **Get** /sapi/v2/portfolio/account | Get Portfolio Margin Pro SPAN Account Info(USER_DATA)
[**PmarginproGetPortfolioCollateralRateV2**](V2API.md#PmarginproGetPortfolioCollateralRateV2) | **Get** /sapi/v2/portfolio/collateralRate | Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)



## PmarginproGetPortfolioAccountV2

> PmarginproGetPortfolioAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro SPAN Account Info(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.V2API.PmarginproGetPortfolioAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.PmarginproGetPortfolioAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioCollateralRateV2

> []PmarginproGetPortfolioCollateralRateV2RespItem PmarginproGetPortfolioCollateralRateV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.PmarginproGetPortfolioCollateralRateV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.PmarginproGetPortfolioCollateralRateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioCollateralRateV2`: []PmarginproGetPortfolioCollateralRateV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `V2API.PmarginproGetPortfolioCollateralRateV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioCollateralRateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginproGetPortfolioCollateralRateV2RespItem**](PmarginproGetPortfolioCollateralRateV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

