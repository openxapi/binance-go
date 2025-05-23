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

// checks if the GetPayTransactionsV1RespDataInnerPayerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPayTransactionsV1RespDataInnerPayerInfo{}

// GetPayTransactionsV1RespDataInnerPayerInfo struct for GetPayTransactionsV1RespDataInnerPayerInfo
type GetPayTransactionsV1RespDataInnerPayerInfo struct {
	AccountId *string `json:"accountId,omitempty"`
	BinanceId *string `json:"binanceId,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGetPayTransactionsV1RespDataInnerPayerInfo instantiates a new GetPayTransactionsV1RespDataInnerPayerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayTransactionsV1RespDataInnerPayerInfo() *GetPayTransactionsV1RespDataInnerPayerInfo {
	this := GetPayTransactionsV1RespDataInnerPayerInfo{}
	return &this
}

// NewGetPayTransactionsV1RespDataInnerPayerInfoWithDefaults instantiates a new GetPayTransactionsV1RespDataInnerPayerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayTransactionsV1RespDataInnerPayerInfoWithDefaults() *GetPayTransactionsV1RespDataInnerPayerInfo {
	this := GetPayTransactionsV1RespDataInnerPayerInfo{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBinanceId returns the BinanceId field value if set, zero value otherwise.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetBinanceId() string {
	if o == nil || IsNil(o.BinanceId) {
		var ret string
		return ret
	}
	return *o.BinanceId
}

// GetBinanceIdOk returns a tuple with the BinanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetBinanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.BinanceId) {
		return nil, false
	}
	return o.BinanceId, true
}

// HasBinanceId returns a boolean if a field has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) HasBinanceId() bool {
	if o != nil && !IsNil(o.BinanceId) {
		return true
	}

	return false
}

// SetBinanceId gets a reference to the given string and assigns it to the BinanceId field.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) SetBinanceId(v string) {
	o.BinanceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetPayTransactionsV1RespDataInnerPayerInfo) SetType(v string) {
	o.Type = &v
}

func (o GetPayTransactionsV1RespDataInnerPayerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayTransactionsV1RespDataInnerPayerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.BinanceId) {
		toSerialize["binanceId"] = o.BinanceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetPayTransactionsV1RespDataInnerPayerInfo struct {
	value *GetPayTransactionsV1RespDataInnerPayerInfo
	isSet bool
}

func (v NullableGetPayTransactionsV1RespDataInnerPayerInfo) Get() *GetPayTransactionsV1RespDataInnerPayerInfo {
	return v.value
}

func (v *NullableGetPayTransactionsV1RespDataInnerPayerInfo) Set(val *GetPayTransactionsV1RespDataInnerPayerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayTransactionsV1RespDataInnerPayerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayTransactionsV1RespDataInnerPayerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayTransactionsV1RespDataInnerPayerInfo(val *GetPayTransactionsV1RespDataInnerPayerInfo) *NullableGetPayTransactionsV1RespDataInnerPayerInfo {
	return &NullableGetPayTransactionsV1RespDataInnerPayerInfo{value: val, isSet: true}
}

func (v NullableGetPayTransactionsV1RespDataInnerPayerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayTransactionsV1RespDataInnerPayerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


