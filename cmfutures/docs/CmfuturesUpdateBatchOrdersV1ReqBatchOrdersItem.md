# CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] [default to ""]
**Price** | Pointer to **string** |  | [optional] [default to ""]
**Quantity** | Pointer to **string** |  | [optional] [default to ""]
**RecvWindow** | Pointer to **int64** |  | [optional] 
**Side** | **string** |  | [default to ""]
**Symbol** | **string** |  | [default to ""]
**Timestamp** | **int64** |  | 

## Methods

### NewCmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem

`func NewCmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem(side string, symbol string, timestamp int64, ) *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem`

NewCmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem instantiates a new CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults

`func NewCmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults() *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem`

NewCmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults instantiates a new CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRecvWindow

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetRecvWindow() int64`

GetRecvWindow returns the RecvWindow field if non-nil, zero value otherwise.

### GetRecvWindowOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetRecvWindowOk() (*int64, bool)`

GetRecvWindowOk returns a tuple with the RecvWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvWindow

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetRecvWindow(v int64)`

SetRecvWindow sets RecvWindow field to given value.

### HasRecvWindow

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasRecvWindow() bool`

HasRecvWindow returns a boolean if a field has been set.

### GetSide

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetSide(v string)`

SetSide sets Side field to given value.


### GetSymbol

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTimestamp

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


