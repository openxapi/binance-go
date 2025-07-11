# GetManagedSubaccountInfoV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagerSubUserInfoVoList** | Pointer to [**[]GetManagedSubaccountInfoV1RespManagerSubUserInfoVoListInner**](GetManagedSubaccountInfoV1RespManagerSubUserInfoVoListInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetManagedSubaccountInfoV1Resp

`func NewGetManagedSubaccountInfoV1Resp() *GetManagedSubaccountInfoV1Resp`

NewGetManagedSubaccountInfoV1Resp instantiates a new GetManagedSubaccountInfoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedSubaccountInfoV1RespWithDefaults

`func NewGetManagedSubaccountInfoV1RespWithDefaults() *GetManagedSubaccountInfoV1Resp`

NewGetManagedSubaccountInfoV1RespWithDefaults instantiates a new GetManagedSubaccountInfoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagerSubUserInfoVoList

`func (o *GetManagedSubaccountInfoV1Resp) GetManagerSubUserInfoVoList() []GetManagedSubaccountInfoV1RespManagerSubUserInfoVoListInner`

GetManagerSubUserInfoVoList returns the ManagerSubUserInfoVoList field if non-nil, zero value otherwise.

### GetManagerSubUserInfoVoListOk

`func (o *GetManagedSubaccountInfoV1Resp) GetManagerSubUserInfoVoListOk() (*[]GetManagedSubaccountInfoV1RespManagerSubUserInfoVoListInner, bool)`

GetManagerSubUserInfoVoListOk returns a tuple with the ManagerSubUserInfoVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerSubUserInfoVoList

`func (o *GetManagedSubaccountInfoV1Resp) SetManagerSubUserInfoVoList(v []GetManagedSubaccountInfoV1RespManagerSubUserInfoVoListInner)`

SetManagerSubUserInfoVoList sets ManagerSubUserInfoVoList field to given value.

### HasManagerSubUserInfoVoList

`func (o *GetManagedSubaccountInfoV1Resp) HasManagerSubUserInfoVoList() bool`

HasManagerSubUserInfoVoList returns a boolean if a field has been set.

### GetTotal

`func (o *GetManagedSubaccountInfoV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetManagedSubaccountInfoV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetManagedSubaccountInfoV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetManagedSubaccountInfoV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


