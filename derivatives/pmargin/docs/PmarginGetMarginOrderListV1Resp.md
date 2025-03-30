# PmarginGetMarginOrderListV1Resp

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

### NewPmarginGetMarginOrderListV1Resp

`func NewPmarginGetMarginOrderListV1Resp() *PmarginGetMarginOrderListV1Resp`

NewPmarginGetMarginOrderListV1Resp instantiates a new PmarginGetMarginOrderListV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetMarginOrderListV1RespWithDefaults

`func NewPmarginGetMarginOrderListV1RespWithDefaults() *PmarginGetMarginOrderListV1Resp`

NewPmarginGetMarginOrderListV1RespWithDefaults instantiates a new PmarginGetMarginOrderListV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *PmarginGetMarginOrderListV1Resp) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *PmarginGetMarginOrderListV1Resp) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *PmarginGetMarginOrderListV1Resp) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *PmarginGetMarginOrderListV1Resp) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *PmarginGetMarginOrderListV1Resp) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *PmarginGetMarginOrderListV1Resp) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *PmarginGetMarginOrderListV1Resp) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *PmarginGetMarginOrderListV1Resp) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *PmarginGetMarginOrderListV1Resp) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *PmarginGetMarginOrderListV1Resp) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *PmarginGetMarginOrderListV1Resp) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *PmarginGetMarginOrderListV1Resp) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *PmarginGetMarginOrderListV1Resp) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *PmarginGetMarginOrderListV1Resp) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *PmarginGetMarginOrderListV1Resp) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *PmarginGetMarginOrderListV1Resp) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetOrderListId

`func (o *PmarginGetMarginOrderListV1Resp) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *PmarginGetMarginOrderListV1Resp) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *PmarginGetMarginOrderListV1Resp) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *PmarginGetMarginOrderListV1Resp) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrders

`func (o *PmarginGetMarginOrderListV1Resp) GetOrders() []PmarginCreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PmarginGetMarginOrderListV1Resp) GetOrdersOk() (*[]PmarginCreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PmarginGetMarginOrderListV1Resp) SetOrders(v []PmarginCreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PmarginGetMarginOrderListV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginGetMarginOrderListV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginGetMarginOrderListV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginGetMarginOrderListV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginGetMarginOrderListV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PmarginGetMarginOrderListV1Resp) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PmarginGetMarginOrderListV1Resp) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PmarginGetMarginOrderListV1Resp) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *PmarginGetMarginOrderListV1Resp) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


