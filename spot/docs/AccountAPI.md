# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpotGetAccountCommissionV3**](AccountAPI.md#SpotGetAccountCommissionV3) | **Get** /api/v3/account/commission | Query Commission Rates (USER_DATA)
[**SpotGetAccountV3**](AccountAPI.md#SpotGetAccountV3) | **Get** /api/v3/account | Account information (USER_DATA)
[**SpotGetMyAllocationsV3**](AccountAPI.md#SpotGetMyAllocationsV3) | **Get** /api/v3/myAllocations | Query Allocations (USER_DATA)
[**SpotGetMyPreventedMatchesV3**](AccountAPI.md#SpotGetMyPreventedMatchesV3) | **Get** /api/v3/myPreventedMatches | Query Prevented Matches (USER_DATA)
[**SpotGetMyTradesV3**](AccountAPI.md#SpotGetMyTradesV3) | **Get** /api/v3/myTrades | Account trade list (USER_DATA)
[**SpotGetRateLimitOrderV3**](AccountAPI.md#SpotGetRateLimitOrderV3) | **Get** /api/v3/rateLimit/order | Query Unfilled Order Count (USER_DATA)



## SpotGetAccountCommissionV3

> SpotGetAccountCommissionV3Resp SpotGetAccountCommissionV3(ctx).Symbol(symbol).Execute()

Query Commission Rates (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.SpotGetAccountCommissionV3(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.SpotGetAccountCommissionV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetAccountCommissionV3`: SpotGetAccountCommissionV3Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.SpotGetAccountCommissionV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetAccountCommissionV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**SpotGetAccountCommissionV3Resp**](SpotGetAccountCommissionV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetAccountV3

> SpotGetAccountV3Resp SpotGetAccountV3(ctx).Timestamp(timestamp).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()

Account information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	omitZeroBalances := true // bool | When set to `true`, emits only the non-zero balances of an account. <br/>Default value: `false` (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.SpotGetAccountV3(context.Background()).Timestamp(timestamp).OmitZeroBalances(omitZeroBalances).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.SpotGetAccountV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetAccountV3`: SpotGetAccountV3Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.SpotGetAccountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **omitZeroBalances** | **bool** | When set to &#x60;true&#x60;, emits only the non-zero balances of an account. &lt;br/&gt;Default value: &#x60;false&#x60; | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**SpotGetAccountV3Resp**](SpotGetAccountV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetMyAllocationsV3

> []SpotGetMyAllocationsV3RespItem SpotGetMyAllocationsV3(ctx).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Timestamp(timestamp).Execute()

Query Allocations (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromAllocationId := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 500;Max 1000 (optional) (default to 500)
	orderId := int64(789) // int64 |  (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000`. (optional)
	timestamp := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.SpotGetMyAllocationsV3(context.Background()).Symbol(symbol).StartTime(startTime).EndTime(endTime).FromAllocationId(fromAllocationId).Limit(limit).OrderId(orderId).RecvWindow(recvWindow).Timestamp(timestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.SpotGetMyAllocationsV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetMyAllocationsV3`: []SpotGetMyAllocationsV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.SpotGetMyAllocationsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetMyAllocationsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromAllocationId** | **int32** |  | 
 **limit** | **int32** | Default 500;Max 1000 | [default to 500]
 **orderId** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60;. | 
 **timestamp** | **int64** |  | 

### Return type

[**[]SpotGetMyAllocationsV3RespItem**](SpotGetMyAllocationsV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetMyPreventedMatchesV3

> []SpotGetMyPreventedMatchesV3RespItem SpotGetMyPreventedMatchesV3(ctx).Symbol(symbol).Timestamp(timestamp).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Prevented Matches (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	preventedMatchId := int64(789) // int64 |  (optional)
	orderId := int64(789) // int64 |  (optional)
	fromPreventedMatchId := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default: `500`; Max: `1000` (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.SpotGetMyPreventedMatchesV3(context.Background()).Symbol(symbol).Timestamp(timestamp).PreventedMatchId(preventedMatchId).OrderId(orderId).FromPreventedMatchId(fromPreventedMatchId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.SpotGetMyPreventedMatchesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetMyPreventedMatchesV3`: []SpotGetMyPreventedMatchesV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.SpotGetMyPreventedMatchesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetMyPreventedMatchesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **preventedMatchId** | **int64** |  | 
 **orderId** | **int64** |  | 
 **fromPreventedMatchId** | **int64** |  | 
 **limit** | **int32** | Default: &#x60;500&#x60;; Max: &#x60;1000&#x60; | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]SpotGetMyPreventedMatchesV3RespItem**](SpotGetMyPreventedMatchesV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetMyTradesV3

> []SpotGetMyTradesV3RespItem SpotGetMyTradesV3(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Account trade list (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | This can only be used in combination with `symbol`. (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.SpotGetMyTradesV3(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.SpotGetMyTradesV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetMyTradesV3`: []SpotGetMyTradesV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.SpotGetMyTradesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetMyTradesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | This can only be used in combination with &#x60;symbol&#x60;. | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]SpotGetMyTradesV3RespItem**](SpotGetMyTradesV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotGetRateLimitOrderV3

> []SpotGetRateLimitOrderV3RespItem SpotGetRateLimitOrderV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Unfilled Order Count (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.SpotGetRateLimitOrderV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.SpotGetRateLimitOrderV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetRateLimitOrderV3`: []SpotGetRateLimitOrderV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.SpotGetRateLimitOrderV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetRateLimitOrderV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]SpotGetRateLimitOrderV3RespItem**](SpotGetRateLimitOrderV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

