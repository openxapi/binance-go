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

// checks if the WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp{}

// WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp struct for WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp
type WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp struct {
	Rows []WalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespRowsInner `json:"rows,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp instantiates a new WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp() *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp {
	this := WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp{}
	return &this
}

// NewWalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespWithDefaults instantiates a new WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespWithDefaults() *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp {
	this := WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) GetRows() []WalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespRowsInner {
	if o == nil || IsNil(o.Rows) {
		var ret []WalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespRowsInner
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) GetRowsOk() ([]WalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespRowsInner, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []WalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespRowsInner and assigns it to the Rows field.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) SetRows(v []WalletGetAssetLedgerTransferCloudMiningQueryByPageV1RespRowsInner) {
	o.Rows = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp struct {
	value *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp
	isSet bool
}

func (v NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) Get() *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp {
	return v.value
}

func (v *NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) Set(val *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp(val *WalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) *NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp {
	return &NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp{value: val, isSet: true}
}

func (v NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetAssetLedgerTransferCloudMiningQueryByPageV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


