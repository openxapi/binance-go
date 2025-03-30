# MarginCreateMarginOrderOtocoV1Resp

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

### NewMarginCreateMarginOrderOtocoV1Resp

`func NewMarginCreateMarginOrderOtocoV1Resp() *MarginCreateMarginOrderOtocoV1Resp`

NewMarginCreateMarginOrderOtocoV1Resp instantiates a new MarginCreateMarginOrderOtocoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginCreateMarginOrderOtocoV1RespWithDefaults

`func NewMarginCreateMarginOrderOtocoV1RespWithDefaults() *MarginCreateMarginOrderOtocoV1Resp`

NewMarginCreateMarginOrderOtocoV1RespWithDefaults instantiates a new MarginCreateMarginOrderOtocoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrderReports

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetOrderReports() []MarginCreateMarginOrderOtoV1RespOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetOrderReportsOk() (*[]MarginCreateMarginOrderOtoV1RespOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetOrderReports(v []MarginCreateMarginOrderOtoV1RespOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.

### GetOrders

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetOrders() []MarginCreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetOrdersOk() (*[]MarginCreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetOrders(v []MarginCreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginCreateMarginOrderOtocoV1Resp) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginCreateMarginOrderOtocoV1Resp) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginCreateMarginOrderOtocoV1Resp) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


