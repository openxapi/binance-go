# SpotDeleteOpenOrdersV3RespInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**TransactTime** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderReports** | Pointer to [**[]SpotDeleteOpenOrdersV3RespOrderItem**](SpotDeleteOpenOrdersV3RespOrderItem.md) |  | [optional] 
**Orders** | Pointer to [**[]CreateMarginOrderOcoV1RespOrdersInner**](CreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpotDeleteOpenOrdersV3RespInner

`func NewSpotDeleteOpenOrdersV3RespInner() *SpotDeleteOpenOrdersV3RespInner`

NewSpotDeleteOpenOrdersV3RespInner instantiates a new SpotDeleteOpenOrdersV3RespInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotDeleteOpenOrdersV3RespInnerWithDefaults

`func NewSpotDeleteOpenOrdersV3RespInnerWithDefaults() *SpotDeleteOpenOrdersV3RespInner`

NewSpotDeleteOpenOrdersV3RespInnerWithDefaults instantiates a new SpotDeleteOpenOrdersV3RespInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *SpotDeleteOpenOrdersV3RespInner) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *SpotDeleteOpenOrdersV3RespInner) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *SpotDeleteOpenOrdersV3RespInner) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *SpotDeleteOpenOrdersV3RespInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *SpotDeleteOpenOrdersV3RespInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *SpotDeleteOpenOrdersV3RespInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *SpotDeleteOpenOrdersV3RespInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *SpotDeleteOpenOrdersV3RespInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *SpotDeleteOpenOrdersV3RespInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *SpotDeleteOpenOrdersV3RespInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *SpotDeleteOpenOrdersV3RespInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SpotDeleteOpenOrdersV3RespInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SpotDeleteOpenOrdersV3RespInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *SpotDeleteOpenOrdersV3RespInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *SpotDeleteOpenOrdersV3RespInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *SpotDeleteOpenOrdersV3RespInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetSide

`func (o *SpotDeleteOpenOrdersV3RespInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *SpotDeleteOpenOrdersV3RespInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *SpotDeleteOpenOrdersV3RespInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *SpotDeleteOpenOrdersV3RespInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpotDeleteOpenOrdersV3RespInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpotDeleteOpenOrdersV3RespInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotDeleteOpenOrdersV3RespInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotDeleteOpenOrdersV3RespInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotDeleteOpenOrdersV3RespInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *SpotDeleteOpenOrdersV3RespInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *SpotDeleteOpenOrdersV3RespInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *SpotDeleteOpenOrdersV3RespInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTransactTime

`func (o *SpotDeleteOpenOrdersV3RespInner) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *SpotDeleteOpenOrdersV3RespInner) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *SpotDeleteOpenOrdersV3RespInner) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetType

`func (o *SpotDeleteOpenOrdersV3RespInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpotDeleteOpenOrdersV3RespInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpotDeleteOpenOrdersV3RespInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContingencyType

`func (o *SpotDeleteOpenOrdersV3RespInner) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *SpotDeleteOpenOrdersV3RespInner) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *SpotDeleteOpenOrdersV3RespInner) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *SpotDeleteOpenOrdersV3RespInner) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *SpotDeleteOpenOrdersV3RespInner) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *SpotDeleteOpenOrdersV3RespInner) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *SpotDeleteOpenOrdersV3RespInner) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *SpotDeleteOpenOrdersV3RespInner) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *SpotDeleteOpenOrdersV3RespInner) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *SpotDeleteOpenOrdersV3RespInner) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderReports

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrderReports() []SpotDeleteOpenOrdersV3RespOrderItem`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrderReportsOk() (*[]SpotDeleteOpenOrdersV3RespOrderItem, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *SpotDeleteOpenOrdersV3RespInner) SetOrderReports(v []SpotDeleteOpenOrdersV3RespOrderItem)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *SpotDeleteOpenOrdersV3RespInner) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.

### GetOrders

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrders() []CreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetOrdersOk() (*[]CreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *SpotDeleteOpenOrdersV3RespInner) SetOrders(v []CreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *SpotDeleteOpenOrdersV3RespInner) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTransactionTime

`func (o *SpotDeleteOpenOrdersV3RespInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *SpotDeleteOpenOrdersV3RespInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *SpotDeleteOpenOrdersV3RespInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *SpotDeleteOpenOrdersV3RespInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


