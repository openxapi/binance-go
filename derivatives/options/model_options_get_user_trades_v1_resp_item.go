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

// checks if the OptionsGetUserTradesV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsGetUserTradesV1RespItem{}

// OptionsGetUserTradesV1RespItem struct for OptionsGetUserTradesV1RespItem
type OptionsGetUserTradesV1RespItem struct {
	Fee *string `json:"fee,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Liquidity *string `json:"liquidity,omitempty"`
	OptionSide *string `json:"optionSide,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	Price *string `json:"price,omitempty"`
	PriceScale *int32 `json:"priceScale,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	QuantityScale *int32 `json:"quantityScale,omitempty"`
	QuoteAsset *string `json:"quoteAsset,omitempty"`
	RealizedProfit *string `json:"realizedProfit,omitempty"`
	Side *string `json:"side,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Time *int64 `json:"time,omitempty"`
	TradeId *int64 `json:"tradeId,omitempty"`
	Type *string `json:"type,omitempty"`
	Volatility *string `json:"volatility,omitempty"`
}

// NewOptionsGetUserTradesV1RespItem instantiates a new OptionsGetUserTradesV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsGetUserTradesV1RespItem() *OptionsGetUserTradesV1RespItem {
	this := OptionsGetUserTradesV1RespItem{}
	return &this
}

// NewOptionsGetUserTradesV1RespItemWithDefaults instantiates a new OptionsGetUserTradesV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsGetUserTradesV1RespItemWithDefaults() *OptionsGetUserTradesV1RespItem {
	this := OptionsGetUserTradesV1RespItem{}
	return &this
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *OptionsGetUserTradesV1RespItem) SetFee(v string) {
	o.Fee = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OptionsGetUserTradesV1RespItem) SetId(v int32) {
	o.Id = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetLiquidity() string {
	if o == nil || IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetLiquidityOk() (*string, bool) {
	if o == nil || IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasLiquidity() bool {
	if o != nil && !IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *OptionsGetUserTradesV1RespItem) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetOptionSide() string {
	if o == nil || IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetOptionSideOk() (*string, bool) {
	if o == nil || IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasOptionSide() bool {
	if o != nil && !IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *OptionsGetUserTradesV1RespItem) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *OptionsGetUserTradesV1RespItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *OptionsGetUserTradesV1RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetPriceScale() int32 {
	if o == nil || IsNil(o.PriceScale) {
		var ret int32
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetPriceScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasPriceScale() bool {
	if o != nil && !IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int32 and assigns it to the PriceScale field.
func (o *OptionsGetUserTradesV1RespItem) SetPriceScale(v int32) {
	o.PriceScale = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *OptionsGetUserTradesV1RespItem) SetQuantity(v string) {
	o.Quantity = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetQuantityScale() int32 {
	if o == nil || IsNil(o.QuantityScale) {
		var ret int32
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetQuantityScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasQuantityScale() bool {
	if o != nil && !IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int32 and assigns it to the QuantityScale field.
func (o *OptionsGetUserTradesV1RespItem) SetQuantityScale(v int32) {
	o.QuantityScale = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetQuoteAsset() string {
	if o == nil || IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetQuoteAssetOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasQuoteAsset() bool {
	if o != nil && !IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *OptionsGetUserTradesV1RespItem) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetRealizedProfit returns the RealizedProfit field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetRealizedProfit() string {
	if o == nil || IsNil(o.RealizedProfit) {
		var ret string
		return ret
	}
	return *o.RealizedProfit
}

// GetRealizedProfitOk returns a tuple with the RealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetRealizedProfitOk() (*string, bool) {
	if o == nil || IsNil(o.RealizedProfit) {
		return nil, false
	}
	return o.RealizedProfit, true
}

// HasRealizedProfit returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasRealizedProfit() bool {
	if o != nil && !IsNil(o.RealizedProfit) {
		return true
	}

	return false
}

// SetRealizedProfit gets a reference to the given string and assigns it to the RealizedProfit field.
func (o *OptionsGetUserTradesV1RespItem) SetRealizedProfit(v string) {
	o.RealizedProfit = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OptionsGetUserTradesV1RespItem) SetSide(v string) {
	o.Side = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionsGetUserTradesV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *OptionsGetUserTradesV1RespItem) SetTime(v int64) {
	o.Time = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetTradeId() int64 {
	if o == nil || IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetTradeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasTradeId() bool {
	if o != nil && !IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *OptionsGetUserTradesV1RespItem) SetTradeId(v int64) {
	o.TradeId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OptionsGetUserTradesV1RespItem) SetType(v string) {
	o.Type = &v
}

// GetVolatility returns the Volatility field value if set, zero value otherwise.
func (o *OptionsGetUserTradesV1RespItem) GetVolatility() string {
	if o == nil || IsNil(o.Volatility) {
		var ret string
		return ret
	}
	return *o.Volatility
}

// GetVolatilityOk returns a tuple with the Volatility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetUserTradesV1RespItem) GetVolatilityOk() (*string, bool) {
	if o == nil || IsNil(o.Volatility) {
		return nil, false
	}
	return o.Volatility, true
}

// HasVolatility returns a boolean if a field has been set.
func (o *OptionsGetUserTradesV1RespItem) HasVolatility() bool {
	if o != nil && !IsNil(o.Volatility) {
		return true
	}

	return false
}

// SetVolatility gets a reference to the given string and assigns it to the Volatility field.
func (o *OptionsGetUserTradesV1RespItem) SetVolatility(v string) {
	o.Volatility = &v
}

func (o OptionsGetUserTradesV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsGetUserTradesV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !IsNil(o.OptionSide) {
		toSerialize["optionSide"] = o.OptionSide
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
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
	if !IsNil(o.RealizedProfit) {
		toSerialize["realizedProfit"] = o.RealizedProfit
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Volatility) {
		toSerialize["volatility"] = o.Volatility
	}
	return toSerialize, nil
}

type NullableOptionsGetUserTradesV1RespItem struct {
	value *OptionsGetUserTradesV1RespItem
	isSet bool
}

func (v NullableOptionsGetUserTradesV1RespItem) Get() *OptionsGetUserTradesV1RespItem {
	return v.value
}

func (v *NullableOptionsGetUserTradesV1RespItem) Set(val *OptionsGetUserTradesV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsGetUserTradesV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsGetUserTradesV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsGetUserTradesV1RespItem(val *OptionsGetUserTradesV1RespItem) *NullableOptionsGetUserTradesV1RespItem {
	return &NullableOptionsGetUserTradesV1RespItem{value: val, isSet: true}
}

func (v NullableOptionsGetUserTradesV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsGetUserTradesV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


