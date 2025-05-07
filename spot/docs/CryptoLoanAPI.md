# \CryptoLoanAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoanFlexibleAdjustLtvV2**](CryptoLoanAPI.md#CreateLoanFlexibleAdjustLtvV2) | **Post** /sapi/v2/loan/flexible/adjust/ltv | Flexible Loan Adjust LTV(TRADE)
[**CreateLoanFlexibleBorrowV2**](CryptoLoanAPI.md#CreateLoanFlexibleBorrowV2) | **Post** /sapi/v2/loan/flexible/borrow | Flexible Loan Borrow(TRADE)
[**CreateLoanFlexibleRepayCollateralV2**](CryptoLoanAPI.md#CreateLoanFlexibleRepayCollateralV2) | **Post** /sapi/v2/loan/flexible/repay/collateral | Flexible Loan Collateral Repayment (TRADE)
[**CreateLoanFlexibleRepayV2**](CryptoLoanAPI.md#CreateLoanFlexibleRepayV2) | **Post** /sapi/v2/loan/flexible/repay | Flexible Loan Repay(TRADE)
[**GetLoanBorrowHistoryV1**](CryptoLoanAPI.md#GetLoanBorrowHistoryV1) | **Get** /sapi/v1/loan/borrow/history | Get Loan Borrow History(USER_DATA)
[**GetLoanFlexibleBorrowHistoryV2**](CryptoLoanAPI.md#GetLoanFlexibleBorrowHistoryV2) | **Get** /sapi/v2/loan/flexible/borrow/history | Get Flexible Loan Borrow History(USER_DATA)
[**GetLoanFlexibleCollateralDataV2**](CryptoLoanAPI.md#GetLoanFlexibleCollateralDataV2) | **Get** /sapi/v2/loan/flexible/collateral/data | Get Flexible Loan Collateral Assets Data(USER_DATA)
[**GetLoanFlexibleLiquidationHistoryV2**](CryptoLoanAPI.md#GetLoanFlexibleLiquidationHistoryV2) | **Get** /sapi/v2/loan/flexible/liquidation/history | Get Flexible Loan Liquidation History (USER_DATA)
[**GetLoanFlexibleLoanableDataV2**](CryptoLoanAPI.md#GetLoanFlexibleLoanableDataV2) | **Get** /sapi/v2/loan/flexible/loanable/data | Get Flexible Loan Assets Data(USER_DATA)
[**GetLoanFlexibleLtvAdjustmentHistoryV2**](CryptoLoanAPI.md#GetLoanFlexibleLtvAdjustmentHistoryV2) | **Get** /sapi/v2/loan/flexible/ltv/adjustment/history | Get Flexible Loan LTV Adjustment History(USER_DATA)
[**GetLoanFlexibleOngoingOrdersV2**](CryptoLoanAPI.md#GetLoanFlexibleOngoingOrdersV2) | **Get** /sapi/v2/loan/flexible/ongoing/orders | Get Flexible Loan Ongoing Orders(USER_DATA)
[**GetLoanFlexibleRepayHistoryV2**](CryptoLoanAPI.md#GetLoanFlexibleRepayHistoryV2) | **Get** /sapi/v2/loan/flexible/repay/history | Get Flexible Loan Repayment History(USER_DATA)
[**GetLoanFlexibleRepayRateV2**](CryptoLoanAPI.md#GetLoanFlexibleRepayRateV2) | **Get** /sapi/v2/loan/flexible/repay/rate | Check Collateral Repay Rate (USER_DATA)
[**GetLoanIncomeV1**](CryptoLoanAPI.md#GetLoanIncomeV1) | **Get** /sapi/v1/loan/income | Get Crypto Loans Income History(USER_DATA)
[**GetLoanLtvAdjustmentHistoryV1**](CryptoLoanAPI.md#GetLoanLtvAdjustmentHistoryV1) | **Get** /sapi/v1/loan/ltv/adjustment/history | Get Loan LTV Adjustment History(USER_DATA)
[**GetLoanRepayHistoryV1**](CryptoLoanAPI.md#GetLoanRepayHistoryV1) | **Get** /sapi/v1/loan/repay/history | Get Loan Repayment History(USER_DATA)



## CreateLoanFlexibleAdjustLtvV2

> CreateLoanFlexibleAdjustLtvV2Resp CreateLoanFlexibleAdjustLtvV2(ctx).AdjustmentAmount(adjustmentAmount).CollateralCoin(collateralCoin).Direction(direction).LoanCoin(loanCoin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Flexible Loan Adjust LTV(TRADE)



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
	adjustmentAmount := "adjustmentAmount_example" // string |  (default to "")
	collateralCoin := "collateralCoin_example" // string |  (default to "")
	direction := "direction_example" // string |  (default to "")
	loanCoin := "loanCoin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.CreateLoanFlexibleAdjustLtvV2(context.Background()).AdjustmentAmount(adjustmentAmount).CollateralCoin(collateralCoin).Direction(direction).LoanCoin(loanCoin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.CreateLoanFlexibleAdjustLtvV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoanFlexibleAdjustLtvV2`: CreateLoanFlexibleAdjustLtvV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.CreateLoanFlexibleAdjustLtvV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoanFlexibleAdjustLtvV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adjustmentAmount** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **direction** | **string** |  | [default to &quot;&quot;]
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateLoanFlexibleAdjustLtvV2Resp**](CreateLoanFlexibleAdjustLtvV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoanFlexibleBorrowV2

> CreateLoanFlexibleBorrowV2Resp CreateLoanFlexibleBorrowV2(ctx).CollateralCoin(collateralCoin).LoanCoin(loanCoin).Timestamp(timestamp).CollateralAmount(collateralAmount).LoanAmount(loanAmount).RecvWindow(recvWindow).Execute()

Flexible Loan Borrow(TRADE)



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
	collateralCoin := "collateralCoin_example" // string |  (default to "")
	loanCoin := "loanCoin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	collateralAmount := "collateralAmount_example" // string |  (optional) (default to "")
	loanAmount := "loanAmount_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.CreateLoanFlexibleBorrowV2(context.Background()).CollateralCoin(collateralCoin).LoanCoin(loanCoin).Timestamp(timestamp).CollateralAmount(collateralAmount).LoanAmount(loanAmount).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.CreateLoanFlexibleBorrowV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoanFlexibleBorrowV2`: CreateLoanFlexibleBorrowV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.CreateLoanFlexibleBorrowV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoanFlexibleBorrowV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **collateralAmount** | **string** |  | [default to &quot;&quot;]
 **loanAmount** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateLoanFlexibleBorrowV2Resp**](CreateLoanFlexibleBorrowV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoanFlexibleRepayCollateralV2

> CreateLoanFlexibleRepayCollateralV2Resp CreateLoanFlexibleRepayCollateralV2(ctx).Execute()

Flexible Loan Collateral Repayment (TRADE)



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
	resp, r, err := apiClient.CryptoLoanAPI.CreateLoanFlexibleRepayCollateralV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.CreateLoanFlexibleRepayCollateralV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoanFlexibleRepayCollateralV2`: CreateLoanFlexibleRepayCollateralV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.CreateLoanFlexibleRepayCollateralV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoanFlexibleRepayCollateralV2Request struct via the builder pattern


### Return type

[**CreateLoanFlexibleRepayCollateralV2Resp**](CreateLoanFlexibleRepayCollateralV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoanFlexibleRepayV2

> CreateLoanFlexibleRepayV2Resp CreateLoanFlexibleRepayV2(ctx).CollateralCoin(collateralCoin).LoanCoin(loanCoin).RepayAmount(repayAmount).Timestamp(timestamp).CollateralReturn(collateralReturn).FullRepayment(fullRepayment).RecvWindow(recvWindow).Execute()

Flexible Loan Repay(TRADE)



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
	collateralCoin := "collateralCoin_example" // string |  (default to "")
	loanCoin := "loanCoin_example" // string |  (default to "")
	repayAmount := "repayAmount_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	collateralReturn := true // bool |  (optional)
	fullRepayment := true // bool |  (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.CreateLoanFlexibleRepayV2(context.Background()).CollateralCoin(collateralCoin).LoanCoin(loanCoin).RepayAmount(repayAmount).Timestamp(timestamp).CollateralReturn(collateralReturn).FullRepayment(fullRepayment).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.CreateLoanFlexibleRepayV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoanFlexibleRepayV2`: CreateLoanFlexibleRepayV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.CreateLoanFlexibleRepayV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoanFlexibleRepayV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **repayAmount** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **collateralReturn** | **bool** |  | 
 **fullRepayment** | **bool** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateLoanFlexibleRepayV2Resp**](CreateLoanFlexibleRepayV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanBorrowHistoryV1

> GetLoanBorrowHistoryV1Resp GetLoanBorrowHistoryV1(ctx).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan Borrow History(USER_DATA)



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
	orderId := int64(789) // int64 | orderId in `POST /sapi/v1/loan/borrow` (optional)
	loanCoin := "loanCoin_example" // string |  (optional) (default to "")
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000. (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100. (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanBorrowHistoryV1(context.Background()).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanBorrowHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanBorrowHistoryV1`: GetLoanBorrowHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanBorrowHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanBorrowHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderId** | **int64** | orderId in &#x60;POST /sapi/v1/loan/borrow&#x60; | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000. | [default to 1]
 **limit** | **int64** | Default: 10; max: 100. | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanBorrowHistoryV1Resp**](GetLoanBorrowHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleBorrowHistoryV2

> GetLoanFlexibleBorrowHistoryV2Resp GetLoanFlexibleBorrowHistoryV2(ctx).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Borrow History(USER_DATA)



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
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleBorrowHistoryV2(context.Background()).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleBorrowHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleBorrowHistoryV2`: GetLoanFlexibleBorrowHistoryV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleBorrowHistoryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleBorrowHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleBorrowHistoryV2Resp**](GetLoanFlexibleBorrowHistoryV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleCollateralDataV2

> GetLoanFlexibleCollateralDataV2Resp GetLoanFlexibleCollateralDataV2(ctx).Timestamp(timestamp).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()

Get Flexible Loan Collateral Assets Data(USER_DATA)



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
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleCollateralDataV2(context.Background()).Timestamp(timestamp).CollateralCoin(collateralCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleCollateralDataV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleCollateralDataV2`: GetLoanFlexibleCollateralDataV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleCollateralDataV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleCollateralDataV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleCollateralDataV2Resp**](GetLoanFlexibleCollateralDataV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleLiquidationHistoryV2

> GetLoanFlexibleLiquidationHistoryV2Resp GetLoanFlexibleLiquidationHistoryV2(ctx).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Liquidation History (USER_DATA)

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
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleLiquidationHistoryV2(context.Background()).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleLiquidationHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleLiquidationHistoryV2`: GetLoanFlexibleLiquidationHistoryV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleLiquidationHistoryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleLiquidationHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleLiquidationHistoryV2Resp**](GetLoanFlexibleLiquidationHistoryV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleLoanableDataV2

> GetLoanFlexibleLoanableDataV2Resp GetLoanFlexibleLoanableDataV2(ctx).Timestamp(timestamp).LoanCoin(loanCoin).RecvWindow(recvWindow).Execute()

Get Flexible Loan Assets Data(USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleLoanableDataV2(context.Background()).Timestamp(timestamp).LoanCoin(loanCoin).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleLoanableDataV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleLoanableDataV2`: GetLoanFlexibleLoanableDataV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleLoanableDataV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleLoanableDataV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleLoanableDataV2Resp**](GetLoanFlexibleLoanableDataV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleLtvAdjustmentHistoryV2

> GetLoanFlexibleLtvAdjustmentHistoryV2Resp GetLoanFlexibleLtvAdjustmentHistoryV2(ctx).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan LTV Adjustment History(USER_DATA)



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
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleLtvAdjustmentHistoryV2(context.Background()).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleLtvAdjustmentHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleLtvAdjustmentHistoryV2`: GetLoanFlexibleLtvAdjustmentHistoryV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleLtvAdjustmentHistoryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleLtvAdjustmentHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleLtvAdjustmentHistoryV2Resp**](GetLoanFlexibleLtvAdjustmentHistoryV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleOngoingOrdersV2

> GetLoanFlexibleOngoingOrdersV2Resp GetLoanFlexibleOngoingOrdersV2(ctx).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Ongoing Orders(USER_DATA)



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
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleOngoingOrdersV2(context.Background()).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleOngoingOrdersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleOngoingOrdersV2`: GetLoanFlexibleOngoingOrdersV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleOngoingOrdersV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleOngoingOrdersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleOngoingOrdersV2Resp**](GetLoanFlexibleOngoingOrdersV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleRepayHistoryV2

> GetLoanFlexibleRepayHistoryV2Resp GetLoanFlexibleRepayHistoryV2(ctx).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Flexible Loan Repayment History(USER_DATA)



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
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleRepayHistoryV2(context.Background()).Timestamp(timestamp).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleRepayHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleRepayHistoryV2`: GetLoanFlexibleRepayHistoryV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleRepayHistoryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleRepayHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleRepayHistoryV2Resp**](GetLoanFlexibleRepayHistoryV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanFlexibleRepayRateV2

> GetLoanFlexibleRepayRateV2Resp GetLoanFlexibleRepayRateV2(ctx).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Check Collateral Repay Rate (USER_DATA)



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
	loanCoin := "loanCoin_example" // string |  (default to "")
	collateralCoin := "collateralCoin_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanFlexibleRepayRateV2(context.Background()).LoanCoin(loanCoin).CollateralCoin(collateralCoin).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanFlexibleRepayRateV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanFlexibleRepayRateV2`: GetLoanFlexibleRepayRateV2Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanFlexibleRepayRateV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanFlexibleRepayRateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanFlexibleRepayRateV2Resp**](GetLoanFlexibleRepayRateV2Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanIncomeV1

> []GetLoanIncomeV1RespItem GetLoanIncomeV1(ctx).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()

Get Crypto Loans Income History(USER_DATA)



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
	asset := "asset_example" // string |  (optional) (default to "")
	type_ := "type__example" // string | All types will be returned by default. Enum：`borrowIn` ,`collateralSpent`, `repayAmount`, `collateralReturn`(Collateral return after repayment), `addCollateral`, `removeCollateral`, `collateralReturnAfterLiquidation` (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	limit := int32(56) // int32 | default 20, max 100 (optional) (default to 20)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanIncomeV1(context.Background()).Timestamp(timestamp).Asset(asset).Type_(type_).StartTime(startTime).EndTime(endTime).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanIncomeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanIncomeV1`: []GetLoanIncomeV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanIncomeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanIncomeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **asset** | **string** |  | [default to &quot;&quot;]
 **type_** | **string** | All types will be returned by default. Enum：&#x60;borrowIn&#x60; ,&#x60;collateralSpent&#x60;, &#x60;repayAmount&#x60;, &#x60;collateralReturn&#x60;(Collateral return after repayment), &#x60;addCollateral&#x60;, &#x60;removeCollateral&#x60;, &#x60;collateralReturnAfterLiquidation&#x60; | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **limit** | **int32** | default 20, max 100 | [default to 20]
 **recvWindow** | **int64** |  | 

### Return type

[**[]GetLoanIncomeV1RespItem**](GetLoanIncomeV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanLtvAdjustmentHistoryV1

> GetLoanLtvAdjustmentHistoryV1Resp GetLoanLtvAdjustmentHistoryV1(ctx).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan LTV Adjustment History(USER_DATA)



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
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanLtvAdjustmentHistoryV1(context.Background()).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanLtvAdjustmentHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanLtvAdjustmentHistoryV1`: GetLoanLtvAdjustmentHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanLtvAdjustmentHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanLtvAdjustmentHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanLtvAdjustmentHistoryV1Resp**](GetLoanLtvAdjustmentHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoanRepayHistoryV1

> GetLoanRepayHistoryV1Resp GetLoanRepayHistoryV1(ctx).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()

Get Loan Repayment History(USER_DATA)



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
	collateralCoin := "collateralCoin_example" // string |  (optional) (default to "")
	startTime := int64(789) // int64 |  (optional)
	endTime := int64(789) // int64 |  (optional)
	current := int64(789) // int64 | Current querying page. Start from 1; default: 1; max: 1000 (optional) (default to 1)
	limit := int64(789) // int64 | Default: 10; max: 100 (optional) (default to 10)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CryptoLoanAPI.GetLoanRepayHistoryV1(context.Background()).Timestamp(timestamp).OrderId(orderId).LoanCoin(loanCoin).CollateralCoin(collateralCoin).StartTime(startTime).EndTime(endTime).Current(current).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CryptoLoanAPI.GetLoanRepayHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoanRepayHistoryV1`: GetLoanRepayHistoryV1Resp
	fmt.Fprintf(os.Stdout, "Response from `CryptoLoanAPI.GetLoanRepayHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoanRepayHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **orderId** | **int64** |  | 
 **loanCoin** | **string** |  | [default to &quot;&quot;]
 **collateralCoin** | **string** |  | [default to &quot;&quot;]
 **startTime** | **int64** |  | 
 **endTime** | **int64** |  | 
 **current** | **int64** | Current querying page. Start from 1; default: 1; max: 1000 | [default to 1]
 **limit** | **int64** | Default: 10; max: 100 | [default to 10]
 **recvWindow** | **int64** |  | 

### Return type

[**GetLoanRepayHistoryV1Resp**](GetLoanRepayHistoryV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

