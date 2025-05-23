/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
)

// checks if the DeleteCmConditionalOrderV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCmConditionalOrderV1Resp{}

// DeleteCmConditionalOrderV1Resp struct for DeleteCmConditionalOrderV1Resp
type DeleteCmConditionalOrderV1Resp struct {
	ActivatePrice *string `json:"activatePrice,omitempty"`
	BookTime *int64 `json:"bookTime,omitempty"`
	NewClientStrategyId *string `json:"newClientStrategyId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
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

// NewDeleteCmConditionalOrderV1Resp instantiates a new DeleteCmConditionalOrderV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCmConditionalOrderV1Resp() *DeleteCmConditionalOrderV1Resp {
	this := DeleteCmConditionalOrderV1Resp{}
	return &this
}

// NewDeleteCmConditionalOrderV1RespWithDefaults instantiates a new DeleteCmConditionalOrderV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCmConditionalOrderV1RespWithDefaults() *DeleteCmConditionalOrderV1Resp {
	this := DeleteCmConditionalOrderV1Resp{}
	return &this
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetActivatePrice() string {
	if o == nil || IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetActivatePriceOk() (*string, bool) {
	if o == nil || IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasActivatePrice() bool {
	if o != nil && !IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *DeleteCmConditionalOrderV1Resp) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetBookTime() int64 {
	if o == nil || IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetBookTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasBookTime() bool {
	if o != nil && !IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *DeleteCmConditionalOrderV1Resp) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetNewClientStrategyId returns the NewClientStrategyId field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetNewClientStrategyId() string {
	if o == nil || IsNil(o.NewClientStrategyId) {
		var ret string
		return ret
	}
	return *o.NewClientStrategyId
}

// GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetNewClientStrategyIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewClientStrategyId) {
		return nil, false
	}
	return o.NewClientStrategyId, true
}

// HasNewClientStrategyId returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasNewClientStrategyId() bool {
	if o != nil && !IsNil(o.NewClientStrategyId) {
		return true
	}

	return false
}

// SetNewClientStrategyId gets a reference to the given string and assigns it to the NewClientStrategyId field.
func (o *DeleteCmConditionalOrderV1Resp) SetNewClientStrategyId(v string) {
	o.NewClientStrategyId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *DeleteCmConditionalOrderV1Resp) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *DeleteCmConditionalOrderV1Resp) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *DeleteCmConditionalOrderV1Resp) SetPrice(v string) {
	o.Price = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetPriceProtect() bool {
	if o == nil || IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetPriceProtectOk() (*bool, bool) {
	if o == nil || IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasPriceProtect() bool {
	if o != nil && !IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *DeleteCmConditionalOrderV1Resp) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetPriceRate() string {
	if o == nil || IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetPriceRateOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasPriceRate() bool {
	if o != nil && !IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *DeleteCmConditionalOrderV1Resp) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *DeleteCmConditionalOrderV1Resp) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *DeleteCmConditionalOrderV1Resp) SetSide(v string) {
	o.Side = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *DeleteCmConditionalOrderV1Resp) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetStrategyId() int64 {
	if o == nil || IsNil(o.StrategyId) {
		var ret int64
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetStrategyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasStrategyId() bool {
	if o != nil && !IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given int64 and assigns it to the StrategyId field.
func (o *DeleteCmConditionalOrderV1Resp) SetStrategyId(v int64) {
	o.StrategyId = &v
}

// GetStrategyStatus returns the StrategyStatus field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetStrategyStatus() string {
	if o == nil || IsNil(o.StrategyStatus) {
		var ret string
		return ret
	}
	return *o.StrategyStatus
}

// GetStrategyStatusOk returns a tuple with the StrategyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetStrategyStatusOk() (*string, bool) {
	if o == nil || IsNil(o.StrategyStatus) {
		return nil, false
	}
	return o.StrategyStatus, true
}

// HasStrategyStatus returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasStrategyStatus() bool {
	if o != nil && !IsNil(o.StrategyStatus) {
		return true
	}

	return false
}

// SetStrategyStatus gets a reference to the given string and assigns it to the StrategyStatus field.
func (o *DeleteCmConditionalOrderV1Resp) SetStrategyStatus(v string) {
	o.StrategyStatus = &v
}

// GetStrategyType returns the StrategyType field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetStrategyType() string {
	if o == nil || IsNil(o.StrategyType) {
		var ret string
		return ret
	}
	return *o.StrategyType
}

// GetStrategyTypeOk returns a tuple with the StrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetStrategyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StrategyType) {
		return nil, false
	}
	return o.StrategyType, true
}

// HasStrategyType returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasStrategyType() bool {
	if o != nil && !IsNil(o.StrategyType) {
		return true
	}

	return false
}

// SetStrategyType gets a reference to the given string and assigns it to the StrategyType field.
func (o *DeleteCmConditionalOrderV1Resp) SetStrategyType(v string) {
	o.StrategyType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *DeleteCmConditionalOrderV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *DeleteCmConditionalOrderV1Resp) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *DeleteCmConditionalOrderV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *DeleteCmConditionalOrderV1Resp) GetWorkingType() string {
	if o == nil || IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCmConditionalOrderV1Resp) GetWorkingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *DeleteCmConditionalOrderV1Resp) HasWorkingType() bool {
	if o != nil && !IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *DeleteCmConditionalOrderV1Resp) SetWorkingType(v string) {
	o.WorkingType = &v
}

func (o DeleteCmConditionalOrderV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCmConditionalOrderV1Resp) ToMap() (map[string]interface{}, error) {
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

type NullableDeleteCmConditionalOrderV1Resp struct {
	value *DeleteCmConditionalOrderV1Resp
	isSet bool
}

func (v NullableDeleteCmConditionalOrderV1Resp) Get() *DeleteCmConditionalOrderV1Resp {
	return v.value
}

func (v *NullableDeleteCmConditionalOrderV1Resp) Set(val *DeleteCmConditionalOrderV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCmConditionalOrderV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCmConditionalOrderV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCmConditionalOrderV1Resp(val *DeleteCmConditionalOrderV1Resp) *NullableDeleteCmConditionalOrderV1Resp {
	return &NullableDeleteCmConditionalOrderV1Resp{value: val, isSet: true}
}

func (v NullableDeleteCmConditionalOrderV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCmConditionalOrderV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


