# \AccountManagementAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreateSubAccountBlvtEnableV1**](AccountManagementAPI.md#SubaccountCreateSubAccountBlvtEnableV1) | **Post** /sapi/v1/sub-account/blvt/enable | Enable Leverage Token for Sub-account(For Master Account)
[**SubaccountCreateSubAccountEoptionsEnableV1**](AccountManagementAPI.md#SubaccountCreateSubAccountEoptionsEnableV1) | **Post** /sapi/v1/sub-account/eoptions/enable | Enable Options for Sub-account(For Master Account)(USER_DATA)
[**SubaccountCreateSubAccountFuturesEnableV1**](AccountManagementAPI.md#SubaccountCreateSubAccountFuturesEnableV1) | **Post** /sapi/v1/sub-account/futures/enable | Enable Futures for Sub-account(For Master Account)
[**SubaccountCreateSubAccountMarginEnableV1**](AccountManagementAPI.md#SubaccountCreateSubAccountMarginEnableV1) | **Post** /sapi/v1/sub-account/margin/enable | Enable Margin for Sub-account(For Master Account)
[**SubaccountCreateSubAccountVirtualSubAccountV1**](AccountManagementAPI.md#SubaccountCreateSubAccountVirtualSubAccountV1) | **Post** /sapi/v1/sub-account/virtualSubAccount | Create a Virtual Sub-account(For Master Account)
[**SubaccountGetSubAccountFuturesPositionRiskV1**](AccountManagementAPI.md#SubaccountGetSubAccountFuturesPositionRiskV1) | **Get** /sapi/v1/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account(For Master Account)
[**SubaccountGetSubAccountFuturesPositionRiskV2**](AccountManagementAPI.md#SubaccountGetSubAccountFuturesPositionRiskV2) | **Get** /sapi/v2/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account V2(For Master Account)
[**SubaccountGetSubAccountListV1**](AccountManagementAPI.md#SubaccountGetSubAccountListV1) | **Get** /sapi/v1/sub-account/list | Query Sub-account List(For Master Account)
[**SubaccountGetSubAccountStatusV1**](AccountManagementAPI.md#SubaccountGetSubAccountStatusV1) | **Get** /sapi/v1/sub-account/status | Get Sub-account&#39;s Status on Margin Or Futures(For Master Account)
[**SubaccountGetSubAccountTransactionStatisticsV1**](AccountManagementAPI.md#SubaccountGetSubAccountTransactionStatisticsV1) | **Get** /sapi/v1/sub-account/transaction-statistics | Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)



## SubaccountCreateSubAccountBlvtEnableV1

> SubaccountCreateSubAccountBlvtEnableV1Resp SubaccountCreateSubAccountBlvtEnableV1(ctx).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Leverage Token for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string |  (default to "")
	enableBlvt := true // bool | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountCreateSubAccountBlvtEnableV1(context.Background()).Email(email).EnableBlvt(enableBlvt).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountCreateSubAccountBlvtEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountBlvtEnableV1`: SubaccountCreateSubAccountBlvtEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountCreateSubAccountBlvtEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountBlvtEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **enableBlvt** | **bool** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountBlvtEnableV1Resp**](SubaccountCreateSubAccountBlvtEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountEoptionsEnableV1

> SubaccountCreateSubAccountEoptionsEnableV1Resp SubaccountCreateSubAccountEoptionsEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Options for Sub-account(For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountCreateSubAccountEoptionsEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountCreateSubAccountEoptionsEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountEoptionsEnableV1`: SubaccountCreateSubAccountEoptionsEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountCreateSubAccountEoptionsEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountEoptionsEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountEoptionsEnableV1Resp**](SubaccountCreateSubAccountEoptionsEnableV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountFuturesEnableV1

> SubaccountCreateSubAccountFuturesEnableV1Resp SubaccountCreateSubAccountFuturesEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Futures for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountCreateSubAccountFuturesEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountCreateSubAccountFuturesEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountFuturesEnableV1`: SubaccountCreateSubAccountFuturesEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountCreateSubAccountFuturesEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountFuturesEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountFuturesEnableV1Resp**](SubaccountCreateSubAccountFuturesEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountMarginEnableV1

> SubaccountCreateSubAccountMarginEnableV1Resp SubaccountCreateSubAccountMarginEnableV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Margin for Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountCreateSubAccountMarginEnableV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountCreateSubAccountMarginEnableV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountMarginEnableV1`: SubaccountCreateSubAccountMarginEnableV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountCreateSubAccountMarginEnableV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountMarginEnableV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountMarginEnableV1Resp**](SubaccountCreateSubAccountMarginEnableV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountCreateSubAccountVirtualSubAccountV1

> SubaccountCreateSubAccountVirtualSubAccountV1Resp SubaccountCreateSubAccountVirtualSubAccountV1(ctx).SubAccountString(subAccountString).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Create a Virtual Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	subAccountString := "subAccountString_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountCreateSubAccountVirtualSubAccountV1(context.Background()).SubAccountString(subAccountString).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountCreateSubAccountVirtualSubAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountVirtualSubAccountV1`: SubaccountCreateSubAccountVirtualSubAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountCreateSubAccountVirtualSubAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountVirtualSubAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAccountString** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountVirtualSubAccountV1Resp**](SubaccountCreateSubAccountVirtualSubAccountV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesPositionRiskV1

> []SubaccountGetSubAccountFuturesPositionRiskV1RespItem SubaccountGetSubAccountFuturesPositionRiskV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Position-Risk of Sub-account(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account#email-address\">Sub-account email</a> (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountGetSubAccountFuturesPositionRiskV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountGetSubAccountFuturesPositionRiskV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesPositionRiskV1`: []SubaccountGetSubAccountFuturesPositionRiskV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountGetSubAccountFuturesPositionRiskV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesPositionRiskV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]SubaccountGetSubAccountFuturesPositionRiskV1RespItem**](SubaccountGetSubAccountFuturesPositionRiskV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesPositionRiskV2

> SubaccountGetSubAccountFuturesPositionRiskV2Resp SubaccountGetSubAccountFuturesPositionRiskV2(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Futures Position-Risk of Sub-account V2(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account-V2#email-address\">Sub-account email</a> (default to "")
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountGetSubAccountFuturesPositionRiskV2(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountGetSubAccountFuturesPositionRiskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesPositionRiskV2`: SubaccountGetSubAccountFuturesPositionRiskV2Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountGetSubAccountFuturesPositionRiskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesPositionRiskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Futures-Position-Risk-of-Sub-account-V2#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesPositionRiskV2Resp**](SubaccountGetSubAccountFuturesPositionRiskV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountListV1

> SubaccountGetSubAccountListV1Resp SubaccountGetSubAccountListV1(ctx).Timestamp(timestamp).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Query Sub-account List(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Query-Sub-account-List#email-address\">Sub-account email</a> (optional) (default to "")
	isFreeze := "isFreeze_example" // string | true or false (optional) (default to "")
	page := int32(56) // int32 | Default value: 1 (optional)
	limit := int32(56) // int32 | Default value: 1, Max value: 200 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountGetSubAccountListV1(context.Background()).Timestamp(timestamp).Email(email).IsFreeze(isFreeze).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountGetSubAccountListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountListV1`: SubaccountGetSubAccountListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountGetSubAccountListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Query-Sub-account-List#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **isFreeze** | **string** | true or false | [default to &quot;&quot;]
 **page** | **int32** | Default value: 1 | 
 **limit** | **int32** | Default value: 1, Max value: 200 | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountListV1Resp**](SubaccountGetSubAccountListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountStatusV1

> []SubaccountGetSubAccountStatusV1RespItem SubaccountGetSubAccountStatusV1(ctx).Timestamp(timestamp).Email(email).RecvWindow(recvWindow).Execute()

Get Sub-account's Status on Margin Or Futures(For Master Account)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	timestamp := int64(789) // int64 | 
	email := "email_example" // string | <a href=\"/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures#email-address\">Sub-account email</a> (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountGetSubAccountStatusV1(context.Background()).Timestamp(timestamp).Email(email).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountGetSubAccountStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountStatusV1`: []SubaccountGetSubAccountStatusV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountGetSubAccountStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/account-management/Get-Sub-accounts-Status-on-Margin-Or-Futures#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]SubaccountGetSubAccountStatusV1RespItem**](SubaccountGetSubAccountStatusV1RespItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountTransactionStatisticsV1

> SubaccountGetSubAccountTransactionStatisticsV1Resp SubaccountGetSubAccountTransactionStatisticsV1(ctx).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Sub-account Transaction Statistics(For Master Account)(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func main() {
	email := "email_example" // string | Sub user email (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountManagementAPI.SubaccountGetSubAccountTransactionStatisticsV1(context.Background()).Email(email).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.SubaccountGetSubAccountTransactionStatisticsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountTransactionStatisticsV1`: SubaccountGetSubAccountTransactionStatisticsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.SubaccountGetSubAccountTransactionStatisticsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountTransactionStatisticsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Sub user email | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountTransactionStatisticsV1Resp**](SubaccountGetSubAccountTransactionStatisticsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

