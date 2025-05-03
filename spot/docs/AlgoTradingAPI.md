# \AlgoTradingAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlgoFuturesNewOrderTwapV1**](AlgoTradingAPI.md#CreateAlgoFuturesNewOrderTwapV1) | **Post** /sapi/v1/algo/futures/newOrderTwap | Time-Weighted Average Price(Twap) New Order(TRADE)
[**CreateAlgoFuturesNewOrderVpV1**](AlgoTradingAPI.md#CreateAlgoFuturesNewOrderVpV1) | **Post** /sapi/v1/algo/futures/newOrderVp | Volume Participation(VP) New Order (TRADE)
[**CreateAlgoSpotNewOrderTwapV1**](AlgoTradingAPI.md#CreateAlgoSpotNewOrderTwapV1) | **Post** /sapi/v1/algo/spot/newOrderTwap | Time-Weighted Average Price(Twap) New Order(TRADE)
[**DeleteAlgoFuturesOrderV1**](AlgoTradingAPI.md#DeleteAlgoFuturesOrderV1) | **Delete** /sapi/v1/algo/futures/order | Cancel Algo Order(TRADE)
[**DeleteAlgoSpotOrderV1**](AlgoTradingAPI.md#DeleteAlgoSpotOrderV1) | **Delete** /sapi/v1/algo/spot/order | Cancel Algo Order(TRADE)
[**GetAlgoFuturesHistoricalOrdersV1**](AlgoTradingAPI.md#GetAlgoFuturesHistoricalOrdersV1) | **Get** /sapi/v1/algo/futures/historicalOrders | Query Historical Algo Orders(USER_DATA)
[**GetAlgoFuturesOpenOrdersV1**](AlgoTradingAPI.md#GetAlgoFuturesOpenOrdersV1) | **Get** /sapi/v1/algo/futures/openOrders | Query Current Algo Open Orders(USER_DATA)
[**GetAlgoFuturesSubOrdersV1**](AlgoTradingAPI.md#GetAlgoFuturesSubOrdersV1) | **Get** /sapi/v1/algo/futures/subOrders | Query Sub Orders(USER_DATA)
[**GetAlgoSpotHistoricalOrdersV1**](AlgoTradingAPI.md#GetAlgoSpotHistoricalOrdersV1) | **Get** /sapi/v1/algo/spot/historicalOrders | Query Historical Algo Orders(USER_DATA)
[**GetAlgoSpotOpenOrdersV1**](AlgoTradingAPI.md#GetAlgoSpotOpenOrdersV1) | **Get** /sapi/v1/algo/spot/openOrders | Query Current Algo Open Orders(USER_DATA)
[**GetAlgoSpotSubOrdersV1**](AlgoTradingAPI.md#GetAlgoSpotSubOrdersV1) | **Get** /sapi/v1/algo/spot/subOrders | Query Sub Orders(USER_DATA)



## CreateAlgoFuturesNewOrderTwapV1

> CreateAlgoFuturesNewOrderTwapV1Resp CreateAlgoFuturesNewOrderTwapV1(ctx).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()

Time-Weighted Average Price(Twap) New Order(TRADE)



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
	resp, r, err := apiClient.AlgoTradingAPI.CreateAlgoFuturesNewOrderTwapV1(context.Background()).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.CreateAlgoFuturesNewOrderTwapV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlgoFuturesNewOrderTwapV1`: CreateAlgoFuturesNewOrderTwapV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.CreateAlgoFuturesNewOrderTwapV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlgoFuturesNewOrderTwapV1Request struct via the builder pattern


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

[**CreateAlgoFuturesNewOrderTwapV1Resp**](CreateAlgoFuturesNewOrderTwapV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlgoFuturesNewOrderVpV1

> CreateAlgoFuturesNewOrderVpV1Resp CreateAlgoFuturesNewOrderVpV1(ctx).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Urgency(urgency).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()

Volume Participation(VP) New Order (TRADE)



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
	resp, r, err := apiClient.AlgoTradingAPI.CreateAlgoFuturesNewOrderVpV1(context.Background()).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).Urgency(urgency).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).PositionSide(positionSide).RecvWindow(recvWindow).ReduceOnly(reduceOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.CreateAlgoFuturesNewOrderVpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlgoFuturesNewOrderVpV1`: CreateAlgoFuturesNewOrderVpV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.CreateAlgoFuturesNewOrderVpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlgoFuturesNewOrderVpV1Request struct via the builder pattern


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

[**CreateAlgoFuturesNewOrderVpV1Resp**](CreateAlgoFuturesNewOrderVpV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlgoSpotNewOrderTwapV1

> CreateAlgoSpotNewOrderTwapV1Resp CreateAlgoSpotNewOrderTwapV1(ctx).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).Execute()

Time-Weighted Average Price(Twap) New Order(TRADE)



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
	duration := int64(789) // int64 | 
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	clientAlgoId := "clientAlgoId_example" // string |  (optional) (default to "")
	limitPrice := "limitPrice_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.CreateAlgoSpotNewOrderTwapV1(context.Background()).Duration(duration).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).ClientAlgoId(clientAlgoId).LimitPrice(limitPrice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.CreateAlgoSpotNewOrderTwapV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlgoSpotNewOrderTwapV1`: CreateAlgoSpotNewOrderTwapV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.CreateAlgoSpotNewOrderTwapV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlgoSpotNewOrderTwapV1Request struct via the builder pattern


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

[**CreateAlgoSpotNewOrderTwapV1Resp**](CreateAlgoSpotNewOrderTwapV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlgoFuturesOrderV1

> DeleteAlgoFuturesOrderV1Resp DeleteAlgoFuturesOrderV1(ctx).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Algo Order(TRADE)



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
	algoId := int64(789) // int64 | eg. 14511
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.DeleteAlgoFuturesOrderV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.DeleteAlgoFuturesOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlgoFuturesOrderV1`: DeleteAlgoFuturesOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.DeleteAlgoFuturesOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlgoFuturesOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** | eg. 14511 | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**DeleteAlgoFuturesOrderV1Resp**](DeleteAlgoFuturesOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlgoSpotOrderV1

> DeleteAlgoSpotOrderV1Resp DeleteAlgoSpotOrderV1(ctx).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Algo Order(TRADE)



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
	algoId := int64(789) // int64 | eg. 14511
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.DeleteAlgoSpotOrderV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.DeleteAlgoSpotOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlgoSpotOrderV1`: DeleteAlgoSpotOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.DeleteAlgoSpotOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlgoSpotOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** | eg. 14511 | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**DeleteAlgoSpotOrderV1Resp**](DeleteAlgoSpotOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlgoFuturesHistoricalOrdersV1

> GetAlgoFuturesHistoricalOrdersV1Resp GetAlgoFuturesHistoricalOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Historical Algo Orders(USER_DATA)



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
	symbol := "symbol_example" // string | Trading symbol eg. BTCUSDT (optional) (default to "")
	side := "side_example" // string | BUY or SELL (optional) (default to "")
	startTime := int64(789) // int64 | in milliseconds  eg.1641522717552 (optional)
	endTime := int64(789) // int64 | in milliseconds  eg.1641522526562 (optional)
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetAlgoFuturesHistoricalOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetAlgoFuturesHistoricalOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlgoFuturesHistoricalOrdersV1`: GetAlgoFuturesHistoricalOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetAlgoFuturesHistoricalOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlgoFuturesHistoricalOrdersV1Request struct via the builder pattern


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

[**GetAlgoFuturesHistoricalOrdersV1Resp**](GetAlgoFuturesHistoricalOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlgoFuturesOpenOrdersV1

> GetAlgoFuturesOpenOrdersV1Resp GetAlgoFuturesOpenOrdersV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Current Algo Open Orders(USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetAlgoFuturesOpenOrdersV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetAlgoFuturesOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlgoFuturesOpenOrdersV1`: GetAlgoFuturesOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetAlgoFuturesOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlgoFuturesOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAlgoFuturesOpenOrdersV1Resp**](GetAlgoFuturesOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlgoFuturesSubOrdersV1

> GetAlgoFuturesSubOrdersV1Resp GetAlgoFuturesSubOrdersV1(ctx).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Sub Orders(USER_DATA)



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
	algoId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetAlgoFuturesSubOrdersV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetAlgoFuturesSubOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlgoFuturesSubOrdersV1`: GetAlgoFuturesSubOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetAlgoFuturesSubOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlgoFuturesSubOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **page** | **int32** | Default is 1 | 
 **pageSize** | **int32** | MIN 1, MAX 100; Default 100 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**GetAlgoFuturesSubOrdersV1Resp**](GetAlgoFuturesSubOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlgoSpotHistoricalOrdersV1

> GetAlgoSpotHistoricalOrdersV1Resp GetAlgoSpotHistoricalOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Historical Algo Orders(USER_DATA)



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
	symbol := "symbol_example" // string | Trading symbol eg. BTCUSDT (optional) (default to "")
	side := "side_example" // string | BUY or SELL (optional) (default to "")
	startTime := int64(789) // int64 | in milliseconds  eg.1641522717552 (optional)
	endTime := int64(789) // int64 | in milliseconds  eg.1641522526562 (optional)
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetAlgoSpotHistoricalOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Side(side).StartTime(startTime).EndTime(endTime).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetAlgoSpotHistoricalOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlgoSpotHistoricalOrdersV1`: GetAlgoSpotHistoricalOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetAlgoSpotHistoricalOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlgoSpotHistoricalOrdersV1Request struct via the builder pattern


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

[**GetAlgoSpotHistoricalOrdersV1Resp**](GetAlgoSpotHistoricalOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlgoSpotOpenOrdersV1

> GetAlgoSpotOpenOrdersV1Resp GetAlgoSpotOpenOrdersV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Current Algo Open Orders(USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetAlgoSpotOpenOrdersV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetAlgoSpotOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlgoSpotOpenOrdersV1`: GetAlgoSpotOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetAlgoSpotOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlgoSpotOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAlgoSpotOpenOrdersV1Resp**](GetAlgoSpotOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlgoSpotSubOrdersV1

> GetAlgoSpotSubOrdersV1Resp GetAlgoSpotSubOrdersV1(ctx).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Query Sub Orders(USER_DATA)



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
	algoId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	page := int32(56) // int32 | Default is 1 (optional)
	pageSize := int32(56) // int32 | MIN 1, MAX 100; Default 100 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetAlgoSpotSubOrdersV1(context.Background()).AlgoId(algoId).Timestamp(timestamp).Page(page).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetAlgoSpotSubOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlgoSpotSubOrdersV1`: GetAlgoSpotSubOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetAlgoSpotSubOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlgoSpotSubOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **page** | **int32** | Default is 1 | 
 **pageSize** | **int32** | MIN 1, MAX 100; Default 100 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**GetAlgoSpotSubOrdersV1Resp**](GetAlgoSpotSubOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

