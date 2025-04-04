/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the PmarginCreateCmConditionalOrderV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PmarginCreateCmConditionalOrderV1Resp{}

// PmarginCreateCmConditionalOrderV1Resp struct for PmarginCreateCmConditionalOrderV1Resp
type PmarginCreateCmConditionalOrderV1Resp struct {
	ActivatePrice *string `json:"activatePrice,omitempty"`
	BookTime *int64 `json:"bookTime,omitempty"`
	NewClientStrategyId *string `json:"newClientStrategyId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	Pair *string `json:"pair,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Price *string `json:"price,omitempty"`
	PriceProtect *bool `json:"priceProtect,omitempty"`
	PriceRate *string `json:"priceRate,omitempty"`
	ReduceOnly *bool `json:"reduceOnly,omitempty"`
	Side *string `json:"side,omitempty"`
	StopPrice *string `json:"stopPrice,omitempty"`
	StrategyId *int64 `json:"strategyId,omitempty"`
	StrategyStatus *string `json:"strategyStatus,omitempty"`
	StrategyType *string `json:"strategyType,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
	WorkingType *string `json:"workingType,omitempty"`
}

// NewPmarginCreateCmConditionalOrderV1Resp instantiates a new PmarginCreateCmConditionalOrderV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPmarginCreateCmConditionalOrderV1Resp() *PmarginCreateCmConditionalOrderV1Resp {
	this := PmarginCreateCmConditionalOrderV1Resp{}
	return &this
}

// NewPmarginCreateCmConditionalOrderV1RespWithDefaults instantiates a new PmarginCreateCmConditionalOrderV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPmarginCreateCmConditionalOrderV1RespWithDefaults() *PmarginCreateCmConditionalOrderV1Resp {
	this := PmarginCreateCmConditionalOrderV1Resp{}
	return &this
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetActivatePrice() string {
	if o == nil || IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetActivatePriceOk() (*string, bool) {
	if o == nil || IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasActivatePrice() bool {
	if o != nil && !IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetBookTime() int64 {
	if o == nil || IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetBookTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasBookTime() bool {
	if o != nil && !IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetNewClientStrategyId returns the NewClientStrategyId field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetNewClientStrategyId() string {
	if o == nil || IsNil(o.NewClientStrategyId) {
		var ret string
		return ret
	}
	return *o.NewClientStrategyId
}

// GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetNewClientStrategyIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewClientStrategyId) {
		return nil, false
	}
	return o.NewClientStrategyId, true
}

// HasNewClientStrategyId returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasNewClientStrategyId() bool {
	if o != nil && !IsNil(o.NewClientStrategyId) {
		return true
	}

	return false
}

// SetNewClientStrategyId gets a reference to the given string and assigns it to the NewClientStrategyId field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetNewClientStrategyId(v string) {
	o.NewClientStrategyId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPair() string {
	if o == nil || IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPairOk() (*string, bool) {
	if o == nil || IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasPair() bool {
	if o != nil && !IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetPair(v string) {
	o.Pair = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetPrice(v string) {
	o.Price = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPriceProtect() bool {
	if o == nil || IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPriceProtectOk() (*bool, bool) {
	if o == nil || IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasPriceProtect() bool {
	if o != nil && !IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPriceRate() string {
	if o == nil || IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetPriceRateOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasPriceRate() bool {
	if o != nil && !IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetSide(v string) {
	o.Side = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStrategyId() int64 {
	if o == nil || IsNil(o.StrategyId) {
		var ret int64
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStrategyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasStrategyId() bool {
	if o != nil && !IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given int64 and assigns it to the StrategyId field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetStrategyId(v int64) {
	o.StrategyId = &v
}

// GetStrategyStatus returns the StrategyStatus field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStrategyStatus() string {
	if o == nil || IsNil(o.StrategyStatus) {
		var ret string
		return ret
	}
	return *o.StrategyStatus
}

// GetStrategyStatusOk returns a tuple with the StrategyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStrategyStatusOk() (*string, bool) {
	if o == nil || IsNil(o.StrategyStatus) {
		return nil, false
	}
	return o.StrategyStatus, true
}

// HasStrategyStatus returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasStrategyStatus() bool {
	if o != nil && !IsNil(o.StrategyStatus) {
		return true
	}

	return false
}

// SetStrategyStatus gets a reference to the given string and assigns it to the StrategyStatus field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetStrategyStatus(v string) {
	o.StrategyStatus = &v
}

// GetStrategyType returns the StrategyType field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStrategyType() string {
	if o == nil || IsNil(o.StrategyType) {
		var ret string
		return ret
	}
	return *o.StrategyType
}

// GetStrategyTypeOk returns a tuple with the StrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetStrategyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StrategyType) {
		return nil, false
	}
	return o.StrategyType, true
}

// HasStrategyType returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasStrategyType() bool {
	if o != nil && !IsNil(o.StrategyType) {
		return true
	}

	return false
}

// SetStrategyType gets a reference to the given string and assigns it to the StrategyType field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetStrategyType(v string) {
	o.StrategyType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetWorkingType() string {
	if o == nil || IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) GetWorkingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *PmarginCreateCmConditionalOrderV1Resp) HasWorkingType() bool {
	if o != nil && !IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *PmarginCreateCmConditionalOrderV1Resp) SetWorkingType(v string) {
	o.WorkingType = &v
}

func (o PmarginCreateCmConditionalOrderV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PmarginCreateCmConditionalOrderV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivatePrice) {
		toSerialize["activatePrice"] = o.ActivatePrice
	}
	if !IsNil(o.BookTime) {
		toSerialize["bookTime"] = o.BookTime
	}
	if !IsNil(o.NewClientStrategyId) {
		toSerialize["newClientStrategyId"] = o.NewClientStrategyId
	}
	if !IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceProtect) {
		toSerialize["priceProtect"] = o.PriceProtect
	}
	if !IsNil(o.PriceRate) {
		toSerialize["priceRate"] = o.PriceRate
	}
	if !IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !IsNil(o.StrategyId) {
		toSerialize["strategyId"] = o.StrategyId
	}
	if !IsNil(o.StrategyStatus) {
		toSerialize["strategyStatus"] = o.StrategyStatus
	}
	if !IsNil(o.StrategyType) {
		toSerialize["strategyType"] = o.StrategyType
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	return toSerialize, nil
}

type NullablePmarginCreateCmConditionalOrderV1Resp struct {
	value *PmarginCreateCmConditionalOrderV1Resp
	isSet bool
}

func (v NullablePmarginCreateCmConditionalOrderV1Resp) Get() *PmarginCreateCmConditionalOrderV1Resp {
	return v.value
}

func (v *NullablePmarginCreateCmConditionalOrderV1Resp) Set(val *PmarginCreateCmConditionalOrderV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginCreateCmConditionalOrderV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginCreateCmConditionalOrderV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginCreateCmConditionalOrderV1Resp(val *PmarginCreateCmConditionalOrderV1Resp) *NullablePmarginCreateCmConditionalOrderV1Resp {
	return &NullablePmarginCreateCmConditionalOrderV1Resp{value: val, isSet: true}
}

func (v NullablePmarginCreateCmConditionalOrderV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginCreateCmConditionalOrderV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


