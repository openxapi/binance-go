# \DualInvestmentAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDciProductAutoCompoundEditStatusV1**](DualInvestmentAPI.md#CreateDciProductAutoCompoundEditStatusV1) | **Post** /sapi/v1/dci/product/auto_compound/edit-status | Change Auto-Compound status(USER_DATA)
[**CreateDciProductSubscribeV1**](DualInvestmentAPI.md#CreateDciProductSubscribeV1) | **Post** /sapi/v1/dci/product/subscribe | Subscribe Dual Investment products(USER_DATA)
[**GetDciProductAccountsV1**](DualInvestmentAPI.md#GetDciProductAccountsV1) | **Get** /sapi/v1/dci/product/accounts | Check Dual Investment accounts(USER_DATA)
[**GetDciProductListV1**](DualInvestmentAPI.md#GetDciProductListV1) | **Get** /sapi/v1/dci/product/list | Get Dual Investment product list
[**GetDciProductPositionsV1**](DualInvestmentAPI.md#GetDciProductPositionsV1) | **Get** /sapi/v1/dci/product/positions | Get Dual Investment positions(USER_DATA)



## CreateDciProductAutoCompoundEditStatusV1

> CreateDciProductAutoCompoundEditStatusV1Resp CreateDciProductAutoCompoundEditStatusV1(ctx).PositionId(positionId).Timestamp(timestamp).AutoCompoundPlan(autoCompoundPlan).RecvWindow(recvWindow).Execute()

Change Auto-Compound status(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	positionId := "positionId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	autoCompoundPlan := "autoCompoundPlan_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DualInvestmentAPI.CreateDciProductAutoCompoundEditStatusV1(context.Background()).PositionId(positionId).Timestamp(timestamp).AutoCompoundPlan(autoCompoundPlan).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DualInvestmentAPI.CreateDciProductAutoCompoundEditStatusV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDciProductAutoCompoundEditStatusV1`: CreateDciProductAutoCompoundEditStatusV1Resp
	fmt.Fprintf(os.Stdout, "Response from `DualInvestmentAPI.CreateDciProductAutoCompoundEditStatusV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDciProductAutoCompoundEditStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **positionId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **autoCompoundPlan** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateDciProductAutoCompoundEditStatusV1Resp**](CreateDciProductAutoCompoundEditStatusV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDciProductSubscribeV1

> CreateDciProductSubscribeV1Resp CreateDciProductSubscribeV1(ctx).AutoCompoundPlan(autoCompoundPlan).DepositAmount(depositAmount).Id(id).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Subscribe Dual Investment products(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	autoCompoundPlan := "autoCompoundPlan_example" // string |  (default to "")
	depositAmount := "depositAmount_example" // string |  (default to "")
	id := "id_example" // string |  (default to "")
	orderId := "orderId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DualInvestmentAPI.CreateDciProductSubscribeV1(context.Background()).AutoCompoundPlan(autoCompoundPlan).DepositAmount(depositAmount).Id(id).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DualInvestmentAPI.CreateDciProductSubscribeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDciProductSubscribeV1`: CreateDciProductSubscribeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `DualInvestmentAPI.CreateDciProductSubscribeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDciProductSubscribeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCompoundPlan** | **string** |  | [default to &quot;&quot;]
 **depositAmount** | **string** |  | [default to &quot;&quot;]
 **id** | **string** |  | [default to &quot;&quot;]
 **orderId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateDciProductSubscribeV1Resp**](CreateDciProductSubscribeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDciProductAccountsV1

> GetDciProductAccountsV1Resp GetDciProductAccountsV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Check Dual Investment accounts(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DualInvestmentAPI.GetDciProductAccountsV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DualInvestmentAPI.GetDciProductAccountsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDciProductAccountsV1`: GetDciProductAccountsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `DualInvestmentAPI.GetDciProductAccountsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDciProductAccountsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetDciProductAccountsV1Resp**](GetDciProductAccountsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDciProductListV1

> GetDciProductListV1Resp GetDciProductListV1(ctx).OptionType(optionType).ExercisedCoin(exercisedCoin).InvestCoin(investCoin).Timestamp(timestamp).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()

Get Dual Investment product list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	optionType := "optionType_example" // string | Input CALL or PUT (default to "")
	exercisedCoin := "exercisedCoin_example" // string | Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType`:CALL,`exercisedCoin`:USDT,`investCoin`:BNB; if you subscribe to a low buy product (put option), you should input: `optionType`:PUT,`exercisedCoin`:BNB,`investCoin`:USDT (default to "")
	investCoin := "investCoin_example" // string | Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: `optionType`:CALL,`exercisedCoin`:USDT,`investCoin`:BNB; if you subscribe to a low buy product (put option), you should input: `optionType`:PUT,`exercisedCoin`:BNB,`investCoin`:USDT (default to "")
	timestamp := int64(789) // int64 | 
	pageSize := int64(789) // int64 | Default: 10, Maximum: 100 (optional) (default to 10)
	pageIndex := int32(56) // int32 | Default: 1 (optional) (default to 1)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DualInvestmentAPI.GetDciProductListV1(context.Background()).OptionType(optionType).ExercisedCoin(exercisedCoin).InvestCoin(investCoin).Timestamp(timestamp).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DualInvestmentAPI.GetDciProductListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDciProductListV1`: GetDciProductListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `DualInvestmentAPI.GetDciProductListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDciProductListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optionType** | **string** | Input CALL or PUT | [default to &quot;&quot;]
 **exercisedCoin** | **string** | Target exercised asset, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType&#x60;:CALL,&#x60;exercisedCoin&#x60;:USDT,&#x60;investCoin&#x60;:BNB; if you subscribe to a low buy product (put option), you should input: &#x60;optionType&#x60;:PUT,&#x60;exercisedCoin&#x60;:BNB,&#x60;investCoin&#x60;:USDT | [default to &quot;&quot;]
 **investCoin** | **string** | Asset used for subscribing, e.g.: if you subscribe to a high sell product (call option), you should input: &#x60;optionType&#x60;:CALL,&#x60;exercisedCoin&#x60;:USDT,&#x60;investCoin&#x60;:BNB; if you subscribe to a low buy product (put option), you should input: &#x60;optionType&#x60;:PUT,&#x60;exercisedCoin&#x60;:BNB,&#x60;investCoin&#x60;:USDT | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **pageSize** | **int64** | Default: 10, Maximum: 100 | [default to 10]
 **pageIndex** | **int32** | Default: 1 | [default to 1]
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetDciProductListV1Resp**](GetDciProductListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDciProductPositionsV1

> GetDciProductPositionsV1Resp GetDciProductPositionsV1(ctx).Timestamp(timestamp).Status(status).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()

Get Dual Investment positions(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/rest/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	status := "status_example" // string | `PENDING`:Products are purchasing, will give results later;`PURCHASE_SUCCESS`:purchase successfully;`SETTLED`: Products are finish settling;`PURCHASE_FAIL`:fail to purchase;`REFUNDING`:refund ongoing;`REFUND_SUCCESS`:refund to spot account successfully; `SETTLING`:Products are settling. If don&#39;t fill this field, will response all the position status. (optional) (default to "")
	pageSize := int64(789) // int64 | Default: 10, Max:100 (optional) (default to 10)
	pageIndex := int32(56) // int32 | Default:1 (optional)
	recvWindow := int64(789) // int64 | The value cannot be greater than 60000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DualInvestmentAPI.GetDciProductPositionsV1(context.Background()).Timestamp(timestamp).Status(status).PageSize(pageSize).PageIndex(pageIndex).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DualInvestmentAPI.GetDciProductPositionsV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDciProductPositionsV1`: GetDciProductPositionsV1Resp
	fmt.Fprintf(os.Stdout, "Response from `DualInvestmentAPI.GetDciProductPositionsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDciProductPositionsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **status** | **string** | &#x60;PENDING&#x60;:Products are purchasing, will give results later;&#x60;PURCHASE_SUCCESS&#x60;:purchase successfully;&#x60;SETTLED&#x60;: Products are finish settling;&#x60;PURCHASE_FAIL&#x60;:fail to purchase;&#x60;REFUNDING&#x60;:refund ongoing;&#x60;REFUND_SUCCESS&#x60;:refund to spot account successfully; &#x60;SETTLING&#x60;:Products are settling. If don&amp;#39;t fill this field, will response all the position status. | [default to &quot;&quot;]
 **pageSize** | **int64** | Default: 10, Max:100 | [default to 10]
 **pageIndex** | **int32** | Default:1 | 
 **recvWindow** | **int64** | The value cannot be greater than 60000 | 

### Return type

[**GetDciProductPositionsV1Resp**](GetDciProductPositionsV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

