# PmarginGetUmAccountConfigV1Resp

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

### NewPmarginGetUmAccountConfigV1Resp

`func NewPmarginGetUmAccountConfigV1Resp() *PmarginGetUmAccountConfigV1Resp`

NewPmarginGetUmAccountConfigV1Resp instantiates a new PmarginGetUmAccountConfigV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetUmAccountConfigV1RespWithDefaults

`func NewPmarginGetUmAccountConfigV1RespWithDefaults() *PmarginGetUmAccountConfigV1Resp`

NewPmarginGetUmAccountConfigV1RespWithDefaults instantiates a new PmarginGetUmAccountConfigV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanDeposit

`func (o *PmarginGetUmAccountConfigV1Resp) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *PmarginGetUmAccountConfigV1Resp) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *PmarginGetUmAccountConfigV1Resp) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanTrade

`func (o *PmarginGetUmAccountConfigV1Resp) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *PmarginGetUmAccountConfigV1Resp) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *PmarginGetUmAccountConfigV1Resp) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *PmarginGetUmAccountConfigV1Resp) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *PmarginGetUmAccountConfigV1Resp) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *PmarginGetUmAccountConfigV1Resp) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetDualSidePosition

`func (o *PmarginGetUmAccountConfigV1Resp) GetDualSidePosition() bool`

GetDualSidePosition returns the DualSidePosition field if non-nil, zero value otherwise.

### GetDualSidePositionOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetDualSidePositionOk() (*bool, bool)`

GetDualSidePositionOk returns a tuple with the DualSidePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualSidePosition

`func (o *PmarginGetUmAccountConfigV1Resp) SetDualSidePosition(v bool)`

SetDualSidePosition sets DualSidePosition field to given value.

### HasDualSidePosition

`func (o *PmarginGetUmAccountConfigV1Resp) HasDualSidePosition() bool`

HasDualSidePosition returns a boolean if a field has been set.

### GetFeeTier

`func (o *PmarginGetUmAccountConfigV1Resp) GetFeeTier() int32`

GetFeeTier returns the FeeTier field if non-nil, zero value otherwise.

### GetFeeTierOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetFeeTierOk() (*int32, bool)`

GetFeeTierOk returns a tuple with the FeeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTier

`func (o *PmarginGetUmAccountConfigV1Resp) SetFeeTier(v int32)`

SetFeeTier sets FeeTier field to given value.

### HasFeeTier

`func (o *PmarginGetUmAccountConfigV1Resp) HasFeeTier() bool`

HasFeeTier returns a boolean if a field has been set.

### GetMultiAssetsMargin

`func (o *PmarginGetUmAccountConfigV1Resp) GetMultiAssetsMargin() bool`

GetMultiAssetsMargin returns the MultiAssetsMargin field if non-nil, zero value otherwise.

### GetMultiAssetsMarginOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetMultiAssetsMarginOk() (*bool, bool)`

GetMultiAssetsMarginOk returns a tuple with the MultiAssetsMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAssetsMargin

`func (o *PmarginGetUmAccountConfigV1Resp) SetMultiAssetsMargin(v bool)`

SetMultiAssetsMargin sets MultiAssetsMargin field to given value.

### HasMultiAssetsMargin

`func (o *PmarginGetUmAccountConfigV1Resp) HasMultiAssetsMargin() bool`

HasMultiAssetsMargin returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *PmarginGetUmAccountConfigV1Resp) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *PmarginGetUmAccountConfigV1Resp) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *PmarginGetUmAccountConfigV1Resp) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PmarginGetUmAccountConfigV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PmarginGetUmAccountConfigV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PmarginGetUmAccountConfigV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PmarginGetUmAccountConfigV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


