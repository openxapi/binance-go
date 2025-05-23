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

// checks if the GetCapitalConfigGetallV1RespItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCapitalConfigGetallV1RespItem{}

// GetCapitalConfigGetallV1RespItem struct for GetCapitalConfigGetallV1RespItem
type GetCapitalConfigGetallV1RespItem struct {
	Coin *string `json:"coin,omitempty"`
	DepositAllEnable *bool `json:"depositAllEnable,omitempty"`
	Free *string `json:"free,omitempty"`
	Freeze *string `json:"freeze,omitempty"`
	Ipoable *string `json:"ipoable,omitempty"`
	Ipoing *string `json:"ipoing,omitempty"`
	IsLegalMoney *bool `json:"isLegalMoney,omitempty"`
	Locked *string `json:"locked,omitempty"`
	Name *string `json:"name,omitempty"`
	NetworkList []GetCapitalConfigGetallV1RespItemNetworkListInner `json:"networkList,omitempty"`
	Storage *string `json:"storage,omitempty"`
	Trading *bool `json:"trading,omitempty"`
	WithdrawAllEnable *bool `json:"withdrawAllEnable,omitempty"`
	Withdrawing *string `json:"withdrawing,omitempty"`
}

// NewGetCapitalConfigGetallV1RespItem instantiates a new GetCapitalConfigGetallV1RespItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCapitalConfigGetallV1RespItem() *GetCapitalConfigGetallV1RespItem {
	this := GetCapitalConfigGetallV1RespItem{}
	return &this
}

// NewGetCapitalConfigGetallV1RespItemWithDefaults instantiates a new GetCapitalConfigGetallV1RespItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCapitalConfigGetallV1RespItemWithDefaults() *GetCapitalConfigGetallV1RespItem {
	this := GetCapitalConfigGetallV1RespItem{}
	return &this
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetCoin() string {
	if o == nil || IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetCoinOk() (*string, bool) {
	if o == nil || IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasCoin() bool {
	if o != nil && !IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetCapitalConfigGetallV1RespItem) SetCoin(v string) {
	o.Coin = &v
}

// GetDepositAllEnable returns the DepositAllEnable field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetDepositAllEnable() bool {
	if o == nil || IsNil(o.DepositAllEnable) {
		var ret bool
		return ret
	}
	return *o.DepositAllEnable
}

// GetDepositAllEnableOk returns a tuple with the DepositAllEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetDepositAllEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.DepositAllEnable) {
		return nil, false
	}
	return o.DepositAllEnable, true
}

// HasDepositAllEnable returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasDepositAllEnable() bool {
	if o != nil && !IsNil(o.DepositAllEnable) {
		return true
	}

	return false
}

// SetDepositAllEnable gets a reference to the given bool and assigns it to the DepositAllEnable field.
func (o *GetCapitalConfigGetallV1RespItem) SetDepositAllEnable(v bool) {
	o.DepositAllEnable = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetFree() string {
	if o == nil || IsNil(o.Free) {
		var ret string
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetFreeOk() (*string, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasFree() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given string and assigns it to the Free field.
func (o *GetCapitalConfigGetallV1RespItem) SetFree(v string) {
	o.Free = &v
}

// GetFreeze returns the Freeze field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetFreeze() string {
	if o == nil || IsNil(o.Freeze) {
		var ret string
		return ret
	}
	return *o.Freeze
}

// GetFreezeOk returns a tuple with the Freeze field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetFreezeOk() (*string, bool) {
	if o == nil || IsNil(o.Freeze) {
		return nil, false
	}
	return o.Freeze, true
}

// HasFreeze returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasFreeze() bool {
	if o != nil && !IsNil(o.Freeze) {
		return true
	}

	return false
}

// SetFreeze gets a reference to the given string and assigns it to the Freeze field.
func (o *GetCapitalConfigGetallV1RespItem) SetFreeze(v string) {
	o.Freeze = &v
}

// GetIpoable returns the Ipoable field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetIpoable() string {
	if o == nil || IsNil(o.Ipoable) {
		var ret string
		return ret
	}
	return *o.Ipoable
}

// GetIpoableOk returns a tuple with the Ipoable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetIpoableOk() (*string, bool) {
	if o == nil || IsNil(o.Ipoable) {
		return nil, false
	}
	return o.Ipoable, true
}

// HasIpoable returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasIpoable() bool {
	if o != nil && !IsNil(o.Ipoable) {
		return true
	}

	return false
}

// SetIpoable gets a reference to the given string and assigns it to the Ipoable field.
func (o *GetCapitalConfigGetallV1RespItem) SetIpoable(v string) {
	o.Ipoable = &v
}

// GetIpoing returns the Ipoing field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetIpoing() string {
	if o == nil || IsNil(o.Ipoing) {
		var ret string
		return ret
	}
	return *o.Ipoing
}

// GetIpoingOk returns a tuple with the Ipoing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetIpoingOk() (*string, bool) {
	if o == nil || IsNil(o.Ipoing) {
		return nil, false
	}
	return o.Ipoing, true
}

// HasIpoing returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasIpoing() bool {
	if o != nil && !IsNil(o.Ipoing) {
		return true
	}

	return false
}

// SetIpoing gets a reference to the given string and assigns it to the Ipoing field.
func (o *GetCapitalConfigGetallV1RespItem) SetIpoing(v string) {
	o.Ipoing = &v
}

// GetIsLegalMoney returns the IsLegalMoney field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetIsLegalMoney() bool {
	if o == nil || IsNil(o.IsLegalMoney) {
		var ret bool
		return ret
	}
	return *o.IsLegalMoney
}

// GetIsLegalMoneyOk returns a tuple with the IsLegalMoney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetIsLegalMoneyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLegalMoney) {
		return nil, false
	}
	return o.IsLegalMoney, true
}

// HasIsLegalMoney returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasIsLegalMoney() bool {
	if o != nil && !IsNil(o.IsLegalMoney) {
		return true
	}

	return false
}

// SetIsLegalMoney gets a reference to the given bool and assigns it to the IsLegalMoney field.
func (o *GetCapitalConfigGetallV1RespItem) SetIsLegalMoney(v bool) {
	o.IsLegalMoney = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetLocked() string {
	if o == nil || IsNil(o.Locked) {
		var ret string
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetLockedOk() (*string, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given string and assigns it to the Locked field.
func (o *GetCapitalConfigGetallV1RespItem) SetLocked(v string) {
	o.Locked = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetCapitalConfigGetallV1RespItem) SetName(v string) {
	o.Name = &v
}

// GetNetworkList returns the NetworkList field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetNetworkList() []GetCapitalConfigGetallV1RespItemNetworkListInner {
	if o == nil || IsNil(o.NetworkList) {
		var ret []GetCapitalConfigGetallV1RespItemNetworkListInner
		return ret
	}
	return o.NetworkList
}

// GetNetworkListOk returns a tuple with the NetworkList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetNetworkListOk() ([]GetCapitalConfigGetallV1RespItemNetworkListInner, bool) {
	if o == nil || IsNil(o.NetworkList) {
		return nil, false
	}
	return o.NetworkList, true
}

// HasNetworkList returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasNetworkList() bool {
	if o != nil && !IsNil(o.NetworkList) {
		return true
	}

	return false
}

// SetNetworkList gets a reference to the given []GetCapitalConfigGetallV1RespItemNetworkListInner and assigns it to the NetworkList field.
func (o *GetCapitalConfigGetallV1RespItem) SetNetworkList(v []GetCapitalConfigGetallV1RespItemNetworkListInner) {
	o.NetworkList = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetStorage() string {
	if o == nil || IsNil(o.Storage) {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetStorageOk() (*string, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *GetCapitalConfigGetallV1RespItem) SetStorage(v string) {
	o.Storage = &v
}

// GetTrading returns the Trading field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetTrading() bool {
	if o == nil || IsNil(o.Trading) {
		var ret bool
		return ret
	}
	return *o.Trading
}

// GetTradingOk returns a tuple with the Trading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetTradingOk() (*bool, bool) {
	if o == nil || IsNil(o.Trading) {
		return nil, false
	}
	return o.Trading, true
}

// HasTrading returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasTrading() bool {
	if o != nil && !IsNil(o.Trading) {
		return true
	}

	return false
}

// SetTrading gets a reference to the given bool and assigns it to the Trading field.
func (o *GetCapitalConfigGetallV1RespItem) SetTrading(v bool) {
	o.Trading = &v
}

// GetWithdrawAllEnable returns the WithdrawAllEnable field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetWithdrawAllEnable() bool {
	if o == nil || IsNil(o.WithdrawAllEnable) {
		var ret bool
		return ret
	}
	return *o.WithdrawAllEnable
}

// GetWithdrawAllEnableOk returns a tuple with the WithdrawAllEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetWithdrawAllEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.WithdrawAllEnable) {
		return nil, false
	}
	return o.WithdrawAllEnable, true
}

// HasWithdrawAllEnable returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasWithdrawAllEnable() bool {
	if o != nil && !IsNil(o.WithdrawAllEnable) {
		return true
	}

	return false
}

// SetWithdrawAllEnable gets a reference to the given bool and assigns it to the WithdrawAllEnable field.
func (o *GetCapitalConfigGetallV1RespItem) SetWithdrawAllEnable(v bool) {
	o.WithdrawAllEnable = &v
}

// GetWithdrawing returns the Withdrawing field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItem) GetWithdrawing() string {
	if o == nil || IsNil(o.Withdrawing) {
		var ret string
		return ret
	}
	return *o.Withdrawing
}

// GetWithdrawingOk returns a tuple with the Withdrawing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItem) GetWithdrawingOk() (*string, bool) {
	if o == nil || IsNil(o.Withdrawing) {
		return nil, false
	}
	return o.Withdrawing, true
}

// HasWithdrawing returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItem) HasWithdrawing() bool {
	if o != nil && !IsNil(o.Withdrawing) {
		return true
	}

	return false
}

// SetWithdrawing gets a reference to the given string and assigns it to the Withdrawing field.
func (o *GetCapitalConfigGetallV1RespItem) SetWithdrawing(v string) {
	o.Withdrawing = &v
}

func (o GetCapitalConfigGetallV1RespItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCapitalConfigGetallV1RespItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !IsNil(o.DepositAllEnable) {
		toSerialize["depositAllEnable"] = o.DepositAllEnable
	}
	if !IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !IsNil(o.Freeze) {
		toSerialize["freeze"] = o.Freeze
	}
	if !IsNil(o.Ipoable) {
		toSerialize["ipoable"] = o.Ipoable
	}
	if !IsNil(o.Ipoing) {
		toSerialize["ipoing"] = o.Ipoing
	}
	if !IsNil(o.IsLegalMoney) {
		toSerialize["isLegalMoney"] = o.IsLegalMoney
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkList) {
		toSerialize["networkList"] = o.NetworkList
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.Trading) {
		toSerialize["trading"] = o.Trading
	}
	if !IsNil(o.WithdrawAllEnable) {
		toSerialize["withdrawAllEnable"] = o.WithdrawAllEnable
	}
	if !IsNil(o.Withdrawing) {
		toSerialize["withdrawing"] = o.Withdrawing
	}
	return toSerialize, nil
}

type NullableGetCapitalConfigGetallV1RespItem struct {
	value *GetCapitalConfigGetallV1RespItem
	isSet bool
}

func (v NullableGetCapitalConfigGetallV1RespItem) Get() *GetCapitalConfigGetallV1RespItem {
	return v.value
}

func (v *NullableGetCapitalConfigGetallV1RespItem) Set(val *GetCapitalConfigGetallV1RespItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCapitalConfigGetallV1RespItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCapitalConfigGetallV1RespItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCapitalConfigGetallV1RespItem(val *GetCapitalConfigGetallV1RespItem) *NullableGetCapitalConfigGetallV1RespItem {
	return &NullableGetCapitalConfigGetallV1RespItem{value: val, isSet: true}
}

func (v NullableGetCapitalConfigGetallV1RespItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCapitalConfigGetallV1RespItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


