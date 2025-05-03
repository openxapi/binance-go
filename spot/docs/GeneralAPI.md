# \GeneralAPI

All URIs are relative to *https://api.binance.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpotGetExchangeInfoV3**](GeneralAPI.md#SpotGetExchangeInfoV3) | **Get** /api/v3/exchangeInfo | Exchange information



## SpotGetExchangeInfoV3

> SpotGetExchangeInfoV3Resp SpotGetExchangeInfoV3(ctx).Symbol(symbol).Symbols(symbols).Permissions(permissions).ShowPermissionSets(showPermissionSets).SymbolStatus(symbolStatus).Execute()

Exchange information



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
	symbol := "symbol_example" // string | Example: curl -X GET &#34;<a href=\"https://api.binance.com/api/v3/exchangeInfo?symbol=BNBBTC\" target=\"_blank\" rel=\"noopener noreferrer\">https://api.binance.com/api/v3/exchangeInfo?symbol=BNBBTC</a>&#34; (optional) (default to "")
	symbols := []string{"Inner_example"} // []string | Examples: curl -X GET &#34;<a href=\"https://api.binance.com/api/v3/exchangeInfo?symbols=%5B%22BNBBTC%22,%22BTCUSDT%22%5D\" target=\"_blank\" rel=\"noopener noreferrer\">https://api.binance.com/api/v3/exchangeInfo?symbols=%5B%22BNBBTC%22,%22BTCUSDT%22%5D</a>&#34; <br/> or <br/> curl -g -X  GET &#39;<a href=\"https://api.binance.com/api/v3/exchangeInfo?symbols=%5B%22BTCUSDT%22,%22BNBBTC\" target=\"_blank\" rel=\"noopener noreferrer\">https://api.binance.com/api/v3/exchangeInfo?symbols=[&#34;BTCUSDT&#34;,&#34;BNBBTC</a>&#34;]&#39; (optional)
	permissions := "permissions_example" // string | Examples: curl -X GET &#34;<a href=\"https://api.binance.com/api/v3/exchangeInfo?permissions=SPOT\" target=\"_blank\" rel=\"noopener noreferrer\">https://api.binance.com/api/v3/exchangeInfo?permissions=SPOT</a>&#34; <br/> or <br/> curl -X GET &#34;<a href=\"https://api.binance.com/api/v3/exchangeInfo?permissions=%5B%22MARGIN%22%2C%22LEVERAGED%22%5D\" target=\"_blank\" rel=\"noopener noreferrer\">https://api.binance.com/api/v3/exchangeInfo?permissions=%5B%22MARGIN%22%2C%22LEVERAGED%22%5D</a>&#34; <br/> or <br/> curl -g -X GET &#39;<a href=\"https://api.binance.com/api/v3/exchangeInfo?permissions=%5B%22MARGIN%22,%22LEVERAGED\" target=\"_blank\" rel=\"noopener noreferrer\">https://api.binance.com/api/v3/exchangeInfo?permissions=[&#34;MARGIN&#34;,&#34;LEVERAGED</a>&#34;]&#39; (optional) (default to "")
	showPermissionSets := true // bool | Controls whether the content of the `permissionSets` field is populated or not. Defaults to `true` (optional)
	symbolStatus := "symbolStatus_example" // string | Filters symbols that have this `tradingStatus`. Valid values: `TRADING`, `HALT`, `BREAK` <br/> Cannot be used in combination with `symbols` or `symbol`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeneralAPI.SpotGetExchangeInfoV3(context.Background()).Symbol(symbol).Symbols(symbols).Permissions(permissions).ShowPermissionSets(showPermissionSets).SymbolStatus(symbolStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeneralAPI.SpotGetExchangeInfoV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpotGetExchangeInfoV3`: SpotGetExchangeInfoV3Resp
	fmt.Fprintf(os.Stdout, "Response from `GeneralAPI.SpotGetExchangeInfoV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpotGetExchangeInfoV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Example: curl -X GET &amp;#34;&lt;a href&#x3D;\&quot;https://api.binance.com/api/v3/exchangeInfo?symbol&#x3D;BNBBTC\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;https://api.binance.com/api/v3/exchangeInfo?symbol&#x3D;BNBBTC&lt;/a&gt;&amp;#34; | [default to &quot;&quot;]
 **symbols** | **[]string** | Examples: curl -X GET &amp;#34;&lt;a href&#x3D;\&quot;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;%5B%22BNBBTC%22,%22BTCUSDT%22%5D\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;%5B%22BNBBTC%22,%22BTCUSDT%22%5D&lt;/a&gt;&amp;#34; &lt;br/&gt; or &lt;br/&gt; curl -g -X  GET &amp;#39;&lt;a href&#x3D;\&quot;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;%5B%22BTCUSDT%22,%22BNBBTC\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;https://api.binance.com/api/v3/exchangeInfo?symbols&#x3D;[&amp;#34;BTCUSDT&amp;#34;,&amp;#34;BNBBTC&lt;/a&gt;&amp;#34;]&amp;#39; | 
 **permissions** | **string** | Examples: curl -X GET &amp;#34;&lt;a href&#x3D;\&quot;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;SPOT\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;SPOT&lt;/a&gt;&amp;#34; &lt;br/&gt; or &lt;br/&gt; curl -X GET &amp;#34;&lt;a href&#x3D;\&quot;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;%5B%22MARGIN%22%2C%22LEVERAGED%22%5D\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;%5B%22MARGIN%22%2C%22LEVERAGED%22%5D&lt;/a&gt;&amp;#34; &lt;br/&gt; or &lt;br/&gt; curl -g -X GET &amp;#39;&lt;a href&#x3D;\&quot;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;%5B%22MARGIN%22,%22LEVERAGED\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;https://api.binance.com/api/v3/exchangeInfo?permissions&#x3D;[&amp;#34;MARGIN&amp;#34;,&amp;#34;LEVERAGED&lt;/a&gt;&amp;#34;]&amp;#39; | [default to &quot;&quot;]
 **showPermissionSets** | **bool** | Controls whether the content of the &#x60;permissionSets&#x60; field is populated or not. Defaults to &#x60;true&#x60; | 
 **symbolStatus** | **string** | Filters symbols that have this &#x60;tradingStatus&#x60;. Valid values: &#x60;TRADING&#x60;, &#x60;HALT&#x60;, &#x60;BREAK&#x60; &lt;br/&gt; Cannot be used in combination with &#x60;symbols&#x60; or &#x60;symbol&#x60;. | [default to &quot;&quot;]

### Return type

[**SpotGetExchangeInfoV3Resp**](SpotGetExchangeInfoV3Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

