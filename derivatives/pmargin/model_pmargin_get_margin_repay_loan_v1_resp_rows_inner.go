/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginGetMarginRepayLoanV1RespRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginGetMarginRepayLoanV1RespRowsInner{}

// PmarginGetMarginRepayLoanV1RespRowsInner struct for PmarginGetMarginRepayLoanV1RespRowsInner
type PmarginGetMarginRepayLoanV1RespRowsInner struct {
	Amount *string `json:"amount,omitempty"`
	Asset *string `json:"asset,omitempty"`
	Interest *string `json:"interest,omitempty"`
	Principal *string `json:"principal,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	TxId *int64 `json:"txId,omitempty"`
}

// NewPmarginGetMarginRepayLoanV1RespRowsInner instantiates a new PmarginGetMarginRepayLoanV1RespRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginGetMarginRepayLoanV1RespRowsInner() *PmarginGetMarginRepayLoanV1RespRowsInner {
	this := PmarginGetMarginRepayLoanV1RespRowsInner{}
	return &this
}

// NewPmarginGetMarginRepayLoanV1RespRowsInnerWithDefaults instantiates a new PmarginGetMarginRepayLoanV1RespRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginGetMarginRepayLoanV1RespRowsInnerWithDefaults() *PmarginGetMarginRepayLoanV1RespRowsInner {
	this := PmarginGetMarginRepayLoanV1RespRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetInterest() string {
	if o == nil || IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetInterestOk() (*string, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) SetInterest(v string) {
	o.Interest = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetTxId() int64 {
	if o == nil || IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) HasTxId() bool {
	if o != nil && !IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *PmarginGetMarginRepayLoanV1RespRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

func (o PmarginGetMarginRepayLoanV1RespRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginGetMarginRepayLoanV1RespRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	return toSerialize, nil
}

type NullablePmarginGetMarginRepayLoanV1RespRowsInner struct {
	value *PmarginGetMarginRepayLoanV1RespRowsInner
	isSet bool
}

func (v NullablePmarginGetMarginRepayLoanV1RespRowsInner) Get() *PmarginGetMarginRepayLoanV1RespRowsInner {
	return v.value
}

func (v *NullablePmarginGetMarginRepayLoanV1RespRowsInner) Set(val *PmarginGetMarginRepayLoanV1RespRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetMarginRepayLoanV1RespRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetMarginRepayLoanV1RespRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetMarginRepayLoanV1RespRowsInner(val *PmarginGetMarginRepayLoanV1RespRowsInner) *NullablePmarginGetMarginRepayLoanV1RespRowsInner {
	return &NullablePmarginGetMarginRepayLoanV1RespRowsInner{value: val, isSet: true}
}

func (v NullablePmarginGetMarginRepayLoanV1RespRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetMarginRepayLoanV1RespRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


