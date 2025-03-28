# \ConvertAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesCreateConvertAcceptQuoteV1**](ConvertAPI.md#UfuturesCreateConvertAcceptQuoteV1) | **Post** /fapi/v1/convert/acceptQuote | Accept the offered quote (USER_DATA)
[**UfuturesCreateConvertGetQuoteV1**](ConvertAPI.md#UfuturesCreateConvertGetQuoteV1) | **Post** /fapi/v1/convert/getQuote | Send Quote Request(USER_DATA)
[**UfuturesGetConvertExchangeInfoV1**](ConvertAPI.md#UfuturesGetConvertExchangeInfoV1) | **Get** /fapi/v1/convert/exchangeInfo | List All Convert Pairs
[**UfuturesGetConvertOrderStatusV1**](ConvertAPI.md#UfuturesGetConvertOrderStatusV1) | **Get** /fapi/v1/convert/orderStatus | Order status(USER_DATA)



## UfuturesCreateConvertAcceptQuoteV1

> UfuturesCreateConvertAcceptQuoteV1Resp UfuturesCreateConvertAcceptQuoteV1(ctx).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept the offered quote (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	quoteId := "quoteId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.UfuturesCreateConvertAcceptQuoteV1(context.Background()).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UfuturesCreateConvertAcceptQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateConvertAcceptQuoteV1`: UfuturesCreateConvertAcceptQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UfuturesCreateConvertAcceptQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateConvertAcceptQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreateConvertAcceptQuoteV1Resp**](UfuturesCreateConvertAcceptQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesCreateConvertGetQuoteV1

> UfuturesCreateConvertGetQuoteV1Resp UfuturesCreateConvertGetQuoteV1(ctx).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).Execute()

Send Quote Request(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	fromAsset := "fromAsset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toAsset := "toAsset_example" // string |  (default to "")
	fromAmount := "fromAmount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	toAmount := "toAmount_example" // string |  (optional) (default to "")
	validTime := "validTime_example" // string |  (optional) (default to "10")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.UfuturesCreateConvertGetQuoteV1(context.Background()).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UfuturesCreateConvertGetQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateConvertGetQuoteV1`: UfuturesCreateConvertGetQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UfuturesCreateConvertGetQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateConvertGetQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toAsset** | **string** |  | [default to &quot;&quot;]
 **fromAmount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **toAmount** | **string** |  | [default to &quot;&quot;]
 **validTime** | **string** |  | [default to &quot;10&quot;]

### Return type

[**UfuturesCreateConvertGetQuoteV1Resp**](UfuturesCreateConvertGetQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetConvertExchangeInfoV1

> []UfuturesGetConvertExchangeInfoV1RespItem UfuturesGetConvertExchangeInfoV1(ctx).FromAsset(fromAsset).ToAsset(toAsset).Execute()

List All Convert Pairs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	fromAsset := "fromAsset_example" // string | User spends coin (optional) (default to "")
	toAsset := "toAsset_example" // string | User receives coin (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.UfuturesGetConvertExchangeInfoV1(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UfuturesGetConvertExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetConvertExchangeInfoV1`: []UfuturesGetConvertExchangeInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UfuturesGetConvertExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetConvertExchangeInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** | User spends coin | [default to &quot;&quot;]
 **toAsset** | **string** | User receives coin | [default to &quot;&quot;]

### Return type

[**[]UfuturesGetConvertExchangeInfoV1RespItem**](UfuturesGetConvertExchangeInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetConvertOrderStatusV1

> UfuturesGetConvertOrderStatusV1Resp UfuturesGetConvertOrderStatusV1(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/ufutures"
)

func main() {
	orderId := "orderId_example" // string | Either orderId or quoteId is required (optional) (default to "")
	quoteId := "quoteId_example" // string | Either orderId or quoteId is required (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.UfuturesGetConvertOrderStatusV1(context.Background()).OrderId(orderId).QuoteId(quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UfuturesGetConvertOrderStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetConvertOrderStatusV1`: UfuturesGetConvertOrderStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UfuturesGetConvertOrderStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetConvertOrderStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]
 **quoteId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]

### Return type

[**UfuturesGetConvertOrderStatusV1Resp**](UfuturesGetConvertOrderStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

