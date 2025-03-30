# \V1API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCreateMarginApiKeyV1**](V1API.md#MarginCreateMarginApiKeyV1) | **Post** /sapi/v1/margin/apiKey | Create Special Key(Low-Latency Trading)(TRADE)
[**MarginCreateMarginBorrowRepayV1**](V1API.md#MarginCreateMarginBorrowRepayV1) | **Post** /sapi/v1/margin/borrow-repay | Margin account borrow/repay(MARGIN)
[**MarginCreateMarginExchangeSmallLiabilityV1**](V1API.md#MarginCreateMarginExchangeSmallLiabilityV1) | **Post** /sapi/v1/margin/exchange-small-liability | Small Liability Exchange (MARGIN)
[**MarginCreateMarginIsolatedAccountV1**](V1API.md#MarginCreateMarginIsolatedAccountV1) | **Post** /sapi/v1/margin/isolated/account | Enable Isolated Margin Account (TRADE)
[**MarginCreateMarginListenKeyV1**](V1API.md#MarginCreateMarginListenKeyV1) | **Post** /sapi/v1/margin/listen-key | Start User Data Stream (USER_STREAM)
[**MarginCreateMarginManualLiquidationV1**](V1API.md#MarginCreateMarginManualLiquidationV1) | **Post** /sapi/v1/margin/manual-liquidation | Margin Manual Liquidation(MARGIN)
[**MarginCreateMarginMaxLeverageV1**](V1API.md#MarginCreateMarginMaxLeverageV1) | **Post** /sapi/v1/margin/max-leverage | Adjust cross margin max leverage (USER_DATA)
[**MarginCreateMarginOrderOcoV1**](V1API.md#MarginCreateMarginOrderOcoV1) | **Post** /sapi/v1/margin/order/oco | Margin Account New OCO (TRADE)
[**MarginCreateMarginOrderOtoV1**](V1API.md#MarginCreateMarginOrderOtoV1) | **Post** /sapi/v1/margin/order/oto | Margin Account New OTO (TRADE)
[**MarginCreateMarginOrderOtocoV1**](V1API.md#MarginCreateMarginOrderOtocoV1) | **Post** /sapi/v1/margin/order/otoco | Margin Account New OTOCO (TRADE)
[**MarginCreateMarginOrderV1**](V1API.md#MarginCreateMarginOrderV1) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**MarginCreateUserDataStreamIsolatedV1**](V1API.md#MarginCreateUserDataStreamIsolatedV1) | **Post** /sapi/v1/userDataStream/isolated | Start Isolated Margin User Data Stream (USER_STREAM)
[**MarginCreateUserDataStreamV1**](V1API.md#MarginCreateUserDataStreamV1) | **Post** /sapi/v1/userDataStream | Start Margin User Data Stream (USER_STREAM)
[**MarginDeleteMarginApiKeyV1**](V1API.md#MarginDeleteMarginApiKeyV1) | **Delete** /sapi/v1/margin/apiKey | Delete Special Key(Low-Latency Trading)(TRADE)
[**MarginDeleteMarginIsolatedAccountV1**](V1API.md#MarginDeleteMarginIsolatedAccountV1) | **Delete** /sapi/v1/margin/isolated/account | Disable Isolated Margin Account (TRADE)
[**MarginDeleteMarginListenKeyV1**](V1API.md#MarginDeleteMarginListenKeyV1) | **Delete** /sapi/v1/margin/listen-key | Close User Data Stream (USER_STREAM)
[**MarginDeleteMarginOpenOrdersV1**](V1API.md#MarginDeleteMarginOpenOrdersV1) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol (TRADE)
[**MarginDeleteMarginOrderListV1**](V1API.md#MarginDeleteMarginOrderListV1) | **Delete** /sapi/v1/margin/orderList | Margin Account Cancel OCO (TRADE)
[**MarginDeleteMarginOrderV1**](V1API.md#MarginDeleteMarginOrderV1) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**MarginDeleteUserDataStreamIsolatedV1**](V1API.md#MarginDeleteUserDataStreamIsolatedV1) | **Delete** /sapi/v1/userDataStream/isolated | Close Isolated Margin User Data Stream (USER_STREAM)
[**MarginDeleteUserDataStreamV1**](V1API.md#MarginDeleteUserDataStreamV1) | **Delete** /sapi/v1/userDataStream | Close Margin User Data Stream (USER_STREAM)
[**MarginGetBnbBurnV1**](V1API.md#MarginGetBnbBurnV1) | **Get** /sapi/v1/bnbBurn | Get BNB Burn Status (USER_DATA)
[**MarginGetMarginAccountV1**](V1API.md#MarginGetMarginAccountV1) | **Get** /sapi/v1/margin/account | Query Cross Margin Account Details (USER_DATA)
[**MarginGetMarginAllAssetsV1**](V1API.md#MarginGetMarginAllAssetsV1) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**MarginGetMarginAllOrderListV1**](V1API.md#MarginGetMarginAllOrderListV1) | **Get** /sapi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**MarginGetMarginAllOrdersV1**](V1API.md#MarginGetMarginAllOrdersV1) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Orders (USER_DATA)
[**MarginGetMarginAllPairsV1**](V1API.md#MarginGetMarginAllPairsV1) | **Get** /sapi/v1/margin/allPairs | Get All Cross Margin Pairs (MARKET_DATA)
[**MarginGetMarginApiKeyListV1**](V1API.md#MarginGetMarginApiKeyListV1) | **Get** /sapi/v1/margin/api-key-list | Query Special key List(Low Latency Trading)(TRADE)
[**MarginGetMarginApiKeyV1**](V1API.md#MarginGetMarginApiKeyV1) | **Get** /sapi/v1/margin/apiKey | Query Special key(Low Latency Trading)(TRADE)
[**MarginGetMarginAvailableInventoryV1**](V1API.md#MarginGetMarginAvailableInventoryV1) | **Get** /sapi/v1/margin/available-inventory | Query Margin Available Inventory(USER_DATA)
[**MarginGetMarginBorrowRepayV1**](V1API.md#MarginGetMarginBorrowRepayV1) | **Get** /sapi/v1/margin/borrow-repay | Query borrow/repay records in Margin account(USER_DATA)
[**MarginGetMarginCapitalFlowV1**](V1API.md#MarginGetMarginCapitalFlowV1) | **Get** /sapi/v1/margin/capital-flow | Query Cross Isolated Margin Capital Flow (USER_DATA)
[**MarginGetMarginCrossMarginCollateralRatioV1**](V1API.md#MarginGetMarginCrossMarginCollateralRatioV1) | **Get** /sapi/v1/margin/crossMarginCollateralRatio | Cross margin collateral ratio (MARKET_DATA)
[**MarginGetMarginCrossMarginDataV1**](V1API.md#MarginGetMarginCrossMarginDataV1) | **Get** /sapi/v1/margin/crossMarginData | Query Cross Margin Fee Data (USER_DATA)
[**MarginGetMarginDelistScheduleV1**](V1API.md#MarginGetMarginDelistScheduleV1) | **Get** /sapi/v1/margin/delist-schedule | Get Delist Schedule (MARKET_DATA)
[**MarginGetMarginExchangeSmallLiabilityHistoryV1**](V1API.md#MarginGetMarginExchangeSmallLiabilityHistoryV1) | **Get** /sapi/v1/margin/exchange-small-liability-history | Get Small Liability Exchange History (USER_DATA)
[**MarginGetMarginExchangeSmallLiabilityV1**](V1API.md#MarginGetMarginExchangeSmallLiabilityV1) | **Get** /sapi/v1/margin/exchange-small-liability | Get Small Liability Exchange Coin List (USER_DATA)
[**MarginGetMarginForceLiquidationRecV1**](V1API.md#MarginGetMarginForceLiquidationRecV1) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**MarginGetMarginInterestHistoryV1**](V1API.md#MarginGetMarginInterestHistoryV1) | **Get** /sapi/v1/margin/interestHistory | Get Interest History (USER_DATA)
[**MarginGetMarginInterestRateHistoryV1**](V1API.md#MarginGetMarginInterestRateHistoryV1) | **Get** /sapi/v1/margin/interestRateHistory | Query Margin Interest Rate History (USER_DATA)
[**MarginGetMarginIsolatedAccountLimitV1**](V1API.md#MarginGetMarginIsolatedAccountLimitV1) | **Get** /sapi/v1/margin/isolated/accountLimit | Query Enabled Isolated Margin Account Limit (USER_DATA)
[**MarginGetMarginIsolatedAccountV1**](V1API.md#MarginGetMarginIsolatedAccountV1) | **Get** /sapi/v1/margin/isolated/account | Query Isolated Margin Account Info (USER_DATA)
[**MarginGetMarginIsolatedAllPairsV1**](V1API.md#MarginGetMarginIsolatedAllPairsV1) | **Get** /sapi/v1/margin/isolated/allPairs | Get All Isolated Margin Symbol(MARKET_DATA)
[**MarginGetMarginIsolatedMarginDataV1**](V1API.md#MarginGetMarginIsolatedMarginDataV1) | **Get** /sapi/v1/margin/isolatedMarginData | Query Isolated Margin Fee Data (USER_DATA)
[**MarginGetMarginIsolatedMarginTierV1**](V1API.md#MarginGetMarginIsolatedMarginTierV1) | **Get** /sapi/v1/margin/isolatedMarginTier | Query Isolated Margin Tier Data (USER_DATA)
[**MarginGetMarginLeverageBracketV1**](V1API.md#MarginGetMarginLeverageBracketV1) | **Get** /sapi/v1/margin/leverageBracket | Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)
[**MarginGetMarginMaxBorrowableV1**](V1API.md#MarginGetMarginMaxBorrowableV1) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)
[**MarginGetMarginMaxTransferableV1**](V1API.md#MarginGetMarginMaxTransferableV1) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)
[**MarginGetMarginMyTradesV1**](V1API.md#MarginGetMarginMyTradesV1) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#39;s Trade List (USER_DATA)
[**MarginGetMarginNextHourlyInterestRateV1**](V1API.md#MarginGetMarginNextHourlyInterestRateV1) | **Get** /sapi/v1/margin/next-hourly-interest-rate | Get future hourly interest rate (USER_DATA)
[**MarginGetMarginOpenOrderListV1**](V1API.md#MarginGetMarginOpenOrderListV1) | **Get** /sapi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**MarginGetMarginOpenOrdersV1**](V1API.md#MarginGetMarginOpenOrdersV1) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)
[**MarginGetMarginOrderListV1**](V1API.md#MarginGetMarginOrderListV1) | **Get** /sapi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**MarginGetMarginOrderV1**](V1API.md#MarginGetMarginOrderV1) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (USER_DATA)
[**MarginGetMarginPriceIndexV1**](V1API.md#MarginGetMarginPriceIndexV1) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)
[**MarginGetMarginRateLimitOrderV1**](V1API.md#MarginGetMarginRateLimitOrderV1) | **Get** /sapi/v1/margin/rateLimit/order | Query Current Margin Order Count Usage (TRADE)
[**MarginGetMarginTradeCoeffV1**](V1API.md#MarginGetMarginTradeCoeffV1) | **Get** /sapi/v1/margin/tradeCoeff | Get Summary of Margin account (USER_DATA)
[**MarginGetMarginTransferV1**](V1API.md#MarginGetMarginTransferV1) | **Get** /sapi/v1/margin/transfer | Get Cross Margin Transfer History (USER_DATA)
[**MarginUpdateMarginApiKeyIpV1**](V1API.md#MarginUpdateMarginApiKeyIpV1) | **Put** /sapi/v1/margin/apiKey/ip | Edit ip for Special Key(Low-Latency Trading)(TRADE)
[**MarginUpdateMarginListenKeyV1**](V1API.md#MarginUpdateMarginListenKeyV1) | **Put** /sapi/v1/margin/listen-key | Keepalive User Data Stream (USER_STREAM)
[**MarginUpdateUserDataStreamIsolatedV1**](V1API.md#MarginUpdateUserDataStreamIsolatedV1) | **Put** /sapi/v1/userDataStream/isolated | Keepalive Isolated Margin User Data Stream (USER_STREAM)
[**MarginUpdateUserDataStreamV1**](V1API.md#MarginUpdateUserDataStreamV1) | **Put** /sapi/v1/userDataStream | Keepalive Margin User Data Stream (USER_STREAM)



## MarginCreateMarginApiKeyV1

> MarginCreateMarginApiKeyV1Resp MarginCreateMarginApiKeyV1(ctx).ApiName(apiName).Timestamp(timestamp).Ip(ip).PermissionMode(permissionMode).PublicKey(publicKey).RecvWindow(recvWindow).Symbol(symbol).Execute()

Create Special Key(Low-Latency Trading)(TRADE)



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
	apiName := "apiName_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ip := "ip_example" // string |  (optional) (default to "")
	permissionMode := "permissionMode_example" // string |  (optional) (default to "")
	publicKey := "publicKey_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginApiKeyV1(context.Background()).ApiName(apiName).Timestamp(timestamp).Ip(ip).PermissionMode(permissionMode).PublicKey(publicKey).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginApiKeyV1`: MarginCreateMarginApiKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiName** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **ip** | **string** |  | [default to &quot;&quot;]
 **permissionMode** | **string** |  | [default to &quot;&quot;]
 **publicKey** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginApiKeyV1Resp**](MarginCreateMarginApiKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.V1API.MarginCreateMarginBorrowRepayV1(context.Background()).Amount(amount).Asset(asset).IsIsolated(isIsolated).Symbol(symbol).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginBorrowRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginBorrowRepayV1`: MarginCreateMarginBorrowRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginBorrowRepayV1`: %v\n", resp)
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


## MarginCreateMarginExchangeSmallLiabilityV1

> MarginCreateMarginExchangeSmallLiabilityV1Resp MarginCreateMarginExchangeSmallLiabilityV1(ctx).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Small Liability Exchange (MARGIN)



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
	assetNames := []string{"Inner_example"} // []string | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginExchangeSmallLiabilityV1(context.Background()).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginExchangeSmallLiabilityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginExchangeSmallLiabilityV1`: MarginCreateMarginExchangeSmallLiabilityV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginExchangeSmallLiabilityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginExchangeSmallLiabilityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetNames** | **[]string** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginCreateMarginExchangeSmallLiabilityV1Resp**](MarginCreateMarginExchangeSmallLiabilityV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginIsolatedAccountV1

> MarginCreateMarginIsolatedAccountV1Resp MarginCreateMarginIsolatedAccountV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Isolated Margin Account (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginIsolatedAccountV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginIsolatedAccountV1`: MarginCreateMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginCreateMarginIsolatedAccountV1Resp**](MarginCreateMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginListenKeyV1

> MarginCreateMarginListenKeyV1Resp MarginCreateMarginListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginListenKeyV1`: MarginCreateMarginListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginListenKeyV1Request struct via the builder pattern


### Return type

[**MarginCreateMarginListenKeyV1Resp**](MarginCreateMarginListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginManualLiquidationV1

> MarginCreateMarginManualLiquidationV1Resp MarginCreateMarginManualLiquidationV1(ctx).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Symbol(symbol).Execute()

Margin Manual Liquidation(MARGIN)



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
	type_ := "type__example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginManualLiquidationV1(context.Background()).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginManualLiquidationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginManualLiquidationV1`: MarginCreateMarginManualLiquidationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginManualLiquidationV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginManualLiquidationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginManualLiquidationV1Resp**](MarginCreateMarginManualLiquidationV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginMaxLeverageV1

> MarginCreateMarginMaxLeverageV1Resp MarginCreateMarginMaxLeverageV1(ctx).MaxLeverage(maxLeverage).Execute()

Adjust cross margin max leverage (USER_DATA)



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
	maxLeverage := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginMaxLeverageV1(context.Background()).MaxLeverage(maxLeverage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginMaxLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginMaxLeverageV1`: MarginCreateMarginMaxLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginMaxLeverageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginMaxLeverageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxLeverage** | **int32** |  | 

### Return type

[**MarginCreateMarginMaxLeverageV1Resp**](MarginCreateMarginMaxLeverageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderOcoV1

> MarginCreateMarginOrderOcoV1Resp MarginCreateMarginOrderOcoV1(ctx).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()

Margin Account New OCO (TRADE)



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
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	stopPrice := "stopPrice_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	autoRepayAtCancel := true // bool |  (optional)
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	limitClientOrderId := "limitClientOrderId_example" // string |  (optional) (default to "")
	limitIcebergQty := "limitIcebergQty_example" // string |  (optional) (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	stopClientOrderId := "stopClientOrderId_example" // string |  (optional) (default to "")
	stopIcebergQty := "stopIcebergQty_example" // string |  (optional) (default to "")
	stopLimitPrice := "stopLimitPrice_example" // string |  (optional) (default to "")
	stopLimitTimeInForce := "stopLimitTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginOrderOcoV1(context.Background()).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginOrderOcoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderOcoV1`: MarginCreateMarginOrderOcoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginOrderOcoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderOcoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **autoRepayAtCancel** | **bool** |  | 
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **limitClientOrderId** | **string** |  | [default to &quot;&quot;]
 **limitIcebergQty** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **stopClientOrderId** | **string** |  | [default to &quot;&quot;]
 **stopIcebergQty** | **string** |  | [default to &quot;&quot;]
 **stopLimitPrice** | **string** |  | [default to &quot;&quot;]
 **stopLimitTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderOcoV1Resp**](MarginCreateMarginOrderOcoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderOtoV1

> MarginCreateMarginOrderOtoV1Resp MarginCreateMarginOrderOtoV1(ctx).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).WorkingIcebergQty(workingIcebergQty).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).Execute()

Margin Account New OTO (TRADE)



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
	pendingQuantity := "pendingQuantity_example" // string |  (default to "")
	pendingSide := "pendingSide_example" // string |  (default to "")
	pendingType := "pendingType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	workingIcebergQty := "workingIcebergQty_example" // string |  (default to "")
	workingPrice := "workingPrice_example" // string |  (default to "")
	workingQuantity := "workingQuantity_example" // string |  (default to "")
	workingSide := "workingSide_example" // string |  (default to "")
	workingType := "workingType_example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	pendingClientOrderId := "pendingClientOrderId_example" // string |  (optional) (default to "")
	pendingIcebergQty := "pendingIcebergQty_example" // string |  (optional) (default to "")
	pendingPrice := "pendingPrice_example" // string |  (optional) (default to "")
	pendingStopPrice := "pendingStopPrice_example" // string |  (optional) (default to "")
	pendingTimeInForce := "pendingTimeInForce_example" // string |  (optional) (default to "")
	pendingTrailingDelta := "pendingTrailingDelta_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	workingClientOrderId := "workingClientOrderId_example" // string |  (optional) (default to "")
	workingTimeInForce := "workingTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginOrderOtoV1(context.Background()).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).WorkingIcebergQty(workingIcebergQty).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginOrderOtoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderOtoV1`: MarginCreateMarginOrderOtoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginOrderOtoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderOtoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingQuantity** | **string** |  | [default to &quot;&quot;]
 **pendingSide** | **string** |  | [default to &quot;&quot;]
 **pendingType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **workingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **workingPrice** | **string** |  | [default to &quot;&quot;]
 **workingQuantity** | **string** |  | [default to &quot;&quot;]
 **workingSide** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **pendingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingPrice** | **string** |  | [default to &quot;&quot;]
 **pendingStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **workingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **workingTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderOtoV1Resp**](MarginCreateMarginOrderOtoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderOtocoV1

> MarginCreateMarginOrderOtocoV1Resp MarginCreateMarginOrderOtocoV1(ctx).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).Execute()

Margin Account New OTOCO (TRADE)



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
	pendingAboveType := "pendingAboveType_example" // string |  (default to "")
	pendingQuantity := "pendingQuantity_example" // string |  (default to "")
	pendingSide := "pendingSide_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	workingPrice := "workingPrice_example" // string |  (default to "")
	workingQuantity := "workingQuantity_example" // string |  (default to "")
	workingSide := "workingSide_example" // string |  (default to "")
	workingType := "workingType_example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	pendingAboveClientOrderId := "pendingAboveClientOrderId_example" // string |  (optional) (default to "")
	pendingAboveIcebergQty := "pendingAboveIcebergQty_example" // string |  (optional) (default to "")
	pendingAbovePrice := "pendingAbovePrice_example" // string |  (optional) (default to "")
	pendingAboveStopPrice := "pendingAboveStopPrice_example" // string |  (optional) (default to "")
	pendingAboveTimeInForce := "pendingAboveTimeInForce_example" // string |  (optional) (default to "")
	pendingAboveTrailingDelta := "pendingAboveTrailingDelta_example" // string |  (optional) (default to "")
	pendingBelowClientOrderId := "pendingBelowClientOrderId_example" // string |  (optional) (default to "")
	pendingBelowIcebergQty := "pendingBelowIcebergQty_example" // string |  (optional) (default to "")
	pendingBelowPrice := "pendingBelowPrice_example" // string |  (optional) (default to "")
	pendingBelowStopPrice := "pendingBelowStopPrice_example" // string |  (optional) (default to "")
	pendingBelowTimeInForce := "pendingBelowTimeInForce_example" // string |  (optional) (default to "")
	pendingBelowTrailingDelta := "pendingBelowTrailingDelta_example" // string |  (optional) (default to "")
	pendingBelowType := "pendingBelowType_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	workingClientOrderId := "workingClientOrderId_example" // string |  (optional) (default to "")
	workingIcebergQty := "workingIcebergQty_example" // string |  (optional) (default to "")
	workingTimeInForce := "workingTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginOrderOtocoV1(context.Background()).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginOrderOtocoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderOtocoV1`: MarginCreateMarginOrderOtocoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginOrderOtocoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderOtocoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingAboveType** | **string** |  | [default to &quot;&quot;]
 **pendingQuantity** | **string** |  | [default to &quot;&quot;]
 **pendingSide** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **workingPrice** | **string** |  | [default to &quot;&quot;]
 **workingQuantity** | **string** |  | [default to &quot;&quot;]
 **workingSide** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **pendingAboveClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingAboveIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingAbovePrice** | **string** |  | [default to &quot;&quot;]
 **pendingAboveStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingAboveTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingAboveTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **pendingBelowClientOrderId** | **string** |  | [default to &quot;&quot;]
 **pendingBelowIcebergQty** | **string** |  | [default to &quot;&quot;]
 **pendingBelowPrice** | **string** |  | [default to &quot;&quot;]
 **pendingBelowStopPrice** | **string** |  | [default to &quot;&quot;]
 **pendingBelowTimeInForce** | **string** |  | [default to &quot;&quot;]
 **pendingBelowTrailingDelta** | **string** |  | [default to &quot;&quot;]
 **pendingBelowType** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **workingClientOrderId** | **string** |  | [default to &quot;&quot;]
 **workingIcebergQty** | **string** |  | [default to &quot;&quot;]
 **workingTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderOtocoV1Resp**](MarginCreateMarginOrderOtocoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginOrderV1

> MarginCreateMarginOrderV1Resp MarginCreateMarginOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()

Margin Account New Order (TRADE)



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
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
	isIsolated := "isIsolated_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	quoteOrderQty := "quoteOrderQty_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateMarginOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginOrderV1`: MarginCreateMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **icebergQty** | **string** |  | [default to &quot;&quot;]
 **isIsolated** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **quoteOrderQty** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateMarginOrderV1Resp**](MarginCreateMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateUserDataStreamIsolatedV1

> MarginCreateUserDataStreamIsolatedV1Resp MarginCreateUserDataStreamIsolatedV1(ctx).Symbol(symbol).Execute()

Start Isolated Margin User Data Stream (USER_STREAM)



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
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateUserDataStreamIsolatedV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateUserDataStreamIsolatedV1`: MarginCreateUserDataStreamIsolatedV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateUserDataStreamIsolatedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateUserDataStreamIsolatedV1Resp**](MarginCreateUserDataStreamIsolatedV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateUserDataStreamV1

> MarginCreateUserDataStreamV1Resp MarginCreateUserDataStreamV1(ctx).Execute()

Start Margin User Data Stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginCreateUserDataStreamV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginCreateUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateUserDataStreamV1`: MarginCreateUserDataStreamV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginCreateUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateUserDataStreamV1Request struct via the builder pattern


### Return type

[**MarginCreateUserDataStreamV1Resp**](MarginCreateUserDataStreamV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginApiKeyV1

> map[string]interface{} MarginDeleteMarginApiKeyV1(ctx).Timestamp(timestamp).ApiKey(apiKey).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()

Delete Special Key(Low-Latency Trading)(TRADE)



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
	apiKey := "apiKey_example" // string |  (optional) (default to "")
	apiName := "apiName_example" // string |  (optional) (default to "")
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteMarginApiKeyV1(context.Background()).Timestamp(timestamp).ApiKey(apiKey).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginApiKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **apiKey** | **string** |  | [default to &quot;&quot;]
 **apiName** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginIsolatedAccountV1

> MarginDeleteMarginIsolatedAccountV1Resp MarginDeleteMarginIsolatedAccountV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Disable Isolated Margin Account (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteMarginIsolatedAccountV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginIsolatedAccountV1`: MarginDeleteMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginDeleteMarginIsolatedAccountV1Resp**](MarginDeleteMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginListenKeyV1

> map[string]interface{} MarginDeleteMarginListenKeyV1(ctx).Execute()

Close User Data Stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteMarginListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginListenKeyV1Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginOpenOrdersV1

> []MarginDeleteMarginOpenOrdersV1RespItem MarginDeleteMarginOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Margin Account Cancel all Open Orders on a Symbol (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteMarginOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginOpenOrdersV1`: []MarginDeleteMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginDeleteMarginOpenOrdersV1RespItem**](MarginDeleteMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginOrderListV1

> MarginDeleteMarginOrderListV1Resp MarginDeleteMarginOrderListV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel OCO (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderListId := int64(789) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "listClientOrderId_example" // string | Either `orderListId` or `listClientOrderId` must be provided (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteMarginOrderListV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginOrderListV1`: MarginDeleteMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginDeleteMarginOrderListV1Resp**](MarginDeleteMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginOrderV1

> MarginDeleteMarginOrderV1Resp MarginDeleteMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel Order (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginOrderV1`: MarginDeleteMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginDeleteMarginOrderV1Resp**](MarginDeleteMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteUserDataStreamIsolatedV1

> map[string]interface{} MarginDeleteUserDataStreamIsolatedV1(ctx).Symbol(symbol).Listenkey(listenkey).Execute()

Close Isolated Margin User Data Stream (USER_STREAM)



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
	symbol := "symbol_example" // string |  (default to "")
	listenkey := "listenkey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteUserDataStreamIsolatedV1(context.Background()).Symbol(symbol).Listenkey(listenkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteUserDataStreamIsolatedV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteUserDataStreamIsolatedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **listenkey** | **string** |  | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteUserDataStreamV1

> map[string]interface{} MarginDeleteUserDataStreamV1(ctx).Listenkey(listenkey).Execute()

Close Margin User Data Stream (USER_STREAM)



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
	listenkey := "listenkey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginDeleteUserDataStreamV1(context.Background()).Listenkey(listenkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginDeleteUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteUserDataStreamV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginDeleteUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteUserDataStreamV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenkey** | **string** |  | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetBnbBurnV1

> MarginGetBnbBurnV1Resp MarginGetBnbBurnV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get BNB Burn Status (USER_DATA)



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
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetBnbBurnV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetBnbBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetBnbBurnV1`: MarginGetBnbBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetBnbBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetBnbBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginGetBnbBurnV1Resp**](MarginGetBnbBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAccountV1

> MarginGetMarginAccountV1Resp MarginGetMarginAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Cross Margin Account Details (USER_DATA)



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
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAccountV1`: MarginGetMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginAccountV1Resp**](MarginGetMarginAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAllAssetsV1

> []MarginGetMarginAllAssetsV1RespItem MarginGetMarginAllAssetsV1(ctx).Asset(asset).Execute()

Get All Margin Assets (MARKET_DATA)



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
	asset := "asset_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginAllAssetsV1(context.Background()).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginAllAssetsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllAssetsV1`: []MarginGetMarginAllAssetsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginAllAssetsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAllAssetsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]MarginGetMarginAllAssetsV1RespItem**](MarginGetMarginAllAssetsV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAllOrderListV1

> []MarginGetMarginAllOrderListV1RespItem MarginGetMarginAllOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	fromId := int64(789) // int64 | If supplied, neither `startTime` or `endTime` can be provided (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default Value: 500; Max Value: 1000 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginAllOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginAllOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllOrderListV1`: []MarginGetMarginAllOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginAllOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAllOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **fromId** | **int64** | If supplied, neither &#x60;startTime&#x60; or &#x60;endTime&#x60; can be provided | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default Value: 500; Max Value: 1000 | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginAllOrderListV1RespItem**](MarginGetMarginAllOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAllOrdersV1

> []MarginGetMarginAllOrdersV1RespItem MarginGetMarginAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's All Orders (USER_DATA)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 500. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllOrdersV1`: []MarginGetMarginAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 500. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginAllOrdersV1RespItem**](MarginGetMarginAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAllPairsV1

> []MarginGetMarginAllPairsV1RespItem MarginGetMarginAllPairsV1(ctx).Symbol(symbol).Execute()

Get All Cross Margin Pairs (MARKET_DATA)



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginAllPairsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginAllPairsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllPairsV1`: []MarginGetMarginAllPairsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginAllPairsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAllPairsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]MarginGetMarginAllPairsV1RespItem**](MarginGetMarginAllPairsV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginApiKeyListV1

> []MarginGetMarginApiKeyListV1RespItem MarginGetMarginApiKeyListV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key List(Low Latency Trading)(TRADE)



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
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginApiKeyListV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginApiKeyListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginApiKeyListV1`: []MarginGetMarginApiKeyListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginApiKeyListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginApiKeyListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginApiKeyListV1RespItem**](MarginGetMarginApiKeyListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginApiKeyV1

> MarginGetMarginApiKeyV1Resp MarginGetMarginApiKeyV1(ctx).ApiKey(apiKey).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key(Low Latency Trading)(TRADE)



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
	apiKey := "apiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginApiKeyV1(context.Background()).ApiKey(apiKey).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginApiKeyV1`: MarginGetMarginApiKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginApiKeyV1Resp**](MarginGetMarginApiKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAvailableInventoryV1

> MarginGetMarginAvailableInventoryV1Resp MarginGetMarginAvailableInventoryV1(ctx).Type_(type_).Execute()

Query Margin Available Inventory(USER_DATA)



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
	type_ := "type__example" // string | MARGIN,ISOLATED (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginAvailableInventoryV1(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginAvailableInventoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAvailableInventoryV1`: MarginGetMarginAvailableInventoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginAvailableInventoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAvailableInventoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | MARGIN,ISOLATED | [default to &quot;&quot;]

### Return type

[**MarginGetMarginAvailableInventoryV1Resp**](MarginGetMarginAvailableInventoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.V1API.MarginGetMarginBorrowRepayV1(context.Background()).Type_(type_).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginBorrowRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginBorrowRepayV1`: MarginGetMarginBorrowRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginBorrowRepayV1`: %v\n", resp)
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


## MarginGetMarginCapitalFlowV1

> []MarginGetMarginCapitalFlowV1RespItem MarginGetMarginCapitalFlowV1(ctx).Timestamp(timestamp).Asset(asset).Symbol(symbol).Type_(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Cross Isolated Margin Capital Flow (USER_DATA)



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
	symbol := "symbol_example" // string | 查询逐仓数据时必填 (optional) (default to "")
	type_ := "type__example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | 只支持查询最近90天的数据 (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | 如设置fromId, 将返回id &gt; fromId的数据。否则将返回最新数据 (optional)
	limit := int64(789) // int64 | 每次返回的数据条数限制。默认 500; 最大 1000. (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginCapitalFlowV1(context.Background()).Timestamp(timestamp).Asset(asset).Symbol(symbol).Type_(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginCapitalFlowV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginCapitalFlowV1`: []MarginGetMarginCapitalFlowV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginCapitalFlowV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginCapitalFlowV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** | 查询逐仓数据时必填 | [default to &quot;&quot;]
 **type_** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | 只支持查询最近90天的数据 | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | 如设置fromId, 将返回id &amp;gt; fromId的数据。否则将返回最新数据 | 
 **limit** | **int64** | 每次返回的数据条数限制。默认 500; 最大 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginCapitalFlowV1RespItem**](MarginGetMarginCapitalFlowV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginCrossMarginCollateralRatioV1

> []MarginGetMarginCrossMarginCollateralRatioV1RespItem MarginGetMarginCrossMarginCollateralRatioV1(ctx).Execute()

Cross margin collateral ratio (MARKET_DATA)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginCrossMarginCollateralRatioV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginCrossMarginCollateralRatioV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginCrossMarginCollateralRatioV1`: []MarginGetMarginCrossMarginCollateralRatioV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginCrossMarginCollateralRatioV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginCrossMarginCollateralRatioV1Request struct via the builder pattern


### Return type

[**[]MarginGetMarginCrossMarginCollateralRatioV1RespItem**](MarginGetMarginCrossMarginCollateralRatioV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginCrossMarginDataV1

> []MarginGetMarginCrossMarginDataV1RespItem MarginGetMarginCrossMarginDataV1(ctx).Timestamp(timestamp).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()

Query Cross Margin Fee Data (USER_DATA)



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
	vipLevel := int32(56) // int32 | User&#39;s current specific margin data will be returned if vipLevel is omitted (optional)
	coin := "coin_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginCrossMarginDataV1(context.Background()).Timestamp(timestamp).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginCrossMarginDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginCrossMarginDataV1`: []MarginGetMarginCrossMarginDataV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginCrossMarginDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginCrossMarginDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | User&amp;#39;s current specific margin data will be returned if vipLevel is omitted | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginCrossMarginDataV1RespItem**](MarginGetMarginCrossMarginDataV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginDelistScheduleV1

> []MarginGetMarginDelistScheduleV1RespItem MarginGetMarginDelistScheduleV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Delist Schedule (MARKET_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginDelistScheduleV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginDelistScheduleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginDelistScheduleV1`: []MarginGetMarginDelistScheduleV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginDelistScheduleV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginDelistScheduleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]MarginGetMarginDelistScheduleV1RespItem**](MarginGetMarginDelistScheduleV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginExchangeSmallLiabilityHistoryV1

> MarginGetMarginExchangeSmallLiabilityHistoryV1Resp MarginGetMarginExchangeSmallLiabilityHistoryV1(ctx).Current(current).Size(size).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange History (USER_DATA)



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
	current := int32(56) // int32 | Currently querying page. Start from 1. Default:1
	size := int32(56) // int32 | Default:10, Max:100
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 | Default: 30 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginExchangeSmallLiabilityHistoryV1(context.Background()).Current(current).Size(size).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginExchangeSmallLiabilityHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginExchangeSmallLiabilityHistoryV1`: MarginGetMarginExchangeSmallLiabilityHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginExchangeSmallLiabilityHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginExchangeSmallLiabilityHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **current** | **int32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10, Max:100 | 
 **timestamp** | **int64** |  | 
 **startTime** | **int64** | Default: 30 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginGetMarginExchangeSmallLiabilityHistoryV1Resp**](MarginGetMarginExchangeSmallLiabilityHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginExchangeSmallLiabilityV1

> []MarginGetMarginExchangeSmallLiabilityV1RespItem MarginGetMarginExchangeSmallLiabilityV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange Coin List (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginExchangeSmallLiabilityV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginExchangeSmallLiabilityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginExchangeSmallLiabilityV1`: []MarginGetMarginExchangeSmallLiabilityV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginExchangeSmallLiabilityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginExchangeSmallLiabilityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]MarginGetMarginExchangeSmallLiabilityV1RespItem**](MarginGetMarginExchangeSmallLiabilityV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginForceLiquidationRecV1

> MarginGetMarginForceLiquidationRecV1Resp MarginGetMarginForceLiquidationRecV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Force Liquidation Record (USER_DATA)



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
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	isolatedSymbol := "isolatedSymbol_example" // string |  (optional) (default to "")
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginForceLiquidationRecV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginForceLiquidationRecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginForceLiquidationRecV1`: MarginGetMarginForceLiquidationRecV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginForceLiquidationRecV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginForceLiquidationRecV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **isolatedSymbol** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginForceLiquidationRecV1Resp**](MarginGetMarginForceLiquidationRecV1Resp.md)

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
	resp, r, err := apiClient.V1API.MarginGetMarginInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginInterestHistoryV1`: MarginGetMarginInterestHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginInterestHistoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.MarginGetMarginInterestRateHistoryV1(context.Background()).Asset(asset).Timestamp(timestamp).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginInterestRateHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginInterestRateHistoryV1`: []MarginGetMarginInterestRateHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginInterestRateHistoryV1`: %v\n", resp)
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


## MarginGetMarginIsolatedAccountLimitV1

> MarginGetMarginIsolatedAccountLimitV1Resp MarginGetMarginIsolatedAccountLimitV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Enabled Isolated Margin Account Limit (USER_DATA)



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
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginIsolatedAccountLimitV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginIsolatedAccountLimitV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedAccountLimitV1`: MarginGetMarginIsolatedAccountLimitV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginIsolatedAccountLimitV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedAccountLimitV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginGetMarginIsolatedAccountLimitV1Resp**](MarginGetMarginIsolatedAccountLimitV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginIsolatedAccountV1

> MarginGetMarginIsolatedAccountV1Resp MarginGetMarginIsolatedAccountV1(ctx).Timestamp(timestamp).Symbols(symbols).RecvWindow(recvWindow).Execute()

Query Isolated Margin Account Info (USER_DATA)



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
	symbols := "symbols_example" // string | Max 5 symbols can be sent; separated by &#34;,&#34;. e.g. &#34;BTCUSDT,BNBUSDT,ADAUSDT&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginIsolatedAccountV1(context.Background()).Timestamp(timestamp).Symbols(symbols).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedAccountV1`: MarginGetMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbols** | **string** | Max 5 symbols can be sent; separated by &amp;#34;,&amp;#34;. e.g. &amp;#34;BTCUSDT,BNBUSDT,ADAUSDT&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginGetMarginIsolatedAccountV1Resp**](MarginGetMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginIsolatedAllPairsV1

> []MarginGetMarginIsolatedAllPairsV1RespItem MarginGetMarginIsolatedAllPairsV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Get All Isolated Margin Symbol(MARKET_DATA)



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
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginIsolatedAllPairsV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginIsolatedAllPairsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedAllPairsV1`: []MarginGetMarginIsolatedAllPairsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginIsolatedAllPairsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedAllPairsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**[]MarginGetMarginIsolatedAllPairsV1RespItem**](MarginGetMarginIsolatedAllPairsV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginIsolatedMarginDataV1

> []MarginGetMarginIsolatedMarginDataV1RespItem MarginGetMarginIsolatedMarginDataV1(ctx).Timestamp(timestamp).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Isolated Margin Fee Data (USER_DATA)



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
	vipLevel := int32(56) // int32 | User&#39;s current specific margin data will be returned if vipLevel is omitted (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginIsolatedMarginDataV1(context.Background()).Timestamp(timestamp).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginIsolatedMarginDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedMarginDataV1`: []MarginGetMarginIsolatedMarginDataV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginIsolatedMarginDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedMarginDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | User&amp;#39;s current specific margin data will be returned if vipLevel is omitted | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginIsolatedMarginDataV1RespItem**](MarginGetMarginIsolatedMarginDataV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginIsolatedMarginTierV1

> []MarginGetMarginIsolatedMarginTierV1RespItem MarginGetMarginIsolatedMarginTierV1(ctx).Symbol(symbol).Timestamp(timestamp).Tier(tier).RecvWindow(recvWindow).Execute()

Query Isolated Margin Tier Data (USER_DATA)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	tier := int32(56) // int32 | All margin tier data will be returned if tier is omitted (optional)
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginIsolatedMarginTierV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Tier(tier).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginIsolatedMarginTierV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedMarginTierV1`: []MarginGetMarginIsolatedMarginTierV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginIsolatedMarginTierV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedMarginTierV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **tier** | **int32** | All margin tier data will be returned if tier is omitted | 
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginIsolatedMarginTierV1RespItem**](MarginGetMarginIsolatedMarginTierV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginLeverageBracketV1

> []MarginGetMarginLeverageBracketV1RespItem MarginGetMarginLeverageBracketV1(ctx).Execute()

Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginLeverageBracketV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginLeverageBracketV1`: []MarginGetMarginLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginLeverageBracketV1Request struct via the builder pattern


### Return type

[**[]MarginGetMarginLeverageBracketV1RespItem**](MarginGetMarginLeverageBracketV1RespItem.md)

### Authorization

No authorization required

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
	resp, r, err := apiClient.V1API.MarginGetMarginMaxBorrowableV1(context.Background()).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginMaxBorrowableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginMaxBorrowableV1`: MarginGetMarginMaxBorrowableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginMaxBorrowableV1`: %v\n", resp)
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


## MarginGetMarginMaxTransferableV1

> MarginGetMarginMaxTransferableV1Resp MarginGetMarginMaxTransferableV1(ctx).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Transfer-Out Amount (USER_DATA)



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
	resp, r, err := apiClient.V1API.MarginGetMarginMaxTransferableV1(context.Background()).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginMaxTransferableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginMaxTransferableV1`: MarginGetMarginMaxTransferableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginMaxTransferableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginMaxTransferableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isolatedSymbol** | **string** | isolated symbol | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginMaxTransferableV1Resp**](MarginGetMarginMaxTransferableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginMyTradesV1

> []MarginGetMarginMyTradesV1RespItem MarginGetMarginMyTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's Trade List (USER_DATA)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginMyTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginMyTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginMyTradesV1`: []MarginGetMarginMyTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginMyTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginMyTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginMyTradesV1RespItem**](MarginGetMarginMyTradesV1RespItem.md)

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
	resp, r, err := apiClient.V1API.MarginGetMarginNextHourlyInterestRateV1(context.Background()).Assets(assets).IsIsolated(isIsolated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginNextHourlyInterestRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginNextHourlyInterestRateV1`: []MarginGetMarginNextHourlyInterestRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginNextHourlyInterestRateV1`: %v\n", resp)
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


## MarginGetMarginOpenOrderListV1

> []MarginGetMarginOpenOrderListV1RespItem MarginGetMarginOpenOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginOpenOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginOpenOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOpenOrderListV1`: []MarginGetMarginOpenOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginOpenOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOpenOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginOpenOrderListV1RespItem**](MarginGetMarginOpenOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOpenOrdersV1

> []MarginGetMarginOpenOrdersV1RespItem MarginGetMarginOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Margin Account's Open Orders (USER_DATA)



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
	symbol := "symbol_example" // string |  (optional) (default to "")
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOpenOrdersV1`: []MarginGetMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginOpenOrdersV1RespItem**](MarginGetMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOrderListV1

> MarginGetMarginOrderListV1Resp MarginGetMarginOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	orderListId := int64(789) // int64 | Either `orderListId` or `origClientOrderId` must be provided (optional)
	origClientOrderId := "origClientOrderId_example" // string | Either `orderListId` or `origClientOrderId` must be provided (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOrderListV1`: MarginGetMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginOrderListV1Resp**](MarginGetMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginOrderV1

> MarginGetMarginOrderV1Resp MarginGetMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's Order (USER_DATA)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginOrderV1`: MarginGetMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginOrderV1Resp**](MarginGetMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginPriceIndexV1

> MarginGetMarginPriceIndexV1Resp MarginGetMarginPriceIndexV1(ctx).Symbol(symbol).Execute()

Query Margin PriceIndex (MARKET_DATA)



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
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginPriceIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginPriceIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginPriceIndexV1`: MarginGetMarginPriceIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginPriceIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginPriceIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginGetMarginPriceIndexV1Resp**](MarginGetMarginPriceIndexV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginRateLimitOrderV1

> []MarginGetMarginRateLimitOrderV1RespItem MarginGetMarginRateLimitOrderV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Current Margin Order Count Usage (TRADE)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | isolated symbol, mandatory for isolated margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginRateLimitOrderV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginRateLimitOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginRateLimitOrderV1`: []MarginGetMarginRateLimitOrderV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginRateLimitOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginRateLimitOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | isolated symbol, mandatory for isolated margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginRateLimitOrderV1RespItem**](MarginGetMarginRateLimitOrderV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginTradeCoeffV1

> MarginGetMarginTradeCoeffV1Resp MarginGetMarginTradeCoeffV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Summary of Margin account (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginTradeCoeffV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginTradeCoeffV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginTradeCoeffV1`: MarginGetMarginTradeCoeffV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginTradeCoeffV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginTradeCoeffV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginGetMarginTradeCoeffV1Resp**](MarginGetMarginTradeCoeffV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginTransferV1

> MarginGetMarginTransferV1Resp MarginGetMarginTransferV1(ctx).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Get Cross Margin Transfer History (USER_DATA)



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
	type_ := "type__example" // string | Transfer Type: ROLL_IN, ROLL_OUT (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	isolatedSymbol := "isolatedSymbol_example" // string | Symbol in Isolated Margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginGetMarginTransferV1(context.Background()).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginGetMarginTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginTransferV1`: MarginGetMarginTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginGetMarginTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **type_** | **string** | Transfer Type: ROLL_IN, ROLL_OUT | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **isolatedSymbol** | **string** | Symbol in Isolated Margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginTransferV1Resp**](MarginGetMarginTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginUpdateMarginApiKeyIpV1

> map[string]interface{} MarginUpdateMarginApiKeyIpV1(ctx).ApiKey(apiKey).Ip(ip).Timestamp(timestamp).RecvWindow(recvWindow).Symbol(symbol).Execute()

Edit ip for Special Key(Low-Latency Trading)(TRADE)



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
	apiKey := "apiKey_example" // string |  (default to "")
	ip := "ip_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginUpdateMarginApiKeyIpV1(context.Background()).ApiKey(apiKey).Ip(ip).Timestamp(timestamp).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginUpdateMarginApiKeyIpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateMarginApiKeyIpV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginUpdateMarginApiKeyIpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateMarginApiKeyIpV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | [default to &quot;&quot;]
 **ip** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]

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


## MarginUpdateMarginListenKeyV1

> map[string]interface{} MarginUpdateMarginListenKeyV1(ctx).ListenKey(listenKey).Execute()

Keepalive User Data Stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginUpdateMarginListenKeyV1(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginUpdateMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateMarginListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginUpdateMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateMarginListenKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]

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


## MarginUpdateUserDataStreamIsolatedV1

> map[string]interface{} MarginUpdateUserDataStreamIsolatedV1(ctx).ListenKey(listenKey).Symbol(symbol).Execute()

Keepalive Isolated Margin User Data Stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginUpdateUserDataStreamIsolatedV1(context.Background()).ListenKey(listenKey).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginUpdateUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateUserDataStreamIsolatedV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginUpdateUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateUserDataStreamIsolatedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]

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


## MarginUpdateUserDataStreamV1

> map[string]interface{} MarginUpdateUserDataStreamV1(ctx).ListenKey(listenKey).Execute()

Keepalive Margin User Data Stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.MarginUpdateUserDataStreamV1(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.MarginUpdateUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateUserDataStreamV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.MarginUpdateUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateUserDataStreamV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]

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

