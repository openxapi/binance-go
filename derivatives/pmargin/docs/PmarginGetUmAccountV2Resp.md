# PmarginGetUmAccountV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]PmarginGetCmAccountV1RespAssetsInner**](PmarginGetCmAccountV1RespAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]PmarginGetUmAccountV2RespPositionsInner**](PmarginGetUmAccountV2RespPositionsInner.md) |  | [optional] 

## Methods

### NewPmarginGetUmAccountV2Resp

`func NewPmarginGetUmAccountV2Resp() *PmarginGetUmAccountV2Resp`

NewPmarginGetUmAccountV2Resp instantiates a new PmarginGetUmAccountV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetUmAccountV2RespWithDefaults

`func NewPmarginGetUmAccountV2RespWithDefaults() *PmarginGetUmAccountV2Resp`

NewPmarginGetUmAccountV2RespWithDefaults instantiates a new PmarginGetUmAccountV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *PmarginGetUmAccountV2Resp) GetAssets() []PmarginGetCmAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PmarginGetUmAccountV2Resp) GetAssetsOk() (*[]PmarginGetCmAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PmarginGetUmAccountV2Resp) SetAssets(v []PmarginGetCmAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *PmarginGetUmAccountV2Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *PmarginGetUmAccountV2Resp) GetPositions() []PmarginGetUmAccountV2RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *PmarginGetUmAccountV2Resp) GetPositionsOk() (*[]PmarginGetUmAccountV2RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *PmarginGetUmAccountV2Resp) SetPositions(v []PmarginGetUmAccountV2RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *PmarginGetUmAccountV2Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


