# GetAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]GetAccountV1RespAssetInner**](GetAccountV1RespAssetInner.md) |  | [optional] 
**Greek** | Pointer to [**[]GetAccountV1RespGreekInner**](GetAccountV1RespGreekInner.md) |  | [optional] 
**RiskLevel** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAccountV1Resp

`func NewGetAccountV1Resp() *GetAccountV1Resp`

NewGetAccountV1Resp instantiates a new GetAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountV1RespWithDefaults

`func NewGetAccountV1RespWithDefaults() *GetAccountV1Resp`

NewGetAccountV1RespWithDefaults instantiates a new GetAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *GetAccountV1Resp) GetAsset() []GetAccountV1RespAssetInner`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetAccountV1Resp) GetAssetOk() (*[]GetAccountV1RespAssetInner, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetAccountV1Resp) SetAsset(v []GetAccountV1RespAssetInner)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetAccountV1Resp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetGreek

`func (o *GetAccountV1Resp) GetGreek() []GetAccountV1RespGreekInner`

GetGreek returns the Greek field if non-nil, zero value otherwise.

### GetGreekOk

`func (o *GetAccountV1Resp) GetGreekOk() (*[]GetAccountV1RespGreekInner, bool)`

GetGreekOk returns a tuple with the Greek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreek

`func (o *GetAccountV1Resp) SetGreek(v []GetAccountV1RespGreekInner)`

SetGreek sets Greek field to given value.

### HasGreek

`func (o *GetAccountV1Resp) HasGreek() bool`

HasGreek returns a boolean if a field has been set.

### GetRiskLevel

`func (o *GetAccountV1Resp) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *GetAccountV1Resp) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *GetAccountV1Resp) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *GetAccountV1Resp) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetTime

`func (o *GetAccountV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetAccountV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetAccountV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetAccountV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


