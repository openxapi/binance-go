# SubaccountGetSubAccountMarginAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**MarginTradeCoeffVo** | Pointer to [**SubaccountGetSubAccountMarginAccountV1RespMarginTradeCoeffVo**](SubaccountGetSubAccountMarginAccountV1RespMarginTradeCoeffVo.md) |  | [optional] 
**MarginUserAssetVoList** | Pointer to [**[]SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner**](SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner.md) |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 

## Methods

### NewSubaccountGetSubAccountMarginAccountV1Resp

`func NewSubaccountGetSubAccountMarginAccountV1Resp() *SubaccountGetSubAccountMarginAccountV1Resp`

NewSubaccountGetSubAccountMarginAccountV1Resp instantiates a new SubaccountGetSubAccountMarginAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubaccountGetSubAccountMarginAccountV1RespWithDefaults

`func NewSubaccountGetSubAccountMarginAccountV1RespWithDefaults() *SubaccountGetSubAccountMarginAccountV1Resp`

NewSubaccountGetSubAccountMarginAccountV1RespWithDefaults instantiates a new SubaccountGetSubAccountMarginAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMarginLevel

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetMarginTradeCoeffVo

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetMarginTradeCoeffVo() SubaccountGetSubAccountMarginAccountV1RespMarginTradeCoeffVo`

GetMarginTradeCoeffVo returns the MarginTradeCoeffVo field if non-nil, zero value otherwise.

### GetMarginTradeCoeffVoOk

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetMarginTradeCoeffVoOk() (*SubaccountGetSubAccountMarginAccountV1RespMarginTradeCoeffVo, bool)`

GetMarginTradeCoeffVoOk returns a tuple with the MarginTradeCoeffVo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginTradeCoeffVo

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) SetMarginTradeCoeffVo(v SubaccountGetSubAccountMarginAccountV1RespMarginTradeCoeffVo)`

SetMarginTradeCoeffVo sets MarginTradeCoeffVo field to given value.

### HasMarginTradeCoeffVo

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) HasMarginTradeCoeffVo() bool`

HasMarginTradeCoeffVo returns a boolean if a field has been set.

### GetMarginUserAssetVoList

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetMarginUserAssetVoList() []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner`

GetMarginUserAssetVoList returns the MarginUserAssetVoList field if non-nil, zero value otherwise.

### GetMarginUserAssetVoListOk

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetMarginUserAssetVoListOk() (*[]SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner, bool)`

GetMarginUserAssetVoListOk returns a tuple with the MarginUserAssetVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginUserAssetVoList

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) SetMarginUserAssetVoList(v []SubaccountGetManagedSubaccountMarginAssetV1RespUserAssetsInner)`

SetMarginUserAssetVoList sets MarginUserAssetVoList field to given value.

### HasMarginUserAssetVoList

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) HasMarginUserAssetVoList() bool`

HasMarginUserAssetVoList returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *SubaccountGetSubAccountMarginAccountV1Resp) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


