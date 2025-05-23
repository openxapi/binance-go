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

// checks if the CreateSubAccountSubAccountApiIpRestrictionV2Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubAccountSubAccountApiIpRestrictionV2Resp{}

// CreateSubAccountSubAccountApiIpRestrictionV2Resp struct for CreateSubAccountSubAccountApiIpRestrictionV2Resp
type CreateSubAccountSubAccountApiIpRestrictionV2Resp struct {
	ApiKey *string `json:"apiKey,omitempty"`
	IpList []string `json:"ipList,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewCreateSubAccountSubAccountApiIpRestrictionV2Resp instantiates a new CreateSubAccountSubAccountApiIpRestrictionV2Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubAccountSubAccountApiIpRestrictionV2Resp() *CreateSubAccountSubAccountApiIpRestrictionV2Resp {
	this := CreateSubAccountSubAccountApiIpRestrictionV2Resp{}
	return &this
}

// NewCreateSubAccountSubAccountApiIpRestrictionV2RespWithDefaults instantiates a new CreateSubAccountSubAccountApiIpRestrictionV2Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubAccountSubAccountApiIpRestrictionV2RespWithDefaults() *CreateSubAccountSubAccountApiIpRestrictionV2Resp {
	this := CreateSubAccountSubAccountApiIpRestrictionV2Resp{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetIpList() []string {
	if o == nil || IsNil(o.IpList) {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetIpListOk() ([]string, bool) {
	if o == nil || IsNil(o.IpList) {
		return nil, false
	}
	return o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) HasIpList() bool {
	if o != nil && !IsNil(o.IpList) {
		return true
	}

	return false
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) SetIpList(v []string) {
	o.IpList = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) SetStatus(v string) {
	o.Status = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *CreateSubAccountSubAccountApiIpRestrictionV2Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o CreateSubAccountSubAccountApiIpRestrictionV2Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubAccountSubAccountApiIpRestrictionV2Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.IpList) {
		toSerialize["ipList"] = o.IpList
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp struct {
	value *CreateSubAccountSubAccountApiIpRestrictionV2Resp
	isSet bool
}

func (v NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp) Get() *CreateSubAccountSubAccountApiIpRestrictionV2Resp {
	return v.value
}

func (v *NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp) Set(val *CreateSubAccountSubAccountApiIpRestrictionV2Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubAccountSubAccountApiIpRestrictionV2Resp(val *CreateSubAccountSubAccountApiIpRestrictionV2Resp) *NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp {
	return &NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp{value: val, isSet: true}
}

func (v NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubAccountSubAccountApiIpRestrictionV2Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


