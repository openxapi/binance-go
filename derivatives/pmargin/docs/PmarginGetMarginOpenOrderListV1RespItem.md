# PmarginGetMarginOpenOrderListV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]PmarginCreateMarginOrderOcoV1RespOrdersInner**](PmarginCreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginGetMarginOpenOrderListV1RespItem

`func NewPmarginGetMarginOpenOrderListV1RespItem() *PmarginGetMarginOpenOrderListV1RespItem`

NewPmarginGetMarginOpenOrderListV1RespItem instantiates a new PmarginGetMarginOpenOrderListV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetMarginOpenOrderListV1RespItemWithDefaults

`func NewPmarginGetMarginOpenOrderListV1RespItemWithDefaults() *PmarginGetMarginOpenOrderListV1RespItem`

NewPmarginGetMarginOpenOrderListV1RespItemWithDefaults instantiates a new PmarginGetMarginOpenOrderListV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrders

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetOrders() []PmarginCreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetOrdersOk() (*[]PmarginCreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetOrders(v []PmarginCreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PmarginGetMarginOpenOrderListV1RespItem) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PmarginGetMarginOpenOrderListV1RespItem) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *PmarginGetMarginOpenOrderListV1RespItem) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


