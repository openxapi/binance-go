# \MarginTradingAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMarginApiKeyV1**](MarginTradingAPI.md#CreateMarginApiKeyV1) | **Post** /sapi/v1/margin/apiKey | Create Special Key(Low-Latency Trading)(TRADE)
[**CreateMarginBorrowRepayV1**](MarginTradingAPI.md#CreateMarginBorrowRepayV1) | **Post** /sapi/v1/margin/borrow-repay | Margin account borrow/repay(MARGIN)
[**CreateMarginExchangeSmallLiabilityV1**](MarginTradingAPI.md#CreateMarginExchangeSmallLiabilityV1) | **Post** /sapi/v1/margin/exchange-small-liability | Small Liability Exchange (MARGIN)
[**CreateMarginIsolatedAccountV1**](MarginTradingAPI.md#CreateMarginIsolatedAccountV1) | **Post** /sapi/v1/margin/isolated/account | Enable Isolated Margin Account (TRADE)
[**CreateMarginListenKeyV1**](MarginTradingAPI.md#CreateMarginListenKeyV1) | **Post** /sapi/v1/margin/listen-key | Start User Data Stream (USER_STREAM)
[**CreateMarginManualLiquidationV1**](MarginTradingAPI.md#CreateMarginManualLiquidationV1) | **Post** /sapi/v1/margin/manual-liquidation | Margin Manual Liquidation(MARGIN)
[**CreateMarginMaxLeverageV1**](MarginTradingAPI.md#CreateMarginMaxLeverageV1) | **Post** /sapi/v1/margin/max-leverage | Adjust cross margin max leverage (USER_DATA)
[**CreateMarginOrderOcoV1**](MarginTradingAPI.md#CreateMarginOrderOcoV1) | **Post** /sapi/v1/margin/order/oco | Margin Account New OCO (TRADE)
[**CreateMarginOrderOtoV1**](MarginTradingAPI.md#CreateMarginOrderOtoV1) | **Post** /sapi/v1/margin/order/oto | Margin Account New OTO (TRADE)
[**CreateMarginOrderOtocoV1**](MarginTradingAPI.md#CreateMarginOrderOtocoV1) | **Post** /sapi/v1/margin/order/otoco | Margin Account New OTOCO (TRADE)
[**CreateMarginOrderV1**](MarginTradingAPI.md#CreateMarginOrderV1) | **Post** /sapi/v1/margin/order | Margin Account New Order (TRADE)
[**CreateUserDataStreamIsolatedV1**](MarginTradingAPI.md#CreateUserDataStreamIsolatedV1) | **Post** /sapi/v1/userDataStream/isolated | Start Isolated Margin User Data Stream (USER_STREAM)
[**CreateUserDataStreamV1**](MarginTradingAPI.md#CreateUserDataStreamV1) | **Post** /sapi/v1/userDataStream | Start Margin User Data Stream (USER_STREAM)
[**DeleteMarginApiKeyV1**](MarginTradingAPI.md#DeleteMarginApiKeyV1) | **Delete** /sapi/v1/margin/apiKey | Delete Special Key(Low-Latency Trading)(TRADE)
[**DeleteMarginIsolatedAccountV1**](MarginTradingAPI.md#DeleteMarginIsolatedAccountV1) | **Delete** /sapi/v1/margin/isolated/account | Disable Isolated Margin Account (TRADE)
[**DeleteMarginListenKeyV1**](MarginTradingAPI.md#DeleteMarginListenKeyV1) | **Delete** /sapi/v1/margin/listen-key | Close User Data Stream (USER_STREAM)
[**DeleteMarginOpenOrdersV1**](MarginTradingAPI.md#DeleteMarginOpenOrdersV1) | **Delete** /sapi/v1/margin/openOrders | Margin Account Cancel all Open Orders on a Symbol (TRADE)
[**DeleteMarginOrderListV1**](MarginTradingAPI.md#DeleteMarginOrderListV1) | **Delete** /sapi/v1/margin/orderList | Margin Account Cancel OCO (TRADE)
[**DeleteMarginOrderV1**](MarginTradingAPI.md#DeleteMarginOrderV1) | **Delete** /sapi/v1/margin/order | Margin Account Cancel Order (TRADE)
[**DeleteUserDataStreamIsolatedV1**](MarginTradingAPI.md#DeleteUserDataStreamIsolatedV1) | **Delete** /sapi/v1/userDataStream/isolated | Close Isolated Margin User Data Stream (USER_STREAM)
[**DeleteUserDataStreamV1**](MarginTradingAPI.md#DeleteUserDataStreamV1) | **Delete** /sapi/v1/userDataStream | Close Margin User Data Stream (USER_STREAM)
[**GetBnbBurnV1**](MarginTradingAPI.md#GetBnbBurnV1) | **Get** /sapi/v1/bnbBurn | Get BNB Burn Status (USER_DATA)
[**GetMarginAccountV1**](MarginTradingAPI.md#GetMarginAccountV1) | **Get** /sapi/v1/margin/account | Query Cross Margin Account Details (USER_DATA)
[**GetMarginAllAssetsV1**](MarginTradingAPI.md#GetMarginAllAssetsV1) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**GetMarginAllOrderListV1**](MarginTradingAPI.md#GetMarginAllOrderListV1) | **Get** /sapi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**GetMarginAllOrdersV1**](MarginTradingAPI.md#GetMarginAllOrdersV1) | **Get** /sapi/v1/margin/allOrders | Query Margin Account&#39;s All Orders (USER_DATA)
[**GetMarginAllPairsV1**](MarginTradingAPI.md#GetMarginAllPairsV1) | **Get** /sapi/v1/margin/allPairs | Get All Cross Margin Pairs (MARKET_DATA)
[**GetMarginApiKeyListV1**](MarginTradingAPI.md#GetMarginApiKeyListV1) | **Get** /sapi/v1/margin/api-key-list | Query Special key List(Low Latency Trading)(TRADE)
[**GetMarginApiKeyV1**](MarginTradingAPI.md#GetMarginApiKeyV1) | **Get** /sapi/v1/margin/apiKey | Query Special key(Low Latency Trading)(TRADE)
[**GetMarginAvailableInventoryV1**](MarginTradingAPI.md#GetMarginAvailableInventoryV1) | **Get** /sapi/v1/margin/available-inventory | Query Margin Available Inventory(USER_DATA)
[**GetMarginBorrowRepayV1**](MarginTradingAPI.md#GetMarginBorrowRepayV1) | **Get** /sapi/v1/margin/borrow-repay | Query borrow/repay records in Margin account(USER_DATA)
[**GetMarginCapitalFlowV1**](MarginTradingAPI.md#GetMarginCapitalFlowV1) | **Get** /sapi/v1/margin/capital-flow | Query Cross Isolated Margin Capital Flow (USER_DATA)
[**GetMarginCrossMarginCollateralRatioV1**](MarginTradingAPI.md#GetMarginCrossMarginCollateralRatioV1) | **Get** /sapi/v1/margin/crossMarginCollateralRatio | Cross margin collateral ratio (MARKET_DATA)
[**GetMarginCrossMarginDataV1**](MarginTradingAPI.md#GetMarginCrossMarginDataV1) | **Get** /sapi/v1/margin/crossMarginData | Query Cross Margin Fee Data (USER_DATA)
[**GetMarginDelistScheduleV1**](MarginTradingAPI.md#GetMarginDelistScheduleV1) | **Get** /sapi/v1/margin/delist-schedule | Get Delist Schedule (MARKET_DATA)
[**GetMarginExchangeSmallLiabilityHistoryV1**](MarginTradingAPI.md#GetMarginExchangeSmallLiabilityHistoryV1) | **Get** /sapi/v1/margin/exchange-small-liability-history | Get Small Liability Exchange History (USER_DATA)
[**GetMarginExchangeSmallLiabilityV1**](MarginTradingAPI.md#GetMarginExchangeSmallLiabilityV1) | **Get** /sapi/v1/margin/exchange-small-liability | Get Small Liability Exchange Coin List (USER_DATA)
[**GetMarginForceLiquidationRecV1**](MarginTradingAPI.md#GetMarginForceLiquidationRecV1) | **Get** /sapi/v1/margin/forceLiquidationRec | Get Force Liquidation Record (USER_DATA)
[**GetMarginInterestHistoryV1**](MarginTradingAPI.md#GetMarginInterestHistoryV1) | **Get** /sapi/v1/margin/interestHistory | Get Interest History (USER_DATA)
[**GetMarginInterestRateHistoryV1**](MarginTradingAPI.md#GetMarginInterestRateHistoryV1) | **Get** /sapi/v1/margin/interestRateHistory | Query Margin Interest Rate History (USER_DATA)
[**GetMarginIsolatedAccountLimitV1**](MarginTradingAPI.md#GetMarginIsolatedAccountLimitV1) | **Get** /sapi/v1/margin/isolated/accountLimit | Query Enabled Isolated Margin Account Limit (USER_DATA)
[**GetMarginIsolatedAccountV1**](MarginTradingAPI.md#GetMarginIsolatedAccountV1) | **Get** /sapi/v1/margin/isolated/account | Query Isolated Margin Account Info (USER_DATA)
[**GetMarginIsolatedAllPairsV1**](MarginTradingAPI.md#GetMarginIsolatedAllPairsV1) | **Get** /sapi/v1/margin/isolated/allPairs | Get All Isolated Margin Symbol(MARKET_DATA)
[**GetMarginIsolatedMarginDataV1**](MarginTradingAPI.md#GetMarginIsolatedMarginDataV1) | **Get** /sapi/v1/margin/isolatedMarginData | Query Isolated Margin Fee Data (USER_DATA)
[**GetMarginIsolatedMarginTierV1**](MarginTradingAPI.md#GetMarginIsolatedMarginTierV1) | **Get** /sapi/v1/margin/isolatedMarginTier | Query Isolated Margin Tier Data (USER_DATA)
[**GetMarginLeverageBracketV1**](MarginTradingAPI.md#GetMarginLeverageBracketV1) | **Get** /sapi/v1/margin/leverageBracket | Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)
[**GetMarginMaxBorrowableV1**](MarginTradingAPI.md#GetMarginMaxBorrowableV1) | **Get** /sapi/v1/margin/maxBorrowable | Query Max Borrow (USER_DATA)
[**GetMarginMaxTransferableV1**](MarginTradingAPI.md#GetMarginMaxTransferableV1) | **Get** /sapi/v1/margin/maxTransferable | Query Max Transfer-Out Amount (USER_DATA)
[**GetMarginMyTradesV1**](MarginTradingAPI.md#GetMarginMyTradesV1) | **Get** /sapi/v1/margin/myTrades | Query Margin Account&#39;s Trade List (USER_DATA)
[**GetMarginNextHourlyInterestRateV1**](MarginTradingAPI.md#GetMarginNextHourlyInterestRateV1) | **Get** /sapi/v1/margin/next-hourly-interest-rate | Get future hourly interest rate (USER_DATA)
[**GetMarginOpenOrderListV1**](MarginTradingAPI.md#GetMarginOpenOrderListV1) | **Get** /sapi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**GetMarginOpenOrdersV1**](MarginTradingAPI.md#GetMarginOpenOrdersV1) | **Get** /sapi/v1/margin/openOrders | Query Margin Account&#39;s Open Orders (USER_DATA)
[**GetMarginOrderListV1**](MarginTradingAPI.md#GetMarginOrderListV1) | **Get** /sapi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**GetMarginOrderV1**](MarginTradingAPI.md#GetMarginOrderV1) | **Get** /sapi/v1/margin/order | Query Margin Account&#39;s Order (USER_DATA)
[**GetMarginPriceIndexV1**](MarginTradingAPI.md#GetMarginPriceIndexV1) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)
[**GetMarginRateLimitOrderV1**](MarginTradingAPI.md#GetMarginRateLimitOrderV1) | **Get** /sapi/v1/margin/rateLimit/order | Query Current Margin Order Count Usage (TRADE)
[**GetMarginTradeCoeffV1**](MarginTradingAPI.md#GetMarginTradeCoeffV1) | **Get** /sapi/v1/margin/tradeCoeff | Get Summary of Margin account (USER_DATA)
[**GetMarginTransferV1**](MarginTradingAPI.md#GetMarginTransferV1) | **Get** /sapi/v1/margin/transfer | Get Cross Margin Transfer History (USER_DATA)
[**UpdateMarginApiKeyIpV1**](MarginTradingAPI.md#UpdateMarginApiKeyIpV1) | **Put** /sapi/v1/margin/apiKey/ip | Edit ip for Special Key(Low-Latency Trading)(TRADE)
[**UpdateMarginListenKeyV1**](MarginTradingAPI.md#UpdateMarginListenKeyV1) | **Put** /sapi/v1/margin/listen-key | Keepalive User Data Stream (USER_STREAM)
[**UpdateUserDataStreamIsolatedV1**](MarginTradingAPI.md#UpdateUserDataStreamIsolatedV1) | **Put** /sapi/v1/userDataStream/isolated | Keepalive Isolated Margin User Data Stream (USER_STREAM)
[**UpdateUserDataStreamV1**](MarginTradingAPI.md#UpdateUserDataStreamV1) | **Put** /sapi/v1/userDataStream | Keepalive Margin User Data Stream (USER_STREAM)



## CreateMarginApiKeyV1

> CreateMarginApiKeyV1Resp CreateMarginApiKeyV1(ctx).ApiName(apiName).Timestamp(timestamp).Ip(ip).PermissionMode(permissionMode).PublicKey(publicKey).RecvWindow(recvWindow).Symbol(symbol).Execute()

Create Special Key(Low-Latency Trading)(TRADE)



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
	apiName := "apiName_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ip := "ip_example" // string |  (optional) (default to "")
	permissionMode := "permissionMode_example" // string |  (optional) (default to "")
	publicKey := "publicKey_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginApiKeyV1(context.Background()).ApiName(apiName).Timestamp(timestamp).Ip(ip).PermissionMode(permissionMode).PublicKey(publicKey).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginApiKeyV1`: CreateMarginApiKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginApiKeyV1Request struct via the builder pattern


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

[**CreateMarginApiKeyV1Resp**](CreateMarginApiKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginBorrowRepayV1

> CreateMarginBorrowRepayV1Resp CreateMarginBorrowRepayV1(ctx).Amount(amount).Asset(asset).IsIsolated(isIsolated).Symbol(symbol).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()

Margin account borrow/repay(MARGIN)



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
	isIsolated := "isIsolated_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginBorrowRepayV1(context.Background()).Amount(amount).Asset(asset).IsIsolated(isIsolated).Symbol(symbol).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginBorrowRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginBorrowRepayV1`: CreateMarginBorrowRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginBorrowRepayV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginBorrowRepayV1Request struct via the builder pattern


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

[**CreateMarginBorrowRepayV1Resp**](CreateMarginBorrowRepayV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginExchangeSmallLiabilityV1

> MarginCreateMarginExchangeSmallLiabilityV1Resp CreateMarginExchangeSmallLiabilityV1(ctx).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Small Liability Exchange (MARGIN)



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
	assetNames := []string{"Inner_example"} // []string | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginExchangeSmallLiabilityV1(context.Background()).AssetNames(assetNames).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginExchangeSmallLiabilityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginExchangeSmallLiabilityV1`: MarginCreateMarginExchangeSmallLiabilityV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginExchangeSmallLiabilityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginExchangeSmallLiabilityV1Request struct via the builder pattern


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


## CreateMarginIsolatedAccountV1

> CreateMarginIsolatedAccountV1Resp CreateMarginIsolatedAccountV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Isolated Margin Account (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginIsolatedAccountV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginIsolatedAccountV1`: CreateMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateMarginIsolatedAccountV1Resp**](CreateMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginListenKeyV1

> CreateMarginListenKeyV1Resp CreateMarginListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginListenKeyV1`: CreateMarginListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginListenKeyV1Request struct via the builder pattern


### Return type

[**CreateMarginListenKeyV1Resp**](CreateMarginListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginManualLiquidationV1

> CreateMarginManualLiquidationV1Resp CreateMarginManualLiquidationV1(ctx).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Symbol(symbol).Execute()

Margin Manual Liquidation(MARGIN)



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
	type_ := "type__example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginManualLiquidationV1(context.Background()).Timestamp(timestamp).Type_(type_).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginManualLiquidationV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginManualLiquidationV1`: CreateMarginManualLiquidationV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginManualLiquidationV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginManualLiquidationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateMarginManualLiquidationV1Resp**](CreateMarginManualLiquidationV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginMaxLeverageV1

> CreateMarginMaxLeverageV1Resp CreateMarginMaxLeverageV1(ctx).MaxLeverage(maxLeverage).Execute()

Adjust cross margin max leverage (USER_DATA)



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
	maxLeverage := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginMaxLeverageV1(context.Background()).MaxLeverage(maxLeverage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginMaxLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginMaxLeverageV1`: CreateMarginMaxLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginMaxLeverageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginMaxLeverageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxLeverage** | **int32** |  | 

### Return type

[**CreateMarginMaxLeverageV1Resp**](CreateMarginMaxLeverageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginOrderOcoV1

> CreateMarginOrderOcoV1Resp CreateMarginOrderOcoV1(ctx).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()

Margin Account New OCO (TRADE)



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
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginOrderOcoV1(context.Background()).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginOrderOcoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginOrderOcoV1`: CreateMarginOrderOcoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginOrderOcoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginOrderOcoV1Request struct via the builder pattern


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

[**CreateMarginOrderOcoV1Resp**](CreateMarginOrderOcoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginOrderOtoV1

> CreateMarginOrderOtoV1Resp CreateMarginOrderOtoV1(ctx).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).WorkingIcebergQty(workingIcebergQty).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).Execute()

Margin Account New OTO (TRADE)



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
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginOrderOtoV1(context.Background()).PendingQuantity(pendingQuantity).PendingSide(pendingSide).PendingType(pendingType).Symbol(symbol).WorkingIcebergQty(workingIcebergQty).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingClientOrderId(pendingClientOrderId).PendingIcebergQty(pendingIcebergQty).PendingPrice(pendingPrice).PendingStopPrice(pendingStopPrice).PendingTimeInForce(pendingTimeInForce).PendingTrailingDelta(pendingTrailingDelta).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginOrderOtoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginOrderOtoV1`: CreateMarginOrderOtoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginOrderOtoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginOrderOtoV1Request struct via the builder pattern


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

[**CreateMarginOrderOtoV1Resp**](CreateMarginOrderOtoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginOrderOtocoV1

> CreateMarginOrderOtocoV1Resp CreateMarginOrderOtocoV1(ctx).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).Execute()

Margin Account New OTOCO (TRADE)



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
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginOrderOtocoV1(context.Background()).PendingAboveType(pendingAboveType).PendingQuantity(pendingQuantity).PendingSide(pendingSide).Symbol(symbol).WorkingPrice(workingPrice).WorkingQuantity(workingQuantity).WorkingSide(workingSide).WorkingType(workingType).AutoRepayAtCancel(autoRepayAtCancel).IsIsolated(isIsolated).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).PendingAboveClientOrderId(pendingAboveClientOrderId).PendingAboveIcebergQty(pendingAboveIcebergQty).PendingAbovePrice(pendingAbovePrice).PendingAboveStopPrice(pendingAboveStopPrice).PendingAboveTimeInForce(pendingAboveTimeInForce).PendingAboveTrailingDelta(pendingAboveTrailingDelta).PendingBelowClientOrderId(pendingBelowClientOrderId).PendingBelowIcebergQty(pendingBelowIcebergQty).PendingBelowPrice(pendingBelowPrice).PendingBelowStopPrice(pendingBelowStopPrice).PendingBelowTimeInForce(pendingBelowTimeInForce).PendingBelowTrailingDelta(pendingBelowTrailingDelta).PendingBelowType(pendingBelowType).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).WorkingClientOrderId(workingClientOrderId).WorkingIcebergQty(workingIcebergQty).WorkingTimeInForce(workingTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginOrderOtocoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginOrderOtocoV1`: CreateMarginOrderOtocoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginOrderOtocoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginOrderOtocoV1Request struct via the builder pattern


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

[**CreateMarginOrderOtocoV1Resp**](CreateMarginOrderOtocoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMarginOrderV1

> MarginCreateMarginOrderV1Resp CreateMarginOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()

Margin Account New Order (TRADE)



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
	resp, r, err := apiClient.MarginTradingAPI.CreateMarginOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).IsIsolated(isIsolated).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarginOrderV1`: MarginCreateMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarginOrderV1Request struct via the builder pattern


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


## CreateUserDataStreamIsolatedV1

> CreateUserDataStreamIsolatedV1Resp CreateUserDataStreamIsolatedV1(ctx).Symbol(symbol).Execute()

Start Isolated Margin User Data Stream (USER_STREAM)



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
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.CreateUserDataStreamIsolatedV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserDataStreamIsolatedV1`: CreateUserDataStreamIsolatedV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserDataStreamIsolatedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateUserDataStreamIsolatedV1Resp**](CreateUserDataStreamIsolatedV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserDataStreamV1

> CreateUserDataStreamV1Resp CreateUserDataStreamV1(ctx).Execute()

Start Margin User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.MarginTradingAPI.CreateUserDataStreamV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.CreateUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserDataStreamV1`: CreateUserDataStreamV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.CreateUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserDataStreamV1Request struct via the builder pattern


### Return type

[**CreateUserDataStreamV1Resp**](CreateUserDataStreamV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMarginApiKeyV1

> map[string]interface{} DeleteMarginApiKeyV1(ctx).Timestamp(timestamp).ApiKey(apiKey).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()

Delete Special Key(Low-Latency Trading)(TRADE)



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
	apiKey := "apiKey_example" // string |  (optional) (default to "")
	apiName := "apiName_example" // string |  (optional) (default to "")
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.DeleteMarginApiKeyV1(context.Background()).Timestamp(timestamp).ApiKey(apiKey).ApiName(apiName).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarginApiKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarginApiKeyV1Request struct via the builder pattern


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


## DeleteMarginIsolatedAccountV1

> DeleteMarginIsolatedAccountV1Resp DeleteMarginIsolatedAccountV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Disable Isolated Margin Account (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.DeleteMarginIsolatedAccountV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarginIsolatedAccountV1`: DeleteMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**DeleteMarginIsolatedAccountV1Resp**](DeleteMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMarginListenKeyV1

> map[string]interface{} DeleteMarginListenKeyV1(ctx).Execute()

Close User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.MarginTradingAPI.DeleteMarginListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarginListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarginListenKeyV1Request struct via the builder pattern


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


## DeleteMarginOpenOrdersV1

> []DeleteMarginOpenOrdersV1RespItem DeleteMarginOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Margin Account Cancel all Open Orders on a Symbol (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.DeleteMarginOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarginOpenOrdersV1`: []DeleteMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]DeleteMarginOpenOrdersV1RespItem**](DeleteMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMarginOrderListV1

> DeleteMarginOrderListV1Resp DeleteMarginOrderListV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel OCO (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderListId := int64(789) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "listClientOrderId_example" // string | Either `orderListId` or `listClientOrderId` must be provided (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.DeleteMarginOrderListV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarginOrderListV1`: DeleteMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarginOrderListV1Request struct via the builder pattern


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

[**DeleteMarginOrderListV1Resp**](DeleteMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMarginOrderV1

> DeleteMarginOrderV1Resp DeleteMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Margin Account Cancel Order (TRADE)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.DeleteMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarginOrderV1`: DeleteMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarginOrderV1Request struct via the builder pattern


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

[**DeleteMarginOrderV1Resp**](DeleteMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserDataStreamIsolatedV1

> map[string]interface{} DeleteUserDataStreamIsolatedV1(ctx).Symbol(symbol).Listenkey(listenkey).Execute()

Close Isolated Margin User Data Stream (USER_STREAM)



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
	symbol := "symbol_example" // string |  (default to "")
	listenkey := "listenkey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.DeleteUserDataStreamIsolatedV1(context.Background()).Symbol(symbol).Listenkey(listenkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserDataStreamIsolatedV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDataStreamIsolatedV1Request struct via the builder pattern


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


## DeleteUserDataStreamV1

> map[string]interface{} DeleteUserDataStreamV1(ctx).Listenkey(listenkey).Execute()

Close Margin User Data Stream (USER_STREAM)



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
	listenkey := "listenkey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.DeleteUserDataStreamV1(context.Background()).Listenkey(listenkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.DeleteUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserDataStreamV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.DeleteUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDataStreamV1Request struct via the builder pattern


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


## GetBnbBurnV1

> GetBnbBurnV1Resp GetBnbBurnV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get BNB Burn Status (USER_DATA)



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
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetBnbBurnV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetBnbBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBnbBurnV1`: GetBnbBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetBnbBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBnbBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetBnbBurnV1Resp**](GetBnbBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginAccountV1

> GetMarginAccountV1Resp GetMarginAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Cross Margin Account Details (USER_DATA)



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
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginAccountV1`: GetMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetMarginAccountV1Resp**](GetMarginAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginAllAssetsV1

> []GetMarginAllAssetsV1RespItem GetMarginAllAssetsV1(ctx).Asset(asset).Execute()

Get All Margin Assets (MARKET_DATA)



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
	asset := "asset_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginAllAssetsV1(context.Background()).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginAllAssetsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginAllAssetsV1`: []GetMarginAllAssetsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginAllAssetsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginAllAssetsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]GetMarginAllAssetsV1RespItem**](GetMarginAllAssetsV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginAllOrderListV1

> []GetMarginAllOrderListV1RespItem GetMarginAllOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	fromId := int64(789) // int64 | If supplied, neither `startTime` or `endTime` can be provided (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default Value: 500; Max Value: 1000 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginAllOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginAllOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginAllOrderListV1`: []GetMarginAllOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginAllOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginAllOrderListV1Request struct via the builder pattern


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

[**[]GetMarginAllOrderListV1RespItem**](GetMarginAllOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginAllOrdersV1

> []GetMarginAllOrdersV1RespItem GetMarginAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's All Orders (USER_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginAllOrdersV1`: []GetMarginAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginAllOrdersV1Request struct via the builder pattern


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

[**[]GetMarginAllOrdersV1RespItem**](GetMarginAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginAllPairsV1

> []GetMarginAllPairsV1RespItem GetMarginAllPairsV1(ctx).Symbol(symbol).Execute()

Get All Cross Margin Pairs (MARKET_DATA)



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
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginAllPairsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginAllPairsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginAllPairsV1`: []GetMarginAllPairsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginAllPairsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginAllPairsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]GetMarginAllPairsV1RespItem**](GetMarginAllPairsV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginApiKeyListV1

> []GetMarginApiKeyListV1RespItem GetMarginApiKeyListV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key List(Low Latency Trading)(TRADE)



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
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginApiKeyListV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginApiKeyListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginApiKeyListV1`: []GetMarginApiKeyListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginApiKeyListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginApiKeyListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetMarginApiKeyListV1RespItem**](GetMarginApiKeyListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginApiKeyV1

> GetMarginApiKeyV1Resp GetMarginApiKeyV1(ctx).ApiKey(apiKey).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Special key(Low Latency Trading)(TRADE)



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
	apiKey := "apiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string | isolated margin pair (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginApiKeyV1(context.Background()).ApiKey(apiKey).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginApiKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginApiKeyV1`: GetMarginApiKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginApiKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **symbol** | **string** | isolated margin pair | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetMarginApiKeyV1Resp**](GetMarginApiKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginAvailableInventoryV1

> MarginGetMarginAvailableInventoryV1Resp GetMarginAvailableInventoryV1(ctx).Type_(type_).Execute()

Query Margin Available Inventory(USER_DATA)



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
	type_ := "type__example" // string | MARGIN,ISOLATED (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginAvailableInventoryV1(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginAvailableInventoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginAvailableInventoryV1`: MarginGetMarginAvailableInventoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginAvailableInventoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginAvailableInventoryV1Request struct via the builder pattern


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


## GetMarginBorrowRepayV1

> GetMarginBorrowRepayV1Resp GetMarginBorrowRepayV1(ctx).Type_(type_).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query borrow/repay records in Margin account(USER_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginBorrowRepayV1(context.Background()).Type_(type_).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginBorrowRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginBorrowRepayV1`: GetMarginBorrowRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginBorrowRepayV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginBorrowRepayV1Request struct via the builder pattern


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

[**GetMarginBorrowRepayV1Resp**](GetMarginBorrowRepayV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginCapitalFlowV1

> []GetMarginCapitalFlowV1RespItem GetMarginCapitalFlowV1(ctx).Timestamp(timestamp).Asset(asset).Symbol(symbol).Type_(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Cross Isolated Margin Capital Flow (USER_DATA)



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
	symbol := "symbol_example" // string | 查询逐仓数据时必填 (optional) (default to "")
	type_ := "type__example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | 只支持查询最近90天的数据 (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | 如设置fromId, 将返回id &gt; fromId的数据。否则将返回最新数据 (optional)
	limit := int64(789) // int64 | 每次返回的数据条数限制。默认 500; 最大 1000. (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginCapitalFlowV1(context.Background()).Timestamp(timestamp).Asset(asset).Symbol(symbol).Type_(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginCapitalFlowV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginCapitalFlowV1`: []GetMarginCapitalFlowV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginCapitalFlowV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginCapitalFlowV1Request struct via the builder pattern


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

[**[]GetMarginCapitalFlowV1RespItem**](GetMarginCapitalFlowV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginCrossMarginCollateralRatioV1

> []GetMarginCrossMarginCollateralRatioV1RespItem GetMarginCrossMarginCollateralRatioV1(ctx).Execute()

Cross margin collateral ratio (MARKET_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginCrossMarginCollateralRatioV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginCrossMarginCollateralRatioV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginCrossMarginCollateralRatioV1`: []GetMarginCrossMarginCollateralRatioV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginCrossMarginCollateralRatioV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginCrossMarginCollateralRatioV1Request struct via the builder pattern


### Return type

[**[]GetMarginCrossMarginCollateralRatioV1RespItem**](GetMarginCrossMarginCollateralRatioV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginCrossMarginDataV1

> []GetMarginCrossMarginDataV1RespItem GetMarginCrossMarginDataV1(ctx).Timestamp(timestamp).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()

Query Cross Margin Fee Data (USER_DATA)



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
	vipLevel := int32(56) // int32 | User&#39;s current specific margin data will be returned if vipLevel is omitted (optional)
	coin := "coin_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginCrossMarginDataV1(context.Background()).Timestamp(timestamp).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginCrossMarginDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginCrossMarginDataV1`: []GetMarginCrossMarginDataV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginCrossMarginDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginCrossMarginDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | User&amp;#39;s current specific margin data will be returned if vipLevel is omitted | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]GetMarginCrossMarginDataV1RespItem**](GetMarginCrossMarginDataV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginDelistScheduleV1

> []GetMarginDelistScheduleV1RespItem GetMarginDelistScheduleV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Delist Schedule (MARKET_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginDelistScheduleV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginDelistScheduleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginDelistScheduleV1`: []GetMarginDelistScheduleV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginDelistScheduleV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginDelistScheduleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetMarginDelistScheduleV1RespItem**](GetMarginDelistScheduleV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginExchangeSmallLiabilityHistoryV1

> GetMarginExchangeSmallLiabilityHistoryV1Resp GetMarginExchangeSmallLiabilityHistoryV1(ctx).Current(current).Size(size).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange History (USER_DATA)



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
	current := int32(56) // int32 | Currently querying page. Start from 1. Default:1
	size := int32(56) // int32 | Default:10, Max:100
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 | Default: 30 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginExchangeSmallLiabilityHistoryV1(context.Background()).Current(current).Size(size).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginExchangeSmallLiabilityHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginExchangeSmallLiabilityHistoryV1`: GetMarginExchangeSmallLiabilityHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginExchangeSmallLiabilityHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginExchangeSmallLiabilityHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **current** | **int32** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int32** | Default:10, Max:100 | 
 **timestamp** | **int64** |  | 
 **startTime** | **int64** | Default: 30 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMarginExchangeSmallLiabilityHistoryV1Resp**](GetMarginExchangeSmallLiabilityHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginExchangeSmallLiabilityV1

> []GetMarginExchangeSmallLiabilityV1RespItem GetMarginExchangeSmallLiabilityV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Small Liability Exchange Coin List (USER_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginExchangeSmallLiabilityV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginExchangeSmallLiabilityV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginExchangeSmallLiabilityV1`: []GetMarginExchangeSmallLiabilityV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginExchangeSmallLiabilityV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginExchangeSmallLiabilityV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetMarginExchangeSmallLiabilityV1RespItem**](GetMarginExchangeSmallLiabilityV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginForceLiquidationRecV1

> GetMarginForceLiquidationRecV1Resp GetMarginForceLiquidationRecV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Force Liquidation Record (USER_DATA)



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
	isolatedSymbol := "isolatedSymbol_example" // string |  (optional) (default to "")
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginForceLiquidationRecV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).IsolatedSymbol(isolatedSymbol).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginForceLiquidationRecV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginForceLiquidationRecV1`: GetMarginForceLiquidationRecV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginForceLiquidationRecV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginForceLiquidationRecV1Request struct via the builder pattern


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

[**GetMarginForceLiquidationRecV1Resp**](GetMarginForceLiquidationRecV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginInterestHistoryV1

> GetMarginInterestHistoryV1Resp GetMarginInterestHistoryV1(ctx).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Interest History (USER_DATA)



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
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).IsolatedSymbol(isolatedSymbol).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginInterestHistoryV1`: GetMarginInterestHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginInterestHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginInterestHistoryV1Request struct via the builder pattern


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

[**GetMarginInterestHistoryV1Resp**](GetMarginInterestHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginInterestRateHistoryV1

> []GetMarginInterestRateHistoryV1RespItem GetMarginInterestRateHistoryV1(ctx).Asset(asset).Timestamp(timestamp).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Query Margin Interest Rate History (USER_DATA)



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
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	vipLevel := int32(56) // int32 | Default: user&#39;s vip level (optional)
	startTime := int64(789) // int64 | Default: 7 days ago (optional)
	endTime := int64(789) // int64 | Default: present. Maximum range: 1 months. (optional)
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginInterestRateHistoryV1(context.Background()).Asset(asset).Timestamp(timestamp).VipLevel(vipLevel).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginInterestRateHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginInterestRateHistoryV1`: []GetMarginInterestRateHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginInterestRateHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginInterestRateHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | Default: user&amp;#39;s vip level | 
 **startTime** | **int64** | Default: 7 days ago | 
 **endTime** | **int64** | Default: present. Maximum range: 1 months. | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**[]GetMarginInterestRateHistoryV1RespItem**](GetMarginInterestRateHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginIsolatedAccountLimitV1

> GetMarginIsolatedAccountLimitV1Resp GetMarginIsolatedAccountLimitV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Enabled Isolated Margin Account Limit (USER_DATA)



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
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginIsolatedAccountLimitV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginIsolatedAccountLimitV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginIsolatedAccountLimitV1`: GetMarginIsolatedAccountLimitV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginIsolatedAccountLimitV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginIsolatedAccountLimitV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetMarginIsolatedAccountLimitV1Resp**](GetMarginIsolatedAccountLimitV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginIsolatedAccountV1

> GetMarginIsolatedAccountV1Resp GetMarginIsolatedAccountV1(ctx).Timestamp(timestamp).Symbols(symbols).RecvWindow(recvWindow).Execute()

Query Isolated Margin Account Info (USER_DATA)



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
	symbols := "symbols_example" // string | Max 5 symbols can be sent; separated by &#34;,&#34;. e.g. &#34;BTCUSDT,BNBUSDT,ADAUSDT&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginIsolatedAccountV1(context.Background()).Timestamp(timestamp).Symbols(symbols).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginIsolatedAccountV1`: GetMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbols** | **string** | Max 5 symbols can be sent; separated by &amp;#34;,&amp;#34;. e.g. &amp;#34;BTCUSDT,BNBUSDT,ADAUSDT&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**GetMarginIsolatedAccountV1Resp**](GetMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginIsolatedAllPairsV1

> []GetMarginIsolatedAllPairsV1RespItem GetMarginIsolatedAllPairsV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Get All Isolated Margin Symbol(MARKET_DATA)



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
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginIsolatedAllPairsV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginIsolatedAllPairsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginIsolatedAllPairsV1`: []GetMarginIsolatedAllPairsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginIsolatedAllPairsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginIsolatedAllPairsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**[]GetMarginIsolatedAllPairsV1RespItem**](GetMarginIsolatedAllPairsV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginIsolatedMarginDataV1

> []GetMarginIsolatedMarginDataV1RespItem GetMarginIsolatedMarginDataV1(ctx).Timestamp(timestamp).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Isolated Margin Fee Data (USER_DATA)



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
	vipLevel := int32(56) // int32 | User&#39;s current specific margin data will be returned if vipLevel is omitted (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginIsolatedMarginDataV1(context.Background()).Timestamp(timestamp).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginIsolatedMarginDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginIsolatedMarginDataV1`: []GetMarginIsolatedMarginDataV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginIsolatedMarginDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginIsolatedMarginDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | User&amp;#39;s current specific margin data will be returned if vipLevel is omitted | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]GetMarginIsolatedMarginDataV1RespItem**](GetMarginIsolatedMarginDataV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginIsolatedMarginTierV1

> []GetMarginIsolatedMarginTierV1RespItem GetMarginIsolatedMarginTierV1(ctx).Symbol(symbol).Timestamp(timestamp).Tier(tier).RecvWindow(recvWindow).Execute()

Query Isolated Margin Tier Data (USER_DATA)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	tier := int32(56) // int32 | All margin tier data will be returned if tier is omitted (optional)
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginIsolatedMarginTierV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Tier(tier).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginIsolatedMarginTierV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginIsolatedMarginTierV1`: []GetMarginIsolatedMarginTierV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginIsolatedMarginTierV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginIsolatedMarginTierV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **tier** | **int32** | All margin tier data will be returned if tier is omitted | 
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]GetMarginIsolatedMarginTierV1RespItem**](GetMarginIsolatedMarginTierV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginLeverageBracketV1

> []GetMarginLeverageBracketV1RespItem GetMarginLeverageBracketV1(ctx).Execute()

Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginLeverageBracketV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginLeverageBracketV1`: []GetMarginLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginLeverageBracketV1Request struct via the builder pattern


### Return type

[**[]GetMarginLeverageBracketV1RespItem**](GetMarginLeverageBracketV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginMaxBorrowableV1

> GetMarginMaxBorrowableV1Resp GetMarginMaxBorrowableV1(ctx).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Borrow (USER_DATA)



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
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginMaxBorrowableV1(context.Background()).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginMaxBorrowableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginMaxBorrowableV1`: GetMarginMaxBorrowableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginMaxBorrowableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginMaxBorrowableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isolatedSymbol** | **string** | isolated symbol | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetMarginMaxBorrowableV1Resp**](GetMarginMaxBorrowableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginMaxTransferableV1

> GetMarginMaxTransferableV1Resp GetMarginMaxTransferableV1(ctx).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Query Max Transfer-Out Amount (USER_DATA)



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
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isolatedSymbol := "isolatedSymbol_example" // string | isolated symbol (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginMaxTransferableV1(context.Background()).Asset(asset).Timestamp(timestamp).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginMaxTransferableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginMaxTransferableV1`: GetMarginMaxTransferableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginMaxTransferableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginMaxTransferableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isolatedSymbol** | **string** | isolated symbol | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetMarginMaxTransferableV1Resp**](GetMarginMaxTransferableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginMyTradesV1

> []GetMarginMyTradesV1RespItem GetMarginMyTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's Trade List (USER_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginMyTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginMyTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginMyTradesV1`: []GetMarginMyTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginMyTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginMyTradesV1Request struct via the builder pattern


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

[**[]GetMarginMyTradesV1RespItem**](GetMarginMyTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginNextHourlyInterestRateV1

> []GetMarginNextHourlyInterestRateV1RespItem GetMarginNextHourlyInterestRateV1(ctx).Assets(assets).IsIsolated(isIsolated).Execute()

Get future hourly interest rate (USER_DATA)



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
	assets := "assets_example" // string | List of assets, separated by commas, up to 20 (default to "")
	isIsolated := true // bool | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginNextHourlyInterestRateV1(context.Background()).Assets(assets).IsIsolated(isIsolated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginNextHourlyInterestRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginNextHourlyInterestRateV1`: []GetMarginNextHourlyInterestRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginNextHourlyInterestRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginNextHourlyInterestRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assets** | **string** | List of assets, separated by commas, up to 20 | [default to &quot;&quot;]
 **isIsolated** | **bool** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34; | 

### Return type

[**[]GetMarginNextHourlyInterestRateV1RespItem**](GetMarginNextHourlyInterestRateV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginOpenOrderListV1

> []GetMarginOpenOrderListV1RespItem GetMarginOpenOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginOpenOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginOpenOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginOpenOrderListV1`: []GetMarginOpenOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginOpenOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginOpenOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetMarginOpenOrderListV1RespItem**](GetMarginOpenOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginOpenOrdersV1

> []MarginGetMarginOpenOrdersV1RespItem GetMarginOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()

Query Margin Account's Open Orders (USER_DATA)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IsIsolated(isIsolated).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginOpenOrdersV1`: []MarginGetMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginOpenOrdersV1Request struct via the builder pattern


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


## GetMarginOrderListV1

> GetMarginOrderListV1Resp GetMarginOrderListV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | mandatory for isolated margin, not supported for cross margin (optional) (default to "")
	orderListId := int64(789) // int64 | Either `orderListId` or `origClientOrderId` must be provided (optional)
	origClientOrderId := "origClientOrderId_example" // string | Either `orderListId` or `origClientOrderId` must be provided (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginOrderListV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginOrderListV1`: GetMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | mandatory for isolated margin, not supported for cross margin | [default to &quot;&quot;]
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided | 
 **origClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;origClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetMarginOrderListV1Resp**](GetMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginOrderV1

> GetMarginOrderV1Resp GetMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's Order (USER_DATA)



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
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).IsIsolated(isIsolated).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginOrderV1`: GetMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetMarginOrderV1Resp**](GetMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginPriceIndexV1

> GetMarginPriceIndexV1Resp GetMarginPriceIndexV1(ctx).Symbol(symbol).Execute()

Query Margin PriceIndex (MARKET_DATA)



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
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginPriceIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginPriceIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginPriceIndexV1`: GetMarginPriceIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginPriceIndexV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginPriceIndexV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**GetMarginPriceIndexV1Resp**](GetMarginPriceIndexV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginRateLimitOrderV1

> []GetMarginRateLimitOrderV1RespItem GetMarginRateLimitOrderV1(ctx).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Current Margin Order Count Usage (TRADE)



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
	isIsolated := "isIsolated_example" // string | for isolated margin or not, &#34;TRUE&#34;, &#34;FALSE&#34;，default &#34;FALSE&#34; (optional) (default to "")
	symbol := "symbol_example" // string | isolated symbol, mandatory for isolated margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginRateLimitOrderV1(context.Background()).Timestamp(timestamp).IsIsolated(isIsolated).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginRateLimitOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginRateLimitOrderV1`: []GetMarginRateLimitOrderV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginRateLimitOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginRateLimitOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **isIsolated** | **string** | for isolated margin or not, &amp;#34;TRUE&amp;#34;, &amp;#34;FALSE&amp;#34;，default &amp;#34;FALSE&amp;#34; | [default to &quot;&quot;]
 **symbol** | **string** | isolated symbol, mandatory for isolated margin | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]GetMarginRateLimitOrderV1RespItem**](GetMarginRateLimitOrderV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginTradeCoeffV1

> GetMarginTradeCoeffV1Resp GetMarginTradeCoeffV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Summary of Margin account (USER_DATA)



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
	resp, r, err := apiClient.MarginTradingAPI.GetMarginTradeCoeffV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginTradeCoeffV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginTradeCoeffV1`: GetMarginTradeCoeffV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginTradeCoeffV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginTradeCoeffV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMarginTradeCoeffV1Resp**](GetMarginTradeCoeffV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarginTransferV1

> GetMarginTransferV1Resp GetMarginTransferV1(ctx).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()

Get Cross Margin Transfer History (USER_DATA)



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
	type_ := "type__example" // string | Transfer Type: ROLL_IN, ROLL_OUT (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	isolatedSymbol := "isolatedSymbol_example" // string | Symbol in Isolated Margin (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.GetMarginTransferV1(context.Background()).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Current(current).Size(size).IsolatedSymbol(isolatedSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.GetMarginTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarginTransferV1`: GetMarginTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.GetMarginTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarginTransferV1Request struct via the builder pattern


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

[**GetMarginTransferV1Resp**](GetMarginTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMarginApiKeyIpV1

> map[string]interface{} UpdateMarginApiKeyIpV1(ctx).ApiKey(apiKey).Ip(ip).Timestamp(timestamp).RecvWindow(recvWindow).Symbol(symbol).Execute()

Edit ip for Special Key(Low-Latency Trading)(TRADE)



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
	apiKey := "apiKey_example" // string |  (default to "")
	ip := "ip_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.UpdateMarginApiKeyIpV1(context.Background()).ApiKey(apiKey).Ip(ip).Timestamp(timestamp).RecvWindow(recvWindow).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.UpdateMarginApiKeyIpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMarginApiKeyIpV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.UpdateMarginApiKeyIpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMarginApiKeyIpV1Request struct via the builder pattern


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


## UpdateMarginListenKeyV1

> map[string]interface{} UpdateMarginListenKeyV1(ctx).ListenKey(listenKey).Execute()

Keepalive User Data Stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.UpdateMarginListenKeyV1(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.UpdateMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMarginListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.UpdateMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMarginListenKeyV1Request struct via the builder pattern


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


## UpdateUserDataStreamIsolatedV1

> map[string]interface{} UpdateUserDataStreamIsolatedV1(ctx).ListenKey(listenKey).Symbol(symbol).Execute()

Keepalive Isolated Margin User Data Stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.UpdateUserDataStreamIsolatedV1(context.Background()).ListenKey(listenKey).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.UpdateUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDataStreamIsolatedV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.UpdateUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDataStreamIsolatedV1Request struct via the builder pattern


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


## UpdateUserDataStreamV1

> map[string]interface{} UpdateUserDataStreamV1(ctx).ListenKey(listenKey).Execute()

Keepalive Margin User Data Stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarginTradingAPI.UpdateUserDataStreamV1(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarginTradingAPI.UpdateUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDataStreamV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarginTradingAPI.UpdateUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDataStreamV1Request struct via the builder pattern


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

