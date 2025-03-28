# SpotCreateSorOrderV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**Fills** | Pointer to [**[]SpotCreateSorOrderV3RespFillsInner**](SpotCreateSorOrderV3RespFillsInner.md) |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**OrigQuoteOrderQty** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**TransactTime** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UsedSor** | Pointer to **bool** |  | [optional] 
**WorkingFloor** | Pointer to **string** |  | [optional] 
**WorkingTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpotCreateSorOrderV3Resp

`func NewSpotCreateSorOrderV3Resp() *SpotCreateSorOrderV3Resp`

NewSpotCreateSorOrderV3Resp instantiates a new SpotCreateSorOrderV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotCreateSorOrderV3RespWithDefaults

`func NewSpotCreateSorOrderV3RespWithDefaults() *SpotCreateSorOrderV3Resp`

NewSpotCreateSorOrderV3RespWithDefaults instantiates a new SpotCreateSorOrderV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *SpotCreateSorOrderV3Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *SpotCreateSorOrderV3Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *SpotCreateSorOrderV3Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *SpotCreateSorOrderV3Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *SpotCreateSorOrderV3Resp) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *SpotCreateSorOrderV3Resp) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *SpotCreateSorOrderV3Resp) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *SpotCreateSorOrderV3Resp) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *SpotCreateSorOrderV3Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *SpotCreateSorOrderV3Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *SpotCreateSorOrderV3Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *SpotCreateSorOrderV3Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFills

`func (o *SpotCreateSorOrderV3Resp) GetFills() []SpotCreateSorOrderV3RespFillsInner`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *SpotCreateSorOrderV3Resp) GetFillsOk() (*[]SpotCreateSorOrderV3RespFillsInner, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *SpotCreateSorOrderV3Resp) SetFills(v []SpotCreateSorOrderV3RespFillsInner)`

SetFills sets Fills field to given value.

### HasFills

`func (o *SpotCreateSorOrderV3Resp) HasFills() bool`

HasFills returns a boolean if a field has been set.

### GetOrderId

`func (o *SpotCreateSorOrderV3Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SpotCreateSorOrderV3Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SpotCreateSorOrderV3Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SpotCreateSorOrderV3Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *SpotCreateSorOrderV3Resp) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *SpotCreateSorOrderV3Resp) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *SpotCreateSorOrderV3Resp) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *SpotCreateSorOrderV3Resp) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrigQty

`func (o *SpotCreateSorOrderV3Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *SpotCreateSorOrderV3Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *SpotCreateSorOrderV3Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *SpotCreateSorOrderV3Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *SpotCreateSorOrderV3Resp) GetOrigQuoteOrderQty() string`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *SpotCreateSorOrderV3Resp) GetOrigQuoteOrderQtyOk() (*string, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *SpotCreateSorOrderV3Resp) SetOrigQuoteOrderQty(v string)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *SpotCreateSorOrderV3Resp) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.

### GetPrice

`func (o *SpotCreateSorOrderV3Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpotCreateSorOrderV3Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SpotCreateSorOrderV3Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SpotCreateSorOrderV3Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *SpotCreateSorOrderV3Resp) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *SpotCreateSorOrderV3Resp) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *SpotCreateSorOrderV3Resp) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *SpotCreateSorOrderV3Resp) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *SpotCreateSorOrderV3Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *SpotCreateSorOrderV3Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *SpotCreateSorOrderV3Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *SpotCreateSorOrderV3Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *SpotCreateSorOrderV3Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpotCreateSorOrderV3Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpotCreateSorOrderV3Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpotCreateSorOrderV3Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotCreateSorOrderV3Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotCreateSorOrderV3Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotCreateSorOrderV3Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotCreateSorOrderV3Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *SpotCreateSorOrderV3Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *SpotCreateSorOrderV3Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *SpotCreateSorOrderV3Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *SpotCreateSorOrderV3Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTransactTime

`func (o *SpotCreateSorOrderV3Resp) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *SpotCreateSorOrderV3Resp) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *SpotCreateSorOrderV3Resp) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *SpotCreateSorOrderV3Resp) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetType

`func (o *SpotCreateSorOrderV3Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpotCreateSorOrderV3Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpotCreateSorOrderV3Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpotCreateSorOrderV3Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsedSor

`func (o *SpotCreateSorOrderV3Resp) GetUsedSor() bool`

GetUsedSor returns the UsedSor field if non-nil, zero value otherwise.

### GetUsedSorOk

`func (o *SpotCreateSorOrderV3Resp) GetUsedSorOk() (*bool, bool)`

GetUsedSorOk returns a tuple with the UsedSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSor

`func (o *SpotCreateSorOrderV3Resp) SetUsedSor(v bool)`

SetUsedSor sets UsedSor field to given value.

### HasUsedSor

`func (o *SpotCreateSorOrderV3Resp) HasUsedSor() bool`

HasUsedSor returns a boolean if a field has been set.

### GetWorkingFloor

`func (o *SpotCreateSorOrderV3Resp) GetWorkingFloor() string`

GetWorkingFloor returns the WorkingFloor field if non-nil, zero value otherwise.

### GetWorkingFloorOk

`func (o *SpotCreateSorOrderV3Resp) GetWorkingFloorOk() (*string, bool)`

GetWorkingFloorOk returns a tuple with the WorkingFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingFloor

`func (o *SpotCreateSorOrderV3Resp) SetWorkingFloor(v string)`

SetWorkingFloor sets WorkingFloor field to given value.

### HasWorkingFloor

`func (o *SpotCreateSorOrderV3Resp) HasWorkingFloor() bool`

HasWorkingFloor returns a boolean if a field has been set.

### GetWorkingTime

`func (o *SpotCreateSorOrderV3Resp) GetWorkingTime() int64`

GetWorkingTime returns the WorkingTime field if non-nil, zero value otherwise.

### GetWorkingTimeOk

`func (o *SpotCreateSorOrderV3Resp) GetWorkingTimeOk() (*int64, bool)`

GetWorkingTimeOk returns a tuple with the WorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTime

`func (o *SpotCreateSorOrderV3Resp) SetWorkingTime(v int64)`

SetWorkingTime sets WorkingTime field to given value.

### HasWorkingTime

`func (o *SpotCreateSorOrderV3Resp) HasWorkingTime() bool`

HasWorkingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


