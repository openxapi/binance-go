/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// UfuturesGetTickerPriceV1Resp - struct for UfuturesGetTickerPriceV1Resp
type UfuturesGetTickerPriceV1Resp struct {
	UfuturesGetTickerPriceV1RespItem *UfuturesGetTickerPriceV1RespItem
	ArrayOfUfuturesGetTickerPriceV1RespItem *[]UfuturesGetTickerPriceV1RespItem
}

// UfuturesGetTickerPriceV1RespItemAsUfuturesGetTickerPriceV1Resp is a convenience function that returns UfuturesGetTickerPriceV1RespItem wrapped in UfuturesGetTickerPriceV1Resp
func UfuturesGetTickerPriceV1RespItemAsUfuturesGetTickerPriceV1Resp(v *UfuturesGetTickerPriceV1RespItem) UfuturesGetTickerPriceV1Resp {
	return UfuturesGetTickerPriceV1Resp{
		UfuturesGetTickerPriceV1RespItem: v,
	}
}

// []UfuturesGetTickerPriceV1RespItemAsUfuturesGetTickerPriceV1Resp is a convenience function that returns []UfuturesGetTickerPriceV1RespItem wrapped in UfuturesGetTickerPriceV1Resp
func ArrayOfUfuturesGetTickerPriceV1RespItemAsUfuturesGetTickerPriceV1Resp(v *[]UfuturesGetTickerPriceV1RespItem) UfuturesGetTickerPriceV1Resp {
	return UfuturesGetTickerPriceV1Resp{
		ArrayOfUfuturesGetTickerPriceV1RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UfuturesGetTickerPriceV1Resp) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UfuturesGetTickerPriceV1RespItem
	err = newStrictDecoder(data).Decode(&dst.UfuturesGetTickerPriceV1RespItem)
	if err == nil {
		jsonUfuturesGetTickerPriceV1RespItem, _ := json.Marshal(dst.UfuturesGetTickerPriceV1RespItem)
		if string(jsonUfuturesGetTickerPriceV1RespItem) == "{}" { // empty struct
			dst.UfuturesGetTickerPriceV1RespItem = nil
		} else {
			if err = validator.Validate(dst.UfuturesGetTickerPriceV1RespItem); err != nil {
				dst.UfuturesGetTickerPriceV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.UfuturesGetTickerPriceV1RespItem = nil
	}

	// try to unmarshal data into ArrayOfUfuturesGetTickerPriceV1RespItem
	err = newStrictDecoder(data).Decode(&dst.ArrayOfUfuturesGetTickerPriceV1RespItem)
	if err == nil {
		jsonArrayOfUfuturesGetTickerPriceV1RespItem, _ := json.Marshal(dst.ArrayOfUfuturesGetTickerPriceV1RespItem)
		if string(jsonArrayOfUfuturesGetTickerPriceV1RespItem) == "{}" { // empty struct
			dst.ArrayOfUfuturesGetTickerPriceV1RespItem = nil
		} else {
			if err = validator.Validate(dst.ArrayOfUfuturesGetTickerPriceV1RespItem); err != nil {
				dst.ArrayOfUfuturesGetTickerPriceV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfUfuturesGetTickerPriceV1RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UfuturesGetTickerPriceV1RespItem = nil
		dst.ArrayOfUfuturesGetTickerPriceV1RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UfuturesGetTickerPriceV1Resp)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UfuturesGetTickerPriceV1Resp)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UfuturesGetTickerPriceV1Resp) MarshalJSON() ([]byte, error) {
	if src.UfuturesGetTickerPriceV1RespItem != nil {
		return json.Marshal(&src.UfuturesGetTickerPriceV1RespItem)
	}

	if src.ArrayOfUfuturesGetTickerPriceV1RespItem != nil {
		return json.Marshal(&src.ArrayOfUfuturesGetTickerPriceV1RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UfuturesGetTickerPriceV1Resp) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UfuturesGetTickerPriceV1RespItem != nil {
		return obj.UfuturesGetTickerPriceV1RespItem
	}

	if obj.ArrayOfUfuturesGetTickerPriceV1RespItem != nil {
		return obj.ArrayOfUfuturesGetTickerPriceV1RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj UfuturesGetTickerPriceV1Resp) GetActualInstanceValue() (interface{}) {
	if obj.UfuturesGetTickerPriceV1RespItem != nil {
		return *obj.UfuturesGetTickerPriceV1RespItem
	}

	if obj.ArrayOfUfuturesGetTickerPriceV1RespItem != nil {
		return *obj.ArrayOfUfuturesGetTickerPriceV1RespItem
	}

	// all schemas are nil
	return nil
}

type NullableUfuturesGetTickerPriceV1Resp struct {
	value *UfuturesGetTickerPriceV1Resp
	isSet bool
}

func (v NullableUfuturesGetTickerPriceV1Resp) Get() *UfuturesGetTickerPriceV1Resp {
	return v.value
}

func (v *NullableUfuturesGetTickerPriceV1Resp) Set(val *UfuturesGetTickerPriceV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetTickerPriceV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetTickerPriceV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetTickerPriceV1Resp(val *UfuturesGetTickerPriceV1Resp) *NullableUfuturesGetTickerPriceV1Resp {
	return &NullableUfuturesGetTickerPriceV1Resp{value: val, isSet: true}
}

func (v NullableUfuturesGetTickerPriceV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetTickerPriceV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


