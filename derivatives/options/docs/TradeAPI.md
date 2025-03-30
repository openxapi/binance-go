# \TradeAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsCreateBatchOrdersV1**](TradeAPI.md#OptionsCreateBatchOrdersV1) | **Post** /eapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**OptionsCreateOrderV1**](TradeAPI.md#OptionsCreateOrderV1) | **Post** /eapi/v1/order | New Order (TRADE)
[**OptionsDeleteAllOpenOrdersByUnderlyingV1**](TradeAPI.md#OptionsDeleteAllOpenOrdersByUnderlyingV1) | **Delete** /eapi/v1/allOpenOrdersByUnderlying | Cancel All Option Orders By Underlying (TRADE)
[**OptionsDeleteAllOpenOrdersV1**](TradeAPI.md#OptionsDeleteAllOpenOrdersV1) | **Delete** /eapi/v1/allOpenOrders | Cancel all Option orders on specific symbol (TRADE)
[**OptionsDeleteBatchOrdersV1**](TradeAPI.md#OptionsDeleteBatchOrdersV1) | **Delete** /eapi/v1/batchOrders | Cancel Multiple Option Orders (TRADE)
[**OptionsDeleteOrderV1**](TradeAPI.md#OptionsDeleteOrderV1) | **Delete** /eapi/v1/order | Cancel Option Order (TRADE)
[**OptionsGetExerciseRecordV1**](TradeAPI.md#OptionsGetExerciseRecordV1) | **Get** /eapi/v1/exerciseRecord | User Exercise Record (USER_DATA)
[**OptionsGetHistoryOrdersV1**](TradeAPI.md#OptionsGetHistoryOrdersV1) | **Get** /eapi/v1/historyOrders | Query Option Order History (TRADE)
[**OptionsGetOpenOrdersV1**](TradeAPI.md#OptionsGetOpenOrdersV1) | **Get** /eapi/v1/openOrders | Query Current Open Option Orders (USER_DATA)
[**OptionsGetOrderV1**](TradeAPI.md#OptionsGetOrderV1) | **Get** /eapi/v1/order | Query Single Order (TRADE)
[**OptionsGetPositionV1**](TradeAPI.md#OptionsGetPositionV1) | **Get** /eapi/v1/position | Option Position Information (USER_DATA)
[**OptionsGetUserTradesV1**](TradeAPI.md#OptionsGetUserTradesV1) | **Get** /eapi/v1/userTrades | Account Trade List (USER_DATA)



## OptionsCreateBatchOrdersV1

> []OptionsCreateBatchOrdersV1RespInner OptionsCreateBatchOrdersV1(ctx).Orders(orders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Place Multiple Orders(TRADE)



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
	orders := []openapiclient.OptionsCreateBatchOrdersV1ReqOrdersItem{*openapiclient.NewOptionsCreateBatchOrdersV1ReqOrdersItem("Quantity_example", "Side_example", "Symbol_example", "Type_example")} // []OptionsCreateBatchOrdersV1ReqOrdersItem | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsCreateBatchOrdersV1(context.Background()).Orders(orders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsCreateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateBatchOrdersV1`: []OptionsCreateBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsCreateBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orders** | [**[]OptionsCreateBatchOrdersV1ReqOrdersItem**](OptionsCreateBatchOrdersV1ReqOrdersItem.md) |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsCreateBatchOrdersV1RespInner**](OptionsCreateBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateOrderV1

> OptionsCreateOrderV1Resp OptionsCreateOrderV1(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ClientOrderId(clientOrderId).IsMmp(isMmp).NewOrderRespType(newOrderRespType).PostOnly(postOnly).Price(price).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()

New Order (TRADE)



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
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	clientOrderId := "clientOrderId_example" // string |  (optional) (default to "")
	isMmp := true // bool |  (optional)
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	postOnly := true // bool |  (optional) (default to false)
	price := "price_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := true // bool |  (optional) (default to false)
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsCreateOrderV1(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ClientOrderId(clientOrderId).IsMmp(isMmp).NewOrderRespType(newOrderRespType).PostOnly(postOnly).Price(price).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsCreateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateOrderV1`: OptionsCreateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsCreateOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **clientOrderId** | **string** |  | [default to &quot;&quot;]
 **isMmp** | **bool** |  | 
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **postOnly** | **bool** |  | [default to false]
 **price** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **bool** |  | [default to false]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**OptionsCreateOrderV1Resp**](OptionsCreateOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteAllOpenOrdersByUnderlyingV1

> OptionsDeleteAllOpenOrdersByUnderlyingV1Resp OptionsDeleteAllOpenOrdersByUnderlyingV1(ctx).Underlying(underlying).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All Option Orders By Underlying (TRADE)



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
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsDeleteAllOpenOrdersByUnderlyingV1(context.Background()).Underlying(underlying).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsDeleteAllOpenOrdersByUnderlyingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteAllOpenOrdersByUnderlyingV1`: OptionsDeleteAllOpenOrdersByUnderlyingV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsDeleteAllOpenOrdersByUnderlyingV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteAllOpenOrdersByUnderlyingV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Option underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsDeleteAllOpenOrdersByUnderlyingV1Resp**](OptionsDeleteAllOpenOrdersByUnderlyingV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteAllOpenOrdersV1

> OptionsDeleteAllOpenOrdersV1Resp OptionsDeleteAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel all Option orders on specific symbol (TRADE)



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
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsDeleteAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsDeleteAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteAllOpenOrdersV1`: OptionsDeleteAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsDeleteAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsDeleteAllOpenOrdersV1Resp**](OptionsDeleteAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteBatchOrdersV1

> []OptionsDeleteBatchOrdersV1RespInner OptionsDeleteBatchOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()

Cancel Multiple Option Orders (TRADE)



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
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	orderIds := []int64{int64(123)} // []int64 | Order ID, e.g [4611875134427365377,4611875134427365378] (optional)
	clientOrderIds := []string{"Inner_example"} // []string | User-defined order ID, e.g [&#34;my_id_1&#34;,&#34;my_id_2&#34;] (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsDeleteBatchOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsDeleteBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteBatchOrdersV1`: []OptionsDeleteBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsDeleteBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderIds** | **[]int64** | Order ID, e.g [4611875134427365377,4611875134427365378] | 
 **clientOrderIds** | **[]string** | User-defined order ID, e.g [&amp;#34;my_id_1&amp;#34;,&amp;#34;my_id_2&amp;#34;] | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsDeleteBatchOrdersV1RespInner**](OptionsDeleteBatchOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsDeleteOrderV1

> OptionsDeleteOrderV1Resp OptionsDeleteOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Cancel Option Order (TRADE)



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
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Order ID, e.g 4611875134427365377 (optional)
	clientOrderId := "clientOrderId_example" // string | User-defined order ID, e.g 10000 (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsDeleteOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsDeleteOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsDeleteOrderV1`: OptionsDeleteOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsDeleteOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsDeleteOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Order ID, e.g 4611875134427365377 | 
 **clientOrderId** | **string** | User-defined order ID, e.g 10000 | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsDeleteOrderV1Resp**](OptionsDeleteOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetExerciseRecordV1

> []OptionsGetExerciseRecordV1RespItem OptionsGetExerciseRecordV1(ctx).Timestamp(timestamp).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

User Exercise Record (USER_DATA)



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
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")
	startTime := int64(789) // int64 | startTime (optional)
	endTime := int64(789) // int64 | endTime (optional)
	limit := int32(56) // int32 | default 1000, max 1000 (optional) (default to 1000)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsGetExerciseRecordV1(context.Background()).Timestamp(timestamp).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsGetExerciseRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetExerciseRecordV1`: []OptionsGetExerciseRecordV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsGetExerciseRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetExerciseRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **startTime** | **int64** | startTime | 
 **endTime** | **int64** | endTime | 
 **limit** | **int32** | default 1000, max 1000 | [default to 1000]
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetExerciseRecordV1RespItem**](OptionsGetExerciseRecordV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetHistoryOrdersV1

> []OptionsGetHistoryOrdersV1RespItem OptionsGetHistoryOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Option Order History (TRADE)



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
	symbol := "symbol_example" // string | Option trading pair (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Returns the orderId and subsequent orders, the most recent order is returned by default (optional)
	startTime := int64(789) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(789) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int32(56) // int32 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsGetHistoryOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsGetHistoryOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetHistoryOrdersV1`: []OptionsGetHistoryOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsGetHistoryOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetHistoryOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Returns the orderId and subsequent orders, the most recent order is returned by default | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int32** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetHistoryOrdersV1RespItem**](OptionsGetHistoryOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetOpenOrdersV1

> []OptionsGetOpenOrdersV1RespItem OptionsGetOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Current Open Option Orders (USER_DATA)



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
	symbol := "symbol_example" // string | return all orders if don&#39;t pass, Option trading pair, e.g BTC-200730-9000-C, (optional) (default to "")
	orderId := int64(789) // int64 | Returns the orderId and subsequent orders, the most recent order is returned by default (optional)
	startTime := int64(789) // int64 | Start Time (optional)
	endTime := int64(789) // int64 | End Time (optional)
	limit := int32(56) // int32 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsGetOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsGetOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetOpenOrdersV1`: []OptionsGetOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsGetOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | return all orders if don&amp;#39;t pass, Option trading pair, e.g BTC-200730-9000-C, | [default to &quot;&quot;]
 **orderId** | **int64** | Returns the orderId and subsequent orders, the most recent order is returned by default | 
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time | 
 **limit** | **int32** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetOpenOrdersV1RespItem**](OptionsGetOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetOrderV1

> OptionsGetOrderV1Resp OptionsGetOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Query Single Order (TRADE)



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
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Order id (optional)
	clientOrderId := "clientOrderId_example" // string | User-defined order ID cannot be repeated in pending orders (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsGetOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsGetOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetOrderV1`: OptionsGetOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsGetOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Order id | 
 **clientOrderId** | **string** | User-defined order ID cannot be repeated in pending orders | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetOrderV1Resp**](OptionsGetOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetPositionV1

> []OptionsGetPositionV1RespItem OptionsGetPositionV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Option Position Information (USER_DATA)



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
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsGetPositionV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsGetPositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetPositionV1`: []OptionsGetPositionV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsGetPositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetPositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetPositionV1RespItem**](OptionsGetPositionV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetUserTradesV1

> []OptionsGetUserTradesV1RespItem OptionsGetUserTradesV1(ctx).Timestamp(timestamp).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)



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
	symbol := "symbol_example" // string | Option symbol, e.g BTC-200730-9000-C (optional) (default to "")
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376 (optional)
	startTime := int64(789) // int64 | Start time, e.g 1593511200000 (optional)
	endTime := int64(789) // int64 | End time, e.g 1593512200000 (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.OptionsGetUserTradesV1(context.Background()).Timestamp(timestamp).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.OptionsGetUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetUserTradesV1`: []OptionsGetUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.OptionsGetUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option symbol, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades, e.g 4611875134427365376 | 
 **startTime** | **int64** | Start time, e.g 1593511200000 | 
 **endTime** | **int64** | End time, e.g 1593512200000 | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetUserTradesV1RespItem**](OptionsGetUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

