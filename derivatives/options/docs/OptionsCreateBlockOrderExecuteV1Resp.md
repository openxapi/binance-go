# OptionsCreateBlockOrderExecuteV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**ExpireTime** | Pointer to **int64** |  | [optional] 
**Legs** | Pointer to [**[]OptionsCreateBlockOrderCreateV1RespLegsInner**](OptionsCreateBlockOrderCreateV1RespLegsInner.md) |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionsCreateBlockOrderExecuteV1Resp

`func NewOptionsCreateBlockOrderExecuteV1Resp() *OptionsCreateBlockOrderExecuteV1Resp`

NewOptionsCreateBlockOrderExecuteV1Resp instantiates a new OptionsCreateBlockOrderExecuteV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsCreateBlockOrderExecuteV1RespWithDefaults

`func NewOptionsCreateBlockOrderExecuteV1RespWithDefaults() *OptionsCreateBlockOrderExecuteV1Resp`

NewOptionsCreateBlockOrderExecuteV1RespWithDefaults instantiates a new OptionsCreateBlockOrderExecuteV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *OptionsCreateBlockOrderExecuteV1Resp) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *OptionsCreateBlockOrderExecuteV1Resp) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetCreateTime

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OptionsCreateBlockOrderExecuteV1Resp) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OptionsCreateBlockOrderExecuteV1Resp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *OptionsCreateBlockOrderExecuteV1Resp) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *OptionsCreateBlockOrderExecuteV1Resp) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLegs

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetLegs() []OptionsCreateBlockOrderCreateV1RespLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetLegsOk() (*[]OptionsCreateBlockOrderCreateV1RespLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *OptionsCreateBlockOrderExecuteV1Resp) SetLegs(v []OptionsCreateBlockOrderCreateV1RespLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *OptionsCreateBlockOrderExecuteV1Resp) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetLiquidity

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *OptionsCreateBlockOrderExecuteV1Resp) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *OptionsCreateBlockOrderExecuteV1Resp) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptionsCreateBlockOrderExecuteV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptionsCreateBlockOrderExecuteV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OptionsCreateBlockOrderExecuteV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


