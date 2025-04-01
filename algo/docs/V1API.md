# \V1API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlgoCreateAlgoFuturesNewOrderTwapV1**](V1API.md#AlgoCreateAlgoFuturesNewOrderTwapV1) | **Post** /sapi/v1/algo/futures/newOrderTwap | Time-Weighted Average Price(Twap) New Order(TRADE)
[**AlgoCreateAlgoFuturesNewOrderVpV1**](V1API.md#AlgoCreateAlgoFuturesNewOrderVpV1) | **Post** /sapi/v1/algo/futures/newOrderVp | Volume Participation(VP) New Order (TRADE)
[**AlgoCreateAlgoSpotNewOrderTwapV1**](V1API.md#AlgoCreateAlgoSpotNewOrderTwapV1) | **Post** /sapi/v1/algo/spot/newOrderTwap | Time-Weighted Average Price(Twap) New Order(TRADE)
[**AlgoDeleteAlgoFuturesOrderV1**](V1API.md#AlgoDeleteAlgoFuturesOrderV1) | **Delete** /sapi/v1/algo/futures/order | Cancel Algo Order(TRADE)
[**AlgoDeleteAlgoSpotOrderV1**](V1API.md#AlgoDeleteAlgoSpotOrderV1) | **Delete** /sapi/v1/algo/spot/order | Cancel Algo Order(TRADE)
[**AlgoGetAlgoFuturesHistoricalOrdersV1**](V1API.md#AlgoGetAlgoFuturesHistoricalOrdersV1) | **Get** /sapi/v1/algo/futures/historicalOrders | Query Historical Algo Orders(USER_DATA)
[**AlgoGetAlgoFuturesOpenOrdersV1**](V1API.md#AlgoGetAlgoFuturesOpenOrdersV1) | **Get** /sapi/v1/algo/futures/openOrders | Query Current Algo Open Orders(USER_DATA)
[**AlgoGetAlgoFuturesSubOrdersV1**](V1API.md#AlgoGetAlgoFuturesSubOrdersV1) | **Get** /sapi/v1/algo/futures/subOrders | Query Sub Orders(USER_DATA)
[**AlgoGetAlgoSpotHistoricalOrdersV1**](V1API.md#AlgoGetAlgoSpotHistoricalOrdersV1) | **Get** /sapi/v1/algo/spot/historicalOrders | Query Historical Algo Orders(USER_DATA)
[**AlgoGetAlgoSpotOpenOrdersV1**](V1API.md#AlgoGetAlgoSpotOpenOrdersV1) | **Get** /sapi/v1/algo/spot/openOrders | Query Current Algo Open Orders(USER_DATA)
[**AlgoGetAlgoSpotSubOrdersV1**](V1API.md#AlgoGetAlgoSpotSubOrdersV1) | **Get** /sapi/v1/algo/spot/subOrders | Query Sub Orders(USER_DATA)



## AlgoCreateAlgoFuturesNewOrderTwapV1

> AlgoCreateAlgoFuturesNewOrderTwapV1Resp AlgoCreateAlgoFuturesNewOrderTwapV1(ctx).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()

Time-Weighted Average Price(Twap) New Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	duration := int64(789) // int64 | 
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	clientAlgoId := "clientAlgoId_example" // string |  (optional) (default to "")
	limitPrice := "limitPrice_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoCreateAlgoFuturesNewOrderTwapV1(context.Background()).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoCreateAlgoFuturesNewOrderTwapV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoCreateAlgoFuturesNewOrderTwapV1`: AlgoCreateAlgoFuturesNewOrderTwapV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoCreateAlgoFuturesNewOrderTwapV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoCreateAlgoFuturesNewOrderTwapV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duration** | **int64** |  | 
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **clientAlgoId** | **string** |  | [default to &quot;&quot;]
 **limitPrice** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **bool** |  | 

### Return type

[**AlgoCreateAlgoFuturesNewOrderTwapV1Resp**](AlgoCreateAlgoFuturesNewOrderTwapV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoCreateAlgoFuturesNewOrderVpV1

> AlgoCreateAlgoFuturesNewOrderVpV1Resp AlgoCreateAlgoFuturesNewOrderVpV1(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Urgency(urgency).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()

Volume Participation(VP) New Order (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	urgency := "urgency_example" // string |  (default to "")
	clientAlgoId := "clientAlgoId_example" // string |  (optional) (default to "")
	limitPrice := "limitPrice_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoCreateAlgoFuturesNewOrderVpV1(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Urgency(urgency).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoCreateAlgoFuturesNewOrderVpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoCreateAlgoFuturesNewOrderVpV1`: AlgoCreateAlgoFuturesNewOrderVpV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoCreateAlgoFuturesNewOrderVpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoCreateAlgoFuturesNewOrderVpV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **urgency** | **string** |  | [default to &quot;&quot;]
 **clientAlgoId** | **string** |  | [default to &quot;&quot;]
 **limitPrice** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **bool** |  | 

### Return type

[**AlgoCreateAlgoFuturesNewOrderVpV1Resp**](AlgoCreateAlgoFuturesNewOrderVpV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoCreateAlgoSpotNewOrderTwapV1

> AlgoCreateAlgoSpotNewOrderTwapV1Resp AlgoCreateAlgoSpotNewOrderTwapV1(ctx).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).Execute()

Time-Weighted Average Price(Twap) New Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	duration := int64(789) // int64 | 
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	clientAlgoId := "clientAlgoId_example" // string |  (optional) (default to "")
	limitPrice := "limitPrice_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoCreateAlgoSpotNewOrderTwapV1(context.Background()).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoCreateAlgoSpotNewOrderTwapV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoCreateAlgoSpotNewOrderTwapV1`: AlgoCreateAlgoSpotNewOrderTwapV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoCreateAlgoSpotNewOrderTwapV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoCreateAlgoSpotNewOrderTwapV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duration** | **int64** |  | 
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **clientAlgoId** | **string** |  | [default to &quot;&quot;]
 **limitPrice** | **string** |  | [default to &quot;&quot;]

### Return type

[**AlgoCreateAlgoSpotNewOrderTwapV1Resp**](AlgoCreateAlgoSpotNewOrderTwapV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoDeleteAlgoFuturesOrderV1

> AlgoDeleteAlgoFuturesOrderV1Resp AlgoDeleteAlgoFuturesOrderV1(ctx).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Algo Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	algoId := int64(789) // int64 | eg. 14511
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoDeleteAlgoFuturesOrderV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoDeleteAlgoFuturesOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoDeleteAlgoFuturesOrderV1`: AlgoDeleteAlgoFuturesOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoDeleteAlgoFuturesOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoDeleteAlgoFuturesOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** | eg. 14511 | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoDeleteAlgoFuturesOrderV1Resp**](AlgoDeleteAlgoFuturesOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoDeleteAlgoSpotOrderV1

> AlgoDeleteAlgoSpotOrderV1Resp AlgoDeleteAlgoSpotOrderV1(ctx).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Algo Order(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	algoId := int64(789) // int64 | eg. 14511
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoDeleteAlgoSpotOrderV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoDeleteAlgoSpotOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoDeleteAlgoSpotOrderV1`: AlgoDeleteAlgoSpotOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoDeleteAlgoSpotOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoDeleteAlgoSpotOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** | eg. 14511 | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoDeleteAlgoSpotOrderV1Resp**](AlgoDeleteAlgoSpotOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoGetAlgoFuturesHistoricalOrdersV1

> AlgoGetAlgoFuturesHistoricalOrdersV1Resp AlgoGetAlgoFuturesHistoricalOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Historical Algo Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | Trading symbol eg. BTCUSDT (optional) (default to "")
	side := "side_example" // string | BUY or SELL (optional) (default to "")
	startTime := int64(789) // int64 | in milliseconds  eg.1641522717552 (optional)
	endTime := int64(789) // int64 | in milliseconds  eg.1641522526562 (optional)
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoGetAlgoFuturesHistoricalOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoGetAlgoFuturesHistoricalOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoGetAlgoFuturesHistoricalOrdersV1`: AlgoGetAlgoFuturesHistoricalOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoGetAlgoFuturesHistoricalOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoGetAlgoFuturesHistoricalOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Trading symbol eg. BTCUSDT | [default to &quot;&quot;]
 **side** | **string** | BUY or SELL | [default to &quot;&quot;]
 **startTime** | **int64** | in milliseconds  eg.1641522717552 | 
 **endTime** | **int64** | in milliseconds  eg.1641522526562 | 
 **page** | **int32** | Default is 1 | 
 **pageSize** | **int32** | MIN 1, MAX 100; Default 100 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoGetAlgoFuturesHistoricalOrdersV1Resp**](AlgoGetAlgoFuturesHistoricalOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoGetAlgoFuturesOpenOrdersV1

> AlgoGetAlgoFuturesOpenOrdersV1Resp AlgoGetAlgoFuturesOpenOrdersV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Current Algo Open Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoGetAlgoFuturesOpenOrdersV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoGetAlgoFuturesOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoGetAlgoFuturesOpenOrdersV1`: AlgoGetAlgoFuturesOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoGetAlgoFuturesOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoGetAlgoFuturesOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoGetAlgoFuturesOpenOrdersV1Resp**](AlgoGetAlgoFuturesOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoGetAlgoFuturesSubOrdersV1

> AlgoGetAlgoFuturesSubOrdersV1Resp AlgoGetAlgoFuturesSubOrdersV1(ctx).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Sub Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	algoId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoGetAlgoFuturesSubOrdersV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoGetAlgoFuturesSubOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoGetAlgoFuturesSubOrdersV1`: AlgoGetAlgoFuturesSubOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoGetAlgoFuturesSubOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoGetAlgoFuturesSubOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **page** | **int32** | Default is 1 | 
 **pageSize** | **int32** | MIN 1, MAX 100; Default 100 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoGetAlgoFuturesSubOrdersV1Resp**](AlgoGetAlgoFuturesSubOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoGetAlgoSpotHistoricalOrdersV1

> AlgoGetAlgoSpotHistoricalOrdersV1Resp AlgoGetAlgoSpotHistoricalOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Historical Algo Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | Trading symbol eg. BTCUSDT (optional) (default to "")
	side := "side_example" // string | BUY or SELL (optional) (default to "")
	startTime := int64(789) // int64 | in milliseconds  eg.1641522717552 (optional)
	endTime := int64(789) // int64 | in milliseconds  eg.1641522526562 (optional)
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoGetAlgoSpotHistoricalOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoGetAlgoSpotHistoricalOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoGetAlgoSpotHistoricalOrdersV1`: AlgoGetAlgoSpotHistoricalOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoGetAlgoSpotHistoricalOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoGetAlgoSpotHistoricalOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | Trading symbol eg. BTCUSDT | [default to &quot;&quot;]
 **side** | **string** | BUY or SELL | [default to &quot;&quot;]
 **startTime** | **int64** | in milliseconds  eg.1641522717552 | 
 **endTime** | **int64** | in milliseconds  eg.1641522526562 | 
 **page** | **int32** | Default is 1 | 
 **pageSize** | **int32** | MIN 1, MAX 100; Default 100 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoGetAlgoSpotHistoricalOrdersV1Resp**](AlgoGetAlgoSpotHistoricalOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoGetAlgoSpotOpenOrdersV1

> AlgoGetAlgoSpotOpenOrdersV1Resp AlgoGetAlgoSpotOpenOrdersV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Current Algo Open Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoGetAlgoSpotOpenOrdersV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoGetAlgoSpotOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoGetAlgoSpotOpenOrdersV1`: AlgoGetAlgoSpotOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoGetAlgoSpotOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoGetAlgoSpotOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoGetAlgoSpotOpenOrdersV1Resp**](AlgoGetAlgoSpotOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlgoGetAlgoSpotSubOrdersV1

> AlgoGetAlgoSpotSubOrdersV1Resp AlgoGetAlgoSpotSubOrdersV1(ctx).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Sub Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/algo"
)

func main() {
	algoId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.AlgoGetAlgoSpotSubOrdersV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.AlgoGetAlgoSpotSubOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlgoGetAlgoSpotSubOrdersV1`: AlgoGetAlgoSpotSubOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.AlgoGetAlgoSpotSubOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlgoGetAlgoSpotSubOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **page** | **int32** | Default is 1 | 
 **pageSize** | **int32** | MIN 1, MAX 100; Default 100 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**AlgoGetAlgoSpotSubOrdersV1Resp**](AlgoGetAlgoSpotSubOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

