# PmarginCreateMarginOrderOcoV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContingencyType** | Pointer to **string** |  | [optional] 
**ListClientOrderId** | Pointer to **string** |  | [optional] 
**ListOrderStatus** | Pointer to **string** |  | [optional] 
**ListStatusType** | Pointer to **string** |  | [optional] 
**MarginBuyBorrowAmount** | Pointer to **string** |  | [optional] 
**MarginBuyBorrowAsset** | Pointer to **string** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**OrderReports** | Pointer to [**[]PmarginCreateMarginOrderOcoV1RespOrderReportsInner**](PmarginCreateMarginOrderOcoV1RespOrderReportsInner.md) |  | [optional] 
**Orders** | Pointer to [**[]PmarginCreateMarginOrderOcoV1RespOrdersInner**](PmarginCreateMarginOrderOcoV1RespOrdersInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginCreateMarginOrderOcoV1Resp

`func NewPmarginCreateMarginOrderOcoV1Resp() *PmarginCreateMarginOrderOcoV1Resp`

NewPmarginCreateMarginOrderOcoV1Resp instantiates a new PmarginCreateMarginOrderOcoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginCreateMarginOrderOcoV1RespWithDefaults

`func NewPmarginCreateMarginOrderOcoV1RespWithDefaults() *PmarginCreateMarginOrderOcoV1Resp`

NewPmarginCreateMarginOrderOcoV1RespWithDefaults instantiates a new PmarginCreateMarginOrderOcoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContingencyType

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetContingencyType() string`

GetContingencyType returns the ContingencyType field if non-nil, zero value otherwise.

### GetContingencyTypeOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetContingencyTypeOk() (*string, bool)`

GetContingencyTypeOk returns a tuple with the ContingencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContingencyType

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetContingencyType(v string)`

SetContingencyType sets ContingencyType field to given value.

### HasContingencyType

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasContingencyType() bool`

HasContingencyType returns a boolean if a field has been set.

### GetListClientOrderId

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetListClientOrderId() string`

GetListClientOrderId returns the ListClientOrderId field if non-nil, zero value otherwise.

### GetListClientOrderIdOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetListClientOrderIdOk() (*string, bool)`

GetListClientOrderIdOk returns a tuple with the ListClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListClientOrderId

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetListClientOrderId(v string)`

SetListClientOrderId sets ListClientOrderId field to given value.

### HasListClientOrderId

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasListClientOrderId() bool`

HasListClientOrderId returns a boolean if a field has been set.

### GetListOrderStatus

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetListOrderStatus() string`

GetListOrderStatus returns the ListOrderStatus field if non-nil, zero value otherwise.

### GetListOrderStatusOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetListOrderStatusOk() (*string, bool)`

GetListOrderStatusOk returns a tuple with the ListOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOrderStatus

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetListOrderStatus(v string)`

SetListOrderStatus sets ListOrderStatus field to given value.

### HasListOrderStatus

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasListOrderStatus() bool`

HasListOrderStatus returns a boolean if a field has been set.

### GetListStatusType

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetListStatusType() string`

GetListStatusType returns the ListStatusType field if non-nil, zero value otherwise.

### GetListStatusTypeOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetListStatusTypeOk() (*string, bool)`

GetListStatusTypeOk returns a tuple with the ListStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatusType

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetListStatusType(v string)`

SetListStatusType sets ListStatusType field to given value.

### HasListStatusType

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasListStatusType() bool`

HasListStatusType returns a boolean if a field has been set.

### GetMarginBuyBorrowAmount

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAmount() string`

GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAmountOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAmountOk() (*string, bool)`

GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAmount

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetMarginBuyBorrowAmount(v string)`

SetMarginBuyBorrowAmount sets MarginBuyBorrowAmount field to given value.

### HasMarginBuyBorrowAmount

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasMarginBuyBorrowAmount() bool`

HasMarginBuyBorrowAmount returns a boolean if a field has been set.

### GetMarginBuyBorrowAsset

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAsset() string`

GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAssetOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAssetOk() (*string, bool)`

GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAsset

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetMarginBuyBorrowAsset(v string)`

SetMarginBuyBorrowAsset sets MarginBuyBorrowAsset field to given value.

### HasMarginBuyBorrowAsset

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasMarginBuyBorrowAsset() bool`

HasMarginBuyBorrowAsset returns a boolean if a field has been set.

### GetOrderListId

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetOrderReports

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetOrderReports() []PmarginCreateMarginOrderOcoV1RespOrderReportsInner`

GetOrderReports returns the OrderReports field if non-nil, zero value otherwise.

### GetOrderReportsOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetOrderReportsOk() (*[]PmarginCreateMarginOrderOcoV1RespOrderReportsInner, bool)`

GetOrderReportsOk returns a tuple with the OrderReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReports

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetOrderReports(v []PmarginCreateMarginOrderOcoV1RespOrderReportsInner)`

SetOrderReports sets OrderReports field to given value.

### HasOrderReports

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasOrderReports() bool`

HasOrderReports returns a boolean if a field has been set.

### GetOrders

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetOrders() []PmarginCreateMarginOrderOcoV1RespOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetOrdersOk() (*[]PmarginCreateMarginOrderOcoV1RespOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetOrders(v []PmarginCreateMarginOrderOcoV1RespOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PmarginCreateMarginOrderOcoV1Resp) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PmarginCreateMarginOrderOcoV1Resp) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *PmarginCreateMarginOrderOcoV1Resp) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


