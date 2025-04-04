/*
Binance Portfolio Margin API

OpenAPI specification for Binance exchange - Pmargin API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pmargin

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// PmarginGetBalanceV1Resp - struct for PmarginGetBalanceV1Resp
type PmarginGetBalanceV1Resp struct {
	PmarginGetBalanceV1RespItem *PmarginGetBalanceV1RespItem
	ArrayOfPmarginGetBalanceV1RespItem *[]PmarginGetBalanceV1RespItem
}

// PmarginGetBalanceV1RespItemAsPmarginGetBalanceV1Resp is a convenience function that returns PmarginGetBalanceV1RespItem wrapped in PmarginGetBalanceV1Resp
func PmarginGetBalanceV1RespItemAsPmarginGetBalanceV1Resp(v *PmarginGetBalanceV1RespItem) PmarginGetBalanceV1Resp {
	return PmarginGetBalanceV1Resp{
		PmarginGetBalanceV1RespItem: v,
	}
}

// []PmarginGetBalanceV1RespItemAsPmarginGetBalanceV1Resp is a convenience function that returns []PmarginGetBalanceV1RespItem wrapped in PmarginGetBalanceV1Resp
func ArrayOfPmarginGetBalanceV1RespItemAsPmarginGetBalanceV1Resp(v *[]PmarginGetBalanceV1RespItem) PmarginGetBalanceV1Resp {
	return PmarginGetBalanceV1Resp{
		ArrayOfPmarginGetBalanceV1RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PmarginGetBalanceV1Resp) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PmarginGetBalanceV1RespItem
	err = newStrictDecoder(data).Decode(&dst.PmarginGetBalanceV1RespItem)
	if err == nil {
		jsonPmarginGetBalanceV1RespItem, _ := json.Marshal(dst.PmarginGetBalanceV1RespItem)
		if string(jsonPmarginGetBalanceV1RespItem) == "{}" { // empty struct
			dst.PmarginGetBalanceV1RespItem = nil
		} else {
			if err = validator.Validate(dst.PmarginGetBalanceV1RespItem); err != nil {
				dst.PmarginGetBalanceV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.PmarginGetBalanceV1RespItem = nil
	}

	// try to unmarshal data into ArrayOfPmarginGetBalanceV1RespItem
	err = newStrictDecoder(data).Decode(&dst.ArrayOfPmarginGetBalanceV1RespItem)
	if err == nil {
		jsonArrayOfPmarginGetBalanceV1RespItem, _ := json.Marshal(dst.ArrayOfPmarginGetBalanceV1RespItem)
		if string(jsonArrayOfPmarginGetBalanceV1RespItem) == "{}" { // empty struct
			dst.ArrayOfPmarginGetBalanceV1RespItem = nil
		} else {
			if err = validator.Validate(dst.ArrayOfPmarginGetBalanceV1RespItem); err != nil {
				dst.ArrayOfPmarginGetBalanceV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfPmarginGetBalanceV1RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PmarginGetBalanceV1RespItem = nil
		dst.ArrayOfPmarginGetBalanceV1RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PmarginGetBalanceV1Resp)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PmarginGetBalanceV1Resp)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PmarginGetBalanceV1Resp) MarshalJSON() ([]byte, error) {
	if src.PmarginGetBalanceV1RespItem != nil {
		return json.Marshal(&src.PmarginGetBalanceV1RespItem)
	}

	if src.ArrayOfPmarginGetBalanceV1RespItem != nil {
		return json.Marshal(&src.ArrayOfPmarginGetBalanceV1RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PmarginGetBalanceV1Resp) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PmarginGetBalanceV1RespItem != nil {
		return obj.PmarginGetBalanceV1RespItem
	}

	if obj.ArrayOfPmarginGetBalanceV1RespItem != nil {
		return obj.ArrayOfPmarginGetBalanceV1RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj PmarginGetBalanceV1Resp) GetActualInstanceValue() (interface{}) {
	if obj.PmarginGetBalanceV1RespItem != nil {
		return *obj.PmarginGetBalanceV1RespItem
	}

	if obj.ArrayOfPmarginGetBalanceV1RespItem != nil {
		return *obj.ArrayOfPmarginGetBalanceV1RespItem
	}

	// all schemas are nil
	return nil
}

type NullablePmarginGetBalanceV1Resp struct {
	value *PmarginGetBalanceV1Resp
	isSet bool
}

func (v NullablePmarginGetBalanceV1Resp) Get() *PmarginGetBalanceV1Resp {
	return v.value
}

func (v *NullablePmarginGetBalanceV1Resp) Set(val *PmarginGetBalanceV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullablePmarginGetBalanceV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullablePmarginGetBalanceV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePmarginGetBalanceV1Resp(val *PmarginGetBalanceV1Resp) *NullablePmarginGetBalanceV1Resp {
	return &NullablePmarginGetBalanceV1Resp{value: val, isSet: true}
}

func (v NullablePmarginGetBalanceV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePmarginGetBalanceV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


