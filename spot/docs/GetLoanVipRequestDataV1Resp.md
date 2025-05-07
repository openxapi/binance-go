# GetLoanVipRequestDataV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetLoanVipRequestDataV1RespRowsInner**](GetLoanVipRequestDataV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetLoanVipRequestDataV1Resp

`func NewGetLoanVipRequestDataV1Resp() *GetLoanVipRequestDataV1Resp`

NewGetLoanVipRequestDataV1Resp instantiates a new GetLoanVipRequestDataV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoanVipRequestDataV1RespWithDefaults

`func NewGetLoanVipRequestDataV1RespWithDefaults() *GetLoanVipRequestDataV1Resp`

NewGetLoanVipRequestDataV1RespWithDefaults instantiates a new GetLoanVipRequestDataV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetLoanVipRequestDataV1Resp) GetRows() []GetLoanVipRequestDataV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetLoanVipRequestDataV1Resp) GetRowsOk() (*[]GetLoanVipRequestDataV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetLoanVipRequestDataV1Resp) SetRows(v []GetLoanVipRequestDataV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetLoanVipRequestDataV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetLoanVipRequestDataV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLoanVipRequestDataV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLoanVipRequestDataV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLoanVipRequestDataV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


