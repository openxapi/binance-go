# MarginDeleteMarginOrderListV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContingencyType** | Pointer to **string** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrderReports** | Pointer to [**[]MarginDeleteMarginOrderListV1RespOrderReportsInner**](MarginDeleteMarginOrderListV1RespOrderReportsInner.md) |  | [optional] 
**Orders** | Pointer to [**[]MarginCreateMarginOrderOcoV1RespOrdersInner**](MarginCreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarginDeleteMarginOrderListV1Resp

`func NewMarginDeleteMarginOrderListV1Resp() *MarginDeleteMarginOrderListV1Resp`

NewMarginDeleteMarginOrderListV1Resp instantiates a new MarginDeleteMarginOrderListV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginDeleteMarginOrderListV1RespWithDefaults

`func NewMarginDeleteMarginOrderListV1RespWithDefaults() *MarginDeleteMarginOrderListV1Resp`

NewMarginDeleteMarginOrderListV1RespWithDefaults instantiates a new MarginDeleteMarginOrderListV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *MarginDeleteMarginOrderListV1Resp) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginDeleteMarginOrderListV1Resp) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginDeleteMarginOrderListV1Resp) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginDeleteMarginOrderListV1Resp) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginDeleteMarginOrderListV1Resp) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginDeleteMarginOrderListV1Resp) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginDeleteMarginOrderListV1Resp) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginDeleteMarginOrderListV1Resp) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginDeleteMarginOrderListV1Resp) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginDeleteMarginOrderListV1Resp) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginDeleteMarginOrderListV1Resp) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginDeleteMarginOrderListV1Resp) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginDeleteMarginOrderListV1Resp) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginDeleteMarginOrderListV1Resp) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginDeleteMarginOrderListV1Resp) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *MarginDeleteMarginOrderListV1Resp) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginDeleteMarginOrderListV1Resp) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginDeleteMarginOrderListV1Resp) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrderReports

`func (o *MarginDeleteMarginOrderListV1Resp) GetOrderReports() []MarginDeleteMarginOrderListV1RespOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetOrderReportsOk() (*[]MarginDeleteMarginOrderListV1RespOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *MarginDeleteMarginOrderListV1Resp) SetOrderReports(v []MarginDeleteMarginOrderListV1RespOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *MarginDeleteMarginOrderListV1Resp) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.

### GetOrders

`func (o *MarginDeleteMarginOrderListV1Resp) GetOrders() []MarginCreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetOrdersOk() (*[]MarginCreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginDeleteMarginOrderListV1Resp) SetOrders(v []MarginCreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginDeleteMarginOrderListV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginDeleteMarginOrderListV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginDeleteMarginOrderListV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginDeleteMarginOrderListV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginDeleteMarginOrderListV1Resp) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginDeleteMarginOrderListV1Resp) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginDeleteMarginOrderListV1Resp) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginDeleteMarginOrderListV1Resp) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


