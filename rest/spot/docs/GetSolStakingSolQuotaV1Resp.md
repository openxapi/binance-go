# GetSolStakingSolQuotaV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calculating** | Pointer to **bool** |  | [optional] 
**CommissionFee** | Pointer to **string** |  | [optional] 
**LeftRedemptionPersonalQuota** | Pointer to **string** |  | [optional] 
**LeftStakingPersonalQuota** | Pointer to **string** |  | [optional] 
**MinRedeemAmount** | Pointer to **string** |  | [optional] 
**MinStakeAmount** | Pointer to **string** |  | [optional] 
**NextEpochTime** | Pointer to **int64** |  | [optional] 
**RedeemPeriod** | Pointer to **int32** |  | [optional] 
**Redeemable** | Pointer to **bool** |  | [optional] 
**SoldOut** | Pointer to **bool** |  | [optional] 
**Stakeable** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetSolStakingSolQuotaV1Resp

`func NewGetSolStakingSolQuotaV1Resp() *GetSolStakingSolQuotaV1Resp`

NewGetSolStakingSolQuotaV1Resp instantiates a new GetSolStakingSolQuotaV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolStakingSolQuotaV1RespWithDefaults

`func NewGetSolStakingSolQuotaV1RespWithDefaults() *GetSolStakingSolQuotaV1Resp`

NewGetSolStakingSolQuotaV1RespWithDefaults instantiates a new GetSolStakingSolQuotaV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculating

`func (o *GetSolStakingSolQuotaV1Resp) GetCalculating() bool`

GetCalculating returns the Calculating field if non-nil, zero value otherwise.

### GetCalculatingOk

`func (o *GetSolStakingSolQuotaV1Resp) GetCalculatingOk() (*bool, bool)`

GetCalculatingOk returns a tuple with the Calculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculating

`func (o *GetSolStakingSolQuotaV1Resp) SetCalculating(v bool)`

SetCalculating sets Calculating field to given value.

### HasCalculating

`func (o *GetSolStakingSolQuotaV1Resp) HasCalculating() bool`

HasCalculating returns a boolean if a field has been set.

### GetCommissionFee

`func (o *GetSolStakingSolQuotaV1Resp) GetCommissionFee() string`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *GetSolStakingSolQuotaV1Resp) GetCommissionFeeOk() (*string, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *GetSolStakingSolQuotaV1Resp) SetCommissionFee(v string)`

SetCommissionFee sets CommissionFee field to given value.

### HasCommissionFee

`func (o *GetSolStakingSolQuotaV1Resp) HasCommissionFee() bool`

HasCommissionFee returns a boolean if a field has been set.

### GetLeftRedemptionPersonalQuota

`func (o *GetSolStakingSolQuotaV1Resp) GetLeftRedemptionPersonalQuota() string`

GetLeftRedemptionPersonalQuota returns the LeftRedemptionPersonalQuota field if non-nil, zero value otherwise.

### GetLeftRedemptionPersonalQuotaOk

`func (o *GetSolStakingSolQuotaV1Resp) GetLeftRedemptionPersonalQuotaOk() (*string, bool)`

GetLeftRedemptionPersonalQuotaOk returns a tuple with the LeftRedemptionPersonalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftRedemptionPersonalQuota

`func (o *GetSolStakingSolQuotaV1Resp) SetLeftRedemptionPersonalQuota(v string)`

SetLeftRedemptionPersonalQuota sets LeftRedemptionPersonalQuota field to given value.

### HasLeftRedemptionPersonalQuota

`func (o *GetSolStakingSolQuotaV1Resp) HasLeftRedemptionPersonalQuota() bool`

HasLeftRedemptionPersonalQuota returns a boolean if a field has been set.

### GetLeftStakingPersonalQuota

`func (o *GetSolStakingSolQuotaV1Resp) GetLeftStakingPersonalQuota() string`

GetLeftStakingPersonalQuota returns the LeftStakingPersonalQuota field if non-nil, zero value otherwise.

### GetLeftStakingPersonalQuotaOk

`func (o *GetSolStakingSolQuotaV1Resp) GetLeftStakingPersonalQuotaOk() (*string, bool)`

GetLeftStakingPersonalQuotaOk returns a tuple with the LeftStakingPersonalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftStakingPersonalQuota

`func (o *GetSolStakingSolQuotaV1Resp) SetLeftStakingPersonalQuota(v string)`

SetLeftStakingPersonalQuota sets LeftStakingPersonalQuota field to given value.

### HasLeftStakingPersonalQuota

`func (o *GetSolStakingSolQuotaV1Resp) HasLeftStakingPersonalQuota() bool`

HasLeftStakingPersonalQuota returns a boolean if a field has been set.

### GetMinRedeemAmount

`func (o *GetSolStakingSolQuotaV1Resp) GetMinRedeemAmount() string`

GetMinRedeemAmount returns the MinRedeemAmount field if non-nil, zero value otherwise.

### GetMinRedeemAmountOk

`func (o *GetSolStakingSolQuotaV1Resp) GetMinRedeemAmountOk() (*string, bool)`

GetMinRedeemAmountOk returns a tuple with the MinRedeemAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRedeemAmount

`func (o *GetSolStakingSolQuotaV1Resp) SetMinRedeemAmount(v string)`

SetMinRedeemAmount sets MinRedeemAmount field to given value.

### HasMinRedeemAmount

`func (o *GetSolStakingSolQuotaV1Resp) HasMinRedeemAmount() bool`

HasMinRedeemAmount returns a boolean if a field has been set.

### GetMinStakeAmount

`func (o *GetSolStakingSolQuotaV1Resp) GetMinStakeAmount() string`

GetMinStakeAmount returns the MinStakeAmount field if non-nil, zero value otherwise.

### GetMinStakeAmountOk

`func (o *GetSolStakingSolQuotaV1Resp) GetMinStakeAmountOk() (*string, bool)`

GetMinStakeAmountOk returns a tuple with the MinStakeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStakeAmount

`func (o *GetSolStakingSolQuotaV1Resp) SetMinStakeAmount(v string)`

SetMinStakeAmount sets MinStakeAmount field to given value.

### HasMinStakeAmount

`func (o *GetSolStakingSolQuotaV1Resp) HasMinStakeAmount() bool`

HasMinStakeAmount returns a boolean if a field has been set.

### GetNextEpochTime

`func (o *GetSolStakingSolQuotaV1Resp) GetNextEpochTime() int64`

GetNextEpochTime returns the NextEpochTime field if non-nil, zero value otherwise.

### GetNextEpochTimeOk

`func (o *GetSolStakingSolQuotaV1Resp) GetNextEpochTimeOk() (*int64, bool)`

GetNextEpochTimeOk returns a tuple with the NextEpochTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEpochTime

`func (o *GetSolStakingSolQuotaV1Resp) SetNextEpochTime(v int64)`

SetNextEpochTime sets NextEpochTime field to given value.

### HasNextEpochTime

`func (o *GetSolStakingSolQuotaV1Resp) HasNextEpochTime() bool`

HasNextEpochTime returns a boolean if a field has been set.

### GetRedeemPeriod

`func (o *GetSolStakingSolQuotaV1Resp) GetRedeemPeriod() int32`

GetRedeemPeriod returns the RedeemPeriod field if non-nil, zero value otherwise.

### GetRedeemPeriodOk

`func (o *GetSolStakingSolQuotaV1Resp) GetRedeemPeriodOk() (*int32, bool)`

GetRedeemPeriodOk returns a tuple with the RedeemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPeriod

`func (o *GetSolStakingSolQuotaV1Resp) SetRedeemPeriod(v int32)`

SetRedeemPeriod sets RedeemPeriod field to given value.

### HasRedeemPeriod

`func (o *GetSolStakingSolQuotaV1Resp) HasRedeemPeriod() bool`

HasRedeemPeriod returns a boolean if a field has been set.

### GetRedeemable

`func (o *GetSolStakingSolQuotaV1Resp) GetRedeemable() bool`

GetRedeemable returns the Redeemable field if non-nil, zero value otherwise.

### GetRedeemableOk

`func (o *GetSolStakingSolQuotaV1Resp) GetRedeemableOk() (*bool, bool)`

GetRedeemableOk returns a tuple with the Redeemable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemable

`func (o *GetSolStakingSolQuotaV1Resp) SetRedeemable(v bool)`

SetRedeemable sets Redeemable field to given value.

### HasRedeemable

`func (o *GetSolStakingSolQuotaV1Resp) HasRedeemable() bool`

HasRedeemable returns a boolean if a field has been set.

### GetSoldOut

`func (o *GetSolStakingSolQuotaV1Resp) GetSoldOut() bool`

GetSoldOut returns the SoldOut field if non-nil, zero value otherwise.

### GetSoldOutOk

`func (o *GetSolStakingSolQuotaV1Resp) GetSoldOutOk() (*bool, bool)`

GetSoldOutOk returns a tuple with the SoldOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldOut

`func (o *GetSolStakingSolQuotaV1Resp) SetSoldOut(v bool)`

SetSoldOut sets SoldOut field to given value.

### HasSoldOut

`func (o *GetSolStakingSolQuotaV1Resp) HasSoldOut() bool`

HasSoldOut returns a boolean if a field has been set.

### GetStakeable

`func (o *GetSolStakingSolQuotaV1Resp) GetStakeable() bool`

GetStakeable returns the Stakeable field if non-nil, zero value otherwise.

### GetStakeableOk

`func (o *GetSolStakingSolQuotaV1Resp) GetStakeableOk() (*bool, bool)`

GetStakeableOk returns a tuple with the Stakeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeable

`func (o *GetSolStakingSolQuotaV1Resp) SetStakeable(v bool)`

SetStakeable sets Stakeable field to given value.

### HasStakeable

`func (o *GetSolStakingSolQuotaV1Resp) HasStakeable() bool`

HasStakeable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


