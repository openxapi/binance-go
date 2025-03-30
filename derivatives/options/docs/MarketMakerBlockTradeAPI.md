# \MarketMakerBlockTradeAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsCreateBlockOrderCreateV1**](MarketMakerBlockTradeAPI.md#OptionsCreateBlockOrderCreateV1) | **Post** /eapi/v1/block/order/create | New Block Trade Order (TRADE)
[**OptionsCreateBlockOrderExecuteV1**](MarketMakerBlockTradeAPI.md#OptionsCreateBlockOrderExecuteV1) | **Post** /eapi/v1/block/order/execute | Accept Block Trade Order (TRADE)
[**OptionsDeleteBlockOrderCreateV1**](MarketMakerBlockTradeAPI.md#OptionsDeleteBlockOrderCreateV1) | **Delete** /eapi/v1/block/order/create | Cancel Block Trade Order (TRADE)
[**OptionsGetBlockOrderExecuteV1**](MarketMakerBlockTradeAPI.md#OptionsGetBlockOrderExecuteV1) | **Get** /eapi/v1/block/order/execute | Query Block Trade Details (USER_DATA)
[**OptionsGetBlockOrderOrdersV1**](MarketMakerBlockTradeAPI.md#OptionsGetBlockOrderOrdersV1) | **Get** /eapi/v1/block/order/orders | Query Block Trade Order (TRADE)
[**OptionsGetBlockUserTradesV1**](MarketMakerBlockTradeAPI.md#OptionsGetBlockUserTradesV1) | **Get** /eapi/v1/block/user-trades | Account Block Trade List (USER_DATA)
[**OptionsUpdateBlockOrderCreateV1**](MarketMakerBlockTradeAPI.md#OptionsUpdateBlockOrderCreateV1) | **Put** /eapi/v1/block/order/create | Extend Block Trade Order (TRADE)



## OptionsCreateBlockOrderCreateV1

> OptionsCreateBlockOrderCreateV1Resp OptionsCreateBlockOrderCreateV1(ctx).Legs(legs).Liquidity(liquidity).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

New Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	legs := []string{"Inner_example"} // []string | 
	liquidity := "liquidity_example" // string |  (default to "")
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int32(56) // int32 | 
	recvWindow := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.OptionsCreateBlockOrderCreateV1(context.Background()).Legs(legs).Liquidity(liquidity).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.OptionsCreateBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateBlockOrderCreateV1`: OptionsCreateBlockOrderCreateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.OptionsCreateBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateBlockOrderCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **legs** | **[]string** |  | 
 **liquidity** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int32** |  | 
 **recvWindow** | **int32** |  | 

### Return type

[**OptionsCreateBlockOrderCreateV1Resp**](OptionsCreateBlockOrderCreateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateBlockOrderExecuteV1

> OptionsCreateBlockOrderExecuteV1Resp OptionsCreateBlockOrderExecuteV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.OptionsCreateBlockOrderExecuteV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.OptionsCreateBlockOrderExecuteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateBlockOrderExecuteV1`: OptionsCreateBlockOrderExecuteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.OptionsCreateBlockOrderExecuteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateBlockOrderExecuteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsCreateBlockOrderExecuteV1Resp**](OptionsCreateBlockOrderExecuteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteBlockOrderCreateV1

> map[string]interface{} OptionsDeleteBlockOrderCreateV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int32(56) // int32 | 
	recvWindow := int32(56) // int32 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.OptionsDeleteBlockOrderCreateV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.OptionsDeleteBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteBlockOrderCreateV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.OptionsDeleteBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteBlockOrderCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int32** |  | 
 **recvWindow** | **int32** | The value cannot be greater than 60000 | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetBlockOrderExecuteV1

> OptionsGetBlockOrderExecuteV1Resp OptionsGetBlockOrderExecuteV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Block Trade Details (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.OptionsGetBlockOrderExecuteV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.OptionsGetBlockOrderExecuteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockOrderExecuteV1`: OptionsGetBlockOrderExecuteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.OptionsGetBlockOrderExecuteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBlockOrderExecuteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**OptionsGetBlockOrderExecuteV1Resp**](OptionsGetBlockOrderExecuteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetBlockOrderOrdersV1

> []OptionsGetBlockOrderOrdersV1RespItem OptionsGetBlockOrderOrdersV1(ctx).Timestamp(timestamp).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Query Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string | If specified, returns the specific block trade associated with the blockOrderMatchingKey (optional) (default to "")
	endTime := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.OptionsGetBlockOrderOrdersV1(context.Background()).Timestamp(timestamp).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.OptionsGetBlockOrderOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockOrderOrdersV1`: []OptionsGetBlockOrderOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.OptionsGetBlockOrderOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBlockOrderOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **blockOrderMatchingKey** | **string** | If specified, returns the specific block trade associated with the blockOrderMatchingKey | [default to &quot;&quot;]
 **endTime** | **int64** |  | 
 **startTime** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]OptionsGetBlockOrderOrdersV1RespItem**](OptionsGetBlockOrderOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetBlockUserTradesV1

> []OptionsGetBlockUserTradesV1RespItem OptionsGetBlockUserTradesV1(ctx).Timestamp(timestamp).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Account Block Trade List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	endTime := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.OptionsGetBlockUserTradesV1(context.Background()).Timestamp(timestamp).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.OptionsGetBlockUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBlockUserTradesV1`: []OptionsGetBlockUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.OptionsGetBlockUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBlockUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **endTime** | **int64** |  | 
 **startTime** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]OptionsGetBlockUserTradesV1RespItem**](OptionsGetBlockUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsUpdateBlockOrderCreateV1

> OptionsUpdateBlockOrderCreateV1Resp OptionsUpdateBlockOrderCreateV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Extend Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int32(56) // int32 | 
	recvWindow := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.OptionsUpdateBlockOrderCreateV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.OptionsUpdateBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsUpdateBlockOrderCreateV1`: OptionsUpdateBlockOrderCreateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.OptionsUpdateBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsUpdateBlockOrderCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int32** |  | 
 **recvWindow** | **int32** |  | 

### Return type

[**OptionsUpdateBlockOrderCreateV1Resp**](OptionsUpdateBlockOrderCreateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

