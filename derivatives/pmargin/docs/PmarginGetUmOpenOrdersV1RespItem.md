# PmarginGetUmOpenOrdersV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CumQuote** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**OrigType** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginGetUmOpenOrdersV1RespItem

`func NewPmarginGetUmOpenOrdersV1RespItem() *PmarginGetUmOpenOrdersV1RespItem`

NewPmarginGetUmOpenOrdersV1RespItem instantiates a new PmarginGetUmOpenOrdersV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetUmOpenOrdersV1RespItemWithDefaults

`func NewPmarginGetUmOpenOrdersV1RespItemWithDefaults() *PmarginGetUmOpenOrdersV1RespItem`

NewPmarginGetUmOpenOrdersV1RespItemWithDefaults instantiates a new PmarginGetUmOpenOrdersV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCumQuote

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetExecutedQty

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetOrderId

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPositionSide

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetReduceOnly

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PmarginGetUmOpenOrdersV1RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PmarginGetUmOpenOrdersV1RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PmarginGetUmOpenOrdersV1RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


