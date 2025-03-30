# OptionsCreateBatchOrdersV1ReqOrdersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] [default to ""]
**IsMmp** | Pointer to **bool** |  | [optional] 
**NewOrderRespType** | Pointer to **string** |  | [optional] [default to ""]
**PostOnly** | Pointer to **bool** |  | [optional] [default to false]
**Price** | Pointer to **string** |  | [optional] [default to ""]
**Quantity** | **string** |  | [default to ""]
**ReduceOnly** | Pointer to **bool** |  | [optional] [default to false]
**Side** | **string** |  | [default to ""]
**Symbol** | **string** |  | [default to ""]
**TimeInForce** | Pointer to **string** |  | [optional] [default to ""]
**Type** | **string** |  | [default to ""]

## Methods

### NewOptionsCreateBatchOrdersV1ReqOrdersItem

`func NewOptionsCreateBatchOrdersV1ReqOrdersItem(quantity string, side string, symbol string, type_ string, ) *OptionsCreateBatchOrdersV1ReqOrdersItem`

NewOptionsCreateBatchOrdersV1ReqOrdersItem instantiates a new OptionsCreateBatchOrdersV1ReqOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsCreateBatchOrdersV1ReqOrdersItemWithDefaults

`func NewOptionsCreateBatchOrdersV1ReqOrdersItemWithDefaults() *OptionsCreateBatchOrdersV1ReqOrdersItem`

NewOptionsCreateBatchOrdersV1ReqOrdersItemWithDefaults instantiates a new OptionsCreateBatchOrdersV1ReqOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetIsMmp

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetIsMmp() bool`

GetIsMmp returns the IsMmp field if non-nil, zero value otherwise.

### GetIsMmpOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetIsMmpOk() (*bool, bool)`

GetIsMmpOk returns a tuple with the IsMmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMmp

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetIsMmp(v bool)`

SetIsMmp sets IsMmp field to given value.

### HasIsMmp

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) HasIsMmp() bool`

HasIsMmp returns a boolean if a field has been set.

### GetNewOrderRespType

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetNewOrderRespType() string`

GetNewOrderRespType returns the NewOrderRespType field if non-nil, zero value otherwise.

### GetNewOrderRespTypeOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetNewOrderRespTypeOk() (*string, bool)`

GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderRespType

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetNewOrderRespType(v string)`

SetNewOrderRespType sets NewOrderRespType field to given value.

### HasNewOrderRespType

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) HasNewOrderRespType() bool`

HasNewOrderRespType returns a boolean if a field has been set.

### GetPostOnly

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetPostOnly() bool`

GetPostOnly returns the PostOnly field if non-nil, zero value otherwise.

### GetPostOnlyOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetPostOnlyOk() (*bool, bool)`

GetPostOnlyOk returns a tuple with the PostOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOnly

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetPostOnly(v bool)`

SetPostOnly sets PostOnly field to given value.

### HasPostOnly

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) HasPostOnly() bool`

HasPostOnly returns a boolean if a field has been set.

### GetPrice

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetReduceOnly

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetSide(v string)`

SetSide sets Side field to given value.


### GetSymbol

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTimeInForce

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionsCreateBatchOrdersV1ReqOrdersItem) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


