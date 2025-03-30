# OptionsGetDepthV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int32** |  | [optional] 
**Asks** | Pointer to **[][]string** |  | [optional] 
**Bids** | Pointer to **[][]string** |  | [optional] 
**U** | Pointer to **int32** |  | [optional] 

## Methods

### NewOptionsGetDepthV1Resp

`func NewOptionsGetDepthV1Resp() *OptionsGetDepthV1Resp`

NewOptionsGetDepthV1Resp instantiates a new OptionsGetDepthV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetDepthV1RespWithDefaults

`func NewOptionsGetDepthV1RespWithDefaults() *OptionsGetDepthV1Resp`

NewOptionsGetDepthV1RespWithDefaults instantiates a new OptionsGetDepthV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *OptionsGetDepthV1Resp) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *OptionsGetDepthV1Resp) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *OptionsGetDepthV1Resp) SetT(v int32)`

SetT sets T field to given value.

### HasT

`func (o *OptionsGetDepthV1Resp) HasT() bool`

HasT returns a boolean if a field has been set.

### GetAsks

`func (o *OptionsGetDepthV1Resp) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *OptionsGetDepthV1Resp) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *OptionsGetDepthV1Resp) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *OptionsGetDepthV1Resp) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *OptionsGetDepthV1Resp) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *OptionsGetDepthV1Resp) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *OptionsGetDepthV1Resp) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *OptionsGetDepthV1Resp) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetU

`func (o *OptionsGetDepthV1Resp) GetU() int32`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *OptionsGetDepthV1Resp) GetUOk() (*int32, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *OptionsGetDepthV1Resp) SetU(v int32)`

SetU sets U field to given value.

### HasU

`func (o *OptionsGetDepthV1Resp) HasU() bool`

HasU returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


