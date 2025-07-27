# GetAssetDribbletV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**UserAssetDribblets** | Pointer to [**[]GetAssetDribbletV1RespUserAssetDribbletsInner**](GetAssetDribbletV1RespUserAssetDribbletsInner.md) |  | [optional] 

## Methods

### NewGetAssetDribbletV1Resp

`func NewGetAssetDribbletV1Resp() *GetAssetDribbletV1Resp`

NewGetAssetDribbletV1Resp instantiates a new GetAssetDribbletV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDribbletV1RespWithDefaults

`func NewGetAssetDribbletV1RespWithDefaults() *GetAssetDribbletV1Resp`

NewGetAssetDribbletV1RespWithDefaults instantiates a new GetAssetDribbletV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetAssetDribbletV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAssetDribbletV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAssetDribbletV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAssetDribbletV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUserAssetDribblets

`func (o *GetAssetDribbletV1Resp) GetUserAssetDribblets() []GetAssetDribbletV1RespUserAssetDribbletsInner`

GetUserAssetDribblets returns the UserAssetDribblets field if non-nil, zero value otherwise.

### GetUserAssetDribbletsOk

`func (o *GetAssetDribbletV1Resp) GetUserAssetDribbletsOk() (*[]GetAssetDribbletV1RespUserAssetDribbletsInner, bool)`

GetUserAssetDribbletsOk returns a tuple with the UserAssetDribblets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssetDribblets

`func (o *GetAssetDribbletV1Resp) SetUserAssetDribblets(v []GetAssetDribbletV1RespUserAssetDribbletsInner)`

SetUserAssetDribblets sets UserAssetDribblets field to given value.

### HasUserAssetDribblets

`func (o *GetAssetDribbletV1Resp) HasUserAssetDribblets() bool`

HasUserAssetDribblets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


