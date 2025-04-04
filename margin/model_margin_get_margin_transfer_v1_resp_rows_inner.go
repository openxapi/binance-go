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

// checks if the MarginGetMarginTransferV1RespRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginTransferV1RespRowsInner{}

// MarginGetMarginTransferV1RespRowsInner struct for MarginGetMarginTransferV1RespRowsInner
type MarginGetMarginTransferV1RespRowsInner struct {
	Amount *string `json:"amount,omitempty"`
	Asset *string `json:"asset,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	TransFrom *string `json:"transFrom,omitempty"`
	TransTo *string `json:"transTo,omitempty"`
	TxId *int64 `json:"txId,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMarginGetMarginTransferV1RespRowsInner instantiates a new MarginGetMarginTransferV1RespRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginTransferV1RespRowsInner() *MarginGetMarginTransferV1RespRowsInner {
	this := MarginGetMarginTransferV1RespRowsInner{}
	return &this
}

// NewMarginGetMarginTransferV1RespRowsInnerWithDefaults instantiates a new MarginGetMarginTransferV1RespRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginTransferV1RespRowsInnerWithDefaults() *MarginGetMarginTransferV1RespRowsInner {
	this := MarginGetMarginTransferV1RespRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTransFrom returns the TransFrom field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTransFrom() string {
	if o == nil || IsNil(o.TransFrom) {
		var ret string
		return ret
	}
	return *o.TransFrom
}

// GetTransFromOk returns a tuple with the TransFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTransFromOk() (*string, bool) {
	if o == nil || IsNil(o.TransFrom) {
		return nil, false
	}
	return o.TransFrom, true
}

// HasTransFrom returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasTransFrom() bool {
	if o != nil && !IsNil(o.TransFrom) {
		return true
	}

	return false
}

// SetTransFrom gets a reference to the given string and assigns it to the TransFrom field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetTransFrom(v string) {
	o.TransFrom = &v
}

// GetTransTo returns the TransTo field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTransTo() string {
	if o == nil || IsNil(o.TransTo) {
		var ret string
		return ret
	}
	return *o.TransTo
}

// GetTransToOk returns a tuple with the TransTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTransToOk() (*string, bool) {
	if o == nil || IsNil(o.TransTo) {
		return nil, false
	}
	return o.TransTo, true
}

// HasTransTo returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasTransTo() bool {
	if o != nil && !IsNil(o.TransTo) {
		return true
	}

	return false
}

// SetTransTo gets a reference to the given string and assigns it to the TransTo field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetTransTo(v string) {
	o.TransTo = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTxId() int64 {
	if o == nil || IsNil(o.TxId) {
		var ret int64
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTxIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasTxId() bool {
	if o != nil && !IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given int64 and assigns it to the TxId field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetTxId(v int64) {
	o.TxId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginGetMarginTransferV1RespRowsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginGetMarginTransferV1RespRowsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginGetMarginTransferV1RespRowsInner) SetType(v string) {
	o.Type = &v
}

func (o MarginGetMarginTransferV1RespRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginTransferV1RespRowsInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TransFrom) {
		toSerialize["transFrom"] = o.TransFrom
	}
	if !IsNil(o.TransTo) {
		toSerialize["transTo"] = o.TransTo
	}
	if !IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMarginGetMarginTransferV1RespRowsInner struct {
	value *MarginGetMarginTransferV1RespRowsInner
	isSet bool
}

func (v NullableMarginGetMarginTransferV1RespRowsInner) Get() *MarginGetMarginTransferV1RespRowsInner {
	return v.value
}

func (v *NullableMarginGetMarginTransferV1RespRowsInner) Set(val *MarginGetMarginTransferV1RespRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginTransferV1RespRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginTransferV1RespRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginTransferV1RespRowsInner(val *MarginGetMarginTransferV1RespRowsInner) *NullableMarginGetMarginTransferV1RespRowsInner {
	return &NullableMarginGetMarginTransferV1RespRowsInner{value: val, isSet: true}
}

func (v NullableMarginGetMarginTransferV1RespRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginTransferV1RespRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


