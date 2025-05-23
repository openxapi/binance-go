/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the CreateCountdownCancelAllV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCountdownCancelAllV1Resp{}

// CreateCountdownCancelAllV1Resp struct for CreateCountdownCancelAllV1Resp
type CreateCountdownCancelAllV1Resp struct {
	CountdownTime *int64 `json:"countdownTime,omitempty"`
	Underlying *string `json:"underlying,omitempty"`
}

// NewCreateCountdownCancelAllV1Resp instantiates a new CreateCountdownCancelAllV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCountdownCancelAllV1Resp() *CreateCountdownCancelAllV1Resp {
	this := CreateCountdownCancelAllV1Resp{}
	return &this
}

// NewCreateCountdownCancelAllV1RespWithDefaults instantiates a new CreateCountdownCancelAllV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCountdownCancelAllV1RespWithDefaults() *CreateCountdownCancelAllV1Resp {
	this := CreateCountdownCancelAllV1Resp{}
	return &this
}

// GetCountdownTime returns the CountdownTime field value if set, zero value otherwise.
func (o *CreateCountdownCancelAllV1Resp) GetCountdownTime() int64 {
	if o == nil || IsNil(o.CountdownTime) {
		var ret int64
		return ret
	}
	return *o.CountdownTime
}

// GetCountdownTimeOk returns a tuple with the CountdownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCountdownCancelAllV1Resp) GetCountdownTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CountdownTime) {
		return nil, false
	}
	return o.CountdownTime, true
}

// HasCountdownTime returns a boolean if a field has been set.
func (o *CreateCountdownCancelAllV1Resp) HasCountdownTime() bool {
	if o != nil && !IsNil(o.CountdownTime) {
		return true
	}

	return false
}

// SetCountdownTime gets a reference to the given int64 and assigns it to the CountdownTime field.
func (o *CreateCountdownCancelAllV1Resp) SetCountdownTime(v int64) {
	o.CountdownTime = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *CreateCountdownCancelAllV1Resp) GetUnderlying() string {
	if o == nil || IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCountdownCancelAllV1Resp) GetUnderlyingOk() (*string, bool) {
	if o == nil || IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *CreateCountdownCancelAllV1Resp) HasUnderlying() bool {
	if o != nil && !IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *CreateCountdownCancelAllV1Resp) SetUnderlying(v string) {
	o.Underlying = &v
}

func (o CreateCountdownCancelAllV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCountdownCancelAllV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountdownTime) {
		toSerialize["countdownTime"] = o.CountdownTime
	}
	if !IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	return toSerialize, nil
}

type NullableCreateCountdownCancelAllV1Resp struct {
	value *CreateCountdownCancelAllV1Resp
	isSet bool
}

func (v NullableCreateCountdownCancelAllV1Resp) Get() *CreateCountdownCancelAllV1Resp {
	return v.value
}

func (v *NullableCreateCountdownCancelAllV1Resp) Set(val *CreateCountdownCancelAllV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCountdownCancelAllV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCountdownCancelAllV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCountdownCancelAllV1Resp(val *CreateCountdownCancelAllV1Resp) *NullableCreateCountdownCancelAllV1Resp {
	return &NullableCreateCountdownCancelAllV1Resp{value: val, isSet: true}
}

func (v NullableCreateCountdownCancelAllV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCountdownCancelAllV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


