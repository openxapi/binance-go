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

// checks if the GetBrokerSubAccountV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBrokerSubAccountV1RespItem{}

// GetBrokerSubAccountV1RespItem struct for GetBrokerSubAccountV1RespItem
type GetBrokerSubAccountV1RespItem struct {
	CreateTime *int64 `json:"createTime,omitempty"`
	Email *string `json:"email,omitempty"`
	MakerCommission *float32 `json:"makerCommission,omitempty"`
	MarginMakerCommission *int32 `json:"marginMakerCommission,omitempty"`
	MarginTakerCommission *int32 `json:"marginTakerCommission,omitempty"`
	SubaccountId *string `json:"subaccountId,omitempty"`
	Tag *string `json:"tag,omitempty"`
	TakerCommission *float32 `json:"takerCommission,omitempty"`
}

// NewGetBrokerSubAccountV1RespItem instantiates a new GetBrokerSubAccountV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBrokerSubAccountV1RespItem() *GetBrokerSubAccountV1RespItem {
	this := GetBrokerSubAccountV1RespItem{}
	return &this
}

// NewGetBrokerSubAccountV1RespItemWithDefaults instantiates a new GetBrokerSubAccountV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBrokerSubAccountV1RespItemWithDefaults() *GetBrokerSubAccountV1RespItem {
	this := GetBrokerSubAccountV1RespItem{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetBrokerSubAccountV1RespItem) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetBrokerSubAccountV1RespItem) SetEmail(v string) {
	o.Email = &v
}

// GetMakerCommission returns the MakerCommission field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetMakerCommission() float32 {
	if o == nil || IsNil(o.MakerCommission) {
		var ret float32
		return ret
	}
	return *o.MakerCommission
}

// GetMakerCommissionOk returns a tuple with the MakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetMakerCommissionOk() (*float32, bool) {
	if o == nil || IsNil(o.MakerCommission) {
		return nil, false
	}
	return o.MakerCommission, true
}

// HasMakerCommission returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasMakerCommission() bool {
	if o != nil && !IsNil(o.MakerCommission) {
		return true
	}

	return false
}

// SetMakerCommission gets a reference to the given float32 and assigns it to the MakerCommission field.
func (o *GetBrokerSubAccountV1RespItem) SetMakerCommission(v float32) {
	o.MakerCommission = &v
}

// GetMarginMakerCommission returns the MarginMakerCommission field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetMarginMakerCommission() int32 {
	if o == nil || IsNil(o.MarginMakerCommission) {
		var ret int32
		return ret
	}
	return *o.MarginMakerCommission
}

// GetMarginMakerCommissionOk returns a tuple with the MarginMakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetMarginMakerCommissionOk() (*int32, bool) {
	if o == nil || IsNil(o.MarginMakerCommission) {
		return nil, false
	}
	return o.MarginMakerCommission, true
}

// HasMarginMakerCommission returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasMarginMakerCommission() bool {
	if o != nil && !IsNil(o.MarginMakerCommission) {
		return true
	}

	return false
}

// SetMarginMakerCommission gets a reference to the given int32 and assigns it to the MarginMakerCommission field.
func (o *GetBrokerSubAccountV1RespItem) SetMarginMakerCommission(v int32) {
	o.MarginMakerCommission = &v
}

// GetMarginTakerCommission returns the MarginTakerCommission field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetMarginTakerCommission() int32 {
	if o == nil || IsNil(o.MarginTakerCommission) {
		var ret int32
		return ret
	}
	return *o.MarginTakerCommission
}

// GetMarginTakerCommissionOk returns a tuple with the MarginTakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetMarginTakerCommissionOk() (*int32, bool) {
	if o == nil || IsNil(o.MarginTakerCommission) {
		return nil, false
	}
	return o.MarginTakerCommission, true
}

// HasMarginTakerCommission returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasMarginTakerCommission() bool {
	if o != nil && !IsNil(o.MarginTakerCommission) {
		return true
	}

	return false
}

// SetMarginTakerCommission gets a reference to the given int32 and assigns it to the MarginTakerCommission field.
func (o *GetBrokerSubAccountV1RespItem) SetMarginTakerCommission(v int32) {
	o.MarginTakerCommission = &v
}

// GetSubaccountId returns the SubaccountId field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetSubaccountId() string {
	if o == nil || IsNil(o.SubaccountId) {
		var ret string
		return ret
	}
	return *o.SubaccountId
}

// GetSubaccountIdOk returns a tuple with the SubaccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetSubaccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubaccountId) {
		return nil, false
	}
	return o.SubaccountId, true
}

// HasSubaccountId returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasSubaccountId() bool {
	if o != nil && !IsNil(o.SubaccountId) {
		return true
	}

	return false
}

// SetSubaccountId gets a reference to the given string and assigns it to the SubaccountId field.
func (o *GetBrokerSubAccountV1RespItem) SetSubaccountId(v string) {
	o.SubaccountId = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetBrokerSubAccountV1RespItem) SetTag(v string) {
	o.Tag = &v
}

// GetTakerCommission returns the TakerCommission field value if set, zero value otherwise.
func (o *GetBrokerSubAccountV1RespItem) GetTakerCommission() float32 {
	if o == nil || IsNil(o.TakerCommission) {
		var ret float32
		return ret
	}
	return *o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBrokerSubAccountV1RespItem) GetTakerCommissionOk() (*float32, bool) {
	if o == nil || IsNil(o.TakerCommission) {
		return nil, false
	}
	return o.TakerCommission, true
}

// HasTakerCommission returns a boolean if a field has been set.
func (o *GetBrokerSubAccountV1RespItem) HasTakerCommission() bool {
	if o != nil && !IsNil(o.TakerCommission) {
		return true
	}

	return false
}

// SetTakerCommission gets a reference to the given float32 and assigns it to the TakerCommission field.
func (o *GetBrokerSubAccountV1RespItem) SetTakerCommission(v float32) {
	o.TakerCommission = &v
}

func (o GetBrokerSubAccountV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBrokerSubAccountV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.MakerCommission) {
		toSerialize["makerCommission"] = o.MakerCommission
	}
	if !IsNil(o.MarginMakerCommission) {
		toSerialize["marginMakerCommission"] = o.MarginMakerCommission
	}
	if !IsNil(o.MarginTakerCommission) {
		toSerialize["marginTakerCommission"] = o.MarginTakerCommission
	}
	if !IsNil(o.SubaccountId) {
		toSerialize["subaccountId"] = o.SubaccountId
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TakerCommission) {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	return toSerialize, nil
}

type NullableGetBrokerSubAccountV1RespItem struct {
	value *GetBrokerSubAccountV1RespItem
	isSet bool
}

func (v NullableGetBrokerSubAccountV1RespItem) Get() *GetBrokerSubAccountV1RespItem {
	return v.value
}

func (v *NullableGetBrokerSubAccountV1RespItem) Set(val *GetBrokerSubAccountV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBrokerSubAccountV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBrokerSubAccountV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBrokerSubAccountV1RespItem(val *GetBrokerSubAccountV1RespItem) *NullableGetBrokerSubAccountV1RespItem {
	return &NullableGetBrokerSubAccountV1RespItem{value: val, isSet: true}
}

func (v NullableGetBrokerSubAccountV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBrokerSubAccountV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


