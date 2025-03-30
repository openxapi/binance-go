# \V2API

All URIs are relative to *https://papi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginGetUmAccountV2**](V2API.md#PmarginGetUmAccountV2) | **Get** /papi/v2/um/account | Get UM Account Detail V2(USER_DATA)



## PmarginGetUmAccountV2

> PmarginGetUmAccountV2Resp PmarginGetUmAccountV2(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get UM Account Detail V2(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/pmargin"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.PmarginGetUmAccountV2(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.PmarginGetUmAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginGetUmAccountV2`: PmarginGetUmAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.PmarginGetUmAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPmarginGetUmAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**PmarginGetUmAccountV2Resp**](PmarginGetUmAccountV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

