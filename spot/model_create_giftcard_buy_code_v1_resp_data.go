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

// checks if the CreateGiftcardBuyCodeV1RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGiftcardBuyCodeV1RespData{}

// CreateGiftcardBuyCodeV1RespData struct for CreateGiftcardBuyCodeV1RespData
type CreateGiftcardBuyCodeV1RespData struct {
	Code *string `json:"code,omitempty"`
	ExpiredTime *int64 `json:"expiredTime,omitempty"`
	ReferenceNo *string `json:"referenceNo,omitempty"`
}

// NewCreateGiftcardBuyCodeV1RespData instantiates a new CreateGiftcardBuyCodeV1RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGiftcardBuyCodeV1RespData() *CreateGiftcardBuyCodeV1RespData {
	this := CreateGiftcardBuyCodeV1RespData{}
	return &this
}

// NewCreateGiftcardBuyCodeV1RespDataWithDefaults instantiates a new CreateGiftcardBuyCodeV1RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGiftcardBuyCodeV1RespDataWithDefaults() *CreateGiftcardBuyCodeV1RespData {
	this := CreateGiftcardBuyCodeV1RespData{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateGiftcardBuyCodeV1RespData) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGiftcardBuyCodeV1RespData) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateGiftcardBuyCodeV1RespData) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateGiftcardBuyCodeV1RespData) SetCode(v string) {
	o.Code = &v
}

// GetExpiredTime returns the ExpiredTime field value if set, zero value otherwise.
func (o *CreateGiftcardBuyCodeV1RespData) GetExpiredTime() int64 {
	if o == nil || IsNil(o.ExpiredTime) {
		var ret int64
		return ret
	}
	return *o.ExpiredTime
}

// GetExpiredTimeOk returns a tuple with the ExpiredTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGiftcardBuyCodeV1RespData) GetExpiredTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiredTime) {
		return nil, false
	}
	return o.ExpiredTime, true
}

// HasExpiredTime returns a boolean if a field has been set.
func (o *CreateGiftcardBuyCodeV1RespData) HasExpiredTime() bool {
	if o != nil && !IsNil(o.ExpiredTime) {
		return true
	}

	return false
}

// SetExpiredTime gets a reference to the given int64 and assigns it to the ExpiredTime field.
func (o *CreateGiftcardBuyCodeV1RespData) SetExpiredTime(v int64) {
	o.ExpiredTime = &v
}

// GetReferenceNo returns the ReferenceNo field value if set, zero value otherwise.
func (o *CreateGiftcardBuyCodeV1RespData) GetReferenceNo() string {
	if o == nil || IsNil(o.ReferenceNo) {
		var ret string
		return ret
	}
	return *o.ReferenceNo
}

// GetReferenceNoOk returns a tuple with the ReferenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGiftcardBuyCodeV1RespData) GetReferenceNoOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNo) {
		return nil, false
	}
	return o.ReferenceNo, true
}

// HasReferenceNo returns a boolean if a field has been set.
func (o *CreateGiftcardBuyCodeV1RespData) HasReferenceNo() bool {
	if o != nil && !IsNil(o.ReferenceNo) {
		return true
	}

	return false
}

// SetReferenceNo gets a reference to the given string and assigns it to the ReferenceNo field.
func (o *CreateGiftcardBuyCodeV1RespData) SetReferenceNo(v string) {
	o.ReferenceNo = &v
}

func (o CreateGiftcardBuyCodeV1RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGiftcardBuyCodeV1RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ExpiredTime) {
		toSerialize["expiredTime"] = o.ExpiredTime
	}
	if !IsNil(o.ReferenceNo) {
		toSerialize["referenceNo"] = o.ReferenceNo
	}
	return toSerialize, nil
}

type NullableCreateGiftcardBuyCodeV1RespData struct {
	value *CreateGiftcardBuyCodeV1RespData
	isSet bool
}

func (v NullableCreateGiftcardBuyCodeV1RespData) Get() *CreateGiftcardBuyCodeV1RespData {
	return v.value
}

func (v *NullableCreateGiftcardBuyCodeV1RespData) Set(val *CreateGiftcardBuyCodeV1RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGiftcardBuyCodeV1RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGiftcardBuyCodeV1RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGiftcardBuyCodeV1RespData(val *CreateGiftcardBuyCodeV1RespData) *NullableCreateGiftcardBuyCodeV1RespData {
	return &NullableCreateGiftcardBuyCodeV1RespData{value: val, isSet: true}
}

func (v NullableCreateGiftcardBuyCodeV1RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGiftcardBuyCodeV1RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


