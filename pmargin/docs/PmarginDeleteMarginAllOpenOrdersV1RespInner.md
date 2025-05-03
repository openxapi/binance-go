# PmarginDeleteMarginAllOpenOrdersV1RespInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**IcebergQty** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderReports** | Pointer to [**[]PmarginDeleteMarginAllOpenOrdersV1RespOrder**](PmarginDeleteMarginAllOpenOrdersV1RespOrder.md) |  | [optional] 
**Orders** | Pointer to [**[]CreateMarginOrderOcoV1RespOrdersInner**](CreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginDeleteMarginAllOpenOrdersV1RespInner

`func NewPmarginDeleteMarginAllOpenOrdersV1RespInner() *PmarginDeleteMarginAllOpenOrdersV1RespInner`

NewPmarginDeleteMarginAllOpenOrdersV1RespInner instantiates a new PmarginDeleteMarginAllOpenOrdersV1RespInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginDeleteMarginAllOpenOrdersV1RespInnerWithDefaults

`func NewPmarginDeleteMarginAllOpenOrdersV1RespInnerWithDefaults() *PmarginDeleteMarginAllOpenOrdersV1RespInner`

NewPmarginDeleteMarginAllOpenOrdersV1RespInnerWithDefaults instantiates a new PmarginDeleteMarginAllOpenOrdersV1RespInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetIcebergQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetIcebergQty() string`

GetIcebergQty returns the IcebergQty field if non-nil, zero value otherwise.

### GetIcebergQtyOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetIcebergQtyOk() (*string, bool)`

GetIcebergQtyOk returns a tuple with the IcebergQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetIcebergQty(v string)`

SetIcebergQty sets IcebergQty field to given value.

### HasIcebergQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasIcebergQty() bool`

HasIcebergQty returns a boolean if a field has been set.

### GetOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSide

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStopPrice

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContingencyType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderReports

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrderReports() []PmarginDeleteMarginAllOpenOrdersV1RespOrder`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrderReportsOk() (*[]PmarginDeleteMarginAllOpenOrdersV1RespOrder, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetOrderReports(v []PmarginDeleteMarginAllOpenOrdersV1RespOrder)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.

### GetOrders

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrders() []CreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetOrdersOk() (*[]CreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetOrders(v []CreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


