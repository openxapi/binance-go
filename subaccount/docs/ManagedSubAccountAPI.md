# \ManagedSubAccountAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreateManagedSubaccountDepositV1**](ManagedSubAccountAPI.md#SubaccountCreateManagedSubaccountDepositV1) | **Post** /sapi/v1/managed-subaccount/deposit | Deposit Assets Into The Managed Sub-account(For Investor Master Account)
[**SubaccountCreateManagedSubaccountWithdrawV1**](ManagedSubAccountAPI.md#SubaccountCreateManagedSubaccountWithdrawV1) | **Post** /sapi/v1/managed-subaccount/withdraw | Withdrawl Assets From The Managed Sub-account(For Investor Master Account)
[**SubaccountGetManagedSubaccountAccountSnapshotV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountAccountSnapshotV1) | **Get** /sapi/v1/managed-subaccount/accountSnapshot | Query Managed Sub-account Snapshot(For Investor Master Account)
[**SubaccountGetManagedSubaccountAssetV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountAssetV1) | **Get** /sapi/v1/managed-subaccount/asset | Query Managed Sub-account Asset Details(For Investor Master Account)
[**SubaccountGetManagedSubaccountDepositAddressV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountDepositAddressV1) | **Get** /sapi/v1/managed-subaccount/deposit/address | Get Managed Sub-account Deposit Address (For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountFetchFutureAssetV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountFetchFutureAssetV1) | **Get** /sapi/v1/managed-subaccount/fetch-future-asset | Query Managed Sub-account Futures Asset Details(For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountInfoV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountInfoV1) | **Get** /sapi/v1/managed-subaccount/info | Query Managed Sub-account List(For Investor)(USER_DATA)
[**SubaccountGetManagedSubaccountMarginAssetV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountMarginAssetV1) | **Get** /sapi/v1/managed-subaccount/marginAsset | Query Managed Sub-account Margin Asset Details(For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountQueryTransLogForInvestorV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountQueryTransLogForInvestorV1) | **Get** /sapi/v1/managed-subaccount/queryTransLogForInvestor | Query Managed Sub Account Transfer Log(For Investor Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1) | **Get** /sapi/v1/managed-subaccount/queryTransLogForTradeParent | Query Managed Sub Account Transfer Log(For Trading Team Master Account)(USER_DATA)
[**SubaccountGetManagedSubaccountQueryTransLogV1**](ManagedSubAccountAPI.md#SubaccountGetManagedSubaccountQueryTransLogV1) | **Get** /sapi/v1/managed-subaccount/query-trans-log | Query Managed Sub Account Transfer Log (For Trading Team Sub Account)(USER_DATA)



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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountCreateManagedSubaccountDepositV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).ToEmail(toEmail).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountCreateManagedSubaccountDepositV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateManagedSubaccountDepositV1`: SubaccountCreateManagedSubaccountDepositV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountCreateManagedSubaccountDepositV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountCreateManagedSubaccountWithdrawV1(context.Background()).Amount(amount).Asset(asset).FromEmail(fromEmail).Timestamp(timestamp).RecvWindow(recvWindow).TransferDate(transferDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountCreateManagedSubaccountWithdrawV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateManagedSubaccountWithdrawV1`: SubaccountCreateManagedSubaccountWithdrawV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountCreateManagedSubaccountWithdrawV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountAccountSnapshotV1(context.Background()).Email(email).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountAccountSnapshotV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountAccountSnapshotV1`: SubaccountGetManagedSubaccountAccountSnapshotV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountAccountSnapshotV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountAssetV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountAssetV1`: []SubaccountGetManagedSubaccountAssetV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountAssetV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountDepositAddressV1(context.Background()).Email(email).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountDepositAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountDepositAddressV1`: SubaccountGetManagedSubaccountDepositAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountDepositAddressV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountFetchFutureAssetV1(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountFetchFutureAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountFetchFutureAssetV1`: SubaccountGetManagedSubaccountFetchFutureAssetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountFetchFutureAssetV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountInfoV1(context.Background()).Timestamp(timestamp).Email(email).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountInfoV1`: SubaccountGetManagedSubaccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountInfoV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountMarginAssetV1(context.Background()).Email(email).AccountType(accountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountMarginAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountMarginAssetV1`: SubaccountGetManagedSubaccountMarginAssetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountMarginAssetV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogForInvestorV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogForInvestorV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountQueryTransLogForInvestorV1`: SubaccountGetManagedSubaccountQueryTransLogForInvestorV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogForInvestorV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1`: SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogForTradeParentV1`: %v\n", resp)
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
	resp, r, err := apiClient.ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogV1(context.Background()).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).Timestamp(timestamp).Transfers(transfers).TransferFunctionAccountType(transferFunctionAccountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetManagedSubaccountQueryTransLogV1`: SubaccountGetManagedSubaccountQueryTransLogV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ManagedSubAccountAPI.SubaccountGetManagedSubaccountQueryTransLogV1`: %v\n", resp)
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

