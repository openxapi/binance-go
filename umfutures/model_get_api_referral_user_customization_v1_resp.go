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

// checks if the GetApiReferralUserCustomizationV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetApiReferralUserCustomizationV1Resp{}

// GetApiReferralUserCustomizationV1Resp struct for GetApiReferralUserCustomizationV1Resp
type GetApiReferralUserCustomizationV1Resp struct {
	BrokerId *string `json:"brokerId,omitempty"`
	CustomerId *string `json:"customerId,omitempty"`
}

// NewGetApiReferralUserCustomizationV1Resp instantiates a new GetApiReferralUserCustomizationV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetApiReferralUserCustomizationV1Resp() *GetApiReferralUserCustomizationV1Resp {
	this := GetApiReferralUserCustomizationV1Resp{}
	return &this
}

// NewGetApiReferralUserCustomizationV1RespWithDefaults instantiates a new GetApiReferralUserCustomizationV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetApiReferralUserCustomizationV1RespWithDefaults() *GetApiReferralUserCustomizationV1Resp {
	this := GetApiReferralUserCustomizationV1Resp{}
	return &this
}

// GetBrokerId returns the BrokerId field value if set, zero value otherwise.
func (o *GetApiReferralUserCustomizationV1Resp) GetBrokerId() string {
	if o == nil || IsNil(o.BrokerId) {
		var ret string
		return ret
	}
	return *o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiReferralUserCustomizationV1Resp) GetBrokerIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerId) {
		return nil, false
	}
	return o.BrokerId, true
}

// HasBrokerId returns a boolean if a field has been set.
func (o *GetApiReferralUserCustomizationV1Resp) HasBrokerId() bool {
	if o != nil && !IsNil(o.BrokerId) {
		return true
	}

	return false
}

// SetBrokerId gets a reference to the given string and assigns it to the BrokerId field.
func (o *GetApiReferralUserCustomizationV1Resp) SetBrokerId(v string) {
	o.BrokerId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *GetApiReferralUserCustomizationV1Resp) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetApiReferralUserCustomizationV1Resp) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *GetApiReferralUserCustomizationV1Resp) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *GetApiReferralUserCustomizationV1Resp) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o GetApiReferralUserCustomizationV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetApiReferralUserCustomizationV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrokerId) {
		toSerialize["brokerId"] = o.BrokerId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	return toSerialize, nil
}

type NullableGetApiReferralUserCustomizationV1Resp struct {
	value *GetApiReferralUserCustomizationV1Resp
	isSet bool
}

func (v NullableGetApiReferralUserCustomizationV1Resp) Get() *GetApiReferralUserCustomizationV1Resp {
	return v.value
}

func (v *NullableGetApiReferralUserCustomizationV1Resp) Set(val *GetApiReferralUserCustomizationV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApiReferralUserCustomizationV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApiReferralUserCustomizationV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApiReferralUserCustomizationV1Resp(val *GetApiReferralUserCustomizationV1Resp) *NullableGetApiReferralUserCustomizationV1Resp {
	return &NullableGetApiReferralUserCustomizationV1Resp{value: val, isSet: true}
}

func (v NullableGetApiReferralUserCustomizationV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApiReferralUserCustomizationV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


