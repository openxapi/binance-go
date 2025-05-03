# GetMyAllocationsV3RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationId** | Pointer to **int64** |  | [optional] 
**AllocationType** | Pointer to **string** |  | [optional] 
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**IsAllocator** | Pointer to **bool** |  | [optional] 
**IsBuyer** | Pointer to **bool** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**QuoteQty** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetMyAllocationsV3RespItem

`func NewGetMyAllocationsV3RespItem() *GetMyAllocationsV3RespItem`

NewGetMyAllocationsV3RespItem instantiates a new GetMyAllocationsV3RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyAllocationsV3RespItemWithDefaults

`func NewGetMyAllocationsV3RespItemWithDefaults() *GetMyAllocationsV3RespItem`

NewGetMyAllocationsV3RespItemWithDefaults instantiates a new GetMyAllocationsV3RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationId

`func (o *GetMyAllocationsV3RespItem) GetAllocationId() int64`

GetAllocationId returns the AllocationId field if non-nil, zero value otherwise.

### GetAllocationIdOk

`func (o *GetMyAllocationsV3RespItem) GetAllocationIdOk() (*int64, bool)`

GetAllocationIdOk returns a tuple with the AllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationId

`func (o *GetMyAllocationsV3RespItem) SetAllocationId(v int64)`

SetAllocationId sets AllocationId field to given value.

### HasAllocationId

`func (o *GetMyAllocationsV3RespItem) HasAllocationId() bool`

HasAllocationId returns a boolean if a field has been set.

### GetAllocationType

`func (o *GetMyAllocationsV3RespItem) GetAllocationType() string`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *GetMyAllocationsV3RespItem) GetAllocationTypeOk() (*string, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *GetMyAllocationsV3RespItem) SetAllocationType(v string)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *GetMyAllocationsV3RespItem) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetCommission

`func (o *GetMyAllocationsV3RespItem) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *GetMyAllocationsV3RespItem) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *GetMyAllocationsV3RespItem) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *GetMyAllocationsV3RespItem) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *GetMyAllocationsV3RespItem) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *GetMyAllocationsV3RespItem) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *GetMyAllocationsV3RespItem) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *GetMyAllocationsV3RespItem) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetIsAllocator

`func (o *GetMyAllocationsV3RespItem) GetIsAllocator() bool`

GetIsAllocator returns the IsAllocator field if non-nil, zero value otherwise.

### GetIsAllocatorOk

`func (o *GetMyAllocationsV3RespItem) GetIsAllocatorOk() (*bool, bool)`

GetIsAllocatorOk returns a tuple with the IsAllocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllocator

`func (o *GetMyAllocationsV3RespItem) SetIsAllocator(v bool)`

SetIsAllocator sets IsAllocator field to given value.

### HasIsAllocator

`func (o *GetMyAllocationsV3RespItem) HasIsAllocator() bool`

HasIsAllocator returns a boolean if a field has been set.

### GetIsBuyer

`func (o *GetMyAllocationsV3RespItem) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *GetMyAllocationsV3RespItem) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *GetMyAllocationsV3RespItem) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.

### HasIsBuyer

`func (o *GetMyAllocationsV3RespItem) HasIsBuyer() bool`

HasIsBuyer returns a boolean if a field has been set.

### GetIsMaker

`func (o *GetMyAllocationsV3RespItem) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *GetMyAllocationsV3RespItem) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *GetMyAllocationsV3RespItem) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *GetMyAllocationsV3RespItem) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetOrderId

`func (o *GetMyAllocationsV3RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetMyAllocationsV3RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetMyAllocationsV3RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetMyAllocationsV3RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *GetMyAllocationsV3RespItem) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *GetMyAllocationsV3RespItem) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *GetMyAllocationsV3RespItem) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *GetMyAllocationsV3RespItem) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetPrice

`func (o *GetMyAllocationsV3RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetMyAllocationsV3RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetMyAllocationsV3RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetMyAllocationsV3RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *GetMyAllocationsV3RespItem) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *GetMyAllocationsV3RespItem) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *GetMyAllocationsV3RespItem) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *GetMyAllocationsV3RespItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetQuoteQty

`func (o *GetMyAllocationsV3RespItem) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *GetMyAllocationsV3RespItem) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *GetMyAllocationsV3RespItem) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *GetMyAllocationsV3RespItem) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetSymbol

`func (o *GetMyAllocationsV3RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetMyAllocationsV3RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetMyAllocationsV3RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetMyAllocationsV3RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *GetMyAllocationsV3RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetMyAllocationsV3RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetMyAllocationsV3RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetMyAllocationsV3RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


