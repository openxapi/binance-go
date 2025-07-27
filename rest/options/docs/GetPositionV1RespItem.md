# GetPositionV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryPrice** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **int64** |  | [optional] 
**MarkPrice** | Pointer to **string** |  | [optional] 
**MarkValue** | Pointer to **string** |  | [optional] 
**OptionSide** | Pointer to **string** |  | [optional] 
**PositionCost** | Pointer to **string** |  | [optional] 
**PriceScale** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**QuantityScale** | Pointer to **int32** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**ReducibleQty** | Pointer to **string** |  | [optional] 
**Ror** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**UnrealizedPNL** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPositionV1RespItem

`func NewGetPositionV1RespItem() *GetPositionV1RespItem`

NewGetPositionV1RespItem instantiates a new GetPositionV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPositionV1RespItemWithDefaults

`func NewGetPositionV1RespItemWithDefaults() *GetPositionV1RespItem`

NewGetPositionV1RespItemWithDefaults instantiates a new GetPositionV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryPrice

`func (o *GetPositionV1RespItem) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *GetPositionV1RespItem) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *GetPositionV1RespItem) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *GetPositionV1RespItem) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetExpiryDate

`func (o *GetPositionV1RespItem) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GetPositionV1RespItem) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GetPositionV1RespItem) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GetPositionV1RespItem) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetMarkPrice

`func (o *GetPositionV1RespItem) GetMarkPrice() string`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *GetPositionV1RespItem) GetMarkPriceOk() (*string, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *GetPositionV1RespItem) SetMarkPrice(v string)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *GetPositionV1RespItem) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetMarkValue

`func (o *GetPositionV1RespItem) GetMarkValue() string`

GetMarkValue returns the MarkValue field if non-nil, zero value otherwise.

### GetMarkValueOk

`func (o *GetPositionV1RespItem) GetMarkValueOk() (*string, bool)`

GetMarkValueOk returns a tuple with the MarkValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkValue

`func (o *GetPositionV1RespItem) SetMarkValue(v string)`

SetMarkValue sets MarkValue field to given value.

### HasMarkValue

`func (o *GetPositionV1RespItem) HasMarkValue() bool`

HasMarkValue returns a boolean if a field has been set.

### GetOptionSide

`func (o *GetPositionV1RespItem) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *GetPositionV1RespItem) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *GetPositionV1RespItem) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *GetPositionV1RespItem) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetPositionCost

`func (o *GetPositionV1RespItem) GetPositionCost() string`

GetPositionCost returns the PositionCost field if non-nil, zero value otherwise.

### GetPositionCostOk

`func (o *GetPositionV1RespItem) GetPositionCostOk() (*string, bool)`

GetPositionCostOk returns a tuple with the PositionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCost

`func (o *GetPositionV1RespItem) SetPositionCost(v string)`

SetPositionCost sets PositionCost field to given value.

### HasPositionCost

`func (o *GetPositionV1RespItem) HasPositionCost() bool`

HasPositionCost returns a boolean if a field has been set.

### GetPriceScale

`func (o *GetPositionV1RespItem) GetPriceScale() int32`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *GetPositionV1RespItem) GetPriceScaleOk() (*int32, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *GetPositionV1RespItem) SetPriceScale(v int32)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *GetPositionV1RespItem) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantity

`func (o *GetPositionV1RespItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetPositionV1RespItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetPositionV1RespItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GetPositionV1RespItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityScale

`func (o *GetPositionV1RespItem) GetQuantityScale() int32`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *GetPositionV1RespItem) GetQuantityScaleOk() (*int32, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *GetPositionV1RespItem) SetQuantityScale(v int32)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *GetPositionV1RespItem) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *GetPositionV1RespItem) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *GetPositionV1RespItem) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *GetPositionV1RespItem) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *GetPositionV1RespItem) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetReducibleQty

`func (o *GetPositionV1RespItem) GetReducibleQty() string`

GetReducibleQty returns the ReducibleQty field if non-nil, zero value otherwise.

### GetReducibleQtyOk

`func (o *GetPositionV1RespItem) GetReducibleQtyOk() (*string, bool)`

GetReducibleQtyOk returns a tuple with the ReducibleQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducibleQty

`func (o *GetPositionV1RespItem) SetReducibleQty(v string)`

SetReducibleQty sets ReducibleQty field to given value.

### HasReducibleQty

`func (o *GetPositionV1RespItem) HasReducibleQty() bool`

HasReducibleQty returns a boolean if a field has been set.

### GetRor

`func (o *GetPositionV1RespItem) GetRor() string`

GetRor returns the Ror field if non-nil, zero value otherwise.

### GetRorOk

`func (o *GetPositionV1RespItem) GetRorOk() (*string, bool)`

GetRorOk returns a tuple with the Ror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRor

`func (o *GetPositionV1RespItem) SetRor(v string)`

SetRor sets Ror field to given value.

### HasRor

`func (o *GetPositionV1RespItem) HasRor() bool`

HasRor returns a boolean if a field has been set.

### GetSide

`func (o *GetPositionV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetPositionV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetPositionV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetPositionV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStrikePrice

`func (o *GetPositionV1RespItem) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *GetPositionV1RespItem) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *GetPositionV1RespItem) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *GetPositionV1RespItem) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetSymbol

`func (o *GetPositionV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetPositionV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetPositionV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetPositionV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnrealizedPNL

`func (o *GetPositionV1RespItem) GetUnrealizedPNL() string`

GetUnrealizedPNL returns the UnrealizedPNL field if non-nil, zero value otherwise.

### GetUnrealizedPNLOk

`func (o *GetPositionV1RespItem) GetUnrealizedPNLOk() (*string, bool)`

GetUnrealizedPNLOk returns a tuple with the UnrealizedPNL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPNL

`func (o *GetPositionV1RespItem) SetUnrealizedPNL(v string)`

SetUnrealizedPNL sets UnrealizedPNL field to given value.

### HasUnrealizedPNL

`func (o *GetPositionV1RespItem) HasUnrealizedPNL() bool`

HasUnrealizedPNL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


