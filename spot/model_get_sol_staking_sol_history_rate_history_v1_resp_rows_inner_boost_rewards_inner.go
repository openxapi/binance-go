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

// checks if the GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner{}

// GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner struct for GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner
type GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner struct {
	BoostAPR *string `json:"boostAPR,omitempty"`
	RewardsAsset *string `json:"rewardsAsset,omitempty"`
}

// NewGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner instantiates a new GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner() *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner {
	this := GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner{}
	return &this
}

// NewGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInnerWithDefaults instantiates a new GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInnerWithDefaults() *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner {
	this := GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner{}
	return &this
}

// GetBoostAPR returns the BoostAPR field value if set, zero value otherwise.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) GetBoostAPR() string {
	if o == nil || IsNil(o.BoostAPR) {
		var ret string
		return ret
	}
	return *o.BoostAPR
}

// GetBoostAPROk returns a tuple with the BoostAPR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) GetBoostAPROk() (*string, bool) {
	if o == nil || IsNil(o.BoostAPR) {
		return nil, false
	}
	return o.BoostAPR, true
}

// HasBoostAPR returns a boolean if a field has been set.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) HasBoostAPR() bool {
	if o != nil && !IsNil(o.BoostAPR) {
		return true
	}

	return false
}

// SetBoostAPR gets a reference to the given string and assigns it to the BoostAPR field.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) SetBoostAPR(v string) {
	o.BoostAPR = &v
}

// GetRewardsAsset returns the RewardsAsset field value if set, zero value otherwise.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) GetRewardsAsset() string {
	if o == nil || IsNil(o.RewardsAsset) {
		var ret string
		return ret
	}
	return *o.RewardsAsset
}

// GetRewardsAssetOk returns a tuple with the RewardsAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) GetRewardsAssetOk() (*string, bool) {
	if o == nil || IsNil(o.RewardsAsset) {
		return nil, false
	}
	return o.RewardsAsset, true
}

// HasRewardsAsset returns a boolean if a field has been set.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) HasRewardsAsset() bool {
	if o != nil && !IsNil(o.RewardsAsset) {
		return true
	}

	return false
}

// SetRewardsAsset gets a reference to the given string and assigns it to the RewardsAsset field.
func (o *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) SetRewardsAsset(v string) {
	o.RewardsAsset = &v
}

func (o GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoostAPR) {
		toSerialize["boostAPR"] = o.BoostAPR
	}
	if !IsNil(o.RewardsAsset) {
		toSerialize["rewardsAsset"] = o.RewardsAsset
	}
	return toSerialize, nil
}

type NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner struct {
	value *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner
	isSet bool
}

func (v NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) Get() *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner {
	return v.value
}

func (v *NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) Set(val *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner(val *GetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) *NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner {
	return &NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner{value: val, isSet: true}
}

func (v NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSolStakingSolHistoryRateHistoryV1RespRowsInnerBoostRewardsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


