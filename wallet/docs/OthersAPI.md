# \OthersAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletGetSystemStatusV1**](OthersAPI.md#WalletGetSystemStatusV1) | **Get** /sapi/v1/system/status | System Status (System)



## WalletGetSystemStatusV1

> WalletGetSystemStatusV1Resp WalletGetSystemStatusV1(ctx).Execute()

System Status (System)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/wallet"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OthersAPI.WalletGetSystemStatusV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OthersAPI.WalletGetSystemStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetSystemStatusV1`: WalletGetSystemStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `OthersAPI.WalletGetSystemStatusV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetSystemStatusV1Request struct via the builder pattern


### Return type

[**WalletGetSystemStatusV1Resp**](WalletGetSystemStatusV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

