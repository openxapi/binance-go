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

// checks if the GetMarginCrossMarginDataV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarginCrossMarginDataV1RespItem{}

// GetMarginCrossMarginDataV1RespItem struct for GetMarginCrossMarginDataV1RespItem
type GetMarginCrossMarginDataV1RespItem struct {
	BorrowLimit *string `json:"borrowLimit,omitempty"`
	Borrowable *bool `json:"borrowable,omitempty"`
	Coin *string `json:"coin,omitempty"`
	DailyInterest *string `json:"dailyInterest,omitempty"`
	MarginablePairs []string `json:"marginablePairs,omitempty"`
	TransferIn *bool `json:"transferIn,omitempty"`
	VipLevel *int32 `json:"vipLevel,omitempty"`
	YearlyInterest *string `json:"yearlyInterest,omitempty"`
}

// NewGetMarginCrossMarginDataV1RespItem instantiates a new GetMarginCrossMarginDataV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginCrossMarginDataV1RespItem() *GetMarginCrossMarginDataV1RespItem {
	this := GetMarginCrossMarginDataV1RespItem{}
	return &this
}

// NewGetMarginCrossMarginDataV1RespItemWithDefaults instantiates a new GetMarginCrossMarginDataV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginCrossMarginDataV1RespItemWithDefaults() *GetMarginCrossMarginDataV1RespItem {
	this := GetMarginCrossMarginDataV1RespItem{}
	return &this
}

// GetBorrowLimit returns the BorrowLimit field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowLimit() string {
	if o == nil || IsNil(o.BorrowLimit) {
		var ret string
		return ret
	}
	return *o.BorrowLimit
}

// GetBorrowLimitOk returns a tuple with the BorrowLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowLimitOk() (*string, bool) {
	if o == nil || IsNil(o.BorrowLimit) {
		return nil, false
	}
	return o.BorrowLimit, true
}

// HasBorrowLimit returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasBorrowLimit() bool {
	if o != nil && !IsNil(o.BorrowLimit) {
		return true
	}

	return false
}

// SetBorrowLimit gets a reference to the given string and assigns it to the BorrowLimit field.
func (o *GetMarginCrossMarginDataV1RespItem) SetBorrowLimit(v string) {
	o.BorrowLimit = &v
}

// GetBorrowable returns the Borrowable field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowable() bool {
	if o == nil || IsNil(o.Borrowable) {
		var ret bool
		return ret
	}
	return *o.Borrowable
}

// GetBorrowableOk returns a tuple with the Borrowable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetBorrowableOk() (*bool, bool) {
	if o == nil || IsNil(o.Borrowable) {
		return nil, false
	}
	return o.Borrowable, true
}

// HasBorrowable returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasBorrowable() bool {
	if o != nil && !IsNil(o.Borrowable) {
		return true
	}

	return false
}

// SetBorrowable gets a reference to the given bool and assigns it to the Borrowable field.
func (o *GetMarginCrossMarginDataV1RespItem) SetBorrowable(v bool) {
	o.Borrowable = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetCoin() string {
	if o == nil || IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetCoinOk() (*string, bool) {
	if o == nil || IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasCoin() bool {
	if o != nil && !IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetMarginCrossMarginDataV1RespItem) SetCoin(v string) {
	o.Coin = &v
}

// GetDailyInterest returns the DailyInterest field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetDailyInterest() string {
	if o == nil || IsNil(o.DailyInterest) {
		var ret string
		return ret
	}
	return *o.DailyInterest
}

// GetDailyInterestOk returns a tuple with the DailyInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetDailyInterestOk() (*string, bool) {
	if o == nil || IsNil(o.DailyInterest) {
		return nil, false
	}
	return o.DailyInterest, true
}

// HasDailyInterest returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasDailyInterest() bool {
	if o != nil && !IsNil(o.DailyInterest) {
		return true
	}

	return false
}

// SetDailyInterest gets a reference to the given string and assigns it to the DailyInterest field.
func (o *GetMarginCrossMarginDataV1RespItem) SetDailyInterest(v string) {
	o.DailyInterest = &v
}

// GetMarginablePairs returns the MarginablePairs field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetMarginablePairs() []string {
	if o == nil || IsNil(o.MarginablePairs) {
		var ret []string
		return ret
	}
	return o.MarginablePairs
}

// GetMarginablePairsOk returns a tuple with the MarginablePairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetMarginablePairsOk() ([]string, bool) {
	if o == nil || IsNil(o.MarginablePairs) {
		return nil, false
	}
	return o.MarginablePairs, true
}

// HasMarginablePairs returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasMarginablePairs() bool {
	if o != nil && !IsNil(o.MarginablePairs) {
		return true
	}

	return false
}

// SetMarginablePairs gets a reference to the given []string and assigns it to the MarginablePairs field.
func (o *GetMarginCrossMarginDataV1RespItem) SetMarginablePairs(v []string) {
	o.MarginablePairs = v
}

// GetTransferIn returns the TransferIn field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetTransferIn() bool {
	if o == nil || IsNil(o.TransferIn) {
		var ret bool
		return ret
	}
	return *o.TransferIn
}

// GetTransferInOk returns a tuple with the TransferIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetTransferInOk() (*bool, bool) {
	if o == nil || IsNil(o.TransferIn) {
		return nil, false
	}
	return o.TransferIn, true
}

// HasTransferIn returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasTransferIn() bool {
	if o != nil && !IsNil(o.TransferIn) {
		return true
	}

	return false
}

// SetTransferIn gets a reference to the given bool and assigns it to the TransferIn field.
func (o *GetMarginCrossMarginDataV1RespItem) SetTransferIn(v bool) {
	o.TransferIn = &v
}

// GetVipLevel returns the VipLevel field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetVipLevel() int32 {
	if o == nil || IsNil(o.VipLevel) {
		var ret int32
		return ret
	}
	return *o.VipLevel
}

// GetVipLevelOk returns a tuple with the VipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetVipLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.VipLevel) {
		return nil, false
	}
	return o.VipLevel, true
}

// HasVipLevel returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasVipLevel() bool {
	if o != nil && !IsNil(o.VipLevel) {
		return true
	}

	return false
}

// SetVipLevel gets a reference to the given int32 and assigns it to the VipLevel field.
func (o *GetMarginCrossMarginDataV1RespItem) SetVipLevel(v int32) {
	o.VipLevel = &v
}

// GetYearlyInterest returns the YearlyInterest field value if set, zero value otherwise.
func (o *GetMarginCrossMarginDataV1RespItem) GetYearlyInterest() string {
	if o == nil || IsNil(o.YearlyInterest) {
		var ret string
		return ret
	}
	return *o.YearlyInterest
}

// GetYearlyInterestOk returns a tuple with the YearlyInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginCrossMarginDataV1RespItem) GetYearlyInterestOk() (*string, bool) {
	if o == nil || IsNil(o.YearlyInterest) {
		return nil, false
	}
	return o.YearlyInterest, true
}

// HasYearlyInterest returns a boolean if a field has been set.
func (o *GetMarginCrossMarginDataV1RespItem) HasYearlyInterest() bool {
	if o != nil && !IsNil(o.YearlyInterest) {
		return true
	}

	return false
}

// SetYearlyInterest gets a reference to the given string and assigns it to the YearlyInterest field.
func (o *GetMarginCrossMarginDataV1RespItem) SetYearlyInterest(v string) {
	o.YearlyInterest = &v
}

func (o GetMarginCrossMarginDataV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginCrossMarginDataV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BorrowLimit) {
		toSerialize["borrowLimit"] = o.BorrowLimit
	}
	if !IsNil(o.Borrowable) {
		toSerialize["borrowable"] = o.Borrowable
	}
	if !IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !IsNil(o.DailyInterest) {
		toSerialize["dailyInterest"] = o.DailyInterest
	}
	if !IsNil(o.MarginablePairs) {
		toSerialize["marginablePairs"] = o.MarginablePairs
	}
	if !IsNil(o.TransferIn) {
		toSerialize["transferIn"] = o.TransferIn
	}
	if !IsNil(o.VipLevel) {
		toSerialize["vipLevel"] = o.VipLevel
	}
	if !IsNil(o.YearlyInterest) {
		toSerialize["yearlyInterest"] = o.YearlyInterest
	}
	return toSerialize, nil
}

type NullableGetMarginCrossMarginDataV1RespItem struct {
	value *GetMarginCrossMarginDataV1RespItem
	isSet bool
}

func (v NullableGetMarginCrossMarginDataV1RespItem) Get() *GetMarginCrossMarginDataV1RespItem {
	return v.value
}

func (v *NullableGetMarginCrossMarginDataV1RespItem) Set(val *GetMarginCrossMarginDataV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginCrossMarginDataV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginCrossMarginDataV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginCrossMarginDataV1RespItem(val *GetMarginCrossMarginDataV1RespItem) *NullableGetMarginCrossMarginDataV1RespItem {
	return &NullableGetMarginCrossMarginDataV1RespItem{value: val, isSet: true}
}

func (v NullableGetMarginCrossMarginDataV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginCrossMarginDataV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


