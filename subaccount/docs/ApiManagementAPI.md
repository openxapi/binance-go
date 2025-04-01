# \ApiManagementAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreateSubAccountSubAccountApiIpRestrictionV2**](ApiManagementAPI.md#SubaccountCreateSubAccountSubAccountApiIpRestrictionV2) | **Post** /sapi/v2/sub-account/subAccountApi/ipRestriction | Add IP Restriction for Sub-Account API key(For Master Account)
[**SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1**](ApiManagementAPI.md#SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1) | **Delete** /sapi/v1/sub-account/subAccountApi/ipRestriction/ipList | Delete IP List For a Sub-account API Key(For Master Account)
[**SubaccountGetSubAccountSubAccountApiIpRestrictionV1**](ApiManagementAPI.md#SubaccountGetSubAccountSubAccountApiIpRestrictionV1) | **Get** /sapi/v1/sub-account/subAccountApi/ipRestriction | Get IP Restriction for a Sub-account API Key(For Master Account)



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
	resp, r, err := apiClient.ApiManagementAPI.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2(context.Background()).Email(email).Status(status).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiManagementAPI.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreateSubAccountSubAccountApiIpRestrictionV2`: SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp
	fmt.Fprintf(os.Stdout, "Response from `ApiManagementAPI.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2`: %v\n", resp)
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


## SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1

> SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()

Delete IP List For a Sub-account API Key(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key#email-address\">Sub-account email</a> (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	ipAddress := "ipAddress_example" // string | Can be added in batches, separated by commas (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiManagementAPI.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).IpAddress(ipAddress).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiManagementAPI.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1`: SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ApiManagementAPI.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **ipAddress** | **string** | Can be added in batches, separated by commas | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp**](SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubaccountGetSubAccountSubAccountApiIpRestrictionV1

> SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp SubaccountGetSubAccountSubAccountApiIpRestrictionV1(ctx).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Get IP Restriction for a Sub-account API Key(For Master Account)



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
	email := "email_example" // string | <a href=\"/docs/sub_account/api-management#email-address\">Sub-account email</a> (default to "")
	subAccountApiKey := "subAccountApiKey_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiManagementAPI.SubaccountGetSubAccountSubAccountApiIpRestrictionV1(context.Background()).Email(email).SubAccountApiKey(subAccountApiKey).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiManagementAPI.SubaccountGetSubAccountSubAccountApiIpRestrictionV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountGetSubAccountSubAccountApiIpRestrictionV1`: SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp
	fmt.Fprintf(os.Stdout, "Response from `ApiManagementAPI.SubaccountGetSubAccountSubAccountApiIpRestrictionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | &lt;a href&#x3D;\&quot;/docs/sub_account/api-management#email-address\&quot;&gt;Sub-account email&lt;/a&gt; | [default to &quot;&quot;]
 **subAccountApiKey** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp**](SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

