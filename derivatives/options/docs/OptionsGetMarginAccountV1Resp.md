# OptionsGetMarginAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]OptionsGetMarginAccountV1RespAssetInner**](OptionsGetMarginAccountV1RespAssetInner.md) |  | [optional] 
**Greek** | Pointer to [**[]OptionsGetAccountV1RespGreekInner**](OptionsGetAccountV1RespGreekInner.md) |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewOptionsGetMarginAccountV1Resp

`func NewOptionsGetMarginAccountV1Resp() *OptionsGetMarginAccountV1Resp`

NewOptionsGetMarginAccountV1Resp instantiates a new OptionsGetMarginAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetMarginAccountV1RespWithDefaults

`func NewOptionsGetMarginAccountV1RespWithDefaults() *OptionsGetMarginAccountV1Resp`

NewOptionsGetMarginAccountV1RespWithDefaults instantiates a new OptionsGetMarginAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *OptionsGetMarginAccountV1Resp) GetAsset() []OptionsGetMarginAccountV1RespAssetInner`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *OptionsGetMarginAccountV1Resp) GetAssetOk() (*[]OptionsGetMarginAccountV1RespAssetInner, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *OptionsGetMarginAccountV1Resp) SetAsset(v []OptionsGetMarginAccountV1RespAssetInner)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *OptionsGetMarginAccountV1Resp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetGreek

`func (o *OptionsGetMarginAccountV1Resp) GetGreek() []OptionsGetAccountV1RespGreekInner`

GetGreek returns the Greek field if non-nil, zero value otherwise.

### GetGreekOk

`func (o *OptionsGetMarginAccountV1Resp) GetGreekOk() (*[]OptionsGetAccountV1RespGreekInner, bool)`

GetGreekOk returns a tuple with the Greek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreek

`func (o *OptionsGetMarginAccountV1Resp) SetGreek(v []OptionsGetAccountV1RespGreekInner)`

SetGreek sets Greek field to given value.

### HasGreek

`func (o *OptionsGetMarginAccountV1Resp) HasGreek() bool`

HasGreek returns a boolean if a field has been set.

### GetTime

`func (o *OptionsGetMarginAccountV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OptionsGetMarginAccountV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OptionsGetMarginAccountV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OptionsGetMarginAccountV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


