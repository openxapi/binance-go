# \AccountManagementAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountGetSubAccountTransactionStatisticsV1**](AccountManagementAPI.md#SubaccountGetSubAccountTransactionStatisticsV1) | **Get** /sapi/v1/sub-account/transaction-statistics | Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)



## SubaccountGetSubAccountTransactionStatisticsV1

> SubaccountGetSubAccountTransactionStatisticsV1Resp SubaccountGetSubAccountTransactionStatisticsV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)



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
	email := "email_example" // string | Sub user email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountGetSubAccountTransactionStatisticsV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountGetSubAccountTransactionStatisticsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountTransactionStatisticsV1`: SubaccountGetSubAccountTransactionStatisticsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountGetSubAccountTransactionStatisticsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountTransactionStatisticsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub user email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountTransactionStatisticsV1Resp**](SubaccountGetSubAccountTransactionStatisticsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

