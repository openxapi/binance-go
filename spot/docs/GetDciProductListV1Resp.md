# GetDciProductListV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetDciProductListV1RespListInner**](GetDciProductListV1RespListInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetDciProductListV1Resp

`func NewGetDciProductListV1Resp() *GetDciProductListV1Resp`

NewGetDciProductListV1Resp instantiates a new GetDciProductListV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDciProductListV1RespWithDefaults

`func NewGetDciProductListV1RespWithDefaults() *GetDciProductListV1Resp`

NewGetDciProductListV1RespWithDefaults instantiates a new GetDciProductListV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetDciProductListV1Resp) GetList() []GetDciProductListV1RespListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetDciProductListV1Resp) GetListOk() (*[]GetDciProductListV1RespListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetDciProductListV1Resp) SetList(v []GetDciProductListV1RespListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetDciProductListV1Resp) HasList() bool`

HasList returns a boolean if a field has been set.

### GetTotal

`func (o *GetDciProductListV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetDciProductListV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetDciProductListV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetDciProductListV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


