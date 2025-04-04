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

// checks if the SpotCreateOrderCancelReplaceV3Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotCreateOrderCancelReplaceV3Data{}

// SpotCreateOrderCancelReplaceV3Data struct for SpotCreateOrderCancelReplaceV3Data
type SpotCreateOrderCancelReplaceV3Data struct {
	CancelResponse *SpotCreateOrderCancelReplaceV3DataCancelResponse `json:"cancelResponse,omitempty"`
	CancelResult *string `json:"cancelResult,omitempty"`
	NewOrderResponse *SpotCreateOrderCancelReplaceV3DataNewOrderResponse `json:"newOrderResponse,omitempty"`
	NewOrderResult *string `json:"newOrderResult,omitempty"`
}

// NewSpotCreateOrderCancelReplaceV3Data instantiates a new SpotCreateOrderCancelReplaceV3Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotCreateOrderCancelReplaceV3Data() *SpotCreateOrderCancelReplaceV3Data {
	this := SpotCreateOrderCancelReplaceV3Data{}
	return &this
}

// NewSpotCreateOrderCancelReplaceV3DataWithDefaults instantiates a new SpotCreateOrderCancelReplaceV3Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotCreateOrderCancelReplaceV3DataWithDefaults() *SpotCreateOrderCancelReplaceV3Data {
	this := SpotCreateOrderCancelReplaceV3Data{}
	return &this
}

// GetCancelResponse returns the CancelResponse field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3Data) GetCancelResponse() SpotCreateOrderCancelReplaceV3DataCancelResponse {
	if o == nil || IsNil(o.CancelResponse) {
		var ret SpotCreateOrderCancelReplaceV3DataCancelResponse
		return ret
	}
	return *o.CancelResponse
}

// GetCancelResponseOk returns a tuple with the CancelResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) GetCancelResponseOk() (*SpotCreateOrderCancelReplaceV3DataCancelResponse, bool) {
	if o == nil || IsNil(o.CancelResponse) {
		return nil, false
	}
	return o.CancelResponse, true
}

// HasCancelResponse returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) HasCancelResponse() bool {
	if o != nil && !IsNil(o.CancelResponse) {
		return true
	}

	return false
}

// SetCancelResponse gets a reference to the given SpotCreateOrderCancelReplaceV3DataCancelResponse and assigns it to the CancelResponse field.
func (o *SpotCreateOrderCancelReplaceV3Data) SetCancelResponse(v SpotCreateOrderCancelReplaceV3DataCancelResponse) {
	o.CancelResponse = &v
}

// GetCancelResult returns the CancelResult field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3Data) GetCancelResult() string {
	if o == nil || IsNil(o.CancelResult) {
		var ret string
		return ret
	}
	return *o.CancelResult
}

// GetCancelResultOk returns a tuple with the CancelResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) GetCancelResultOk() (*string, bool) {
	if o == nil || IsNil(o.CancelResult) {
		return nil, false
	}
	return o.CancelResult, true
}

// HasCancelResult returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) HasCancelResult() bool {
	if o != nil && !IsNil(o.CancelResult) {
		return true
	}

	return false
}

// SetCancelResult gets a reference to the given string and assigns it to the CancelResult field.
func (o *SpotCreateOrderCancelReplaceV3Data) SetCancelResult(v string) {
	o.CancelResult = &v
}

// GetNewOrderResponse returns the NewOrderResponse field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3Data) GetNewOrderResponse() SpotCreateOrderCancelReplaceV3DataNewOrderResponse {
	if o == nil || IsNil(o.NewOrderResponse) {
		var ret SpotCreateOrderCancelReplaceV3DataNewOrderResponse
		return ret
	}
	return *o.NewOrderResponse
}

// GetNewOrderResponseOk returns a tuple with the NewOrderResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) GetNewOrderResponseOk() (*SpotCreateOrderCancelReplaceV3DataNewOrderResponse, bool) {
	if o == nil || IsNil(o.NewOrderResponse) {
		return nil, false
	}
	return o.NewOrderResponse, true
}

// HasNewOrderResponse returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) HasNewOrderResponse() bool {
	if o != nil && !IsNil(o.NewOrderResponse) {
		return true
	}

	return false
}

// SetNewOrderResponse gets a reference to the given SpotCreateOrderCancelReplaceV3DataNewOrderResponse and assigns it to the NewOrderResponse field.
func (o *SpotCreateOrderCancelReplaceV3Data) SetNewOrderResponse(v SpotCreateOrderCancelReplaceV3DataNewOrderResponse) {
	o.NewOrderResponse = &v
}

// GetNewOrderResult returns the NewOrderResult field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3Data) GetNewOrderResult() string {
	if o == nil || IsNil(o.NewOrderResult) {
		var ret string
		return ret
	}
	return *o.NewOrderResult
}

// GetNewOrderResultOk returns a tuple with the NewOrderResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) GetNewOrderResultOk() (*string, bool) {
	if o == nil || IsNil(o.NewOrderResult) {
		return nil, false
	}
	return o.NewOrderResult, true
}

// HasNewOrderResult returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3Data) HasNewOrderResult() bool {
	if o != nil && !IsNil(o.NewOrderResult) {
		return true
	}

	return false
}

// SetNewOrderResult gets a reference to the given string and assigns it to the NewOrderResult field.
func (o *SpotCreateOrderCancelReplaceV3Data) SetNewOrderResult(v string) {
	o.NewOrderResult = &v
}

func (o SpotCreateOrderCancelReplaceV3Data) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotCreateOrderCancelReplaceV3Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CancelResponse) {
		toSerialize["cancelResponse"] = o.CancelResponse
	}
	if !IsNil(o.CancelResult) {
		toSerialize["cancelResult"] = o.CancelResult
	}
	if !IsNil(o.NewOrderResponse) {
		toSerialize["newOrderResponse"] = o.NewOrderResponse
	}
	if !IsNil(o.NewOrderResult) {
		toSerialize["newOrderResult"] = o.NewOrderResult
	}
	return toSerialize, nil
}

type NullableSpotCreateOrderCancelReplaceV3Data struct {
	value *SpotCreateOrderCancelReplaceV3Data
	isSet bool
}

func (v NullableSpotCreateOrderCancelReplaceV3Data) Get() *SpotCreateOrderCancelReplaceV3Data {
	return v.value
}

func (v *NullableSpotCreateOrderCancelReplaceV3Data) Set(val *SpotCreateOrderCancelReplaceV3Data) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotCreateOrderCancelReplaceV3Data) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotCreateOrderCancelReplaceV3Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotCreateOrderCancelReplaceV3Data(val *SpotCreateOrderCancelReplaceV3Data) *NullableSpotCreateOrderCancelReplaceV3Data {
	return &NullableSpotCreateOrderCancelReplaceV3Data{value: val, isSet: true}
}

func (v NullableSpotCreateOrderCancelReplaceV3Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotCreateOrderCancelReplaceV3Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


