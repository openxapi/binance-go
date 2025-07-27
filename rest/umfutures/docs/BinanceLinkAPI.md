# \BinanceLinkAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiReferralCustomizationV1**](BinanceLinkAPI.md#CreateApiReferralCustomizationV1) | **Post** /fapi/v1/apiReferral/customization | Customize Id For Client (USER DATA)(For Partner)
[**CreateApiReferralUserCustomizationPAPIV1**](BinanceLinkAPI.md#CreateApiReferralUserCustomizationPAPIV1) | **Post** /papi/v1/apiReferral/userCustomization | Customize Id For Client  (USER DATA)(For client)(PAPI)
[**CreateApiReferralUserCustomizationV1**](BinanceLinkAPI.md#CreateApiReferralUserCustomizationV1) | **Post** /fapi/v1/apiReferral/userCustomization | Customize Id For Client  (USER DATA)(For client)
[**GetApiReferralCustomizationV1**](BinanceLinkAPI.md#GetApiReferralCustomizationV1) | **Get** /fapi/v1/apiReferral/customization | Get Client Email Customized Id (USER DATA)
[**GetApiReferralIfNewUserPAPIV1**](BinanceLinkAPI.md#GetApiReferralIfNewUserPAPIV1) | **Get** /papi/v1/apiReferral/ifNewUser | Query Client If The New User (USER DATA)(PAPI)
[**GetApiReferralIfNewUserV1**](BinanceLinkAPI.md#GetApiReferralIfNewUserV1) | **Get** /fapi/v1/apiReferral/ifNewUser | Query Client If The New User (USER DATA)
[**GetApiReferralOverviewV1**](BinanceLinkAPI.md#GetApiReferralOverviewV1) | **Get** /fapi/v1/apiReferral/overview | Get Rebate Data Overview (USER DATA)
[**GetApiReferralRebateVolV1**](BinanceLinkAPI.md#GetApiReferralRebateVolV1) | **Get** /fapi/v1/apiReferral/rebateVol | Get Rebate Volume (USER DATA)
[**GetApiReferralTradeVolV1**](BinanceLinkAPI.md#GetApiReferralTradeVolV1) | **Get** /fapi/v1/apiReferral/tradeVol | Get User Trade Volume (USER DATA)
[**GetApiReferralTraderNumV1**](BinanceLinkAPI.md#GetApiReferralTraderNumV1) | **Get** /fapi/v1/apiReferral/traderNum | Get Trader Number (USER DATA)
[**GetApiReferralTraderSummaryV1**](BinanceLinkAPI.md#GetApiReferralTraderSummaryV1) | **Get** /fapi/v1/apiReferral/traderSummary | Get Trader Detail (USER DATA)
[**GetApiReferralUserCustomizationPAPIV1**](BinanceLinkAPI.md#GetApiReferralUserCustomizationPAPIV1) | **Get** /papi/v1/apiReferral/userCustomization | Get User’s Customize Id (USER DATA)(PAPI)
[**GetApiReferralUserCustomizationV1**](BinanceLinkAPI.md#GetApiReferralUserCustomizationV1) | **Get** /fapi/v1/apiReferral/userCustomization | Get User’s Customize Id (USER DATA)
[**GetIncomeV1**](BinanceLinkAPI.md#GetIncomeV1) | **Get** /fapi/v1/income | Get Income History(USER DATA)



## CreateApiReferralCustomizationV1

> CreateApiReferralCustomizationV1Resp CreateApiReferralCustomizationV1(ctx).CustomerId(customerId).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Customize Id For Client (USER DATA)(For Partner)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	customerId := "customerId_example" // string |  (default to "")
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.CreateApiReferralCustomizationV1(context.Background()).CustomerId(customerId).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.CreateApiReferralCustomizationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiReferralCustomizationV1`: CreateApiReferralCustomizationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.CreateApiReferralCustomizationV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiReferralCustomizationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** |  | [default to &quot;&quot;]
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateApiReferralCustomizationV1Resp**](CreateApiReferralCustomizationV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiReferralUserCustomizationPAPIV1

> CreateApiReferralUserCustomizationV1Resp CreateApiReferralUserCustomizationPAPIV1(ctx).BrokerId(brokerId).CustomerId(customerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Customize Id For Client  (USER DATA)(For client)(PAPI)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	brokerId := "brokerId_example" // string |  (default to "")
	customerId := "customerId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.CreateApiReferralUserCustomizationPAPIV1(context.Background()).BrokerId(brokerId).CustomerId(customerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.CreateApiReferralUserCustomizationPAPIV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiReferralUserCustomizationPAPIV1`: CreateApiReferralUserCustomizationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.CreateApiReferralUserCustomizationPAPIV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiReferralUserCustomizationPAPIV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brokerId** | **string** |  | [default to &quot;&quot;]
 **customerId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateApiReferralUserCustomizationV1Resp**](CreateApiReferralUserCustomizationV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiReferralUserCustomizationV1

> CreateApiReferralUserCustomizationV1Resp CreateApiReferralUserCustomizationV1(ctx).BrokerId(brokerId).CustomerId(customerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Customize Id For Client  (USER DATA)(For client)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	brokerId := "brokerId_example" // string |  (default to "")
	customerId := "customerId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.CreateApiReferralUserCustomizationV1(context.Background()).BrokerId(brokerId).CustomerId(customerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.CreateApiReferralUserCustomizationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiReferralUserCustomizationV1`: CreateApiReferralUserCustomizationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.CreateApiReferralUserCustomizationV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiReferralUserCustomizationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brokerId** | **string** |  | [default to &quot;&quot;]
 **customerId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateApiReferralUserCustomizationV1Resp**](CreateApiReferralUserCustomizationV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralCustomizationV1

> []GetApiReferralCustomizationV1RespItem GetApiReferralCustomizationV1(ctx).Timestamp(timestamp).CustomerId(customerId).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Client Email Customized Id (USER DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	customerId := "customerId_example" // string |  (optional) (default to "")
	email := "email_example" // string |  (optional) (default to "")
	page := int64(789) // int64 | default 1 (optional) (default to 1)
	limit := int64(789) // int64 | items num of one page，default 100，max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralCustomizationV1(context.Background()).Timestamp(timestamp).CustomerId(customerId).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralCustomizationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralCustomizationV1`: []GetApiReferralCustomizationV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralCustomizationV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralCustomizationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **customerId** | **string** |  | [default to &quot;&quot;]
 **email** | **string** |  | [default to &quot;&quot;]
 **page** | **int64** | default 1 | [default to 1]
 **limit** | **int64** | items num of one page，default 100，max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetApiReferralCustomizationV1RespItem**](GetApiReferralCustomizationV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralIfNewUserPAPIV1

> GetApiReferralIfNewUserV1Resp GetApiReferralIfNewUserPAPIV1(ctx).BrokerId(brokerId).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Query Client If The New User (USER DATA)(PAPI)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	brokerId := "brokerId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1:USDT-margined Futures,  2: Coin-margined Futures ; Default：1:USDT-margined Futures (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralIfNewUserPAPIV1(context.Background()).BrokerId(brokerId).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralIfNewUserPAPIV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralIfNewUserPAPIV1`: GetApiReferralIfNewUserV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralIfNewUserPAPIV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralIfNewUserPAPIV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brokerId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1:USDT-margined Futures,  2: Coin-margined Futures ; Default：1:USDT-margined Futures | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetApiReferralIfNewUserV1Resp**](GetApiReferralIfNewUserV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralIfNewUserV1

> GetApiReferralIfNewUserV1Resp GetApiReferralIfNewUserV1(ctx).BrokerId(brokerId).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Query Client If The New User (USER DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	brokerId := "brokerId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1:USDT-margined Futures，2: Coin-margined Futures; Default：1:USDT-margined Futures (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralIfNewUserV1(context.Background()).BrokerId(brokerId).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralIfNewUserV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralIfNewUserV1`: GetApiReferralIfNewUserV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralIfNewUserV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralIfNewUserV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brokerId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1:USDT-margined Futures，2: Coin-margined Futures; Default：1:USDT-margined Futures | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetApiReferralIfNewUserV1Resp**](GetApiReferralIfNewUserV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralOverviewV1

> GetApiReferralOverviewV1Resp GetApiReferralOverviewV1(ctx).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Get Rebate Data Overview (USER DATA)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralOverviewV1(context.Background()).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralOverviewV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralOverviewV1`: GetApiReferralOverviewV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralOverviewV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralOverviewV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetApiReferralOverviewV1Resp**](GetApiReferralOverviewV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralRebateVolV1

> []GetApiReferralRebateVolV1RespItem GetApiReferralRebateVolV1(ctx).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Rebate Volume (USER DATA)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | default 500, max 1000 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralRebateVolV1(context.Background()).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralRebateVolV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralRebateVolV1`: []GetApiReferralRebateVolV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralRebateVolV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralRebateVolV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | default 500, max 1000 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetApiReferralRebateVolV1RespItem**](GetApiReferralRebateVolV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralTradeVolV1

> []GetApiReferralTradeVolV1RespItem GetApiReferralTradeVolV1(ctx).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get User Trade Volume (USER DATA)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | default 500, max 1000 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralTradeVolV1(context.Background()).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralTradeVolV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralTradeVolV1`: []GetApiReferralTradeVolV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralTradeVolV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralTradeVolV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | default 500, max 1000 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetApiReferralTradeVolV1RespItem**](GetApiReferralTradeVolV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralTraderNumV1

> []GetApiReferralTraderNumV1RespItem GetApiReferralTraderNumV1(ctx).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Trader Number (USER DATA)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | default 500, max 1000 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralTraderNumV1(context.Background()).Timestamp(timestamp).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralTraderNumV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralTraderNumV1`: []GetApiReferralTraderNumV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralTraderNumV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralTraderNumV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **type_** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | default 500, max 1000 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetApiReferralTraderNumV1RespItem**](GetApiReferralTraderNumV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralTraderSummaryV1

> []GetApiReferralTraderSummaryV1RespItem GetApiReferralTraderSummaryV1(ctx).Timestamp(timestamp).CustomerId(customerId).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Trader Detail (USER DATA)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	customerId := "customerId_example" // string |  (optional) (default to "")
	type_ := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | default 500, max 1000 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralTraderSummaryV1(context.Background()).Timestamp(timestamp).CustomerId(customerId).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralTraderSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralTraderSummaryV1`: []GetApiReferralTraderSummaryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralTraderSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralTraderSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **customerId** | **string** |  | [default to &quot;&quot;]
 **type_** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | default 500, max 1000 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetApiReferralTraderSummaryV1RespItem**](GetApiReferralTraderSummaryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralUserCustomizationPAPIV1

> GetApiReferralUserCustomizationV1Resp GetApiReferralUserCustomizationPAPIV1(ctx).BrokerId(brokerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get User’s Customize Id (USER DATA)(PAPI)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	brokerId := "brokerId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralUserCustomizationPAPIV1(context.Background()).BrokerId(brokerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralUserCustomizationPAPIV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralUserCustomizationPAPIV1`: GetApiReferralUserCustomizationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralUserCustomizationPAPIV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralUserCustomizationPAPIV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brokerId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetApiReferralUserCustomizationV1Resp**](GetApiReferralUserCustomizationV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiReferralUserCustomizationV1

> GetApiReferralUserCustomizationV1Resp GetApiReferralUserCustomizationV1(ctx).BrokerId(brokerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get User’s Customize Id (USER DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	brokerId := "brokerId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetApiReferralUserCustomizationV1(context.Background()).BrokerId(brokerId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetApiReferralUserCustomizationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiReferralUserCustomizationV1`: GetApiReferralUserCustomizationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetApiReferralUserCustomizationV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiReferralUserCustomizationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brokerId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetApiReferralUserCustomizationV1Resp**](GetApiReferralUserCustomizationV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomeV1

> []GetIncomeV1RespItem GetIncomeV1(ctx).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Income History(USER DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/umfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	incomeType := "incomeType_example" // string | &#34;TRANSFER&#34;，&#34;WELCOME_BONUS&#34;, &#34;REALIZED_PNL&#34;，&#34;FUNDING_FEE&#34;, &#34;COMMISSION&#34;, and &#34;INSURANCE_CLEAR&#34; (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinanceLinkAPI.GetIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinanceLinkAPI.GetIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomeV1`: []GetIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `BinanceLinkAPI.GetIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **incomeType** | **string** | &amp;#34;TRANSFER&amp;#34;，&amp;#34;WELCOME_BONUS&amp;#34;, &amp;#34;REALIZED_PNL&amp;#34;，&amp;#34;FUNDING_FEE&amp;#34;, &amp;#34;COMMISSION&amp;#34;, and &amp;#34;INSURANCE_CLEAR&amp;#34; | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetIncomeV1RespItem**](GetIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

