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

// checks if the GetSubAccountTransferSubUserHistoryV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubAccountTransferSubUserHistoryV1RespItem{}

// GetSubAccountTransferSubUserHistoryV1RespItem struct for GetSubAccountTransferSubUserHistoryV1RespItem
type GetSubAccountTransferSubUserHistoryV1RespItem struct {
	Asset *string `json:"asset,omitempty"`
	CounterParty *string `json:"counterParty,omitempty"`
	Email *string `json:"email,omitempty"`
	FromAccountType *string `json:"fromAccountType,omitempty"`
	Qty *string `json:"qty,omitempty"`
	Status *string `json:"status,omitempty"`
	Time *int64 `json:"time,omitempty"`
	ToAccountType *string `json:"toAccountType,omitempty"`
	TranId *int64 `json:"tranId,omitempty"`
	Type *int32 `json:"type,omitempty"`
}

// NewGetSubAccountTransferSubUserHistoryV1RespItem instantiates a new GetSubAccountTransferSubUserHistoryV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountTransferSubUserHistoryV1RespItem() *GetSubAccountTransferSubUserHistoryV1RespItem {
	this := GetSubAccountTransferSubUserHistoryV1RespItem{}
	return &this
}

// NewGetSubAccountTransferSubUserHistoryV1RespItemWithDefaults instantiates a new GetSubAccountTransferSubUserHistoryV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountTransferSubUserHistoryV1RespItemWithDefaults() *GetSubAccountTransferSubUserHistoryV1RespItem {
	this := GetSubAccountTransferSubUserHistoryV1RespItem{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetAsset() string {
	if o == nil || IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetAssetOk() (*string, bool) {
	if o == nil || IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasAsset() bool {
	if o != nil && !IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetAsset(v string) {
	o.Asset = &v
}

// GetCounterParty returns the CounterParty field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetCounterParty() string {
	if o == nil || IsNil(o.CounterParty) {
		var ret string
		return ret
	}
	return *o.CounterParty
}

// GetCounterPartyOk returns a tuple with the CounterParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetCounterPartyOk() (*string, bool) {
	if o == nil || IsNil(o.CounterParty) {
		return nil, false
	}
	return o.CounterParty, true
}

// HasCounterParty returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasCounterParty() bool {
	if o != nil && !IsNil(o.CounterParty) {
		return true
	}

	return false
}

// SetCounterParty gets a reference to the given string and assigns it to the CounterParty field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetCounterParty(v string) {
	o.CounterParty = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetEmail(v string) {
	o.Email = &v
}

// GetFromAccountType returns the FromAccountType field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetFromAccountType() string {
	if o == nil || IsNil(o.FromAccountType) {
		var ret string
		return ret
	}
	return *o.FromAccountType
}

// GetFromAccountTypeOk returns a tuple with the FromAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetFromAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FromAccountType) {
		return nil, false
	}
	return o.FromAccountType, true
}

// HasFromAccountType returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasFromAccountType() bool {
	if o != nil && !IsNil(o.FromAccountType) {
		return true
	}

	return false
}

// SetFromAccountType gets a reference to the given string and assigns it to the FromAccountType field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetFromAccountType(v string) {
	o.FromAccountType = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetQty() string {
	if o == nil || IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetQtyOk() (*string, bool) {
	if o == nil || IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasQty() bool {
	if o != nil && !IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetQty(v string) {
	o.Qty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetStatus(v string) {
	o.Status = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetTime() int64 {
	if o == nil || IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetTime(v int64) {
	o.Time = &v
}

// GetToAccountType returns the ToAccountType field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetToAccountType() string {
	if o == nil || IsNil(o.ToAccountType) {
		var ret string
		return ret
	}
	return *o.ToAccountType
}

// GetToAccountTypeOk returns a tuple with the ToAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetToAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ToAccountType) {
		return nil, false
	}
	return o.ToAccountType, true
}

// HasToAccountType returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasToAccountType() bool {
	if o != nil && !IsNil(o.ToAccountType) {
		return true
	}

	return false
}

// SetToAccountType gets a reference to the given string and assigns it to the ToAccountType field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetToAccountType(v string) {
	o.ToAccountType = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetTranId() int64 {
	if o == nil || IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetTranIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasTranId() bool {
	if o != nil && !IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetTranId(v int64) {
	o.TranId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *GetSubAccountTransferSubUserHistoryV1RespItem) SetType(v int32) {
	o.Type = &v
}

func (o GetSubAccountTransferSubUserHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountTransferSubUserHistoryV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !IsNil(o.CounterParty) {
		toSerialize["counterParty"] = o.CounterParty
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FromAccountType) {
		toSerialize["fromAccountType"] = o.FromAccountType
	}
	if !IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ToAccountType) {
		toSerialize["toAccountType"] = o.ToAccountType
	}
	if !IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetSubAccountTransferSubUserHistoryV1RespItem struct {
	value *GetSubAccountTransferSubUserHistoryV1RespItem
	isSet bool
}

func (v NullableGetSubAccountTransferSubUserHistoryV1RespItem) Get() *GetSubAccountTransferSubUserHistoryV1RespItem {
	return v.value
}

func (v *NullableGetSubAccountTransferSubUserHistoryV1RespItem) Set(val *GetSubAccountTransferSubUserHistoryV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountTransferSubUserHistoryV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountTransferSubUserHistoryV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountTransferSubUserHistoryV1RespItem(val *GetSubAccountTransferSubUserHistoryV1RespItem) *NullableGetSubAccountTransferSubUserHistoryV1RespItem {
	return &NullableGetSubAccountTransferSubUserHistoryV1RespItem{value: val, isSet: true}
}

func (v NullableGetSubAccountTransferSubUserHistoryV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountTransferSubUserHistoryV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


