# AlgoGetAlgoSpotOpenOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner**](AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlgoGetAlgoSpotOpenOrdersV1Resp

`func NewAlgoGetAlgoSpotOpenOrdersV1Resp() *AlgoGetAlgoSpotOpenOrdersV1Resp`

NewAlgoGetAlgoSpotOpenOrdersV1Resp instantiates a new AlgoGetAlgoSpotOpenOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoGetAlgoSpotOpenOrdersV1RespWithDefaults

`func NewAlgoGetAlgoSpotOpenOrdersV1RespWithDefaults() *AlgoGetAlgoSpotOpenOrdersV1Resp`

NewAlgoGetAlgoSpotOpenOrdersV1RespWithDefaults instantiates a new AlgoGetAlgoSpotOpenOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) GetOrders() []AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) GetOrdersOk() (*[]AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) SetOrders(v []AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTotal

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AlgoGetAlgoSpotOpenOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


