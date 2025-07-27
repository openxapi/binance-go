# \StakingAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEthStakingEthRedeemV1**](StakingAPI.md#CreateEthStakingEthRedeemV1) | **Post** /sapi/v1/eth-staking/eth/redeem | Redeem ETH(TRADE)
[**CreateEthStakingEthStakeV2**](StakingAPI.md#CreateEthStakingEthStakeV2) | **Post** /sapi/v2/eth-staking/eth/stake | Subscribe ETH Staking(TRADE)
[**CreateEthStakingWbethWrapV1**](StakingAPI.md#CreateEthStakingWbethWrapV1) | **Post** /sapi/v1/eth-staking/wbeth/wrap | Wrap BETH(TRADE)
[**CreateSolStakingSolClaimV1**](StakingAPI.md#CreateSolStakingSolClaimV1) | **Post** /sapi/v1/sol-staking/sol/claim | Claim Boost Rewards(TRADE)
[**CreateSolStakingSolRedeemV1**](StakingAPI.md#CreateSolStakingSolRedeemV1) | **Post** /sapi/v1/sol-staking/sol/redeem | Redeem SOL(TRADE)
[**CreateSolStakingSolStakeV1**](StakingAPI.md#CreateSolStakingSolStakeV1) | **Post** /sapi/v1/sol-staking/sol/stake | Subscribe SOL Staking(TRADE)
[**GetEthStakingAccountV2**](StakingAPI.md#GetEthStakingAccountV2) | **Get** /sapi/v2/eth-staking/account | ETH Staking account(USER_DATA)
[**GetEthStakingEthHistoryRateHistoryV1**](StakingAPI.md#GetEthStakingEthHistoryRateHistoryV1) | **Get** /sapi/v1/eth-staking/eth/history/rateHistory | Get WBETH Rate History(USER_DATA)
[**GetEthStakingEthHistoryRedemptionHistoryV1**](StakingAPI.md#GetEthStakingEthHistoryRedemptionHistoryV1) | **Get** /sapi/v1/eth-staking/eth/history/redemptionHistory | Get ETH redemption history(USER_DATA)
[**GetEthStakingEthHistoryRewardsHistoryV1**](StakingAPI.md#GetEthStakingEthHistoryRewardsHistoryV1) | **Get** /sapi/v1/eth-staking/eth/history/rewardsHistory | Get BETH rewards distribution history(USER_DATA)
[**GetEthStakingEthHistoryStakingHistoryV1**](StakingAPI.md#GetEthStakingEthHistoryStakingHistoryV1) | **Get** /sapi/v1/eth-staking/eth/history/stakingHistory | Get ETH staking history(USER_DATA)
[**GetEthStakingEthHistoryWbethRewardsHistoryV1**](StakingAPI.md#GetEthStakingEthHistoryWbethRewardsHistoryV1) | **Get** /sapi/v1/eth-staking/eth/history/wbethRewardsHistory | Get WBETH rewards history(USER_DATA)
[**GetEthStakingEthQuotaV1**](StakingAPI.md#GetEthStakingEthQuotaV1) | **Get** /sapi/v1/eth-staking/eth/quota | Get current ETH staking quota(USER_DATA)
[**GetEthStakingWbethHistoryUnwrapHistoryV1**](StakingAPI.md#GetEthStakingWbethHistoryUnwrapHistoryV1) | **Get** /sapi/v1/eth-staking/wbeth/history/unwrapHistory | Get WBETH unwrap history(USER_DATA)
[**GetEthStakingWbethHistoryWrapHistoryV1**](StakingAPI.md#GetEthStakingWbethHistoryWrapHistoryV1) | **Get** /sapi/v1/eth-staking/wbeth/history/wrapHistory | Get WBETH wrap history(USER_DATA)
[**GetSolStakingAccountV1**](StakingAPI.md#GetSolStakingAccountV1) | **Get** /sapi/v1/sol-staking/account | SOL Staking account(USER_DATA)
[**GetSolStakingSolHistoryBnsolRewardsHistoryV1**](StakingAPI.md#GetSolStakingSolHistoryBnsolRewardsHistoryV1) | **Get** /sapi/v1/sol-staking/sol/history/bnsolRewardsHistory | Get BNSOL rewards history(USER_DATA)
[**GetSolStakingSolHistoryBoostRewardsHistoryV1**](StakingAPI.md#GetSolStakingSolHistoryBoostRewardsHistoryV1) | **Get** /sapi/v1/sol-staking/sol/history/boostRewardsHistory | Get Boost Rewards History(USER_DATA)
[**GetSolStakingSolHistoryRateHistoryV1**](StakingAPI.md#GetSolStakingSolHistoryRateHistoryV1) | **Get** /sapi/v1/sol-staking/sol/history/rateHistory | Get BNSOL Rate History(USER_DATA)
[**GetSolStakingSolHistoryRedemptionHistoryV1**](StakingAPI.md#GetSolStakingSolHistoryRedemptionHistoryV1) | **Get** /sapi/v1/sol-staking/sol/history/redemptionHistory | Get SOL redemption history(USER_DATA)
[**GetSolStakingSolHistoryStakingHistoryV1**](StakingAPI.md#GetSolStakingSolHistoryStakingHistoryV1) | **Get** /sapi/v1/sol-staking/sol/history/stakingHistory | Get SOL staking history(USER_DATA)
[**GetSolStakingSolHistoryUnclaimedRewardsV1**](StakingAPI.md#GetSolStakingSolHistoryUnclaimedRewardsV1) | **Get** /sapi/v1/sol-staking/sol/history/unclaimedRewards | Get Unclaimed Rewards(USER_DATA)
[**GetSolStakingSolQuotaV1**](StakingAPI.md#GetSolStakingSolQuotaV1) | **Get** /sapi/v1/sol-staking/sol/quota | Get SOL staking quota details(USER_DATA)



## CreateEthStakingEthRedeemV1

> CreateEthStakingEthRedeemV1Resp CreateEthStakingEthRedeemV1(ctx).Amount(amount).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()

Redeem ETH(TRADE)



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
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.CreateEthStakingEthRedeemV1(context.Background()).Amount(amount).Timestamp(timestamp).Asset(asset).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.CreateEthStakingEthRedeemV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEthStakingEthRedeemV1`: CreateEthStakingEthRedeemV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.CreateEthStakingEthRedeemV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEthStakingEthRedeemV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateEthStakingEthRedeemV1Resp**](CreateEthStakingEthRedeemV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEthStakingEthStakeV2

> CreateEthStakingEthStakeV2Resp CreateEthStakingEthStakeV2(ctx).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Subscribe ETH Staking(TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.CreateEthStakingEthStakeV2(context.Background()).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.CreateEthStakingEthStakeV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEthStakingEthStakeV2`: CreateEthStakingEthStakeV2Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.CreateEthStakingEthStakeV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEthStakingEthStakeV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateEthStakingEthStakeV2Resp**](CreateEthStakingEthStakeV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEthStakingWbethWrapV1

> CreateEthStakingWbethWrapV1Resp CreateEthStakingWbethWrapV1(ctx).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Wrap BETH(TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.CreateEthStakingWbethWrapV1(context.Background()).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.CreateEthStakingWbethWrapV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEthStakingWbethWrapV1`: CreateEthStakingWbethWrapV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.CreateEthStakingWbethWrapV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEthStakingWbethWrapV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateEthStakingWbethWrapV1Resp**](CreateEthStakingWbethWrapV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSolStakingSolClaimV1

> CreateSolStakingSolClaimV1Resp CreateSolStakingSolClaimV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Claim Boost Rewards(TRADE)



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
	resp, r, err := apiClient.StakingAPI.CreateSolStakingSolClaimV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.CreateSolStakingSolClaimV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSolStakingSolClaimV1`: CreateSolStakingSolClaimV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.CreateSolStakingSolClaimV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSolStakingSolClaimV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSolStakingSolClaimV1Resp**](CreateSolStakingSolClaimV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSolStakingSolRedeemV1

> CreateSolStakingSolRedeemV1Resp CreateSolStakingSolRedeemV1(ctx).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Redeem SOL(TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.CreateSolStakingSolRedeemV1(context.Background()).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.CreateSolStakingSolRedeemV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSolStakingSolRedeemV1`: CreateSolStakingSolRedeemV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.CreateSolStakingSolRedeemV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSolStakingSolRedeemV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSolStakingSolRedeemV1Resp**](CreateSolStakingSolRedeemV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSolStakingSolStakeV1

> CreateSolStakingSolStakeV1Resp CreateSolStakingSolStakeV1(ctx).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Subscribe SOL Staking(TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.CreateSolStakingSolStakeV1(context.Background()).Amount(amount).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.CreateSolStakingSolStakeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSolStakingSolStakeV1`: CreateSolStakingSolStakeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.CreateSolStakingSolStakeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSolStakingSolStakeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateSolStakingSolStakeV1Resp**](CreateSolStakingSolStakeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingAccountV2

> GetEthStakingAccountV2Resp GetEthStakingAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

ETH Staking account(USER_DATA)



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
	resp, r, err := apiClient.StakingAPI.GetEthStakingAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingAccountV2`: GetEthStakingAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingAccountV2Resp**](GetEthStakingAccountV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingEthHistoryRateHistoryV1

> GetEthStakingEthHistoryRateHistoryV1Resp GetEthStakingEthHistoryRateHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH Rate History(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetEthStakingEthHistoryRateHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingEthHistoryRateHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingEthHistoryRateHistoryV1`: GetEthStakingEthHistoryRateHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingEthHistoryRateHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingEthHistoryRateHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingEthHistoryRateHistoryV1Resp**](GetEthStakingEthHistoryRateHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingEthHistoryRedemptionHistoryV1

> GetEthStakingEthHistoryRedemptionHistoryV1Resp GetEthStakingEthHistoryRedemptionHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get ETH redemption history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetEthStakingEthHistoryRedemptionHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingEthHistoryRedemptionHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingEthHistoryRedemptionHistoryV1`: GetEthStakingEthHistoryRedemptionHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingEthHistoryRedemptionHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingEthHistoryRedemptionHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingEthHistoryRedemptionHistoryV1Resp**](GetEthStakingEthHistoryRedemptionHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingEthHistoryRewardsHistoryV1

> GetEthStakingEthHistoryRewardsHistoryV1Resp GetEthStakingEthHistoryRewardsHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BETH rewards distribution history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetEthStakingEthHistoryRewardsHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingEthHistoryRewardsHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingEthHistoryRewardsHistoryV1`: GetEthStakingEthHistoryRewardsHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingEthHistoryRewardsHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingEthHistoryRewardsHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingEthHistoryRewardsHistoryV1Resp**](GetEthStakingEthHistoryRewardsHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingEthHistoryStakingHistoryV1

> GetEthStakingEthHistoryStakingHistoryV1Resp GetEthStakingEthHistoryStakingHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get ETH staking history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetEthStakingEthHistoryStakingHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingEthHistoryStakingHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingEthHistoryStakingHistoryV1`: GetEthStakingEthHistoryStakingHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingEthHistoryStakingHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingEthHistoryStakingHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingEthHistoryStakingHistoryV1Resp**](GetEthStakingEthHistoryStakingHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingEthHistoryWbethRewardsHistoryV1

> GetEthStakingEthHistoryWbethRewardsHistoryV1Resp GetEthStakingEthHistoryWbethRewardsHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH rewards history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetEthStakingEthHistoryWbethRewardsHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingEthHistoryWbethRewardsHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingEthHistoryWbethRewardsHistoryV1`: GetEthStakingEthHistoryWbethRewardsHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingEthHistoryWbethRewardsHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingEthHistoryWbethRewardsHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingEthHistoryWbethRewardsHistoryV1Resp**](GetEthStakingEthHistoryWbethRewardsHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingEthQuotaV1

> GetEthStakingEthQuotaV1Resp GetEthStakingEthQuotaV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get current ETH staking quota(USER_DATA)



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
	resp, r, err := apiClient.StakingAPI.GetEthStakingEthQuotaV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingEthQuotaV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingEthQuotaV1`: GetEthStakingEthQuotaV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingEthQuotaV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingEthQuotaV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingEthQuotaV1Resp**](GetEthStakingEthQuotaV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingWbethHistoryUnwrapHistoryV1

> GetEthStakingWbethHistoryUnwrapHistoryV1Resp GetEthStakingWbethHistoryUnwrapHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH unwrap history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetEthStakingWbethHistoryUnwrapHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingWbethHistoryUnwrapHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingWbethHistoryUnwrapHistoryV1`: GetEthStakingWbethHistoryUnwrapHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingWbethHistoryUnwrapHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingWbethHistoryUnwrapHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingWbethHistoryUnwrapHistoryV1Resp**](GetEthStakingWbethHistoryUnwrapHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthStakingWbethHistoryWrapHistoryV1

> GetEthStakingWbethHistoryWrapHistoryV1Resp GetEthStakingWbethHistoryWrapHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get WBETH wrap history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetEthStakingWbethHistoryWrapHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetEthStakingWbethHistoryWrapHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthStakingWbethHistoryWrapHistoryV1`: GetEthStakingWbethHistoryWrapHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetEthStakingWbethHistoryWrapHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEthStakingWbethHistoryWrapHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetEthStakingWbethHistoryWrapHistoryV1Resp**](GetEthStakingWbethHistoryWrapHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingAccountV1

> GetSolStakingAccountV1Resp GetSolStakingAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

SOL Staking account(USER_DATA)



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
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetSolStakingAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingAccountV1`: GetSolStakingAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetSolStakingAccountV1Resp**](GetSolStakingAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingSolHistoryBnsolRewardsHistoryV1

> GetSolStakingSolHistoryBnsolRewardsHistoryV1Resp GetSolStakingSolHistoryBnsolRewardsHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BNSOL rewards history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetSolStakingSolHistoryBnsolRewardsHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingSolHistoryBnsolRewardsHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingSolHistoryBnsolRewardsHistoryV1`: GetSolStakingSolHistoryBnsolRewardsHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingSolHistoryBnsolRewardsHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingSolHistoryBnsolRewardsHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetSolStakingSolHistoryBnsolRewardsHistoryV1Resp**](GetSolStakingSolHistoryBnsolRewardsHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingSolHistoryBoostRewardsHistoryV1

> GetSolStakingSolHistoryBoostRewardsHistoryV1Resp GetSolStakingSolHistoryBoostRewardsHistoryV1(ctx).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get Boost Rewards History(USER_DATA)



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
	type_ := "type__example" // string | &#34;CLAIM&#34;, &#34;DISTRIBUTE&#34;, default &#34;CLAIM&#34; (default to "")
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetSolStakingSolHistoryBoostRewardsHistoryV1(context.Background()).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingSolHistoryBoostRewardsHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingSolHistoryBoostRewardsHistoryV1`: GetSolStakingSolHistoryBoostRewardsHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingSolHistoryBoostRewardsHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingSolHistoryBoostRewardsHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &amp;#34;CLAIM&amp;#34;, &amp;#34;DISTRIBUTE&amp;#34;, default &amp;#34;CLAIM&amp;#34; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetSolStakingSolHistoryBoostRewardsHistoryV1Resp**](GetSolStakingSolHistoryBoostRewardsHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingSolHistoryRateHistoryV1

> GetSolStakingSolHistoryRateHistoryV1Resp GetSolStakingSolHistoryRateHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get BNSOL Rate History(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default:1 (optional)
	size := int64(789) // int64 | Default:10, Max:100 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetSolStakingSolHistoryRateHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingSolHistoryRateHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingSolHistoryRateHistoryV1`: GetSolStakingSolHistoryRateHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingSolHistoryRateHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingSolHistoryRateHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default:1 | 
 **size** | **int64** | Default:10, Max:100 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetSolStakingSolHistoryRateHistoryV1Resp**](GetSolStakingSolHistoryRateHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingSolHistoryRedemptionHistoryV1

> GetSolStakingSolHistoryRedemptionHistoryV1Resp GetSolStakingSolHistoryRedemptionHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get SOL redemption history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetSolStakingSolHistoryRedemptionHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingSolHistoryRedemptionHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingSolHistoryRedemptionHistoryV1`: GetSolStakingSolHistoryRedemptionHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingSolHistoryRedemptionHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingSolHistoryRedemptionHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetSolStakingSolHistoryRedemptionHistoryV1Resp**](GetSolStakingSolHistoryRedemptionHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingSolHistoryStakingHistoryV1

> GetSolStakingSolHistoryStakingHistoryV1Resp GetSolStakingSolHistoryStakingHistoryV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()

Get SOL staking history(USER_DATA)



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
	current := int64(789) // int64 | Currently querying page. Start from 1. Default: 1 (optional) (default to 1)
	size := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetSolStakingSolHistoryStakingHistoryV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Size(size).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingSolHistoryStakingHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingSolHistoryStakingHistoryV1`: GetSolStakingSolHistoryStakingHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingSolHistoryStakingHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingSolHistoryStakingHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1. Default: 1 | [default to 1]
 **size** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetSolStakingSolHistoryStakingHistoryV1Resp**](GetSolStakingSolHistoryStakingHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingSolHistoryUnclaimedRewardsV1

> []GetSolStakingSolHistoryUnclaimedRewardsV1RespItem GetSolStakingSolHistoryUnclaimedRewardsV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Unclaimed Rewards(USER_DATA)



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
	resp, r, err := apiClient.StakingAPI.GetSolStakingSolHistoryUnclaimedRewardsV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingSolHistoryUnclaimedRewardsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingSolHistoryUnclaimedRewardsV1`: []GetSolStakingSolHistoryUnclaimedRewardsV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingSolHistoryUnclaimedRewardsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingSolHistoryUnclaimedRewardsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetSolStakingSolHistoryUnclaimedRewardsV1RespItem**](GetSolStakingSolHistoryUnclaimedRewardsV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSolStakingSolQuotaV1

> GetSolStakingSolQuotaV1Resp GetSolStakingSolQuotaV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get SOL staking quota details(USER_DATA)



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
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetSolStakingSolQuotaV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetSolStakingSolQuotaV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSolStakingSolQuotaV1`: GetSolStakingSolQuotaV1Resp
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetSolStakingSolQuotaV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSolStakingSolQuotaV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetSolStakingSolQuotaV1Resp**](GetSolStakingSolQuotaV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

