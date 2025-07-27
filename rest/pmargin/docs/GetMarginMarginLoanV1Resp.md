# GetMarginMarginLoanV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetMarginMarginLoanV1RespRowsInner**](GetMarginMarginLoanV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMarginMarginLoanV1Resp

`func NewGetMarginMarginLoanV1Resp() *GetMarginMarginLoanV1Resp`

NewGetMarginMarginLoanV1Resp instantiates a new GetMarginMarginLoanV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarginMarginLoanV1RespWithDefaults

`func NewGetMarginMarginLoanV1RespWithDefaults() *GetMarginMarginLoanV1Resp`

NewGetMarginMarginLoanV1RespWithDefaults instantiates a new GetMarginMarginLoanV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetMarginMarginLoanV1Resp) GetRows() []GetMarginMarginLoanV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetMarginMarginLoanV1Resp) GetRowsOk() (*[]GetMarginMarginLoanV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetMarginMarginLoanV1Resp) SetRows(v []GetMarginMarginLoanV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetMarginMarginLoanV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetMarginMarginLoanV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMarginMarginLoanV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMarginMarginLoanV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMarginMarginLoanV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


