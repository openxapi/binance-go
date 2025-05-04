# \SubAccountAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagedSubaccountDepositV1**](SubAccountAPI.md#CreateManagedSubaccountDepositV1) | **Post** /sapi/v1/managed-subaccount/deposit | Deposit Assets Into The Managed Sub-account(For Investor Master Account)
[**CreateManagedSubaccountWithdrawV1**](SubAccountAPI.md#CreateManagedSubaccountWithdrawV1) | **Post** /sapi/v1/managed-subaccount/withdraw | Withdrawl Assets From The Managed Sub-account(For Investor Master Account)
[**CreateSubAccountBlvtEnableV1**](SubAccountAPI.md#CreateSubAccountBlvtEnableV1) | **Post** /sapi/v1/sub-account/blvt/enable | Enable Leverage Token for Sub-account(For Master Account)
[**CreateSubAccountEoptionsEnableV1**](SubAccountAPI.md#CreateSubAccountEoptionsEnableV1) | **Post** /sapi/v1/sub-account/eoptions/enable | Enable Options for Sub-account(For Master Account)(USER_DATA)
[**CreateSubAccountFuturesEnableV1**](SubAccountAPI.md#CreateSubAccountFuturesEnableV1) | **Post** /sapi/v1/sub-account/futures/enable | Enable Futures for Sub-account(For Master Account)
[**CreateSubAccountFuturesInternalTransferV1**](SubAccountAPI.md#CreateSubAccountFuturesInternalTransferV1) | **Post** /sapi/v1/sub-account/futures/internalTransfer | Sub-account Futures Asset Transfer(For Master Account)
[**CreateSubAccountFuturesMovePositionV1**](SubAccountAPI.md#CreateSubAccountFuturesMovePositionV1) | **Post** /sapi/v1/sub-account/futures/move-position | Move Position for Sub-account (For Master Account)
[**CreateSubAccountFuturesTransferV1**](SubAccountAPI.md#CreateSubAccountFuturesTransferV1) | **Post** /sapi/v1/sub-account/futures/transfer | Futures Transfer for Sub-account(For Master Account)
[**CreateSubAccountMarginEnableV1**](SubAccountAPI.md#CreateSubAccountMarginEnableV1) | **Post** /sapi/v1/sub-account/margin/enable | Enable Margin for Sub-account(For Master Account)
[**CreateSubAccountMarginTransferV1**](SubAccountAPI.md#CreateSubAccountMarginTransferV1) | **Post** /sapi/v1/sub-account/margin/transfer | Margin Transfer for Sub-account(For Master Account)
[**CreateSubAccountSubAccountApiIpRestrictionV2**](SubAccountAPI.md#CreateSubAccountSubAccountApiIpRestrictionV2) | **Post** /sapi/v2/sub-account/subAccountApi/ipRestriction | Add IP Restriction for Sub-Account API key(For Master Account)
[**CreateSubAccountTransferSubToMasterV1**](SubAccountAPI.md#CreateSubAccountTransferSubToMasterV1) | **Post** /sapi/v1/sub-account/transfer/subToMaster | Transfer to Master(For Sub-account)
[**CreateSubAccountTransferSubToSubV1**](SubAccountAPI.md#CreateSubAccountTransferSubToSubV1) | **Post** /sapi/v1/sub-account/transfer/subToSub | Transfer to Sub-account of Same Master(For Sub-account)
[**CreateSubAccountUniversalTransferV1**](SubAccountAPI.md#CreateSubAccountUniversalTransferV1) | **Post** /sapi/v1/sub-account/universalTransfer | Universal Transfer(For Master Account)
[**CreateSubAccountVirtualSubAccountV1**](SubAccountAPI.md#CreateSubAccountVirtualSubAccountV1) | **Post** /sapi/v1/sub-account/virtualSubAccount | Create a Virtual Sub-account(For Master Account)
[**DeleteSubAccountSubAccountApiIpRestrictionIpListV1**](SubAccountAPI.md#DeleteSubAccountSubAccountApiIpRestrictionIpListV1) | **Delete** /sapi/v1/sub-account/subAccountApi/ipRestriction/ipList | Delete IP List For a Sub-account API Key(For Master Account)
[**GetCapitalDepositSubAddressV1**](SubAccountAPI.md#GetCapitalDepositSubAddressV1) | **Get** /sapi/v1/capital/deposit/subAddress | Get Sub-account Deposit Address(For Master Account)
[**GetCapitalDepositSubHisrecV1**](SubAccountAPI.md#GetCapitalDepositSubHisrecV1) | **Get** /sapi/v1/capital/deposit/subHisrec | Get Sub-account Deposit History(For Master Account)
[**GetManagedSubaccountAccountSnapshotV1**](SubAccountAPI.md#GetManagedSubaccountAccountSnapshotV1) | **Get** /sapi/v1/managed-subaccount/accountSnapshot | Query Managed Sub-account Snapshot(For Investor Master Account)
[**GetManagedSubaccountAssetV1**](SubAccountAPI.md#GetManagedSubaccountAssetV1) | **Get** /sapi/v1/managed-subaccount/asset | Query Managed Sub-account Asset Details(For Investor Master Account)
[**GetManagedSubaccountDepositAddressV1**](SubAccountAPI.md#GetManagedSubaccountDepositAddressV1) | **Get** /sapi/v1/managed-subaccount/deposit/address | Get Managed Sub-account Deposit Address (For Investor Master Account)(USER_DATA)
[**GetManagedSubaccountFetchFutureAssetV1**](SubAccountAPI.md#GetManagedSubaccountFetchFutureAssetV1) | **Get** /sapi/v1/managed-subaccount/fetch-future-asset | Query Managed Sub-account Futures Asset Details(For Investor Master Account)(USER_DATA)
[**GetManagedSubaccountInfoV1**](SubAccountAPI.md#GetManagedSubaccountInfoV1) | **Get** /sapi/v1/managed-subaccount/info | Query Managed Sub-account List(For Investor)(USER_DATA)
[**GetManagedSubaccountMarginAssetV1**](SubAccountAPI.md#GetManagedSubaccountMarginAssetV1) | **Get** /sapi/v1/managed-subaccount/marginAsset | Query Managed Sub-account Margin Asset Details(For Investor Master Account)(USER_DATA)
[**GetManagedSubaccountQueryTransLogForInvestorV1**](SubAccountAPI.md#GetManagedSubaccountQueryTransLogForInvestorV1) | **Get** /sapi/v1/managed-subaccount/queryTransLogForInvestor | Query Managed Sub Account Transfer Log(For Investor Master Account)(USER_DATA)
[**GetManagedSubaccountQueryTransLogForTradeParentV1**](SubAccountAPI.md#GetManagedSubaccountQueryTransLogForTradeParentV1) | **Get** /sapi/v1/managed-subaccount/queryTransLogForTradeParent | Query Managed Sub Account Transfer Log(For Trading Team Master Account)(USER_DATA)
[**GetManagedSubaccountQueryTransLogV1**](SubAccountAPI.md#GetManagedSubaccountQueryTransLogV1) | **Get** /sapi/v1/managed-subaccount/query-trans-log | Query Managed Sub Account Transfer Log (For Trading Team Sub Account)(USER_DATA)
[**GetSubAccountAssetsV3**](SubAccountAPI.md#GetSubAccountAssetsV3) | **Get** /sapi/v3/sub-account/assets | Query Sub-account Assets(For Master Account)
[**GetSubAccountAssetsV4**](SubAccountAPI.md#GetSubAccountAssetsV4) | **Get** /sapi/v4/sub-account/assets | Query Sub-account Assets (For Master Account)(USER_DATA)
[**GetSubAccountFuturesAccountSummaryV1**](SubAccountAPI.md#GetSubAccountFuturesAccountSummaryV1) | **Get** /sapi/v1/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account(For Master Account)
[**GetSubAccountFuturesAccountSummaryV2**](SubAccountAPI.md#GetSubAccountFuturesAccountSummaryV2) | **Get** /sapi/v2/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account V2(For Master Account)
[**GetSubAccountFuturesAccountV1**](SubAccountAPI.md#GetSubAccountFuturesAccountV1) | **Get** /sapi/v1/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account(For Master Account)
[**GetSubAccountFuturesAccountV2**](SubAccountAPI.md#GetSubAccountFuturesAccountV2) | **Get** /sapi/v2/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account V2(For Master Account)
[**GetSubAccountFuturesInternalTransferV1**](SubAccountAPI.md#GetSubAccountFuturesInternalTransferV1) | **Get** /sapi/v1/sub-account/futures/internalTransfer | Query Sub-account Futures Asset Transfer History(For Master Account)
[**GetSubAccountFuturesMovePositionV1**](SubAccountAPI.md#GetSubAccountFuturesMovePositionV1) | **Get** /sapi/v1/sub-account/futures/move-position | Get Move Position History for Sub-account (For Master Account)
[**GetSubAccountFuturesPositionRiskV1**](SubAccountAPI.md#GetSubAccountFuturesPositionRiskV1) | **Get** /sapi/v1/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account(For Master Account)
[**GetSubAccountFuturesPositionRiskV2**](SubAccountAPI.md#GetSubAccountFuturesPositionRiskV2) | **Get** /sapi/v2/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account V2(For Master Account)
[**GetSubAccountListV1**](SubAccountAPI.md#GetSubAccountListV1) | **Get** /sapi/v1/sub-account/list | Query Sub-account List(For Master Account)
[**GetSubAccountMarginAccountSummaryV1**](SubAccountAPI.md#GetSubAccountMarginAccountSummaryV1) | **Get** /sapi/v1/sub-account/margin/accountSummary | Get Summary of Sub-account&#39;s Margin Account(For Master Account)
[**GetSubAccountMarginAccountV1**](SubAccountAPI.md#GetSubAccountMarginAccountV1) | **Get** /sapi/v1/sub-account/margin/account | Get Detail on Sub-account&#39;s Margin Account(For Master Account)
[**GetSubAccountSpotSummaryV1**](SubAccountAPI.md#GetSubAccountSpotSummaryV1) | **Get** /sapi/v1/sub-account/spotSummary | Query Sub-account Spot Assets Summary(For Master Account)
[**GetSubAccountStatusV1**](SubAccountAPI.md#GetSubAccountStatusV1) | **Get** /sapi/v1/sub-account/status | Get Sub-account&#39;s Status on Margin Or Futures(For Master Account)
[**GetSubAccountSubAccountApiIpRestrictionV1**](SubAccountAPI.md#GetSubAccountSubAccountApiIpRestrictionV1) | **Get** /sapi/v1/sub-account/subAccountApi/ipRestriction | Get IP Restriction for a Sub-account API Key(For Master Account)
[**GetSubAccountSubTransferHistoryV1**](SubAccountAPI.md#GetSubAccountSubTransferHistoryV1) | **Get** /sapi/v1/sub-account/sub/transfer/history | Query Sub-account Spot Asset Transfer History(For Master Account)
[**GetSubAccountTransactionStatisticsV1**](SubAccountAPI.md#GetSubAccountTransactionStatisticsV1) | **Get** /sapi/v1/sub-account/transaction-statistics | Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)
[**GetSubAccountTransferSubUserHistoryV1**](SubAccountAPI.md#GetSubAccountTransferSubUserHistoryV1) | **Get** /sapi/v1/sub-account/transfer/subUserHistory | Sub-account Transfer History(For Sub-account)
[**GetSubAccountUniversalTransferV1**](SubAccountAPI.md#GetSubAccountUniversalTransferV1) | **Get** /sapi/v1/sub-account/universalTransfer | Query Universal Transfer History(For Master Account)



## CreateManagedSubaccountDepositV1

> CreateManagedSubaccountDepositV1Resp CreateManagedSubaccountDepositV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()

Deposit Assets Into The Managed Sub-account(For Investor Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toEmail := "toEmail_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateManagedSubaccountDepositV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateManagedSubaccountDepositV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManagedSubaccountDepositV1`: CreateManagedSubaccountDepositV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateManagedSubaccountDepositV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedSubaccountDepositV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateManagedSubaccountDepositV1Resp**](CreateManagedSubaccountDepositV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManagedSubaccountWithdrawV1

> CreateManagedSubaccountWithdrawV1Resp CreateManagedSubaccountWithdrawV1(ctx).Amount(amount).Asset(asset).FromEmail(fromEmail).Timestamp(timestamp).RecvWindow(recvWindow).TransferDate(transferDate).Execute()

Withdrawl Assets From The Managed Sub-account(For Investor Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	fromEmail := "fromEmail_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	transferDate := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateManagedSubaccountWithdrawV1(context.Background()).Amount(amount).Asset(asset).FromEmail(fromEmail).Timestamp(timestamp).RecvWindow(recvWindow).TransferDate(transferDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateManagedSubaccountWithdrawV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManagedSubaccountWithdrawV1`: CreateManagedSubaccountWithdrawV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateManagedSubaccountWithdrawV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManagedSubaccountWithdrawV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **fromEmail** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **transferDate** | **int64** |  | 

### Return type

[**CreateManagedSubaccountWithdrawV1Resp**](CreateManagedSubaccountWithdrawV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountBlvtEnableV1

> CreateSubAccountBlvtEnableV1Resp CreateSubAccountBlvtEnableV1(ctx).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Leverage Token for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	enableBlvt := true // bool | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountBlvtEnableV1(context.Background()).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountBlvtEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountBlvtEnableV1`: CreateSubAccountBlvtEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountBlvtEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountBlvtEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **enableBlvt** | **bool** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountBlvtEnableV1Resp**](CreateSubAccountBlvtEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountEoptionsEnableV1

> CreateSubAccountEoptionsEnableV1Resp CreateSubAccountEoptionsEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Options for Sub-account(For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountEoptionsEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountEoptionsEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountEoptionsEnableV1`: CreateSubAccountEoptionsEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountEoptionsEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountEoptionsEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountEoptionsEnableV1Resp**](CreateSubAccountEoptionsEnableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountFuturesEnableV1

> CreateSubAccountFuturesEnableV1Resp CreateSubAccountFuturesEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Futures for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountFuturesEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountFuturesEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountFuturesEnableV1`: CreateSubAccountFuturesEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountFuturesEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountFuturesEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountFuturesEnableV1Resp**](CreateSubAccountFuturesEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountFuturesInternalTransferV1

> CreateSubAccountFuturesInternalTransferV1Resp CreateSubAccountFuturesInternalTransferV1(ctx).Amount(amount).Asset(asset).FromEmail(fromEmail).FuturesType(futuresType).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()

Sub-account Futures Asset Transfer(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountFuturesInternalTransferV1(context.Background()).Amount(amount).Asset(asset).FromEmail(fromEmail).FuturesType(futuresType).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountFuturesInternalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountFuturesInternalTransferV1`: CreateSubAccountFuturesInternalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountFuturesInternalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountFuturesInternalTransferV1Request struct via the builder pattern


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

[**CreateSubAccountFuturesInternalTransferV1Resp**](CreateSubAccountFuturesInternalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountFuturesMovePositionV1

> CreateSubAccountFuturesMovePositionV1Resp CreateSubAccountFuturesMovePositionV1(ctx).FromUserEmail(fromUserEmail).OrderArgs(orderArgs).ProductType(productType).Timestamp(timestamp).ToUserEmail(toUserEmail).RecvWindow(recvWindow).Execute()

Move Position for Sub-account (For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountFuturesMovePositionV1(context.Background()).FromUserEmail(fromUserEmail).OrderArgs(orderArgs).ProductType(productType).Timestamp(timestamp).ToUserEmail(toUserEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountFuturesMovePositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountFuturesMovePositionV1`: CreateSubAccountFuturesMovePositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountFuturesMovePositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountFuturesMovePositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromUserEmail** | **string** |  | [default to &quot;&quot;]
 **orderArgs** | **[]map[string]interface{}** |  | 
 **productType** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toUserEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountFuturesMovePositionV1Resp**](CreateSubAccountFuturesMovePositionV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountFuturesTransferV1

> CreateSubAccountFuturesTransferV1Resp CreateSubAccountFuturesTransferV1(ctx).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Futures Transfer for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountFuturesTransferV1(context.Background()).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountFuturesTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountFuturesTransferV1`: CreateSubAccountFuturesTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountFuturesTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountFuturesTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountFuturesTransferV1Resp**](CreateSubAccountFuturesTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountMarginEnableV1

> CreateSubAccountMarginEnableV1Resp CreateSubAccountMarginEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Margin for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountMarginEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountMarginEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountMarginEnableV1`: CreateSubAccountMarginEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountMarginEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountMarginEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountMarginEnableV1Resp**](CreateSubAccountMarginEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountMarginTransferV1

> CreateSubAccountMarginTransferV1Resp CreateSubAccountMarginTransferV1(ctx).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Margin Transfer for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountMarginTransferV1(context.Background()).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountMarginTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountMarginTransferV1`: CreateSubAccountMarginTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountMarginTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountMarginTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **int32** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountMarginTransferV1Resp**](CreateSubAccountMarginTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountSubAccountApiIpRestrictionV2

> CreateSubAccountSubAccountApiIpRestrictionV2Resp CreateSubAccountSubAccountApiIpRestrictionV2(ctx).Email(email).Status(status).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()

Add IP Restriction for Sub-Account API key(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	status := "status_example" // string |  (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ipAddress := "ipAddress_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountSubAccountApiIpRestrictionV2(context.Background()).Email(email).Status(status).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountSubAccountApiIpRestrictionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountSubAccountApiIpRestrictionV2`: CreateSubAccountSubAccountApiIpRestrictionV2Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountSubAccountApiIpRestrictionV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountSubAccountApiIpRestrictionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **status** | **string** |  | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **ipAddress** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountSubAccountApiIpRestrictionV2Resp**](CreateSubAccountSubAccountApiIpRestrictionV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountTransferSubToMasterV1

> CreateSubAccountTransferSubToMasterV1Resp CreateSubAccountTransferSubToMasterV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Transfer to Master(For Sub-account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountTransferSubToMasterV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountTransferSubToMasterV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountTransferSubToMasterV1`: CreateSubAccountTransferSubToMasterV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountTransferSubToMasterV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountTransferSubToMasterV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountTransferSubToMasterV1Resp**](CreateSubAccountTransferSubToMasterV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountTransferSubToSubV1

> CreateSubAccountTransferSubToSubV1Resp CreateSubAccountTransferSubToSubV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()

Transfer to Sub-account of Same Master(For Sub-account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	toEmail := "toEmail_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountTransferSubToSubV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountTransferSubToSubV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountTransferSubToSubV1`: CreateSubAccountTransferSubToSubV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountTransferSubToSubV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountTransferSubToSubV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountTransferSubToSubV1Resp**](CreateSubAccountTransferSubToSubV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountUniversalTransferV1

> CreateSubAccountUniversalTransferV1Resp CreateSubAccountUniversalTransferV1(ctx).Amount(amount).Asset(asset).FromAccountType(fromAccountType).Timestamp(timestamp).ToAccountType(toAccountType).ClientTranId(clientTranId).FromEmail(fromEmail).RecvWindow(recvWindow).Symbol(symbol).ToEmail(toEmail).Execute()

Universal Transfer(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountUniversalTransferV1(context.Background()).Amount(amount).Asset(asset).FromAccountType(fromAccountType).Timestamp(timestamp).ToAccountType(toAccountType).ClientTranId(clientTranId).FromEmail(fromEmail).RecvWindow(recvWindow).Symbol(symbol).ToEmail(toEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountUniversalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountUniversalTransferV1`: CreateSubAccountUniversalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountUniversalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountUniversalTransferV1Request struct via the builder pattern


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

[**CreateSubAccountUniversalTransferV1Resp**](CreateSubAccountUniversalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubAccountVirtualSubAccountV1

> CreateSubAccountVirtualSubAccountV1Resp CreateSubAccountVirtualSubAccountV1(ctx).SubAccountString(subAccountString).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Create a Virtual Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	subAccountString := "subAccountString_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateSubAccountVirtualSubAccountV1(context.Background()).SubAccountString(subAccountString).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateSubAccountVirtualSubAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubAccountVirtualSubAccountV1`: CreateSubAccountVirtualSubAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateSubAccountVirtualSubAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubAccountVirtualSubAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAccountString** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSubAccountVirtualSubAccountV1Resp**](CreateSubAccountVirtualSubAccountV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubAccountSubAccountApiIpRestrictionIpListV1

> DeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp DeleteSubAccountSubAccountApiIpRestrictionIpListV1(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()

Delete IP List For a Sub-account API Key(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key#email-address\">Sub-account email</a> (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ipAddress := "ipAddress_example" // string | Can be added in batches, separated by commas (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.DeleteSubAccountSubAccountApiIpRestrictionIpListV1(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.DeleteSubAccountSubAccountApiIpRestrictionIpListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSubAccountSubAccountApiIpRestrictionIpListV1`: DeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.DeleteSubAccountSubAccountApiIpRestrictionIpListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **ipAddress** | **string** | Can be added in batches, separated by commas | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**DeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp**](DeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalDepositSubAddressV1

> GetCapitalDepositSubAddressV1Resp GetCapitalDepositSubAddressV1(ctx).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Get Sub-account Deposit Address(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.GetCapitalDepositSubAddressV1(context.Background()).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetCapitalDepositSubAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalDepositSubAddressV1`: GetCapitalDepositSubAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetCapitalDepositSubAddressV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalDepositSubAddressV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetCapitalDepositSubAddressV1Resp**](GetCapitalDepositSubAddressV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalDepositSubHisrecV1

> []GetCapitalDepositSubHisrecV1RespItem GetCapitalDepositSubHisrecV1(ctx).Email(email).Timestamp(timestamp).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).TxId(txId).Execute()

Get Sub-account Deposit History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.GetCapitalDepositSubHisrecV1(context.Background()).Email(email).Timestamp(timestamp).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).TxId(txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetCapitalDepositSubHisrecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalDepositSubHisrecV1`: []GetCapitalDepositSubHisrecV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetCapitalDepositSubHisrecV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalDepositSubHisrecV1Request struct via the builder pattern


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

[**[]GetCapitalDepositSubHisrecV1RespItem**](GetCapitalDepositSubHisrecV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountAccountSnapshotV1

> GetManagedSubaccountAccountSnapshotV1Resp GetManagedSubaccountAccountSnapshotV1(ctx).Email(email).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Managed Sub-account Snapshot(For Investor Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	type_ := "type__example" // string | &#34;SPOT&#34;, &#34;MARGIN&#34;（cross）, &#34;FUTURES&#34;（UM） (default to "")
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | min 7, max 30, default 7 (optional) (default to 7)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountAccountSnapshotV1(context.Background()).Email(email).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountAccountSnapshotV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountAccountSnapshotV1`: GetManagedSubaccountAccountSnapshotV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountAccountSnapshotV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountAccountSnapshotV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **type_** | **string** | &amp;#34;SPOT&amp;#34;, &amp;#34;MARGIN&amp;#34;（cross）, &amp;#34;FUTURES&amp;#34;（UM） | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | min 7, max 30, default 7 | [default to 7]
 **recvWindow** | **int64** |  | 

### Return type

[**GetManagedSubaccountAccountSnapshotV1Resp**](GetManagedSubaccountAccountSnapshotV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountAssetV1

> []GetManagedSubaccountAssetV1RespItem GetManagedSubaccountAssetV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Managed Sub-account Asset Details(For Investor Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountAssetV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountAssetV1`: []GetManagedSubaccountAssetV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetManagedSubaccountAssetV1RespItem**](GetManagedSubaccountAssetV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountDepositAddressV1

> GetManagedSubaccountDepositAddressV1Resp GetManagedSubaccountDepositAddressV1(ctx).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Get Managed Sub-account Deposit Address (For Investor Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Sub user email (default to "")
	coin := "coin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	network := "network_example" // string | networks can be found in `GET /sapi/v1/capital/deposit/address` (optional) (default to "")
	amount := "amount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountDepositAddressV1(context.Background()).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountDepositAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountDepositAddressV1`: GetManagedSubaccountDepositAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountDepositAddressV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountDepositAddressV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub user email | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** | networks can be found in &#x60;GET /sapi/v1/capital/deposit/address&#x60; | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetManagedSubaccountDepositAddressV1Resp**](GetManagedSubaccountDepositAddressV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountFetchFutureAssetV1

> GetManagedSubaccountFetchFutureAssetV1Resp GetManagedSubaccountFetchFutureAssetV1(ctx).Email(email).AccountType(accountType).Execute()

Query Managed Sub-account Futures Asset Details(For Investor Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Managed Sub Account Email (default to "")
	accountType := "accountType_example" // string | No input or input &#34;USDT_FUTURE&#34; to get UM Futures account details. Input &#34;COIN_FUTURE&#34; to get CM Futures account details. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountFetchFutureAssetV1(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountFetchFutureAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountFetchFutureAssetV1`: GetManagedSubaccountFetchFutureAssetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountFetchFutureAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountFetchFutureAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed Sub Account Email | [default to &quot;&quot;]
 **accountType** | **string** | No input or input &amp;#34;USDT_FUTURE&amp;#34; to get UM Futures account details. Input &amp;#34;COIN_FUTURE&amp;#34; to get CM Futures account details. | [default to &quot;&quot;]

### Return type

[**GetManagedSubaccountFetchFutureAssetV1Resp**](GetManagedSubaccountFetchFutureAssetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountInfoV1

> GetManagedSubaccountInfoV1Resp GetManagedSubaccountInfoV1(ctx).Timestamp(timestamp).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Managed Sub-account List(For Investor)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	email := "email_example" // string | Managed sub-account email (optional) (default to "")
	page := int32(56) // int32 | Default value: 1 (optional)
	limit := int32(56) // int32 | Default value: 20, Max value: 20 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountInfoV1(context.Background()).Timestamp(timestamp).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountInfoV1`: GetManagedSubaccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | Managed sub-account email | [default to &quot;&quot;]
 **page** | **int32** | Default value: 1 | 
 **limit** | **int32** | Default value: 20, Max value: 20 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetManagedSubaccountInfoV1Resp**](GetManagedSubaccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountMarginAssetV1

> GetManagedSubaccountMarginAssetV1Resp GetManagedSubaccountMarginAssetV1(ctx).Email(email).AccountType(accountType).Execute()

Query Managed Sub-account Margin Asset Details(For Investor Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Managed Sub Account Email (default to "")
	accountType := "accountType_example" // string | No input or input &#34;MARGIN&#34; to get Cross Margin account details. Input &#34;ISOLATED_MARGIN&#34; to get Isolated Margin account details. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountMarginAssetV1(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountMarginAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountMarginAssetV1`: GetManagedSubaccountMarginAssetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountMarginAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountMarginAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed Sub Account Email | [default to &quot;&quot;]
 **accountType** | **string** | No input or input &amp;#34;MARGIN&amp;#34; to get Cross Margin account details. Input &amp;#34;ISOLATED_MARGIN&amp;#34; to get Isolated Margin account details. | [default to &quot;&quot;]

### Return type

[**GetManagedSubaccountMarginAssetV1Resp**](GetManagedSubaccountMarginAssetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountQueryTransLogForInvestorV1

> GetManagedSubaccountQueryTransLogForInvestorV1Resp GetManagedSubaccountQueryTransLogForInvestorV1(ctx).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()

Query Managed Sub Account Transfer Log(For Investor Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Managed Sub Account Email (default to "")
	startTime := int64(789) // int64 | Start Time
	endTime := int64(789) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int32(56) // int32 | Page
	limit := int32(56) // int32 | Limit (Max: 500)
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional) (default to "")
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountQueryTransLogForInvestorV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountQueryTransLogForInvestorV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountQueryTransLogForInvestorV1`: GetManagedSubaccountQueryTransLogForInvestorV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountQueryTransLogForInvestorV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountQueryTransLogForInvestorV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed Sub Account Email | [default to &quot;&quot;]
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int32** | Page | 
 **limit** | **int32** | Limit (Max: 500) | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | [default to &quot;&quot;]
 **transferFunctionAccountType** | **string** | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) | [default to &quot;&quot;]

### Return type

[**GetManagedSubaccountQueryTransLogForInvestorV1Resp**](GetManagedSubaccountQueryTransLogForInvestorV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountQueryTransLogForTradeParentV1

> GetManagedSubaccountQueryTransLogForTradeParentV1Resp GetManagedSubaccountQueryTransLogForTradeParentV1(ctx).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()

Query Managed Sub Account Transfer Log(For Trading Team Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Managed Sub Account Email (default to "")
	startTime := int64(789) // int64 | Start Time
	endTime := int64(789) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int32(56) // int32 | Page
	limit := int32(56) // int32 | Limit (Max: 500)
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional) (default to "")
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountQueryTransLogForTradeParentV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountQueryTransLogForTradeParentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountQueryTransLogForTradeParentV1`: GetManagedSubaccountQueryTransLogForTradeParentV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountQueryTransLogForTradeParentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountQueryTransLogForTradeParentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed Sub Account Email | [default to &quot;&quot;]
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int32** | Page | 
 **limit** | **int32** | Limit (Max: 500) | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | [default to &quot;&quot;]
 **transferFunctionAccountType** | **string** | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) | [default to &quot;&quot;]

### Return type

[**GetManagedSubaccountQueryTransLogForTradeParentV1Resp**](GetManagedSubaccountQueryTransLogForTradeParentV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSubaccountQueryTransLogV1

> GetManagedSubaccountQueryTransLogV1Resp GetManagedSubaccountQueryTransLogV1(ctx).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Timestamp(timestamp).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).RecvWindow(recvWindow).Execute()

Query Managed Sub Account Transfer Log (For Trading Team Sub Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	startTime := int64(789) // int64 | Start Time
	endTime := int64(789) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int32(56) // int32 | Page
	limit := int32(56) // int32 | Limit (Max: 500)
	timestamp := int64(789) // int64 | 
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional) (default to "")
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetManagedSubaccountQueryTransLogV1(context.Background()).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Timestamp(timestamp).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetManagedSubaccountQueryTransLogV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSubaccountQueryTransLogV1`: GetManagedSubaccountQueryTransLogV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetManagedSubaccountQueryTransLogV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSubaccountQueryTransLogV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Start Time | 
 **endTime** | **int64** | End Time (The start time and end time interval cannot exceed half a year) | 
 **page** | **int32** | Page | 
 **limit** | **int32** | Limit (Max: 500) | 
 **timestamp** | **int64** |  | 
 **transfers** | **string** | Transfer Direction (FROM/TO) | [default to &quot;&quot;]
 **transferFunctionAccountType** | **string** | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetManagedSubaccountQueryTransLogV1Resp**](GetManagedSubaccountQueryTransLogV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountAssetsV3

> GetSubAccountAssetsV3Resp GetSubAccountAssetsV3(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Assets(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Sub account email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountAssetsV3(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountAssetsV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountAssetsV3`: GetSubAccountAssetsV3Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountAssetsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountAssetsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountAssetsV3Resp**](GetSubAccountAssetsV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountAssetsV4

> GetSubAccountAssetsV4Resp GetSubAccountAssetsV4(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Assets (For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Sub Account Email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountAssetsV4(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountAssetsV4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountAssetsV4`: GetSubAccountAssetsV4Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountAssetsV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountAssetsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub Account Email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountAssetsV4Resp**](GetSubAccountAssetsV4Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesAccountSummaryV1

> GetSubAccountFuturesAccountSummaryV1Resp GetSubAccountFuturesAccountSummaryV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Futures Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesAccountSummaryV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesAccountSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesAccountSummaryV1`: GetSubAccountFuturesAccountSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesAccountSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesAccountSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountFuturesAccountSummaryV1Resp**](GetSubAccountFuturesAccountSummaryV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesAccountSummaryV2

> GetSubAccountFuturesAccountSummaryV2Resp GetSubAccountFuturesAccountSummaryV2(ctx).FuturesType(futuresType).Timestamp(timestamp).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Futures Account V2(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	page := int32(56) // int32 | default:1 (optional)
	limit := int32(56) // int32 | default:10, max:20 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesAccountSummaryV2(context.Background()).FuturesType(futuresType).Timestamp(timestamp).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesAccountSummaryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesAccountSummaryV2`: GetSubAccountFuturesAccountSummaryV2Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesAccountSummaryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesAccountSummaryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **page** | **int32** | default:1 | 
 **limit** | **int32** | default:10, max:20 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountFuturesAccountSummaryV2Resp**](GetSubAccountFuturesAccountSummaryV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesAccountV1

> GetSubAccountFuturesAccountV1Resp GetSubAccountFuturesAccountV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Futures Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account#email-address\">Sub-account email</a> (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesAccountV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesAccountV1`: GetSubAccountFuturesAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountFuturesAccountV1Resp**](GetSubAccountFuturesAccountV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesAccountV2

> GetSubAccountFuturesAccountV2Resp GetSubAccountFuturesAccountV2(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Futures Account V2(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2#email-address\">Sub-account email</a> (default to "")
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesAccountV2(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesAccountV2`: GetSubAccountFuturesAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountFuturesAccountV2Resp**](GetSubAccountFuturesAccountV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesInternalTransferV1

> GetSubAccountFuturesInternalTransferV1Resp GetSubAccountFuturesInternalTransferV1(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account Futures Asset Transfer History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesInternalTransferV1(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesInternalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesInternalTransferV1`: GetSubAccountFuturesInternalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesInternalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesInternalTransferV1Request struct via the builder pattern


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

[**GetSubAccountFuturesInternalTransferV1Resp**](GetSubAccountFuturesInternalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesMovePositionV1

> GetSubAccountFuturesMovePositionV1Resp GetSubAccountFuturesMovePositionV1(ctx).Symbol(symbol).Page(page).Row(row).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Move Position History for Sub-account (For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesMovePositionV1(context.Background()).Symbol(symbol).Page(page).Row(row).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesMovePositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesMovePositionV1`: GetSubAccountFuturesMovePositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesMovePositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesMovePositionV1Request struct via the builder pattern


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

[**GetSubAccountFuturesMovePositionV1Resp**](GetSubAccountFuturesMovePositionV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesPositionRiskV1

> []GetSubAccountFuturesPositionRiskV1RespItem GetSubAccountFuturesPositionRiskV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Position-Risk of Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account#email-address\">Sub-account email</a> (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesPositionRiskV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesPositionRiskV1`: []GetSubAccountFuturesPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesPositionRiskV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesPositionRiskV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetSubAccountFuturesPositionRiskV1RespItem**](GetSubAccountFuturesPositionRiskV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountFuturesPositionRiskV2

> GetSubAccountFuturesPositionRiskV2Resp GetSubAccountFuturesPositionRiskV2(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Position-Risk of Sub-account V2(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account-V2#email-address\">Sub-account email</a> (default to "")
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountFuturesPositionRiskV2(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountFuturesPositionRiskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountFuturesPositionRiskV2`: GetSubAccountFuturesPositionRiskV2Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountFuturesPositionRiskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountFuturesPositionRiskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account-V2#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountFuturesPositionRiskV2Resp**](GetSubAccountFuturesPositionRiskV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountListV1

> GetSubAccountListV1Resp GetSubAccountListV1(ctx).Timestamp(timestamp).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account List(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Query-Sub-account-List#email-address\">Sub-account email</a> (optional) (default to "")
	isFreeze := "isFreeze_example" // string | true or false (optional) (default to "")
	page := int32(56) // int32 | Default value: 1 (optional)
	limit := int32(56) // int32 | Default value: 1, Max value: 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountListV1(context.Background()).Timestamp(timestamp).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountListV1`: GetSubAccountListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Query-Sub-account-List#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **isFreeze** | **string** | true or false | [default to &quot;&quot;]
 **page** | **int32** | Default value: 1 | 
 **limit** | **int32** | Default value: 1, Max value: 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountListV1Resp**](GetSubAccountListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountMarginAccountSummaryV1

> GetSubAccountMarginAccountSummaryV1Resp GetSubAccountMarginAccountSummaryV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Margin Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountMarginAccountSummaryV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountMarginAccountSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountMarginAccountSummaryV1`: GetSubAccountMarginAccountSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountMarginAccountSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountMarginAccountSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountMarginAccountSummaryV1Resp**](GetSubAccountMarginAccountSummaryV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountMarginAccountV1

> GetSubAccountMarginAccountV1Resp GetSubAccountMarginAccountV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Margin Account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Margin-Account#email-address\">Sub-account email</a> (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountMarginAccountV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountMarginAccountV1`: GetSubAccountMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Margin-Account#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountMarginAccountV1Resp**](GetSubAccountMarginAccountV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountSpotSummaryV1

> GetSubAccountSpotSummaryV1Resp GetSubAccountSpotSummaryV1(ctx).Timestamp(timestamp).Email(email).Page(page).Size(size).RecvWindow(recvWindow).Execute()

Query Sub-account Spot Assets Summary(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	email := "email_example" // string | Sub account email (optional) (default to "")
	page := int64(789) // int64 | default 1 (optional) (default to 1)
	size := int64(789) // int64 | default 10, max 20 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountSpotSummaryV1(context.Background()).Timestamp(timestamp).Email(email).Page(page).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountSpotSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountSpotSummaryV1`: GetSubAccountSpotSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountSpotSummaryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountSpotSummaryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **page** | **int64** | default 1 | [default to 1]
 **size** | **int64** | default 10, max 20 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountSpotSummaryV1Resp**](GetSubAccountSpotSummaryV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountStatusV1

> []GetSubAccountStatusV1RespItem GetSubAccountStatusV1(ctx).Timestamp(timestamp).Email(email).RecvWindow(recvWindow).Execute()

Get Sub-account's Status on Margin Or Futures(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures#email-address\">Sub-account email</a> (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountStatusV1(context.Background()).Timestamp(timestamp).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountStatusV1`: []GetSubAccountStatusV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetSubAccountStatusV1RespItem**](GetSubAccountStatusV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountSubAccountApiIpRestrictionV1

> GetSubAccountSubAccountApiIpRestrictionV1Resp GetSubAccountSubAccountApiIpRestrictionV1(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get IP Restriction for a Sub-account API Key(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/api-management#email-address\">Sub-account email</a> (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountSubAccountApiIpRestrictionV1(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountSubAccountApiIpRestrictionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountSubAccountApiIpRestrictionV1`: GetSubAccountSubAccountApiIpRestrictionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountSubAccountApiIpRestrictionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountSubAccountApiIpRestrictionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/api-management#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSubAccountSubAccountApiIpRestrictionV1Resp**](GetSubAccountSubAccountApiIpRestrictionV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountSubTransferHistoryV1

> []GetSubAccountSubTransferHistoryV1RespItem GetSubAccountSubTransferHistoryV1(ctx).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account Spot Asset Transfer History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountSubTransferHistoryV1(context.Background()).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountSubTransferHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountSubTransferHistoryV1`: []GetSubAccountSubTransferHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountSubTransferHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountSubTransferHistoryV1Request struct via the builder pattern


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

[**[]GetSubAccountSubTransferHistoryV1RespItem**](GetSubAccountSubTransferHistoryV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountTransactionStatisticsV1

> SubaccountGetSubAccountTransactionStatisticsV1Resp GetSubAccountTransactionStatisticsV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	email := "email_example" // string | Sub user email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountTransactionStatisticsV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountTransactionStatisticsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountTransactionStatisticsV1`: SubaccountGetSubAccountTransactionStatisticsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountTransactionStatisticsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountTransactionStatisticsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub user email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountTransactionStatisticsV1Resp**](SubaccountGetSubAccountTransactionStatisticsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountTransferSubUserHistoryV1

> []GetSubAccountTransferSubUserHistoryV1RespItem GetSubAccountTransferSubUserHistoryV1(ctx).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).ReturnFailHistory(returnFailHistory).RecvWindow(recvWindow).Execute()

Sub-account Transfer History(For Sub-account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountTransferSubUserHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).ReturnFailHistory(returnFailHistory).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountTransferSubUserHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountTransferSubUserHistoryV1`: []GetSubAccountTransferSubUserHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountTransferSubUserHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountTransferSubUserHistoryV1Request struct via the builder pattern


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

[**[]GetSubAccountTransferSubUserHistoryV1RespItem**](GetSubAccountTransferSubUserHistoryV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubAccountUniversalTransferV1

> GetSubAccountUniversalTransferV1Resp GetSubAccountUniversalTransferV1(ctx).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Universal Transfer History(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
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
	resp, r, err := apiClient.SubAccountAPI.GetSubAccountUniversalTransferV1(context.Background()).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetSubAccountUniversalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubAccountUniversalTransferV1`: GetSubAccountUniversalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetSubAccountUniversalTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSubAccountUniversalTransferV1Request struct via the builder pattern


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

[**GetSubAccountUniversalTransferV1Resp**](GetSubAccountUniversalTransferV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

