# GetSimpleEarnLockedPositionV1RespRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APY** | Pointer to **string** |  | [optional] 
**AccrualDays** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**AutoSubscribe** | Pointer to **bool** |  | [optional] 
**BoostApr** | Pointer to **string** |  | [optional] 
**BoostRewardAsset** | Pointer to **string** |  | [optional] 
**CanFastRedemption** | Pointer to **bool** |  | [optional] 
**CanReStake** | Pointer to **bool** |  | [optional] 
**CanRedeemEarly** | Pointer to **bool** |  | [optional] 
**DeliverDate** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**EstExtraRewardAmt** | Pointer to **string** |  | [optional] 
**ExtraRewardAPR** | Pointer to **string** |  | [optional] 
**ExtraRewardAsset** | Pointer to **string** |  | [optional] 
**NextPay** | Pointer to **string** |  | [optional] 
**NextPayDate** | Pointer to **string** |  | [optional] 
**ParentPositionId** | Pointer to **int64** |  | [optional] 
**PartialAmtDeliverDate** | Pointer to **string** |  | [optional] 
**PayPeriod** | Pointer to **string** |  | [optional] 
**PositionId** | Pointer to **int64** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**PurchaseTime** | Pointer to **string** |  | [optional] 
**RedeemAmountEarly** | Pointer to **string** |  | [optional] 
**RedeemPeriod** | Pointer to **string** |  | [optional] 
**RedeemTo** | Pointer to **string** |  | [optional] 
**RedeemingAmt** | Pointer to **string** |  | [optional] 
**RewardAmt** | Pointer to **string** |  | [optional] 
**RewardAsset** | Pointer to **string** |  | [optional] 
**RewardsEndDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalBoostRewardAmt** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSimpleEarnLockedPositionV1RespRowsInner

`func NewGetSimpleEarnLockedPositionV1RespRowsInner() *GetSimpleEarnLockedPositionV1RespRowsInner`

NewGetSimpleEarnLockedPositionV1RespRowsInner instantiates a new GetSimpleEarnLockedPositionV1RespRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSimpleEarnLockedPositionV1RespRowsInnerWithDefaults

`func NewGetSimpleEarnLockedPositionV1RespRowsInnerWithDefaults() *GetSimpleEarnLockedPositionV1RespRowsInner`

NewGetSimpleEarnLockedPositionV1RespRowsInnerWithDefaults instantiates a new GetSimpleEarnLockedPositionV1RespRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPY

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAPY() string`

GetAPY returns the APY field if non-nil, zero value otherwise.

### GetAPYOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAPYOk() (*string, bool)`

GetAPYOk returns a tuple with the APY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPY

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetAPY(v string)`

SetAPY sets APY field to given value.

### HasAPY

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasAPY() bool`

HasAPY returns a boolean if a field has been set.

### GetAccrualDays

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAccrualDays() string`

GetAccrualDays returns the AccrualDays field if non-nil, zero value otherwise.

### GetAccrualDaysOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAccrualDaysOk() (*string, bool)`

GetAccrualDaysOk returns a tuple with the AccrualDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualDays

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetAccrualDays(v string)`

SetAccrualDays sets AccrualDays field to given value.

### HasAccrualDays

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasAccrualDays() bool`

HasAccrualDays returns a boolean if a field has been set.

### GetAmount

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetAutoSubscribe

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAutoSubscribe() bool`

GetAutoSubscribe returns the AutoSubscribe field if non-nil, zero value otherwise.

### GetAutoSubscribeOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetAutoSubscribeOk() (*bool, bool)`

GetAutoSubscribeOk returns a tuple with the AutoSubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSubscribe

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetAutoSubscribe(v bool)`

SetAutoSubscribe sets AutoSubscribe field to given value.

### HasAutoSubscribe

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasAutoSubscribe() bool`

HasAutoSubscribe returns a boolean if a field has been set.

### GetBoostApr

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetBoostApr() string`

GetBoostApr returns the BoostApr field if non-nil, zero value otherwise.

### GetBoostAprOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetBoostAprOk() (*string, bool)`

GetBoostAprOk returns a tuple with the BoostApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostApr

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetBoostApr(v string)`

SetBoostApr sets BoostApr field to given value.

### HasBoostApr

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasBoostApr() bool`

HasBoostApr returns a boolean if a field has been set.

### GetBoostRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetBoostRewardAsset() string`

GetBoostRewardAsset returns the BoostRewardAsset field if non-nil, zero value otherwise.

### GetBoostRewardAssetOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetBoostRewardAssetOk() (*string, bool)`

GetBoostRewardAssetOk returns a tuple with the BoostRewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetBoostRewardAsset(v string)`

SetBoostRewardAsset sets BoostRewardAsset field to given value.

### HasBoostRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasBoostRewardAsset() bool`

HasBoostRewardAsset returns a boolean if a field has been set.

### GetCanFastRedemption

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetCanFastRedemption() bool`

GetCanFastRedemption returns the CanFastRedemption field if non-nil, zero value otherwise.

### GetCanFastRedemptionOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetCanFastRedemptionOk() (*bool, bool)`

GetCanFastRedemptionOk returns a tuple with the CanFastRedemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFastRedemption

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetCanFastRedemption(v bool)`

SetCanFastRedemption sets CanFastRedemption field to given value.

### HasCanFastRedemption

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasCanFastRedemption() bool`

HasCanFastRedemption returns a boolean if a field has been set.

### GetCanReStake

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetCanReStake() bool`

GetCanReStake returns the CanReStake field if non-nil, zero value otherwise.

### GetCanReStakeOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetCanReStakeOk() (*bool, bool)`

GetCanReStakeOk returns a tuple with the CanReStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReStake

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetCanReStake(v bool)`

SetCanReStake sets CanReStake field to given value.

### HasCanReStake

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasCanReStake() bool`

HasCanReStake returns a boolean if a field has been set.

### GetCanRedeemEarly

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetCanRedeemEarly() bool`

GetCanRedeemEarly returns the CanRedeemEarly field if non-nil, zero value otherwise.

### GetCanRedeemEarlyOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetCanRedeemEarlyOk() (*bool, bool)`

GetCanRedeemEarlyOk returns a tuple with the CanRedeemEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRedeemEarly

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetCanRedeemEarly(v bool)`

SetCanRedeemEarly sets CanRedeemEarly field to given value.

### HasCanRedeemEarly

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasCanRedeemEarly() bool`

HasCanRedeemEarly returns a boolean if a field has been set.

### GetDeliverDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetDeliverDate() string`

GetDeliverDate returns the DeliverDate field if non-nil, zero value otherwise.

### GetDeliverDateOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetDeliverDateOk() (*string, bool)`

GetDeliverDateOk returns a tuple with the DeliverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetDeliverDate(v string)`

SetDeliverDate sets DeliverDate field to given value.

### HasDeliverDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasDeliverDate() bool`

HasDeliverDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEstExtraRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetEstExtraRewardAmt() string`

GetEstExtraRewardAmt returns the EstExtraRewardAmt field if non-nil, zero value otherwise.

### GetEstExtraRewardAmtOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetEstExtraRewardAmtOk() (*string, bool)`

GetEstExtraRewardAmtOk returns a tuple with the EstExtraRewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstExtraRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetEstExtraRewardAmt(v string)`

SetEstExtraRewardAmt sets EstExtraRewardAmt field to given value.

### HasEstExtraRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasEstExtraRewardAmt() bool`

HasEstExtraRewardAmt returns a boolean if a field has been set.

### GetExtraRewardAPR

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetExtraRewardAPR() string`

GetExtraRewardAPR returns the ExtraRewardAPR field if non-nil, zero value otherwise.

### GetExtraRewardAPROk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetExtraRewardAPROk() (*string, bool)`

GetExtraRewardAPROk returns a tuple with the ExtraRewardAPR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRewardAPR

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetExtraRewardAPR(v string)`

SetExtraRewardAPR sets ExtraRewardAPR field to given value.

### HasExtraRewardAPR

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasExtraRewardAPR() bool`

HasExtraRewardAPR returns a boolean if a field has been set.

### GetExtraRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetExtraRewardAsset() string`

GetExtraRewardAsset returns the ExtraRewardAsset field if non-nil, zero value otherwise.

### GetExtraRewardAssetOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetExtraRewardAssetOk() (*string, bool)`

GetExtraRewardAssetOk returns a tuple with the ExtraRewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetExtraRewardAsset(v string)`

SetExtraRewardAsset sets ExtraRewardAsset field to given value.

### HasExtraRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasExtraRewardAsset() bool`

HasExtraRewardAsset returns a boolean if a field has been set.

### GetNextPay

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetNextPay() string`

GetNextPay returns the NextPay field if non-nil, zero value otherwise.

### GetNextPayOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetNextPayOk() (*string, bool)`

GetNextPayOk returns a tuple with the NextPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPay

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetNextPay(v string)`

SetNextPay sets NextPay field to given value.

### HasNextPay

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasNextPay() bool`

HasNextPay returns a boolean if a field has been set.

### GetNextPayDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetNextPayDate() string`

GetNextPayDate returns the NextPayDate field if non-nil, zero value otherwise.

### GetNextPayDateOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetNextPayDateOk() (*string, bool)`

GetNextPayDateOk returns a tuple with the NextPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPayDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetNextPayDate(v string)`

SetNextPayDate sets NextPayDate field to given value.

### HasNextPayDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasNextPayDate() bool`

HasNextPayDate returns a boolean if a field has been set.

### GetParentPositionId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetParentPositionId() int64`

GetParentPositionId returns the ParentPositionId field if non-nil, zero value otherwise.

### GetParentPositionIdOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetParentPositionIdOk() (*int64, bool)`

GetParentPositionIdOk returns a tuple with the ParentPositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPositionId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetParentPositionId(v int64)`

SetParentPositionId sets ParentPositionId field to given value.

### HasParentPositionId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasParentPositionId() bool`

HasParentPositionId returns a boolean if a field has been set.

### GetPartialAmtDeliverDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPartialAmtDeliverDate() string`

GetPartialAmtDeliverDate returns the PartialAmtDeliverDate field if non-nil, zero value otherwise.

### GetPartialAmtDeliverDateOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPartialAmtDeliverDateOk() (*string, bool)`

GetPartialAmtDeliverDateOk returns a tuple with the PartialAmtDeliverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialAmtDeliverDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetPartialAmtDeliverDate(v string)`

SetPartialAmtDeliverDate sets PartialAmtDeliverDate field to given value.

### HasPartialAmtDeliverDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasPartialAmtDeliverDate() bool`

HasPartialAmtDeliverDate returns a boolean if a field has been set.

### GetPayPeriod

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPayPeriod() string`

GetPayPeriod returns the PayPeriod field if non-nil, zero value otherwise.

### GetPayPeriodOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPayPeriodOk() (*string, bool)`

GetPayPeriodOk returns a tuple with the PayPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriod

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetPayPeriod(v string)`

SetPayPeriod sets PayPeriod field to given value.

### HasPayPeriod

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasPayPeriod() bool`

HasPayPeriod returns a boolean if a field has been set.

### GetPositionId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPositionId() int64`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPositionIdOk() (*int64, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetPositionId(v int64)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetProjectId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPurchaseTime

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPurchaseTime() string`

GetPurchaseTime returns the PurchaseTime field if non-nil, zero value otherwise.

### GetPurchaseTimeOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetPurchaseTimeOk() (*string, bool)`

GetPurchaseTimeOk returns a tuple with the PurchaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTime

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetPurchaseTime(v string)`

SetPurchaseTime sets PurchaseTime field to given value.

### HasPurchaseTime

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasPurchaseTime() bool`

HasPurchaseTime returns a boolean if a field has been set.

### GetRedeemAmountEarly

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemAmountEarly() string`

GetRedeemAmountEarly returns the RedeemAmountEarly field if non-nil, zero value otherwise.

### GetRedeemAmountEarlyOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemAmountEarlyOk() (*string, bool)`

GetRedeemAmountEarlyOk returns a tuple with the RedeemAmountEarly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemAmountEarly

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetRedeemAmountEarly(v string)`

SetRedeemAmountEarly sets RedeemAmountEarly field to given value.

### HasRedeemAmountEarly

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasRedeemAmountEarly() bool`

HasRedeemAmountEarly returns a boolean if a field has been set.

### GetRedeemPeriod

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemPeriod() string`

GetRedeemPeriod returns the RedeemPeriod field if non-nil, zero value otherwise.

### GetRedeemPeriodOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemPeriodOk() (*string, bool)`

GetRedeemPeriodOk returns a tuple with the RedeemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPeriod

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetRedeemPeriod(v string)`

SetRedeemPeriod sets RedeemPeriod field to given value.

### HasRedeemPeriod

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasRedeemPeriod() bool`

HasRedeemPeriod returns a boolean if a field has been set.

### GetRedeemTo

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemTo() string`

GetRedeemTo returns the RedeemTo field if non-nil, zero value otherwise.

### GetRedeemToOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemToOk() (*string, bool)`

GetRedeemToOk returns a tuple with the RedeemTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemTo

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetRedeemTo(v string)`

SetRedeemTo sets RedeemTo field to given value.

### HasRedeemTo

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasRedeemTo() bool`

HasRedeemTo returns a boolean if a field has been set.

### GetRedeemingAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemingAmt() string`

GetRedeemingAmt returns the RedeemingAmt field if non-nil, zero value otherwise.

### GetRedeemingAmtOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRedeemingAmtOk() (*string, bool)`

GetRedeemingAmtOk returns a tuple with the RedeemingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemingAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetRedeemingAmt(v string)`

SetRedeemingAmt sets RedeemingAmt field to given value.

### HasRedeemingAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasRedeemingAmt() bool`

HasRedeemingAmt returns a boolean if a field has been set.

### GetRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRewardAmt() string`

GetRewardAmt returns the RewardAmt field if non-nil, zero value otherwise.

### GetRewardAmtOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRewardAmtOk() (*string, bool)`

GetRewardAmtOk returns a tuple with the RewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetRewardAmt(v string)`

SetRewardAmt sets RewardAmt field to given value.

### HasRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasRewardAmt() bool`

HasRewardAmt returns a boolean if a field has been set.

### GetRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRewardAsset() string`

GetRewardAsset returns the RewardAsset field if non-nil, zero value otherwise.

### GetRewardAssetOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRewardAssetOk() (*string, bool)`

GetRewardAssetOk returns a tuple with the RewardAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetRewardAsset(v string)`

SetRewardAsset sets RewardAsset field to given value.

### HasRewardAsset

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasRewardAsset() bool`

HasRewardAsset returns a boolean if a field has been set.

### GetRewardsEndDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRewardsEndDate() string`

GetRewardsEndDate returns the RewardsEndDate field if non-nil, zero value otherwise.

### GetRewardsEndDateOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetRewardsEndDateOk() (*string, bool)`

GetRewardsEndDateOk returns a tuple with the RewardsEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardsEndDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetRewardsEndDate(v string)`

SetRewardsEndDate sets RewardsEndDate field to given value.

### HasRewardsEndDate

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasRewardsEndDate() bool`

HasRewardsEndDate returns a boolean if a field has been set.

### GetStatus

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalBoostRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetTotalBoostRewardAmt() string`

GetTotalBoostRewardAmt returns the TotalBoostRewardAmt field if non-nil, zero value otherwise.

### GetTotalBoostRewardAmtOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetTotalBoostRewardAmtOk() (*string, bool)`

GetTotalBoostRewardAmtOk returns a tuple with the TotalBoostRewardAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBoostRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetTotalBoostRewardAmt(v string)`

SetTotalBoostRewardAmt sets TotalBoostRewardAmt field to given value.

### HasTotalBoostRewardAmt

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasTotalBoostRewardAmt() bool`

HasTotalBoostRewardAmt returns a boolean if a field has been set.

### GetType

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSimpleEarnLockedPositionV1RespRowsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


