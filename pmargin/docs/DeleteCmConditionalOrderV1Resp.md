# DeleteCmConditionalOrderV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatePrice** | Pointer to **string** |  | [optional] 
**BookTime** | Pointer to **int64** |  | [optional] 
**NewClientStrategyId** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**StrategyId** | Pointer to **int64** |  | [optional] 
**StrategyStatus** | Pointer to **string** |  | [optional] 
**StrategyType** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteCmConditionalOrderV1Resp

`func NewDeleteCmConditionalOrderV1Resp() *DeleteCmConditionalOrderV1Resp`

NewDeleteCmConditionalOrderV1Resp instantiates a new DeleteCmConditionalOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCmConditionalOrderV1RespWithDefaults

`func NewDeleteCmConditionalOrderV1RespWithDefaults() *DeleteCmConditionalOrderV1Resp`

NewDeleteCmConditionalOrderV1RespWithDefaults instantiates a new DeleteCmConditionalOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatePrice

`func (o *DeleteCmConditionalOrderV1Resp) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *DeleteCmConditionalOrderV1Resp) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *DeleteCmConditionalOrderV1Resp) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *DeleteCmConditionalOrderV1Resp) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetBookTime

`func (o *DeleteCmConditionalOrderV1Resp) GetBookTime() int64`

GetBookTime returns the BookTime field if non-nil, zero value otherwise.

### GetBookTimeOk

`func (o *DeleteCmConditionalOrderV1Resp) GetBookTimeOk() (*int64, bool)`

GetBookTimeOk returns a tuple with the BookTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTime

`func (o *DeleteCmConditionalOrderV1Resp) SetBookTime(v int64)`

SetBookTime sets BookTime field to given value.

### HasBookTime

`func (o *DeleteCmConditionalOrderV1Resp) HasBookTime() bool`

HasBookTime returns a boolean if a field has been set.

### GetNewClientStrategyId

`func (o *DeleteCmConditionalOrderV1Resp) GetNewClientStrategyId() string`

GetNewClientStrategyId returns the NewClientStrategyId field if non-nil, zero value otherwise.

### GetNewClientStrategyIdOk

`func (o *DeleteCmConditionalOrderV1Resp) GetNewClientStrategyIdOk() (*string, bool)`

GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientStrategyId

`func (o *DeleteCmConditionalOrderV1Resp) SetNewClientStrategyId(v string)`

SetNewClientStrategyId sets NewClientStrategyId field to given value.

### HasNewClientStrategyId

`func (o *DeleteCmConditionalOrderV1Resp) HasNewClientStrategyId() bool`

HasNewClientStrategyId returns a boolean if a field has been set.

### GetOrigQty

`func (o *DeleteCmConditionalOrderV1Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *DeleteCmConditionalOrderV1Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *DeleteCmConditionalOrderV1Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *DeleteCmConditionalOrderV1Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPositionSide

`func (o *DeleteCmConditionalOrderV1Resp) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *DeleteCmConditionalOrderV1Resp) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *DeleteCmConditionalOrderV1Resp) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *DeleteCmConditionalOrderV1Resp) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *DeleteCmConditionalOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DeleteCmConditionalOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DeleteCmConditionalOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DeleteCmConditionalOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceProtect

`func (o *DeleteCmConditionalOrderV1Resp) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *DeleteCmConditionalOrderV1Resp) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *DeleteCmConditionalOrderV1Resp) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *DeleteCmConditionalOrderV1Resp) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceRate

`func (o *DeleteCmConditionalOrderV1Resp) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *DeleteCmConditionalOrderV1Resp) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *DeleteCmConditionalOrderV1Resp) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *DeleteCmConditionalOrderV1Resp) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetReduceOnly

`func (o *DeleteCmConditionalOrderV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *DeleteCmConditionalOrderV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *DeleteCmConditionalOrderV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *DeleteCmConditionalOrderV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *DeleteCmConditionalOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *DeleteCmConditionalOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *DeleteCmConditionalOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *DeleteCmConditionalOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStopPrice

`func (o *DeleteCmConditionalOrderV1Resp) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *DeleteCmConditionalOrderV1Resp) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *DeleteCmConditionalOrderV1Resp) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *DeleteCmConditionalOrderV1Resp) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStrategyId

`func (o *DeleteCmConditionalOrderV1Resp) GetStrategyId() int64`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *DeleteCmConditionalOrderV1Resp) GetStrategyIdOk() (*int64, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *DeleteCmConditionalOrderV1Resp) SetStrategyId(v int64)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *DeleteCmConditionalOrderV1Resp) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.

### GetStrategyStatus

`func (o *DeleteCmConditionalOrderV1Resp) GetStrategyStatus() string`

GetStrategyStatus returns the StrategyStatus field if non-nil, zero value otherwise.

### GetStrategyStatusOk

`func (o *DeleteCmConditionalOrderV1Resp) GetStrategyStatusOk() (*string, bool)`

GetStrategyStatusOk returns a tuple with the StrategyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyStatus

`func (o *DeleteCmConditionalOrderV1Resp) SetStrategyStatus(v string)`

SetStrategyStatus sets StrategyStatus field to given value.

### HasStrategyStatus

`func (o *DeleteCmConditionalOrderV1Resp) HasStrategyStatus() bool`

HasStrategyStatus returns a boolean if a field has been set.

### GetStrategyType

`func (o *DeleteCmConditionalOrderV1Resp) GetStrategyType() string`

GetStrategyType returns the StrategyType field if non-nil, zero value otherwise.

### GetStrategyTypeOk

`func (o *DeleteCmConditionalOrderV1Resp) GetStrategyTypeOk() (*string, bool)`

GetStrategyTypeOk returns a tuple with the StrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyType

`func (o *DeleteCmConditionalOrderV1Resp) SetStrategyType(v string)`

SetStrategyType sets StrategyType field to given value.

### HasStrategyType

`func (o *DeleteCmConditionalOrderV1Resp) HasStrategyType() bool`

HasStrategyType returns a boolean if a field has been set.

### GetSymbol

`func (o *DeleteCmConditionalOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *DeleteCmConditionalOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *DeleteCmConditionalOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *DeleteCmConditionalOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *DeleteCmConditionalOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *DeleteCmConditionalOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *DeleteCmConditionalOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *DeleteCmConditionalOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetUpdateTime

`func (o *DeleteCmConditionalOrderV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *DeleteCmConditionalOrderV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *DeleteCmConditionalOrderV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *DeleteCmConditionalOrderV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *DeleteCmConditionalOrderV1Resp) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *DeleteCmConditionalOrderV1Resp) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *DeleteCmConditionalOrderV1Resp) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *DeleteCmConditionalOrderV1Resp) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


