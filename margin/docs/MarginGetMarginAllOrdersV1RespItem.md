# MarginGetMarginAllOrdersV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**IcebergQty** | Pointer to **string** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
**IsWorking** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarginGetMarginAllOrdersV1RespItem

`func NewMarginGetMarginAllOrdersV1RespItem() *MarginGetMarginAllOrdersV1RespItem`

NewMarginGetMarginAllOrdersV1RespItem instantiates a new MarginGetMarginAllOrdersV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginGetMarginAllOrdersV1RespItemWithDefaults

`func NewMarginGetMarginAllOrdersV1RespItemWithDefaults() *MarginGetMarginAllOrdersV1RespItem`

NewMarginGetMarginAllOrdersV1RespItemWithDefaults instantiates a new MarginGetMarginAllOrdersV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *MarginGetMarginAllOrdersV1RespItem) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *MarginGetMarginAllOrdersV1RespItem) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *MarginGetMarginAllOrdersV1RespItem) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *MarginGetMarginAllOrdersV1RespItem) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *MarginGetMarginAllOrdersV1RespItem) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *MarginGetMarginAllOrdersV1RespItem) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *MarginGetMarginAllOrdersV1RespItem) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *MarginGetMarginAllOrdersV1RespItem) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *MarginGetMarginAllOrdersV1RespItem) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetIcebergQty

`func (o *MarginGetMarginAllOrdersV1RespItem) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *MarginGetMarginAllOrdersV1RespItem) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *MarginGetMarginAllOrdersV1RespItem) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginGetMarginAllOrdersV1RespItem) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginGetMarginAllOrdersV1RespItem) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginGetMarginAllOrdersV1RespItem) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetIsWorking

`func (o *MarginGetMarginAllOrdersV1RespItem) GetIsWorking() bool`

GetIsWorking returns the IsWorking field if non-nil, zero value otherwise.

### GetIsWorkingOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetIsWorkingOk() (*bool, bool)`

GetIsWorkingOk returns a tuple with the IsWorking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorking

`func (o *MarginGetMarginAllOrdersV1RespItem) SetIsWorking(v bool)`

SetIsWorking sets IsWorking field to given value.

### HasIsWorking

`func (o *MarginGetMarginAllOrdersV1RespItem) HasIsWorking() bool`

HasIsWorking returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginGetMarginAllOrdersV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginGetMarginAllOrdersV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginGetMarginAllOrdersV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *MarginGetMarginAllOrdersV1RespItem) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *MarginGetMarginAllOrdersV1RespItem) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *MarginGetMarginAllOrdersV1RespItem) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *MarginGetMarginAllOrdersV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginGetMarginAllOrdersV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginGetMarginAllOrdersV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *MarginGetMarginAllOrdersV1RespItem) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *MarginGetMarginAllOrdersV1RespItem) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *MarginGetMarginAllOrdersV1RespItem) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *MarginGetMarginAllOrdersV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *MarginGetMarginAllOrdersV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *MarginGetMarginAllOrdersV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *MarginGetMarginAllOrdersV1RespItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarginGetMarginAllOrdersV1RespItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarginGetMarginAllOrdersV1RespItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *MarginGetMarginAllOrdersV1RespItem) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *MarginGetMarginAllOrdersV1RespItem) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *MarginGetMarginAllOrdersV1RespItem) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginGetMarginAllOrdersV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginGetMarginAllOrdersV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginGetMarginAllOrdersV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *MarginGetMarginAllOrdersV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarginGetMarginAllOrdersV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *MarginGetMarginAllOrdersV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *MarginGetMarginAllOrdersV1RespItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *MarginGetMarginAllOrdersV1RespItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *MarginGetMarginAllOrdersV1RespItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *MarginGetMarginAllOrdersV1RespItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarginGetMarginAllOrdersV1RespItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MarginGetMarginAllOrdersV1RespItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *MarginGetMarginAllOrdersV1RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *MarginGetMarginAllOrdersV1RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *MarginGetMarginAllOrdersV1RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *MarginGetMarginAllOrdersV1RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


