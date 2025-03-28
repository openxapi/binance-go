# \V3API

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesGetAccountV3**](V3API.md#UfuturesGetAccountV3) | **Get** /fapi/v3/account | Account Information V3(USER_DATA)
[**UfuturesGetBalanceV3**](V3API.md#UfuturesGetBalanceV3) | **Get** /fapi/v3/balance | Futures Account Balance V3 (USER_DATA)
[**UfuturesGetPositionRiskV3**](V3API.md#UfuturesGetPositionRiskV3) | **Get** /fapi/v3/positionRisk | Position Information V3 (USER_DATA)



## UfuturesGetAccountV3

> UfuturesGetAccountV3Resp UfuturesGetAccountV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information V3(USER_DATA)



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
	resp, r, err := apiClient.V3API.UfuturesGetAccountV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.UfuturesGetAccountV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAccountV3`: UfuturesGetAccountV3Resp
	fmt.Fprintf(os.Stdout, "Response from `V3API.UfuturesGetAccountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetAccountV3Resp**](UfuturesGetAccountV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetBalanceV3

> []UfuturesGetBalanceV3RespItem UfuturesGetBalanceV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance V3 (USER_DATA)



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
	resp, r, err := apiClient.V3API.UfuturesGetBalanceV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.UfuturesGetBalanceV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetBalanceV3`: []UfuturesGetBalanceV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `V3API.UfuturesGetBalanceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetBalanceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetBalanceV3RespItem**](UfuturesGetBalanceV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPositionRiskV3

> []UfuturesGetPositionRiskV3RespItem UfuturesGetPositionRiskV3(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Position Information V3 (USER_DATA)



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
	resp, r, err := apiClient.V3API.UfuturesGetPositionRiskV3(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.UfuturesGetPositionRiskV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionRiskV3`: []UfuturesGetPositionRiskV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `V3API.UfuturesGetPositionRiskV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPositionRiskV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetPositionRiskV3RespItem**](UfuturesGetPositionRiskV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

