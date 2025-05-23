/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the CreateUserDataStreamIsolatedV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserDataStreamIsolatedV1Resp{}

// CreateUserDataStreamIsolatedV1Resp struct for CreateUserDataStreamIsolatedV1Resp
type CreateUserDataStreamIsolatedV1Resp struct {
	ListenKey *string `json:"listenKey,omitempty"`
}

// NewCreateUserDataStreamIsolatedV1Resp instantiates a new CreateUserDataStreamIsolatedV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserDataStreamIsolatedV1Resp() *CreateUserDataStreamIsolatedV1Resp {
	this := CreateUserDataStreamIsolatedV1Resp{}
	return &this
}

// NewCreateUserDataStreamIsolatedV1RespWithDefaults instantiates a new CreateUserDataStreamIsolatedV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserDataStreamIsolatedV1RespWithDefaults() *CreateUserDataStreamIsolatedV1Resp {
	this := CreateUserDataStreamIsolatedV1Resp{}
	return &this
}

// GetListenKey returns the ListenKey field value if set, zero value otherwise.
func (o *CreateUserDataStreamIsolatedV1Resp) GetListenKey() string {
	if o == nil || IsNil(o.ListenKey) {
		var ret string
		return ret
	}
	return *o.ListenKey
}

// GetListenKeyOk returns a tuple with the ListenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserDataStreamIsolatedV1Resp) GetListenKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ListenKey) {
		return nil, false
	}
	return o.ListenKey, true
}

// HasListenKey returns a boolean if a field has been set.
func (o *CreateUserDataStreamIsolatedV1Resp) HasListenKey() bool {
	if o != nil && !IsNil(o.ListenKey) {
		return true
	}

	return false
}

// SetListenKey gets a reference to the given string and assigns it to the ListenKey field.
func (o *CreateUserDataStreamIsolatedV1Resp) SetListenKey(v string) {
	o.ListenKey = &v
}

func (o CreateUserDataStreamIsolatedV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserDataStreamIsolatedV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListenKey) {
		toSerialize["listenKey"] = o.ListenKey
	}
	return toSerialize, nil
}

type NullableCreateUserDataStreamIsolatedV1Resp struct {
	value *CreateUserDataStreamIsolatedV1Resp
	isSet bool
}

func (v NullableCreateUserDataStreamIsolatedV1Resp) Get() *CreateUserDataStreamIsolatedV1Resp {
	return v.value
}

func (v *NullableCreateUserDataStreamIsolatedV1Resp) Set(val *CreateUserDataStreamIsolatedV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserDataStreamIsolatedV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserDataStreamIsolatedV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserDataStreamIsolatedV1Resp(val *CreateUserDataStreamIsolatedV1Resp) *NullableCreateUserDataStreamIsolatedV1Resp {
	return &NullableCreateUserDataStreamIsolatedV1Resp{value: val, isSet: true}
}

func (v NullableCreateUserDataStreamIsolatedV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserDataStreamIsolatedV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


