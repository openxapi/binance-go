/*
Binance Cfutures API

OpenAPI specification for Binance cryptocurrency exchange - Cfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the CfuturesCreateCountdownCancelAllV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CfuturesCreateCountdownCancelAllV1Resp{}

// CfuturesCreateCountdownCancelAllV1Resp struct for CfuturesCreateCountdownCancelAllV1Resp
type CfuturesCreateCountdownCancelAllV1Resp struct {
	CountdownTime *string `json:"countdownTime,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewCfuturesCreateCountdownCancelAllV1Resp instantiates a new CfuturesCreateCountdownCancelAllV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCfuturesCreateCountdownCancelAllV1Resp() *CfuturesCreateCountdownCancelAllV1Resp {
	this := CfuturesCreateCountdownCancelAllV1Resp{}
	return &this
}

// NewCfuturesCreateCountdownCancelAllV1RespWithDefaults instantiates a new CfuturesCreateCountdownCancelAllV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCfuturesCreateCountdownCancelAllV1RespWithDefaults() *CfuturesCreateCountdownCancelAllV1Resp {
	this := CfuturesCreateCountdownCancelAllV1Resp{}
	return &this
}

// GetCountdownTime returns the CountdownTime field value if set, zero value otherwise.
func (o *CfuturesCreateCountdownCancelAllV1Resp) GetCountdownTime() string {
	if o == nil || IsNil(o.CountdownTime) {
		var ret string
		return ret
	}
	return *o.CountdownTime
}

// GetCountdownTimeOk returns a tuple with the CountdownTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesCreateCountdownCancelAllV1Resp) GetCountdownTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CountdownTime) {
		return nil, false
	}
	return o.CountdownTime, true
}

// HasCountdownTime returns a boolean if a field has been set.
func (o *CfuturesCreateCountdownCancelAllV1Resp) HasCountdownTime() bool {
	if o != nil && !IsNil(o.CountdownTime) {
		return true
	}

	return false
}

// SetCountdownTime gets a reference to the given string and assigns it to the CountdownTime field.
func (o *CfuturesCreateCountdownCancelAllV1Resp) SetCountdownTime(v string) {
	o.CountdownTime = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CfuturesCreateCountdownCancelAllV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CfuturesCreateCountdownCancelAllV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CfuturesCreateCountdownCancelAllV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CfuturesCreateCountdownCancelAllV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

func (o CfuturesCreateCountdownCancelAllV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CfuturesCreateCountdownCancelAllV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountdownTime) {
		toSerialize["countdownTime"] = o.CountdownTime
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableCfuturesCreateCountdownCancelAllV1Resp struct {
	value *CfuturesCreateCountdownCancelAllV1Resp
	isSet bool
}

func (v NullableCfuturesCreateCountdownCancelAllV1Resp) Get() *CfuturesCreateCountdownCancelAllV1Resp {
	return v.value
}

func (v *NullableCfuturesCreateCountdownCancelAllV1Resp) Set(val *CfuturesCreateCountdownCancelAllV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCfuturesCreateCountdownCancelAllV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCfuturesCreateCountdownCancelAllV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCfuturesCreateCountdownCancelAllV1Resp(val *CfuturesCreateCountdownCancelAllV1Resp) *NullableCfuturesCreateCountdownCancelAllV1Resp {
	return &NullableCfuturesCreateCountdownCancelAllV1Resp{value: val, isSet: true}
}

func (v NullableCfuturesCreateCountdownCancelAllV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCfuturesCreateCountdownCancelAllV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


