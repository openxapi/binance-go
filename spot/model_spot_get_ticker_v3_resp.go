/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// SpotGetTickerV3Resp - struct for SpotGetTickerV3Resp
type SpotGetTickerV3Resp struct {
	SpotGetTickerV3RespItem *SpotGetTickerV3RespItem
	ArrayOfSpotGetTickerV3RespItem *[]SpotGetTickerV3RespItem
}

// SpotGetTickerV3RespItemAsSpotGetTickerV3Resp is a convenience function that returns SpotGetTickerV3RespItem wrapped in SpotGetTickerV3Resp
func SpotGetTickerV3RespItemAsSpotGetTickerV3Resp(v *SpotGetTickerV3RespItem) SpotGetTickerV3Resp {
	return SpotGetTickerV3Resp{
		SpotGetTickerV3RespItem: v,
	}
}

// []SpotGetTickerV3RespItemAsSpotGetTickerV3Resp is a convenience function that returns []SpotGetTickerV3RespItem wrapped in SpotGetTickerV3Resp
func ArrayOfSpotGetTickerV3RespItemAsSpotGetTickerV3Resp(v *[]SpotGetTickerV3RespItem) SpotGetTickerV3Resp {
	return SpotGetTickerV3Resp{
		ArrayOfSpotGetTickerV3RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SpotGetTickerV3Resp) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SpotGetTickerV3RespItem
	err = newStrictDecoder(data).Decode(&dst.SpotGetTickerV3RespItem)
	if err == nil {
		jsonSpotGetTickerV3RespItem, _ := json.Marshal(dst.SpotGetTickerV3RespItem)
		if string(jsonSpotGetTickerV3RespItem) == "{}" { // empty struct
			dst.SpotGetTickerV3RespItem = nil
		} else {
			if err = validator.Validate(dst.SpotGetTickerV3RespItem); err != nil {
				dst.SpotGetTickerV3RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.SpotGetTickerV3RespItem = nil
	}

	// try to unmarshal data into ArrayOfSpotGetTickerV3RespItem
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSpotGetTickerV3RespItem)
	if err == nil {
		jsonArrayOfSpotGetTickerV3RespItem, _ := json.Marshal(dst.ArrayOfSpotGetTickerV3RespItem)
		if string(jsonArrayOfSpotGetTickerV3RespItem) == "{}" { // empty struct
			dst.ArrayOfSpotGetTickerV3RespItem = nil
		} else {
			if err = validator.Validate(dst.ArrayOfSpotGetTickerV3RespItem); err != nil {
				dst.ArrayOfSpotGetTickerV3RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfSpotGetTickerV3RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SpotGetTickerV3RespItem = nil
		dst.ArrayOfSpotGetTickerV3RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SpotGetTickerV3Resp)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SpotGetTickerV3Resp)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SpotGetTickerV3Resp) MarshalJSON() ([]byte, error) {
	if src.SpotGetTickerV3RespItem != nil {
		return json.Marshal(&src.SpotGetTickerV3RespItem)
	}

	if src.ArrayOfSpotGetTickerV3RespItem != nil {
		return json.Marshal(&src.ArrayOfSpotGetTickerV3RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SpotGetTickerV3Resp) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SpotGetTickerV3RespItem != nil {
		return obj.SpotGetTickerV3RespItem
	}

	if obj.ArrayOfSpotGetTickerV3RespItem != nil {
		return obj.ArrayOfSpotGetTickerV3RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SpotGetTickerV3Resp) GetActualInstanceValue() (interface{}) {
	if obj.SpotGetTickerV3RespItem != nil {
		return *obj.SpotGetTickerV3RespItem
	}

	if obj.ArrayOfSpotGetTickerV3RespItem != nil {
		return *obj.ArrayOfSpotGetTickerV3RespItem
	}

	// all schemas are nil
	return nil
}

type NullableSpotGetTickerV3Resp struct {
	value *SpotGetTickerV3Resp
	isSet bool
}

func (v NullableSpotGetTickerV3Resp) Get() *SpotGetTickerV3Resp {
	return v.value
}

func (v *NullableSpotGetTickerV3Resp) Set(val *SpotGetTickerV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotGetTickerV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotGetTickerV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotGetTickerV3Resp(val *SpotGetTickerV3Resp) *NullableSpotGetTickerV3Resp {
	return &NullableSpotGetTickerV3Resp{value: val, isSet: true}
}

func (v NullableSpotGetTickerV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotGetTickerV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


