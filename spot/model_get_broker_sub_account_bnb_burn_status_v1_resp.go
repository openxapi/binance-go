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

// checks if the GetBrokerSubAccountBnbBurnStatusV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBrokerSubAccountBnbBurnStatusV1Resp{}

// GetBrokerSubAccountBnbBurnStatusV1Resp struct for GetBrokerSubAccountBnbBurnStatusV1Resp
type GetBrokerSubAccountBnbBurnStatusV1Resp struct {
	InterestBNBBurn *bool `json:"interestBNBBurn,omitempty"`
	SpotBNBBurn *bool `json:"spotBNBBurn,omitempty"`
	SubAccountId *int64 `json:"subAccountId,omitempty"`
}

// NewGetBrokerSubAccountBnbBurnStatusV1Resp instantiates a new GetBrokerSubAccountBnbBurnStatusV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBrokerSubAccountBnbBurnStatusV1Resp() *GetBrokerSubAccountBnbBurnStatusV1Resp {
	this := GetBrokerSubAccountBnbBurnStatusV1Resp{}
	return &this
}

// NewGetBrokerSubAccountBnbBurnStatusV1RespWithDefaults instantiates a new GetBrokerSubAccountBnbBurnStatusV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBrokerSubAccountBnbBurnStatusV1RespWithDefaults() *GetBrokerSubAccountBnbBurnStatusV1Resp {
	this := GetBrokerSubAccountBnbBurnStatusV1Resp{}
	return &this
}

// GetInterestBNBBurn returns the InterestBNBBurn field value if set, zero value otherwise.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) GetInterestBNBBurn() bool {
	if o == nil || IsNil(o.InterestBNBBurn) {
		var ret bool
		return ret
	}
	return *o.InterestBNBBurn
}

// GetInterestBNBBurnOk returns a tuple with the InterestBNBBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) GetInterestBNBBurnOk() (*bool, bool) {
	if o == nil || IsNil(o.InterestBNBBurn) {
		return nil, false
	}
	return o.InterestBNBBurn, true
}

// HasInterestBNBBurn returns a boolean if a field has been set.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) HasInterestBNBBurn() bool {
	if o != nil && !IsNil(o.InterestBNBBurn) {
		return true
	}

	return false
}

// SetInterestBNBBurn gets a reference to the given bool and assigns it to the InterestBNBBurn field.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) SetInterestBNBBurn(v bool) {
	o.InterestBNBBurn = &v
}

// GetSpotBNBBurn returns the SpotBNBBurn field value if set, zero value otherwise.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) GetSpotBNBBurn() bool {
	if o == nil || IsNil(o.SpotBNBBurn) {
		var ret bool
		return ret
	}
	return *o.SpotBNBBurn
}

// GetSpotBNBBurnOk returns a tuple with the SpotBNBBurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) GetSpotBNBBurnOk() (*bool, bool) {
	if o == nil || IsNil(o.SpotBNBBurn) {
		return nil, false
	}
	return o.SpotBNBBurn, true
}

// HasSpotBNBBurn returns a boolean if a field has been set.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) HasSpotBNBBurn() bool {
	if o != nil && !IsNil(o.SpotBNBBurn) {
		return true
	}

	return false
}

// SetSpotBNBBurn gets a reference to the given bool and assigns it to the SpotBNBBurn field.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) SetSpotBNBBurn(v bool) {
	o.SpotBNBBurn = &v
}

// GetSubAccountId returns the SubAccountId field value if set, zero value otherwise.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) GetSubAccountId() int64 {
	if o == nil || IsNil(o.SubAccountId) {
		var ret int64
		return ret
	}
	return *o.SubAccountId
}

// GetSubAccountIdOk returns a tuple with the SubAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) GetSubAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SubAccountId) {
		return nil, false
	}
	return o.SubAccountId, true
}

// HasSubAccountId returns a boolean if a field has been set.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) HasSubAccountId() bool {
	if o != nil && !IsNil(o.SubAccountId) {
		return true
	}

	return false
}

// SetSubAccountId gets a reference to the given int64 and assigns it to the SubAccountId field.
func (o *GetBrokerSubAccountBnbBurnStatusV1Resp) SetSubAccountId(v int64) {
	o.SubAccountId = &v
}

func (o GetBrokerSubAccountBnbBurnStatusV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBrokerSubAccountBnbBurnStatusV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterestBNBBurn) {
		toSerialize["interestBNBBurn"] = o.InterestBNBBurn
	}
	if !IsNil(o.SpotBNBBurn) {
		toSerialize["spotBNBBurn"] = o.SpotBNBBurn
	}
	if !IsNil(o.SubAccountId) {
		toSerialize["subAccountId"] = o.SubAccountId
	}
	return toSerialize, nil
}

type NullableGetBrokerSubAccountBnbBurnStatusV1Resp struct {
	value *GetBrokerSubAccountBnbBurnStatusV1Resp
	isSet bool
}

func (v NullableGetBrokerSubAccountBnbBurnStatusV1Resp) Get() *GetBrokerSubAccountBnbBurnStatusV1Resp {
	return v.value
}

func (v *NullableGetBrokerSubAccountBnbBurnStatusV1Resp) Set(val *GetBrokerSubAccountBnbBurnStatusV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBrokerSubAccountBnbBurnStatusV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBrokerSubAccountBnbBurnStatusV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBrokerSubAccountBnbBurnStatusV1Resp(val *GetBrokerSubAccountBnbBurnStatusV1Resp) *NullableGetBrokerSubAccountBnbBurnStatusV1Resp {
	return &NullableGetBrokerSubAccountBnbBurnStatusV1Resp{value: val, isSet: true}
}

func (v NullableGetBrokerSubAccountBnbBurnStatusV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBrokerSubAccountBnbBurnStatusV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


