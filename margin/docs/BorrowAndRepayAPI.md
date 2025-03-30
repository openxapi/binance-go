# \BorrowAndRepayAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCreateMarginBorrowRepayV1**](BorrowAndRepayAPI.md#MarginCreateMarginBorrowRepayV1) | **Post** /sapi/v1/margin/borrow-repay | Margin account borrow/repay(MARGIN)
[**MarginGetMarginBorrowRepayV1**](BorrowAndRepayAPI.md#MarginGetMarginBorrowRepayV1) | **Get** /sapi/v1/margin/borrow-repay | Query borrow/repay records in Margin account(USER_DATA)
[**MarginGetMarginInterestHistoryV1**](BorrowAndRepayAPI.md#MarginGetMarginInterestHistoryV1) | **Get** /sapi/v1/margin/interestHistory | Get Interest History (USER_DATA)
[**MarginGetMarginInterestRateHistoryV1**](BorrowAndRepayAPI.md#MarginGetMarginInterestRateHistoryV1) | **Get** /sapi/v1/margin/interestRateHistory | Query Margin Interest Rate History (USER_DATA)
[**MarginGetMarginMaxBorrowableV1**](BorrowAndRepayAPI.md#MarginGetMarginMaxBorrowableV1) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)
[**MarginGetMarginNextHourlyInterestRateV1**](BorrowAndRepayAPI.md#MarginGetMarginNextHourlyInterestRateV1) | **Get** /sapi/v1/margin/next-hourly-interest-rate | Get future hourly interest rate (USER_DATA)



## MarginCreateMarginBorrowRepayV1

> MarginCreateMarginBorrowRepayV1Resp MarginCreateMarginBorrowRepayV1(ctx).Amount(amount).Asset(asset).IsIsolated(isIsolated).Symbol(symbol).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Margin account borrow/repay(MARGIN)



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
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	isIsolated := "isIsolated_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BorrowAndRepayAPI.MarginCreateMarginBorrowRepayV1(context.Background()).Amount(amount).Asset(asset).IsIsolated(isIsolated).Symbol(symbol).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BorrowAndRepayAPI.MarginCreateMarginBorrowRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginBorrowRepayV1`: MarginCreateMarginBorrowRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BorrowAndRepayAPI.MarginCreateMarginBorrowRepayV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginBorrowRepayV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**MarginCreateMarginBorrowRepayV1Resp**](MarginCreateMarginBorrowRepayV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginBorrowRepayV1

> MarginGetMarginBorrowRepayV1Resp MarginGetMarginBorrowRepayV1(ctx).Type_(type_).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query borrow/repay records in Margin account(USER_DATA)



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
	type_ := "type__example" // string | `BORROW` or `REPAY` (default to "")
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	isolatedSymbol := "isolatedSymbol_example" // string | Symbol in Isolated Margin (optional) (default to "")
	txId := int64(789) // int64 | `tranId` in `POST /sapi/v1/margin/loan` (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BorrowAndRepayAPI.MarginGetMarginBorrowRepayV1(context.Background()).Type_(type_).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BorrowAndRepayAPI.MarginGetMarginBorrowRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginBorrowRepayV1`: MarginGetMarginBorrowRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BorrowAndRepayAPI.MarginGetMarginBorrowRepayV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginBorrowRepayV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &#x60;BORROW&#x60; or &#x60;REPAY&#x60; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **isolatedSymbol** | **string** | Symbol in Isolated Margin | [default to &quot;&quot;]
 **txId** | **int64** | &#x60;tranId&#x60; in &#x60;POST /sapi/v1/margin/loan&#x60; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**MarginGetMarginBorrowRepayV1Resp**](MarginGetMarginBorrowRepayV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginInterestHistoryV1

> MarginGetMarginInterestHistoryV1Resp MarginGetMarginInterestHistoryV1(ctx).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Interest History (USER_DATA)



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
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BorrowAndRepayAPI.MarginGetMarginInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BorrowAndRepayAPI.MarginGetMarginInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginInterestHistoryV1`: MarginGetMarginInterestHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BorrowAndRepayAPI.MarginGetMarginInterestHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginInterestHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **isolatedSymbol** | **string** | isolated symbol | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginInterestHistoryV1Resp**](MarginGetMarginInterestHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginInterestRateHistoryV1

> []MarginGetMarginInterestRateHistoryV1RespItem MarginGetMarginInterestRateHistoryV1(ctx).Asset(asset).Timestamp(timestamp).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Query Margin Interest Rate History (USER_DATA)



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
	vipLevel := int32(56) // int32 | Default: user&#39;s vip level (optional)
	startTime := int64(789) // int64 | Default: 7 days ago (optional)
	endTime := int64(789) // int64 | Default: present. Maximum range: 1 months. (optional)
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BorrowAndRepayAPI.MarginGetMarginInterestRateHistoryV1(context.Background()).Asset(asset).Timestamp(timestamp).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BorrowAndRepayAPI.MarginGetMarginInterestRateHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginInterestRateHistoryV1`: []MarginGetMarginInterestRateHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BorrowAndRepayAPI.MarginGetMarginInterestRateHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginInterestRateHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | Default: user&amp;#39;s vip level | 
 **startTime** | **int64** | Default: 7 days ago | 
 **endTime** | **int64** | Default: present. Maximum range: 1 months. | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**[]MarginGetMarginInterestRateHistoryV1RespItem**](MarginGetMarginInterestRateHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginMaxBorrowableV1

> MarginGetMarginMaxBorrowableV1Resp MarginGetMarginMaxBorrowableV1(ctx).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Borrow (USER_DATA)



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
	resp, r, err := apiClient.BorrowAndRepayAPI.MarginGetMarginMaxBorrowableV1(context.Background()).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BorrowAndRepayAPI.MarginGetMarginMaxBorrowableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginMaxBorrowableV1`: MarginGetMarginMaxBorrowableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BorrowAndRepayAPI.MarginGetMarginMaxBorrowableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginMaxBorrowableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isolatedSymbol** | **string** | isolated symbol | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginMaxBorrowableV1Resp**](MarginGetMarginMaxBorrowableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginNextHourlyInterestRateV1

> []MarginGetMarginNextHourlyInterestRateV1RespItem MarginGetMarginNextHourlyInterestRateV1(ctx).Assets(assets).IsIsolated(isIsolated).Execute()

Get future hourly interest rate (USER_DATA)



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
	assets := "assets_example" // string | List of assets, separated by commas, up to 20 (default to "")
	isIsolated := true // bool | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BorrowAndRepayAPI.MarginGetMarginNextHourlyInterestRateV1(context.Background()).Assets(assets).IsIsolated(isIsolated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BorrowAndRepayAPI.MarginGetMarginNextHourlyInterestRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginNextHourlyInterestRateV1`: []MarginGetMarginNextHourlyInterestRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BorrowAndRepayAPI.MarginGetMarginNextHourlyInterestRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginNextHourlyInterestRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assets** | **string** | List of assets, separated by commas, up to 20 | [default to &quot;&quot;]
 **isIsolated** | **bool** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34; | 

### Return type

[**[]MarginGetMarginNextHourlyInterestRateV1RespItem**](MarginGetMarginNextHourlyInterestRateV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

