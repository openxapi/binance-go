/*
Binance Cmfutures API

OpenAPI specification for Binance cryptocurrency exchange - Cmfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// CmfuturesDeleteBatchOrdersV1RespInner - struct for CmfuturesDeleteBatchOrdersV1RespInner
type CmfuturesDeleteBatchOrdersV1RespInner struct {
	APIError *APIError
	CmfuturesDeleteBatchOrdersV1RespItem *CmfuturesDeleteBatchOrdersV1RespItem
}

// APIErrorAsCmfuturesDeleteBatchOrdersV1RespInner is a convenience function that returns APIError wrapped in CmfuturesDeleteBatchOrdersV1RespInner
func APIErrorAsCmfuturesDeleteBatchOrdersV1RespInner(v *APIError) CmfuturesDeleteBatchOrdersV1RespInner {
	return CmfuturesDeleteBatchOrdersV1RespInner{
		APIError: v,
	}
}

// CmfuturesDeleteBatchOrdersV1RespItemAsCmfuturesDeleteBatchOrdersV1RespInner is a convenience function that returns CmfuturesDeleteBatchOrdersV1RespItem wrapped in CmfuturesDeleteBatchOrdersV1RespInner
func CmfuturesDeleteBatchOrdersV1RespItemAsCmfuturesDeleteBatchOrdersV1RespInner(v *CmfuturesDeleteBatchOrdersV1RespItem) CmfuturesDeleteBatchOrdersV1RespInner {
	return CmfuturesDeleteBatchOrdersV1RespInner{
		CmfuturesDeleteBatchOrdersV1RespItem: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CmfuturesDeleteBatchOrdersV1RespInner) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into CmfuturesDeleteBatchOrdersV1RespItem
	err = newStrictDecoder(data).Decode(&dst.CmfuturesDeleteBatchOrdersV1RespItem)
	if err == nil {
		jsonCmfuturesDeleteBatchOrdersV1RespItem, _ := json.Marshal(dst.CmfuturesDeleteBatchOrdersV1RespItem)
		if string(jsonCmfuturesDeleteBatchOrdersV1RespItem) == "{}" { // empty struct
			dst.CmfuturesDeleteBatchOrdersV1RespItem = nil
		} else {
			if err = validator.Validate(dst.CmfuturesDeleteBatchOrdersV1RespItem); err != nil {
				dst.CmfuturesDeleteBatchOrdersV1RespItem = nil
			} else {
				match++
			}
		}
	} else {
		dst.CmfuturesDeleteBatchOrdersV1RespItem = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.APIError = nil
		dst.CmfuturesDeleteBatchOrdersV1RespItem = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CmfuturesDeleteBatchOrdersV1RespInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CmfuturesDeleteBatchOrdersV1RespInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CmfuturesDeleteBatchOrdersV1RespInner) MarshalJSON() ([]byte, error) {
	if src.APIError != nil {
		return json.Marshal(&src.APIError)
	}

	if src.CmfuturesDeleteBatchOrdersV1RespItem != nil {
		return json.Marshal(&src.CmfuturesDeleteBatchOrdersV1RespItem)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CmfuturesDeleteBatchOrdersV1RespInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.APIError != nil {
		return obj.APIError
	}

	if obj.CmfuturesDeleteBatchOrdersV1RespItem != nil {
		return obj.CmfuturesDeleteBatchOrdersV1RespItem
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CmfuturesDeleteBatchOrdersV1RespInner) GetActualInstanceValue() (interface{}) {
	if obj.APIError != nil {
		return *obj.APIError
	}

	if obj.CmfuturesDeleteBatchOrdersV1RespItem != nil {
		return *obj.CmfuturesDeleteBatchOrdersV1RespItem
	}

	// all schemas are nil
	return nil
}

type NullableCmfuturesDeleteBatchOrdersV1RespInner struct {
	value *CmfuturesDeleteBatchOrdersV1RespInner
	isSet bool
}

func (v NullableCmfuturesDeleteBatchOrdersV1RespInner) Get() *CmfuturesDeleteBatchOrdersV1RespInner {
	return v.value
}

func (v *NullableCmfuturesDeleteBatchOrdersV1RespInner) Set(val *CmfuturesDeleteBatchOrdersV1RespInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCmfuturesDeleteBatchOrdersV1RespInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCmfuturesDeleteBatchOrdersV1RespInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmfuturesDeleteBatchOrdersV1RespInner(val *CmfuturesDeleteBatchOrdersV1RespInner) *NullableCmfuturesDeleteBatchOrdersV1RespInner {
	return &NullableCmfuturesDeleteBatchOrdersV1RespInner{value: val, isSet: true}
}

func (v NullableCmfuturesDeleteBatchOrdersV1RespInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmfuturesDeleteBatchOrdersV1RespInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


