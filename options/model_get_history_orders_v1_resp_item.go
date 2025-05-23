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

// checks if the GetHistoryOrdersV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetHistoryOrdersV1RespItem{}

// GetHistoryOrdersV1RespItem struct for GetHistoryOrdersV1RespItem
type GetHistoryOrdersV1RespItem struct {
	AvgPrice *string `json:"avgPrice,omitempty"`
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	CreateTime *int64 `json:"createTime,omitempty"`
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
	Reason *string `json:"reason,omitempty"`
	ReduceOnly *bool `json:"reduceOnly,omitempty"`
	Side *string `json:"side,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	TimeInForce *string `json:"timeInForce,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// NewGetHistoryOrdersV1RespItem instantiates a new GetHistoryOrdersV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHistoryOrdersV1RespItem() *GetHistoryOrdersV1RespItem {
	this := GetHistoryOrdersV1RespItem{}
	return &this
}

// NewGetHistoryOrdersV1RespItemWithDefaults instantiates a new GetHistoryOrdersV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHistoryOrdersV1RespItemWithDefaults() *GetHistoryOrdersV1RespItem {
	this := GetHistoryOrdersV1RespItem{}
	return &this
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetAvgPrice() string {
	if o == nil || IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetAvgPriceOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasAvgPrice() bool {
	if o != nil && !IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *GetHistoryOrdersV1RespItem) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetClientOrderId() string {
	if o == nil || IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasClientOrderId() bool {
	if o != nil && !IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *GetHistoryOrdersV1RespItem) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetHistoryOrdersV1RespItem) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *GetHistoryOrdersV1RespItem) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *GetHistoryOrdersV1RespItem) SetFee(v string) {
	o.Fee = &v
}

// GetMmp returns the Mmp field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetMmp() bool {
	if o == nil || IsNil(o.Mmp) {
		var ret bool
		return ret
	}
	return *o.Mmp
}

// GetMmpOk returns a tuple with the Mmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetMmpOk() (*bool, bool) {
	if o == nil || IsNil(o.Mmp) {
		return nil, false
	}
	return o.Mmp, true
}

// HasMmp returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasMmp() bool {
	if o != nil && !IsNil(o.Mmp) {
		return true
	}

	return false
}

// SetMmp gets a reference to the given bool and assigns it to the Mmp field.
func (o *GetHistoryOrdersV1RespItem) SetMmp(v bool) {
	o.Mmp = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetOptionSide() string {
	if o == nil || IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetOptionSideOk() (*string, bool) {
	if o == nil || IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasOptionSide() bool {
	if o != nil && !IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *GetHistoryOrdersV1RespItem) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetHistoryOrdersV1RespItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetPostOnly() bool {
	if o == nil || IsNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetPostOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.PostOnly) {
		return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasPostOnly() bool {
	if o != nil && !IsNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *GetHistoryOrdersV1RespItem) SetPostOnly(v bool) {
	o.PostOnly = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetHistoryOrdersV1RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetPriceScale() int32 {
	if o == nil || IsNil(o.PriceScale) {
		var ret int32
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetPriceScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasPriceScale() bool {
	if o != nil && !IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int32 and assigns it to the PriceScale field.
func (o *GetHistoryOrdersV1RespItem) SetPriceScale(v int32) {
	o.PriceScale = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *GetHistoryOrdersV1RespItem) SetQuantity(v string) {
	o.Quantity = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetQuantityScale() int32 {
	if o == nil || IsNil(o.QuantityScale) {
		var ret int32
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetQuantityScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasQuantityScale() bool {
	if o != nil && !IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int32 and assigns it to the QuantityScale field.
func (o *GetHistoryOrdersV1RespItem) SetQuantityScale(v int32) {
	o.QuantityScale = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetQuoteAsset() string {
	if o == nil || IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetQuoteAssetOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasQuoteAsset() bool {
	if o != nil && !IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *GetHistoryOrdersV1RespItem) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *GetHistoryOrdersV1RespItem) SetReason(v string) {
	o.Reason = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetReduceOnly() bool {
	if o == nil || IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasReduceOnly() bool {
	if o != nil && !IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *GetHistoryOrdersV1RespItem) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetHistoryOrdersV1RespItem) SetSide(v string) {
	o.Side = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *GetHistoryOrdersV1RespItem) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetHistoryOrdersV1RespItem) SetStatus(v string) {
	o.Status = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetHistoryOrdersV1RespItem) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetTimeInForce() string {
	if o == nil || IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetTimeInForceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasTimeInForce() bool {
	if o != nil && !IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *GetHistoryOrdersV1RespItem) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetHistoryOrdersV1RespItem) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetHistoryOrdersV1RespItem) GetUpdateTime() int64 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoryOrdersV1RespItem) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetHistoryOrdersV1RespItem) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *GetHistoryOrdersV1RespItem) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o GetHistoryOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetHistoryOrdersV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
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
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
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

type NullableGetHistoryOrdersV1RespItem struct {
	value *GetHistoryOrdersV1RespItem
	isSet bool
}

func (v NullableGetHistoryOrdersV1RespItem) Get() *GetHistoryOrdersV1RespItem {
	return v.value
}

func (v *NullableGetHistoryOrdersV1RespItem) Set(val *GetHistoryOrdersV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHistoryOrdersV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHistoryOrdersV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHistoryOrdersV1RespItem(val *GetHistoryOrdersV1RespItem) *NullableGetHistoryOrdersV1RespItem {
	return &NullableGetHistoryOrdersV1RespItem{value: val, isSet: true}
}

func (v NullableGetHistoryOrdersV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHistoryOrdersV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


