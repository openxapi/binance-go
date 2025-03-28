/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetAllOrdersV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetAllOrdersV1RespItem{}

// UfuturesGetAllOrdersV1RespItem struct for UfuturesGetAllOrdersV1RespItem
type UfuturesGetAllOrdersV1RespItem struct {
	ActivatePrice *string `json:"activatePrice,omitempty"`
	AvgPrice *string `json:"avgPrice,omitempty"`
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	ClosePosition *bool `json:"closePosition,omitempty"`
	CumQuote *string `json:"cumQuote,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	GoodTillDate *int64 `json:"goodTillDate,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	OrigQty *string `json:"origQty,omitempty"`
	OrigType *string `json:"origType,omitempty"`
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

// NewUfuturesGetAllOrdersV1RespItem instantiates a new UfuturesGetAllOrdersV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetAllOrdersV1RespItem() *UfuturesGetAllOrdersV1RespItem {
	this := UfuturesGetAllOrdersV1RespItem{}
	return &this
}

// NewUfuturesGetAllOrdersV1RespItemWithDefaults instantiates a new UfuturesGetAllOrdersV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetAllOrdersV1RespItemWithDefaults() *UfuturesGetAllOrdersV1RespItem {
	this := UfuturesGetAllOrdersV1RespItem{}
	return &this
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetActivatePrice() string {
	if o == nil || IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetActivatePriceOk() (*string, bool) {
	if o == nil || IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasActivatePrice() bool {
	if o != nil && !IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *UfuturesGetAllOrdersV1RespItem) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetAvgPrice() string {
	if o == nil || IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetAvgPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasAvgPrice() bool {
	if o != nil && !IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *UfuturesGetAllOrdersV1RespItem) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *UfuturesGetAllOrdersV1RespItem) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetClosePosition() bool {
	if o == nil || IsNil(o.ClosePosition) {
		var ret bool
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetClosePositionOk() (*bool, bool) {
	if o == nil || IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasClosePosition() bool {
	if o != nil && !IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given bool and assigns it to the ClosePosition field.
func (o *UfuturesGetAllOrdersV1RespItem) SetClosePosition(v bool) {
	o.ClosePosition = &v
}

// GetCumQuote returns the CumQuote field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetCumQuote() string {
	if o == nil || IsNil(o.CumQuote) {
		var ret string
		return ret
	}
	return *o.CumQuote
}

// GetCumQuoteOk returns a tuple with the CumQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetCumQuoteOk() (*string, bool) {
	if o == nil || IsNil(o.CumQuote) {
		return nil, false
	}
	return o.CumQuote, true
}

// HasCumQuote returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasCumQuote() bool {
	if o != nil && !IsNil(o.CumQuote) {
		return true
	}

	return false
}

// SetCumQuote gets a reference to the given string and assigns it to the CumQuote field.
func (o *UfuturesGetAllOrdersV1RespItem) SetCumQuote(v string) {
	o.CumQuote = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *UfuturesGetAllOrdersV1RespItem) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetGoodTillDate returns the GoodTillDate field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetGoodTillDate() int64 {
	if o == nil || IsNil(o.GoodTillDate) {
		var ret int64
		return ret
	}
	return *o.GoodTillDate
}

// GetGoodTillDateOk returns a tuple with the GoodTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetGoodTillDateOk() (*int64, bool) {
	if o == nil || IsNil(o.GoodTillDate) {
		return nil, false
	}
	return o.GoodTillDate, true
}

// HasGoodTillDate returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasGoodTillDate() bool {
	if o != nil && !IsNil(o.GoodTillDate) {
		return true
	}

	return false
}

// SetGoodTillDate gets a reference to the given int64 and assigns it to the GoodTillDate field.
func (o *UfuturesGetAllOrdersV1RespItem) SetGoodTillDate(v int64) {
	o.GoodTillDate = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *UfuturesGetAllOrdersV1RespItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetOrigQty() string {
	if o == nil || IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetOrigQtyOk() (*string, bool) {
	if o == nil || IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasOrigQty() bool {
	if o != nil && !IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *UfuturesGetAllOrdersV1RespItem) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetOrigType returns the OrigType field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetOrigType() string {
	if o == nil || IsNil(o.OrigType) {
		var ret string
		return ret
	}
	return *o.OrigType
}

// GetOrigTypeOk returns a tuple with the OrigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetOrigTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrigType) {
		return nil, false
	}
	return o.OrigType, true
}

// HasOrigType returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasOrigType() bool {
	if o != nil && !IsNil(o.OrigType) {
		return true
	}

	return false
}

// SetOrigType gets a reference to the given string and assigns it to the OrigType field.
func (o *UfuturesGetAllOrdersV1RespItem) SetOrigType(v string) {
	o.OrigType = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *UfuturesGetAllOrdersV1RespItem) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *UfuturesGetAllOrdersV1RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetPriceMatch() string {
	if o == nil || IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetPriceMatchOk() (*string, bool) {
	if o == nil || IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasPriceMatch() bool {
	if o != nil && !IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *UfuturesGetAllOrdersV1RespItem) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetPriceProtect() bool {
	if o == nil || IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetPriceProtectOk() (*bool, bool) {
	if o == nil || IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasPriceProtect() bool {
	if o != nil && !IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *UfuturesGetAllOrdersV1RespItem) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetPriceRate() string {
	if o == nil || IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetPriceRateOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasPriceRate() bool {
	if o != nil && !IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *UfuturesGetAllOrdersV1RespItem) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *UfuturesGetAllOrdersV1RespItem) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *UfuturesGetAllOrdersV1RespItem) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *UfuturesGetAllOrdersV1RespItem) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UfuturesGetAllOrdersV1RespItem) SetStatus(v string) {
	o.Status = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *UfuturesGetAllOrdersV1RespItem) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UfuturesGetAllOrdersV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *UfuturesGetAllOrdersV1RespItem) SetTime(v int64) {
	o.Time = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *UfuturesGetAllOrdersV1RespItem) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UfuturesGetAllOrdersV1RespItem) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *UfuturesGetAllOrdersV1RespItem) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *UfuturesGetAllOrdersV1RespItem) GetWorkingType() string {
	if o == nil || IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetAllOrdersV1RespItem) GetWorkingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *UfuturesGetAllOrdersV1RespItem) HasWorkingType() bool {
	if o != nil && !IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *UfuturesGetAllOrdersV1RespItem) SetWorkingType(v string) {
	o.WorkingType = &v
}

func (o UfuturesGetAllOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetAllOrdersV1RespItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CumQuote) {
		toSerialize["cumQuote"] = o.CumQuote
	}
	if !IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !IsNil(o.GoodTillDate) {
		toSerialize["goodTillDate"] = o.GoodTillDate
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

type NullableUfuturesGetAllOrdersV1RespItem struct {
	value *UfuturesGetAllOrdersV1RespItem
	isSet bool
}

func (v NullableUfuturesGetAllOrdersV1RespItem) Get() *UfuturesGetAllOrdersV1RespItem {
	return v.value
}

func (v *NullableUfuturesGetAllOrdersV1RespItem) Set(val *UfuturesGetAllOrdersV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetAllOrdersV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetAllOrdersV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetAllOrdersV1RespItem(val *UfuturesGetAllOrdersV1RespItem) *NullableUfuturesGetAllOrdersV1RespItem {
	return &NullableUfuturesGetAllOrdersV1RespItem{value: val, isSet: true}
}

func (v NullableUfuturesGetAllOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetAllOrdersV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


