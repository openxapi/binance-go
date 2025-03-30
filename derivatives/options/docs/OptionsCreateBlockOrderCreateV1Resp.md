# OptionsCreateBlockOrderCreateV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **int64** |  | [optional] 
**Legs** | Pointer to [**[]OptionsCreateBlockOrderCreateV1RespLegsInner**](OptionsCreateBlockOrderCreateV1RespLegsInner.md) |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionsCreateBlockOrderCreateV1Resp

`func NewOptionsCreateBlockOrderCreateV1Resp() *OptionsCreateBlockOrderCreateV1Resp`

NewOptionsCreateBlockOrderCreateV1Resp instantiates a new OptionsCreateBlockOrderCreateV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsCreateBlockOrderCreateV1RespWithDefaults

`func NewOptionsCreateBlockOrderCreateV1RespWithDefaults() *OptionsCreateBlockOrderCreateV1Resp`

NewOptionsCreateBlockOrderCreateV1RespWithDefaults instantiates a new OptionsCreateBlockOrderCreateV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *OptionsCreateBlockOrderCreateV1Resp) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *OptionsCreateBlockOrderCreateV1Resp) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetExpireTime

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *OptionsCreateBlockOrderCreateV1Resp) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *OptionsCreateBlockOrderCreateV1Resp) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLegs

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetLegs() []OptionsCreateBlockOrderCreateV1RespLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetLegsOk() (*[]OptionsCreateBlockOrderCreateV1RespLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *OptionsCreateBlockOrderCreateV1Resp) SetLegs(v []OptionsCreateBlockOrderCreateV1RespLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *OptionsCreateBlockOrderCreateV1Resp) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetLiquidity

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *OptionsCreateBlockOrderCreateV1Resp) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *OptionsCreateBlockOrderCreateV1Resp) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptionsCreateBlockOrderCreateV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptionsCreateBlockOrderCreateV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OptionsCreateBlockOrderCreateV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


