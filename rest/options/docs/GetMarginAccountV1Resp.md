# GetMarginAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]GetMarginAccountV1RespAssetInner**](GetMarginAccountV1RespAssetInner.md) |  | [optional] 
**Greek** | Pointer to [**[]GetAccountV1RespGreekInner**](GetAccountV1RespGreekInner.md) |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetMarginAccountV1Resp

`func NewGetMarginAccountV1Resp() *GetMarginAccountV1Resp`

NewGetMarginAccountV1Resp instantiates a new GetMarginAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarginAccountV1RespWithDefaults

`func NewGetMarginAccountV1RespWithDefaults() *GetMarginAccountV1Resp`

NewGetMarginAccountV1RespWithDefaults instantiates a new GetMarginAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *GetMarginAccountV1Resp) GetAsset() []GetMarginAccountV1RespAssetInner`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetMarginAccountV1Resp) GetAssetOk() (*[]GetMarginAccountV1RespAssetInner, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetMarginAccountV1Resp) SetAsset(v []GetMarginAccountV1RespAssetInner)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetMarginAccountV1Resp) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetGreek

`func (o *GetMarginAccountV1Resp) GetGreek() []GetAccountV1RespGreekInner`

GetGreek returns the Greek field if non-nil, zero value otherwise.

### GetGreekOk

`func (o *GetMarginAccountV1Resp) GetGreekOk() (*[]GetAccountV1RespGreekInner, bool)`

GetGreekOk returns a tuple with the Greek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreek

`func (o *GetMarginAccountV1Resp) SetGreek(v []GetAccountV1RespGreekInner)`

SetGreek sets Greek field to given value.

### HasGreek

`func (o *GetMarginAccountV1Resp) HasGreek() bool`

HasGreek returns a boolean if a field has been set.

### GetTime

`func (o *GetMarginAccountV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetMarginAccountV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetMarginAccountV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetMarginAccountV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


