# MarginGetMarginAllOrderListV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContingencyType** | Pointer to **string** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]MarginCreateMarginOrderOcoV1RespOrdersInner**](MarginCreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarginGetMarginAllOrderListV1RespItem

`func NewMarginGetMarginAllOrderListV1RespItem() *MarginGetMarginAllOrderListV1RespItem`

NewMarginGetMarginAllOrderListV1RespItem instantiates a new MarginGetMarginAllOrderListV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginGetMarginAllOrderListV1RespItemWithDefaults

`func NewMarginGetMarginAllOrderListV1RespItemWithDefaults() *MarginGetMarginAllOrderListV1RespItem`

NewMarginGetMarginAllOrderListV1RespItemWithDefaults instantiates a new MarginGetMarginAllOrderListV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *MarginGetMarginAllOrderListV1RespItem) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *MarginGetMarginAllOrderListV1RespItem) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *MarginGetMarginAllOrderListV1RespItem) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginGetMarginAllOrderListV1RespItem) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginGetMarginAllOrderListV1RespItem) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginGetMarginAllOrderListV1RespItem) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *MarginGetMarginAllOrderListV1RespItem) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *MarginGetMarginAllOrderListV1RespItem) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *MarginGetMarginAllOrderListV1RespItem) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *MarginGetMarginAllOrderListV1RespItem) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *MarginGetMarginAllOrderListV1RespItem) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *MarginGetMarginAllOrderListV1RespItem) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *MarginGetMarginAllOrderListV1RespItem) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *MarginGetMarginAllOrderListV1RespItem) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *MarginGetMarginAllOrderListV1RespItem) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *MarginGetMarginAllOrderListV1RespItem) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MarginGetMarginAllOrderListV1RespItem) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MarginGetMarginAllOrderListV1RespItem) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrders

`func (o *MarginGetMarginAllOrderListV1RespItem) GetOrders() []MarginCreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetOrdersOk() (*[]MarginCreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MarginGetMarginAllOrderListV1RespItem) SetOrders(v []MarginCreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MarginGetMarginAllOrderListV1RespItem) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginGetMarginAllOrderListV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginGetMarginAllOrderListV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginGetMarginAllOrderListV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *MarginGetMarginAllOrderListV1RespItem) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *MarginGetMarginAllOrderListV1RespItem) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *MarginGetMarginAllOrderListV1RespItem) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *MarginGetMarginAllOrderListV1RespItem) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


