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

// checks if the GetAlgoSpotSubOrdersV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlgoSpotSubOrdersV1Resp{}

// GetAlgoSpotSubOrdersV1Resp struct for GetAlgoSpotSubOrdersV1Resp
type GetAlgoSpotSubOrdersV1Resp struct {
	ExecutedAmt *string `json:"executedAmt,omitempty"`
	ExecutedQty *string `json:"executedQty,omitempty"`
	SubOrders []GetAlgoFuturesSubOrdersV1RespSubOrdersInner `json:"subOrders,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewGetAlgoSpotSubOrdersV1Resp instantiates a new GetAlgoSpotSubOrdersV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlgoSpotSubOrdersV1Resp() *GetAlgoSpotSubOrdersV1Resp {
	this := GetAlgoSpotSubOrdersV1Resp{}
	return &this
}

// NewGetAlgoSpotSubOrdersV1RespWithDefaults instantiates a new GetAlgoSpotSubOrdersV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlgoSpotSubOrdersV1RespWithDefaults() *GetAlgoSpotSubOrdersV1Resp {
	this := GetAlgoSpotSubOrdersV1Resp{}
	return &this
}

// GetExecutedAmt returns the ExecutedAmt field value if set, zero value otherwise.
func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedAmt() string {
	if o == nil || IsNil(o.ExecutedAmt) {
		var ret string
		return ret
	}
	return *o.ExecutedAmt
}

// GetExecutedAmtOk returns a tuple with the ExecutedAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedAmtOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedAmt) {
		return nil, false
	}
	return o.ExecutedAmt, true
}

// HasExecutedAmt returns a boolean if a field has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) HasExecutedAmt() bool {
	if o != nil && !IsNil(o.ExecutedAmt) {
		return true
	}

	return false
}

// SetExecutedAmt gets a reference to the given string and assigns it to the ExecutedAmt field.
func (o *GetAlgoSpotSubOrdersV1Resp) SetExecutedAmt(v string) {
	o.ExecutedAmt = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedQty() string {
	if o == nil || IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) GetExecutedQtyOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) HasExecutedQty() bool {
	if o != nil && !IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *GetAlgoSpotSubOrdersV1Resp) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetSubOrders returns the SubOrders field value if set, zero value otherwise.
func (o *GetAlgoSpotSubOrdersV1Resp) GetSubOrders() []GetAlgoFuturesSubOrdersV1RespSubOrdersInner {
	if o == nil || IsNil(o.SubOrders) {
		var ret []GetAlgoFuturesSubOrdersV1RespSubOrdersInner
		return ret
	}
	return o.SubOrders
}

// GetSubOrdersOk returns a tuple with the SubOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) GetSubOrdersOk() ([]GetAlgoFuturesSubOrdersV1RespSubOrdersInner, bool) {
	if o == nil || IsNil(o.SubOrders) {
		return nil, false
	}
	return o.SubOrders, true
}

// HasSubOrders returns a boolean if a field has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) HasSubOrders() bool {
	if o != nil && !IsNil(o.SubOrders) {
		return true
	}

	return false
}

// SetSubOrders gets a reference to the given []GetAlgoFuturesSubOrdersV1RespSubOrdersInner and assigns it to the SubOrders field.
func (o *GetAlgoSpotSubOrdersV1Resp) SetSubOrders(v []GetAlgoFuturesSubOrdersV1RespSubOrdersInner) {
	o.SubOrders = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetAlgoSpotSubOrdersV1Resp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetAlgoSpotSubOrdersV1Resp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetAlgoSpotSubOrdersV1Resp) SetTotal(v int32) {
	o.Total = &v
}

func (o GetAlgoSpotSubOrdersV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlgoSpotSubOrdersV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutedAmt) {
		toSerialize["executedAmt"] = o.ExecutedAmt
	}
	if !IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !IsNil(o.SubOrders) {
		toSerialize["subOrders"] = o.SubOrders
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableGetAlgoSpotSubOrdersV1Resp struct {
	value *GetAlgoSpotSubOrdersV1Resp
	isSet bool
}

func (v NullableGetAlgoSpotSubOrdersV1Resp) Get() *GetAlgoSpotSubOrdersV1Resp {
	return v.value
}

func (v *NullableGetAlgoSpotSubOrdersV1Resp) Set(val *GetAlgoSpotSubOrdersV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAlgoSpotSubOrdersV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAlgoSpotSubOrdersV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAlgoSpotSubOrdersV1Resp(val *GetAlgoSpotSubOrdersV1Resp) *NullableGetAlgoSpotSubOrdersV1Resp {
	return &NullableGetAlgoSpotSubOrdersV1Resp{value: val, isSet: true}
}

func (v NullableGetAlgoSpotSubOrdersV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAlgoSpotSubOrdersV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


