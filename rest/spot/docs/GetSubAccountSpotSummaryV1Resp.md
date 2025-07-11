# GetSubAccountSpotSummaryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterAccountTotalAsset** | Pointer to **string** |  | [optional] 
**SpotSubUserAssetBtcVoList** | Pointer to [**[]GetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner**](GetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetSubAccountSpotSummaryV1Resp

`func NewGetSubAccountSpotSummaryV1Resp() *GetSubAccountSpotSummaryV1Resp`

NewGetSubAccountSpotSummaryV1Resp instantiates a new GetSubAccountSpotSummaryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubAccountSpotSummaryV1RespWithDefaults

`func NewGetSubAccountSpotSummaryV1RespWithDefaults() *GetSubAccountSpotSummaryV1Resp`

NewGetSubAccountSpotSummaryV1RespWithDefaults instantiates a new GetSubAccountSpotSummaryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterAccountTotalAsset

`func (o *GetSubAccountSpotSummaryV1Resp) GetMasterAccountTotalAsset() string`

GetMasterAccountTotalAsset returns the MasterAccountTotalAsset field if non-nil, zero value otherwise.

### GetMasterAccountTotalAssetOk

`func (o *GetSubAccountSpotSummaryV1Resp) GetMasterAccountTotalAssetOk() (*string, bool)`

GetMasterAccountTotalAssetOk returns a tuple with the MasterAccountTotalAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccountTotalAsset

`func (o *GetSubAccountSpotSummaryV1Resp) SetMasterAccountTotalAsset(v string)`

SetMasterAccountTotalAsset sets MasterAccountTotalAsset field to given value.

### HasMasterAccountTotalAsset

`func (o *GetSubAccountSpotSummaryV1Resp) HasMasterAccountTotalAsset() bool`

HasMasterAccountTotalAsset returns a boolean if a field has been set.

### GetSpotSubUserAssetBtcVoList

`func (o *GetSubAccountSpotSummaryV1Resp) GetSpotSubUserAssetBtcVoList() []GetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner`

GetSpotSubUserAssetBtcVoList returns the SpotSubUserAssetBtcVoList field if non-nil, zero value otherwise.

### GetSpotSubUserAssetBtcVoListOk

`func (o *GetSubAccountSpotSummaryV1Resp) GetSpotSubUserAssetBtcVoListOk() (*[]GetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner, bool)`

GetSpotSubUserAssetBtcVoListOk returns a tuple with the SpotSubUserAssetBtcVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotSubUserAssetBtcVoList

`func (o *GetSubAccountSpotSummaryV1Resp) SetSpotSubUserAssetBtcVoList(v []GetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner)`

SetSpotSubUserAssetBtcVoList sets SpotSubUserAssetBtcVoList field to given value.

### HasSpotSubUserAssetBtcVoList

`func (o *GetSubAccountSpotSummaryV1Resp) HasSpotSubUserAssetBtcVoList() bool`

HasSpotSubUserAssetBtcVoList returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetSubAccountSpotSummaryV1Resp) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetSubAccountSpotSummaryV1Resp) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetSubAccountSpotSummaryV1Resp) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetSubAccountSpotSummaryV1Resp) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


