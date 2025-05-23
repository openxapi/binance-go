/*
Binance COIN-M Futures API

OpenAPI specification for Binance exchange - Cmfutures API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the GetFuturesDataTakerBuySellVolRespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFuturesDataTakerBuySellVolRespItem{}

// GetFuturesDataTakerBuySellVolRespItem struct for GetFuturesDataTakerBuySellVolRespItem
type GetFuturesDataTakerBuySellVolRespItem struct {
	ContractType *string `json:"contractType,omitempty"`
	Pair *string `json:"pair,omitempty"`
	TakerBuyVol *string `json:"takerBuyVol,omitempty"`
	TakerBuyVolValue *string `json:"takerBuyVolValue,omitempty"`
	TakerSellVol *string `json:"takerSellVol,omitempty"`
	TakerSellVolValue *string `json:"takerSellVolValue,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// NewGetFuturesDataTakerBuySellVolRespItem instantiates a new GetFuturesDataTakerBuySellVolRespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesDataTakerBuySellVolRespItem() *GetFuturesDataTakerBuySellVolRespItem {
	this := GetFuturesDataTakerBuySellVolRespItem{}
	return &this
}

// NewGetFuturesDataTakerBuySellVolRespItemWithDefaults instantiates a new GetFuturesDataTakerBuySellVolRespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesDataTakerBuySellVolRespItemWithDefaults() *GetFuturesDataTakerBuySellVolRespItem {
	this := GetFuturesDataTakerBuySellVolRespItem{}
	return &this
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetContractType() string {
	if o == nil || IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetContractTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *GetFuturesDataTakerBuySellVolRespItem) SetContractType(v string) {
	o.ContractType = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetPair() string {
	if o == nil || IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetPairOk() (*string, bool) {
	if o == nil || IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) HasPair() bool {
	if o != nil && !IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *GetFuturesDataTakerBuySellVolRespItem) SetPair(v string) {
	o.Pair = &v
}

// GetTakerBuyVol returns the TakerBuyVol field value if set, zero value otherwise.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerBuyVol() string {
	if o == nil || IsNil(o.TakerBuyVol) {
		var ret string
		return ret
	}
	return *o.TakerBuyVol
}

// GetTakerBuyVolOk returns a tuple with the TakerBuyVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerBuyVolOk() (*string, bool) {
	if o == nil || IsNil(o.TakerBuyVol) {
		return nil, false
	}
	return o.TakerBuyVol, true
}

// HasTakerBuyVol returns a boolean if a field has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) HasTakerBuyVol() bool {
	if o != nil && !IsNil(o.TakerBuyVol) {
		return true
	}

	return false
}

// SetTakerBuyVol gets a reference to the given string and assigns it to the TakerBuyVol field.
func (o *GetFuturesDataTakerBuySellVolRespItem) SetTakerBuyVol(v string) {
	o.TakerBuyVol = &v
}

// GetTakerBuyVolValue returns the TakerBuyVolValue field value if set, zero value otherwise.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerBuyVolValue() string {
	if o == nil || IsNil(o.TakerBuyVolValue) {
		var ret string
		return ret
	}
	return *o.TakerBuyVolValue
}

// GetTakerBuyVolValueOk returns a tuple with the TakerBuyVolValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerBuyVolValueOk() (*string, bool) {
	if o == nil || IsNil(o.TakerBuyVolValue) {
		return nil, false
	}
	return o.TakerBuyVolValue, true
}

// HasTakerBuyVolValue returns a boolean if a field has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) HasTakerBuyVolValue() bool {
	if o != nil && !IsNil(o.TakerBuyVolValue) {
		return true
	}

	return false
}

// SetTakerBuyVolValue gets a reference to the given string and assigns it to the TakerBuyVolValue field.
func (o *GetFuturesDataTakerBuySellVolRespItem) SetTakerBuyVolValue(v string) {
	o.TakerBuyVolValue = &v
}

// GetTakerSellVol returns the TakerSellVol field value if set, zero value otherwise.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerSellVol() string {
	if o == nil || IsNil(o.TakerSellVol) {
		var ret string
		return ret
	}
	return *o.TakerSellVol
}

// GetTakerSellVolOk returns a tuple with the TakerSellVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerSellVolOk() (*string, bool) {
	if o == nil || IsNil(o.TakerSellVol) {
		return nil, false
	}
	return o.TakerSellVol, true
}

// HasTakerSellVol returns a boolean if a field has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) HasTakerSellVol() bool {
	if o != nil && !IsNil(o.TakerSellVol) {
		return true
	}

	return false
}

// SetTakerSellVol gets a reference to the given string and assigns it to the TakerSellVol field.
func (o *GetFuturesDataTakerBuySellVolRespItem) SetTakerSellVol(v string) {
	o.TakerSellVol = &v
}

// GetTakerSellVolValue returns the TakerSellVolValue field value if set, zero value otherwise.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerSellVolValue() string {
	if o == nil || IsNil(o.TakerSellVolValue) {
		var ret string
		return ret
	}
	return *o.TakerSellVolValue
}

// GetTakerSellVolValueOk returns a tuple with the TakerSellVolValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTakerSellVolValueOk() (*string, bool) {
	if o == nil || IsNil(o.TakerSellVolValue) {
		return nil, false
	}
	return o.TakerSellVolValue, true
}

// HasTakerSellVolValue returns a boolean if a field has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) HasTakerSellVolValue() bool {
	if o != nil && !IsNil(o.TakerSellVolValue) {
		return true
	}

	return false
}

// SetTakerSellVolValue gets a reference to the given string and assigns it to the TakerSellVolValue field.
func (o *GetFuturesDataTakerBuySellVolRespItem) SetTakerSellVolValue(v string) {
	o.TakerSellVolValue = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetFuturesDataTakerBuySellVolRespItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetFuturesDataTakerBuySellVolRespItem) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o GetFuturesDataTakerBuySellVolRespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesDataTakerBuySellVolRespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractType) {
		toSerialize["contractType"] = o.ContractType
	}
	if !IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !IsNil(o.TakerBuyVol) {
		toSerialize["takerBuyVol"] = o.TakerBuyVol
	}
	if !IsNil(o.TakerBuyVolValue) {
		toSerialize["takerBuyVolValue"] = o.TakerBuyVolValue
	}
	if !IsNil(o.TakerSellVol) {
		toSerialize["takerSellVol"] = o.TakerSellVol
	}
	if !IsNil(o.TakerSellVolValue) {
		toSerialize["takerSellVolValue"] = o.TakerSellVolValue
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableGetFuturesDataTakerBuySellVolRespItem struct {
	value *GetFuturesDataTakerBuySellVolRespItem
	isSet bool
}

func (v NullableGetFuturesDataTakerBuySellVolRespItem) Get() *GetFuturesDataTakerBuySellVolRespItem {
	return v.value
}

func (v *NullableGetFuturesDataTakerBuySellVolRespItem) Set(val *GetFuturesDataTakerBuySellVolRespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesDataTakerBuySellVolRespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesDataTakerBuySellVolRespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesDataTakerBuySellVolRespItem(val *GetFuturesDataTakerBuySellVolRespItem) *NullableGetFuturesDataTakerBuySellVolRespItem {
	return &NullableGetFuturesDataTakerBuySellVolRespItem{value: val, isSet: true}
}

func (v NullableGetFuturesDataTakerBuySellVolRespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesDataTakerBuySellVolRespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


