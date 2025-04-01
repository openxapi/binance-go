# \V2API

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletGetLocalentityWithdrawHistoryV2**](V2API.md#WalletGetLocalentityWithdrawHistoryV2) | **Get** /sapi/v2/localentity/withdraw/history | Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)



## WalletGetLocalentityWithdrawHistoryV2

> []WalletGetLocalentityWithdrawHistoryV2RespItem WalletGetLocalentityWithdrawHistoryV2(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)



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
	trId := "trId_example" // string | Comma(,) separated list of travel rule record Ids. (optional) (default to "")
	txId := "txId_example" // string | Comma(,) separated list of transaction Ids. (optional) (default to "")
	withdrawOrderId := "withdrawOrderId_example" // string | Withdraw ID defined by the client (i.e. client&#39;s internal withdrawID). (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	coin := "coin_example" // string |  (optional) (default to "")
	travelRuleStatus := int32(56) // int32 | 0:Completed,1:Pending,2:Failed (optional)
	offset := int32(56) // int32 | Default: 0 (optional) (default to 0)
	limit := int32(56) // int32 | Default: 1000, Max: 1000 (optional) (default to 1000)
	startTime := int64(789) // int64 | Default: 90 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.WalletGetLocalentityWithdrawHistoryV2(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.WalletGetLocalentityWithdrawHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityWithdrawHistoryV2`: []WalletGetLocalentityWithdrawHistoryV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `V2API.WalletGetLocalentityWithdrawHistoryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetLocalentityWithdrawHistoryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | [default to &quot;&quot;]
 **txId** | **string** | Comma(,) separated list of transaction Ids. | [default to &quot;&quot;]
 **withdrawOrderId** | **string** | Withdraw ID defined by the client (i.e. client&amp;#39;s internal withdrawID). | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **travelRuleStatus** | **int32** | 0:Completed,1:Pending,2:Failed | 
 **offset** | **int32** | Default: 0 | [default to 0]
 **limit** | **int32** | Default: 1000, Max: 1000 | [default to 1000]
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletGetLocalentityWithdrawHistoryV2RespItem**](WalletGetLocalentityWithdrawHistoryV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

