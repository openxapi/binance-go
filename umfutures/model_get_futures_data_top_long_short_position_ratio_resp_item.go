/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the GetFuturesDataTopLongShortPositionRatioRespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFuturesDataTopLongShortPositionRatioRespItem{}

// GetFuturesDataTopLongShortPositionRatioRespItem struct for GetFuturesDataTopLongShortPositionRatioRespItem
type GetFuturesDataTopLongShortPositionRatioRespItem struct {
	LongAccount *string `json:"longAccount,omitempty"`
	LongShortRatio *string `json:"longShortRatio,omitempty"`
	ShortAccount *string `json:"shortAccount,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewGetFuturesDataTopLongShortPositionRatioRespItem instantiates a new GetFuturesDataTopLongShortPositionRatioRespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesDataTopLongShortPositionRatioRespItem() *GetFuturesDataTopLongShortPositionRatioRespItem {
	this := GetFuturesDataTopLongShortPositionRatioRespItem{}
	return &this
}

// NewGetFuturesDataTopLongShortPositionRatioRespItemWithDefaults instantiates a new GetFuturesDataTopLongShortPositionRatioRespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesDataTopLongShortPositionRatioRespItemWithDefaults() *GetFuturesDataTopLongShortPositionRatioRespItem {
	this := GetFuturesDataTopLongShortPositionRatioRespItem{}
	return &this
}

// GetLongAccount returns the LongAccount field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetLongAccount() string {
	if o == nil || IsNil(o.LongAccount) {
		var ret string
		return ret
	}
	return *o.LongAccount
}

// GetLongAccountOk returns a tuple with the LongAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetLongAccountOk() (*string, bool) {
	if o == nil || IsNil(o.LongAccount) {
		return nil, false
	}
	return o.LongAccount, true
}

// HasLongAccount returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) HasLongAccount() bool {
	if o != nil && !IsNil(o.LongAccount) {
		return true
	}

	return false
}

// SetLongAccount gets a reference to the given string and assigns it to the LongAccount field.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) SetLongAccount(v string) {
	o.LongAccount = &v
}

// GetLongShortRatio returns the LongShortRatio field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetLongShortRatio() string {
	if o == nil || IsNil(o.LongShortRatio) {
		var ret string
		return ret
	}
	return *o.LongShortRatio
}

// GetLongShortRatioOk returns a tuple with the LongShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetLongShortRatioOk() (*string, bool) {
	if o == nil || IsNil(o.LongShortRatio) {
		return nil, false
	}
	return o.LongShortRatio, true
}

// HasLongShortRatio returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) HasLongShortRatio() bool {
	if o != nil && !IsNil(o.LongShortRatio) {
		return true
	}

	return false
}

// SetLongShortRatio gets a reference to the given string and assigns it to the LongShortRatio field.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) SetLongShortRatio(v string) {
	o.LongShortRatio = &v
}

// GetShortAccount returns the ShortAccount field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetShortAccount() string {
	if o == nil || IsNil(o.ShortAccount) {
		var ret string
		return ret
	}
	return *o.ShortAccount
}

// GetShortAccountOk returns a tuple with the ShortAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetShortAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ShortAccount) {
		return nil, false
	}
	return o.ShortAccount, true
}

// HasShortAccount returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) HasShortAccount() bool {
	if o != nil && !IsNil(o.ShortAccount) {
		return true
	}

	return false
}

// SetShortAccount gets a reference to the given string and assigns it to the ShortAccount field.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) SetShortAccount(v string) {
	o.ShortAccount = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *GetFuturesDataTopLongShortPositionRatioRespItem) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o GetFuturesDataTopLongShortPositionRatioRespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesDataTopLongShortPositionRatioRespItem) ToMap() (map[string]interface{}, error) {
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

type NullableGetFuturesDataTopLongShortPositionRatioRespItem struct {
	value *GetFuturesDataTopLongShortPositionRatioRespItem
	isSet bool
}

func (v NullableGetFuturesDataTopLongShortPositionRatioRespItem) Get() *GetFuturesDataTopLongShortPositionRatioRespItem {
	return v.value
}

func (v *NullableGetFuturesDataTopLongShortPositionRatioRespItem) Set(val *GetFuturesDataTopLongShortPositionRatioRespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesDataTopLongShortPositionRatioRespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesDataTopLongShortPositionRatioRespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesDataTopLongShortPositionRatioRespItem(val *GetFuturesDataTopLongShortPositionRatioRespItem) *NullableGetFuturesDataTopLongShortPositionRatioRespItem {
	return &NullableGetFuturesDataTopLongShortPositionRatioRespItem{value: val, isSet: true}
}

func (v NullableGetFuturesDataTopLongShortPositionRatioRespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesDataTopLongShortPositionRatioRespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


