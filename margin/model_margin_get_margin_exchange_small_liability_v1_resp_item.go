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

// checks if the MarginGetMarginExchangeSmallLiabilityV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginGetMarginExchangeSmallLiabilityV1RespItem{}

// MarginGetMarginExchangeSmallLiabilityV1RespItem struct for MarginGetMarginExchangeSmallLiabilityV1RespItem
type MarginGetMarginExchangeSmallLiabilityV1RespItem struct {
	Asset *string `json:"asset,omitempty"`
	Interest *string `json:"interest,omitempty"`
	LiabilityAsset *string `json:"liabilityAsset,omitempty"`
	LiabilityQty *float32 `json:"liabilityQty,omitempty"`
	Principal *string `json:"principal,omitempty"`
}

// NewMarginGetMarginExchangeSmallLiabilityV1RespItem instantiates a new MarginGetMarginExchangeSmallLiabilityV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginGetMarginExchangeSmallLiabilityV1RespItem() *MarginGetMarginExchangeSmallLiabilityV1RespItem {
	this := MarginGetMarginExchangeSmallLiabilityV1RespItem{}
	return &this
}

// NewMarginGetMarginExchangeSmallLiabilityV1RespItemWithDefaults instantiates a new MarginGetMarginExchangeSmallLiabilityV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginGetMarginExchangeSmallLiabilityV1RespItemWithDefaults() *MarginGetMarginExchangeSmallLiabilityV1RespItem {
	this := MarginGetMarginExchangeSmallLiabilityV1RespItem{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) SetAsset(v string) {
	o.Asset = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetInterest() string {
	if o == nil || IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetInterestOk() (*string, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) SetInterest(v string) {
	o.Interest = &v
}

// GetLiabilityAsset returns the LiabilityAsset field value if set, zero value otherwise.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetLiabilityAsset() string {
	if o == nil || IsNil(o.LiabilityAsset) {
		var ret string
		return ret
	}
	return *o.LiabilityAsset
}

// GetLiabilityAssetOk returns a tuple with the LiabilityAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetLiabilityAssetOk() (*string, bool) {
	if o == nil || IsNil(o.LiabilityAsset) {
		return nil, false
	}
	return o.LiabilityAsset, true
}

// HasLiabilityAsset returns a boolean if a field has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) HasLiabilityAsset() bool {
	if o != nil && !IsNil(o.LiabilityAsset) {
		return true
	}

	return false
}

// SetLiabilityAsset gets a reference to the given string and assigns it to the LiabilityAsset field.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) SetLiabilityAsset(v string) {
	o.LiabilityAsset = &v
}

// GetLiabilityQty returns the LiabilityQty field value if set, zero value otherwise.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetLiabilityQty() float32 {
	if o == nil || IsNil(o.LiabilityQty) {
		var ret float32
		return ret
	}
	return *o.LiabilityQty
}

// GetLiabilityQtyOk returns a tuple with the LiabilityQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetLiabilityQtyOk() (*float32, bool) {
	if o == nil || IsNil(o.LiabilityQty) {
		return nil, false
	}
	return o.LiabilityQty, true
}

// HasLiabilityQty returns a boolean if a field has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) HasLiabilityQty() bool {
	if o != nil && !IsNil(o.LiabilityQty) {
		return true
	}

	return false
}

// SetLiabilityQty gets a reference to the given float32 and assigns it to the LiabilityQty field.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) SetLiabilityQty(v float32) {
	o.LiabilityQty = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *MarginGetMarginExchangeSmallLiabilityV1RespItem) SetPrincipal(v string) {
	o.Principal = &v
}

func (o MarginGetMarginExchangeSmallLiabilityV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginGetMarginExchangeSmallLiabilityV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.LiabilityAsset) {
		toSerialize["liabilityAsset"] = o.LiabilityAsset
	}
	if !IsNil(o.LiabilityQty) {
		toSerialize["liabilityQty"] = o.LiabilityQty
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	return toSerialize, nil
}

type NullableMarginGetMarginExchangeSmallLiabilityV1RespItem struct {
	value *MarginGetMarginExchangeSmallLiabilityV1RespItem
	isSet bool
}

func (v NullableMarginGetMarginExchangeSmallLiabilityV1RespItem) Get() *MarginGetMarginExchangeSmallLiabilityV1RespItem {
	return v.value
}

func (v *NullableMarginGetMarginExchangeSmallLiabilityV1RespItem) Set(val *MarginGetMarginExchangeSmallLiabilityV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginGetMarginExchangeSmallLiabilityV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginGetMarginExchangeSmallLiabilityV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginGetMarginExchangeSmallLiabilityV1RespItem(val *MarginGetMarginExchangeSmallLiabilityV1RespItem) *NullableMarginGetMarginExchangeSmallLiabilityV1RespItem {
	return &NullableMarginGetMarginExchangeSmallLiabilityV1RespItem{value: val, isSet: true}
}

func (v NullableMarginGetMarginExchangeSmallLiabilityV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginGetMarginExchangeSmallLiabilityV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


