# \V4API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountGetSubAccountAssetsV4**](V4API.md#SubaccountGetSubAccountAssetsV4) | **Get** /sapi/v4/sub-account/assets | Query Sub-account Assets (For Master Account)(USER_DATA)



## SubaccountGetSubAccountAssetsV4

> SubaccountGetSubAccountAssetsV4Resp SubaccountGetSubAccountAssetsV4(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Assets (For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | Sub Account Email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V4API.SubaccountGetSubAccountAssetsV4(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V4API.SubaccountGetSubAccountAssetsV4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountAssetsV4`: SubaccountGetSubAccountAssetsV4Resp
	fmt.Fprintf(os.Stdout, "Response from `V4API.SubaccountGetSubAccountAssetsV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountAssetsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub Account Email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountAssetsV4Resp**](SubaccountGetSubAccountAssetsV4Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

