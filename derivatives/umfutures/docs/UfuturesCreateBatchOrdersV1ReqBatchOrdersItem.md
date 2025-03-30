# UfuturesCreateBatchOrdersV1ReqBatchOrdersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationPrice** | Pointer to **string** |  | [optional] [default to ""]
**CallbackRate** | Pointer to **string** |  | [optional] [default to ""]
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**NewClientOrderId** | Pointer to **string** |  | [optional] [default to ""]
**NewOrderRespType** | Pointer to **string** |  | [optional] [default to ""]
**PositionSide** | Pointer to **string** |  | [optional] [default to ""]
**Price** | Pointer to **string** |  | [optional] [default to ""]
**PriceMatch** | Pointer to **string** |  | [optional] [default to ""]
**PriceProtect** | Pointer to **string** |  | [optional] [default to ""]
**Quantity** | **string** |  | [default to ""]
**ReduceOnly** | Pointer to **string** |  | [optional] [default to ""]
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] [default to ""]
**Side** | **string** |  | [default to ""]
**StopPrice** | Pointer to **string** |  | [optional] [default to ""]
**Symbol** | **string** |  | [default to ""]
**TimeInForce** | Pointer to **string** |  | [optional] [default to ""]
**Type** | **string** |  | [default to ""]
**WorkingType** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewUfuturesCreateBatchOrdersV1ReqBatchOrdersItem

`func NewUfuturesCreateBatchOrdersV1ReqBatchOrdersItem(quantity string, side string, symbol string, type_ string, ) *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem`

NewUfuturesCreateBatchOrdersV1ReqBatchOrdersItem instantiates a new UfuturesCreateBatchOrdersV1ReqBatchOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesCreateBatchOrdersV1ReqBatchOrdersItemWithDefaults

`func NewUfuturesCreateBatchOrdersV1ReqBatchOrdersItemWithDefaults() *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem`

NewUfuturesCreateBatchOrdersV1ReqBatchOrdersItemWithDefaults instantiates a new UfuturesCreateBatchOrdersV1ReqBatchOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetActivationPrice() string`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetActivationPriceOk() (*string, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetActivationPrice(v string)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCallbackRate

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetCallbackRate() string`

GetCallbackRate returns the CallbackRate field if non-nil, zero value otherwise.

### GetCallbackRateOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetCallbackRateOk() (*string, bool)`

GetCallbackRateOk returns a tuple with the CallbackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRate

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetCallbackRate(v string)`

SetCallbackRate sets CallbackRate field to given value.

### HasCallbackRate

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasCallbackRate() bool`

HasCallbackRate returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetNewClientOrderId

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewClientOrderId() string`

GetNewClientOrderId returns the NewClientOrderId field if non-nil, zero value otherwise.

### GetNewClientOrderIdOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewClientOrderIdOk() (*string, bool)`

GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientOrderId

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetNewClientOrderId(v string)`

SetNewClientOrderId sets NewClientOrderId field to given value.

### HasNewClientOrderId

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasNewClientOrderId() bool`

HasNewClientOrderId returns a boolean if a field has been set.

### GetNewOrderRespType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewOrderRespType() string`

GetNewOrderRespType returns the NewOrderRespType field if non-nil, zero value otherwise.

### GetNewOrderRespTypeOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewOrderRespTypeOk() (*string, bool)`

GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderRespType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetNewOrderRespType(v string)`

SetNewOrderRespType sets NewOrderRespType field to given value.

### HasNewOrderRespType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasNewOrderRespType() bool`

HasNewOrderRespType returns a boolean if a field has been set.

### GetPositionSide

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetPriceProtect

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceProtect() string`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceProtectOk() (*string, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPriceProtect(v string)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetQuantity

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetReduceOnly

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetSide(v string)`

SetSide sets Side field to given value.


### GetStopPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTimeInForce

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetType(v string)`

SetType sets Type field to given value.


### GetWorkingType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *UfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


