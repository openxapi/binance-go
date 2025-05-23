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

// checks if the GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner{}

// GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner struct for GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner
type GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner struct {
	Amount *string `json:"amount,omitempty"`
	Asset *string `json:"asset,omitempty"`
	DestAccount *string `json:"destAccount,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	RedeemId *int64 `json:"redeemId,omitempty"`
	Status *string `json:"status,omitempty"`
	Time *int64 `json:"time,omitempty"`
}

// NewGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner instantiates a new GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner() *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner {
	this := GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner{}
	return &this
}

// NewGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInnerWithDefaults instantiates a new GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInnerWithDefaults() *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner {
	this := GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) SetAmount(v string) {
	o.Amount = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetDestAccount returns the DestAccount field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetDestAccount() string {
	if o == nil || IsNil(o.DestAccount) {
		var ret string
		return ret
	}
	return *o.DestAccount
}

// GetDestAccountOk returns a tuple with the DestAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetDestAccountOk() (*string, bool) {
	if o == nil || IsNil(o.DestAccount) {
		return nil, false
	}
	return o.DestAccount, true
}

// HasDestAccount returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) HasDestAccount() bool {
	if o != nil && !IsNil(o.DestAccount) {
		return true
	}

	return false
}

// SetDestAccount gets a reference to the given string and assigns it to the DestAccount field.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) SetDestAccount(v string) {
	o.DestAccount = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRedeemId returns the RedeemId field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetRedeemId() int64 {
	if o == nil || IsNil(o.RedeemId) {
		var ret int64
		return ret
	}
	return *o.RedeemId
}

// GetRedeemIdOk returns a tuple with the RedeemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetRedeemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RedeemId) {
		return nil, false
	}
	return o.RedeemId, true
}

// HasRedeemId returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) HasRedeemId() bool {
	if o != nil && !IsNil(o.RedeemId) {
		return true
	}

	return false
}

// SetRedeemId gets a reference to the given int64 and assigns it to the RedeemId field.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) SetRedeemId(v int64) {
	o.RedeemId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) SetStatus(v string) {
	o.Status = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) SetTime(v int64) {
	o.Time = &v
}

func (o GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.DestAccount) {
		toSerialize["destAccount"] = o.DestAccount
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.RedeemId) {
		toSerialize["redeemId"] = o.RedeemId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner struct {
	value *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner
	isSet bool
}

func (v NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) Get() *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner {
	return v.value
}

func (v *NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) Set(val *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner(val *GetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) *NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner {
	return &NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner{value: val, isSet: true}
}

func (v NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimpleEarnFlexibleHistoryRedemptionRecordV1RespRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


