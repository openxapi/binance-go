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

// checks if the GetMarginLeverageBracketV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarginLeverageBracketV1RespItem{}

// GetMarginLeverageBracketV1RespItem struct for GetMarginLeverageBracketV1RespItem
type GetMarginLeverageBracketV1RespItem struct {
	AssetNames []string `json:"assetNames,omitempty"`
	Brackets []GetMarginLeverageBracketV1RespItemBracketsInner `json:"brackets,omitempty"`
	Rank *int32 `json:"rank,omitempty"`
}

// NewGetMarginLeverageBracketV1RespItem instantiates a new GetMarginLeverageBracketV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarginLeverageBracketV1RespItem() *GetMarginLeverageBracketV1RespItem {
	this := GetMarginLeverageBracketV1RespItem{}
	return &this
}

// NewGetMarginLeverageBracketV1RespItemWithDefaults instantiates a new GetMarginLeverageBracketV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarginLeverageBracketV1RespItemWithDefaults() *GetMarginLeverageBracketV1RespItem {
	this := GetMarginLeverageBracketV1RespItem{}
	return &this
}

// GetAssetNames returns the AssetNames field value if set, zero value otherwise.
func (o *GetMarginLeverageBracketV1RespItem) GetAssetNames() []string {
	if o == nil || IsNil(o.AssetNames) {
		var ret []string
		return ret
	}
	return o.AssetNames
}

// GetAssetNamesOk returns a tuple with the AssetNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginLeverageBracketV1RespItem) GetAssetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AssetNames) {
		return nil, false
	}
	return o.AssetNames, true
}

// HasAssetNames returns a boolean if a field has been set.
func (o *GetMarginLeverageBracketV1RespItem) HasAssetNames() bool {
	if o != nil && !IsNil(o.AssetNames) {
		return true
	}

	return false
}

// SetAssetNames gets a reference to the given []string and assigns it to the AssetNames field.
func (o *GetMarginLeverageBracketV1RespItem) SetAssetNames(v []string) {
	o.AssetNames = v
}

// GetBrackets returns the Brackets field value if set, zero value otherwise.
func (o *GetMarginLeverageBracketV1RespItem) GetBrackets() []GetMarginLeverageBracketV1RespItemBracketsInner {
	if o == nil || IsNil(o.Brackets) {
		var ret []GetMarginLeverageBracketV1RespItemBracketsInner
		return ret
	}
	return o.Brackets
}

// GetBracketsOk returns a tuple with the Brackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginLeverageBracketV1RespItem) GetBracketsOk() ([]GetMarginLeverageBracketV1RespItemBracketsInner, bool) {
	if o == nil || IsNil(o.Brackets) {
		return nil, false
	}
	return o.Brackets, true
}

// HasBrackets returns a boolean if a field has been set.
func (o *GetMarginLeverageBracketV1RespItem) HasBrackets() bool {
	if o != nil && !IsNil(o.Brackets) {
		return true
	}

	return false
}

// SetBrackets gets a reference to the given []GetMarginLeverageBracketV1RespItemBracketsInner and assigns it to the Brackets field.
func (o *GetMarginLeverageBracketV1RespItem) SetBrackets(v []GetMarginLeverageBracketV1RespItemBracketsInner) {
	o.Brackets = v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *GetMarginLeverageBracketV1RespItem) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarginLeverageBracketV1RespItem) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *GetMarginLeverageBracketV1RespItem) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *GetMarginLeverageBracketV1RespItem) SetRank(v int32) {
	o.Rank = &v
}

func (o GetMarginLeverageBracketV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarginLeverageBracketV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetNames) {
		toSerialize["assetNames"] = o.AssetNames
	}
	if !IsNil(o.Brackets) {
		toSerialize["brackets"] = o.Brackets
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

type NullableGetMarginLeverageBracketV1RespItem struct {
	value *GetMarginLeverageBracketV1RespItem
	isSet bool
}

func (v NullableGetMarginLeverageBracketV1RespItem) Get() *GetMarginLeverageBracketV1RespItem {
	return v.value
}

func (v *NullableGetMarginLeverageBracketV1RespItem) Set(val *GetMarginLeverageBracketV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarginLeverageBracketV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarginLeverageBracketV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarginLeverageBracketV1RespItem(val *GetMarginLeverageBracketV1RespItem) *NullableGetMarginLeverageBracketV1RespItem {
	return &NullableGetMarginLeverageBracketV1RespItem{value: val, isSet: true}
}

func (v NullableGetMarginLeverageBracketV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarginLeverageBracketV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


