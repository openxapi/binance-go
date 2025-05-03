# \TradeAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCreateMarginExchangeSmallLiabilityV1**](TradeAPI.md#MarginCreateMarginExchangeSmallLiabilityV1) | **Post** /sapi/v1/margin/exchange-small-liability | Small Liability Exchange (MARGIN)
[**MarginCreateMarginOrderV1**](TradeAPI.md#MarginCreateMarginOrderV1) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**MarginGetMarginOpenOrdersV1**](TradeAPI.md#MarginGetMarginOpenOrdersV1) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)



## MarginCreateMarginExchangeSmallLiabilityV1

> MarginCreateMarginExchangeSmallLiabilityV1Resp MarginCreateMarginExchangeSmallLiabilityV1(ctx).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Small Liability Exchange (MARGIN)



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
	assetNames := []string{"Inner_example"} // []string | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginExchangeSmallLiabilityV1(context.Background()).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginExchangeSmallLiabilityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginExchangeSmallLiabilityV1`: MarginCreateMarginExchangeSmallLiabilityV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginExchangeSmallLiabilityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginExchangeSmallLiabilityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetNames** | **[]string** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginCreateMarginExchangeSmallLiabilityV1Resp**](MarginCreateMarginExchangeSmallLiabilityV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderV1

> MarginCreateMarginOrderV1Resp MarginCreateMarginOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()

Margin Account New Order (TRADE)



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
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	quoteOrderQty := "quoteOrderQty_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginCreateMarginOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginCreateMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderV1`: MarginCreateMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginCreateMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **icebergQty** | **string** |  | [default to &quot;&quot;]
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **quoteOrderQty** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderV1Resp**](MarginCreateMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOpenOrdersV1

> []MarginGetMarginOpenOrdersV1RespItem MarginGetMarginOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Margin Account's Open Orders (USER_DATA)



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
	symbol := "symbol_example" // string |  (optional) (default to "")
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.MarginGetMarginOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.MarginGetMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOpenOrdersV1`: []MarginGetMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.MarginGetMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginOpenOrdersV1RespItem**](MarginGetMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

