/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the OptionsDeleteOrderV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsDeleteOrderV1Resp{}

// OptionsDeleteOrderV1Resp struct for OptionsDeleteOrderV1Resp
type OptionsDeleteOrderV1Resp struct {
	AvgPrice *string `json:"avgPrice,omitempty"`
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CreateDate *int32 `json:"createDate,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	Fee *string `json:"fee,omitempty"`
	Mmp *bool `json:"mmp,omitempty"`
	OptionSide *string `json:"optionSide,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	PostOnly *bool `json:"postOnly,omitempty"`
	Price *string `json:"price,omitempty"`
	PriceScale *int32 `json:"priceScale,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	QuantityScale *int32 `json:"quantityScale,omitempty"`
	QuoteAsset *string `json:"quoteAsset,omitempty"`
	ReduceOnly *bool `json:"reduceOnly,omitempty"`
	Side *string `json:"side,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewOptionsDeleteOrderV1Resp instantiates a new OptionsDeleteOrderV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsDeleteOrderV1Resp() *OptionsDeleteOrderV1Resp {
	this := OptionsDeleteOrderV1Resp{}
	return &this
}

// NewOptionsDeleteOrderV1RespWithDefaults instantiates a new OptionsDeleteOrderV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsDeleteOrderV1RespWithDefaults() *OptionsDeleteOrderV1Resp {
	this := OptionsDeleteOrderV1Resp{}
	return &this
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetAvgPrice() string {
	if o == nil || IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetAvgPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasAvgPrice() bool {
	if o != nil && !IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *OptionsDeleteOrderV1Resp) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *OptionsDeleteOrderV1Resp) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetCreateDate() int32 {
	if o == nil || IsNil(o.CreateDate) {
		var ret int32
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetCreateDateOk() (*int32, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasCreateDate() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given int32 and assigns it to the CreateDate field.
func (o *OptionsDeleteOrderV1Resp) SetCreateDate(v int32) {
	o.CreateDate = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *OptionsDeleteOrderV1Resp) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *OptionsDeleteOrderV1Resp) SetFee(v string) {
	o.Fee = &v
}

// GetMmp returns the Mmp field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetMmp() bool {
	if o == nil || IsNil(o.Mmp) {
		var ret bool
		return ret
	}
	return *o.Mmp
}

// GetMmpOk returns a tuple with the Mmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetMmpOk() (*bool, bool) {
	if o == nil || IsNil(o.Mmp) {
		return nil, false
	}
	return o.Mmp, true
}

// HasMmp returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasMmp() bool {
	if o != nil && !IsNil(o.Mmp) {
		return true
	}

	return false
}

// SetMmp gets a reference to the given bool and assigns it to the Mmp field.
func (o *OptionsDeleteOrderV1Resp) SetMmp(v bool) {
	o.Mmp = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetOptionSide() string {
	if o == nil || IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetOptionSideOk() (*string, bool) {
	if o == nil || IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasOptionSide() bool {
	if o != nil && !IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *OptionsDeleteOrderV1Resp) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OptionsDeleteOrderV1Resp) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetPostOnly() bool {
	if o == nil || IsNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetPostOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.PostOnly) {
		return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasPostOnly() bool {
	if o != nil && !IsNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *OptionsDeleteOrderV1Resp) SetPostOnly(v bool) {
	o.PostOnly = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OptionsDeleteOrderV1Resp) SetPrice(v string) {
	o.Price = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetPriceScale() int32 {
	if o == nil || IsNil(o.PriceScale) {
		var ret int32
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetPriceScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasPriceScale() bool {
	if o != nil && !IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int32 and assigns it to the PriceScale field.
func (o *OptionsDeleteOrderV1Resp) SetPriceScale(v int32) {
	o.PriceScale = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *OptionsDeleteOrderV1Resp) SetQuantity(v string) {
	o.Quantity = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetQuantityScale() int32 {
	if o == nil || IsNil(o.QuantityScale) {
		var ret int32
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetQuantityScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasQuantityScale() bool {
	if o != nil && !IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int32 and assigns it to the QuantityScale field.
func (o *OptionsDeleteOrderV1Resp) SetQuantityScale(v int32) {
	o.QuantityScale = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetQuoteAsset() string {
	if o == nil || IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetQuoteAssetOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasQuoteAsset() bool {
	if o != nil && !IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *OptionsDeleteOrderV1Resp) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *OptionsDeleteOrderV1Resp) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OptionsDeleteOrderV1Resp) SetSide(v string) {
	o.Side = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *OptionsDeleteOrderV1Resp) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OptionsDeleteOrderV1Resp) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionsDeleteOrderV1Resp) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *OptionsDeleteOrderV1Resp) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OptionsDeleteOrderV1Resp) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *OptionsDeleteOrderV1Resp) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsDeleteOrderV1Resp) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *OptionsDeleteOrderV1Resp) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *OptionsDeleteOrderV1Resp) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o OptionsDeleteOrderV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsDeleteOrderV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
	}
	if !IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Mmp) {
		toSerialize["mmp"] = o.Mmp
	}
	if !IsNil(o.OptionSide) {
		toSerialize["optionSide"] = o.OptionSide
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.PostOnly) {
		toSerialize["postOnly"] = o.PostOnly
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceScale) {
		toSerialize["priceScale"] = o.PriceScale
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.QuantityScale) {
		toSerialize["quantityScale"] = o.QuantityScale
	}
	if !IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
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
	return toSerialize, nil
}

type NullableOptionsDeleteOrderV1Resp struct {
	value *OptionsDeleteOrderV1Resp
	isSet bool
}

func (v NullableOptionsDeleteOrderV1Resp) Get() *OptionsDeleteOrderV1Resp {
	return v.value
}

func (v *NullableOptionsDeleteOrderV1Resp) Set(val *OptionsDeleteOrderV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsDeleteOrderV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsDeleteOrderV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsDeleteOrderV1Resp(val *OptionsDeleteOrderV1Resp) *NullableOptionsDeleteOrderV1Resp {
	return &NullableOptionsDeleteOrderV1Resp{value: val, isSet: true}
}

func (v NullableOptionsDeleteOrderV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsDeleteOrderV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


