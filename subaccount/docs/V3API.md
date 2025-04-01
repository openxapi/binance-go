# \V3API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountGetSubAccountAssetsV3**](V3API.md#SubaccountGetSubAccountAssetsV3) | **Get** /sapi/v3/sub-account/assets | Query Sub-account Assets(For Master Account)



## SubaccountGetSubAccountAssetsV3

> SubaccountGetSubAccountAssetsV3Resp SubaccountGetSubAccountAssetsV3(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Assets(For Master Account)



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
	email := "email_example" // string | Sub account email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V3API.SubaccountGetSubAccountAssetsV3(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.SubaccountGetSubAccountAssetsV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountAssetsV3`: SubaccountGetSubAccountAssetsV3Resp
	fmt.Fprintf(os.Stdout, "Response from `V3API.SubaccountGetSubAccountAssetsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountAssetsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub account email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountAssetsV3Resp**](SubaccountGetSubAccountAssetsV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

