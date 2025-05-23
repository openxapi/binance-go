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

// checks if the GetBrokerRebateRecentRecordV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBrokerRebateRecentRecordV1RespItem{}

// GetBrokerRebateRecentRecordV1RespItem struct for GetBrokerRebateRecentRecordV1RespItem
type GetBrokerRebateRecentRecordV1RespItem struct {
	Asset *string `json:"asset,omitempty"`
	Income *string `json:"income,omitempty"`
	Status *int32 `json:"status,omitempty"`
	SubaccountId *string `json:"subaccountId,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
	TradeId *int64 `json:"tradeId,omitempty"`
}

// NewGetBrokerRebateRecentRecordV1RespItem instantiates a new GetBrokerRebateRecentRecordV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBrokerRebateRecentRecordV1RespItem() *GetBrokerRebateRecentRecordV1RespItem {
	this := GetBrokerRebateRecentRecordV1RespItem{}
	return &this
}

// NewGetBrokerRebateRecentRecordV1RespItemWithDefaults instantiates a new GetBrokerRebateRecentRecordV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBrokerRebateRecentRecordV1RespItemWithDefaults() *GetBrokerRebateRecentRecordV1RespItem {
	this := GetBrokerRebateRecentRecordV1RespItem{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetBrokerRebateRecentRecordV1RespItem) SetAsset(v string) {
	o.Asset = &v
}

// GetIncome returns the Income field value if set, zero value otherwise.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetIncome() string {
	if o == nil || IsNil(o.Income) {
		var ret string
		return ret
	}
	return *o.Income
}

// GetIncomeOk returns a tuple with the Income field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetIncomeOk() (*string, bool) {
	if o == nil || IsNil(o.Income) {
		return nil, false
	}
	return o.Income, true
}

// HasIncome returns a boolean if a field has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) HasIncome() bool {
	if o != nil && !IsNil(o.Income) {
		return true
	}

	return false
}

// SetIncome gets a reference to the given string and assigns it to the Income field.
func (o *GetBrokerRebateRecentRecordV1RespItem) SetIncome(v string) {
	o.Income = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *GetBrokerRebateRecentRecordV1RespItem) SetStatus(v int32) {
	o.Status = &v
}

// GetSubaccountId returns the SubaccountId field value if set, zero value otherwise.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetSubaccountId() string {
	if o == nil || IsNil(o.SubaccountId) {
		var ret string
		return ret
	}
	return *o.SubaccountId
}

// GetSubaccountIdOk returns a tuple with the SubaccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetSubaccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubaccountId) {
		return nil, false
	}
	return o.SubaccountId, true
}

// HasSubaccountId returns a boolean if a field has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) HasSubaccountId() bool {
	if o != nil && !IsNil(o.SubaccountId) {
		return true
	}

	return false
}

// SetSubaccountId gets a reference to the given string and assigns it to the SubaccountId field.
func (o *GetBrokerRebateRecentRecordV1RespItem) SetSubaccountId(v string) {
	o.SubaccountId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetBrokerRebateRecentRecordV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetBrokerRebateRecentRecordV1RespItem) SetTime(v int64) {
	o.Time = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetTradeId() int64 {
	if o == nil || IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) GetTradeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *GetBrokerRebateRecentRecordV1RespItem) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *GetBrokerRebateRecentRecordV1RespItem) SetTradeId(v int64) {
	o.TradeId = &v
}

func (o GetBrokerRebateRecentRecordV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBrokerRebateRecentRecordV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.Income) {
		toSerialize["income"] = o.Income
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubaccountId) {
		toSerialize["subaccountId"] = o.SubaccountId
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	return toSerialize, nil
}

type NullableGetBrokerRebateRecentRecordV1RespItem struct {
	value *GetBrokerRebateRecentRecordV1RespItem
	isSet bool
}

func (v NullableGetBrokerRebateRecentRecordV1RespItem) Get() *GetBrokerRebateRecentRecordV1RespItem {
	return v.value
}

func (v *NullableGetBrokerRebateRecentRecordV1RespItem) Set(val *GetBrokerRebateRecentRecordV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBrokerRebateRecentRecordV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBrokerRebateRecentRecordV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBrokerRebateRecentRecordV1RespItem(val *GetBrokerRebateRecentRecordV1RespItem) *NullableGetBrokerRebateRecentRecordV1RespItem {
	return &NullableGetBrokerRebateRecentRecordV1RespItem{value: val, isSet: true}
}

func (v NullableGetBrokerRebateRecentRecordV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBrokerRebateRecentRecordV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


