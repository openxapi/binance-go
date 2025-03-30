# OptionsGetAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]OptionsGetAccountV1RespAssetInner**](OptionsGetAccountV1RespAssetInner.md) |  | [optional] 
**Greek** | Pointer to [**[]OptionsGetAccountV1RespGreekInner**](OptionsGetAccountV1RespGreekInner.md) |  | [optional] 
**RiskLevel** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewOptionsGetAccountV1Resp

`func NewOptionsGetAccountV1Resp() *OptionsGetAccountV1Resp`

NewOptionsGetAccountV1Resp instantiates a new OptionsGetAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetAccountV1RespWithDefaults

`func NewOptionsGetAccountV1RespWithDefaults() *OptionsGetAccountV1Resp`

NewOptionsGetAccountV1RespWithDefaults instantiates a new OptionsGetAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *OptionsGetAccountV1Resp) GetAsset() []OptionsGetAccountV1RespAssetInner`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *OptionsGetAccountV1Resp) GetAssetOk() (*[]OptionsGetAccountV1RespAssetInner, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *OptionsGetAccountV1Resp) SetAsset(v []OptionsGetAccountV1RespAssetInner)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *OptionsGetAccountV1Resp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetGreek

`func (o *OptionsGetAccountV1Resp) GetGreek() []OptionsGetAccountV1RespGreekInner`

GetGreek returns the Greek field if non-nil, zero value otherwise.

### GetGreekOk

`func (o *OptionsGetAccountV1Resp) GetGreekOk() (*[]OptionsGetAccountV1RespGreekInner, bool)`

GetGreekOk returns a tuple with the Greek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreek

`func (o *OptionsGetAccountV1Resp) SetGreek(v []OptionsGetAccountV1RespGreekInner)`

SetGreek sets Greek field to given value.

### HasGreek

`func (o *OptionsGetAccountV1Resp) HasGreek() bool`

HasGreek returns a boolean if a field has been set.

### GetRiskLevel

`func (o *OptionsGetAccountV1Resp) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *OptionsGetAccountV1Resp) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *OptionsGetAccountV1Resp) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *OptionsGetAccountV1Resp) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetTime

`func (o *OptionsGetAccountV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OptionsGetAccountV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OptionsGetAccountV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OptionsGetAccountV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


