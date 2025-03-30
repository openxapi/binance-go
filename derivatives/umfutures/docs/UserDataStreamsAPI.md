# \UserDataStreamsAPI

All URIs are relative to *https://fapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UfuturesCreateListenKeyV1**](UserDataStreamsAPI.md#UfuturesCreateListenKeyV1) | **Post** /fapi/v1/listenKey | Start User Data Stream (USER_STREAM)
[**UfuturesDeleteListenKeyV1**](UserDataStreamsAPI.md#UfuturesDeleteListenKeyV1) | **Delete** /fapi/v1/listenKey | Close User Data Stream (USER_STREAM)
[**UfuturesUpdateListenKeyV1**](UserDataStreamsAPI.md#UfuturesUpdateListenKeyV1) | **Put** /fapi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)



## UfuturesCreateListenKeyV1

> UfuturesCreateListenKeyV1Resp UfuturesCreateListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.UfuturesCreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.UfuturesCreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesCreateListenKeyV1`: UfuturesCreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.UfuturesCreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesCreateListenKeyV1Request struct via the builder pattern


### Return type

[**UfuturesCreateListenKeyV1Resp**](UfuturesCreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UfuturesDeleteListenKeyV1

> map[string]interface{} UfuturesDeleteListenKeyV1(ctx).Execute()

Close User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.UfuturesDeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.UfuturesDeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesDeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.UfuturesDeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesDeleteListenKeyV1Request struct via the builder pattern


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


## UfuturesUpdateListenKeyV1

> UfuturesUpdateListenKeyV1Resp UfuturesUpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.UfuturesUpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.UfuturesUpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UfuturesUpdateListenKeyV1`: UfuturesUpdateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.UfuturesUpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUfuturesUpdateListenKeyV1Request struct via the builder pattern


### Return type

[**UfuturesUpdateListenKeyV1Resp**](UfuturesUpdateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

