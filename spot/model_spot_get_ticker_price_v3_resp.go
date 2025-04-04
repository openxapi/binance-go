/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// SpotGetTickerPriceV3Resp - struct for SpotGetTickerPriceV3Resp
type SpotGetTickerPriceV3Resp struct {
	SpotGetTickerPriceV3RespItem *SpotGetTickerPriceV3RespItem
	ArrayOfSpotGetTickerPriceV3RespItem *[]SpotGetTickerPriceV3RespItem
}

// SpotGetTickerPriceV3RespItemAsSpotGetTickerPriceV3Resp is a convenience function that returns SpotGetTickerPriceV3RespItem wrapped in SpotGetTickerPriceV3Resp
func SpotGetTickerPriceV3RespItemAsSpotGetTickerPriceV3Resp(v *SpotGetTickerPriceV3RespItem) SpotGetTickerPriceV3Resp {
	return SpotGetTickerPriceV3Resp{
		SpotGetTickerPriceV3RespItem: v,
	}
}

// []SpotGetTickerPriceV3RespItemAsSpotGetTickerPriceV3Resp is a convenience function that returns []SpotGetTickerPriceV3RespItem wrapped in SpotGetTickerPriceV3Resp
func ArrayOfSpotGetTickerPriceV3RespItemAsSpotGetTickerPriceV3Resp(v *[]SpotGetTickerPriceV3RespItem) SpotGetTickerPriceV3Resp {
	return SpotGetTickerPriceV3Resp{
		ArrayOfSpotGetTickerPriceV3RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SpotGetTickerPriceV3Resp) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SpotGetTickerPriceV3RespItem
	err = newStrictDecoder(data).Decode(&dst.SpotGetTickerPriceV3RespItem)
	if err == nil {
		jsonSpotGetTickerPriceV3RespItem, _ := json.Marshal(dst.SpotGetTickerPriceV3RespItem)
		if string(jsonSpotGetTickerPriceV3RespItem) == "{}" { // empty struct
			dst.SpotGetTickerPriceV3RespItem = nil
		} else {
			if err = validator.Validate(dst.SpotGetTickerPriceV3RespItem); err != nil {
				dst.SpotGetTickerPriceV3RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.SpotGetTickerPriceV3RespItem = nil
	}

	// try to unmarshal data into ArrayOfSpotGetTickerPriceV3RespItem
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSpotGetTickerPriceV3RespItem)
	if err == nil {
		jsonArrayOfSpotGetTickerPriceV3RespItem, _ := json.Marshal(dst.ArrayOfSpotGetTickerPriceV3RespItem)
		if string(jsonArrayOfSpotGetTickerPriceV3RespItem) == "{}" { // empty struct
			dst.ArrayOfSpotGetTickerPriceV3RespItem = nil
		} else {
			if err = validator.Validate(dst.ArrayOfSpotGetTickerPriceV3RespItem); err != nil {
				dst.ArrayOfSpotGetTickerPriceV3RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfSpotGetTickerPriceV3RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SpotGetTickerPriceV3RespItem = nil
		dst.ArrayOfSpotGetTickerPriceV3RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SpotGetTickerPriceV3Resp)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SpotGetTickerPriceV3Resp)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SpotGetTickerPriceV3Resp) MarshalJSON() ([]byte, error) {
	if src.SpotGetTickerPriceV3RespItem != nil {
		return json.Marshal(&src.SpotGetTickerPriceV3RespItem)
	}

	if src.ArrayOfSpotGetTickerPriceV3RespItem != nil {
		return json.Marshal(&src.ArrayOfSpotGetTickerPriceV3RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SpotGetTickerPriceV3Resp) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SpotGetTickerPriceV3RespItem != nil {
		return obj.SpotGetTickerPriceV3RespItem
	}

	if obj.ArrayOfSpotGetTickerPriceV3RespItem != nil {
		return obj.ArrayOfSpotGetTickerPriceV3RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SpotGetTickerPriceV3Resp) GetActualInstanceValue() (interface{}) {
	if obj.SpotGetTickerPriceV3RespItem != nil {
		return *obj.SpotGetTickerPriceV3RespItem
	}

	if obj.ArrayOfSpotGetTickerPriceV3RespItem != nil {
		return *obj.ArrayOfSpotGetTickerPriceV3RespItem
	}

	// all schemas are nil
	return nil
}

type NullableSpotGetTickerPriceV3Resp struct {
	value *SpotGetTickerPriceV3Resp
	isSet bool
}

func (v NullableSpotGetTickerPriceV3Resp) Get() *SpotGetTickerPriceV3Resp {
	return v.value
}

func (v *NullableSpotGetTickerPriceV3Resp) Set(val *SpotGetTickerPriceV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotGetTickerPriceV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotGetTickerPriceV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotGetTickerPriceV3Resp(val *SpotGetTickerPriceV3Resp) *NullableSpotGetTickerPriceV3Resp {
	return &NullableSpotGetTickerPriceV3Resp{value: val, isSet: true}
}

func (v NullableSpotGetTickerPriceV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotGetTickerPriceV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


