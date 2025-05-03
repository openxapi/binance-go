# PmarginDeleteMarginAllOpenOrdersV1RespOrderList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrderReports** | Pointer to [**[]PmarginDeleteMarginAllOpenOrdersV1RespOrder**](PmarginDeleteMarginAllOpenOrdersV1RespOrder.md) |  | [optional] 
**Orders** | Pointer to [**[]CreateMarginOrderOcoV1RespOrdersInner**](CreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginDeleteMarginAllOpenOrdersV1RespOrderList

`func NewPmarginDeleteMarginAllOpenOrdersV1RespOrderList() *PmarginDeleteMarginAllOpenOrdersV1RespOrderList`

NewPmarginDeleteMarginAllOpenOrdersV1RespOrderList instantiates a new PmarginDeleteMarginAllOpenOrdersV1RespOrderList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginDeleteMarginAllOpenOrdersV1RespOrderListWithDefaults

`func NewPmarginDeleteMarginAllOpenOrdersV1RespOrderListWithDefaults() *PmarginDeleteMarginAllOpenOrdersV1RespOrderList`

NewPmarginDeleteMarginAllOpenOrdersV1RespOrderListWithDefaults instantiates a new PmarginDeleteMarginAllOpenOrdersV1RespOrderList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrderReports

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetOrderReports() []PmarginDeleteMarginAllOpenOrdersV1RespOrder`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetOrderReportsOk() (*[]PmarginDeleteMarginAllOpenOrdersV1RespOrder, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetOrderReports(v []PmarginDeleteMarginAllOpenOrdersV1RespOrder)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.

### GetOrders

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetOrders() []CreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetOrdersOk() (*[]CreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetOrders(v []CreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *PmarginDeleteMarginAllOpenOrdersV1RespOrderList) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


