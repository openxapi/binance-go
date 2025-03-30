# \UserDataStreamsAPI

All URIs are relative to *https://papi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PmarginCreateListenKeyV1**](UserDataStreamsAPI.md#PmarginCreateListenKeyV1) | **Post** /papi/v1/listenKey | Start User Data Stream(USER_STREAM)
[**PmarginDeleteListenKeyV1**](UserDataStreamsAPI.md#PmarginDeleteListenKeyV1) | **Delete** /papi/v1/listenKey | Close User Data Stream(USER_STREAM)
[**PmarginUpdateListenKeyV1**](UserDataStreamsAPI.md#PmarginUpdateListenKeyV1) | **Put** /papi/v1/listenKey | Keepalive User Data Stream (USER_STREAM)



## PmarginCreateListenKeyV1

> PmarginCreateListenKeyV1Resp PmarginCreateListenKeyV1(ctx).Execute()

Start User Data Stream(USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.PmarginCreateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.PmarginCreateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginCreateListenKeyV1`: PmarginCreateListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.PmarginCreateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginCreateListenKeyV1Request struct via the builder pattern


### Return type

[**PmarginCreateListenKeyV1Resp**](PmarginCreateListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PmarginDeleteListenKeyV1

> map[string]interface{} PmarginDeleteListenKeyV1(ctx).Execute()

Close User Data Stream(USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.PmarginDeleteListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.PmarginDeleteListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginDeleteListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.PmarginDeleteListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginDeleteListenKeyV1Request struct via the builder pattern


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


## PmarginUpdateListenKeyV1

> map[string]interface{} PmarginUpdateListenKeyV1(ctx).Execute()

Keepalive User Data Stream (USER_STREAM)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataStreamsAPI.PmarginUpdateListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataStreamsAPI.PmarginUpdateListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PmarginUpdateListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UserDataStreamsAPI.PmarginUpdateListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPmarginUpdateListenKeyV1Request struct via the builder pattern


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

