# GetUmAccountV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetCmAccountV1RespAssetsInner**](GetCmAccountV1RespAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]GetUmAccountV2RespPositionsInner**](GetUmAccountV2RespPositionsInner.md) |  | [optional] 

## Methods

### NewGetUmAccountV2Resp

`func NewGetUmAccountV2Resp() *GetUmAccountV2Resp`

NewGetUmAccountV2Resp instantiates a new GetUmAccountV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmAccountV2RespWithDefaults

`func NewGetUmAccountV2RespWithDefaults() *GetUmAccountV2Resp`

NewGetUmAccountV2RespWithDefaults instantiates a new GetUmAccountV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetUmAccountV2Resp) GetAssets() []GetCmAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetUmAccountV2Resp) GetAssetsOk() (*[]GetCmAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetUmAccountV2Resp) SetAssets(v []GetCmAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetUmAccountV2Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *GetUmAccountV2Resp) GetPositions() []GetUmAccountV2RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetUmAccountV2Resp) GetPositionsOk() (*[]GetUmAccountV2RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetUmAccountV2Resp) SetPositions(v []GetUmAccountV2RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetUmAccountV2Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


