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

// checks if the GetSubAccountMarginAccountSummaryV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubAccountMarginAccountSummaryV1Resp{}

// GetSubAccountMarginAccountSummaryV1Resp struct for GetSubAccountMarginAccountSummaryV1Resp
type GetSubAccountMarginAccountSummaryV1Resp struct {
	SubAccountList []GetSubAccountMarginAccountSummaryV1RespSubAccountListInner `json:"subAccountList,omitempty"`
	TotalAssetOfBtc *string `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc *string `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc *string `json:"totalNetAssetOfBtc,omitempty"`
}

// NewGetSubAccountMarginAccountSummaryV1Resp instantiates a new GetSubAccountMarginAccountSummaryV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountMarginAccountSummaryV1Resp() *GetSubAccountMarginAccountSummaryV1Resp {
	this := GetSubAccountMarginAccountSummaryV1Resp{}
	return &this
}

// NewGetSubAccountMarginAccountSummaryV1RespWithDefaults instantiates a new GetSubAccountMarginAccountSummaryV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountMarginAccountSummaryV1RespWithDefaults() *GetSubAccountMarginAccountSummaryV1Resp {
	this := GetSubAccountMarginAccountSummaryV1Resp{}
	return &this
}

// GetSubAccountList returns the SubAccountList field value if set, zero value otherwise.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetSubAccountList() []GetSubAccountMarginAccountSummaryV1RespSubAccountListInner {
	if o == nil || IsNil(o.SubAccountList) {
		var ret []GetSubAccountMarginAccountSummaryV1RespSubAccountListInner
		return ret
	}
	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetSubAccountListOk() ([]GetSubAccountMarginAccountSummaryV1RespSubAccountListInner, bool) {
	if o == nil || IsNil(o.SubAccountList) {
		return nil, false
	}
	return o.SubAccountList, true
}

// HasSubAccountList returns a boolean if a field has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) HasSubAccountList() bool {
	if o != nil && !IsNil(o.SubAccountList) {
		return true
	}

	return false
}

// SetSubAccountList gets a reference to the given []GetSubAccountMarginAccountSummaryV1RespSubAccountListInner and assigns it to the SubAccountList field.
func (o *GetSubAccountMarginAccountSummaryV1Resp) SetSubAccountList(v []GetSubAccountMarginAccountSummaryV1RespSubAccountListInner) {
	o.SubAccountList = v
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetTotalAssetOfBtc() string {
	if o == nil || IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) HasTotalAssetOfBtc() bool {
	if o != nil && !IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *GetSubAccountMarginAccountSummaryV1Resp) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetTotalLiabilityOfBtc() string {
	if o == nil || IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) HasTotalLiabilityOfBtc() bool {
	if o != nil && !IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *GetSubAccountMarginAccountSummaryV1Resp) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetTotalNetAssetOfBtc() string {
	if o == nil || IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *GetSubAccountMarginAccountSummaryV1Resp) HasTotalNetAssetOfBtc() bool {
	if o != nil && !IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *GetSubAccountMarginAccountSummaryV1Resp) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

func (o GetSubAccountMarginAccountSummaryV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountMarginAccountSummaryV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubAccountList) {
		toSerialize["subAccountList"] = o.SubAccountList
	}
	if !IsNil(o.TotalAssetOfBtc) {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if !IsNil(o.TotalLiabilityOfBtc) {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if !IsNil(o.TotalNetAssetOfBtc) {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	return toSerialize, nil
}

type NullableGetSubAccountMarginAccountSummaryV1Resp struct {
	value *GetSubAccountMarginAccountSummaryV1Resp
	isSet bool
}

func (v NullableGetSubAccountMarginAccountSummaryV1Resp) Get() *GetSubAccountMarginAccountSummaryV1Resp {
	return v.value
}

func (v *NullableGetSubAccountMarginAccountSummaryV1Resp) Set(val *GetSubAccountMarginAccountSummaryV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountMarginAccountSummaryV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountMarginAccountSummaryV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountMarginAccountSummaryV1Resp(val *GetSubAccountMarginAccountSummaryV1Resp) *NullableGetSubAccountMarginAccountSummaryV1Resp {
	return &NullableGetSubAccountMarginAccountSummaryV1Resp{value: val, isSet: true}
}

func (v NullableGetSubAccountMarginAccountSummaryV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountMarginAccountSummaryV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


