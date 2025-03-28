# UfuturesDeleteOrderV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatePrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**CumQty** | Pointer to **string** |  | [optional] 
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
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesDeleteOrderV1Resp

`func NewUfuturesDeleteOrderV1Resp() *UfuturesDeleteOrderV1Resp`

NewUfuturesDeleteOrderV1Resp instantiates a new UfuturesDeleteOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesDeleteOrderV1RespWithDefaults

`func NewUfuturesDeleteOrderV1RespWithDefaults() *UfuturesDeleteOrderV1Resp`

NewUfuturesDeleteOrderV1RespWithDefaults instantiates a new UfuturesDeleteOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatePrice

`func (o *UfuturesDeleteOrderV1Resp) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *UfuturesDeleteOrderV1Resp) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *UfuturesDeleteOrderV1Resp) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *UfuturesDeleteOrderV1Resp) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *UfuturesDeleteOrderV1Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *UfuturesDeleteOrderV1Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *UfuturesDeleteOrderV1Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *UfuturesDeleteOrderV1Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetClosePosition

`func (o *UfuturesDeleteOrderV1Resp) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *UfuturesDeleteOrderV1Resp) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *UfuturesDeleteOrderV1Resp) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *UfuturesDeleteOrderV1Resp) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetCumQty

`func (o *UfuturesDeleteOrderV1Resp) GetCumQty() string`

GetCumQty returns the CumQty field if non-nil, zero value otherwise.

### GetCumQtyOk

`func (o *UfuturesDeleteOrderV1Resp) GetCumQtyOk() (*string, bool)`

GetCumQtyOk returns a tuple with the CumQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQty

`func (o *UfuturesDeleteOrderV1Resp) SetCumQty(v string)`

SetCumQty sets CumQty field to given value.

### HasCumQty

`func (o *UfuturesDeleteOrderV1Resp) HasCumQty() bool`

HasCumQty returns a boolean if a field has been set.

### GetCumQuote

`func (o *UfuturesDeleteOrderV1Resp) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *UfuturesDeleteOrderV1Resp) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *UfuturesDeleteOrderV1Resp) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *UfuturesDeleteOrderV1Resp) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetExecutedQty

`func (o *UfuturesDeleteOrderV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *UfuturesDeleteOrderV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *UfuturesDeleteOrderV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *UfuturesDeleteOrderV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *UfuturesDeleteOrderV1Resp) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *UfuturesDeleteOrderV1Resp) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *UfuturesDeleteOrderV1Resp) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *UfuturesDeleteOrderV1Resp) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetOrderId

`func (o *UfuturesDeleteOrderV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UfuturesDeleteOrderV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UfuturesDeleteOrderV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UfuturesDeleteOrderV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *UfuturesDeleteOrderV1Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *UfuturesDeleteOrderV1Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *UfuturesDeleteOrderV1Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *UfuturesDeleteOrderV1Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *UfuturesDeleteOrderV1Resp) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *UfuturesDeleteOrderV1Resp) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *UfuturesDeleteOrderV1Resp) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *UfuturesDeleteOrderV1Resp) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPositionSide

`func (o *UfuturesDeleteOrderV1Resp) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UfuturesDeleteOrderV1Resp) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UfuturesDeleteOrderV1Resp) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UfuturesDeleteOrderV1Resp) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *UfuturesDeleteOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UfuturesDeleteOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UfuturesDeleteOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UfuturesDeleteOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *UfuturesDeleteOrderV1Resp) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *UfuturesDeleteOrderV1Resp) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *UfuturesDeleteOrderV1Resp) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *UfuturesDeleteOrderV1Resp) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetPriceProtect

`func (o *UfuturesDeleteOrderV1Resp) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *UfuturesDeleteOrderV1Resp) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *UfuturesDeleteOrderV1Resp) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *UfuturesDeleteOrderV1Resp) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceRate

`func (o *UfuturesDeleteOrderV1Resp) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *UfuturesDeleteOrderV1Resp) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *UfuturesDeleteOrderV1Resp) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *UfuturesDeleteOrderV1Resp) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetReduceOnly

`func (o *UfuturesDeleteOrderV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *UfuturesDeleteOrderV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *UfuturesDeleteOrderV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *UfuturesDeleteOrderV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *UfuturesDeleteOrderV1Resp) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *UfuturesDeleteOrderV1Resp) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *UfuturesDeleteOrderV1Resp) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *UfuturesDeleteOrderV1Resp) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *UfuturesDeleteOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UfuturesDeleteOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UfuturesDeleteOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *UfuturesDeleteOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *UfuturesDeleteOrderV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UfuturesDeleteOrderV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UfuturesDeleteOrderV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UfuturesDeleteOrderV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *UfuturesDeleteOrderV1Resp) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *UfuturesDeleteOrderV1Resp) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *UfuturesDeleteOrderV1Resp) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *UfuturesDeleteOrderV1Resp) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesDeleteOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesDeleteOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesDeleteOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UfuturesDeleteOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *UfuturesDeleteOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UfuturesDeleteOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UfuturesDeleteOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UfuturesDeleteOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *UfuturesDeleteOrderV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UfuturesDeleteOrderV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UfuturesDeleteOrderV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UfuturesDeleteOrderV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UfuturesDeleteOrderV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UfuturesDeleteOrderV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UfuturesDeleteOrderV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UfuturesDeleteOrderV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *UfuturesDeleteOrderV1Resp) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *UfuturesDeleteOrderV1Resp) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *UfuturesDeleteOrderV1Resp) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *UfuturesDeleteOrderV1Resp) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


