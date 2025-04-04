/*
Binance COIN-M Futures API

OpenAPI specification for Binance exchange - Cmfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CmfuturesCreateBatchOrderV1ReqBatchOrdersItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CmfuturesCreateBatchOrderV1ReqBatchOrdersItem{}

// CmfuturesCreateBatchOrderV1ReqBatchOrdersItem struct for CmfuturesCreateBatchOrderV1ReqBatchOrdersItem
type CmfuturesCreateBatchOrderV1ReqBatchOrdersItem struct {
	ActivationPrice *string `json:"activationPrice,omitempty"`
	CallbackRate *string `json:"callbackRate,omitempty"`
	ClosePosition *string `json:"closePosition,omitempty"`
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	NewOrderRespType *string `json:"newOrderRespType,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Price *string `json:"price,omitempty"`
	PriceMatch *string `json:"priceMatch,omitempty"`
	PriceProtect *string `json:"priceProtect,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	ReduceOnly *string `json:"reduceOnly,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Side string `json:"side"`
	StopPrice *string `json:"stopPrice,omitempty"`
	Symbol string `json:"symbol"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Type string `json:"type"`
	WorkingType *string `json:"workingType,omitempty"`
}

type _CmfuturesCreateBatchOrderV1ReqBatchOrdersItem CmfuturesCreateBatchOrderV1ReqBatchOrdersItem

// NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItem instantiates a new CmfuturesCreateBatchOrderV1ReqBatchOrdersItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItem(side string, symbol string, type_ string) *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem {
	this := CmfuturesCreateBatchOrderV1ReqBatchOrdersItem{}
	var activationPrice string = ""
	this.ActivationPrice = &activationPrice
	var callbackRate string = ""
	this.CallbackRate = &callbackRate
	var closePosition string = ""
	this.ClosePosition = &closePosition
	var newClientOrderId string = ""
	this.NewClientOrderId = &newClientOrderId
	var newOrderRespType string = ""
	this.NewOrderRespType = &newOrderRespType
	var positionSide string = ""
	this.PositionSide = &positionSide
	var price string = ""
	this.Price = &price
	var priceMatch string = ""
	this.PriceMatch = &priceMatch
	var priceProtect string = ""
	this.PriceProtect = &priceProtect
	var quantity string = ""
	this.Quantity = &quantity
	var reduceOnly string = ""
	this.ReduceOnly = &reduceOnly
	var selfTradePreventionMode string = ""
	this.SelfTradePreventionMode = &selfTradePreventionMode
	this.Side = side
	var stopPrice string = ""
	this.StopPrice = &stopPrice
	this.Symbol = symbol
	var timeInForce string = ""
	this.TimeInForce = &timeInForce
	this.Type = type_
	var workingType string = ""
	this.WorkingType = &workingType
	return &this
}

// NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItemWithDefaults instantiates a new CmfuturesCreateBatchOrderV1ReqBatchOrdersItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmfuturesCreateBatchOrderV1ReqBatchOrdersItemWithDefaults() *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem {
	this := CmfuturesCreateBatchOrderV1ReqBatchOrdersItem{}
	var activationPrice string = ""
	this.ActivationPrice = &activationPrice
	var callbackRate string = ""
	this.CallbackRate = &callbackRate
	var closePosition string = ""
	this.ClosePosition = &closePosition
	var newClientOrderId string = ""
	this.NewClientOrderId = &newClientOrderId
	var newOrderRespType string = ""
	this.NewOrderRespType = &newOrderRespType
	var positionSide string = ""
	this.PositionSide = &positionSide
	var price string = ""
	this.Price = &price
	var priceMatch string = ""
	this.PriceMatch = &priceMatch
	var priceProtect string = ""
	this.PriceProtect = &priceProtect
	var quantity string = ""
	this.Quantity = &quantity
	var reduceOnly string = ""
	this.ReduceOnly = &reduceOnly
	var selfTradePreventionMode string = ""
	this.SelfTradePreventionMode = &selfTradePreventionMode
	var side string = ""
	this.Side = side
	var stopPrice string = ""
	this.StopPrice = &stopPrice
	var symbol string = ""
	this.Symbol = symbol
	var timeInForce string = ""
	this.TimeInForce = &timeInForce
	var type_ string = ""
	this.Type = type_
	var workingType string = ""
	this.WorkingType = &workingType
	return &this
}

// GetActivationPrice returns the ActivationPrice field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetActivationPrice() string {
	if o == nil || IsNil(o.ActivationPrice) {
		var ret string
		return ret
	}
	return *o.ActivationPrice
}

// GetActivationPriceOk returns a tuple with the ActivationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetActivationPriceOk() (*string, bool) {
	if o == nil || IsNil(o.ActivationPrice) {
		return nil, false
	}
	return o.ActivationPrice, true
}

// HasActivationPrice returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasActivationPrice() bool {
	if o != nil && !IsNil(o.ActivationPrice) {
		return true
	}

	return false
}

// SetActivationPrice gets a reference to the given string and assigns it to the ActivationPrice field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetActivationPrice(v string) {
	o.ActivationPrice = &v
}

// GetCallbackRate returns the CallbackRate field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetCallbackRate() string {
	if o == nil || IsNil(o.CallbackRate) {
		var ret string
		return ret
	}
	return *o.CallbackRate
}

// GetCallbackRateOk returns a tuple with the CallbackRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetCallbackRateOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackRate) {
		return nil, false
	}
	return o.CallbackRate, true
}

// HasCallbackRate returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasCallbackRate() bool {
	if o != nil && !IsNil(o.CallbackRate) {
		return true
	}

	return false
}

// SetCallbackRate gets a reference to the given string and assigns it to the CallbackRate field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetCallbackRate(v string) {
	o.CallbackRate = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetClosePosition() string {
	if o == nil || IsNil(o.ClosePosition) {
		var ret string
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetClosePositionOk() (*string, bool) {
	if o == nil || IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasClosePosition() bool {
	if o != nil && !IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given string and assigns it to the ClosePosition field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetClosePosition(v string) {
	o.ClosePosition = &v
}

// GetNewClientOrderId returns the NewClientOrderId field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewClientOrderId() string {
	if o == nil || IsNil(o.NewClientOrderId) {
		var ret string
		return ret
	}
	return *o.NewClientOrderId
}

// GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewClientOrderId) {
		return nil, false
	}
	return o.NewClientOrderId, true
}

// HasNewClientOrderId returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasNewClientOrderId() bool {
	if o != nil && !IsNil(o.NewClientOrderId) {
		return true
	}

	return false
}

// SetNewClientOrderId gets a reference to the given string and assigns it to the NewClientOrderId field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetNewClientOrderId(v string) {
	o.NewClientOrderId = &v
}

// GetNewOrderRespType returns the NewOrderRespType field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewOrderRespType() string {
	if o == nil || IsNil(o.NewOrderRespType) {
		var ret string
		return ret
	}
	return *o.NewOrderRespType
}

// GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetNewOrderRespTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NewOrderRespType) {
		return nil, false
	}
	return o.NewOrderRespType, true
}

// HasNewOrderRespType returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasNewOrderRespType() bool {
	if o != nil && !IsNil(o.NewOrderRespType) {
		return true
	}

	return false
}

// SetNewOrderRespType gets a reference to the given string and assigns it to the NewOrderRespType field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetNewOrderRespType(v string) {
	o.NewOrderRespType = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPrice(v string) {
	o.Price = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceMatch() string {
	if o == nil || IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceMatchOk() (*string, bool) {
	if o == nil || IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPriceMatch() bool {
	if o != nil && !IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceProtect() string {
	if o == nil || IsNil(o.PriceProtect) {
		var ret string
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetPriceProtectOk() (*string, bool) {
	if o == nil || IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasPriceProtect() bool {
	if o != nil && !IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given string and assigns it to the PriceProtect field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetPriceProtect(v string) {
	o.PriceProtect = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetQuantity(v string) {
	o.Quantity = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetReduceOnly() string {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret string
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetReduceOnlyOk() (*string, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given string and assigns it to the ReduceOnly field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetReduceOnly(v string) {
	o.ReduceOnly = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSelfTradePreventionMode() string {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasSelfTradePreventionMode() bool {
	if o != nil && !IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetSide returns the Side field value
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetSide(v string) {
	o.Side = v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetStopPrice() string {
	if o == nil || IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetStopPriceOk() (*string, bool) {
	if o == nil || IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasStopPrice() bool {
	if o != nil && !IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetType(v string) {
	o.Type = v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetWorkingType() string {
	if o == nil || IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) GetWorkingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) HasWorkingType() bool {
	if o != nil && !IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) SetWorkingType(v string) {
	o.WorkingType = &v
}

func (o CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivationPrice) {
		toSerialize["activationPrice"] = o.ActivationPrice
	}
	if !IsNil(o.CallbackRate) {
		toSerialize["callbackRate"] = o.CallbackRate
	}
	if !IsNil(o.ClosePosition) {
		toSerialize["closePosition"] = o.ClosePosition
	}
	if !IsNil(o.NewClientOrderId) {
		toSerialize["newClientOrderId"] = o.NewClientOrderId
	}
	if !IsNil(o.NewOrderRespType) {
		toSerialize["newOrderRespType"] = o.NewOrderRespType
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
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	toSerialize["side"] = o.Side
	if !IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	toSerialize["symbol"] = o.Symbol
	if !IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	return toSerialize, nil
}

func (o *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"side",
		"symbol",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCmfuturesCreateBatchOrderV1ReqBatchOrdersItem := _CmfuturesCreateBatchOrderV1ReqBatchOrdersItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCmfuturesCreateBatchOrderV1ReqBatchOrdersItem)

	if err != nil {
		return err
	}

	*o = CmfuturesCreateBatchOrderV1ReqBatchOrdersItem(varCmfuturesCreateBatchOrderV1ReqBatchOrdersItem)

	return err
}

type NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem struct {
	value *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem
	isSet bool
}

func (v NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem) Get() *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem {
	return v.value
}

func (v *NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem) Set(val *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem(val *CmfuturesCreateBatchOrderV1ReqBatchOrdersItem) *NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem {
	return &NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem{value: val, isSet: true}
}

func (v NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmfuturesCreateBatchOrderV1ReqBatchOrdersItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


