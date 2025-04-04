/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotSymbolFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotSymbolFilter{}

// SpotSymbolFilter struct for SpotSymbolFilter
type SpotSymbolFilter struct {
	ApplyMaxToMarket *bool `json:"applyMaxToMarket,omitempty"`
	ApplyMinToMarket *bool `json:"applyMinToMarket,omitempty"`
	ApplyToMarket *bool `json:"applyToMarket,omitempty"`
	AskMultiplierDown *string `json:"askMultiplierDown,omitempty"`
	AskMultiplierUp *string `json:"askMultiplierUp,omitempty"`
	AvgPriceMins *int32 `json:"avgPriceMins,omitempty"`
	BidMultiplierDown *string `json:"bidMultiplierDown,omitempty"`
	BidMultiplierUp *string `json:"bidMultiplierUp,omitempty"`
	FilterType *string `json:"filterType,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	MaxNotionalValue *string `json:"maxNotionalValue,omitempty"`
	MaxNumAlgoOrders *int32 `json:"maxNumAlgoOrders,omitempty"`
	MaxNumIcebergOrders *int32 `json:"maxNumIcebergOrders,omitempty"`
	MaxNumOrders *int32 `json:"maxNumOrders,omitempty"`
	MaxPosition *string `json:"maxPosition,omitempty"`
	MaxPrice *string `json:"maxPrice,omitempty"`
	MaxQty *string `json:"maxQty,omitempty"`
	MaxTrailingAboveDelta *int32 `json:"maxTrailingAboveDelta,omitempty"`
	MaxTrailingBelowDelta *int32 `json:"maxTrailingBelowDelta,omitempty"`
	MinNotionalValue *string `json:"minNotionalValue,omitempty"`
	MinPrice *string `json:"minPrice,omitempty"`
	MinQty *string `json:"minQty,omitempty"`
	MinTrailingAboveDelta *int32 `json:"minTrailingAboveDelta,omitempty"`
	MinTrailingBelowDelta *int32 `json:"minTrailingBelowDelta,omitempty"`
	MultiplierDown *string `json:"multiplierDown,omitempty"`
	MultiplierUp *string `json:"multiplierUp,omitempty"`
	StepSize *string `json:"stepSize,omitempty"`
	TickSize *string `json:"tickSize,omitempty"`
}

// NewSpotSymbolFilter instantiates a new SpotSymbolFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotSymbolFilter() *SpotSymbolFilter {
	this := SpotSymbolFilter{}
	return &this
}

// NewSpotSymbolFilterWithDefaults instantiates a new SpotSymbolFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotSymbolFilterWithDefaults() *SpotSymbolFilter {
	this := SpotSymbolFilter{}
	return &this
}

// GetApplyMaxToMarket returns the ApplyMaxToMarket field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetApplyMaxToMarket() bool {
	if o == nil || IsNil(o.ApplyMaxToMarket) {
		var ret bool
		return ret
	}
	return *o.ApplyMaxToMarket
}

// GetApplyMaxToMarketOk returns a tuple with the ApplyMaxToMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetApplyMaxToMarketOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyMaxToMarket) {
		return nil, false
	}
	return o.ApplyMaxToMarket, true
}

// HasApplyMaxToMarket returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasApplyMaxToMarket() bool {
	if o != nil && !IsNil(o.ApplyMaxToMarket) {
		return true
	}

	return false
}

// SetApplyMaxToMarket gets a reference to the given bool and assigns it to the ApplyMaxToMarket field.
func (o *SpotSymbolFilter) SetApplyMaxToMarket(v bool) {
	o.ApplyMaxToMarket = &v
}

// GetApplyMinToMarket returns the ApplyMinToMarket field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetApplyMinToMarket() bool {
	if o == nil || IsNil(o.ApplyMinToMarket) {
		var ret bool
		return ret
	}
	return *o.ApplyMinToMarket
}

// GetApplyMinToMarketOk returns a tuple with the ApplyMinToMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetApplyMinToMarketOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyMinToMarket) {
		return nil, false
	}
	return o.ApplyMinToMarket, true
}

// HasApplyMinToMarket returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasApplyMinToMarket() bool {
	if o != nil && !IsNil(o.ApplyMinToMarket) {
		return true
	}

	return false
}

// SetApplyMinToMarket gets a reference to the given bool and assigns it to the ApplyMinToMarket field.
func (o *SpotSymbolFilter) SetApplyMinToMarket(v bool) {
	o.ApplyMinToMarket = &v
}

// GetApplyToMarket returns the ApplyToMarket field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetApplyToMarket() bool {
	if o == nil || IsNil(o.ApplyToMarket) {
		var ret bool
		return ret
	}
	return *o.ApplyToMarket
}

// GetApplyToMarketOk returns a tuple with the ApplyToMarket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetApplyToMarketOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyToMarket) {
		return nil, false
	}
	return o.ApplyToMarket, true
}

// HasApplyToMarket returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasApplyToMarket() bool {
	if o != nil && !IsNil(o.ApplyToMarket) {
		return true
	}

	return false
}

// SetApplyToMarket gets a reference to the given bool and assigns it to the ApplyToMarket field.
func (o *SpotSymbolFilter) SetApplyToMarket(v bool) {
	o.ApplyToMarket = &v
}

// GetAskMultiplierDown returns the AskMultiplierDown field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetAskMultiplierDown() string {
	if o == nil || IsNil(o.AskMultiplierDown) {
		var ret string
		return ret
	}
	return *o.AskMultiplierDown
}

// GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetAskMultiplierDownOk() (*string, bool) {
	if o == nil || IsNil(o.AskMultiplierDown) {
		return nil, false
	}
	return o.AskMultiplierDown, true
}

// HasAskMultiplierDown returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasAskMultiplierDown() bool {
	if o != nil && !IsNil(o.AskMultiplierDown) {
		return true
	}

	return false
}

// SetAskMultiplierDown gets a reference to the given string and assigns it to the AskMultiplierDown field.
func (o *SpotSymbolFilter) SetAskMultiplierDown(v string) {
	o.AskMultiplierDown = &v
}

// GetAskMultiplierUp returns the AskMultiplierUp field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetAskMultiplierUp() string {
	if o == nil || IsNil(o.AskMultiplierUp) {
		var ret string
		return ret
	}
	return *o.AskMultiplierUp
}

// GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetAskMultiplierUpOk() (*string, bool) {
	if o == nil || IsNil(o.AskMultiplierUp) {
		return nil, false
	}
	return o.AskMultiplierUp, true
}

// HasAskMultiplierUp returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasAskMultiplierUp() bool {
	if o != nil && !IsNil(o.AskMultiplierUp) {
		return true
	}

	return false
}

// SetAskMultiplierUp gets a reference to the given string and assigns it to the AskMultiplierUp field.
func (o *SpotSymbolFilter) SetAskMultiplierUp(v string) {
	o.AskMultiplierUp = &v
}

// GetAvgPriceMins returns the AvgPriceMins field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetAvgPriceMins() int32 {
	if o == nil || IsNil(o.AvgPriceMins) {
		var ret int32
		return ret
	}
	return *o.AvgPriceMins
}

// GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetAvgPriceMinsOk() (*int32, bool) {
	if o == nil || IsNil(o.AvgPriceMins) {
		return nil, false
	}
	return o.AvgPriceMins, true
}

// HasAvgPriceMins returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasAvgPriceMins() bool {
	if o != nil && !IsNil(o.AvgPriceMins) {
		return true
	}

	return false
}

// SetAvgPriceMins gets a reference to the given int32 and assigns it to the AvgPriceMins field.
func (o *SpotSymbolFilter) SetAvgPriceMins(v int32) {
	o.AvgPriceMins = &v
}

// GetBidMultiplierDown returns the BidMultiplierDown field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetBidMultiplierDown() string {
	if o == nil || IsNil(o.BidMultiplierDown) {
		var ret string
		return ret
	}
	return *o.BidMultiplierDown
}

// GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetBidMultiplierDownOk() (*string, bool) {
	if o == nil || IsNil(o.BidMultiplierDown) {
		return nil, false
	}
	return o.BidMultiplierDown, true
}

// HasBidMultiplierDown returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasBidMultiplierDown() bool {
	if o != nil && !IsNil(o.BidMultiplierDown) {
		return true
	}

	return false
}

// SetBidMultiplierDown gets a reference to the given string and assigns it to the BidMultiplierDown field.
func (o *SpotSymbolFilter) SetBidMultiplierDown(v string) {
	o.BidMultiplierDown = &v
}

// GetBidMultiplierUp returns the BidMultiplierUp field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetBidMultiplierUp() string {
	if o == nil || IsNil(o.BidMultiplierUp) {
		var ret string
		return ret
	}
	return *o.BidMultiplierUp
}

// GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetBidMultiplierUpOk() (*string, bool) {
	if o == nil || IsNil(o.BidMultiplierUp) {
		return nil, false
	}
	return o.BidMultiplierUp, true
}

// HasBidMultiplierUp returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasBidMultiplierUp() bool {
	if o != nil && !IsNil(o.BidMultiplierUp) {
		return true
	}

	return false
}

// SetBidMultiplierUp gets a reference to the given string and assigns it to the BidMultiplierUp field.
func (o *SpotSymbolFilter) SetBidMultiplierUp(v string) {
	o.BidMultiplierUp = &v
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetFilterType() string {
	if o == nil || IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *SpotSymbolFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SpotSymbolFilter) SetLimit(v int32) {
	o.Limit = &v
}

// GetMaxNotionalValue returns the MaxNotionalValue field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxNotionalValue() string {
	if o == nil || IsNil(o.MaxNotionalValue) {
		var ret string
		return ret
	}
	return *o.MaxNotionalValue
}

// GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxNotionalValueOk() (*string, bool) {
	if o == nil || IsNil(o.MaxNotionalValue) {
		return nil, false
	}
	return o.MaxNotionalValue, true
}

// HasMaxNotionalValue returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxNotionalValue() bool {
	if o != nil && !IsNil(o.MaxNotionalValue) {
		return true
	}

	return false
}

// SetMaxNotionalValue gets a reference to the given string and assigns it to the MaxNotionalValue field.
func (o *SpotSymbolFilter) SetMaxNotionalValue(v string) {
	o.MaxNotionalValue = &v
}

// GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxNumAlgoOrders() int32 {
	if o == nil || IsNil(o.MaxNumAlgoOrders) {
		var ret int32
		return ret
	}
	return *o.MaxNumAlgoOrders
}

// GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxNumAlgoOrdersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumAlgoOrders) {
		return nil, false
	}
	return o.MaxNumAlgoOrders, true
}

// HasMaxNumAlgoOrders returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxNumAlgoOrders() bool {
	if o != nil && !IsNil(o.MaxNumAlgoOrders) {
		return true
	}

	return false
}

// SetMaxNumAlgoOrders gets a reference to the given int32 and assigns it to the MaxNumAlgoOrders field.
func (o *SpotSymbolFilter) SetMaxNumAlgoOrders(v int32) {
	o.MaxNumAlgoOrders = &v
}

// GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxNumIcebergOrders() int32 {
	if o == nil || IsNil(o.MaxNumIcebergOrders) {
		var ret int32
		return ret
	}
	return *o.MaxNumIcebergOrders
}

// GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxNumIcebergOrdersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumIcebergOrders) {
		return nil, false
	}
	return o.MaxNumIcebergOrders, true
}

// HasMaxNumIcebergOrders returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxNumIcebergOrders() bool {
	if o != nil && !IsNil(o.MaxNumIcebergOrders) {
		return true
	}

	return false
}

// SetMaxNumIcebergOrders gets a reference to the given int32 and assigns it to the MaxNumIcebergOrders field.
func (o *SpotSymbolFilter) SetMaxNumIcebergOrders(v int32) {
	o.MaxNumIcebergOrders = &v
}

// GetMaxNumOrders returns the MaxNumOrders field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxNumOrders() int32 {
	if o == nil || IsNil(o.MaxNumOrders) {
		var ret int32
		return ret
	}
	return *o.MaxNumOrders
}

// GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxNumOrdersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumOrders) {
		return nil, false
	}
	return o.MaxNumOrders, true
}

// HasMaxNumOrders returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxNumOrders() bool {
	if o != nil && !IsNil(o.MaxNumOrders) {
		return true
	}

	return false
}

// SetMaxNumOrders gets a reference to the given int32 and assigns it to the MaxNumOrders field.
func (o *SpotSymbolFilter) SetMaxNumOrders(v int32) {
	o.MaxNumOrders = &v
}

// GetMaxPosition returns the MaxPosition field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxPosition() string {
	if o == nil || IsNil(o.MaxPosition) {
		var ret string
		return ret
	}
	return *o.MaxPosition
}

// GetMaxPositionOk returns a tuple with the MaxPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxPositionOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPosition) {
		return nil, false
	}
	return o.MaxPosition, true
}

// HasMaxPosition returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxPosition() bool {
	if o != nil && !IsNil(o.MaxPosition) {
		return true
	}

	return false
}

// SetMaxPosition gets a reference to the given string and assigns it to the MaxPosition field.
func (o *SpotSymbolFilter) SetMaxPosition(v string) {
	o.MaxPosition = &v
}

// GetMaxPrice returns the MaxPrice field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxPrice() string {
	if o == nil || IsNil(o.MaxPrice) {
		var ret string
		return ret
	}
	return *o.MaxPrice
}

// GetMaxPriceOk returns a tuple with the MaxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxPriceOk() (*string, bool) {
	if o == nil || IsNil(o.MaxPrice) {
		return nil, false
	}
	return o.MaxPrice, true
}

// HasMaxPrice returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxPrice() bool {
	if o != nil && !IsNil(o.MaxPrice) {
		return true
	}

	return false
}

// SetMaxPrice gets a reference to the given string and assigns it to the MaxPrice field.
func (o *SpotSymbolFilter) SetMaxPrice(v string) {
	o.MaxPrice = &v
}

// GetMaxQty returns the MaxQty field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxQty() string {
	if o == nil || IsNil(o.MaxQty) {
		var ret string
		return ret
	}
	return *o.MaxQty
}

// GetMaxQtyOk returns a tuple with the MaxQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxQtyOk() (*string, bool) {
	if o == nil || IsNil(o.MaxQty) {
		return nil, false
	}
	return o.MaxQty, true
}

// HasMaxQty returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxQty() bool {
	if o != nil && !IsNil(o.MaxQty) {
		return true
	}

	return false
}

// SetMaxQty gets a reference to the given string and assigns it to the MaxQty field.
func (o *SpotSymbolFilter) SetMaxQty(v string) {
	o.MaxQty = &v
}

// GetMaxTrailingAboveDelta returns the MaxTrailingAboveDelta field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxTrailingAboveDelta() int32 {
	if o == nil || IsNil(o.MaxTrailingAboveDelta) {
		var ret int32
		return ret
	}
	return *o.MaxTrailingAboveDelta
}

// GetMaxTrailingAboveDeltaOk returns a tuple with the MaxTrailingAboveDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxTrailingAboveDeltaOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTrailingAboveDelta) {
		return nil, false
	}
	return o.MaxTrailingAboveDelta, true
}

// HasMaxTrailingAboveDelta returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxTrailingAboveDelta() bool {
	if o != nil && !IsNil(o.MaxTrailingAboveDelta) {
		return true
	}

	return false
}

// SetMaxTrailingAboveDelta gets a reference to the given int32 and assigns it to the MaxTrailingAboveDelta field.
func (o *SpotSymbolFilter) SetMaxTrailingAboveDelta(v int32) {
	o.MaxTrailingAboveDelta = &v
}

// GetMaxTrailingBelowDelta returns the MaxTrailingBelowDelta field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMaxTrailingBelowDelta() int32 {
	if o == nil || IsNil(o.MaxTrailingBelowDelta) {
		var ret int32
		return ret
	}
	return *o.MaxTrailingBelowDelta
}

// GetMaxTrailingBelowDeltaOk returns a tuple with the MaxTrailingBelowDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMaxTrailingBelowDeltaOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTrailingBelowDelta) {
		return nil, false
	}
	return o.MaxTrailingBelowDelta, true
}

// HasMaxTrailingBelowDelta returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMaxTrailingBelowDelta() bool {
	if o != nil && !IsNil(o.MaxTrailingBelowDelta) {
		return true
	}

	return false
}

// SetMaxTrailingBelowDelta gets a reference to the given int32 and assigns it to the MaxTrailingBelowDelta field.
func (o *SpotSymbolFilter) SetMaxTrailingBelowDelta(v int32) {
	o.MaxTrailingBelowDelta = &v
}

// GetMinNotionalValue returns the MinNotionalValue field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMinNotionalValue() string {
	if o == nil || IsNil(o.MinNotionalValue) {
		var ret string
		return ret
	}
	return *o.MinNotionalValue
}

// GetMinNotionalValueOk returns a tuple with the MinNotionalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMinNotionalValueOk() (*string, bool) {
	if o == nil || IsNil(o.MinNotionalValue) {
		return nil, false
	}
	return o.MinNotionalValue, true
}

// HasMinNotionalValue returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMinNotionalValue() bool {
	if o != nil && !IsNil(o.MinNotionalValue) {
		return true
	}

	return false
}

// SetMinNotionalValue gets a reference to the given string and assigns it to the MinNotionalValue field.
func (o *SpotSymbolFilter) SetMinNotionalValue(v string) {
	o.MinNotionalValue = &v
}

// GetMinPrice returns the MinPrice field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMinPrice() string {
	if o == nil || IsNil(o.MinPrice) {
		var ret string
		return ret
	}
	return *o.MinPrice
}

// GetMinPriceOk returns a tuple with the MinPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMinPriceOk() (*string, bool) {
	if o == nil || IsNil(o.MinPrice) {
		return nil, false
	}
	return o.MinPrice, true
}

// HasMinPrice returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMinPrice() bool {
	if o != nil && !IsNil(o.MinPrice) {
		return true
	}

	return false
}

// SetMinPrice gets a reference to the given string and assigns it to the MinPrice field.
func (o *SpotSymbolFilter) SetMinPrice(v string) {
	o.MinPrice = &v
}

// GetMinQty returns the MinQty field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMinQty() string {
	if o == nil || IsNil(o.MinQty) {
		var ret string
		return ret
	}
	return *o.MinQty
}

// GetMinQtyOk returns a tuple with the MinQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMinQtyOk() (*string, bool) {
	if o == nil || IsNil(o.MinQty) {
		return nil, false
	}
	return o.MinQty, true
}

// HasMinQty returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMinQty() bool {
	if o != nil && !IsNil(o.MinQty) {
		return true
	}

	return false
}

// SetMinQty gets a reference to the given string and assigns it to the MinQty field.
func (o *SpotSymbolFilter) SetMinQty(v string) {
	o.MinQty = &v
}

// GetMinTrailingAboveDelta returns the MinTrailingAboveDelta field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMinTrailingAboveDelta() int32 {
	if o == nil || IsNil(o.MinTrailingAboveDelta) {
		var ret int32
		return ret
	}
	return *o.MinTrailingAboveDelta
}

// GetMinTrailingAboveDeltaOk returns a tuple with the MinTrailingAboveDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMinTrailingAboveDeltaOk() (*int32, bool) {
	if o == nil || IsNil(o.MinTrailingAboveDelta) {
		return nil, false
	}
	return o.MinTrailingAboveDelta, true
}

// HasMinTrailingAboveDelta returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMinTrailingAboveDelta() bool {
	if o != nil && !IsNil(o.MinTrailingAboveDelta) {
		return true
	}

	return false
}

// SetMinTrailingAboveDelta gets a reference to the given int32 and assigns it to the MinTrailingAboveDelta field.
func (o *SpotSymbolFilter) SetMinTrailingAboveDelta(v int32) {
	o.MinTrailingAboveDelta = &v
}

// GetMinTrailingBelowDelta returns the MinTrailingBelowDelta field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMinTrailingBelowDelta() int32 {
	if o == nil || IsNil(o.MinTrailingBelowDelta) {
		var ret int32
		return ret
	}
	return *o.MinTrailingBelowDelta
}

// GetMinTrailingBelowDeltaOk returns a tuple with the MinTrailingBelowDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMinTrailingBelowDeltaOk() (*int32, bool) {
	if o == nil || IsNil(o.MinTrailingBelowDelta) {
		return nil, false
	}
	return o.MinTrailingBelowDelta, true
}

// HasMinTrailingBelowDelta returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMinTrailingBelowDelta() bool {
	if o != nil && !IsNil(o.MinTrailingBelowDelta) {
		return true
	}

	return false
}

// SetMinTrailingBelowDelta gets a reference to the given int32 and assigns it to the MinTrailingBelowDelta field.
func (o *SpotSymbolFilter) SetMinTrailingBelowDelta(v int32) {
	o.MinTrailingBelowDelta = &v
}

// GetMultiplierDown returns the MultiplierDown field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMultiplierDown() string {
	if o == nil || IsNil(o.MultiplierDown) {
		var ret string
		return ret
	}
	return *o.MultiplierDown
}

// GetMultiplierDownOk returns a tuple with the MultiplierDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMultiplierDownOk() (*string, bool) {
	if o == nil || IsNil(o.MultiplierDown) {
		return nil, false
	}
	return o.MultiplierDown, true
}

// HasMultiplierDown returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMultiplierDown() bool {
	if o != nil && !IsNil(o.MultiplierDown) {
		return true
	}

	return false
}

// SetMultiplierDown gets a reference to the given string and assigns it to the MultiplierDown field.
func (o *SpotSymbolFilter) SetMultiplierDown(v string) {
	o.MultiplierDown = &v
}

// GetMultiplierUp returns the MultiplierUp field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetMultiplierUp() string {
	if o == nil || IsNil(o.MultiplierUp) {
		var ret string
		return ret
	}
	return *o.MultiplierUp
}

// GetMultiplierUpOk returns a tuple with the MultiplierUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetMultiplierUpOk() (*string, bool) {
	if o == nil || IsNil(o.MultiplierUp) {
		return nil, false
	}
	return o.MultiplierUp, true
}

// HasMultiplierUp returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasMultiplierUp() bool {
	if o != nil && !IsNil(o.MultiplierUp) {
		return true
	}

	return false
}

// SetMultiplierUp gets a reference to the given string and assigns it to the MultiplierUp field.
func (o *SpotSymbolFilter) SetMultiplierUp(v string) {
	o.MultiplierUp = &v
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetStepSize() string {
	if o == nil || IsNil(o.StepSize) {
		var ret string
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetStepSizeOk() (*string, bool) {
	if o == nil || IsNil(o.StepSize) {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasStepSize() bool {
	if o != nil && !IsNil(o.StepSize) {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given string and assigns it to the StepSize field.
func (o *SpotSymbolFilter) SetStepSize(v string) {
	o.StepSize = &v
}

// GetTickSize returns the TickSize field value if set, zero value otherwise.
func (o *SpotSymbolFilter) GetTickSize() string {
	if o == nil || IsNil(o.TickSize) {
		var ret string
		return ret
	}
	return *o.TickSize
}

// GetTickSizeOk returns a tuple with the TickSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotSymbolFilter) GetTickSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TickSize) {
		return nil, false
	}
	return o.TickSize, true
}

// HasTickSize returns a boolean if a field has been set.
func (o *SpotSymbolFilter) HasTickSize() bool {
	if o != nil && !IsNil(o.TickSize) {
		return true
	}

	return false
}

// SetTickSize gets a reference to the given string and assigns it to the TickSize field.
func (o *SpotSymbolFilter) SetTickSize(v string) {
	o.TickSize = &v
}

func (o SpotSymbolFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotSymbolFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyMaxToMarket) {
		toSerialize["applyMaxToMarket"] = o.ApplyMaxToMarket
	}
	if !IsNil(o.ApplyMinToMarket) {
		toSerialize["applyMinToMarket"] = o.ApplyMinToMarket
	}
	if !IsNil(o.ApplyToMarket) {
		toSerialize["applyToMarket"] = o.ApplyToMarket
	}
	if !IsNil(o.AskMultiplierDown) {
		toSerialize["askMultiplierDown"] = o.AskMultiplierDown
	}
	if !IsNil(o.AskMultiplierUp) {
		toSerialize["askMultiplierUp"] = o.AskMultiplierUp
	}
	if !IsNil(o.AvgPriceMins) {
		toSerialize["avgPriceMins"] = o.AvgPriceMins
	}
	if !IsNil(o.BidMultiplierDown) {
		toSerialize["bidMultiplierDown"] = o.BidMultiplierDown
	}
	if !IsNil(o.BidMultiplierUp) {
		toSerialize["bidMultiplierUp"] = o.BidMultiplierUp
	}
	if !IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.MaxNotionalValue) {
		toSerialize["maxNotionalValue"] = o.MaxNotionalValue
	}
	if !IsNil(o.MaxNumAlgoOrders) {
		toSerialize["maxNumAlgoOrders"] = o.MaxNumAlgoOrders
	}
	if !IsNil(o.MaxNumIcebergOrders) {
		toSerialize["maxNumIcebergOrders"] = o.MaxNumIcebergOrders
	}
	if !IsNil(o.MaxNumOrders) {
		toSerialize["maxNumOrders"] = o.MaxNumOrders
	}
	if !IsNil(o.MaxPosition) {
		toSerialize["maxPosition"] = o.MaxPosition
	}
	if !IsNil(o.MaxPrice) {
		toSerialize["maxPrice"] = o.MaxPrice
	}
	if !IsNil(o.MaxQty) {
		toSerialize["maxQty"] = o.MaxQty
	}
	if !IsNil(o.MaxTrailingAboveDelta) {
		toSerialize["maxTrailingAboveDelta"] = o.MaxTrailingAboveDelta
	}
	if !IsNil(o.MaxTrailingBelowDelta) {
		toSerialize["maxTrailingBelowDelta"] = o.MaxTrailingBelowDelta
	}
	if !IsNil(o.MinNotionalValue) {
		toSerialize["minNotionalValue"] = o.MinNotionalValue
	}
	if !IsNil(o.MinPrice) {
		toSerialize["minPrice"] = o.MinPrice
	}
	if !IsNil(o.MinQty) {
		toSerialize["minQty"] = o.MinQty
	}
	if !IsNil(o.MinTrailingAboveDelta) {
		toSerialize["minTrailingAboveDelta"] = o.MinTrailingAboveDelta
	}
	if !IsNil(o.MinTrailingBelowDelta) {
		toSerialize["minTrailingBelowDelta"] = o.MinTrailingBelowDelta
	}
	if !IsNil(o.MultiplierDown) {
		toSerialize["multiplierDown"] = o.MultiplierDown
	}
	if !IsNil(o.MultiplierUp) {
		toSerialize["multiplierUp"] = o.MultiplierUp
	}
	if !IsNil(o.StepSize) {
		toSerialize["stepSize"] = o.StepSize
	}
	if !IsNil(o.TickSize) {
		toSerialize["tickSize"] = o.TickSize
	}
	return toSerialize, nil
}

type NullableSpotSymbolFilter struct {
	value *SpotSymbolFilter
	isSet bool
}

func (v NullableSpotSymbolFilter) Get() *SpotSymbolFilter {
	return v.value
}

func (v *NullableSpotSymbolFilter) Set(val *SpotSymbolFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotSymbolFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotSymbolFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotSymbolFilter(val *SpotSymbolFilter) *NullableSpotSymbolFilter {
	return &NullableSpotSymbolFilter{value: val, isSet: true}
}

func (v NullableSpotSymbolFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotSymbolFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


