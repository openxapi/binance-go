# PmarginGetUmAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]PmarginGetCmAccountV1RespAssetsInner**](PmarginGetCmAccountV1RespAssetsInner.md) |  | [optional] 
**Positions** | Pointer to [**[]PmarginGetUmAccountV1RespPositionsInner**](PmarginGetUmAccountV1RespPositionsInner.md) |  | [optional] 

## Methods

### NewPmarginGetUmAccountV1Resp

`func NewPmarginGetUmAccountV1Resp() *PmarginGetUmAccountV1Resp`

NewPmarginGetUmAccountV1Resp instantiates a new PmarginGetUmAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetUmAccountV1RespWithDefaults

`func NewPmarginGetUmAccountV1RespWithDefaults() *PmarginGetUmAccountV1Resp`

NewPmarginGetUmAccountV1RespWithDefaults instantiates a new PmarginGetUmAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *PmarginGetUmAccountV1Resp) GetAssets() []PmarginGetCmAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *PmarginGetUmAccountV1Resp) GetAssetsOk() (*[]PmarginGetCmAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *PmarginGetUmAccountV1Resp) SetAssets(v []PmarginGetCmAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *PmarginGetUmAccountV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetPositions

`func (o *PmarginGetUmAccountV1Resp) GetPositions() []PmarginGetUmAccountV1RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *PmarginGetUmAccountV1Resp) GetPositionsOk() (*[]PmarginGetUmAccountV1RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *PmarginGetUmAccountV1Resp) SetPositions(v []PmarginGetUmAccountV1RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *PmarginGetUmAccountV1Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


