# \MarketDataAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginGetMarginAllAssetsV1**](MarketDataAPI.md#MarginGetMarginAllAssetsV1) | **Get** /sapi/v1/margin/allAssets | Get All Margin Assets (MARKET_DATA)
[**MarginGetMarginAllPairsV1**](MarketDataAPI.md#MarginGetMarginAllPairsV1) | **Get** /sapi/v1/margin/allPairs | Get All Cross Margin Pairs (MARKET_DATA)
[**MarginGetMarginAvailableInventoryV1**](MarketDataAPI.md#MarginGetMarginAvailableInventoryV1) | **Get** /sapi/v1/margin/available-inventory | Query Margin Available Inventory(USER_DATA)
[**MarginGetMarginCrossMarginCollateralRatioV1**](MarketDataAPI.md#MarginGetMarginCrossMarginCollateralRatioV1) | **Get** /sapi/v1/margin/crossMarginCollateralRatio | Cross margin collateral ratio (MARKET_DATA)
[**MarginGetMarginDelistScheduleV1**](MarketDataAPI.md#MarginGetMarginDelistScheduleV1) | **Get** /sapi/v1/margin/delist-schedule | Get Delist Schedule (MARKET_DATA)
[**MarginGetMarginIsolatedAllPairsV1**](MarketDataAPI.md#MarginGetMarginIsolatedAllPairsV1) | **Get** /sapi/v1/margin/isolated/allPairs | Get All Isolated Margin Symbol(MARKET_DATA)
[**MarginGetMarginIsolatedMarginTierV1**](MarketDataAPI.md#MarginGetMarginIsolatedMarginTierV1) | **Get** /sapi/v1/margin/isolatedMarginTier | Query Isolated Margin Tier Data (USER_DATA)
[**MarginGetMarginLeverageBracketV1**](MarketDataAPI.md#MarginGetMarginLeverageBracketV1) | **Get** /sapi/v1/margin/leverageBracket | Query Liability Coin Leverage Bracket in Cross Margin Pro Mode(MARKET_DATA)
[**MarginGetMarginPriceIndexV1**](MarketDataAPI.md#MarginGetMarginPriceIndexV1) | **Get** /sapi/v1/margin/priceIndex | Query Margin PriceIndex (MARKET_DATA)



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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginAllAssetsV1(context.Background()).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginAllAssetsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllAssetsV1`: []MarginGetMarginAllAssetsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginAllAssetsV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginAllPairsV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginAllPairsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAllPairsV1`: []MarginGetMarginAllPairsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginAllPairsV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginAvailableInventoryV1(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginAvailableInventoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAvailableInventoryV1`: MarginGetMarginAvailableInventoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginAvailableInventoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginCrossMarginCollateralRatioV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginCrossMarginCollateralRatioV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginCrossMarginCollateralRatioV1`: []MarginGetMarginCrossMarginCollateralRatioV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginCrossMarginCollateralRatioV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginDelistScheduleV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginDelistScheduleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginDelistScheduleV1`: []MarginGetMarginDelistScheduleV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginDelistScheduleV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginIsolatedAllPairsV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginIsolatedAllPairsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedAllPairsV1`: []MarginGetMarginIsolatedAllPairsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginIsolatedAllPairsV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginIsolatedMarginTierV1(context.Background()).Symbol(symbol).Timestamp(timestamp).Tier(tier).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginIsolatedMarginTierV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedMarginTierV1`: []MarginGetMarginIsolatedMarginTierV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginIsolatedMarginTierV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginLeverageBracketV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginLeverageBracketV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginLeverageBracketV1`: []MarginGetMarginLeverageBracketV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginLeverageBracketV1`: %v\n", resp)
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
	resp, r, err := apiClient.MarketDataAPI.MarginGetMarginPriceIndexV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.MarginGetMarginPriceIndexV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginPriceIndexV1`: MarginGetMarginPriceIndexV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.MarginGetMarginPriceIndexV1`: %v\n", resp)
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

