/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the GetAccountV3RespPositionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountV3RespPositionsInner{}

// GetAccountV3RespPositionsInner struct for GetAccountV3RespPositionsInner
type GetAccountV3RespPositionsInner struct {
	InitialMargin *string `json:"initialMargin,omitempty"`
	IsolatedMargin *string `json:"isolatedMargin,omitempty"`
	IsolatedWallet *string `json:"isolatedWallet,omitempty"`
	MaintMargin *string `json:"maintMargin,omitempty"`
	Notional *string `json:"notional,omitempty"`
	PositionAmt *string `json:"positionAmt,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	UnrealizedProfit *string `json:"unrealizedProfit,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewGetAccountV3RespPositionsInner instantiates a new GetAccountV3RespPositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountV3RespPositionsInner() *GetAccountV3RespPositionsInner {
	this := GetAccountV3RespPositionsInner{}
	return &this
}

// NewGetAccountV3RespPositionsInnerWithDefaults instantiates a new GetAccountV3RespPositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountV3RespPositionsInnerWithDefaults() *GetAccountV3RespPositionsInner {
	this := GetAccountV3RespPositionsInner{}
	return &this
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetInitialMargin() string {
	if o == nil || IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasInitialMargin() bool {
	if o != nil && !IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *GetAccountV3RespPositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetIsolatedMargin returns the IsolatedMargin field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetIsolatedMargin() string {
	if o == nil || IsNil(o.IsolatedMargin) {
		var ret string
		return ret
	}
	return *o.IsolatedMargin
}

// GetIsolatedMarginOk returns a tuple with the IsolatedMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetIsolatedMarginOk() (*string, bool) {
	if o == nil || IsNil(o.IsolatedMargin) {
		return nil, false
	}
	return o.IsolatedMargin, true
}

// HasIsolatedMargin returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasIsolatedMargin() bool {
	if o != nil && !IsNil(o.IsolatedMargin) {
		return true
	}

	return false
}

// SetIsolatedMargin gets a reference to the given string and assigns it to the IsolatedMargin field.
func (o *GetAccountV3RespPositionsInner) SetIsolatedMargin(v string) {
	o.IsolatedMargin = &v
}

// GetIsolatedWallet returns the IsolatedWallet field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetIsolatedWallet() string {
	if o == nil || IsNil(o.IsolatedWallet) {
		var ret string
		return ret
	}
	return *o.IsolatedWallet
}

// GetIsolatedWalletOk returns a tuple with the IsolatedWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetIsolatedWalletOk() (*string, bool) {
	if o == nil || IsNil(o.IsolatedWallet) {
		return nil, false
	}
	return o.IsolatedWallet, true
}

// HasIsolatedWallet returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasIsolatedWallet() bool {
	if o != nil && !IsNil(o.IsolatedWallet) {
		return true
	}

	return false
}

// SetIsolatedWallet gets a reference to the given string and assigns it to the IsolatedWallet field.
func (o *GetAccountV3RespPositionsInner) SetIsolatedWallet(v string) {
	o.IsolatedWallet = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetMaintMargin() string {
	if o == nil || IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasMaintMargin() bool {
	if o != nil && !IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *GetAccountV3RespPositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetNotional() string {
	if o == nil || IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetNotionalOk() (*string, bool) {
	if o == nil || IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasNotional() bool {
	if o != nil && !IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given string and assigns it to the Notional field.
func (o *GetAccountV3RespPositionsInner) SetNotional(v string) {
	o.Notional = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetPositionAmt() string {
	if o == nil || IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasPositionAmt() bool {
	if o != nil && !IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *GetAccountV3RespPositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetAccountV3RespPositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetAccountV3RespPositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetUnrealizedProfit() string {
	if o == nil || IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *GetAccountV3RespPositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetAccountV3RespPositionsInner) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountV3RespPositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetAccountV3RespPositionsInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetAccountV3RespPositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetAccountV3RespPositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountV3RespPositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !IsNil(o.IsolatedMargin) {
		toSerialize["isolatedMargin"] = o.IsolatedMargin
	}
	if !IsNil(o.IsolatedWallet) {
		toSerialize["isolatedWallet"] = o.IsolatedWallet
	}
	if !IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
	}
	if !IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.UnrealizedProfit) {
		toSerialize["unrealizedProfit"] = o.UnrealizedProfit
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableGetAccountV3RespPositionsInner struct {
	value *GetAccountV3RespPositionsInner
	isSet bool
}

func (v NullableGetAccountV3RespPositionsInner) Get() *GetAccountV3RespPositionsInner {
	return v.value
}

func (v *NullableGetAccountV3RespPositionsInner) Set(val *GetAccountV3RespPositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountV3RespPositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountV3RespPositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountV3RespPositionsInner(val *GetAccountV3RespPositionsInner) *NullableGetAccountV3RespPositionsInner {
	return &NullableGetAccountV3RespPositionsInner{value: val, isSet: true}
}

func (v NullableGetAccountV3RespPositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountV3RespPositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


