# \V1API

All URIs are relative to *https://papi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginCreateAssetCollectionV1**](V1API.md#PmarginCreateAssetCollectionV1) | **Post** /papi/v1/asset-collection | Fund Collection by Asset(TRADE)
[**PmarginCreateAutoCollectionV1**](V1API.md#PmarginCreateAutoCollectionV1) | **Post** /papi/v1/auto-collection | Fund Auto-collection(TRADE)
[**PmarginCreateBnbTransferV1**](V1API.md#PmarginCreateBnbTransferV1) | **Post** /papi/v1/bnb-transfer | BNB transfer (TRADE)
[**PmarginCreateCmConditionalOrderV1**](V1API.md#PmarginCreateCmConditionalOrderV1) | **Post** /papi/v1/cm/conditional/order | New CM Conditional Order(TRADE)
[**PmarginCreateCmLeverageV1**](V1API.md#PmarginCreateCmLeverageV1) | **Post** /papi/v1/cm/leverage | Change CM Initial Leverage (TRADE)
[**PmarginCreateCmOrderV1**](V1API.md#PmarginCreateCmOrderV1) | **Post** /papi/v1/cm/order | New CM Order(TRADE)
[**PmarginCreateCmPositionSideDualV1**](V1API.md#PmarginCreateCmPositionSideDualV1) | **Post** /papi/v1/cm/positionSide/dual | Change CM Position Mode(TRADE)
[**PmarginCreateListenKeyV1**](V1API.md#PmarginCreateListenKeyV1) | **Post** /papi/v1/listenKey | Start User Data Stream(USER_STREAM)
[**PmarginCreateMarginLoanV1**](V1API.md#PmarginCreateMarginLoanV1) | **Post** /papi/v1/marginLoan | Margin Account Borrow(MARGIN)
[**PmarginCreateMarginOrderOcoV1**](V1API.md#PmarginCreateMarginOrderOcoV1) | **Post** /papi/v1/margin/order/oco | Margin Account New OCO(TRADE)
[**PmarginCreateMarginOrderV1**](V1API.md#PmarginCreateMarginOrderV1) | **Post** /papi/v1/margin/order | New Margin Order(TRADE)
[**PmarginCreateMarginRepayDebtV1**](V1API.md#PmarginCreateMarginRepayDebtV1) | **Post** /papi/v1/margin/repay-debt | Margin Account Repay Debt(TRADE)
[**PmarginCreateRepayFuturesNegativeBalanceV1**](V1API.md#PmarginCreateRepayFuturesNegativeBalanceV1) | **Post** /papi/v1/repay-futures-negative-balance | Repay futures Negative Balance(USER_DATA)
[**PmarginCreateRepayFuturesSwitchV1**](V1API.md#PmarginCreateRepayFuturesSwitchV1) | **Post** /papi/v1/repay-futures-switch | Change Auto-repay-futures Status(TRADE)
[**PmarginCreateRepayLoanV1**](V1API.md#PmarginCreateRepayLoanV1) | **Post** /papi/v1/repayLoan | Margin Account Repay(MARGIN)
[**PmarginCreateUmConditionalOrderV1**](V1API.md#PmarginCreateUmConditionalOrderV1) | **Post** /papi/v1/um/conditional/order | New UM Conditional Order (TRADE)
[**PmarginCreateUmFeeBurnV1**](V1API.md#PmarginCreateUmFeeBurnV1) | **Post** /papi/v1/um/feeBurn | Toggle BNB Burn On UM Futures Trade (TRADE)
[**PmarginCreateUmLeverageV1**](V1API.md#PmarginCreateUmLeverageV1) | **Post** /papi/v1/um/leverage | Change UM Initial Leverage(TRADE)
[**PmarginCreateUmOrderV1**](V1API.md#PmarginCreateUmOrderV1) | **Post** /papi/v1/um/order | New UM Order (TRADE)
[**PmarginCreateUmPositionSideDualV1**](V1API.md#PmarginCreateUmPositionSideDualV1) | **Post** /papi/v1/um/positionSide/dual | Change UM Position Mode(TRADE)
[**PmarginDeleteCmAllOpenOrdersV1**](V1API.md#PmarginDeleteCmAllOpenOrdersV1) | **Delete** /papi/v1/cm/allOpenOrders | Cancel All CM Open Orders(TRADE)
[**PmarginDeleteCmConditionalAllOpenOrdersV1**](V1API.md#PmarginDeleteCmConditionalAllOpenOrdersV1) | **Delete** /papi/v1/cm/conditional/allOpenOrders | Cancel All CM Open Conditional Orders(TRADE)
[**PmarginDeleteCmConditionalOrderV1**](V1API.md#PmarginDeleteCmConditionalOrderV1) | **Delete** /papi/v1/cm/conditional/order | Cancel CM Conditional Order(TRADE)
[**PmarginDeleteCmOrderV1**](V1API.md#PmarginDeleteCmOrderV1) | **Delete** /papi/v1/cm/order | Cancel CM Order(TRADE)
[**PmarginDeleteListenKeyV1**](V1API.md#PmarginDeleteListenKeyV1) | **Delete** /papi/v1/listenKey | Close User Data Stream(USER_STREAM)
[**PmarginDeleteMarginAllOpenOrdersV1**](V1API.md#PmarginDeleteMarginAllOpenOrdersV1) | **Delete** /papi/v1/margin/allOpenOrders | Cancel Margin Account All Open Orders on a Symbol(TRADE)
[**PmarginDeleteMarginOrderListV1**](V1API.md#PmarginDeleteMarginOrderListV1) | **Delete** /papi/v1/margin/orderList | Cancel Margin Account OCO Orders(TRADE)
[**PmarginDeleteMarginOrderV1**](V1API.md#PmarginDeleteMarginOrderV1) | **Delete** /papi/v1/margin/order | Cancel Margin Account Order(TRADE)
[**PmarginDeleteUmAllOpenOrdersV1**](V1API.md#PmarginDeleteUmAllOpenOrdersV1) | **Delete** /papi/v1/um/allOpenOrders | Cancel All UM Open Orders(TRADE)
[**PmarginDeleteUmConditionalAllOpenOrdersV1**](V1API.md#PmarginDeleteUmConditionalAllOpenOrdersV1) | **Delete** /papi/v1/um/conditional/allOpenOrders | Cancel All UM Open Conditional Orders (TRADE)
[**PmarginDeleteUmConditionalOrderV1**](V1API.md#PmarginDeleteUmConditionalOrderV1) | **Delete** /papi/v1/um/conditional/order | Cancel UM Conditional Order(TRADE)
[**PmarginDeleteUmOrderV1**](V1API.md#PmarginDeleteUmOrderV1) | **Delete** /papi/v1/um/order | Cancel UM Order(TRADE)
[**PmarginGetAccountV1**](V1API.md#PmarginGetAccountV1) | **Get** /papi/v1/account | Account Information(USER_DATA)
[**PmarginGetBalanceV1**](V1API.md#PmarginGetBalanceV1) | **Get** /papi/v1/balance | Account Balance(USER_DATA)
[**PmarginGetCmAccountV1**](V1API.md#PmarginGetCmAccountV1) | **Get** /papi/v1/cm/account | Get CM Account Detail(USER_DATA)
[**PmarginGetCmAdlQuantileV1**](V1API.md#PmarginGetCmAdlQuantileV1) | **Get** /papi/v1/cm/adlQuantile | CM Position ADL Quantile Estimation(USER_DATA)
[**PmarginGetCmAllOrdersV1**](V1API.md#PmarginGetCmAllOrdersV1) | **Get** /papi/v1/cm/allOrders | Query All CM Orders (USER_DATA)
[**PmarginGetCmCommissionRateV1**](V1API.md#PmarginGetCmCommissionRateV1) | **Get** /papi/v1/cm/commissionRate | Get User Commission Rate for CM(USER_DATA)
[**PmarginGetCmConditionalAllOrdersV1**](V1API.md#PmarginGetCmConditionalAllOrdersV1) | **Get** /papi/v1/cm/conditional/allOrders | Query All CM Conditional Orders(USER_DATA)
[**PmarginGetCmConditionalOpenOrderV1**](V1API.md#PmarginGetCmConditionalOpenOrderV1) | **Get** /papi/v1/cm/conditional/openOrder | Query Current CM Open Conditional Order(USER_DATA)
[**PmarginGetCmConditionalOpenOrdersV1**](V1API.md#PmarginGetCmConditionalOpenOrdersV1) | **Get** /papi/v1/cm/conditional/openOrders | Query All Current CM Open Conditional Orders (USER_DATA)
[**PmarginGetCmConditionalOrderHistoryV1**](V1API.md#PmarginGetCmConditionalOrderHistoryV1) | **Get** /papi/v1/cm/conditional/orderHistory | Query CM Conditional Order History(USER_DATA)
[**PmarginGetCmForceOrdersV1**](V1API.md#PmarginGetCmForceOrdersV1) | **Get** /papi/v1/cm/forceOrders | Query User&#39;s CM Force Orders(USER_DATA)
[**PmarginGetCmIncomeV1**](V1API.md#PmarginGetCmIncomeV1) | **Get** /papi/v1/cm/income | Get CM Income History(USER_DATA)
[**PmarginGetCmLeverageBracketV1**](V1API.md#PmarginGetCmLeverageBracketV1) | **Get** /papi/v1/cm/leverageBracket | CM Notional and Leverage Brackets(USER_DATA)
[**PmarginGetCmOpenOrderV1**](V1API.md#PmarginGetCmOpenOrderV1) | **Get** /papi/v1/cm/openOrder | Query Current CM Open Order (USER_DATA)
[**PmarginGetCmOpenOrdersV1**](V1API.md#PmarginGetCmOpenOrdersV1) | **Get** /papi/v1/cm/openOrders | Query All Current CM Open Orders(USER_DATA)
[**PmarginGetCmOrderAmendmentV1**](V1API.md#PmarginGetCmOrderAmendmentV1) | **Get** /papi/v1/cm/orderAmendment | Query CM Modify Order History(TRADE)
[**PmarginGetCmOrderV1**](V1API.md#PmarginGetCmOrderV1) | **Get** /papi/v1/cm/order | Query CM Order(USER_DATA)
[**PmarginGetCmPositionRiskV1**](V1API.md#PmarginGetCmPositionRiskV1) | **Get** /papi/v1/cm/positionRisk | Query CM Position Information(USER_DATA)
[**PmarginGetCmPositionSideDualV1**](V1API.md#PmarginGetCmPositionSideDualV1) | **Get** /papi/v1/cm/positionSide/dual | Get CM Current Position Mode(USER_DATA)
[**PmarginGetCmUserTradesV1**](V1API.md#PmarginGetCmUserTradesV1) | **Get** /papi/v1/cm/userTrades | CM Account Trade List(USER_DATA)
[**PmarginGetMarginAllOrderListV1**](V1API.md#PmarginGetMarginAllOrderListV1) | **Get** /papi/v1/margin/allOrderList | Query Margin Account&#39;s all OCO (USER_DATA)
[**PmarginGetMarginAllOrdersV1**](V1API.md#PmarginGetMarginAllOrdersV1) | **Get** /papi/v1/margin/allOrders | Query All Margin Account Orders (USER_DATA)
[**PmarginGetMarginForceOrdersV1**](V1API.md#PmarginGetMarginForceOrdersV1) | **Get** /papi/v1/margin/forceOrders | Query User&#39;s Margin Force Orders(USER_DATA)
[**PmarginGetMarginMarginInterestHistoryV1**](V1API.md#PmarginGetMarginMarginInterestHistoryV1) | **Get** /papi/v1/margin/marginInterestHistory | Get Margin Borrow/Loan Interest History(USER_DATA)
[**PmarginGetMarginMarginLoanV1**](V1API.md#PmarginGetMarginMarginLoanV1) | **Get** /papi/v1/margin/marginLoan | Query Margin Loan Record(USER_DATA)
[**PmarginGetMarginMaxBorrowableV1**](V1API.md#PmarginGetMarginMaxBorrowableV1) | **Get** /papi/v1/margin/maxBorrowable | Margin Max Borrow(USER_DATA)
[**PmarginGetMarginMaxWithdrawV1**](V1API.md#PmarginGetMarginMaxWithdrawV1) | **Get** /papi/v1/margin/maxWithdraw | Query Margin Max Withdraw(USER_DATA)
[**PmarginGetMarginMyTradesV1**](V1API.md#PmarginGetMarginMyTradesV1) | **Get** /papi/v1/margin/myTrades | Margin Account Trade List (USER_DATA)
[**PmarginGetMarginOpenOrderListV1**](V1API.md#PmarginGetMarginOpenOrderListV1) | **Get** /papi/v1/margin/openOrderList | Query Margin Account&#39;s Open OCO (USER_DATA)
[**PmarginGetMarginOpenOrdersV1**](V1API.md#PmarginGetMarginOpenOrdersV1) | **Get** /papi/v1/margin/openOrders | Query Current Margin Open Order (USER_DATA)
[**PmarginGetMarginOrderListV1**](V1API.md#PmarginGetMarginOrderListV1) | **Get** /papi/v1/margin/orderList | Query Margin Account&#39;s OCO (USER_DATA)
[**PmarginGetMarginOrderV1**](V1API.md#PmarginGetMarginOrderV1) | **Get** /papi/v1/margin/order | Query Margin Account Order (USER_DATA)
[**PmarginGetMarginRepayLoanV1**](V1API.md#PmarginGetMarginRepayLoanV1) | **Get** /papi/v1/margin/repayLoan | Query Margin repay Record(USER_DATA)
[**PmarginGetPingV1**](V1API.md#PmarginGetPingV1) | **Get** /papi/v1/ping | Test Connectivity
[**PmarginGetPortfolioInterestHistoryV1**](V1API.md#PmarginGetPortfolioInterestHistoryV1) | **Get** /papi/v1/portfolio/interest-history | Query Portfolio Margin Negative Balance Interest History(USER_DATA)
[**PmarginGetPortfolioNegativeBalanceExchangeRecordV1**](V1API.md#PmarginGetPortfolioNegativeBalanceExchangeRecordV1) | **Get** /papi/v1/portfolio/negative-balance-exchange-record | Query User Negative Balance Auto Exchange Record (USER_DATA)
[**PmarginGetRateLimitOrderV1**](V1API.md#PmarginGetRateLimitOrderV1) | **Get** /papi/v1/rateLimit/order | Query User Rate Limit (USER_DATA)
[**PmarginGetRepayFuturesSwitchV1**](V1API.md#PmarginGetRepayFuturesSwitchV1) | **Get** /papi/v1/repay-futures-switch | Get Auto-repay-futures Status(USER_DATA)
[**PmarginGetUmAccountConfigV1**](V1API.md#PmarginGetUmAccountConfigV1) | **Get** /papi/v1/um/accountConfig | UM Futures Account Configuration(USER_DATA)
[**PmarginGetUmAccountV1**](V1API.md#PmarginGetUmAccountV1) | **Get** /papi/v1/um/account | Get UM Account Detail(USER_DATA)
[**PmarginGetUmAdlQuantileV1**](V1API.md#PmarginGetUmAdlQuantileV1) | **Get** /papi/v1/um/adlQuantile | UM Position ADL Quantile Estimation(USER_DATA)
[**PmarginGetUmAllOrdersV1**](V1API.md#PmarginGetUmAllOrdersV1) | **Get** /papi/v1/um/allOrders | Query All UM Orders(USER_DATA)
[**PmarginGetUmApiTradingStatusV1**](V1API.md#PmarginGetUmApiTradingStatusV1) | **Get** /papi/v1/um/apiTradingStatus | Portfolio Margin UM Trading Quantitative Rules Indicators(USER_DATA)
[**PmarginGetUmCommissionRateV1**](V1API.md#PmarginGetUmCommissionRateV1) | **Get** /papi/v1/um/commissionRate | Get User Commission Rate for UM(USER_DATA)
[**PmarginGetUmConditionalAllOrdersV1**](V1API.md#PmarginGetUmConditionalAllOrdersV1) | **Get** /papi/v1/um/conditional/allOrders | Query All UM Conditional Orders(USER_DATA)
[**PmarginGetUmConditionalOpenOrderV1**](V1API.md#PmarginGetUmConditionalOpenOrderV1) | **Get** /papi/v1/um/conditional/openOrder | Query Current UM Open Conditional Order(USER_DATA)
[**PmarginGetUmConditionalOpenOrdersV1**](V1API.md#PmarginGetUmConditionalOpenOrdersV1) | **Get** /papi/v1/um/conditional/openOrders | Query All Current UM Open Conditional Orders(USER_DATA)
[**PmarginGetUmConditionalOrderHistoryV1**](V1API.md#PmarginGetUmConditionalOrderHistoryV1) | **Get** /papi/v1/um/conditional/orderHistory | Query UM Conditional Order History(USER_DATA)
[**PmarginGetUmFeeBurnV1**](V1API.md#PmarginGetUmFeeBurnV1) | **Get** /papi/v1/um/feeBurn | Get UM Futures BNB Burn Status (USER_DATA)
[**PmarginGetUmForceOrdersV1**](V1API.md#PmarginGetUmForceOrdersV1) | **Get** /papi/v1/um/forceOrders | Query User&#39;s UM Force Orders (USER_DATA)
[**PmarginGetUmIncomeAsynIdV1**](V1API.md#PmarginGetUmIncomeAsynIdV1) | **Get** /papi/v1/um/income/asyn/id | Get UM Futures Transaction Download Link by Id(USER_DATA)
[**PmarginGetUmIncomeAsynV1**](V1API.md#PmarginGetUmIncomeAsynV1) | **Get** /papi/v1/um/income/asyn | Get Download Id For UM Futures Transaction History (USER_DATA)
[**PmarginGetUmIncomeV1**](V1API.md#PmarginGetUmIncomeV1) | **Get** /papi/v1/um/income | Get UM Income History(USER_DATA)
[**PmarginGetUmLeverageBracketV1**](V1API.md#PmarginGetUmLeverageBracketV1) | **Get** /papi/v1/um/leverageBracket | UM Notional and Leverage Brackets (USER_DATA)
[**PmarginGetUmOpenOrderV1**](V1API.md#PmarginGetUmOpenOrderV1) | **Get** /papi/v1/um/openOrder | Query Current UM Open Order(USER_DATA)
[**PmarginGetUmOpenOrdersV1**](V1API.md#PmarginGetUmOpenOrdersV1) | **Get** /papi/v1/um/openOrders | Query All Current UM Open Orders(USER_DATA)
[**PmarginGetUmOrderAmendmentV1**](V1API.md#PmarginGetUmOrderAmendmentV1) | **Get** /papi/v1/um/orderAmendment | Query UM Modify Order History(TRADE)
[**PmarginGetUmOrderAsynIdV1**](V1API.md#PmarginGetUmOrderAsynIdV1) | **Get** /papi/v1/um/order/asyn/id | Get UM Futures Order Download Link by Id(USER_DATA)
[**PmarginGetUmOrderAsynV1**](V1API.md#PmarginGetUmOrderAsynV1) | **Get** /papi/v1/um/order/asyn | Get Download Id For UM Futures Order History (USER_DATA)
[**PmarginGetUmOrderV1**](V1API.md#PmarginGetUmOrderV1) | **Get** /papi/v1/um/order | Query UM Order (USER_DATA)
[**PmarginGetUmPositionRiskV1**](V1API.md#PmarginGetUmPositionRiskV1) | **Get** /papi/v1/um/positionRisk | Query UM Position Information(USER_DATA)
[**PmarginGetUmPositionSideDualV1**](V1API.md#PmarginGetUmPositionSideDualV1) | **Get** /papi/v1/um/positionSide/dual | Get UM Current Position Mode(USER_DATA)
[**PmarginGetUmSymbolConfigV1**](V1API.md#PmarginGetUmSymbolConfigV1) | **Get** /papi/v1/um/symbolConfig | UM Futures Symbol Configuration(USER_DATA)
[**PmarginGetUmTradeAsynIdV1**](V1API.md#PmarginGetUmTradeAsynIdV1) | **Get** /papi/v1/um/trade/asyn/id | Get UM Futures Trade Download Link by Id(USER_DATA)
[**PmarginGetUmTradeAsynV1**](V1API.md#PmarginGetUmTradeAsynV1) | **Get** /papi/v1/um/trade/asyn | Get Download Id For UM Futures Trade History (USER_DATA)
[**PmarginGetUmUserTradesV1**](V1API.md#PmarginGetUmUserTradesV1) | **Get** /papi/v1/um/userTrades | UM Account Trade List(USER_DATA)
[**PmarginUpdateCmOrderV1**](V1API.md#PmarginUpdateCmOrderV1) | **Put** /papi/v1/cm/order | Modify CM Order(TRADE)
[**PmarginUpdateListenKeyV1**](V1API.md#PmarginUpdateListenKeyV1) | **Put** /papi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)
[**PmarginUpdateUmOrderV1**](V1API.md#PmarginUpdateUmOrderV1) | **Put** /papi/v1/um/order | Modify UM Order(TRADE)



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
	resp, r, err := apiClient.V1API.PmarginCreateAssetCollectionV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateAssetCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateAssetCollectionV1`: PmarginCreateAssetCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateAssetCollectionV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginCreateAutoCollectionV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateAutoCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateAutoCollectionV1`: PmarginCreateAutoCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateAutoCollectionV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginCreateBnbTransferV1(context.Background()).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateBnbTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateBnbTransferV1`: PmarginCreateBnbTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateBnbTransferV1`: %v\n", resp)
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


## PmarginCreateCmConditionalOrderV1

> PmarginCreateCmConditionalOrderV1Resp PmarginCreateCmConditionalOrderV1(ctx).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()

New CM Conditional Order(TRADE)



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
	side := "side_example" // string |  (default to "")
	strategyType := "strategyType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	activationPrice := "activationPrice_example" // string |  (optional) (default to "")
	callbackRate := "callbackRate_example" // string |  (optional) (default to "")
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceProtect := "priceProtect_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	workingType := "workingType_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateCmConditionalOrderV1(context.Background()).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateCmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmConditionalOrderV1`: PmarginCreateCmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateCmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateCmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **strategyType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **activationPrice** | **string** |  | [default to &quot;&quot;]
 **callbackRate** | **string** |  | [default to &quot;&quot;]
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceProtect** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateCmConditionalOrderV1Resp**](PmarginCreateCmConditionalOrderV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginCreateCmLeverageV1(context.Background()).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateCmLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmLeverageV1`: PmarginCreateCmLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateCmLeverageV1`: %v\n", resp)
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


## PmarginCreateCmOrderV1

> PmarginCreateCmOrderV1Resp PmarginCreateCmOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()

New CM Order(TRADE)



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
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateCmOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmOrderV1`: PmarginCreateCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateCmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateCmOrderV1Resp**](PmarginCreateCmOrderV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginCreateCmPositionSideDualV1(context.Background()).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateCmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateCmPositionSideDualV1`: PmarginCreateCmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateCmPositionSideDualV1`: %v\n", resp)
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


## PmarginCreateListenKeyV1

> PmarginCreateListenKeyV1Resp PmarginCreateListenKeyV1(ctx).Execute()

Start User Data Stream(USER_STREAM)



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
	resp, r, err := apiClient.V1API.PmarginCreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateListenKeyV1`: PmarginCreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateListenKeyV1Request struct via the builder pattern


### Return type

[**PmarginCreateListenKeyV1Resp**](PmarginCreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginLoanV1

> PmarginCreateMarginLoanV1Resp PmarginCreateMarginLoanV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Margin Account Borrow(MARGIN)



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
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateMarginLoanV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateMarginLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginLoanV1`: PmarginCreateMarginLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateMarginLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateMarginLoanV1Resp**](PmarginCreateMarginLoanV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginOrderOcoV1

> PmarginCreateMarginOrderOcoV1Resp PmarginCreateMarginOrderOcoV1(ctx).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()

Margin Account New OCO(TRADE)



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
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	stopPrice := "stopPrice_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	limitClientOrderId := "limitClientOrderId_example" // string |  (optional) (default to "")
	limitIcebergQty := "limitIcebergQty_example" // string |  (optional) (default to "")
	listClientOrderId := "listClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	sideEffectType := "sideEffectType_example" // string |  (optional) (default to "")
	stopClientOrderId := "stopClientOrderId_example" // string |  (optional) (default to "")
	stopIcebergQty := "stopIcebergQty_example" // string |  (optional) (default to "")
	stopLimitPrice := "stopLimitPrice_example" // string |  (optional) (default to "")
	stopLimitTimeInForce := "stopLimitTimeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateMarginOrderOcoV1(context.Background()).Price(price).Quantity(quantity).Side(side).StopPrice(stopPrice).Symbol(symbol).Timestamp(timestamp).LimitClientOrderId(limitClientOrderId).LimitIcebergQty(limitIcebergQty).ListClientOrderId(listClientOrderId).NewOrderRespType(newOrderRespType).RecvWindow(recvWindow).SideEffectType(sideEffectType).StopClientOrderId(stopClientOrderId).StopIcebergQty(stopIcebergQty).StopLimitPrice(stopLimitPrice).StopLimitTimeInForce(stopLimitTimeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateMarginOrderOcoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginOrderOcoV1`: PmarginCreateMarginOrderOcoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateMarginOrderOcoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginOrderOcoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **limitClientOrderId** | **string** |  | [default to &quot;&quot;]
 **limitIcebergQty** | **string** |  | [default to &quot;&quot;]
 **listClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **sideEffectType** | **string** |  | [default to &quot;&quot;]
 **stopClientOrderId** | **string** |  | [default to &quot;&quot;]
 **stopIcebergQty** | **string** |  | [default to &quot;&quot;]
 **stopLimitPrice** | **string** |  | [default to &quot;&quot;]
 **stopLimitTimeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateMarginOrderOcoV1Resp**](PmarginCreateMarginOrderOcoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginOrderV1

> PmarginCreateMarginOrderV1Resp PmarginCreateMarginOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()

New Margin Order(TRADE)



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
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	autoRepayAtCancel := true // bool |  (optional)
	icebergQty := "icebergQty_example" // string |  (optional) (default to "")
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
	resp, r, err := apiClient.V1API.PmarginCreateMarginOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).AutoRepayAtCancel(autoRepayAtCancel).IcebergQty(icebergQty).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).Price(price).Quantity(quantity).QuoteOrderQty(quoteOrderQty).RecvWindow(recvWindow).SelfTradePreventionMode(selfTradePreventionMode).SideEffectType(sideEffectType).StopPrice(stopPrice).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginOrderV1`: PmarginCreateMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **autoRepayAtCancel** | **bool** |  | 
 **icebergQty** | **string** |  | [default to &quot;&quot;]
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

[**PmarginCreateMarginOrderV1Resp**](PmarginCreateMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateMarginRepayDebtV1

> PmarginCreateMarginRepayDebtV1Resp PmarginCreateMarginRepayDebtV1(ctx).Asset(asset).Timestamp(timestamp).Amount(amount).RecvWindow(recvWindow).SpecifyRepayAssets(specifyRepayAssets).Execute()

Margin Account Repay Debt(TRADE)



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
	amount := "amount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	specifyRepayAssets := "specifyRepayAssets_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateMarginRepayDebtV1(context.Background()).Asset(asset).Timestamp(timestamp).Amount(amount).RecvWindow(recvWindow).SpecifyRepayAssets(specifyRepayAssets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateMarginRepayDebtV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateMarginRepayDebtV1`: PmarginCreateMarginRepayDebtV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateMarginRepayDebtV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateMarginRepayDebtV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **amount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **specifyRepayAssets** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateMarginRepayDebtV1Resp**](PmarginCreateMarginRepayDebtV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginCreateRepayFuturesNegativeBalanceV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateRepayFuturesNegativeBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateRepayFuturesNegativeBalanceV1`: PmarginCreateRepayFuturesNegativeBalanceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateRepayFuturesNegativeBalanceV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginCreateRepayFuturesSwitchV1(context.Background()).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateRepayFuturesSwitchV1`: PmarginCreateRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateRepayFuturesSwitchV1`: %v\n", resp)
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


## PmarginCreateRepayLoanV1

> PmarginCreateRepayLoanV1Resp PmarginCreateRepayLoanV1(ctx).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Margin Account Repay(MARGIN)



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
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateRepayLoanV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateRepayLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateRepayLoanV1`: PmarginCreateRepayLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateRepayLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateRepayLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateRepayLoanV1Resp**](PmarginCreateRepayLoanV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateUmConditionalOrderV1

> PmarginCreateUmConditionalOrderV1Resp PmarginCreateUmConditionalOrderV1(ctx).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).GoodTillDate(goodTillDate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()

New UM Conditional Order (TRADE)



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
	side := "side_example" // string |  (default to "")
	strategyType := "strategyType_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	activationPrice := "activationPrice_example" // string |  (optional) (default to "")
	callbackRate := "callbackRate_example" // string |  (optional) (default to "")
	goodTillDate := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	priceProtect := "priceProtect_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	stopPrice := "stopPrice_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")
	workingType := "workingType_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateUmConditionalOrderV1(context.Background()).Side(side).StrategyType(strategyType).Symbol(symbol).Timestamp(timestamp).ActivationPrice(activationPrice).CallbackRate(callbackRate).GoodTillDate(goodTillDate).NewClientStrategyId(newClientStrategyId).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).PriceProtect(priceProtect).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).StopPrice(stopPrice).TimeInForce(timeInForce).WorkingType(workingType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateUmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmConditionalOrderV1`: PmarginCreateUmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateUmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **strategyType** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **activationPrice** | **string** |  | [default to &quot;&quot;]
 **callbackRate** | **string** |  | [default to &quot;&quot;]
 **goodTillDate** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **priceProtect** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **stopPrice** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]
 **workingType** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateUmConditionalOrderV1Resp**](PmarginCreateUmConditionalOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginCreateUmFeeBurnV1

> PmarginCreateUmFeeBurnV1Resp PmarginCreateUmFeeBurnV1(ctx).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Toggle BNB Burn On UM Futures Trade (TRADE)



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
	feeBurn := "feeBurn_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateUmFeeBurnV1(context.Background()).FeeBurn(feeBurn).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateUmFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmFeeBurnV1`: PmarginCreateUmFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateUmFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feeBurn** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginCreateUmFeeBurnV1Resp**](PmarginCreateUmFeeBurnV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginCreateUmLeverageV1(context.Background()).Leverage(leverage).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateUmLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmLeverageV1`: PmarginCreateUmLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateUmLeverageV1`: %v\n", resp)
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


## PmarginCreateUmOrderV1

> PmarginCreateUmOrderV1Resp PmarginCreateUmOrderV1(ctx).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).TimeInForce(timeInForce).Execute()

New UM Order (TRADE)



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
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	type_ := "type__example" // string |  (default to "")
	goodTillDate := int64(789) // int64 |  (optional)
	newClientOrderId := "newClientOrderId_example" // string |  (optional) (default to "")
	newOrderRespType := "newOrderRespType_example" // string |  (optional) (default to "")
	positionSide := "positionSide_example" // string |  (optional) (default to "")
	price := "price_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	quantity := "quantity_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	reduceOnly := "reduceOnly_example" // string |  (optional) (default to "")
	selfTradePreventionMode := "selfTradePreventionMode_example" // string |  (optional) (default to "")
	timeInForce := "timeInForce_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginCreateUmOrderV1(context.Background()).Side(side).Symbol(symbol).Timestamp(timestamp).Type_(type_).GoodTillDate(goodTillDate).NewClientOrderId(newClientOrderId).NewOrderRespType(newOrderRespType).PositionSide(positionSide).Price(price).PriceMatch(priceMatch).Quantity(quantity).RecvWindow(recvWindow).ReduceOnly(reduceOnly).SelfTradePreventionMode(selfTradePreventionMode).TimeInForce(timeInForce).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmOrderV1`: PmarginCreateUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateUmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **type_** | **string** |  | [default to &quot;&quot;]
 **goodTillDate** | **int64** |  | 
 **newClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newOrderRespType** | **string** |  | [default to &quot;&quot;]
 **positionSide** | **string** |  | [default to &quot;&quot;]
 **price** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **reduceOnly** | **string** |  | [default to &quot;&quot;]
 **selfTradePreventionMode** | **string** |  | [default to &quot;&quot;]
 **timeInForce** | **string** |  | [default to &quot;&quot;]

### Return type

[**PmarginCreateUmOrderV1Resp**](PmarginCreateUmOrderV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginCreateUmPositionSideDualV1(context.Background()).DualSidePosition(dualSidePosition).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginCreateUmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateUmPositionSideDualV1`: PmarginCreateUmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginCreateUmPositionSideDualV1`: %v\n", resp)
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


## PmarginDeleteCmAllOpenOrdersV1

> PmarginDeleteCmAllOpenOrdersV1Resp PmarginDeleteCmAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All CM Open Orders(TRADE)



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
	resp, r, err := apiClient.V1API.PmarginDeleteCmAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteCmAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmAllOpenOrdersV1`: PmarginDeleteCmAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteCmAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmAllOpenOrdersV1Resp**](PmarginDeleteCmAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteCmConditionalAllOpenOrdersV1

> PmarginDeleteCmConditionalAllOpenOrdersV1Resp PmarginDeleteCmConditionalAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All CM Open Conditional Orders(TRADE)



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
	resp, r, err := apiClient.V1API.PmarginDeleteCmConditionalAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteCmConditionalAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmConditionalAllOpenOrdersV1`: PmarginDeleteCmConditionalAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteCmConditionalAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmConditionalAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmConditionalAllOpenOrdersV1Resp**](PmarginDeleteCmConditionalAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteCmConditionalOrderV1

> PmarginDeleteCmConditionalOrderV1Resp PmarginDeleteCmConditionalOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Cancel CM Conditional Order(TRADE)



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
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginDeleteCmConditionalOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteCmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmConditionalOrderV1`: PmarginDeleteCmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteCmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmConditionalOrderV1Resp**](PmarginDeleteCmConditionalOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteCmOrderV1

> PmarginDeleteCmOrderV1Resp PmarginDeleteCmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel CM Order(TRADE)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginDeleteCmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteCmOrderV1`: PmarginDeleteCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteCmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteCmOrderV1Resp**](PmarginDeleteCmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteListenKeyV1

> map[string]interface{} PmarginDeleteListenKeyV1(ctx).Execute()

Close User Data Stream(USER_STREAM)



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
	resp, r, err := apiClient.V1API.PmarginDeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteListenKeyV1Request struct via the builder pattern


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


## PmarginDeleteMarginAllOpenOrdersV1

> []PmarginDeleteMarginAllOpenOrdersV1RespInner PmarginDeleteMarginAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel Margin Account All Open Orders on a Symbol(TRADE)



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
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginDeleteMarginAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteMarginAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteMarginAllOpenOrdersV1`: []PmarginDeleteMarginAllOpenOrdersV1RespInner
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteMarginAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteMarginAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginDeleteMarginAllOpenOrdersV1RespInner**](PmarginDeleteMarginAllOpenOrdersV1RespInner.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteMarginOrderListV1

> PmarginDeleteMarginOrderListV1Resp PmarginDeleteMarginOrderListV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Margin Account OCO Orders(TRADE)



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
	orderListId := int64(789) // int64 | Either `orderListId` or `listClientOrderId` must be provided (optional)
	listClientOrderId := "listClientOrderId_example" // string | Either `orderListId` or `listClientOrderId` must be provided (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginDeleteMarginOrderListV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderListId(orderListId).ListClientOrderId(listClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteMarginOrderListV1`: PmarginDeleteMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | 
 **listClientOrderId** | **string** | Either &#x60;orderListId&#x60; or &#x60;listClientOrderId&#x60; must be provided | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginDeleteMarginOrderListV1Resp**](PmarginDeleteMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteMarginOrderV1

> PmarginDeleteMarginOrderV1Resp PmarginDeleteMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()

Cancel Margin Account Order(TRADE)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	newClientOrderId := "newClientOrderId_example" // string | Used to uniquely identify this cancel. Automatically generated by default. (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginDeleteMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).NewClientOrderId(newClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteMarginOrderV1`: PmarginDeleteMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **newClientOrderId** | **string** | Used to uniquely identify this cancel. Automatically generated by default. | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**PmarginDeleteMarginOrderV1Resp**](PmarginDeleteMarginOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmAllOpenOrdersV1

> PmarginDeleteUmAllOpenOrdersV1Resp PmarginDeleteUmAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All UM Open Orders(TRADE)



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
	resp, r, err := apiClient.V1API.PmarginDeleteUmAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteUmAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmAllOpenOrdersV1`: PmarginDeleteUmAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteUmAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmAllOpenOrdersV1Resp**](PmarginDeleteUmAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmConditionalAllOpenOrdersV1

> PmarginDeleteUmConditionalAllOpenOrdersV1Resp PmarginDeleteUmConditionalAllOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Cancel All UM Open Conditional Orders (TRADE)



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
	resp, r, err := apiClient.V1API.PmarginDeleteUmConditionalAllOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteUmConditionalAllOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmConditionalAllOpenOrdersV1`: PmarginDeleteUmConditionalAllOpenOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteUmConditionalAllOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmConditionalAllOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmConditionalAllOpenOrdersV1Resp**](PmarginDeleteUmConditionalAllOpenOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmConditionalOrderV1

> PmarginDeleteUmConditionalOrderV1Resp PmarginDeleteUmConditionalOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Cancel UM Conditional Order(TRADE)



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
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginDeleteUmConditionalOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteUmConditionalOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmConditionalOrderV1`: PmarginDeleteUmConditionalOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteUmConditionalOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmConditionalOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmConditionalOrderV1Resp**](PmarginDeleteUmConditionalOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteUmOrderV1

> PmarginDeleteUmOrderV1Resp PmarginDeleteUmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Cancel UM Order(TRADE)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginDeleteUmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginDeleteUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteUmOrderV1`: PmarginDeleteUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginDeleteUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteUmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginDeleteUmOrderV1Resp**](PmarginDeleteUmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.V1API.PmarginGetAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetAccountV1`: PmarginGetAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetAccountV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetBalanceV1(context.Background()).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetBalanceV1`: PmarginGetBalanceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetBalanceV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetCmAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmAccountV1`: PmarginGetCmAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmAccountV1`: %v\n", resp)
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


## PmarginGetCmAdlQuantileV1

> []PmarginGetCmAdlQuantileV1RespItem PmarginGetCmAdlQuantileV1(ctx).Execute()

CM Position ADL Quantile Estimation(USER_DATA)



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
	resp, r, err := apiClient.V1API.PmarginGetCmAdlQuantileV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmAdlQuantileV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmAdlQuantileV1`: []PmarginGetCmAdlQuantileV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmAdlQuantileV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmAdlQuantileV1Request struct via the builder pattern


### Return type

[**[]PmarginGetCmAdlQuantileV1RespItem**](PmarginGetCmAdlQuantileV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmAllOrdersV1

> []PmarginGetCmAllOrdersV1RespItem PmarginGetCmAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All CM Orders (USER_DATA)



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
	pair := "pair_example" // string |  (optional) (default to "")
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Pair(pair).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmAllOrdersV1`: []PmarginGetCmAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **pair** | **string** |  | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmAllOrdersV1RespItem**](PmarginGetCmAllOrdersV1RespItem.md)

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
	resp, r, err := apiClient.V1API.PmarginGetCmCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmCommissionRateV1`: PmarginGetCmCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmCommissionRateV1`: %v\n", resp)
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


## PmarginGetCmConditionalAllOrdersV1

> []PmarginGetCmConditionalAllOrdersV1RespItem PmarginGetCmConditionalAllOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All CM Conditional Orders(USER_DATA)



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
	strategyId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmConditionalAllOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmConditionalAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalAllOrdersV1`: []PmarginGetCmConditionalAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmConditionalAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmConditionalAllOrdersV1RespItem**](PmarginGetCmConditionalAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmConditionalOpenOrderV1

> PmarginGetCmConditionalOpenOrderV1Resp PmarginGetCmConditionalOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query Current CM Open Conditional Order(USER_DATA)



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
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmConditionalOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmConditionalOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalOpenOrderV1`: PmarginGetCmConditionalOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmConditionalOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmConditionalOpenOrderV1Resp**](PmarginGetCmConditionalOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmConditionalOpenOrdersV1

> []PmarginGetCmConditionalOpenOrdersV1RespItem PmarginGetCmConditionalOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query All Current CM Open Conditional Orders (USER_DATA)



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
	resp, r, err := apiClient.V1API.PmarginGetCmConditionalOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmConditionalOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalOpenOrdersV1`: []PmarginGetCmConditionalOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmConditionalOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmConditionalOpenOrdersV1RespItem**](PmarginGetCmConditionalOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmConditionalOrderHistoryV1

> PmarginGetCmConditionalOrderHistoryV1Resp PmarginGetCmConditionalOrderHistoryV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query CM Conditional Order History(USER_DATA)



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
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmConditionalOrderHistoryV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmConditionalOrderHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmConditionalOrderHistoryV1`: PmarginGetCmConditionalOrderHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmConditionalOrderHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmConditionalOrderHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmConditionalOrderHistoryV1Resp**](PmarginGetCmConditionalOrderHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmForceOrdersV1

> []PmarginGetCmForceOrdersV1RespItem PmarginGetCmForceOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query User's CM Force Orders(USER_DATA)



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
	autoCloseType := "autoCloseType_example" // string | &#34;LIQUIDATION&#34; for liquidation orders, &#34;ADL&#34; for ADL orders. (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmForceOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmForceOrdersV1`: []PmarginGetCmForceOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **autoCloseType** | **string** | &amp;#34;LIQUIDATION&amp;#34; for liquidation orders, &amp;#34;ADL&amp;#34; for ADL orders. | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetCmForceOrdersV1RespItem**](PmarginGetCmForceOrdersV1RespItem.md)

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
	resp, r, err := apiClient.V1API.PmarginGetCmIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmIncomeV1`: []PmarginGetCmIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmIncomeV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetCmLeverageBracketV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmLeverageBracketV1`: []PmarginGetCmLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmLeverageBracketV1`: %v\n", resp)
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


## PmarginGetCmOpenOrderV1

> PmarginGetCmOpenOrderV1Resp PmarginGetCmOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current CM Open Order (USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOpenOrderV1`: PmarginGetCmOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmOpenOrderV1Resp**](PmarginGetCmOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmOpenOrdersV1

> []PmarginGetCmOpenOrdersV1RespItem PmarginGetCmOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()

Query All Current CM Open Orders(USER_DATA)



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
	pair := "pair_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOpenOrdersV1`: []PmarginGetCmOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmOpenOrdersV1RespItem**](PmarginGetCmOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmOrderAmendmentV1

> []PmarginGetCmOrderAmendmentV1RespItem PmarginGetCmOrderAmendmentV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query CM Modify Order History(TRADE)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get modification history from INCLUSIVE (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get modification history until INCLUSIVE (optional)
	limit := int32(56) // int32 | Default 50, max 100 (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmOrderAmendmentV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmOrderAmendmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOrderAmendmentV1`: []PmarginGetCmOrderAmendmentV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmOrderAmendmentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOrderAmendmentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get modification history from INCLUSIVE | 
 **endTime** | **int64** | Timestamp in ms to get modification history until INCLUSIVE | 
 **limit** | **int32** | Default 50, max 100 | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmOrderAmendmentV1RespItem**](PmarginGetCmOrderAmendmentV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetCmOrderV1

> PmarginGetCmOrderV1Resp PmarginGetCmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query CM Order(USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmOrderV1`: PmarginGetCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetCmOrderV1Resp**](PmarginGetCmOrderV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginGetCmPositionRiskV1(context.Background()).Timestamp(timestamp).MarginAsset(marginAsset).Pair(pair).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmPositionRiskV1`: []PmarginGetCmPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmPositionRiskV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetCmPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmPositionSideDualV1`: PmarginGetCmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmPositionSideDualV1`: %v\n", resp)
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


## PmarginGetCmUserTradesV1

> []PmarginGetCmUserTradesV1RespItem PmarginGetCmUserTradesV1(ctx).Timestamp(timestamp).Symbol(symbol).Pair(pair).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

CM Account Trade List(USER_DATA)



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
	pair := "pair_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 50; max 1000. (optional) (default to 50)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetCmUserTradesV1(context.Background()).Timestamp(timestamp).Symbol(symbol).Pair(pair).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetCmUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetCmUserTradesV1`: []PmarginGetCmUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetCmUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetCmUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **pair** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 50; max 1000. | [default to 50]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetCmUserTradesV1RespItem**](PmarginGetCmUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginAllOrderListV1

> []PmarginGetMarginAllOrderListV1RespItem PmarginGetMarginAllOrderListV1(ctx).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query Margin Account's all OCO (USER_DATA)



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
	fromId := int64(789) // int64 | If supplied, neither startTime or endTime can be provided (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 500. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginAllOrderListV1(context.Background()).Timestamp(timestamp).FromId(fromId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginAllOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginAllOrderListV1`: []PmarginGetMarginAllOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginAllOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginAllOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **fromId** | **int64** | If supplied, neither startTime or endTime can be provided | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 500. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginAllOrderListV1RespItem**](PmarginGetMarginAllOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginAllOrdersV1

> []PmarginGetMarginAllOrdersV1RespItem PmarginGetMarginAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All Margin Account Orders (USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 500. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginAllOrdersV1`: []PmarginGetMarginAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 500. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginAllOrdersV1RespItem**](PmarginGetMarginAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginForceOrdersV1

> PmarginGetMarginForceOrdersV1Resp PmarginGetMarginForceOrdersV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query User's Margin Force Orders(USER_DATA)



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
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginForceOrdersV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginForceOrdersV1`: PmarginGetMarginForceOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginForceOrdersV1Resp**](PmarginGetMarginForceOrdersV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginGetMarginMarginInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginMarginInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMarginInterestHistoryV1`: PmarginGetMarginMarginInterestHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginMarginInterestHistoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetMarginMarginLoanV1(context.Background()).Asset(asset).Timestamp(timestamp).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginMarginLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMarginLoanV1`: PmarginGetMarginMarginLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginMarginLoanV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetMarginMaxBorrowableV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginMaxBorrowableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMaxBorrowableV1`: PmarginGetMarginMaxBorrowableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginMaxBorrowableV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetMarginMaxWithdrawV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginMaxWithdrawV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMaxWithdrawV1`: PmarginGetMarginMaxWithdrawV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginMaxWithdrawV1`: %v\n", resp)
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


## PmarginGetMarginMyTradesV1

> []PmarginGetMarginMyTradesV1RespItem PmarginGetMarginMyTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Margin Account Trade List (USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | TradeId to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginMyTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginMyTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginMyTradesV1`: []PmarginGetMarginMyTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginMyTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginMyTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | TradeId to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginMyTradesV1RespItem**](PmarginGetMarginMyTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOpenOrderListV1

> []PmarginGetMarginOpenOrderListV1RespItem PmarginGetMarginOpenOrderListV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Margin Account's Open OCO (USER_DATA)



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
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginOpenOrderListV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginOpenOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOpenOrderListV1`: []PmarginGetMarginOpenOrderListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginOpenOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOpenOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginOpenOrderListV1RespItem**](PmarginGetMarginOpenOrderListV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOpenOrdersV1

> []PmarginGetMarginOpenOrdersV1RespItem PmarginGetMarginOpenOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Current Margin Open Order (USER_DATA)



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
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginOpenOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOpenOrdersV1`: []PmarginGetMarginOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetMarginOpenOrdersV1RespItem**](PmarginGetMarginOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOrderListV1

> PmarginGetMarginOrderListV1Resp PmarginGetMarginOrderListV1(ctx).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account's OCO (USER_DATA)



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
	orderListId := int64(789) // int64 | Either orderListId or origClientOrderId must be provided (optional)
	origClientOrderId := "origClientOrderId_example" // string | Either orderListId or origClientOrderId must be provided (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginOrderListV1(context.Background()).Timestamp(timestamp).OrderListId(orderListId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginOrderListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOrderListV1`: PmarginGetMarginOrderListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginOrderListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOrderListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderListId** | **int64** | Either orderListId or origClientOrderId must be provided | 
 **origClientOrderId** | **string** | Either orderListId or origClientOrderId must be provided | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginOrderListV1Resp**](PmarginGetMarginOrderListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetMarginOrderV1

> PmarginGetMarginOrderV1Resp PmarginGetMarginOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Margin Account Order (USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetMarginOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginOrderV1`: PmarginGetMarginOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetMarginOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**PmarginGetMarginOrderV1Resp**](PmarginGetMarginOrderV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginGetMarginRepayLoanV1(context.Background()).Asset(asset).Timestamp(timestamp).TxId(txId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Archived(archived).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetMarginRepayLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetMarginRepayLoanV1`: PmarginGetMarginRepayLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetMarginRepayLoanV1`: %v\n", resp)
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


## PmarginGetPingV1

> map[string]interface{} PmarginGetPingV1(ctx).Execute()

Test Connectivity



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
	resp, r, err := apiClient.V1API.PmarginGetPingV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetPingV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetPingV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetPingV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetPingV1Request struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

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
	resp, r, err := apiClient.V1API.PmarginGetPortfolioInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetPortfolioInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetPortfolioInterestHistoryV1`: []PmarginGetPortfolioInterestHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetPortfolioInterestHistoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetPortfolioNegativeBalanceExchangeRecordV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetPortfolioNegativeBalanceExchangeRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetPortfolioNegativeBalanceExchangeRecordV1`: PmarginGetPortfolioNegativeBalanceExchangeRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetPortfolioNegativeBalanceExchangeRecordV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetRateLimitOrderV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetRateLimitOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetRateLimitOrderV1`: []PmarginGetRateLimitOrderV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetRateLimitOrderV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetRepayFuturesSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetRepayFuturesSwitchV1`: PmarginGetRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetRepayFuturesSwitchV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmAccountConfigV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmAccountConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAccountConfigV1`: PmarginGetUmAccountConfigV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmAccountConfigV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAccountV1`: PmarginGetUmAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmAccountV1`: %v\n", resp)
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


## PmarginGetUmAdlQuantileV1

> []PmarginGetUmAdlQuantileV1RespItem PmarginGetUmAdlQuantileV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

UM Position ADL Quantile Estimation(USER_DATA)



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
	resp, r, err := apiClient.V1API.PmarginGetUmAdlQuantileV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmAdlQuantileV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAdlQuantileV1`: []PmarginGetUmAdlQuantileV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmAdlQuantileV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAdlQuantileV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmAdlQuantileV1RespItem**](PmarginGetUmAdlQuantileV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmAllOrdersV1

> []PmarginGetUmAllOrdersV1RespItem PmarginGetUmAllOrdersV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All UM Orders(USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmAllOrdersV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAllOrdersV1`: []PmarginGetUmAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmAllOrdersV1RespItem**](PmarginGetUmAllOrdersV1RespItem.md)

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
	resp, r, err := apiClient.V1API.PmarginGetUmApiTradingStatusV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmApiTradingStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmApiTradingStatusV1`: PmarginGetUmApiTradingStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmApiTradingStatusV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmCommissionRateV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmCommissionRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmCommissionRateV1`: PmarginGetUmCommissionRateV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmCommissionRateV1`: %v\n", resp)
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


## PmarginGetUmConditionalAllOrdersV1

> []PmarginGetUmConditionalAllOrdersV1RespItem PmarginGetUmConditionalAllOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query All UM Conditional Orders(USER_DATA)



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
	strategyId := int64(789) // int64 |  (optional)
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmConditionalAllOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).StrategyId(strategyId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmConditionalAllOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalAllOrdersV1`: []PmarginGetUmConditionalAllOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmConditionalAllOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalAllOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **strategyId** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmConditionalAllOrdersV1RespItem**](PmarginGetUmConditionalAllOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmConditionalOpenOrderV1

> PmarginGetUmConditionalOpenOrderV1Resp PmarginGetUmConditionalOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query Current UM Open Conditional Order(USER_DATA)



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
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmConditionalOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmConditionalOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalOpenOrderV1`: PmarginGetUmConditionalOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmConditionalOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmConditionalOpenOrderV1Resp**](PmarginGetUmConditionalOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmConditionalOpenOrdersV1

> []PmarginGetUmConditionalOpenOrdersV1RespItem PmarginGetUmConditionalOpenOrdersV1(ctx).Execute()

Query All Current UM Open Conditional Orders(USER_DATA)



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
	resp, r, err := apiClient.V1API.PmarginGetUmConditionalOpenOrdersV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmConditionalOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalOpenOrdersV1`: []PmarginGetUmConditionalOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmConditionalOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalOpenOrdersV1Request struct via the builder pattern


### Return type

[**[]PmarginGetUmConditionalOpenOrdersV1RespItem**](PmarginGetUmConditionalOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmConditionalOrderHistoryV1

> PmarginGetUmConditionalOrderHistoryV1Resp PmarginGetUmConditionalOrderHistoryV1(ctx).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()

Query UM Conditional Order History(USER_DATA)



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
	strategyId := int64(789) // int64 |  (optional)
	newClientStrategyId := "newClientStrategyId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmConditionalOrderHistoryV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StrategyId(strategyId).NewClientStrategyId(newClientStrategyId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmConditionalOrderHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmConditionalOrderHistoryV1`: PmarginGetUmConditionalOrderHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmConditionalOrderHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmConditionalOrderHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **strategyId** | **int64** |  | 
 **newClientStrategyId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmConditionalOrderHistoryV1Resp**](PmarginGetUmConditionalOrderHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmFeeBurnV1

> PmarginGetUmFeeBurnV1Resp PmarginGetUmFeeBurnV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Futures BNB Burn Status (USER_DATA)



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
	resp, r, err := apiClient.V1API.PmarginGetUmFeeBurnV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmFeeBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmFeeBurnV1`: PmarginGetUmFeeBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmFeeBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmFeeBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmFeeBurnV1Resp**](PmarginGetUmFeeBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmForceOrdersV1

> []PmarginGetUmForceOrdersV1RespItem PmarginGetUmForceOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query User's UM Force Orders (USER_DATA)



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
	autoCloseType := "autoCloseType_example" // string | `LIQUIDATION` for liquidation orders, `ADL` for ADL orders. (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50; max 100. (optional) (default to 50)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmForceOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).AutoCloseType(autoCloseType).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmForceOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmForceOrdersV1`: []PmarginGetUmForceOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmForceOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmForceOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **autoCloseType** | **string** | &#x60;LIQUIDATION&#x60; for liquidation orders, &#x60;ADL&#x60; for ADL orders. | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50; max 100. | [default to 50]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**[]PmarginGetUmForceOrdersV1RespItem**](PmarginGetUmForceOrdersV1RespItem.md)

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
	resp, r, err := apiClient.V1API.PmarginGetUmIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmIncomeAsynIdV1`: PmarginGetUmIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmIncomeAsynIdV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmIncomeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmIncomeAsynV1`: PmarginGetUmIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmIncomeAsynV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmIncomeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).IncomeType(incomeType).StartTime(startTime).EndTime(endTime).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmIncomeV1`: []PmarginGetUmIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmIncomeV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmLeverageBracketV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmLeverageBracketV1`: []PmarginGetUmLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmLeverageBracketV1`: %v\n", resp)
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


## PmarginGetUmOpenOrderV1

> PmarginGetUmOpenOrderV1Resp PmarginGetUmOpenOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query Current UM Open Order(USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmOpenOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmOpenOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOpenOrderV1`: PmarginGetUmOpenOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmOpenOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOpenOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmOpenOrderV1Resp**](PmarginGetUmOpenOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOpenOrdersV1

> []PmarginGetUmOpenOrdersV1RespItem PmarginGetUmOpenOrdersV1(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query All Current UM Open Orders(USER_DATA)



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
	resp, r, err := apiClient.V1API.PmarginGetUmOpenOrdersV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmOpenOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOpenOrdersV1`: []PmarginGetUmOpenOrdersV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmOpenOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOpenOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmOpenOrdersV1RespItem**](PmarginGetUmOpenOrdersV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginGetUmOrderAmendmentV1

> []PmarginGetUmOrderAmendmentV1RespItem PmarginGetUmOrderAmendmentV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Query UM Modify Order History(TRADE)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | Timestamp in ms to get modification history from INCLUSIVE (optional)
	endTime := int64(789) // int64 | Timestamp in ms to get modification history until INCLUSIVE (optional)
	limit := int32(56) // int32 | Default 500, max 1000 (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmOrderAmendmentV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmOrderAmendmentV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderAmendmentV1`: []PmarginGetUmOrderAmendmentV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmOrderAmendmentV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOrderAmendmentV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | Timestamp in ms to get modification history from INCLUSIVE | 
 **endTime** | **int64** | Timestamp in ms to get modification history until INCLUSIVE | 
 **limit** | **int32** | Default 500, max 1000 | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmOrderAmendmentV1RespItem**](PmarginGetUmOrderAmendmentV1RespItem.md)

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
	resp, r, err := apiClient.V1API.PmarginGetUmOrderAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmOrderAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderAsynIdV1`: PmarginGetUmOrderAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmOrderAsynIdV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmOrderAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmOrderAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderAsynV1`: PmarginGetUmOrderAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmOrderAsynV1`: %v\n", resp)
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


## PmarginGetUmOrderV1

> PmarginGetUmOrderV1Resp PmarginGetUmOrderV1(ctx).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()

Query UM Order (USER_DATA)



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
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmOrderV1(context.Background()).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmOrderV1`: PmarginGetUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmOrderV1Resp**](PmarginGetUmOrderV1Resp.md)

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
	resp, r, err := apiClient.V1API.PmarginGetUmPositionRiskV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmPositionRiskV1`: []PmarginGetUmPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmPositionRiskV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmPositionSideDualV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmPositionSideDualV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmPositionSideDualV1`: PmarginGetUmPositionSideDualV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmPositionSideDualV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmSymbolConfigV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmSymbolConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmSymbolConfigV1`: []PmarginGetUmSymbolConfigV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmSymbolConfigV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmTradeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmTradeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmTradeAsynIdV1`: PmarginGetUmTradeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmTradeAsynIdV1`: %v\n", resp)
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
	resp, r, err := apiClient.V1API.PmarginGetUmTradeAsynV1(context.Background()).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmTradeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmTradeAsynV1`: PmarginGetUmTradeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmTradeAsynV1`: %v\n", resp)
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


## PmarginGetUmUserTradesV1

> []PmarginGetUmUserTradesV1RespItem PmarginGetUmUserTradesV1(ctx).Symbol(symbol).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

UM Account Trade List(USER_DATA)



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
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | Trade id to fetch from. Default gets most recent trades. (optional)
	limit := int32(56) // int32 | Default 500; max 1000. (optional) (default to 500)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginGetUmUserTradesV1(context.Background()).Symbol(symbol).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginGetUmUserTradesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmUserTradesV1`: []PmarginGetUmUserTradesV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginGetUmUserTradesV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmUserTradesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | Trade id to fetch from. Default gets most recent trades. | 
 **limit** | **int32** | Default 500; max 1000. | [default to 500]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginGetUmUserTradesV1RespItem**](PmarginGetUmUserTradesV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginUpdateCmOrderV1

> PmarginUpdateCmOrderV1Resp PmarginUpdateCmOrderV1(ctx).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify CM Order(TRADE)



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
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginUpdateCmOrderV1(context.Background()).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginUpdateCmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginUpdateCmOrderV1`: PmarginUpdateCmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginUpdateCmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginUpdateCmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginUpdateCmOrderV1Resp**](PmarginUpdateCmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginUpdateListenKeyV1

> map[string]interface{} PmarginUpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.V1API.PmarginUpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginUpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginUpdateListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginUpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginUpdateListenKeyV1Request struct via the builder pattern


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


## PmarginUpdateUmOrderV1

> PmarginUpdateUmOrderV1Resp PmarginUpdateUmOrderV1(ctx).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()

Modify UM Order(TRADE)



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
	price := "price_example" // string |  (default to "")
	quantity := "quantity_example" // string |  (default to "")
	side := "side_example" // string |  (default to "")
	symbol := "symbol_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	origClientOrderId := "origClientOrderId_example" // string |  (optional) (default to "")
	priceMatch := "priceMatch_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginUpdateUmOrderV1(context.Background()).Price(price).Quantity(quantity).Side(side).Symbol(symbol).Timestamp(timestamp).OrderId(orderId).OrigClientOrderId(origClientOrderId).PriceMatch(priceMatch).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginUpdateUmOrderV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginUpdateUmOrderV1`: PmarginUpdateUmOrderV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginUpdateUmOrderV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginUpdateUmOrderV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | **string** |  | [default to &quot;&quot;]
 **quantity** | **string** |  | [default to &quot;&quot;]
 **side** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **origClientOrderId** | **string** |  | [default to &quot;&quot;]
 **priceMatch** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginUpdateUmOrderV1Resp**](PmarginUpdateUmOrderV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

