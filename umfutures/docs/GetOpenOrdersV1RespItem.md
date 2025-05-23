# GetOpenOrdersV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatePrice** | Pointer to **string** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**CumQuote** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**OrigType** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**PriceRate** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOpenOrdersV1RespItem

`func NewGetOpenOrdersV1RespItem() *GetOpenOrdersV1RespItem`

NewGetOpenOrdersV1RespItem instantiates a new GetOpenOrdersV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOpenOrdersV1RespItemWithDefaults

`func NewGetOpenOrdersV1RespItemWithDefaults() *GetOpenOrdersV1RespItem`

NewGetOpenOrdersV1RespItemWithDefaults instantiates a new GetOpenOrdersV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatePrice

`func (o *GetOpenOrdersV1RespItem) GetActivatePrice() string`

GetActivatePrice returns the ActivatePrice field if non-nil, zero value otherwise.

### GetActivatePriceOk

`func (o *GetOpenOrdersV1RespItem) GetActivatePriceOk() (*string, bool)`

GetActivatePriceOk returns a tuple with the ActivatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatePrice

`func (o *GetOpenOrdersV1RespItem) SetActivatePrice(v string)`

SetActivatePrice sets ActivatePrice field to given value.

### HasActivatePrice

`func (o *GetOpenOrdersV1RespItem) HasActivatePrice() bool`

HasActivatePrice returns a boolean if a field has been set.

### GetAvgPrice

`func (o *GetOpenOrdersV1RespItem) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *GetOpenOrdersV1RespItem) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *GetOpenOrdersV1RespItem) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *GetOpenOrdersV1RespItem) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *GetOpenOrdersV1RespItem) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *GetOpenOrdersV1RespItem) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *GetOpenOrdersV1RespItem) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *GetOpenOrdersV1RespItem) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetClosePosition

`func (o *GetOpenOrdersV1RespItem) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *GetOpenOrdersV1RespItem) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *GetOpenOrdersV1RespItem) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *GetOpenOrdersV1RespItem) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetCumQuote

`func (o *GetOpenOrdersV1RespItem) GetCumQuote() string`

GetCumQuote returns the CumQuote field if non-nil, zero value otherwise.

### GetCumQuoteOk

`func (o *GetOpenOrdersV1RespItem) GetCumQuoteOk() (*string, bool)`

GetCumQuoteOk returns a tuple with the CumQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQuote

`func (o *GetOpenOrdersV1RespItem) SetCumQuote(v string)`

SetCumQuote sets CumQuote field to given value.

### HasCumQuote

`func (o *GetOpenOrdersV1RespItem) HasCumQuote() bool`

HasCumQuote returns a boolean if a field has been set.

### GetExecutedQty

`func (o *GetOpenOrdersV1RespItem) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *GetOpenOrdersV1RespItem) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *GetOpenOrdersV1RespItem) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *GetOpenOrdersV1RespItem) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *GetOpenOrdersV1RespItem) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *GetOpenOrdersV1RespItem) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *GetOpenOrdersV1RespItem) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *GetOpenOrdersV1RespItem) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOpenOrdersV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOpenOrdersV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOpenOrdersV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOpenOrdersV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *GetOpenOrdersV1RespItem) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *GetOpenOrdersV1RespItem) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *GetOpenOrdersV1RespItem) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *GetOpenOrdersV1RespItem) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetOrigType

`func (o *GetOpenOrdersV1RespItem) GetOrigType() string`

GetOrigType returns the OrigType field if non-nil, zero value otherwise.

### GetOrigTypeOk

`func (o *GetOpenOrdersV1RespItem) GetOrigTypeOk() (*string, bool)`

GetOrigTypeOk returns a tuple with the OrigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigType

`func (o *GetOpenOrdersV1RespItem) SetOrigType(v string)`

SetOrigType sets OrigType field to given value.

### HasOrigType

`func (o *GetOpenOrdersV1RespItem) HasOrigType() bool`

HasOrigType returns a boolean if a field has been set.

### GetPositionSide

`func (o *GetOpenOrdersV1RespItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *GetOpenOrdersV1RespItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *GetOpenOrdersV1RespItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *GetOpenOrdersV1RespItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *GetOpenOrdersV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetOpenOrdersV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetOpenOrdersV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetOpenOrdersV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceMatch

`func (o *GetOpenOrdersV1RespItem) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *GetOpenOrdersV1RespItem) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *GetOpenOrdersV1RespItem) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *GetOpenOrdersV1RespItem) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetPriceProtect

`func (o *GetOpenOrdersV1RespItem) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *GetOpenOrdersV1RespItem) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *GetOpenOrdersV1RespItem) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *GetOpenOrdersV1RespItem) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetPriceRate

`func (o *GetOpenOrdersV1RespItem) GetPriceRate() string`

GetPriceRate returns the PriceRate field if non-nil, zero value otherwise.

### GetPriceRateOk

`func (o *GetOpenOrdersV1RespItem) GetPriceRateOk() (*string, bool)`

GetPriceRateOk returns a tuple with the PriceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRate

`func (o *GetOpenOrdersV1RespItem) SetPriceRate(v string)`

SetPriceRate sets PriceRate field to given value.

### HasPriceRate

`func (o *GetOpenOrdersV1RespItem) HasPriceRate() bool`

HasPriceRate returns a boolean if a field has been set.

### GetReduceOnly

`func (o *GetOpenOrdersV1RespItem) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *GetOpenOrdersV1RespItem) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *GetOpenOrdersV1RespItem) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *GetOpenOrdersV1RespItem) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *GetOpenOrdersV1RespItem) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *GetOpenOrdersV1RespItem) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *GetOpenOrdersV1RespItem) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *GetOpenOrdersV1RespItem) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *GetOpenOrdersV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetOpenOrdersV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetOpenOrdersV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetOpenOrdersV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *GetOpenOrdersV1RespItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOpenOrdersV1RespItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOpenOrdersV1RespItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOpenOrdersV1RespItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *GetOpenOrdersV1RespItem) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *GetOpenOrdersV1RespItem) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *GetOpenOrdersV1RespItem) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *GetOpenOrdersV1RespItem) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *GetOpenOrdersV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetOpenOrdersV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetOpenOrdersV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetOpenOrdersV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *GetOpenOrdersV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetOpenOrdersV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetOpenOrdersV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetOpenOrdersV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeInForce

`func (o *GetOpenOrdersV1RespItem) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *GetOpenOrdersV1RespItem) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *GetOpenOrdersV1RespItem) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *GetOpenOrdersV1RespItem) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *GetOpenOrdersV1RespItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOpenOrdersV1RespItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOpenOrdersV1RespItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOpenOrdersV1RespItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetOpenOrdersV1RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetOpenOrdersV1RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetOpenOrdersV1RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetOpenOrdersV1RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetWorkingType

`func (o *GetOpenOrdersV1RespItem) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *GetOpenOrdersV1RespItem) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *GetOpenOrdersV1RespItem) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *GetOpenOrdersV1RespItem) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


