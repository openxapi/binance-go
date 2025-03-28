# SpotGetAvgPriceV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseTime** | Pointer to **int64** |  | [optional] 
**Mins** | Pointer to **int32** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 

## Methods

### NewSpotGetAvgPriceV3Resp

`func NewSpotGetAvgPriceV3Resp() *SpotGetAvgPriceV3Resp`

NewSpotGetAvgPriceV3Resp instantiates a new SpotGetAvgPriceV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetAvgPriceV3RespWithDefaults

`func NewSpotGetAvgPriceV3RespWithDefaults() *SpotGetAvgPriceV3Resp`

NewSpotGetAvgPriceV3RespWithDefaults instantiates a new SpotGetAvgPriceV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseTime

`func (o *SpotGetAvgPriceV3Resp) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *SpotGetAvgPriceV3Resp) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *SpotGetAvgPriceV3Resp) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *SpotGetAvgPriceV3Resp) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetMins

`func (o *SpotGetAvgPriceV3Resp) GetMins() int32`

GetMins returns the Mins field if non-nil, zero value otherwise.

### GetMinsOk

`func (o *SpotGetAvgPriceV3Resp) GetMinsOk() (*int32, bool)`

GetMinsOk returns a tuple with the Mins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMins

`func (o *SpotGetAvgPriceV3Resp) SetMins(v int32)`

SetMins sets Mins field to given value.

### HasMins

`func (o *SpotGetAvgPriceV3Resp) HasMins() bool`

HasMins returns a boolean if a field has been set.

### GetPrice

`func (o *SpotGetAvgPriceV3Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpotGetAvgPriceV3Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SpotGetAvgPriceV3Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SpotGetAvgPriceV3Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


