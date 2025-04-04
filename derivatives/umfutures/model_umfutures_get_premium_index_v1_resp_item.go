/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the UmfuturesGetPremiumIndexV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesGetPremiumIndexV1RespItem{}

// UmfuturesGetPremiumIndexV1RespItem struct for UmfuturesGetPremiumIndexV1RespItem
type UmfuturesGetPremiumIndexV1RespItem struct {
	EstimatedSettlePrice *string `json:"estimatedSettlePrice,omitempty"`
	IndexPrice *string `json:"indexPrice,omitempty"`
	InterestRate *string `json:"interestRate,omitempty"`
	LastFundingRate *string `json:"lastFundingRate,omitempty"`
	MarkPrice *string `json:"markPrice,omitempty"`
	NextFundingTime *int64 `json:"nextFundingTime,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewUmfuturesGetPremiumIndexV1RespItem instantiates a new UmfuturesGetPremiumIndexV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesGetPremiumIndexV1RespItem() *UmfuturesGetPremiumIndexV1RespItem {
	this := UmfuturesGetPremiumIndexV1RespItem{}
	return &this
}

// NewUmfuturesGetPremiumIndexV1RespItemWithDefaults instantiates a new UmfuturesGetPremiumIndexV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesGetPremiumIndexV1RespItemWithDefaults() *UmfuturesGetPremiumIndexV1RespItem {
	this := UmfuturesGetPremiumIndexV1RespItem{}
	return &this
}

// GetEstimatedSettlePrice returns the EstimatedSettlePrice field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetEstimatedSettlePrice() string {
	if o == nil || IsNil(o.EstimatedSettlePrice) {
		var ret string
		return ret
	}
	return *o.EstimatedSettlePrice
}

// GetEstimatedSettlePriceOk returns a tuple with the EstimatedSettlePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetEstimatedSettlePriceOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedSettlePrice) {
		return nil, false
	}
	return o.EstimatedSettlePrice, true
}

// HasEstimatedSettlePrice returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasEstimatedSettlePrice() bool {
	if o != nil && !IsNil(o.EstimatedSettlePrice) {
		return true
	}

	return false
}

// SetEstimatedSettlePrice gets a reference to the given string and assigns it to the EstimatedSettlePrice field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetEstimatedSettlePrice(v string) {
	o.EstimatedSettlePrice = &v
}

// GetIndexPrice returns the IndexPrice field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetIndexPrice() string {
	if o == nil || IsNil(o.IndexPrice) {
		var ret string
		return ret
	}
	return *o.IndexPrice
}

// GetIndexPriceOk returns a tuple with the IndexPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetIndexPriceOk() (*string, bool) {
	if o == nil || IsNil(o.IndexPrice) {
		return nil, false
	}
	return o.IndexPrice, true
}

// HasIndexPrice returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasIndexPrice() bool {
	if o != nil && !IsNil(o.IndexPrice) {
		return true
	}

	return false
}

// SetIndexPrice gets a reference to the given string and assigns it to the IndexPrice field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetIndexPrice(v string) {
	o.IndexPrice = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetInterestRate() string {
	if o == nil || IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetInterestRateOk() (*string, bool) {
	if o == nil || IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasInterestRate() bool {
	if o != nil && !IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetLastFundingRate returns the LastFundingRate field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetLastFundingRate() string {
	if o == nil || IsNil(o.LastFundingRate) {
		var ret string
		return ret
	}
	return *o.LastFundingRate
}

// GetLastFundingRateOk returns a tuple with the LastFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetLastFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.LastFundingRate) {
		return nil, false
	}
	return o.LastFundingRate, true
}

// HasLastFundingRate returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasLastFundingRate() bool {
	if o != nil && !IsNil(o.LastFundingRate) {
		return true
	}

	return false
}

// SetLastFundingRate gets a reference to the given string and assigns it to the LastFundingRate field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetLastFundingRate(v string) {
	o.LastFundingRate = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetMarkPrice() string {
	if o == nil || IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetMarkPriceOk() (*string, bool) {
	if o == nil || IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasMarkPrice() bool {
	if o != nil && !IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetNextFundingTime returns the NextFundingTime field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetNextFundingTime() int64 {
	if o == nil || IsNil(o.NextFundingTime) {
		var ret int64
		return ret
	}
	return *o.NextFundingTime
}

// GetNextFundingTimeOk returns a tuple with the NextFundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetNextFundingTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.NextFundingTime) {
		return nil, false
	}
	return o.NextFundingTime, true
}

// HasNextFundingTime returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasNextFundingTime() bool {
	if o != nil && !IsNil(o.NextFundingTime) {
		return true
	}

	return false
}

// SetNextFundingTime gets a reference to the given int64 and assigns it to the NextFundingTime field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetNextFundingTime(v int64) {
	o.NextFundingTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UmfuturesGetPremiumIndexV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *UmfuturesGetPremiumIndexV1RespItem) SetTime(v int64) {
	o.Time = &v
}

func (o UmfuturesGetPremiumIndexV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesGetPremiumIndexV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EstimatedSettlePrice) {
		toSerialize["estimatedSettlePrice"] = o.EstimatedSettlePrice
	}
	if !IsNil(o.IndexPrice) {
		toSerialize["indexPrice"] = o.IndexPrice
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !IsNil(o.LastFundingRate) {
		toSerialize["lastFundingRate"] = o.LastFundingRate
	}
	if !IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !IsNil(o.NextFundingTime) {
		toSerialize["nextFundingTime"] = o.NextFundingTime
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableUmfuturesGetPremiumIndexV1RespItem struct {
	value *UmfuturesGetPremiumIndexV1RespItem
	isSet bool
}

func (v NullableUmfuturesGetPremiumIndexV1RespItem) Get() *UmfuturesGetPremiumIndexV1RespItem {
	return v.value
}

func (v *NullableUmfuturesGetPremiumIndexV1RespItem) Set(val *UmfuturesGetPremiumIndexV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetPremiumIndexV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetPremiumIndexV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetPremiumIndexV1RespItem(val *UmfuturesGetPremiumIndexV1RespItem) *NullableUmfuturesGetPremiumIndexV1RespItem {
	return &NullableUmfuturesGetPremiumIndexV1RespItem{value: val, isSet: true}
}

func (v NullableUmfuturesGetPremiumIndexV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetPremiumIndexV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


