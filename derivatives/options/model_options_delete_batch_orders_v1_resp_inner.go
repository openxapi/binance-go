/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// OptionsDeleteBatchOrdersV1RespInner - struct for OptionsDeleteBatchOrdersV1RespInner
type OptionsDeleteBatchOrdersV1RespInner struct {
	APIError *APIError
	OptionsDeleteBatchOrdersV1RespItem *OptionsDeleteBatchOrdersV1RespItem
}

// APIErrorAsOptionsDeleteBatchOrdersV1RespInner is a convenience function that returns APIError wrapped in OptionsDeleteBatchOrdersV1RespInner
func APIErrorAsOptionsDeleteBatchOrdersV1RespInner(v *APIError) OptionsDeleteBatchOrdersV1RespInner {
	return OptionsDeleteBatchOrdersV1RespInner{
		APIError: v,
	}
}

// OptionsDeleteBatchOrdersV1RespItemAsOptionsDeleteBatchOrdersV1RespInner is a convenience function that returns OptionsDeleteBatchOrdersV1RespItem wrapped in OptionsDeleteBatchOrdersV1RespInner
func OptionsDeleteBatchOrdersV1RespItemAsOptionsDeleteBatchOrdersV1RespInner(v *OptionsDeleteBatchOrdersV1RespItem) OptionsDeleteBatchOrdersV1RespInner {
	return OptionsDeleteBatchOrdersV1RespInner{
		OptionsDeleteBatchOrdersV1RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OptionsDeleteBatchOrdersV1RespInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into APIError
	err = newStrictDecoder(data).Decode(&dst.APIError)
	if err == nil {
		jsonAPIError, _ := json.Marshal(dst.APIError)
		if string(jsonAPIError) == "{}" { // empty struct
			dst.APIError = nil
		} else {
			if err = validator.Validate(dst.APIError); err != nil {
				dst.APIError = nil
			} else {
				match++
			}
		}
	} else {
		dst.APIError = nil
	}

	// try to unmarshal data into OptionsDeleteBatchOrdersV1RespItem
	err = newStrictDecoder(data).Decode(&dst.OptionsDeleteBatchOrdersV1RespItem)
	if err == nil {
		jsonOptionsDeleteBatchOrdersV1RespItem, _ := json.Marshal(dst.OptionsDeleteBatchOrdersV1RespItem)
		if string(jsonOptionsDeleteBatchOrdersV1RespItem) == "{}" { // empty struct
			dst.OptionsDeleteBatchOrdersV1RespItem = nil
		} else {
			if err = validator.Validate(dst.OptionsDeleteBatchOrdersV1RespItem); err != nil {
				dst.OptionsDeleteBatchOrdersV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.OptionsDeleteBatchOrdersV1RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.APIError = nil
		dst.OptionsDeleteBatchOrdersV1RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OptionsDeleteBatchOrdersV1RespInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OptionsDeleteBatchOrdersV1RespInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OptionsDeleteBatchOrdersV1RespInner) MarshalJSON() ([]byte, error) {
	if src.APIError != nil {
		return json.Marshal(&src.APIError)
	}

	if src.OptionsDeleteBatchOrdersV1RespItem != nil {
		return json.Marshal(&src.OptionsDeleteBatchOrdersV1RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OptionsDeleteBatchOrdersV1RespInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.APIError != nil {
		return obj.APIError
	}

	if obj.OptionsDeleteBatchOrdersV1RespItem != nil {
		return obj.OptionsDeleteBatchOrdersV1RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj OptionsDeleteBatchOrdersV1RespInner) GetActualInstanceValue() (interface{}) {
	if obj.APIError != nil {
		return *obj.APIError
	}

	if obj.OptionsDeleteBatchOrdersV1RespItem != nil {
		return *obj.OptionsDeleteBatchOrdersV1RespItem
	}

	// all schemas are nil
	return nil
}

type NullableOptionsDeleteBatchOrdersV1RespInner struct {
	value *OptionsDeleteBatchOrdersV1RespInner
	isSet bool
}

func (v NullableOptionsDeleteBatchOrdersV1RespInner) Get() *OptionsDeleteBatchOrdersV1RespInner {
	return v.value
}

func (v *NullableOptionsDeleteBatchOrdersV1RespInner) Set(val *OptionsDeleteBatchOrdersV1RespInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsDeleteBatchOrdersV1RespInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsDeleteBatchOrdersV1RespInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsDeleteBatchOrdersV1RespInner(val *OptionsDeleteBatchOrdersV1RespInner) *NullableOptionsDeleteBatchOrdersV1RespInner {
	return &NullableOptionsDeleteBatchOrdersV1RespInner{value: val, isSet: true}
}

func (v NullableOptionsDeleteBatchOrdersV1RespInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsDeleteBatchOrdersV1RespInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


