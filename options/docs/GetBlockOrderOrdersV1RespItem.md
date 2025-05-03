# GetBlockOrderOrdersV1RespItem

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

### NewGetBlockOrderOrdersV1RespItem

`func NewGetBlockOrderOrdersV1RespItem() *GetBlockOrderOrdersV1RespItem`

NewGetBlockOrderOrdersV1RespItem instantiates a new GetBlockOrderOrdersV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockOrderOrdersV1RespItemWithDefaults

`func NewGetBlockOrderOrdersV1RespItemWithDefaults() *GetBlockOrderOrdersV1RespItem`

NewGetBlockOrderOrdersV1RespItemWithDefaults instantiates a new GetBlockOrderOrdersV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *GetBlockOrderOrdersV1RespItem) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *GetBlockOrderOrdersV1RespItem) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *GetBlockOrderOrdersV1RespItem) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *GetBlockOrderOrdersV1RespItem) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetBlockOrderOrdersV1RespItem) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetBlockOrderOrdersV1RespItem) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetBlockOrderOrdersV1RespItem) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetBlockOrderOrdersV1RespItem) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *GetBlockOrderOrdersV1RespItem) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *GetBlockOrderOrdersV1RespItem) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *GetBlockOrderOrdersV1RespItem) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *GetBlockOrderOrdersV1RespItem) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLegs

`func (o *GetBlockOrderOrdersV1RespItem) GetLegs() []CreateBlockOrderExecuteV1RespLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetBlockOrderOrdersV1RespItem) GetLegsOk() (*[]CreateBlockOrderExecuteV1RespLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetBlockOrderOrdersV1RespItem) SetLegs(v []CreateBlockOrderExecuteV1RespLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetBlockOrderOrdersV1RespItem) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetLiquidity

`func (o *GetBlockOrderOrdersV1RespItem) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *GetBlockOrderOrdersV1RespItem) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *GetBlockOrderOrdersV1RespItem) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *GetBlockOrderOrdersV1RespItem) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *GetBlockOrderOrdersV1RespItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBlockOrderOrdersV1RespItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBlockOrderOrdersV1RespItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBlockOrderOrdersV1RespItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


