/*
Binance Options API

OpenAPI specification for Binance exchange - Options API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"encoding/json"
)

// checks if the OptionsCreateBlockOrderCreateV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsCreateBlockOrderCreateV1Resp{}

// OptionsCreateBlockOrderCreateV1Resp struct for OptionsCreateBlockOrderCreateV1Resp
type OptionsCreateBlockOrderCreateV1Resp struct {
	BlockTradeSettlementKey *string `json:"blockTradeSettlementKey,omitempty"`
	ExpireTime *int64 `json:"expireTime,omitempty"`
	Legs []CreateBlockOrderExecuteV1RespLegsInner `json:"legs,omitempty"`
	Liquidity *string `json:"liquidity,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewOptionsCreateBlockOrderCreateV1Resp instantiates a new OptionsCreateBlockOrderCreateV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsCreateBlockOrderCreateV1Resp() *OptionsCreateBlockOrderCreateV1Resp {
	this := OptionsCreateBlockOrderCreateV1Resp{}
	return &this
}

// NewOptionsCreateBlockOrderCreateV1RespWithDefaults instantiates a new OptionsCreateBlockOrderCreateV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsCreateBlockOrderCreateV1RespWithDefaults() *OptionsCreateBlockOrderCreateV1Resp {
	this := OptionsCreateBlockOrderCreateV1Resp{}
	return &this
}

// GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field value if set, zero value otherwise.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetBlockTradeSettlementKey() string {
	if o == nil || IsNil(o.BlockTradeSettlementKey) {
		var ret string
		return ret
	}
	return *o.BlockTradeSettlementKey
}

// GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetBlockTradeSettlementKeyOk() (*string, bool) {
	if o == nil || IsNil(o.BlockTradeSettlementKey) {
		return nil, false
	}
	return o.BlockTradeSettlementKey, true
}

// HasBlockTradeSettlementKey returns a boolean if a field has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) HasBlockTradeSettlementKey() bool {
	if o != nil && !IsNil(o.BlockTradeSettlementKey) {
		return true
	}

	return false
}

// SetBlockTradeSettlementKey gets a reference to the given string and assigns it to the BlockTradeSettlementKey field.
func (o *OptionsCreateBlockOrderCreateV1Resp) SetBlockTradeSettlementKey(v string) {
	o.BlockTradeSettlementKey = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetExpireTime() int64 {
	if o == nil || IsNil(o.ExpireTime) {
		var ret int64
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetExpireTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given int64 and assigns it to the ExpireTime field.
func (o *OptionsCreateBlockOrderCreateV1Resp) SetExpireTime(v int64) {
	o.ExpireTime = &v
}

// GetLegs returns the Legs field value if set, zero value otherwise.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetLegs() []CreateBlockOrderExecuteV1RespLegsInner {
	if o == nil || IsNil(o.Legs) {
		var ret []CreateBlockOrderExecuteV1RespLegsInner
		return ret
	}
	return o.Legs
}

// GetLegsOk returns a tuple with the Legs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetLegsOk() ([]CreateBlockOrderExecuteV1RespLegsInner, bool) {
	if o == nil || IsNil(o.Legs) {
		return nil, false
	}
	return o.Legs, true
}

// HasLegs returns a boolean if a field has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) HasLegs() bool {
	if o != nil && !IsNil(o.Legs) {
		return true
	}

	return false
}

// SetLegs gets a reference to the given []CreateBlockOrderExecuteV1RespLegsInner and assigns it to the Legs field.
func (o *OptionsCreateBlockOrderCreateV1Resp) SetLegs(v []CreateBlockOrderExecuteV1RespLegsInner) {
	o.Legs = v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetLiquidity() string {
	if o == nil || IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetLiquidityOk() (*string, bool) {
	if o == nil || IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) HasLiquidity() bool {
	if o != nil && !IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *OptionsCreateBlockOrderCreateV1Resp) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OptionsCreateBlockOrderCreateV1Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OptionsCreateBlockOrderCreateV1Resp) SetStatus(v string) {
	o.Status = &v
}

func (o OptionsCreateBlockOrderCreateV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsCreateBlockOrderCreateV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockTradeSettlementKey) {
		toSerialize["blockTradeSettlementKey"] = o.BlockTradeSettlementKey
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expireTime"] = o.ExpireTime
	}
	if !IsNil(o.Legs) {
		toSerialize["legs"] = o.Legs
	}
	if !IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableOptionsCreateBlockOrderCreateV1Resp struct {
	value *OptionsCreateBlockOrderCreateV1Resp
	isSet bool
}

func (v NullableOptionsCreateBlockOrderCreateV1Resp) Get() *OptionsCreateBlockOrderCreateV1Resp {
	return v.value
}

func (v *NullableOptionsCreateBlockOrderCreateV1Resp) Set(val *OptionsCreateBlockOrderCreateV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsCreateBlockOrderCreateV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsCreateBlockOrderCreateV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsCreateBlockOrderCreateV1Resp(val *OptionsCreateBlockOrderCreateV1Resp) *NullableOptionsCreateBlockOrderCreateV1Resp {
	return &NullableOptionsCreateBlockOrderCreateV1Resp{value: val, isSet: true}
}

func (v NullableOptionsCreateBlockOrderCreateV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsCreateBlockOrderCreateV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


