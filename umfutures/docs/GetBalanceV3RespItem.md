# GetBalanceV3RespItem

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

### NewGetBalanceV3RespItem

`func NewGetBalanceV3RespItem() *GetBalanceV3RespItem`

NewGetBalanceV3RespItem instantiates a new GetBalanceV3RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceV3RespItemWithDefaults

`func NewGetBalanceV3RespItemWithDefaults() *GetBalanceV3RespItem`

NewGetBalanceV3RespItemWithDefaults instantiates a new GetBalanceV3RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *GetBalanceV3RespItem) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *GetBalanceV3RespItem) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *GetBalanceV3RespItem) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *GetBalanceV3RespItem) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetAsset

`func (o *GetBalanceV3RespItem) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetBalanceV3RespItem) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetBalanceV3RespItem) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetBalanceV3RespItem) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *GetBalanceV3RespItem) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *GetBalanceV3RespItem) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *GetBalanceV3RespItem) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *GetBalanceV3RespItem) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetBalance

`func (o *GetBalanceV3RespItem) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetBalanceV3RespItem) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetBalanceV3RespItem) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *GetBalanceV3RespItem) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCrossUnPnl

`func (o *GetBalanceV3RespItem) GetCrossUnPnl() string`

GetCrossUnPnl returns the CrossUnPnl field if non-nil, zero value otherwise.

### GetCrossUnPnlOk

`func (o *GetBalanceV3RespItem) GetCrossUnPnlOk() (*string, bool)`

GetCrossUnPnlOk returns a tuple with the CrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossUnPnl

`func (o *GetBalanceV3RespItem) SetCrossUnPnl(v string)`

SetCrossUnPnl sets CrossUnPnl field to given value.

### HasCrossUnPnl

`func (o *GetBalanceV3RespItem) HasCrossUnPnl() bool`

HasCrossUnPnl returns a boolean if a field has been set.

### GetCrossWalletBalance

`func (o *GetBalanceV3RespItem) GetCrossWalletBalance() string`

GetCrossWalletBalance returns the CrossWalletBalance field if non-nil, zero value otherwise.

### GetCrossWalletBalanceOk

`func (o *GetBalanceV3RespItem) GetCrossWalletBalanceOk() (*string, bool)`

GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossWalletBalance

`func (o *GetBalanceV3RespItem) SetCrossWalletBalance(v string)`

SetCrossWalletBalance sets CrossWalletBalance field to given value.

### HasCrossWalletBalance

`func (o *GetBalanceV3RespItem) HasCrossWalletBalance() bool`

HasCrossWalletBalance returns a boolean if a field has been set.

### GetMarginAvailable

`func (o *GetBalanceV3RespItem) GetMarginAvailable() bool`

GetMarginAvailable returns the MarginAvailable field if non-nil, zero value otherwise.

### GetMarginAvailableOk

`func (o *GetBalanceV3RespItem) GetMarginAvailableOk() (*bool, bool)`

GetMarginAvailableOk returns a tuple with the MarginAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAvailable

`func (o *GetBalanceV3RespItem) SetMarginAvailable(v bool)`

SetMarginAvailable sets MarginAvailable field to given value.

### HasMarginAvailable

`func (o *GetBalanceV3RespItem) HasMarginAvailable() bool`

HasMarginAvailable returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *GetBalanceV3RespItem) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *GetBalanceV3RespItem) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *GetBalanceV3RespItem) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *GetBalanceV3RespItem) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetBalanceV3RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetBalanceV3RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetBalanceV3RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetBalanceV3RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


