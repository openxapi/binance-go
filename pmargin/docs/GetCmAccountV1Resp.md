# GetCmAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetCmAccountV1RespAssetsInner**](GetCmAccountV1RespAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]GetCmAccountV1RespPositionsInner**](GetCmAccountV1RespPositionsInner.md) |  | [optional] 

## Methods

### NewGetCmAccountV1Resp

`func NewGetCmAccountV1Resp() *GetCmAccountV1Resp`

NewGetCmAccountV1Resp instantiates a new GetCmAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCmAccountV1RespWithDefaults

`func NewGetCmAccountV1RespWithDefaults() *GetCmAccountV1Resp`

NewGetCmAccountV1RespWithDefaults instantiates a new GetCmAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetCmAccountV1Resp) GetAssets() []GetCmAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetCmAccountV1Resp) GetAssetsOk() (*[]GetCmAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetCmAccountV1Resp) SetAssets(v []GetCmAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetCmAccountV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *GetCmAccountV1Resp) GetPositions() []GetCmAccountV1RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetCmAccountV1Resp) GetPositionsOk() (*[]GetCmAccountV1RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetCmAccountV1Resp) SetPositions(v []GetCmAccountV1RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetCmAccountV1Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


