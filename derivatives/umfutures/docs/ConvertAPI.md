# \ConvertAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UmfuturesCreateConvertAcceptQuoteV1**](ConvertAPI.md#UmfuturesCreateConvertAcceptQuoteV1) | **Post** /fapi/v1/convert/acceptQuote | Accept the offered quote (USER_DATA)
[**UmfuturesCreateConvertGetQuoteV1**](ConvertAPI.md#UmfuturesCreateConvertGetQuoteV1) | **Post** /fapi/v1/convert/getQuote | Send Quote Request(USER_DATA)
[**UmfuturesGetConvertExchangeInfoV1**](ConvertAPI.md#UmfuturesGetConvertExchangeInfoV1) | **Get** /fapi/v1/convert/exchangeInfo | List All Convert Pairs
[**UmfuturesGetConvertOrderStatusV1**](ConvertAPI.md#UmfuturesGetConvertOrderStatusV1) | **Get** /fapi/v1/convert/orderStatus | Order status(USER_DATA)



## UmfuturesCreateConvertAcceptQuoteV1

> UmfuturesCreateConvertAcceptQuoteV1Resp UmfuturesCreateConvertAcceptQuoteV1(ctx).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Accept the offered quote (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	quoteId := "quoteId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.UmfuturesCreateConvertAcceptQuoteV1(context.Background()).QuoteId(quoteId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UmfuturesCreateConvertAcceptQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesCreateConvertAcceptQuoteV1`: UmfuturesCreateConvertAcceptQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UmfuturesCreateConvertAcceptQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesCreateConvertAcceptQuoteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UmfuturesCreateConvertAcceptQuoteV1Resp**](UmfuturesCreateConvertAcceptQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesCreateConvertGetQuoteV1

> UmfuturesCreateConvertGetQuoteV1Resp UmfuturesCreateConvertGetQuoteV1(ctx).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).Execute()

Send Quote Request(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
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
	resp, r, err := apiClient.ConvertAPI.UmfuturesCreateConvertGetQuoteV1(context.Background()).FromAsset(fromAsset).Timestamp(timestamp).ToAsset(toAsset).FromAmount(fromAmount).RecvWindow(recvWindow).ToAmount(toAmount).ValidTime(validTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UmfuturesCreateConvertGetQuoteV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesCreateConvertGetQuoteV1`: UmfuturesCreateConvertGetQuoteV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UmfuturesCreateConvertGetQuoteV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesCreateConvertGetQuoteV1Request struct via the builder pattern


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

[**UmfuturesCreateConvertGetQuoteV1Resp**](UmfuturesCreateConvertGetQuoteV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetConvertExchangeInfoV1

> []UmfuturesGetConvertExchangeInfoV1RespItem UmfuturesGetConvertExchangeInfoV1(ctx).FromAsset(fromAsset).ToAsset(toAsset).Execute()

List All Convert Pairs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	fromAsset := "fromAsset_example" // string | User spends coin (optional) (default to "")
	toAsset := "toAsset_example" // string | User receives coin (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.UmfuturesGetConvertExchangeInfoV1(context.Background()).FromAsset(fromAsset).ToAsset(toAsset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UmfuturesGetConvertExchangeInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetConvertExchangeInfoV1`: []UmfuturesGetConvertExchangeInfoV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UmfuturesGetConvertExchangeInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetConvertExchangeInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromAsset** | **string** | User spends coin | [default to &quot;&quot;]
 **toAsset** | **string** | User receives coin | [default to &quot;&quot;]

### Return type

[**[]UmfuturesGetConvertExchangeInfoV1RespItem**](UmfuturesGetConvertExchangeInfoV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UmfuturesGetConvertOrderStatusV1

> UmfuturesGetConvertOrderStatusV1Resp UmfuturesGetConvertOrderStatusV1(ctx).OrderId(orderId).QuoteId(quoteId).Execute()

Order status(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {
	orderId := "orderId_example" // string | Either orderId or quoteId is required (optional) (default to "")
	quoteId := "quoteId_example" // string | Either orderId or quoteId is required (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConvertAPI.UmfuturesGetConvertOrderStatusV1(context.Background()).OrderId(orderId).QuoteId(quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConvertAPI.UmfuturesGetConvertOrderStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UmfuturesGetConvertOrderStatusV1`: UmfuturesGetConvertOrderStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ConvertAPI.UmfuturesGetConvertOrderStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUmfuturesGetConvertOrderStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]
 **quoteId** | **string** | Either orderId or quoteId is required | [default to &quot;&quot;]

### Return type

[**UmfuturesGetConvertOrderStatusV1Resp**](UmfuturesGetConvertOrderStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

