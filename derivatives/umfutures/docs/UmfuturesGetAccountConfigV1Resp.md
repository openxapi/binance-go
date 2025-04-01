# UmfuturesGetAccountConfigV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**DualSidePosition** | Pointer to **bool** |  | [optional] 
**FeeTier** | Pointer to **int32** |  | [optional] 
**MultiAssetsMargin** | Pointer to **bool** |  | [optional] 
**TradeGroupId** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewUmfuturesGetAccountConfigV1Resp

`func NewUmfuturesGetAccountConfigV1Resp() *UmfuturesGetAccountConfigV1Resp`

NewUmfuturesGetAccountConfigV1Resp instantiates a new UmfuturesGetAccountConfigV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesGetAccountConfigV1RespWithDefaults

`func NewUmfuturesGetAccountConfigV1RespWithDefaults() *UmfuturesGetAccountConfigV1Resp`

NewUmfuturesGetAccountConfigV1RespWithDefaults instantiates a new UmfuturesGetAccountConfigV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanDeposit

`func (o *UmfuturesGetAccountConfigV1Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *UmfuturesGetAccountConfigV1Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *UmfuturesGetAccountConfigV1Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *UmfuturesGetAccountConfigV1Resp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *UmfuturesGetAccountConfigV1Resp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *UmfuturesGetAccountConfigV1Resp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *UmfuturesGetAccountConfigV1Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *UmfuturesGetAccountConfigV1Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *UmfuturesGetAccountConfigV1Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetDualSidePosition

`func (o *UmfuturesGetAccountConfigV1Resp) GetDualSidePosition() bool`

GetDualSidePosition returns the DualSidePosition field if non-nil, zero value otherwise.

### GetDualSidePositionOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetDualSidePositionOk() (*bool, bool)`

GetDualSidePositionOk returns a tuple with the DualSidePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualSidePosition

`func (o *UmfuturesGetAccountConfigV1Resp) SetDualSidePosition(v bool)`

SetDualSidePosition sets DualSidePosition field to given value.

### HasDualSidePosition

`func (o *UmfuturesGetAccountConfigV1Resp) HasDualSidePosition() bool`

HasDualSidePosition returns a boolean if a field has been set.

### GetFeeTier

`func (o *UmfuturesGetAccountConfigV1Resp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *UmfuturesGetAccountConfigV1Resp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *UmfuturesGetAccountConfigV1Resp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetMultiAssetsMargin

`func (o *UmfuturesGetAccountConfigV1Resp) GetMultiAssetsMargin() bool`

GetMultiAssetsMargin returns the MultiAssetsMargin field if non-nil, zero value otherwise.

### GetMultiAssetsMarginOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetMultiAssetsMarginOk() (*bool, bool)`

GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAssetsMargin

`func (o *UmfuturesGetAccountConfigV1Resp) SetMultiAssetsMargin(v bool)`

SetMultiAssetsMargin sets MultiAssetsMargin field to given value.

### HasMultiAssetsMargin

`func (o *UmfuturesGetAccountConfigV1Resp) HasMultiAssetsMargin() bool`

HasMultiAssetsMargin returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *UmfuturesGetAccountConfigV1Resp) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *UmfuturesGetAccountConfigV1Resp) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *UmfuturesGetAccountConfigV1Resp) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UmfuturesGetAccountConfigV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UmfuturesGetAccountConfigV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UmfuturesGetAccountConfigV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UmfuturesGetAccountConfigV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


