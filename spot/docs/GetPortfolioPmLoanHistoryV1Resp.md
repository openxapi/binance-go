# GetPortfolioPmLoanHistoryV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetPortfolioPmLoanHistoryV1RespRowsInner**](GetPortfolioPmLoanHistoryV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetPortfolioPmLoanHistoryV1Resp

`func NewGetPortfolioPmLoanHistoryV1Resp() *GetPortfolioPmLoanHistoryV1Resp`

NewGetPortfolioPmLoanHistoryV1Resp instantiates a new GetPortfolioPmLoanHistoryV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPortfolioPmLoanHistoryV1RespWithDefaults

`func NewGetPortfolioPmLoanHistoryV1RespWithDefaults() *GetPortfolioPmLoanHistoryV1Resp`

NewGetPortfolioPmLoanHistoryV1RespWithDefaults instantiates a new GetPortfolioPmLoanHistoryV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetPortfolioPmLoanHistoryV1Resp) GetRows() []GetPortfolioPmLoanHistoryV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetPortfolioPmLoanHistoryV1Resp) GetRowsOk() (*[]GetPortfolioPmLoanHistoryV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetPortfolioPmLoanHistoryV1Resp) SetRows(v []GetPortfolioPmLoanHistoryV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetPortfolioPmLoanHistoryV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetPortfolioPmLoanHistoryV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPortfolioPmLoanHistoryV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPortfolioPmLoanHistoryV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetPortfolioPmLoanHistoryV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


