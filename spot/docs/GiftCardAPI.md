# \GiftCardAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGiftcardBuyCodeV1**](GiftCardAPI.md#CreateGiftcardBuyCodeV1) | **Post** /sapi/v1/giftcard/buyCode | Create a dual-token gift card(fixed value, discount feature)(TRADE)
[**CreateGiftcardCreateCodeV1**](GiftCardAPI.md#CreateGiftcardCreateCodeV1) | **Post** /sapi/v1/giftcard/createCode | Create a single-token gift card (USER_DATA)
[**CreateGiftcardRedeemCodeV1**](GiftCardAPI.md#CreateGiftcardRedeemCodeV1) | **Post** /sapi/v1/giftcard/redeemCode | Redeem a Binance Gift Card(USER_DATA)
[**GetGiftcardBuyCodeTokenLimitV1**](GiftCardAPI.md#GetGiftcardBuyCodeTokenLimitV1) | **Get** /sapi/v1/giftcard/buyCode/token-limit | Fetch Token Limit(USER_DATA)
[**GetGiftcardCryptographyRsaPublicKeyV1**](GiftCardAPI.md#GetGiftcardCryptographyRsaPublicKeyV1) | **Get** /sapi/v1/giftcard/cryptography/rsa-public-key | Fetch RSA Public Key(USER_DATA)
[**GetGiftcardVerifyV1**](GiftCardAPI.md#GetGiftcardVerifyV1) | **Get** /sapi/v1/giftcard/verify | Verify Binance Gift Card by Gift Card Number(USER_DATA)



## CreateGiftcardBuyCodeV1

> CreateGiftcardBuyCodeV1Resp CreateGiftcardBuyCodeV1(ctx).BaseToken(baseToken).BaseTokenAmount(baseTokenAmount).FaceToken(faceToken).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Create a dual-token gift card(fixed value, discount feature)(TRADE)



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
	baseToken := "baseToken_example" // string |  (default to "")
	baseTokenAmount := float32(8.14) // float32 | 
	faceToken := "faceToken_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardAPI.CreateGiftcardBuyCodeV1(context.Background()).BaseToken(baseToken).BaseTokenAmount(baseTokenAmount).FaceToken(faceToken).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardAPI.CreateGiftcardBuyCodeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGiftcardBuyCodeV1`: CreateGiftcardBuyCodeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `GiftCardAPI.CreateGiftcardBuyCodeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGiftcardBuyCodeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseToken** | **string** |  | [default to &quot;&quot;]
 **baseTokenAmount** | **float32** |  | 
 **faceToken** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**CreateGiftcardBuyCodeV1Resp**](CreateGiftcardBuyCodeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGiftcardCreateCodeV1

> CreateGiftcardCreateCodeV1Resp CreateGiftcardCreateCodeV1(ctx).Amount(amount).Timestamp(timestamp).Token(token).RecvWindow(recvWindow).Execute()

Create a single-token gift card (USER_DATA)



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
	amount := float32(8.14) // float32 | 
	timestamp := int64(789) // int64 | 
	token := "token_example" // string |  (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardAPI.CreateGiftcardCreateCodeV1(context.Background()).Amount(amount).Timestamp(timestamp).Token(token).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardAPI.CreateGiftcardCreateCodeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGiftcardCreateCodeV1`: CreateGiftcardCreateCodeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `GiftCardAPI.CreateGiftcardCreateCodeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGiftcardCreateCodeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float32** |  | 
 **timestamp** | **int64** |  | 
 **token** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateGiftcardCreateCodeV1Resp**](CreateGiftcardCreateCodeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGiftcardRedeemCodeV1

> CreateGiftcardRedeemCodeV1Resp CreateGiftcardRedeemCodeV1(ctx).Code(code).Timestamp(timestamp).ExternalUid(externalUid).RecvWindow(recvWindow).Execute()

Redeem a Binance Gift Card(USER_DATA)



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
	code := "code_example" // string |  (default to "")
	timestamp := int64(789) // int64 | 
	externalUid := "externalUid_example" // string |  (optional) (default to "")
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardAPI.CreateGiftcardRedeemCodeV1(context.Background()).Code(code).Timestamp(timestamp).ExternalUid(externalUid).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardAPI.CreateGiftcardRedeemCodeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGiftcardRedeemCodeV1`: CreateGiftcardRedeemCodeV1Resp
	fmt.Fprintf(os.Stdout, "Response from `GiftCardAPI.CreateGiftcardRedeemCodeV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGiftcardRedeemCodeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** |  | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **externalUid** | **string** |  | [default to &quot;&quot;]
 **recvWindow** | **int64** |  | 

### Return type

[**CreateGiftcardRedeemCodeV1Resp**](CreateGiftcardRedeemCodeV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGiftcardBuyCodeTokenLimitV1

> GetGiftcardBuyCodeTokenLimitV1Resp GetGiftcardBuyCodeTokenLimitV1(ctx).BaseToken(baseToken).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fetch Token Limit(USER_DATA)



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
	baseToken := "baseToken_example" // string | The token you want to pay, example: BUSD (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardAPI.GetGiftcardBuyCodeTokenLimitV1(context.Background()).BaseToken(baseToken).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardAPI.GetGiftcardBuyCodeTokenLimitV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGiftcardBuyCodeTokenLimitV1`: GetGiftcardBuyCodeTokenLimitV1Resp
	fmt.Fprintf(os.Stdout, "Response from `GiftCardAPI.GetGiftcardBuyCodeTokenLimitV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftcardBuyCodeTokenLimitV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseToken** | **string** | The token you want to pay, example: BUSD | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetGiftcardBuyCodeTokenLimitV1Resp**](GetGiftcardBuyCodeTokenLimitV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGiftcardCryptographyRsaPublicKeyV1

> GetGiftcardCryptographyRsaPublicKeyV1Resp GetGiftcardCryptographyRsaPublicKeyV1(ctx).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Fetch RSA Public Key(USER_DATA)



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
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardAPI.GetGiftcardCryptographyRsaPublicKeyV1(context.Background()).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardAPI.GetGiftcardCryptographyRsaPublicKeyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGiftcardCryptographyRsaPublicKeyV1`: GetGiftcardCryptographyRsaPublicKeyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `GiftCardAPI.GetGiftcardCryptographyRsaPublicKeyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftcardCryptographyRsaPublicKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetGiftcardCryptographyRsaPublicKeyV1Resp**](GetGiftcardCryptographyRsaPublicKeyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGiftcardVerifyV1

> GetGiftcardVerifyV1Resp GetGiftcardVerifyV1(ctx).ReferenceNo(referenceNo).Timestamp(timestamp).RecvWindow(recvWindow).Execute()

Verify Binance Gift Card by Gift Card Number(USER_DATA)



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
	referenceNo := "referenceNo_example" // string | Enter the Gift Card Number (default to "")
	timestamp := int64(789) // int64 | 
	recvWindow := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardAPI.GetGiftcardVerifyV1(context.Background()).ReferenceNo(referenceNo).Timestamp(timestamp).RecvWindow(recvWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardAPI.GetGiftcardVerifyV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGiftcardVerifyV1`: GetGiftcardVerifyV1Resp
	fmt.Fprintf(os.Stdout, "Response from `GiftCardAPI.GetGiftcardVerifyV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftcardVerifyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referenceNo** | **string** | Enter the Gift Card Number | [default to &quot;&quot;]
 **timestamp** | **int64** |  | 
 **recvWindow** | **int64** |  | 

### Return type

[**GetGiftcardVerifyV1Resp**](GetGiftcardVerifyV1Resp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

