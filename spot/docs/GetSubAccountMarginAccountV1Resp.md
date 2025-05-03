# GetSubAccountMarginAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**MarginTradeCoeffVo** | Pointer to [**GetSubAccountMarginAccountV1RespMarginTradeCoeffVo**](GetSubAccountMarginAccountV1RespMarginTradeCoeffVo.md) |  | [optional] 
**MarginUserAssetVoList** | Pointer to [**[]GetManagedSubaccountMarginAssetV1RespUserAssetsInner**](GetManagedSubaccountMarginAssetV1RespUserAssetsInner.md) |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSubAccountMarginAccountV1Resp

`func NewGetSubAccountMarginAccountV1Resp() *GetSubAccountMarginAccountV1Resp`

NewGetSubAccountMarginAccountV1Resp instantiates a new GetSubAccountMarginAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubAccountMarginAccountV1RespWithDefaults

`func NewGetSubAccountMarginAccountV1RespWithDefaults() *GetSubAccountMarginAccountV1Resp`

NewGetSubAccountMarginAccountV1RespWithDefaults instantiates a new GetSubAccountMarginAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetSubAccountMarginAccountV1Resp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetSubAccountMarginAccountV1Resp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetSubAccountMarginAccountV1Resp) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetSubAccountMarginAccountV1Resp) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMarginLevel

`func (o *GetSubAccountMarginAccountV1Resp) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *GetSubAccountMarginAccountV1Resp) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *GetSubAccountMarginAccountV1Resp) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *GetSubAccountMarginAccountV1Resp) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetMarginTradeCoeffVo

`func (o *GetSubAccountMarginAccountV1Resp) GetMarginTradeCoeffVo() GetSubAccountMarginAccountV1RespMarginTradeCoeffVo`

GetMarginTradeCoeffVo returns the MarginTradeCoeffVo field if non-nil, zero value otherwise.

### GetMarginTradeCoeffVoOk

`func (o *GetSubAccountMarginAccountV1Resp) GetMarginTradeCoeffVoOk() (*GetSubAccountMarginAccountV1RespMarginTradeCoeffVo, bool)`

GetMarginTradeCoeffVoOk returns a tuple with the MarginTradeCoeffVo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginTradeCoeffVo

`func (o *GetSubAccountMarginAccountV1Resp) SetMarginTradeCoeffVo(v GetSubAccountMarginAccountV1RespMarginTradeCoeffVo)`

SetMarginTradeCoeffVo sets MarginTradeCoeffVo field to given value.

### HasMarginTradeCoeffVo

`func (o *GetSubAccountMarginAccountV1Resp) HasMarginTradeCoeffVo() bool`

HasMarginTradeCoeffVo returns a boolean if a field has been set.

### GetMarginUserAssetVoList

`func (o *GetSubAccountMarginAccountV1Resp) GetMarginUserAssetVoList() []GetManagedSubaccountMarginAssetV1RespUserAssetsInner`

GetMarginUserAssetVoList returns the MarginUserAssetVoList field if non-nil, zero value otherwise.

### GetMarginUserAssetVoListOk

`func (o *GetSubAccountMarginAccountV1Resp) GetMarginUserAssetVoListOk() (*[]GetManagedSubaccountMarginAssetV1RespUserAssetsInner, bool)`

GetMarginUserAssetVoListOk returns a tuple with the MarginUserAssetVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginUserAssetVoList

`func (o *GetSubAccountMarginAccountV1Resp) SetMarginUserAssetVoList(v []GetManagedSubaccountMarginAssetV1RespUserAssetsInner)`

SetMarginUserAssetVoList sets MarginUserAssetVoList field to given value.

### HasMarginUserAssetVoList

`func (o *GetSubAccountMarginAccountV1Resp) HasMarginUserAssetVoList() bool`

HasMarginUserAssetVoList returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *GetSubAccountMarginAccountV1Resp) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *GetSubAccountMarginAccountV1Resp) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *GetSubAccountMarginAccountV1Resp) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *GetSubAccountMarginAccountV1Resp) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


