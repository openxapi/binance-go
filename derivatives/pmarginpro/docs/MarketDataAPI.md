# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginproGetPortfolioAssetIndexPriceV1**](MarketDataAPI.md#PmarginproGetPortfolioAssetIndexPriceV1) | **Get** /sapi/v1/portfolio/asset-index-price | Query Portfolio Margin Asset Index Price (MARKET_DATA)
[**PmarginproGetPortfolioCollateralRateV1**](MarketDataAPI.md#PmarginproGetPortfolioCollateralRateV1) | **Get** /sapi/v1/portfolio/collateralRate | Portfolio Margin Collateral Rate(MARKET_DATA)
[**PmarginproGetPortfolioCollateralRateV2**](MarketDataAPI.md#PmarginproGetPortfolioCollateralRateV2) | **Get** /sapi/v2/portfolio/collateralRate | Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)
[**PmarginproGetPortfolioMarginAssetLeverageV1**](MarketDataAPI.md#PmarginproGetPortfolioMarginAssetLeverageV1) | **Get** /sapi/v1/portfolio/margin-asset-leverage | Get Portfolio Margin Asset Leverage(USER_DATA)



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
	resp, r, err := apiClient.MarketDataAPI.PmarginproGetPortfolioAssetIndexPriceV1(context.Background()).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.PmarginproGetPortfolioAssetIndexPriceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioAssetIndexPriceV1`: []PmarginproGetPortfolioAssetIndexPriceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.PmarginproGetPortfolioAssetIndexPriceV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.PmarginproGetPortfolioCollateralRateV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.PmarginproGetPortfolioCollateralRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioCollateralRateV1`: []PmarginproGetPortfolioCollateralRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.PmarginproGetPortfolioCollateralRateV1`: %v\n", resp)
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


## PmarginproGetPortfolioCollateralRateV2

> []PmarginproGetPortfolioCollateralRateV2RespItem PmarginproGetPortfolioCollateralRateV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Portfolio Margin Pro Tiered Collateral Rate(USER_DATA)



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
	resp, r, err := apiClient.MarketDataAPI.PmarginproGetPortfolioCollateralRateV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.PmarginproGetPortfolioCollateralRateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioCollateralRateV2`: []PmarginproGetPortfolioCollateralRateV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.PmarginproGetPortfolioCollateralRateV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginproGetPortfolioCollateralRateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]PmarginproGetPortfolioCollateralRateV2RespItem**](PmarginproGetPortfolioCollateralRateV2RespItem.md)

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
	resp, r, err := apiClient.MarketDataAPI.PmarginproGetPortfolioMarginAssetLeverageV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.PmarginproGetPortfolioMarginAssetLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginproGetPortfolioMarginAssetLeverageV1`: []PmarginproGetPortfolioMarginAssetLeverageV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.PmarginproGetPortfolioMarginAssetLeverageV1`: %v\n", resp)
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

