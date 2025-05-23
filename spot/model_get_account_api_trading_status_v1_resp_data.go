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

// checks if the GetAccountApiTradingStatusV1RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountApiTradingStatusV1RespData{}

// GetAccountApiTradingStatusV1RespData struct for GetAccountApiTradingStatusV1RespData
type GetAccountApiTradingStatusV1RespData struct {
	IsLocked *bool `json:"isLocked,omitempty"`
	PlannedRecoverTime *int64 `json:"plannedRecoverTime,omitempty"`
	TriggerCondition *GetAccountApiTradingStatusV1RespDataTriggerCondition `json:"triggerCondition,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewGetAccountApiTradingStatusV1RespData instantiates a new GetAccountApiTradingStatusV1RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountApiTradingStatusV1RespData() *GetAccountApiTradingStatusV1RespData {
	this := GetAccountApiTradingStatusV1RespData{}
	return &this
}

// NewGetAccountApiTradingStatusV1RespDataWithDefaults instantiates a new GetAccountApiTradingStatusV1RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountApiTradingStatusV1RespDataWithDefaults() *GetAccountApiTradingStatusV1RespData {
	this := GetAccountApiTradingStatusV1RespData{}
	return &this
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *GetAccountApiTradingStatusV1RespData) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountApiTradingStatusV1RespData) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *GetAccountApiTradingStatusV1RespData) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *GetAccountApiTradingStatusV1RespData) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetPlannedRecoverTime returns the PlannedRecoverTime field value if set, zero value otherwise.
func (o *GetAccountApiTradingStatusV1RespData) GetPlannedRecoverTime() int64 {
	if o == nil || IsNil(o.PlannedRecoverTime) {
		var ret int64
		return ret
	}
	return *o.PlannedRecoverTime
}

// GetPlannedRecoverTimeOk returns a tuple with the PlannedRecoverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountApiTradingStatusV1RespData) GetPlannedRecoverTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PlannedRecoverTime) {
		return nil, false
	}
	return o.PlannedRecoverTime, true
}

// HasPlannedRecoverTime returns a boolean if a field has been set.
func (o *GetAccountApiTradingStatusV1RespData) HasPlannedRecoverTime() bool {
	if o != nil && !IsNil(o.PlannedRecoverTime) {
		return true
	}

	return false
}

// SetPlannedRecoverTime gets a reference to the given int64 and assigns it to the PlannedRecoverTime field.
func (o *GetAccountApiTradingStatusV1RespData) SetPlannedRecoverTime(v int64) {
	o.PlannedRecoverTime = &v
}

// GetTriggerCondition returns the TriggerCondition field value if set, zero value otherwise.
func (o *GetAccountApiTradingStatusV1RespData) GetTriggerCondition() GetAccountApiTradingStatusV1RespDataTriggerCondition {
	if o == nil || IsNil(o.TriggerCondition) {
		var ret GetAccountApiTradingStatusV1RespDataTriggerCondition
		return ret
	}
	return *o.TriggerCondition
}

// GetTriggerConditionOk returns a tuple with the TriggerCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountApiTradingStatusV1RespData) GetTriggerConditionOk() (*GetAccountApiTradingStatusV1RespDataTriggerCondition, bool) {
	if o == nil || IsNil(o.TriggerCondition) {
		return nil, false
	}
	return o.TriggerCondition, true
}

// HasTriggerCondition returns a boolean if a field has been set.
func (o *GetAccountApiTradingStatusV1RespData) HasTriggerCondition() bool {
	if o != nil && !IsNil(o.TriggerCondition) {
		return true
	}

	return false
}

// SetTriggerCondition gets a reference to the given GetAccountApiTradingStatusV1RespDataTriggerCondition and assigns it to the TriggerCondition field.
func (o *GetAccountApiTradingStatusV1RespData) SetTriggerCondition(v GetAccountApiTradingStatusV1RespDataTriggerCondition) {
	o.TriggerCondition = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetAccountApiTradingStatusV1RespData) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountApiTradingStatusV1RespData) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetAccountApiTradingStatusV1RespData) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetAccountApiTradingStatusV1RespData) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetAccountApiTradingStatusV1RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountApiTradingStatusV1RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.PlannedRecoverTime) {
		toSerialize["plannedRecoverTime"] = o.PlannedRecoverTime
	}
	if !IsNil(o.TriggerCondition) {
		toSerialize["triggerCondition"] = o.TriggerCondition
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableGetAccountApiTradingStatusV1RespData struct {
	value *GetAccountApiTradingStatusV1RespData
	isSet bool
}

func (v NullableGetAccountApiTradingStatusV1RespData) Get() *GetAccountApiTradingStatusV1RespData {
	return v.value
}

func (v *NullableGetAccountApiTradingStatusV1RespData) Set(val *GetAccountApiTradingStatusV1RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountApiTradingStatusV1RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountApiTradingStatusV1RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountApiTradingStatusV1RespData(val *GetAccountApiTradingStatusV1RespData) *NullableGetAccountApiTradingStatusV1RespData {
	return &NullableGetAccountApiTradingStatusV1RespData{value: val, isSet: true}
}

func (v NullableGetAccountApiTradingStatusV1RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountApiTradingStatusV1RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


