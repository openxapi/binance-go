# \WalletAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountDisableFastWithdrawSwitchV1**](WalletAPI.md#CreateAccountDisableFastWithdrawSwitchV1) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**CreateAccountEnableFastWithdrawSwitchV1**](WalletAPI.md#CreateAccountEnableFastWithdrawSwitchV1) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**CreateAssetDustBtcV1**](WalletAPI.md#CreateAssetDustBtcV1) | **Post** /sapi/v1/asset/dust-btc | Get Assets That Can Be Converted Into BNB (USER_DATA)
[**CreateAssetDustV1**](WalletAPI.md#CreateAssetDustV1) | **Post** /sapi/v1/asset/dust | Dust Transfer (USER_DATA)
[**CreateAssetGetFundingAssetV1**](WalletAPI.md#CreateAssetGetFundingAssetV1) | **Post** /sapi/v1/asset/get-funding-asset | Funding Wallet (USER_DATA)
[**CreateAssetGetUserAssetV3**](WalletAPI.md#CreateAssetGetUserAssetV3) | **Post** /sapi/v3/asset/getUserAsset | User Asset (USER_DATA)
[**CreateAssetTransferV1**](WalletAPI.md#CreateAssetTransferV1) | **Post** /sapi/v1/asset/transfer | User Universal Transfer (USER_DATA)
[**CreateBnbBurnV1**](WalletAPI.md#CreateBnbBurnV1) | **Post** /sapi/v1/bnbBurn | Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
[**CreateCapitalDepositCreditApplyV1**](WalletAPI.md#CreateCapitalDepositCreditApplyV1) | **Post** /sapi/v1/capital/deposit/credit-apply | One click arrival deposit apply (for expired address deposit) (USER_DATA)
[**CreateCapitalWithdrawApplyV1**](WalletAPI.md#CreateCapitalWithdrawApplyV1) | **Post** /sapi/v1/capital/withdraw/apply | Withdraw(USER_DATA)
[**CreateLocalentityBrokerWithdrawApplyV1**](WalletAPI.md#CreateLocalentityBrokerWithdrawApplyV1) | **Post** /sapi/v1/localentity/broker/withdraw/apply | Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)
[**CreateLocalentityWithdrawApplyV1**](WalletAPI.md#CreateLocalentityWithdrawApplyV1) | **Post** /sapi/v1/localentity/withdraw/apply | Withdraw (for local entities that require travel rule) (USER_DATA)
[**GetAccountApiRestrictionsV1**](WalletAPI.md#GetAccountApiRestrictionsV1) | **Get** /sapi/v1/account/apiRestrictions | Get API Key Permission (USER_DATA)
[**GetAccountApiTradingStatusV1**](WalletAPI.md#GetAccountApiTradingStatusV1) | **Get** /sapi/v1/account/apiTradingStatus | Account API Trading Status (USER_DATA)
[**GetAccountInfoV1**](WalletAPI.md#GetAccountInfoV1) | **Get** /sapi/v1/account/info | Account info (USER_DATA)
[**GetAccountSnapshotV1**](WalletAPI.md#GetAccountSnapshotV1) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**GetAccountStatusV1**](WalletAPI.md#GetAccountStatusV1) | **Get** /sapi/v1/account/status | Account Status (USER_DATA)
[**GetAssetAssetDetailV1**](WalletAPI.md#GetAssetAssetDetailV1) | **Get** /sapi/v1/asset/assetDetail | Asset Detail (USER_DATA)
[**GetAssetAssetDividendV1**](WalletAPI.md#GetAssetAssetDividendV1) | **Get** /sapi/v1/asset/assetDividend | Asset Dividend Record (USER_DATA)
[**GetAssetCustodyTransferHistoryV1**](WalletAPI.md#GetAssetCustodyTransferHistoryV1) | **Get** /sapi/v1/asset/custody/transfer-history | Query User Delegation History(For Master Account)(USER_DATA)
[**GetAssetDribbletV1**](WalletAPI.md#GetAssetDribbletV1) | **Get** /sapi/v1/asset/dribblet | DustLog(USER_DATA)
[**GetAssetLedgerTransferCloudMiningQueryByPageV1**](WalletAPI.md#GetAssetLedgerTransferCloudMiningQueryByPageV1) | **Get** /sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage | Get Cloud-Mining payment and refund history (USER_DATA)
[**GetAssetTradeFeeV1**](WalletAPI.md#GetAssetTradeFeeV1) | **Get** /sapi/v1/asset/tradeFee | Trade Fee (USER_DATA)
[**GetAssetTransferV1**](WalletAPI.md#GetAssetTransferV1) | **Get** /sapi/v1/asset/transfer | Query User Universal Transfer History(USER_DATA)
[**GetAssetWalletBalanceV1**](WalletAPI.md#GetAssetWalletBalanceV1) | **Get** /sapi/v1/asset/wallet/balance | Query User Wallet Balance (USER_DATA)
[**GetCapitalConfigGetallV1**](WalletAPI.md#GetCapitalConfigGetallV1) | **Get** /sapi/v1/capital/config/getall | All Coins&#39; Information (USER_DATA)
[**GetCapitalDepositAddressListV1**](WalletAPI.md#GetCapitalDepositAddressListV1) | **Get** /sapi/v1/capital/deposit/address/list | Fetch deposit address list with network(USER_DATA)
[**GetCapitalDepositAddressV1**](WalletAPI.md#GetCapitalDepositAddressV1) | **Get** /sapi/v1/capital/deposit/address | Deposit Address(supporting network) (USER_DATA)
[**GetCapitalDepositHisrecV1**](WalletAPI.md#GetCapitalDepositHisrecV1) | **Get** /sapi/v1/capital/deposit/hisrec | Deposit History (supporting network) (USER_DATA)
[**GetCapitalWithdrawAddressListV1**](WalletAPI.md#GetCapitalWithdrawAddressListV1) | **Get** /sapi/v1/capital/withdraw/address/list | Fetch withdraw address list (USER_DATA)
[**GetCapitalWithdrawHistoryV1**](WalletAPI.md#GetCapitalWithdrawHistoryV1) | **Get** /sapi/v1/capital/withdraw/history | Withdraw History (supporting network) (USER_DATA)
[**GetLocalentityDepositHistoryV1**](WalletAPI.md#GetLocalentityDepositHistoryV1) | **Get** /sapi/v1/localentity/deposit/history | Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)
[**GetLocalentityVaspV1**](WalletAPI.md#GetLocalentityVaspV1) | **Get** /sapi/v1/localentity/vasp | Onboarded VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)
[**GetLocalentityWithdrawHistoryV1**](WalletAPI.md#GetLocalentityWithdrawHistoryV1) | **Get** /sapi/v1/localentity/withdraw/history | Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)
[**GetLocalentityWithdrawHistoryV2**](WalletAPI.md#GetLocalentityWithdrawHistoryV2) | **Get** /sapi/v2/localentity/withdraw/history | Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)
[**GetSpotDelistScheduleV1**](WalletAPI.md#GetSpotDelistScheduleV1) | **Get** /sapi/v1/spot/delist-schedule | Get Spot Delist Schedule (MARKET_DATA)
[**GetSpotOpenSymbolListV1**](WalletAPI.md#GetSpotOpenSymbolListV1) | **Get** /sapi/v1/spot/open-symbol-list | Get Open Symbol List (MARKET_DATA)
[**GetSystemStatusV1**](WalletAPI.md#GetSystemStatusV1) | **Get** /sapi/v1/system/status | System Status (System)
[**UpdateLocalentityBrokerDepositProvideInfoV1**](WalletAPI.md#UpdateLocalentityBrokerDepositProvideInfoV1) | **Put** /sapi/v1/localentity/broker/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)
[**UpdateLocalentityDepositProvideInfoV1**](WalletAPI.md#UpdateLocalentityDepositProvideInfoV1) | **Put** /sapi/v1/localentity/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



## CreateAccountDisableFastWithdrawSwitchV1

> map[string]interface{} CreateAccountDisableFastWithdrawSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Disable Fast Withdraw Switch (USER_DATA)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateAccountDisableFastWithdrawSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateAccountDisableFastWithdrawSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountDisableFastWithdrawSwitchV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateAccountDisableFastWithdrawSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountDisableFastWithdrawSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountEnableFastWithdrawSwitchV1

> map[string]interface{} CreateAccountEnableFastWithdrawSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Fast Withdraw Switch (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateAccountEnableFastWithdrawSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateAccountEnableFastWithdrawSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountEnableFastWithdrawSwitchV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateAccountEnableFastWithdrawSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountEnableFastWithdrawSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetDustBtcV1

> CreateAssetDustBtcV1Resp CreateAssetDustBtcV1(ctx).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()

Get Assets That Can Be Converted Into BNB (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	accountType := "accountType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateAssetDustBtcV1(context.Background()).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateAssetDustBtcV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetDustBtcV1`: CreateAssetDustBtcV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateAssetDustBtcV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetDustBtcV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **accountType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateAssetDustBtcV1Resp**](CreateAssetDustBtcV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetDustV1

> CreateAssetDustV1Resp CreateAssetDustV1(ctx).Asset(asset).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()

Dust Transfer (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	asset := []string{"Inner_example"} // []string | 
	timestamp := int64(789) // int64 | 
	accountType := "accountType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateAssetDustV1(context.Background()).Asset(asset).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateAssetDustV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetDustV1`: CreateAssetDustV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateAssetDustV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetDustV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **[]string** |  | 
 **timestamp** | **int64** |  | 
 **accountType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateAssetDustV1Resp**](CreateAssetDustV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetGetFundingAssetV1

> []CreateAssetGetFundingAssetV1RespItem CreateAssetGetFundingAssetV1(ctx).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

Funding Wallet (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	needBtcValuation := "needBtcValuation_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateAssetGetFundingAssetV1(context.Background()).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateAssetGetFundingAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetGetFundingAssetV1`: []CreateAssetGetFundingAssetV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateAssetGetFundingAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetGetFundingAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **needBtcValuation** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CreateAssetGetFundingAssetV1RespItem**](CreateAssetGetFundingAssetV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetGetUserAssetV3

> []CreateAssetGetUserAssetV3RespItem CreateAssetGetUserAssetV3(ctx).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

User Asset (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	needBtcValuation := true // bool |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateAssetGetUserAssetV3(context.Background()).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateAssetGetUserAssetV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetGetUserAssetV3`: []CreateAssetGetUserAssetV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateAssetGetUserAssetV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetGetUserAssetV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **needBtcValuation** | **bool** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]CreateAssetGetUserAssetV3RespItem**](CreateAssetGetUserAssetV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetTransferV1

> CreateAssetTransferV1Resp CreateAssetTransferV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).Type_(type_).FromSymbol(fromSymbol).RecvWindow(recvWindow).ToSymbol(toSymbol).Execute()

User Universal Transfer (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	fromSymbol := "fromSymbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	toSymbol := "toSymbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateAssetTransferV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).Type_(type_).FromSymbol(fromSymbol).RecvWindow(recvWindow).ToSymbol(toSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateAssetTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetTransferV1`: CreateAssetTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateAssetTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **fromSymbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **toSymbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateAssetTransferV1Resp**](CreateAssetTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBnbBurnV1

> CreateBnbBurnV1Resp CreateBnbBurnV1(ctx).Timestamp(timestamp).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).SpotBNBBurn(spotBNBBurn).Execute()

Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	interestBNBBurn := "interestBNBBurn_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	spotBNBBurn := "spotBNBBurn_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateBnbBurnV1(context.Background()).Timestamp(timestamp).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).SpotBNBBurn(spotBNBBurn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateBnbBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBnbBurnV1`: CreateBnbBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateBnbBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBnbBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **interestBNBBurn** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **spotBNBBurn** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateBnbBurnV1Resp**](CreateBnbBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCapitalDepositCreditApplyV1

> CreateCapitalDepositCreditApplyV1Resp CreateCapitalDepositCreditApplyV1(ctx).DepositId(depositId).SubAccountId(subAccountId).SubUserId(subUserId).TxId(txId).Execute()

One click arrival deposit apply (for expired address deposit) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	depositId := int64(789) // int64 |  (optional)
	subAccountId := int64(789) // int64 |  (optional)
	subUserId := int64(789) // int64 |  (optional)
	txId := "txId_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateCapitalDepositCreditApplyV1(context.Background()).DepositId(depositId).SubAccountId(subAccountId).SubUserId(subUserId).TxId(txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateCapitalDepositCreditApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCapitalDepositCreditApplyV1`: CreateCapitalDepositCreditApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateCapitalDepositCreditApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCapitalDepositCreditApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositId** | **int64** |  | 
 **subAccountId** | **int64** |  | 
 **subUserId** | **int64** |  | 
 **txId** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateCapitalDepositCreditApplyV1Resp**](CreateCapitalDepositCreditApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCapitalWithdrawApplyV1

> CreateCapitalWithdrawApplyV1Resp CreateCapitalWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()

Withdraw(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	address := "address_example" // string |  (default to "")
	amount := "amount_example" // string |  (default to "")
	coin := "coin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	addressTag := "addressTag_example" // string |  (optional) (default to "")
	name := "name_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	transactionFeeFlag := true // bool |  (optional) (default to false)
	walletType := int32(56) // int32 |  (optional)
	withdrawOrderId := "withdrawOrderId_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateCapitalWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateCapitalWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCapitalWithdrawApplyV1`: CreateCapitalWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateCapitalWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCapitalWithdrawApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **addressTag** | **string** |  | [default to &quot;&quot;]
 **name** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **transactionFeeFlag** | **bool** |  | [default to false]
 **walletType** | **int32** |  | 
 **withdrawOrderId** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateCapitalWithdrawApplyV1Resp**](CreateCapitalWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocalentityBrokerWithdrawApplyV1

> CreateLocalentityBrokerWithdrawApplyV1Resp CreateLocalentityBrokerWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).OriginatorPii(originatorPii).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).WithdrawOrderId(withdrawOrderId).AddressName(addressName).AddressTag(addressTag).Network(network).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()

Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	address := "address_example" // string |  (default to "")
	amount := "amount_example" // string |  (default to "")
	coin := "coin_example" // string |  (default to "")
	originatorPii := "originatorPii_example" // string |  (default to "")
	questionnaire := "questionnaire_example" // string |  (default to "")
	signature := "signature_example" // string |  (default to "")
	subAccountId := "subAccountId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	withdrawOrderId := "withdrawOrderId_example" // string |  (default to "")
	addressName := "addressName_example" // string |  (optional) (default to "")
	addressTag := "addressTag_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	transactionFeeFlag := true // bool |  (optional) (default to false)
	walletType := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateLocalentityBrokerWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).OriginatorPii(originatorPii).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).WithdrawOrderId(withdrawOrderId).AddressName(addressName).AddressTag(addressTag).Network(network).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateLocalentityBrokerWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLocalentityBrokerWithdrawApplyV1`: CreateLocalentityBrokerWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateLocalentityBrokerWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalentityBrokerWithdrawApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **originatorPii** | **string** |  | [default to &quot;&quot;]
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **signature** | **string** |  | [default to &quot;&quot;]
 **subAccountId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **withdrawOrderId** | **string** |  | [default to &quot;&quot;]
 **addressName** | **string** |  | [default to &quot;&quot;]
 **addressTag** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **transactionFeeFlag** | **bool** |  | [default to false]
 **walletType** | **int32** |  | 

### Return type

[**CreateLocalentityBrokerWithdrawApplyV1Resp**](CreateLocalentityBrokerWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocalentityWithdrawApplyV1

> CreateLocalentityWithdrawApplyV1Resp CreateLocalentityWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).Questionnaire(questionnaire).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()

Withdraw (for local entities that require travel rule) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	address := "address_example" // string |  (default to "")
	amount := "amount_example" // string |  (default to "")
	coin := "coin_example" // string |  (default to "")
	questionnaire := "questionnaire_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	addressTag := "addressTag_example" // string |  (optional) (default to "")
	name := "name_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	transactionFeeFlag := true // bool |  (optional) (default to false)
	walletType := int32(56) // int32 |  (optional)
	withdrawOrderId := "withdrawOrderId_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.CreateLocalentityWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).Questionnaire(questionnaire).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.CreateLocalentityWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLocalentityWithdrawApplyV1`: CreateLocalentityWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.CreateLocalentityWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalentityWithdrawApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **addressTag** | **string** |  | [default to &quot;&quot;]
 **name** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **transactionFeeFlag** | **bool** |  | [default to false]
 **walletType** | **int32** |  | 
 **withdrawOrderId** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateLocalentityWithdrawApplyV1Resp**](CreateLocalentityWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountApiRestrictionsV1

> GetAccountApiRestrictionsV1Resp GetAccountApiRestrictionsV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get API Key Permission (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAccountApiRestrictionsV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAccountApiRestrictionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountApiRestrictionsV1`: GetAccountApiRestrictionsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAccountApiRestrictionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountApiRestrictionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAccountApiRestrictionsV1Resp**](GetAccountApiRestrictionsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountApiTradingStatusV1

> GetAccountApiTradingStatusV1Resp GetAccountApiTradingStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account API Trading Status (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAccountApiTradingStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAccountApiTradingStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountApiTradingStatusV1`: GetAccountApiTradingStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAccountApiTradingStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountApiTradingStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAccountApiTradingStatusV1Resp**](GetAccountApiTradingStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInfoV1

> GetAccountInfoV1Resp GetAccountInfoV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account info (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAccountInfoV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInfoV1`: GetAccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAccountInfoV1Resp**](GetAccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSnapshotV1

> GetAccountSnapshotV1Resp GetAccountSnapshotV1(ctx).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Daily Account Snapshot (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	type_ := "type__example" // string | &#34;SPOT&#34;, &#34;MARGIN&#34;, &#34;FUTURES&#34; (default to "")
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | min 7, max 30, default 7 (optional) (default to 7)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAccountSnapshotV1(context.Background()).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAccountSnapshotV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSnapshotV1`: GetAccountSnapshotV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAccountSnapshotV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSnapshotV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &amp;#34;SPOT&amp;#34;, &amp;#34;MARGIN&amp;#34;, &amp;#34;FUTURES&amp;#34; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | min 7, max 30, default 7 | [default to 7]
 **recvWindow** | **int64** |  | 

### Return type

[**GetAccountSnapshotV1Resp**](GetAccountSnapshotV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountStatusV1

> GetAccountStatusV1Resp GetAccountStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Status (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAccountStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAccountStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountStatusV1`: GetAccountStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAccountStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAccountStatusV1Resp**](GetAccountStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetAssetDetailV1

> map[string]WalletGetAssetAssetDetailV1RespValue GetAssetAssetDetailV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Asset Detail (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetAssetDetailV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetAssetDetailV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetAssetDetailV1`: map[string]WalletGetAssetAssetDetailV1RespValue
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetAssetDetailV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetAssetDetailV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**map[string]WalletGetAssetAssetDetailV1RespValue**](WalletGetAssetAssetDetailV1RespValue.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetAssetDividendV1

> GetAssetAssetDividendV1Resp GetAssetAssetDividendV1(ctx).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Asset Dividend Record (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 20, max 500 (optional) (default to 20)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetAssetDividendV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetAssetDividendV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetAssetDividendV1`: GetAssetAssetDividendV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetAssetDividendV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetAssetDividendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 20, max 500 | [default to 20]
 **recvWindow** | **int64** |  | 

### Return type

[**GetAssetAssetDividendV1Resp**](GetAssetAssetDividendV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetCustodyTransferHistoryV1

> GetAssetCustodyTransferHistoryV1Resp GetAssetCustodyTransferHistoryV1(ctx).Email(email).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Type_(type_).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query User Delegation History(For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	email := "email_example" // string |  (default to "")
	startTime := int64(789) // int64 | 
	endTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string | Delegate/Undelegate (optional) (default to "")
	asset := "asset_example" // string |  (optional) (default to "")
	current := int32(56) // int32 | default 1 (optional) (default to 1)
	size := int32(56) // int32 | default 10, max 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetCustodyTransferHistoryV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Type_(type_).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetCustodyTransferHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetCustodyTransferHistoryV1`: GetAssetCustodyTransferHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetCustodyTransferHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetCustodyTransferHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **type_** | **string** | Delegate/Undelegate | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **current** | **int32** | default 1 | [default to 1]
 **size** | **int32** | default 10, max 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetAssetCustodyTransferHistoryV1Resp**](GetAssetCustodyTransferHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDribbletV1

> GetAssetDribbletV1Resp GetAssetDribbletV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

DustLog(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetDribbletV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetDribbletV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetDribbletV1`: GetAssetDribbletV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetDribbletV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDribbletV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetAssetDribbletV1Resp**](GetAssetDribbletV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetLedgerTransferCloudMiningQueryByPageV1

> GetAssetLedgerTransferCloudMiningQueryByPageV1Resp GetAssetLedgerTransferCloudMiningQueryByPageV1(ctx).StartTime(startTime).EndTime(endTime).TranId(tranId).ClientTranId(clientTranId).Asset(asset).Current(current).Size(size).Execute()

Get Cloud-Mining payment and refund history (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	startTime := int64(789) // int64 | inclusive, unit: ms
	endTime := int64(789) // int64 | exclusive, unit: ms
	tranId := int64(789) // int64 | The transaction id (optional)
	clientTranId := "clientTranId_example" // string | The unique flag (optional) (default to "")
	asset := "asset_example" // string | If it is blank, we will query all assets (optional) (default to "")
	current := int32(56) // int32 | current page, default 1, the min value is 1 (optional) (default to 1)
	size := int32(56) // int32 | page size, default 10, the max value is 100 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetLedgerTransferCloudMiningQueryByPageV1(context.Background()).StartTime(startTime).EndTime(endTime).TranId(tranId).ClientTranId(clientTranId).Asset(asset).Current(current).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetLedgerTransferCloudMiningQueryByPageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetLedgerTransferCloudMiningQueryByPageV1`: GetAssetLedgerTransferCloudMiningQueryByPageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetLedgerTransferCloudMiningQueryByPageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetLedgerTransferCloudMiningQueryByPageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | inclusive, unit: ms | 
 **endTime** | **int64** | exclusive, unit: ms | 
 **tranId** | **int64** | The transaction id | 
 **clientTranId** | **string** | The unique flag | [default to &quot;&quot;]
 **asset** | **string** | If it is blank, we will query all assets | [default to &quot;&quot;]
 **current** | **int32** | current page, default 1, the min value is 1 | [default to 1]
 **size** | **int32** | page size, default 10, the max value is 100 | [default to 10]

### Return type

[**GetAssetLedgerTransferCloudMiningQueryByPageV1Resp**](GetAssetLedgerTransferCloudMiningQueryByPageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTradeFeeV1

> []GetAssetTradeFeeV1RespItem GetAssetTradeFeeV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Trade Fee (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetTradeFeeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetTradeFeeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTradeFeeV1`: []GetAssetTradeFeeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetTradeFeeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTradeFeeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetAssetTradeFeeV1RespItem**](GetAssetTradeFeeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTransferV1

> GetAssetTransferV1Resp GetAssetTransferV1(ctx).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

Query User Universal Transfer History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	type_ := "type__example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int32(56) // int32 | Default 1 (optional) (default to 1)
	size := int32(56) // int32 | Default 10, Max 100 (optional) (default to 10)
	fromSymbol := "fromSymbol_example" // string |  (optional) (default to "")
	toSymbol := "toSymbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetTransferV1(context.Background()).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTransferV1`: GetAssetTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int32** | Default 1 | [default to 1]
 **size** | **int32** | Default 10, Max 100 | [default to 10]
 **fromSymbol** | **string** |  | [default to &quot;&quot;]
 **toSymbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetAssetTransferV1Resp**](GetAssetTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetWalletBalanceV1

> []GetAssetWalletBalanceV1RespItem GetAssetWalletBalanceV1(ctx).Timestamp(timestamp).QuoteAsset(quoteAsset).RecvWindow(recvWindow).Execute()

Query User Wallet Balance (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	quoteAsset := "quoteAsset_example" // string | `USDT`, `ETH`, `USDC`, `BNB`, etc. default `BTC` (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetAssetWalletBalanceV1(context.Background()).Timestamp(timestamp).QuoteAsset(quoteAsset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAssetWalletBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetWalletBalanceV1`: []GetAssetWalletBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAssetWalletBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetWalletBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **quoteAsset** | **string** | &#x60;USDT&#x60;, &#x60;ETH&#x60;, &#x60;USDC&#x60;, &#x60;BNB&#x60;, etc. default &#x60;BTC&#x60; | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetAssetWalletBalanceV1RespItem**](GetAssetWalletBalanceV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalConfigGetallV1

> []GetCapitalConfigGetallV1RespItem GetCapitalConfigGetallV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

All Coins' Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetCapitalConfigGetallV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetCapitalConfigGetallV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalConfigGetallV1`: []GetCapitalConfigGetallV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetCapitalConfigGetallV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalConfigGetallV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetCapitalConfigGetallV1RespItem**](GetCapitalConfigGetallV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalDepositAddressListV1

> []GetCapitalDepositAddressListV1RespItem GetCapitalDepositAddressListV1(ctx).Coin(coin).Timestamp(timestamp).Network(network).Execute()

Fetch deposit address list with network(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	coin := "coin_example" // string | `coin` refers to the parent network address format that the address is using (default to "")
	timestamp := int64(789) // int64 | 
	network := "network_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetCapitalDepositAddressListV1(context.Background()).Coin(coin).Timestamp(timestamp).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetCapitalDepositAddressListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalDepositAddressListV1`: []GetCapitalDepositAddressListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetCapitalDepositAddressListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalDepositAddressListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | &#x60;coin&#x60; refers to the parent network address format that the address is using | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]GetCapitalDepositAddressListV1RespItem**](GetCapitalDepositAddressListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalDepositAddressV1

> GetCapitalDepositAddressV1Resp GetCapitalDepositAddressV1(ctx).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Deposit Address(supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	coin := "coin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	network := "network_example" // string |  (optional) (default to "")
	amount := "amount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetCapitalDepositAddressV1(context.Background()).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetCapitalDepositAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalDepositAddressV1`: GetCapitalDepositAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetCapitalDepositAddressV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalDepositAddressV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetCapitalDepositAddressV1Resp**](GetCapitalDepositAddressV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalDepositHisrecV1

> []GetCapitalDepositHisrecV1RespItem GetCapitalDepositHisrecV1(ctx).Timestamp(timestamp).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).TxId(txId).Execute()

Deposit History (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	includeSource := true // bool | Default: `false`, return `sourceAddress`field when set to `true` (optional) (default to false)
	coin := "coin_example" // string |  (optional) (default to "")
	status := int32(56) // int32 | 0(0:pending, 6:credited but cannot withdraw, 7:Wrong Deposit, 8:Waiting User confirm, 1:success, 2:rejected) (optional)
	startTime := int64(789) // int64 | Default: 90 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	offset := int32(56) // int32 | Default:0 (optional)
	limit := int32(56) // int32 | Default:1000, Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)
	txId := "txId_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetCapitalDepositHisrecV1(context.Background()).Timestamp(timestamp).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).TxId(txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetCapitalDepositHisrecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalDepositHisrecV1`: []GetCapitalDepositHisrecV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetCapitalDepositHisrecV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalDepositHisrecV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **includeSource** | **bool** | Default: &#x60;false&#x60;, return &#x60;sourceAddress&#x60;field when set to &#x60;true&#x60; | [default to false]
 **coin** | **string** |  | [default to &quot;&quot;]
 **status** | **int32** | 0(0:pending, 6:credited but cannot withdraw, 7:Wrong Deposit, 8:Waiting User confirm, 1:success, 2:rejected) | 
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **offset** | **int32** | Default:0 | 
 **limit** | **int32** | Default:1000, Max:1000 | 
 **recvWindow** | **int64** |  | 
 **txId** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]GetCapitalDepositHisrecV1RespItem**](GetCapitalDepositHisrecV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalWithdrawAddressListV1

> []GetCapitalWithdrawAddressListV1RespItem GetCapitalWithdrawAddressListV1(ctx).Execute()

Fetch withdraw address list (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetCapitalWithdrawAddressListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetCapitalWithdrawAddressListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalWithdrawAddressListV1`: []GetCapitalWithdrawAddressListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetCapitalWithdrawAddressListV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalWithdrawAddressListV1Request struct via the builder pattern


### Return type

[**[]GetCapitalWithdrawAddressListV1RespItem**](GetCapitalWithdrawAddressListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapitalWithdrawHistoryV1

> []GetCapitalWithdrawHistoryV1RespItem GetCapitalWithdrawHistoryV1(ctx).Timestamp(timestamp).Coin(coin).WithdrawOrderId(withdrawOrderId).Status(status).Offset(offset).Limit(limit).IdList(idList).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	coin := "coin_example" // string |  (optional) (default to "")
	withdrawOrderId := "withdrawOrderId_example" // string |  (optional) (default to "")
	status := int32(56) // int32 | 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed) (optional)
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default: 1000, Max: 1000 (optional) (default to 1000)
	idList := "idList_example" // string | id list returned in the response of POST `/sapi/v1/capital/withdraw/apply`, separated by `,` (optional) (default to "")
	startTime := int64(789) // int64 | Default: 90 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetCapitalWithdrawHistoryV1(context.Background()).Timestamp(timestamp).Coin(coin).WithdrawOrderId(withdrawOrderId).Status(status).Offset(offset).Limit(limit).IdList(idList).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetCapitalWithdrawHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapitalWithdrawHistoryV1`: []GetCapitalWithdrawHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetCapitalWithdrawHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCapitalWithdrawHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **withdrawOrderId** | **string** |  | [default to &quot;&quot;]
 **status** | **int32** | 0(0:Email Sent, 2:Awaiting Approval 3:Rejected 4:Processing 6:Completed) | 
 **offset** | **int32** |  | 
 **limit** | **int32** | Default: 1000, Max: 1000 | [default to 1000]
 **idList** | **string** | id list returned in the response of POST &#x60;/sapi/v1/capital/withdraw/apply&#x60;, separated by &#x60;,&#x60; | [default to &quot;&quot;]
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetCapitalWithdrawHistoryV1RespItem**](GetCapitalWithdrawHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalentityDepositHistoryV1

> []GetLocalentityDepositHistoryV1RespItem GetLocalentityDepositHistoryV1(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()

Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	trId := "trId_example" // string | Comma(,) separated list of travel rule record Ids. (optional) (default to "")
	txId := "txId_example" // string | Comma(,) separated list of transaction Ids. (optional) (default to "")
	tranId := "tranId_example" // string | Comma(,) separated list of wallet tran Ids. (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	coin := "coin_example" // string |  (optional) (default to "")
	travelRuleStatus := int32(56) // int32 | 0:Completed,1:Pending,2:Failed (optional)
	pendingQuestionnaire := true // bool | true: Only return records that pending deposit questionnaire. false/not provided: return all records. (optional)
	startTime := int64(789) // int64 | Default: 90 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	offset := int32(56) // int32 | Default:0 (optional)
	limit := int32(56) // int32 | Default:1000, Max:1000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetLocalentityDepositHistoryV1(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetLocalentityDepositHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalentityDepositHistoryV1`: []GetLocalentityDepositHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetLocalentityDepositHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalentityDepositHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | [default to &quot;&quot;]
 **txId** | **string** | Comma(,) separated list of transaction Ids. | [default to &quot;&quot;]
 **tranId** | **string** | Comma(,) separated list of wallet tran Ids. | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **travelRuleStatus** | **int32** | 0:Completed,1:Pending,2:Failed | 
 **pendingQuestionnaire** | **bool** | true: Only return records that pending deposit questionnaire. false/not provided: return all records. | 
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **offset** | **int32** | Default:0 | 
 **limit** | **int32** | Default:1000, Max:1000 | 

### Return type

[**[]GetLocalentityDepositHistoryV1RespItem**](GetLocalentityDepositHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalentityVaspV1

> []GetLocalentityVaspV1RespItem GetLocalentityVaspV1(ctx).Execute()

Onboarded VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetLocalentityVaspV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetLocalentityVaspV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalentityVaspV1`: []GetLocalentityVaspV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetLocalentityVaspV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalentityVaspV1Request struct via the builder pattern


### Return type

[**[]GetLocalentityVaspV1RespItem**](GetLocalentityVaspV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalentityWithdrawHistoryV1

> []GetLocalentityWithdrawHistoryV1RespItem GetLocalentityWithdrawHistoryV1(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	trId := "trId_example" // string | Comma(,) separated list of travel rule record Ids. (optional) (default to "")
	txId := "txId_example" // string | Comma(,) separated list of transaction Ids. (optional) (default to "")
	withdrawOrderId := "withdrawOrderId_example" // string | Comma(,) separated list of withdrawID defined by the client (i.e. client&#39;s internal withdrawID). (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	coin := "coin_example" // string |  (optional) (default to "")
	travelRuleStatus := int32(56) // int32 | 0:Completed,1:Pending,2:Failed (optional)
	offset := int32(56) // int32 | Default: 0 (optional) (default to 0)
	limit := int32(56) // int32 | Default: 1000, Max: 1000 (optional) (default to 1000)
	startTime := int64(789) // int64 | Default: 90 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetLocalentityWithdrawHistoryV1(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetLocalentityWithdrawHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalentityWithdrawHistoryV1`: []GetLocalentityWithdrawHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetLocalentityWithdrawHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalentityWithdrawHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | [default to &quot;&quot;]
 **txId** | **string** | Comma(,) separated list of transaction Ids. | [default to &quot;&quot;]
 **withdrawOrderId** | **string** | Comma(,) separated list of withdrawID defined by the client (i.e. client&amp;#39;s internal withdrawID). | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **travelRuleStatus** | **int32** | 0:Completed,1:Pending,2:Failed | 
 **offset** | **int32** | Default: 0 | [default to 0]
 **limit** | **int32** | Default: 1000, Max: 1000 | [default to 1000]
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetLocalentityWithdrawHistoryV1RespItem**](GetLocalentityWithdrawHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalentityWithdrawHistoryV2

> []GetLocalentityWithdrawHistoryV2RespItem GetLocalentityWithdrawHistoryV2(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	trId := "trId_example" // string | Comma(,) separated list of travel rule record Ids. (optional) (default to "")
	txId := "txId_example" // string | Comma(,) separated list of transaction Ids. (optional) (default to "")
	withdrawOrderId := "withdrawOrderId_example" // string | Withdraw ID defined by the client (i.e. client&#39;s internal withdrawID). (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	coin := "coin_example" // string |  (optional) (default to "")
	travelRuleStatus := int32(56) // int32 | 0:Completed,1:Pending,2:Failed (optional)
	offset := int32(56) // int32 | Default: 0 (optional) (default to 0)
	limit := int32(56) // int32 | Default: 1000, Max: 1000 (optional) (default to 1000)
	startTime := int64(789) // int64 | Default: 90 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetLocalentityWithdrawHistoryV2(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetLocalentityWithdrawHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalentityWithdrawHistoryV2`: []GetLocalentityWithdrawHistoryV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetLocalentityWithdrawHistoryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalentityWithdrawHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | [default to &quot;&quot;]
 **txId** | **string** | Comma(,) separated list of transaction Ids. | [default to &quot;&quot;]
 **withdrawOrderId** | **string** | Withdraw ID defined by the client (i.e. client&amp;#39;s internal withdrawID). | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **travelRuleStatus** | **int32** | 0:Completed,1:Pending,2:Failed | 
 **offset** | **int32** | Default: 0 | [default to 0]
 **limit** | **int32** | Default: 1000, Max: 1000 | [default to 1000]
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetLocalentityWithdrawHistoryV2RespItem**](GetLocalentityWithdrawHistoryV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpotDelistScheduleV1

> []GetSpotDelistScheduleV1RespItem GetSpotDelistScheduleV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Spot Delist Schedule (MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetSpotDelistScheduleV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetSpotDelistScheduleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpotDelistScheduleV1`: []GetSpotDelistScheduleV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetSpotDelistScheduleV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpotDelistScheduleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetSpotDelistScheduleV1RespItem**](GetSpotDelistScheduleV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpotOpenSymbolListV1

> []GetSpotOpenSymbolListV1RespItem GetSpotOpenSymbolListV1(ctx).Execute()

Get Open Symbol List (MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetSpotOpenSymbolListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetSpotOpenSymbolListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpotOpenSymbolListV1`: []GetSpotOpenSymbolListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetSpotOpenSymbolListV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpotOpenSymbolListV1Request struct via the builder pattern


### Return type

[**[]GetSpotOpenSymbolListV1RespItem**](GetSpotOpenSymbolListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemStatusV1

> GetSystemStatusV1Resp GetSystemStatusV1(ctx).Execute()

System Status (System)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetSystemStatusV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetSystemStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemStatusV1`: GetSystemStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetSystemStatusV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemStatusV1Request struct via the builder pattern


### Return type

[**GetSystemStatusV1Resp**](GetSystemStatusV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalentityBrokerDepositProvideInfoV1

> UpdateLocalentityBrokerDepositProvideInfoV1Resp UpdateLocalentityBrokerDepositProvideInfoV1(ctx).BeneficiaryPii(beneficiaryPii).DepositId(depositId).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).Address(address).AddressTag(addressTag).Amount(amount).Coin(coin).Network(network).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	beneficiaryPii := "beneficiaryPii_example" // string |  (default to "")
	depositId := "depositId_example" // string |  (default to "")
	questionnaire := "questionnaire_example" // string |  (default to "")
	signature := "signature_example" // string |  (default to "")
	subAccountId := "subAccountId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	address := "address_example" // string |  (optional) (default to "")
	addressTag := "addressTag_example" // string |  (optional) (default to "")
	amount := "amount_example" // string |  (optional) (default to "")
	coin := "coin_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.UpdateLocalentityBrokerDepositProvideInfoV1(context.Background()).BeneficiaryPii(beneficiaryPii).DepositId(depositId).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).Address(address).AddressTag(addressTag).Amount(amount).Coin(coin).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.UpdateLocalentityBrokerDepositProvideInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLocalentityBrokerDepositProvideInfoV1`: UpdateLocalentityBrokerDepositProvideInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.UpdateLocalentityBrokerDepositProvideInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalentityBrokerDepositProvideInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beneficiaryPii** | **string** |  | [default to &quot;&quot;]
 **depositId** | **string** |  | [default to &quot;&quot;]
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **signature** | **string** |  | [default to &quot;&quot;]
 **subAccountId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **address** | **string** |  | [default to &quot;&quot;]
 **addressTag** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]

### Return type

[**UpdateLocalentityBrokerDepositProvideInfoV1Resp**](UpdateLocalentityBrokerDepositProvideInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalentityDepositProvideInfoV1

> UpdateLocalentityDepositProvideInfoV1Resp UpdateLocalentityDepositProvideInfoV1(ctx).Questionnaire(questionnaire).Timestamp(timestamp).TranId(tranId).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	questionnaire := "questionnaire_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	tranId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.UpdateLocalentityDepositProvideInfoV1(context.Background()).Questionnaire(questionnaire).Timestamp(timestamp).TranId(tranId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.UpdateLocalentityDepositProvideInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLocalentityDepositProvideInfoV1`: UpdateLocalentityDepositProvideInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.UpdateLocalentityDepositProvideInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalentityDepositProvideInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **tranId** | **int64** |  | 

### Return type

[**UpdateLocalentityDepositProvideInfoV1Resp**](UpdateLocalentityDepositProvideInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

