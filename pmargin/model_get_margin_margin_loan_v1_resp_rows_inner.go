/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the GetMarginMarginLoanV1RespRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarginMarginLoanV1RespRowsInner{}

// GetMarginMarginLoanV1RespRowsInner struct for GetMarginMarginLoanV1RespRowsInner
type GetMarginMarginLoanV1RespRowsInner struct {
	Asset *string `json:"asset,omitempty"`
	Principal *string `json:"principal,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	TxId *int64 `json:"txId,omitempty"`
}

// NewGetMarginMarginLoanV1RespRowsInner instantiates a new GetMarginMarginLoanV1RespRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginMarginLoanV1RespRowsInner() *GetMarginMarginLoanV1RespRowsInner {
	this := GetMarginMarginLoanV1RespRowsInner{}
	return &this
}

// NewGetMarginMarginLoanV1RespRowsInnerWithDefaults instantiates a new GetMarginMarginLoanV1RespRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginMarginLoanV1RespRowsInnerWithDefaults() *GetMarginMarginLoanV1RespRowsInner {
	this := GetMarginMarginLoanV1RespRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetMarginMarginLoanV1RespRowsInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetMarginMarginLoanV1RespRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *GetMarginMarginLoanV1RespRowsInner) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *GetMarginMarginLoanV1RespRowsInner) SetPrincipal(v string) {
	o.Principal = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetMarginMarginLoanV1RespRowsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetMarginMarginLoanV1RespRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetMarginMarginLoanV1RespRowsInner) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetMarginMarginLoanV1RespRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *GetMarginMarginLoanV1RespRowsInner) GetTxId() int64 {
	if o == nil || IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *GetMarginMarginLoanV1RespRowsInner) HasTxId() bool {
	if o != nil && !IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *GetMarginMarginLoanV1RespRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

func (o GetMarginMarginLoanV1RespRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginMarginLoanV1RespRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
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

type NullableGetMarginMarginLoanV1RespRowsInner struct {
	value *GetMarginMarginLoanV1RespRowsInner
	isSet bool
}

func (v NullableGetMarginMarginLoanV1RespRowsInner) Get() *GetMarginMarginLoanV1RespRowsInner {
	return v.value
}

func (v *NullableGetMarginMarginLoanV1RespRowsInner) Set(val *GetMarginMarginLoanV1RespRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginMarginLoanV1RespRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginMarginLoanV1RespRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginMarginLoanV1RespRowsInner(val *GetMarginMarginLoanV1RespRowsInner) *NullableGetMarginMarginLoanV1RespRowsInner {
	return &NullableGetMarginMarginLoanV1RespRowsInner{value: val, isSet: true}
}

func (v NullableGetMarginMarginLoanV1RespRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginMarginLoanV1RespRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


