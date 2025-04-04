/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the OptionsGetHistoricalTradesV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsGetHistoricalTradesV1RespItem{}

// OptionsGetHistoricalTradesV1RespItem struct for OptionsGetHistoricalTradesV1RespItem
type OptionsGetHistoricalTradesV1RespItem struct {
	Id *string `json:"id,omitempty"`
	Price *string `json:"price,omitempty"`
	Qty *string `json:"qty,omitempty"`
	QuoteQty *string `json:"quoteQty,omitempty"`
	Side *int32 `json:"side,omitempty"`
	Time *int64 `json:"time,omitempty"`
	TradeId *string `json:"tradeId,omitempty"`
}

// NewOptionsGetHistoricalTradesV1RespItem instantiates a new OptionsGetHistoricalTradesV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsGetHistoricalTradesV1RespItem() *OptionsGetHistoricalTradesV1RespItem {
	this := OptionsGetHistoricalTradesV1RespItem{}
	return &this
}

// NewOptionsGetHistoricalTradesV1RespItemWithDefaults instantiates a new OptionsGetHistoricalTradesV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsGetHistoricalTradesV1RespItemWithDefaults() *OptionsGetHistoricalTradesV1RespItem {
	this := OptionsGetHistoricalTradesV1RespItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionsGetHistoricalTradesV1RespItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OptionsGetHistoricalTradesV1RespItem) SetId(v string) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OptionsGetHistoricalTradesV1RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OptionsGetHistoricalTradesV1RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *OptionsGetHistoricalTradesV1RespItem) GetQty() string {
	if o == nil || IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) GetQtyOk() (*string, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *OptionsGetHistoricalTradesV1RespItem) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *OptionsGetHistoricalTradesV1RespItem) GetQuoteQty() string {
	if o == nil || IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) GetQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) HasQuoteQty() bool {
	if o != nil && !IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *OptionsGetHistoricalTradesV1RespItem) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OptionsGetHistoricalTradesV1RespItem) GetSide() int32 {
	if o == nil || IsNil(o.Side) {
		var ret int32
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) GetSideOk() (*int32, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given int32 and assigns it to the Side field.
func (o *OptionsGetHistoricalTradesV1RespItem) SetSide(v int32) {
	o.Side = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OptionsGetHistoricalTradesV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OptionsGetHistoricalTradesV1RespItem) SetTime(v int64) {
	o.Time = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *OptionsGetHistoricalTradesV1RespItem) GetTradeId() string {
	if o == nil || IsNil(o.TradeId) {
		var ret string
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) GetTradeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *OptionsGetHistoricalTradesV1RespItem) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given string and assigns it to the TradeId field.
func (o *OptionsGetHistoricalTradesV1RespItem) SetTradeId(v string) {
	o.TradeId = &v
}

func (o OptionsGetHistoricalTradesV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsGetHistoricalTradesV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !IsNil(o.QuoteQty) {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	return toSerialize, nil
}

type NullableOptionsGetHistoricalTradesV1RespItem struct {
	value *OptionsGetHistoricalTradesV1RespItem
	isSet bool
}

func (v NullableOptionsGetHistoricalTradesV1RespItem) Get() *OptionsGetHistoricalTradesV1RespItem {
	return v.value
}

func (v *NullableOptionsGetHistoricalTradesV1RespItem) Set(val *OptionsGetHistoricalTradesV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsGetHistoricalTradesV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsGetHistoricalTradesV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsGetHistoricalTradesV1RespItem(val *OptionsGetHistoricalTradesV1RespItem) *NullableOptionsGetHistoricalTradesV1RespItem {
	return &NullableOptionsGetHistoricalTradesV1RespItem{value: val, isSet: true}
}

func (v NullableOptionsGetHistoricalTradesV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsGetHistoricalTradesV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


