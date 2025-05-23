/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"encoding/json"
)

// checks if the GetLoanRepayHistoryV1RespRowsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLoanRepayHistoryV1RespRowsInner{}

// GetLoanRepayHistoryV1RespRowsInner struct for GetLoanRepayHistoryV1RespRowsInner
type GetLoanRepayHistoryV1RespRowsInner struct {
	CollateralCoin *string `json:"collateralCoin,omitempty"`
	CollateralReturn *string `json:"collateralReturn,omitempty"`
	CollateralUsed *string `json:"collateralUsed,omitempty"`
	LoanCoin *string `json:"loanCoin,omitempty"`
	OrderId *int64 `json:"orderId,omitempty"`
	RepayAmount *string `json:"repayAmount,omitempty"`
	RepayStatus *string `json:"repayStatus,omitempty"`
	RepayTime *int64 `json:"repayTime,omitempty"`
	RepayType *string `json:"repayType,omitempty"`
}

// NewGetLoanRepayHistoryV1RespRowsInner instantiates a new GetLoanRepayHistoryV1RespRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLoanRepayHistoryV1RespRowsInner() *GetLoanRepayHistoryV1RespRowsInner {
	this := GetLoanRepayHistoryV1RespRowsInner{}
	return &this
}

// NewGetLoanRepayHistoryV1RespRowsInnerWithDefaults instantiates a new GetLoanRepayHistoryV1RespRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLoanRepayHistoryV1RespRowsInnerWithDefaults() *GetLoanRepayHistoryV1RespRowsInner {
	this := GetLoanRepayHistoryV1RespRowsInner{}
	return &this
}

// GetCollateralCoin returns the CollateralCoin field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetCollateralCoin() string {
	if o == nil || IsNil(o.CollateralCoin) {
		var ret string
		return ret
	}
	return *o.CollateralCoin
}

// GetCollateralCoinOk returns a tuple with the CollateralCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetCollateralCoinOk() (*string, bool) {
	if o == nil || IsNil(o.CollateralCoin) {
		return nil, false
	}
	return o.CollateralCoin, true
}

// HasCollateralCoin returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasCollateralCoin() bool {
	if o != nil && !IsNil(o.CollateralCoin) {
		return true
	}

	return false
}

// SetCollateralCoin gets a reference to the given string and assigns it to the CollateralCoin field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetCollateralCoin(v string) {
	o.CollateralCoin = &v
}

// GetCollateralReturn returns the CollateralReturn field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetCollateralReturn() string {
	if o == nil || IsNil(o.CollateralReturn) {
		var ret string
		return ret
	}
	return *o.CollateralReturn
}

// GetCollateralReturnOk returns a tuple with the CollateralReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetCollateralReturnOk() (*string, bool) {
	if o == nil || IsNil(o.CollateralReturn) {
		return nil, false
	}
	return o.CollateralReturn, true
}

// HasCollateralReturn returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasCollateralReturn() bool {
	if o != nil && !IsNil(o.CollateralReturn) {
		return true
	}

	return false
}

// SetCollateralReturn gets a reference to the given string and assigns it to the CollateralReturn field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetCollateralReturn(v string) {
	o.CollateralReturn = &v
}

// GetCollateralUsed returns the CollateralUsed field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetCollateralUsed() string {
	if o == nil || IsNil(o.CollateralUsed) {
		var ret string
		return ret
	}
	return *o.CollateralUsed
}

// GetCollateralUsedOk returns a tuple with the CollateralUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetCollateralUsedOk() (*string, bool) {
	if o == nil || IsNil(o.CollateralUsed) {
		return nil, false
	}
	return o.CollateralUsed, true
}

// HasCollateralUsed returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasCollateralUsed() bool {
	if o != nil && !IsNil(o.CollateralUsed) {
		return true
	}

	return false
}

// SetCollateralUsed gets a reference to the given string and assigns it to the CollateralUsed field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetCollateralUsed(v string) {
	o.CollateralUsed = &v
}

// GetLoanCoin returns the LoanCoin field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetLoanCoin() string {
	if o == nil || IsNil(o.LoanCoin) {
		var ret string
		return ret
	}
	return *o.LoanCoin
}

// GetLoanCoinOk returns a tuple with the LoanCoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetLoanCoinOk() (*string, bool) {
	if o == nil || IsNil(o.LoanCoin) {
		return nil, false
	}
	return o.LoanCoin, true
}

// HasLoanCoin returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasLoanCoin() bool {
	if o != nil && !IsNil(o.LoanCoin) {
		return true
	}

	return false
}

// SetLoanCoin gets a reference to the given string and assigns it to the LoanCoin field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetLoanCoin(v string) {
	o.LoanCoin = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetOrderId() int64 {
	if o == nil || IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetRepayAmount returns the RepayAmount field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayAmount() string {
	if o == nil || IsNil(o.RepayAmount) {
		var ret string
		return ret
	}
	return *o.RepayAmount
}

// GetRepayAmountOk returns a tuple with the RepayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.RepayAmount) {
		return nil, false
	}
	return o.RepayAmount, true
}

// HasRepayAmount returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasRepayAmount() bool {
	if o != nil && !IsNil(o.RepayAmount) {
		return true
	}

	return false
}

// SetRepayAmount gets a reference to the given string and assigns it to the RepayAmount field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetRepayAmount(v string) {
	o.RepayAmount = &v
}

// GetRepayStatus returns the RepayStatus field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayStatus() string {
	if o == nil || IsNil(o.RepayStatus) {
		var ret string
		return ret
	}
	return *o.RepayStatus
}

// GetRepayStatusOk returns a tuple with the RepayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RepayStatus) {
		return nil, false
	}
	return o.RepayStatus, true
}

// HasRepayStatus returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasRepayStatus() bool {
	if o != nil && !IsNil(o.RepayStatus) {
		return true
	}

	return false
}

// SetRepayStatus gets a reference to the given string and assigns it to the RepayStatus field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetRepayStatus(v string) {
	o.RepayStatus = &v
}

// GetRepayTime returns the RepayTime field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayTime() int64 {
	if o == nil || IsNil(o.RepayTime) {
		var ret int64
		return ret
	}
	return *o.RepayTime
}

// GetRepayTimeOk returns a tuple with the RepayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.RepayTime) {
		return nil, false
	}
	return o.RepayTime, true
}

// HasRepayTime returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasRepayTime() bool {
	if o != nil && !IsNil(o.RepayTime) {
		return true
	}

	return false
}

// SetRepayTime gets a reference to the given int64 and assigns it to the RepayTime field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetRepayTime(v int64) {
	o.RepayTime = &v
}

// GetRepayType returns the RepayType field value if set, zero value otherwise.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayType() string {
	if o == nil || IsNil(o.RepayType) {
		var ret string
		return ret
	}
	return *o.RepayType
}

// GetRepayTypeOk returns a tuple with the RepayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) GetRepayTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RepayType) {
		return nil, false
	}
	return o.RepayType, true
}

// HasRepayType returns a boolean if a field has been set.
func (o *GetLoanRepayHistoryV1RespRowsInner) HasRepayType() bool {
	if o != nil && !IsNil(o.RepayType) {
		return true
	}

	return false
}

// SetRepayType gets a reference to the given string and assigns it to the RepayType field.
func (o *GetLoanRepayHistoryV1RespRowsInner) SetRepayType(v string) {
	o.RepayType = &v
}

func (o GetLoanRepayHistoryV1RespRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLoanRepayHistoryV1RespRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollateralCoin) {
		toSerialize["collateralCoin"] = o.CollateralCoin
	}
	if !IsNil(o.CollateralReturn) {
		toSerialize["collateralReturn"] = o.CollateralReturn
	}
	if !IsNil(o.CollateralUsed) {
		toSerialize["collateralUsed"] = o.CollateralUsed
	}
	if !IsNil(o.LoanCoin) {
		toSerialize["loanCoin"] = o.LoanCoin
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.RepayAmount) {
		toSerialize["repayAmount"] = o.RepayAmount
	}
	if !IsNil(o.RepayStatus) {
		toSerialize["repayStatus"] = o.RepayStatus
	}
	if !IsNil(o.RepayTime) {
		toSerialize["repayTime"] = o.RepayTime
	}
	if !IsNil(o.RepayType) {
		toSerialize["repayType"] = o.RepayType
	}
	return toSerialize, nil
}

type NullableGetLoanRepayHistoryV1RespRowsInner struct {
	value *GetLoanRepayHistoryV1RespRowsInner
	isSet bool
}

func (v NullableGetLoanRepayHistoryV1RespRowsInner) Get() *GetLoanRepayHistoryV1RespRowsInner {
	return v.value
}

func (v *NullableGetLoanRepayHistoryV1RespRowsInner) Set(val *GetLoanRepayHistoryV1RespRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLoanRepayHistoryV1RespRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLoanRepayHistoryV1RespRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLoanRepayHistoryV1RespRowsInner(val *GetLoanRepayHistoryV1RespRowsInner) *NullableGetLoanRepayHistoryV1RespRowsInner {
	return &NullableGetLoanRepayHistoryV1RespRowsInner{value: val, isSet: true}
}

func (v NullableGetLoanRepayHistoryV1RespRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLoanRepayHistoryV1RespRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


