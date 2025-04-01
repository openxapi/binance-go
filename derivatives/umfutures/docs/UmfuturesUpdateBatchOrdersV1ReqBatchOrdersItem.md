# UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] [default to ""]
**Price** | **string** |  | [default to ""]
**PriceMatch** | Pointer to **string** |  | [optional] [default to ""]
**Quantity** | **string** |  | [default to ""]
**Side** | **string** |  | [default to ""]
**Symbol** | **string** |  | [default to ""]

## Methods

### NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem

`func NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem(price string, quantity string, side string, symbol string, ) *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem`

NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem instantiates a new UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults

`func NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults() *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem`

NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults instantiates a new UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetPrice(v string)`

SetPrice sets Price field to given value.


### GetPriceMatch

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetQuantity

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetSide

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetSide(v string)`

SetSide sets Side field to given value.


### GetSymbol

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


