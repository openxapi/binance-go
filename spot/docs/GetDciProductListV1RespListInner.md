# GetDciProductListV1RespListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apr** | Pointer to **string** |  | [optional] 
**AutoCompoundPlanList** | Pointer to **[]string** |  | [optional] 
**CanPurchase** | Pointer to **bool** |  | [optional] 
**CreateTimestamp** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**ExercisedCoin** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvestCoin** | Pointer to **string** |  | [optional] 
**IsAutoCompoundEnable** | Pointer to **bool** |  | [optional] 
**MaxAmount** | Pointer to **string** |  | [optional] 
**MinAmount** | Pointer to **string** |  | [optional] 
**OptionType** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**PurchaseDecimal** | Pointer to **int32** |  | [optional] 
**PurchaseEndTime** | Pointer to **int64** |  | [optional] 
**SettleDate** | Pointer to **int32** |  | [optional] 
**StrikePrice** | Pointer to **string** |  | [optional] 

## Methods

### NewGetDciProductListV1RespListInner

`func NewGetDciProductListV1RespListInner() *GetDciProductListV1RespListInner`

NewGetDciProductListV1RespListInner instantiates a new GetDciProductListV1RespListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDciProductListV1RespListInnerWithDefaults

`func NewGetDciProductListV1RespListInnerWithDefaults() *GetDciProductListV1RespListInner`

NewGetDciProductListV1RespListInnerWithDefaults instantiates a new GetDciProductListV1RespListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApr

`func (o *GetDciProductListV1RespListInner) GetApr() string`

GetApr returns the Apr field if non-nil, zero value otherwise.

### GetAprOk

`func (o *GetDciProductListV1RespListInner) GetAprOk() (*string, bool)`

GetAprOk returns a tuple with the Apr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr

`func (o *GetDciProductListV1RespListInner) SetApr(v string)`

SetApr sets Apr field to given value.

### HasApr

`func (o *GetDciProductListV1RespListInner) HasApr() bool`

HasApr returns a boolean if a field has been set.

### GetAutoCompoundPlanList

`func (o *GetDciProductListV1RespListInner) GetAutoCompoundPlanList() []string`

GetAutoCompoundPlanList returns the AutoCompoundPlanList field if non-nil, zero value otherwise.

### GetAutoCompoundPlanListOk

`func (o *GetDciProductListV1RespListInner) GetAutoCompoundPlanListOk() (*[]string, bool)`

GetAutoCompoundPlanListOk returns a tuple with the AutoCompoundPlanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompoundPlanList

`func (o *GetDciProductListV1RespListInner) SetAutoCompoundPlanList(v []string)`

SetAutoCompoundPlanList sets AutoCompoundPlanList field to given value.

### HasAutoCompoundPlanList

`func (o *GetDciProductListV1RespListInner) HasAutoCompoundPlanList() bool`

HasAutoCompoundPlanList returns a boolean if a field has been set.

### GetCanPurchase

`func (o *GetDciProductListV1RespListInner) GetCanPurchase() bool`

GetCanPurchase returns the CanPurchase field if non-nil, zero value otherwise.

### GetCanPurchaseOk

`func (o *GetDciProductListV1RespListInner) GetCanPurchaseOk() (*bool, bool)`

GetCanPurchaseOk returns a tuple with the CanPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPurchase

`func (o *GetDciProductListV1RespListInner) SetCanPurchase(v bool)`

SetCanPurchase sets CanPurchase field to given value.

### HasCanPurchase

`func (o *GetDciProductListV1RespListInner) HasCanPurchase() bool`

HasCanPurchase returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *GetDciProductListV1RespListInner) GetCreateTimestamp() int64`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *GetDciProductListV1RespListInner) GetCreateTimestampOk() (*int64, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *GetDciProductListV1RespListInner) SetCreateTimestamp(v int64)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *GetDciProductListV1RespListInner) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetDuration

`func (o *GetDciProductListV1RespListInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetDciProductListV1RespListInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetDciProductListV1RespListInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetDciProductListV1RespListInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetExercisedCoin

`func (o *GetDciProductListV1RespListInner) GetExercisedCoin() string`

GetExercisedCoin returns the ExercisedCoin field if non-nil, zero value otherwise.

### GetExercisedCoinOk

`func (o *GetDciProductListV1RespListInner) GetExercisedCoinOk() (*string, bool)`

GetExercisedCoinOk returns a tuple with the ExercisedCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExercisedCoin

`func (o *GetDciProductListV1RespListInner) SetExercisedCoin(v string)`

SetExercisedCoin sets ExercisedCoin field to given value.

### HasExercisedCoin

`func (o *GetDciProductListV1RespListInner) HasExercisedCoin() bool`

HasExercisedCoin returns a boolean if a field has been set.

### GetId

`func (o *GetDciProductListV1RespListInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDciProductListV1RespListInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDciProductListV1RespListInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDciProductListV1RespListInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvestCoin

`func (o *GetDciProductListV1RespListInner) GetInvestCoin() string`

GetInvestCoin returns the InvestCoin field if non-nil, zero value otherwise.

### GetInvestCoinOk

`func (o *GetDciProductListV1RespListInner) GetInvestCoinOk() (*string, bool)`

GetInvestCoinOk returns a tuple with the InvestCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestCoin

`func (o *GetDciProductListV1RespListInner) SetInvestCoin(v string)`

SetInvestCoin sets InvestCoin field to given value.

### HasInvestCoin

`func (o *GetDciProductListV1RespListInner) HasInvestCoin() bool`

HasInvestCoin returns a boolean if a field has been set.

### GetIsAutoCompoundEnable

`func (o *GetDciProductListV1RespListInner) GetIsAutoCompoundEnable() bool`

GetIsAutoCompoundEnable returns the IsAutoCompoundEnable field if non-nil, zero value otherwise.

### GetIsAutoCompoundEnableOk

`func (o *GetDciProductListV1RespListInner) GetIsAutoCompoundEnableOk() (*bool, bool)`

GetIsAutoCompoundEnableOk returns a tuple with the IsAutoCompoundEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoCompoundEnable

`func (o *GetDciProductListV1RespListInner) SetIsAutoCompoundEnable(v bool)`

SetIsAutoCompoundEnable sets IsAutoCompoundEnable field to given value.

### HasIsAutoCompoundEnable

`func (o *GetDciProductListV1RespListInner) HasIsAutoCompoundEnable() bool`

HasIsAutoCompoundEnable returns a boolean if a field has been set.

### GetMaxAmount

`func (o *GetDciProductListV1RespListInner) GetMaxAmount() string`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *GetDciProductListV1RespListInner) GetMaxAmountOk() (*string, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *GetDciProductListV1RespListInner) SetMaxAmount(v string)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *GetDciProductListV1RespListInner) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetMinAmount

`func (o *GetDciProductListV1RespListInner) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *GetDciProductListV1RespListInner) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *GetDciProductListV1RespListInner) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *GetDciProductListV1RespListInner) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetOptionType

`func (o *GetDciProductListV1RespListInner) GetOptionType() string`

GetOptionType returns the OptionType field if non-nil, zero value otherwise.

### GetOptionTypeOk

`func (o *GetDciProductListV1RespListInner) GetOptionTypeOk() (*string, bool)`

GetOptionTypeOk returns a tuple with the OptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionType

`func (o *GetDciProductListV1RespListInner) SetOptionType(v string)`

SetOptionType sets OptionType field to given value.

### HasOptionType

`func (o *GetDciProductListV1RespListInner) HasOptionType() bool`

HasOptionType returns a boolean if a field has been set.

### GetOrderId

`func (o *GetDciProductListV1RespListInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetDciProductListV1RespListInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetDciProductListV1RespListInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetDciProductListV1RespListInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPurchaseDecimal

`func (o *GetDciProductListV1RespListInner) GetPurchaseDecimal() int32`

GetPurchaseDecimal returns the PurchaseDecimal field if non-nil, zero value otherwise.

### GetPurchaseDecimalOk

`func (o *GetDciProductListV1RespListInner) GetPurchaseDecimalOk() (*int32, bool)`

GetPurchaseDecimalOk returns a tuple with the PurchaseDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDecimal

`func (o *GetDciProductListV1RespListInner) SetPurchaseDecimal(v int32)`

SetPurchaseDecimal sets PurchaseDecimal field to given value.

### HasPurchaseDecimal

`func (o *GetDciProductListV1RespListInner) HasPurchaseDecimal() bool`

HasPurchaseDecimal returns a boolean if a field has been set.

### GetPurchaseEndTime

`func (o *GetDciProductListV1RespListInner) GetPurchaseEndTime() int64`

GetPurchaseEndTime returns the PurchaseEndTime field if non-nil, zero value otherwise.

### GetPurchaseEndTimeOk

`func (o *GetDciProductListV1RespListInner) GetPurchaseEndTimeOk() (*int64, bool)`

GetPurchaseEndTimeOk returns a tuple with the PurchaseEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseEndTime

`func (o *GetDciProductListV1RespListInner) SetPurchaseEndTime(v int64)`

SetPurchaseEndTime sets PurchaseEndTime field to given value.

### HasPurchaseEndTime

`func (o *GetDciProductListV1RespListInner) HasPurchaseEndTime() bool`

HasPurchaseEndTime returns a boolean if a field has been set.

### GetSettleDate

`func (o *GetDciProductListV1RespListInner) GetSettleDate() int32`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *GetDciProductListV1RespListInner) GetSettleDateOk() (*int32, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *GetDciProductListV1RespListInner) SetSettleDate(v int32)`

SetSettleDate sets SettleDate field to given value.

### HasSettleDate

`func (o *GetDciProductListV1RespListInner) HasSettleDate() bool`

HasSettleDate returns a boolean if a field has been set.

### GetStrikePrice

`func (o *GetDciProductListV1RespListInner) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *GetDciProductListV1RespListInner) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *GetDciProductListV1RespListInner) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *GetDciProductListV1RespListInner) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


