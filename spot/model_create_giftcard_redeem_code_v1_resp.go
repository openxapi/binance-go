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

// checks if the CreateGiftcardRedeemCodeV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGiftcardRedeemCodeV1Resp{}

// CreateGiftcardRedeemCodeV1Resp struct for CreateGiftcardRedeemCodeV1Resp
type CreateGiftcardRedeemCodeV1Resp struct {
	Code *string `json:"code,omitempty"`
	Data *CreateGiftcardRedeemCodeV1RespData `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	Success *bool `json:"success,omitempty"`
}

// NewCreateGiftcardRedeemCodeV1Resp instantiates a new CreateGiftcardRedeemCodeV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGiftcardRedeemCodeV1Resp() *CreateGiftcardRedeemCodeV1Resp {
	this := CreateGiftcardRedeemCodeV1Resp{}
	return &this
}

// NewCreateGiftcardRedeemCodeV1RespWithDefaults instantiates a new CreateGiftcardRedeemCodeV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGiftcardRedeemCodeV1RespWithDefaults() *CreateGiftcardRedeemCodeV1Resp {
	this := CreateGiftcardRedeemCodeV1Resp{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateGiftcardRedeemCodeV1Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateGiftcardRedeemCodeV1Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateGiftcardRedeemCodeV1Resp) GetData() CreateGiftcardRedeemCodeV1RespData {
	if o == nil || IsNil(o.Data) {
		var ret CreateGiftcardRedeemCodeV1RespData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) GetDataOk() (*CreateGiftcardRedeemCodeV1RespData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateGiftcardRedeemCodeV1RespData and assigns it to the Data field.
func (o *CreateGiftcardRedeemCodeV1Resp) SetData(v CreateGiftcardRedeemCodeV1RespData) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateGiftcardRedeemCodeV1Resp) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateGiftcardRedeemCodeV1Resp) SetMessage(v string) {
	o.Message = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateGiftcardRedeemCodeV1Resp) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateGiftcardRedeemCodeV1Resp) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateGiftcardRedeemCodeV1Resp) SetSuccess(v bool) {
	o.Success = &v
}

func (o CreateGiftcardRedeemCodeV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGiftcardRedeemCodeV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableCreateGiftcardRedeemCodeV1Resp struct {
	value *CreateGiftcardRedeemCodeV1Resp
	isSet bool
}

func (v NullableCreateGiftcardRedeemCodeV1Resp) Get() *CreateGiftcardRedeemCodeV1Resp {
	return v.value
}

func (v *NullableCreateGiftcardRedeemCodeV1Resp) Set(val *CreateGiftcardRedeemCodeV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGiftcardRedeemCodeV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGiftcardRedeemCodeV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGiftcardRedeemCodeV1Resp(val *CreateGiftcardRedeemCodeV1Resp) *NullableCreateGiftcardRedeemCodeV1Resp {
	return &NullableCreateGiftcardRedeemCodeV1Resp{value: val, isSet: true}
}

func (v NullableCreateGiftcardRedeemCodeV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGiftcardRedeemCodeV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


