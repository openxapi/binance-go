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

// checks if the GetHistoricalTradesV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetHistoricalTradesV1RespItem{}

// GetHistoricalTradesV1RespItem struct for GetHistoricalTradesV1RespItem
type GetHistoricalTradesV1RespItem struct {
	Id *int32 `json:"id,omitempty"`
	IsBuyerMaker *bool `json:"isBuyerMaker,omitempty"`
	Price *string `json:"price,omitempty"`
	Qty *string `json:"qty,omitempty"`
	QuoteQty *string `json:"quoteQty,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewGetHistoricalTradesV1RespItem instantiates a new GetHistoricalTradesV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetHistoricalTradesV1RespItem() *GetHistoricalTradesV1RespItem {
	this := GetHistoricalTradesV1RespItem{}
	return &this
}

// NewGetHistoricalTradesV1RespItemWithDefaults instantiates a new GetHistoricalTradesV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetHistoricalTradesV1RespItemWithDefaults() *GetHistoricalTradesV1RespItem {
	this := GetHistoricalTradesV1RespItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetHistoricalTradesV1RespItem) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoricalTradesV1RespItem) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetHistoricalTradesV1RespItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetHistoricalTradesV1RespItem) SetId(v int32) {
	o.Id = &v
}

// GetIsBuyerMaker returns the IsBuyerMaker field value if set, zero value otherwise.
func (o *GetHistoricalTradesV1RespItem) GetIsBuyerMaker() bool {
	if o == nil || IsNil(o.IsBuyerMaker) {
		var ret bool
		return ret
	}
	return *o.IsBuyerMaker
}

// GetIsBuyerMakerOk returns a tuple with the IsBuyerMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoricalTradesV1RespItem) GetIsBuyerMakerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBuyerMaker) {
		return nil, false
	}
	return o.IsBuyerMaker, true
}

// HasIsBuyerMaker returns a boolean if a field has been set.
func (o *GetHistoricalTradesV1RespItem) HasIsBuyerMaker() bool {
	if o != nil && !IsNil(o.IsBuyerMaker) {
		return true
	}

	return false
}

// SetIsBuyerMaker gets a reference to the given bool and assigns it to the IsBuyerMaker field.
func (o *GetHistoricalTradesV1RespItem) SetIsBuyerMaker(v bool) {
	o.IsBuyerMaker = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetHistoricalTradesV1RespItem) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoricalTradesV1RespItem) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GetHistoricalTradesV1RespItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *GetHistoricalTradesV1RespItem) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *GetHistoricalTradesV1RespItem) GetQty() string {
	if o == nil || IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoricalTradesV1RespItem) GetQtyOk() (*string, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *GetHistoricalTradesV1RespItem) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *GetHistoricalTradesV1RespItem) SetQty(v string) {
	o.Qty = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *GetHistoricalTradesV1RespItem) GetQuoteQty() string {
	if o == nil || IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoricalTradesV1RespItem) GetQuoteQtyOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *GetHistoricalTradesV1RespItem) HasQuoteQty() bool {
	if o != nil && !IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *GetHistoricalTradesV1RespItem) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetHistoricalTradesV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetHistoricalTradesV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetHistoricalTradesV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetHistoricalTradesV1RespItem) SetTime(v int64) {
	o.Time = &v
}

func (o GetHistoricalTradesV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetHistoricalTradesV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsBuyerMaker) {
		toSerialize["isBuyerMaker"] = o.IsBuyerMaker
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !IsNil(o.QuoteQty) {
		toSerialize["quoteQty"] = o.QuoteQty
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableGetHistoricalTradesV1RespItem struct {
	value *GetHistoricalTradesV1RespItem
	isSet bool
}

func (v NullableGetHistoricalTradesV1RespItem) Get() *GetHistoricalTradesV1RespItem {
	return v.value
}

func (v *NullableGetHistoricalTradesV1RespItem) Set(val *GetHistoricalTradesV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHistoricalTradesV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHistoricalTradesV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHistoricalTradesV1RespItem(val *GetHistoricalTradesV1RespItem) *NullableGetHistoricalTradesV1RespItem {
	return &NullableGetHistoricalTradesV1RespItem{value: val, isSet: true}
}

func (v NullableGetHistoricalTradesV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHistoricalTradesV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


