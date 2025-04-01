# CmfuturesGetAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]CmfuturesGetAccountV1RespAssetsInner**](CmfuturesGetAccountV1RespAssetsInner.md) |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int32** |  | [optional] 
**Positions** | Pointer to [**[]CmfuturesGetAccountV1RespPositionsInner**](CmfuturesGetAccountV1RespPositionsInner.md) |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewCmfuturesGetAccountV1Resp

`func NewCmfuturesGetAccountV1Resp() *CmfuturesGetAccountV1Resp`

NewCmfuturesGetAccountV1Resp instantiates a new CmfuturesGetAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmfuturesGetAccountV1RespWithDefaults

`func NewCmfuturesGetAccountV1RespWithDefaults() *CmfuturesGetAccountV1Resp`

NewCmfuturesGetAccountV1RespWithDefaults instantiates a new CmfuturesGetAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *CmfuturesGetAccountV1Resp) GetAssets() []CmfuturesGetAccountV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *CmfuturesGetAccountV1Resp) GetAssetsOk() (*[]CmfuturesGetAccountV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *CmfuturesGetAccountV1Resp) SetAssets(v []CmfuturesGetAccountV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *CmfuturesGetAccountV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCanDeposit

`func (o *CmfuturesGetAccountV1Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *CmfuturesGetAccountV1Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *CmfuturesGetAccountV1Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *CmfuturesGetAccountV1Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *CmfuturesGetAccountV1Resp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *CmfuturesGetAccountV1Resp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *CmfuturesGetAccountV1Resp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *CmfuturesGetAccountV1Resp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *CmfuturesGetAccountV1Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *CmfuturesGetAccountV1Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *CmfuturesGetAccountV1Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *CmfuturesGetAccountV1Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetFeeTier

`func (o *CmfuturesGetAccountV1Resp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *CmfuturesGetAccountV1Resp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *CmfuturesGetAccountV1Resp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *CmfuturesGetAccountV1Resp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetPositions

`func (o *CmfuturesGetAccountV1Resp) GetPositions() []CmfuturesGetAccountV1RespPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *CmfuturesGetAccountV1Resp) GetPositionsOk() (*[]CmfuturesGetAccountV1RespPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *CmfuturesGetAccountV1Resp) SetPositions(v []CmfuturesGetAccountV1RespPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *CmfuturesGetAccountV1Resp) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CmfuturesGetAccountV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CmfuturesGetAccountV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CmfuturesGetAccountV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CmfuturesGetAccountV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


