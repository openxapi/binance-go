# AlgoGetAlgoFuturesHistoricalOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]AlgoGetAlgoFuturesHistoricalOrdersV1RespOrdersInner**](AlgoGetAlgoFuturesHistoricalOrdersV1RespOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlgoGetAlgoFuturesHistoricalOrdersV1Resp

`func NewAlgoGetAlgoFuturesHistoricalOrdersV1Resp() *AlgoGetAlgoFuturesHistoricalOrdersV1Resp`

NewAlgoGetAlgoFuturesHistoricalOrdersV1Resp instantiates a new AlgoGetAlgoFuturesHistoricalOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoGetAlgoFuturesHistoricalOrdersV1RespWithDefaults

`func NewAlgoGetAlgoFuturesHistoricalOrdersV1RespWithDefaults() *AlgoGetAlgoFuturesHistoricalOrdersV1Resp`

NewAlgoGetAlgoFuturesHistoricalOrdersV1RespWithDefaults instantiates a new AlgoGetAlgoFuturesHistoricalOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) GetOrders() []AlgoGetAlgoFuturesHistoricalOrdersV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) GetOrdersOk() (*[]AlgoGetAlgoFuturesHistoricalOrdersV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) SetOrders(v []AlgoGetAlgoFuturesHistoricalOrdersV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTotal

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AlgoGetAlgoFuturesHistoricalOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


