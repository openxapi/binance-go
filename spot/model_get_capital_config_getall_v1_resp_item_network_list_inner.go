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

// checks if the GetCapitalConfigGetallV1RespItemNetworkListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCapitalConfigGetallV1RespItemNetworkListInner{}

// GetCapitalConfigGetallV1RespItemNetworkListInner struct for GetCapitalConfigGetallV1RespItemNetworkListInner
type GetCapitalConfigGetallV1RespItemNetworkListInner struct {
	AddressRegex *string `json:"addressRegex,omitempty"`
	Busy *bool `json:"busy,omitempty"`
	Coin *string `json:"coin,omitempty"`
	ContractAddress *string `json:"contractAddress,omitempty"`
	ContractAddressUrl *string `json:"contractAddressUrl,omitempty"`
	DepositDesc *string `json:"depositDesc,omitempty"`
	DepositEnable *bool `json:"depositEnable,omitempty"`
	EstimatedArrivalTime *int64 `json:"estimatedArrivalTime,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	MemoRegex *string `json:"memoRegex,omitempty"`
	MinConfirm *int32 `json:"minConfirm,omitempty"`
	Name *string `json:"name,omitempty"`
	Network *string `json:"network,omitempty"`
	SameAddress *bool `json:"sameAddress,omitempty"`
	SpecialTips *string `json:"specialTips,omitempty"`
	UnLockConfirm *int32 `json:"unLockConfirm,omitempty"`
	WithdrawDesc *string `json:"withdrawDesc,omitempty"`
	WithdrawEnable *bool `json:"withdrawEnable,omitempty"`
	WithdrawFee *string `json:"withdrawFee,omitempty"`
	WithdrawIntegerMultiple *string `json:"withdrawIntegerMultiple,omitempty"`
	WithdrawInternalMin *string `json:"withdrawInternalMin,omitempty"`
	WithdrawMax *string `json:"withdrawMax,omitempty"`
	WithdrawMin *string `json:"withdrawMin,omitempty"`
}

// NewGetCapitalConfigGetallV1RespItemNetworkListInner instantiates a new GetCapitalConfigGetallV1RespItemNetworkListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCapitalConfigGetallV1RespItemNetworkListInner() *GetCapitalConfigGetallV1RespItemNetworkListInner {
	this := GetCapitalConfigGetallV1RespItemNetworkListInner{}
	return &this
}

// NewGetCapitalConfigGetallV1RespItemNetworkListInnerWithDefaults instantiates a new GetCapitalConfigGetallV1RespItemNetworkListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCapitalConfigGetallV1RespItemNetworkListInnerWithDefaults() *GetCapitalConfigGetallV1RespItemNetworkListInner {
	this := GetCapitalConfigGetallV1RespItemNetworkListInner{}
	return &this
}

// GetAddressRegex returns the AddressRegex field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetAddressRegex() string {
	if o == nil || IsNil(o.AddressRegex) {
		var ret string
		return ret
	}
	return *o.AddressRegex
}

// GetAddressRegexOk returns a tuple with the AddressRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetAddressRegexOk() (*string, bool) {
	if o == nil || IsNil(o.AddressRegex) {
		return nil, false
	}
	return o.AddressRegex, true
}

// HasAddressRegex returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasAddressRegex() bool {
	if o != nil && !IsNil(o.AddressRegex) {
		return true
	}

	return false
}

// SetAddressRegex gets a reference to the given string and assigns it to the AddressRegex field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetAddressRegex(v string) {
	o.AddressRegex = &v
}

// GetBusy returns the Busy field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetBusy() bool {
	if o == nil || IsNil(o.Busy) {
		var ret bool
		return ret
	}
	return *o.Busy
}

// GetBusyOk returns a tuple with the Busy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetBusyOk() (*bool, bool) {
	if o == nil || IsNil(o.Busy) {
		return nil, false
	}
	return o.Busy, true
}

// HasBusy returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasBusy() bool {
	if o != nil && !IsNil(o.Busy) {
		return true
	}

	return false
}

// SetBusy gets a reference to the given bool and assigns it to the Busy field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetBusy(v bool) {
	o.Busy = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetCoin() string {
	if o == nil || IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetCoinOk() (*string, bool) {
	if o == nil || IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasCoin() bool {
	if o != nil && !IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetCoin(v string) {
	o.Coin = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetContractAddressUrl returns the ContractAddressUrl field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetContractAddressUrl() string {
	if o == nil || IsNil(o.ContractAddressUrl) {
		var ret string
		return ret
	}
	return *o.ContractAddressUrl
}

// GetContractAddressUrlOk returns a tuple with the ContractAddressUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetContractAddressUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddressUrl) {
		return nil, false
	}
	return o.ContractAddressUrl, true
}

// HasContractAddressUrl returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasContractAddressUrl() bool {
	if o != nil && !IsNil(o.ContractAddressUrl) {
		return true
	}

	return false
}

// SetContractAddressUrl gets a reference to the given string and assigns it to the ContractAddressUrl field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetContractAddressUrl(v string) {
	o.ContractAddressUrl = &v
}

// GetDepositDesc returns the DepositDesc field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetDepositDesc() string {
	if o == nil || IsNil(o.DepositDesc) {
		var ret string
		return ret
	}
	return *o.DepositDesc
}

// GetDepositDescOk returns a tuple with the DepositDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetDepositDescOk() (*string, bool) {
	if o == nil || IsNil(o.DepositDesc) {
		return nil, false
	}
	return o.DepositDesc, true
}

// HasDepositDesc returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasDepositDesc() bool {
	if o != nil && !IsNil(o.DepositDesc) {
		return true
	}

	return false
}

// SetDepositDesc gets a reference to the given string and assigns it to the DepositDesc field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetDepositDesc(v string) {
	o.DepositDesc = &v
}

// GetDepositEnable returns the DepositEnable field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetDepositEnable() bool {
	if o == nil || IsNil(o.DepositEnable) {
		var ret bool
		return ret
	}
	return *o.DepositEnable
}

// GetDepositEnableOk returns a tuple with the DepositEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetDepositEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.DepositEnable) {
		return nil, false
	}
	return o.DepositEnable, true
}

// HasDepositEnable returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasDepositEnable() bool {
	if o != nil && !IsNil(o.DepositEnable) {
		return true
	}

	return false
}

// SetDepositEnable gets a reference to the given bool and assigns it to the DepositEnable field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetDepositEnable(v bool) {
	o.DepositEnable = &v
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetEstimatedArrivalTime() int64 {
	if o == nil || IsNil(o.EstimatedArrivalTime) {
		var ret int64
		return ret
	}
	return *o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetEstimatedArrivalTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EstimatedArrivalTime) {
		return nil, false
	}
	return o.EstimatedArrivalTime, true
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasEstimatedArrivalTime() bool {
	if o != nil && !IsNil(o.EstimatedArrivalTime) {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given int64 and assigns it to the EstimatedArrivalTime field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetEstimatedArrivalTime(v int64) {
	o.EstimatedArrivalTime = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetMemoRegex returns the MemoRegex field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetMemoRegex() string {
	if o == nil || IsNil(o.MemoRegex) {
		var ret string
		return ret
	}
	return *o.MemoRegex
}

// GetMemoRegexOk returns a tuple with the MemoRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetMemoRegexOk() (*string, bool) {
	if o == nil || IsNil(o.MemoRegex) {
		return nil, false
	}
	return o.MemoRegex, true
}

// HasMemoRegex returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasMemoRegex() bool {
	if o != nil && !IsNil(o.MemoRegex) {
		return true
	}

	return false
}

// SetMemoRegex gets a reference to the given string and assigns it to the MemoRegex field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetMemoRegex(v string) {
	o.MemoRegex = &v
}

// GetMinConfirm returns the MinConfirm field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetMinConfirm() int32 {
	if o == nil || IsNil(o.MinConfirm) {
		var ret int32
		return ret
	}
	return *o.MinConfirm
}

// GetMinConfirmOk returns a tuple with the MinConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetMinConfirmOk() (*int32, bool) {
	if o == nil || IsNil(o.MinConfirm) {
		return nil, false
	}
	return o.MinConfirm, true
}

// HasMinConfirm returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasMinConfirm() bool {
	if o != nil && !IsNil(o.MinConfirm) {
		return true
	}

	return false
}

// SetMinConfirm gets a reference to the given int32 and assigns it to the MinConfirm field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetMinConfirm(v int32) {
	o.MinConfirm = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetNetwork(v string) {
	o.Network = &v
}

// GetSameAddress returns the SameAddress field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetSameAddress() bool {
	if o == nil || IsNil(o.SameAddress) {
		var ret bool
		return ret
	}
	return *o.SameAddress
}

// GetSameAddressOk returns a tuple with the SameAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetSameAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.SameAddress) {
		return nil, false
	}
	return o.SameAddress, true
}

// HasSameAddress returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasSameAddress() bool {
	if o != nil && !IsNil(o.SameAddress) {
		return true
	}

	return false
}

// SetSameAddress gets a reference to the given bool and assigns it to the SameAddress field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetSameAddress(v bool) {
	o.SameAddress = &v
}

// GetSpecialTips returns the SpecialTips field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetSpecialTips() string {
	if o == nil || IsNil(o.SpecialTips) {
		var ret string
		return ret
	}
	return *o.SpecialTips
}

// GetSpecialTipsOk returns a tuple with the SpecialTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetSpecialTipsOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialTips) {
		return nil, false
	}
	return o.SpecialTips, true
}

// HasSpecialTips returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasSpecialTips() bool {
	if o != nil && !IsNil(o.SpecialTips) {
		return true
	}

	return false
}

// SetSpecialTips gets a reference to the given string and assigns it to the SpecialTips field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetSpecialTips(v string) {
	o.SpecialTips = &v
}

// GetUnLockConfirm returns the UnLockConfirm field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetUnLockConfirm() int32 {
	if o == nil || IsNil(o.UnLockConfirm) {
		var ret int32
		return ret
	}
	return *o.UnLockConfirm
}

// GetUnLockConfirmOk returns a tuple with the UnLockConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetUnLockConfirmOk() (*int32, bool) {
	if o == nil || IsNil(o.UnLockConfirm) {
		return nil, false
	}
	return o.UnLockConfirm, true
}

// HasUnLockConfirm returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasUnLockConfirm() bool {
	if o != nil && !IsNil(o.UnLockConfirm) {
		return true
	}

	return false
}

// SetUnLockConfirm gets a reference to the given int32 and assigns it to the UnLockConfirm field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetUnLockConfirm(v int32) {
	o.UnLockConfirm = &v
}

// GetWithdrawDesc returns the WithdrawDesc field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawDesc() string {
	if o == nil || IsNil(o.WithdrawDesc) {
		var ret string
		return ret
	}
	return *o.WithdrawDesc
}

// GetWithdrawDescOk returns a tuple with the WithdrawDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawDescOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawDesc) {
		return nil, false
	}
	return o.WithdrawDesc, true
}

// HasWithdrawDesc returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasWithdrawDesc() bool {
	if o != nil && !IsNil(o.WithdrawDesc) {
		return true
	}

	return false
}

// SetWithdrawDesc gets a reference to the given string and assigns it to the WithdrawDesc field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetWithdrawDesc(v string) {
	o.WithdrawDesc = &v
}

// GetWithdrawEnable returns the WithdrawEnable field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawEnable() bool {
	if o == nil || IsNil(o.WithdrawEnable) {
		var ret bool
		return ret
	}
	return *o.WithdrawEnable
}

// GetWithdrawEnableOk returns a tuple with the WithdrawEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.WithdrawEnable) {
		return nil, false
	}
	return o.WithdrawEnable, true
}

// HasWithdrawEnable returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasWithdrawEnable() bool {
	if o != nil && !IsNil(o.WithdrawEnable) {
		return true
	}

	return false
}

// SetWithdrawEnable gets a reference to the given bool and assigns it to the WithdrawEnable field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetWithdrawEnable(v bool) {
	o.WithdrawEnable = &v
}

// GetWithdrawFee returns the WithdrawFee field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawFee() string {
	if o == nil || IsNil(o.WithdrawFee) {
		var ret string
		return ret
	}
	return *o.WithdrawFee
}

// GetWithdrawFeeOk returns a tuple with the WithdrawFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawFeeOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawFee) {
		return nil, false
	}
	return o.WithdrawFee, true
}

// HasWithdrawFee returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasWithdrawFee() bool {
	if o != nil && !IsNil(o.WithdrawFee) {
		return true
	}

	return false
}

// SetWithdrawFee gets a reference to the given string and assigns it to the WithdrawFee field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetWithdrawFee(v string) {
	o.WithdrawFee = &v
}

// GetWithdrawIntegerMultiple returns the WithdrawIntegerMultiple field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawIntegerMultiple() string {
	if o == nil || IsNil(o.WithdrawIntegerMultiple) {
		var ret string
		return ret
	}
	return *o.WithdrawIntegerMultiple
}

// GetWithdrawIntegerMultipleOk returns a tuple with the WithdrawIntegerMultiple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawIntegerMultipleOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawIntegerMultiple) {
		return nil, false
	}
	return o.WithdrawIntegerMultiple, true
}

// HasWithdrawIntegerMultiple returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasWithdrawIntegerMultiple() bool {
	if o != nil && !IsNil(o.WithdrawIntegerMultiple) {
		return true
	}

	return false
}

// SetWithdrawIntegerMultiple gets a reference to the given string and assigns it to the WithdrawIntegerMultiple field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetWithdrawIntegerMultiple(v string) {
	o.WithdrawIntegerMultiple = &v
}

// GetWithdrawInternalMin returns the WithdrawInternalMin field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawInternalMin() string {
	if o == nil || IsNil(o.WithdrawInternalMin) {
		var ret string
		return ret
	}
	return *o.WithdrawInternalMin
}

// GetWithdrawInternalMinOk returns a tuple with the WithdrawInternalMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawInternalMinOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawInternalMin) {
		return nil, false
	}
	return o.WithdrawInternalMin, true
}

// HasWithdrawInternalMin returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasWithdrawInternalMin() bool {
	if o != nil && !IsNil(o.WithdrawInternalMin) {
		return true
	}

	return false
}

// SetWithdrawInternalMin gets a reference to the given string and assigns it to the WithdrawInternalMin field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetWithdrawInternalMin(v string) {
	o.WithdrawInternalMin = &v
}

// GetWithdrawMax returns the WithdrawMax field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawMax() string {
	if o == nil || IsNil(o.WithdrawMax) {
		var ret string
		return ret
	}
	return *o.WithdrawMax
}

// GetWithdrawMaxOk returns a tuple with the WithdrawMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawMaxOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawMax) {
		return nil, false
	}
	return o.WithdrawMax, true
}

// HasWithdrawMax returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasWithdrawMax() bool {
	if o != nil && !IsNil(o.WithdrawMax) {
		return true
	}

	return false
}

// SetWithdrawMax gets a reference to the given string and assigns it to the WithdrawMax field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetWithdrawMax(v string) {
	o.WithdrawMax = &v
}

// GetWithdrawMin returns the WithdrawMin field value if set, zero value otherwise.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawMin() string {
	if o == nil || IsNil(o.WithdrawMin) {
		var ret string
		return ret
	}
	return *o.WithdrawMin
}

// GetWithdrawMinOk returns a tuple with the WithdrawMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) GetWithdrawMinOk() (*string, bool) {
	if o == nil || IsNil(o.WithdrawMin) {
		return nil, false
	}
	return o.WithdrawMin, true
}

// HasWithdrawMin returns a boolean if a field has been set.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) HasWithdrawMin() bool {
	if o != nil && !IsNil(o.WithdrawMin) {
		return true
	}

	return false
}

// SetWithdrawMin gets a reference to the given string and assigns it to the WithdrawMin field.
func (o *GetCapitalConfigGetallV1RespItemNetworkListInner) SetWithdrawMin(v string) {
	o.WithdrawMin = &v
}

func (o GetCapitalConfigGetallV1RespItemNetworkListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCapitalConfigGetallV1RespItemNetworkListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressRegex) {
		toSerialize["addressRegex"] = o.AddressRegex
	}
	if !IsNil(o.Busy) {
		toSerialize["busy"] = o.Busy
	}
	if !IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if !IsNil(o.ContractAddressUrl) {
		toSerialize["contractAddressUrl"] = o.ContractAddressUrl
	}
	if !IsNil(o.DepositDesc) {
		toSerialize["depositDesc"] = o.DepositDesc
	}
	if !IsNil(o.DepositEnable) {
		toSerialize["depositEnable"] = o.DepositEnable
	}
	if !IsNil(o.EstimatedArrivalTime) {
		toSerialize["estimatedArrivalTime"] = o.EstimatedArrivalTime
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.MemoRegex) {
		toSerialize["memoRegex"] = o.MemoRegex
	}
	if !IsNil(o.MinConfirm) {
		toSerialize["minConfirm"] = o.MinConfirm
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.SameAddress) {
		toSerialize["sameAddress"] = o.SameAddress
	}
	if !IsNil(o.SpecialTips) {
		toSerialize["specialTips"] = o.SpecialTips
	}
	if !IsNil(o.UnLockConfirm) {
		toSerialize["unLockConfirm"] = o.UnLockConfirm
	}
	if !IsNil(o.WithdrawDesc) {
		toSerialize["withdrawDesc"] = o.WithdrawDesc
	}
	if !IsNil(o.WithdrawEnable) {
		toSerialize["withdrawEnable"] = o.WithdrawEnable
	}
	if !IsNil(o.WithdrawFee) {
		toSerialize["withdrawFee"] = o.WithdrawFee
	}
	if !IsNil(o.WithdrawIntegerMultiple) {
		toSerialize["withdrawIntegerMultiple"] = o.WithdrawIntegerMultiple
	}
	if !IsNil(o.WithdrawInternalMin) {
		toSerialize["withdrawInternalMin"] = o.WithdrawInternalMin
	}
	if !IsNil(o.WithdrawMax) {
		toSerialize["withdrawMax"] = o.WithdrawMax
	}
	if !IsNil(o.WithdrawMin) {
		toSerialize["withdrawMin"] = o.WithdrawMin
	}
	return toSerialize, nil
}

type NullableGetCapitalConfigGetallV1RespItemNetworkListInner struct {
	value *GetCapitalConfigGetallV1RespItemNetworkListInner
	isSet bool
}

func (v NullableGetCapitalConfigGetallV1RespItemNetworkListInner) Get() *GetCapitalConfigGetallV1RespItemNetworkListInner {
	return v.value
}

func (v *NullableGetCapitalConfigGetallV1RespItemNetworkListInner) Set(val *GetCapitalConfigGetallV1RespItemNetworkListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCapitalConfigGetallV1RespItemNetworkListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCapitalConfigGetallV1RespItemNetworkListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCapitalConfigGetallV1RespItemNetworkListInner(val *GetCapitalConfigGetallV1RespItemNetworkListInner) *NullableGetCapitalConfigGetallV1RespItemNetworkListInner {
	return &NullableGetCapitalConfigGetallV1RespItemNetworkListInner{value: val, isSet: true}
}

func (v NullableGetCapitalConfigGetallV1RespItemNetworkListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCapitalConfigGetallV1RespItemNetworkListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


