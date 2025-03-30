# MarginCreateMarginOrderOtoV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContingencyType** | Pointer to **string** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrderReports** | Pointer to [**[]MarginCreateMarginOrderOtoV1RespOrderReportsInner**](MarginCreateMarginOrderOtoV1RespOrderReportsInner.md) |  | [optional] 
**Orders** | Pointer to [**[]MarginCreateMarginOrderOcoV1RespOrdersInner**](MarginCreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarginCreateMarginOrderOtoV1Resp

`func NewMarginCreateMarginOrderOtoV1Resp() *MarginCreateMarginOrderOtoV1Resp`

NewMarginCreateMarginOrderOtoV1Resp instantiates a new MarginCreateMarginOrderOtoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginCreateMarginOrderOtoV1RespWithDefaults

`func NewMarginCreateMarginOrderOtoV1RespWithDefaults() *MarginCreateMarginOrderOtoV1Resp`

NewMarginCreateMarginOrderOtoV1RespWithDefaults instantiates a new MarginCreateMarginOrderOtoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *MarginCreateMarginOrderOtoV1Resp) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginCreateMarginOrderOtoV1Resp) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginCreateMarginOrderOtoV1Resp) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginCreateMarginOrderOtoV1Resp) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginCreateMarginOrderOtoV1Resp) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginCreateMarginOrderOtoV1Resp) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginCreateMarginOrderOtoV1Resp) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginCreateMarginOrderOtoV1Resp) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginCreateMarginOrderOtoV1Resp) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginCreateMarginOrderOtoV1Resp) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginCreateMarginOrderOtoV1Resp) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginCreateMarginOrderOtoV1Resp) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginCreateMarginOrderOtoV1Resp) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginCreateMarginOrderOtoV1Resp) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginCreateMarginOrderOtoV1Resp) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *MarginCreateMarginOrderOtoV1Resp) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginCreateMarginOrderOtoV1Resp) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginCreateMarginOrderOtoV1Resp) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrderReports

`func (o *MarginCreateMarginOrderOtoV1Resp) GetOrderReports() []MarginCreateMarginOrderOtoV1RespOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetOrderReportsOk() (*[]MarginCreateMarginOrderOtoV1RespOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginCreateMarginOrderOtoV1Resp) SetOrderReports(v []MarginCreateMarginOrderOtoV1RespOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *MarginCreateMarginOrderOtoV1Resp) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.

### GetOrders

`func (o *MarginCreateMarginOrderOtoV1Resp) GetOrders() []MarginCreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetOrdersOk() (*[]MarginCreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginCreateMarginOrderOtoV1Resp) SetOrders(v []MarginCreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginCreateMarginOrderOtoV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginCreateMarginOrderOtoV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginCreateMarginOrderOtoV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginCreateMarginOrderOtoV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginCreateMarginOrderOtoV1Resp) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginCreateMarginOrderOtoV1Resp) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginCreateMarginOrderOtoV1Resp) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginCreateMarginOrderOtoV1Resp) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


