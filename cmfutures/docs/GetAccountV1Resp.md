# GetAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]GetAccountV1RespAssetsInner**](GetAccountV1RespAssetsInner.md) |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int32** |  | [optional] 
**Positions** | Pointer to [**[]GetAccountV1RespPositionsInner**](GetAccountV1RespPositionsInner.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

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

### GetAssets

`func (o *GetAccountV1Resp) GetAssets() []GetAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetAccountV1Resp) GetAssetsOk() (*[]GetAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetAccountV1Resp) SetAssets(v []GetAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetAccountV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCanDeposit

`func (o *GetAccountV1Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *GetAccountV1Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *GetAccountV1Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *GetAccountV1Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *GetAccountV1Resp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *GetAccountV1Resp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *GetAccountV1Resp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *GetAccountV1Resp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *GetAccountV1Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *GetAccountV1Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *GetAccountV1Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *GetAccountV1Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeTier

`func (o *GetAccountV1Resp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *GetAccountV1Resp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *GetAccountV1Resp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *GetAccountV1Resp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetPositions

`func (o *GetAccountV1Resp) GetPositions() []GetAccountV1RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *GetAccountV1Resp) GetPositionsOk() (*[]GetAccountV1RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *GetAccountV1Resp) SetPositions(v []GetAccountV1RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *GetAccountV1Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetAccountV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetAccountV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetAccountV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetAccountV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


