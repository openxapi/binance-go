# \V3API

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UmfuturesGetAccountV3**](V3API.md#UmfuturesGetAccountV3) | **Get** /fapi/v3/account | Account Information V3(USER_DATA)
[**UmfuturesGetBalanceV3**](V3API.md#UmfuturesGetBalanceV3) | **Get** /fapi/v3/balance | Futures Account Balance V3 (USER_DATA)
[**UmfuturesGetPositionRiskV3**](V3API.md#UmfuturesGetPositionRiskV3) | **Get** /fapi/v3/positionRisk | Position Information V3 (USER_DATA)



## UmfuturesGetAccountV3

> UmfuturesGetAccountV3Resp UmfuturesGetAccountV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information V3(USER_DATA)



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
	resp, r, err := apiClient.V3API.UmfuturesGetAccountV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.UmfuturesGetAccountV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetAccountV3`: UmfuturesGetAccountV3Resp
	fmt.Fprintf(os.Stdout, "Response from `V3API.UmfuturesGetAccountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UmfuturesGetAccountV3Resp**](UmfuturesGetAccountV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetBalanceV3

> []UmfuturesGetBalanceV3RespItem UmfuturesGetBalanceV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance V3 (USER_DATA)



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
	resp, r, err := apiClient.V3API.UmfuturesGetBalanceV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.UmfuturesGetBalanceV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetBalanceV3`: []UmfuturesGetBalanceV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `V3API.UmfuturesGetBalanceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetBalanceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UmfuturesGetBalanceV3RespItem**](UmfuturesGetBalanceV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetPositionRiskV3

> []UmfuturesGetPositionRiskV3RespItem UmfuturesGetPositionRiskV3(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V3 (USER_DATA)



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
	resp, r, err := apiClient.V3API.UmfuturesGetPositionRiskV3(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.UmfuturesGetPositionRiskV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetPositionRiskV3`: []UmfuturesGetPositionRiskV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `V3API.UmfuturesGetPositionRiskV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetPositionRiskV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UmfuturesGetPositionRiskV3RespItem**](UmfuturesGetPositionRiskV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

