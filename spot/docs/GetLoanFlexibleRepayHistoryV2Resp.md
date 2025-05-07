# GetLoanFlexibleRepayHistoryV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetLoanFlexibleRepayHistoryV2RespRowsInner**](GetLoanFlexibleRepayHistoryV2RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetLoanFlexibleRepayHistoryV2Resp

`func NewGetLoanFlexibleRepayHistoryV2Resp() *GetLoanFlexibleRepayHistoryV2Resp`

NewGetLoanFlexibleRepayHistoryV2Resp instantiates a new GetLoanFlexibleRepayHistoryV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoanFlexibleRepayHistoryV2RespWithDefaults

`func NewGetLoanFlexibleRepayHistoryV2RespWithDefaults() *GetLoanFlexibleRepayHistoryV2Resp`

NewGetLoanFlexibleRepayHistoryV2RespWithDefaults instantiates a new GetLoanFlexibleRepayHistoryV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetLoanFlexibleRepayHistoryV2Resp) GetRows() []GetLoanFlexibleRepayHistoryV2RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetLoanFlexibleRepayHistoryV2Resp) GetRowsOk() (*[]GetLoanFlexibleRepayHistoryV2RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetLoanFlexibleRepayHistoryV2Resp) SetRows(v []GetLoanFlexibleRepayHistoryV2RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetLoanFlexibleRepayHistoryV2Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetLoanFlexibleRepayHistoryV2Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLoanFlexibleRepayHistoryV2Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLoanFlexibleRepayHistoryV2Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLoanFlexibleRepayHistoryV2Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


