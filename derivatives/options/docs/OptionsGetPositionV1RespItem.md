# OptionsGetPositionV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryPrice** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **int32** |  | [optional] 
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

### NewOptionsGetPositionV1RespItem

`func NewOptionsGetPositionV1RespItem() *OptionsGetPositionV1RespItem`

NewOptionsGetPositionV1RespItem instantiates a new OptionsGetPositionV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetPositionV1RespItemWithDefaults

`func NewOptionsGetPositionV1RespItemWithDefaults() *OptionsGetPositionV1RespItem`

NewOptionsGetPositionV1RespItemWithDefaults instantiates a new OptionsGetPositionV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryPrice

`func (o *OptionsGetPositionV1RespItem) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *OptionsGetPositionV1RespItem) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *OptionsGetPositionV1RespItem) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *OptionsGetPositionV1RespItem) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetExpiryDate

`func (o *OptionsGetPositionV1RespItem) GetExpiryDate() int32`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *OptionsGetPositionV1RespItem) GetExpiryDateOk() (*int32, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *OptionsGetPositionV1RespItem) SetExpiryDate(v int32)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *OptionsGetPositionV1RespItem) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetMarkPrice

`func (o *OptionsGetPositionV1RespItem) GetMarkPrice() string`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *OptionsGetPositionV1RespItem) GetMarkPriceOk() (*string, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *OptionsGetPositionV1RespItem) SetMarkPrice(v string)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *OptionsGetPositionV1RespItem) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetMarkValue

`func (o *OptionsGetPositionV1RespItem) GetMarkValue() string`

GetMarkValue returns the MarkValue field if non-nil, zero value otherwise.

### GetMarkValueOk

`func (o *OptionsGetPositionV1RespItem) GetMarkValueOk() (*string, bool)`

GetMarkValueOk returns a tuple with the MarkValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkValue

`func (o *OptionsGetPositionV1RespItem) SetMarkValue(v string)`

SetMarkValue sets MarkValue field to given value.

### HasMarkValue

`func (o *OptionsGetPositionV1RespItem) HasMarkValue() bool`

HasMarkValue returns a boolean if a field has been set.

### GetOptionSide

`func (o *OptionsGetPositionV1RespItem) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *OptionsGetPositionV1RespItem) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *OptionsGetPositionV1RespItem) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *OptionsGetPositionV1RespItem) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetPositionCost

`func (o *OptionsGetPositionV1RespItem) GetPositionCost() string`

GetPositionCost returns the PositionCost field if non-nil, zero value otherwise.

### GetPositionCostOk

`func (o *OptionsGetPositionV1RespItem) GetPositionCostOk() (*string, bool)`

GetPositionCostOk returns a tuple with the PositionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionCost

`func (o *OptionsGetPositionV1RespItem) SetPositionCost(v string)`

SetPositionCost sets PositionCost field to given value.

### HasPositionCost

`func (o *OptionsGetPositionV1RespItem) HasPositionCost() bool`

HasPositionCost returns a boolean if a field has been set.

### GetPriceScale

`func (o *OptionsGetPositionV1RespItem) GetPriceScale() int32`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *OptionsGetPositionV1RespItem) GetPriceScaleOk() (*int32, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *OptionsGetPositionV1RespItem) SetPriceScale(v int32)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *OptionsGetPositionV1RespItem) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionsGetPositionV1RespItem) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionsGetPositionV1RespItem) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionsGetPositionV1RespItem) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OptionsGetPositionV1RespItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityScale

`func (o *OptionsGetPositionV1RespItem) GetQuantityScale() int32`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *OptionsGetPositionV1RespItem) GetQuantityScaleOk() (*int32, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *OptionsGetPositionV1RespItem) SetQuantityScale(v int32)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *OptionsGetPositionV1RespItem) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *OptionsGetPositionV1RespItem) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *OptionsGetPositionV1RespItem) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *OptionsGetPositionV1RespItem) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *OptionsGetPositionV1RespItem) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetReducibleQty

`func (o *OptionsGetPositionV1RespItem) GetReducibleQty() string`

GetReducibleQty returns the ReducibleQty field if non-nil, zero value otherwise.

### GetReducibleQtyOk

`func (o *OptionsGetPositionV1RespItem) GetReducibleQtyOk() (*string, bool)`

GetReducibleQtyOk returns a tuple with the ReducibleQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducibleQty

`func (o *OptionsGetPositionV1RespItem) SetReducibleQty(v string)`

SetReducibleQty sets ReducibleQty field to given value.

### HasReducibleQty

`func (o *OptionsGetPositionV1RespItem) HasReducibleQty() bool`

HasReducibleQty returns a boolean if a field has been set.

### GetRor

`func (o *OptionsGetPositionV1RespItem) GetRor() string`

GetRor returns the Ror field if non-nil, zero value otherwise.

### GetRorOk

`func (o *OptionsGetPositionV1RespItem) GetRorOk() (*string, bool)`

GetRorOk returns a tuple with the Ror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRor

`func (o *OptionsGetPositionV1RespItem) SetRor(v string)`

SetRor sets Ror field to given value.

### HasRor

`func (o *OptionsGetPositionV1RespItem) HasRor() bool`

HasRor returns a boolean if a field has been set.

### GetSide

`func (o *OptionsGetPositionV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OptionsGetPositionV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OptionsGetPositionV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OptionsGetPositionV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStrikePrice

`func (o *OptionsGetPositionV1RespItem) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *OptionsGetPositionV1RespItem) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *OptionsGetPositionV1RespItem) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *OptionsGetPositionV1RespItem) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionsGetPositionV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionsGetPositionV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionsGetPositionV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionsGetPositionV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnrealizedPNL

`func (o *OptionsGetPositionV1RespItem) GetUnrealizedPNL() string`

GetUnrealizedPNL returns the UnrealizedPNL field if non-nil, zero value otherwise.

### GetUnrealizedPNLOk

`func (o *OptionsGetPositionV1RespItem) GetUnrealizedPNLOk() (*string, bool)`

GetUnrealizedPNLOk returns a tuple with the UnrealizedPNL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPNL

`func (o *OptionsGetPositionV1RespItem) SetUnrealizedPNL(v string)`

SetUnrealizedPNL sets UnrealizedPNL field to given value.

### HasUnrealizedPNL

`func (o *OptionsGetPositionV1RespItem) HasUnrealizedPNL() bool`

HasUnrealizedPNL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


