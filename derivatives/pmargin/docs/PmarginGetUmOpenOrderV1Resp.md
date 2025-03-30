# PmarginGetUmOpenOrderV1Resp

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

### NewPmarginGetUmOpenOrderV1Resp

`func NewPmarginGetUmOpenOrderV1Resp() *PmarginGetUmOpenOrderV1Resp`

NewPmarginGetUmOpenOrderV1Resp instantiates a new PmarginGetUmOpenOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetUmOpenOrderV1RespWithDefaults

`func NewPmarginGetUmOpenOrderV1RespWithDefaults() *PmarginGetUmOpenOrderV1Resp`

NewPmarginGetUmOpenOrderV1RespWithDefaults instantiates a new PmarginGetUmOpenOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *PmarginGetUmOpenOrderV1Resp) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *PmarginGetUmOpenOrderV1Resp) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *PmarginGetUmOpenOrderV1Resp) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PmarginGetUmOpenOrderV1Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PmarginGetUmOpenOrderV1Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PmarginGetUmOpenOrderV1Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCumQuote

`func (o *PmarginGetUmOpenOrderV1Resp) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *PmarginGetUmOpenOrderV1Resp) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *PmarginGetUmOpenOrderV1Resp) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetExecutedQty

`func (o *PmarginGetUmOpenOrderV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *PmarginGetUmOpenOrderV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *PmarginGetUmOpenOrderV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *PmarginGetUmOpenOrderV1Resp) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *PmarginGetUmOpenOrderV1Resp) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *PmarginGetUmOpenOrderV1Resp) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetOrderId

`func (o *PmarginGetUmOpenOrderV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PmarginGetUmOpenOrderV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PmarginGetUmOpenOrderV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *PmarginGetUmOpenOrderV1Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *PmarginGetUmOpenOrderV1Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *PmarginGetUmOpenOrderV1Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *PmarginGetUmOpenOrderV1Resp) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *PmarginGetUmOpenOrderV1Resp) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *PmarginGetUmOpenOrderV1Resp) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPositionSide

`func (o *PmarginGetUmOpenOrderV1Resp) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *PmarginGetUmOpenOrderV1Resp) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *PmarginGetUmOpenOrderV1Resp) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *PmarginGetUmOpenOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PmarginGetUmOpenOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PmarginGetUmOpenOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *PmarginGetUmOpenOrderV1Resp) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *PmarginGetUmOpenOrderV1Resp) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *PmarginGetUmOpenOrderV1Resp) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetReduceOnly

`func (o *PmarginGetUmOpenOrderV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PmarginGetUmOpenOrderV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PmarginGetUmOpenOrderV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *PmarginGetUmOpenOrderV1Resp) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *PmarginGetUmOpenOrderV1Resp) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *PmarginGetUmOpenOrderV1Resp) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *PmarginGetUmOpenOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PmarginGetUmOpenOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PmarginGetUmOpenOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *PmarginGetUmOpenOrderV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PmarginGetUmOpenOrderV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PmarginGetUmOpenOrderV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginGetUmOpenOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginGetUmOpenOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginGetUmOpenOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *PmarginGetUmOpenOrderV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PmarginGetUmOpenOrderV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *PmarginGetUmOpenOrderV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PmarginGetUmOpenOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PmarginGetUmOpenOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PmarginGetUmOpenOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *PmarginGetUmOpenOrderV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PmarginGetUmOpenOrderV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PmarginGetUmOpenOrderV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PmarginGetUmOpenOrderV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PmarginGetUmOpenOrderV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PmarginGetUmOpenOrderV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PmarginGetUmOpenOrderV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


