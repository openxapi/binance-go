# \UserDataStreamsAPI

All URIs are relative to *https://dapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfuturesCreateListenKeyV1**](UserDataStreamsAPI.md#CfuturesCreateListenKeyV1) | **Post** /dapi/v1/listenKey | Start User Data Stream (USER_STREAM)
[**CfuturesDeleteListenKeyV1**](UserDataStreamsAPI.md#CfuturesDeleteListenKeyV1) | **Delete** /dapi/v1/listenKey | Close User Data Stream(USER_STREAM)
[**CfuturesUpdateListenKeyV1**](UserDataStreamsAPI.md#CfuturesUpdateListenKeyV1) | **Put** /dapi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)



## CfuturesCreateListenKeyV1

> CfuturesCreateListenKeyV1Resp CfuturesCreateListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.CfuturesCreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.CfuturesCreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesCreateListenKeyV1`: CfuturesCreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.CfuturesCreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesCreateListenKeyV1Request struct via the builder pattern


### Return type

[**CfuturesCreateListenKeyV1Resp**](CfuturesCreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfuturesDeleteListenKeyV1

> map[string]interface{} CfuturesDeleteListenKeyV1(ctx).Execute()

Close User Data Stream(USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.CfuturesDeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.CfuturesDeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesDeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.CfuturesDeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesDeleteListenKeyV1Request struct via the builder pattern


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


## CfuturesUpdateListenKeyV1

> map[string]interface{} CfuturesUpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.CfuturesUpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.CfuturesUpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesUpdateListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.CfuturesUpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesUpdateListenKeyV1Request struct via the builder pattern


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

