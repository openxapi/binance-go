# UfuturesGetOrderV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatePrice** | Pointer to **string** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**CumQuote** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**OrigType** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesGetOrderV1Resp

`func NewUfuturesGetOrderV1Resp() *UfuturesGetOrderV1Resp`

NewUfuturesGetOrderV1Resp instantiates a new UfuturesGetOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetOrderV1RespWithDefaults

`func NewUfuturesGetOrderV1RespWithDefaults() *UfuturesGetOrderV1Resp`

NewUfuturesGetOrderV1RespWithDefaults instantiates a new UfuturesGetOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatePrice

`func (o *UfuturesGetOrderV1Resp) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *UfuturesGetOrderV1Resp) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *UfuturesGetOrderV1Resp) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *UfuturesGetOrderV1Resp) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetAvgPrice

`func (o *UfuturesGetOrderV1Resp) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *UfuturesGetOrderV1Resp) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *UfuturesGetOrderV1Resp) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *UfuturesGetOrderV1Resp) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *UfuturesGetOrderV1Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *UfuturesGetOrderV1Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *UfuturesGetOrderV1Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *UfuturesGetOrderV1Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetClosePosition

`func (o *UfuturesGetOrderV1Resp) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *UfuturesGetOrderV1Resp) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *UfuturesGetOrderV1Resp) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *UfuturesGetOrderV1Resp) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetCumQuote

`func (o *UfuturesGetOrderV1Resp) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *UfuturesGetOrderV1Resp) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *UfuturesGetOrderV1Resp) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *UfuturesGetOrderV1Resp) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetExecutedQty

`func (o *UfuturesGetOrderV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *UfuturesGetOrderV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *UfuturesGetOrderV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *UfuturesGetOrderV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *UfuturesGetOrderV1Resp) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *UfuturesGetOrderV1Resp) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *UfuturesGetOrderV1Resp) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *UfuturesGetOrderV1Resp) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetOrderId

`func (o *UfuturesGetOrderV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UfuturesGetOrderV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UfuturesGetOrderV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UfuturesGetOrderV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *UfuturesGetOrderV1Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *UfuturesGetOrderV1Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *UfuturesGetOrderV1Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *UfuturesGetOrderV1Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *UfuturesGetOrderV1Resp) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *UfuturesGetOrderV1Resp) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *UfuturesGetOrderV1Resp) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *UfuturesGetOrderV1Resp) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPositionSide

`func (o *UfuturesGetOrderV1Resp) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UfuturesGetOrderV1Resp) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UfuturesGetOrderV1Resp) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UfuturesGetOrderV1Resp) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *UfuturesGetOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UfuturesGetOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UfuturesGetOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UfuturesGetOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *UfuturesGetOrderV1Resp) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *UfuturesGetOrderV1Resp) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *UfuturesGetOrderV1Resp) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *UfuturesGetOrderV1Resp) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetPriceProtect

`func (o *UfuturesGetOrderV1Resp) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *UfuturesGetOrderV1Resp) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *UfuturesGetOrderV1Resp) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *UfuturesGetOrderV1Resp) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceRate

`func (o *UfuturesGetOrderV1Resp) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *UfuturesGetOrderV1Resp) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *UfuturesGetOrderV1Resp) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *UfuturesGetOrderV1Resp) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetReduceOnly

`func (o *UfuturesGetOrderV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *UfuturesGetOrderV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *UfuturesGetOrderV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *UfuturesGetOrderV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *UfuturesGetOrderV1Resp) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *UfuturesGetOrderV1Resp) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *UfuturesGetOrderV1Resp) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *UfuturesGetOrderV1Resp) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *UfuturesGetOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UfuturesGetOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UfuturesGetOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *UfuturesGetOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *UfuturesGetOrderV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UfuturesGetOrderV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UfuturesGetOrderV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UfuturesGetOrderV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *UfuturesGetOrderV1Resp) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *UfuturesGetOrderV1Resp) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *UfuturesGetOrderV1Resp) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *UfuturesGetOrderV1Resp) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesGetOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesGetOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesGetOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UfuturesGetOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *UfuturesGetOrderV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UfuturesGetOrderV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UfuturesGetOrderV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UfuturesGetOrderV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *UfuturesGetOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UfuturesGetOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UfuturesGetOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UfuturesGetOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *UfuturesGetOrderV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UfuturesGetOrderV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UfuturesGetOrderV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UfuturesGetOrderV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UfuturesGetOrderV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UfuturesGetOrderV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UfuturesGetOrderV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UfuturesGetOrderV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *UfuturesGetOrderV1Resp) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *UfuturesGetOrderV1Resp) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *UfuturesGetOrderV1Resp) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *UfuturesGetOrderV1Resp) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


