/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotCreateOrderV3RespFillsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotCreateOrderV3RespFillsInner{}

// SpotCreateOrderV3RespFillsInner struct for SpotCreateOrderV3RespFillsInner
type SpotCreateOrderV3RespFillsInner struct {
	Commission *string `json:"commission,omitempty"`
	CommissionAsset *string `json:"commissionAsset,omitempty"`
	Price *string `json:"price,omitempty"`
	Qty *string `json:"qty,omitempty"`
	TradeId *int64 `json:"tradeId,omitempty"`
}

// NewSpotCreateOrderV3RespFillsInner instantiates a new SpotCreateOrderV3RespFillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotCreateOrderV3RespFillsInner() *SpotCreateOrderV3RespFillsInner {
	this := SpotCreateOrderV3RespFillsInner{}
	return &this
}

// NewSpotCreateOrderV3RespFillsInnerWithDefaults instantiates a new SpotCreateOrderV3RespFillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotCreateOrderV3RespFillsInnerWithDefaults() *SpotCreateOrderV3RespFillsInner {
	this := SpotCreateOrderV3RespFillsInner{}
	return &this
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *SpotCreateOrderV3RespFillsInner) GetCommission() string {
	if o == nil || IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderV3RespFillsInner) GetCommissionOk() (*string, bool) {
	if o == nil || IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *SpotCreateOrderV3RespFillsInner) HasCommission() bool {
	if o != nil && !IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *SpotCreateOrderV3RespFillsInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *SpotCreateOrderV3RespFillsInner) GetCommissionAsset() string {
	if o == nil || IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderV3RespFillsInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *SpotCreateOrderV3RespFillsInner) HasCommissionAsset() bool {
	if o != nil && !IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *SpotCreateOrderV3RespFillsInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpotCreateOrderV3RespFillsInner) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderV3RespFillsInner) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpotCreateOrderV3RespFillsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SpotCreateOrderV3RespFillsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *SpotCreateOrderV3RespFillsInner) GetQty() string {
	if o == nil || IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderV3RespFillsInner) GetQtyOk() (*string, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *SpotCreateOrderV3RespFillsInner) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *SpotCreateOrderV3RespFillsInner) SetQty(v string) {
	o.Qty = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *SpotCreateOrderV3RespFillsInner) GetTradeId() int64 {
	if o == nil || IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderV3RespFillsInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *SpotCreateOrderV3RespFillsInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *SpotCreateOrderV3RespFillsInner) SetTradeId(v int64) {
	o.TradeId = &v
}

func (o SpotCreateOrderV3RespFillsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotCreateOrderV3RespFillsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	return toSerialize, nil
}

type NullableSpotCreateOrderV3RespFillsInner struct {
	value *SpotCreateOrderV3RespFillsInner
	isSet bool
}

func (v NullableSpotCreateOrderV3RespFillsInner) Get() *SpotCreateOrderV3RespFillsInner {
	return v.value
}

func (v *NullableSpotCreateOrderV3RespFillsInner) Set(val *SpotCreateOrderV3RespFillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotCreateOrderV3RespFillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotCreateOrderV3RespFillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotCreateOrderV3RespFillsInner(val *SpotCreateOrderV3RespFillsInner) *NullableSpotCreateOrderV3RespFillsInner {
	return &NullableSpotCreateOrderV3RespFillsInner{value: val, isSet: true}
}

func (v NullableSpotCreateOrderV3RespFillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotCreateOrderV3RespFillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


