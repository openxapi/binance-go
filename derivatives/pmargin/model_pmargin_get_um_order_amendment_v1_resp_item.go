/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginGetUmOrderAmendmentV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetUmOrderAmendmentV1RespItem{}

// PmarginGetUmOrderAmendmentV1RespItem struct for PmarginGetUmOrderAmendmentV1RespItem
type PmarginGetUmOrderAmendmentV1RespItem struct {
	Amendment *PmarginGetCmOrderAmendmentV1RespItemAmendment `json:"amendment,omitempty"`
	AmendmentId *int64 `json:"amendmentId,omitempty"`
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	Pair *string `json:"pair,omitempty"`
	PriceMatch *string `json:"priceMatch,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewPmarginGetUmOrderAmendmentV1RespItem instantiates a new PmarginGetUmOrderAmendmentV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetUmOrderAmendmentV1RespItem() *PmarginGetUmOrderAmendmentV1RespItem {
	this := PmarginGetUmOrderAmendmentV1RespItem{}
	return &this
}

// NewPmarginGetUmOrderAmendmentV1RespItemWithDefaults instantiates a new PmarginGetUmOrderAmendmentV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetUmOrderAmendmentV1RespItemWithDefaults() *PmarginGetUmOrderAmendmentV1RespItem {
	this := PmarginGetUmOrderAmendmentV1RespItem{}
	return &this
}

// GetAmendment returns the Amendment field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetAmendment() PmarginGetCmOrderAmendmentV1RespItemAmendment {
	if o == nil || IsNil(o.Amendment) {
		var ret PmarginGetCmOrderAmendmentV1RespItemAmendment
		return ret
	}
	return *o.Amendment
}

// GetAmendmentOk returns a tuple with the Amendment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetAmendmentOk() (*PmarginGetCmOrderAmendmentV1RespItemAmendment, bool) {
	if o == nil || IsNil(o.Amendment) {
		return nil, false
	}
	return o.Amendment, true
}

// HasAmendment returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasAmendment() bool {
	if o != nil && !IsNil(o.Amendment) {
		return true
	}

	return false
}

// SetAmendment gets a reference to the given PmarginGetCmOrderAmendmentV1RespItemAmendment and assigns it to the Amendment field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetAmendment(v PmarginGetCmOrderAmendmentV1RespItemAmendment) {
	o.Amendment = &v
}

// GetAmendmentId returns the AmendmentId field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetAmendmentId() int64 {
	if o == nil || IsNil(o.AmendmentId) {
		var ret int64
		return ret
	}
	return *o.AmendmentId
}

// GetAmendmentIdOk returns a tuple with the AmendmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetAmendmentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AmendmentId) {
		return nil, false
	}
	return o.AmendmentId, true
}

// HasAmendmentId returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasAmendmentId() bool {
	if o != nil && !IsNil(o.AmendmentId) {
		return true
	}

	return false
}

// SetAmendmentId gets a reference to the given int64 and assigns it to the AmendmentId field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetAmendmentId(v int64) {
	o.AmendmentId = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetPair() string {
	if o == nil || IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetPairOk() (*string, bool) {
	if o == nil || IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasPair() bool {
	if o != nil && !IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetPair(v string) {
	o.Pair = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetPriceMatch() string {
	if o == nil || IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetPriceMatchOk() (*string, bool) {
	if o == nil || IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasPriceMatch() bool {
	if o != nil && !IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *PmarginGetUmOrderAmendmentV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *PmarginGetUmOrderAmendmentV1RespItem) SetTime(v int64) {
	o.Time = &v
}

func (o PmarginGetUmOrderAmendmentV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetUmOrderAmendmentV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amendment) {
		toSerialize["amendment"] = o.Amendment
	}
	if !IsNil(o.AmendmentId) {
		toSerialize["amendmentId"] = o.AmendmentId
	}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullablePmarginGetUmOrderAmendmentV1RespItem struct {
	value *PmarginGetUmOrderAmendmentV1RespItem
	isSet bool
}

func (v NullablePmarginGetUmOrderAmendmentV1RespItem) Get() *PmarginGetUmOrderAmendmentV1RespItem {
	return v.value
}

func (v *NullablePmarginGetUmOrderAmendmentV1RespItem) Set(val *PmarginGetUmOrderAmendmentV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetUmOrderAmendmentV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetUmOrderAmendmentV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetUmOrderAmendmentV1RespItem(val *PmarginGetUmOrderAmendmentV1RespItem) *NullablePmarginGetUmOrderAmendmentV1RespItem {
	return &NullablePmarginGetUmOrderAmendmentV1RespItem{value: val, isSet: true}
}

func (v NullablePmarginGetUmOrderAmendmentV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetUmOrderAmendmentV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


