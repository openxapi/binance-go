# WalletGetLocalentityDepositHistoryV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AddressTag** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**ConfirmTimes** | Pointer to **string** |  | [optional] 
**DepositStatus** | Pointer to **int32** |  | [optional] 
**InsertTime** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Questionnaire** | Pointer to **map[string]interface{}** |  | [optional] 
**RequireQuestionnaire** | Pointer to **bool** |  | [optional] 
**TrId** | Pointer to **int64** |  | [optional] 
**TranId** | Pointer to **int64** |  | [optional] 
**TransferType** | Pointer to **int32** |  | [optional] 
**TravelRuleStatus** | Pointer to **int32** |  | [optional] 
**TxId** | Pointer to **string** |  | [optional] 
**UnlockConfirm** | Pointer to **int32** |  | [optional] 
**WalletType** | Pointer to **int32** |  | [optional] 

## Methods

### NewWalletGetLocalentityDepositHistoryV1RespItem

`func NewWalletGetLocalentityDepositHistoryV1RespItem() *WalletGetLocalentityDepositHistoryV1RespItem`

NewWalletGetLocalentityDepositHistoryV1RespItem instantiates a new WalletGetLocalentityDepositHistoryV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletGetLocalentityDepositHistoryV1RespItemWithDefaults

`func NewWalletGetLocalentityDepositHistoryV1RespItemWithDefaults() *WalletGetLocalentityDepositHistoryV1RespItem`

NewWalletGetLocalentityDepositHistoryV1RespItemWithDefaults instantiates a new WalletGetLocalentityDepositHistoryV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressTag

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetAddressTag() string`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetAddressTagOk() (*string, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetAddressTag(v string)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

### GetAmount

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCoin

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetConfirmTimes

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetConfirmTimes() string`

GetConfirmTimes returns the ConfirmTimes field if non-nil, zero value otherwise.

### GetConfirmTimesOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetConfirmTimesOk() (*string, bool)`

GetConfirmTimesOk returns a tuple with the ConfirmTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTimes

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetConfirmTimes(v string)`

SetConfirmTimes sets ConfirmTimes field to given value.

### HasConfirmTimes

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasConfirmTimes() bool`

HasConfirmTimes returns a boolean if a field has been set.

### GetDepositStatus

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetDepositStatus() int32`

GetDepositStatus returns the DepositStatus field if non-nil, zero value otherwise.

### GetDepositStatusOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetDepositStatusOk() (*int32, bool)`

GetDepositStatusOk returns a tuple with the DepositStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositStatus

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetDepositStatus(v int32)`

SetDepositStatus sets DepositStatus field to given value.

### HasDepositStatus

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasDepositStatus() bool`

HasDepositStatus returns a boolean if a field has been set.

### GetInsertTime

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetInsertTime() int64`

GetInsertTime returns the InsertTime field if non-nil, zero value otherwise.

### GetInsertTimeOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetInsertTimeOk() (*int64, bool)`

GetInsertTimeOk returns a tuple with the InsertTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertTime

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetInsertTime(v int64)`

SetInsertTime sets InsertTime field to given value.

### HasInsertTime

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasInsertTime() bool`

HasInsertTime returns a boolean if a field has been set.

### GetNetwork

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetQuestionnaire

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetQuestionnaire() map[string]interface{}`

GetQuestionnaire returns the Questionnaire field if non-nil, zero value otherwise.

### GetQuestionnaireOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetQuestionnaireOk() (*map[string]interface{}, bool)`

GetQuestionnaireOk returns a tuple with the Questionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaire

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetQuestionnaire(v map[string]interface{})`

SetQuestionnaire sets Questionnaire field to given value.

### HasQuestionnaire

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasQuestionnaire() bool`

HasQuestionnaire returns a boolean if a field has been set.

### SetQuestionnaireNil

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetQuestionnaireNil(b bool)`

 SetQuestionnaireNil sets the value for Questionnaire to be an explicit nil

### UnsetQuestionnaire
`func (o *WalletGetLocalentityDepositHistoryV1RespItem) UnsetQuestionnaire()`

UnsetQuestionnaire ensures that no value is present for Questionnaire, not even an explicit nil
### GetRequireQuestionnaire

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetRequireQuestionnaire() bool`

GetRequireQuestionnaire returns the RequireQuestionnaire field if non-nil, zero value otherwise.

### GetRequireQuestionnaireOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetRequireQuestionnaireOk() (*bool, bool)`

GetRequireQuestionnaireOk returns a tuple with the RequireQuestionnaire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireQuestionnaire

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetRequireQuestionnaire(v bool)`

SetRequireQuestionnaire sets RequireQuestionnaire field to given value.

### HasRequireQuestionnaire

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasRequireQuestionnaire() bool`

HasRequireQuestionnaire returns a boolean if a field has been set.

### GetTrId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTrId() int64`

GetTrId returns the TrId field if non-nil, zero value otherwise.

### GetTrIdOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTrIdOk() (*int64, bool)`

GetTrIdOk returns a tuple with the TrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetTrId(v int64)`

SetTrId sets TrId field to given value.

### HasTrId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasTrId() bool`

HasTrId returns a boolean if a field has been set.

### GetTranId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTranId() int64`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTranIdOk() (*int64, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetTranId(v int64)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTransferType

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTransferType() int32`

GetTransferType returns the TransferType field if non-nil, zero value otherwise.

### GetTransferTypeOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTransferTypeOk() (*int32, bool)`

GetTransferTypeOk returns a tuple with the TransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferType

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetTransferType(v int32)`

SetTransferType sets TransferType field to given value.

### HasTransferType

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasTransferType() bool`

HasTransferType returns a boolean if a field has been set.

### GetTravelRuleStatus

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTravelRuleStatus() int32`

GetTravelRuleStatus returns the TravelRuleStatus field if non-nil, zero value otherwise.

### GetTravelRuleStatusOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTravelRuleStatusOk() (*int32, bool)`

GetTravelRuleStatusOk returns a tuple with the TravelRuleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelRuleStatus

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetTravelRuleStatus(v int32)`

SetTravelRuleStatus sets TravelRuleStatus field to given value.

### HasTravelRuleStatus

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasTravelRuleStatus() bool`

HasTravelRuleStatus returns a boolean if a field has been set.

### GetTxId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetUnlockConfirm

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetUnlockConfirm() int32`

GetUnlockConfirm returns the UnlockConfirm field if non-nil, zero value otherwise.

### GetUnlockConfirmOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetUnlockConfirmOk() (*int32, bool)`

GetUnlockConfirmOk returns a tuple with the UnlockConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockConfirm

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetUnlockConfirm(v int32)`

SetUnlockConfirm sets UnlockConfirm field to given value.

### HasUnlockConfirm

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasUnlockConfirm() bool`

HasUnlockConfirm returns a boolean if a field has been set.

### GetWalletType

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetWalletType() int32`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) GetWalletTypeOk() (*int32, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) SetWalletType(v int32)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *WalletGetLocalentityDepositHistoryV1RespItem) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


