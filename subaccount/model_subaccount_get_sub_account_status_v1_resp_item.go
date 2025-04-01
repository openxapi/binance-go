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

// checks if the SubaccountGetSubAccountStatusV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountGetSubAccountStatusV1RespItem{}

// SubaccountGetSubAccountStatusV1RespItem struct for SubaccountGetSubAccountStatusV1RespItem
type SubaccountGetSubAccountStatusV1RespItem struct {
	Email *string `json:"email,omitempty"`
	InsertTime *int64 `json:"insertTime,omitempty"`
	IsFutureEnabled *bool `json:"isFutureEnabled,omitempty"`
	IsMarginEnabled *bool `json:"isMarginEnabled,omitempty"`
	IsSubUserEnabled *bool `json:"isSubUserEnabled,omitempty"`
	IsUserActive *bool `json:"isUserActive,omitempty"`
	Mobile *int32 `json:"mobile,omitempty"`
}

// NewSubaccountGetSubAccountStatusV1RespItem instantiates a new SubaccountGetSubAccountStatusV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountGetSubAccountStatusV1RespItem() *SubaccountGetSubAccountStatusV1RespItem {
	this := SubaccountGetSubAccountStatusV1RespItem{}
	return &this
}

// NewSubaccountGetSubAccountStatusV1RespItemWithDefaults instantiates a new SubaccountGetSubAccountStatusV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountGetSubAccountStatusV1RespItemWithDefaults() *SubaccountGetSubAccountStatusV1RespItem {
	this := SubaccountGetSubAccountStatusV1RespItem{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SubaccountGetSubAccountStatusV1RespItem) SetEmail(v string) {
	o.Email = &v
}

// GetInsertTime returns the InsertTime field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetInsertTime() int64 {
	if o == nil || IsNil(o.InsertTime) {
		var ret int64
		return ret
	}
	return *o.InsertTime
}

// GetInsertTimeOk returns a tuple with the InsertTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetInsertTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.InsertTime) {
		return nil, false
	}
	return o.InsertTime, true
}

// HasInsertTime returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) HasInsertTime() bool {
	if o != nil && !IsNil(o.InsertTime) {
		return true
	}

	return false
}

// SetInsertTime gets a reference to the given int64 and assigns it to the InsertTime field.
func (o *SubaccountGetSubAccountStatusV1RespItem) SetInsertTime(v int64) {
	o.InsertTime = &v
}

// GetIsFutureEnabled returns the IsFutureEnabled field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsFutureEnabled() bool {
	if o == nil || IsNil(o.IsFutureEnabled) {
		var ret bool
		return ret
	}
	return *o.IsFutureEnabled
}

// GetIsFutureEnabledOk returns a tuple with the IsFutureEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsFutureEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFutureEnabled) {
		return nil, false
	}
	return o.IsFutureEnabled, true
}

// HasIsFutureEnabled returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) HasIsFutureEnabled() bool {
	if o != nil && !IsNil(o.IsFutureEnabled) {
		return true
	}

	return false
}

// SetIsFutureEnabled gets a reference to the given bool and assigns it to the IsFutureEnabled field.
func (o *SubaccountGetSubAccountStatusV1RespItem) SetIsFutureEnabled(v bool) {
	o.IsFutureEnabled = &v
}

// GetIsMarginEnabled returns the IsMarginEnabled field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsMarginEnabled() bool {
	if o == nil || IsNil(o.IsMarginEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMarginEnabled
}

// GetIsMarginEnabledOk returns a tuple with the IsMarginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsMarginEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMarginEnabled) {
		return nil, false
	}
	return o.IsMarginEnabled, true
}

// HasIsMarginEnabled returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) HasIsMarginEnabled() bool {
	if o != nil && !IsNil(o.IsMarginEnabled) {
		return true
	}

	return false
}

// SetIsMarginEnabled gets a reference to the given bool and assigns it to the IsMarginEnabled field.
func (o *SubaccountGetSubAccountStatusV1RespItem) SetIsMarginEnabled(v bool) {
	o.IsMarginEnabled = &v
}

// GetIsSubUserEnabled returns the IsSubUserEnabled field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsSubUserEnabled() bool {
	if o == nil || IsNil(o.IsSubUserEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSubUserEnabled
}

// GetIsSubUserEnabledOk returns a tuple with the IsSubUserEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsSubUserEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubUserEnabled) {
		return nil, false
	}
	return o.IsSubUserEnabled, true
}

// HasIsSubUserEnabled returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) HasIsSubUserEnabled() bool {
	if o != nil && !IsNil(o.IsSubUserEnabled) {
		return true
	}

	return false
}

// SetIsSubUserEnabled gets a reference to the given bool and assigns it to the IsSubUserEnabled field.
func (o *SubaccountGetSubAccountStatusV1RespItem) SetIsSubUserEnabled(v bool) {
	o.IsSubUserEnabled = &v
}

// GetIsUserActive returns the IsUserActive field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsUserActive() bool {
	if o == nil || IsNil(o.IsUserActive) {
		var ret bool
		return ret
	}
	return *o.IsUserActive
}

// GetIsUserActiveOk returns a tuple with the IsUserActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetIsUserActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUserActive) {
		return nil, false
	}
	return o.IsUserActive, true
}

// HasIsUserActive returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) HasIsUserActive() bool {
	if o != nil && !IsNil(o.IsUserActive) {
		return true
	}

	return false
}

// SetIsUserActive gets a reference to the given bool and assigns it to the IsUserActive field.
func (o *SubaccountGetSubAccountStatusV1RespItem) SetIsUserActive(v bool) {
	o.IsUserActive = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetMobile() int32 {
	if o == nil || IsNil(o.Mobile) {
		var ret int32
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) GetMobileOk() (*int32, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountStatusV1RespItem) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given int32 and assigns it to the Mobile field.
func (o *SubaccountGetSubAccountStatusV1RespItem) SetMobile(v int32) {
	o.Mobile = &v
}

func (o SubaccountGetSubAccountStatusV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountGetSubAccountStatusV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.InsertTime) {
		toSerialize["insertTime"] = o.InsertTime
	}
	if !IsNil(o.IsFutureEnabled) {
		toSerialize["isFutureEnabled"] = o.IsFutureEnabled
	}
	if !IsNil(o.IsMarginEnabled) {
		toSerialize["isMarginEnabled"] = o.IsMarginEnabled
	}
	if !IsNil(o.IsSubUserEnabled) {
		toSerialize["isSubUserEnabled"] = o.IsSubUserEnabled
	}
	if !IsNil(o.IsUserActive) {
		toSerialize["isUserActive"] = o.IsUserActive
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	return toSerialize, nil
}

type NullableSubaccountGetSubAccountStatusV1RespItem struct {
	value *SubaccountGetSubAccountStatusV1RespItem
	isSet bool
}

func (v NullableSubaccountGetSubAccountStatusV1RespItem) Get() *SubaccountGetSubAccountStatusV1RespItem {
	return v.value
}

func (v *NullableSubaccountGetSubAccountStatusV1RespItem) Set(val *SubaccountGetSubAccountStatusV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountGetSubAccountStatusV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountGetSubAccountStatusV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountGetSubAccountStatusV1RespItem(val *SubaccountGetSubAccountStatusV1RespItem) *NullableSubaccountGetSubAccountStatusV1RespItem {
	return &NullableSubaccountGetSubAccountStatusV1RespItem{value: val, isSet: true}
}

func (v NullableSubaccountGetSubAccountStatusV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountGetSubAccountStatusV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


