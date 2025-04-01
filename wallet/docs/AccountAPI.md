# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletCreateAccountDisableFastWithdrawSwitchV1**](AccountAPI.md#WalletCreateAccountDisableFastWithdrawSwitchV1) | **Post** /sapi/v1/account/disableFastWithdrawSwitch | Disable Fast Withdraw Switch (USER_DATA)
[**WalletCreateAccountEnableFastWithdrawSwitchV1**](AccountAPI.md#WalletCreateAccountEnableFastWithdrawSwitchV1) | **Post** /sapi/v1/account/enableFastWithdrawSwitch | Enable Fast Withdraw Switch (USER_DATA)
[**WalletGetAccountApiRestrictionsV1**](AccountAPI.md#WalletGetAccountApiRestrictionsV1) | **Get** /sapi/v1/account/apiRestrictions | Get API Key Permission (USER_DATA)
[**WalletGetAccountApiTradingStatusV1**](AccountAPI.md#WalletGetAccountApiTradingStatusV1) | **Get** /sapi/v1/account/apiTradingStatus | Account API Trading Status (USER_DATA)
[**WalletGetAccountInfoV1**](AccountAPI.md#WalletGetAccountInfoV1) | **Get** /sapi/v1/account/info | Account info (USER_DATA)
[**WalletGetAccountSnapshotV1**](AccountAPI.md#WalletGetAccountSnapshotV1) | **Get** /sapi/v1/accountSnapshot | Daily Account Snapshot (USER_DATA)
[**WalletGetAccountStatusV1**](AccountAPI.md#WalletGetAccountStatusV1) | **Get** /sapi/v1/account/status | Account Status (USER_DATA)



## WalletCreateAccountDisableFastWithdrawSwitchV1

> map[string]interface{} WalletCreateAccountDisableFastWithdrawSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Disable Fast Withdraw Switch (USER_DATA)

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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.WalletCreateAccountDisableFastWithdrawSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.WalletCreateAccountDisableFastWithdrawSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAccountDisableFastWithdrawSwitchV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.WalletCreateAccountDisableFastWithdrawSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAccountDisableFastWithdrawSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

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


## WalletCreateAccountEnableFastWithdrawSwitchV1

> map[string]interface{} WalletCreateAccountEnableFastWithdrawSwitchV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Fast Withdraw Switch (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.WalletCreateAccountEnableFastWithdrawSwitchV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.WalletCreateAccountEnableFastWithdrawSwitchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateAccountEnableFastWithdrawSwitchV1`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.WalletCreateAccountEnableFastWithdrawSwitchV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateAccountEnableFastWithdrawSwitchV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

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


## WalletGetAccountApiRestrictionsV1

> WalletGetAccountApiRestrictionsV1Resp WalletGetAccountApiRestrictionsV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get API Key Permission (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.WalletGetAccountApiRestrictionsV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.WalletGetAccountApiRestrictionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountApiRestrictionsV1`: WalletGetAccountApiRestrictionsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.WalletGetAccountApiRestrictionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountApiRestrictionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountApiRestrictionsV1Resp**](WalletGetAccountApiRestrictionsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountApiTradingStatusV1

> WalletGetAccountApiTradingStatusV1Resp WalletGetAccountApiTradingStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account API Trading Status (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.WalletGetAccountApiTradingStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.WalletGetAccountApiTradingStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountApiTradingStatusV1`: WalletGetAccountApiTradingStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.WalletGetAccountApiTradingStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountApiTradingStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountApiTradingStatusV1Resp**](WalletGetAccountApiTradingStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountInfoV1

> WalletGetAccountInfoV1Resp WalletGetAccountInfoV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account info (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.WalletGetAccountInfoV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.WalletGetAccountInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountInfoV1`: WalletGetAccountInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.WalletGetAccountInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountInfoV1Resp**](WalletGetAccountInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountSnapshotV1

> WalletGetAccountSnapshotV1Resp WalletGetAccountSnapshotV1(ctx).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Daily Account Snapshot (USER_DATA)



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
	type_ := "type__example" // string | &#34;SPOT&#34;, &#34;MARGIN&#34;, &#34;FUTURES&#34; (default to "")
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | min 7, max 30, default 7 (optional) (default to 7)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.WalletGetAccountSnapshotV1(context.Background()).Type_(type_).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.WalletGetAccountSnapshotV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountSnapshotV1`: WalletGetAccountSnapshotV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.WalletGetAccountSnapshotV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountSnapshotV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | &amp;#34;SPOT&amp;#34;, &amp;#34;MARGIN&amp;#34;, &amp;#34;FUTURES&amp;#34; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | min 7, max 30, default 7 | [default to 7]
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountSnapshotV1Resp**](WalletGetAccountSnapshotV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetAccountStatusV1

> WalletGetAccountStatusV1Resp WalletGetAccountStatusV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Account Status (USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.WalletGetAccountStatusV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.WalletGetAccountStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetAccountStatusV1`: WalletGetAccountStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.WalletGetAccountStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetAccountStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**WalletGetAccountStatusV1Resp**](WalletGetAccountStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

