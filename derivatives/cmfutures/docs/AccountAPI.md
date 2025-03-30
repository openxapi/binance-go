# \AccountAPI

All URIs are relative to *https://dapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfuturesGetAccountV1**](AccountAPI.md#CfuturesGetAccountV1) | **Get** /dapi/v1/account | Account Information (USER_DATA)
[**CfuturesGetBalanceV1**](AccountAPI.md#CfuturesGetBalanceV1) | **Get** /dapi/v1/balance | Futures Account Balance (USER_DATA)
[**CfuturesGetCommissionRateV1**](AccountAPI.md#CfuturesGetCommissionRateV1) | **Get** /dapi/v1/commissionRate | User Commission Rate (USER_DATA)
[**CfuturesGetIncomeAsynIdV1**](AccountAPI.md#CfuturesGetIncomeAsynIdV1) | **Get** /dapi/v1/income/asyn/id | Get Futures Transaction History Download Link by Id (USER_DATA)
[**CfuturesGetIncomeAsynV1**](AccountAPI.md#CfuturesGetIncomeAsynV1) | **Get** /dapi/v1/income/asyn | Get Download Id For Futures Transaction History(USER_DATA)
[**CfuturesGetIncomeV1**](AccountAPI.md#CfuturesGetIncomeV1) | **Get** /dapi/v1/income | Get Income History(USER_DATA)
[**CfuturesGetLeverageBracketV1**](AccountAPI.md#CfuturesGetLeverageBracketV1) | **Get** /dapi/v1/leverageBracket | Notional Bracket for Pair(USER_DATA)
[**CfuturesGetLeverageBracketV2**](AccountAPI.md#CfuturesGetLeverageBracketV2) | **Get** /dapi/v2/leverageBracket | Notional Bracket for Symbol(USER_DATA)
[**CfuturesGetOrderAsynIdV1**](AccountAPI.md#CfuturesGetOrderAsynIdV1) | **Get** /dapi/v1/order/asyn/id | Get Futures Order History Download Link by Id (USER_DATA)
[**CfuturesGetOrderAsynV1**](AccountAPI.md#CfuturesGetOrderAsynV1) | **Get** /dapi/v1/order/asyn | Get Download Id For Futures Order History (USER_DATA)
[**CfuturesGetPositionSideDualV1**](AccountAPI.md#CfuturesGetPositionSideDualV1) | **Get** /dapi/v1/positionSide/dual | Get Current Position Mode(USER_DATA)
[**CfuturesGetTradeAsynIdV1**](AccountAPI.md#CfuturesGetTradeAsynIdV1) | **Get** /dapi/v1/trade/asyn/id | Get Futures Trade Download Link by Id(USER_DATA)
[**CfuturesGetTradeAsynV1**](AccountAPI.md#CfuturesGetTradeAsynV1) | **Get** /dapi/v1/trade/asyn | Get Download Id For Futures Trade History (USER_DATA)



## CfuturesGetAccountV1

> CfuturesGetAccountV1Resp CfuturesGetAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetAccountV1`: CfuturesGetAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetAccountV1Resp**](CfuturesGetAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetBalanceV1

> []CfuturesGetBalanceV1RespItem CfuturesGetBalanceV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetBalanceV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetBalanceV1`: []CfuturesGetBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetBalanceV1RespItem**](CfuturesGetBalanceV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetCommissionRateV1

> CfuturesGetCommissionRateV1Resp CfuturesGetCommissionRateV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

User Commission Rate (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetCommissionRateV1`: CfuturesGetCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetCommissionRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetCommissionRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetCommissionRateV1Resp**](CfuturesGetCommissionRateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetIncomeAsynIdV1

> CfuturesGetIncomeAsynIdV1Resp CfuturesGetIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Transaction History Download Link by Id (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIncomeAsynIdV1`: CfuturesGetIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetIncomeAsynIdV1Resp**](CfuturesGetIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetIncomeAsynV1

> CfuturesGetIncomeAsynV1Resp CfuturesGetIncomeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Transaction History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetIncomeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIncomeAsynV1`: CfuturesGetIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetIncomeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetIncomeAsynV1Resp**](CfuturesGetIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetIncomeV1

> []CfuturesGetIncomeV1RespItem CfuturesGetIncomeV1(ctx).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Income History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	incomeType := "incomeType_example" // string | &#34;TRANSFER&#34;,&#34;WELCOME_BONUS&#34;, &#34;FUNDING_FEE&#34;, &#34;REALIZED_PNL&#34;, &#34;COMMISSION&#34;, &#34;INSURANCE_CLEAR&#34;, and &#34;DELIVERED_SETTELMENT&#34; (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetIncomeV1`: []CfuturesGetIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **incomeType** | **string** | &amp;#34;TRANSFER&amp;#34;,&amp;#34;WELCOME_BONUS&amp;#34;, &amp;#34;FUNDING_FEE&amp;#34;, &amp;#34;REALIZED_PNL&amp;#34;, &amp;#34;COMMISSION&amp;#34;, &amp;#34;INSURANCE_CLEAR&amp;#34;, and &amp;#34;DELIVERED_SETTELMENT&amp;#34; | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int32** |  | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetIncomeV1RespItem**](CfuturesGetIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetLeverageBracketV1

> []CfuturesGetLeverageBracketV1RespItem CfuturesGetLeverageBracketV1(ctx).Timestamp(timestamp).Pair(pair).RecvWindow(recvWindow).Execute()

Notional Bracket for Pair(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	pair := "pair_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetLeverageBracketV1(context.Background()).Timestamp(timestamp).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetLeverageBracketV1`: []CfuturesGetLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetLeverageBracketV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **pair** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetLeverageBracketV1RespItem**](CfuturesGetLeverageBracketV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetLeverageBracketV2

> []CfuturesGetLeverageBracketV2RespItem CfuturesGetLeverageBracketV2(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Notional Bracket for Symbol(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetLeverageBracketV2(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetLeverageBracketV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetLeverageBracketV2`: []CfuturesGetLeverageBracketV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetLeverageBracketV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetLeverageBracketV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetLeverageBracketV2RespItem**](CfuturesGetLeverageBracketV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOrderAsynIdV1

> CfuturesGetOrderAsynIdV1Resp CfuturesGetOrderAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Order History Download Link by Id (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetOrderAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetOrderAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOrderAsynIdV1`: CfuturesGetOrderAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetOrderAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOrderAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetOrderAsynIdV1Resp**](CfuturesGetOrderAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetOrderAsynV1

> CfuturesGetOrderAsynV1Resp CfuturesGetOrderAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Order History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetOrderAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetOrderAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetOrderAsynV1`: CfuturesGetOrderAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetOrderAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetOrderAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetOrderAsynV1Resp**](CfuturesGetOrderAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetPositionSideDualV1

> CfuturesGetPositionSideDualV1Resp CfuturesGetPositionSideDualV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Current Position Mode(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetPositionSideDualV1`: CfuturesGetPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetPositionSideDualV1Resp**](CfuturesGetPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTradeAsynIdV1

> CfuturesGetTradeAsynIdV1Resp CfuturesGetTradeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Trade Download Link by Id(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetTradeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetTradeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTradeAsynIdV1`: CfuturesGetTradeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetTradeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTradeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetTradeAsynIdV1Resp**](CfuturesGetTradeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesGetTradeAsynV1

> CfuturesGetTradeAsynV1Resp CfuturesGetTradeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Trade History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.CfuturesGetTradeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CfuturesGetTradeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetTradeAsynV1`: CfuturesGetTradeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CfuturesGetTradeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetTradeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CfuturesGetTradeAsynV1Resp**](CfuturesGetTradeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

