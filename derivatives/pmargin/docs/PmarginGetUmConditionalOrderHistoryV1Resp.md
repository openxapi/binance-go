# PmarginGetUmConditionalOrderHistoryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatePrice** | Pointer to **string** |  | [optional] 
**BookTime** | Pointer to **int64** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**NewClientStrategyId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**StrategyId** | Pointer to **int64** |  | [optional] 
**StrategyStatus** | Pointer to **string** |  | [optional] 
**StrategyType** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**TriggerTime** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 

## Methods

### NewPmarginGetUmConditionalOrderHistoryV1Resp

`func NewPmarginGetUmConditionalOrderHistoryV1Resp() *PmarginGetUmConditionalOrderHistoryV1Resp`

NewPmarginGetUmConditionalOrderHistoryV1Resp instantiates a new PmarginGetUmConditionalOrderHistoryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetUmConditionalOrderHistoryV1RespWithDefaults

`func NewPmarginGetUmConditionalOrderHistoryV1RespWithDefaults() *PmarginGetUmConditionalOrderHistoryV1Resp`

NewPmarginGetUmConditionalOrderHistoryV1RespWithDefaults instantiates a new PmarginGetUmConditionalOrderHistoryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatePrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetBookTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetBookTime() int64`

GetBookTime returns the BookTime field if non-nil, zero value otherwise.

### GetBookTimeOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetBookTimeOk() (*int64, bool)`

GetBookTimeOk returns a tuple with the BookTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetBookTime(v int64)`

SetBookTime sets BookTime field to given value.

### HasBookTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasBookTime() bool`

HasBookTime returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetNewClientStrategyId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetNewClientStrategyId() string`

GetNewClientStrategyId returns the NewClientStrategyId field if non-nil, zero value otherwise.

### GetNewClientStrategyIdOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetNewClientStrategyIdOk() (*string, bool)`

GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientStrategyId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetNewClientStrategyId(v string)`

SetNewClientStrategyId sets NewClientStrategyId field to given value.

### HasNewClientStrategyId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasNewClientStrategyId() bool`

HasNewClientStrategyId returns a boolean if a field has been set.

### GetOrderId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPositionSide

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceProtect

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceRate

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetReduceOnly

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStrategyId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyStatus

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStrategyStatus() string`

GetStrategyStatus returns the StrategyStatus field if non-nil, zero value otherwise.

### GetStrategyStatusOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStrategyStatusOk() (*string, bool)`

GetStrategyStatusOk returns a tuple with the StrategyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyStatus

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetStrategyStatus(v string)`

SetStrategyStatus sets StrategyStatus field to given value.

### HasStrategyStatus

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasStrategyStatus() bool`

HasStrategyStatus returns a boolean if a field has been set.

### GetStrategyType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStrategyType() string`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetStrategyTypeOk() (*string, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetStrategyType(v string)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTriggerTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetTriggerTime() int64`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetTriggerTimeOk() (*int64, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetTriggerTime(v int64)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *PmarginGetUmConditionalOrderHistoryV1Resp) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


