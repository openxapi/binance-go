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

// checks if the WalletCreateAssetDustV1RespTransferResultInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletCreateAssetDustV1RespTransferResultInner{}

// WalletCreateAssetDustV1RespTransferResultInner struct for WalletCreateAssetDustV1RespTransferResultInner
type WalletCreateAssetDustV1RespTransferResultInner struct {
	Amount *string `json:"amount,omitempty"`
	FromAsset *string `json:"fromAsset,omitempty"`
	OperateTime *int64 `json:"operateTime,omitempty"`
	ServiceChargeAmount *string `json:"serviceChargeAmount,omitempty"`
	TranId *int64 `json:"tranId,omitempty"`
	TransferedAmount *string `json:"transferedAmount,omitempty"`
}

// NewWalletCreateAssetDustV1RespTransferResultInner instantiates a new WalletCreateAssetDustV1RespTransferResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletCreateAssetDustV1RespTransferResultInner() *WalletCreateAssetDustV1RespTransferResultInner {
	this := WalletCreateAssetDustV1RespTransferResultInner{}
	return &this
}

// NewWalletCreateAssetDustV1RespTransferResultInnerWithDefaults instantiates a new WalletCreateAssetDustV1RespTransferResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletCreateAssetDustV1RespTransferResultInnerWithDefaults() *WalletCreateAssetDustV1RespTransferResultInner {
	this := WalletCreateAssetDustV1RespTransferResultInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *WalletCreateAssetDustV1RespTransferResultInner) SetAmount(v string) {
	o.Amount = &v
}

// GetFromAsset returns the FromAsset field value if set, zero value otherwise.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetFromAsset() string {
	if o == nil || IsNil(o.FromAsset) {
		var ret string
		return ret
	}
	return *o.FromAsset
}

// GetFromAssetOk returns a tuple with the FromAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetFromAssetOk() (*string, bool) {
	if o == nil || IsNil(o.FromAsset) {
		return nil, false
	}
	return o.FromAsset, true
}

// HasFromAsset returns a boolean if a field has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) HasFromAsset() bool {
	if o != nil && !IsNil(o.FromAsset) {
		return true
	}

	return false
}

// SetFromAsset gets a reference to the given string and assigns it to the FromAsset field.
func (o *WalletCreateAssetDustV1RespTransferResultInner) SetFromAsset(v string) {
	o.FromAsset = &v
}

// GetOperateTime returns the OperateTime field value if set, zero value otherwise.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetOperateTime() int64 {
	if o == nil || IsNil(o.OperateTime) {
		var ret int64
		return ret
	}
	return *o.OperateTime
}

// GetOperateTimeOk returns a tuple with the OperateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetOperateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.OperateTime) {
		return nil, false
	}
	return o.OperateTime, true
}

// HasOperateTime returns a boolean if a field has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) HasOperateTime() bool {
	if o != nil && !IsNil(o.OperateTime) {
		return true
	}

	return false
}

// SetOperateTime gets a reference to the given int64 and assigns it to the OperateTime field.
func (o *WalletCreateAssetDustV1RespTransferResultInner) SetOperateTime(v int64) {
	o.OperateTime = &v
}

// GetServiceChargeAmount returns the ServiceChargeAmount field value if set, zero value otherwise.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetServiceChargeAmount() string {
	if o == nil || IsNil(o.ServiceChargeAmount) {
		var ret string
		return ret
	}
	return *o.ServiceChargeAmount
}

// GetServiceChargeAmountOk returns a tuple with the ServiceChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetServiceChargeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceChargeAmount) {
		return nil, false
	}
	return o.ServiceChargeAmount, true
}

// HasServiceChargeAmount returns a boolean if a field has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) HasServiceChargeAmount() bool {
	if o != nil && !IsNil(o.ServiceChargeAmount) {
		return true
	}

	return false
}

// SetServiceChargeAmount gets a reference to the given string and assigns it to the ServiceChargeAmount field.
func (o *WalletCreateAssetDustV1RespTransferResultInner) SetServiceChargeAmount(v string) {
	o.ServiceChargeAmount = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetTranId() int64 {
	if o == nil || IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetTranIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) HasTranId() bool {
	if o != nil && !IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *WalletCreateAssetDustV1RespTransferResultInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetTransferedAmount returns the TransferedAmount field value if set, zero value otherwise.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetTransferedAmount() string {
	if o == nil || IsNil(o.TransferedAmount) {
		var ret string
		return ret
	}
	return *o.TransferedAmount
}

// GetTransferedAmountOk returns a tuple with the TransferedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) GetTransferedAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TransferedAmount) {
		return nil, false
	}
	return o.TransferedAmount, true
}

// HasTransferedAmount returns a boolean if a field has been set.
func (o *WalletCreateAssetDustV1RespTransferResultInner) HasTransferedAmount() bool {
	if o != nil && !IsNil(o.TransferedAmount) {
		return true
	}

	return false
}

// SetTransferedAmount gets a reference to the given string and assigns it to the TransferedAmount field.
func (o *WalletCreateAssetDustV1RespTransferResultInner) SetTransferedAmount(v string) {
	o.TransferedAmount = &v
}

func (o WalletCreateAssetDustV1RespTransferResultInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletCreateAssetDustV1RespTransferResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.FromAsset) {
		toSerialize["fromAsset"] = o.FromAsset
	}
	if !IsNil(o.OperateTime) {
		toSerialize["operateTime"] = o.OperateTime
	}
	if !IsNil(o.ServiceChargeAmount) {
		toSerialize["serviceChargeAmount"] = o.ServiceChargeAmount
	}
	if !IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !IsNil(o.TransferedAmount) {
		toSerialize["transferedAmount"] = o.TransferedAmount
	}
	return toSerialize, nil
}

type NullableWalletCreateAssetDustV1RespTransferResultInner struct {
	value *WalletCreateAssetDustV1RespTransferResultInner
	isSet bool
}

func (v NullableWalletCreateAssetDustV1RespTransferResultInner) Get() *WalletCreateAssetDustV1RespTransferResultInner {
	return v.value
}

func (v *NullableWalletCreateAssetDustV1RespTransferResultInner) Set(val *WalletCreateAssetDustV1RespTransferResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletCreateAssetDustV1RespTransferResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletCreateAssetDustV1RespTransferResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletCreateAssetDustV1RespTransferResultInner(val *WalletCreateAssetDustV1RespTransferResultInner) *NullableWalletCreateAssetDustV1RespTransferResultInner {
	return &NullableWalletCreateAssetDustV1RespTransferResultInner{value: val, isSet: true}
}

func (v NullableWalletCreateAssetDustV1RespTransferResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletCreateAssetDustV1RespTransferResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


