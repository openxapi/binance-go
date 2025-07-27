# GetUmAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetCmAccountV1RespAssetsInner**](GetCmAccountV1RespAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]GetUmAccountV1RespPositionsInner**](GetUmAccountV1RespPositionsInner.md) |  | [optional] 

## Methods

### NewGetUmAccountV1Resp

`func NewGetUmAccountV1Resp() *GetUmAccountV1Resp`

NewGetUmAccountV1Resp instantiates a new GetUmAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmAccountV1RespWithDefaults

`func NewGetUmAccountV1RespWithDefaults() *GetUmAccountV1Resp`

NewGetUmAccountV1RespWithDefaults instantiates a new GetUmAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *GetUmAccountV1Resp) GetAssets() []GetCmAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetUmAccountV1Resp) GetAssetsOk() (*[]GetCmAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetUmAccountV1Resp) SetAssets(v []GetCmAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetUmAccountV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *GetUmAccountV1Resp) GetPositions() []GetUmAccountV1RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetUmAccountV1Resp) GetPositionsOk() (*[]GetUmAccountV1RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetUmAccountV1Resp) SetPositions(v []GetUmAccountV1RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetUmAccountV1Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


