# GetMiningPaymentListV1RespDataAccountProfitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoinName** | Pointer to **string** |  | [optional] 
**DayHashRate** | Pointer to **int32** |  | [optional] 
**HashTransfer** | Pointer to **map[string]interface{}** |  | [optional] 
**ProfitAmount** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TransferAmount** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMiningPaymentListV1RespDataAccountProfitsInner

`func NewGetMiningPaymentListV1RespDataAccountProfitsInner() *GetMiningPaymentListV1RespDataAccountProfitsInner`

NewGetMiningPaymentListV1RespDataAccountProfitsInner instantiates a new GetMiningPaymentListV1RespDataAccountProfitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiningPaymentListV1RespDataAccountProfitsInnerWithDefaults

`func NewGetMiningPaymentListV1RespDataAccountProfitsInnerWithDefaults() *GetMiningPaymentListV1RespDataAccountProfitsInner`

NewGetMiningPaymentListV1RespDataAccountProfitsInnerWithDefaults instantiates a new GetMiningPaymentListV1RespDataAccountProfitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoinName

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetCoinName() string`

GetCoinName returns the CoinName field if non-nil, zero value otherwise.

### GetCoinNameOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetCoinNameOk() (*string, bool)`

GetCoinNameOk returns a tuple with the CoinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinName

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetCoinName(v string)`

SetCoinName sets CoinName field to given value.

### HasCoinName

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasCoinName() bool`

HasCoinName returns a boolean if a field has been set.

### GetDayHashRate

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetDayHashRate() int32`

GetDayHashRate returns the DayHashRate field if non-nil, zero value otherwise.

### GetDayHashRateOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetDayHashRateOk() (*int32, bool)`

GetDayHashRateOk returns a tuple with the DayHashRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayHashRate

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetDayHashRate(v int32)`

SetDayHashRate sets DayHashRate field to given value.

### HasDayHashRate

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasDayHashRate() bool`

HasDayHashRate returns a boolean if a field has been set.

### GetHashTransfer

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetHashTransfer() map[string]interface{}`

GetHashTransfer returns the HashTransfer field if non-nil, zero value otherwise.

### GetHashTransferOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetHashTransferOk() (*map[string]interface{}, bool)`

GetHashTransferOk returns a tuple with the HashTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashTransfer

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetHashTransfer(v map[string]interface{})`

SetHashTransfer sets HashTransfer field to given value.

### HasHashTransfer

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasHashTransfer() bool`

HasHashTransfer returns a boolean if a field has been set.

### SetHashTransferNil

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetHashTransferNil(b bool)`

 SetHashTransferNil sets the value for HashTransfer to be an explicit nil

### UnsetHashTransfer
`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) UnsetHashTransfer()`

UnsetHashTransfer ensures that no value is present for HashTransfer, not even an explicit nil
### GetProfitAmount

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetProfitAmount() float32`

GetProfitAmount returns the ProfitAmount field if non-nil, zero value otherwise.

### GetProfitAmountOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetProfitAmountOk() (*float32, bool)`

GetProfitAmountOk returns a tuple with the ProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitAmount

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetProfitAmount(v float32)`

SetProfitAmount sets ProfitAmount field to given value.

### HasProfitAmount

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasProfitAmount() bool`

HasProfitAmount returns a boolean if a field has been set.

### GetStatus

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTransferAmount

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetTransferAmount() map[string]interface{}`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetTransferAmountOk() (*map[string]interface{}, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetTransferAmount(v map[string]interface{})`

SetTransferAmount sets TransferAmount field to given value.

### HasTransferAmount

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasTransferAmount() bool`

HasTransferAmount returns a boolean if a field has been set.

### SetTransferAmountNil

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetTransferAmountNil(b bool)`

 SetTransferAmountNil sets the value for TransferAmount to be an explicit nil

### UnsetTransferAmount
`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) UnsetTransferAmount()`

UnsetTransferAmount ensures that no value is present for TransferAmount, not even an explicit nil
### GetType

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GetMiningPaymentListV1RespDataAccountProfitsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


