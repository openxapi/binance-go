# \MarketMakerBlockTradeAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsCreateBlockOrderCreateV1**](MarketMakerBlockTradeAPI.md#OptionsCreateBlockOrderCreateV1) | **Post** /eapi/v1/block/order/create | New Block Trade Order (TRADE)
[**OptionsDeleteBlockOrderCreateV1**](MarketMakerBlockTradeAPI.md#OptionsDeleteBlockOrderCreateV1) | **Delete** /eapi/v1/block/order/create | Cancel Block Trade Order (TRADE)



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
	openapiclient "github.com/openxapi/binance-go/options"
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
	openapiclient "github.com/openxapi/binance-go/options"
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

