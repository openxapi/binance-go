/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotDeleteOrderListV3Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotDeleteOrderListV3Resp{}

// SpotDeleteOrderListV3Resp struct for SpotDeleteOrderListV3Resp
type SpotDeleteOrderListV3Resp struct {
	ContingencyType *string `json:"contingencyType,omitempty"`
	ListClientOrderId *string `json:"listClientOrderId,omitempty"`
	ListOrderStatus *string `json:"listOrderStatus,omitempty"`
	ListStatusType *string `json:"listStatusType,omitempty"`
	OrderListId *int64 `json:"orderListId,omitempty"`
	OrderReports []SpotDeleteOrderListV3RespOrderReportsInner `json:"orderReports,omitempty"`
	Orders []SpotCreateOrderListOcoV3RespOrdersInner `json:"orders,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TransactionTime *int64 `json:"transactionTime,omitempty"`
}

// NewSpotDeleteOrderListV3Resp instantiates a new SpotDeleteOrderListV3Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotDeleteOrderListV3Resp() *SpotDeleteOrderListV3Resp {
	this := SpotDeleteOrderListV3Resp{}
	return &this
}

// NewSpotDeleteOrderListV3RespWithDefaults instantiates a new SpotDeleteOrderListV3Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotDeleteOrderListV3RespWithDefaults() *SpotDeleteOrderListV3Resp {
	this := SpotDeleteOrderListV3Resp{}
	return &this
}

// GetContingencyType returns the ContingencyType field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetContingencyType() string {
	if o == nil || IsNil(o.ContingencyType) {
		var ret string
		return ret
	}
	return *o.ContingencyType
}

// GetContingencyTypeOk returns a tuple with the ContingencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetContingencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContingencyType) {
		return nil, false
	}
	return o.ContingencyType, true
}

// HasContingencyType returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasContingencyType() bool {
	if o != nil && !IsNil(o.ContingencyType) {
		return true
	}

	return false
}

// SetContingencyType gets a reference to the given string and assigns it to the ContingencyType field.
func (o *SpotDeleteOrderListV3Resp) SetContingencyType(v string) {
	o.ContingencyType = &v
}

// GetListClientOrderId returns the ListClientOrderId field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetListClientOrderId() string {
	if o == nil || IsNil(o.ListClientOrderId) {
		var ret string
		return ret
	}
	return *o.ListClientOrderId
}

// GetListClientOrderIdOk returns a tuple with the ListClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetListClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListClientOrderId) {
		return nil, false
	}
	return o.ListClientOrderId, true
}

// HasListClientOrderId returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasListClientOrderId() bool {
	if o != nil && !IsNil(o.ListClientOrderId) {
		return true
	}

	return false
}

// SetListClientOrderId gets a reference to the given string and assigns it to the ListClientOrderId field.
func (o *SpotDeleteOrderListV3Resp) SetListClientOrderId(v string) {
	o.ListClientOrderId = &v
}

// GetListOrderStatus returns the ListOrderStatus field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetListOrderStatus() string {
	if o == nil || IsNil(o.ListOrderStatus) {
		var ret string
		return ret
	}
	return *o.ListOrderStatus
}

// GetListOrderStatusOk returns a tuple with the ListOrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetListOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ListOrderStatus) {
		return nil, false
	}
	return o.ListOrderStatus, true
}

// HasListOrderStatus returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasListOrderStatus() bool {
	if o != nil && !IsNil(o.ListOrderStatus) {
		return true
	}

	return false
}

// SetListOrderStatus gets a reference to the given string and assigns it to the ListOrderStatus field.
func (o *SpotDeleteOrderListV3Resp) SetListOrderStatus(v string) {
	o.ListOrderStatus = &v
}

// GetListStatusType returns the ListStatusType field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetListStatusType() string {
	if o == nil || IsNil(o.ListStatusType) {
		var ret string
		return ret
	}
	return *o.ListStatusType
}

// GetListStatusTypeOk returns a tuple with the ListStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetListStatusTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ListStatusType) {
		return nil, false
	}
	return o.ListStatusType, true
}

// HasListStatusType returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasListStatusType() bool {
	if o != nil && !IsNil(o.ListStatusType) {
		return true
	}

	return false
}

// SetListStatusType gets a reference to the given string and assigns it to the ListStatusType field.
func (o *SpotDeleteOrderListV3Resp) SetListStatusType(v string) {
	o.ListStatusType = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *SpotDeleteOrderListV3Resp) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrderReports returns the OrderReports field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetOrderReports() []SpotDeleteOrderListV3RespOrderReportsInner {
	if o == nil || IsNil(o.OrderReports) {
		var ret []SpotDeleteOrderListV3RespOrderReportsInner
		return ret
	}
	return o.OrderReports
}

// GetOrderReportsOk returns a tuple with the OrderReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetOrderReportsOk() ([]SpotDeleteOrderListV3RespOrderReportsInner, bool) {
	if o == nil || IsNil(o.OrderReports) {
		return nil, false
	}
	return o.OrderReports, true
}

// HasOrderReports returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasOrderReports() bool {
	if o != nil && !IsNil(o.OrderReports) {
		return true
	}

	return false
}

// SetOrderReports gets a reference to the given []SpotDeleteOrderListV3RespOrderReportsInner and assigns it to the OrderReports field.
func (o *SpotDeleteOrderListV3Resp) SetOrderReports(v []SpotDeleteOrderListV3RespOrderReportsInner) {
	o.OrderReports = v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetOrders() []SpotCreateOrderListOcoV3RespOrdersInner {
	if o == nil || IsNil(o.Orders) {
		var ret []SpotCreateOrderListOcoV3RespOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetOrdersOk() ([]SpotCreateOrderListOcoV3RespOrdersInner, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []SpotCreateOrderListOcoV3RespOrdersInner and assigns it to the Orders field.
func (o *SpotDeleteOrderListV3Resp) SetOrders(v []SpotCreateOrderListOcoV3RespOrdersInner) {
	o.Orders = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SpotDeleteOrderListV3Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *SpotDeleteOrderListV3Resp) GetTransactionTime() int64 {
	if o == nil || IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotDeleteOrderListV3Resp) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *SpotDeleteOrderListV3Resp) HasTransactionTime() bool {
	if o != nil && !IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *SpotDeleteOrderListV3Resp) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

func (o SpotDeleteOrderListV3Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotDeleteOrderListV3Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContingencyType) {
		toSerialize["contingencyType"] = o.ContingencyType
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

type NullableSpotDeleteOrderListV3Resp struct {
	value *SpotDeleteOrderListV3Resp
	isSet bool
}

func (v NullableSpotDeleteOrderListV3Resp) Get() *SpotDeleteOrderListV3Resp {
	return v.value
}

func (v *NullableSpotDeleteOrderListV3Resp) Set(val *SpotDeleteOrderListV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotDeleteOrderListV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotDeleteOrderListV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotDeleteOrderListV3Resp(val *SpotDeleteOrderListV3Resp) *NullableSpotDeleteOrderListV3Resp {
	return &NullableSpotDeleteOrderListV3Resp{value: val, isSet: true}
}

func (v NullableSpotDeleteOrderListV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotDeleteOrderListV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


