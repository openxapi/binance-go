# AlgoGetAlgoSpotSubOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutedAmt** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**SubOrders** | Pointer to [**[]AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner**](AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlgoGetAlgoSpotSubOrdersV1Resp

`func NewAlgoGetAlgoSpotSubOrdersV1Resp() *AlgoGetAlgoSpotSubOrdersV1Resp`

NewAlgoGetAlgoSpotSubOrdersV1Resp instantiates a new AlgoGetAlgoSpotSubOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoGetAlgoSpotSubOrdersV1RespWithDefaults

`func NewAlgoGetAlgoSpotSubOrdersV1RespWithDefaults() *AlgoGetAlgoSpotSubOrdersV1Resp`

NewAlgoGetAlgoSpotSubOrdersV1RespWithDefaults instantiates a new AlgoGetAlgoSpotSubOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutedAmt

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetExecutedAmt() string`

GetExecutedAmt returns the ExecutedAmt field if non-nil, zero value otherwise.

### GetExecutedAmtOk

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetExecutedAmtOk() (*string, bool)`

GetExecutedAmtOk returns a tuple with the ExecutedAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAmt

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) SetExecutedAmt(v string)`

SetExecutedAmt sets ExecutedAmt field to given value.

### HasExecutedAmt

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) HasExecutedAmt() bool`

HasExecutedAmt returns a boolean if a field has been set.

### GetExecutedQty

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetSubOrders

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetSubOrders() []AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner`

GetSubOrders returns the SubOrders field if non-nil, zero value otherwise.

### GetSubOrdersOk

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetSubOrdersOk() (*[]AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner, bool)`

GetSubOrdersOk returns a tuple with the SubOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrders

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) SetSubOrders(v []AlgoGetAlgoFuturesSubOrdersV1RespSubOrdersInner)`

SetSubOrders sets SubOrders field to given value.

### HasSubOrders

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) HasSubOrders() bool`

HasSubOrders returns a boolean if a field has been set.

### GetTotal

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AlgoGetAlgoSpotSubOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


