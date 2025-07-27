# GetAlgoSpotSubOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutedAmt** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**SubOrders** | Pointer to [**[]GetAlgoFuturesSubOrdersV1RespSubOrdersInner**](GetAlgoFuturesSubOrdersV1RespSubOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAlgoSpotSubOrdersV1Resp

`func NewGetAlgoSpotSubOrdersV1Resp() *GetAlgoSpotSubOrdersV1Resp`

NewGetAlgoSpotSubOrdersV1Resp instantiates a new GetAlgoSpotSubOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlgoSpotSubOrdersV1RespWithDefaults

`func NewGetAlgoSpotSubOrdersV1RespWithDefaults() *GetAlgoSpotSubOrdersV1Resp`

NewGetAlgoSpotSubOrdersV1RespWithDefaults instantiates a new GetAlgoSpotSubOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutedAmt

`func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedAmt() string`

GetExecutedAmt returns the ExecutedAmt field if non-nil, zero value otherwise.

### GetExecutedAmtOk

`func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedAmtOk() (*string, bool)`

GetExecutedAmtOk returns a tuple with the ExecutedAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAmt

`func (o *GetAlgoSpotSubOrdersV1Resp) SetExecutedAmt(v string)`

SetExecutedAmt sets ExecutedAmt field to given value.

### HasExecutedAmt

`func (o *GetAlgoSpotSubOrdersV1Resp) HasExecutedAmt() bool`

HasExecutedAmt returns a boolean if a field has been set.

### GetExecutedQty

`func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *GetAlgoSpotSubOrdersV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *GetAlgoSpotSubOrdersV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetSubOrders

`func (o *GetAlgoSpotSubOrdersV1Resp) GetSubOrders() []GetAlgoFuturesSubOrdersV1RespSubOrdersInner`

GetSubOrders returns the SubOrders field if non-nil, zero value otherwise.

### GetSubOrdersOk

`func (o *GetAlgoSpotSubOrdersV1Resp) GetSubOrdersOk() (*[]GetAlgoFuturesSubOrdersV1RespSubOrdersInner, bool)`

GetSubOrdersOk returns a tuple with the SubOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrders

`func (o *GetAlgoSpotSubOrdersV1Resp) SetSubOrders(v []GetAlgoFuturesSubOrdersV1RespSubOrdersInner)`

SetSubOrders sets SubOrders field to given value.

### HasSubOrders

`func (o *GetAlgoSpotSubOrdersV1Resp) HasSubOrders() bool`

HasSubOrders returns a boolean if a field has been set.

### GetTotal

`func (o *GetAlgoSpotSubOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAlgoSpotSubOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAlgoSpotSubOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAlgoSpotSubOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


