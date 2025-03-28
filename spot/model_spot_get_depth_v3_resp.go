/*
Binance Spot API

OpenAPI specification for Binance cryptocurrency exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the SpotGetDepthV3Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotGetDepthV3Resp{}

// SpotGetDepthV3Resp struct for SpotGetDepthV3Resp
type SpotGetDepthV3Resp struct {
	Asks [][]string `json:"asks,omitempty"`
	Bids [][]string `json:"bids,omitempty"`
	LastUpdateId *int64 `json:"lastUpdateId,omitempty"`
}

// NewSpotGetDepthV3Resp instantiates a new SpotGetDepthV3Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotGetDepthV3Resp() *SpotGetDepthV3Resp {
	this := SpotGetDepthV3Resp{}
	return &this
}

// NewSpotGetDepthV3RespWithDefaults instantiates a new SpotGetDepthV3Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotGetDepthV3RespWithDefaults() *SpotGetDepthV3Resp {
	this := SpotGetDepthV3Resp{}
	return &this
}

// GetAsks returns the Asks field value if set, zero value otherwise.
func (o *SpotGetDepthV3Resp) GetAsks() [][]string {
	if o == nil || IsNil(o.Asks) {
		var ret [][]string
		return ret
	}
	return o.Asks
}

// GetAsksOk returns a tuple with the Asks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetDepthV3Resp) GetAsksOk() ([][]string, bool) {
	if o == nil || IsNil(o.Asks) {
		return nil, false
	}
	return o.Asks, true
}

// HasAsks returns a boolean if a field has been set.
func (o *SpotGetDepthV3Resp) HasAsks() bool {
	if o != nil && !IsNil(o.Asks) {
		return true
	}

	return false
}

// SetAsks gets a reference to the given [][]string and assigns it to the Asks field.
func (o *SpotGetDepthV3Resp) SetAsks(v [][]string) {
	o.Asks = v
}

// GetBids returns the Bids field value if set, zero value otherwise.
func (o *SpotGetDepthV3Resp) GetBids() [][]string {
	if o == nil || IsNil(o.Bids) {
		var ret [][]string
		return ret
	}
	return o.Bids
}

// GetBidsOk returns a tuple with the Bids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetDepthV3Resp) GetBidsOk() ([][]string, bool) {
	if o == nil || IsNil(o.Bids) {
		return nil, false
	}
	return o.Bids, true
}

// HasBids returns a boolean if a field has been set.
func (o *SpotGetDepthV3Resp) HasBids() bool {
	if o != nil && !IsNil(o.Bids) {
		return true
	}

	return false
}

// SetBids gets a reference to the given [][]string and assigns it to the Bids field.
func (o *SpotGetDepthV3Resp) SetBids(v [][]string) {
	o.Bids = v
}

// GetLastUpdateId returns the LastUpdateId field value if set, zero value otherwise.
func (o *SpotGetDepthV3Resp) GetLastUpdateId() int64 {
	if o == nil || IsNil(o.LastUpdateId) {
		var ret int64
		return ret
	}
	return *o.LastUpdateId
}

// GetLastUpdateIdOk returns a tuple with the LastUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotGetDepthV3Resp) GetLastUpdateIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdateId) {
		return nil, false
	}
	return o.LastUpdateId, true
}

// HasLastUpdateId returns a boolean if a field has been set.
func (o *SpotGetDepthV3Resp) HasLastUpdateId() bool {
	if o != nil && !IsNil(o.LastUpdateId) {
		return true
	}

	return false
}

// SetLastUpdateId gets a reference to the given int64 and assigns it to the LastUpdateId field.
func (o *SpotGetDepthV3Resp) SetLastUpdateId(v int64) {
	o.LastUpdateId = &v
}

func (o SpotGetDepthV3Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotGetDepthV3Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asks) {
		toSerialize["asks"] = o.Asks
	}
	if !IsNil(o.Bids) {
		toSerialize["bids"] = o.Bids
	}
	if !IsNil(o.LastUpdateId) {
		toSerialize["lastUpdateId"] = o.LastUpdateId
	}
	return toSerialize, nil
}

type NullableSpotGetDepthV3Resp struct {
	value *SpotGetDepthV3Resp
	isSet bool
}

func (v NullableSpotGetDepthV3Resp) Get() *SpotGetDepthV3Resp {
	return v.value
}

func (v *NullableSpotGetDepthV3Resp) Set(val *SpotGetDepthV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotGetDepthV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotGetDepthV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotGetDepthV3Resp(val *SpotGetDepthV3Resp) *NullableSpotGetDepthV3Resp {
	return &NullableSpotGetDepthV3Resp{value: val, isSet: true}
}

func (v NullableSpotGetDepthV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotGetDepthV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


