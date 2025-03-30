# OptionsGetBlockOrderOrdersV1RespItem

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

### NewOptionsGetBlockOrderOrdersV1RespItem

`func NewOptionsGetBlockOrderOrdersV1RespItem() *OptionsGetBlockOrderOrdersV1RespItem`

NewOptionsGetBlockOrderOrdersV1RespItem instantiates a new OptionsGetBlockOrderOrdersV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetBlockOrderOrdersV1RespItemWithDefaults

`func NewOptionsGetBlockOrderOrdersV1RespItemWithDefaults() *OptionsGetBlockOrderOrdersV1RespItem`

NewOptionsGetBlockOrderOrdersV1RespItemWithDefaults instantiates a new OptionsGetBlockOrderOrdersV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *OptionsGetBlockOrderOrdersV1RespItem) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *OptionsGetBlockOrderOrdersV1RespItem) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetCreateTime

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OptionsGetBlockOrderOrdersV1RespItem) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OptionsGetBlockOrderOrdersV1RespItem) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *OptionsGetBlockOrderOrdersV1RespItem) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *OptionsGetBlockOrderOrdersV1RespItem) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLegs

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetLegs() []OptionsCreateBlockOrderCreateV1RespLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetLegsOk() (*[]OptionsCreateBlockOrderCreateV1RespLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *OptionsGetBlockOrderOrdersV1RespItem) SetLegs(v []OptionsCreateBlockOrderCreateV1RespLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *OptionsGetBlockOrderOrdersV1RespItem) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetLiquidity

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *OptionsGetBlockOrderOrdersV1RespItem) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *OptionsGetBlockOrderOrdersV1RespItem) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptionsGetBlockOrderOrdersV1RespItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptionsGetBlockOrderOrdersV1RespItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OptionsGetBlockOrderOrdersV1RespItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


