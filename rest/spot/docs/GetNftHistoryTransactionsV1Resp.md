# GetNftHistoryTransactionsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetNftHistoryTransactionsV1RespListInner**](GetNftHistoryTransactionsV1RespListInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetNftHistoryTransactionsV1Resp

`func NewGetNftHistoryTransactionsV1Resp() *GetNftHistoryTransactionsV1Resp`

NewGetNftHistoryTransactionsV1Resp instantiates a new GetNftHistoryTransactionsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNftHistoryTransactionsV1RespWithDefaults

`func NewGetNftHistoryTransactionsV1RespWithDefaults() *GetNftHistoryTransactionsV1Resp`

NewGetNftHistoryTransactionsV1RespWithDefaults instantiates a new GetNftHistoryTransactionsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetNftHistoryTransactionsV1Resp) GetList() []GetNftHistoryTransactionsV1RespListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetNftHistoryTransactionsV1Resp) GetListOk() (*[]GetNftHistoryTransactionsV1RespListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetNftHistoryTransactionsV1Resp) SetList(v []GetNftHistoryTransactionsV1RespListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetNftHistoryTransactionsV1Resp) HasList() bool`

HasList returns a boolean if a field has been set.

### GetTotal

`func (o *GetNftHistoryTransactionsV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetNftHistoryTransactionsV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetNftHistoryTransactionsV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetNftHistoryTransactionsV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


