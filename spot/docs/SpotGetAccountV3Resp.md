# SpotGetAccountV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **string** |  | [optional] 
**Balances** | Pointer to [**[]SpotGetAccountV3RespBalancesInner**](SpotGetAccountV3RespBalancesInner.md) |  | [optional] 
**Brokered** | Pointer to **bool** |  | [optional] 
**BuyerCommission** | Pointer to **int32** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**CommissionRates** | Pointer to [**SpotGetAccountCommissionV3RespStandardCommission**](SpotGetAccountCommissionV3RespStandardCommission.md) |  | [optional] 
**MakerCommission** | Pointer to **int32** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**PreventSor** | Pointer to **bool** |  | [optional] 
**RequireSelfTradePrevention** | Pointer to **bool** |  | [optional] 
**SellerCommission** | Pointer to **int32** |  | [optional] 
**TakerCommission** | Pointer to **int32** |  | [optional] 
**Uid** | Pointer to **int32** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpotGetAccountV3Resp

`func NewSpotGetAccountV3Resp() *SpotGetAccountV3Resp`

NewSpotGetAccountV3Resp instantiates a new SpotGetAccountV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetAccountV3RespWithDefaults

`func NewSpotGetAccountV3RespWithDefaults() *SpotGetAccountV3Resp`

NewSpotGetAccountV3RespWithDefaults instantiates a new SpotGetAccountV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *SpotGetAccountV3Resp) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SpotGetAccountV3Resp) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SpotGetAccountV3Resp) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *SpotGetAccountV3Resp) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalances

`func (o *SpotGetAccountV3Resp) GetBalances() []SpotGetAccountV3RespBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *SpotGetAccountV3Resp) GetBalancesOk() (*[]SpotGetAccountV3RespBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *SpotGetAccountV3Resp) SetBalances(v []SpotGetAccountV3RespBalancesInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *SpotGetAccountV3Resp) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBrokered

`func (o *SpotGetAccountV3Resp) GetBrokered() bool`

GetBrokered returns the Brokered field if non-nil, zero value otherwise.

### GetBrokeredOk

`func (o *SpotGetAccountV3Resp) GetBrokeredOk() (*bool, bool)`

GetBrokeredOk returns a tuple with the Brokered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokered

`func (o *SpotGetAccountV3Resp) SetBrokered(v bool)`

SetBrokered sets Brokered field to given value.

### HasBrokered

`func (o *SpotGetAccountV3Resp) HasBrokered() bool`

HasBrokered returns a boolean if a field has been set.

### GetBuyerCommission

`func (o *SpotGetAccountV3Resp) GetBuyerCommission() int32`

GetBuyerCommission returns the BuyerCommission field if non-nil, zero value otherwise.

### GetBuyerCommissionOk

`func (o *SpotGetAccountV3Resp) GetBuyerCommissionOk() (*int32, bool)`

GetBuyerCommissionOk returns a tuple with the BuyerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCommission

`func (o *SpotGetAccountV3Resp) SetBuyerCommission(v int32)`

SetBuyerCommission sets BuyerCommission field to given value.

### HasBuyerCommission

`func (o *SpotGetAccountV3Resp) HasBuyerCommission() bool`

HasBuyerCommission returns a boolean if a field has been set.

### GetCanDeposit

`func (o *SpotGetAccountV3Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *SpotGetAccountV3Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *SpotGetAccountV3Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *SpotGetAccountV3Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *SpotGetAccountV3Resp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *SpotGetAccountV3Resp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *SpotGetAccountV3Resp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *SpotGetAccountV3Resp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *SpotGetAccountV3Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *SpotGetAccountV3Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *SpotGetAccountV3Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *SpotGetAccountV3Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetCommissionRates

`func (o *SpotGetAccountV3Resp) GetCommissionRates() SpotGetAccountCommissionV3RespStandardCommission`

GetCommissionRates returns the CommissionRates field if non-nil, zero value otherwise.

### GetCommissionRatesOk

`func (o *SpotGetAccountV3Resp) GetCommissionRatesOk() (*SpotGetAccountCommissionV3RespStandardCommission, bool)`

GetCommissionRatesOk returns a tuple with the CommissionRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionRates

`func (o *SpotGetAccountV3Resp) SetCommissionRates(v SpotGetAccountCommissionV3RespStandardCommission)`

SetCommissionRates sets CommissionRates field to given value.

### HasCommissionRates

`func (o *SpotGetAccountV3Resp) HasCommissionRates() bool`

HasCommissionRates returns a boolean if a field has been set.

### GetMakerCommission

`func (o *SpotGetAccountV3Resp) GetMakerCommission() int32`

GetMakerCommission returns the MakerCommission field if non-nil, zero value otherwise.

### GetMakerCommissionOk

`func (o *SpotGetAccountV3Resp) GetMakerCommissionOk() (*int32, bool)`

GetMakerCommissionOk returns a tuple with the MakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerCommission

`func (o *SpotGetAccountV3Resp) SetMakerCommission(v int32)`

SetMakerCommission sets MakerCommission field to given value.

### HasMakerCommission

`func (o *SpotGetAccountV3Resp) HasMakerCommission() bool`

HasMakerCommission returns a boolean if a field has been set.

### GetPermissions

`func (o *SpotGetAccountV3Resp) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SpotGetAccountV3Resp) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SpotGetAccountV3Resp) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SpotGetAccountV3Resp) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPreventSor

`func (o *SpotGetAccountV3Resp) GetPreventSor() bool`

GetPreventSor returns the PreventSor field if non-nil, zero value otherwise.

### GetPreventSorOk

`func (o *SpotGetAccountV3Resp) GetPreventSorOk() (*bool, bool)`

GetPreventSorOk returns a tuple with the PreventSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventSor

`func (o *SpotGetAccountV3Resp) SetPreventSor(v bool)`

SetPreventSor sets PreventSor field to given value.

### HasPreventSor

`func (o *SpotGetAccountV3Resp) HasPreventSor() bool`

HasPreventSor returns a boolean if a field has been set.

### GetRequireSelfTradePrevention

`func (o *SpotGetAccountV3Resp) GetRequireSelfTradePrevention() bool`

GetRequireSelfTradePrevention returns the RequireSelfTradePrevention field if non-nil, zero value otherwise.

### GetRequireSelfTradePreventionOk

`func (o *SpotGetAccountV3Resp) GetRequireSelfTradePreventionOk() (*bool, bool)`

GetRequireSelfTradePreventionOk returns a tuple with the RequireSelfTradePrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSelfTradePrevention

`func (o *SpotGetAccountV3Resp) SetRequireSelfTradePrevention(v bool)`

SetRequireSelfTradePrevention sets RequireSelfTradePrevention field to given value.

### HasRequireSelfTradePrevention

`func (o *SpotGetAccountV3Resp) HasRequireSelfTradePrevention() bool`

HasRequireSelfTradePrevention returns a boolean if a field has been set.

### GetSellerCommission

`func (o *SpotGetAccountV3Resp) GetSellerCommission() int32`

GetSellerCommission returns the SellerCommission field if non-nil, zero value otherwise.

### GetSellerCommissionOk

`func (o *SpotGetAccountV3Resp) GetSellerCommissionOk() (*int32, bool)`

GetSellerCommissionOk returns a tuple with the SellerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerCommission

`func (o *SpotGetAccountV3Resp) SetSellerCommission(v int32)`

SetSellerCommission sets SellerCommission field to given value.

### HasSellerCommission

`func (o *SpotGetAccountV3Resp) HasSellerCommission() bool`

HasSellerCommission returns a boolean if a field has been set.

### GetTakerCommission

`func (o *SpotGetAccountV3Resp) GetTakerCommission() int32`

GetTakerCommission returns the TakerCommission field if non-nil, zero value otherwise.

### GetTakerCommissionOk

`func (o *SpotGetAccountV3Resp) GetTakerCommissionOk() (*int32, bool)`

GetTakerCommissionOk returns a tuple with the TakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommission

`func (o *SpotGetAccountV3Resp) SetTakerCommission(v int32)`

SetTakerCommission sets TakerCommission field to given value.

### HasTakerCommission

`func (o *SpotGetAccountV3Resp) HasTakerCommission() bool`

HasTakerCommission returns a boolean if a field has been set.

### GetUid

`func (o *SpotGetAccountV3Resp) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SpotGetAccountV3Resp) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SpotGetAccountV3Resp) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SpotGetAccountV3Resp) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdateTime

`func (o *SpotGetAccountV3Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *SpotGetAccountV3Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *SpotGetAccountV3Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *SpotGetAccountV3Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


