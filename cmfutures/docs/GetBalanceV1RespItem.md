# GetBalanceV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**AvailableBalance** | Pointer to **string** |  | [optional] 
**Balance** | Pointer to **string** |  | [optional] 
**CrossUnPnl** | Pointer to **string** |  | [optional] 
**CrossWalletBalance** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WithdrawAvailable** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBalanceV1RespItem

`func NewGetBalanceV1RespItem() *GetBalanceV1RespItem`

NewGetBalanceV1RespItem instantiates a new GetBalanceV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceV1RespItemWithDefaults

`func NewGetBalanceV1RespItemWithDefaults() *GetBalanceV1RespItem`

NewGetBalanceV1RespItemWithDefaults instantiates a new GetBalanceV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *GetBalanceV1RespItem) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *GetBalanceV1RespItem) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *GetBalanceV1RespItem) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *GetBalanceV1RespItem) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetAsset

`func (o *GetBalanceV1RespItem) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetBalanceV1RespItem) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetBalanceV1RespItem) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetBalanceV1RespItem) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *GetBalanceV1RespItem) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *GetBalanceV1RespItem) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *GetBalanceV1RespItem) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *GetBalanceV1RespItem) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetBalance

`func (o *GetBalanceV1RespItem) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetBalanceV1RespItem) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetBalanceV1RespItem) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *GetBalanceV1RespItem) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCrossUnPnl

`func (o *GetBalanceV1RespItem) GetCrossUnPnl() string`

GetCrossUnPnl returns the CrossUnPnl field if non-nil, zero value otherwise.

### GetCrossUnPnlOk

`func (o *GetBalanceV1RespItem) GetCrossUnPnlOk() (*string, bool)`

GetCrossUnPnlOk returns a tuple with the CrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossUnPnl

`func (o *GetBalanceV1RespItem) SetCrossUnPnl(v string)`

SetCrossUnPnl sets CrossUnPnl field to given value.

### HasCrossUnPnl

`func (o *GetBalanceV1RespItem) HasCrossUnPnl() bool`

HasCrossUnPnl returns a boolean if a field has been set.

### GetCrossWalletBalance

`func (o *GetBalanceV1RespItem) GetCrossWalletBalance() string`

GetCrossWalletBalance returns the CrossWalletBalance field if non-nil, zero value otherwise.

### GetCrossWalletBalanceOk

`func (o *GetBalanceV1RespItem) GetCrossWalletBalanceOk() (*string, bool)`

GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossWalletBalance

`func (o *GetBalanceV1RespItem) SetCrossWalletBalance(v string)`

SetCrossWalletBalance sets CrossWalletBalance field to given value.

### HasCrossWalletBalance

`func (o *GetBalanceV1RespItem) HasCrossWalletBalance() bool`

HasCrossWalletBalance returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetBalanceV1RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetBalanceV1RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetBalanceV1RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetBalanceV1RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWithdrawAvailable

`func (o *GetBalanceV1RespItem) GetWithdrawAvailable() string`

GetWithdrawAvailable returns the WithdrawAvailable field if non-nil, zero value otherwise.

### GetWithdrawAvailableOk

`func (o *GetBalanceV1RespItem) GetWithdrawAvailableOk() (*string, bool)`

GetWithdrawAvailableOk returns a tuple with the WithdrawAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawAvailable

`func (o *GetBalanceV1RespItem) SetWithdrawAvailable(v string)`

SetWithdrawAvailable sets WithdrawAvailable field to given value.

### HasWithdrawAvailable

`func (o *GetBalanceV1RespItem) HasWithdrawAvailable() bool`

HasWithdrawAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


