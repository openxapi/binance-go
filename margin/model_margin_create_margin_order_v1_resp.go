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

// checks if the MarginCreateMarginOrderV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginCreateMarginOrderV1Resp{}

// MarginCreateMarginOrderV1Resp struct for MarginCreateMarginOrderV1Resp
type MarginCreateMarginOrderV1Resp struct {
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CummulativeQuoteQty *string `json:"cummulativeQuoteQty,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	Fills []MarginCreateMarginOrderV1RespFillsInner `json:"fills,omitempty"`
	IsIsolated *bool `json:"isIsolated,omitempty"`
	MarginBuyBorrowAmount *int64 `json:"marginBuyBorrowAmount,omitempty"`
	MarginBuyBorrowAsset *string `json:"marginBuyBorrowAsset,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	Price *string `json:"price,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	TransactTime *int64 `json:"transactTime,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMarginCreateMarginOrderV1Resp instantiates a new MarginCreateMarginOrderV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCreateMarginOrderV1Resp() *MarginCreateMarginOrderV1Resp {
	this := MarginCreateMarginOrderV1Resp{}
	return &this
}

// NewMarginCreateMarginOrderV1RespWithDefaults instantiates a new MarginCreateMarginOrderV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCreateMarginOrderV1RespWithDefaults() *MarginCreateMarginOrderV1Resp {
	this := MarginCreateMarginOrderV1Resp{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *MarginCreateMarginOrderV1Resp) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetCummulativeQuoteQty() string {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasCummulativeQuoteQty() bool {
	if o != nil && !IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *MarginCreateMarginOrderV1Resp) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *MarginCreateMarginOrderV1Resp) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetFills returns the Fills field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetFills() []MarginCreateMarginOrderV1RespFillsInner {
	if o == nil || IsNil(o.Fills) {
		var ret []MarginCreateMarginOrderV1RespFillsInner
		return ret
	}
	return o.Fills
}

// GetFillsOk returns a tuple with the Fills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetFillsOk() ([]MarginCreateMarginOrderV1RespFillsInner, bool) {
	if o == nil || IsNil(o.Fills) {
		return nil, false
	}
	return o.Fills, true
}

// HasFills returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasFills() bool {
	if o != nil && !IsNil(o.Fills) {
		return true
	}

	return false
}

// SetFills gets a reference to the given []MarginCreateMarginOrderV1RespFillsInner and assigns it to the Fills field.
func (o *MarginCreateMarginOrderV1Resp) SetFills(v []MarginCreateMarginOrderV1RespFillsInner) {
	o.Fills = v
}

// GetIsIsolated returns the IsIsolated field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetIsIsolated() bool {
	if o == nil || IsNil(o.IsIsolated) {
		var ret bool
		return ret
	}
	return *o.IsIsolated
}

// GetIsIsolatedOk returns a tuple with the IsIsolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetIsIsolatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIsolated) {
		return nil, false
	}
	return o.IsIsolated, true
}

// HasIsIsolated returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasIsIsolated() bool {
	if o != nil && !IsNil(o.IsIsolated) {
		return true
	}

	return false
}

// SetIsIsolated gets a reference to the given bool and assigns it to the IsIsolated field.
func (o *MarginCreateMarginOrderV1Resp) SetIsIsolated(v bool) {
	o.IsIsolated = &v
}

// GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAmount() int64 {
	if o == nil || IsNil(o.MarginBuyBorrowAmount) {
		var ret int64
		return ret
	}
	return *o.MarginBuyBorrowAmount
}

// GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.MarginBuyBorrowAmount) {
		return nil, false
	}
	return o.MarginBuyBorrowAmount, true
}

// HasMarginBuyBorrowAmount returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasMarginBuyBorrowAmount() bool {
	if o != nil && !IsNil(o.MarginBuyBorrowAmount) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAmount gets a reference to the given int64 and assigns it to the MarginBuyBorrowAmount field.
func (o *MarginCreateMarginOrderV1Resp) SetMarginBuyBorrowAmount(v int64) {
	o.MarginBuyBorrowAmount = &v
}

// GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAsset() string {
	if o == nil || IsNil(o.MarginBuyBorrowAsset) {
		var ret string
		return ret
	}
	return *o.MarginBuyBorrowAsset
}

// GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAssetOk() (*string, bool) {
	if o == nil || IsNil(o.MarginBuyBorrowAsset) {
		return nil, false
	}
	return o.MarginBuyBorrowAsset, true
}

// HasMarginBuyBorrowAsset returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasMarginBuyBorrowAsset() bool {
	if o != nil && !IsNil(o.MarginBuyBorrowAsset) {
		return true
	}

	return false
}

// SetMarginBuyBorrowAsset gets a reference to the given string and assigns it to the MarginBuyBorrowAsset field.
func (o *MarginCreateMarginOrderV1Resp) SetMarginBuyBorrowAsset(v string) {
	o.MarginBuyBorrowAsset = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *MarginCreateMarginOrderV1Resp) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *MarginCreateMarginOrderV1Resp) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *MarginCreateMarginOrderV1Resp) SetPrice(v string) {
	o.Price = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *MarginCreateMarginOrderV1Resp) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *MarginCreateMarginOrderV1Resp) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MarginCreateMarginOrderV1Resp) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MarginCreateMarginOrderV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *MarginCreateMarginOrderV1Resp) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetTransactTime() int64 {
	if o == nil || IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetTransactTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasTransactTime() bool {
	if o != nil && !IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *MarginCreateMarginOrderV1Resp) SetTransactTime(v int64) {
	o.TransactTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginCreateMarginOrderV1Resp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCreateMarginOrderV1Resp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginCreateMarginOrderV1Resp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginCreateMarginOrderV1Resp) SetType(v string) {
	o.Type = &v
}

func (o MarginCreateMarginOrderV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginCreateMarginOrderV1Resp) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Fills) {
		toSerialize["fills"] = o.Fills
	}
	if !IsNil(o.IsIsolated) {
		toSerialize["isIsolated"] = o.IsIsolated
	}
	if !IsNil(o.MarginBuyBorrowAmount) {
		toSerialize["marginBuyBorrowAmount"] = o.MarginBuyBorrowAmount
	}
	if !IsNil(o.MarginBuyBorrowAsset) {
		toSerialize["marginBuyBorrowAsset"] = o.MarginBuyBorrowAsset
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
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
	if !IsNil(o.TransactTime) {
		toSerialize["transactTime"] = o.TransactTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMarginCreateMarginOrderV1Resp struct {
	value *MarginCreateMarginOrderV1Resp
	isSet bool
}

func (v NullableMarginCreateMarginOrderV1Resp) Get() *MarginCreateMarginOrderV1Resp {
	return v.value
}

func (v *NullableMarginCreateMarginOrderV1Resp) Set(val *MarginCreateMarginOrderV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCreateMarginOrderV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCreateMarginOrderV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCreateMarginOrderV1Resp(val *MarginCreateMarginOrderV1Resp) *NullableMarginCreateMarginOrderV1Resp {
	return &NullableMarginCreateMarginOrderV1Resp{value: val, isSet: true}
}

func (v NullableMarginCreateMarginOrderV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCreateMarginOrderV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


