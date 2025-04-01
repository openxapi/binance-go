# \V1API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletCreateAccountDisableFastWithdrawSwitchV1**](V1API.md#WalletCreateAccountDisableFastWithdrawSwitchV1) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**WalletCreateAccountEnableFastWithdrawSwitchV1**](V1API.md#WalletCreateAccountEnableFastWithdrawSwitchV1) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**WalletCreateAssetDustBtcV1**](V1API.md#WalletCreateAssetDustBtcV1) | **Post** /sapi/v1/asset/dust-btc | Get Assets That Can Be Converted Into BNB (USER_DATA)
[**WalletCreateAssetDustV1**](V1API.md#WalletCreateAssetDustV1) | **Post** /sapi/v1/asset/dust | Dust Transfer (USER_DATA)
[**WalletCreateAssetGetFundingAssetV1**](V1API.md#WalletCreateAssetGetFundingAssetV1) | **Post** /sapi/v1/asset/get-funding-asset | Funding Wallet (USER_DATA)
[**WalletCreateAssetTransferV1**](V1API.md#WalletCreateAssetTransferV1) | **Post** /sapi/v1/asset/transfer | User Universal Transfer (USER_DATA)
[**WalletCreateBnbBurnV1**](V1API.md#WalletCreateBnbBurnV1) | **Post** /sapi/v1/bnbBurn | Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
[**WalletCreateCapitalDepositCreditApplyV1**](V1API.md#WalletCreateCapitalDepositCreditApplyV1) | **Post** /sapi/v1/capital/deposit/credit-apply | One click arrival deposit apply (for expired address deposit) (USER_DATA)
[**WalletCreateCapitalWithdrawApplyV1**](V1API.md#WalletCreateCapitalWithdrawApplyV1) | **Post** /sapi/v1/capital/withdraw/apply | Withdraw(USER_DATA)
[**WalletCreateLocalentityBrokerWithdrawApplyV1**](V1API.md#WalletCreateLocalentityBrokerWithdrawApplyV1) | **Post** /sapi/v1/localentity/broker/withdraw/apply | Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)
[**WalletCreateLocalentityWithdrawApplyV1**](V1API.md#WalletCreateLocalentityWithdrawApplyV1) | **Post** /sapi/v1/localentity/withdraw/apply | Withdraw (for local entities that require travel rule) (USER_DATA)
[**WalletGetAccountApiRestrictionsV1**](V1API.md#WalletGetAccountApiRestrictionsV1) | **Get** /sapi/v1/account/apiRestrictions | Get API Key Permission (USER_DATA)
[**WalletGetAccountApiTradingStatusV1**](V1API.md#WalletGetAccountApiTradingStatusV1) | **Get** /sapi/v1/account/apiTradingStatus | Account API Trading Status (USER_DATA)
[**WalletGetAccountInfoV1**](V1API.md#WalletGetAccountInfoV1) | **Get** /sapi/v1/account/info | Account info (USER_DATA)
[**WalletGetAccountSnapshotV1**](V1API.md#WalletGetAccountSnapshotV1) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**WalletGetAccountStatusV1**](V1API.md#WalletGetAccountStatusV1) | **Get** /sapi/v1/account/status | Account Status (USER_DATA)
[**WalletGetAssetAssetDetailV1**](V1API.md#WalletGetAssetAssetDetailV1) | **Get** /sapi/v1/asset/assetDetail | Asset Detail (USER_DATA)
[**WalletGetAssetAssetDividendV1**](V1API.md#WalletGetAssetAssetDividendV1) | **Get** /sapi/v1/asset/assetDividend | Asset Dividend Record (USER_DATA)
[**WalletGetAssetCustodyTransferHistoryV1**](V1API.md#WalletGetAssetCustodyTransferHistoryV1) | **Get** /sapi/v1/asset/custody/transfer-history | Query User Delegation History(For Master Account)(USER_DATA)
[**WalletGetAssetDribbletV1**](V1API.md#WalletGetAssetDribbletV1) | **Get** /sapi/v1/asset/dribblet | DustLog(USER_DATA)
[**WalletGetAssetLedgerTransferCloudMiningQueryByPageV1**](V1API.md#WalletGetAssetLedgerTransferCloudMiningQueryByPageV1) | **Get** /sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage | Get Cloud-Mining payment and refund history (USER_DATA)
[**WalletGetAssetTradeFeeV1**](V1API.md#WalletGetAssetTradeFeeV1) | **Get** /sapi/v1/asset/tradeFee | Trade Fee (USER_DATA)
[**WalletGetAssetTransferV1**](V1API.md#WalletGetAssetTransferV1) | **Get** /sapi/v1/asset/transfer | Query User Universal Transfer History(USER_DATA)
[**WalletGetAssetWalletBalanceV1**](V1API.md#WalletGetAssetWalletBalanceV1) | **Get** /sapi/v1/asset/wallet/balance | Query User Wallet Balance (USER_DATA)
[**WalletGetCapitalConfigGetallV1**](V1API.md#WalletGetCapitalConfigGetallV1) | **Get** /sapi/v1/capital/config/getall | All Coins&#39; Information (USER_DATA)
[**WalletGetCapitalDepositAddressListV1**](V1API.md#WalletGetCapitalDepositAddressListV1) | **Get** /sapi/v1/capital/deposit/address/list | Fetch deposit address list with network(USER_DATA)
[**WalletGetCapitalDepositAddressV1**](V1API.md#WalletGetCapitalDepositAddressV1) | **Get** /sapi/v1/capital/deposit/address | Deposit Address(supporting network) (USER_DATA)
[**WalletGetCapitalDepositHisrecV1**](V1API.md#WalletGetCapitalDepositHisrecV1) | **Get** /sapi/v1/capital/deposit/hisrec | Deposit History (supporting network) (USER_DATA)
[**WalletGetCapitalWithdrawAddressListV1**](V1API.md#WalletGetCapitalWithdrawAddressListV1) | **Get** /sapi/v1/capital/withdraw/address/list | Fetch withdraw address list (USER_DATA)
[**WalletGetCapitalWithdrawHistoryV1**](V1API.md#WalletGetCapitalWithdrawHistoryV1) | **Get** /sapi/v1/capital/withdraw/history | Withdraw History (supporting network) (USER_DATA)
[**WalletGetLocalentityDepositHistoryV1**](V1API.md#WalletGetLocalentityDepositHistoryV1) | **Get** /sapi/v1/localentity/deposit/history | Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)
[**WalletGetLocalentityVaspV1**](V1API.md#WalletGetLocalentityVaspV1) | **Get** /sapi/v1/localentity/vasp | Onboarded VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WalletGetLocalentityWithdrawHistoryV1**](V1API.md#WalletGetLocalentityWithdrawHistoryV1) | **Get** /sapi/v1/localentity/withdraw/history | Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WalletGetSpotDelistScheduleV1**](V1API.md#WalletGetSpotDelistScheduleV1) | **Get** /sapi/v1/spot/delist-schedule | Get Spot Delist Schedule (MARKET_DATA)
[**WalletGetSpotOpenSymbolListV1**](V1API.md#WalletGetSpotOpenSymbolListV1) | **Get** /sapi/v1/spot/open-symbol-list | Get Open Symbol List (MARKET_DATA)
[**WalletGetSystemStatusV1**](V1API.md#WalletGetSystemStatusV1) | **Get** /sapi/v1/system/status | System Status (System)
[**WalletUpdateLocalentityBrokerDepositProvideInfoV1**](V1API.md#WalletUpdateLocalentityBrokerDepositProvideInfoV1) | **Put** /sapi/v1/localentity/broker/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)
[**WalletUpdateLocalentityDepositProvideInfoV1**](V1API.md#WalletUpdateLocalentityDepositProvideInfoV1) | **Put** /sapi/v1/localentity/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



## WalletCreateAccountDisableFastWithdrawSwitchV1

> map[string]interface{} WalletCreateAccountDisableFastWithdrawSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Disable Fast Withdraw Switch (USER_DATA)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletCreateAccountDisableFastWithdrawSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateAccountDisableFastWithdrawSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAccountDisableFastWithdrawSwitchV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateAccountDisableFastWithdrawSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAccountDisableFastWithdrawSwitchV1Request struct via the builder pattern


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


## WalletCreateAccountEnableFastWithdrawSwitchV1

> map[string]interface{} WalletCreateAccountEnableFastWithdrawSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Fast Withdraw Switch (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletCreateAccountEnableFastWithdrawSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateAccountEnableFastWithdrawSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAccountEnableFastWithdrawSwitchV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateAccountEnableFastWithdrawSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAccountEnableFastWithdrawSwitchV1Request struct via the builder pattern


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


## WalletCreateAssetDustBtcV1

> WalletCreateAssetDustBtcV1Resp WalletCreateAssetDustBtcV1(ctx).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()

Get Assets That Can Be Converted Into BNB (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	accountType := "accountType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletCreateAssetDustBtcV1(context.Background()).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateAssetDustBtcV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetDustBtcV1`: WalletCreateAssetDustBtcV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateAssetDustBtcV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAssetDustBtcV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **accountType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**WalletCreateAssetDustBtcV1Resp**](WalletCreateAssetDustBtcV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateAssetDustV1

> WalletCreateAssetDustV1Resp WalletCreateAssetDustV1(ctx).Asset(asset).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()

Dust Transfer (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	asset := []string{"Inner_example"} // []string | 
	timestamp := int64(789) // int64 | 
	accountType := "accountType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletCreateAssetDustV1(context.Background()).Asset(asset).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateAssetDustV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetDustV1`: WalletCreateAssetDustV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateAssetDustV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAssetDustV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **[]string** |  | 
 **timestamp** | **int64** |  | 
 **accountType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**WalletCreateAssetDustV1Resp**](WalletCreateAssetDustV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateAssetGetFundingAssetV1

> []WalletCreateAssetGetFundingAssetV1RespItem WalletCreateAssetGetFundingAssetV1(ctx).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

Funding Wallet (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	needBtcValuation := "needBtcValuation_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletCreateAssetGetFundingAssetV1(context.Background()).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateAssetGetFundingAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetGetFundingAssetV1`: []WalletCreateAssetGetFundingAssetV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateAssetGetFundingAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAssetGetFundingAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **needBtcValuation** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletCreateAssetGetFundingAssetV1RespItem**](WalletCreateAssetGetFundingAssetV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateAssetTransferV1

> WalletCreateAssetTransferV1Resp WalletCreateAssetTransferV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).Type_(type_).FromSymbol(fromSymbol).RecvWindow(recvWindow).ToSymbol(toSymbol).Execute()

User Universal Transfer (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletCreateAssetTransferV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).Type_(type_).FromSymbol(fromSymbol).RecvWindow(recvWindow).ToSymbol(toSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateAssetTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetTransferV1`: WalletCreateAssetTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateAssetTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAssetTransferV1Request struct via the builder pattern


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

[**WalletCreateAssetTransferV1Resp**](WalletCreateAssetTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateBnbBurnV1

> WalletCreateBnbBurnV1Resp WalletCreateBnbBurnV1(ctx).Timestamp(timestamp).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).SpotBNBBurn(spotBNBBurn).Execute()

Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	interestBNBBurn := "interestBNBBurn_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	spotBNBBurn := "spotBNBBurn_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletCreateBnbBurnV1(context.Background()).Timestamp(timestamp).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).SpotBNBBurn(spotBNBBurn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateBnbBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateBnbBurnV1`: WalletCreateBnbBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateBnbBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateBnbBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **interestBNBBurn** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **spotBNBBurn** | **string** |  | [default to &quot;&quot;]

### Return type

[**WalletCreateBnbBurnV1Resp**](WalletCreateBnbBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateCapitalDepositCreditApplyV1

> WalletCreateCapitalDepositCreditApplyV1Resp WalletCreateCapitalDepositCreditApplyV1(ctx).DepositId(depositId).SubAccountId(subAccountId).SubUserId(subUserId).TxId(txId).Execute()

One click arrival deposit apply (for expired address deposit) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	depositId := int64(789) // int64 |  (optional)
	subAccountId := int64(789) // int64 |  (optional)
	subUserId := int64(789) // int64 |  (optional)
	txId := "txId_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletCreateCapitalDepositCreditApplyV1(context.Background()).DepositId(depositId).SubAccountId(subAccountId).SubUserId(subUserId).TxId(txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateCapitalDepositCreditApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateCapitalDepositCreditApplyV1`: WalletCreateCapitalDepositCreditApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateCapitalDepositCreditApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateCapitalDepositCreditApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositId** | **int64** |  | 
 **subAccountId** | **int64** |  | 
 **subUserId** | **int64** |  | 
 **txId** | **string** |  | [default to &quot;&quot;]

### Return type

[**WalletCreateCapitalDepositCreditApplyV1Resp**](WalletCreateCapitalDepositCreditApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateCapitalWithdrawApplyV1

> WalletCreateCapitalWithdrawApplyV1Resp WalletCreateCapitalWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()

Withdraw(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletCreateCapitalWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateCapitalWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateCapitalWithdrawApplyV1`: WalletCreateCapitalWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateCapitalWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateCapitalWithdrawApplyV1Request struct via the builder pattern


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

[**WalletCreateCapitalWithdrawApplyV1Resp**](WalletCreateCapitalWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateLocalentityBrokerWithdrawApplyV1

> WalletCreateLocalentityBrokerWithdrawApplyV1Resp WalletCreateLocalentityBrokerWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).OriginatorPii(originatorPii).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).WithdrawOrderId(withdrawOrderId).AddressName(addressName).AddressTag(addressTag).Network(network).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()

Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	address := "address_example" // string |  (default to "")
	amount := map[string]interface{}{ ... } // map[string]interface{} | 
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
	resp, r, err := apiClient.V1API.WalletCreateLocalentityBrokerWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).OriginatorPii(originatorPii).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).WithdrawOrderId(withdrawOrderId).AddressName(addressName).AddressTag(addressTag).Network(network).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateLocalentityBrokerWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateLocalentityBrokerWithdrawApplyV1`: WalletCreateLocalentityBrokerWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateLocalentityBrokerWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateLocalentityBrokerWithdrawApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | [default to &quot;&quot;]
 **amount** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
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

[**WalletCreateLocalentityBrokerWithdrawApplyV1Resp**](WalletCreateLocalentityBrokerWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateLocalentityWithdrawApplyV1

> WalletCreateLocalentityWithdrawApplyV1Resp WalletCreateLocalentityWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).Questionnaire(questionnaire).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()

Withdraw (for local entities that require travel rule) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletCreateLocalentityWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).Questionnaire(questionnaire).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletCreateLocalentityWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateLocalentityWithdrawApplyV1`: WalletCreateLocalentityWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletCreateLocalentityWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateLocalentityWithdrawApplyV1Request struct via the builder pattern


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

[**WalletCreateLocalentityWithdrawApplyV1Resp**](WalletCreateLocalentityWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountApiRestrictionsV1

> WalletGetAccountApiRestrictionsV1Resp WalletGetAccountApiRestrictionsV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get API Key Permission (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAccountApiRestrictionsV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAccountApiRestrictionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountApiRestrictionsV1`: WalletGetAccountApiRestrictionsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAccountApiRestrictionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountApiRestrictionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountApiRestrictionsV1Resp**](WalletGetAccountApiRestrictionsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountApiTradingStatusV1

> WalletGetAccountApiTradingStatusV1Resp WalletGetAccountApiTradingStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account API Trading Status (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAccountApiTradingStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAccountApiTradingStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountApiTradingStatusV1`: WalletGetAccountApiTradingStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAccountApiTradingStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountApiTradingStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountApiTradingStatusV1Resp**](WalletGetAccountApiTradingStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountInfoV1

> WalletGetAccountInfoV1Resp WalletGetAccountInfoV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account info (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAccountInfoV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountInfoV1`: WalletGetAccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountInfoV1Resp**](WalletGetAccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountSnapshotV1

> WalletGetAccountSnapshotV1Resp WalletGetAccountSnapshotV1(ctx).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Daily Account Snapshot (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetAccountSnapshotV1(context.Background()).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAccountSnapshotV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountSnapshotV1`: WalletGetAccountSnapshotV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAccountSnapshotV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountSnapshotV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &amp;#34;SPOT&amp;#34;, &amp;#34;MARGIN&amp;#34;, &amp;#34;FUTURES&amp;#34; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | min 7, max 30, default 7 | [default to 7]
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountSnapshotV1Resp**](WalletGetAccountSnapshotV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountStatusV1

> WalletGetAccountStatusV1Resp WalletGetAccountStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Status (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAccountStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAccountStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountStatusV1`: WalletGetAccountStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAccountStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountStatusV1Resp**](WalletGetAccountStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAssetAssetDetailV1

> map[string]WalletGetAssetAssetDetailV1RespValue WalletGetAssetAssetDetailV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Asset Detail (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAssetAssetDetailV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetAssetDetailV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetAssetDetailV1`: map[string]WalletGetAssetAssetDetailV1RespValue
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetAssetDetailV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetAssetDetailV1Request struct via the builder pattern


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


## WalletGetAssetAssetDividendV1

> WalletGetAssetAssetDividendV1Resp WalletGetAssetAssetDividendV1(ctx).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Asset Dividend Record (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetAssetAssetDividendV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetAssetDividendV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetAssetDividendV1`: WalletGetAssetAssetDividendV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetAssetDividendV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetAssetDividendV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 20, max 500 | [default to 20]
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAssetAssetDividendV1Resp**](WalletGetAssetAssetDividendV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAssetCustodyTransferHistoryV1

> WalletGetAssetCustodyTransferHistoryV1Resp WalletGetAssetCustodyTransferHistoryV1(ctx).Email(email).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Type_(type_).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query User Delegation History(For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetAssetCustodyTransferHistoryV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Type_(type_).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetCustodyTransferHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetCustodyTransferHistoryV1`: WalletGetAssetCustodyTransferHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetCustodyTransferHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetCustodyTransferHistoryV1Request struct via the builder pattern


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

[**WalletGetAssetCustodyTransferHistoryV1Resp**](WalletGetAssetCustodyTransferHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAssetDribbletV1

> WalletGetAssetDribbletV1Resp WalletGetAssetDribbletV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

DustLog(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAssetDribbletV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetDribbletV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetDribbletV1`: WalletGetAssetDribbletV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetDribbletV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetDribbletV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAssetDribbletV1Resp**](WalletGetAssetDribbletV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAssetLedgerTransferCloudMiningQueryByPageV1

> WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp WalletGetAssetLedgerTransferCloudMiningQueryByPageV1(ctx).StartTime(startTime).EndTime(endTime).TranId(tranId).ClientTranId(clientTranId).Asset(asset).Current(current).Size(size).Execute()

Get Cloud-Mining payment and refund history (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetAssetLedgerTransferCloudMiningQueryByPageV1(context.Background()).StartTime(startTime).EndTime(endTime).TranId(tranId).ClientTranId(clientTranId).Asset(asset).Current(current).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetLedgerTransferCloudMiningQueryByPageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetLedgerTransferCloudMiningQueryByPageV1`: WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetLedgerTransferCloudMiningQueryByPageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Request struct via the builder pattern


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

[**WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp**](WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAssetTradeFeeV1

> []WalletGetAssetTradeFeeV1RespItem WalletGetAssetTradeFeeV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Trade Fee (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAssetTradeFeeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetTradeFeeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetTradeFeeV1`: []WalletGetAssetTradeFeeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetTradeFeeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetTradeFeeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletGetAssetTradeFeeV1RespItem**](WalletGetAssetTradeFeeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAssetTransferV1

> WalletGetAssetTransferV1Resp WalletGetAssetTransferV1(ctx).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()

Query User Universal Transfer History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetAssetTransferV1(context.Background()).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetTransferV1`: WalletGetAssetTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetTransferV1Request struct via the builder pattern


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

[**WalletGetAssetTransferV1Resp**](WalletGetAssetTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAssetWalletBalanceV1

> []WalletGetAssetWalletBalanceV1RespItem WalletGetAssetWalletBalanceV1(ctx).Timestamp(timestamp).QuoteAsset(quoteAsset).RecvWindow(recvWindow).Execute()

Query User Wallet Balance (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	quoteAsset := "quoteAsset_example" // string | `USDT`, `ETH`, `USDC`, `BNB`, etc. default `BTC` (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetAssetWalletBalanceV1(context.Background()).Timestamp(timestamp).QuoteAsset(quoteAsset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetAssetWalletBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetWalletBalanceV1`: []WalletGetAssetWalletBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetAssetWalletBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetWalletBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **quoteAsset** | **string** | &#x60;USDT&#x60;, &#x60;ETH&#x60;, &#x60;USDC&#x60;, &#x60;BNB&#x60;, etc. default &#x60;BTC&#x60; | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletGetAssetWalletBalanceV1RespItem**](WalletGetAssetWalletBalanceV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetCapitalConfigGetallV1

> []WalletGetCapitalConfigGetallV1RespItem WalletGetCapitalConfigGetallV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

All Coins' Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetCapitalConfigGetallV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetCapitalConfigGetallV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetCapitalConfigGetallV1`: []WalletGetCapitalConfigGetallV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetCapitalConfigGetallV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetCapitalConfigGetallV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletGetCapitalConfigGetallV1RespItem**](WalletGetCapitalConfigGetallV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetCapitalDepositAddressListV1

> []WalletGetCapitalDepositAddressListV1RespItem WalletGetCapitalDepositAddressListV1(ctx).Coin(coin).Timestamp(timestamp).Network(network).Execute()

Fetch deposit address list with network(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	coin := "coin_example" // string | `coin` refers to the parent network address format that the address is using (default to "")
	timestamp := int64(789) // int64 | 
	network := "network_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetCapitalDepositAddressListV1(context.Background()).Coin(coin).Timestamp(timestamp).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetCapitalDepositAddressListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetCapitalDepositAddressListV1`: []WalletGetCapitalDepositAddressListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetCapitalDepositAddressListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetCapitalDepositAddressListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** | &#x60;coin&#x60; refers to the parent network address format that the address is using | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]WalletGetCapitalDepositAddressListV1RespItem**](WalletGetCapitalDepositAddressListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetCapitalDepositAddressV1

> WalletGetCapitalDepositAddressV1Resp WalletGetCapitalDepositAddressV1(ctx).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()

Deposit Address(supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	coin := "coin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	network := "network_example" // string |  (optional) (default to "")
	amount := "amount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetCapitalDepositAddressV1(context.Background()).Coin(coin).Timestamp(timestamp).Network(network).Amount(amount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetCapitalDepositAddressV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetCapitalDepositAddressV1`: WalletGetCapitalDepositAddressV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetCapitalDepositAddressV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetCapitalDepositAddressV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **network** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetCapitalDepositAddressV1Resp**](WalletGetCapitalDepositAddressV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetCapitalDepositHisrecV1

> []WalletGetCapitalDepositHisrecV1RespItem WalletGetCapitalDepositHisrecV1(ctx).Timestamp(timestamp).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).TxId(txId).Execute()

Deposit History (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetCapitalDepositHisrecV1(context.Background()).Timestamp(timestamp).IncludeSource(includeSource).Coin(coin).Status(status).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).RecvWindow(recvWindow).TxId(txId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetCapitalDepositHisrecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetCapitalDepositHisrecV1`: []WalletGetCapitalDepositHisrecV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetCapitalDepositHisrecV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetCapitalDepositHisrecV1Request struct via the builder pattern


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

[**[]WalletGetCapitalDepositHisrecV1RespItem**](WalletGetCapitalDepositHisrecV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetCapitalWithdrawAddressListV1

> []WalletGetCapitalWithdrawAddressListV1RespItem WalletGetCapitalWithdrawAddressListV1(ctx).Execute()

Fetch withdraw address list (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetCapitalWithdrawAddressListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetCapitalWithdrawAddressListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetCapitalWithdrawAddressListV1`: []WalletGetCapitalWithdrawAddressListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetCapitalWithdrawAddressListV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetCapitalWithdrawAddressListV1Request struct via the builder pattern


### Return type

[**[]WalletGetCapitalWithdrawAddressListV1RespItem**](WalletGetCapitalWithdrawAddressListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetCapitalWithdrawHistoryV1

> []WalletGetCapitalWithdrawHistoryV1RespItem WalletGetCapitalWithdrawHistoryV1(ctx).Timestamp(timestamp).Coin(coin).WithdrawOrderId(withdrawOrderId).Status(status).Offset(offset).Limit(limit).IdList(idList).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetCapitalWithdrawHistoryV1(context.Background()).Timestamp(timestamp).Coin(coin).WithdrawOrderId(withdrawOrderId).Status(status).Offset(offset).Limit(limit).IdList(idList).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetCapitalWithdrawHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetCapitalWithdrawHistoryV1`: []WalletGetCapitalWithdrawHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetCapitalWithdrawHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetCapitalWithdrawHistoryV1Request struct via the builder pattern


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

[**[]WalletGetCapitalWithdrawHistoryV1RespItem**](WalletGetCapitalWithdrawHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetLocalentityDepositHistoryV1

> []WalletGetLocalentityDepositHistoryV1RespItem WalletGetLocalentityDepositHistoryV1(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()

Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetLocalentityDepositHistoryV1(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetLocalentityDepositHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityDepositHistoryV1`: []WalletGetLocalentityDepositHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetLocalentityDepositHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetLocalentityDepositHistoryV1Request struct via the builder pattern


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

[**[]WalletGetLocalentityDepositHistoryV1RespItem**](WalletGetLocalentityDepositHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetLocalentityVaspV1

> []WalletGetLocalentityVaspV1RespItem WalletGetLocalentityVaspV1(ctx).Execute()

Onboarded VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetLocalentityVaspV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetLocalentityVaspV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityVaspV1`: []WalletGetLocalentityVaspV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetLocalentityVaspV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetLocalentityVaspV1Request struct via the builder pattern


### Return type

[**[]WalletGetLocalentityVaspV1RespItem**](WalletGetLocalentityVaspV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetLocalentityWithdrawHistoryV1

> []WalletGetLocalentityWithdrawHistoryV1RespItem WalletGetLocalentityWithdrawHistoryV1(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	resp, r, err := apiClient.V1API.WalletGetLocalentityWithdrawHistoryV1(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetLocalentityWithdrawHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityWithdrawHistoryV1`: []WalletGetLocalentityWithdrawHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetLocalentityWithdrawHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetLocalentityWithdrawHistoryV1Request struct via the builder pattern


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

[**[]WalletGetLocalentityWithdrawHistoryV1RespItem**](WalletGetLocalentityWithdrawHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetSpotDelistScheduleV1

> []WalletGetSpotDelistScheduleV1RespItem WalletGetSpotDelistScheduleV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Spot Delist Schedule (MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetSpotDelistScheduleV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetSpotDelistScheduleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetSpotDelistScheduleV1`: []WalletGetSpotDelistScheduleV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetSpotDelistScheduleV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetSpotDelistScheduleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletGetSpotDelistScheduleV1RespItem**](WalletGetSpotDelistScheduleV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetSpotOpenSymbolListV1

> []WalletGetSpotOpenSymbolListV1RespItem WalletGetSpotOpenSymbolListV1(ctx).Execute()

Get Open Symbol List (MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetSpotOpenSymbolListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetSpotOpenSymbolListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetSpotOpenSymbolListV1`: []WalletGetSpotOpenSymbolListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetSpotOpenSymbolListV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetSpotOpenSymbolListV1Request struct via the builder pattern


### Return type

[**[]WalletGetSpotOpenSymbolListV1RespItem**](WalletGetSpotOpenSymbolListV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetSystemStatusV1

> WalletGetSystemStatusV1Resp WalletGetSystemStatusV1(ctx).Execute()

System Status (System)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletGetSystemStatusV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletGetSystemStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetSystemStatusV1`: WalletGetSystemStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletGetSystemStatusV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetSystemStatusV1Request struct via the builder pattern


### Return type

[**WalletGetSystemStatusV1Resp**](WalletGetSystemStatusV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletUpdateLocalentityBrokerDepositProvideInfoV1

> WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp WalletUpdateLocalentityBrokerDepositProvideInfoV1(ctx).BeneficiaryPii(beneficiaryPii).DepositId(depositId).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).Address(address).AddressTag(addressTag).Amount(amount).Coin(coin).Network(network).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
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
	amount := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
	coin := "coin_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletUpdateLocalentityBrokerDepositProvideInfoV1(context.Background()).BeneficiaryPii(beneficiaryPii).DepositId(depositId).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).Address(address).AddressTag(addressTag).Amount(amount).Coin(coin).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletUpdateLocalentityBrokerDepositProvideInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletUpdateLocalentityBrokerDepositProvideInfoV1`: WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletUpdateLocalentityBrokerDepositProvideInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletUpdateLocalentityBrokerDepositProvideInfoV1Request struct via the builder pattern


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
 **amount** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]

### Return type

[**WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp**](WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletUpdateLocalentityDepositProvideInfoV1

> WalletUpdateLocalentityDepositProvideInfoV1Resp WalletUpdateLocalentityDepositProvideInfoV1(ctx).Questionnaire(questionnaire).Timestamp(timestamp).TranId(tranId).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {
	questionnaire := "questionnaire_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	tranId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.WalletUpdateLocalentityDepositProvideInfoV1(context.Background()).Questionnaire(questionnaire).Timestamp(timestamp).TranId(tranId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.WalletUpdateLocalentityDepositProvideInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletUpdateLocalentityDepositProvideInfoV1`: WalletUpdateLocalentityDepositProvideInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.WalletUpdateLocalentityDepositProvideInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletUpdateLocalentityDepositProvideInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **tranId** | **int64** |  | 

### Return type

[**WalletUpdateLocalentityDepositProvideInfoV1Resp**](WalletUpdateLocalentityDepositProvideInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

