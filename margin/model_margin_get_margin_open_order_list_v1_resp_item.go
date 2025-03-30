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

// checks if the MarginGetMarginOpenOrderListV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginOpenOrderListV1RespItem{}

// MarginGetMarginOpenOrderListV1RespItem struct for MarginGetMarginOpenOrderListV1RespItem
type MarginGetMarginOpenOrderListV1RespItem struct {
	ContingencyType *string `json:"contingencyType,omitempty"`
	IsIsolated *bool `json:"isIsolated,omitempty"`
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	ListOrderStatus *string `json:"listOrderStatus,omitempty"`
	ListStatusType *string `json:"listStatusType,omitempty"`
	OrderListId *int64 `json:"orderListId,omitempty"`
	Orders []MarginCreateMarginOrderOcoV1RespOrdersInner `json:"orders,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TransactionTime *int64 `json:"transactionTime,omitempty"`
}

// NewMarginGetMarginOpenOrderListV1RespItem instantiates a new MarginGetMarginOpenOrderListV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginOpenOrderListV1RespItem() *MarginGetMarginOpenOrderListV1RespItem {
	this := MarginGetMarginOpenOrderListV1RespItem{}
	return &this
}

// NewMarginGetMarginOpenOrderListV1RespItemWithDefaults instantiates a new MarginGetMarginOpenOrderListV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginOpenOrderListV1RespItemWithDefaults() *MarginGetMarginOpenOrderListV1RespItem {
	this := MarginGetMarginOpenOrderListV1RespItem{}
	return &this
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetContingencyType() string {
	if o == nil || IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetContingencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasContingencyType() bool {
	if o != nil && !IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetIsIsolated() bool {
	if o == nil || IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasIsIsolated() bool {
	if o != nil && !IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetListClientOrderId() string {
	if o == nil || IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasListClientOrderId() bool {
	if o != nil && !IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetListOrderStatus() string {
	if o == nil || IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetListOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasListOrderStatus() bool {
	if o != nil && !IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetListStatusType() string {
	if o == nil || IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetListStatusTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasListStatusType() bool {
	if o != nil && !IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetOrders() []MarginCreateMarginOrderOcoV1RespOrdersInner {
	if o == nil || IsNil(o.Orders) {
		var ret []MarginCreateMarginOrderOcoV1RespOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetOrdersOk() ([]MarginCreateMarginOrderOcoV1RespOrdersInner, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []MarginCreateMarginOrderOcoV1RespOrdersInner and assigns it to the Orders field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetOrders(v []MarginCreateMarginOrderOcoV1RespOrdersInner) {
	o.Orders = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetTransactionTime() int64 {
	if o == nil || IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *MarginGetMarginOpenOrderListV1RespItem) HasTransactionTime() bool {
	if o != nil && !IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *MarginGetMarginOpenOrderListV1RespItem) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

func (o MarginGetMarginOpenOrderListV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginOpenOrderListV1RespItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
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

type NullableMarginGetMarginOpenOrderListV1RespItem struct {
	value *MarginGetMarginOpenOrderListV1RespItem
	isSet bool
}

func (v NullableMarginGetMarginOpenOrderListV1RespItem) Get() *MarginGetMarginOpenOrderListV1RespItem {
	return v.value
}

func (v *NullableMarginGetMarginOpenOrderListV1RespItem) Set(val *MarginGetMarginOpenOrderListV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginOpenOrderListV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginOpenOrderListV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginOpenOrderListV1RespItem(val *MarginGetMarginOpenOrderListV1RespItem) *NullableMarginGetMarginOpenOrderListV1RespItem {
	return &NullableMarginGetMarginOpenOrderListV1RespItem{value: val, isSet: true}
}

func (v NullableMarginGetMarginOpenOrderListV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginOpenOrderListV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


