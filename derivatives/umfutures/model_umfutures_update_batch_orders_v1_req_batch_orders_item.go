/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem{}

// UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem struct for UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem
type UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem struct {
	OrderId *int64 `json:"orderId,omitempty"`
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"`
	Price string `json:"price"`
	PriceMatch *string `json:"priceMatch,omitempty"`
	Quantity string `json:"quantity"`
	Side string `json:"side"`
	Symbol string `json:"symbol"`
}

type _UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem

// NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem instantiates a new UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem(price string, quantity string, side string, symbol string) *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem {
	this := UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem{}
	var origClientOrderId string = ""
	this.OrigClientOrderId = &origClientOrderId
	this.Price = price
	var priceMatch string = ""
	this.PriceMatch = &priceMatch
	this.Quantity = quantity
	this.Side = side
	this.Symbol = symbol
	return &this
}

// NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults instantiates a new UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItemWithDefaults() *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem {
	this := UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem{}
	var origClientOrderId string = ""
	this.OrigClientOrderId = &origClientOrderId
	var price string = ""
	this.Price = price
	var priceMatch string = ""
	this.PriceMatch = &priceMatch
	var quantity string = ""
	this.Quantity = quantity
	var side string = ""
	this.Side = side
	var symbol string = ""
	this.Symbol = symbol
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigClientOrderId returns the OrigClientOrderId field value if set, zero value otherwise.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrigClientOrderId() string {
	if o == nil || IsNil(o.OrigClientOrderId) {
		var ret string
		return ret
	}
	return *o.OrigClientOrderId
}

// GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetOrigClientOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrigClientOrderId) {
		return nil, false
	}
	return o.OrigClientOrderId, true
}

// HasOrigClientOrderId returns a boolean if a field has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasOrigClientOrderId() bool {
	if o != nil && !IsNil(o.OrigClientOrderId) {
		return true
	}

	return false
}

// SetOrigClientOrderId gets a reference to the given string and assigns it to the OrigClientOrderId field.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetOrigClientOrderId(v string) {
	o.OrigClientOrderId = &v
}

// GetPrice returns the Price field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetPrice(v string) {
	o.Price = v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatch() string {
	if o == nil || IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetPriceMatchOk() (*string, bool) {
	if o == nil || IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) HasPriceMatch() bool {
	if o != nil && !IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetQuantity returns the Quantity field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetQuantity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetQuantity(v string) {
	o.Quantity = v
}

// GetSide returns the Side field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSide() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetSide(v string) {
	o.Side = v
}

// GetSymbol returns the Symbol field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) SetSymbol(v string) {
	o.Symbol = v
}

func (o UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrigClientOrderId) {
		toSerialize["origClientOrderId"] = o.OrigClientOrderId
	}
	toSerialize["price"] = o.Price
	if !IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}
	toSerialize["quantity"] = o.Quantity
	toSerialize["side"] = o.Side
	toSerialize["symbol"] = o.Symbol
	return toSerialize, nil
}

func (o *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"price",
		"quantity",
		"side",
		"symbol",
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

	varUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem := _UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem)

	if err != nil {
		return err
	}

	*o = UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem(varUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem)

	return err
}

type NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem struct {
	value *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem
	isSet bool
}

func (v NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) Get() *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem {
	return v.value
}

func (v *NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) Set(val *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem(val *UmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) *NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem {
	return &NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem{value: val, isSet: true}
}

func (v NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesUpdateBatchOrdersV1ReqBatchOrdersItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


