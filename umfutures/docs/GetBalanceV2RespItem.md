# GetBalanceV2RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**AvailableBalance** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to **string** |  | [optional] 
**CrossUnPnl** | Pointer to **string** |  | [optional] 
**CrossWalletBalance** | Pointer to **string** |  | [optional] 
**MarginAvailable** | Pointer to **bool** |  | [optional] 
**MaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetBalanceV2RespItem

`func NewGetBalanceV2RespItem() *GetBalanceV2RespItem`

NewGetBalanceV2RespItem instantiates a new GetBalanceV2RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceV2RespItemWithDefaults

`func NewGetBalanceV2RespItemWithDefaults() *GetBalanceV2RespItem`

NewGetBalanceV2RespItemWithDefaults instantiates a new GetBalanceV2RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *GetBalanceV2RespItem) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *GetBalanceV2RespItem) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *GetBalanceV2RespItem) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *GetBalanceV2RespItem) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetAsset

`func (o *GetBalanceV2RespItem) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetBalanceV2RespItem) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetBalanceV2RespItem) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetBalanceV2RespItem) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *GetBalanceV2RespItem) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *GetBalanceV2RespItem) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *GetBalanceV2RespItem) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *GetBalanceV2RespItem) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetBalance

`func (o *GetBalanceV2RespItem) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetBalanceV2RespItem) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetBalanceV2RespItem) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *GetBalanceV2RespItem) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCrossUnPnl

`func (o *GetBalanceV2RespItem) GetCrossUnPnl() string`

GetCrossUnPnl returns the CrossUnPnl field if non-nil, zero value otherwise.

### GetCrossUnPnlOk

`func (o *GetBalanceV2RespItem) GetCrossUnPnlOk() (*string, bool)`

GetCrossUnPnlOk returns a tuple with the CrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossUnPnl

`func (o *GetBalanceV2RespItem) SetCrossUnPnl(v string)`

SetCrossUnPnl sets CrossUnPnl field to given value.

### HasCrossUnPnl

`func (o *GetBalanceV2RespItem) HasCrossUnPnl() bool`

HasCrossUnPnl returns a boolean if a field has been set.

### GetCrossWalletBalance

`func (o *GetBalanceV2RespItem) GetCrossWalletBalance() string`

GetCrossWalletBalance returns the CrossWalletBalance field if non-nil, zero value otherwise.

### GetCrossWalletBalanceOk

`func (o *GetBalanceV2RespItem) GetCrossWalletBalanceOk() (*string, bool)`

GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossWalletBalance

`func (o *GetBalanceV2RespItem) SetCrossWalletBalance(v string)`

SetCrossWalletBalance sets CrossWalletBalance field to given value.

### HasCrossWalletBalance

`func (o *GetBalanceV2RespItem) HasCrossWalletBalance() bool`

HasCrossWalletBalance returns a boolean if a field has been set.

### GetMarginAvailable

`func (o *GetBalanceV2RespItem) GetMarginAvailable() bool`

GetMarginAvailable returns the MarginAvailable field if non-nil, zero value otherwise.

### GetMarginAvailableOk

`func (o *GetBalanceV2RespItem) GetMarginAvailableOk() (*bool, bool)`

GetMarginAvailableOk returns a tuple with the MarginAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAvailable

`func (o *GetBalanceV2RespItem) SetMarginAvailable(v bool)`

SetMarginAvailable sets MarginAvailable field to given value.

### HasMarginAvailable

`func (o *GetBalanceV2RespItem) HasMarginAvailable() bool`

HasMarginAvailable returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *GetBalanceV2RespItem) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *GetBalanceV2RespItem) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *GetBalanceV2RespItem) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *GetBalanceV2RespItem) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetBalanceV2RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetBalanceV2RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetBalanceV2RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetBalanceV2RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


