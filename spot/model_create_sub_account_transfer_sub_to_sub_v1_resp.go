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

// checks if the CreateSubAccountTransferSubToSubV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubAccountTransferSubToSubV1Resp{}

// CreateSubAccountTransferSubToSubV1Resp struct for CreateSubAccountTransferSubToSubV1Resp
type CreateSubAccountTransferSubToSubV1Resp struct {
	TxnId *string `json:"txnId,omitempty"`
}

// NewCreateSubAccountTransferSubToSubV1Resp instantiates a new CreateSubAccountTransferSubToSubV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubAccountTransferSubToSubV1Resp() *CreateSubAccountTransferSubToSubV1Resp {
	this := CreateSubAccountTransferSubToSubV1Resp{}
	return &this
}

// NewCreateSubAccountTransferSubToSubV1RespWithDefaults instantiates a new CreateSubAccountTransferSubToSubV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubAccountTransferSubToSubV1RespWithDefaults() *CreateSubAccountTransferSubToSubV1Resp {
	this := CreateSubAccountTransferSubToSubV1Resp{}
	return &this
}

// GetTxnId returns the TxnId field value if set, zero value otherwise.
func (o *CreateSubAccountTransferSubToSubV1Resp) GetTxnId() string {
	if o == nil || IsNil(o.TxnId) {
		var ret string
		return ret
	}
	return *o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountTransferSubToSubV1Resp) GetTxnIdOk() (*string, bool) {
	if o == nil || IsNil(o.TxnId) {
		return nil, false
	}
	return o.TxnId, true
}

// HasTxnId returns a boolean if a field has been set.
func (o *CreateSubAccountTransferSubToSubV1Resp) HasTxnId() bool {
	if o != nil && !IsNil(o.TxnId) {
		return true
	}

	return false
}

// SetTxnId gets a reference to the given string and assigns it to the TxnId field.
func (o *CreateSubAccountTransferSubToSubV1Resp) SetTxnId(v string) {
	o.TxnId = &v
}

func (o CreateSubAccountTransferSubToSubV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubAccountTransferSubToSubV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TxnId) {
		toSerialize["txnId"] = o.TxnId
	}
	return toSerialize, nil
}

type NullableCreateSubAccountTransferSubToSubV1Resp struct {
	value *CreateSubAccountTransferSubToSubV1Resp
	isSet bool
}

func (v NullableCreateSubAccountTransferSubToSubV1Resp) Get() *CreateSubAccountTransferSubToSubV1Resp {
	return v.value
}

func (v *NullableCreateSubAccountTransferSubToSubV1Resp) Set(val *CreateSubAccountTransferSubToSubV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubAccountTransferSubToSubV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubAccountTransferSubToSubV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubAccountTransferSubToSubV1Resp(val *CreateSubAccountTransferSubToSubV1Resp) *NullableCreateSubAccountTransferSubToSubV1Resp {
	return &NullableCreateSubAccountTransferSubToSubV1Resp{value: val, isSet: true}
}

func (v NullableCreateSubAccountTransferSubToSubV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubAccountTransferSubToSubV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


