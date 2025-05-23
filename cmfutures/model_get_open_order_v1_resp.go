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

// checks if the GetOpenOrderV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOpenOrderV1Resp{}

// GetOpenOrderV1Resp struct for GetOpenOrderV1Resp
type GetOpenOrderV1Resp struct {
	ActivatePrice *string `json:"activatePrice,omitempty"`
	AvgPrice *string `json:"avgPrice,omitempty"`
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	ClosePosition *bool `json:"closePosition,omitempty"`
	CumBase *string `json:"cumBase,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	OrigType *string `json:"origType,omitempty"`
	Pair *string `json:"pair,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Price *string `json:"price,omitempty"`
	PriceMatch *string `json:"priceMatch,omitempty"`
	PriceProtect *bool `json:"priceProtect,omitempty"`
	PriceRate *string `json:"priceRate,omitempty"`
	ReduceOnly *bool `json:"reduceOnly,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Side *string `json:"side,omitempty"`
	Status *string `json:"status,omitempty"`
	StopPrice *string `json:"stopPrice,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
	WorkingType *string `json:"workingType,omitempty"`
}

// NewGetOpenOrderV1Resp instantiates a new GetOpenOrderV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOpenOrderV1Resp() *GetOpenOrderV1Resp {
	this := GetOpenOrderV1Resp{}
	return &this
}

// NewGetOpenOrderV1RespWithDefaults instantiates a new GetOpenOrderV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOpenOrderV1RespWithDefaults() *GetOpenOrderV1Resp {
	this := GetOpenOrderV1Resp{}
	return &this
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetActivatePrice() string {
	if o == nil || IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetActivatePriceOk() (*string, bool) {
	if o == nil || IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasActivatePrice() bool {
	if o != nil && !IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *GetOpenOrderV1Resp) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetAvgPrice() string {
	if o == nil || IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetAvgPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasAvgPrice() bool {
	if o != nil && !IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *GetOpenOrderV1Resp) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *GetOpenOrderV1Resp) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetClosePosition() bool {
	if o == nil || IsNil(o.ClosePosition) {
		var ret bool
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetClosePositionOk() (*bool, bool) {
	if o == nil || IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasClosePosition() bool {
	if o != nil && !IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given bool and assigns it to the ClosePosition field.
func (o *GetOpenOrderV1Resp) SetClosePosition(v bool) {
	o.ClosePosition = &v
}

// GetCumBase returns the CumBase field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetCumBase() string {
	if o == nil || IsNil(o.CumBase) {
		var ret string
		return ret
	}
	return *o.CumBase
}

// GetCumBaseOk returns a tuple with the CumBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetCumBaseOk() (*string, bool) {
	if o == nil || IsNil(o.CumBase) {
		return nil, false
	}
	return o.CumBase, true
}

// HasCumBase returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasCumBase() bool {
	if o != nil && !IsNil(o.CumBase) {
		return true
	}

	return false
}

// SetCumBase gets a reference to the given string and assigns it to the CumBase field.
func (o *GetOpenOrderV1Resp) SetCumBase(v string) {
	o.CumBase = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *GetOpenOrderV1Resp) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetOpenOrderV1Resp) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *GetOpenOrderV1Resp) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetOrigType returns the OrigType field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetOrigType() string {
	if o == nil || IsNil(o.OrigType) {
		var ret string
		return ret
	}
	return *o.OrigType
}

// GetOrigTypeOk returns a tuple with the OrigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetOrigTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrigType) {
		return nil, false
	}
	return o.OrigType, true
}

// HasOrigType returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasOrigType() bool {
	if o != nil && !IsNil(o.OrigType) {
		return true
	}

	return false
}

// SetOrigType gets a reference to the given string and assigns it to the OrigType field.
func (o *GetOpenOrderV1Resp) SetOrigType(v string) {
	o.OrigType = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetPair() string {
	if o == nil || IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetPairOk() (*string, bool) {
	if o == nil || IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasPair() bool {
	if o != nil && !IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *GetOpenOrderV1Resp) SetPair(v string) {
	o.Pair = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetOpenOrderV1Resp) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetOpenOrderV1Resp) SetPrice(v string) {
	o.Price = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetPriceMatch() string {
	if o == nil || IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetPriceMatchOk() (*string, bool) {
	if o == nil || IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasPriceMatch() bool {
	if o != nil && !IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *GetOpenOrderV1Resp) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetPriceProtect() bool {
	if o == nil || IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetPriceProtectOk() (*bool, bool) {
	if o == nil || IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasPriceProtect() bool {
	if o != nil && !IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *GetOpenOrderV1Resp) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetPriceRate() string {
	if o == nil || IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetPriceRateOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasPriceRate() bool {
	if o != nil && !IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *GetOpenOrderV1Resp) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *GetOpenOrderV1Resp) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *GetOpenOrderV1Resp) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetOpenOrderV1Resp) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOpenOrderV1Resp) SetStatus(v string) {
	o.Status = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *GetOpenOrderV1Resp) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetOpenOrderV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetOpenOrderV1Resp) SetTime(v int64) {
	o.Time = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *GetOpenOrderV1Resp) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOpenOrderV1Resp) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetOpenOrderV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *GetOpenOrderV1Resp) GetWorkingType() string {
	if o == nil || IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOpenOrderV1Resp) GetWorkingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *GetOpenOrderV1Resp) HasWorkingType() bool {
	if o != nil && !IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *GetOpenOrderV1Resp) SetWorkingType(v string) {
	o.WorkingType = &v
}

func (o GetOpenOrderV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOpenOrderV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivatePrice) {
		toSerialize["activatePrice"] = o.ActivatePrice
	}
	if !IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.ClosePosition) {
		toSerialize["closePosition"] = o.ClosePosition
	}
	if !IsNil(o.CumBase) {
		toSerialize["cumBase"] = o.CumBase
	}
	if !IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !IsNil(o.OrigType) {
		toSerialize["origType"] = o.OrigType
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
	if !IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
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
	if !IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	return toSerialize, nil
}

type NullableGetOpenOrderV1Resp struct {
	value *GetOpenOrderV1Resp
	isSet bool
}

func (v NullableGetOpenOrderV1Resp) Get() *GetOpenOrderV1Resp {
	return v.value
}

func (v *NullableGetOpenOrderV1Resp) Set(val *GetOpenOrderV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOpenOrderV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOpenOrderV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOpenOrderV1Resp(val *GetOpenOrderV1Resp) *NullableGetOpenOrderV1Resp {
	return &NullableGetOpenOrderV1Resp{value: val, isSet: true}
}

func (v NullableGetOpenOrderV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOpenOrderV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


