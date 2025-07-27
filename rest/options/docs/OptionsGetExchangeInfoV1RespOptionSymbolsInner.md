# OptionsGetExchangeInfoV1RespOptionSymbolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryDate** | Pointer to **int64** |  | [optional] 
**Filters** | Pointer to [**[]OptionsSymbolFilter**](OptionsSymbolFilter.md) |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**MaintenanceMargin** | Pointer to **string** |  | [optional] 
**MakerFeeRate** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**MinInitialMargin** | Pointer to **string** |  | [optional] 
**MinMaintenanceMargin** | Pointer to **string** |  | [optional] 
**MinQty** | Pointer to **string** |  | [optional] 
**PriceScale** | Pointer to **int32** |  | [optional] 
**QuantityScale** | Pointer to **int32** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TakerFeeRate** | Pointer to **string** |  | [optional] 
**Underlying** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **int32** |  | [optional] 

## Methods

### NewOptionsGetExchangeInfoV1RespOptionSymbolsInner

`func NewOptionsGetExchangeInfoV1RespOptionSymbolsInner() *OptionsGetExchangeInfoV1RespOptionSymbolsInner`

NewOptionsGetExchangeInfoV1RespOptionSymbolsInner instantiates a new OptionsGetExchangeInfoV1RespOptionSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetExchangeInfoV1RespOptionSymbolsInnerWithDefaults

`func NewOptionsGetExchangeInfoV1RespOptionSymbolsInnerWithDefaults() *OptionsGetExchangeInfoV1RespOptionSymbolsInner`

NewOptionsGetExchangeInfoV1RespOptionSymbolsInnerWithDefaults instantiates a new OptionsGetExchangeInfoV1RespOptionSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryDate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetFilters

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetFilters() []OptionsSymbolFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetFiltersOk() (*[]OptionsSymbolFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetFilters(v []OptionsSymbolFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetInitialMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetMaintenanceMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaintenanceMargin() string`

GetMaintenanceMargin returns the MaintenanceMargin field if non-nil, zero value otherwise.

### GetMaintenanceMarginOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaintenanceMarginOk() (*string, bool)`

GetMaintenanceMarginOk returns a tuple with the MaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMaintenanceMargin(v string)`

SetMaintenanceMargin sets MaintenanceMargin field to given value.

### HasMaintenanceMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMaintenanceMargin() bool`

HasMaintenanceMargin returns a boolean if a field has been set.

### GetMakerFeeRate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMakerFeeRate() string`

GetMakerFeeRate returns the MakerFeeRate field if non-nil, zero value otherwise.

### GetMakerFeeRateOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMakerFeeRateOk() (*string, bool)`

GetMakerFeeRateOk returns a tuple with the MakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerFeeRate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMakerFeeRate(v string)`

SetMakerFeeRate sets MakerFeeRate field to given value.

### HasMakerFeeRate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMakerFeeRate() bool`

HasMakerFeeRate returns a boolean if a field has been set.

### GetMaxQty

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetMinInitialMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinInitialMargin() string`

GetMinInitialMargin returns the MinInitialMargin field if non-nil, zero value otherwise.

### GetMinInitialMarginOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinInitialMarginOk() (*string, bool)`

GetMinInitialMarginOk returns a tuple with the MinInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInitialMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMinInitialMargin(v string)`

SetMinInitialMargin sets MinInitialMargin field to given value.

### HasMinInitialMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMinInitialMargin() bool`

HasMinInitialMargin returns a boolean if a field has been set.

### GetMinMaintenanceMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinMaintenanceMargin() string`

GetMinMaintenanceMargin returns the MinMaintenanceMargin field if non-nil, zero value otherwise.

### GetMinMaintenanceMarginOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinMaintenanceMarginOk() (*string, bool)`

GetMinMaintenanceMarginOk returns a tuple with the MinMaintenanceMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMaintenanceMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMinMaintenanceMargin(v string)`

SetMinMaintenanceMargin sets MinMaintenanceMargin field to given value.

### HasMinMaintenanceMargin

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMinMaintenanceMargin() bool`

HasMinMaintenanceMargin returns a boolean if a field has been set.

### GetMinQty

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetPriceScale

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetPriceScale() int32`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetPriceScaleOk() (*int32, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetPriceScale(v int32)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuantityScale() int32`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuantityScaleOk() (*int32, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetQuantityScale(v int32)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetSide

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStrikePrice

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTakerFeeRate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetTakerFeeRate() string`

GetTakerFeeRate returns the TakerFeeRate field if non-nil, zero value otherwise.

### GetTakerFeeRateOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetTakerFeeRateOk() (*string, bool)`

GetTakerFeeRateOk returns a tuple with the TakerFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerFeeRate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetTakerFeeRate(v string)`

SetTakerFeeRate sets TakerFeeRate field to given value.

### HasTakerFeeRate

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasTakerFeeRate() bool`

HasTakerFeeRate returns a boolean if a field has been set.

### GetUnderlying

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.

### GetUnit

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OptionsGetExchangeInfoV1RespOptionSymbolsInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


