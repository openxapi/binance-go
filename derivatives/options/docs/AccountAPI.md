# \AccountAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsGetAccountV1**](AccountAPI.md#OptionsGetAccountV1) | **Get** /eapi/v1/account | Option Account Information(TRADE)
[**OptionsGetBillV1**](AccountAPI.md#OptionsGetBillV1) | **Get** /eapi/v1/bill | Account Funding Flow (USER_DATA)
[**OptionsGetIncomeAsynIdV1**](AccountAPI.md#OptionsGetIncomeAsynIdV1) | **Get** /eapi/v1/income/asyn/id | Get Option Transaction History Download Link by Id (USER_DATA)
[**OptionsGetIncomeAsynV1**](AccountAPI.md#OptionsGetIncomeAsynV1) | **Get** /eapi/v1/income/asyn | Get Download Id For Option Transaction History (USER_DATA)



## OptionsGetAccountV1

> OptionsGetAccountV1Resp OptionsGetAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Option Account Information(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.OptionsGetAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.OptionsGetAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetAccountV1`: OptionsGetAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.OptionsGetAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetAccountV1Resp**](OptionsGetAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetBillV1

> []OptionsGetBillV1RespItem OptionsGetBillV1(ctx).Currency(currency).Timestamp(timestamp).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Account Funding Flow (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	currency := "currency_example" // string | Asset type, only support USDT  as of now (default to "")
	timestamp := int64(789) // int64 | 
	recordId := int64(789) // int64 | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 (optional)
	startTime := int64(789) // int64 | Start Time, e.g 1593511200000 (optional)
	endTime := int64(789) // int64 | End Time, e.g 1593512200000 (optional)
	limit := int32(56) // int32 | Number of result sets returned Default:100 Max:1000 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.OptionsGetBillV1(context.Background()).Currency(currency).Timestamp(timestamp).RecordId(recordId).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.OptionsGetBillV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetBillV1`: []OptionsGetBillV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.OptionsGetBillV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetBillV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Asset type, only support USDT  as of now | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recordId** | **int64** | Return the recordId and subsequent data, the latest data is returned by default, e.g 100000 | 
 **startTime** | **int64** | Start Time, e.g 1593511200000 | 
 **endTime** | **int64** | End Time, e.g 1593512200000 | 
 **limit** | **int32** | Number of result sets returned Default:100 Max:1000 | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]OptionsGetBillV1RespItem**](OptionsGetBillV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetIncomeAsynIdV1

> OptionsGetIncomeAsynIdV1Resp OptionsGetIncomeAsynIdV1(ctx).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Option Transaction History Download Link by Id (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	downloadId := "downloadId_example" // string | get by download id api (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.OptionsGetIncomeAsynIdV1(context.Background()).DownloadId(downloadId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.OptionsGetIncomeAsynIdV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetIncomeAsynIdV1`: OptionsGetIncomeAsynIdV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.OptionsGetIncomeAsynIdV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetIncomeAsynIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadId** | **string** | get by download id api | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetIncomeAsynIdV1Resp**](OptionsGetIncomeAsynIdV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetIncomeAsynV1

> OptionsGetIncomeAsynV1Resp OptionsGetIncomeAsynV1(ctx).Execute()

Get Download Id For Option Transaction History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.OptionsGetIncomeAsynV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.OptionsGetIncomeAsynV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetIncomeAsynV1`: OptionsGetIncomeAsynV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.OptionsGetIncomeAsynV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetIncomeAsynV1Request struct via the builder pattern


### Return type

[**OptionsGetIncomeAsynV1Resp**](OptionsGetIncomeAsynV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

