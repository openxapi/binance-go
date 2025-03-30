# \TradeDataStreamAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCreateUserDataStreamIsolatedV1**](TradeDataStreamAPI.md#MarginCreateUserDataStreamIsolatedV1) | **Post** /sapi/v1/userDataStream/isolated | Start Isolated Margin User Data Stream (USER_STREAM)
[**MarginCreateUserDataStreamV1**](TradeDataStreamAPI.md#MarginCreateUserDataStreamV1) | **Post** /sapi/v1/userDataStream | Start Margin User Data Stream (USER_STREAM)
[**MarginDeleteUserDataStreamIsolatedV1**](TradeDataStreamAPI.md#MarginDeleteUserDataStreamIsolatedV1) | **Delete** /sapi/v1/userDataStream/isolated | Close Isolated Margin User Data Stream (USER_STREAM)
[**MarginDeleteUserDataStreamV1**](TradeDataStreamAPI.md#MarginDeleteUserDataStreamV1) | **Delete** /sapi/v1/userDataStream | Close Margin User Data Stream (USER_STREAM)
[**MarginUpdateUserDataStreamIsolatedV1**](TradeDataStreamAPI.md#MarginUpdateUserDataStreamIsolatedV1) | **Put** /sapi/v1/userDataStream/isolated | Keepalive Isolated Margin User Data Stream (USER_STREAM)
[**MarginUpdateUserDataStreamV1**](TradeDataStreamAPI.md#MarginUpdateUserDataStreamV1) | **Put** /sapi/v1/userDataStream | Keepalive Margin User Data Stream (USER_STREAM)



## MarginCreateUserDataStreamIsolatedV1

> MarginCreateUserDataStreamIsolatedV1Resp MarginCreateUserDataStreamIsolatedV1(ctx).Symbol(symbol).Execute()

Start Isolated Margin User Data Stream (USER_STREAM)



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
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeDataStreamAPI.MarginCreateUserDataStreamIsolatedV1(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeDataStreamAPI.MarginCreateUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateUserDataStreamIsolatedV1`: MarginCreateUserDataStreamIsolatedV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeDataStreamAPI.MarginCreateUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateUserDataStreamIsolatedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]

### Return type

[**MarginCreateUserDataStreamIsolatedV1Resp**](MarginCreateUserDataStreamIsolatedV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateUserDataStreamV1

> MarginCreateUserDataStreamV1Resp MarginCreateUserDataStreamV1(ctx).Execute()

Start Margin User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.TradeDataStreamAPI.MarginCreateUserDataStreamV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeDataStreamAPI.MarginCreateUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateUserDataStreamV1`: MarginCreateUserDataStreamV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeDataStreamAPI.MarginCreateUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateUserDataStreamV1Request struct via the builder pattern


### Return type

[**MarginCreateUserDataStreamV1Resp**](MarginCreateUserDataStreamV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteUserDataStreamIsolatedV1

> map[string]interface{} MarginDeleteUserDataStreamIsolatedV1(ctx).Symbol(symbol).Listenkey(listenkey).Execute()

Close Isolated Margin User Data Stream (USER_STREAM)



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
	symbol := "symbol_example" // string |  (default to "")
	listenkey := "listenkey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeDataStreamAPI.MarginDeleteUserDataStreamIsolatedV1(context.Background()).Symbol(symbol).Listenkey(listenkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeDataStreamAPI.MarginDeleteUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteUserDataStreamIsolatedV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TradeDataStreamAPI.MarginDeleteUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteUserDataStreamIsolatedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **listenkey** | **string** |  | [default to &quot;&quot;]

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


## MarginDeleteUserDataStreamV1

> map[string]interface{} MarginDeleteUserDataStreamV1(ctx).Listenkey(listenkey).Execute()

Close Margin User Data Stream (USER_STREAM)



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
	listenkey := "listenkey_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeDataStreamAPI.MarginDeleteUserDataStreamV1(context.Background()).Listenkey(listenkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeDataStreamAPI.MarginDeleteUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteUserDataStreamV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TradeDataStreamAPI.MarginDeleteUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteUserDataStreamV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenkey** | **string** |  | [default to &quot;&quot;]

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


## MarginUpdateUserDataStreamIsolatedV1

> map[string]interface{} MarginUpdateUserDataStreamIsolatedV1(ctx).ListenKey(listenKey).Symbol(symbol).Execute()

Keepalive Isolated Margin User Data Stream (USER_STREAM)



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
	symbol := "symbol_example" // string |  (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeDataStreamAPI.MarginUpdateUserDataStreamIsolatedV1(context.Background()).ListenKey(listenKey).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeDataStreamAPI.MarginUpdateUserDataStreamIsolatedV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateUserDataStreamIsolatedV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TradeDataStreamAPI.MarginUpdateUserDataStreamIsolatedV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateUserDataStreamIsolatedV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listenKey** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** |  | [default to &quot;&quot;]

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


## MarginUpdateUserDataStreamV1

> map[string]interface{} MarginUpdateUserDataStreamV1(ctx).ListenKey(listenKey).Execute()

Keepalive Margin User Data Stream (USER_STREAM)



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
	resp, r, err := apiClient.TradeDataStreamAPI.MarginUpdateUserDataStreamV1(context.Background()).ListenKey(listenKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeDataStreamAPI.MarginUpdateUserDataStreamV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginUpdateUserDataStreamV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TradeDataStreamAPI.MarginUpdateUserDataStreamV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginUpdateUserDataStreamV1Request struct via the builder pattern


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

