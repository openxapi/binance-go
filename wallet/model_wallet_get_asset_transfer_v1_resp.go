/*
Binance Wallet API

OpenAPI specification for Binance exchange - Wallet API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package wallet

import (
	"encoding/json"
)

// checks if the WalletGetAssetTransferV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetAssetTransferV1Resp{}

// WalletGetAssetTransferV1Resp struct for WalletGetAssetTransferV1Resp
type WalletGetAssetTransferV1Resp struct {
	Rows []WalletGetAssetTransferV1RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewWalletGetAssetTransferV1Resp instantiates a new WalletGetAssetTransferV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetAssetTransferV1Resp() *WalletGetAssetTransferV1Resp {
	this := WalletGetAssetTransferV1Resp{}
	return &this
}

// NewWalletGetAssetTransferV1RespWithDefaults instantiates a new WalletGetAssetTransferV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetAssetTransferV1RespWithDefaults() *WalletGetAssetTransferV1Resp {
	this := WalletGetAssetTransferV1Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1Resp) GetRows() []WalletGetAssetTransferV1RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []WalletGetAssetTransferV1RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1Resp) GetRowsOk() ([]WalletGetAssetTransferV1RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []WalletGetAssetTransferV1RespRowsInner and assigns it to the Rows field.
func (o *WalletGetAssetTransferV1Resp) SetRows(v []WalletGetAssetTransferV1RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *WalletGetAssetTransferV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o WalletGetAssetTransferV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetAssetTransferV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableWalletGetAssetTransferV1Resp struct {
	value *WalletGetAssetTransferV1Resp
	isSet bool
}

func (v NullableWalletGetAssetTransferV1Resp) Get() *WalletGetAssetTransferV1Resp {
	return v.value
}

func (v *NullableWalletGetAssetTransferV1Resp) Set(val *WalletGetAssetTransferV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetAssetTransferV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetAssetTransferV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetAssetTransferV1Resp(val *WalletGetAssetTransferV1Resp) *NullableWalletGetAssetTransferV1Resp {
	return &NullableWalletGetAssetTransferV1Resp{value: val, isSet: true}
}

func (v NullableWalletGetAssetTransferV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetAssetTransferV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


