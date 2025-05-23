/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
)

// checks if the GetPositionRiskV3RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPositionRiskV3RespItem{}

// GetPositionRiskV3RespItem struct for GetPositionRiskV3RespItem
type GetPositionRiskV3RespItem struct {
	Adl *int32 `json:"adl,omitempty"`
	AskNotional *string `json:"askNotional,omitempty"`
	BidNotional *string `json:"bidNotional,omitempty"`
	BreakEvenPrice *string `json:"breakEvenPrice,omitempty"`
	EntryPrice *string `json:"entryPrice,omitempty"`
	InitialMargin *string `json:"initialMargin,omitempty"`
	IsolatedMargin *string `json:"isolatedMargin,omitempty"`
	IsolatedWallet *string `json:"isolatedWallet,omitempty"`
	LiquidationPrice *string `json:"liquidationPrice,omitempty"`
	MaintMargin *string `json:"maintMargin,omitempty"`
	MarginAsset *string `json:"marginAsset,omitempty"`
	MarkPrice *string `json:"markPrice,omitempty"`
	Notional *string `json:"notional,omitempty"`
	OpenOrderInitialMargin *string `json:"openOrderInitialMargin,omitempty"`
	PositionAmt *string `json:"positionAmt,omitempty"`
	PositionInitialMargin *string `json:"positionInitialMargin,omitempty"`
	PositionSide *string `json:"positionSide,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	UnRealizedProfit *string `json:"unRealizedProfit,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewGetPositionRiskV3RespItem instantiates a new GetPositionRiskV3RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPositionRiskV3RespItem() *GetPositionRiskV3RespItem {
	this := GetPositionRiskV3RespItem{}
	return &this
}

// NewGetPositionRiskV3RespItemWithDefaults instantiates a new GetPositionRiskV3RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPositionRiskV3RespItemWithDefaults() *GetPositionRiskV3RespItem {
	this := GetPositionRiskV3RespItem{}
	return &this
}

// GetAdl returns the Adl field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetAdl() int32 {
	if o == nil || IsNil(o.Adl) {
		var ret int32
		return ret
	}
	return *o.Adl
}

// GetAdlOk returns a tuple with the Adl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetAdlOk() (*int32, bool) {
	if o == nil || IsNil(o.Adl) {
		return nil, false
	}
	return o.Adl, true
}

// HasAdl returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasAdl() bool {
	if o != nil && !IsNil(o.Adl) {
		return true
	}

	return false
}

// SetAdl gets a reference to the given int32 and assigns it to the Adl field.
func (o *GetPositionRiskV3RespItem) SetAdl(v int32) {
	o.Adl = &v
}

// GetAskNotional returns the AskNotional field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetAskNotional() string {
	if o == nil || IsNil(o.AskNotional) {
		var ret string
		return ret
	}
	return *o.AskNotional
}

// GetAskNotionalOk returns a tuple with the AskNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetAskNotionalOk() (*string, bool) {
	if o == nil || IsNil(o.AskNotional) {
		return nil, false
	}
	return o.AskNotional, true
}

// HasAskNotional returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasAskNotional() bool {
	if o != nil && !IsNil(o.AskNotional) {
		return true
	}

	return false
}

// SetAskNotional gets a reference to the given string and assigns it to the AskNotional field.
func (o *GetPositionRiskV3RespItem) SetAskNotional(v string) {
	o.AskNotional = &v
}

// GetBidNotional returns the BidNotional field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetBidNotional() string {
	if o == nil || IsNil(o.BidNotional) {
		var ret string
		return ret
	}
	return *o.BidNotional
}

// GetBidNotionalOk returns a tuple with the BidNotional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetBidNotionalOk() (*string, bool) {
	if o == nil || IsNil(o.BidNotional) {
		return nil, false
	}
	return o.BidNotional, true
}

// HasBidNotional returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasBidNotional() bool {
	if o != nil && !IsNil(o.BidNotional) {
		return true
	}

	return false
}

// SetBidNotional gets a reference to the given string and assigns it to the BidNotional field.
func (o *GetPositionRiskV3RespItem) SetBidNotional(v string) {
	o.BidNotional = &v
}

// GetBreakEvenPrice returns the BreakEvenPrice field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetBreakEvenPrice() string {
	if o == nil || IsNil(o.BreakEvenPrice) {
		var ret string
		return ret
	}
	return *o.BreakEvenPrice
}

// GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetBreakEvenPriceOk() (*string, bool) {
	if o == nil || IsNil(o.BreakEvenPrice) {
		return nil, false
	}
	return o.BreakEvenPrice, true
}

// HasBreakEvenPrice returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasBreakEvenPrice() bool {
	if o != nil && !IsNil(o.BreakEvenPrice) {
		return true
	}

	return false
}

// SetBreakEvenPrice gets a reference to the given string and assigns it to the BreakEvenPrice field.
func (o *GetPositionRiskV3RespItem) SetBreakEvenPrice(v string) {
	o.BreakEvenPrice = &v
}

// GetEntryPrice returns the EntryPrice field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetEntryPrice() string {
	if o == nil || IsNil(o.EntryPrice) {
		var ret string
		return ret
	}
	return *o.EntryPrice
}

// GetEntryPriceOk returns a tuple with the EntryPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetEntryPriceOk() (*string, bool) {
	if o == nil || IsNil(o.EntryPrice) {
		return nil, false
	}
	return o.EntryPrice, true
}

// HasEntryPrice returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasEntryPrice() bool {
	if o != nil && !IsNil(o.EntryPrice) {
		return true
	}

	return false
}

// SetEntryPrice gets a reference to the given string and assigns it to the EntryPrice field.
func (o *GetPositionRiskV3RespItem) SetEntryPrice(v string) {
	o.EntryPrice = &v
}

// GetInitialMargin returns the InitialMargin field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetInitialMargin() string {
	if o == nil || IsNil(o.InitialMargin) {
		var ret string
		return ret
	}
	return *o.InitialMargin
}

// GetInitialMarginOk returns a tuple with the InitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.InitialMargin) {
		return nil, false
	}
	return o.InitialMargin, true
}

// HasInitialMargin returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasInitialMargin() bool {
	if o != nil && !IsNil(o.InitialMargin) {
		return true
	}

	return false
}

// SetInitialMargin gets a reference to the given string and assigns it to the InitialMargin field.
func (o *GetPositionRiskV3RespItem) SetInitialMargin(v string) {
	o.InitialMargin = &v
}

// GetIsolatedMargin returns the IsolatedMargin field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetIsolatedMargin() string {
	if o == nil || IsNil(o.IsolatedMargin) {
		var ret string
		return ret
	}
	return *o.IsolatedMargin
}

// GetIsolatedMarginOk returns a tuple with the IsolatedMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetIsolatedMarginOk() (*string, bool) {
	if o == nil || IsNil(o.IsolatedMargin) {
		return nil, false
	}
	return o.IsolatedMargin, true
}

// HasIsolatedMargin returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasIsolatedMargin() bool {
	if o != nil && !IsNil(o.IsolatedMargin) {
		return true
	}

	return false
}

// SetIsolatedMargin gets a reference to the given string and assigns it to the IsolatedMargin field.
func (o *GetPositionRiskV3RespItem) SetIsolatedMargin(v string) {
	o.IsolatedMargin = &v
}

// GetIsolatedWallet returns the IsolatedWallet field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetIsolatedWallet() string {
	if o == nil || IsNil(o.IsolatedWallet) {
		var ret string
		return ret
	}
	return *o.IsolatedWallet
}

// GetIsolatedWalletOk returns a tuple with the IsolatedWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetIsolatedWalletOk() (*string, bool) {
	if o == nil || IsNil(o.IsolatedWallet) {
		return nil, false
	}
	return o.IsolatedWallet, true
}

// HasIsolatedWallet returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasIsolatedWallet() bool {
	if o != nil && !IsNil(o.IsolatedWallet) {
		return true
	}

	return false
}

// SetIsolatedWallet gets a reference to the given string and assigns it to the IsolatedWallet field.
func (o *GetPositionRiskV3RespItem) SetIsolatedWallet(v string) {
	o.IsolatedWallet = &v
}

// GetLiquidationPrice returns the LiquidationPrice field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetLiquidationPrice() string {
	if o == nil || IsNil(o.LiquidationPrice) {
		var ret string
		return ret
	}
	return *o.LiquidationPrice
}

// GetLiquidationPriceOk returns a tuple with the LiquidationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetLiquidationPriceOk() (*string, bool) {
	if o == nil || IsNil(o.LiquidationPrice) {
		return nil, false
	}
	return o.LiquidationPrice, true
}

// HasLiquidationPrice returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasLiquidationPrice() bool {
	if o != nil && !IsNil(o.LiquidationPrice) {
		return true
	}

	return false
}

// SetLiquidationPrice gets a reference to the given string and assigns it to the LiquidationPrice field.
func (o *GetPositionRiskV3RespItem) SetLiquidationPrice(v string) {
	o.LiquidationPrice = &v
}

// GetMaintMargin returns the MaintMargin field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetMaintMargin() string {
	if o == nil || IsNil(o.MaintMargin) {
		var ret string
		return ret
	}
	return *o.MaintMargin
}

// GetMaintMarginOk returns a tuple with the MaintMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetMaintMarginOk() (*string, bool) {
	if o == nil || IsNil(o.MaintMargin) {
		return nil, false
	}
	return o.MaintMargin, true
}

// HasMaintMargin returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasMaintMargin() bool {
	if o != nil && !IsNil(o.MaintMargin) {
		return true
	}

	return false
}

// SetMaintMargin gets a reference to the given string and assigns it to the MaintMargin field.
func (o *GetPositionRiskV3RespItem) SetMaintMargin(v string) {
	o.MaintMargin = &v
}

// GetMarginAsset returns the MarginAsset field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetMarginAsset() string {
	if o == nil || IsNil(o.MarginAsset) {
		var ret string
		return ret
	}
	return *o.MarginAsset
}

// GetMarginAssetOk returns a tuple with the MarginAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetMarginAssetOk() (*string, bool) {
	if o == nil || IsNil(o.MarginAsset) {
		return nil, false
	}
	return o.MarginAsset, true
}

// HasMarginAsset returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasMarginAsset() bool {
	if o != nil && !IsNil(o.MarginAsset) {
		return true
	}

	return false
}

// SetMarginAsset gets a reference to the given string and assigns it to the MarginAsset field.
func (o *GetPositionRiskV3RespItem) SetMarginAsset(v string) {
	o.MarginAsset = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetMarkPrice() string {
	if o == nil || IsNil(o.MarkPrice) {
		var ret string
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetMarkPriceOk() (*string, bool) {
	if o == nil || IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasMarkPrice() bool {
	if o != nil && !IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given string and assigns it to the MarkPrice field.
func (o *GetPositionRiskV3RespItem) SetMarkPrice(v string) {
	o.MarkPrice = &v
}

// GetNotional returns the Notional field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetNotional() string {
	if o == nil || IsNil(o.Notional) {
		var ret string
		return ret
	}
	return *o.Notional
}

// GetNotionalOk returns a tuple with the Notional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetNotionalOk() (*string, bool) {
	if o == nil || IsNil(o.Notional) {
		return nil, false
	}
	return o.Notional, true
}

// HasNotional returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasNotional() bool {
	if o != nil && !IsNil(o.Notional) {
		return true
	}

	return false
}

// SetNotional gets a reference to the given string and assigns it to the Notional field.
func (o *GetPositionRiskV3RespItem) SetNotional(v string) {
	o.Notional = &v
}

// GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetOpenOrderInitialMargin() string {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.OpenOrderInitialMargin
}

// GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.OpenOrderInitialMargin) {
		return nil, false
	}
	return o.OpenOrderInitialMargin, true
}

// HasOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasOpenOrderInitialMargin() bool {
	if o != nil && !IsNil(o.OpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetOpenOrderInitialMargin gets a reference to the given string and assigns it to the OpenOrderInitialMargin field.
func (o *GetPositionRiskV3RespItem) SetOpenOrderInitialMargin(v string) {
	o.OpenOrderInitialMargin = &v
}

// GetPositionAmt returns the PositionAmt field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetPositionAmt() string {
	if o == nil || IsNil(o.PositionAmt) {
		var ret string
		return ret
	}
	return *o.PositionAmt
}

// GetPositionAmtOk returns a tuple with the PositionAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetPositionAmtOk() (*string, bool) {
	if o == nil || IsNil(o.PositionAmt) {
		return nil, false
	}
	return o.PositionAmt, true
}

// HasPositionAmt returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasPositionAmt() bool {
	if o != nil && !IsNil(o.PositionAmt) {
		return true
	}

	return false
}

// SetPositionAmt gets a reference to the given string and assigns it to the PositionAmt field.
func (o *GetPositionRiskV3RespItem) SetPositionAmt(v string) {
	o.PositionAmt = &v
}

// GetPositionInitialMargin returns the PositionInitialMargin field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetPositionInitialMargin() string {
	if o == nil || IsNil(o.PositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.PositionInitialMargin
}

// GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetPositionInitialMarginOk() (*string, bool) {
	if o == nil || IsNil(o.PositionInitialMargin) {
		return nil, false
	}
	return o.PositionInitialMargin, true
}

// HasPositionInitialMargin returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasPositionInitialMargin() bool {
	if o != nil && !IsNil(o.PositionInitialMargin) {
		return true
	}

	return false
}

// SetPositionInitialMargin gets a reference to the given string and assigns it to the PositionInitialMargin field.
func (o *GetPositionRiskV3RespItem) SetPositionInitialMargin(v string) {
	o.PositionInitialMargin = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetPositionSide() string {
	if o == nil || IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetPositionSideOk() (*string, bool) {
	if o == nil || IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasPositionSide() bool {
	if o != nil && !IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *GetPositionRiskV3RespItem) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetPositionRiskV3RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnRealizedProfit returns the UnRealizedProfit field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetUnRealizedProfit() string {
	if o == nil || IsNil(o.UnRealizedProfit) {
		var ret string
		return ret
	}
	return *o.UnRealizedProfit
}

// GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetUnRealizedProfitOk() (*string, bool) {
	if o == nil || IsNil(o.UnRealizedProfit) {
		return nil, false
	}
	return o.UnRealizedProfit, true
}

// HasUnRealizedProfit returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasUnRealizedProfit() bool {
	if o != nil && !IsNil(o.UnRealizedProfit) {
		return true
	}

	return false
}

// SetUnRealizedProfit gets a reference to the given string and assigns it to the UnRealizedProfit field.
func (o *GetPositionRiskV3RespItem) SetUnRealizedProfit(v string) {
	o.UnRealizedProfit = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetPositionRiskV3RespItem) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPositionRiskV3RespItem) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetPositionRiskV3RespItem) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetPositionRiskV3RespItem) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetPositionRiskV3RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPositionRiskV3RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adl) {
		toSerialize["adl"] = o.Adl
	}
	if !IsNil(o.AskNotional) {
		toSerialize["askNotional"] = o.AskNotional
	}
	if !IsNil(o.BidNotional) {
		toSerialize["bidNotional"] = o.BidNotional
	}
	if !IsNil(o.BreakEvenPrice) {
		toSerialize["breakEvenPrice"] = o.BreakEvenPrice
	}
	if !IsNil(o.EntryPrice) {
		toSerialize["entryPrice"] = o.EntryPrice
	}
	if !IsNil(o.InitialMargin) {
		toSerialize["initialMargin"] = o.InitialMargin
	}
	if !IsNil(o.IsolatedMargin) {
		toSerialize["isolatedMargin"] = o.IsolatedMargin
	}
	if !IsNil(o.IsolatedWallet) {
		toSerialize["isolatedWallet"] = o.IsolatedWallet
	}
	if !IsNil(o.LiquidationPrice) {
		toSerialize["liquidationPrice"] = o.LiquidationPrice
	}
	if !IsNil(o.MaintMargin) {
		toSerialize["maintMargin"] = o.MaintMargin
	}
	if !IsNil(o.MarginAsset) {
		toSerialize["marginAsset"] = o.MarginAsset
	}
	if !IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !IsNil(o.Notional) {
		toSerialize["notional"] = o.Notional
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
	if !IsNil(o.UnRealizedProfit) {
		toSerialize["unRealizedProfit"] = o.UnRealizedProfit
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableGetPositionRiskV3RespItem struct {
	value *GetPositionRiskV3RespItem
	isSet bool
}

func (v NullableGetPositionRiskV3RespItem) Get() *GetPositionRiskV3RespItem {
	return v.value
}

func (v *NullableGetPositionRiskV3RespItem) Set(val *GetPositionRiskV3RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPositionRiskV3RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPositionRiskV3RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPositionRiskV3RespItem(val *GetPositionRiskV3RespItem) *NullableGetPositionRiskV3RespItem {
	return &NullableGetPositionRiskV3RespItem{value: val, isSet: true}
}

func (v NullableGetPositionRiskV3RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPositionRiskV3RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


