# GetAlgoFuturesOpenOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]GetAlgoFuturesHistoricalOrdersV1RespOrdersInner**](GetAlgoFuturesHistoricalOrdersV1RespOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAlgoFuturesOpenOrdersV1Resp

`func NewGetAlgoFuturesOpenOrdersV1Resp() *GetAlgoFuturesOpenOrdersV1Resp`

NewGetAlgoFuturesOpenOrdersV1Resp instantiates a new GetAlgoFuturesOpenOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlgoFuturesOpenOrdersV1RespWithDefaults

`func NewGetAlgoFuturesOpenOrdersV1RespWithDefaults() *GetAlgoFuturesOpenOrdersV1Resp`

NewGetAlgoFuturesOpenOrdersV1RespWithDefaults instantiates a new GetAlgoFuturesOpenOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *GetAlgoFuturesOpenOrdersV1Resp) GetOrders() []GetAlgoFuturesHistoricalOrdersV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetAlgoFuturesOpenOrdersV1Resp) GetOrdersOk() (*[]GetAlgoFuturesHistoricalOrdersV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetAlgoFuturesOpenOrdersV1Resp) SetOrders(v []GetAlgoFuturesHistoricalOrdersV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GetAlgoFuturesOpenOrdersV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTotal

`func (o *GetAlgoFuturesOpenOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAlgoFuturesOpenOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAlgoFuturesOpenOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAlgoFuturesOpenOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


