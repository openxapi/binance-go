# \MarketMakerEndpointsAPI

All URIs are relative to *https://eapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OptionsCreateCountdownCancelAllHeartBeatV1**](MarketMakerEndpointsAPI.md#OptionsCreateCountdownCancelAllHeartBeatV1) | **Post** /eapi/v1/countdownCancelAllHeartBeat | Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)
[**OptionsCreateCountdownCancelAllV1**](MarketMakerEndpointsAPI.md#OptionsCreateCountdownCancelAllV1) | **Post** /eapi/v1/countdownCancelAll | Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**OptionsCreateMmpResetV1**](MarketMakerEndpointsAPI.md#OptionsCreateMmpResetV1) | **Post** /eapi/v1/mmpReset | Reset Market Maker Protection Config (TRADE)
[**OptionsCreateMmpSetV1**](MarketMakerEndpointsAPI.md#OptionsCreateMmpSetV1) | **Post** /eapi/v1/mmpSet | Set Market Maker Protection Config (TRADE)
[**OptionsGetCountdownCancelAllV1**](MarketMakerEndpointsAPI.md#OptionsGetCountdownCancelAllV1) | **Get** /eapi/v1/countdownCancelAll | Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)
[**OptionsGetMarginAccountV1**](MarketMakerEndpointsAPI.md#OptionsGetMarginAccountV1) | **Get** /eapi/v1/marginAccount | Option Margin Account Information (USER_DATA)
[**OptionsGetMmpV1**](MarketMakerEndpointsAPI.md#OptionsGetMmpV1) | **Get** /eapi/v1/mmp | Get Market Maker Protection Config (TRADE)



## OptionsCreateCountdownCancelAllHeartBeatV1

> OptionsCreateCountdownCancelAllHeartBeatV1Resp OptionsCreateCountdownCancelAllHeartBeatV1(ctx).Timestamp(timestamp).Underlyings(underlyings).RecvWindow(recvWindow).Execute()

Auto-Cancel All Open Orders (Kill-Switch) Heartbeat (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	underlyings := "underlyings_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerEndpointsAPI.OptionsCreateCountdownCancelAllHeartBeatV1(context.Background()).Timestamp(timestamp).Underlyings(underlyings).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.OptionsCreateCountdownCancelAllHeartBeatV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateCountdownCancelAllHeartBeatV1`: OptionsCreateCountdownCancelAllHeartBeatV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerEndpointsAPI.OptionsCreateCountdownCancelAllHeartBeatV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateCountdownCancelAllHeartBeatV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlyings** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsCreateCountdownCancelAllHeartBeatV1Resp**](OptionsCreateCountdownCancelAllHeartBeatV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateCountdownCancelAllV1

> OptionsCreateCountdownCancelAllV1Resp OptionsCreateCountdownCancelAllV1(ctx).CountdownTime(countdownTime).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Set Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	countdownTime := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerEndpointsAPI.OptionsCreateCountdownCancelAllV1(context.Background()).CountdownTime(countdownTime).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.OptionsCreateCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateCountdownCancelAllV1`: OptionsCreateCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerEndpointsAPI.OptionsCreateCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countdownTime** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsCreateCountdownCancelAllV1Resp**](OptionsCreateCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateMmpResetV1

> OptionsCreateMmpResetV1Resp OptionsCreateMmpResetV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Underlying(underlying).Execute()

Reset Market Maker Protection Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerEndpointsAPI.OptionsCreateMmpResetV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Underlying(underlying).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.OptionsCreateMmpResetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateMmpResetV1`: OptionsCreateMmpResetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerEndpointsAPI.OptionsCreateMmpResetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateMmpResetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]

### Return type

[**OptionsCreateMmpResetV1Resp**](OptionsCreateMmpResetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsCreateMmpSetV1

> OptionsCreateMmpSetV1Resp OptionsCreateMmpSetV1(ctx).Timestamp(timestamp).DeltaLimit(deltaLimit).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).RecvWindow(recvWindow).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).Execute()

Set Market Maker Protection Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	deltaLimit := "deltaLimit_example" // string |  (optional) (default to "")
	frozenTimeInMilliseconds := int64(789) // int64 |  (optional)
	qtyLimit := "qtyLimit_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	underlying := "underlying_example" // string |  (optional) (default to "")
	windowTimeInMilliseconds := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerEndpointsAPI.OptionsCreateMmpSetV1(context.Background()).Timestamp(timestamp).DeltaLimit(deltaLimit).FrozenTimeInMilliseconds(frozenTimeInMilliseconds).QtyLimit(qtyLimit).RecvWindow(recvWindow).Underlying(underlying).WindowTimeInMilliseconds(windowTimeInMilliseconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.OptionsCreateMmpSetV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsCreateMmpSetV1`: OptionsCreateMmpSetV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerEndpointsAPI.OptionsCreateMmpSetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsCreateMmpSetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **deltaLimit** | **string** |  | [default to &quot;&quot;]
 **frozenTimeInMilliseconds** | **int64** |  | 
 **qtyLimit** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **underlying** | **string** |  | [default to &quot;&quot;]
 **windowTimeInMilliseconds** | **int64** |  | 

### Return type

[**OptionsCreateMmpSetV1Resp**](OptionsCreateMmpSetV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetCountdownCancelAllV1

> OptionsGetCountdownCancelAllV1Resp OptionsGetCountdownCancelAllV1(ctx).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Auto-Cancel All Open Orders (Kill-Switch) Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string | Option underlying, e.g BTCUSDT (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerEndpointsAPI.OptionsGetCountdownCancelAllV1(context.Background()).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.OptionsGetCountdownCancelAllV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetCountdownCancelAllV1`: OptionsGetCountdownCancelAllV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerEndpointsAPI.OptionsGetCountdownCancelAllV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetCountdownCancelAllV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlying** | **string** | Option underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetCountdownCancelAllV1Resp**](OptionsGetCountdownCancelAllV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetMarginAccountV1

> OptionsGetMarginAccountV1Resp OptionsGetMarginAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Option Margin Account Information (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerEndpointsAPI.OptionsGetMarginAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.OptionsGetMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetMarginAccountV1`: OptionsGetMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerEndpointsAPI.OptionsGetMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetMarginAccountV1Resp**](OptionsGetMarginAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsGetMmpV1

> OptionsGetMmpV1Resp OptionsGetMmpV1(ctx).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()

Get Market Maker Protection Config (TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/options"
)

func main() {
	timestamp := int64(789) // int64 | 
	underlying := "underlying_example" // string | underlying, e.g BTCUSDT (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketMakerEndpointsAPI.OptionsGetMmpV1(context.Background()).Timestamp(timestamp).Underlying(underlying).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketMakerEndpointsAPI.OptionsGetMmpV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OptionsGetMmpV1`: OptionsGetMmpV1Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketMakerEndpointsAPI.OptionsGetMmpV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOptionsGetMmpV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **underlying** | **string** | underlying, e.g BTCUSDT | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**OptionsGetMmpV1Resp**](OptionsGetMmpV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

