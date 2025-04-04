/*
Binance COIN-M Futures API

OpenAPI specification for Binance exchange - Cmfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
)

// checks if the CmfuturesGetAccountV1RespPositionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmfuturesGetAccountV1RespPositionsInner{}

// CmfuturesGetAccountV1RespPositionsInner struct for CmfuturesGetAccountV1RespPositionsInner
type CmfuturesGetAccountV1RespPositionsInner struct {
	BreakEvenPrice *string `json:"breakEvenPrice,omitempty"`
	EntryPrice *string `json:"entryPrice,omitempty"`
	InitialMargin *string `json:"initialMargin,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Leverage *string `json:"leverage,omitempty"`
	MaintMargin *string `json:"maintMargin,omitempty"`
	MaxQty *string `json:"maxQty,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	PositionAmt *string `json:"positionAmt,omitempty"`
	PositionInitialMargin *string `json:"positionInitialMargin,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	UnrealizedProfit *string `json:"unrealizedProfit,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewCmfuturesGetAccountV1RespPositionsInner instantiates a new CmfuturesGetAccountV1RespPositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmfuturesGetAccountV1RespPositionsInner() *CmfuturesGetAccountV1RespPositionsInner {
	this := CmfuturesGetAccountV1RespPositionsInner{}
	return &this
}

// NewCmfuturesGetAccountV1RespPositionsInnerWithDefaults instantiates a new CmfuturesGetAccountV1RespPositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmfuturesGetAccountV1RespPositionsInnerWithDefaults() *CmfuturesGetAccountV1RespPositionsInner {
	this := CmfuturesGetAccountV1RespPositionsInner{}
	return &this
}

// GetBreakEvenPrice returns the BreakEvenPrice field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetBreakEvenPrice() string {
	if o == nil || IsNil(o.BreakEvenPrice) {
		var ret string
		return ret
	}
	return *o.BreakEvenPrice
}

// GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetBreakEvenPriceOk() (*string, bool) {
	if o == nil || IsNil(o.BreakEvenPrice) {
		return nil, false
	}
	return o.BreakEvenPrice, true
}

// HasBreakEvenPrice returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasBreakEvenPrice() bool {
	if o != nil && !IsNil(o.BreakEvenPrice) {
		return true
	}

	return false
}

// SetBreakEvenPrice gets a reference to the given string and assigns it to the BreakEvenPrice field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetBreakEvenPrice(v string) {
	o.BreakEvenPrice = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetEntryPrice() string {
	if o == nil || IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetEntryPriceOk() (*string, bool) {
	if o == nil || IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasEntryPrice() bool {
	if o != nil && !IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetInitialMargin() string {
	if o == nil || IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasInitialMargin() bool {
	if o != nil && !IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetIsolated returns the Isolated field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetIsolated() bool {
	if o == nil || IsNil(o.Isolated) {
		var ret bool
		return ret
	}
	return *o.Isolated
}

// GetIsolatedOk returns a tuple with the Isolated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetIsolatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Isolated) {
		return nil, false
	}
	return o.Isolated, true
}

// HasIsolated returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasIsolated() bool {
	if o != nil && !IsNil(o.Isolated) {
		return true
	}

	return false
}

// SetIsolated gets a reference to the given bool and assigns it to the Isolated field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetIsolated(v bool) {
	o.Isolated = &v
}

// GetLeverage returns the Leverage field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetLeverage() string {
	if o == nil || IsNil(o.Leverage) {
		var ret string
		return ret
	}
	return *o.Leverage
}

// GetLeverageOk returns a tuple with the Leverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetLeverageOk() (*string, bool) {
	if o == nil || IsNil(o.Leverage) {
		return nil, false
	}
	return o.Leverage, true
}

// HasLeverage returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasLeverage() bool {
	if o != nil && !IsNil(o.Leverage) {
		return true
	}

	return false
}

// SetLeverage gets a reference to the given string and assigns it to the Leverage field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetLeverage(v string) {
	o.Leverage = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetMaintMargin() string {
	if o == nil || IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetMaintMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasMaintMargin() bool {
	if o != nil && !IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetMaxQty() string {
	if o == nil || IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasMaxQty() bool {
	if o != nil && !IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetOpenOrderInitialMargin() string {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasOpenOrderInitialMargin() bool {
	if o != nil && !IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetPositionAmt() string {
	if o == nil || IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetPositionAmtOk() (*string, bool) {
	if o == nil || IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasPositionAmt() bool {
	if o != nil && !IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetPositionInitialMargin() string {
	if o == nil || IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasPositionInitialMargin() bool {
	if o != nil && !IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnrealizedProfit returns the UnrealizedProfit field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetUnrealizedProfit() string {
	if o == nil || IsNil(o.UnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnrealizedProfit
}

// GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetUnrealizedProfitOk() (*string, bool) {
	if o == nil || IsNil(o.UnrealizedProfit) {
		return nil, false
	}
	return o.UnrealizedProfit, true
}

// HasUnrealizedProfit returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasUnrealizedProfit() bool {
	if o != nil && !IsNil(o.UnrealizedProfit) {
		return true
	}

	return false
}

// SetUnrealizedProfit gets a reference to the given string and assigns it to the UnrealizedProfit field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetUnrealizedProfit(v string) {
	o.UnrealizedProfit = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *CmfuturesGetAccountV1RespPositionsInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *CmfuturesGetAccountV1RespPositionsInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o CmfuturesGetAccountV1RespPositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmfuturesGetAccountV1RespPositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BreakEvenPrice) {
		toSerialize["breakEvenPrice"] = o.BreakEvenPrice
	}
	if !IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !IsNil(o.Isolated) {
		toSerialize["isolated"] = o.Isolated
	}
	if !IsNil(o.Leverage) {
		toSerialize["leverage"] = o.Leverage
	}
	if !IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !IsNil(o.OpenOrderInitialMargin) {
		toSerialize["openOrderInitialMargin"] = o.OpenOrderInitialMargin
	}
	if !IsNil(o.PositionAmt) {
		toSerialize["positionAmt"] = o.PositionAmt
	}
	if !IsNil(o.PositionInitialMargin) {
		toSerialize["positionInitialMargin"] = o.PositionInitialMargin
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

type NullableCmfuturesGetAccountV1RespPositionsInner struct {
	value *CmfuturesGetAccountV1RespPositionsInner
	isSet bool
}

func (v NullableCmfuturesGetAccountV1RespPositionsInner) Get() *CmfuturesGetAccountV1RespPositionsInner {
	return v.value
}

func (v *NullableCmfuturesGetAccountV1RespPositionsInner) Set(val *CmfuturesGetAccountV1RespPositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCmfuturesGetAccountV1RespPositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCmfuturesGetAccountV1RespPositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmfuturesGetAccountV1RespPositionsInner(val *CmfuturesGetAccountV1RespPositionsInner) *NullableCmfuturesGetAccountV1RespPositionsInner {
	return &NullableCmfuturesGetAccountV1RespPositionsInner{value: val, isSet: true}
}

func (v NullableCmfuturesGetAccountV1RespPositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmfuturesGetAccountV1RespPositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


