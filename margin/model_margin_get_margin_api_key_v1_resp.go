/*
Binance Margin API

OpenAPI specification for Binance cryptocurrency exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginGetMarginApiKeyV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginApiKeyV1Resp{}

// MarginGetMarginApiKeyV1Resp struct for MarginGetMarginApiKeyV1Resp
type MarginGetMarginApiKeyV1Resp struct {
	ApiKey *string `json:"apiKey,omitempty"`
	ApiName *string `json:"apiName,omitempty"`
	Ip *string `json:"ip,omitempty"`
	PermissionMode *string `json:"permissionMode,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMarginGetMarginApiKeyV1Resp instantiates a new MarginGetMarginApiKeyV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginApiKeyV1Resp() *MarginGetMarginApiKeyV1Resp {
	this := MarginGetMarginApiKeyV1Resp{}
	return &this
}

// NewMarginGetMarginApiKeyV1RespWithDefaults instantiates a new MarginGetMarginApiKeyV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginApiKeyV1RespWithDefaults() *MarginGetMarginApiKeyV1Resp {
	this := MarginGetMarginApiKeyV1Resp{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyV1Resp) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyV1Resp) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyV1Resp) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *MarginGetMarginApiKeyV1Resp) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyV1Resp) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyV1Resp) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyV1Resp) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *MarginGetMarginApiKeyV1Resp) SetApiName(v string) {
	o.ApiName = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyV1Resp) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyV1Resp) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyV1Resp) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *MarginGetMarginApiKeyV1Resp) SetIp(v string) {
	o.Ip = &v
}

// GetPermissionMode returns the PermissionMode field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyV1Resp) GetPermissionMode() string {
	if o == nil || IsNil(o.PermissionMode) {
		var ret string
		return ret
	}
	return *o.PermissionMode
}

// GetPermissionModeOk returns a tuple with the PermissionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyV1Resp) GetPermissionModeOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionMode) {
		return nil, false
	}
	return o.PermissionMode, true
}

// HasPermissionMode returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyV1Resp) HasPermissionMode() bool {
	if o != nil && !IsNil(o.PermissionMode) {
		return true
	}

	return false
}

// SetPermissionMode gets a reference to the given string and assigns it to the PermissionMode field.
func (o *MarginGetMarginApiKeyV1Resp) SetPermissionMode(v string) {
	o.PermissionMode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyV1Resp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyV1Resp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyV1Resp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginGetMarginApiKeyV1Resp) SetType(v string) {
	o.Type = &v
}

func (o MarginGetMarginApiKeyV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginApiKeyV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.ApiName) {
		toSerialize["apiName"] = o.ApiName
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.PermissionMode) {
		toSerialize["permissionMode"] = o.PermissionMode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMarginGetMarginApiKeyV1Resp struct {
	value *MarginGetMarginApiKeyV1Resp
	isSet bool
}

func (v NullableMarginGetMarginApiKeyV1Resp) Get() *MarginGetMarginApiKeyV1Resp {
	return v.value
}

func (v *NullableMarginGetMarginApiKeyV1Resp) Set(val *MarginGetMarginApiKeyV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginApiKeyV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginApiKeyV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginApiKeyV1Resp(val *MarginGetMarginApiKeyV1Resp) *NullableMarginGetMarginApiKeyV1Resp {
	return &NullableMarginGetMarginApiKeyV1Resp{value: val, isSet: true}
}

func (v NullableMarginGetMarginApiKeyV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginApiKeyV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


