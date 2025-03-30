# PmarginGetMarginOpenOrdersV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**IcebergQty** | Pointer to **string** |  | [optional] 
**IsWorking** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**PreventedMatchId** | Pointer to **map[string]interface{}** |  | [optional] 
**PreventedQuantity** | Pointer to **map[string]interface{}** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginGetMarginOpenOrdersV1RespItem

`func NewPmarginGetMarginOpenOrdersV1RespItem() *PmarginGetMarginOpenOrdersV1RespItem`

NewPmarginGetMarginOpenOrdersV1RespItem instantiates a new PmarginGetMarginOpenOrdersV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetMarginOpenOrdersV1RespItemWithDefaults

`func NewPmarginGetMarginOpenOrdersV1RespItemWithDefaults() *PmarginGetMarginOpenOrdersV1RespItem`

NewPmarginGetMarginOpenOrdersV1RespItemWithDefaults instantiates a new PmarginGetMarginOpenOrdersV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetIcebergQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetIsWorking

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetIsWorking() bool`

GetIsWorking returns the IsWorking field if non-nil, zero value otherwise.

### GetIsWorkingOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetIsWorkingOk() (*bool, bool)`

GetIsWorkingOk returns a tuple with the IsWorking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorking

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetIsWorking(v bool)`

SetIsWorking sets IsWorking field to given value.

### HasIsWorking

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasIsWorking() bool`

HasIsWorking returns a boolean if a field has been set.

### GetOrderId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetPreventedMatchId() map[string]interface{}`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetPreventedMatchIdOk() (*map[string]interface{}, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetPreventedMatchId(v map[string]interface{})`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### SetPreventedMatchIdNil

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetPreventedMatchIdNil(b bool)`

 SetPreventedMatchIdNil sets the value for PreventedMatchId to be an explicit nil

### UnsetPreventedMatchId
`func (o *PmarginGetMarginOpenOrdersV1RespItem) UnsetPreventedMatchId()`

UnsetPreventedMatchId ensures that no value is present for PreventedMatchId, not even an explicit nil
### GetPreventedQuantity

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetPreventedQuantity() map[string]interface{}`

GetPreventedQuantity returns the PreventedQuantity field if non-nil, zero value otherwise.

### GetPreventedQuantityOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetPreventedQuantityOk() (*map[string]interface{}, bool)`

GetPreventedQuantityOk returns a tuple with the PreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedQuantity

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetPreventedQuantity(v map[string]interface{})`

SetPreventedQuantity sets PreventedQuantity field to given value.

### HasPreventedQuantity

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasPreventedQuantity() bool`

HasPreventedQuantity returns a boolean if a field has been set.

### SetPreventedQuantityNil

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetPreventedQuantityNil(b bool)`

 SetPreventedQuantityNil sets the value for PreventedQuantity to be an explicit nil

### UnsetPreventedQuantity
`func (o *PmarginGetMarginOpenOrdersV1RespItem) UnsetPreventedQuantity()`

UnsetPreventedQuantity ensures that no value is present for PreventedQuantity, not even an explicit nil
### GetPrice

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PmarginGetMarginOpenOrdersV1RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PmarginGetMarginOpenOrdersV1RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PmarginGetMarginOpenOrdersV1RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


