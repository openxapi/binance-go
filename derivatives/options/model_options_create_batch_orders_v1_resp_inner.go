/*
Binance Options API

OpenAPI specification for Binance cryptocurrency exchange - Options API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// OptionsCreateBatchOrdersV1RespInner - struct for OptionsCreateBatchOrdersV1RespInner
type OptionsCreateBatchOrdersV1RespInner struct {
	APIError *APIError
	OptionsCreateBatchOrdersV1RespItem *OptionsCreateBatchOrdersV1RespItem
}

// APIErrorAsOptionsCreateBatchOrdersV1RespInner is a convenience function that returns APIError wrapped in OptionsCreateBatchOrdersV1RespInner
func APIErrorAsOptionsCreateBatchOrdersV1RespInner(v *APIError) OptionsCreateBatchOrdersV1RespInner {
	return OptionsCreateBatchOrdersV1RespInner{
		APIError: v,
	}
}

// OptionsCreateBatchOrdersV1RespItemAsOptionsCreateBatchOrdersV1RespInner is a convenience function that returns OptionsCreateBatchOrdersV1RespItem wrapped in OptionsCreateBatchOrdersV1RespInner
func OptionsCreateBatchOrdersV1RespItemAsOptionsCreateBatchOrdersV1RespInner(v *OptionsCreateBatchOrdersV1RespItem) OptionsCreateBatchOrdersV1RespInner {
	return OptionsCreateBatchOrdersV1RespInner{
		OptionsCreateBatchOrdersV1RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OptionsCreateBatchOrdersV1RespInner) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into OptionsCreateBatchOrdersV1RespItem
	err = newStrictDecoder(data).Decode(&dst.OptionsCreateBatchOrdersV1RespItem)
	if err == nil {
		jsonOptionsCreateBatchOrdersV1RespItem, _ := json.Marshal(dst.OptionsCreateBatchOrdersV1RespItem)
		if string(jsonOptionsCreateBatchOrdersV1RespItem) == "{}" { // empty struct
			dst.OptionsCreateBatchOrdersV1RespItem = nil
		} else {
			if err = validator.Validate(dst.OptionsCreateBatchOrdersV1RespItem); err != nil {
				dst.OptionsCreateBatchOrdersV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.OptionsCreateBatchOrdersV1RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.APIError = nil
		dst.OptionsCreateBatchOrdersV1RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OptionsCreateBatchOrdersV1RespInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OptionsCreateBatchOrdersV1RespInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OptionsCreateBatchOrdersV1RespInner) MarshalJSON() ([]byte, error) {
	if src.APIError != nil {
		return json.Marshal(&src.APIError)
	}

	if src.OptionsCreateBatchOrdersV1RespItem != nil {
		return json.Marshal(&src.OptionsCreateBatchOrdersV1RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OptionsCreateBatchOrdersV1RespInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.APIError != nil {
		return obj.APIError
	}

	if obj.OptionsCreateBatchOrdersV1RespItem != nil {
		return obj.OptionsCreateBatchOrdersV1RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj OptionsCreateBatchOrdersV1RespInner) GetActualInstanceValue() (interface{}) {
	if obj.APIError != nil {
		return *obj.APIError
	}

	if obj.OptionsCreateBatchOrdersV1RespItem != nil {
		return *obj.OptionsCreateBatchOrdersV1RespItem
	}

	// all schemas are nil
	return nil
}

type NullableOptionsCreateBatchOrdersV1RespInner struct {
	value *OptionsCreateBatchOrdersV1RespInner
	isSet bool
}

func (v NullableOptionsCreateBatchOrdersV1RespInner) Get() *OptionsCreateBatchOrdersV1RespInner {
	return v.value
}

func (v *NullableOptionsCreateBatchOrdersV1RespInner) Set(val *OptionsCreateBatchOrdersV1RespInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsCreateBatchOrdersV1RespInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsCreateBatchOrdersV1RespInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsCreateBatchOrdersV1RespInner(val *OptionsCreateBatchOrdersV1RespInner) *NullableOptionsCreateBatchOrdersV1RespInner {
	return &NullableOptionsCreateBatchOrdersV1RespInner{value: val, isSet: true}
}

func (v NullableOptionsCreateBatchOrdersV1RespInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsCreateBatchOrdersV1RespInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


