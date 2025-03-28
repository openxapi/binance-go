/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetAggTradesV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetAggTradesV1RespItem{}

// UfuturesGetAggTradesV1RespItem struct for UfuturesGetAggTradesV1RespItem
type UfuturesGetAggTradesV1RespItem struct {
	T *int64 `json:"T,omitempty"`
	A *int32 `json:"a,omitempty"`
	F *int32 `json:"f,omitempty"`
	L *int32 `json:"l,omitempty"`
	M *bool `json:"m,omitempty"`
	P *string `json:"p,omitempty"`
	Q *string `json:"q,omitempty"`
}

// NewUfuturesGetAggTradesV1RespItem instantiates a new UfuturesGetAggTradesV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetAggTradesV1RespItem() *UfuturesGetAggTradesV1RespItem {
	this := UfuturesGetAggTradesV1RespItem{}
	return &this
}

// NewUfuturesGetAggTradesV1RespItemWithDefaults instantiates a new UfuturesGetAggTradesV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetAggTradesV1RespItemWithDefaults() *UfuturesGetAggTradesV1RespItem {
	this := UfuturesGetAggTradesV1RespItem{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *UfuturesGetAggTradesV1RespItem) GetT() int64 {
	if o == nil || IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAggTradesV1RespItem) GetTOk() (*int64, bool) {
	if o == nil || IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *UfuturesGetAggTradesV1RespItem) HasT() bool {
	if o != nil && !IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *UfuturesGetAggTradesV1RespItem) SetT(v int64) {
	o.T = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *UfuturesGetAggTradesV1RespItem) GetA() int32 {
	if o == nil || IsNil(o.A) {
		var ret int32
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAggTradesV1RespItem) GetAOk() (*int32, bool) {
	if o == nil || IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *UfuturesGetAggTradesV1RespItem) HasA() bool {
	if o != nil && !IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given int32 and assigns it to the A field.
func (o *UfuturesGetAggTradesV1RespItem) SetA(v int32) {
	o.A = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *UfuturesGetAggTradesV1RespItem) GetF() int32 {
	if o == nil || IsNil(o.F) {
		var ret int32
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAggTradesV1RespItem) GetFOk() (*int32, bool) {
	if o == nil || IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *UfuturesGetAggTradesV1RespItem) HasF() bool {
	if o != nil && !IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int32 and assigns it to the F field.
func (o *UfuturesGetAggTradesV1RespItem) SetF(v int32) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *UfuturesGetAggTradesV1RespItem) GetL() int32 {
	if o == nil || IsNil(o.L) {
		var ret int32
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAggTradesV1RespItem) GetLOk() (*int32, bool) {
	if o == nil || IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *UfuturesGetAggTradesV1RespItem) HasL() bool {
	if o != nil && !IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int32 and assigns it to the L field.
func (o *UfuturesGetAggTradesV1RespItem) SetL(v int32) {
	o.L = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *UfuturesGetAggTradesV1RespItem) GetM() bool {
	if o == nil || IsNil(o.M) {
		var ret bool
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAggTradesV1RespItem) GetMOk() (*bool, bool) {
	if o == nil || IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *UfuturesGetAggTradesV1RespItem) HasM() bool {
	if o != nil && !IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *UfuturesGetAggTradesV1RespItem) SetM(v bool) {
	o.M = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *UfuturesGetAggTradesV1RespItem) GetP() string {
	if o == nil || IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAggTradesV1RespItem) GetPOk() (*string, bool) {
	if o == nil || IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *UfuturesGetAggTradesV1RespItem) HasP() bool {
	if o != nil && !IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *UfuturesGetAggTradesV1RespItem) SetP(v string) {
	o.P = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *UfuturesGetAggTradesV1RespItem) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAggTradesV1RespItem) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *UfuturesGetAggTradesV1RespItem) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *UfuturesGetAggTradesV1RespItem) SetQ(v string) {
	o.Q = &v
}

func (o UfuturesGetAggTradesV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetAggTradesV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	if !IsNil(o.F) {
		toSerialize["f"] = o.F
	}
	if !IsNil(o.L) {
		toSerialize["l"] = o.L
	}
	if !IsNil(o.M) {
		toSerialize["m"] = o.M
	}
	if !IsNil(o.P) {
		toSerialize["p"] = o.P
	}
	if !IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}
	return toSerialize, nil
}

type NullableUfuturesGetAggTradesV1RespItem struct {
	value *UfuturesGetAggTradesV1RespItem
	isSet bool
}

func (v NullableUfuturesGetAggTradesV1RespItem) Get() *UfuturesGetAggTradesV1RespItem {
	return v.value
}

func (v *NullableUfuturesGetAggTradesV1RespItem) Set(val *UfuturesGetAggTradesV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetAggTradesV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetAggTradesV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetAggTradesV1RespItem(val *UfuturesGetAggTradesV1RespItem) *NullableUfuturesGetAggTradesV1RespItem {
	return &NullableUfuturesGetAggTradesV1RespItem{value: val, isSet: true}
}

func (v NullableUfuturesGetAggTradesV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetAggTradesV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


