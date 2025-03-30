# \V1API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginproCreatePortfolioAssetCollectionV1**](V1API.md#PmarginproCreatePortfolioAssetCollectionV1) | **Post** /sapi/v1/portfolio/asset-collection | Fund Collection by Asset(USER_DATA)
[**PmarginproCreatePortfolioAutoCollectionV1**](V1API.md#PmarginproCreatePortfolioAutoCollectionV1) | **Post** /sapi/v1/portfolio/auto-collection | Fund Auto-collection(USER_DATA)
[**PmarginproCreatePortfolioBnbTransferV1**](V1API.md#PmarginproCreatePortfolioBnbTransferV1) | **Post** /sapi/v1/portfolio/bnb-transfer | BNB transfer(USER_DATA)
[**PmarginproCreatePortfolioMintV1**](V1API.md#PmarginproCreatePortfolioMintV1) | **Post** /sapi/v1/portfolio/mint | Mint BFUSD for Portfolio Margin(TRADE)
[**PmarginproCreatePortfolioRedeemV1**](V1API.md#PmarginproCreatePortfolioRedeemV1) | **Post** /sapi/v1/portfolio/redeem | Redeem BFUSD for Portfolio Margin(TRADE)
[**PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1**](V1API.md#PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1) | **Post** /sapi/v1/portfolio/repay-futures-negative-balance | Repay futures Negative Balance(USER_DATA)
[**PmarginproCreatePortfolioRepayFuturesSwitchV1**](V1API.md#PmarginproCreatePortfolioRepayFuturesSwitchV1) | **Post** /sapi/v1/portfolio/repay-futures-switch | Change Auto-repay-futures Status(TRADE)
[**PmarginproCreatePortfolioRepayV1**](V1API.md#PmarginproCreatePortfolioRepayV1) | **Post** /sapi/v1/portfolio/repay | Portfolio Margin Pro Bankruptcy Loan Repay
[**PmarginproGetPortfolioAccountV1**](V1API.md#PmarginproGetPortfolioAccountV1) | **Get** /sapi/v1/portfolio/account | Get Portfolio Margin Pro Account Info(USER_DATA)
[**PmarginproGetPortfolioAssetIndexPriceV1**](V1API.md#PmarginproGetPortfolioAssetIndexPriceV1) | **Get** /sapi/v1/portfolio/asset-index-price | Query Portfolio Margin Asset Index Price (MARKET_DATA)
[**PmarginproGetPortfolioBalanceV1**](V1API.md#PmarginproGetPortfolioBalanceV1) | **Get** /sapi/v1/portfolio/balance | Get Portfolio Margin Pro Account Balance(USER_DATA)
[**PmarginproGetPortfolioCollateralRateV1**](V1API.md#PmarginproGetPortfolioCollateralRateV1) | **Get** /sapi/v1/portfolio/collateralRate | Portfolio Margin Collateral Rate(MARKET_DATA)
[**PmarginproGetPortfolioInterestHistoryV1**](V1API.md#PmarginproGetPortfolioInterestHistoryV1) | **Get** /sapi/v1/portfolio/interest-history | Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)
[**PmarginproGetPortfolioMarginAssetLeverageV1**](V1API.md#PmarginproGetPortfolioMarginAssetLeverageV1) | **Get** /sapi/v1/portfolio/margin-asset-leverage | Get Portfolio Margin Asset Leverage(USER_DATA)
[**PmarginproGetPortfolioPmLoanHistoryV1**](V1API.md#PmarginproGetPortfolioPmLoanHistoryV1) | **Get** /sapi/v1/portfolio/pmLoan-history | Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)
[**PmarginproGetPortfolioPmLoanV1**](V1API.md#PmarginproGetPortfolioPmLoanV1) | **Get** /sapi/v1/portfolio/pmLoan | Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)
[**PmarginproGetPortfolioRepayFuturesSwitchV1**](V1API.md#PmarginproGetPortfolioRepayFuturesSwitchV1) | **Get** /sapi/v1/portfolio/repay-futures-switch | Get Auto-repay-futures Status(USER_DATA)



## PmarginproCreatePortfolioAssetCollectionV1

> PmarginproCreatePortfolioAssetCollectionV1Resp PmarginproCreatePortfolioAssetCollectionV1(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fund Collection by Asset(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioAssetCollectionV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioAssetCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioAssetCollectionV1`: PmarginproCreatePortfolioAssetCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioAssetCollectionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioAssetCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioAssetCollectionV1Resp**](PmarginproCreatePortfolioAssetCollectionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproCreatePortfolioAutoCollectionV1

> PmarginproCreatePortfolioAutoCollectionV1Resp PmarginproCreatePortfolioAutoCollectionV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fund Auto-collection(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioAutoCollectionV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioAutoCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioAutoCollectionV1`: PmarginproCreatePortfolioAutoCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioAutoCollectionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioAutoCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioAutoCollectionV1Resp**](PmarginproCreatePortfolioAutoCollectionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproCreatePortfolioBnbTransferV1

> PmarginproCreatePortfolioBnbTransferV1Resp PmarginproCreatePortfolioBnbTransferV1(ctx).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()

BNB transfer(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	transferSide := "transferSide_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioBnbTransferV1(context.Background()).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioBnbTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioBnbTransferV1`: PmarginproCreatePortfolioBnbTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioBnbTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioBnbTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **transferSide** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioBnbTransferV1Resp**](PmarginproCreatePortfolioBnbTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproCreatePortfolioMintV1

> PmarginproCreatePortfolioMintV1Resp PmarginproCreatePortfolioMintV1(ctx).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Mint BFUSD for Portfolio Margin(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	fromAsset := "fromAsset_example" // string |  (default to "")
	targetAsset := "targetAsset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioMintV1(context.Background()).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioMintV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioMintV1`: PmarginproCreatePortfolioMintV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioMintV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioMintV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **fromAsset** | **string** |  | [default to &quot;&quot;]
 **targetAsset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioMintV1Resp**](PmarginproCreatePortfolioMintV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproCreatePortfolioRedeemV1

> PmarginproCreatePortfolioRedeemV1Resp PmarginproCreatePortfolioRedeemV1(ctx).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Redeem BFUSD for Portfolio Margin(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	fromAsset := "fromAsset_example" // string |  (default to "")
	targetAsset := "targetAsset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioRedeemV1(context.Background()).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioRedeemV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRedeemV1`: PmarginproCreatePortfolioRedeemV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioRedeemV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioRedeemV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **fromAsset** | **string** |  | [default to &quot;&quot;]
 **targetAsset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioRedeemV1Resp**](PmarginproCreatePortfolioRedeemV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1

> PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1Resp PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1(ctx).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()

Repay futures Negative Balance(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	from := "from_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1(context.Background()).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1`: PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioRepayFuturesNegativeBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **from** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1Resp**](PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproCreatePortfolioRepayFuturesSwitchV1

> PmarginproCreatePortfolioRepayFuturesSwitchV1Resp PmarginproCreatePortfolioRepayFuturesSwitchV1(ctx).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Auto-repay-futures Status(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	autoRepay := "autoRepay_example" // string |  (default to "true")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioRepayFuturesSwitchV1(context.Background()).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRepayFuturesSwitchV1`: PmarginproCreatePortfolioRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioRepayFuturesSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioRepayFuturesSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoRepay** | **string** |  | [default to &quot;true&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioRepayFuturesSwitchV1Resp**](PmarginproCreatePortfolioRepayFuturesSwitchV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproCreatePortfolioRepayV1

> PmarginproCreatePortfolioRepayV1Resp PmarginproCreatePortfolioRepayV1(ctx).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()

Portfolio Margin Pro Bankruptcy Loan Repay



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	from := "from_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproCreatePortfolioRepayV1(context.Background()).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproCreatePortfolioRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRepayV1`: PmarginproCreatePortfolioRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproCreatePortfolioRepayV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproCreatePortfolioRepayV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **from** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproCreatePortfolioRepayV1Resp**](PmarginproCreatePortfolioRepayV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioAccountV1

> PmarginproGetPortfolioAccountV1Resp PmarginproGetPortfolioAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro Account Info(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioAccountV1`: PmarginproGetPortfolioAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproGetPortfolioAccountV1Resp**](PmarginproGetPortfolioAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioAssetIndexPriceV1

> []PmarginproGetPortfolioAssetIndexPriceV1RespItem PmarginproGetPortfolioAssetIndexPriceV1(ctx).Asset(asset).Execute()

Query Portfolio Margin Asset Index Price (MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	asset := "asset_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioAssetIndexPriceV1(context.Background()).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioAssetIndexPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioAssetIndexPriceV1`: []PmarginproGetPortfolioAssetIndexPriceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioAssetIndexPriceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioAssetIndexPriceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]PmarginproGetPortfolioAssetIndexPriceV1RespItem**](PmarginproGetPortfolioAssetIndexPriceV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioBalanceV1

> []PmarginproGetPortfolioBalanceV1RespItem PmarginproGetPortfolioBalanceV1(ctx).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro Account Balance(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioBalanceV1(context.Background()).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioBalanceV1`: []PmarginproGetPortfolioBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginproGetPortfolioBalanceV1RespItem**](PmarginproGetPortfolioBalanceV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioCollateralRateV1

> []PmarginproGetPortfolioCollateralRateV1RespItem PmarginproGetPortfolioCollateralRateV1(ctx).Execute()

Portfolio Margin Collateral Rate(MARKET_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioCollateralRateV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioCollateralRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioCollateralRateV1`: []PmarginproGetPortfolioCollateralRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioCollateralRateV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioCollateralRateV1Request struct via the builder pattern


### Return type

[**[]PmarginproGetPortfolioCollateralRateV1RespItem**](PmarginproGetPortfolioCollateralRateV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioInterestHistoryV1

> []PmarginproGetPortfolioInterestHistoryV1RespItem PmarginproGetPortfolioInterestHistoryV1(ctx).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
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
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioInterestHistoryV1`: []PmarginproGetPortfolioInterestHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioInterestHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioInterestHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginproGetPortfolioInterestHistoryV1RespItem**](PmarginproGetPortfolioInterestHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioMarginAssetLeverageV1

> []PmarginproGetPortfolioMarginAssetLeverageV1RespItem PmarginproGetPortfolioMarginAssetLeverageV1(ctx).Execute()

Get Portfolio Margin Asset Leverage(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioMarginAssetLeverageV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioMarginAssetLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioMarginAssetLeverageV1`: []PmarginproGetPortfolioMarginAssetLeverageV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioMarginAssetLeverageV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioMarginAssetLeverageV1Request struct via the builder pattern


### Return type

[**[]PmarginproGetPortfolioMarginAssetLeverageV1RespItem**](PmarginproGetPortfolioMarginAssetLeverageV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioPmLoanHistoryV1

> PmarginproGetPortfolioPmLoanHistoryV1Resp PmarginproGetPortfolioPmLoanHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioPmLoanHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioPmLoanHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioPmLoanHistoryV1`: PmarginproGetPortfolioPmLoanHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioPmLoanHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioPmLoanHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproGetPortfolioPmLoanHistoryV1Resp**](PmarginproGetPortfolioPmLoanHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioPmLoanV1

> PmarginproGetPortfolioPmLoanV1Resp PmarginproGetPortfolioPmLoanV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioPmLoanV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioPmLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioPmLoanV1`: PmarginproGetPortfolioPmLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioPmLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioPmLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproGetPortfolioPmLoanV1Resp**](PmarginproGetPortfolioPmLoanV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginproGetPortfolioRepayFuturesSwitchV1

> PmarginproGetPortfolioRepayFuturesSwitchV1Resp PmarginproGetPortfolioRepayFuturesSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Auto-repay-futures Status(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmarginpro"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1API.PmarginproGetPortfolioRepayFuturesSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1API.PmarginproGetPortfolioRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioRepayFuturesSwitchV1`: PmarginproGetPortfolioRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `V1API.PmarginproGetPortfolioRepayFuturesSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioRepayFuturesSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginproGetPortfolioRepayFuturesSwitchV1Resp**](PmarginproGetPortfolioRepayFuturesSwitchV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

