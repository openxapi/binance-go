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

// checks if the GetMmpV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMmpV1Resp{}

// GetMmpV1Resp struct for GetMmpV1Resp
type GetMmpV1Resp struct {
	DeltaLimit *string `json:"deltaLimit,omitempty"`
	FrozenTimeInMilliseconds *int32 `json:"frozenTimeInMilliseconds,omitempty"`
	LastTriggerTime *int64 `json:"lastTriggerTime,omitempty"`
	QtyLimit *string `json:"qtyLimit,omitempty"`
	Underlying *string `json:"underlying,omitempty"`
	UnderlyingId *int64 `json:"underlyingId,omitempty"`
	WindowTimeInMilliseconds *int32 `json:"windowTimeInMilliseconds,omitempty"`
}

// NewGetMmpV1Resp instantiates a new GetMmpV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMmpV1Resp() *GetMmpV1Resp {
	this := GetMmpV1Resp{}
	return &this
}

// NewGetMmpV1RespWithDefaults instantiates a new GetMmpV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMmpV1RespWithDefaults() *GetMmpV1Resp {
	this := GetMmpV1Resp{}
	return &this
}

// GetDeltaLimit returns the DeltaLimit field value if set, zero value otherwise.
func (o *GetMmpV1Resp) GetDeltaLimit() string {
	if o == nil || IsNil(o.DeltaLimit) {
		var ret string
		return ret
	}
	return *o.DeltaLimit
}

// GetDeltaLimitOk returns a tuple with the DeltaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMmpV1Resp) GetDeltaLimitOk() (*string, bool) {
	if o == nil || IsNil(o.DeltaLimit) {
		return nil, false
	}
	return o.DeltaLimit, true
}

// HasDeltaLimit returns a boolean if a field has been set.
func (o *GetMmpV1Resp) HasDeltaLimit() bool {
	if o != nil && !IsNil(o.DeltaLimit) {
		return true
	}

	return false
}

// SetDeltaLimit gets a reference to the given string and assigns it to the DeltaLimit field.
func (o *GetMmpV1Resp) SetDeltaLimit(v string) {
	o.DeltaLimit = &v
}

// GetFrozenTimeInMilliseconds returns the FrozenTimeInMilliseconds field value if set, zero value otherwise.
func (o *GetMmpV1Resp) GetFrozenTimeInMilliseconds() int32 {
	if o == nil || IsNil(o.FrozenTimeInMilliseconds) {
		var ret int32
		return ret
	}
	return *o.FrozenTimeInMilliseconds
}

// GetFrozenTimeInMillisecondsOk returns a tuple with the FrozenTimeInMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMmpV1Resp) GetFrozenTimeInMillisecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.FrozenTimeInMilliseconds) {
		return nil, false
	}
	return o.FrozenTimeInMilliseconds, true
}

// HasFrozenTimeInMilliseconds returns a boolean if a field has been set.
func (o *GetMmpV1Resp) HasFrozenTimeInMilliseconds() bool {
	if o != nil && !IsNil(o.FrozenTimeInMilliseconds) {
		return true
	}

	return false
}

// SetFrozenTimeInMilliseconds gets a reference to the given int32 and assigns it to the FrozenTimeInMilliseconds field.
func (o *GetMmpV1Resp) SetFrozenTimeInMilliseconds(v int32) {
	o.FrozenTimeInMilliseconds = &v
}

// GetLastTriggerTime returns the LastTriggerTime field value if set, zero value otherwise.
func (o *GetMmpV1Resp) GetLastTriggerTime() int64 {
	if o == nil || IsNil(o.LastTriggerTime) {
		var ret int64
		return ret
	}
	return *o.LastTriggerTime
}

// GetLastTriggerTimeOk returns a tuple with the LastTriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMmpV1Resp) GetLastTriggerTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.LastTriggerTime) {
		return nil, false
	}
	return o.LastTriggerTime, true
}

// HasLastTriggerTime returns a boolean if a field has been set.
func (o *GetMmpV1Resp) HasLastTriggerTime() bool {
	if o != nil && !IsNil(o.LastTriggerTime) {
		return true
	}

	return false
}

// SetLastTriggerTime gets a reference to the given int64 and assigns it to the LastTriggerTime field.
func (o *GetMmpV1Resp) SetLastTriggerTime(v int64) {
	o.LastTriggerTime = &v
}

// GetQtyLimit returns the QtyLimit field value if set, zero value otherwise.
func (o *GetMmpV1Resp) GetQtyLimit() string {
	if o == nil || IsNil(o.QtyLimit) {
		var ret string
		return ret
	}
	return *o.QtyLimit
}

// GetQtyLimitOk returns a tuple with the QtyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMmpV1Resp) GetQtyLimitOk() (*string, bool) {
	if o == nil || IsNil(o.QtyLimit) {
		return nil, false
	}
	return o.QtyLimit, true
}

// HasQtyLimit returns a boolean if a field has been set.
func (o *GetMmpV1Resp) HasQtyLimit() bool {
	if o != nil && !IsNil(o.QtyLimit) {
		return true
	}

	return false
}

// SetQtyLimit gets a reference to the given string and assigns it to the QtyLimit field.
func (o *GetMmpV1Resp) SetQtyLimit(v string) {
	o.QtyLimit = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *GetMmpV1Resp) GetUnderlying() string {
	if o == nil || IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMmpV1Resp) GetUnderlyingOk() (*string, bool) {
	if o == nil || IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *GetMmpV1Resp) HasUnderlying() bool {
	if o != nil && !IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *GetMmpV1Resp) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetUnderlyingId returns the UnderlyingId field value if set, zero value otherwise.
func (o *GetMmpV1Resp) GetUnderlyingId() int64 {
	if o == nil || IsNil(o.UnderlyingId) {
		var ret int64
		return ret
	}
	return *o.UnderlyingId
}

// GetUnderlyingIdOk returns a tuple with the UnderlyingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMmpV1Resp) GetUnderlyingIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UnderlyingId) {
		return nil, false
	}
	return o.UnderlyingId, true
}

// HasUnderlyingId returns a boolean if a field has been set.
func (o *GetMmpV1Resp) HasUnderlyingId() bool {
	if o != nil && !IsNil(o.UnderlyingId) {
		return true
	}

	return false
}

// SetUnderlyingId gets a reference to the given int64 and assigns it to the UnderlyingId field.
func (o *GetMmpV1Resp) SetUnderlyingId(v int64) {
	o.UnderlyingId = &v
}

// GetWindowTimeInMilliseconds returns the WindowTimeInMilliseconds field value if set, zero value otherwise.
func (o *GetMmpV1Resp) GetWindowTimeInMilliseconds() int32 {
	if o == nil || IsNil(o.WindowTimeInMilliseconds) {
		var ret int32
		return ret
	}
	return *o.WindowTimeInMilliseconds
}

// GetWindowTimeInMillisecondsOk returns a tuple with the WindowTimeInMilliseconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMmpV1Resp) GetWindowTimeInMillisecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.WindowTimeInMilliseconds) {
		return nil, false
	}
	return o.WindowTimeInMilliseconds, true
}

// HasWindowTimeInMilliseconds returns a boolean if a field has been set.
func (o *GetMmpV1Resp) HasWindowTimeInMilliseconds() bool {
	if o != nil && !IsNil(o.WindowTimeInMilliseconds) {
		return true
	}

	return false
}

// SetWindowTimeInMilliseconds gets a reference to the given int32 and assigns it to the WindowTimeInMilliseconds field.
func (o *GetMmpV1Resp) SetWindowTimeInMilliseconds(v int32) {
	o.WindowTimeInMilliseconds = &v
}

func (o GetMmpV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMmpV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeltaLimit) {
		toSerialize["deltaLimit"] = o.DeltaLimit
	}
	if !IsNil(o.FrozenTimeInMilliseconds) {
		toSerialize["frozenTimeInMilliseconds"] = o.FrozenTimeInMilliseconds
	}
	if !IsNil(o.LastTriggerTime) {
		toSerialize["lastTriggerTime"] = o.LastTriggerTime
	}
	if !IsNil(o.QtyLimit) {
		toSerialize["qtyLimit"] = o.QtyLimit
	}
	if !IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !IsNil(o.UnderlyingId) {
		toSerialize["underlyingId"] = o.UnderlyingId
	}
	if !IsNil(o.WindowTimeInMilliseconds) {
		toSerialize["windowTimeInMilliseconds"] = o.WindowTimeInMilliseconds
	}
	return toSerialize, nil
}

type NullableGetMmpV1Resp struct {
	value *GetMmpV1Resp
	isSet bool
}

func (v NullableGetMmpV1Resp) Get() *GetMmpV1Resp {
	return v.value
}

func (v *NullableGetMmpV1Resp) Set(val *GetMmpV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMmpV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMmpV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMmpV1Resp(val *GetMmpV1Resp) *NullableGetMmpV1Resp {
	return &NullableGetMmpV1Resp{value: val, isSet: true}
}

func (v NullableGetMmpV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMmpV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


