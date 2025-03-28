# SpotGetAllOrderListV3RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]SpotCreateOrderListOcoV3RespOrdersInner**](SpotCreateOrderListOcoV3RespOrdersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpotGetAllOrderListV3RespItem

`func NewSpotGetAllOrderListV3RespItem() *SpotGetAllOrderListV3RespItem`

NewSpotGetAllOrderListV3RespItem instantiates a new SpotGetAllOrderListV3RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetAllOrderListV3RespItemWithDefaults

`func NewSpotGetAllOrderListV3RespItemWithDefaults() *SpotGetAllOrderListV3RespItem`

NewSpotGetAllOrderListV3RespItemWithDefaults instantiates a new SpotGetAllOrderListV3RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *SpotGetAllOrderListV3RespItem) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *SpotGetAllOrderListV3RespItem) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *SpotGetAllOrderListV3RespItem) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *SpotGetAllOrderListV3RespItem) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *SpotGetAllOrderListV3RespItem) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *SpotGetAllOrderListV3RespItem) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *SpotGetAllOrderListV3RespItem) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *SpotGetAllOrderListV3RespItem) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *SpotGetAllOrderListV3RespItem) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *SpotGetAllOrderListV3RespItem) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *SpotGetAllOrderListV3RespItem) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *SpotGetAllOrderListV3RespItem) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *SpotGetAllOrderListV3RespItem) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *SpotGetAllOrderListV3RespItem) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *SpotGetAllOrderListV3RespItem) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *SpotGetAllOrderListV3RespItem) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *SpotGetAllOrderListV3RespItem) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *SpotGetAllOrderListV3RespItem) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *SpotGetAllOrderListV3RespItem) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *SpotGetAllOrderListV3RespItem) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrders

`func (o *SpotGetAllOrderListV3RespItem) GetOrders() []SpotCreateOrderListOcoV3RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *SpotGetAllOrderListV3RespItem) GetOrdersOk() (*[]SpotCreateOrderListOcoV3RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *SpotGetAllOrderListV3RespItem) SetOrders(v []SpotCreateOrderListOcoV3RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *SpotGetAllOrderListV3RespItem) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotGetAllOrderListV3RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotGetAllOrderListV3RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotGetAllOrderListV3RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotGetAllOrderListV3RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *SpotGetAllOrderListV3RespItem) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *SpotGetAllOrderListV3RespItem) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *SpotGetAllOrderListV3RespItem) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *SpotGetAllOrderListV3RespItem) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


