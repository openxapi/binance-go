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

// checks if the MarginGetMarginAvailableInventoryV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginAvailableInventoryV1Resp{}

// MarginGetMarginAvailableInventoryV1Resp struct for MarginGetMarginAvailableInventoryV1Resp
type MarginGetMarginAvailableInventoryV1Resp struct {
	Assets *map[string]string `json:"assets,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewMarginGetMarginAvailableInventoryV1Resp instantiates a new MarginGetMarginAvailableInventoryV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginAvailableInventoryV1Resp() *MarginGetMarginAvailableInventoryV1Resp {
	this := MarginGetMarginAvailableInventoryV1Resp{}
	return &this
}

// NewMarginGetMarginAvailableInventoryV1RespWithDefaults instantiates a new MarginGetMarginAvailableInventoryV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginAvailableInventoryV1RespWithDefaults() *MarginGetMarginAvailableInventoryV1Resp {
	this := MarginGetMarginAvailableInventoryV1Resp{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *MarginGetMarginAvailableInventoryV1Resp) GetAssets() map[string]string {
	if o == nil || IsNil(o.Assets) {
		var ret map[string]string
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAvailableInventoryV1Resp) GetAssetsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *MarginGetMarginAvailableInventoryV1Resp) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given map[string]string and assigns it to the Assets field.
func (o *MarginGetMarginAvailableInventoryV1Resp) SetAssets(v map[string]string) {
	o.Assets = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *MarginGetMarginAvailableInventoryV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginAvailableInventoryV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *MarginGetMarginAvailableInventoryV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *MarginGetMarginAvailableInventoryV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o MarginGetMarginAvailableInventoryV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginAvailableInventoryV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableMarginGetMarginAvailableInventoryV1Resp struct {
	value *MarginGetMarginAvailableInventoryV1Resp
	isSet bool
}

func (v NullableMarginGetMarginAvailableInventoryV1Resp) Get() *MarginGetMarginAvailableInventoryV1Resp {
	return v.value
}

func (v *NullableMarginGetMarginAvailableInventoryV1Resp) Set(val *MarginGetMarginAvailableInventoryV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginAvailableInventoryV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginAvailableInventoryV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginAvailableInventoryV1Resp(val *MarginGetMarginAvailableInventoryV1Resp) *NullableMarginGetMarginAvailableInventoryV1Resp {
	return &NullableMarginGetMarginAvailableInventoryV1Resp{value: val, isSet: true}
}

func (v NullableMarginGetMarginAvailableInventoryV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginAvailableInventoryV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


