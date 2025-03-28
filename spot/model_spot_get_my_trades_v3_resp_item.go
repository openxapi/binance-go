/*
Binance Spot API

OpenAPI specification for Binance cryptocurrency exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotGetMyTradesV3RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotGetMyTradesV3RespItem{}

// SpotGetMyTradesV3RespItem struct for SpotGetMyTradesV3RespItem
type SpotGetMyTradesV3RespItem struct {
	Commission *string `json:"commission,omitempty"`
	CommissionAsset *string `json:"commissionAsset,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsBestMatch *bool `json:"isBestMatch,omitempty"`
	IsBuyer *bool `json:"isBuyer,omitempty"`
	IsMaker *bool `json:"isMaker,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrderListId *int64 `json:"orderListId,omitempty"`
	Price *string `json:"price,omitempty"`
	Qty *string `json:"qty,omitempty"`
	QuoteQty *string `json:"quoteQty,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewSpotGetMyTradesV3RespItem instantiates a new SpotGetMyTradesV3RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotGetMyTradesV3RespItem() *SpotGetMyTradesV3RespItem {
	this := SpotGetMyTradesV3RespItem{}
	return &this
}

// NewSpotGetMyTradesV3RespItemWithDefaults instantiates a new SpotGetMyTradesV3RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotGetMyTradesV3RespItemWithDefaults() *SpotGetMyTradesV3RespItem {
	this := SpotGetMyTradesV3RespItem{}
	return &this
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetCommission() string {
	if o == nil || IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetCommissionOk() (*string, bool) {
	if o == nil || IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasCommission() bool {
	if o != nil && !IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *SpotGetMyTradesV3RespItem) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetCommissionAsset() string {
	if o == nil || IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetCommissionAssetOk() (*string, bool) {
	if o == nil || IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasCommissionAsset() bool {
	if o != nil && !IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *SpotGetMyTradesV3RespItem) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SpotGetMyTradesV3RespItem) SetId(v int32) {
	o.Id = &v
}

// GetIsBestMatch returns the IsBestMatch field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetIsBestMatch() bool {
	if o == nil || IsNil(o.IsBestMatch) {
		var ret bool
		return ret
	}
	return *o.IsBestMatch
}

// GetIsBestMatchOk returns a tuple with the IsBestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetIsBestMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBestMatch) {
		return nil, false
	}
	return o.IsBestMatch, true
}

// HasIsBestMatch returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasIsBestMatch() bool {
	if o != nil && !IsNil(o.IsBestMatch) {
		return true
	}

	return false
}

// SetIsBestMatch gets a reference to the given bool and assigns it to the IsBestMatch field.
func (o *SpotGetMyTradesV3RespItem) SetIsBestMatch(v bool) {
	o.IsBestMatch = &v
}

// GetIsBuyer returns the IsBuyer field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetIsBuyer() bool {
	if o == nil || IsNil(o.IsBuyer) {
		var ret bool
		return ret
	}
	return *o.IsBuyer
}

// GetIsBuyerOk returns a tuple with the IsBuyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetIsBuyerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBuyer) {
		return nil, false
	}
	return o.IsBuyer, true
}

// HasIsBuyer returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasIsBuyer() bool {
	if o != nil && !IsNil(o.IsBuyer) {
		return true
	}

	return false
}

// SetIsBuyer gets a reference to the given bool and assigns it to the IsBuyer field.
func (o *SpotGetMyTradesV3RespItem) SetIsBuyer(v bool) {
	o.IsBuyer = &v
}

// GetIsMaker returns the IsMaker field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetIsMaker() bool {
	if o == nil || IsNil(o.IsMaker) {
		var ret bool
		return ret
	}
	return *o.IsMaker
}

// GetIsMakerOk returns a tuple with the IsMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetIsMakerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMaker) {
		return nil, false
	}
	return o.IsMaker, true
}

// HasIsMaker returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasIsMaker() bool {
	if o != nil && !IsNil(o.IsMaker) {
		return true
	}

	return false
}

// SetIsMaker gets a reference to the given bool and assigns it to the IsMaker field.
func (o *SpotGetMyTradesV3RespItem) SetIsMaker(v bool) {
	o.IsMaker = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *SpotGetMyTradesV3RespItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrderListId returns the OrderListId field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetOrderListId() int64 {
	if o == nil || IsNil(o.OrderListId) {
		var ret int64
		return ret
	}
	return *o.OrderListId
}

// GetOrderListIdOk returns a tuple with the OrderListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetOrderListIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderListId) {
		return nil, false
	}
	return o.OrderListId, true
}

// HasOrderListId returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasOrderListId() bool {
	if o != nil && !IsNil(o.OrderListId) {
		return true
	}

	return false
}

// SetOrderListId gets a reference to the given int64 and assigns it to the OrderListId field.
func (o *SpotGetMyTradesV3RespItem) SetOrderListId(v int64) {
	o.OrderListId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *SpotGetMyTradesV3RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetQty() string {
	if o == nil || IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetQtyOk() (*string, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *SpotGetMyTradesV3RespItem) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetQuoteQty() string {
	if o == nil || IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasQuoteQty() bool {
	if o != nil && !IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *SpotGetMyTradesV3RespItem) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *SpotGetMyTradesV3RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SpotGetMyTradesV3RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetMyTradesV3RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SpotGetMyTradesV3RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *SpotGetMyTradesV3RespItem) SetTime(v int64) {
	o.Time = &v
}

func (o SpotGetMyTradesV3RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotGetMyTradesV3RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsBestMatch) {
		toSerialize["isBestMatch"] = o.IsBestMatch
	}
	if !IsNil(o.IsBuyer) {
		toSerialize["isBuyer"] = o.IsBuyer
	}
	if !IsNil(o.IsMaker) {
		toSerialize["isMaker"] = o.IsMaker
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrderListId) {
		toSerialize["orderListId"] = o.OrderListId
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !IsNil(o.QuoteQty) {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableSpotGetMyTradesV3RespItem struct {
	value *SpotGetMyTradesV3RespItem
	isSet bool
}

func (v NullableSpotGetMyTradesV3RespItem) Get() *SpotGetMyTradesV3RespItem {
	return v.value
}

func (v *NullableSpotGetMyTradesV3RespItem) Set(val *SpotGetMyTradesV3RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotGetMyTradesV3RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotGetMyTradesV3RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotGetMyTradesV3RespItem(val *SpotGetMyTradesV3RespItem) *NullableSpotGetMyTradesV3RespItem {
	return &NullableSpotGetMyTradesV3RespItem{value: val, isSet: true}
}

func (v NullableSpotGetMyTradesV3RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotGetMyTradesV3RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


