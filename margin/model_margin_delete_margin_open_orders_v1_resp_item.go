/*
Binance Margin Trading API

OpenAPI specification for Binance exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginDeleteMarginOpenOrdersV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginDeleteMarginOpenOrdersV1RespItem{}

// MarginDeleteMarginOpenOrdersV1RespItem struct for MarginDeleteMarginOpenOrdersV1RespItem
type MarginDeleteMarginOpenOrdersV1RespItem struct {
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CummulativeQuoteQty *string `json:"cummulativeQuoteQty,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	IsIsolated *bool `json:"isIsolated,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrderListId *int64 `json:"orderListId,omitempty"`
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	Price *string `json:"price,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMarginDeleteMarginOpenOrdersV1RespItem instantiates a new MarginDeleteMarginOpenOrdersV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginDeleteMarginOpenOrdersV1RespItem() *MarginDeleteMarginOpenOrdersV1RespItem {
	this := MarginDeleteMarginOpenOrdersV1RespItem{}
	return &this
}

// NewMarginDeleteMarginOpenOrdersV1RespItemWithDefaults instantiates a new MarginDeleteMarginOpenOrdersV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginDeleteMarginOpenOrdersV1RespItemWithDefaults() *MarginDeleteMarginOpenOrdersV1RespItem {
	this := MarginDeleteMarginOpenOrdersV1RespItem{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetCummulativeQuoteQty() string {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasCummulativeQuoteQty() bool {
	if o != nil && !IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetIsIsolated() bool {
	if o == nil || IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasIsIsolated() bool {
	if o != nil && !IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrigClientOrderId() string {
	if o == nil || IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasOrigClientOrderId() bool {
	if o != nil && !IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginDeleteMarginOpenOrdersV1RespItem) SetType(v string) {
	o.Type = &v
}

func (o MarginDeleteMarginOpenOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginDeleteMarginOpenOrdersV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.CummulativeQuoteQty) {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if !IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !IsNil(o.IsIsolated) {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !IsNil(o.OrigClientOrderId) {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	if !IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMarginDeleteMarginOpenOrdersV1RespItem struct {
	value *MarginDeleteMarginOpenOrdersV1RespItem
	isSet bool
}

func (v NullableMarginDeleteMarginOpenOrdersV1RespItem) Get() *MarginDeleteMarginOpenOrdersV1RespItem {
	return v.value
}

func (v *NullableMarginDeleteMarginOpenOrdersV1RespItem) Set(val *MarginDeleteMarginOpenOrdersV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginDeleteMarginOpenOrdersV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginDeleteMarginOpenOrdersV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginDeleteMarginOpenOrdersV1RespItem(val *MarginDeleteMarginOpenOrdersV1RespItem) *NullableMarginDeleteMarginOpenOrdersV1RespItem {
	return &NullableMarginDeleteMarginOpenOrdersV1RespItem{value: val, isSet: true}
}

func (v NullableMarginDeleteMarginOpenOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginDeleteMarginOpenOrdersV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


