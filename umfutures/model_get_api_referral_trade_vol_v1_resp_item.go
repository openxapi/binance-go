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

// checks if the GetApiReferralTradeVolV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiReferralTradeVolV1RespItem{}

// GetApiReferralTradeVolV1RespItem struct for GetApiReferralTradeVolV1RespItem
type GetApiReferralTradeVolV1RespItem struct {
	Time *int64 `json:"time,omitempty"`
	TradeVol *string `json:"tradeVol,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewGetApiReferralTradeVolV1RespItem instantiates a new GetApiReferralTradeVolV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiReferralTradeVolV1RespItem() *GetApiReferralTradeVolV1RespItem {
	this := GetApiReferralTradeVolV1RespItem{}
	return &this
}

// NewGetApiReferralTradeVolV1RespItemWithDefaults instantiates a new GetApiReferralTradeVolV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiReferralTradeVolV1RespItemWithDefaults() *GetApiReferralTradeVolV1RespItem {
	this := GetApiReferralTradeVolV1RespItem{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetApiReferralTradeVolV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiReferralTradeVolV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetApiReferralTradeVolV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetApiReferralTradeVolV1RespItem) SetTime(v int64) {
	o.Time = &v
}

// GetTradeVol returns the TradeVol field value if set, zero value otherwise.
func (o *GetApiReferralTradeVolV1RespItem) GetTradeVol() string {
	if o == nil || IsNil(o.TradeVol) {
		var ret string
		return ret
	}
	return *o.TradeVol
}

// GetTradeVolOk returns a tuple with the TradeVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiReferralTradeVolV1RespItem) GetTradeVolOk() (*string, bool) {
	if o == nil || IsNil(o.TradeVol) {
		return nil, false
	}
	return o.TradeVol, true
}

// HasTradeVol returns a boolean if a field has been set.
func (o *GetApiReferralTradeVolV1RespItem) HasTradeVol() bool {
	if o != nil && !IsNil(o.TradeVol) {
		return true
	}

	return false
}

// SetTradeVol gets a reference to the given string and assigns it to the TradeVol field.
func (o *GetApiReferralTradeVolV1RespItem) SetTradeVol(v string) {
	o.TradeVol = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GetApiReferralTradeVolV1RespItem) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiReferralTradeVolV1RespItem) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GetApiReferralTradeVolV1RespItem) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GetApiReferralTradeVolV1RespItem) SetUnit(v string) {
	o.Unit = &v
}

func (o GetApiReferralTradeVolV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiReferralTradeVolV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TradeVol) {
		toSerialize["tradeVol"] = o.TradeVol
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableGetApiReferralTradeVolV1RespItem struct {
	value *GetApiReferralTradeVolV1RespItem
	isSet bool
}

func (v NullableGetApiReferralTradeVolV1RespItem) Get() *GetApiReferralTradeVolV1RespItem {
	return v.value
}

func (v *NullableGetApiReferralTradeVolV1RespItem) Set(val *GetApiReferralTradeVolV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiReferralTradeVolV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiReferralTradeVolV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiReferralTradeVolV1RespItem(val *GetApiReferralTradeVolV1RespItem) *NullableGetApiReferralTradeVolV1RespItem {
	return &NullableGetApiReferralTradeVolV1RespItem{value: val, isSet: true}
}

func (v NullableGetApiReferralTradeVolV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiReferralTradeVolV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


