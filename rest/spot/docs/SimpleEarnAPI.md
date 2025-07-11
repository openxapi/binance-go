# \SimpleEarnAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSimpleEarnFlexibleRedeemV1**](SimpleEarnAPI.md#CreateSimpleEarnFlexibleRedeemV1) | **Post** /sapi/v1/simple-earn/flexible/redeem | Redeem Flexible Product(TRADE)
[**CreateSimpleEarnFlexibleSetAutoSubscribeV1**](SimpleEarnAPI.md#CreateSimpleEarnFlexibleSetAutoSubscribeV1) | **Post** /sapi/v1/simple-earn/flexible/setAutoSubscribe | Set Flexible Auto Subscribe(USER_DATA)
[**CreateSimpleEarnFlexibleSubscribeV1**](SimpleEarnAPI.md#CreateSimpleEarnFlexibleSubscribeV1) | **Post** /sapi/v1/simple-earn/flexible/subscribe | Subscribe Flexible Product(TRADE)
[**CreateSimpleEarnLockedRedeemV1**](SimpleEarnAPI.md#CreateSimpleEarnLockedRedeemV1) | **Post** /sapi/v1/simple-earn/locked/redeem | Redeem Locked Product(TRADE)
[**CreateSimpleEarnLockedSetAutoSubscribeV1**](SimpleEarnAPI.md#CreateSimpleEarnLockedSetAutoSubscribeV1) | **Post** /sapi/v1/simple-earn/locked/setAutoSubscribe | Set Locked Auto Subscribe(USER_DATA)
[**CreateSimpleEarnLockedSetRedeemOptionV1**](SimpleEarnAPI.md#CreateSimpleEarnLockedSetRedeemOptionV1) | **Post** /sapi/v1/simple-earn/locked/setRedeemOption | Set Locked Product Redeem Option(USER_DATA)
[**CreateSimpleEarnLockedSubscribeV1**](SimpleEarnAPI.md#CreateSimpleEarnLockedSubscribeV1) | **Post** /sapi/v1/simple-earn/locked/subscribe | Subscribe Locked Product(TRADE)
[**GetSimpleEarnAccountV1**](SimpleEarnAPI.md#GetSimpleEarnAccountV1) | **Get** /sapi/v1/simple-earn/account | Simple Account(USER_DATA)
[**GetSimpleEarnFlexibleHistoryCollateralRecordV1**](SimpleEarnAPI.md#GetSimpleEarnFlexibleHistoryCollateralRecordV1) | **Get** /sapi/v1/simple-earn/flexible/history/collateralRecord | Get Collateral Record(USER_DATA)
[**GetSimpleEarnFlexibleHistoryRateHistoryV1**](SimpleEarnAPI.md#GetSimpleEarnFlexibleHistoryRateHistoryV1) | **Get** /sapi/v1/simple-earn/flexible/history/rateHistory | Get Rate History(USER_DATA)
[**GetSimpleEarnFlexibleHistoryRedemptionRecordV1**](SimpleEarnAPI.md#GetSimpleEarnFlexibleHistoryRedemptionRecordV1) | **Get** /sapi/v1/simple-earn/flexible/history/redemptionRecord | Get Flexible Redemption Record(USER_DATA)
[**GetSimpleEarnFlexibleHistoryRewardsRecordV1**](SimpleEarnAPI.md#GetSimpleEarnFlexibleHistoryRewardsRecordV1) | **Get** /sapi/v1/simple-earn/flexible/history/rewardsRecord | Get Flexible Rewards History(USER_DATA)
[**GetSimpleEarnFlexibleHistorySubscriptionRecordV1**](SimpleEarnAPI.md#GetSimpleEarnFlexibleHistorySubscriptionRecordV1) | **Get** /sapi/v1/simple-earn/flexible/history/subscriptionRecord | Get Flexible Subscription Record(USER_DATA)
[**GetSimpleEarnFlexibleListV1**](SimpleEarnAPI.md#GetSimpleEarnFlexibleListV1) | **Get** /sapi/v1/simple-earn/flexible/list | Get Simple Earn Flexible Product List(USER_DATA)
[**GetSimpleEarnFlexiblePersonalLeftQuotaV1**](SimpleEarnAPI.md#GetSimpleEarnFlexiblePersonalLeftQuotaV1) | **Get** /sapi/v1/simple-earn/flexible/personalLeftQuota | Get Flexible Personal Left Quota(USER_DATA)
[**GetSimpleEarnFlexiblePositionV1**](SimpleEarnAPI.md#GetSimpleEarnFlexiblePositionV1) | **Get** /sapi/v1/simple-earn/flexible/position | Get Flexible Product Position(USER_DATA)
[**GetSimpleEarnFlexibleSubscriptionPreviewV1**](SimpleEarnAPI.md#GetSimpleEarnFlexibleSubscriptionPreviewV1) | **Get** /sapi/v1/simple-earn/flexible/subscriptionPreview | Get Flexible Subscription Preview(USER_DATA)
[**GetSimpleEarnLockedHistoryRedemptionRecordV1**](SimpleEarnAPI.md#GetSimpleEarnLockedHistoryRedemptionRecordV1) | **Get** /sapi/v1/simple-earn/locked/history/redemptionRecord | Get Locked Redemption Record(USER_DATA)
[**GetSimpleEarnLockedHistoryRewardsRecordV1**](SimpleEarnAPI.md#GetSimpleEarnLockedHistoryRewardsRecordV1) | **Get** /sapi/v1/simple-earn/locked/history/rewardsRecord | Get Locked Rewards History(USER_DATA)
[**GetSimpleEarnLockedHistorySubscriptionRecordV1**](SimpleEarnAPI.md#GetSimpleEarnLockedHistorySubscriptionRecordV1) | **Get** /sapi/v1/simple-earn/locked/history/subscriptionRecord | Get Locked Subscription Record(USER_DATA)
[**GetSimpleEarnLockedListV1**](SimpleEarnAPI.md#GetSimpleEarnLockedListV1) | **Get** /sapi/v1/simple-earn/locked/list | Get Simple Earn Locked Product List(USER_DATA)
[**GetSimpleEarnLockedPersonalLeftQuotaV1**](SimpleEarnAPI.md#GetSimpleEarnLockedPersonalLeftQuotaV1) | **Get** /sapi/v1/simple-earn/locked/personalLeftQuota | Get Locked Personal Left Quota(USER_DATA)
[**GetSimpleEarnLockedPositionV1**](SimpleEarnAPI.md#GetSimpleEarnLockedPositionV1) | **Get** /sapi/v1/simple-earn/locked/position | Get Locked Product Position
[**GetSimpleEarnLockedSubscriptionPreviewV1**](SimpleEarnAPI.md#GetSimpleEarnLockedSubscriptionPreviewV1) | **Get** /sapi/v1/simple-earn/locked/subscriptionPreview | Get Locked Subscription Preview(USER_DATA)



## CreateSimpleEarnFlexibleRedeemV1

> CreateSimpleEarnFlexibleRedeemV1Resp CreateSimpleEarnFlexibleRedeemV1(ctx).ProductId(productId).Timestamp(timestamp).Amount(amount).DestAccount(destAccount).RecvWindow(recvWindow).RedeemAll(redeemAll).Execute()

Redeem Flexible Product(TRADE)



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
	productId := "productId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	amount := "amount_example" // string |  (optional) (default to "")
	destAccount := "destAccount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	redeemAll := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.CreateSimpleEarnFlexibleRedeemV1(context.Background()).ProductId(productId).Timestamp(timestamp).Amount(amount).DestAccount(destAccount).RecvWindow(recvWindow).RedeemAll(redeemAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.CreateSimpleEarnFlexibleRedeemV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimpleEarnFlexibleRedeemV1`: CreateSimpleEarnFlexibleRedeemV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.CreateSimpleEarnFlexibleRedeemV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimpleEarnFlexibleRedeemV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **amount** | **string** |  | [default to &quot;&quot;]
 **destAccount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **redeemAll** | **bool** |  | 

### Return type

[**CreateSimpleEarnFlexibleRedeemV1Resp**](CreateSimpleEarnFlexibleRedeemV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSimpleEarnFlexibleSetAutoSubscribeV1

> CreateSimpleEarnFlexibleSetAutoSubscribeV1Resp CreateSimpleEarnFlexibleSetAutoSubscribeV1(ctx).AutoSubscribe(autoSubscribe).ProductId(productId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Set Flexible Auto Subscribe(USER_DATA)



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
	autoSubscribe := true // bool | 
	productId := "productId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.CreateSimpleEarnFlexibleSetAutoSubscribeV1(context.Background()).AutoSubscribe(autoSubscribe).ProductId(productId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.CreateSimpleEarnFlexibleSetAutoSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimpleEarnFlexibleSetAutoSubscribeV1`: CreateSimpleEarnFlexibleSetAutoSubscribeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.CreateSimpleEarnFlexibleSetAutoSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimpleEarnFlexibleSetAutoSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoSubscribe** | **bool** |  | 
 **productId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSimpleEarnFlexibleSetAutoSubscribeV1Resp**](CreateSimpleEarnFlexibleSetAutoSubscribeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSimpleEarnFlexibleSubscribeV1

> CreateSimpleEarnFlexibleSubscribeV1Resp CreateSimpleEarnFlexibleSubscribeV1(ctx).Amount(amount).ProductId(productId).Timestamp(timestamp).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).SourceAccount(sourceAccount).Execute()

Subscribe Flexible Product(TRADE)



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
	productId := "productId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	autoSubscribe := true // bool |  (optional) (default to true)
	recvWindow := int64(789) // int64 |  (optional)
	sourceAccount := "sourceAccount_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.CreateSimpleEarnFlexibleSubscribeV1(context.Background()).Amount(amount).ProductId(productId).Timestamp(timestamp).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).SourceAccount(sourceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.CreateSimpleEarnFlexibleSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimpleEarnFlexibleSubscribeV1`: CreateSimpleEarnFlexibleSubscribeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.CreateSimpleEarnFlexibleSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimpleEarnFlexibleSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **productId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **autoSubscribe** | **bool** |  | [default to true]
 **recvWindow** | **int64** |  | 
 **sourceAccount** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateSimpleEarnFlexibleSubscribeV1Resp**](CreateSimpleEarnFlexibleSubscribeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSimpleEarnLockedRedeemV1

> CreateSimpleEarnLockedRedeemV1Resp CreateSimpleEarnLockedRedeemV1(ctx).PositionId(positionId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Redeem Locked Product(TRADE)



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
	positionId := int32(56) // int32 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.CreateSimpleEarnLockedRedeemV1(context.Background()).PositionId(positionId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.CreateSimpleEarnLockedRedeemV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimpleEarnLockedRedeemV1`: CreateSimpleEarnLockedRedeemV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.CreateSimpleEarnLockedRedeemV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimpleEarnLockedRedeemV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **int32** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSimpleEarnLockedRedeemV1Resp**](CreateSimpleEarnLockedRedeemV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSimpleEarnLockedSetAutoSubscribeV1

> CreateSimpleEarnLockedSetAutoSubscribeV1Resp CreateSimpleEarnLockedSetAutoSubscribeV1(ctx).AutoSubscribe(autoSubscribe).PositionId(positionId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Set Locked Auto Subscribe(USER_DATA)



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
	autoSubscribe := true // bool | 
	positionId := int32(56) // int32 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.CreateSimpleEarnLockedSetAutoSubscribeV1(context.Background()).AutoSubscribe(autoSubscribe).PositionId(positionId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.CreateSimpleEarnLockedSetAutoSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimpleEarnLockedSetAutoSubscribeV1`: CreateSimpleEarnLockedSetAutoSubscribeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.CreateSimpleEarnLockedSetAutoSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimpleEarnLockedSetAutoSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoSubscribe** | **bool** |  | 
 **positionId** | **int32** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSimpleEarnLockedSetAutoSubscribeV1Resp**](CreateSimpleEarnLockedSetAutoSubscribeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSimpleEarnLockedSetRedeemOptionV1

> CreateSimpleEarnLockedSetRedeemOptionV1Resp CreateSimpleEarnLockedSetRedeemOptionV1(ctx).PositionId(positionId).RedeemTo(redeemTo).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Set Locked Product Redeem Option(USER_DATA)



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
	positionId := "positionId_example" // string |  (default to "")
	redeemTo := "redeemTo_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.CreateSimpleEarnLockedSetRedeemOptionV1(context.Background()).PositionId(positionId).RedeemTo(redeemTo).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.CreateSimpleEarnLockedSetRedeemOptionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimpleEarnLockedSetRedeemOptionV1`: CreateSimpleEarnLockedSetRedeemOptionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.CreateSimpleEarnLockedSetRedeemOptionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimpleEarnLockedSetRedeemOptionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** |  | [default to &quot;&quot;]
 **redeemTo** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSimpleEarnLockedSetRedeemOptionV1Resp**](CreateSimpleEarnLockedSetRedeemOptionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSimpleEarnLockedSubscribeV1

> CreateSimpleEarnLockedSubscribeV1Resp CreateSimpleEarnLockedSubscribeV1(ctx).Amount(amount).ProjectId(projectId).Timestamp(timestamp).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).RedeemTo(redeemTo).SourceAccount(sourceAccount).Execute()

Subscribe Locked Product(TRADE)



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
	projectId := "projectId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	autoSubscribe := true // bool |  (optional) (default to false)
	recvWindow := int64(789) // int64 |  (optional)
	redeemTo := "redeemTo_example" // string |  (optional) (default to "")
	sourceAccount := "sourceAccount_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.CreateSimpleEarnLockedSubscribeV1(context.Background()).Amount(amount).ProjectId(projectId).Timestamp(timestamp).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).RedeemTo(redeemTo).SourceAccount(sourceAccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.CreateSimpleEarnLockedSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSimpleEarnLockedSubscribeV1`: CreateSimpleEarnLockedSubscribeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.CreateSimpleEarnLockedSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimpleEarnLockedSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **projectId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **autoSubscribe** | **bool** |  | [default to false]
 **recvWindow** | **int64** |  | 
 **redeemTo** | **string** |  | [default to &quot;&quot;]
 **sourceAccount** | **string** |  | [default to &quot;&quot;]

### Return type

[**CreateSimpleEarnLockedSubscribeV1Resp**](CreateSimpleEarnLockedSubscribeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnAccountV1

> GetSimpleEarnAccountV1Resp GetSimpleEarnAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Simple Account(USER_DATA)



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
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnAccountV1`: GetSimpleEarnAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**GetSimpleEarnAccountV1Resp**](GetSimpleEarnAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexibleHistoryCollateralRecordV1

> GetSimpleEarnFlexibleHistoryCollateralRecordV1Resp GetSimpleEarnFlexibleHistoryCollateralRecordV1(ctx).Timestamp(timestamp).ProductId(productId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Collateral Record(USER_DATA)



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
	productId := "productId_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexibleHistoryCollateralRecordV1(context.Background()).Timestamp(timestamp).ProductId(productId).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryCollateralRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexibleHistoryCollateralRecordV1`: GetSimpleEarnFlexibleHistoryCollateralRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryCollateralRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexibleHistoryCollateralRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **productId** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetSimpleEarnFlexibleHistoryCollateralRecordV1Resp**](GetSimpleEarnFlexibleHistoryCollateralRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexibleHistoryRateHistoryV1

> GetSimpleEarnFlexibleHistoryRateHistoryV1Resp GetSimpleEarnFlexibleHistoryRateHistoryV1(ctx).ProductId(productId).Timestamp(timestamp).AprPeriod(aprPeriod).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Rate History(USER_DATA)



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
	productId := "productId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	aprPeriod := "aprPeriod_example" // string | &#34;DAY&#34;,&#34;YEAR&#34;,default&#34;DAY&#34; (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRateHistoryV1(context.Background()).ProductId(productId).Timestamp(timestamp).AprPeriod(aprPeriod).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRateHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexibleHistoryRateHistoryV1`: GetSimpleEarnFlexibleHistoryRateHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRateHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexibleHistoryRateHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **aprPeriod** | **string** | &amp;#34;DAY&amp;#34;,&amp;#34;YEAR&amp;#34;,default&amp;#34;DAY&amp;#34; | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetSimpleEarnFlexibleHistoryRateHistoryV1Resp**](GetSimpleEarnFlexibleHistoryRateHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexibleHistoryRedemptionRecordV1

> GetSimpleEarnFlexibleHistoryRedemptionRecordV1Resp GetSimpleEarnFlexibleHistoryRedemptionRecordV1(ctx).ProductId(productId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Execute()

Get Flexible Redemption Record(USER_DATA)



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
	productId := "productId_example" // string |  (optional) (default to "")
	redeemId := "redeemId_example" // string |  (optional) (default to "")
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying the page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRedemptionRecordV1(context.Background()).ProductId(productId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRedemptionRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexibleHistoryRedemptionRecordV1`: GetSimpleEarnFlexibleHistoryRedemptionRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRedemptionRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexibleHistoryRedemptionRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | [default to &quot;&quot;]
 **redeemId** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying the page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 

### Return type

[**GetSimpleEarnFlexibleHistoryRedemptionRecordV1Resp**](GetSimpleEarnFlexibleHistoryRedemptionRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexibleHistoryRewardsRecordV1

> GetSimpleEarnFlexibleHistoryRewardsRecordV1Resp GetSimpleEarnFlexibleHistoryRewardsRecordV1(ctx).Type_(type_).Timestamp(timestamp).ProductId(productId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Rewards History(USER_DATA)



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
	type_ := "type__example" // string | `Bonus` - Bonus tiered APR, `REALTIME` Real-time APR, `REWARDS` Historical rewards,`ALL`(set to default) (default to "")
	timestamp := int64(789) // int64 | 
	productId := "productId_example" // string |  (optional) (default to "")
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying the page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRewardsRecordV1(context.Background()).Type_(type_).Timestamp(timestamp).ProductId(productId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRewardsRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexibleHistoryRewardsRecordV1`: GetSimpleEarnFlexibleHistoryRewardsRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexibleHistoryRewardsRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexibleHistoryRewardsRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &#x60;Bonus&#x60; - Bonus tiered APR, &#x60;REALTIME&#x60; Real-time APR, &#x60;REWARDS&#x60; Historical rewards,&#x60;ALL&#x60;(set to default) | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **productId** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying the page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnFlexibleHistoryRewardsRecordV1Resp**](GetSimpleEarnFlexibleHistoryRewardsRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexibleHistorySubscriptionRecordV1

> GetSimpleEarnFlexibleHistorySubscriptionRecordV1Resp GetSimpleEarnFlexibleHistorySubscriptionRecordV1(ctx).Timestamp(timestamp).ProductId(productId).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Subscription Record(USER_DATA)



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
	productId := "productId_example" // string |  (optional) (default to "")
	purchaseId := "purchaseId_example" // string |  (optional) (default to "")
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying the page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexibleHistorySubscriptionRecordV1(context.Background()).Timestamp(timestamp).ProductId(productId).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexibleHistorySubscriptionRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexibleHistorySubscriptionRecordV1`: GetSimpleEarnFlexibleHistorySubscriptionRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexibleHistorySubscriptionRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexibleHistorySubscriptionRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **productId** | **string** |  | [default to &quot;&quot;]
 **purchaseId** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying the page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnFlexibleHistorySubscriptionRecordV1Resp**](GetSimpleEarnFlexibleHistorySubscriptionRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexibleListV1

> GetSimpleEarnFlexibleListV1Resp GetSimpleEarnFlexibleListV1(ctx).Timestamp(timestamp).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Simple Earn Flexible Product List(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexibleListV1(context.Background()).Timestamp(timestamp).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexibleListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexibleListV1`: GetSimpleEarnFlexibleListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexibleListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexibleListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnFlexibleListV1Resp**](GetSimpleEarnFlexibleListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexiblePersonalLeftQuotaV1

> GetSimpleEarnFlexiblePersonalLeftQuotaV1Resp GetSimpleEarnFlexiblePersonalLeftQuotaV1(ctx).ProductId(productId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Flexible Personal Left Quota(USER_DATA)



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
	productId := "productId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexiblePersonalLeftQuotaV1(context.Background()).ProductId(productId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexiblePersonalLeftQuotaV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexiblePersonalLeftQuotaV1`: GetSimpleEarnFlexiblePersonalLeftQuotaV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexiblePersonalLeftQuotaV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexiblePersonalLeftQuotaV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnFlexiblePersonalLeftQuotaV1Resp**](GetSimpleEarnFlexiblePersonalLeftQuotaV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexiblePositionV1

> GetSimpleEarnFlexiblePositionV1Resp GetSimpleEarnFlexiblePositionV1(ctx).Timestamp(timestamp).Asset(asset).ProductId(productId).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Flexible Product Position(USER_DATA)



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
	productId := "productId_example" // string |  (optional) (default to "")
	current := int64(789) // int64 | Currently querying the page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexiblePositionV1(context.Background()).Timestamp(timestamp).Asset(asset).ProductId(productId).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexiblePositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexiblePositionV1`: GetSimpleEarnFlexiblePositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexiblePositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexiblePositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **productId** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Currently querying the page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnFlexiblePositionV1Resp**](GetSimpleEarnFlexiblePositionV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnFlexibleSubscriptionPreviewV1

> GetSimpleEarnFlexibleSubscriptionPreviewV1Resp GetSimpleEarnFlexibleSubscriptionPreviewV1(ctx).ProductId(productId).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Flexible Subscription Preview(USER_DATA)



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
	productId := "productId_example" // string |  (default to "")
	amount := "amount_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnFlexibleSubscriptionPreviewV1(context.Background()).ProductId(productId).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnFlexibleSubscriptionPreviewV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnFlexibleSubscriptionPreviewV1`: GetSimpleEarnFlexibleSubscriptionPreviewV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnFlexibleSubscriptionPreviewV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnFlexibleSubscriptionPreviewV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnFlexibleSubscriptionPreviewV1Resp**](GetSimpleEarnFlexibleSubscriptionPreviewV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnLockedHistoryRedemptionRecordV1

> GetSimpleEarnLockedHistoryRedemptionRecordV1Resp GetSimpleEarnLockedHistoryRedemptionRecordV1(ctx).Timestamp(timestamp).PositionId(positionId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Redemption Record(USER_DATA)



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
	positionId := int32(56) // int32 |  (optional)
	redeemId := "redeemId_example" // string |  (optional) (default to "")
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying the page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnLockedHistoryRedemptionRecordV1(context.Background()).Timestamp(timestamp).PositionId(positionId).RedeemId(redeemId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnLockedHistoryRedemptionRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnLockedHistoryRedemptionRecordV1`: GetSimpleEarnLockedHistoryRedemptionRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnLockedHistoryRedemptionRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnLockedHistoryRedemptionRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **positionId** | **int32** |  | 
 **redeemId** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying the page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnLockedHistoryRedemptionRecordV1Resp**](GetSimpleEarnLockedHistoryRedemptionRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnLockedHistoryRewardsRecordV1

> GetSimpleEarnLockedHistoryRewardsRecordV1Resp GetSimpleEarnLockedHistoryRewardsRecordV1(ctx).Timestamp(timestamp).PositionId(positionId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Rewards History(USER_DATA)



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
	positionId := int32(56) // int32 |  (optional)
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying the page. Start from 1, Default:1, Max: 1,000 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnLockedHistoryRewardsRecordV1(context.Background()).Timestamp(timestamp).PositionId(positionId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnLockedHistoryRewardsRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnLockedHistoryRewardsRecordV1`: GetSimpleEarnLockedHistoryRewardsRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnLockedHistoryRewardsRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnLockedHistoryRewardsRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **positionId** | **int32** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying the page. Start from 1, Default:1, Max: 1,000 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnLockedHistoryRewardsRecordV1Resp**](GetSimpleEarnLockedHistoryRewardsRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnLockedHistorySubscriptionRecordV1

> GetSimpleEarnLockedHistorySubscriptionRecordV1Resp GetSimpleEarnLockedHistorySubscriptionRecordV1(ctx).Timestamp(timestamp).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Subscription Record(USER_DATA)



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
	purchaseId := "purchaseId_example" // string |  (optional) (default to "")
	asset := "asset_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying the page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnLockedHistorySubscriptionRecordV1(context.Background()).Timestamp(timestamp).PurchaseId(purchaseId).Asset(asset).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnLockedHistorySubscriptionRecordV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnLockedHistorySubscriptionRecordV1`: GetSimpleEarnLockedHistorySubscriptionRecordV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnLockedHistorySubscriptionRecordV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnLockedHistorySubscriptionRecordV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **purchaseId** | **string** |  | [default to &quot;&quot;]
 **asset** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying the page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnLockedHistorySubscriptionRecordV1Resp**](GetSimpleEarnLockedHistorySubscriptionRecordV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnLockedListV1

> GetSimpleEarnLockedListV1Resp GetSimpleEarnLockedListV1(ctx).Timestamp(timestamp).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Simple Earn Locked Product List(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnLockedListV1(context.Background()).Timestamp(timestamp).Asset(asset).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnLockedListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnLockedListV1`: GetSimpleEarnLockedListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnLockedListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnLockedListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnLockedListV1Resp**](GetSimpleEarnLockedListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnLockedPersonalLeftQuotaV1

> GetSimpleEarnLockedPersonalLeftQuotaV1Resp GetSimpleEarnLockedPersonalLeftQuotaV1(ctx).ProjectId(projectId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Locked Personal Left Quota(USER_DATA)



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
	projectId := "projectId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnLockedPersonalLeftQuotaV1(context.Background()).ProjectId(projectId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnLockedPersonalLeftQuotaV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnLockedPersonalLeftQuotaV1`: GetSimpleEarnLockedPersonalLeftQuotaV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnLockedPersonalLeftQuotaV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnLockedPersonalLeftQuotaV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnLockedPersonalLeftQuotaV1Resp**](GetSimpleEarnLockedPersonalLeftQuotaV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnLockedPositionV1

> GetSimpleEarnLockedPositionV1Resp GetSimpleEarnLockedPositionV1(ctx).Timestamp(timestamp).Asset(asset).PositionId(positionId).ProjectId(projectId).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Locked Product Position



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
	positionId := int32(56) // int32 |  (optional)
	projectId := "projectId_example" // string |  (optional) (default to "")
	current := int64(789) // int64 | Currently querying the page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnLockedPositionV1(context.Background()).Timestamp(timestamp).Asset(asset).PositionId(positionId).ProjectId(projectId).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnLockedPositionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnLockedPositionV1`: GetSimpleEarnLockedPositionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnLockedPositionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnLockedPositionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **positionId** | **int32** |  | 
 **projectId** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Currently querying the page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetSimpleEarnLockedPositionV1Resp**](GetSimpleEarnLockedPositionV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleEarnLockedSubscriptionPreviewV1

> []GetSimpleEarnLockedSubscriptionPreviewV1RespItem GetSimpleEarnLockedSubscriptionPreviewV1(ctx).ProjectId(projectId).Amount(amount).Timestamp(timestamp).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()

Get Locked Subscription Preview(USER_DATA)



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
	projectId := "projectId_example" // string |  (default to "")
	amount := "amount_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	autoSubscribe := true // bool | true or false, default true. (optional) (default to true)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnAPI.GetSimpleEarnLockedSubscriptionPreviewV1(context.Background()).ProjectId(projectId).Amount(amount).Timestamp(timestamp).AutoSubscribe(autoSubscribe).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnAPI.GetSimpleEarnLockedSubscriptionPreviewV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleEarnLockedSubscriptionPreviewV1`: []GetSimpleEarnLockedSubscriptionPreviewV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnAPI.GetSimpleEarnLockedSubscriptionPreviewV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleEarnLockedSubscriptionPreviewV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **autoSubscribe** | **bool** | true or false, default true. | [default to true]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetSimpleEarnLockedSubscriptionPreviewV1RespItem**](GetSimpleEarnLockedSubscriptionPreviewV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

