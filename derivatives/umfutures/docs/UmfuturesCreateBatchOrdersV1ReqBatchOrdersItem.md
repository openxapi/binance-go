# UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem

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

### NewUmfuturesCreateBatchOrdersV1ReqBatchOrdersItem

`func NewUmfuturesCreateBatchOrdersV1ReqBatchOrdersItem(quantity string, side string, symbol string, type_ string, ) *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem`

NewUmfuturesCreateBatchOrdersV1ReqBatchOrdersItem instantiates a new UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesCreateBatchOrdersV1ReqBatchOrdersItemWithDefaults

`func NewUmfuturesCreateBatchOrdersV1ReqBatchOrdersItemWithDefaults() *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem`

NewUmfuturesCreateBatchOrdersV1ReqBatchOrdersItemWithDefaults instantiates a new UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetActivationPrice() string`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetActivationPriceOk() (*string, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetActivationPrice(v string)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetCallbackRate

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetCallbackRate() string`

GetCallbackRate returns the CallbackRate field if non-nil, zero value otherwise.

### GetCallbackRateOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetCallbackRateOk() (*string, bool)`

GetCallbackRateOk returns a tuple with the CallbackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackRate

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetCallbackRate(v string)`

SetCallbackRate sets CallbackRate field to given value.

### HasCallbackRate

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasCallbackRate() bool`

HasCallbackRate returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetNewClientOrderId

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewClientOrderId() string`

GetNewClientOrderId returns the NewClientOrderId field if non-nil, zero value otherwise.

### GetNewClientOrderIdOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewClientOrderIdOk() (*string, bool)`

GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientOrderId

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetNewClientOrderId(v string)`

SetNewClientOrderId sets NewClientOrderId field to given value.

### HasNewClientOrderId

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasNewClientOrderId() bool`

HasNewClientOrderId returns a boolean if a field has been set.

### GetNewOrderRespType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewOrderRespType() string`

GetNewOrderRespType returns the NewOrderRespType field if non-nil, zero value otherwise.

### GetNewOrderRespTypeOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetNewOrderRespTypeOk() (*string, bool)`

GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderRespType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetNewOrderRespType(v string)`

SetNewOrderRespType sets NewOrderRespType field to given value.

### HasNewOrderRespType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasNewOrderRespType() bool`

HasNewOrderRespType returns a boolean if a field has been set.

### GetPositionSide

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetPriceProtect

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceProtect() string`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetPriceProtectOk() (*string, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetPriceProtect(v string)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetQuantity

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetReduceOnly

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetSide(v string)`

SetSide sets Side field to given value.


### GetStopPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTimeInForce

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetType(v string)`

SetType sets Type field to given value.


### GetWorkingType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *UmfuturesCreateBatchOrdersV1ReqBatchOrdersItem) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


