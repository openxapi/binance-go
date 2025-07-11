# GetAlgoSpotHistoricalOrdersV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]GetAlgoSpotHistoricalOrdersV1RespOrdersInner**](GetAlgoSpotHistoricalOrdersV1RespOrdersInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetAlgoSpotHistoricalOrdersV1Resp

`func NewGetAlgoSpotHistoricalOrdersV1Resp() *GetAlgoSpotHistoricalOrdersV1Resp`

NewGetAlgoSpotHistoricalOrdersV1Resp instantiates a new GetAlgoSpotHistoricalOrdersV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlgoSpotHistoricalOrdersV1RespWithDefaults

`func NewGetAlgoSpotHistoricalOrdersV1RespWithDefaults() *GetAlgoSpotHistoricalOrdersV1Resp`

NewGetAlgoSpotHistoricalOrdersV1RespWithDefaults instantiates a new GetAlgoSpotHistoricalOrdersV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetOrders() []GetAlgoSpotHistoricalOrdersV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetOrdersOk() (*[]GetAlgoSpotHistoricalOrdersV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) SetOrders(v []GetAlgoSpotHistoricalOrdersV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTotal

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAlgoSpotHistoricalOrdersV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


