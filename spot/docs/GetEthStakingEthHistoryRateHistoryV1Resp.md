# GetEthStakingEthHistoryRateHistoryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetEthStakingEthHistoryRateHistoryV1RespRowsInner**](GetEthStakingEthHistoryRateHistoryV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEthStakingEthHistoryRateHistoryV1Resp

`func NewGetEthStakingEthHistoryRateHistoryV1Resp() *GetEthStakingEthHistoryRateHistoryV1Resp`

NewGetEthStakingEthHistoryRateHistoryV1Resp instantiates a new GetEthStakingEthHistoryRateHistoryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthStakingEthHistoryRateHistoryV1RespWithDefaults

`func NewGetEthStakingEthHistoryRateHistoryV1RespWithDefaults() *GetEthStakingEthHistoryRateHistoryV1Resp`

NewGetEthStakingEthHistoryRateHistoryV1RespWithDefaults instantiates a new GetEthStakingEthHistoryRateHistoryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) GetRows() []GetEthStakingEthHistoryRateHistoryV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) GetRowsOk() (*[]GetEthStakingEthHistoryRateHistoryV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) SetRows(v []GetEthStakingEthHistoryRateHistoryV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetEthStakingEthHistoryRateHistoryV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


