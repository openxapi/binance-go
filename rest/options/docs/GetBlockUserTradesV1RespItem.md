# GetBlockUserTradesV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**CrossType** | Pointer to **string** |  | [optional] 
**Legs** | Pointer to [**[]GetBlockUserTradesV1RespItemLegsInner**](GetBlockUserTradesV1RespItemLegsInner.md) |  | [optional] 
**ParentOrderId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBlockUserTradesV1RespItem

`func NewGetBlockUserTradesV1RespItem() *GetBlockUserTradesV1RespItem`

NewGetBlockUserTradesV1RespItem instantiates a new GetBlockUserTradesV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockUserTradesV1RespItemWithDefaults

`func NewGetBlockUserTradesV1RespItemWithDefaults() *GetBlockUserTradesV1RespItem`

NewGetBlockUserTradesV1RespItemWithDefaults instantiates a new GetBlockUserTradesV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *GetBlockUserTradesV1RespItem) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *GetBlockUserTradesV1RespItem) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *GetBlockUserTradesV1RespItem) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *GetBlockUserTradesV1RespItem) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetCrossType

`func (o *GetBlockUserTradesV1RespItem) GetCrossType() string`

GetCrossType returns the CrossType field if non-nil, zero value otherwise.

### GetCrossTypeOk

`func (o *GetBlockUserTradesV1RespItem) GetCrossTypeOk() (*string, bool)`

GetCrossTypeOk returns a tuple with the CrossType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossType

`func (o *GetBlockUserTradesV1RespItem) SetCrossType(v string)`

SetCrossType sets CrossType field to given value.

### HasCrossType

`func (o *GetBlockUserTradesV1RespItem) HasCrossType() bool`

HasCrossType returns a boolean if a field has been set.

### GetLegs

`func (o *GetBlockUserTradesV1RespItem) GetLegs() []GetBlockUserTradesV1RespItemLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetBlockUserTradesV1RespItem) GetLegsOk() (*[]GetBlockUserTradesV1RespItemLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetBlockUserTradesV1RespItem) SetLegs(v []GetBlockUserTradesV1RespItemLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetBlockUserTradesV1RespItem) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetParentOrderId

`func (o *GetBlockUserTradesV1RespItem) GetParentOrderId() string`

GetParentOrderId returns the ParentOrderId field if non-nil, zero value otherwise.

### GetParentOrderIdOk

`func (o *GetBlockUserTradesV1RespItem) GetParentOrderIdOk() (*string, bool)`

GetParentOrderIdOk returns a tuple with the ParentOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrderId

`func (o *GetBlockUserTradesV1RespItem) SetParentOrderId(v string)`

SetParentOrderId sets ParentOrderId field to given value.

### HasParentOrderId

`func (o *GetBlockUserTradesV1RespItem) HasParentOrderId() bool`

HasParentOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


