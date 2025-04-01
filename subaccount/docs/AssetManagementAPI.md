# \AssetManagementAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreateSubAccountFuturesInternalTransferV1**](AssetManagementAPI.md#SubaccountCreateSubAccountFuturesInternalTransferV1) | **Post** /sapi/v1/sub-account/futures/internalTransfer | Sub-account Futures Asset Transfer(For Master Account)
[**SubaccountCreateSubAccountFuturesMovePositionV1**](AssetManagementAPI.md#SubaccountCreateSubAccountFuturesMovePositionV1) | **Post** /sapi/v1/sub-account/futures/move-position | Move Position for Sub-account (For Master Account)
[**SubaccountCreateSubAccountFuturesTransferV1**](AssetManagementAPI.md#SubaccountCreateSubAccountFuturesTransferV1) | **Post** /sapi/v1/sub-account/futures/transfer | Futures Transfer for Sub-account(For Master Account)
[**SubaccountCreateSubAccountMarginTransferV1**](AssetManagementAPI.md#SubaccountCreateSubAccountMarginTransferV1) | **Post** /sapi/v1/sub-account/margin/transfer | Margin Transfer for Sub-account(For Master Account)
[**SubaccountCreateSubAccountTransferSubToMasterV1**](AssetManagementAPI.md#SubaccountCreateSubAccountTransferSubToMasterV1) | **Post** /sapi/v1/sub-account/transfer/subToMaster | Transfer to Master(For Sub-account)
[**SubaccountCreateSubAccountTransferSubToSubV1**](AssetManagementAPI.md#SubaccountCreateSubAccountTransferSubToSubV1) | **Post** /sapi/v1/sub-account/transfer/subToSub | Transfer to Sub-account of Same Master(For Sub-account)
[**SubaccountCreateSubAccountUniversalTransferV1**](AssetManagementAPI.md#SubaccountCreateSubAccountUniversalTransferV1) | **Post** /sapi/v1/sub-account/universalTransfer | Universal Transfer(For Master Account)
[**SubaccountGetCapitalDepositSubAddressV1**](AssetManagementAPI.md#SubaccountGetCapitalDepositSubAddressV1) | **Get** /sapi/v1/capital/deposit/subAddress | Get Sub-account Deposit Address(For Master Account)
[**SubaccountGetCapitalDepositSubHisrecV1**](AssetManagementAPI.md#SubaccountGetCapitalDepositSubHisrecV1) | **Get** /sapi/v1/capital/deposit/subHisrec | Get Sub-account Deposit History(For Master Account)
[**SubaccountGetSubAccountAssetsV3**](AssetManagementAPI.md#SubaccountGetSubAccountAssetsV3) | **Get** /sapi/v3/sub-account/assets | Query Sub-account Assets(For Master Account)
[**SubaccountGetSubAccountAssetsV4**](AssetManagementAPI.md#SubaccountGetSubAccountAssetsV4) | **Get** /sapi/v4/sub-account/assets | Query Sub-account Assets (For Master Account)(USER_DATA)
[**SubaccountGetSubAccountFuturesAccountSummaryV1**](AssetManagementAPI.md#SubaccountGetSubAccountFuturesAccountSummaryV1) | **Get** /sapi/v1/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account(For Master Account)
[**SubaccountGetSubAccountFuturesAccountSummaryV2**](AssetManagementAPI.md#SubaccountGetSubAccountFuturesAccountSummaryV2) | **Get** /sapi/v2/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account V2(For Master Account)
[**SubaccountGetSubAccountFuturesAccountV1**](AssetManagementAPI.md#SubaccountGetSubAccountFuturesAccountV1) | **Get** /sapi/v1/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account(For Master Account)
[**SubaccountGetSubAccountFuturesAccountV2**](AssetManagementAPI.md#SubaccountGetSubAccountFuturesAccountV2) | **Get** /sapi/v2/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account V2(For Master Account)
[**SubaccountGetSubAccountFuturesInternalTransferV1**](AssetManagementAPI.md#SubaccountGetSubAccountFuturesInternalTransferV1) | **Get** /sapi/v1/sub-account/futures/internalTransfer | Query Sub-account Futures Asset Transfer History(For Master Account)
[**SubaccountGetSubAccountFuturesMovePositionV1**](AssetManagementAPI.md#SubaccountGetSubAccountFuturesMovePositionV1) | **Get** /sapi/v1/sub-account/futures/move-position | Get Move Position History for Sub-account (For Master Account)
[**SubaccountGetSubAccountMarginAccountSummaryV1**](AssetManagementAPI.md#SubaccountGetSubAccountMarginAccountSummaryV1) | **Get** /sapi/v1/sub-account/margin/accountSummary | Get Summary of Sub-account&#39;s Margin Account(For Master Account)
[**SubaccountGetSubAccountMarginAccountV1**](AssetManagementAPI.md#SubaccountGetSubAccountMarginAccountV1) | **Get** /sapi/v1/sub-account/margin/account | Get Detail on Sub-account&#39;s Margin Account(For Master Account)
[**SubaccountGetSubAccountSpotSummaryV1**](AssetManagementAPI.md#SubaccountGetSubAccountSpotSummaryV1) | **Get** /sapi/v1/sub-account/spotSummary | Query Sub-account Spot Assets Summary(For Master Account)
[**SubaccountGetSubAccountSubTransferHistoryV1**](AssetManagementAPI.md#SubaccountGetSubAccountSubTransferHistoryV1) | **Get** /sapi/v1/sub-account/sub/transfer/history | Query Sub-account Spot Asset Transfer History(For Master Account)
[**SubaccountGetSubAccountTransferSubUserHistoryV1**](AssetManagementAPI.md#SubaccountGetSubAccountTransferSubUserHistoryV1) | **Get** /sapi/v1/sub-account/transfer/subUserHistory | Sub-account Transfer History(For Sub-account)
[**SubaccountGetSubAccountUniversalTransferV1**](AssetManagementAPI.md#SubaccountGetSubAccountUniversalTransferV1) | **Get** /sapi/v1/sub-account/universalTransfer | Query Universal Transfer History(For Master Account)



## SubaccountCreateSubAccountFuturesInternalTransferV1

> SubaccountCreateSubAccountFuturesInternalTransferV1Resp SubaccountCreateSubAccountFuturesInternalTransferV1(ctx).Amount(amount).Asset(asset).FromEmail(fromEmail).FuturesType(futuresType).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()

Sub-account Futures Asset Transfer(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	fromEmail := "fromEmail_example" // string |  (default to "")
	futuresType := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	toEmail := "toEmail_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountCreateSubAccountFuturesInternalTransferV1(context.Background()).Amount(amount).Asset(asset).FromEmail(fromEmail).FuturesType(futuresType).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountCreateSubAccountFuturesInternalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesInternalTransferV1`: SubaccountCreateSubAccountFuturesInternalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountCreateSubAccountFuturesInternalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountFuturesInternalTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **fromEmail** | **string** |  | [default to &quot;&quot;]
 **futuresType** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **toEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountFuturesInternalTransferV1Resp**](SubaccountCreateSubAccountFuturesInternalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountFuturesMovePositionV1

> SubaccountCreateSubAccountFuturesMovePositionV1Resp SubaccountCreateSubAccountFuturesMovePositionV1(ctx).FromUserEmail(fromUserEmail).OrderArgs(orderArgs).ProductType(productType).Timestamp(timestamp).ToUserEmail(toUserEmail).RecvWindow(recvWindow).Execute()

Move Position for Sub-account (For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	fromUserEmail := "fromUserEmail_example" // string |  (default to "")
	orderArgs := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | 
	productType := "productType_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toUserEmail := "toUserEmail_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountCreateSubAccountFuturesMovePositionV1(context.Background()).FromUserEmail(fromUserEmail).OrderArgs(orderArgs).ProductType(productType).Timestamp(timestamp).ToUserEmail(toUserEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountCreateSubAccountFuturesMovePositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesMovePositionV1`: SubaccountCreateSubAccountFuturesMovePositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountCreateSubAccountFuturesMovePositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountFuturesMovePositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromUserEmail** | **string** |  | [default to &quot;&quot;]
 **orderArgs** | **[]map[string]interface{}** |  | 
 **productType** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toUserEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountFuturesMovePositionV1Resp**](SubaccountCreateSubAccountFuturesMovePositionV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountFuturesTransferV1

> SubaccountCreateSubAccountFuturesTransferV1Resp SubaccountCreateSubAccountFuturesTransferV1(ctx).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Futures Transfer for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountCreateSubAccountFuturesTransferV1(context.Background()).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountCreateSubAccountFuturesTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesTransferV1`: SubaccountCreateSubAccountFuturesTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountCreateSubAccountFuturesTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountFuturesTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountFuturesTransferV1Resp**](SubaccountCreateSubAccountFuturesTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountMarginTransferV1

> SubaccountCreateSubAccountMarginTransferV1Resp SubaccountCreateSubAccountMarginTransferV1(ctx).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Margin Transfer for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := int32(56) // int32 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountCreateSubAccountMarginTransferV1(context.Background()).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountCreateSubAccountMarginTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountMarginTransferV1`: SubaccountCreateSubAccountMarginTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountCreateSubAccountMarginTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountMarginTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountMarginTransferV1Resp**](SubaccountCreateSubAccountMarginTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountTransferSubToMasterV1

> SubaccountCreateSubAccountTransferSubToMasterV1Resp SubaccountCreateSubAccountTransferSubToMasterV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Transfer to Master(For Sub-account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountCreateSubAccountTransferSubToMasterV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountCreateSubAccountTransferSubToMasterV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountTransferSubToMasterV1`: SubaccountCreateSubAccountTransferSubToMasterV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountCreateSubAccountTransferSubToMasterV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountTransferSubToMasterV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountTransferSubToMasterV1Resp**](SubaccountCreateSubAccountTransferSubToMasterV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountTransferSubToSubV1

> SubaccountCreateSubAccountTransferSubToSubV1Resp SubaccountCreateSubAccountTransferSubToSubV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()

Transfer to Sub-account of Same Master(For Sub-account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toEmail := "toEmail_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountCreateSubAccountTransferSubToSubV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountCreateSubAccountTransferSubToSubV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountTransferSubToSubV1`: SubaccountCreateSubAccountTransferSubToSubV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountCreateSubAccountTransferSubToSubV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountTransferSubToSubV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountTransferSubToSubV1Resp**](SubaccountCreateSubAccountTransferSubToSubV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountUniversalTransferV1

> SubaccountCreateSubAccountUniversalTransferV1Resp SubaccountCreateSubAccountUniversalTransferV1(ctx).Amount(amount).Asset(asset).FromAccountType(fromAccountType).Timestamp(timestamp).ToAccountType(toAccountType).ClientTranId(clientTranId).FromEmail(fromEmail).RecvWindow(recvWindow).Symbol(symbol).ToEmail(toEmail).Execute()

Universal Transfer(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	fromAccountType := "fromAccountType_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toAccountType := "toAccountType_example" // string |  (default to "")
	clientTranId := "clientTranId_example" // string |  (optional) (default to "")
	fromEmail := "fromEmail_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")
	toEmail := "toEmail_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountCreateSubAccountUniversalTransferV1(context.Background()).Amount(amount).Asset(asset).FromAccountType(fromAccountType).Timestamp(timestamp).ToAccountType(toAccountType).ClientTranId(clientTranId).FromEmail(fromEmail).RecvWindow(recvWindow).Symbol(symbol).ToEmail(toEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountCreateSubAccountUniversalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountUniversalTransferV1`: SubaccountCreateSubAccountUniversalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountCreateSubAccountUniversalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountUniversalTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **fromAccountType** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toAccountType** | **string** |  | [default to &quot;&quot;]
 **clientTranId** | **string** |  | [default to &quot;&quot;]
 **fromEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **toEmail** | **string** |  | [default to &quot;&quot;]

### Return type

[**SubaccountCreateSubAccountUniversalTransferV1Resp**](SubaccountCreateSubAccountUniversalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetCapitalDepositSubAddressV1

> SubaccountGetCapitalDepositSubAddressV1Resp SubaccountGetCapitalDepositSubAddressV1(ctx).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Get Sub-account Deposit Address(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | Sub account email (default to "")
	coin := "coin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	network := "network_example" // string |  (optional) (default to "")
	amount := "amount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetCapitalDepositSubAddressV1(context.Background()).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetCapitalDepositSubAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetCapitalDepositSubAddressV1`: SubaccountGetCapitalDepositSubAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetCapitalDepositSubAddressV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetCapitalDepositSubAddressV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetCapitalDepositSubAddressV1Resp**](SubaccountGetCapitalDepositSubAddressV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetCapitalDepositSubHisrecV1

> []SubaccountGetCapitalDepositSubHisrecV1RespItem SubaccountGetCapitalDepositSubHisrecV1(ctx).Email(email).Timestamp(timestamp).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).TxId(txId).Execute()

Get Sub-account Deposit History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | Sub account email (default to "")
	timestamp := int64(789) // int64 | 
	coin := "coin_example" // string |  (optional) (default to "")
	status := int32(56) // int32 | 0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success) (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 | default:0 (optional)
	recvWindow := int64(789) // int64 |  (optional)
	txId := "txId_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetCapitalDepositSubHisrecV1(context.Background()).Email(email).Timestamp(timestamp).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).TxId(txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetCapitalDepositSubHisrecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetCapitalDepositSubHisrecV1`: []SubaccountGetCapitalDepositSubHisrecV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetCapitalDepositSubHisrecV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetCapitalDepositSubHisrecV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **status** | **int32** | 0(0:pending,6: credited but cannot withdraw,7:Wrong Deposit,8:Waiting User confirm,1:success) | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** | default:0 | 
 **recvWindow** | **int64** |  | 
 **txId** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]SubaccountGetCapitalDepositSubHisrecV1RespItem**](SubaccountGetCapitalDepositSubHisrecV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountAssetsV3

> SubaccountGetSubAccountAssetsV3Resp SubaccountGetSubAccountAssetsV3(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Assets(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | Sub account email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountAssetsV3(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountAssetsV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountAssetsV3`: SubaccountGetSubAccountAssetsV3Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountAssetsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountAssetsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountAssetsV3Resp**](SubaccountGetSubAccountAssetsV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountAssetsV4

> SubaccountGetSubAccountAssetsV4Resp SubaccountGetSubAccountAssetsV4(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Assets (For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | Sub Account Email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountAssetsV4(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountAssetsV4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountAssetsV4`: SubaccountGetSubAccountAssetsV4Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountAssetsV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountAssetsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub Account Email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountAssetsV4Resp**](SubaccountGetSubAccountAssetsV4Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesAccountSummaryV1

> SubaccountGetSubAccountFuturesAccountSummaryV1Resp SubaccountGetSubAccountFuturesAccountSummaryV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Futures Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountFuturesAccountSummaryV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountSummaryV1`: SubaccountGetSubAccountFuturesAccountSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesAccountSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesAccountSummaryV1Resp**](SubaccountGetSubAccountFuturesAccountSummaryV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesAccountSummaryV2

> SubaccountGetSubAccountFuturesAccountSummaryV2Resp SubaccountGetSubAccountFuturesAccountSummaryV2(ctx).FuturesType(futuresType).Timestamp(timestamp).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Futures Account V2(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	page := int32(56) // int32 | default:1 (optional)
	limit := int32(56) // int32 | default:10, max:20 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountFuturesAccountSummaryV2(context.Background()).FuturesType(futuresType).Timestamp(timestamp).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountSummaryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountSummaryV2`: SubaccountGetSubAccountFuturesAccountSummaryV2Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountSummaryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesAccountSummaryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **page** | **int32** | default:1 | 
 **limit** | **int32** | default:10, max:20 | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesAccountSummaryV2Resp**](SubaccountGetSubAccountFuturesAccountSummaryV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesAccountV1

> SubaccountGetSubAccountFuturesAccountV1Resp SubaccountGetSubAccountFuturesAccountV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Futures Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account#email-address\">Sub-account email</a> (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountFuturesAccountV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountV1`: SubaccountGetSubAccountFuturesAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesAccountV1Resp**](SubaccountGetSubAccountFuturesAccountV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesAccountV2

> SubaccountGetSubAccountFuturesAccountV2Resp SubaccountGetSubAccountFuturesAccountV2(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Futures Account V2(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2#email-address\">Sub-account email</a> (default to "")
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountFuturesAccountV2(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountV2`: SubaccountGetSubAccountFuturesAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountFuturesAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesAccountV2Resp**](SubaccountGetSubAccountFuturesAccountV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesInternalTransferV1

> SubaccountGetSubAccountFuturesInternalTransferV1Resp SubaccountGetSubAccountFuturesInternalTransferV1(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account Futures Asset Transfer History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Query-Sub-account-Futures-Asset-Transfer-History#email-address\">Sub-account email</a> (default to "")
	futuresType := int64(789) // int64 | 1:USDT-margined Futures，2: Coin-margined Futures
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 | Cannot be earlier than 1 month ago (optional)
	endTime := int64(789) // int64 |  (optional)
	page := int32(56) // int32 | Default value: 1 (optional)
	limit := int32(56) // int32 | Default value: 50, Max value: 500 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountFuturesInternalTransferV1(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountFuturesInternalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesInternalTransferV1`: SubaccountGetSubAccountFuturesInternalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountFuturesInternalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesInternalTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Query-Sub-account-Futures-Asset-Transfer-History#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **futuresType** | **int64** | 1:USDT-margined Futures，2: Coin-margined Futures | 
 **timestamp** | **int64** |  | 
 **startTime** | **int64** | Cannot be earlier than 1 month ago | 
 **endTime** | **int64** |  | 
 **page** | **int32** | Default value: 1 | 
 **limit** | **int32** | Default value: 50, Max value: 500 | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesInternalTransferV1Resp**](SubaccountGetSubAccountFuturesInternalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesMovePositionV1

> SubaccountGetSubAccountFuturesMovePositionV1Resp SubaccountGetSubAccountFuturesMovePositionV1(ctx).Symbol(symbol).Page(page).Row(row).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Move Position History for Sub-account (For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	page := int32(56) // int32 | 
	row := int32(56) // int32 | 
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountFuturesMovePositionV1(context.Background()).Symbol(symbol).Page(page).Row(row).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountFuturesMovePositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesMovePositionV1`: SubaccountGetSubAccountFuturesMovePositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountFuturesMovePositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesMovePositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **page** | **int32** |  | 
 **row** | **int32** |  | 
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesMovePositionV1Resp**](SubaccountGetSubAccountFuturesMovePositionV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountMarginAccountSummaryV1

> SubaccountGetSubAccountMarginAccountSummaryV1Resp SubaccountGetSubAccountMarginAccountSummaryV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Margin Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountMarginAccountSummaryV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountMarginAccountSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountMarginAccountSummaryV1`: SubaccountGetSubAccountMarginAccountSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountMarginAccountSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountMarginAccountSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountMarginAccountSummaryV1Resp**](SubaccountGetSubAccountMarginAccountSummaryV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountMarginAccountV1

> SubaccountGetSubAccountMarginAccountV1Resp SubaccountGetSubAccountMarginAccountV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Margin Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Margin-Account#email-address\">Sub-account email</a> (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountMarginAccountV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountMarginAccountV1`: SubaccountGetSubAccountMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Margin-Account#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountMarginAccountV1Resp**](SubaccountGetSubAccountMarginAccountV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountSpotSummaryV1

> SubaccountGetSubAccountSpotSummaryV1Resp SubaccountGetSubAccountSpotSummaryV1(ctx).Timestamp(timestamp).Email(email).Page(page).Size(size).RecvWindow(recvWindow).Execute()

Query Sub-account Spot Assets Summary(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	email := "email_example" // string | Sub account email (optional) (default to "")
	page := int64(789) // int64 | default 1 (optional) (default to 1)
	size := int64(789) // int64 | default 10, max 20 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountSpotSummaryV1(context.Background()).Timestamp(timestamp).Email(email).Page(page).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountSpotSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountSpotSummaryV1`: SubaccountGetSubAccountSpotSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountSpotSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountSpotSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **page** | **int64** | default 1 | [default to 1]
 **size** | **int64** | default 10, max 20 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountSpotSummaryV1Resp**](SubaccountGetSubAccountSpotSummaryV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountSubTransferHistoryV1

> []SubaccountGetSubAccountSubTransferHistoryV1RespItem SubaccountGetSubAccountSubTransferHistoryV1(ctx).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account Spot Asset Transfer History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	fromEmail := "fromEmail_example" // string |  (optional) (default to "")
	toEmail := "toEmail_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	page := int32(56) // int32 | Default value: 1 (optional)
	limit := int32(56) // int32 | Default value: 500 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountSubTransferHistoryV1(context.Background()).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountSubTransferHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountSubTransferHistoryV1`: []SubaccountGetSubAccountSubTransferHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountSubTransferHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountSubTransferHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **fromEmail** | **string** |  | [default to &quot;&quot;]
 **toEmail** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int32** | Default value: 1 | 
 **limit** | **int32** | Default value: 500 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]SubaccountGetSubAccountSubTransferHistoryV1RespItem**](SubaccountGetSubAccountSubTransferHistoryV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountTransferSubUserHistoryV1

> []SubaccountGetSubAccountTransferSubUserHistoryV1RespItem SubaccountGetSubAccountTransferSubUserHistoryV1(ctx).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).ReturnFailHistory(returnFailHistory).RecvWindow(recvWindow).Execute()

Sub-account Transfer History(For Sub-account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string | If not sent, result of all assets will be returned (optional) (default to "")
	type_ := int32(56) // int32 | 1: transfer in, 2: transfer out (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500 (optional) (default to 500)
	returnFailHistory := true // bool | Default `False`, return PROCESS and SUCCESS status history; If `True`,return PROCESS and SUCCESS and FAILURE status history (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountTransferSubUserHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).ReturnFailHistory(returnFailHistory).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountTransferSubUserHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountTransferSubUserHistoryV1`: []SubaccountGetSubAccountTransferSubUserHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountTransferSubUserHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountTransferSubUserHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** | If not sent, result of all assets will be returned | [default to &quot;&quot;]
 **type_** | **int32** | 1: transfer in, 2: transfer out | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500 | [default to 500]
 **returnFailHistory** | **bool** | Default &#x60;False&#x60;, return PROCESS and SUCCESS status history; If &#x60;True&#x60;,return PROCESS and SUCCESS and FAILURE status history | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]SubaccountGetSubAccountTransferSubUserHistoryV1RespItem**](SubaccountGetSubAccountTransferSubUserHistoryV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountUniversalTransferV1

> SubaccountGetSubAccountUniversalTransferV1Resp SubaccountGetSubAccountUniversalTransferV1(ctx).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Universal Transfer History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	fromEmail := "fromEmail_example" // string |  (optional) (default to "")
	toEmail := "toEmail_example" // string |  (optional) (default to "")
	clientTranId := "clientTranId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	page := int32(56) // int32 | Default 1 (optional) (default to 1)
	limit := int32(56) // int32 | Default 500, Max 500 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetManagementAPI.SubaccountGetSubAccountUniversalTransferV1(context.Background()).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetManagementAPI.SubaccountGetSubAccountUniversalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountUniversalTransferV1`: SubaccountGetSubAccountUniversalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetManagementAPI.SubaccountGetSubAccountUniversalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountUniversalTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **fromEmail** | **string** |  | [default to &quot;&quot;]
 **toEmail** | **string** |  | [default to &quot;&quot;]
 **clientTranId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int32** | Default 1 | [default to 1]
 **limit** | **int32** | Default 500, Max 500 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountUniversalTransferV1Resp**](SubaccountGetSubAccountUniversalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

