/*
Binance Sub Account API

OpenAPI specification for Binance exchange - Subaccount API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subaccount

import (
	"encoding/json"
)

// checks if the SubaccountCreateSubAccountUniversalTransferV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountCreateSubAccountUniversalTransferV1Resp{}

// SubaccountCreateSubAccountUniversalTransferV1Resp struct for SubaccountCreateSubAccountUniversalTransferV1Resp
type SubaccountCreateSubAccountUniversalTransferV1Resp struct {
	ClientTranId *string `json:"clientTranId,omitempty"`
	TranId *int64 `json:"tranId,omitempty"`
}

// NewSubaccountCreateSubAccountUniversalTransferV1Resp instantiates a new SubaccountCreateSubAccountUniversalTransferV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountCreateSubAccountUniversalTransferV1Resp() *SubaccountCreateSubAccountUniversalTransferV1Resp {
	this := SubaccountCreateSubAccountUniversalTransferV1Resp{}
	return &this
}

// NewSubaccountCreateSubAccountUniversalTransferV1RespWithDefaults instantiates a new SubaccountCreateSubAccountUniversalTransferV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountCreateSubAccountUniversalTransferV1RespWithDefaults() *SubaccountCreateSubAccountUniversalTransferV1Resp {
	this := SubaccountCreateSubAccountUniversalTransferV1Resp{}
	return &this
}

// GetClientTranId returns the ClientTranId field value if set, zero value otherwise.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) GetClientTranId() string {
	if o == nil || IsNil(o.ClientTranId) {
		var ret string
		return ret
	}
	return *o.ClientTranId
}

// GetClientTranIdOk returns a tuple with the ClientTranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) GetClientTranIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientTranId) {
		return nil, false
	}
	return o.ClientTranId, true
}

// HasClientTranId returns a boolean if a field has been set.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) HasClientTranId() bool {
	if o != nil && !IsNil(o.ClientTranId) {
		return true
	}

	return false
}

// SetClientTranId gets a reference to the given string and assigns it to the ClientTranId field.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) SetClientTranId(v string) {
	o.ClientTranId = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) GetTranId() int64 {
	if o == nil || IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) GetTranIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) HasTranId() bool {
	if o != nil && !IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *SubaccountCreateSubAccountUniversalTransferV1Resp) SetTranId(v int64) {
	o.TranId = &v
}

func (o SubaccountCreateSubAccountUniversalTransferV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountCreateSubAccountUniversalTransferV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientTranId) {
		toSerialize["clientTranId"] = o.ClientTranId
	}
	if !IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	return toSerialize, nil
}

type NullableSubaccountCreateSubAccountUniversalTransferV1Resp struct {
	value *SubaccountCreateSubAccountUniversalTransferV1Resp
	isSet bool
}

func (v NullableSubaccountCreateSubAccountUniversalTransferV1Resp) Get() *SubaccountCreateSubAccountUniversalTransferV1Resp {
	return v.value
}

func (v *NullableSubaccountCreateSubAccountUniversalTransferV1Resp) Set(val *SubaccountCreateSubAccountUniversalTransferV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountCreateSubAccountUniversalTransferV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountCreateSubAccountUniversalTransferV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountCreateSubAccountUniversalTransferV1Resp(val *SubaccountCreateSubAccountUniversalTransferV1Resp) *NullableSubaccountCreateSubAccountUniversalTransferV1Resp {
	return &NullableSubaccountCreateSubAccountUniversalTransferV1Resp{value: val, isSet: true}
}

func (v NullableSubaccountCreateSubAccountUniversalTransferV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountCreateSubAccountUniversalTransferV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


