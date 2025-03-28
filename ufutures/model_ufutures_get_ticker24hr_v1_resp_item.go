/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetTicker24hrV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetTicker24hrV1RespItem{}

// UfuturesGetTicker24hrV1RespItem struct for UfuturesGetTicker24hrV1RespItem
type UfuturesGetTicker24hrV1RespItem struct {
	CloseTime *int64 `json:"closeTime,omitempty"`
	Count *int32 `json:"count,omitempty"`
	FirstId *int32 `json:"firstId,omitempty"`
	HighPrice *string `json:"highPrice,omitempty"`
	LastId *int32 `json:"lastId,omitempty"`
	LastPrice *string `json:"lastPrice,omitempty"`
	LastQty *string `json:"lastQty,omitempty"`
	LowPrice *string `json:"lowPrice,omitempty"`
	OpenPrice *string `json:"openPrice,omitempty"`
	OpenTime *int64 `json:"openTime,omitempty"`
	PriceChange *string `json:"priceChange,omitempty"`
	PriceChangePercent *string `json:"priceChangePercent,omitempty"`
	QuoteVolume *string `json:"quoteVolume,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Volume *string `json:"volume,omitempty"`
	WeightedAvgPrice *string `json:"weightedAvgPrice,omitempty"`
}

// NewUfuturesGetTicker24hrV1RespItem instantiates a new UfuturesGetTicker24hrV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetTicker24hrV1RespItem() *UfuturesGetTicker24hrV1RespItem {
	this := UfuturesGetTicker24hrV1RespItem{}
	return &this
}

// NewUfuturesGetTicker24hrV1RespItemWithDefaults instantiates a new UfuturesGetTicker24hrV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetTicker24hrV1RespItemWithDefaults() *UfuturesGetTicker24hrV1RespItem {
	this := UfuturesGetTicker24hrV1RespItem{}
	return &this
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetCloseTime() int64 {
	if o == nil || IsNil(o.CloseTime) {
		var ret int64
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetCloseTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasCloseTime() bool {
	if o != nil && !IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given int64 and assigns it to the CloseTime field.
func (o *UfuturesGetTicker24hrV1RespItem) SetCloseTime(v int64) {
	o.CloseTime = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UfuturesGetTicker24hrV1RespItem) SetCount(v int32) {
	o.Count = &v
}

// GetFirstId returns the FirstId field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetFirstId() int32 {
	if o == nil || IsNil(o.FirstId) {
		var ret int32
		return ret
	}
	return *o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetFirstIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstId) {
		return nil, false
	}
	return o.FirstId, true
}

// HasFirstId returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasFirstId() bool {
	if o != nil && !IsNil(o.FirstId) {
		return true
	}

	return false
}

// SetFirstId gets a reference to the given int32 and assigns it to the FirstId field.
func (o *UfuturesGetTicker24hrV1RespItem) SetFirstId(v int32) {
	o.FirstId = &v
}

// GetHighPrice returns the HighPrice field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetHighPrice() string {
	if o == nil || IsNil(o.HighPrice) {
		var ret string
		return ret
	}
	return *o.HighPrice
}

// GetHighPriceOk returns a tuple with the HighPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetHighPriceOk() (*string, bool) {
	if o == nil || IsNil(o.HighPrice) {
		return nil, false
	}
	return o.HighPrice, true
}

// HasHighPrice returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasHighPrice() bool {
	if o != nil && !IsNil(o.HighPrice) {
		return true
	}

	return false
}

// SetHighPrice gets a reference to the given string and assigns it to the HighPrice field.
func (o *UfuturesGetTicker24hrV1RespItem) SetHighPrice(v string) {
	o.HighPrice = &v
}

// GetLastId returns the LastId field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetLastId() int32 {
	if o == nil || IsNil(o.LastId) {
		var ret int32
		return ret
	}
	return *o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetLastIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LastId) {
		return nil, false
	}
	return o.LastId, true
}

// HasLastId returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasLastId() bool {
	if o != nil && !IsNil(o.LastId) {
		return true
	}

	return false
}

// SetLastId gets a reference to the given int32 and assigns it to the LastId field.
func (o *UfuturesGetTicker24hrV1RespItem) SetLastId(v int32) {
	o.LastId = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetLastPrice() string {
	if o == nil || IsNil(o.LastPrice) {
		var ret string
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetLastPriceOk() (*string, bool) {
	if o == nil || IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasLastPrice() bool {
	if o != nil && !IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given string and assigns it to the LastPrice field.
func (o *UfuturesGetTicker24hrV1RespItem) SetLastPrice(v string) {
	o.LastPrice = &v
}

// GetLastQty returns the LastQty field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetLastQty() string {
	if o == nil || IsNil(o.LastQty) {
		var ret string
		return ret
	}
	return *o.LastQty
}

// GetLastQtyOk returns a tuple with the LastQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetLastQtyOk() (*string, bool) {
	if o == nil || IsNil(o.LastQty) {
		return nil, false
	}
	return o.LastQty, true
}

// HasLastQty returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasLastQty() bool {
	if o != nil && !IsNil(o.LastQty) {
		return true
	}

	return false
}

// SetLastQty gets a reference to the given string and assigns it to the LastQty field.
func (o *UfuturesGetTicker24hrV1RespItem) SetLastQty(v string) {
	o.LastQty = &v
}

// GetLowPrice returns the LowPrice field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetLowPrice() string {
	if o == nil || IsNil(o.LowPrice) {
		var ret string
		return ret
	}
	return *o.LowPrice
}

// GetLowPriceOk returns a tuple with the LowPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetLowPriceOk() (*string, bool) {
	if o == nil || IsNil(o.LowPrice) {
		return nil, false
	}
	return o.LowPrice, true
}

// HasLowPrice returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasLowPrice() bool {
	if o != nil && !IsNil(o.LowPrice) {
		return true
	}

	return false
}

// SetLowPrice gets a reference to the given string and assigns it to the LowPrice field.
func (o *UfuturesGetTicker24hrV1RespItem) SetLowPrice(v string) {
	o.LowPrice = &v
}

// GetOpenPrice returns the OpenPrice field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetOpenPrice() string {
	if o == nil || IsNil(o.OpenPrice) {
		var ret string
		return ret
	}
	return *o.OpenPrice
}

// GetOpenPriceOk returns a tuple with the OpenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetOpenPriceOk() (*string, bool) {
	if o == nil || IsNil(o.OpenPrice) {
		return nil, false
	}
	return o.OpenPrice, true
}

// HasOpenPrice returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasOpenPrice() bool {
	if o != nil && !IsNil(o.OpenPrice) {
		return true
	}

	return false
}

// SetOpenPrice gets a reference to the given string and assigns it to the OpenPrice field.
func (o *UfuturesGetTicker24hrV1RespItem) SetOpenPrice(v string) {
	o.OpenPrice = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetOpenTime() int64 {
	if o == nil || IsNil(o.OpenTime) {
		var ret int64
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetOpenTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasOpenTime() bool {
	if o != nil && !IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given int64 and assigns it to the OpenTime field.
func (o *UfuturesGetTicker24hrV1RespItem) SetOpenTime(v int64) {
	o.OpenTime = &v
}

// GetPriceChange returns the PriceChange field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetPriceChange() string {
	if o == nil || IsNil(o.PriceChange) {
		var ret string
		return ret
	}
	return *o.PriceChange
}

// GetPriceChangeOk returns a tuple with the PriceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetPriceChangeOk() (*string, bool) {
	if o == nil || IsNil(o.PriceChange) {
		return nil, false
	}
	return o.PriceChange, true
}

// HasPriceChange returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasPriceChange() bool {
	if o != nil && !IsNil(o.PriceChange) {
		return true
	}

	return false
}

// SetPriceChange gets a reference to the given string and assigns it to the PriceChange field.
func (o *UfuturesGetTicker24hrV1RespItem) SetPriceChange(v string) {
	o.PriceChange = &v
}

// GetPriceChangePercent returns the PriceChangePercent field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetPriceChangePercent() string {
	if o == nil || IsNil(o.PriceChangePercent) {
		var ret string
		return ret
	}
	return *o.PriceChangePercent
}

// GetPriceChangePercentOk returns a tuple with the PriceChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetPriceChangePercentOk() (*string, bool) {
	if o == nil || IsNil(o.PriceChangePercent) {
		return nil, false
	}
	return o.PriceChangePercent, true
}

// HasPriceChangePercent returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasPriceChangePercent() bool {
	if o != nil && !IsNil(o.PriceChangePercent) {
		return true
	}

	return false
}

// SetPriceChangePercent gets a reference to the given string and assigns it to the PriceChangePercent field.
func (o *UfuturesGetTicker24hrV1RespItem) SetPriceChangePercent(v string) {
	o.PriceChangePercent = &v
}

// GetQuoteVolume returns the QuoteVolume field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetQuoteVolume() string {
	if o == nil || IsNil(o.QuoteVolume) {
		var ret string
		return ret
	}
	return *o.QuoteVolume
}

// GetQuoteVolumeOk returns a tuple with the QuoteVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetQuoteVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteVolume) {
		return nil, false
	}
	return o.QuoteVolume, true
}

// HasQuoteVolume returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasQuoteVolume() bool {
	if o != nil && !IsNil(o.QuoteVolume) {
		return true
	}

	return false
}

// SetQuoteVolume gets a reference to the given string and assigns it to the QuoteVolume field.
func (o *UfuturesGetTicker24hrV1RespItem) SetQuoteVolume(v string) {
	o.QuoteVolume = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UfuturesGetTicker24hrV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *UfuturesGetTicker24hrV1RespItem) SetVolume(v string) {
	o.Volume = &v
}

// GetWeightedAvgPrice returns the WeightedAvgPrice field value if set, zero value otherwise.
func (o *UfuturesGetTicker24hrV1RespItem) GetWeightedAvgPrice() string {
	if o == nil || IsNil(o.WeightedAvgPrice) {
		var ret string
		return ret
	}
	return *o.WeightedAvgPrice
}

// GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetTicker24hrV1RespItem) GetWeightedAvgPriceOk() (*string, bool) {
	if o == nil || IsNil(o.WeightedAvgPrice) {
		return nil, false
	}
	return o.WeightedAvgPrice, true
}

// HasWeightedAvgPrice returns a boolean if a field has been set.
func (o *UfuturesGetTicker24hrV1RespItem) HasWeightedAvgPrice() bool {
	if o != nil && !IsNil(o.WeightedAvgPrice) {
		return true
	}

	return false
}

// SetWeightedAvgPrice gets a reference to the given string and assigns it to the WeightedAvgPrice field.
func (o *UfuturesGetTicker24hrV1RespItem) SetWeightedAvgPrice(v string) {
	o.WeightedAvgPrice = &v
}

func (o UfuturesGetTicker24hrV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetTicker24hrV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.FirstId) {
		toSerialize["firstId"] = o.FirstId
	}
	if !IsNil(o.HighPrice) {
		toSerialize["highPrice"] = o.HighPrice
	}
	if !IsNil(o.LastId) {
		toSerialize["lastId"] = o.LastId
	}
	if !IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !IsNil(o.LastQty) {
		toSerialize["lastQty"] = o.LastQty
	}
	if !IsNil(o.LowPrice) {
		toSerialize["lowPrice"] = o.LowPrice
	}
	if !IsNil(o.OpenPrice) {
		toSerialize["openPrice"] = o.OpenPrice
	}
	if !IsNil(o.OpenTime) {
		toSerialize["openTime"] = o.OpenTime
	}
	if !IsNil(o.PriceChange) {
		toSerialize["priceChange"] = o.PriceChange
	}
	if !IsNil(o.PriceChangePercent) {
		toSerialize["priceChangePercent"] = o.PriceChangePercent
	}
	if !IsNil(o.QuoteVolume) {
		toSerialize["quoteVolume"] = o.QuoteVolume
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.WeightedAvgPrice) {
		toSerialize["weightedAvgPrice"] = o.WeightedAvgPrice
	}
	return toSerialize, nil
}

type NullableUfuturesGetTicker24hrV1RespItem struct {
	value *UfuturesGetTicker24hrV1RespItem
	isSet bool
}

func (v NullableUfuturesGetTicker24hrV1RespItem) Get() *UfuturesGetTicker24hrV1RespItem {
	return v.value
}

func (v *NullableUfuturesGetTicker24hrV1RespItem) Set(val *UfuturesGetTicker24hrV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetTicker24hrV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetTicker24hrV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetTicker24hrV1RespItem(val *UfuturesGetTicker24hrV1RespItem) *NullableUfuturesGetTicker24hrV1RespItem {
	return &NullableUfuturesGetTicker24hrV1RespItem{value: val, isSet: true}
}

func (v NullableUfuturesGetTicker24hrV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetTicker24hrV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


