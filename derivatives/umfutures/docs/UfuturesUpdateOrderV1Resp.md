# UfuturesUpdateOrderV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**CumBase** | Pointer to **string** |  | [optional] 
**CumQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**OrigType** | Pointer to **string** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesUpdateOrderV1Resp

`func NewUfuturesUpdateOrderV1Resp() *UfuturesUpdateOrderV1Resp`

NewUfuturesUpdateOrderV1Resp instantiates a new UfuturesUpdateOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesUpdateOrderV1RespWithDefaults

`func NewUfuturesUpdateOrderV1RespWithDefaults() *UfuturesUpdateOrderV1Resp`

NewUfuturesUpdateOrderV1RespWithDefaults instantiates a new UfuturesUpdateOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *UfuturesUpdateOrderV1Resp) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *UfuturesUpdateOrderV1Resp) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *UfuturesUpdateOrderV1Resp) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *UfuturesUpdateOrderV1Resp) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *UfuturesUpdateOrderV1Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *UfuturesUpdateOrderV1Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *UfuturesUpdateOrderV1Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *UfuturesUpdateOrderV1Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetClosePosition

`func (o *UfuturesUpdateOrderV1Resp) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *UfuturesUpdateOrderV1Resp) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *UfuturesUpdateOrderV1Resp) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *UfuturesUpdateOrderV1Resp) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetCumBase

`func (o *UfuturesUpdateOrderV1Resp) GetCumBase() string`

GetCumBase returns the CumBase field if non-nil, zero value otherwise.

### GetCumBaseOk

`func (o *UfuturesUpdateOrderV1Resp) GetCumBaseOk() (*string, bool)`

GetCumBaseOk returns a tuple with the CumBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumBase

`func (o *UfuturesUpdateOrderV1Resp) SetCumBase(v string)`

SetCumBase sets CumBase field to given value.

### HasCumBase

`func (o *UfuturesUpdateOrderV1Resp) HasCumBase() bool`

HasCumBase returns a boolean if a field has been set.

### GetCumQty

`func (o *UfuturesUpdateOrderV1Resp) GetCumQty() string`

GetCumQty returns the CumQty field if non-nil, zero value otherwise.

### GetCumQtyOk

`func (o *UfuturesUpdateOrderV1Resp) GetCumQtyOk() (*string, bool)`

GetCumQtyOk returns a tuple with the CumQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQty

`func (o *UfuturesUpdateOrderV1Resp) SetCumQty(v string)`

SetCumQty sets CumQty field to given value.

### HasCumQty

`func (o *UfuturesUpdateOrderV1Resp) HasCumQty() bool`

HasCumQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *UfuturesUpdateOrderV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *UfuturesUpdateOrderV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *UfuturesUpdateOrderV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *UfuturesUpdateOrderV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *UfuturesUpdateOrderV1Resp) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *UfuturesUpdateOrderV1Resp) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *UfuturesUpdateOrderV1Resp) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *UfuturesUpdateOrderV1Resp) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetOrderId

`func (o *UfuturesUpdateOrderV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UfuturesUpdateOrderV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UfuturesUpdateOrderV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UfuturesUpdateOrderV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *UfuturesUpdateOrderV1Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *UfuturesUpdateOrderV1Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *UfuturesUpdateOrderV1Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *UfuturesUpdateOrderV1Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *UfuturesUpdateOrderV1Resp) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *UfuturesUpdateOrderV1Resp) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *UfuturesUpdateOrderV1Resp) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *UfuturesUpdateOrderV1Resp) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPair

`func (o *UfuturesUpdateOrderV1Resp) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *UfuturesUpdateOrderV1Resp) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *UfuturesUpdateOrderV1Resp) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *UfuturesUpdateOrderV1Resp) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetPositionSide

`func (o *UfuturesUpdateOrderV1Resp) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UfuturesUpdateOrderV1Resp) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UfuturesUpdateOrderV1Resp) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UfuturesUpdateOrderV1Resp) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *UfuturesUpdateOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UfuturesUpdateOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UfuturesUpdateOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UfuturesUpdateOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *UfuturesUpdateOrderV1Resp) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *UfuturesUpdateOrderV1Resp) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *UfuturesUpdateOrderV1Resp) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *UfuturesUpdateOrderV1Resp) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetPriceProtect

`func (o *UfuturesUpdateOrderV1Resp) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *UfuturesUpdateOrderV1Resp) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *UfuturesUpdateOrderV1Resp) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *UfuturesUpdateOrderV1Resp) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetReduceOnly

`func (o *UfuturesUpdateOrderV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *UfuturesUpdateOrderV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *UfuturesUpdateOrderV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *UfuturesUpdateOrderV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *UfuturesUpdateOrderV1Resp) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *UfuturesUpdateOrderV1Resp) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *UfuturesUpdateOrderV1Resp) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *UfuturesUpdateOrderV1Resp) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *UfuturesUpdateOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UfuturesUpdateOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UfuturesUpdateOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *UfuturesUpdateOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *UfuturesUpdateOrderV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UfuturesUpdateOrderV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UfuturesUpdateOrderV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UfuturesUpdateOrderV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *UfuturesUpdateOrderV1Resp) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *UfuturesUpdateOrderV1Resp) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *UfuturesUpdateOrderV1Resp) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *UfuturesUpdateOrderV1Resp) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesUpdateOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesUpdateOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesUpdateOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UfuturesUpdateOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *UfuturesUpdateOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UfuturesUpdateOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UfuturesUpdateOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UfuturesUpdateOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *UfuturesUpdateOrderV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UfuturesUpdateOrderV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UfuturesUpdateOrderV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UfuturesUpdateOrderV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *UfuturesUpdateOrderV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *UfuturesUpdateOrderV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *UfuturesUpdateOrderV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *UfuturesUpdateOrderV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *UfuturesUpdateOrderV1Resp) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *UfuturesUpdateOrderV1Resp) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *UfuturesUpdateOrderV1Resp) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *UfuturesUpdateOrderV1Resp) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


