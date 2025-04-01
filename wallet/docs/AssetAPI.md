# \AssetAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletCreateAssetDustBtcV1**](AssetAPI.md#WalletCreateAssetDustBtcV1) | **Post** /sapi/v1/asset/dust-btc | Get Assets That Can Be Converted Into BNB (USER_DATA)
[**WalletCreateAssetDustV1**](AssetAPI.md#WalletCreateAssetDustV1) | **Post** /sapi/v1/asset/dust | Dust Transfer (USER_DATA)
[**WalletCreateAssetGetFundingAssetV1**](AssetAPI.md#WalletCreateAssetGetFundingAssetV1) | **Post** /sapi/v1/asset/get-funding-asset | Funding Wallet (USER_DATA)
[**WalletCreateAssetGetUserAssetV3**](AssetAPI.md#WalletCreateAssetGetUserAssetV3) | **Post** /sapi/v3/asset/getUserAsset | User Asset (USER_DATA)
[**WalletCreateAssetTransferV1**](AssetAPI.md#WalletCreateAssetTransferV1) | **Post** /sapi/v1/asset/transfer | User Universal Transfer (USER_DATA)
[**WalletCreateBnbBurnV1**](AssetAPI.md#WalletCreateBnbBurnV1) | **Post** /sapi/v1/bnbBurn | Toggle BNB Burn On Spot Trade And Margin Interest (USER_DATA)
[**WalletGetAssetAssetDetailV1**](AssetAPI.md#WalletGetAssetAssetDetailV1) | **Get** /sapi/v1/asset/assetDetail | Asset Detail (USER_DATA)
[**WalletGetAssetAssetDividendV1**](AssetAPI.md#WalletGetAssetAssetDividendV1) | **Get** /sapi/v1/asset/assetDividend | Asset Dividend Record (USER_DATA)
[**WalletGetAssetCustodyTransferHistoryV1**](AssetAPI.md#WalletGetAssetCustodyTransferHistoryV1) | **Get** /sapi/v1/asset/custody/transfer-history | Query User Delegation History(For Master Account)(USER_DATA)
[**WalletGetAssetDribbletV1**](AssetAPI.md#WalletGetAssetDribbletV1) | **Get** /sapi/v1/asset/dribblet | DustLog(USER_DATA)
[**WalletGetAssetLedgerTransferCloudMiningQueryByPageV1**](AssetAPI.md#WalletGetAssetLedgerTransferCloudMiningQueryByPageV1) | **Get** /sapi/v1/asset/ledger-transfer/cloud-mining/queryByPage | Get Cloud-Mining payment and refund history (USER_DATA)
[**WalletGetAssetTradeFeeV1**](AssetAPI.md#WalletGetAssetTradeFeeV1) | **Get** /sapi/v1/asset/tradeFee | Trade Fee (USER_DATA)
[**WalletGetAssetTransferV1**](AssetAPI.md#WalletGetAssetTransferV1) | **Get** /sapi/v1/asset/transfer | Query User Universal Transfer History(USER_DATA)
[**WalletGetAssetWalletBalanceV1**](AssetAPI.md#WalletGetAssetWalletBalanceV1) | **Get** /sapi/v1/asset/wallet/balance | Query User Wallet Balance (USER_DATA)
[**WalletGetSpotDelistScheduleV1**](AssetAPI.md#WalletGetSpotDelistScheduleV1) | **Get** /sapi/v1/spot/delist-schedule | Get Spot Delist Schedule (MARKET_DATA)
[**WalletGetSpotOpenSymbolListV1**](AssetAPI.md#WalletGetSpotOpenSymbolListV1) | **Get** /sapi/v1/spot/open-symbol-list | Get Open Symbol List (MARKET_DATA)



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
	resp, r, err := apiClient.AssetAPI.WalletCreateAssetDustBtcV1(context.Background()).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletCreateAssetDustBtcV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetDustBtcV1`: WalletCreateAssetDustBtcV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletCreateAssetDustBtcV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletCreateAssetDustV1(context.Background()).Asset(asset).Timestamp(timestamp).AccountType(accountType).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletCreateAssetDustV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetDustV1`: WalletCreateAssetDustV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletCreateAssetDustV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletCreateAssetGetFundingAssetV1(context.Background()).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletCreateAssetGetFundingAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetGetFundingAssetV1`: []WalletCreateAssetGetFundingAssetV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletCreateAssetGetFundingAssetV1`: %v\n", resp)
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


## WalletCreateAssetGetUserAssetV3

> []WalletCreateAssetGetUserAssetV3RespItem WalletCreateAssetGetUserAssetV3(ctx).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

User Asset (USER_DATA)



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
	needBtcValuation := true // bool |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetAPI.WalletCreateAssetGetUserAssetV3(context.Background()).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletCreateAssetGetUserAssetV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetGetUserAssetV3`: []WalletCreateAssetGetUserAssetV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletCreateAssetGetUserAssetV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAssetGetUserAssetV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **needBtcValuation** | **bool** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletCreateAssetGetUserAssetV3RespItem**](WalletCreateAssetGetUserAssetV3RespItem.md)

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
	resp, r, err := apiClient.AssetAPI.WalletCreateAssetTransferV1(context.Background()).Amount(amount).Asset(asset).Timestamp(timestamp).Type_(type_).FromSymbol(fromSymbol).RecvWindow(recvWindow).ToSymbol(toSymbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletCreateAssetTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetTransferV1`: WalletCreateAssetTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletCreateAssetTransferV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletCreateBnbBurnV1(context.Background()).Timestamp(timestamp).InterestBNBBurn(interestBNBBurn).RecvWindow(recvWindow).SpotBNBBurn(spotBNBBurn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletCreateBnbBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateBnbBurnV1`: WalletCreateBnbBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletCreateBnbBurnV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetAssetDetailV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetAssetDetailV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetAssetDetailV1`: map[string]WalletGetAssetAssetDetailV1RespValue
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetAssetDetailV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetAssetDividendV1(context.Background()).Timestamp(timestamp).Asset(asset).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetAssetDividendV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetAssetDividendV1`: WalletGetAssetAssetDividendV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetAssetDividendV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetCustodyTransferHistoryV1(context.Background()).Email(email).StartTime(startTime).EndTime(endTime).Timestamp(timestamp).Type_(type_).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetCustodyTransferHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetCustodyTransferHistoryV1`: WalletGetAssetCustodyTransferHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetCustodyTransferHistoryV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetDribbletV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetDribbletV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetDribbletV1`: WalletGetAssetDribbletV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetDribbletV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetLedgerTransferCloudMiningQueryByPageV1(context.Background()).StartTime(startTime).EndTime(endTime).TranId(tranId).ClientTranId(clientTranId).Asset(asset).Current(current).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetLedgerTransferCloudMiningQueryByPageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetLedgerTransferCloudMiningQueryByPageV1`: WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetLedgerTransferCloudMiningQueryByPageV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetTradeFeeV1(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetTradeFeeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetTradeFeeV1`: []WalletGetAssetTradeFeeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetTradeFeeV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetTransferV1(context.Background()).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).FromSymbol(fromSymbol).ToSymbol(toSymbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetTransferV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetTransferV1`: WalletGetAssetTransferV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetTransferV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetAssetWalletBalanceV1(context.Background()).Timestamp(timestamp).QuoteAsset(quoteAsset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetWalletBalanceV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetWalletBalanceV1`: []WalletGetAssetWalletBalanceV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetWalletBalanceV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetSpotDelistScheduleV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetSpotDelistScheduleV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetSpotDelistScheduleV1`: []WalletGetSpotDelistScheduleV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetSpotDelistScheduleV1`: %v\n", resp)
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
	resp, r, err := apiClient.AssetAPI.WalletGetSpotOpenSymbolListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetSpotOpenSymbolListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetSpotOpenSymbolListV1`: []WalletGetSpotOpenSymbolListV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetSpotOpenSymbolListV1`: %v\n", resp)
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

