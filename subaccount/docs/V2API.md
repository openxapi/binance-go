# \V2API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreateSubAccountSubAccountApiIpRestrictionV2**](V2API.md#SubaccountCreateSubAccountSubAccountApiIpRestrictionV2) | **Post** /sapi/v2/sub-account/subAccountApi/ipRestriction | Add IP Restriction for Sub-Account API key(For Master Account)
[**SubaccountGetSubAccountFuturesAccountSummaryV2**](V2API.md#SubaccountGetSubAccountFuturesAccountSummaryV2) | **Get** /sapi/v2/sub-account/futures/accountSummary | Get Summary of Sub-account&#39;s Futures Account V2(For Master Account)
[**SubaccountGetSubAccountFuturesAccountV2**](V2API.md#SubaccountGetSubAccountFuturesAccountV2) | **Get** /sapi/v2/sub-account/futures/account | Get Detail on Sub-account&#39;s Futures Account V2(For Master Account)
[**SubaccountGetSubAccountFuturesPositionRiskV2**](V2API.md#SubaccountGetSubAccountFuturesPositionRiskV2) | **Get** /sapi/v2/sub-account/futures/positionRisk | Get Futures Position-Risk of Sub-account V2(For Master Account)



## SubaccountCreateSubAccountSubAccountApiIpRestrictionV2

> SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp SubaccountCreateSubAccountSubAccountApiIpRestrictionV2(ctx).Email(email).Status(status).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()

Add IP Restriction for Sub-Account API key(For Master Account)



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
	status := "status_example" // string |  (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ipAddress := "ipAddress_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2(context.Background()).Email(email).Status(status).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountSubAccountApiIpRestrictionV2`: SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | [default to &quot;&quot;]
 **status** | **string** |  | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **ipAddress** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp**](SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesAccountSummaryV2

> SubaccountGetSubAccountFuturesAccountSummaryV2Resp SubaccountGetSubAccountFuturesAccountSummaryV2(ctx).FuturesType(futuresType).Timestamp(timestamp).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()

Get Summary of Sub-account's Futures Account V2(For Master Account)



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
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	page := int32(56) // int32 | default:1 (optional)
	limit := int32(56) // int32 | default:10, max:20 (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.SubaccountGetSubAccountFuturesAccountSummaryV2(context.Background()).FuturesType(futuresType).Timestamp(timestamp).Page(page).Limit(limit).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.SubaccountGetSubAccountFuturesAccountSummaryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountSummaryV2`: SubaccountGetSubAccountFuturesAccountSummaryV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.SubaccountGetSubAccountFuturesAccountSummaryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesAccountSummaryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **page** | **int32** | default:1 | 
 **limit** | **int32** | default:10, max:20 | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesAccountSummaryV2Resp**](SubaccountGetSubAccountFuturesAccountSummaryV2Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountFuturesAccountV2

> SubaccountGetSubAccountFuturesAccountV2Resp SubaccountGetSubAccountFuturesAccountV2(ctx).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get Detail on Sub-account's Futures Account V2(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2#email-address\">Sub-account email</a> (default to "")
	futuresType := int32(56) // int32 | 1:USDT Margined Futures, 2:COIN Margined Futures
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.SubaccountGetSubAccountFuturesAccountV2(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.SubaccountGetSubAccountFuturesAccountV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesAccountV2`: SubaccountGetSubAccountFuturesAccountV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.SubaccountGetSubAccountFuturesAccountV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountFuturesAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/asset-management/Get-Detail-on-Sub-accounts-Futures-Account-V2#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **futuresType** | **int32** | 1:USDT Margined Futures, 2:COIN Margined Futures | 
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountFuturesAccountV2Resp**](SubaccountGetSubAccountFuturesAccountV2Resp.md)

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
	resp, r, err := apiClient.V2API.SubaccountGetSubAccountFuturesPositionRiskV2(context.Background()).Email(email).FuturesType(futuresType).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.SubaccountGetSubAccountFuturesPositionRiskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountFuturesPositionRiskV2`: SubaccountGetSubAccountFuturesPositionRiskV2Resp
	fmt.Fprintf(os.Stdout, "Response from `V2API.SubaccountGetSubAccountFuturesPositionRiskV2`: %v\n", resp)
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

