# \UserDataStreamAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpotCreateUserDataStreamV3**](UserDataStreamAPI.md#SpotCreateUserDataStreamV3) | **Post** /api/v3/userDataStream | Start user data stream (USER_STREAM)
[**SpotDeleteUserDataStreamV3**](UserDataStreamAPI.md#SpotDeleteUserDataStreamV3) | **Delete** /api/v3/userDataStream | Close user data stream (USER_STREAM)
[**SpotUpdateUserDataStreamV3**](UserDataStreamAPI.md#SpotUpdateUserDataStreamV3) | **Put** /api/v3/userDataStream | Keepalive user data stream (USER_STREAM)



## SpotCreateUserDataStreamV3

> SpotCreateUserDataStreamV3Resp SpotCreateUserDataStreamV3(ctx).Execute()

Start user data stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamAPI.SpotCreateUserDataStreamV3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamAPI.SpotCreateUserDataStreamV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotCreateUserDataStreamV3`: SpotCreateUserDataStreamV3Resp
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamAPI.SpotCreateUserDataStreamV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpotCreateUserDataStreamV3Request struct via the builder pattern


### Return type

[**SpotCreateUserDataStreamV3Resp**](SpotCreateUserDataStreamV3Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotDeleteUserDataStreamV3

> map[string]interface{} SpotDeleteUserDataStreamV3(ctx).ListenKey(listenKey).Execute()

Close user data stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamAPI.SpotDeleteUserDataStreamV3(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamAPI.SpotDeleteUserDataStreamV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotDeleteUserDataStreamV3`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamAPI.SpotDeleteUserDataStreamV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotDeleteUserDataStreamV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpotUpdateUserDataStreamV3

> map[string]interface{} SpotUpdateUserDataStreamV3(ctx).ListenKey(listenKey).Execute()

Keepalive user data stream (USER_STREAM)



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
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamAPI.SpotUpdateUserDataStreamV3(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamAPI.SpotUpdateUserDataStreamV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotUpdateUserDataStreamV3`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamAPI.SpotUpdateUserDataStreamV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotUpdateUserDataStreamV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

