# \MarketMakerBlockTradeAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockOrderCreateV1**](MarketMakerBlockTradeAPI.md#CreateBlockOrderCreateV1) | **Post** /eapi/v1/block/order/create | New Block Trade Order (TRADE)
[**DeleteBlockOrderCreateV1**](MarketMakerBlockTradeAPI.md#DeleteBlockOrderCreateV1) | **Delete** /eapi/v1/block/order/create | Cancel Block Trade Order (TRADE)



## CreateBlockOrderCreateV1

> OptionsCreateBlockOrderCreateV1Resp CreateBlockOrderCreateV1(ctx).Legs(legs).Liquidity(liquidity).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.CreateBlockOrderCreateV1(context.Background()).Legs(legs).Liquidity(liquidity).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.CreateBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlockOrderCreateV1`: OptionsCreateBlockOrderCreateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.CreateBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlockOrderCreateV1Request struct via the builder pattern


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


## DeleteBlockOrderCreateV1

> map[string]interface{} DeleteBlockOrderCreateV1(ctx).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

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
	resp, r, err := apiClient.MarketMakerBlockTradeAPI.DeleteBlockOrderCreateV1(context.Background()).BlockOrderMatchingKey(blockOrderMatchingKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerBlockTradeAPI.DeleteBlockOrderCreateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlockOrderCreateV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerBlockTradeAPI.DeleteBlockOrderCreateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockOrderCreateV1Request struct via the builder pattern


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

