# GetBrokerSubAccountV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**MakerCommission** | Pointer to **float32** |  | [optional] 
**MarginMakerCommission** | Pointer to **int32** |  | [optional] 
**MarginTakerCommission** | Pointer to **int32** |  | [optional] 
**SubaccountId** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**TakerCommission** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetBrokerSubAccountV1RespItem

`func NewGetBrokerSubAccountV1RespItem() *GetBrokerSubAccountV1RespItem`

NewGetBrokerSubAccountV1RespItem instantiates a new GetBrokerSubAccountV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBrokerSubAccountV1RespItemWithDefaults

`func NewGetBrokerSubAccountV1RespItemWithDefaults() *GetBrokerSubAccountV1RespItem`

NewGetBrokerSubAccountV1RespItemWithDefaults instantiates a new GetBrokerSubAccountV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *GetBrokerSubAccountV1RespItem) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetBrokerSubAccountV1RespItem) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetBrokerSubAccountV1RespItem) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetBrokerSubAccountV1RespItem) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEmail

`func (o *GetBrokerSubAccountV1RespItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetBrokerSubAccountV1RespItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetBrokerSubAccountV1RespItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetBrokerSubAccountV1RespItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMakerCommission

`func (o *GetBrokerSubAccountV1RespItem) GetMakerCommission() float32`

GetMakerCommission returns the MakerCommission field if non-nil, zero value otherwise.

### GetMakerCommissionOk

`func (o *GetBrokerSubAccountV1RespItem) GetMakerCommissionOk() (*float32, bool)`

GetMakerCommissionOk returns a tuple with the MakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerCommission

`func (o *GetBrokerSubAccountV1RespItem) SetMakerCommission(v float32)`

SetMakerCommission sets MakerCommission field to given value.

### HasMakerCommission

`func (o *GetBrokerSubAccountV1RespItem) HasMakerCommission() bool`

HasMakerCommission returns a boolean if a field has been set.

### GetMarginMakerCommission

`func (o *GetBrokerSubAccountV1RespItem) GetMarginMakerCommission() int32`

GetMarginMakerCommission returns the MarginMakerCommission field if non-nil, zero value otherwise.

### GetMarginMakerCommissionOk

`func (o *GetBrokerSubAccountV1RespItem) GetMarginMakerCommissionOk() (*int32, bool)`

GetMarginMakerCommissionOk returns a tuple with the MarginMakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginMakerCommission

`func (o *GetBrokerSubAccountV1RespItem) SetMarginMakerCommission(v int32)`

SetMarginMakerCommission sets MarginMakerCommission field to given value.

### HasMarginMakerCommission

`func (o *GetBrokerSubAccountV1RespItem) HasMarginMakerCommission() bool`

HasMarginMakerCommission returns a boolean if a field has been set.

### GetMarginTakerCommission

`func (o *GetBrokerSubAccountV1RespItem) GetMarginTakerCommission() int32`

GetMarginTakerCommission returns the MarginTakerCommission field if non-nil, zero value otherwise.

### GetMarginTakerCommissionOk

`func (o *GetBrokerSubAccountV1RespItem) GetMarginTakerCommissionOk() (*int32, bool)`

GetMarginTakerCommissionOk returns a tuple with the MarginTakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginTakerCommission

`func (o *GetBrokerSubAccountV1RespItem) SetMarginTakerCommission(v int32)`

SetMarginTakerCommission sets MarginTakerCommission field to given value.

### HasMarginTakerCommission

`func (o *GetBrokerSubAccountV1RespItem) HasMarginTakerCommission() bool`

HasMarginTakerCommission returns a boolean if a field has been set.

### GetSubaccountId

`func (o *GetBrokerSubAccountV1RespItem) GetSubaccountId() string`

GetSubaccountId returns the SubaccountId field if non-nil, zero value otherwise.

### GetSubaccountIdOk

`func (o *GetBrokerSubAccountV1RespItem) GetSubaccountIdOk() (*string, bool)`

GetSubaccountIdOk returns a tuple with the SubaccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubaccountId

`func (o *GetBrokerSubAccountV1RespItem) SetSubaccountId(v string)`

SetSubaccountId sets SubaccountId field to given value.

### HasSubaccountId

`func (o *GetBrokerSubAccountV1RespItem) HasSubaccountId() bool`

HasSubaccountId returns a boolean if a field has been set.

### GetTag

`func (o *GetBrokerSubAccountV1RespItem) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetBrokerSubAccountV1RespItem) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetBrokerSubAccountV1RespItem) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetBrokerSubAccountV1RespItem) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTakerCommission

`func (o *GetBrokerSubAccountV1RespItem) GetTakerCommission() float32`

GetTakerCommission returns the TakerCommission field if non-nil, zero value otherwise.

### GetTakerCommissionOk

`func (o *GetBrokerSubAccountV1RespItem) GetTakerCommissionOk() (*float32, bool)`

GetTakerCommissionOk returns a tuple with the TakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommission

`func (o *GetBrokerSubAccountV1RespItem) SetTakerCommission(v float32)`

SetTakerCommission sets TakerCommission field to given value.

### HasTakerCommission

`func (o *GetBrokerSubAccountV1RespItem) HasTakerCommission() bool`

HasTakerCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


