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

// checks if the GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner{}

// GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner struct for GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner
type GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner struct {
	EntryPrice *int32 `json:"entryPrice,omitempty"`
	MarkPrice *int32 `json:"markPrice,omitempty"`
	PositionAmt *float32 `json:"positionAmt,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner instantiates a new GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner() *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner {
	this := GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner{}
	return &this
}

// NewGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInnerWithDefaults instantiates a new GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInnerWithDefaults() *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner {
	this := GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner{}
	return &this
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetEntryPrice() int32 {
	if o == nil || IsNil(o.EntryPrice) {
		var ret int32
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetEntryPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) HasEntryPrice() bool {
	if o != nil && !IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given int32 and assigns it to the EntryPrice field.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) SetEntryPrice(v int32) {
	o.EntryPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetMarkPrice() int32 {
	if o == nil || IsNil(o.MarkPrice) {
		var ret int32
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetMarkPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) HasMarkPrice() bool {
	if o != nil && !IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given int32 and assigns it to the MarkPrice field.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) SetMarkPrice(v int32) {
	o.MarkPrice = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetPositionAmt() float32 {
	if o == nil || IsNil(o.PositionAmt) {
		var ret float32
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetPositionAmtOk() (*float32, bool) {
	if o == nil || IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) HasPositionAmt() bool {
	if o != nil && !IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given float32 and assigns it to the PositionAmt field.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) SetPositionAmt(v float32) {
	o.PositionAmt = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) SetSymbol(v string) {
	o.Symbol = &v
}

func (o GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner struct {
	value *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner
	isSet bool
}

func (v NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) Get() *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner {
	return v.value
}

func (v *NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) Set(val *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner(val *GetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) *NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner {
	return &NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner{value: val, isSet: true}
}

func (v NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagedSubaccountFetchFutureAssetV1RespSnapshotVosInnerDataPositionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


