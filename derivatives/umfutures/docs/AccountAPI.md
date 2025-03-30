# \AccountAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesCreateFeeBurnV1**](AccountAPI.md#UfuturesCreateFeeBurnV1) | **Post** /fapi/v1/feeBurn | Toggle BNB Burn On Futures Trade (TRADE)
[**UfuturesGetAccountConfigV1**](AccountAPI.md#UfuturesGetAccountConfigV1) | **Get** /fapi/v1/accountConfig | Futures Account Configuration(USER_DATA)
[**UfuturesGetAccountV2**](AccountAPI.md#UfuturesGetAccountV2) | **Get** /fapi/v2/account | Account Information V2(USER_DATA)
[**UfuturesGetAccountV3**](AccountAPI.md#UfuturesGetAccountV3) | **Get** /fapi/v3/account | Account Information V3(USER_DATA)
[**UfuturesGetApiTradingStatusV1**](AccountAPI.md#UfuturesGetApiTradingStatusV1) | **Get** /fapi/v1/apiTradingStatus | Futures Trading Quantitative Rules Indicators (USER_DATA)
[**UfuturesGetBalanceV2**](AccountAPI.md#UfuturesGetBalanceV2) | **Get** /fapi/v2/balance | Futures Account Balance V2 (USER_DATA)
[**UfuturesGetBalanceV3**](AccountAPI.md#UfuturesGetBalanceV3) | **Get** /fapi/v3/balance | Futures Account Balance V3 (USER_DATA)
[**UfuturesGetCommissionRateV1**](AccountAPI.md#UfuturesGetCommissionRateV1) | **Get** /fapi/v1/commissionRate | User Commission Rate (USER_DATA)
[**UfuturesGetFeeBurnV1**](AccountAPI.md#UfuturesGetFeeBurnV1) | **Get** /fapi/v1/feeBurn | Get BNB Burn Status (USER_DATA)
[**UfuturesGetIncomeAsynIdV1**](AccountAPI.md#UfuturesGetIncomeAsynIdV1) | **Get** /fapi/v1/income/asyn/id | Get Futures Transaction History Download Link by Id (USER_DATA)
[**UfuturesGetIncomeAsynV1**](AccountAPI.md#UfuturesGetIncomeAsynV1) | **Get** /fapi/v1/income/asyn | Get Download Id For Futures Transaction History(USER_DATA)
[**UfuturesGetIncomeV1**](AccountAPI.md#UfuturesGetIncomeV1) | **Get** /fapi/v1/income | Get Income History (USER_DATA)
[**UfuturesGetLeverageBracketV1**](AccountAPI.md#UfuturesGetLeverageBracketV1) | **Get** /fapi/v1/leverageBracket | Notional and Leverage Brackets (USER_DATA)
[**UfuturesGetMultiAssetsMarginV1**](AccountAPI.md#UfuturesGetMultiAssetsMarginV1) | **Get** /fapi/v1/multiAssetsMargin | Get Current Multi-Assets Mode (USER_DATA)
[**UfuturesGetOrderAsynIdV1**](AccountAPI.md#UfuturesGetOrderAsynIdV1) | **Get** /fapi/v1/order/asyn/id | Get Futures Order History Download Link by Id (USER_DATA)
[**UfuturesGetOrderAsynV1**](AccountAPI.md#UfuturesGetOrderAsynV1) | **Get** /fapi/v1/order/asyn | Get Download Id For Futures Order History (USER_DATA)
[**UfuturesGetPositionSideDualV1**](AccountAPI.md#UfuturesGetPositionSideDualV1) | **Get** /fapi/v1/positionSide/dual | Get Current Position Mode(USER_DATA)
[**UfuturesGetRateLimitOrderV1**](AccountAPI.md#UfuturesGetRateLimitOrderV1) | **Get** /fapi/v1/rateLimit/order | Query User Rate Limit (USER_DATA)
[**UfuturesGetSymbolConfigV1**](AccountAPI.md#UfuturesGetSymbolConfigV1) | **Get** /fapi/v1/symbolConfig | Symbol Configuration(USER_DATA)
[**UfuturesGetTradeAsynIdV1**](AccountAPI.md#UfuturesGetTradeAsynIdV1) | **Get** /fapi/v1/trade/asyn/id | Get Futures Trade Download Link by Id(USER_DATA)
[**UfuturesGetTradeAsynV1**](AccountAPI.md#UfuturesGetTradeAsynV1) | **Get** /fapi/v1/trade/asyn | Get Download Id For Futures Trade History (USER_DATA)



## UfuturesCreateFeeBurnV1

> UfuturesCreateFeeBurnV1Resp UfuturesCreateFeeBurnV1(ctx).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On Futures Trade (TRADE)



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
	feeBurn := "feeBurn_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesCreateFeeBurnV1(context.Background()).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesCreateFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateFeeBurnV1`: UfuturesCreateFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesCreateFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeBurn** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesCreateFeeBurnV1Resp**](UfuturesCreateFeeBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetAccountConfigV1

> UfuturesGetAccountConfigV1Resp UfuturesGetAccountConfigV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Configuration(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetAccountConfigV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetAccountConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAccountConfigV1`: UfuturesGetAccountConfigV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetAccountConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAccountConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetAccountConfigV1Resp**](UfuturesGetAccountConfigV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetAccountV2

> UfuturesGetAccountV2Resp UfuturesGetAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information V2(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAccountV2`: UfuturesGetAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetAccountV2Resp**](UfuturesGetAccountV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetAccountV3

> UfuturesGetAccountV3Resp UfuturesGetAccountV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information V3(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetAccountV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetAccountV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetAccountV3`: UfuturesGetAccountV3Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetAccountV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetAccountV3Resp**](UfuturesGetAccountV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetApiTradingStatusV1

> UfuturesGetApiTradingStatusV1Resp UfuturesGetApiTradingStatusV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Futures Trading Quantitative Rules Indicators (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetApiTradingStatusV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetApiTradingStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetApiTradingStatusV1`: UfuturesGetApiTradingStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetApiTradingStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetApiTradingStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetApiTradingStatusV1Resp**](UfuturesGetApiTradingStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetBalanceV2

> []UfuturesGetBalanceV2RespItem UfuturesGetBalanceV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance V2 (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetBalanceV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetBalanceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetBalanceV2`: []UfuturesGetBalanceV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetBalanceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetBalanceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetBalanceV2RespItem**](UfuturesGetBalanceV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetBalanceV3

> []UfuturesGetBalanceV3RespItem UfuturesGetBalanceV3(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Futures Account Balance V3 (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetBalanceV3(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetBalanceV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetBalanceV3`: []UfuturesGetBalanceV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetBalanceV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetBalanceV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetBalanceV3RespItem**](UfuturesGetBalanceV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetCommissionRateV1

> UfuturesGetCommissionRateV1Resp UfuturesGetCommissionRateV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

User Commission Rate (USER_DATA)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetCommissionRateV1`: UfuturesGetCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetCommissionRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetCommissionRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetCommissionRateV1Resp**](UfuturesGetCommissionRateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetFeeBurnV1

> UfuturesGetFeeBurnV1Resp UfuturesGetFeeBurnV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get BNB Burn Status (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetFeeBurnV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetFeeBurnV1`: UfuturesGetFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetFeeBurnV1Resp**](UfuturesGetFeeBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIncomeAsynIdV1

> UfuturesGetIncomeAsynIdV1Resp UfuturesGetIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Transaction History Download Link by Id (USER_DATA)



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
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIncomeAsynIdV1`: UfuturesGetIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetIncomeAsynIdV1Resp**](UfuturesGetIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIncomeAsynV1

> UfuturesGetIncomeAsynV1Resp UfuturesGetIncomeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Transaction History(USER_DATA)



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
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetIncomeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIncomeAsynV1`: UfuturesGetIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIncomeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetIncomeAsynV1Resp**](UfuturesGetIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetIncomeV1

> []UfuturesGetIncomeV1RespItem UfuturesGetIncomeV1(ctx).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Income History (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	incomeType := "incomeType_example" // string | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetIncomeV1`: []UfuturesGetIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **incomeType** | **string** | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE, STRATEGY_UMFUTURES_TRANSFER，FEE_RETURN，BFUSD_REWARD | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int32** |  | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetIncomeV1RespItem**](UfuturesGetIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetLeverageBracketV1

> UfuturesGetLeverageBracketV1Resp UfuturesGetLeverageBracketV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Notional and Leverage Brackets (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetLeverageBracketV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetLeverageBracketV1`: UfuturesGetLeverageBracketV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetLeverageBracketV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetLeverageBracketV1Resp**](UfuturesGetLeverageBracketV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetMultiAssetsMarginV1

> UfuturesGetMultiAssetsMarginV1Resp UfuturesGetMultiAssetsMarginV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Current Multi-Assets Mode (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetMultiAssetsMarginV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetMultiAssetsMarginV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetMultiAssetsMarginV1`: UfuturesGetMultiAssetsMarginV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetMultiAssetsMarginV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetMultiAssetsMarginV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetMultiAssetsMarginV1Resp**](UfuturesGetMultiAssetsMarginV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOrderAsynIdV1

> UfuturesGetOrderAsynIdV1Resp UfuturesGetOrderAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Order History Download Link by Id (USER_DATA)



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
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetOrderAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetOrderAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderAsynIdV1`: UfuturesGetOrderAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetOrderAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOrderAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetOrderAsynIdV1Resp**](UfuturesGetOrderAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetOrderAsynV1

> UfuturesGetOrderAsynV1Resp UfuturesGetOrderAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Order History (USER_DATA)



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
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetOrderAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetOrderAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetOrderAsynV1`: UfuturesGetOrderAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetOrderAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetOrderAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetOrderAsynV1Resp**](UfuturesGetOrderAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetPositionSideDualV1

> UfuturesGetPositionSideDualV1Resp UfuturesGetPositionSideDualV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Current Position Mode(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetPositionSideDualV1`: UfuturesGetPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetPositionSideDualV1Resp**](UfuturesGetPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetRateLimitOrderV1

> []UfuturesGetRateLimitOrderV1RespItem UfuturesGetRateLimitOrderV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query User Rate Limit (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetRateLimitOrderV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetRateLimitOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetRateLimitOrderV1`: []UfuturesGetRateLimitOrderV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetRateLimitOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetRateLimitOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetRateLimitOrderV1RespItem**](UfuturesGetRateLimitOrderV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetSymbolConfigV1

> []UfuturesGetSymbolConfigV1RespItem UfuturesGetSymbolConfigV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Symbol Configuration(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetSymbolConfigV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetSymbolConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetSymbolConfigV1`: []UfuturesGetSymbolConfigV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetSymbolConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetSymbolConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]UfuturesGetSymbolConfigV1RespItem**](UfuturesGetSymbolConfigV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTradeAsynIdV1

> UfuturesGetTradeAsynIdV1Resp UfuturesGetTradeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Trade Download Link by Id(USER_DATA)



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
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetTradeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetTradeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTradeAsynIdV1`: UfuturesGetTradeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetTradeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTradeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetTradeAsynIdV1Resp**](UfuturesGetTradeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesGetTradeAsynV1

> UfuturesGetTradeAsynV1Resp UfuturesGetTradeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For Futures Trade History (USER_DATA)



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
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.UfuturesGetTradeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UfuturesGetTradeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesGetTradeAsynV1`: UfuturesGetTradeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UfuturesGetTradeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesGetTradeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**UfuturesGetTradeAsynV1Resp**](UfuturesGetTradeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

