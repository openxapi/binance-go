# \MiningAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMiningHashTransferConfigCancelV1**](MiningAPI.md#CreateMiningHashTransferConfigCancelV1) | **Post** /sapi/v1/mining/hash-transfer/config/cancel | Cancel hashrate resale configuration(USER_DATA)
[**CreateMiningHashTransferConfigV1**](MiningAPI.md#CreateMiningHashTransferConfigV1) | **Post** /sapi/v1/mining/hash-transfer/config | Hashrate Resale Request(USER_DATA)
[**GetMiningHashTransferConfigDetailsListV1**](MiningAPI.md#GetMiningHashTransferConfigDetailsListV1) | **Get** /sapi/v1/mining/hash-transfer/config/details/list | Hashrate Resale List (USER_DATA)
[**GetMiningHashTransferProfitDetailsV1**](MiningAPI.md#GetMiningHashTransferProfitDetailsV1) | **Get** /sapi/v1/mining/hash-transfer/profit/details | Hashrate Resale Detail(USER_DATA)
[**GetMiningPaymentListV1**](MiningAPI.md#GetMiningPaymentListV1) | **Get** /sapi/v1/mining/payment/list | Earnings List(USER_DATA)
[**GetMiningPaymentOtherV1**](MiningAPI.md#GetMiningPaymentOtherV1) | **Get** /sapi/v1/mining/payment/other | Extra Bonus List(USER_DATA)
[**GetMiningPaymentUidV1**](MiningAPI.md#GetMiningPaymentUidV1) | **Get** /sapi/v1/mining/payment/uid | Mining Account Earning(USER_DATA)
[**GetMiningPubAlgoListV1**](MiningAPI.md#GetMiningPubAlgoListV1) | **Get** /sapi/v1/mining/pub/algoList | Acquiring Algorithm(MARKET_DATA)
[**GetMiningPubCoinListV1**](MiningAPI.md#GetMiningPubCoinListV1) | **Get** /sapi/v1/mining/pub/coinList | Acquiring CoinName(MARKET_DATA)
[**GetMiningStatisticsUserListV1**](MiningAPI.md#GetMiningStatisticsUserListV1) | **Get** /sapi/v1/mining/statistics/user/list | Account List(USER_DATA)
[**GetMiningStatisticsUserStatusV1**](MiningAPI.md#GetMiningStatisticsUserStatusV1) | **Get** /sapi/v1/mining/statistics/user/status | Statistic List(USER_DATA)
[**GetMiningWorkerDetailV1**](MiningAPI.md#GetMiningWorkerDetailV1) | **Get** /sapi/v1/mining/worker/detail | Request for Detail Miner List(USER_DATA)
[**GetMiningWorkerListV1**](MiningAPI.md#GetMiningWorkerListV1) | **Get** /sapi/v1/mining/worker/list | Request for Miner List(USER_DATA)



## CreateMiningHashTransferConfigCancelV1

> CreateMiningHashTransferConfigCancelV1Resp CreateMiningHashTransferConfigCancelV1(ctx).ConfigId(configId).Timestamp(timestamp).UserName(userName).RecvWindow(recvWindow).Execute()

Cancel hashrate resale configuration(USER_DATA)

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
	configId := int32(56) // int32 | 
	timestamp := int64(789) // int64 | 
	userName := "userName_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.CreateMiningHashTransferConfigCancelV1(context.Background()).ConfigId(configId).Timestamp(timestamp).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.CreateMiningHashTransferConfigCancelV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMiningHashTransferConfigCancelV1`: CreateMiningHashTransferConfigCancelV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.CreateMiningHashTransferConfigCancelV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMiningHashTransferConfigCancelV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **int32** |  | 
 **timestamp** | **int64** |  | 
 **userName** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateMiningHashTransferConfigCancelV1Resp**](CreateMiningHashTransferConfigCancelV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMiningHashTransferConfigV1

> CreateMiningHashTransferConfigV1Resp CreateMiningHashTransferConfigV1(ctx).Algo(algo).EndDate(endDate).HashRate(hashRate).StartDate(startDate).Timestamp(timestamp).ToPoolUser(toPoolUser).UserName(userName).RecvWindow(recvWindow).Execute()

Hashrate Resale Request(USER_DATA)



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
	algo := "algo_example" // string |  (default to "")
	endDate := int64(789) // int64 | 
	hashRate := int64(789) // int64 | 
	startDate := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	toPoolUser := "toPoolUser_example" // string |  (default to "")
	userName := "userName_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.CreateMiningHashTransferConfigV1(context.Background()).Algo(algo).EndDate(endDate).HashRate(hashRate).StartDate(startDate).Timestamp(timestamp).ToPoolUser(toPoolUser).UserName(userName).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.CreateMiningHashTransferConfigV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMiningHashTransferConfigV1`: CreateMiningHashTransferConfigV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.CreateMiningHashTransferConfigV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMiningHashTransferConfigV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** |  | [default to &quot;&quot;]
 **endDate** | **int64** |  | 
 **hashRate** | **int64** |  | 
 **startDate** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **toPoolUser** | **string** |  | [default to &quot;&quot;]
 **userName** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateMiningHashTransferConfigV1Resp**](CreateMiningHashTransferConfigV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningHashTransferConfigDetailsListV1

> GetMiningHashTransferConfigDetailsListV1Resp GetMiningHashTransferConfigDetailsListV1(ctx).Timestamp(timestamp).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale List (USER_DATA)



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
	pageIndex := int32(56) // int32 | Page number, empty default first page, starting from 1 (optional)
	pageSize := int32(56) // int32 | Number of pages, minimum 10, maximum 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningHashTransferConfigDetailsListV1(context.Background()).Timestamp(timestamp).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningHashTransferConfigDetailsListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningHashTransferConfigDetailsListV1`: GetMiningHashTransferConfigDetailsListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningHashTransferConfigDetailsListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningHashTransferConfigDetailsListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **pageIndex** | **int32** | Page number, empty default first page, starting from 1 | 
 **pageSize** | **int32** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningHashTransferConfigDetailsListV1Resp**](GetMiningHashTransferConfigDetailsListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningHashTransferProfitDetailsV1

> GetMiningHashTransferProfitDetailsV1Resp GetMiningHashTransferProfitDetailsV1(ctx).ConfigId(configId).UserName(userName).Timestamp(timestamp).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Hashrate Resale Detail(USER_DATA)



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
	configId := int32(56) // int32 | Mining ID
	userName := "userName_example" // string | Mining Account (default to "")
	timestamp := int64(789) // int64 | 
	pageIndex := int32(56) // int32 | Page number, empty default first page, starting from 1 (optional)
	pageSize := int32(56) // int32 | Number of pages, minimum 10, maximum 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningHashTransferProfitDetailsV1(context.Background()).ConfigId(configId).UserName(userName).Timestamp(timestamp).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningHashTransferProfitDetailsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningHashTransferProfitDetailsV1`: GetMiningHashTransferProfitDetailsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningHashTransferProfitDetailsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningHashTransferProfitDetailsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configId** | **int32** | Mining ID | 
 **userName** | **string** | Mining Account | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **pageIndex** | **int32** | Page number, empty default first page, starting from 1 | 
 **pageSize** | **int32** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningHashTransferProfitDetailsV1Resp**](GetMiningHashTransferProfitDetailsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningPaymentListV1

> GetMiningPaymentListV1Resp GetMiningPaymentListV1(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Earnings List(USER_DATA)



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
	algo := "algo_example" // string | Transfer algorithm(sha256) (default to "")
	userName := "userName_example" // string | Mining account (default to "")
	timestamp := int64(789) // int64 | 
	coin := "coin_example" // string | Coin name (optional) (default to "")
	startDate := int64(789) // int64 | Search date, millisecond timestamp, while empty query all (optional)
	endDate := int64(789) // int64 | Search date, millisecond timestamp, while empty query all (optional)
	pageIndex := int32(56) // int32 | Page number, empty default first page, starting from 1 (optional)
	pageSize := int32(56) // int32 | Number of pages, minimum 10, maximum 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningPaymentListV1(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningPaymentListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningPaymentListV1`: GetMiningPaymentListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningPaymentListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningPaymentListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Transfer algorithm(sha256) | [default to &quot;&quot;]
 **userName** | **string** | Mining account | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **coin** | **string** | Coin name | [default to &quot;&quot;]
 **startDate** | **int64** | Search date, millisecond timestamp, while empty query all | 
 **endDate** | **int64** | Search date, millisecond timestamp, while empty query all | 
 **pageIndex** | **int32** | Page number, empty default first page, starting from 1 | 
 **pageSize** | **int32** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningPaymentListV1Resp**](GetMiningPaymentListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningPaymentOtherV1

> GetMiningPaymentOtherV1Resp GetMiningPaymentOtherV1(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Extra Bonus List(USER_DATA)



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
	algo := "algo_example" // string | Transfer algorithm(sha256) (default to "")
	userName := "userName_example" // string | Mining Account (default to "")
	timestamp := int64(789) // int64 | 
	coin := "coin_example" // string | Coin Name (optional) (default to "")
	startDate := int64(789) // int64 | Search date, millisecond timestamp, while empty query all (optional)
	endDate := int64(789) // int64 | Search date, millisecond timestamp, while empty query all (optional)
	pageIndex := int32(56) // int32 | Page number, empty default first page, starting from 1 (optional)
	pageSize := int32(56) // int32 | Number of pages, minimum 10, maximum 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningPaymentOtherV1(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).Coin(coin).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningPaymentOtherV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningPaymentOtherV1`: GetMiningPaymentOtherV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningPaymentOtherV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningPaymentOtherV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Transfer algorithm(sha256) | [default to &quot;&quot;]
 **userName** | **string** | Mining Account | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **coin** | **string** | Coin Name | [default to &quot;&quot;]
 **startDate** | **int64** | Search date, millisecond timestamp, while empty query all | 
 **endDate** | **int64** | Search date, millisecond timestamp, while empty query all | 
 **pageIndex** | **int32** | Page number, empty default first page, starting from 1 | 
 **pageSize** | **int32** | Number of pages, minimum 10, maximum 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningPaymentOtherV1Resp**](GetMiningPaymentOtherV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningPaymentUidV1

> GetMiningPaymentUidV1Resp GetMiningPaymentUidV1(ctx).Algo(algo).Timestamp(timestamp).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()

Mining Account Earning(USER_DATA)



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
	algo := "algo_example" // string | Algorithm(sha256) (default to "")
	timestamp := int64(789) // int64 | 
	startDate := int64(789) // int64 | Millisecond timestamp (optional)
	endDate := int64(789) // int64 | Millisecond timestamp (optional)
	pageIndex := int32(56) // int32 | Default 1 (optional) (default to 1)
	pageSize := int32(56) // int32 | Min 10,Max 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningPaymentUidV1(context.Background()).Algo(algo).Timestamp(timestamp).StartDate(startDate).EndDate(endDate).PageIndex(pageIndex).PageSize(pageSize).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningPaymentUidV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningPaymentUidV1`: GetMiningPaymentUidV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningPaymentUidV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningPaymentUidV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startDate** | **int64** | Millisecond timestamp | 
 **endDate** | **int64** | Millisecond timestamp | 
 **pageIndex** | **int32** | Default 1 | [default to 1]
 **pageSize** | **int32** | Min 10,Max 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningPaymentUidV1Resp**](GetMiningPaymentUidV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningPubAlgoListV1

> GetMiningPubAlgoListV1Resp GetMiningPubAlgoListV1(ctx).Execute()

Acquiring Algorithm(MARKET_DATA)



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
	resp, r, err := apiClient.MiningAPI.GetMiningPubAlgoListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningPubAlgoListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningPubAlgoListV1`: GetMiningPubAlgoListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningPubAlgoListV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningPubAlgoListV1Request struct via the builder pattern


### Return type

[**GetMiningPubAlgoListV1Resp**](GetMiningPubAlgoListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningPubCoinListV1

> GetMiningPubCoinListV1Resp GetMiningPubCoinListV1(ctx).Execute()

Acquiring CoinName(MARKET_DATA)



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
	resp, r, err := apiClient.MiningAPI.GetMiningPubCoinListV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningPubCoinListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningPubCoinListV1`: GetMiningPubCoinListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningPubCoinListV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningPubCoinListV1Request struct via the builder pattern


### Return type

[**GetMiningPubCoinListV1Resp**](GetMiningPubCoinListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningStatisticsUserListV1

> GetMiningStatisticsUserListV1Resp GetMiningStatisticsUserListV1(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account List(USER_DATA)



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
	algo := "algo_example" // string | Algorithm(sha256) (default to "")
	userName := "userName_example" // string | Mining account (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningStatisticsUserListV1(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningStatisticsUserListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningStatisticsUserListV1`: GetMiningStatisticsUserListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningStatisticsUserListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningStatisticsUserListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | [default to &quot;&quot;]
 **userName** | **string** | Mining account | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningStatisticsUserListV1Resp**](GetMiningStatisticsUserListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningStatisticsUserStatusV1

> GetMiningStatisticsUserStatusV1Resp GetMiningStatisticsUserStatusV1(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Statistic List(USER_DATA)



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
	algo := "algo_example" // string | Algorithm(sha256) (default to "")
	userName := "userName_example" // string | Mining account (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningStatisticsUserStatusV1(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningStatisticsUserStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningStatisticsUserStatusV1`: GetMiningStatisticsUserStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningStatisticsUserStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningStatisticsUserStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | [default to &quot;&quot;]
 **userName** | **string** | Mining account | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningStatisticsUserStatusV1Resp**](GetMiningStatisticsUserStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningWorkerDetailV1

> GetMiningWorkerDetailV1Resp GetMiningWorkerDetailV1(ctx).Algo(algo).UserName(userName).WorkerName(workerName).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Request for Detail Miner List(USER_DATA)



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
	algo := "algo_example" // string | Algorithm(sha256) (default to "")
	userName := "userName_example" // string | Mining account (default to "")
	workerName := "workerName_example" // string | Miner’s name(required) (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningWorkerDetailV1(context.Background()).Algo(algo).UserName(userName).WorkerName(workerName).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningWorkerDetailV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningWorkerDetailV1`: GetMiningWorkerDetailV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningWorkerDetailV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningWorkerDetailV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | [default to &quot;&quot;]
 **userName** | **string** | Mining account | [default to &quot;&quot;]
 **workerName** | **string** | Miner’s name(required) | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningWorkerDetailV1Resp**](GetMiningWorkerDetailV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMiningWorkerListV1

> GetMiningWorkerListV1Resp GetMiningWorkerListV1(ctx).Algo(algo).UserName(userName).Timestamp(timestamp).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()

Request for Miner List(USER_DATA)



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
	algo := "algo_example" // string | Algorithm(sha256) (default to "")
	userName := "userName_example" // string | Mining account (default to "")
	timestamp := int64(789) // int64 | 
	pageIndex := int32(56) // int32 | Page number，default is first page，start form 1 (optional)
	sort := int32(56) // int32 | sort sequence(default=0)0 positive sequence，1 negative sequence (optional)
	sortColumn := int32(56) // int32 | Sort by( default 1): <br/><br/>1: miner name, <br/><br/>2: real-time computing power, <br/><br/>3: daily average computing power, <br/><br/>4: real-time rejection rate, <br/><br/>5: last submission time (optional) (default to 1)
	workerStatus := int32(56) // int32 | miners status(default=0),0 all，1 valid，2 invalid，3 failure (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiningAPI.GetMiningWorkerListV1(context.Background()).Algo(algo).UserName(userName).Timestamp(timestamp).PageIndex(pageIndex).Sort(sort).SortColumn(sortColumn).WorkerStatus(workerStatus).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiningAPI.GetMiningWorkerListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMiningWorkerListV1`: GetMiningWorkerListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MiningAPI.GetMiningWorkerListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMiningWorkerListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algo** | **string** | Algorithm(sha256) | [default to &quot;&quot;]
 **userName** | **string** | Mining account | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **pageIndex** | **int32** | Page number，default is first page，start form 1 | 
 **sort** | **int32** | sort sequence(default&#x3D;0)0 positive sequence，1 negative sequence | 
 **sortColumn** | **int32** | Sort by( default 1): &lt;br/&gt;&lt;br/&gt;1: miner name, &lt;br/&gt;&lt;br/&gt;2: real-time computing power, &lt;br/&gt;&lt;br/&gt;3: daily average computing power, &lt;br/&gt;&lt;br/&gt;4: real-time rejection rate, &lt;br/&gt;&lt;br/&gt;5: last submission time | [default to 1]
 **workerStatus** | **int32** | miners status(default&#x3D;0),0 all，1 valid，2 invalid，3 failure | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetMiningWorkerListV1Resp**](GetMiningWorkerListV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

