# \AssetAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletGetAssetAssetDetailV1**](AssetAPI.md#WalletGetAssetAssetDetailV1) | **Get** /sapi/v1/asset/assetDetail | Asset Detail (USER_DATA)



## WalletGetAssetAssetDetailV1

> map[string]WalletGetAssetAssetDetailV1RespValue WalletGetAssetAssetDetailV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Asset Detail (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetAPI.WalletGetAssetAssetDetailV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetAPI.WalletGetAssetAssetDetailV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAssetAssetDetailV1`: map[string]WalletGetAssetAssetDetailV1RespValue
	fmt.Fprintf(os.Stdout, "Response from `AssetAPI.WalletGetAssetAssetDetailV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAssetAssetDetailV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**map[string]WalletGetAssetAssetDetailV1RespValue**](WalletGetAssetAssetDetailV1RespValue.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

