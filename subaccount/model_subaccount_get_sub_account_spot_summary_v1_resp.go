/*
Binance Sub Account API

OpenAPI specification for Binance exchange - Subaccount API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subaccount

import (
	"encoding/json"
)

// checks if the SubaccountGetSubAccountSpotSummaryV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountGetSubAccountSpotSummaryV1Resp{}

// SubaccountGetSubAccountSpotSummaryV1Resp struct for SubaccountGetSubAccountSpotSummaryV1Resp
type SubaccountGetSubAccountSpotSummaryV1Resp struct {
	MasterAccountTotalAsset *string `json:"masterAccountTotalAsset,omitempty"`
	SpotSubUserAssetBtcVoList []SubaccountGetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner `json:"spotSubUserAssetBtcVoList,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewSubaccountGetSubAccountSpotSummaryV1Resp instantiates a new SubaccountGetSubAccountSpotSummaryV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountGetSubAccountSpotSummaryV1Resp() *SubaccountGetSubAccountSpotSummaryV1Resp {
	this := SubaccountGetSubAccountSpotSummaryV1Resp{}
	return &this
}

// NewSubaccountGetSubAccountSpotSummaryV1RespWithDefaults instantiates a new SubaccountGetSubAccountSpotSummaryV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountGetSubAccountSpotSummaryV1RespWithDefaults() *SubaccountGetSubAccountSpotSummaryV1Resp {
	this := SubaccountGetSubAccountSpotSummaryV1Resp{}
	return &this
}

// GetMasterAccountTotalAsset returns the MasterAccountTotalAsset field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) GetMasterAccountTotalAsset() string {
	if o == nil || IsNil(o.MasterAccountTotalAsset) {
		var ret string
		return ret
	}
	return *o.MasterAccountTotalAsset
}

// GetMasterAccountTotalAssetOk returns a tuple with the MasterAccountTotalAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) GetMasterAccountTotalAssetOk() (*string, bool) {
	if o == nil || IsNil(o.MasterAccountTotalAsset) {
		return nil, false
	}
	return o.MasterAccountTotalAsset, true
}

// HasMasterAccountTotalAsset returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) HasMasterAccountTotalAsset() bool {
	if o != nil && !IsNil(o.MasterAccountTotalAsset) {
		return true
	}

	return false
}

// SetMasterAccountTotalAsset gets a reference to the given string and assigns it to the MasterAccountTotalAsset field.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) SetMasterAccountTotalAsset(v string) {
	o.MasterAccountTotalAsset = &v
}

// GetSpotSubUserAssetBtcVoList returns the SpotSubUserAssetBtcVoList field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) GetSpotSubUserAssetBtcVoList() []SubaccountGetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner {
	if o == nil || IsNil(o.SpotSubUserAssetBtcVoList) {
		var ret []SubaccountGetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner
		return ret
	}
	return o.SpotSubUserAssetBtcVoList
}

// GetSpotSubUserAssetBtcVoListOk returns a tuple with the SpotSubUserAssetBtcVoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) GetSpotSubUserAssetBtcVoListOk() ([]SubaccountGetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner, bool) {
	if o == nil || IsNil(o.SpotSubUserAssetBtcVoList) {
		return nil, false
	}
	return o.SpotSubUserAssetBtcVoList, true
}

// HasSpotSubUserAssetBtcVoList returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) HasSpotSubUserAssetBtcVoList() bool {
	if o != nil && !IsNil(o.SpotSubUserAssetBtcVoList) {
		return true
	}

	return false
}

// SetSpotSubUserAssetBtcVoList gets a reference to the given []SubaccountGetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner and assigns it to the SpotSubUserAssetBtcVoList field.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) SetSpotSubUserAssetBtcVoList(v []SubaccountGetSubAccountSpotSummaryV1RespSpotSubUserAssetBtcVoListInner) {
	o.SpotSubUserAssetBtcVoList = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *SubaccountGetSubAccountSpotSummaryV1Resp) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o SubaccountGetSubAccountSpotSummaryV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountGetSubAccountSpotSummaryV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MasterAccountTotalAsset) {
		toSerialize["masterAccountTotalAsset"] = o.MasterAccountTotalAsset
	}
	if !IsNil(o.SpotSubUserAssetBtcVoList) {
		toSerialize["spotSubUserAssetBtcVoList"] = o.SpotSubUserAssetBtcVoList
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableSubaccountGetSubAccountSpotSummaryV1Resp struct {
	value *SubaccountGetSubAccountSpotSummaryV1Resp
	isSet bool
}

func (v NullableSubaccountGetSubAccountSpotSummaryV1Resp) Get() *SubaccountGetSubAccountSpotSummaryV1Resp {
	return v.value
}

func (v *NullableSubaccountGetSubAccountSpotSummaryV1Resp) Set(val *SubaccountGetSubAccountSpotSummaryV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountGetSubAccountSpotSummaryV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountGetSubAccountSpotSummaryV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountGetSubAccountSpotSummaryV1Resp(val *SubaccountGetSubAccountSpotSummaryV1Resp) *NullableSubaccountGetSubAccountSpotSummaryV1Resp {
	return &NullableSubaccountGetSubAccountSpotSummaryV1Resp{value: val, isSet: true}
}

func (v NullableSubaccountGetSubAccountSpotSummaryV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountGetSubAccountSpotSummaryV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


