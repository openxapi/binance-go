# \V2API

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UmfuturesGetAccountV2**](V2API.md#UmfuturesGetAccountV2) | **Get** /fapi/v2/account | Account Information V2(USER_DATA)
[**UmfuturesGetBalanceV2**](V2API.md#UmfuturesGetBalanceV2) | **Get** /fapi/v2/balance | Futures Account Balance V2 (USER_DATA)
[**UmfuturesGetPositionRiskV2**](V2API.md#UmfuturesGetPositionRiskV2) | **Get** /fapi/v2/positionRisk | Position Information V2 (USER_DATA)
[**UmfuturesGetTickerPriceV2**](V2API.md#UmfuturesGetTickerPriceV2) | **Get** /fapi/v2/ticker/price | Symbol Price Ticker V2



## UmfuturesGetAccountV2

> UmfuturesGetAccountV2Resp UmfuturesGetAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information V2(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UmfuturesGetAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UmfuturesGetAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetAccountV2`: UmfuturesGetAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.UmfuturesGetAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UmfuturesGetAccountV2Resp**](UmfuturesGetAccountV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetBalanceV2

> []UmfuturesGetBalanceV2RespItem UmfuturesGetBalanceV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance V2 (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UmfuturesGetBalanceV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UmfuturesGetBalanceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetBalanceV2`: []UmfuturesGetBalanceV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `V2API.UmfuturesGetBalanceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetBalanceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UmfuturesGetBalanceV2RespItem**](UmfuturesGetBalanceV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetPositionRiskV2

> []UmfuturesGetPositionRiskV2RespItem UmfuturesGetPositionRiskV2(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V2 (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UmfuturesGetPositionRiskV2(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UmfuturesGetPositionRiskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetPositionRiskV2`: []UmfuturesGetPositionRiskV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `V2API.UmfuturesGetPositionRiskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetPositionRiskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UmfuturesGetPositionRiskV2RespItem**](UmfuturesGetPositionRiskV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetTickerPriceV2

> UmfuturesGetTickerPriceV2Resp UmfuturesGetTickerPriceV2(ctx).Symbol(symbol).Execute()

Symbol Price Ticker V2



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.UmfuturesGetTickerPriceV2(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.UmfuturesGetTickerPriceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetTickerPriceV2`: UmfuturesGetTickerPriceV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.UmfuturesGetTickerPriceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetTickerPriceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**UmfuturesGetTickerPriceV2Resp**](UmfuturesGetTickerPriceV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

