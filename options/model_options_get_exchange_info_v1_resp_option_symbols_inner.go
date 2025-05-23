/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the OptionsGetExchangeInfoV1RespOptionSymbolsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsGetExchangeInfoV1RespOptionSymbolsInner{}

// OptionsGetExchangeInfoV1RespOptionSymbolsInner struct for OptionsGetExchangeInfoV1RespOptionSymbolsInner
type OptionsGetExchangeInfoV1RespOptionSymbolsInner struct {
	ExpiryDate *int32 `json:"expiryDate,omitempty"`
	Filters []OptionsSymbolFilter `json:"filters,omitempty"`
	InitialMargin *string `json:"initialMargin,omitempty"`
	MaintenanceMargin *string `json:"maintenanceMargin,omitempty"`
	MakerFeeRate *string `json:"makerFeeRate,omitempty"`
	MaxQty *string `json:"maxQty,omitempty"`
	MinInitialMargin *string `json:"minInitialMargin,omitempty"`
	MinMaintenanceMargin *string `json:"minMaintenanceMargin,omitempty"`
	MinQty *string `json:"minQty,omitempty"`
	PriceScale *int32 `json:"priceScale,omitempty"`
	QuantityScale *int32 `json:"quantityScale,omitempty"`
	QuoteAsset *string `json:"quoteAsset,omitempty"`
	Side *string `json:"side,omitempty"`
	StrikePrice *string `json:"strikePrice,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TakerFeeRate *string `json:"takerFeeRate,omitempty"`
	Underlying *string `json:"underlying,omitempty"`
	Unit *int32 `json:"unit,omitempty"`
}

// NewOptionsGetExchangeInfoV1RespOptionSymbolsInner instantiates a new OptionsGetExchangeInfoV1RespOptionSymbolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsGetExchangeInfoV1RespOptionSymbolsInner() *OptionsGetExchangeInfoV1RespOptionSymbolsInner {
	this := OptionsGetExchangeInfoV1RespOptionSymbolsInner{}
	return &this
}

// NewOptionsGetExchangeInfoV1RespOptionSymbolsInnerWithDefaults instantiates a new OptionsGetExchangeInfoV1RespOptionSymbolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsGetExchangeInfoV1RespOptionSymbolsInnerWithDefaults() *OptionsGetExchangeInfoV1RespOptionSymbolsInner {
	this := OptionsGetExchangeInfoV1RespOptionSymbolsInner{}
	return &this
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetExpiryDate() int32 {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret int32
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetExpiryDateOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given int32 and assigns it to the ExpiryDate field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetExpiryDate(v int32) {
	o.ExpiryDate = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetFilters() []OptionsSymbolFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []OptionsSymbolFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetFiltersOk() ([]OptionsSymbolFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []OptionsSymbolFilter and assigns it to the Filters field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetFilters(v []OptionsSymbolFilter) {
	o.Filters = v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetInitialMargin() string {
	if o == nil || IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasInitialMargin() bool {
	if o != nil && !IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetMaintenanceMargin returns the MaintenanceMargin field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaintenanceMargin() string {
	if o == nil || IsNil(o.MaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.MaintenanceMargin
}

// GetMaintenanceMarginOk returns a tuple with the MaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaintenanceMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MaintenanceMargin) {
		return nil, false
	}
	return o.MaintenanceMargin, true
}

// HasMaintenanceMargin returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMaintenanceMargin() bool {
	if o != nil && !IsNil(o.MaintenanceMargin) {
		return true
	}

	return false
}

// SetMaintenanceMargin gets a reference to the given string and assigns it to the MaintenanceMargin field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMaintenanceMargin(v string) {
	o.MaintenanceMargin = &v
}

// GetMakerFeeRate returns the MakerFeeRate field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMakerFeeRate() string {
	if o == nil || IsNil(o.MakerFeeRate) {
		var ret string
		return ret
	}
	return *o.MakerFeeRate
}

// GetMakerFeeRateOk returns a tuple with the MakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMakerFeeRateOk() (*string, bool) {
	if o == nil || IsNil(o.MakerFeeRate) {
		return nil, false
	}
	return o.MakerFeeRate, true
}

// HasMakerFeeRate returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMakerFeeRate() bool {
	if o != nil && !IsNil(o.MakerFeeRate) {
		return true
	}

	return false
}

// SetMakerFeeRate gets a reference to the given string and assigns it to the MakerFeeRate field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMakerFeeRate(v string) {
	o.MakerFeeRate = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaxQty() string {
	if o == nil || IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaxQtyOk() (*string, bool) {
	if o == nil || IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMaxQty() bool {
	if o != nil && !IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetMinInitialMargin returns the MinInitialMargin field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinInitialMargin() string {
	if o == nil || IsNil(o.MinInitialMargin) {
		var ret string
		return ret
	}
	return *o.MinInitialMargin
}

// GetMinInitialMarginOk returns a tuple with the MinInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MinInitialMargin) {
		return nil, false
	}
	return o.MinInitialMargin, true
}

// HasMinInitialMargin returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMinInitialMargin() bool {
	if o != nil && !IsNil(o.MinInitialMargin) {
		return true
	}

	return false
}

// SetMinInitialMargin gets a reference to the given string and assigns it to the MinInitialMargin field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMinInitialMargin(v string) {
	o.MinInitialMargin = &v
}

// GetMinMaintenanceMargin returns the MinMaintenanceMargin field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinMaintenanceMargin() string {
	if o == nil || IsNil(o.MinMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.MinMaintenanceMargin
}

// GetMinMaintenanceMarginOk returns a tuple with the MinMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinMaintenanceMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MinMaintenanceMargin) {
		return nil, false
	}
	return o.MinMaintenanceMargin, true
}

// HasMinMaintenanceMargin returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMinMaintenanceMargin() bool {
	if o != nil && !IsNil(o.MinMaintenanceMargin) {
		return true
	}

	return false
}

// SetMinMaintenanceMargin gets a reference to the given string and assigns it to the MinMaintenanceMargin field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMinMaintenanceMargin(v string) {
	o.MinMaintenanceMargin = &v
}

// GetMinQty returns the MinQty field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinQty() string {
	if o == nil || IsNil(o.MinQty) {
		var ret string
		return ret
	}
	return *o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinQtyOk() (*string, bool) {
	if o == nil || IsNil(o.MinQty) {
		return nil, false
	}
	return o.MinQty, true
}

// HasMinQty returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMinQty() bool {
	if o != nil && !IsNil(o.MinQty) {
		return true
	}

	return false
}

// SetMinQty gets a reference to the given string and assigns it to the MinQty field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMinQty(v string) {
	o.MinQty = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetPriceScale() int32 {
	if o == nil || IsNil(o.PriceScale) {
		var ret int32
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetPriceScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasPriceScale() bool {
	if o != nil && !IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int32 and assigns it to the PriceScale field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetPriceScale(v int32) {
	o.PriceScale = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuantityScale() int32 {
	if o == nil || IsNil(o.QuantityScale) {
		var ret int32
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuantityScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasQuantityScale() bool {
	if o != nil && !IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int32 and assigns it to the QuantityScale field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetQuantityScale(v int32) {
	o.QuantityScale = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuoteAsset() string {
	if o == nil || IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasQuoteAsset() bool {
	if o != nil && !IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetSide(v string) {
	o.Side = &v
}

// GetStrikePrice returns the StrikePrice field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetStrikePrice() string {
	if o == nil || IsNil(o.StrikePrice) {
		var ret string
		return ret
	}
	return *o.StrikePrice
}

// GetStrikePriceOk returns a tuple with the StrikePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetStrikePriceOk() (*string, bool) {
	if o == nil || IsNil(o.StrikePrice) {
		return nil, false
	}
	return o.StrikePrice, true
}

// HasStrikePrice returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasStrikePrice() bool {
	if o != nil && !IsNil(o.StrikePrice) {
		return true
	}

	return false
}

// SetStrikePrice gets a reference to the given string and assigns it to the StrikePrice field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetStrikePrice(v string) {
	o.StrikePrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTakerFeeRate returns the TakerFeeRate field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetTakerFeeRate() string {
	if o == nil || IsNil(o.TakerFeeRate) {
		var ret string
		return ret
	}
	return *o.TakerFeeRate
}

// GetTakerFeeRateOk returns a tuple with the TakerFeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetTakerFeeRateOk() (*string, bool) {
	if o == nil || IsNil(o.TakerFeeRate) {
		return nil, false
	}
	return o.TakerFeeRate, true
}

// HasTakerFeeRate returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasTakerFeeRate() bool {
	if o != nil && !IsNil(o.TakerFeeRate) {
		return true
	}

	return false
}

// SetTakerFeeRate gets a reference to the given string and assigns it to the TakerFeeRate field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetTakerFeeRate(v string) {
	o.TakerFeeRate = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnderlying() string {
	if o == nil || IsNil(o.Underlying) {
		var ret string
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnderlyingOk() (*string, bool) {
	if o == nil || IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasUnderlying() bool {
	if o != nil && !IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given string and assigns it to the Underlying field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetUnderlying(v string) {
	o.Underlying = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnit() int32 {
	if o == nil || IsNil(o.Unit) {
		var ret int32
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given int32 and assigns it to the Unit field.
func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetUnit(v int32) {
	o.Unit = &v
}

func (o OptionsGetExchangeInfoV1RespOptionSymbolsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsGetExchangeInfoV1RespOptionSymbolsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !IsNil(o.MaintenanceMargin) {
		toSerialize["maintenanceMargin"] = o.MaintenanceMargin
	}
	if !IsNil(o.MakerFeeRate) {
		toSerialize["makerFeeRate"] = o.MakerFeeRate
	}
	if !IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !IsNil(o.MinInitialMargin) {
		toSerialize["minInitialMargin"] = o.MinInitialMargin
	}
	if !IsNil(o.MinMaintenanceMargin) {
		toSerialize["minMaintenanceMargin"] = o.MinMaintenanceMargin
	}
	if !IsNil(o.MinQty) {
		toSerialize["minQty"] = o.MinQty
	}
	if !IsNil(o.PriceScale) {
		toSerialize["priceScale"] = o.PriceScale
	}
	if !IsNil(o.QuantityScale) {
		toSerialize["quantityScale"] = o.QuantityScale
	}
	if !IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.StrikePrice) {
		toSerialize["strikePrice"] = o.StrikePrice
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.TakerFeeRate) {
		toSerialize["takerFeeRate"] = o.TakerFeeRate
	}
	if !IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner struct {
	value *OptionsGetExchangeInfoV1RespOptionSymbolsInner
	isSet bool
}

func (v NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner) Get() *OptionsGetExchangeInfoV1RespOptionSymbolsInner {
	return v.value
}

func (v *NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner) Set(val *OptionsGetExchangeInfoV1RespOptionSymbolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsGetExchangeInfoV1RespOptionSymbolsInner(val *OptionsGetExchangeInfoV1RespOptionSymbolsInner) *NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner {
	return &NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner{value: val, isSet: true}
}

func (v NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsGetExchangeInfoV1RespOptionSymbolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


