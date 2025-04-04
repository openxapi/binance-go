/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem{}

// UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem struct for UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem
type UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem struct {
	LongAccount *string `json:"longAccount,omitempty"`
	LongShortRatio *string `json:"longShortRatio,omitempty"`
	ShortAccount *string `json:"shortAccount,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem instantiates a new UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem() *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem {
	this := UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem{}
	return &this
}

// NewUmfuturesGetFuturesDataTopLongShortAccountRatioRespItemWithDefaults instantiates a new UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesGetFuturesDataTopLongShortAccountRatioRespItemWithDefaults() *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem {
	this := UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem{}
	return &this
}

// GetLongAccount returns the LongAccount field value if set, zero value otherwise.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetLongAccount() string {
	if o == nil || IsNil(o.LongAccount) {
		var ret string
		return ret
	}
	return *o.LongAccount
}

// GetLongAccountOk returns a tuple with the LongAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetLongAccountOk() (*string, bool) {
	if o == nil || IsNil(o.LongAccount) {
		return nil, false
	}
	return o.LongAccount, true
}

// HasLongAccount returns a boolean if a field has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) HasLongAccount() bool {
	if o != nil && !IsNil(o.LongAccount) {
		return true
	}

	return false
}

// SetLongAccount gets a reference to the given string and assigns it to the LongAccount field.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) SetLongAccount(v string) {
	o.LongAccount = &v
}

// GetLongShortRatio returns the LongShortRatio field value if set, zero value otherwise.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetLongShortRatio() string {
	if o == nil || IsNil(o.LongShortRatio) {
		var ret string
		return ret
	}
	return *o.LongShortRatio
}

// GetLongShortRatioOk returns a tuple with the LongShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetLongShortRatioOk() (*string, bool) {
	if o == nil || IsNil(o.LongShortRatio) {
		return nil, false
	}
	return o.LongShortRatio, true
}

// HasLongShortRatio returns a boolean if a field has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) HasLongShortRatio() bool {
	if o != nil && !IsNil(o.LongShortRatio) {
		return true
	}

	return false
}

// SetLongShortRatio gets a reference to the given string and assigns it to the LongShortRatio field.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) SetLongShortRatio(v string) {
	o.LongShortRatio = &v
}

// GetShortAccount returns the ShortAccount field value if set, zero value otherwise.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetShortAccount() string {
	if o == nil || IsNil(o.ShortAccount) {
		var ret string
		return ret
	}
	return *o.ShortAccount
}

// GetShortAccountOk returns a tuple with the ShortAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetShortAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ShortAccount) {
		return nil, false
	}
	return o.ShortAccount, true
}

// HasShortAccount returns a boolean if a field has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) HasShortAccount() bool {
	if o != nil && !IsNil(o.ShortAccount) {
		return true
	}

	return false
}

// SetShortAccount gets a reference to the given string and assigns it to the ShortAccount field.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) SetShortAccount(v string) {
	o.ShortAccount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongAccount) {
		toSerialize["longAccount"] = o.LongAccount
	}
	if !IsNil(o.LongShortRatio) {
		toSerialize["longShortRatio"] = o.LongShortRatio
	}
	if !IsNil(o.ShortAccount) {
		toSerialize["shortAccount"] = o.ShortAccount
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem struct {
	value *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem
	isSet bool
}

func (v NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) Get() *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem {
	return v.value
}

func (v *NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) Set(val *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem(val *UmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) *NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem {
	return &NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem{value: val, isSet: true}
}

func (v NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetFuturesDataTopLongShortAccountRatioRespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


