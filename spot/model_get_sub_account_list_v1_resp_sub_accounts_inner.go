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

// checks if the GetSubAccountListV1RespSubAccountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubAccountListV1RespSubAccountsInner{}

// GetSubAccountListV1RespSubAccountsInner struct for GetSubAccountListV1RespSubAccountsInner
type GetSubAccountListV1RespSubAccountsInner struct {
	CreateTime *int64 `json:"createTime,omitempty"`
	Email *string `json:"email,omitempty"`
	IsAssetManagementSubAccount *bool `json:"isAssetManagementSubAccount,omitempty"`
	IsFreeze *bool `json:"isFreeze,omitempty"`
	IsManagedSubAccount *bool `json:"isManagedSubAccount,omitempty"`
}

// NewGetSubAccountListV1RespSubAccountsInner instantiates a new GetSubAccountListV1RespSubAccountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountListV1RespSubAccountsInner() *GetSubAccountListV1RespSubAccountsInner {
	this := GetSubAccountListV1RespSubAccountsInner{}
	return &this
}

// NewGetSubAccountListV1RespSubAccountsInnerWithDefaults instantiates a new GetSubAccountListV1RespSubAccountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountListV1RespSubAccountsInnerWithDefaults() *GetSubAccountListV1RespSubAccountsInner {
	this := GetSubAccountListV1RespSubAccountsInner{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetSubAccountListV1RespSubAccountsInner) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetSubAccountListV1RespSubAccountsInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSubAccountListV1RespSubAccountsInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSubAccountListV1RespSubAccountsInner) SetEmail(v string) {
	o.Email = &v
}

// GetIsAssetManagementSubAccount returns the IsAssetManagementSubAccount field value if set, zero value otherwise.
func (o *GetSubAccountListV1RespSubAccountsInner) GetIsAssetManagementSubAccount() bool {
	if o == nil || IsNil(o.IsAssetManagementSubAccount) {
		var ret bool
		return ret
	}
	return *o.IsAssetManagementSubAccount
}

// GetIsAssetManagementSubAccountOk returns a tuple with the IsAssetManagementSubAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) GetIsAssetManagementSubAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAssetManagementSubAccount) {
		return nil, false
	}
	return o.IsAssetManagementSubAccount, true
}

// HasIsAssetManagementSubAccount returns a boolean if a field has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) HasIsAssetManagementSubAccount() bool {
	if o != nil && !IsNil(o.IsAssetManagementSubAccount) {
		return true
	}

	return false
}

// SetIsAssetManagementSubAccount gets a reference to the given bool and assigns it to the IsAssetManagementSubAccount field.
func (o *GetSubAccountListV1RespSubAccountsInner) SetIsAssetManagementSubAccount(v bool) {
	o.IsAssetManagementSubAccount = &v
}

// GetIsFreeze returns the IsFreeze field value if set, zero value otherwise.
func (o *GetSubAccountListV1RespSubAccountsInner) GetIsFreeze() bool {
	if o == nil || IsNil(o.IsFreeze) {
		var ret bool
		return ret
	}
	return *o.IsFreeze
}

// GetIsFreezeOk returns a tuple with the IsFreeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) GetIsFreezeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFreeze) {
		return nil, false
	}
	return o.IsFreeze, true
}

// HasIsFreeze returns a boolean if a field has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) HasIsFreeze() bool {
	if o != nil && !IsNil(o.IsFreeze) {
		return true
	}

	return false
}

// SetIsFreeze gets a reference to the given bool and assigns it to the IsFreeze field.
func (o *GetSubAccountListV1RespSubAccountsInner) SetIsFreeze(v bool) {
	o.IsFreeze = &v
}

// GetIsManagedSubAccount returns the IsManagedSubAccount field value if set, zero value otherwise.
func (o *GetSubAccountListV1RespSubAccountsInner) GetIsManagedSubAccount() bool {
	if o == nil || IsNil(o.IsManagedSubAccount) {
		var ret bool
		return ret
	}
	return *o.IsManagedSubAccount
}

// GetIsManagedSubAccountOk returns a tuple with the IsManagedSubAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) GetIsManagedSubAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManagedSubAccount) {
		return nil, false
	}
	return o.IsManagedSubAccount, true
}

// HasIsManagedSubAccount returns a boolean if a field has been set.
func (o *GetSubAccountListV1RespSubAccountsInner) HasIsManagedSubAccount() bool {
	if o != nil && !IsNil(o.IsManagedSubAccount) {
		return true
	}

	return false
}

// SetIsManagedSubAccount gets a reference to the given bool and assigns it to the IsManagedSubAccount field.
func (o *GetSubAccountListV1RespSubAccountsInner) SetIsManagedSubAccount(v bool) {
	o.IsManagedSubAccount = &v
}

func (o GetSubAccountListV1RespSubAccountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountListV1RespSubAccountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IsAssetManagementSubAccount) {
		toSerialize["isAssetManagementSubAccount"] = o.IsAssetManagementSubAccount
	}
	if !IsNil(o.IsFreeze) {
		toSerialize["isFreeze"] = o.IsFreeze
	}
	if !IsNil(o.IsManagedSubAccount) {
		toSerialize["isManagedSubAccount"] = o.IsManagedSubAccount
	}
	return toSerialize, nil
}

type NullableGetSubAccountListV1RespSubAccountsInner struct {
	value *GetSubAccountListV1RespSubAccountsInner
	isSet bool
}

func (v NullableGetSubAccountListV1RespSubAccountsInner) Get() *GetSubAccountListV1RespSubAccountsInner {
	return v.value
}

func (v *NullableGetSubAccountListV1RespSubAccountsInner) Set(val *GetSubAccountListV1RespSubAccountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountListV1RespSubAccountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountListV1RespSubAccountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountListV1RespSubAccountsInner(val *GetSubAccountListV1RespSubAccountsInner) *NullableGetSubAccountListV1RespSubAccountsInner {
	return &NullableGetSubAccountListV1RespSubAccountsInner{value: val, isSet: true}
}

func (v NullableGetSubAccountListV1RespSubAccountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountListV1RespSubAccountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


