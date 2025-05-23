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

// checks if the CreateSimpleEarnFlexibleSubscribeV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSimpleEarnFlexibleSubscribeV1Resp{}

// CreateSimpleEarnFlexibleSubscribeV1Resp struct for CreateSimpleEarnFlexibleSubscribeV1Resp
type CreateSimpleEarnFlexibleSubscribeV1Resp struct {
	PurchaseId *int64 `json:"purchaseId,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewCreateSimpleEarnFlexibleSubscribeV1Resp instantiates a new CreateSimpleEarnFlexibleSubscribeV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSimpleEarnFlexibleSubscribeV1Resp() *CreateSimpleEarnFlexibleSubscribeV1Resp {
	this := CreateSimpleEarnFlexibleSubscribeV1Resp{}
	return &this
}

// NewCreateSimpleEarnFlexibleSubscribeV1RespWithDefaults instantiates a new CreateSimpleEarnFlexibleSubscribeV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSimpleEarnFlexibleSubscribeV1RespWithDefaults() *CreateSimpleEarnFlexibleSubscribeV1Resp {
	this := CreateSimpleEarnFlexibleSubscribeV1Resp{}
	return &this
}

// GetPurchaseId returns the PurchaseId field value if set, zero value otherwise.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) GetPurchaseId() int64 {
	if o == nil || IsNil(o.PurchaseId) {
		var ret int64
		return ret
	}
	return *o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) GetPurchaseIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PurchaseId) {
		return nil, false
	}
	return o.PurchaseId, true
}

// HasPurchaseId returns a boolean if a field has been set.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) HasPurchaseId() bool {
	if o != nil && !IsNil(o.PurchaseId) {
		return true
	}

	return false
}

// SetPurchaseId gets a reference to the given int64 and assigns it to the PurchaseId field.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) SetPurchaseId(v int64) {
	o.PurchaseId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateSimpleEarnFlexibleSubscribeV1Resp) SetSuccess(v bool) {
	o.Success = &v
}

func (o CreateSimpleEarnFlexibleSubscribeV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSimpleEarnFlexibleSubscribeV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PurchaseId) {
		toSerialize["purchaseId"] = o.PurchaseId
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableCreateSimpleEarnFlexibleSubscribeV1Resp struct {
	value *CreateSimpleEarnFlexibleSubscribeV1Resp
	isSet bool
}

func (v NullableCreateSimpleEarnFlexibleSubscribeV1Resp) Get() *CreateSimpleEarnFlexibleSubscribeV1Resp {
	return v.value
}

func (v *NullableCreateSimpleEarnFlexibleSubscribeV1Resp) Set(val *CreateSimpleEarnFlexibleSubscribeV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSimpleEarnFlexibleSubscribeV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSimpleEarnFlexibleSubscribeV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSimpleEarnFlexibleSubscribeV1Resp(val *CreateSimpleEarnFlexibleSubscribeV1Resp) *NullableCreateSimpleEarnFlexibleSubscribeV1Resp {
	return &NullableCreateSimpleEarnFlexibleSubscribeV1Resp{value: val, isSet: true}
}

func (v NullableCreateSimpleEarnFlexibleSubscribeV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSimpleEarnFlexibleSubscribeV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


