# \V2API

All URIs are relative to *https://dapi.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfuturesGetLeverageBracketV2**](V2API.md#CfuturesGetLeverageBracketV2) | **Get** /dapi/v2/leverageBracket | Notional Bracket for Symbol(USER_DATA)



## CfuturesGetLeverageBracketV2

> []CfuturesGetLeverageBracketV2RespItem CfuturesGetLeverageBracketV2(ctx).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()

Notional Bracket for Symbol(USER_DATA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func main() {
	timestamp := int64(789) // int64 | 
	symbol := "symbol_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V2API.CfuturesGetLeverageBracketV2(context.Background()).Timestamp(timestamp).Symbol(symbol).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V2API.CfuturesGetLeverageBracketV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CfuturesGetLeverageBracketV2`: []CfuturesGetLeverageBracketV2RespItem
	fmt.Fprintf(os.Stdout, "Response from `V2API.CfuturesGetLeverageBracketV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCfuturesGetLeverageBracketV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **symbol** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**[]CfuturesGetLeverageBracketV2RespItem**](CfuturesGetLeverageBracketV2RespItem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

