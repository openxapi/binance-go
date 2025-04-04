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

// checks if the WalletGetAssetTransferV1RespRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletGetAssetTransferV1RespRowsInner{}

// WalletGetAssetTransferV1RespRowsInner struct for WalletGetAssetTransferV1RespRowsInner
type WalletGetAssetTransferV1RespRowsInner struct {
	Amount *string `json:"amount,omitempty"`
	Asset *string `json:"asset,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	TranId *int64 `json:"tranId,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewWalletGetAssetTransferV1RespRowsInner instantiates a new WalletGetAssetTransferV1RespRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetAssetTransferV1RespRowsInner() *WalletGetAssetTransferV1RespRowsInner {
	this := WalletGetAssetTransferV1RespRowsInner{}
	return &this
}

// NewWalletGetAssetTransferV1RespRowsInnerWithDefaults instantiates a new WalletGetAssetTransferV1RespRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetAssetTransferV1RespRowsInnerWithDefaults() *WalletGetAssetTransferV1RespRowsInner {
	this := WalletGetAssetTransferV1RespRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1RespRowsInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *WalletGetAssetTransferV1RespRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1RespRowsInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *WalletGetAssetTransferV1RespRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1RespRowsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WalletGetAssetTransferV1RespRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1RespRowsInner) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *WalletGetAssetTransferV1RespRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1RespRowsInner) GetTranId() int64 {
	if o == nil || IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) GetTranIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) HasTranId() bool {
	if o != nil && !IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *WalletGetAssetTransferV1RespRowsInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WalletGetAssetTransferV1RespRowsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WalletGetAssetTransferV1RespRowsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WalletGetAssetTransferV1RespRowsInner) SetType(v string) {
	o.Type = &v
}

func (o WalletGetAssetTransferV1RespRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletGetAssetTransferV1RespRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableWalletGetAssetTransferV1RespRowsInner struct {
	value *WalletGetAssetTransferV1RespRowsInner
	isSet bool
}

func (v NullableWalletGetAssetTransferV1RespRowsInner) Get() *WalletGetAssetTransferV1RespRowsInner {
	return v.value
}

func (v *NullableWalletGetAssetTransferV1RespRowsInner) Set(val *WalletGetAssetTransferV1RespRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetAssetTransferV1RespRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetAssetTransferV1RespRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetAssetTransferV1RespRowsInner(val *WalletGetAssetTransferV1RespRowsInner) *NullableWalletGetAssetTransferV1RespRowsInner {
	return &NullableWalletGetAssetTransferV1RespRowsInner{value: val, isSet: true}
}

func (v NullableWalletGetAssetTransferV1RespRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetAssetTransferV1RespRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


