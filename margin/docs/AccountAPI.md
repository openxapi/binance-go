# \AccountAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCreateMarginIsolatedAccountV1**](AccountAPI.md#MarginCreateMarginIsolatedAccountV1) | **Post** /sapi/v1/margin/isolated/account | Enable Isolated Margin Account (TRADE)
[**MarginCreateMarginMaxLeverageV1**](AccountAPI.md#MarginCreateMarginMaxLeverageV1) | **Post** /sapi/v1/margin/max-leverage | Adjust cross margin max leverage (USER_DATA)
[**MarginDeleteMarginIsolatedAccountV1**](AccountAPI.md#MarginDeleteMarginIsolatedAccountV1) | **Delete** /sapi/v1/margin/isolated/account | Disable Isolated Margin Account (TRADE)
[**MarginGetBnbBurnV1**](AccountAPI.md#MarginGetBnbBurnV1) | **Get** /sapi/v1/bnbBurn | Get BNB Burn Status (USER_DATA)
[**MarginGetMarginAccountV1**](AccountAPI.md#MarginGetMarginAccountV1) | **Get** /sapi/v1/margin/account | Query Cross Margin Account Details (USER_DATA)
[**MarginGetMarginCapitalFlowV1**](AccountAPI.md#MarginGetMarginCapitalFlowV1) | **Get** /sapi/v1/margin/capital-flow | Query Cross Isolated Margin Capital Flow (USER_DATA)
[**MarginGetMarginCrossMarginDataV1**](AccountAPI.md#MarginGetMarginCrossMarginDataV1) | **Get** /sapi/v1/margin/crossMarginData | Query Cross Margin Fee Data (USER_DATA)
[**MarginGetMarginIsolatedAccountLimitV1**](AccountAPI.md#MarginGetMarginIsolatedAccountLimitV1) | **Get** /sapi/v1/margin/isolated/accountLimit | Query Enabled Isolated Margin Account Limit (USER_DATA)
[**MarginGetMarginIsolatedAccountV1**](AccountAPI.md#MarginGetMarginIsolatedAccountV1) | **Get** /sapi/v1/margin/isolated/account | Query Isolated Margin Account Info (USER_DATA)
[**MarginGetMarginIsolatedMarginDataV1**](AccountAPI.md#MarginGetMarginIsolatedMarginDataV1) | **Get** /sapi/v1/margin/isolatedMarginData | Query Isolated Margin Fee Data (USER_DATA)
[**MarginGetMarginTradeCoeffV1**](AccountAPI.md#MarginGetMarginTradeCoeffV1) | **Get** /sapi/v1/margin/tradeCoeff | Get Summary of Margin account (USER_DATA)



## MarginCreateMarginIsolatedAccountV1

> MarginCreateMarginIsolatedAccountV1Resp MarginCreateMarginIsolatedAccountV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Enable Isolated Margin Account (TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginCreateMarginIsolatedAccountV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginCreateMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginIsolatedAccountV1`: MarginCreateMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginCreateMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginCreateMarginIsolatedAccountV1Resp**](MarginCreateMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginCreateMarginMaxLeverageV1

> MarginCreateMarginMaxLeverageV1Resp MarginCreateMarginMaxLeverageV1(ctx).MaxLeverage(maxLeverage).Execute()

Adjust cross margin max leverage (USER_DATA)



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
	maxLeverage := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginCreateMarginMaxLeverageV1(context.Background()).MaxLeverage(maxLeverage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginCreateMarginMaxLeverageV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginCreateMarginMaxLeverageV1`: MarginCreateMarginMaxLeverageV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginCreateMarginMaxLeverageV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginCreateMarginMaxLeverageV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxLeverage** | **int32** |  | 

### Return type

[**MarginCreateMarginMaxLeverageV1Resp**](MarginCreateMarginMaxLeverageV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginDeleteMarginIsolatedAccountV1

> MarginDeleteMarginIsolatedAccountV1Resp MarginDeleteMarginIsolatedAccountV1(ctx).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Disable Isolated Margin Account (TRADE)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginDeleteMarginIsolatedAccountV1(context.Background()).Symbol(symbol).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginDeleteMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginDeleteMarginIsolatedAccountV1`: MarginDeleteMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginDeleteMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginDeleteMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginDeleteMarginIsolatedAccountV1Resp**](MarginDeleteMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetBnbBurnV1

> MarginGetBnbBurnV1Resp MarginGetBnbBurnV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get BNB Burn Status (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetBnbBurnV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetBnbBurnV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetBnbBurnV1`: MarginGetBnbBurnV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetBnbBurnV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetBnbBurnV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginGetBnbBurnV1Resp**](MarginGetBnbBurnV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginAccountV1

> MarginGetMarginAccountV1Resp MarginGetMarginAccountV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Cross Margin Account Details (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetMarginAccountV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetMarginAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginAccountV1`: MarginGetMarginAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetMarginAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**MarginGetMarginAccountV1Resp**](MarginGetMarginAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginCapitalFlowV1

> []MarginGetMarginCapitalFlowV1RespItem MarginGetMarginCapitalFlowV1(ctx).Timestamp(timestamp).Asset(asset).Symbol(symbol).Type_(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()

Query Cross Isolated Margin Capital Flow (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	asset := "asset_example" // string |  (optional) (default to "")
	symbol := "symbol_example" // string | 查询逐仓数据时必填 (optional) (default to "")
	type_ := "type__example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 | 只支持查询最近90天的数据 (optional)
	endTime := int64(789) // int64 |  (optional)
	fromId := int64(789) // int64 | 如设置fromId, 将返回id &gt; fromId的数据。否则将返回最新数据 (optional)
	limit := int64(789) // int64 | 每次返回的数据条数限制。默认 500; 最大 1000. (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetMarginCapitalFlowV1(context.Background()).Timestamp(timestamp).Asset(asset).Symbol(symbol).Type_(type_).StartTime(startTime).EndTime(endTime).FromId(fromId).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetMarginCapitalFlowV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginCapitalFlowV1`: []MarginGetMarginCapitalFlowV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetMarginCapitalFlowV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginCapitalFlowV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **symbol** | **string** | 查询逐仓数据时必填 | [default to &quot;&quot;]
 **type_** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** | 只支持查询最近90天的数据 | 
 **endTime** | **int64** |  | 
 **fromId** | **int64** | 如设置fromId, 将返回id &amp;gt; fromId的数据。否则将返回最新数据 | 
 **limit** | **int64** | 每次返回的数据条数限制。默认 500; 最大 1000. | 
 **recvWindow** | **int64** | The value cannot be greater than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginCapitalFlowV1RespItem**](MarginGetMarginCapitalFlowV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginCrossMarginDataV1

> []MarginGetMarginCrossMarginDataV1RespItem MarginGetMarginCrossMarginDataV1(ctx).Timestamp(timestamp).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()

Query Cross Margin Fee Data (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	vipLevel := int32(56) // int32 | User&#39;s current specific margin data will be returned if vipLevel is omitted (optional)
	coin := "coin_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetMarginCrossMarginDataV1(context.Background()).Timestamp(timestamp).VipLevel(vipLevel).Coin(coin).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetMarginCrossMarginDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginCrossMarginDataV1`: []MarginGetMarginCrossMarginDataV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetMarginCrossMarginDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginCrossMarginDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | User&amp;#39;s current specific margin data will be returned if vipLevel is omitted | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginCrossMarginDataV1RespItem**](MarginGetMarginCrossMarginDataV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginIsolatedAccountLimitV1

> MarginGetMarginIsolatedAccountLimitV1Resp MarginGetMarginIsolatedAccountLimitV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Query Enabled Isolated Margin Account Limit (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetMarginIsolatedAccountLimitV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetMarginIsolatedAccountLimitV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedAccountLimitV1`: MarginGetMarginIsolatedAccountLimitV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetMarginIsolatedAccountLimitV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedAccountLimitV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginGetMarginIsolatedAccountLimitV1Resp**](MarginGetMarginIsolatedAccountLimitV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginIsolatedAccountV1

> MarginGetMarginIsolatedAccountV1Resp MarginGetMarginIsolatedAccountV1(ctx).Timestamp(timestamp).Symbols(symbols).RecvWindow(recvWindow).Execute()

Query Isolated Margin Account Info (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	symbols := "symbols_example" // string | Max 5 symbols can be sent; separated by &#34;,&#34;. e.g. &#34;BTCUSDT,BNBUSDT,ADAUSDT&#34; (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetMarginIsolatedAccountV1(context.Background()).Timestamp(timestamp).Symbols(symbols).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetMarginIsolatedAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedAccountV1`: MarginGetMarginIsolatedAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetMarginIsolatedAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbols** | **string** | Max 5 symbols can be sent; separated by &amp;#34;,&amp;#34;. e.g. &amp;#34;BTCUSDT,BNBUSDT,ADAUSDT&amp;#34; | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than 60000 | 

### Return type

[**MarginGetMarginIsolatedAccountV1Resp**](MarginGetMarginIsolatedAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginIsolatedMarginDataV1

> []MarginGetMarginIsolatedMarginDataV1RespItem MarginGetMarginIsolatedMarginDataV1(ctx).Timestamp(timestamp).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()

Query Isolated Margin Fee Data (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	vipLevel := int32(56) // int32 | User&#39;s current specific margin data will be returned if vipLevel is omitted (optional)
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 | No more than `60000` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetMarginIsolatedMarginDataV1(context.Background()).Timestamp(timestamp).VipLevel(vipLevel).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetMarginIsolatedMarginDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginIsolatedMarginDataV1`: []MarginGetMarginIsolatedMarginDataV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetMarginIsolatedMarginDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginIsolatedMarginDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **vipLevel** | **int32** | User&amp;#39;s current specific margin data will be returned if vipLevel is omitted | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** | No more than &#x60;60000&#x60; | 

### Return type

[**[]MarginGetMarginIsolatedMarginDataV1RespItem**](MarginGetMarginIsolatedMarginDataV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarginGetMarginTradeCoeffV1

> MarginGetMarginTradeCoeffV1Resp MarginGetMarginTradeCoeffV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Summary of Margin account (USER_DATA)



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
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.MarginGetMarginTradeCoeffV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.MarginGetMarginTradeCoeffV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarginGetMarginTradeCoeffV1`: MarginGetMarginTradeCoeffV1Resp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.MarginGetMarginTradeCoeffV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarginGetMarginTradeCoeffV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**MarginGetMarginTradeCoeffV1Resp**](MarginGetMarginTradeCoeffV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

