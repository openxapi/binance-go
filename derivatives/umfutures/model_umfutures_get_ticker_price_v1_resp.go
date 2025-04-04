/*
Binance USD-M Futures API

OpenAPI specification for Binance exchange - Umfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package umfutures

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// UmfuturesGetTickerPriceV1Resp - struct for UmfuturesGetTickerPriceV1Resp
type UmfuturesGetTickerPriceV1Resp struct {
	UmfuturesGetTickerPriceV1RespItem *UmfuturesGetTickerPriceV1RespItem
	ArrayOfUmfuturesGetTickerPriceV1RespItem *[]UmfuturesGetTickerPriceV1RespItem
}

// UmfuturesGetTickerPriceV1RespItemAsUmfuturesGetTickerPriceV1Resp is a convenience function that returns UmfuturesGetTickerPriceV1RespItem wrapped in UmfuturesGetTickerPriceV1Resp
func UmfuturesGetTickerPriceV1RespItemAsUmfuturesGetTickerPriceV1Resp(v *UmfuturesGetTickerPriceV1RespItem) UmfuturesGetTickerPriceV1Resp {
	return UmfuturesGetTickerPriceV1Resp{
		UmfuturesGetTickerPriceV1RespItem: v,
	}
}

// []UmfuturesGetTickerPriceV1RespItemAsUmfuturesGetTickerPriceV1Resp is a convenience function that returns []UmfuturesGetTickerPriceV1RespItem wrapped in UmfuturesGetTickerPriceV1Resp
func ArrayOfUmfuturesGetTickerPriceV1RespItemAsUmfuturesGetTickerPriceV1Resp(v *[]UmfuturesGetTickerPriceV1RespItem) UmfuturesGetTickerPriceV1Resp {
	return UmfuturesGetTickerPriceV1Resp{
		ArrayOfUmfuturesGetTickerPriceV1RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UmfuturesGetTickerPriceV1Resp) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UmfuturesGetTickerPriceV1RespItem
	err = newStrictDecoder(data).Decode(&dst.UmfuturesGetTickerPriceV1RespItem)
	if err == nil {
		jsonUmfuturesGetTickerPriceV1RespItem, _ := json.Marshal(dst.UmfuturesGetTickerPriceV1RespItem)
		if string(jsonUmfuturesGetTickerPriceV1RespItem) == "{}" { // empty struct
			dst.UmfuturesGetTickerPriceV1RespItem = nil
		} else {
			if err = validator.Validate(dst.UmfuturesGetTickerPriceV1RespItem); err != nil {
				dst.UmfuturesGetTickerPriceV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.UmfuturesGetTickerPriceV1RespItem = nil
	}

	// try to unmarshal data into ArrayOfUmfuturesGetTickerPriceV1RespItem
	err = newStrictDecoder(data).Decode(&dst.ArrayOfUmfuturesGetTickerPriceV1RespItem)
	if err == nil {
		jsonArrayOfUmfuturesGetTickerPriceV1RespItem, _ := json.Marshal(dst.ArrayOfUmfuturesGetTickerPriceV1RespItem)
		if string(jsonArrayOfUmfuturesGetTickerPriceV1RespItem) == "{}" { // empty struct
			dst.ArrayOfUmfuturesGetTickerPriceV1RespItem = nil
		} else {
			if err = validator.Validate(dst.ArrayOfUmfuturesGetTickerPriceV1RespItem); err != nil {
				dst.ArrayOfUmfuturesGetTickerPriceV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfUmfuturesGetTickerPriceV1RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UmfuturesGetTickerPriceV1RespItem = nil
		dst.ArrayOfUmfuturesGetTickerPriceV1RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UmfuturesGetTickerPriceV1Resp)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UmfuturesGetTickerPriceV1Resp)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UmfuturesGetTickerPriceV1Resp) MarshalJSON() ([]byte, error) {
	if src.UmfuturesGetTickerPriceV1RespItem != nil {
		return json.Marshal(&src.UmfuturesGetTickerPriceV1RespItem)
	}

	if src.ArrayOfUmfuturesGetTickerPriceV1RespItem != nil {
		return json.Marshal(&src.ArrayOfUmfuturesGetTickerPriceV1RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UmfuturesGetTickerPriceV1Resp) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UmfuturesGetTickerPriceV1RespItem != nil {
		return obj.UmfuturesGetTickerPriceV1RespItem
	}

	if obj.ArrayOfUmfuturesGetTickerPriceV1RespItem != nil {
		return obj.ArrayOfUmfuturesGetTickerPriceV1RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj UmfuturesGetTickerPriceV1Resp) GetActualInstanceValue() (interface{}) {
	if obj.UmfuturesGetTickerPriceV1RespItem != nil {
		return *obj.UmfuturesGetTickerPriceV1RespItem
	}

	if obj.ArrayOfUmfuturesGetTickerPriceV1RespItem != nil {
		return *obj.ArrayOfUmfuturesGetTickerPriceV1RespItem
	}

	// all schemas are nil
	return nil
}

type NullableUmfuturesGetTickerPriceV1Resp struct {
	value *UmfuturesGetTickerPriceV1Resp
	isSet bool
}

func (v NullableUmfuturesGetTickerPriceV1Resp) Get() *UmfuturesGetTickerPriceV1Resp {
	return v.value
}

func (v *NullableUmfuturesGetTickerPriceV1Resp) Set(val *UmfuturesGetTickerPriceV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUmfuturesGetTickerPriceV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUmfuturesGetTickerPriceV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmfuturesGetTickerPriceV1Resp(val *UmfuturesGetTickerPriceV1Resp) *NullableUmfuturesGetTickerPriceV1Resp {
	return &NullableUmfuturesGetTickerPriceV1Resp{value: val, isSet: true}
}

func (v NullableUmfuturesGetTickerPriceV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmfuturesGetTickerPriceV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


