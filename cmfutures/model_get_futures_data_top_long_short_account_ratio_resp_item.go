/*
Binance COIN-M Futures API

OpenAPI specification for Binance exchange - Cmfutures API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the GetFuturesDataTopLongShortAccountRatioRespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFuturesDataTopLongShortAccountRatioRespItem{}

// GetFuturesDataTopLongShortAccountRatioRespItem struct for GetFuturesDataTopLongShortAccountRatioRespItem
type GetFuturesDataTopLongShortAccountRatioRespItem struct {
	LongAccount *string `json:"longAccount,omitempty"`
	LongShortRatio *string `json:"longShortRatio,omitempty"`
	Pair *string `json:"pair,omitempty"`
	ShortAccount *string `json:"shortAccount,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// NewGetFuturesDataTopLongShortAccountRatioRespItem instantiates a new GetFuturesDataTopLongShortAccountRatioRespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesDataTopLongShortAccountRatioRespItem() *GetFuturesDataTopLongShortAccountRatioRespItem {
	this := GetFuturesDataTopLongShortAccountRatioRespItem{}
	return &this
}

// NewGetFuturesDataTopLongShortAccountRatioRespItemWithDefaults instantiates a new GetFuturesDataTopLongShortAccountRatioRespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesDataTopLongShortAccountRatioRespItemWithDefaults() *GetFuturesDataTopLongShortAccountRatioRespItem {
	this := GetFuturesDataTopLongShortAccountRatioRespItem{}
	return &this
}

// GetLongAccount returns the LongAccount field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetLongAccount() string {
	if o == nil || IsNil(o.LongAccount) {
		var ret string
		return ret
	}
	return *o.LongAccount
}

// GetLongAccountOk returns a tuple with the LongAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetLongAccountOk() (*string, bool) {
	if o == nil || IsNil(o.LongAccount) {
		return nil, false
	}
	return o.LongAccount, true
}

// HasLongAccount returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) HasLongAccount() bool {
	if o != nil && !IsNil(o.LongAccount) {
		return true
	}

	return false
}

// SetLongAccount gets a reference to the given string and assigns it to the LongAccount field.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) SetLongAccount(v string) {
	o.LongAccount = &v
}

// GetLongShortRatio returns the LongShortRatio field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetLongShortRatio() string {
	if o == nil || IsNil(o.LongShortRatio) {
		var ret string
		return ret
	}
	return *o.LongShortRatio
}

// GetLongShortRatioOk returns a tuple with the LongShortRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetLongShortRatioOk() (*string, bool) {
	if o == nil || IsNil(o.LongShortRatio) {
		return nil, false
	}
	return o.LongShortRatio, true
}

// HasLongShortRatio returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) HasLongShortRatio() bool {
	if o != nil && !IsNil(o.LongShortRatio) {
		return true
	}

	return false
}

// SetLongShortRatio gets a reference to the given string and assigns it to the LongShortRatio field.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) SetLongShortRatio(v string) {
	o.LongShortRatio = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetPair() string {
	if o == nil || IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetPairOk() (*string, bool) {
	if o == nil || IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) HasPair() bool {
	if o != nil && !IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) SetPair(v string) {
	o.Pair = &v
}

// GetShortAccount returns the ShortAccount field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetShortAccount() string {
	if o == nil || IsNil(o.ShortAccount) {
		var ret string
		return ret
	}
	return *o.ShortAccount
}

// GetShortAccountOk returns a tuple with the ShortAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetShortAccountOk() (*string, bool) {
	if o == nil || IsNil(o.ShortAccount) {
		return nil, false
	}
	return o.ShortAccount, true
}

// HasShortAccount returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) HasShortAccount() bool {
	if o != nil && !IsNil(o.ShortAccount) {
		return true
	}

	return false
}

// SetShortAccount gets a reference to the given string and assigns it to the ShortAccount field.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) SetShortAccount(v string) {
	o.ShortAccount = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetFuturesDataTopLongShortAccountRatioRespItem) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o GetFuturesDataTopLongShortAccountRatioRespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesDataTopLongShortAccountRatioRespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongAccount) {
		toSerialize["longAccount"] = o.LongAccount
	}
	if !IsNil(o.LongShortRatio) {
		toSerialize["longShortRatio"] = o.LongShortRatio
	}
	if !IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !IsNil(o.ShortAccount) {
		toSerialize["shortAccount"] = o.ShortAccount
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableGetFuturesDataTopLongShortAccountRatioRespItem struct {
	value *GetFuturesDataTopLongShortAccountRatioRespItem
	isSet bool
}

func (v NullableGetFuturesDataTopLongShortAccountRatioRespItem) Get() *GetFuturesDataTopLongShortAccountRatioRespItem {
	return v.value
}

func (v *NullableGetFuturesDataTopLongShortAccountRatioRespItem) Set(val *GetFuturesDataTopLongShortAccountRatioRespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesDataTopLongShortAccountRatioRespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesDataTopLongShortAccountRatioRespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesDataTopLongShortAccountRatioRespItem(val *GetFuturesDataTopLongShortAccountRatioRespItem) *NullableGetFuturesDataTopLongShortAccountRatioRespItem {
	return &NullableGetFuturesDataTopLongShortAccountRatioRespItem{value: val, isSet: true}
}

func (v NullableGetFuturesDataTopLongShortAccountRatioRespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesDataTopLongShortAccountRatioRespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


