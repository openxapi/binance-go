# \FiatAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFiatOrdersV1**](FiatAPI.md#GetFiatOrdersV1) | **Get** /sapi/v1/fiat/orders | Get Fiat Deposit/Withdraw History (USER_DATA)
[**GetFiatPaymentsV1**](FiatAPI.md#GetFiatPaymentsV1) | **Get** /sapi/v1/fiat/payments | Get Fiat Payments History (USER_DATA)



## GetFiatOrdersV1

> GetFiatOrdersV1Resp GetFiatOrdersV1(ctx).TransactionType(transactionType).Timestamp(timestamp).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()

Get Fiat Deposit/Withdraw History (USER_DATA)



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
	transactionType := "transactionType_example" // string | 0-deposit,1-withdraw (default to "")
	timestamp := int64(789) // int64 | 
	beginTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	page := int32(56) // int32 | default 1 (optional) (default to 1)
	rows := int32(56) // int32 | default 100, max 500 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatAPI.GetFiatOrdersV1(context.Background()).TransactionType(transactionType).Timestamp(timestamp).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatAPI.GetFiatOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatOrdersV1`: GetFiatOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `FiatAPI.GetFiatOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionType** | **string** | 0-deposit,1-withdraw | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int32** | default 1 | [default to 1]
 **rows** | **int32** | default 100, max 500 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**GetFiatOrdersV1Resp**](GetFiatOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatPaymentsV1

> GetFiatPaymentsV1Resp GetFiatPaymentsV1(ctx).TransactionType(transactionType).Timestamp(timestamp).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()

Get Fiat Payments History (USER_DATA)



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
	transactionType := "transactionType_example" // string | 0-buy,1-sell (default to "")
	timestamp := int64(789) // int64 | 
	beginTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	page := int32(56) // int32 | default 1 (optional) (default to 1)
	rows := int32(56) // int32 | default 100, max 500 (optional) (default to 100)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiatAPI.GetFiatPaymentsV1(context.Background()).TransactionType(transactionType).Timestamp(timestamp).BeginTime(beginTime).EndTime(endTime).Page(page).Rows(rows).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiatAPI.GetFiatPaymentsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatPaymentsV1`: GetFiatPaymentsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `FiatAPI.GetFiatPaymentsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatPaymentsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionType** | **string** | 0-buy,1-sell | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **beginTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **page** | **int32** | default 1 | [default to 1]
 **rows** | **int32** | default 100, max 500 | [default to 100]
 **recvWindow** | **int64** |  | 

### Return type

[**GetFiatPaymentsV1Resp**](GetFiatPaymentsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

