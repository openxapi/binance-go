# \TravelRuleAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletCreateLocalentityBrokerWithdrawApplyV1**](TravelRuleAPI.md#WalletCreateLocalentityBrokerWithdrawApplyV1) | **Post** /sapi/v1/localentity/broker/withdraw/apply | Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)
[**WalletCreateLocalentityWithdrawApplyV1**](TravelRuleAPI.md#WalletCreateLocalentityWithdrawApplyV1) | **Post** /sapi/v1/localentity/withdraw/apply | Withdraw (for local entities that require travel rule) (USER_DATA)
[**WalletGetLocalentityDepositHistoryV1**](TravelRuleAPI.md#WalletGetLocalentityDepositHistoryV1) | **Get** /sapi/v1/localentity/deposit/history | Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)
[**WalletGetLocalentityVaspV1**](TravelRuleAPI.md#WalletGetLocalentityVaspV1) | **Get** /sapi/v1/localentity/vasp | Onboarded VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WalletGetLocalentityWithdrawHistoryV1**](TravelRuleAPI.md#WalletGetLocalentityWithdrawHistoryV1) | **Get** /sapi/v1/localentity/withdraw/history | Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WalletGetLocalentityWithdrawHistoryV2**](TravelRuleAPI.md#WalletGetLocalentityWithdrawHistoryV2) | **Get** /sapi/v2/localentity/withdraw/history | Withdraw History V2 (for local entities that require travel rule) (supporting network) (USER_DATA)
[**WalletUpdateLocalentityBrokerDepositProvideInfoV1**](TravelRuleAPI.md#WalletUpdateLocalentityBrokerDepositProvideInfoV1) | **Put** /sapi/v1/localentity/broker/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)
[**WalletUpdateLocalentityDepositProvideInfoV1**](TravelRuleAPI.md#WalletUpdateLocalentityDepositProvideInfoV1) | **Put** /sapi/v1/localentity/deposit/provide-info | Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



## WalletCreateLocalentityBrokerWithdrawApplyV1

> WalletCreateLocalentityBrokerWithdrawApplyV1Resp WalletCreateLocalentityBrokerWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).OriginatorPii(originatorPii).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).WithdrawOrderId(withdrawOrderId).AddressName(addressName).AddressTag(addressTag).Network(network).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()

Broker Withdraw (for brokers of local entities that require travel rule) (USER_DATA)



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
	address := "address_example" // string |  (default to "")
	amount := map[string]interface{}{ ... } // map[string]interface{} | 
	coin := "coin_example" // string |  (default to "")
	originatorPii := "originatorPii_example" // string |  (default to "")
	questionnaire := "questionnaire_example" // string |  (default to "")
	signature := "signature_example" // string |  (default to "")
	subAccountId := "subAccountId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	withdrawOrderId := "withdrawOrderId_example" // string |  (default to "")
	addressName := "addressName_example" // string |  (optional) (default to "")
	addressTag := "addressTag_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	transactionFeeFlag := true // bool |  (optional) (default to false)
	walletType := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TravelRuleAPI.WalletCreateLocalentityBrokerWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).OriginatorPii(originatorPii).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).WithdrawOrderId(withdrawOrderId).AddressName(addressName).AddressTag(addressTag).Network(network).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletCreateLocalentityBrokerWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateLocalentityBrokerWithdrawApplyV1`: WalletCreateLocalentityBrokerWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletCreateLocalentityBrokerWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateLocalentityBrokerWithdrawApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | [default to &quot;&quot;]
 **amount** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **originatorPii** | **string** |  | [default to &quot;&quot;]
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **signature** | **string** |  | [default to &quot;&quot;]
 **subAccountId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **withdrawOrderId** | **string** |  | [default to &quot;&quot;]
 **addressName** | **string** |  | [default to &quot;&quot;]
 **addressTag** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **transactionFeeFlag** | **bool** |  | [default to false]
 **walletType** | **int32** |  | 

### Return type

[**WalletCreateLocalentityBrokerWithdrawApplyV1Resp**](WalletCreateLocalentityBrokerWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCreateLocalentityWithdrawApplyV1

> WalletCreateLocalentityWithdrawApplyV1Resp WalletCreateLocalentityWithdrawApplyV1(ctx).Address(address).Amount(amount).Coin(coin).Questionnaire(questionnaire).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()

Withdraw (for local entities that require travel rule) (USER_DATA)



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
	address := "address_example" // string |  (default to "")
	amount := "amount_example" // string |  (default to "")
	coin := "coin_example" // string |  (default to "")
	questionnaire := "questionnaire_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	addressTag := "addressTag_example" // string |  (optional) (default to "")
	name := "name_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)
	transactionFeeFlag := true // bool |  (optional) (default to false)
	walletType := int32(56) // int32 |  (optional)
	withdrawOrderId := "withdrawOrderId_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TravelRuleAPI.WalletCreateLocalentityWithdrawApplyV1(context.Background()).Address(address).Amount(amount).Coin(coin).Questionnaire(questionnaire).Timestamp(timestamp).AddressTag(addressTag).Name(name).Network(network).RecvWindow(recvWindow).TransactionFeeFlag(transactionFeeFlag).WalletType(walletType).WithdrawOrderId(withdrawOrderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletCreateLocalentityWithdrawApplyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCreateLocalentityWithdrawApplyV1`: WalletCreateLocalentityWithdrawApplyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletCreateLocalentityWithdrawApplyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCreateLocalentityWithdrawApplyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** |  | [default to &quot;&quot;]
 **amount** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **addressTag** | **string** |  | [default to &quot;&quot;]
 **name** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 
 **transactionFeeFlag** | **bool** |  | [default to false]
 **walletType** | **int32** |  | 
 **withdrawOrderId** | **string** |  | [default to &quot;&quot;]

### Return type

[**WalletCreateLocalentityWithdrawApplyV1Resp**](WalletCreateLocalentityWithdrawApplyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetLocalentityDepositHistoryV1

> []WalletGetLocalentityDepositHistoryV1RespItem WalletGetLocalentityDepositHistoryV1(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()

Deposit History (for local entities that required travel rule) (supporting network) (USER_DATA)



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
	tranId := "tranId_example" // string | Comma(,) separated list of wallet tran Ids. (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")
	coin := "coin_example" // string |  (optional) (default to "")
	travelRuleStatus := int32(56) // int32 | 0:Completed,1:Pending,2:Failed (optional)
	pendingQuestionnaire := true // bool | true: Only return records that pending deposit questionnaire. false/not provided: return all records. (optional)
	startTime := int64(789) // int64 | Default: 90 days from current timestamp (optional)
	endTime := int64(789) // int64 | Default: present timestamp (optional)
	offset := int32(56) // int32 | Default:0 (optional)
	limit := int32(56) // int32 | Default:1000, Max:1000 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TravelRuleAPI.WalletGetLocalentityDepositHistoryV1(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).TranId(tranId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).PendingQuestionnaire(pendingQuestionnaire).StartTime(startTime).EndTime(endTime).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletGetLocalentityDepositHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityDepositHistoryV1`: []WalletGetLocalentityDepositHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletGetLocalentityDepositHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetLocalentityDepositHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | [default to &quot;&quot;]
 **txId** | **string** | Comma(,) separated list of transaction Ids. | [default to &quot;&quot;]
 **tranId** | **string** | Comma(,) separated list of wallet tran Ids. | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **travelRuleStatus** | **int32** | 0:Completed,1:Pending,2:Failed | 
 **pendingQuestionnaire** | **bool** | true: Only return records that pending deposit questionnaire. false/not provided: return all records. | 
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **offset** | **int32** | Default:0 | 
 **limit** | **int32** | Default:1000, Max:1000 | 

### Return type

[**[]WalletGetLocalentityDepositHistoryV1RespItem**](WalletGetLocalentityDepositHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetLocalentityVaspV1

> []WalletGetLocalentityVaspV1RespItem WalletGetLocalentityVaspV1(ctx).Execute()

Onboarded VASP list (for local entities that require travel rule) (supporting network) (USER_DATA)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TravelRuleAPI.WalletGetLocalentityVaspV1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletGetLocalentityVaspV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityVaspV1`: []WalletGetLocalentityVaspV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletGetLocalentityVaspV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetLocalentityVaspV1Request struct via the builder pattern


### Return type

[**[]WalletGetLocalentityVaspV1RespItem**](WalletGetLocalentityVaspV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetLocalentityWithdrawHistoryV1

> []WalletGetLocalentityWithdrawHistoryV1RespItem WalletGetLocalentityWithdrawHistoryV1(ctx).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()

Withdraw History (for local entities that require travel rule) (supporting network) (USER_DATA)



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
	withdrawOrderId := "withdrawOrderId_example" // string | Comma(,) separated list of withdrawID defined by the client (i.e. client&#39;s internal withdrawID). (optional) (default to "")
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
	resp, r, err := apiClient.TravelRuleAPI.WalletGetLocalentityWithdrawHistoryV1(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletGetLocalentityWithdrawHistoryV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityWithdrawHistoryV1`: []WalletGetLocalentityWithdrawHistoryV1RespItem
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletGetLocalentityWithdrawHistoryV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetLocalentityWithdrawHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **trId** | **string** | Comma(,) separated list of travel rule record Ids. | [default to &quot;&quot;]
 **txId** | **string** | Comma(,) separated list of transaction Ids. | [default to &quot;&quot;]
 **withdrawOrderId** | **string** | Comma(,) separated list of withdrawID defined by the client (i.e. client&amp;#39;s internal withdrawID). | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]
 **coin** | **string** |  | [default to &quot;&quot;]
 **travelRuleStatus** | **int32** | 0:Completed,1:Pending,2:Failed | 
 **offset** | **int32** | Default: 0 | [default to 0]
 **limit** | **int32** | Default: 1000, Max: 1000 | [default to 1000]
 **startTime** | **int64** | Default: 90 days from current timestamp | 
 **endTime** | **int64** | Default: present timestamp | 
 **recvWindow** | **int64** |  | 

### Return type

[**[]WalletGetLocalentityWithdrawHistoryV1RespItem**](WalletGetLocalentityWithdrawHistoryV1RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.TravelRuleAPI.WalletGetLocalentityWithdrawHistoryV2(context.Background()).Timestamp(timestamp).TrId(trId).TxId(txId).WithdrawOrderId(withdrawOrderId).Network(network).Coin(coin).TravelRuleStatus(travelRuleStatus).Offset(offset).Limit(limit).StartTime(startTime).EndTime(endTime).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletGetLocalentityWithdrawHistoryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletGetLocalentityWithdrawHistoryV2`: []WalletGetLocalentityWithdrawHistoryV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletGetLocalentityWithdrawHistoryV2`: %v\n", resp)
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


## WalletUpdateLocalentityBrokerDepositProvideInfoV1

> WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp WalletUpdateLocalentityBrokerDepositProvideInfoV1(ctx).BeneficiaryPii(beneficiaryPii).DepositId(depositId).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).Address(address).AddressTag(addressTag).Amount(amount).Coin(coin).Network(network).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



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
	beneficiaryPii := "beneficiaryPii_example" // string |  (default to "")
	depositId := "depositId_example" // string |  (default to "")
	questionnaire := "questionnaire_example" // string |  (default to "")
	signature := "signature_example" // string |  (default to "")
	subAccountId := "subAccountId_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	address := "address_example" // string |  (optional) (default to "")
	addressTag := "addressTag_example" // string |  (optional) (default to "")
	amount := map[string]interface{}{ ... } // map[string]interface{} |  (optional)
	coin := "coin_example" // string |  (optional) (default to "")
	network := "network_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TravelRuleAPI.WalletUpdateLocalentityBrokerDepositProvideInfoV1(context.Background()).BeneficiaryPii(beneficiaryPii).DepositId(depositId).Questionnaire(questionnaire).Signature(signature).SubAccountId(subAccountId).Timestamp(timestamp).Address(address).AddressTag(addressTag).Amount(amount).Coin(coin).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletUpdateLocalentityBrokerDepositProvideInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletUpdateLocalentityBrokerDepositProvideInfoV1`: WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletUpdateLocalentityBrokerDepositProvideInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletUpdateLocalentityBrokerDepositProvideInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beneficiaryPii** | **string** |  | [default to &quot;&quot;]
 **depositId** | **string** |  | [default to &quot;&quot;]
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **signature** | **string** |  | [default to &quot;&quot;]
 **subAccountId** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **address** | **string** |  | [default to &quot;&quot;]
 **addressTag** | **string** |  | [default to &quot;&quot;]
 **amount** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **coin** | **string** |  | [default to &quot;&quot;]
 **network** | **string** |  | [default to &quot;&quot;]

### Return type

[**WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp**](WalletUpdateLocalentityBrokerDepositProvideInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletUpdateLocalentityDepositProvideInfoV1

> WalletUpdateLocalentityDepositProvideInfoV1Resp WalletUpdateLocalentityDepositProvideInfoV1(ctx).Questionnaire(questionnaire).Timestamp(timestamp).TranId(tranId).Execute()

Submit Deposit Questionnaire (For local entities that require travel rule) (supporting network) (USER_DATA)



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
	questionnaire := "questionnaire_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	tranId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TravelRuleAPI.WalletUpdateLocalentityDepositProvideInfoV1(context.Background()).Questionnaire(questionnaire).Timestamp(timestamp).TranId(tranId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TravelRuleAPI.WalletUpdateLocalentityDepositProvideInfoV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletUpdateLocalentityDepositProvideInfoV1`: WalletUpdateLocalentityDepositProvideInfoV1Resp
	fmt.Fprintf(os.Stdout, "Response from `TravelRuleAPI.WalletUpdateLocalentityDepositProvideInfoV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletUpdateLocalentityDepositProvideInfoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionnaire** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **tranId** | **int64** |  | 

### Return type

[**WalletUpdateLocalentityDepositProvideInfoV1Resp**](WalletUpdateLocalentityDepositProvideInfoV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

