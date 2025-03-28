# SpotGetDepthV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asks** | Pointer to **[][]string** |  | [optional] 
**Bids** | Pointer to **[][]string** |  | [optional] 
**LastUpdateId** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpotGetDepthV3Resp

`func NewSpotGetDepthV3Resp() *SpotGetDepthV3Resp`

NewSpotGetDepthV3Resp instantiates a new SpotGetDepthV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetDepthV3RespWithDefaults

`func NewSpotGetDepthV3RespWithDefaults() *SpotGetDepthV3Resp`

NewSpotGetDepthV3RespWithDefaults instantiates a new SpotGetDepthV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsks

`func (o *SpotGetDepthV3Resp) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *SpotGetDepthV3Resp) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *SpotGetDepthV3Resp) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *SpotGetDepthV3Resp) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *SpotGetDepthV3Resp) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *SpotGetDepthV3Resp) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *SpotGetDepthV3Resp) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *SpotGetDepthV3Resp) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetLastUpdateId

`func (o *SpotGetDepthV3Resp) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *SpotGetDepthV3Resp) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *SpotGetDepthV3Resp) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *SpotGetDepthV3Resp) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


