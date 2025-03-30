# \TransferAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginGetMarginMaxTransferableV1**](TransferAPI.md#MarginGetMarginMaxTransferableV1) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)
[**MarginGetMarginTransferV1**](TransferAPI.md#MarginGetMarginTransferV1) | **Get** /sapi/v1/margin/transfer | Get Cross Margin Transfer History (USER_DATA)



## MarginGetMarginMaxTransferableV1

> MarginGetMarginMaxTransferableV1Resp MarginGetMarginMaxTransferableV1(ctx).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Transfer-Out Amount (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.MarginGetMarginMaxTransferableV1(context.Background()).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.MarginGetMarginMaxTransferableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginMaxTransferableV1`: MarginGetMarginMaxTransferableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.MarginGetMarginMaxTransferableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginMaxTransferableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isolatedSymbol** | **string** | isolated symbol | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginMaxTransferableV1Resp**](MarginGetMarginMaxTransferableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginTransferV1

> MarginGetMarginTransferV1Resp MarginGetMarginTransferV1(ctx).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Get Cross Margin Transfer History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	type_ := "type__example" // string | Transfer Type: ROLL_IN, ROLL_OUT (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	isolatedSymbol := "isolatedSymbol_example" // string | Symbol in Isolated Margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.MarginGetMarginTransferV1(context.Background()).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.MarginGetMarginTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginTransferV1`: MarginGetMarginTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.MarginGetMarginTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **type_** | **string** | Transfer Type: ROLL_IN, ROLL_OUT | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **isolatedSymbol** | **string** | Symbol in Isolated Margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginTransferV1Resp**](MarginGetMarginTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

