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

// checks if the SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner{}

// SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner struct for SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner
type SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner struct {
	AllocId interface{} `json:"allocId,omitempty"`
	Commission *string `json:"commission,omitempty"`
	CommissionAsset *string `json:"commissionAsset,omitempty"`
	MatchType *string `json:"matchType,omitempty"`
	Price *string `json:"price,omitempty"`
	Qty *string `json:"qty,omitempty"`
	TradeId *int64 `json:"tradeId,omitempty"`
}

// NewSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner instantiates a new SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner() *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner {
	this := SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner{}
	return &this
}

// NewSpotCreateOrderCancelReplaceV3NewOrderRespFillsInnerWithDefaults instantiates a new SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotCreateOrderCancelReplaceV3NewOrderRespFillsInnerWithDefaults() *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner {
	this := SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner{}
	return &this
}

// GetAllocId returns the AllocId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetAllocId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllocId
}

// GetAllocIdOk returns a tuple with the AllocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetAllocIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllocId) {
		return nil, false
	}
	return &o.AllocId, true
}

// HasAllocId returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) HasAllocId() bool {
	if o != nil && !IsNil(o.AllocId) {
		return true
	}

	return false
}

// SetAllocId gets a reference to the given interface{} and assigns it to the AllocId field.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) SetAllocId(v interface{}) {
	o.AllocId = v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetCommission() string {
	if o == nil || IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetCommissionOk() (*string, bool) {
	if o == nil || IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) HasCommission() bool {
	if o != nil && !IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetCommissionAsset() string {
	if o == nil || IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) HasCommissionAsset() bool {
	if o != nil && !IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) SetMatchType(v string) {
	o.MatchType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetQty() string {
	if o == nil || IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetQtyOk() (*string, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) SetQty(v string) {
	o.Qty = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetTradeId() int64 {
	if o == nil || IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) SetTradeId(v int64) {
	o.TradeId = &v
}

func (o SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllocId != nil {
		toSerialize["allocId"] = o.AllocId
	}
	if !IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
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

type NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner struct {
	value *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner
	isSet bool
}

func (v NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) Get() *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner {
	return v.value
}

func (v *NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) Set(val *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner(val *SpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) *NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner {
	return &NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner{value: val, isSet: true}
}

func (v NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotCreateOrderCancelReplaceV3NewOrderRespFillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


