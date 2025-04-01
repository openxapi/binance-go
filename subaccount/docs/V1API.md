# \V1API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreateManagedSubaccountDepositV1**](V1API.md#SubaccountCreateManagedSubaccountDepositV1) | **Post** /sapi/v1/managed-subaccount/deposit | Deposit Assets Into The Managed Sub-account(For Investor Master Account)
[**SubaccountCreateManagedSubaccountWithdrawV1**](V1API.md#SubaccountCreateManagedSubaccountWithdrawV1) | **Post** /sapi/v1/managed-subaccount/withdraw | Withdrawl Assets From The Managed Sub-account(For Investor Master Account)
[**SubaccountCreateSubAccountBlvtEnableV1**](V1API.md#SubaccountCreateSubAccountBlvtEnableV1) | **Post** /sapi/v1/sub-account/blvt/enable | Enable Leverage Token for Sub-account(For Master Account)
[**SubaccountCreateSubAccountEoptionsEnableV1**](V1API.md#SubaccountCreateSubAccountEoptionsEnableV1) | **Post** /sapi/v1/sub-account/eoptions/enable | Enable Options for Sub-account(For Master Account)(USER_DATA)
[**SubaccountCreateSubAccountFuturesEnableV1**](V1API.md#SubaccountCreateSubAccountFuturesEnableV1) | **Post** /sapi/v1/sub-account/futures/enable | Enable Futures for Sub-account(For Master Account)
[**SubaccountCreateSubAccountFuturesInternalTransferV1**](V1API.md#SubaccountCreateSubAccountFuturesInternalTransferV1) | **Post** /sapi/v1/sub-account/futures/internalTransfer | Sub-account Futures Asset Transfer(For Master Account)
[**SubaccountCreateSubAccountFuturesMovePositionV1**](V1API.md#SubaccountCreateSubAccountFuturesMovePositionV1) | **Post** /sapi/v1/sub-account/futures/move-position | Move Position for Sub-account (For Master Account)
[**SubaccountCreateSubAccountFuturesTransferV1**](V1API.md#SubaccountCreateSubAccountFuturesTransferV1) | **Post** /sapi/v1/sub-account/futures/transfer | Futures Transfer for Sub-account(For Master Account)
[**SubaccountCreateSubAccountMarginEnableV1**](V1API.md#SubaccountCreateSubAccountMarginEnableV1) | **Post** /sapi/v1/sub-account/margin/enable | Enable Margin for Sub-account(For Master Account)
[**SubaccountCreateSubAccountMarginTransferV1**](V1API.md#SubaccountCreateSubAccountMarginTransferV1) | **Post** /sapi/v1/sub-account/margin/transfer | Margin Transfer for Sub-account(For Master Account)
[**SubaccountCreateSubAccountTransferSubToMasterV1**](V1API.md#SubaccountCreateSubAccountTransferSubToMasterV1) | **Post** /sapi/v1/sub-account/transfer/subToMaster | Transfer to Master(For Sub-account)
[**SubaccountCreateSubAccountTransferSubToSubV1**](V1API.md#SubaccountCreateSubAccountTransferSubToSubV1) | **Post** /sapi/v1/sub-account/transfer/subToSub | Transfer to Sub-account of Same Master(For Sub-account)
[**SubaccountCreateSubAccountUniversalTransferV1**](V1API.md#SubaccountCreateSubAccountUniversalTransferV1) | **Post** /sapi/v1/sub-account/universalTransfer | Universal Transfer(For Master Account)
[**SubaccountCreateSubAccountVirtualSubAccountV1**](V1API.md#SubaccountCreateSubAccountVirtualSubAccountV1) | **Post** /sapi/v1/sub-account/virtualSubAccount | Create a Virtual Sub-account(For Master Account)
[**SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1**](V1API.md#SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1) | **Delete** /sapi/v1/sub-account/subAccountApi/ipRestriction/ipList | Delete IP List For a Sub-account API Key(For Master Account)
[**SubaccountGetCapitalDepositSubAddressV1**](V1API.md#SubaccountGetCapitalDepositSubAddressV1) | **Get** /sapi/v1/capital/deposit/subAddress | Get Sub-account Deposit Address(For Master Account)
[**SubaccountGetCapitalDepositSubHisrecV1**](V1API.md#SubaccountGetCapitalDepositSubHisrecV1) | **Get** /sapi/v1/capital/deposit/subHisrec | Get Sub-account Deposit History(For Master Account)
[**SubaccountGetManagedSubaccountAccountSnapshotV1**](V1API.md#SubaccountGetManagedSubaccountAccountSnapshotV1) | **Get** /sapi/v1/managed-subaccount/accountSnapshot | Query Managed Sub-account Snapshot(For Investor Master Account)
[**SubaccountGetManagedSubaccountAssetV1**](V1API.md#SubaccountGetManagedSubaccountAssetV1) | **Get** /sapi/v1/managed-subaccount/asset | Query Managed Sub-account Asset Details(For Investor Master Account)
[**SubaccountGetManagedSubaccountDepositAddressV1**](V1API.md#SubaccountGetManagedSubaccountDepositAddressV1) | **Get** /sapi/v1/managed-subaccount/deposit/address | Get Managed Sub-account Deposit Address (For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountFetchFutureAssetV1**](V1API.md#SubaccountGetManagedSubaccountFetchFutureAssetV1) | **Get** /sapi/v1/managed-subaccount/fetch-future-asset | Query Managed Sub-account Futures Asset Details(For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountInfoV1**](V1API.md#SubaccountGetManagedSubaccountInfoV1) | **Get** /sapi/v1/managed-subaccount/info | Query Managed Sub-account List(For Investor)(USER_DATA)
[**SubaccountGetManagedSubaccountMarginAssetV1**](V1API.md#SubaccountGetManagedSubaccountMarginAssetV1) | **Get** /sapi/v1/managed-subaccount/marginAsset | Query Managed Sub-account Margin Asset Details(For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountQueryTransLogForInvestorV1**](V1API.md#SubaccountGetManagedSubaccountQueryTransLogForInvestorV1) | **Get** /sapi/v1/managed-subaccount/queryTransLogForInvestor | Query Managed Sub Account Transfer Log(For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1**](V1API.md#SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1) | **Get** /sapi/v1/managed-subaccount/queryTransLogForTradeParent | Query Managed Sub Account Transfer Log(For Trading Team Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountQueryTransLogV1**](V1API.md#SubaccountGetManagedSubaccountQueryTransLogV1) | **Get** /sapi/v1/managed-subaccount/query-trans-log | Query Managed Sub Account Transfer Log (For Trading Team Sub Account)(USER_DATA)
[**SubaccountGetSubAccountFuturesAccountSummaryV1**](V1API.md#SubaccountGetSubAccountFuturesAccountSummaryV1) | **Get** /sapi/v1/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account(For Master Account)
[**SubaccountGetSubAccountFuturesAccountV1**](V1API.md#SubaccountGetSubAccountFuturesAccountV1) | **Get** /sapi/v1/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account(For Master Account)
[**SubaccountGetSubAccountFuturesInternalTransferV1**](V1API.md#SubaccountGetSubAccountFuturesInternalTransferV1) | **Get** /sapi/v1/sub-account/futures/internalTransfer | Query Sub-account Futures Asset Transfer History(For Master Account)
[**SubaccountGetSubAccountFuturesMovePositionV1**](V1API.md#SubaccountGetSubAccountFuturesMovePositionV1) | **Get** /sapi/v1/sub-account/futures/move-position | Get Move Position History for Sub-account (For Master Account)
[**SubaccountGetSubAccountFuturesPositionRiskV1**](V1API.md#SubaccountGetSubAccountFuturesPositionRiskV1) | **Get** /sapi/v1/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account(For Master Account)
[**SubaccountGetSubAccountListV1**](V1API.md#SubaccountGetSubAccountListV1) | **Get** /sapi/v1/sub-account/list | Query Sub-account List(For Master Account)
[**SubaccountGetSubAccountMarginAccountSummaryV1**](V1API.md#SubaccountGetSubAccountMarginAccountSummaryV1) | **Get** /sapi/v1/sub-account/margin/accountSummary | Get Summary of Sub-account&#39;s Margin Account(For Master Account)
[**SubaccountGetSubAccountMarginAccountV1**](V1API.md#SubaccountGetSubAccountMarginAccountV1) | **Get** /sapi/v1/sub-account/margin/account | Get Detail on Sub-account&#39;s Margin Account(For Master Account)
[**SubaccountGetSubAccountSpotSummaryV1**](V1API.md#SubaccountGetSubAccountSpotSummaryV1) | **Get** /sapi/v1/sub-account/spotSummary | Query Sub-account Spot Assets Summary(For Master Account)
[**SubaccountGetSubAccountStatusV1**](V1API.md#SubaccountGetSubAccountStatusV1) | **Get** /sapi/v1/sub-account/status | Get Sub-account&#39;s Status on Margin Or Futures(For Master Account)
[**SubaccountGetSubAccountSubAccountApiIpRestrictionV1**](V1API.md#SubaccountGetSubAccountSubAccountApiIpRestrictionV1) | **Get** /sapi/v1/sub-account/subAccountApi/ipRestriction | Get IP Restriction for a Sub-account API Key(For Master Account)
[**SubaccountGetSubAccountSubTransferHistoryV1**](V1API.md#SubaccountGetSubAccountSubTransferHistoryV1) | **Get** /sapi/v1/sub-account/sub/transfer/history | Query Sub-account Spot Asset Transfer History(For Master Account)
[**SubaccountGetSubAccountTransactionStatisticsV1**](V1API.md#SubaccountGetSubAccountTransactionStatisticsV1) | **Get** /sapi/v1/sub-account/transaction-statistics | Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)
[**SubaccountGetSubAccountTransferSubUserHistoryV1**](V1API.md#SubaccountGetSubAccountTransferSubUserHistoryV1) | **Get** /sapi/v1/sub-account/transfer/subUserHistory | Sub-account Transfer History(For Sub-account)
[**SubaccountGetSubAccountUniversalTransferV1**](V1API.md#SubaccountGetSubAccountUniversalTransferV1) | **Get** /sapi/v1/sub-account/universalTransfer | Query Universal Transfer History(For Master Account)



## SubaccountCreateManagedSubaccountDepositV1

> SubaccountCreateManagedSubaccountDepositV1Resp SubaccountCreateManagedSubaccountDepositV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()

Deposit Assets Into The Managed Sub-account(For Investor Master Account)



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
	resp, r, err := apiClient.V1API.SubaccountCreateManagedSubaccountDepositV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateManagedSubaccountDepositV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateManagedSubaccountDepositV1`: SubaccountCreateManagedSubaccountDepositV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateManagedSubaccountDepositV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateManagedSubaccountDepositV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **toEmail** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateManagedSubaccountDepositV1Resp**](SubaccountCreateManagedSubaccountDepositV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateManagedSubaccountWithdrawV1

> SubaccountCreateManagedSubaccountWithdrawV1Resp SubaccountCreateManagedSubaccountWithdrawV1(ctx).Amount(amount).Asset(asset).FromEmail(fromEmail).Timestamp(timestamp).RecvWindow(recvWindow).TransferDate(transferDate).Execute()

Withdrawl Assets From The Managed Sub-account(For Investor Master Account)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	transferDate := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountCreateManagedSubaccountWithdrawV1(context.Background()).Amount(amount).Asset(asset).FromEmail(fromEmail).Timestamp(timestamp).RecvWindow(recvWindow).TransferDate(transferDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateManagedSubaccountWithdrawV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateManagedSubaccountWithdrawV1`: SubaccountCreateManagedSubaccountWithdrawV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateManagedSubaccountWithdrawV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateManagedSubaccountWithdrawV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **fromEmail** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **transferDate** | **int64** |  | 

### Return type

[**SubaccountCreateManagedSubaccountWithdrawV1Resp**](SubaccountCreateManagedSubaccountWithdrawV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountBlvtEnableV1

> SubaccountCreateSubAccountBlvtEnableV1Resp SubaccountCreateSubAccountBlvtEnableV1(ctx).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Leverage Token for Sub-account(For Master Account)



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
	email := "email_example" // string |  (default to "")
	enableBlvt := true // bool | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountBlvtEnableV1(context.Background()).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountBlvtEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountBlvtEnableV1`: SubaccountCreateSubAccountBlvtEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountBlvtEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountBlvtEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **enableBlvt** | **bool** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountBlvtEnableV1Resp**](SubaccountCreateSubAccountBlvtEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountEoptionsEnableV1

> SubaccountCreateSubAccountEoptionsEnableV1Resp SubaccountCreateSubAccountEoptionsEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Options for Sub-account(For Master Account)(USER_DATA)



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
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountEoptionsEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountEoptionsEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountEoptionsEnableV1`: SubaccountCreateSubAccountEoptionsEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountEoptionsEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountEoptionsEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountEoptionsEnableV1Resp**](SubaccountCreateSubAccountEoptionsEnableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountFuturesEnableV1

> SubaccountCreateSubAccountFuturesEnableV1Resp SubaccountCreateSubAccountFuturesEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Futures for Sub-account(For Master Account)



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
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountFuturesEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountFuturesEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesEnableV1`: SubaccountCreateSubAccountFuturesEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountFuturesEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountFuturesEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountFuturesEnableV1Resp**](SubaccountCreateSubAccountFuturesEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountFuturesInternalTransferV1(context.Background()).Amount(amount).Asset(asset).FromEmail(fromEmail).FuturesType(futuresType).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountFuturesInternalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesInternalTransferV1`: SubaccountCreateSubAccountFuturesInternalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountFuturesInternalTransferV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountFuturesMovePositionV1(context.Background()).FromUserEmail(fromUserEmail).OrderArgs(orderArgs).ProductType(productType).Timestamp(timestamp).ToUserEmail(toUserEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountFuturesMovePositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesMovePositionV1`: SubaccountCreateSubAccountFuturesMovePositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountFuturesMovePositionV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountFuturesTransferV1(context.Background()).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountFuturesTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesTransferV1`: SubaccountCreateSubAccountFuturesTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountFuturesTransferV1`: %v\n", resp)
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


## SubaccountCreateSubAccountMarginEnableV1

> SubaccountCreateSubAccountMarginEnableV1Resp SubaccountCreateSubAccountMarginEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Margin for Sub-account(For Master Account)



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
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountMarginEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountMarginEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountMarginEnableV1`: SubaccountCreateSubAccountMarginEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountMarginEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountMarginEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountMarginEnableV1Resp**](SubaccountCreateSubAccountMarginEnableV1Resp.md)

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
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountMarginTransferV1(context.Background()).Amount(amount).Asset(asset).Email(email).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountMarginTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountMarginTransferV1`: SubaccountCreateSubAccountMarginTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountMarginTransferV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountTransferSubToMasterV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountTransferSubToMasterV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountTransferSubToMasterV1`: SubaccountCreateSubAccountTransferSubToMasterV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountTransferSubToMasterV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountTransferSubToSubV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountTransferSubToSubV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountTransferSubToSubV1`: SubaccountCreateSubAccountTransferSubToSubV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountTransferSubToSubV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountUniversalTransferV1(context.Background()).Amount(amount).Asset(asset).FromAccountType(fromAccountType).Timestamp(timestamp).ToAccountType(toAccountType).ClientTranId(clientTranId).FromEmail(fromEmail).RecvWindow(recvWindow).Symbol(symbol).ToEmail(toEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountUniversalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountUniversalTransferV1`: SubaccountCreateSubAccountUniversalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountUniversalTransferV1`: %v\n", resp)
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


## SubaccountCreateSubAccountVirtualSubAccountV1

> SubaccountCreateSubAccountVirtualSubAccountV1Resp SubaccountCreateSubAccountVirtualSubAccountV1(ctx).SubAccountString(subAccountString).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Create a Virtual Sub-account(For Master Account)



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
	subAccountString := "subAccountString_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountCreateSubAccountVirtualSubAccountV1(context.Background()).SubAccountString(subAccountString).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountCreateSubAccountVirtualSubAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountVirtualSubAccountV1`: SubaccountCreateSubAccountVirtualSubAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountCreateSubAccountVirtualSubAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountVirtualSubAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAccountString** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountVirtualSubAccountV1Resp**](SubaccountCreateSubAccountVirtualSubAccountV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1

> SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()

Delete IP List For a Sub-account API Key(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key#email-address\">Sub-account email</a> (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ipAddress := "ipAddress_example" // string | Can be added in batches, separated by commas (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1`: SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **ipAddress** | **string** | Can be added in batches, separated by commas | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp**](SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.V1API.SubaccountGetCapitalDepositSubAddressV1(context.Background()).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetCapitalDepositSubAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetCapitalDepositSubAddressV1`: SubaccountGetCapitalDepositSubAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetCapitalDepositSubAddressV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountGetCapitalDepositSubHisrecV1(context.Background()).Email(email).Timestamp(timestamp).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Limit(limit).Offset(offset).RecvWindow(recvWindow).TxId(txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetCapitalDepositSubHisrecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetCapitalDepositSubHisrecV1`: []SubaccountGetCapitalDepositSubHisrecV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetCapitalDepositSubHisrecV1`: %v\n", resp)
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


## SubaccountGetManagedSubaccountAccountSnapshotV1

> SubaccountGetManagedSubaccountAccountSnapshotV1Resp SubaccountGetManagedSubaccountAccountSnapshotV1(ctx).Email(email).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Managed Sub-account Snapshot(For Investor Master Account)



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
	email := "email_example" // string |  (default to "")
	type_ := "type__example" // string | &#34;SPOT&#34;, &#34;MARGIN&#34;（cross）, &#34;FUTURES&#34;（UM） (default to "")
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | min 7, max 30, default 7 (optional) (default to 7)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountAccountSnapshotV1(context.Background()).Email(email).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountAccountSnapshotV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountAccountSnapshotV1`: SubaccountGetManagedSubaccountAccountSnapshotV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountAccountSnapshotV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountAccountSnapshotV1Request struct via the builder pattern


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

[**SubaccountGetManagedSubaccountAccountSnapshotV1Resp**](SubaccountGetManagedSubaccountAccountSnapshotV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountAssetV1

> []SubaccountGetManagedSubaccountAssetV1RespItem SubaccountGetManagedSubaccountAssetV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Managed Sub-account Asset Details(For Investor Master Account)



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
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountAssetV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountAssetV1`: []SubaccountGetManagedSubaccountAssetV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]SubaccountGetManagedSubaccountAssetV1RespItem**](SubaccountGetManagedSubaccountAssetV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountDepositAddressV1

> SubaccountGetManagedSubaccountDepositAddressV1Resp SubaccountGetManagedSubaccountDepositAddressV1(ctx).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Get Managed Sub-account Deposit Address (For Investor Master Account)(USER_DATA)



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
	email := "email_example" // string | Sub user email (default to "")
	coin := "coin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	network := "network_example" // string | networks can be found in `GET /sapi/v1/capital/deposit/address` (optional) (default to "")
	amount := "amount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountDepositAddressV1(context.Background()).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountDepositAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountDepositAddressV1`: SubaccountGetManagedSubaccountDepositAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountDepositAddressV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountDepositAddressV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub user email | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** | networks can be found in &#x60;GET /sapi/v1/capital/deposit/address&#x60; | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetManagedSubaccountDepositAddressV1Resp**](SubaccountGetManagedSubaccountDepositAddressV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountFetchFutureAssetV1

> SubaccountGetManagedSubaccountFetchFutureAssetV1Resp SubaccountGetManagedSubaccountFetchFutureAssetV1(ctx).Email(email).AccountType(accountType).Execute()

Query Managed Sub-account Futures Asset Details(For Investor Master Account)(USER_DATA)



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
	email := "email_example" // string | Managed Sub Account Email (default to "")
	accountType := "accountType_example" // string | No input or input &#34;USDT_FUTURE&#34; to get UM Futures account details. Input &#34;COIN_FUTURE&#34; to get CM Futures account details. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountFetchFutureAssetV1(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountFetchFutureAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountFetchFutureAssetV1`: SubaccountGetManagedSubaccountFetchFutureAssetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountFetchFutureAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountFetchFutureAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed Sub Account Email | [default to &quot;&quot;]
 **accountType** | **string** | No input or input &amp;#34;USDT_FUTURE&amp;#34; to get UM Futures account details. Input &amp;#34;COIN_FUTURE&amp;#34; to get CM Futures account details. | [default to &quot;&quot;]

### Return type

[**SubaccountGetManagedSubaccountFetchFutureAssetV1Resp**](SubaccountGetManagedSubaccountFetchFutureAssetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountInfoV1

> SubaccountGetManagedSubaccountInfoV1Resp SubaccountGetManagedSubaccountInfoV1(ctx).Timestamp(timestamp).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Managed Sub-account List(For Investor)(USER_DATA)



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
	email := "email_example" // string | Managed sub-account email (optional) (default to "")
	page := int32(56) // int32 | Default value: 1 (optional)
	limit := int32(56) // int32 | Default value: 20, Max value: 20 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountInfoV1(context.Background()).Timestamp(timestamp).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountInfoV1`: SubaccountGetManagedSubaccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | Managed sub-account email | [default to &quot;&quot;]
 **page** | **int32** | Default value: 1 | 
 **limit** | **int32** | Default value: 20, Max value: 20 | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetManagedSubaccountInfoV1Resp**](SubaccountGetManagedSubaccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountMarginAssetV1

> SubaccountGetManagedSubaccountMarginAssetV1Resp SubaccountGetManagedSubaccountMarginAssetV1(ctx).Email(email).AccountType(accountType).Execute()

Query Managed Sub-account Margin Asset Details(For Investor Master Account)(USER_DATA)



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
	email := "email_example" // string | Managed Sub Account Email (default to "")
	accountType := "accountType_example" // string | No input or input &#34;MARGIN&#34; to get Cross Margin account details. Input &#34;ISOLATED_MARGIN&#34; to get Isolated Margin account details. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountMarginAssetV1(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountMarginAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountMarginAssetV1`: SubaccountGetManagedSubaccountMarginAssetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountMarginAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountMarginAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Managed Sub Account Email | [default to &quot;&quot;]
 **accountType** | **string** | No input or input &amp;#34;MARGIN&amp;#34; to get Cross Margin account details. Input &amp;#34;ISOLATED_MARGIN&amp;#34; to get Isolated Margin account details. | [default to &quot;&quot;]

### Return type

[**SubaccountGetManagedSubaccountMarginAssetV1Resp**](SubaccountGetManagedSubaccountMarginAssetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountQueryTransLogForInvestorV1

> SubaccountGetManagedSubaccountQueryTransLogForInvestorV1Resp SubaccountGetManagedSubaccountQueryTransLogForInvestorV1(ctx).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()

Query Managed Sub Account Transfer Log(For Investor Master Account)(USER_DATA)



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
	email := "email_example" // string | Managed Sub Account Email (default to "")
	startTime := int64(789) // int64 | Start Time
	endTime := int64(789) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int32(56) // int32 | Page
	limit := int32(56) // int32 | Limit (Max: 500)
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional) (default to "")
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountQueryTransLogForInvestorV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountQueryTransLogForInvestorV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountQueryTransLogForInvestorV1`: SubaccountGetManagedSubaccountQueryTransLogForInvestorV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountQueryTransLogForInvestorV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountQueryTransLogForInvestorV1Request struct via the builder pattern


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

[**SubaccountGetManagedSubaccountQueryTransLogForInvestorV1Resp**](SubaccountGetManagedSubaccountQueryTransLogForInvestorV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1

> SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1Resp SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1(ctx).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()

Query Managed Sub Account Transfer Log(For Trading Team Master Account)(USER_DATA)



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
	email := "email_example" // string | Managed Sub Account Email (default to "")
	startTime := int64(789) // int64 | Start Time
	endTime := int64(789) // int64 | End Time (The start time and end time interval cannot exceed half a year)
	page := int32(56) // int32 | Page
	limit := int32(56) // int32 | Limit (Max: 500)
	transfers := "transfers_example" // string | Transfer Direction (FROM/TO) (optional) (default to "")
	transferFunctionAccountType := "transferFunctionAccountType_example" // string | Transfer function account type (SPOT/MARGIN/ISOLATED_MARGIN/USDT_FUTURE/COIN_FUTURE) (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1`: SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountQueryTransLogForTradeParentV1Request struct via the builder pattern


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

[**SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1Resp**](SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetManagedSubaccountQueryTransLogV1

> SubaccountGetManagedSubaccountQueryTransLogV1Resp SubaccountGetManagedSubaccountQueryTransLogV1(ctx).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Timestamp(timestamp).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).RecvWindow(recvWindow).Execute()

Query Managed Sub Account Transfer Log (For Trading Team Sub Account)(USER_DATA)



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
	resp, r, err := apiClient.V1API.SubaccountGetManagedSubaccountQueryTransLogV1(context.Background()).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Timestamp(timestamp).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetManagedSubaccountQueryTransLogV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountQueryTransLogV1`: SubaccountGetManagedSubaccountQueryTransLogV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetManagedSubaccountQueryTransLogV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetManagedSubaccountQueryTransLogV1Request struct via the builder pattern


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

[**SubaccountGetManagedSubaccountQueryTransLogV1Resp**](SubaccountGetManagedSubaccountQueryTransLogV1Resp.md)

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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountFuturesAccountSummaryV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountFuturesAccountSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountSummaryV1`: SubaccountGetSubAccountFuturesAccountSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountFuturesAccountSummaryV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountFuturesAccountV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountFuturesAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountV1`: SubaccountGetSubAccountFuturesAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountFuturesAccountV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountFuturesInternalTransferV1(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountFuturesInternalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesInternalTransferV1`: SubaccountGetSubAccountFuturesInternalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountFuturesInternalTransferV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountFuturesMovePositionV1(context.Background()).Symbol(symbol).Page(page).Row(row).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountFuturesMovePositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesMovePositionV1`: SubaccountGetSubAccountFuturesMovePositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountFuturesMovePositionV1`: %v\n", resp)
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


## SubaccountGetSubAccountFuturesPositionRiskV1

> []SubaccountGetSubAccountFuturesPositionRiskV1RespItem SubaccountGetSubAccountFuturesPositionRiskV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Position-Risk of Sub-account(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account#email-address\">Sub-account email</a> (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountFuturesPositionRiskV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountFuturesPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesPositionRiskV1`: []SubaccountGetSubAccountFuturesPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountFuturesPositionRiskV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesPositionRiskV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]SubaccountGetSubAccountFuturesPositionRiskV1RespItem**](SubaccountGetSubAccountFuturesPositionRiskV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountListV1

> SubaccountGetSubAccountListV1Resp SubaccountGetSubAccountListV1(ctx).Timestamp(timestamp).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account List(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Query-Sub-account-List#email-address\">Sub-account email</a> (optional) (default to "")
	isFreeze := "isFreeze_example" // string | true or false (optional) (default to "")
	page := int32(56) // int32 | Default value: 1 (optional)
	limit := int32(56) // int32 | Default value: 1, Max value: 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountListV1(context.Background()).Timestamp(timestamp).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountListV1`: SubaccountGetSubAccountListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Query-Sub-account-List#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **isFreeze** | **string** | true or false | [default to &quot;&quot;]
 **page** | **int32** | Default value: 1 | 
 **limit** | **int32** | Default value: 1, Max value: 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountListV1Resp**](SubaccountGetSubAccountListV1Resp.md)

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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountMarginAccountSummaryV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountMarginAccountSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountMarginAccountSummaryV1`: SubaccountGetSubAccountMarginAccountSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountMarginAccountSummaryV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountMarginAccountV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountMarginAccountV1`: SubaccountGetSubAccountMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountMarginAccountV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountSpotSummaryV1(context.Background()).Timestamp(timestamp).Email(email).Page(page).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountSpotSummaryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountSpotSummaryV1`: SubaccountGetSubAccountSpotSummaryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountSpotSummaryV1`: %v\n", resp)
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


## SubaccountGetSubAccountStatusV1

> []SubaccountGetSubAccountStatusV1RespItem SubaccountGetSubAccountStatusV1(ctx).Timestamp(timestamp).Email(email).RecvWindow(recvWindow).Execute()

Get Sub-account's Status on Margin Or Futures(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures#email-address\">Sub-account email</a> (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountStatusV1(context.Background()).Timestamp(timestamp).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountStatusV1`: []SubaccountGetSubAccountStatusV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]SubaccountGetSubAccountStatusV1RespItem**](SubaccountGetSubAccountStatusV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountSubAccountApiIpRestrictionV1

> SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp SubaccountGetSubAccountSubAccountApiIpRestrictionV1(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get IP Restriction for a Sub-account API Key(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/api-management#email-address\">Sub-account email</a> (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountSubAccountApiIpRestrictionV1(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountSubAccountApiIpRestrictionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountSubAccountApiIpRestrictionV1`: SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountSubAccountApiIpRestrictionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/api-management#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp**](SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp.md)

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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountSubTransferHistoryV1(context.Background()).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountSubTransferHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountSubTransferHistoryV1`: []SubaccountGetSubAccountSubTransferHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountSubTransferHistoryV1`: %v\n", resp)
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


## SubaccountGetSubAccountTransactionStatisticsV1

> SubaccountGetSubAccountTransactionStatisticsV1Resp SubaccountGetSubAccountTransactionStatisticsV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)



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
	email := "email_example" // string | Sub user email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountTransactionStatisticsV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountTransactionStatisticsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountTransactionStatisticsV1`: SubaccountGetSubAccountTransactionStatisticsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountTransactionStatisticsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountTransactionStatisticsV1Request struct via the builder pattern


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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountTransferSubUserHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).ReturnFailHistory(returnFailHistory).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountTransferSubUserHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountTransferSubUserHistoryV1`: []SubaccountGetSubAccountTransferSubUserHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountTransferSubUserHistoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.SubaccountGetSubAccountUniversalTransferV1(context.Background()).Timestamp(timestamp).FromEmail(fromEmail).ToEmail(toEmail).ClientTranId(clientTranId).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.SubaccountGetSubAccountUniversalTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountUniversalTransferV1`: SubaccountGetSubAccountUniversalTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.SubaccountGetSubAccountUniversalTransferV1`: %v\n", resp)
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

