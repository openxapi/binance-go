/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the GetAlgoSpotHistoricalOrdersV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlgoSpotHistoricalOrdersV1Resp{}

// GetAlgoSpotHistoricalOrdersV1Resp struct for GetAlgoSpotHistoricalOrdersV1Resp
type GetAlgoSpotHistoricalOrdersV1Resp struct {
	Orders []GetAlgoSpotHistoricalOrdersV1RespOrdersInner `json:"orders,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewGetAlgoSpotHistoricalOrdersV1Resp instantiates a new GetAlgoSpotHistoricalOrdersV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlgoSpotHistoricalOrdersV1Resp() *GetAlgoSpotHistoricalOrdersV1Resp {
	this := GetAlgoSpotHistoricalOrdersV1Resp{}
	return &this
}

// NewGetAlgoSpotHistoricalOrdersV1RespWithDefaults instantiates a new GetAlgoSpotHistoricalOrdersV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlgoSpotHistoricalOrdersV1RespWithDefaults() *GetAlgoSpotHistoricalOrdersV1Resp {
	this := GetAlgoSpotHistoricalOrdersV1Resp{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetOrders() []GetAlgoSpotHistoricalOrdersV1RespOrdersInner {
	if o == nil || IsNil(o.Orders) {
		var ret []GetAlgoSpotHistoricalOrdersV1RespOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetOrdersOk() ([]GetAlgoSpotHistoricalOrdersV1RespOrdersInner, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []GetAlgoSpotHistoricalOrdersV1RespOrdersInner and assigns it to the Orders field.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) SetOrders(v []GetAlgoSpotHistoricalOrdersV1RespOrdersInner) {
	o.Orders = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetAlgoSpotHistoricalOrdersV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o GetAlgoSpotHistoricalOrdersV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlgoSpotHistoricalOrdersV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableGetAlgoSpotHistoricalOrdersV1Resp struct {
	value *GetAlgoSpotHistoricalOrdersV1Resp
	isSet bool
}

func (v NullableGetAlgoSpotHistoricalOrdersV1Resp) Get() *GetAlgoSpotHistoricalOrdersV1Resp {
	return v.value
}

func (v *NullableGetAlgoSpotHistoricalOrdersV1Resp) Set(val *GetAlgoSpotHistoricalOrdersV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAlgoSpotHistoricalOrdersV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAlgoSpotHistoricalOrdersV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAlgoSpotHistoricalOrdersV1Resp(val *GetAlgoSpotHistoricalOrdersV1Resp) *NullableGetAlgoSpotHistoricalOrdersV1Resp {
	return &NullableGetAlgoSpotHistoricalOrdersV1Resp{value: val, isSet: true}
}

func (v NullableGetAlgoSpotHistoricalOrdersV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAlgoSpotHistoricalOrdersV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


