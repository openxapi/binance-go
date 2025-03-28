/*
Binance Ufutures API

OpenAPI specification for Binance cryptocurrency exchange - Ufutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ufutures

import (
	"encoding/json"
)

// checks if the UfuturesGetOrderAsynIdV1Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UfuturesGetOrderAsynIdV1Resp{}

// UfuturesGetOrderAsynIdV1Resp struct for UfuturesGetOrderAsynIdV1Resp
type UfuturesGetOrderAsynIdV1Resp struct {
	DownloadId *string `json:"downloadId,omitempty"`
	ExpirationTimestamp *int64 `json:"expirationTimestamp,omitempty"`
	IsExpired map[string]interface{} `json:"isExpired,omitempty"`
	Notified *bool `json:"notified,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewUfuturesGetOrderAsynIdV1Resp instantiates a new UfuturesGetOrderAsynIdV1Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUfuturesGetOrderAsynIdV1Resp() *UfuturesGetOrderAsynIdV1Resp {
	this := UfuturesGetOrderAsynIdV1Resp{}
	return &this
}

// NewUfuturesGetOrderAsynIdV1RespWithDefaults instantiates a new UfuturesGetOrderAsynIdV1Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUfuturesGetOrderAsynIdV1RespWithDefaults() *UfuturesGetOrderAsynIdV1Resp {
	this := UfuturesGetOrderAsynIdV1Resp{}
	return &this
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise.
func (o *UfuturesGetOrderAsynIdV1Resp) GetDownloadId() string {
	if o == nil || IsNil(o.DownloadId) {
		var ret string
		return ret
	}
	return *o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) GetDownloadIdOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadId) {
		return nil, false
	}
	return o.DownloadId, true
}

// HasDownloadId returns a boolean if a field has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) HasDownloadId() bool {
	if o != nil && !IsNil(o.DownloadId) {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given string and assigns it to the DownloadId field.
func (o *UfuturesGetOrderAsynIdV1Resp) SetDownloadId(v string) {
	o.DownloadId = &v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *UfuturesGetOrderAsynIdV1Resp) GetExpirationTimestamp() int64 {
	if o == nil || IsNil(o.ExpirationTimestamp) {
		var ret int64
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) GetExpirationTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpirationTimestamp) {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) HasExpirationTimestamp() bool {
	if o != nil && !IsNil(o.ExpirationTimestamp) {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int64 and assigns it to the ExpirationTimestamp field.
func (o *UfuturesGetOrderAsynIdV1Resp) SetExpirationTimestamp(v int64) {
	o.ExpirationTimestamp = &v
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UfuturesGetOrderAsynIdV1Resp) GetIsExpired() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IsExpired
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UfuturesGetOrderAsynIdV1Resp) GetIsExpiredOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IsExpired) {
		return map[string]interface{}{}, false
	}
	return o.IsExpired, true
}

// HasIsExpired returns a boolean if a field has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) HasIsExpired() bool {
	if o != nil && !IsNil(o.IsExpired) {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given map[string]interface{} and assigns it to the IsExpired field.
func (o *UfuturesGetOrderAsynIdV1Resp) SetIsExpired(v map[string]interface{}) {
	o.IsExpired = v
}

// GetNotified returns the Notified field value if set, zero value otherwise.
func (o *UfuturesGetOrderAsynIdV1Resp) GetNotified() bool {
	if o == nil || IsNil(o.Notified) {
		var ret bool
		return ret
	}
	return *o.Notified
}

// GetNotifiedOk returns a tuple with the Notified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) GetNotifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Notified) {
		return nil, false
	}
	return o.Notified, true
}

// HasNotified returns a boolean if a field has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) HasNotified() bool {
	if o != nil && !IsNil(o.Notified) {
		return true
	}

	return false
}

// SetNotified gets a reference to the given bool and assigns it to the Notified field.
func (o *UfuturesGetOrderAsynIdV1Resp) SetNotified(v bool) {
	o.Notified = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UfuturesGetOrderAsynIdV1Resp) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UfuturesGetOrderAsynIdV1Resp) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UfuturesGetOrderAsynIdV1Resp) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UfuturesGetOrderAsynIdV1Resp) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UfuturesGetOrderAsynIdV1Resp) SetUrl(v string) {
	o.Url = &v
}

func (o UfuturesGetOrderAsynIdV1Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UfuturesGetOrderAsynIdV1Resp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadId) {
		toSerialize["downloadId"] = o.DownloadId
	}
	if !IsNil(o.ExpirationTimestamp) {
		toSerialize["expirationTimestamp"] = o.ExpirationTimestamp
	}
	if o.IsExpired != nil {
		toSerialize["isExpired"] = o.IsExpired
	}
	if !IsNil(o.Notified) {
		toSerialize["notified"] = o.Notified
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableUfuturesGetOrderAsynIdV1Resp struct {
	value *UfuturesGetOrderAsynIdV1Resp
	isSet bool
}

func (v NullableUfuturesGetOrderAsynIdV1Resp) Get() *UfuturesGetOrderAsynIdV1Resp {
	return v.value
}

func (v *NullableUfuturesGetOrderAsynIdV1Resp) Set(val *UfuturesGetOrderAsynIdV1Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableUfuturesGetOrderAsynIdV1Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableUfuturesGetOrderAsynIdV1Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUfuturesGetOrderAsynIdV1Resp(val *UfuturesGetOrderAsynIdV1Resp) *NullableUfuturesGetOrderAsynIdV1Resp {
	return &NullableUfuturesGetOrderAsynIdV1Resp{value: val, isSet: true}
}

func (v NullableUfuturesGetOrderAsynIdV1Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUfuturesGetOrderAsynIdV1Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


