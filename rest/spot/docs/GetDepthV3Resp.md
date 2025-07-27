# GetDepthV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asks** | Pointer to **[][]string** |  | [optional] 
**Bids** | Pointer to **[][]string** |  | [optional] 
**LastUpdateId** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetDepthV3Resp

`func NewGetDepthV3Resp() *GetDepthV3Resp`

NewGetDepthV3Resp instantiates a new GetDepthV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDepthV3RespWithDefaults

`func NewGetDepthV3RespWithDefaults() *GetDepthV3Resp`

NewGetDepthV3RespWithDefaults instantiates a new GetDepthV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsks

`func (o *GetDepthV3Resp) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *GetDepthV3Resp) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *GetDepthV3Resp) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *GetDepthV3Resp) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *GetDepthV3Resp) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *GetDepthV3Resp) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *GetDepthV3Resp) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *GetDepthV3Resp) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetLastUpdateId

`func (o *GetDepthV3Resp) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *GetDepthV3Resp) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *GetDepthV3Resp) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *GetDepthV3Resp) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


