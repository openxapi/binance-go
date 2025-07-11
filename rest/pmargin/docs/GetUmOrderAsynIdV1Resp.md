# GetUmOrderAsynIdV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadId** | Pointer to **string** |  | [optional] 
**ExpirationTimestamp** | Pointer to **int64** |  | [optional] 
**IsExpired** | Pointer to **map[string]interface{}** |  | [optional] 
**Notified** | Pointer to **bool** |  | [optional] 
**S3Link** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUmOrderAsynIdV1Resp

`func NewGetUmOrderAsynIdV1Resp() *GetUmOrderAsynIdV1Resp`

NewGetUmOrderAsynIdV1Resp instantiates a new GetUmOrderAsynIdV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmOrderAsynIdV1RespWithDefaults

`func NewGetUmOrderAsynIdV1RespWithDefaults() *GetUmOrderAsynIdV1Resp`

NewGetUmOrderAsynIdV1RespWithDefaults instantiates a new GetUmOrderAsynIdV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadId

`func (o *GetUmOrderAsynIdV1Resp) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *GetUmOrderAsynIdV1Resp) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *GetUmOrderAsynIdV1Resp) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *GetUmOrderAsynIdV1Resp) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *GetUmOrderAsynIdV1Resp) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetUmOrderAsynIdV1Resp) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetUmOrderAsynIdV1Resp) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetUmOrderAsynIdV1Resp) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetIsExpired

`func (o *GetUmOrderAsynIdV1Resp) GetIsExpired() map[string]interface{}`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *GetUmOrderAsynIdV1Resp) GetIsExpiredOk() (*map[string]interface{}, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *GetUmOrderAsynIdV1Resp) SetIsExpired(v map[string]interface{})`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *GetUmOrderAsynIdV1Resp) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### SetIsExpiredNil

`func (o *GetUmOrderAsynIdV1Resp) SetIsExpiredNil(b bool)`

 SetIsExpiredNil sets the value for IsExpired to be an explicit nil

### UnsetIsExpired
`func (o *GetUmOrderAsynIdV1Resp) UnsetIsExpired()`

UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
### GetNotified

`func (o *GetUmOrderAsynIdV1Resp) GetNotified() bool`

GetNotified returns the Notified field if non-nil, zero value otherwise.

### GetNotifiedOk

`func (o *GetUmOrderAsynIdV1Resp) GetNotifiedOk() (*bool, bool)`

GetNotifiedOk returns a tuple with the Notified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotified

`func (o *GetUmOrderAsynIdV1Resp) SetNotified(v bool)`

SetNotified sets Notified field to given value.

### HasNotified

`func (o *GetUmOrderAsynIdV1Resp) HasNotified() bool`

HasNotified returns a boolean if a field has been set.

### GetS3Link

`func (o *GetUmOrderAsynIdV1Resp) GetS3Link() map[string]interface{}`

GetS3Link returns the S3Link field if non-nil, zero value otherwise.

### GetS3LinkOk

`func (o *GetUmOrderAsynIdV1Resp) GetS3LinkOk() (*map[string]interface{}, bool)`

GetS3LinkOk returns a tuple with the S3Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Link

`func (o *GetUmOrderAsynIdV1Resp) SetS3Link(v map[string]interface{})`

SetS3Link sets S3Link field to given value.

### HasS3Link

`func (o *GetUmOrderAsynIdV1Resp) HasS3Link() bool`

HasS3Link returns a boolean if a field has been set.

### SetS3LinkNil

`func (o *GetUmOrderAsynIdV1Resp) SetS3LinkNil(b bool)`

 SetS3LinkNil sets the value for S3Link to be an explicit nil

### UnsetS3Link
`func (o *GetUmOrderAsynIdV1Resp) UnsetS3Link()`

UnsetS3Link ensures that no value is present for S3Link, not even an explicit nil
### GetStatus

`func (o *GetUmOrderAsynIdV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUmOrderAsynIdV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUmOrderAsynIdV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUmOrderAsynIdV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *GetUmOrderAsynIdV1Resp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetUmOrderAsynIdV1Resp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetUmOrderAsynIdV1Resp) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetUmOrderAsynIdV1Resp) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


