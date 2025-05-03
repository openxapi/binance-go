# \TradeAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsCreateBatchOrdersV1**](TradeAPI.md#OptionsCreateBatchOrdersV1) | **Post** /eapi/v1/batchOrders | Place Multiple Orders(TRADE)
[**OptionsCreateOrderV1**](TradeAPI.md#OptionsCreateOrderV1) | **Post** /eapi/v1/order | New Order (TRADE)
[**OptionsDeleteBatchOrdersV1**](TradeAPI.md#OptionsDeleteBatchOrdersV1) | **Delete** /eapi/v1/batchOrders | Cancel Multiple Option Orders (TRADE)



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
	openapiclient "github.com/openxapi/binance-go/options"
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
	openapiclient "github.com/openxapi/binance-go/options"
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
	openapiclient "github.com/openxapi/binance-go/options"
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

