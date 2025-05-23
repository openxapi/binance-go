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

// SpotCreateOrderCancelReplaceV3Resp - struct for SpotCreateOrderCancelReplaceV3Resp
type SpotCreateOrderCancelReplaceV3Resp struct {
	SpotCreateOrderCancelReplaceV3Data *SpotCreateOrderCancelReplaceV3Data
	SpotCreateOrderCancelReplaceV3FailResp *SpotCreateOrderCancelReplaceV3FailResp
}

// SpotCreateOrderCancelReplaceV3DataAsSpotCreateOrderCancelReplaceV3Resp is a convenience function that returns SpotCreateOrderCancelReplaceV3Data wrapped in SpotCreateOrderCancelReplaceV3Resp
func SpotCreateOrderCancelReplaceV3DataAsSpotCreateOrderCancelReplaceV3Resp(v *SpotCreateOrderCancelReplaceV3Data) SpotCreateOrderCancelReplaceV3Resp {
	return SpotCreateOrderCancelReplaceV3Resp{
		SpotCreateOrderCancelReplaceV3Data: v,
	}
}

// SpotCreateOrderCancelReplaceV3FailRespAsSpotCreateOrderCancelReplaceV3Resp is a convenience function that returns SpotCreateOrderCancelReplaceV3FailResp wrapped in SpotCreateOrderCancelReplaceV3Resp
func SpotCreateOrderCancelReplaceV3FailRespAsSpotCreateOrderCancelReplaceV3Resp(v *SpotCreateOrderCancelReplaceV3FailResp) SpotCreateOrderCancelReplaceV3Resp {
	return SpotCreateOrderCancelReplaceV3Resp{
		SpotCreateOrderCancelReplaceV3FailResp: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SpotCreateOrderCancelReplaceV3Resp) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SpotCreateOrderCancelReplaceV3Data
	err = newStrictDecoder(data).Decode(&dst.SpotCreateOrderCancelReplaceV3Data)
	if err == nil {
		jsonSpotCreateOrderCancelReplaceV3Data, _ := json.Marshal(dst.SpotCreateOrderCancelReplaceV3Data)
		if string(jsonSpotCreateOrderCancelReplaceV3Data) == "{}" { // empty struct
			dst.SpotCreateOrderCancelReplaceV3Data = nil
		} else {
			if err = validator.Validate(dst.SpotCreateOrderCancelReplaceV3Data); err != nil {
				dst.SpotCreateOrderCancelReplaceV3Data = nil
			} else {
				match++
			}
		}
	} else {
		dst.SpotCreateOrderCancelReplaceV3Data = nil
	}

	// try to unmarshal data into SpotCreateOrderCancelReplaceV3FailResp
	err = newStrictDecoder(data).Decode(&dst.SpotCreateOrderCancelReplaceV3FailResp)
	if err == nil {
		jsonSpotCreateOrderCancelReplaceV3FailResp, _ := json.Marshal(dst.SpotCreateOrderCancelReplaceV3FailResp)
		if string(jsonSpotCreateOrderCancelReplaceV3FailResp) == "{}" { // empty struct
			dst.SpotCreateOrderCancelReplaceV3FailResp = nil
		} else {
			if err = validator.Validate(dst.SpotCreateOrderCancelReplaceV3FailResp); err != nil {
				dst.SpotCreateOrderCancelReplaceV3FailResp = nil
			} else {
				match++
			}
		}
	} else {
		dst.SpotCreateOrderCancelReplaceV3FailResp = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SpotCreateOrderCancelReplaceV3Data = nil
		dst.SpotCreateOrderCancelReplaceV3FailResp = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SpotCreateOrderCancelReplaceV3Resp)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SpotCreateOrderCancelReplaceV3Resp)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SpotCreateOrderCancelReplaceV3Resp) MarshalJSON() ([]byte, error) {
	if src.SpotCreateOrderCancelReplaceV3Data != nil {
		return json.Marshal(&src.SpotCreateOrderCancelReplaceV3Data)
	}

	if src.SpotCreateOrderCancelReplaceV3FailResp != nil {
		return json.Marshal(&src.SpotCreateOrderCancelReplaceV3FailResp)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SpotCreateOrderCancelReplaceV3Resp) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SpotCreateOrderCancelReplaceV3Data != nil {
		return obj.SpotCreateOrderCancelReplaceV3Data
	}

	if obj.SpotCreateOrderCancelReplaceV3FailResp != nil {
		return obj.SpotCreateOrderCancelReplaceV3FailResp
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SpotCreateOrderCancelReplaceV3Resp) GetActualInstanceValue() (interface{}) {
	if obj.SpotCreateOrderCancelReplaceV3Data != nil {
		return *obj.SpotCreateOrderCancelReplaceV3Data
	}

	if obj.SpotCreateOrderCancelReplaceV3FailResp != nil {
		return *obj.SpotCreateOrderCancelReplaceV3FailResp
	}

	// all schemas are nil
	return nil
}

type NullableSpotCreateOrderCancelReplaceV3Resp struct {
	value *SpotCreateOrderCancelReplaceV3Resp
	isSet bool
}

func (v NullableSpotCreateOrderCancelReplaceV3Resp) Get() *SpotCreateOrderCancelReplaceV3Resp {
	return v.value
}

func (v *NullableSpotCreateOrderCancelReplaceV3Resp) Set(val *SpotCreateOrderCancelReplaceV3Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotCreateOrderCancelReplaceV3Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotCreateOrderCancelReplaceV3Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotCreateOrderCancelReplaceV3Resp(val *SpotCreateOrderCancelReplaceV3Resp) *NullableSpotCreateOrderCancelReplaceV3Resp {
	return &NullableSpotCreateOrderCancelReplaceV3Resp{value: val, isSet: true}
}

func (v NullableSpotCreateOrderCancelReplaceV3Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotCreateOrderCancelReplaceV3Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


