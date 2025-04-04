/*
Binance Margin Trading API

OpenAPI specification for Binance exchange - Margin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package margin

import (
	"encoding/json"
)

// checks if the MarginGetMarginApiKeyListV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginApiKeyListV1RespItem{}

// MarginGetMarginApiKeyListV1RespItem struct for MarginGetMarginApiKeyListV1RespItem
type MarginGetMarginApiKeyListV1RespItem struct {
	ApiKey *string `json:"apiKey,omitempty"`
	ApiName *string `json:"apiName,omitempty"`
	Ip *string `json:"ip,omitempty"`
	PermissionMode *string `json:"permissionMode,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMarginGetMarginApiKeyListV1RespItem instantiates a new MarginGetMarginApiKeyListV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginApiKeyListV1RespItem() *MarginGetMarginApiKeyListV1RespItem {
	this := MarginGetMarginApiKeyListV1RespItem{}
	return &this
}

// NewMarginGetMarginApiKeyListV1RespItemWithDefaults instantiates a new MarginGetMarginApiKeyListV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginApiKeyListV1RespItemWithDefaults() *MarginGetMarginApiKeyListV1RespItem {
	this := MarginGetMarginApiKeyListV1RespItem{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyListV1RespItem) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *MarginGetMarginApiKeyListV1RespItem) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiName returns the ApiName field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyListV1RespItem) GetApiName() string {
	if o == nil || IsNil(o.ApiName) {
		var ret string
		return ret
	}
	return *o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) GetApiNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApiName) {
		return nil, false
	}
	return o.ApiName, true
}

// HasApiName returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) HasApiName() bool {
	if o != nil && !IsNil(o.ApiName) {
		return true
	}

	return false
}

// SetApiName gets a reference to the given string and assigns it to the ApiName field.
func (o *MarginGetMarginApiKeyListV1RespItem) SetApiName(v string) {
	o.ApiName = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyListV1RespItem) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *MarginGetMarginApiKeyListV1RespItem) SetIp(v string) {
	o.Ip = &v
}

// GetPermissionMode returns the PermissionMode field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyListV1RespItem) GetPermissionMode() string {
	if o == nil || IsNil(o.PermissionMode) {
		var ret string
		return ret
	}
	return *o.PermissionMode
}

// GetPermissionModeOk returns a tuple with the PermissionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) GetPermissionModeOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionMode) {
		return nil, false
	}
	return o.PermissionMode, true
}

// HasPermissionMode returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) HasPermissionMode() bool {
	if o != nil && !IsNil(o.PermissionMode) {
		return true
	}

	return false
}

// SetPermissionMode gets a reference to the given string and assigns it to the PermissionMode field.
func (o *MarginGetMarginApiKeyListV1RespItem) SetPermissionMode(v string) {
	o.PermissionMode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginGetMarginApiKeyListV1RespItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginGetMarginApiKeyListV1RespItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginGetMarginApiKeyListV1RespItem) SetType(v string) {
	o.Type = &v
}

func (o MarginGetMarginApiKeyListV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginApiKeyListV1RespItem) ToMap() (map[string]interface{}, error) {
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

type NullableMarginGetMarginApiKeyListV1RespItem struct {
	value *MarginGetMarginApiKeyListV1RespItem
	isSet bool
}

func (v NullableMarginGetMarginApiKeyListV1RespItem) Get() *MarginGetMarginApiKeyListV1RespItem {
	return v.value
}

func (v *NullableMarginGetMarginApiKeyListV1RespItem) Set(val *MarginGetMarginApiKeyListV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginApiKeyListV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginApiKeyListV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginApiKeyListV1RespItem(val *MarginGetMarginApiKeyListV1RespItem) *NullableMarginGetMarginApiKeyListV1RespItem {
	return &NullableMarginGetMarginApiKeyListV1RespItem{value: val, isSet: true}
}

func (v NullableMarginGetMarginApiKeyListV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginApiKeyListV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


