# GetDepthV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **int32** |  | [optional] 
**T** | Pointer to **int32** |  | [optional] 
**Asks** | Pointer to **[][]string** |  | [optional] 
**Bids** | Pointer to **[][]string** |  | [optional] 
**LastUpdateId** | Pointer to **int64** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewGetDepthV1Resp

`func NewGetDepthV1Resp() *GetDepthV1Resp`

NewGetDepthV1Resp instantiates a new GetDepthV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDepthV1RespWithDefaults

`func NewGetDepthV1RespWithDefaults() *GetDepthV1Resp`

NewGetDepthV1RespWithDefaults instantiates a new GetDepthV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *GetDepthV1Resp) GetE() int32`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *GetDepthV1Resp) GetEOk() (*int32, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *GetDepthV1Resp) SetE(v int32)`

SetE sets E field to given value.

### HasE

`func (o *GetDepthV1Resp) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *GetDepthV1Resp) GetT() int32`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *GetDepthV1Resp) GetTOk() (*int32, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *GetDepthV1Resp) SetT(v int32)`

SetT sets T field to given value.

### HasT

`func (o *GetDepthV1Resp) HasT() bool`

HasT returns a boolean if a field has been set.

### GetAsks

`func (o *GetDepthV1Resp) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *GetDepthV1Resp) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *GetDepthV1Resp) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *GetDepthV1Resp) HasAsks() bool`

HasAsks returns a boolean if a field has been set.

### GetBids

`func (o *GetDepthV1Resp) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *GetDepthV1Resp) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *GetDepthV1Resp) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *GetDepthV1Resp) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetLastUpdateId

`func (o *GetDepthV1Resp) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *GetDepthV1Resp) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *GetDepthV1Resp) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *GetDepthV1Resp) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.

### GetPair

`func (o *GetDepthV1Resp) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *GetDepthV1Resp) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *GetDepthV1Resp) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *GetDepthV1Resp) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetSymbol

`func (o *GetDepthV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetDepthV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetDepthV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetDepthV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


