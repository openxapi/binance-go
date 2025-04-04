/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotCreateUserDataStreamV3Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotCreateUserDataStreamV3Resp{}

// SpotCreateUserDataStreamV3Resp struct for SpotCreateUserDataStreamV3Resp
type SpotCreateUserDataStreamV3Resp struct {
	ListenKey *string `json:"listenKey,omitempty"`
}

// NewSpotCreateUserDataStreamV3Resp instantiates a new SpotCreateUserDataStreamV3Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotCreateUserDataStreamV3Resp() *SpotCreateUserDataStreamV3Resp {
	this := SpotCreateUserDataStreamV3Resp{}
	return &this
}

// NewSpotCreateUserDataStreamV3RespWithDefaults instantiates a new SpotCreateUserDataStreamV3Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotCreateUserDataStreamV3RespWithDefaults() *SpotCreateUserDataStreamV3Resp {
	this := SpotCreateUserDataStreamV3Resp{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *SpotCreateUserDataStreamV3Resp) GetListenKey() string {
	if o == nil || IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateUserDataStreamV3Resp) GetListenKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *SpotCreateUserDataStreamV3Resp) HasListenKey() bool {
	if o != nil && !IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *SpotCreateUserDataStreamV3Resp) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o SpotCreateUserDataStreamV3Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotCreateUserDataStreamV3Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}
	return toSerialize, nil
}

type NullableSpotCreateUserDataStreamV3Resp struct {
	value *SpotCreateUserDataStreamV3Resp
	isSet bool
}

func (v NullableSpotCreateUserDataStreamV3Resp) Get() *SpotCreateUserDataStreamV3Resp {
	return v.value
}

func (v *NullableSpotCreateUserDataStreamV3Resp) Set(val *SpotCreateUserDataStreamV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotCreateUserDataStreamV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotCreateUserDataStreamV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotCreateUserDataStreamV3Resp(val *SpotCreateUserDataStreamV3Resp) *NullableSpotCreateUserDataStreamV3Resp {
	return &NullableSpotCreateUserDataStreamV3Resp{value: val, isSet: true}
}

func (v NullableSpotCreateUserDataStreamV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotCreateUserDataStreamV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


