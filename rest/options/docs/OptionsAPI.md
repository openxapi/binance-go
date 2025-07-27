# \OptionsAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchOrdersV1**](OptionsAPI.md#CreateBatchOrdersV1) | **Post** /eapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**CreateBlockOrderExecuteV1**](OptionsAPI.md#CreateBlockOrderExecuteV1) | **Post** /eapi/v1/block/order/execute | Accept Block Trade Order (TRADE)
[**CreateCountdownCancelAllHeartBeatV1**](OptionsAPI.md#CreateCountdownCancelAllHeartBeatV1) | **Post** /eapi/v1/countdownCancelAllHeartBeat | Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)
[**CreateCountdownCancelAllV1**](OptionsAPI.md#CreateCountdownCancelAllV1) | **Post** /eapi/v1/countdownCancelAll | Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**CreateListenKeyV1**](OptionsAPI.md#CreateListenKeyV1) | **Post** /eapi/v1/listenKey | Start User Data Stream (USER_STREAM)
[**CreateMmpResetV1**](OptionsAPI.md#CreateMmpResetV1) | **Post** /eapi/v1/mmpReset | Reset Market Maker Protection Config (TRADE)
[**CreateMmpSetV1**](OptionsAPI.md#CreateMmpSetV1) | **Post** /eapi/v1/mmpSet | Set Market Maker Protection Config (TRADE)
[**CreateOrderV1**](OptionsAPI.md#CreateOrderV1) | **Post** /eapi/v1/order | New Order (TRADE)
[**DeleteAllOpenOrdersByUnderlyingV1**](OptionsAPI.md#DeleteAllOpenOrdersByUnderlyingV1) | **Delete** /eapi/v1/allOpenOrdersByUnderlying | Cancel All Option Orders By Underlying (TRADE)
[**DeleteAllOpenOrdersV1**](OptionsAPI.md#DeleteAllOpenOrdersV1) | **Delete** /eapi/v1/allOpenOrders | Cancel all Option orders on specific symbol (TRADE)
[**DeleteBatchOrdersV1**](OptionsAPI.md#DeleteBatchOrdersV1) | **Delete** /eapi/v1/batchOrders | Cancel Multiple Option Orders (TRADE)
[**DeleteListenKeyV1**](OptionsAPI.md#DeleteListenKeyV1) | **Delete** /eapi/v1/listenKey | Close User Data Stream (USER_STREAM)
[**DeleteOrderV1**](OptionsAPI.md#DeleteOrderV1) | **Delete** /eapi/v1/order | Cancel Option Order (TRADE)
[**GetAccountV1**](OptionsAPI.md#GetAccountV1) | **Get** /eapi/v1/account | Option Account Information(TRADE)
[**GetBillV1**](OptionsAPI.md#GetBillV1) | **Get** /eapi/v1/bill | Account Funding Flow (USER_DATA)
[**GetBlockOrderExecuteV1**](OptionsAPI.md#GetBlockOrderExecuteV1) | **Get** /eapi/v1/block/order/execute | Query Block Trade Details (USER_DATA)
[**GetBlockOrderOrdersV1**](OptionsAPI.md#GetBlockOrderOrdersV1) | **Get** /eapi/v1/block/order/orders | Query Block Trade Order (TRADE)
[**GetBlockTradesV1**](OptionsAPI.md#GetBlockTradesV1) | **Get** /eapi/v1/blockTrades | Recent Block Trades List
[**GetBlockUserTradesV1**](OptionsAPI.md#GetBlockUserTradesV1) | **Get** /eapi/v1/block/user-trades | Account Block Trade List (USER_DATA)
[**GetCountdownCancelAllV1**](OptionsAPI.md#GetCountdownCancelAllV1) | **Get** /eapi/v1/countdownCancelAll | Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**GetDepthV1**](OptionsAPI.md#GetDepthV1) | **Get** /eapi/v1/depth | Order Book
[**GetExchangeInfoV1**](OptionsAPI.md#GetExchangeInfoV1) | **Get** /eapi/v1/exchangeInfo | Exchange Information
[**GetExerciseHistoryV1**](OptionsAPI.md#GetExerciseHistoryV1) | **Get** /eapi/v1/exerciseHistory | Historical Exercise Records
[**GetExerciseRecordV1**](OptionsAPI.md#GetExerciseRecordV1) | **Get** /eapi/v1/exerciseRecord | User Exercise Record (USER_DATA)
[**GetHistoricalTradesV1**](OptionsAPI.md#GetHistoricalTradesV1) | **Get** /eapi/v1/historicalTrades | Old Trades Lookup (MARKET_DATA)
[**GetHistoryOrdersV1**](OptionsAPI.md#GetHistoryOrdersV1) | **Get** /eapi/v1/historyOrders | Query Option Order History (TRADE)
[**GetIncomeAsynIdV1**](OptionsAPI.md#GetIncomeAsynIdV1) | **Get** /eapi/v1/income/asyn/id | Get Option Transaction History Download Link by Id (USER_DATA)
[**GetIncomeAsynV1**](OptionsAPI.md#GetIncomeAsynV1) | **Get** /eapi/v1/income/asyn | Get Download Id For Option Transaction History (USER_DATA)
[**GetIndexV1**](OptionsAPI.md#GetIndexV1) | **Get** /eapi/v1/index | Symbol Price Ticker
[**GetKlinesV1**](OptionsAPI.md#GetKlinesV1) | **Get** /eapi/v1/klines | Kline/Candlestick Data
[**GetMarginAccountV1**](OptionsAPI.md#GetMarginAccountV1) | **Get** /eapi/v1/marginAccount | Option Margin Account Information (USER_DATA)
[**GetMarkV1**](OptionsAPI.md#GetMarkV1) | **Get** /eapi/v1/mark | Option Mark Price
[**GetMmpV1**](OptionsAPI.md#GetMmpV1) | **Get** /eapi/v1/mmp | Get Market Maker Protection Config (TRADE)
[**GetOpenInterestV1**](OptionsAPI.md#GetOpenInterestV1) | **Get** /eapi/v1/openInterest | Open Interest
[**GetOpenOrdersV1**](OptionsAPI.md#GetOpenOrdersV1) | **Get** /eapi/v1/openOrders | Query Current Open Option Orders (USER_DATA)
[**GetOrderV1**](OptionsAPI.md#GetOrderV1) | **Get** /eapi/v1/order | Query Single Order (TRADE)
[**GetPingV1**](OptionsAPI.md#GetPingV1) | **Get** /eapi/v1/ping | Test Connectivity
[**GetPositionV1**](OptionsAPI.md#GetPositionV1) | **Get** /eapi/v1/position | Option Position Information (USER_DATA)
[**GetTickerV1**](OptionsAPI.md#GetTickerV1) | **Get** /eapi/v1/ticker | 24hr Ticker Price Change Statistics
[**GetTimeV1**](OptionsAPI.md#GetTimeV1) | **Get** /eapi/v1/time | Check Server Time
[**GetTradesV1**](OptionsAPI.md#GetTradesV1) | **Get** /eapi/v1/trades | Recent Trades List
[**GetUserTradesV1**](OptionsAPI.md#GetUserTradesV1) | **Get** /eapi/v1/userTrades | Account Trade List (USER_DATA)
[**UpdateBlockOrderCreateV1**](OptionsAPI.md#UpdateBlockOrderCreateV1) | **Put** /eapi/v1/block/order/create | Extend Block Trade Order (TRADE)
[**UpdateListenKeyV1**](OptionsAPI.md#UpdateListenKeyV1) | **Put** /eapi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)



## CreateBatchOrdersV1

> []OptionsCreateBatchOrdersV1RespInner CreateBatchOrdersV1(ctx).Orders(orders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Place Multiple Orders(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	orders := "orders_example" // string | JSON string containing array of order objects. Max 20 orders.
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateBatchOrdersV1(context.Background()).Orders(orders).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchOrdersV1`: []OptionsCreateBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orders** | **string** | JSON string containing array of order objects. Max 20 orders. | 
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


## CreateBlockOrderExecuteV1

> CreateBlockOrderExecuteV1Resp CreateBlockOrderExecuteV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateBlockOrderExecuteV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateBlockOrderExecuteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockOrderExecuteV1`: CreateBlockOrderExecuteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateBlockOrderExecuteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockOrderExecuteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateBlockOrderExecuteV1Resp**](CreateBlockOrderExecuteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCountdownCancelAllHeartBeatV1

> CreateCountdownCancelAllHeartBeatV1Resp CreateCountdownCancelAllHeartBeatV1(ctx).Timestamp(timestamp).Underlyings(underlyings).RecvWindow(recvWindow).Execute()

Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	underlyings := "underlyings_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateCountdownCancelAllHeartBeatV1(context.Background()).Timestamp(timestamp).Underlyings(underlyings).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateCountdownCancelAllHeartBeatV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCountdownCancelAllHeartBeatV1`: CreateCountdownCancelAllHeartBeatV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateCountdownCancelAllHeartBeatV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCountdownCancelAllHeartBeatV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlyings** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateCountdownCancelAllHeartBeatV1Resp**](CreateCountdownCancelAllHeartBeatV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCountdownCancelAllV1

> CreateCountdownCancelAllV1Resp CreateCountdownCancelAllV1(ctx).CountdownTime(countdownTime).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	countdownTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateCountdownCancelAllV1(context.Background()).CountdownTime(countdownTime).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCountdownCancelAllV1`: CreateCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countdownTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateCountdownCancelAllV1Resp**](CreateCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListenKeyV1

> CreateListenKeyV1Resp CreateListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateListenKeyV1`: CreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListenKeyV1Request struct via the builder pattern


### Return type

[**CreateListenKeyV1Resp**](CreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMmpResetV1

> CreateMmpResetV1Resp CreateMmpResetV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Underlying(underlying).Execute()

Reset Market Maker Protection Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateMmpResetV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Underlying(underlying).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateMmpResetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMmpResetV1`: CreateMmpResetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateMmpResetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMmpResetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateMmpResetV1Resp**](CreateMmpResetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMmpSetV1

> CreateMmpSetV1Resp CreateMmpSetV1(ctx).Timestamp(timestamp).DeltaLimit(deltaLimit).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).RecvWindow(recvWindow).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).Execute()

Set Market Maker Protection Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	deltaLimit := "deltaLimit_example" // string |  (optional) (default to "")
	frozenTimeInMilliseconds := int64(789) // int64 |  (optional)
	qtyLimit := "qtyLimit_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	windowTimeInMilliseconds := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateMmpSetV1(context.Background()).Timestamp(timestamp).DeltaLimit(deltaLimit).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).RecvWindow(recvWindow).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateMmpSetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMmpSetV1`: CreateMmpSetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateMmpSetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMmpSetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **deltaLimit** | **string** |  | [default to &quot;&quot;]
 **frozenTimeInMilliseconds** | **int64** |  | 
 **qtyLimit** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **windowTimeInMilliseconds** | **int64** |  | 

### Return type

[**CreateMmpSetV1Resp**](CreateMmpSetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrderV1

> OptionsCreateOrderV1Resp CreateOrderV1(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ClientOrderId(clientOrderId).IsMmp(isMmp).NewOrderRespType(newOrderRespType).PostOnly(postOnly).Price(price).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()

New Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
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
	resp, r, err := apiClient.OptionsAPI.CreateOrderV1(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).ClientOrderId(clientOrderId).IsMmp(isMmp).NewOrderRespType(newOrderRespType).PostOnly(postOnly).Price(price).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderV1`: OptionsCreateOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderV1Request struct via the builder pattern


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


## DeleteAllOpenOrdersByUnderlyingV1

> DeleteAllOpenOrdersByUnderlyingV1Resp DeleteAllOpenOrdersByUnderlyingV1(ctx).Underlying(underlying).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All Option Orders By Underlying (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.DeleteAllOpenOrdersByUnderlyingV1(context.Background()).Underlying(underlying).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.DeleteAllOpenOrdersByUnderlyingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllOpenOrdersByUnderlyingV1`: DeleteAllOpenOrdersByUnderlyingV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.DeleteAllOpenOrdersByUnderlyingV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllOpenOrdersByUnderlyingV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Option underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**DeleteAllOpenOrdersByUnderlyingV1Resp**](DeleteAllOpenOrdersByUnderlyingV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllOpenOrdersV1

> DeleteAllOpenOrdersV1Resp DeleteAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel all Option orders on specific symbol (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.DeleteAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.DeleteAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllOpenOrdersV1`: DeleteAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.DeleteAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**DeleteAllOpenOrdersV1Resp**](DeleteAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBatchOrdersV1

> []OptionsDeleteBatchOrdersV1RespInner DeleteBatchOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()

Cancel Multiple Option Orders (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	orderIds := "orderIds_example" // string | Order ID, e.g [4611875134427365377,4611875134427365378] (optional)
	clientOrderIds := "clientOrderIds_example" // string | User-defined order ID, e.g [&#34;my_id_1&#34;,&#34;my_id_2&#34;] (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.DeleteBatchOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderIds(orderIds).ClientOrderIds(clientOrderIds).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.DeleteBatchOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBatchOrdersV1`: []OptionsDeleteBatchOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.DeleteBatchOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderIds** | **string** | Order ID, e.g [4611875134427365377,4611875134427365378] | 
 **clientOrderIds** | **string** | User-defined order ID, e.g [&amp;#34;my_id_1&amp;#34;,&amp;#34;my_id_2&amp;#34;] | 
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


## DeleteListenKeyV1

> map[string]interface{} DeleteListenKeyV1(ctx).Execute()

Close User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.DeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.DeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.DeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListenKeyV1Request struct via the builder pattern


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


## DeleteOrderV1

> DeleteOrderV1Resp DeleteOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Cancel Option Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Order ID, e.g 4611875134427365377 (optional)
	clientOrderId := "clientOrderId_example" // string | User-defined order ID, e.g 10000 (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.DeleteOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.DeleteOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrderV1`: DeleteOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.DeleteOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Order ID, e.g 4611875134427365377 | 
 **clientOrderId** | **string** | User-defined order ID, e.g 10000 | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**DeleteOrderV1Resp**](DeleteOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountV1

> GetAccountV1Resp GetAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Option Account Information(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountV1`: GetAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAccountV1Resp**](GetAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillV1

> []GetBillV1RespItem GetBillV1(ctx).Currency(currency).Timestamp(timestamp).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Funding Flow (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	currency := "currency_example" // string | Asset type, only support USDT  as of now (default to "")
	timestamp := int64(789) // int64 | 
	recordId := int64(789) // int64 | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 (optional)
	startTime := int64(789) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(789) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int32(56) // int32 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetBillV1(context.Background()).Currency(currency).Timestamp(timestamp).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetBillV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillV1`: []GetBillV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetBillV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Asset type, only support USDT  as of now | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recordId** | **int64** | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int32** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetBillV1RespItem**](GetBillV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockOrderExecuteV1

> GetBlockOrderExecuteV1Resp GetBlockOrderExecuteV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Block Trade Details (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetBlockOrderExecuteV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetBlockOrderExecuteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockOrderExecuteV1`: GetBlockOrderExecuteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetBlockOrderExecuteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockOrderExecuteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetBlockOrderExecuteV1Resp**](GetBlockOrderExecuteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockOrderOrdersV1

> []GetBlockOrderOrdersV1RespItem GetBlockOrderOrdersV1(ctx).Timestamp(timestamp).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Query Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
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
	resp, r, err := apiClient.OptionsAPI.GetBlockOrderOrdersV1(context.Background()).Timestamp(timestamp).BlockOrderMatchingKey(blockOrderMatchingKey).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetBlockOrderOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockOrderOrdersV1`: []GetBlockOrderOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetBlockOrderOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockOrderOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **blockOrderMatchingKey** | **string** | If specified, returns the specific block trade associated with the blockOrderMatchingKey | [default to &quot;&quot;]
 **endTime** | **int64** |  | 
 **startTime** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]GetBlockOrderOrdersV1RespItem**](GetBlockOrderOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockTradesV1

> []GetBlockTradesV1RespItem GetBlockTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Block Trades List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g. BTC-200730-9000-C (optional) (default to "")
	limit := int32(56) // int32 | Number of records; Default: 100 and Max: 500 (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetBlockTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetBlockTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockTradesV1`: []GetBlockTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetBlockTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g. BTC-200730-9000-C | [default to &quot;&quot;]
 **limit** | **int32** | Number of records; Default: 100 and Max: 500 | [default to 100]

### Return type

[**[]GetBlockTradesV1RespItem**](GetBlockTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockUserTradesV1

> []GetBlockUserTradesV1RespItem GetBlockUserTradesV1(ctx).Timestamp(timestamp).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()

Account Block Trade List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	endTime := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetBlockUserTradesV1(context.Background()).Timestamp(timestamp).EndTime(endTime).StartTime(startTime).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetBlockUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockUserTradesV1`: []GetBlockUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetBlockUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **endTime** | **int64** |  | 
 **startTime** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]GetBlockUserTradesV1RespItem**](GetBlockUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountdownCancelAllV1

> GetCountdownCancelAllV1Resp GetCountdownCancelAllV1(ctx).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetCountdownCancelAllV1(context.Background()).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountdownCancelAllV1`: GetCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlying** | **string** | Option underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetCountdownCancelAllV1Resp**](GetCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepthV1

> GetDepthV1Resp GetDepthV1(ctx).Symbol(symbol).Limit(limit).Execute()

Order Book



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	limit := int32(56) // int32 | Default:100 Max:1000.Optional value:[10, 20, 50, 100, 500, 1000] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetDepthV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetDepthV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepthV1`: GetDepthV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetDepthV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepthV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **limit** | **int32** | Default:100 Max:1000.Optional value:[10, 20, 50, 100, 500, 1000] | 

### Return type

[**GetDepthV1Resp**](GetDepthV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeInfoV1

> OptionsGetExchangeInfoV1Resp GetExchangeInfoV1(ctx).Execute()

Exchange Information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetExchangeInfoV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeInfoV1`: OptionsGetExchangeInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeInfoV1Request struct via the builder pattern


### Return type

[**OptionsGetExchangeInfoV1Resp**](OptionsGetExchangeInfoV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExerciseHistoryV1

> []GetExerciseHistoryV1RespItem GetExerciseHistoryV1(ctx).Underlying(underlying).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Historical Exercise Records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	underlying := "underlying_example" // string | Underlying index like BTCUSDT (optional) (default to "")
	startTime := int64(789) // int64 | Start Time (optional)
	endTime := int64(789) // int64 | End Time (optional)
	limit := int32(56) // int32 | Number of records Default:100 Max:100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetExerciseHistoryV1(context.Background()).Underlying(underlying).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetExerciseHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExerciseHistoryV1`: []GetExerciseHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetExerciseHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExerciseHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Underlying index like BTCUSDT | [default to &quot;&quot;]
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time | 
 **limit** | **int32** | Number of records Default:100 Max:100 | 

### Return type

[**[]GetExerciseHistoryV1RespItem**](GetExerciseHistoryV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExerciseRecordV1

> []GetExerciseRecordV1RespItem GetExerciseRecordV1(ctx).Timestamp(timestamp).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

User Exercise Record (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
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
	resp, r, err := apiClient.OptionsAPI.GetExerciseRecordV1(context.Background()).Timestamp(timestamp).Symbol(symbol).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetExerciseRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExerciseRecordV1`: []GetExerciseRecordV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetExerciseRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExerciseRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **startTime** | **int64** | startTime | 
 **endTime** | **int64** | endTime | 
 **limit** | **int32** | default 1000, max 1000 | [default to 1000]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetExerciseRecordV1RespItem**](GetExerciseRecordV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalTradesV1

> []GetHistoricalTradesV1RespItem GetHistoricalTradesV1(ctx).Symbol(symbol).FromId(fromId).Limit(limit).Execute()

Old Trades Lookup (MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	fromId := int64(789) // int64 | The UniqueId ID from which to return. The latest deal record is returned by default (optional)
	limit := int32(56) // int32 | Number of records Default:100 Max:500 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetHistoricalTradesV1(context.Background()).Symbol(symbol).FromId(fromId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetHistoricalTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricalTradesV1`: []GetHistoricalTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetHistoricalTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **fromId** | **int64** | The UniqueId ID from which to return. The latest deal record is returned by default | 
 **limit** | **int32** | Number of records Default:100 Max:500 | 

### Return type

[**[]GetHistoricalTradesV1RespItem**](GetHistoricalTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoryOrdersV1

> []GetHistoryOrdersV1RespItem GetHistoryOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Option Order History (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
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
	resp, r, err := apiClient.OptionsAPI.GetHistoryOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetHistoryOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoryOrdersV1`: []GetHistoryOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetHistoryOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryOrdersV1Request struct via the builder pattern


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

[**[]GetHistoryOrdersV1RespItem**](GetHistoryOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomeAsynIdV1

> GetIncomeAsynIdV1Resp GetIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Option Transaction History Download Link by Id (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomeAsynIdV1`: GetIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetIncomeAsynIdV1Resp**](GetIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomeAsynV1

> GetIncomeAsynV1Resp GetIncomeAsynV1(ctx).Execute()

Get Download Id For Option Transaction History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetIncomeAsynV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomeAsynV1`: GetIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomeAsynV1Request struct via the builder pattern


### Return type

[**GetIncomeAsynV1Resp**](GetIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexV1

> GetIndexV1Resp GetIndexV1(ctx).Underlying(underlying).Execute()

Symbol Price Ticker



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	underlying := "underlying_example" // string | Spot pair（Option contract underlying asset, e.g BTCUSDT) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetIndexV1(context.Background()).Underlying(underlying).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexV1`: GetIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlying** | **string** | Spot pair（Option contract underlying asset, e.g BTCUSDT) | [default to &quot;&quot;]

### Return type

[**GetIndexV1Resp**](GetIndexV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKlinesV1

> []GetKlinesV1RespItem GetKlinesV1(ctx).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Kline/Candlestick Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	interval := "interval_example" // string | Time interval (default to "")
	startTime := int64(789) // int64 | Start Time  1592317127349 (optional)
	endTime := int64(789) // int64 | End Time (optional)
	limit := int32(56) // int32 | Number of records Default:500 Max:1500 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetKlinesV1(context.Background()).Symbol(symbol).Interval(interval).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetKlinesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKlinesV1`: []GetKlinesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetKlinesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKlinesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **interval** | **string** | Time interval | [default to &quot;&quot;]
 **startTime** | **int64** | Start Time  1592317127349 | 
 **endTime** | **int64** | End Time | 
 **limit** | **int32** | Number of records Default:500 Max:1500 | 

### Return type

[**[]GetKlinesV1RespItem**](GetKlinesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginAccountV1

> GetMarginAccountV1Resp GetMarginAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Option Margin Account Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetMarginAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginAccountV1`: GetMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMarginAccountV1Resp**](GetMarginAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarkV1

> []GetMarkV1RespItem GetMarkV1(ctx).Symbol(symbol).Execute()

Option Mark Price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetMarkV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetMarkV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarkV1`: []GetMarkV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetMarkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]

### Return type

[**[]GetMarkV1RespItem**](GetMarkV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMmpV1

> GetMmpV1Resp GetMmpV1(ctx).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Market Maker Protection Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetMmpV1(context.Background()).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetMmpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMmpV1`: GetMmpV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetMmpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMmpV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlying** | **string** | underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetMmpV1Resp**](GetMmpV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenInterestV1

> []GetOpenInterestV1RespItem GetOpenInterestV1(ctx).UnderlyingAsset(underlyingAsset).Expiration(expiration).Execute()

Open Interest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	underlyingAsset := "underlyingAsset_example" // string | underlying asset, e.g ETH/BTC (default to "")
	expiration := "expiration_example" // string | expiration date, e.g 221225 (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetOpenInterestV1(context.Background()).UnderlyingAsset(underlyingAsset).Expiration(expiration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetOpenInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenInterestV1`: []GetOpenInterestV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetOpenInterestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenInterestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **underlyingAsset** | **string** | underlying asset, e.g ETH/BTC | [default to &quot;&quot;]
 **expiration** | **string** | expiration date, e.g 221225 | [default to &quot;&quot;]

### Return type

[**[]GetOpenInterestV1RespItem**](GetOpenInterestV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOpenOrdersV1

> []GetOpenOrdersV1RespItem GetOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Current Open Option Orders (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
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
	resp, r, err := apiClient.OptionsAPI.GetOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenOrdersV1`: []GetOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenOrdersV1Request struct via the builder pattern


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

[**[]GetOpenOrdersV1RespItem**](GetOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderV1

> GetOrderV1Resp GetOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()

Query Single Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 | Order id (optional)
	clientOrderId := "clientOrderId_example" // string | User-defined order ID cannot be repeated in pending orders (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).ClientOrderId(clientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderV1`: GetOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | Order id | 
 **clientOrderId** | **string** | User-defined order ID cannot be repeated in pending orders | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetOrderV1Resp**](GetOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPingV1

> map[string]interface{} GetPingV1(ctx).Execute()

Test Connectivity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPingV1Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPositionV1

> []GetPositionV1RespItem GetPositionV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Option Position Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetPositionV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetPositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPositionV1`: []GetPositionV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetPositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetPositionV1RespItem**](GetPositionV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTickerV1

> []GetTickerV1RespItem GetTickerV1(ctx).Symbol(symbol).Execute()

24hr Ticker Price Change Statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetTickerV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetTickerV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTickerV1`: []GetTickerV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetTickerV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTickerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]

### Return type

[**[]GetTickerV1RespItem**](GetTickerV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeV1

> GetTimeV1Resp GetTimeV1(ctx).Execute()

Check Server Time



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetTimeV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetTimeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeV1`: GetTimeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetTimeV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeV1Request struct via the builder pattern


### Return type

[**GetTimeV1Resp**](GetTimeV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradesV1

> []GetTradesV1RespItem GetTradesV1(ctx).Symbol(symbol).Limit(limit).Execute()

Recent Trades List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	symbol := "symbol_example" // string | Option trading pair, e.g BTC-200730-9000-C (default to "")
	limit := int32(56) // int32 | Number of records Default:100 Max:500 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetTradesV1(context.Background()).Symbol(symbol).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradesV1`: []GetTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Option trading pair, e.g BTC-200730-9000-C | [default to &quot;&quot;]
 **limit** | **int32** | Number of records Default:100 Max:500 | 

### Return type

[**[]GetTradesV1RespItem**](GetTradesV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTradesV1

> []GetUserTradesV1RespItem GetUserTradesV1(ctx).Timestamp(timestamp).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Trade List (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
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
	resp, r, err := apiClient.OptionsAPI.GetUserTradesV1(context.Background()).Timestamp(timestamp).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTradesV1`: []GetUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTradesV1Request struct via the builder pattern


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

[**[]GetUserTradesV1RespItem**](GetUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockOrderCreateV1

> UpdateBlockOrderCreateV1Resp UpdateBlockOrderCreateV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Extend Block Trade Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {
	blockOrderMatchingKey := "blockOrderMatchingKey_example" // string |  (default to "")
	timestamp := int32(56) // int32 | 
	recvWindow := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.UpdateBlockOrderCreateV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.UpdateBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlockOrderCreateV1`: UpdateBlockOrderCreateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.UpdateBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlockOrderCreateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockOrderMatchingKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int32** |  | 
 **recvWindow** | **int32** |  | 

### Return type

[**UpdateBlockOrderCreateV1Resp**](UpdateBlockOrderCreateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateListenKeyV1

> map[string]interface{} UpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.UpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.UpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.UpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListenKeyV1Request struct via the builder pattern


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

