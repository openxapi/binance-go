# GetDciProductPositionsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetDciProductPositionsV1RespListInner**](GetDciProductPositionsV1RespListInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetDciProductPositionsV1Resp

`func NewGetDciProductPositionsV1Resp() *GetDciProductPositionsV1Resp`

NewGetDciProductPositionsV1Resp instantiates a new GetDciProductPositionsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDciProductPositionsV1RespWithDefaults

`func NewGetDciProductPositionsV1RespWithDefaults() *GetDciProductPositionsV1Resp`

NewGetDciProductPositionsV1RespWithDefaults instantiates a new GetDciProductPositionsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetDciProductPositionsV1Resp) GetList() []GetDciProductPositionsV1RespListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetDciProductPositionsV1Resp) GetListOk() (*[]GetDciProductPositionsV1RespListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetDciProductPositionsV1Resp) SetList(v []GetDciProductPositionsV1RespListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetDciProductPositionsV1Resp) HasList() bool`

HasList returns a boolean if a field has been set.

### GetTotal

`func (o *GetDciProductPositionsV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetDciProductPositionsV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetDciProductPositionsV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetDciProductPositionsV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


