# \AccountAPI

All URIs are relative to *https://papi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginCreateAssetCollectionV1**](AccountAPI.md#PmarginCreateAssetCollectionV1) | **Post** /papi/v1/asset-collection | Fund Collection by Asset(TRADE)
[**PmarginCreateAutoCollectionV1**](AccountAPI.md#PmarginCreateAutoCollectionV1) | **Post** /papi/v1/auto-collection | Fund Auto-collection(TRADE)
[**PmarginCreateBnbTransferV1**](AccountAPI.md#PmarginCreateBnbTransferV1) | **Post** /papi/v1/bnb-transfer | BNB transfer (TRADE)
[**PmarginCreateCmLeverageV1**](AccountAPI.md#PmarginCreateCmLeverageV1) | **Post** /papi/v1/cm/leverage | Change CM Initial Leverage (TRADE)
[**PmarginCreateCmPositionSideDualV1**](AccountAPI.md#PmarginCreateCmPositionSideDualV1) | **Post** /papi/v1/cm/positionSide/dual | Change CM Position Mode(TRADE)
[**PmarginCreateRepayFuturesNegativeBalanceV1**](AccountAPI.md#PmarginCreateRepayFuturesNegativeBalanceV1) | **Post** /papi/v1/repay-futures-negative-balance | Repay futures Negative Balance(USER_DATA)
[**PmarginCreateRepayFuturesSwitchV1**](AccountAPI.md#PmarginCreateRepayFuturesSwitchV1) | **Post** /papi/v1/repay-futures-switch | Change Auto-repay-futures Status(TRADE)
[**PmarginCreateUmLeverageV1**](AccountAPI.md#PmarginCreateUmLeverageV1) | **Post** /papi/v1/um/leverage | Change UM Initial Leverage(TRADE)
[**PmarginCreateUmPositionSideDualV1**](AccountAPI.md#PmarginCreateUmPositionSideDualV1) | **Post** /papi/v1/um/positionSide/dual | Change UM Position Mode(TRADE)
[**PmarginGetAccountV1**](AccountAPI.md#PmarginGetAccountV1) | **Get** /papi/v1/account | Account Information(USER_DATA)
[**PmarginGetBalanceV1**](AccountAPI.md#PmarginGetBalanceV1) | **Get** /papi/v1/balance | Account Balance(USER_DATA)
[**PmarginGetCmAccountV1**](AccountAPI.md#PmarginGetCmAccountV1) | **Get** /papi/v1/cm/account | Get CM Account Detail(USER_DATA)
[**PmarginGetCmCommissionRateV1**](AccountAPI.md#PmarginGetCmCommissionRateV1) | **Get** /papi/v1/cm/commissionRate | Get User Commission Rate for CM(USER_DATA)
[**PmarginGetCmIncomeV1**](AccountAPI.md#PmarginGetCmIncomeV1) | **Get** /papi/v1/cm/income | Get CM Income History(USER_DATA)
[**PmarginGetCmLeverageBracketV1**](AccountAPI.md#PmarginGetCmLeverageBracketV1) | **Get** /papi/v1/cm/leverageBracket | CM Notional and Leverage Brackets(USER_DATA)
[**PmarginGetCmPositionRiskV1**](AccountAPI.md#PmarginGetCmPositionRiskV1) | **Get** /papi/v1/cm/positionRisk | Query CM Position Information(USER_DATA)
[**PmarginGetCmPositionSideDualV1**](AccountAPI.md#PmarginGetCmPositionSideDualV1) | **Get** /papi/v1/cm/positionSide/dual | Get CM Current Position Mode(USER_DATA)
[**PmarginGetMarginMarginInterestHistoryV1**](AccountAPI.md#PmarginGetMarginMarginInterestHistoryV1) | **Get** /papi/v1/margin/marginInterestHistory | Get Margin Borrow/Loan Interest History(USER_DATA)
[**PmarginGetMarginMarginLoanV1**](AccountAPI.md#PmarginGetMarginMarginLoanV1) | **Get** /papi/v1/margin/marginLoan | Query Margin Loan Record(USER_DATA)
[**PmarginGetMarginMaxBorrowableV1**](AccountAPI.md#PmarginGetMarginMaxBorrowableV1) | **Get** /papi/v1/margin/maxBorrowable | Margin Max Borrow(USER_DATA)
[**PmarginGetMarginMaxWithdrawV1**](AccountAPI.md#PmarginGetMarginMaxWithdrawV1) | **Get** /papi/v1/margin/maxWithdraw | Query Margin Max Withdraw(USER_DATA)
[**PmarginGetMarginRepayLoanV1**](AccountAPI.md#PmarginGetMarginRepayLoanV1) | **Get** /papi/v1/margin/repayLoan | Query Margin repay Record(USER_DATA)
[**PmarginGetPortfolioInterestHistoryV1**](AccountAPI.md#PmarginGetPortfolioInterestHistoryV1) | **Get** /papi/v1/portfolio/interest-history | Query Portfolio Margin Negative Balance Interest History(USER_DATA)
[**PmarginGetPortfolioNegativeBalanceExchangeRecordV1**](AccountAPI.md#PmarginGetPortfolioNegativeBalanceExchangeRecordV1) | **Get** /papi/v1/portfolio/negative-balance-exchange-record | Query User Negative Balance Auto Exchange Record (USER_DATA)
[**PmarginGetRateLimitOrderV1**](AccountAPI.md#PmarginGetRateLimitOrderV1) | **Get** /papi/v1/rateLimit/order | Query User Rate Limit (USER_DATA)
[**PmarginGetRepayFuturesSwitchV1**](AccountAPI.md#PmarginGetRepayFuturesSwitchV1) | **Get** /papi/v1/repay-futures-switch | Get Auto-repay-futures Status(USER_DATA)
[**PmarginGetUmAccountConfigV1**](AccountAPI.md#PmarginGetUmAccountConfigV1) | **Get** /papi/v1/um/accountConfig | UM Futures Account Configuration(USER_DATA)
[**PmarginGetUmAccountV1**](AccountAPI.md#PmarginGetUmAccountV1) | **Get** /papi/v1/um/account | Get UM Account Detail(USER_DATA)
[**PmarginGetUmAccountV2**](AccountAPI.md#PmarginGetUmAccountV2) | **Get** /papi/v2/um/account | Get UM Account Detail V2(USER_DATA)
[**PmarginGetUmApiTradingStatusV1**](AccountAPI.md#PmarginGetUmApiTradingStatusV1) | **Get** /papi/v1/um/apiTradingStatus | Portfolio Margin UM Trading Quantitative Rules Indicators(USER_DATA)
[**PmarginGetUmCommissionRateV1**](AccountAPI.md#PmarginGetUmCommissionRateV1) | **Get** /papi/v1/um/commissionRate | Get User Commission Rate for UM(USER_DATA)
[**PmarginGetUmIncomeAsynIdV1**](AccountAPI.md#PmarginGetUmIncomeAsynIdV1) | **Get** /papi/v1/um/income/asyn/id | Get UM Futures Transaction Download Link by Id(USER_DATA)
[**PmarginGetUmIncomeAsynV1**](AccountAPI.md#PmarginGetUmIncomeAsynV1) | **Get** /papi/v1/um/income/asyn | Get Download Id For UM Futures Transaction History (USER_DATA)
[**PmarginGetUmIncomeV1**](AccountAPI.md#PmarginGetUmIncomeV1) | **Get** /papi/v1/um/income | Get UM Income History(USER_DATA)
[**PmarginGetUmLeverageBracketV1**](AccountAPI.md#PmarginGetUmLeverageBracketV1) | **Get** /papi/v1/um/leverageBracket | UM Notional and Leverage Brackets (USER_DATA)
[**PmarginGetUmOrderAsynIdV1**](AccountAPI.md#PmarginGetUmOrderAsynIdV1) | **Get** /papi/v1/um/order/asyn/id | Get UM Futures Order Download Link by Id(USER_DATA)
[**PmarginGetUmOrderAsynV1**](AccountAPI.md#PmarginGetUmOrderAsynV1) | **Get** /papi/v1/um/order/asyn | Get Download Id For UM Futures Order History (USER_DATA)
[**PmarginGetUmPositionRiskV1**](AccountAPI.md#PmarginGetUmPositionRiskV1) | **Get** /papi/v1/um/positionRisk | Query UM Position Information(USER_DATA)
[**PmarginGetUmPositionSideDualV1**](AccountAPI.md#PmarginGetUmPositionSideDualV1) | **Get** /papi/v1/um/positionSide/dual | Get UM Current Position Mode(USER_DATA)
[**PmarginGetUmSymbolConfigV1**](AccountAPI.md#PmarginGetUmSymbolConfigV1) | **Get** /papi/v1/um/symbolConfig | UM Futures Symbol Configuration(USER_DATA)
[**PmarginGetUmTradeAsynIdV1**](AccountAPI.md#PmarginGetUmTradeAsynIdV1) | **Get** /papi/v1/um/trade/asyn/id | Get UM Futures Trade Download Link by Id(USER_DATA)
[**PmarginGetUmTradeAsynV1**](AccountAPI.md#PmarginGetUmTradeAsynV1) | **Get** /papi/v1/um/trade/asyn | Get Download Id For UM Futures Trade History (USER_DATA)



## PmarginCreateAssetCollectionV1

> PmarginCreateAssetCollectionV1Resp PmarginCreateAssetCollectionV1(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fund Collection by Asset(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateAssetCollectionV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateAssetCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateAssetCollectionV1`: PmarginCreateAssetCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateAssetCollectionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateAssetCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateAssetCollectionV1Resp**](PmarginCreateAssetCollectionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateAutoCollectionV1

> PmarginCreateAutoCollectionV1Resp PmarginCreateAutoCollectionV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fund Auto-collection(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateAutoCollectionV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateAutoCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateAutoCollectionV1`: PmarginCreateAutoCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateAutoCollectionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateAutoCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateAutoCollectionV1Resp**](PmarginCreateAutoCollectionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateBnbTransferV1

> PmarginCreateBnbTransferV1Resp PmarginCreateBnbTransferV1(ctx).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()

BNB transfer (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	transferSide := "transferSide_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateBnbTransferV1(context.Background()).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateBnbTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateBnbTransferV1`: PmarginCreateBnbTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateBnbTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateBnbTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **transferSide** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateBnbTransferV1Resp**](PmarginCreateBnbTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateCmLeverageV1

> PmarginCreateCmLeverageV1Resp PmarginCreateCmLeverageV1(ctx).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change CM Initial Leverage (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	leverage := int32(56) // int32 | 
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateCmLeverageV1(context.Background()).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateCmLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmLeverageV1`: PmarginCreateCmLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateCmLeverageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateCmLeverageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leverage** | **int32** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateCmLeverageV1Resp**](PmarginCreateCmLeverageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateCmPositionSideDualV1

> PmarginCreateCmPositionSideDualV1Resp PmarginCreateCmPositionSideDualV1(ctx).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change CM Position Mode(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	dualSidePosition := "dualSidePosition_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateCmPositionSideDualV1(context.Background()).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateCmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmPositionSideDualV1`: PmarginCreateCmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateCmPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateCmPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dualSidePosition** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateCmPositionSideDualV1Resp**](PmarginCreateCmPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateRepayFuturesNegativeBalanceV1

> PmarginCreateRepayFuturesNegativeBalanceV1Resp PmarginCreateRepayFuturesNegativeBalanceV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Repay futures Negative Balance(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateRepayFuturesNegativeBalanceV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateRepayFuturesNegativeBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateRepayFuturesNegativeBalanceV1`: PmarginCreateRepayFuturesNegativeBalanceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateRepayFuturesNegativeBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateRepayFuturesNegativeBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateRepayFuturesNegativeBalanceV1Resp**](PmarginCreateRepayFuturesNegativeBalanceV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateRepayFuturesSwitchV1

> PmarginCreateRepayFuturesSwitchV1Resp PmarginCreateRepayFuturesSwitchV1(ctx).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Auto-repay-futures Status(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	autoRepay := "autoRepay_example" // string |  (default to "true")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateRepayFuturesSwitchV1(context.Background()).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateRepayFuturesSwitchV1`: PmarginCreateRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateRepayFuturesSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateRepayFuturesSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoRepay** | **string** |  | [default to &quot;true&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateRepayFuturesSwitchV1Resp**](PmarginCreateRepayFuturesSwitchV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateUmLeverageV1

> PmarginCreateUmLeverageV1Resp PmarginCreateUmLeverageV1(ctx).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change UM Initial Leverage(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	leverage := int32(56) // int32 | 
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateUmLeverageV1(context.Background()).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateUmLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmLeverageV1`: PmarginCreateUmLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateUmLeverageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmLeverageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **leverage** | **int32** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateUmLeverageV1Resp**](PmarginCreateUmLeverageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateUmPositionSideDualV1

> PmarginCreateUmPositionSideDualV1Resp PmarginCreateUmPositionSideDualV1(ctx).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change UM Position Mode(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	dualSidePosition := "dualSidePosition_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginCreateUmPositionSideDualV1(context.Background()).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginCreateUmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmPositionSideDualV1`: PmarginCreateUmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginCreateUmPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dualSidePosition** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateUmPositionSideDualV1Resp**](PmarginCreateUmPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetAccountV1

> PmarginGetAccountV1Resp PmarginGetAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Information(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetAccountV1`: PmarginGetAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetAccountV1Resp**](PmarginGetAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetBalanceV1

> PmarginGetBalanceV1Resp PmarginGetBalanceV1(ctx).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()

Account Balance(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetBalanceV1(context.Background()).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetBalanceV1`: PmarginGetBalanceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetBalanceV1Resp**](PmarginGetBalanceV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmAccountV1

> PmarginGetCmAccountV1Resp PmarginGetCmAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get CM Account Detail(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetCmAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetCmAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmAccountV1`: PmarginGetCmAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetCmAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmAccountV1Resp**](PmarginGetCmAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmCommissionRateV1

> PmarginGetCmCommissionRateV1Resp PmarginGetCmCommissionRateV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get User Commission Rate for CM(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetCmCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetCmCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmCommissionRateV1`: PmarginGetCmCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetCmCommissionRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmCommissionRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmCommissionRateV1Resp**](PmarginGetCmCommissionRateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmIncomeV1

> []PmarginGetCmIncomeV1RespItem PmarginGetCmIncomeV1(ctx).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get CM Income History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	incomeType := "incomeType_example" // string | &#34;TRANSFER&#34;,&#34;WELCOME_BONUS&#34;, &#34;FUNDING_FEE&#34;, &#34;REALIZED_PNL&#34;, &#34;COMMISSION&#34;, &#34;INSURANCE_CLEAR&#34;, and &#34;DELIVERED_SETTELMENT&#34; (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetCmIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetCmIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmIncomeV1`: []PmarginGetCmIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetCmIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **incomeType** | **string** | &amp;#34;TRANSFER&amp;#34;,&amp;#34;WELCOME_BONUS&amp;#34;, &amp;#34;FUNDING_FEE&amp;#34;, &amp;#34;REALIZED_PNL&amp;#34;, &amp;#34;COMMISSION&amp;#34;, &amp;#34;INSURANCE_CLEAR&amp;#34;, and &amp;#34;DELIVERED_SETTELMENT&amp;#34; | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int32** |  | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmIncomeV1RespItem**](PmarginGetCmIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmLeverageBracketV1

> []PmarginGetCmLeverageBracketV1RespItem PmarginGetCmLeverageBracketV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

CM Notional and Leverage Brackets(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetCmLeverageBracketV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetCmLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmLeverageBracketV1`: []PmarginGetCmLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetCmLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmLeverageBracketV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmLeverageBracketV1RespItem**](PmarginGetCmLeverageBracketV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmPositionRiskV1

> []PmarginGetCmPositionRiskV1RespItem PmarginGetCmPositionRiskV1(ctx).Timestamp(timestamp).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()

Query CM Position Information(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	marginAsset := "marginAsset_example" // string |  (optional) (default to "")
	pair := "pair_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetCmPositionRiskV1(context.Background()).Timestamp(timestamp).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetCmPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmPositionRiskV1`: []PmarginGetCmPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetCmPositionRiskV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmPositionRiskV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **marginAsset** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmPositionRiskV1RespItem**](PmarginGetCmPositionRiskV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmPositionSideDualV1

> PmarginGetCmPositionSideDualV1Resp PmarginGetCmPositionSideDualV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get CM Current Position Mode(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetCmPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetCmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmPositionSideDualV1`: PmarginGetCmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetCmPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmPositionSideDualV1Resp**](PmarginGetCmPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginMarginInterestHistoryV1

> PmarginGetMarginMarginInterestHistoryV1Resp PmarginGetMarginMarginInterestHistoryV1(ctx).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Get Margin Borrow/Loan Interest History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	archived := "archived_example" // string | Default: `false`. Set to `true` for archived data from 6 months ago (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetMarginMarginInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetMarginMarginInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMarginInterestHistoryV1`: PmarginGetMarginMarginInterestHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetMarginMarginInterestHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginMarginInterestHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **archived** | **string** | Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**PmarginGetMarginMarginInterestHistoryV1Resp**](PmarginGetMarginMarginInterestHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginMarginLoanV1

> PmarginGetMarginMarginLoanV1Resp PmarginGetMarginMarginLoanV1(ctx).Asset(asset).Timestamp(timestamp).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Query Margin Loan Record(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	txId := int64(789) // int64 | the `tranId` in `POST/papi/v1/marginLoan` (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	archived := "archived_example" // string | Default: `false`. Set to `true` for archived data from 6 months ago (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetMarginMarginLoanV1(context.Background()).Asset(asset).Timestamp(timestamp).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetMarginMarginLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMarginLoanV1`: PmarginGetMarginMarginLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetMarginMarginLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginMarginLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **txId** | **int64** | the &#x60;tranId&#x60; in &#x60;POST/papi/v1/marginLoan&#x60; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **archived** | **string** | Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginMarginLoanV1Resp**](PmarginGetMarginMarginLoanV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginMaxBorrowableV1

> PmarginGetMarginMaxBorrowableV1Resp PmarginGetMarginMaxBorrowableV1(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Margin Max Borrow(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetMarginMaxBorrowableV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetMarginMaxBorrowableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMaxBorrowableV1`: PmarginGetMarginMaxBorrowableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetMarginMaxBorrowableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginMaxBorrowableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**PmarginGetMarginMaxBorrowableV1Resp**](PmarginGetMarginMaxBorrowableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginMaxWithdrawV1

> PmarginGetMarginMaxWithdrawV1Resp PmarginGetMarginMaxWithdrawV1(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Margin Max Withdraw(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetMarginMaxWithdrawV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetMarginMaxWithdrawV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMaxWithdrawV1`: PmarginGetMarginMaxWithdrawV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetMarginMaxWithdrawV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginMaxWithdrawV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**PmarginGetMarginMaxWithdrawV1Resp**](PmarginGetMarginMaxWithdrawV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginRepayLoanV1

> PmarginGetMarginRepayLoanV1Resp PmarginGetMarginRepayLoanV1(ctx).Asset(asset).Timestamp(timestamp).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()

Query Margin repay Record(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	txId := int64(789) // int64 | the tranId in `POST/papi/v1/repayLoan` (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	archived := "archived_example" // string | Default: `false`. Set to `true` for archived data from 6 months ago (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetMarginRepayLoanV1(context.Background()).Asset(asset).Timestamp(timestamp).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetMarginRepayLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginRepayLoanV1`: PmarginGetMarginRepayLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetMarginRepayLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginRepayLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **txId** | **int64** | the tranId in &#x60;POST/papi/v1/repayLoan&#x60; | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **archived** | **string** | Default: &#x60;false&#x60;. Set to &#x60;true&#x60; for archived data from 6 months ago | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginRepayLoanV1Resp**](PmarginGetMarginRepayLoanV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetPortfolioInterestHistoryV1

> []PmarginGetPortfolioInterestHistoryV1RespItem PmarginGetPortfolioInterestHistoryV1(ctx).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Negative Balance Interest History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetPortfolioInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetPortfolioInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetPortfolioInterestHistoryV1`: []PmarginGetPortfolioInterestHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetPortfolioInterestHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetPortfolioInterestHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetPortfolioInterestHistoryV1RespItem**](PmarginGetPortfolioInterestHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetPortfolioNegativeBalanceExchangeRecordV1

> PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp PmarginGetPortfolioNegativeBalanceExchangeRecordV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query User Negative Balance Auto Exchange Record (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	startTime := int64(789) // int64 | 
	endTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetPortfolioNegativeBalanceExchangeRecordV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetPortfolioNegativeBalanceExchangeRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetPortfolioNegativeBalanceExchangeRecordV1`: PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetPortfolioNegativeBalanceExchangeRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetPortfolioNegativeBalanceExchangeRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp**](PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetRateLimitOrderV1

> []PmarginGetRateLimitOrderV1RespItem PmarginGetRateLimitOrderV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query User Rate Limit (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetRateLimitOrderV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetRateLimitOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetRateLimitOrderV1`: []PmarginGetRateLimitOrderV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetRateLimitOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetRateLimitOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetRateLimitOrderV1RespItem**](PmarginGetRateLimitOrderV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetRepayFuturesSwitchV1

> PmarginGetRepayFuturesSwitchV1Resp PmarginGetRepayFuturesSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Auto-repay-futures Status(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetRepayFuturesSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetRepayFuturesSwitchV1`: PmarginGetRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetRepayFuturesSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetRepayFuturesSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetRepayFuturesSwitchV1Resp**](PmarginGetRepayFuturesSwitchV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmAccountConfigV1

> PmarginGetUmAccountConfigV1Resp PmarginGetUmAccountConfigV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

UM Futures Account Configuration(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmAccountConfigV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmAccountConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAccountConfigV1`: PmarginGetUmAccountConfigV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmAccountConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAccountConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmAccountConfigV1Resp**](PmarginGetUmAccountConfigV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmAccountV1

> PmarginGetUmAccountV1Resp PmarginGetUmAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Account Detail(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAccountV1`: PmarginGetUmAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmAccountV1Resp**](PmarginGetUmAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmAccountV2

> PmarginGetUmAccountV2Resp PmarginGetUmAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Account Detail V2(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAccountV2`: PmarginGetUmAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmAccountV2Resp**](PmarginGetUmAccountV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmApiTradingStatusV1

> PmarginGetUmApiTradingStatusV1Resp PmarginGetUmApiTradingStatusV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Portfolio Margin UM Trading Quantitative Rules Indicators(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmApiTradingStatusV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmApiTradingStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmApiTradingStatusV1`: PmarginGetUmApiTradingStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmApiTradingStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmApiTradingStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmApiTradingStatusV1Resp**](PmarginGetUmApiTradingStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmCommissionRateV1

> PmarginGetUmCommissionRateV1Resp PmarginGetUmCommissionRateV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get User Commission Rate for UM(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmCommissionRateV1`: PmarginGetUmCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmCommissionRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmCommissionRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmCommissionRateV1Resp**](PmarginGetUmCommissionRateV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmIncomeAsynIdV1

> PmarginGetUmIncomeAsynIdV1Resp PmarginGetUmIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Futures Transaction Download Link by Id(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmIncomeAsynIdV1`: PmarginGetUmIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmIncomeAsynIdV1Resp**](PmarginGetUmIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmIncomeAsynV1

> PmarginGetUmIncomeAsynV1Resp PmarginGetUmIncomeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For UM Futures Transaction History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmIncomeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmIncomeAsynV1`: PmarginGetUmIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmIncomeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmIncomeAsynV1Resp**](PmarginGetUmIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmIncomeV1

> []PmarginGetUmIncomeV1RespItem PmarginGetUmIncomeV1(ctx).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get UM Income History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	incomeType := "incomeType_example" // string | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get funding from INCLUSIVE. (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get funding until INCLUSIVE. (optional)
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Default 100; max 1000 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmIncomeV1`: []PmarginGetUmIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **incomeType** | **string** | TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get funding from INCLUSIVE. | 
 **endTime** | **int64** | Timestamp in ms to get funding until INCLUSIVE. | 
 **page** | **int32** |  | 
 **limit** | **int32** | Default 100; max 1000 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmIncomeV1RespItem**](PmarginGetUmIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmLeverageBracketV1

> []PmarginGetUmLeverageBracketV1RespItem PmarginGetUmLeverageBracketV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

UM Notional and Leverage Brackets (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmLeverageBracketV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmLeverageBracketV1`: []PmarginGetUmLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmLeverageBracketV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmLeverageBracketV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmLeverageBracketV1RespItem**](PmarginGetUmLeverageBracketV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOrderAsynIdV1

> PmarginGetUmOrderAsynIdV1Resp PmarginGetUmOrderAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Futures Order Download Link by Id(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmOrderAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmOrderAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderAsynIdV1`: PmarginGetUmOrderAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmOrderAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOrderAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmOrderAsynIdV1Resp**](PmarginGetUmOrderAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOrderAsynV1

> PmarginGetUmOrderAsynV1Resp PmarginGetUmOrderAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For UM Futures Order History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmOrderAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmOrderAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderAsynV1`: PmarginGetUmOrderAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmOrderAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOrderAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmOrderAsynV1Resp**](PmarginGetUmOrderAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmPositionRiskV1

> []PmarginGetUmPositionRiskV1RespItem PmarginGetUmPositionRiskV1(ctx).Execute()

Query UM Position Information(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmPositionRiskV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmPositionRiskV1`: []PmarginGetUmPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmPositionRiskV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmPositionRiskV1Request struct via the builder pattern


### Return type

[**[]PmarginGetUmPositionRiskV1RespItem**](PmarginGetUmPositionRiskV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmPositionSideDualV1

> PmarginGetUmPositionSideDualV1Resp PmarginGetUmPositionSideDualV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Current Position Mode(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmPositionSideDualV1`: PmarginGetUmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmPositionSideDualV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmPositionSideDualV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmPositionSideDualV1Resp**](PmarginGetUmPositionSideDualV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmSymbolConfigV1

> []PmarginGetUmSymbolConfigV1RespItem PmarginGetUmSymbolConfigV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

UM Futures Symbol Configuration(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmSymbolConfigV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmSymbolConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmSymbolConfigV1`: []PmarginGetUmSymbolConfigV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmSymbolConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmSymbolConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmSymbolConfigV1RespItem**](PmarginGetUmSymbolConfigV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmTradeAsynIdV1

> PmarginGetUmTradeAsynIdV1Resp PmarginGetUmTradeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Futures Trade Download Link by Id(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmTradeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmTradeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmTradeAsynIdV1`: PmarginGetUmTradeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmTradeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmTradeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmTradeAsynIdV1Resp**](PmarginGetUmTradeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmTradeAsynV1

> PmarginGetUmTradeAsynV1Resp PmarginGetUmTradeAsynV1(ctx).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Download Id For UM Futures Trade History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	startTime := int64(789) // int64 | Timestamp in ms
	endTime := int64(789) // int64 | Timestamp in ms
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PmarginGetUmTradeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginGetUmTradeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmTradeAsynV1`: PmarginGetUmTradeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginGetUmTradeAsynV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmTradeAsynV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **int64** | Timestamp in ms | 
 **endTime** | **int64** | Timestamp in ms | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmTradeAsynV1Resp**](PmarginGetUmTradeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

