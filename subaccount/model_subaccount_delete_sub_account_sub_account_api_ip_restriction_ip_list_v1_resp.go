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

// checks if the SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp{}

// SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp struct for SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp
type SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp struct {
	ApiKey *string `json:"apiKey,omitempty"`
	IpList []string `json:"ipList,omitempty"`
	IpRestrict *string `json:"ipRestrict,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp instantiates a new SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp() *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp {
	this := SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp{}
	return &this
}

// NewSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1RespWithDefaults instantiates a new SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1RespWithDefaults() *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp {
	this := SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetIpList() []string {
	if o == nil || IsNil(o.IpList) {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetIpListOk() ([]string, bool) {
	if o == nil || IsNil(o.IpList) {
		return nil, false
	}
	return o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) HasIpList() bool {
	if o != nil && !IsNil(o.IpList) {
		return true
	}

	return false
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) SetIpList(v []string) {
	o.IpList = v
}

// GetIpRestrict returns the IpRestrict field value if set, zero value otherwise.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetIpRestrict() string {
	if o == nil || IsNil(o.IpRestrict) {
		var ret string
		return ret
	}
	return *o.IpRestrict
}

// GetIpRestrictOk returns a tuple with the IpRestrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetIpRestrictOk() (*string, bool) {
	if o == nil || IsNil(o.IpRestrict) {
		return nil, false
	}
	return o.IpRestrict, true
}

// HasIpRestrict returns a boolean if a field has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) HasIpRestrict() bool {
	if o != nil && !IsNil(o.IpRestrict) {
		return true
	}

	return false
}

// SetIpRestrict gets a reference to the given string and assigns it to the IpRestrict field.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) SetIpRestrict(v string) {
	o.IpRestrict = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.IpList) {
		toSerialize["ipList"] = o.IpList
	}
	if !IsNil(o.IpRestrict) {
		toSerialize["ipRestrict"] = o.IpRestrict
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp struct {
	value *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp
	isSet bool
}

func (v NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) Get() *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp {
	return v.value
}

func (v *NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) Set(val *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp(val *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) *NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp {
	return &NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp{value: val, isSet: true}
}

func (v NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


