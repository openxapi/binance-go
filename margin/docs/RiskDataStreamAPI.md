# \RiskDataStreamAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCreateMarginListenKeyV1**](RiskDataStreamAPI.md#MarginCreateMarginListenKeyV1) | **Post** /sapi/v1/margin/listen-key | Start User Data Stream (USER_STREAM)
[**MarginDeleteMarginListenKeyV1**](RiskDataStreamAPI.md#MarginDeleteMarginListenKeyV1) | **Delete** /sapi/v1/margin/listen-key | Close User Data Stream (USER_STREAM)
[**MarginUpdateMarginListenKeyV1**](RiskDataStreamAPI.md#MarginUpdateMarginListenKeyV1) | **Put** /sapi/v1/margin/listen-key | Keepalive User Data Stream (USER_STREAM)



## MarginCreateMarginListenKeyV1

> MarginCreateMarginListenKeyV1Resp MarginCreateMarginListenKeyV1(ctx).Execute()

Start User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskDataStreamAPI.MarginCreateMarginListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskDataStreamAPI.MarginCreateMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginListenKeyV1`: MarginCreateMarginListenKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `RiskDataStreamAPI.MarginCreateMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginListenKeyV1Request struct via the builder pattern


### Return type

[**MarginCreateMarginListenKeyV1Resp**](MarginCreateMarginListenKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginListenKeyV1

> map[string]interface{} MarginDeleteMarginListenKeyV1(ctx).Execute()

Close User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskDataStreamAPI.MarginDeleteMarginListenKeyV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskDataStreamAPI.MarginDeleteMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RiskDataStreamAPI.MarginDeleteMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginListenKeyV1Request struct via the builder pattern


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


## MarginUpdateMarginListenKeyV1

> map[string]interface{} MarginUpdateMarginListenKeyV1(ctx).ListenKey(listenKey).Execute()

Keepalive User Data Stream (USER_STREAM)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func main() {
	listenKey := "listenKey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskDataStreamAPI.MarginUpdateMarginListenKeyV1(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskDataStreamAPI.MarginUpdateMarginListenKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateMarginListenKeyV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RiskDataStreamAPI.MarginUpdateMarginListenKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateMarginListenKeyV1Request struct via the builder pattern


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

