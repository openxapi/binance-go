# \V3API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletCreateAssetGetUserAssetV3**](V3API.md#WalletCreateAssetGetUserAssetV3) | **Post** /sapi/v3/asset/getUserAsset | User Asset (USER_DATA)



## WalletCreateAssetGetUserAssetV3

> []WalletCreateAssetGetUserAssetV3RespItem WalletCreateAssetGetUserAssetV3(ctx).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()

User Asset (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	needBtcValuation := true // bool |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V3API.WalletCreateAssetGetUserAssetV3(context.Background()).Timestamp(timestamp).Asset(asset).NeedBtcValuation(needBtcValuation).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V3API.WalletCreateAssetGetUserAssetV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAssetGetUserAssetV3`: []WalletCreateAssetGetUserAssetV3RespItem
	fmt.Fprintf(os.Stdout, "Response from `V3API.WalletCreateAssetGetUserAssetV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAssetGetUserAssetV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **needBtcValuation** | **bool** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletCreateAssetGetUserAssetV3RespItem**](WalletCreateAssetGetUserAssetV3RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

