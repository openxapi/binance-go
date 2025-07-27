# \NftAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNftHistoryDepositV1**](NftAPI.md#GetNftHistoryDepositV1) | **Get** /sapi/v1/nft/history/deposit | Get NFT Deposit History(USER_DATA)
[**GetNftHistoryTransactionsV1**](NftAPI.md#GetNftHistoryTransactionsV1) | **Get** /sapi/v1/nft/history/transactions | Get NFT Transaction History(USER_DATA)
[**GetNftHistoryWithdrawV1**](NftAPI.md#GetNftHistoryWithdrawV1) | **Get** /sapi/v1/nft/history/withdraw | Get NFT Withdraw History(USER_DATA)
[**GetNftUserGetAssetV1**](NftAPI.md#GetNftUserGetAssetV1) | **Get** /sapi/v1/nft/user/getAsset | Get NFT Asset(USER_DATA)



## GetNftHistoryDepositV1

> GetNftHistoryDepositV1Resp GetNftHistoryDepositV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Deposit History(USER_DATA)



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
	limit := int32(56) // int32 | Default 50, Max 50 (optional) (default to 50)
	page := int32(56) // int32 | Default 1 (optional) (default to 1)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NftAPI.GetNftHistoryDepositV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NftAPI.GetNftHistoryDepositV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftHistoryDepositV1`: GetNftHistoryDepositV1Resp
	fmt.Fprintf(os.Stdout, "Response from `NftAPI.GetNftHistoryDepositV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHistoryDepositV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50, Max 50 | [default to 50]
 **page** | **int32** | Default 1 | [default to 1]
 **recvWindow** | **int64** |  | 

### Return type

[**GetNftHistoryDepositV1Resp**](GetNftHistoryDepositV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftHistoryTransactionsV1

> GetNftHistoryTransactionsV1Resp GetNftHistoryTransactionsV1(ctx).OrderType(orderType).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Transaction History(USER_DATA)



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
	orderType := int32(56) // int32 | 0: purchase order, 1: sell order, 2: royalty income, 3: primary market order, 4: mint fee
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | Default 50, Max 50 (optional) (default to 50)
	page := int32(56) // int32 | Default 1 (optional) (default to 1)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NftAPI.GetNftHistoryTransactionsV1(context.Background()).OrderType(orderType).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NftAPI.GetNftHistoryTransactionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftHistoryTransactionsV1`: GetNftHistoryTransactionsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `NftAPI.GetNftHistoryTransactionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHistoryTransactionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderType** | **int32** | 0: purchase order, 1: sell order, 2: royalty income, 3: primary market order, 4: mint fee | 
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50, Max 50 | [default to 50]
 **page** | **int32** | Default 1 | [default to 1]
 **recvWindow** | **int64** |  | 

### Return type

[**GetNftHistoryTransactionsV1Resp**](GetNftHistoryTransactionsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftHistoryWithdrawV1

> GetNftHistoryWithdrawV1Resp GetNftHistoryWithdrawV1(ctx).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Withdraw History(USER_DATA)



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
	limit := int32(56) // int32 | Default 50, Max 50 (optional) (default to 50)
	page := int32(56) // int32 | Default 1 (optional) (default to 1)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NftAPI.GetNftHistoryWithdrawV1(context.Background()).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NftAPI.GetNftHistoryWithdrawV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftHistoryWithdrawV1`: GetNftHistoryWithdrawV1Resp
	fmt.Fprintf(os.Stdout, "Response from `NftAPI.GetNftHistoryWithdrawV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHistoryWithdrawV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | Default 50, Max 50 | [default to 50]
 **page** | **int32** | Default 1 | [default to 1]
 **recvWindow** | **int64** |  | 

### Return type

[**GetNftHistoryWithdrawV1Resp**](GetNftHistoryWithdrawV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftUserGetAssetV1

> GetNftUserGetAssetV1Resp GetNftUserGetAssetV1(ctx).Timestamp(timestamp).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()

Get NFT Asset(USER_DATA)



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
	limit := int32(56) // int32 | Default 50, Max 50 (optional) (default to 50)
	page := int32(56) // int32 | Default 1 (optional) (default to 1)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NftAPI.GetNftUserGetAssetV1(context.Background()).Timestamp(timestamp).Limit(limit).Page(page).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NftAPI.GetNftUserGetAssetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftUserGetAssetV1`: GetNftUserGetAssetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `NftAPI.GetNftUserGetAssetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftUserGetAssetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **limit** | **int32** | Default 50, Max 50 | [default to 50]
 **page** | **int32** | Default 1 | [default to 1]
 **recvWindow** | **int64** |  | 

### Return type

[**GetNftUserGetAssetV1Resp**](GetNftUserGetAssetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

