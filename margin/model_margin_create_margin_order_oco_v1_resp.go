/*
Binance Margin API

OpenAPI specification for Binance cryptocurrency exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginCreateMarginOrderOcoV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginCreateMarginOrderOcoV1Resp{}

// MarginCreateMarginOrderOcoV1Resp struct for MarginCreateMarginOrderOcoV1Resp
type MarginCreateMarginOrderOcoV1Resp struct {
	ContingencyType *string `json:"contingencyType,omitempty"`
	IsIsolated *bool `json:"isIsolated,omitempty"`
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	ListOrderStatus *string `json:"listOrderStatus,omitempty"`
	ListStatusType *string `json:"listStatusType,omitempty"`
	MarginBuyBorrowAmount *string `json:"marginBuyBorrowAmount,omitempty"`
	MarginBuyBorrowAsset *string `json:"marginBuyBorrowAsset,omitempty"`
	OrderListId *int64 `json:"orderListId,omitempty"`
	OrderReports []MarginCreateMarginOrderOcoV1RespOrderReportsInner `json:"orderReports,omitempty"`
	Orders []MarginCreateMarginOrderOcoV1RespOrdersInner `json:"orders,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TransactionTime *int64 `json:"transactionTime,omitempty"`
}

// NewMarginCreateMarginOrderOcoV1Resp instantiates a new MarginCreateMarginOrderOcoV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCreateMarginOrderOcoV1Resp() *MarginCreateMarginOrderOcoV1Resp {
	this := MarginCreateMarginOrderOcoV1Resp{}
	return &this
}

// NewMarginCreateMarginOrderOcoV1RespWithDefaults instantiates a new MarginCreateMarginOrderOcoV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCreateMarginOrderOcoV1RespWithDefaults() *MarginCreateMarginOrderOcoV1Resp {
	this := MarginCreateMarginOrderOcoV1Resp{}
	return &this
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetContingencyType() string {
	if o == nil || IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetContingencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasContingencyType() bool {
	if o != nil && !IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetIsIsolated() bool {
	if o == nil || IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasIsIsolated() bool {
	if o != nil && !IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetListClientOrderId() string {
	if o == nil || IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasListClientOrderId() bool {
	if o != nil && !IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetListOrderStatus() string {
	if o == nil || IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetListOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasListOrderStatus() bool {
	if o != nil && !IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetListStatusType() string {
	if o == nil || IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetListStatusTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasListStatusType() bool {
	if o != nil && !IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAmount() string {
	if o == nil || IsNil(o.MarginBuyBorrowAmount) {
		var ret string
		return ret
	}
	return *o.MarginBuyBorrowAmount
}

// GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAmountOk() (*string, bool) {
	if o == nil || IsNil(o.MarginBuyBorrowAmount) {
		return nil, false
	}
	return o.MarginBuyBorrowAmount, true
}

// HasMarginBuyBorrowAmount returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasMarginBuyBorrowAmount() bool {
	if o != nil && !IsNil(o.MarginBuyBorrowAmount) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAmount gets a reference to the given string and assigns it to the MarginBuyBorrowAmount field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetMarginBuyBorrowAmount(v string) {
	o.MarginBuyBorrowAmount = &v
}

// GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAsset() string {
	if o == nil || IsNil(o.MarginBuyBorrowAsset) {
		var ret string
		return ret
	}
	return *o.MarginBuyBorrowAsset
}

// GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetMarginBuyBorrowAssetOk() (*string, bool) {
	if o == nil || IsNil(o.MarginBuyBorrowAsset) {
		return nil, false
	}
	return o.MarginBuyBorrowAsset, true
}

// HasMarginBuyBorrowAsset returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasMarginBuyBorrowAsset() bool {
	if o != nil && !IsNil(o.MarginBuyBorrowAsset) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAsset gets a reference to the given string and assigns it to the MarginBuyBorrowAsset field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetMarginBuyBorrowAsset(v string) {
	o.MarginBuyBorrowAsset = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetOrderReports() []MarginCreateMarginOrderOcoV1RespOrderReportsInner {
	if o == nil || IsNil(o.OrderReports) {
		var ret []MarginCreateMarginOrderOcoV1RespOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetOrderReportsOk() ([]MarginCreateMarginOrderOcoV1RespOrderReportsInner, bool) {
	if o == nil || IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasOrderReports() bool {
	if o != nil && !IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []MarginCreateMarginOrderOcoV1RespOrderReportsInner and assigns it to the OrderReports field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetOrderReports(v []MarginCreateMarginOrderOcoV1RespOrderReportsInner) {
	o.OrderReports = v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetOrders() []MarginCreateMarginOrderOcoV1RespOrdersInner {
	if o == nil || IsNil(o.Orders) {
		var ret []MarginCreateMarginOrderOcoV1RespOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetOrdersOk() ([]MarginCreateMarginOrderOcoV1RespOrdersInner, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []MarginCreateMarginOrderOcoV1RespOrdersInner and assigns it to the Orders field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetOrders(v []MarginCreateMarginOrderOcoV1RespOrdersInner) {
	o.Orders = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderOcoV1Resp) GetTransactionTime() int64 {
	if o == nil || IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderOcoV1Resp) HasTransactionTime() bool {
	if o != nil && !IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *MarginCreateMarginOrderOcoV1Resp) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

func (o MarginCreateMarginOrderOcoV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginCreateMarginOrderOcoV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContingencyType) {
		toSerialize["contingencyType"] = o.ContingencyType
	}
	if !IsNil(o.IsIsolated) {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if !IsNil(o.ListClientOrderId) {
		toSerialize["listClientOrderId"] = o.ListClientOrderId
	}
	if !IsNil(o.ListOrderStatus) {
		toSerialize["listOrderStatus"] = o.ListOrderStatus
	}
	if !IsNil(o.ListStatusType) {
		toSerialize["listStatusType"] = o.ListStatusType
	}
	if !IsNil(o.MarginBuyBorrowAmount) {
		toSerialize["marginBuyBorrowAmount"] = o.MarginBuyBorrowAmount
	}
	if !IsNil(o.MarginBuyBorrowAsset) {
		toSerialize["marginBuyBorrowAsset"] = o.MarginBuyBorrowAsset
	}
	if !IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !IsNil(o.OrderReports) {
		toSerialize["orderReports"] = o.OrderReports
	}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TransactionTime) {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	return toSerialize, nil
}

type NullableMarginCreateMarginOrderOcoV1Resp struct {
	value *MarginCreateMarginOrderOcoV1Resp
	isSet bool
}

func (v NullableMarginCreateMarginOrderOcoV1Resp) Get() *MarginCreateMarginOrderOcoV1Resp {
	return v.value
}

func (v *NullableMarginCreateMarginOrderOcoV1Resp) Set(val *MarginCreateMarginOrderOcoV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCreateMarginOrderOcoV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCreateMarginOrderOcoV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCreateMarginOrderOcoV1Resp(val *MarginCreateMarginOrderOcoV1Resp) *NullableMarginCreateMarginOrderOcoV1Resp {
	return &NullableMarginCreateMarginOrderOcoV1Resp{value: val, isSet: true}
}

func (v NullableMarginCreateMarginOrderOcoV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCreateMarginOrderOcoV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


