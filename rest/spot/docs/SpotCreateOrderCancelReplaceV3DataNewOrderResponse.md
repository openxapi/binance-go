# SpotCreateOrderCancelReplaceV3DataNewOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**Fills** | Pointer to [**[]SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner**](SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner.md) |  | [optional] 
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
**WorkingTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpotCreateOrderCancelReplaceV3DataNewOrderResponse

`func NewSpotCreateOrderCancelReplaceV3DataNewOrderResponse() *SpotCreateOrderCancelReplaceV3DataNewOrderResponse`

NewSpotCreateOrderCancelReplaceV3DataNewOrderResponse instantiates a new SpotCreateOrderCancelReplaceV3DataNewOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotCreateOrderCancelReplaceV3DataNewOrderResponseWithDefaults

`func NewSpotCreateOrderCancelReplaceV3DataNewOrderResponseWithDefaults() *SpotCreateOrderCancelReplaceV3DataNewOrderResponse`

NewSpotCreateOrderCancelReplaceV3DataNewOrderResponseWithDefaults instantiates a new SpotCreateOrderCancelReplaceV3DataNewOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetClientOrderId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFills

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetFills() []SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetFillsOk() (*[]SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetFills(v []SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner)`

SetFills sets Fills field to given value.

### HasFills

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasFills() bool`

HasFills returns a boolean if a field has been set.

### GetOrderId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrigQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigQuoteOrderQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrigQuoteOrderQty() string`

GetOrigQuoteOrderQty returns the OrigQuoteOrderQty field if non-nil, zero value otherwise.

### GetOrigQuoteOrderQtyOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetOrigQuoteOrderQtyOk() (*string, bool)`

GetOrigQuoteOrderQtyOk returns a tuple with the OrigQuoteOrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQuoteOrderQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetOrigQuoteOrderQty(v string)`

SetOrigQuoteOrderQty sets OrigQuoteOrderQty field to given value.

### HasOrigQuoteOrderQty

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasOrigQuoteOrderQty() bool`

HasOrigQuoteOrderQty returns a boolean if a field has been set.

### GetPrice

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTransactTime

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetType

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkingTime

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetWorkingTime() int64`

GetWorkingTime returns the WorkingTime field if non-nil, zero value otherwise.

### GetWorkingTimeOk

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) GetWorkingTimeOk() (*int64, bool)`

GetWorkingTimeOk returns a tuple with the WorkingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTime

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) SetWorkingTime(v int64)`

SetWorkingTime sets WorkingTime field to given value.

### HasWorkingTime

`func (o *SpotCreateOrderCancelReplaceV3DataNewOrderResponse) HasWorkingTime() bool`

HasWorkingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


