# AlgoGetAlgoSpotHistoricalOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner**](AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewAlgoGetAlgoSpotHistoricalOrdersV1Resp

`func NewAlgoGetAlgoSpotHistoricalOrdersV1Resp() *AlgoGetAlgoSpotHistoricalOrdersV1Resp`

NewAlgoGetAlgoSpotHistoricalOrdersV1Resp instantiates a new AlgoGetAlgoSpotHistoricalOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoGetAlgoSpotHistoricalOrdersV1RespWithDefaults

`func NewAlgoGetAlgoSpotHistoricalOrdersV1RespWithDefaults() *AlgoGetAlgoSpotHistoricalOrdersV1Resp`

NewAlgoGetAlgoSpotHistoricalOrdersV1RespWithDefaults instantiates a new AlgoGetAlgoSpotHistoricalOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) GetOrders() []AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) GetOrdersOk() (*[]AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) SetOrders(v []AlgoGetAlgoSpotHistoricalOrdersV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTotal

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AlgoGetAlgoSpotHistoricalOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


