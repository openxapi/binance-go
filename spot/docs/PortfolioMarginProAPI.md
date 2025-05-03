# \PortfolioMarginProAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortfolioAssetCollectionV1**](PortfolioMarginProAPI.md#CreatePortfolioAssetCollectionV1) | **Post** /sapi/v1/portfolio/asset-collection | Fund Collection by Asset(USER_DATA)
[**CreatePortfolioAutoCollectionV1**](PortfolioMarginProAPI.md#CreatePortfolioAutoCollectionV1) | **Post** /sapi/v1/portfolio/auto-collection | Fund Auto-collection(USER_DATA)
[**CreatePortfolioBnbTransferV1**](PortfolioMarginProAPI.md#CreatePortfolioBnbTransferV1) | **Post** /sapi/v1/portfolio/bnb-transfer | BNB transfer(USER_DATA)
[**CreatePortfolioMintV1**](PortfolioMarginProAPI.md#CreatePortfolioMintV1) | **Post** /sapi/v1/portfolio/mint | Mint BFUSD for Portfolio Margin(TRADE)
[**CreatePortfolioRedeemV1**](PortfolioMarginProAPI.md#CreatePortfolioRedeemV1) | **Post** /sapi/v1/portfolio/redeem | Redeem BFUSD for Portfolio Margin(TRADE)
[**CreatePortfolioRepayFuturesNegativeBalanceV1**](PortfolioMarginProAPI.md#CreatePortfolioRepayFuturesNegativeBalanceV1) | **Post** /sapi/v1/portfolio/repay-futures-negative-balance | Repay futures Negative Balance(USER_DATA)
[**CreatePortfolioRepayFuturesSwitchV1**](PortfolioMarginProAPI.md#CreatePortfolioRepayFuturesSwitchV1) | **Post** /sapi/v1/portfolio/repay-futures-switch | Change Auto-repay-futures Status(TRADE)
[**CreatePortfolioRepayV1**](PortfolioMarginProAPI.md#CreatePortfolioRepayV1) | **Post** /sapi/v1/portfolio/repay | Portfolio Margin Pro Bankruptcy Loan Repay
[**GetPortfolioAccountV1**](PortfolioMarginProAPI.md#GetPortfolioAccountV1) | **Get** /sapi/v1/portfolio/account | Get Portfolio Margin Pro Account Info(USER_DATA)
[**GetPortfolioAccountV2**](PortfolioMarginProAPI.md#GetPortfolioAccountV2) | **Get** /sapi/v2/portfolio/account | Get Portfolio Margin Pro SPAN Account Info(USER_DATA)
[**GetPortfolioAssetIndexPriceV1**](PortfolioMarginProAPI.md#GetPortfolioAssetIndexPriceV1) | **Get** /sapi/v1/portfolio/asset-index-price | Query Portfolio Margin Asset Index Price (MARKET_DATA)
[**GetPortfolioBalanceV1**](PortfolioMarginProAPI.md#GetPortfolioBalanceV1) | **Get** /sapi/v1/portfolio/balance | Get Portfolio Margin Pro Account Balance(USER_DATA)
[**GetPortfolioCollateralRateV1**](PortfolioMarginProAPI.md#GetPortfolioCollateralRateV1) | **Get** /sapi/v1/portfolio/collateralRate | Portfolio Margin Collateral Rate(MARKET_DATA)
[**GetPortfolioCollateralRateV2**](PortfolioMarginProAPI.md#GetPortfolioCollateralRateV2) | **Get** /sapi/v2/portfolio/collateralRate | Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)
[**GetPortfolioInterestHistoryV1**](PortfolioMarginProAPI.md#GetPortfolioInterestHistoryV1) | **Get** /sapi/v1/portfolio/interest-history | Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)
[**GetPortfolioMarginAssetLeverageV1**](PortfolioMarginProAPI.md#GetPortfolioMarginAssetLeverageV1) | **Get** /sapi/v1/portfolio/margin-asset-leverage | Get Portfolio Margin Asset Leverage(USER_DATA)
[**GetPortfolioPmLoanHistoryV1**](PortfolioMarginProAPI.md#GetPortfolioPmLoanHistoryV1) | **Get** /sapi/v1/portfolio/pmLoan-history | Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)
[**GetPortfolioPmLoanV1**](PortfolioMarginProAPI.md#GetPortfolioPmLoanV1) | **Get** /sapi/v1/portfolio/pmLoan | Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)
[**GetPortfolioRepayFuturesSwitchV1**](PortfolioMarginProAPI.md#GetPortfolioRepayFuturesSwitchV1) | **Get** /sapi/v1/portfolio/repay-futures-switch | Get Auto-repay-futures Status(USER_DATA)



## CreatePortfolioAssetCollectionV1

> CreatePortfolioAssetCollectionV1Resp CreatePortfolioAssetCollectionV1(ctx).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fund Collection by Asset(USER_DATA)



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
	asset := "asset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioAssetCollectionV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioAssetCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioAssetCollectionV1`: CreatePortfolioAssetCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioAssetCollectionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioAssetCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioAssetCollectionV1Resp**](CreatePortfolioAssetCollectionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortfolioAutoCollectionV1

> CreatePortfolioAutoCollectionV1Resp CreatePortfolioAutoCollectionV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fund Auto-collection(USER_DATA)



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
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioAutoCollectionV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioAutoCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioAutoCollectionV1`: CreatePortfolioAutoCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioAutoCollectionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioAutoCollectionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioAutoCollectionV1Resp**](CreatePortfolioAutoCollectionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortfolioBnbTransferV1

> CreatePortfolioBnbTransferV1Resp CreatePortfolioBnbTransferV1(ctx).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()

BNB transfer(USER_DATA)



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
	timestamp := int64(789) // int64 | 
	transferSide := "transferSide_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioBnbTransferV1(context.Background()).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioBnbTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioBnbTransferV1`: CreatePortfolioBnbTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioBnbTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioBnbTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **transferSide** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioBnbTransferV1Resp**](CreatePortfolioBnbTransferV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortfolioMintV1

> CreatePortfolioMintV1Resp CreatePortfolioMintV1(ctx).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Mint BFUSD for Portfolio Margin(TRADE)



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
	fromAsset := "fromAsset_example" // string |  (default to "")
	targetAsset := "targetAsset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioMintV1(context.Background()).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioMintV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioMintV1`: CreatePortfolioMintV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioMintV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioMintV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **fromAsset** | **string** |  | [default to &quot;&quot;]
 **targetAsset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioMintV1Resp**](CreatePortfolioMintV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortfolioRedeemV1

> CreatePortfolioRedeemV1Resp CreatePortfolioRedeemV1(ctx).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Redeem BFUSD for Portfolio Margin(TRADE)



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
	fromAsset := "fromAsset_example" // string |  (default to "")
	targetAsset := "targetAsset_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioRedeemV1(context.Background()).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioRedeemV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioRedeemV1`: CreatePortfolioRedeemV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioRedeemV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioRedeemV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **fromAsset** | **string** |  | [default to &quot;&quot;]
 **targetAsset** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioRedeemV1Resp**](CreatePortfolioRedeemV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortfolioRepayFuturesNegativeBalanceV1

> CreatePortfolioRepayFuturesNegativeBalanceV1Resp CreatePortfolioRepayFuturesNegativeBalanceV1(ctx).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()

Repay futures Negative Balance(USER_DATA)



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
	from := "from_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioRepayFuturesNegativeBalanceV1(context.Background()).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioRepayFuturesNegativeBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioRepayFuturesNegativeBalanceV1`: CreatePortfolioRepayFuturesNegativeBalanceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioRepayFuturesNegativeBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioRepayFuturesNegativeBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **from** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioRepayFuturesNegativeBalanceV1Resp**](CreatePortfolioRepayFuturesNegativeBalanceV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortfolioRepayFuturesSwitchV1

> CreatePortfolioRepayFuturesSwitchV1Resp CreatePortfolioRepayFuturesSwitchV1(ctx).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Change Auto-repay-futures Status(TRADE)



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
	autoRepay := "autoRepay_example" // string |  (default to "true")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioRepayFuturesSwitchV1(context.Background()).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioRepayFuturesSwitchV1`: CreatePortfolioRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioRepayFuturesSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioRepayFuturesSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoRepay** | **string** |  | [default to &quot;true&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioRepayFuturesSwitchV1Resp**](CreatePortfolioRepayFuturesSwitchV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePortfolioRepayV1

> CreatePortfolioRepayV1Resp CreatePortfolioRepayV1(ctx).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()

Portfolio Margin Pro Bankruptcy Loan Repay



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
	from := "from_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.CreatePortfolioRepayV1(context.Background()).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.CreatePortfolioRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePortfolioRepayV1`: CreatePortfolioRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.CreatePortfolioRepayV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortfolioRepayV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **from** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreatePortfolioRepayV1Resp**](CreatePortfolioRepayV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioAccountV1

> GetPortfolioAccountV1Resp GetPortfolioAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro Account Info(USER_DATA)



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
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioAccountV1`: GetPortfolioAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetPortfolioAccountV1Resp**](GetPortfolioAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioAccountV2

> GetPortfolioAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro SPAN Account Info(USER_DATA)



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
	r, err := apiClient.PortfolioMarginProAPI.GetPortfolioAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioAssetIndexPriceV1

> []GetPortfolioAssetIndexPriceV1RespItem GetPortfolioAssetIndexPriceV1(ctx).Asset(asset).Execute()

Query Portfolio Margin Asset Index Price (MARKET_DATA)



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
	asset := "asset_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioAssetIndexPriceV1(context.Background()).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioAssetIndexPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioAssetIndexPriceV1`: []GetPortfolioAssetIndexPriceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioAssetIndexPriceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioAssetIndexPriceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asset** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]GetPortfolioAssetIndexPriceV1RespItem**](GetPortfolioAssetIndexPriceV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioBalanceV1

> []GetPortfolioBalanceV1RespItem GetPortfolioBalanceV1(ctx).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro Account Balance(USER_DATA)



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
	asset := "asset_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioBalanceV1(context.Background()).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioBalanceV1`: []GetPortfolioBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioBalanceV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioBalanceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetPortfolioBalanceV1RespItem**](GetPortfolioBalanceV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioCollateralRateV1

> []GetPortfolioCollateralRateV1RespItem GetPortfolioCollateralRateV1(ctx).Execute()

Portfolio Margin Collateral Rate(MARKET_DATA)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioCollateralRateV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioCollateralRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioCollateralRateV1`: []GetPortfolioCollateralRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioCollateralRateV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioCollateralRateV1Request struct via the builder pattern


### Return type

[**[]GetPortfolioCollateralRateV1RespItem**](GetPortfolioCollateralRateV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioCollateralRateV2

> []GetPortfolioCollateralRateV2RespItem GetPortfolioCollateralRateV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)



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
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioCollateralRateV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioCollateralRateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioCollateralRateV2`: []GetPortfolioCollateralRateV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioCollateralRateV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioCollateralRateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetPortfolioCollateralRateV2RespItem**](GetPortfolioCollateralRateV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioInterestHistoryV1

> []GetPortfolioInterestHistoryV1RespItem GetPortfolioInterestHistoryV1(ctx).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)



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
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioInterestHistoryV1`: []GetPortfolioInterestHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioInterestHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioInterestHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetPortfolioInterestHistoryV1RespItem**](GetPortfolioInterestHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioMarginAssetLeverageV1

> []GetPortfolioMarginAssetLeverageV1RespItem GetPortfolioMarginAssetLeverageV1(ctx).Execute()

Get Portfolio Margin Asset Leverage(USER_DATA)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioMarginAssetLeverageV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioMarginAssetLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioMarginAssetLeverageV1`: []GetPortfolioMarginAssetLeverageV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioMarginAssetLeverageV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioMarginAssetLeverageV1Request struct via the builder pattern


### Return type

[**[]GetPortfolioMarginAssetLeverageV1RespItem**](GetPortfolioMarginAssetLeverageV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioPmLoanHistoryV1

> GetPortfolioPmLoanHistoryV1Resp GetPortfolioPmLoanHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)



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
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10 Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioPmLoanHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioPmLoanHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioPmLoanHistoryV1`: GetPortfolioPmLoanHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioPmLoanHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioPmLoanHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10 Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetPortfolioPmLoanHistoryV1Resp**](GetPortfolioPmLoanHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioPmLoanV1

> GetPortfolioPmLoanV1Resp GetPortfolioPmLoanV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)



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
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioPmLoanV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioPmLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioPmLoanV1`: GetPortfolioPmLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioPmLoanV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioPmLoanV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetPortfolioPmLoanV1Resp**](GetPortfolioPmLoanV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolioRepayFuturesSwitchV1

> GetPortfolioRepayFuturesSwitchV1Resp GetPortfolioRepayFuturesSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Auto-repay-futures Status(USER_DATA)



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
	resp, r, err := apiClient.PortfolioMarginProAPI.GetPortfolioRepayFuturesSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioMarginProAPI.GetPortfolioRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolioRepayFuturesSwitchV1`: GetPortfolioRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `PortfolioMarginProAPI.GetPortfolioRepayFuturesSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioRepayFuturesSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetPortfolioRepayFuturesSwitchV1Resp**](GetPortfolioRepayFuturesSwitchV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

