# \VipLoanAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoanVipBorrowV1**](VipLoanAPI.md#CreateLoanVipBorrowV1) | **Post** /sapi/v1/loan/vip/borrow | VIP Loan Borrow(TRADE)
[**CreateLoanVipRenewV1**](VipLoanAPI.md#CreateLoanVipRenewV1) | **Post** /sapi/v1/loan/vip/renew | VIP Loan Renew(TRADE)
[**CreateLoanVipRepayV1**](VipLoanAPI.md#CreateLoanVipRepayV1) | **Post** /sapi/v1/loan/vip/repay | VIP Loan Repay(TRADE)
[**GetLoanVipAccruedInterestV1**](VipLoanAPI.md#GetLoanVipAccruedInterestV1) | **Get** /sapi/v1/loan/vip/accruedInterest | Get VIP Loan Accrued Interest(USER_DATA)
[**GetLoanVipCollateralAccountV1**](VipLoanAPI.md#GetLoanVipCollateralAccountV1) | **Get** /sapi/v1/loan/vip/collateral/account | Check VIP Loan Collateral Account (USER_DATA)
[**GetLoanVipCollateralDataV1**](VipLoanAPI.md#GetLoanVipCollateralDataV1) | **Get** /sapi/v1/loan/vip/collateral/data | Get Collateral Asset Data(USER_DATA)
[**GetLoanVipInterestRateHistoryV1**](VipLoanAPI.md#GetLoanVipInterestRateHistoryV1) | **Get** /sapi/v1/loan/vip/interestRateHistory | Get VIP Loan Interest Rate History (USER_DATA)
[**GetLoanVipLoanableDataV1**](VipLoanAPI.md#GetLoanVipLoanableDataV1) | **Get** /sapi/v1/loan/vip/loanable/data | Get Loanable Assets Data(USER_DATA)
[**GetLoanVipOngoingOrdersV1**](VipLoanAPI.md#GetLoanVipOngoingOrdersV1) | **Get** /sapi/v1/loan/vip/ongoing/orders | Get VIP Loan Ongoing Orders(USER_DATA)
[**GetLoanVipRepayHistoryV1**](VipLoanAPI.md#GetLoanVipRepayHistoryV1) | **Get** /sapi/v1/loan/vip/repay/history | Get VIP Loan Repayment History(USER_DATA)
[**GetLoanVipRequestDataV1**](VipLoanAPI.md#GetLoanVipRequestDataV1) | **Get** /sapi/v1/loan/vip/request/data | Query Application Status(USER_DATA)
[**GetLoanVipRequestInterestRateV1**](VipLoanAPI.md#GetLoanVipRequestInterestRateV1) | **Get** /sapi/v1/loan/vip/request/interestRate | Get Borrow Interest Rate(USER_DATA)



## CreateLoanVipBorrowV1

> CreateLoanVipBorrowV1Resp CreateLoanVipBorrowV1(ctx).Execute()

VIP Loan Borrow(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.CreateLoanVipBorrowV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.CreateLoanVipBorrowV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoanVipBorrowV1`: CreateLoanVipBorrowV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.CreateLoanVipBorrowV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoanVipBorrowV1Request struct via the builder pattern


### Return type

[**CreateLoanVipBorrowV1Resp**](CreateLoanVipBorrowV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoanVipRenewV1

> CreateLoanVipRenewV1Resp CreateLoanVipRenewV1(ctx).LoanTerm(loanTerm).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

VIP Loan Renew(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	loanTerm := int32(56) // int32 | 
	orderId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.CreateLoanVipRenewV1(context.Background()).LoanTerm(loanTerm).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.CreateLoanVipRenewV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoanVipRenewV1`: CreateLoanVipRenewV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.CreateLoanVipRenewV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoanVipRenewV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loanTerm** | **int32** |  | 
 **orderId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateLoanVipRenewV1Resp**](CreateLoanVipRenewV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoanVipRepayV1

> CreateLoanVipRepayV1Resp CreateLoanVipRepayV1(ctx).Amount(amount).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

VIP Loan Repay(TRADE)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	amount := "amount_example" // string |  (default to "")
	orderId := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.CreateLoanVipRepayV1(context.Background()).Amount(amount).OrderId(orderId).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.CreateLoanVipRepayV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoanVipRepayV1`: CreateLoanVipRepayV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.CreateLoanVipRepayV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoanVipRepayV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **string** |  | [default to &quot;&quot;]
 **orderId** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateLoanVipRepayV1Resp**](CreateLoanVipRepayV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipAccruedInterestV1

> GetLoanVipAccruedInterestV1Resp GetLoanVipAccruedInterestV1(ctx).RecvWindow(recvWindow).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()

Get VIP Loan Accrued Interest(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	recvWindow := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	loanCoin := "loanCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipAccruedInterestV1(context.Background()).RecvWindow(recvWindow).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipAccruedInterestV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipAccruedInterestV1`: GetLoanVipAccruedInterestV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipAccruedInterestV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipAccruedInterestV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recvWindow** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]

### Return type

[**GetLoanVipAccruedInterestV1Resp**](GetLoanVipAccruedInterestV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipCollateralAccountV1

> GetLoanVipCollateralAccountV1Resp GetLoanVipCollateralAccountV1(ctx).Timestamp(timestamp).OrderId(orderId).CollateralAccountId(collateralAccountId).RecvWindow(recvWindow).Execute()

Check VIP Loan Collateral Account (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	collateralAccountId := int64(789) // int64 |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipCollateralAccountV1(context.Background()).Timestamp(timestamp).OrderId(orderId).CollateralAccountId(collateralAccountId).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipCollateralAccountV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipCollateralAccountV1`: GetLoanVipCollateralAccountV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipCollateralAccountV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipCollateralAccountV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **collateralAccountId** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanVipCollateralAccountV1Resp**](GetLoanVipCollateralAccountV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipCollateralDataV1

> GetLoanVipCollateralDataV1Resp GetLoanVipCollateralDataV1(ctx).Timestamp(timestamp).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()

Get Collateral Asset Data(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipCollateralDataV1(context.Background()).Timestamp(timestamp).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipCollateralDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipCollateralDataV1`: GetLoanVipCollateralDataV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipCollateralDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipCollateralDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanVipCollateralDataV1Resp**](GetLoanVipCollateralDataV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipInterestRateHistoryV1

> GetLoanVipInterestRateHistoryV1Resp GetLoanVipInterestRateHistoryV1(ctx).Coin(coin).RecvWindow(recvWindow).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()

Get VIP Loan Interest Rate History (USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	coin := "coin_example" // string |  (default to "")
	recvWindow := int64(789) // int64 | 
	timestamp := int64(789) // int64 | 
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Check current querying page, start from 1. Default：1；Max：1000. (optional)
	limit := int64(789) // int64 | Default：10; Max：100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipInterestRateHistoryV1(context.Background()).Coin(coin).RecvWindow(recvWindow).Timestamp(timestamp).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipInterestRateHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipInterestRateHistoryV1`: GetLoanVipInterestRateHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipInterestRateHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipInterestRateHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coin** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **timestamp** | **int64** |  | 
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Check current querying page, start from 1. Default：1；Max：1000. | 
 **limit** | **int64** | Default：10; Max：100. | 

### Return type

[**GetLoanVipInterestRateHistoryV1Resp**](GetLoanVipInterestRateHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipLoanableDataV1

> GetLoanVipLoanableDataV1Resp GetLoanVipLoanableDataV1(ctx).Timestamp(timestamp).LoanCoin(loanCoin).VipLevel(vipLevel).RecvWindow(recvWindow).Execute()

Get Loanable Assets Data(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	loanCoin := "loanCoin_example" // string |  (optional) (default to "")
	vipLevel := int32(56) // int32 | default:user&#39;s vip level (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipLoanableDataV1(context.Background()).Timestamp(timestamp).LoanCoin(loanCoin).VipLevel(vipLevel).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipLoanableDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipLoanableDataV1`: GetLoanVipLoanableDataV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipLoanableDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipLoanableDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **vipLevel** | **int32** | default:user&amp;#39;s vip level | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanVipLoanableDataV1Resp**](GetLoanVipLoanableDataV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipOngoingOrdersV1

> GetLoanVipOngoingOrdersV1Resp GetLoanVipOngoingOrdersV1(ctx).Timestamp(timestamp).OrderId(orderId).CollateralAccountId(collateralAccountId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get VIP Loan Ongoing Orders(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	collateralAccountId := int64(789) // int64 |  (optional)
	loanCoin := "loanCoin_example" // string |  (optional) (default to "")
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	current := int64(789) // int64 | Currently querying page. Start from 1, Default:1, Max: 1000. (optional)
	limit := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipOngoingOrdersV1(context.Background()).Timestamp(timestamp).OrderId(orderId).CollateralAccountId(collateralAccountId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipOngoingOrdersV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipOngoingOrdersV1`: GetLoanVipOngoingOrdersV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipOngoingOrdersV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipOngoingOrdersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **collateralAccountId** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Currently querying page. Start from 1, Default:1, Max: 1000. | 
 **limit** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanVipOngoingOrdersV1Resp**](GetLoanVipOngoingOrdersV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipRepayHistoryV1

> GetLoanVipRepayHistoryV1Resp GetLoanVipRepayHistoryV1(ctx).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get VIP Loan Repayment History(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	orderId := int64(789) // int64 |  (optional)
	loanCoin := "loanCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Currently querying page. Start from 1, Default:1, Max: 1000 (optional)
	limit := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipRepayHistoryV1(context.Background()).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipRepayHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipRepayHistoryV1`: GetLoanVipRepayHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipRepayHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipRepayHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1, Default:1, Max: 1000 | 
 **limit** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanVipRepayHistoryV1Resp**](GetLoanVipRepayHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipRequestDataV1

> GetLoanVipRequestDataV1Resp GetLoanVipRequestDataV1(ctx).Timestamp(timestamp).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Query Application Status(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	timestamp := int64(789) // int64 | 
	current := int64(789) // int64 | Currently querying page. Start from 1, Default:1, Max: 1000 (optional)
	limit := int64(789) // int64 | Default: 10, Max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipRequestDataV1(context.Background()).Timestamp(timestamp).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipRequestDataV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipRequestDataV1`: GetLoanVipRequestDataV1Resp
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipRequestDataV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipRequestDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **current** | **int64** | Currently querying page. Start from 1, Default:1, Max: 1000 | 
 **limit** | **int64** | Default: 10, Max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanVipRequestDataV1Resp**](GetLoanVipRequestDataV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanVipRequestInterestRateV1

> []GetLoanVipRequestInterestRateV1RespItem GetLoanVipRequestInterestRateV1(ctx).LoanCoin(loanCoin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Borrow Interest Rate(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/spot"
)

func main() {
	loanCoin := "loanCoin_example" // string | Max 10 assets, Multiple split by &#34;,&#34; (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VipLoanAPI.GetLoanVipRequestInterestRateV1(context.Background()).LoanCoin(loanCoin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VipLoanAPI.GetLoanVipRequestInterestRateV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanVipRequestInterestRateV1`: []GetLoanVipRequestInterestRateV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `VipLoanAPI.GetLoanVipRequestInterestRateV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanVipRequestInterestRateV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** | Max 10 assets, Multiple split by &amp;#34;,&amp;#34; | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetLoanVipRequestInterestRateV1RespItem**](GetLoanVipRequestInterestRateV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

