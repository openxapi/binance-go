/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetTickerBookTickerV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetTickerBookTickerV1RespItem{}

// UfuturesGetTickerBookTickerV1RespItem struct for UfuturesGetTickerBookTickerV1RespItem
type UfuturesGetTickerBookTickerV1RespItem struct {
	AskPrice *string `json:"askPrice,omitempty"`
	AskQty *string `json:"askQty,omitempty"`
	BidPrice *string `json:"bidPrice,omitempty"`
	BidQty *string `json:"bidQty,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewUfuturesGetTickerBookTickerV1RespItem instantiates a new UfuturesGetTickerBookTickerV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetTickerBookTickerV1RespItem() *UfuturesGetTickerBookTickerV1RespItem {
	this := UfuturesGetTickerBookTickerV1RespItem{}
	return &this
}

// NewUfuturesGetTickerBookTickerV1RespItemWithDefaults instantiates a new UfuturesGetTickerBookTickerV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetTickerBookTickerV1RespItemWithDefaults() *UfuturesGetTickerBookTickerV1RespItem {
	this := UfuturesGetTickerBookTickerV1RespItem{}
	return &this
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetAskPrice() string {
	if o == nil || IsNil(o.AskPrice) {
		var ret string
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetAskPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) HasAskPrice() bool {
	if o != nil && !IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given string and assigns it to the AskPrice field.
func (o *UfuturesGetTickerBookTickerV1RespItem) SetAskPrice(v string) {
	o.AskPrice = &v
}

// GetAskQty returns the AskQty field value if set, zero value otherwise.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetAskQty() string {
	if o == nil || IsNil(o.AskQty) {
		var ret string
		return ret
	}
	return *o.AskQty
}

// GetAskQtyOk returns a tuple with the AskQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetAskQtyOk() (*string, bool) {
	if o == nil || IsNil(o.AskQty) {
		return nil, false
	}
	return o.AskQty, true
}

// HasAskQty returns a boolean if a field has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) HasAskQty() bool {
	if o != nil && !IsNil(o.AskQty) {
		return true
	}

	return false
}

// SetAskQty gets a reference to the given string and assigns it to the AskQty field.
func (o *UfuturesGetTickerBookTickerV1RespItem) SetAskQty(v string) {
	o.AskQty = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetBidPrice() string {
	if o == nil || IsNil(o.BidPrice) {
		var ret string
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetBidPriceOk() (*string, bool) {
	if o == nil || IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) HasBidPrice() bool {
	if o != nil && !IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given string and assigns it to the BidPrice field.
func (o *UfuturesGetTickerBookTickerV1RespItem) SetBidPrice(v string) {
	o.BidPrice = &v
}

// GetBidQty returns the BidQty field value if set, zero value otherwise.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetBidQty() string {
	if o == nil || IsNil(o.BidQty) {
		var ret string
		return ret
	}
	return *o.BidQty
}

// GetBidQtyOk returns a tuple with the BidQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetBidQtyOk() (*string, bool) {
	if o == nil || IsNil(o.BidQty) {
		return nil, false
	}
	return o.BidQty, true
}

// HasBidQty returns a boolean if a field has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) HasBidQty() bool {
	if o != nil && !IsNil(o.BidQty) {
		return true
	}

	return false
}

// SetBidQty gets a reference to the given string and assigns it to the BidQty field.
func (o *UfuturesGetTickerBookTickerV1RespItem) SetBidQty(v string) {
	o.BidQty = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UfuturesGetTickerBookTickerV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UfuturesGetTickerBookTickerV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *UfuturesGetTickerBookTickerV1RespItem) SetTime(v int64) {
	o.Time = &v
}

func (o UfuturesGetTickerBookTickerV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetTickerBookTickerV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AskPrice) {
		toSerialize["askPrice"] = o.AskPrice
	}
	if !IsNil(o.AskQty) {
		toSerialize["askQty"] = o.AskQty
	}
	if !IsNil(o.BidPrice) {
		toSerialize["bidPrice"] = o.BidPrice
	}
	if !IsNil(o.BidQty) {
		toSerialize["bidQty"] = o.BidQty
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableUfuturesGetTickerBookTickerV1RespItem struct {
	value *UfuturesGetTickerBookTickerV1RespItem
	isSet bool
}

func (v NullableUfuturesGetTickerBookTickerV1RespItem) Get() *UfuturesGetTickerBookTickerV1RespItem {
	return v.value
}

func (v *NullableUfuturesGetTickerBookTickerV1RespItem) Set(val *UfuturesGetTickerBookTickerV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetTickerBookTickerV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetTickerBookTickerV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetTickerBookTickerV1RespItem(val *UfuturesGetTickerBookTickerV1RespItem) *NullableUfuturesGetTickerBookTickerV1RespItem {
	return &NullableUfuturesGetTickerBookTickerV1RespItem{value: val, isSet: true}
}

func (v NullableUfuturesGetTickerBookTickerV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetTickerBookTickerV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


