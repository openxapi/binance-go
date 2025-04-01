# AlgoGetAlgoFuturesSubOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutedAmt** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**SubOrders** | Pointer to [**[]AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner**](AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlgoGetAlgoFuturesSubOrdersV1Resp

`func NewAlgoGetAlgoFuturesSubOrdersV1Resp() *AlgoGetAlgoFuturesSubOrdersV1Resp`

NewAlgoGetAlgoFuturesSubOrdersV1Resp instantiates a new AlgoGetAlgoFuturesSubOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoGetAlgoFuturesSubOrdersV1RespWithDefaults

`func NewAlgoGetAlgoFuturesSubOrdersV1RespWithDefaults() *AlgoGetAlgoFuturesSubOrdersV1Resp`

NewAlgoGetAlgoFuturesSubOrdersV1RespWithDefaults instantiates a new AlgoGetAlgoFuturesSubOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutedAmt

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetExecutedAmt() string`

GetExecutedAmt returns the ExecutedAmt field if non-nil, zero value otherwise.

### GetExecutedAmtOk

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetExecutedAmtOk() (*string, bool)`

GetExecutedAmtOk returns a tuple with the ExecutedAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAmt

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) SetExecutedAmt(v string)`

SetExecutedAmt sets ExecutedAmt field to given value.

### HasExecutedAmt

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) HasExecutedAmt() bool`

HasExecutedAmt returns a boolean if a field has been set.

### GetExecutedQty

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetSubOrders

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetSubOrders() []AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner`

GetSubOrders returns the SubOrders field if non-nil, zero value otherwise.

### GetSubOrdersOk

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetSubOrdersOk() (*[]AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner, bool)`

GetSubOrdersOk returns a tuple with the SubOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrders

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) SetSubOrders(v []AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner)`

SetSubOrders sets SubOrders field to given value.

### HasSubOrders

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) HasSubOrders() bool`

HasSubOrders returns a boolean if a field has been set.

### GetTotal

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AlgoGetAlgoFuturesSubOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


