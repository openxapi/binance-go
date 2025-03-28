# \V2API

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesGetAccountV2**](V2API.md#UfuturesGetAccountV2) | **Get** /fapi/v2/account | Account Information V2(USER_DATA)
[**UfuturesGetBalanceV2**](V2API.md#UfuturesGetBalanceV2) | **Get** /fapi/v2/balance | Futures Account Balance V2 (USER_DATA)
[**UfuturesGetPositionRiskV2**](V2API.md#UfuturesGetPositionRiskV2) | **Get** /fapi/v2/positionRisk | Position Information V2 (USER_DATA)
[**UfuturesGetTickerPriceV2**](V2API.md#UfuturesGetTickerPriceV2) | **Get** /fapi/v2/ticker/price | Symbol Price Ticker V2



## UfuturesGetAccountV2

> UfuturesGetAccountV2Resp UfuturesGetAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information V2(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UfuturesGetAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UfuturesGetAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAccountV2`: UfuturesGetAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.UfuturesGetAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetAccountV2Resp**](UfuturesGetAccountV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetBalanceV2

> []UfuturesGetBalanceV2RespItem UfuturesGetBalanceV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance V2 (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UfuturesGetBalanceV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UfuturesGetBalanceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetBalanceV2`: []UfuturesGetBalanceV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `V2API.UfuturesGetBalanceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetBalanceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetBalanceV2RespItem**](UfuturesGetBalanceV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPositionRiskV2

> []UfuturesGetPositionRiskV2RespItem UfuturesGetPositionRiskV2(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V2 (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UfuturesGetPositionRiskV2(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UfuturesGetPositionRiskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionRiskV2`: []UfuturesGetPositionRiskV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `V2API.UfuturesGetPositionRiskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPositionRiskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetPositionRiskV2RespItem**](UfuturesGetPositionRiskV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTickerPriceV2

> UfuturesGetTickerPriceV2Resp UfuturesGetTickerPriceV2(ctx).Symbol(symbol).Execute()

Symbol Price Ticker V2



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UfuturesGetTickerPriceV2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UfuturesGetTickerPriceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTickerPriceV2`: UfuturesGetTickerPriceV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.UfuturesGetTickerPriceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTickerPriceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UfuturesGetTickerPriceV2Resp**](UfuturesGetTickerPriceV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

