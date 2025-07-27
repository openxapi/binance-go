# CmfuturesCreateBatchOrderV1ReqBatchOrdersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationPrice** | Pointer to **string** |  | [optional] [default to ""]
**CallbackRate** | Pointer to **string** |  | [optional] [default to ""]
**ClosePosition** | Pointer to **string** |  | [optional] [default to ""]
**NewClientOrderId** | Pointer to **string** |  | [optional] [default to ""]
**NewOrderRespType** | Pointer to **string** |  | [optional] [default to ""]
**PositionSide** | Pointer to **string** |  | [optional] [default to ""]
**Price** | Pointer to **string** |  | [optional] [default to ""]
**PriceMatch** | Pointer to **string** |  | [optional] [default to ""]
**PriceProtect** | Pointer to **string** |  | [optional] [default to ""]
**Quantity** | Pointer to **string** |  | [optional] [default to ""]
**ReduceOnly** | Pointer to **string** |  | [optional] [default to ""]
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] [default to ""]
**Side** | **string** |  | [default to ""]
**StopPrice** | Pointer to **string** |  | [optional] [default to ""]
**Symbol** | **string** |  | [default to ""]
**TimeInForce** | Pointer to **string** |  | [optional] [default to ""]
**Type** | **string** |  | [default to ""]
**WorkingType** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItem

`func NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItem(side string, symbol string, type_ string, ) *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem`

NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItem instantiates a new CmfuturesCreateBatchOrderV1ReqBatchOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItemWithDefaults

`func NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItemWithDefaults() *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem`

NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItemWithDefaults instantiates a new CmfuturesCreateBatchOrderV1ReqBatchOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetActivationPrice() string`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetActivationPriceOk() (*string, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetActivationPrice(v string)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCallbackRate

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetCallbackRate() string`

GetCallbackRate returns the CallbackRate field if non-nil, zero value otherwise.

### GetCallbackRateOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetCallbackRateOk() (*string, bool)`

GetCallbackRateOk returns a tuple with the CallbackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRate

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetCallbackRate(v string)`

SetCallbackRate sets CallbackRate field to given value.

### HasCallbackRate

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasCallbackRate() bool`

HasCallbackRate returns a boolean if a field has been set.

### GetClosePosition

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetClosePosition() string`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetClosePositionOk() (*string, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetClosePosition(v string)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetNewClientOrderId

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewClientOrderId() string`

GetNewClientOrderId returns the NewClientOrderId field if non-nil, zero value otherwise.

### GetNewClientOrderIdOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewClientOrderIdOk() (*string, bool)`

GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientOrderId

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetNewClientOrderId(v string)`

SetNewClientOrderId sets NewClientOrderId field to given value.

### HasNewClientOrderId

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasNewClientOrderId() bool`

HasNewClientOrderId returns a boolean if a field has been set.

### GetNewOrderRespType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewOrderRespType() string`

GetNewOrderRespType returns the NewOrderRespType field if non-nil, zero value otherwise.

### GetNewOrderRespTypeOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewOrderRespTypeOk() (*string, bool)`

GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderRespType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetNewOrderRespType(v string)`

SetNewOrderRespType sets NewOrderRespType field to given value.

### HasNewOrderRespType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasNewOrderRespType() bool`

HasNewOrderRespType returns a boolean if a field has been set.

### GetPositionSide

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetPriceProtect

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceProtect() string`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceProtectOk() (*string, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPriceProtect(v string)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetQuantity

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetSide(v string)`

SetSide sets Side field to given value.


### GetStopPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTimeInForce

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetType(v string)`

SetType sets Type field to given value.


### GetWorkingType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


