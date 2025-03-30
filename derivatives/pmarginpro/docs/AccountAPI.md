# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginproCreatePortfolioAssetCollectionV1**](AccountAPI.md#PmarginproCreatePortfolioAssetCollectionV1) | **Post** /sapi/v1/portfolio/asset-collection | Fund Collection by Asset(USER_DATA)
[**PmarginproCreatePortfolioAutoCollectionV1**](AccountAPI.md#PmarginproCreatePortfolioAutoCollectionV1) | **Post** /sapi/v1/portfolio/auto-collection | Fund Auto-collection(USER_DATA)
[**PmarginproCreatePortfolioBnbTransferV1**](AccountAPI.md#PmarginproCreatePortfolioBnbTransferV1) | **Post** /sapi/v1/portfolio/bnb-transfer | BNB transfer(USER_DATA)
[**PmarginproCreatePortfolioMintV1**](AccountAPI.md#PmarginproCreatePortfolioMintV1) | **Post** /sapi/v1/portfolio/mint | Mint BFUSD for Portfolio Margin(TRADE)
[**PmarginproCreatePortfolioRedeemV1**](AccountAPI.md#PmarginproCreatePortfolioRedeemV1) | **Post** /sapi/v1/portfolio/redeem | Redeem BFUSD for Portfolio Margin(TRADE)
[**PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1**](AccountAPI.md#PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1) | **Post** /sapi/v1/portfolio/repay-futures-negative-balance | Repay futures Negative Balance(USER_DATA)
[**PmarginproCreatePortfolioRepayFuturesSwitchV1**](AccountAPI.md#PmarginproCreatePortfolioRepayFuturesSwitchV1) | **Post** /sapi/v1/portfolio/repay-futures-switch | Change Auto-repay-futures Status(TRADE)
[**PmarginproCreatePortfolioRepayV1**](AccountAPI.md#PmarginproCreatePortfolioRepayV1) | **Post** /sapi/v1/portfolio/repay | Portfolio Margin Pro Bankruptcy Loan Repay
[**PmarginproGetPortfolioAccountV1**](AccountAPI.md#PmarginproGetPortfolioAccountV1) | **Get** /sapi/v1/portfolio/account | Get Portfolio Margin Pro Account Info(USER_DATA)
[**PmarginproGetPortfolioAccountV2**](AccountAPI.md#PmarginproGetPortfolioAccountV2) | **Get** /sapi/v2/portfolio/account | Get Portfolio Margin Pro SPAN Account Info(USER_DATA)
[**PmarginproGetPortfolioBalanceV1**](AccountAPI.md#PmarginproGetPortfolioBalanceV1) | **Get** /sapi/v1/portfolio/balance | Get Portfolio Margin Pro Account Balance(USER_DATA)
[**PmarginproGetPortfolioInterestHistoryV1**](AccountAPI.md#PmarginproGetPortfolioInterestHistoryV1) | **Get** /sapi/v1/portfolio/interest-history | Query Portfolio Margin Pro Negative Balance Interest History(USER_DATA)
[**PmarginproGetPortfolioPmLoanHistoryV1**](AccountAPI.md#PmarginproGetPortfolioPmLoanHistoryV1) | **Get** /sapi/v1/portfolio/pmLoan-history | Query Portfolio Margin Pro Bankruptcy Loan Repay History(USER_DATA)
[**PmarginproGetPortfolioPmLoanV1**](AccountAPI.md#PmarginproGetPortfolioPmLoanV1) | **Get** /sapi/v1/portfolio/pmLoan | Query Portfolio Margin Pro Bankruptcy Loan Amount(USER_DATA)
[**PmarginproGetPortfolioRepayFuturesSwitchV1**](AccountAPI.md#PmarginproGetPortfolioRepayFuturesSwitchV1) | **Get** /sapi/v1/portfolio/repay-futures-switch | Get Auto-repay-futures Status(USER_DATA)



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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioAssetCollectionV1(context.Background()).Asset(asset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioAssetCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioAssetCollectionV1`: PmarginproCreatePortfolioAssetCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioAssetCollectionV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioAutoCollectionV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioAutoCollectionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioAutoCollectionV1`: PmarginproCreatePortfolioAutoCollectionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioAutoCollectionV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioBnbTransferV1(context.Background()).Amount(amount).Timestamp(timestamp).TransferSide(transferSide).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioBnbTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioBnbTransferV1`: PmarginproCreatePortfolioBnbTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioBnbTransferV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioMintV1(context.Background()).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioMintV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioMintV1`: PmarginproCreatePortfolioMintV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioMintV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioRedeemV1(context.Background()).Amount(amount).FromAsset(fromAsset).TargetAsset(targetAsset).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioRedeemV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRedeemV1`: PmarginproCreatePortfolioRedeemV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioRedeemV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1(context.Background()).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1`: PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioRepayFuturesNegativeBalanceV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioRepayFuturesSwitchV1(context.Background()).AutoRepay(autoRepay).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRepayFuturesSwitchV1`: PmarginproCreatePortfolioRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioRepayFuturesSwitchV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproCreatePortfolioRepayV1(context.Background()).Timestamp(timestamp).From(from).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproCreatePortfolioRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproCreatePortfolioRepayV1`: PmarginproCreatePortfolioRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproCreatePortfolioRepayV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproGetPortfolioAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproGetPortfolioAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioAccountV1`: PmarginproGetPortfolioAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproGetPortfolioAccountV1`: %v\n", resp)
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


## PmarginproGetPortfolioAccountV2

> PmarginproGetPortfolioAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Portfolio Margin Pro SPAN Account Info(USER_DATA)



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
	r, err := apiClient.AccountAPI.PmarginproGetPortfolioAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproGetPortfolioAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioAccountV2Request struct via the builder pattern


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
	resp, r, err := apiClient.AccountAPI.PmarginproGetPortfolioBalanceV1(context.Background()).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproGetPortfolioBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioBalanceV1`: []PmarginproGetPortfolioBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproGetPortfolioBalanceV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproGetPortfolioInterestHistoryV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproGetPortfolioInterestHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioInterestHistoryV1`: []PmarginproGetPortfolioInterestHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproGetPortfolioInterestHistoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproGetPortfolioPmLoanHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproGetPortfolioPmLoanHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioPmLoanHistoryV1`: PmarginproGetPortfolioPmLoanHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproGetPortfolioPmLoanHistoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproGetPortfolioPmLoanV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproGetPortfolioPmLoanV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioPmLoanV1`: PmarginproGetPortfolioPmLoanV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproGetPortfolioPmLoanV1`: %v\n", resp)
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
	resp, r, err := apiClient.AccountAPI.PmarginproGetPortfolioRepayFuturesSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PmarginproGetPortfolioRepayFuturesSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioRepayFuturesSwitchV1`: PmarginproGetPortfolioRepayFuturesSwitchV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PmarginproGetPortfolioRepayFuturesSwitchV1`: %v\n", resp)
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

