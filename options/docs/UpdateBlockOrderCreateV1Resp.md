# UpdateBlockOrderCreateV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**ExpireTime** | Pointer to **int64** |  | [optional] 
**Legs** | Pointer to [**[]CreateBlockOrderExecuteV1RespLegsInner**](CreateBlockOrderExecuteV1RespLegsInner.md) |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateBlockOrderCreateV1Resp

`func NewUpdateBlockOrderCreateV1Resp() *UpdateBlockOrderCreateV1Resp`

NewUpdateBlockOrderCreateV1Resp instantiates a new UpdateBlockOrderCreateV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlockOrderCreateV1RespWithDefaults

`func NewUpdateBlockOrderCreateV1RespWithDefaults() *UpdateBlockOrderCreateV1Resp`

NewUpdateBlockOrderCreateV1RespWithDefaults instantiates a new UpdateBlockOrderCreateV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *UpdateBlockOrderCreateV1Resp) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *UpdateBlockOrderCreateV1Resp) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *UpdateBlockOrderCreateV1Resp) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *UpdateBlockOrderCreateV1Resp) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetCreateTime

`func (o *UpdateBlockOrderCreateV1Resp) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UpdateBlockOrderCreateV1Resp) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UpdateBlockOrderCreateV1Resp) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UpdateBlockOrderCreateV1Resp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *UpdateBlockOrderCreateV1Resp) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *UpdateBlockOrderCreateV1Resp) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *UpdateBlockOrderCreateV1Resp) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *UpdateBlockOrderCreateV1Resp) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLegs

`func (o *UpdateBlockOrderCreateV1Resp) GetLegs() []CreateBlockOrderExecuteV1RespLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *UpdateBlockOrderCreateV1Resp) GetLegsOk() (*[]CreateBlockOrderExecuteV1RespLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *UpdateBlockOrderCreateV1Resp) SetLegs(v []CreateBlockOrderExecuteV1RespLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *UpdateBlockOrderCreateV1Resp) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetLiquidity

`func (o *UpdateBlockOrderCreateV1Resp) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *UpdateBlockOrderCreateV1Resp) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *UpdateBlockOrderCreateV1Resp) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *UpdateBlockOrderCreateV1Resp) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateBlockOrderCreateV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateBlockOrderCreateV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateBlockOrderCreateV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateBlockOrderCreateV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


