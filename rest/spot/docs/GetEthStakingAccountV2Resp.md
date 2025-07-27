# GetEthStakingAccountV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoldingInETH** | Pointer to **string** |  | [optional] 
**Holdings** | Pointer to [**GetEthStakingAccountV2RespHoldings**](GetEthStakingAccountV2RespHoldings.md) |  | [optional] 
**Profit** | Pointer to [**GetEthStakingAccountV2RespProfit**](GetEthStakingAccountV2RespProfit.md) |  | [optional] 
**ThirtyDaysProfitInETH** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEthStakingAccountV2Resp

`func NewGetEthStakingAccountV2Resp() *GetEthStakingAccountV2Resp`

NewGetEthStakingAccountV2Resp instantiates a new GetEthStakingAccountV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthStakingAccountV2RespWithDefaults

`func NewGetEthStakingAccountV2RespWithDefaults() *GetEthStakingAccountV2Resp`

NewGetEthStakingAccountV2RespWithDefaults instantiates a new GetEthStakingAccountV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoldingInETH

`func (o *GetEthStakingAccountV2Resp) GetHoldingInETH() string`

GetHoldingInETH returns the HoldingInETH field if non-nil, zero value otherwise.

### GetHoldingInETHOk

`func (o *GetEthStakingAccountV2Resp) GetHoldingInETHOk() (*string, bool)`

GetHoldingInETHOk returns a tuple with the HoldingInETH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingInETH

`func (o *GetEthStakingAccountV2Resp) SetHoldingInETH(v string)`

SetHoldingInETH sets HoldingInETH field to given value.

### HasHoldingInETH

`func (o *GetEthStakingAccountV2Resp) HasHoldingInETH() bool`

HasHoldingInETH returns a boolean if a field has been set.

### GetHoldings

`func (o *GetEthStakingAccountV2Resp) GetHoldings() GetEthStakingAccountV2RespHoldings`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *GetEthStakingAccountV2Resp) GetHoldingsOk() (*GetEthStakingAccountV2RespHoldings, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *GetEthStakingAccountV2Resp) SetHoldings(v GetEthStakingAccountV2RespHoldings)`

SetHoldings sets Holdings field to given value.

### HasHoldings

`func (o *GetEthStakingAccountV2Resp) HasHoldings() bool`

HasHoldings returns a boolean if a field has been set.

### GetProfit

`func (o *GetEthStakingAccountV2Resp) GetProfit() GetEthStakingAccountV2RespProfit`

GetProfit returns the Profit field if non-nil, zero value otherwise.

### GetProfitOk

`func (o *GetEthStakingAccountV2Resp) GetProfitOk() (*GetEthStakingAccountV2RespProfit, bool)`

GetProfitOk returns a tuple with the Profit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfit

`func (o *GetEthStakingAccountV2Resp) SetProfit(v GetEthStakingAccountV2RespProfit)`

SetProfit sets Profit field to given value.

### HasProfit

`func (o *GetEthStakingAccountV2Resp) HasProfit() bool`

HasProfit returns a boolean if a field has been set.

### GetThirtyDaysProfitInETH

`func (o *GetEthStakingAccountV2Resp) GetThirtyDaysProfitInETH() string`

GetThirtyDaysProfitInETH returns the ThirtyDaysProfitInETH field if non-nil, zero value otherwise.

### GetThirtyDaysProfitInETHOk

`func (o *GetEthStakingAccountV2Resp) GetThirtyDaysProfitInETHOk() (*string, bool)`

GetThirtyDaysProfitInETHOk returns a tuple with the ThirtyDaysProfitInETH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirtyDaysProfitInETH

`func (o *GetEthStakingAccountV2Resp) SetThirtyDaysProfitInETH(v string)`

SetThirtyDaysProfitInETH sets ThirtyDaysProfitInETH field to given value.

### HasThirtyDaysProfitInETH

`func (o *GetEthStakingAccountV2Resp) HasThirtyDaysProfitInETH() bool`

HasThirtyDaysProfitInETH returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


