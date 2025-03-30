# CfuturesGetAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]CfuturesGetAccountV1RespAssetsInner**](CfuturesGetAccountV1RespAssetsInner.md) |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int32** |  | [optional] 
**Positions** | Pointer to [**[]CfuturesGetAccountV1RespPositionsInner**](CfuturesGetAccountV1RespPositionsInner.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewCfuturesGetAccountV1Resp

`func NewCfuturesGetAccountV1Resp() *CfuturesGetAccountV1Resp`

NewCfuturesGetAccountV1Resp instantiates a new CfuturesGetAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCfuturesGetAccountV1RespWithDefaults

`func NewCfuturesGetAccountV1RespWithDefaults() *CfuturesGetAccountV1Resp`

NewCfuturesGetAccountV1RespWithDefaults instantiates a new CfuturesGetAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *CfuturesGetAccountV1Resp) GetAssets() []CfuturesGetAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CfuturesGetAccountV1Resp) GetAssetsOk() (*[]CfuturesGetAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CfuturesGetAccountV1Resp) SetAssets(v []CfuturesGetAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CfuturesGetAccountV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCanDeposit

`func (o *CfuturesGetAccountV1Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *CfuturesGetAccountV1Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *CfuturesGetAccountV1Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *CfuturesGetAccountV1Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *CfuturesGetAccountV1Resp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *CfuturesGetAccountV1Resp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *CfuturesGetAccountV1Resp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *CfuturesGetAccountV1Resp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *CfuturesGetAccountV1Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *CfuturesGetAccountV1Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *CfuturesGetAccountV1Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *CfuturesGetAccountV1Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeTier

`func (o *CfuturesGetAccountV1Resp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *CfuturesGetAccountV1Resp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *CfuturesGetAccountV1Resp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *CfuturesGetAccountV1Resp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetPositions

`func (o *CfuturesGetAccountV1Resp) GetPositions() []CfuturesGetAccountV1RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *CfuturesGetAccountV1Resp) GetPositionsOk() (*[]CfuturesGetAccountV1RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *CfuturesGetAccountV1Resp) SetPositions(v []CfuturesGetAccountV1RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *CfuturesGetAccountV1Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CfuturesGetAccountV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CfuturesGetAccountV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CfuturesGetAccountV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CfuturesGetAccountV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


