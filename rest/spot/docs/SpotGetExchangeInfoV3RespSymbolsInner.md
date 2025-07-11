# SpotGetExchangeInfoV3RespSymbolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowTrailingStop** | Pointer to **bool** |  | [optional] 
**AllowedSelfTradePreventionModes** | Pointer to **[]string** |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**BaseAssetPrecision** | Pointer to **int32** |  | [optional] 
**BaseCommissionPrecision** | Pointer to **int32** |  | [optional] 
**CancelReplaceAllowed** | Pointer to **bool** |  | [optional] 
**DefaultSelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]SpotSymbolFilter**](SpotSymbolFilter.md) |  | [optional] 
**IcebergAllowed** | Pointer to **bool** |  | [optional] 
**IsMarginTradingAllowed** | Pointer to **bool** |  | [optional] 
**IsSpotTradingAllowed** | Pointer to **bool** |  | [optional] 
**OcoAllowed** | Pointer to **bool** |  | [optional] 
**OrderTypes** | Pointer to **[]string** |  | [optional] 
**OtoAllowed** | Pointer to **bool** |  | [optional] 
**PermissionSets** | Pointer to **[][]string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**QuoteAssetPrecision** | Pointer to **int32** |  | [optional] 
**QuoteCommissionPrecision** | Pointer to **int32** |  | [optional] 
**QuoteOrderQtyMarketAllowed** | Pointer to **bool** |  | [optional] 
**QuotePrecision** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewSpotGetExchangeInfoV3RespSymbolsInner

`func NewSpotGetExchangeInfoV3RespSymbolsInner() *SpotGetExchangeInfoV3RespSymbolsInner`

NewSpotGetExchangeInfoV3RespSymbolsInner instantiates a new SpotGetExchangeInfoV3RespSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetExchangeInfoV3RespSymbolsInnerWithDefaults

`func NewSpotGetExchangeInfoV3RespSymbolsInnerWithDefaults() *SpotGetExchangeInfoV3RespSymbolsInner`

NewSpotGetExchangeInfoV3RespSymbolsInnerWithDefaults instantiates a new SpotGetExchangeInfoV3RespSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowTrailingStop

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetAllowTrailingStop() bool`

GetAllowTrailingStop returns the AllowTrailingStop field if non-nil, zero value otherwise.

### GetAllowTrailingStopOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetAllowTrailingStopOk() (*bool, bool)`

GetAllowTrailingStopOk returns a tuple with the AllowTrailingStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrailingStop

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetAllowTrailingStop(v bool)`

SetAllowTrailingStop sets AllowTrailingStop field to given value.

### HasAllowTrailingStop

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasAllowTrailingStop() bool`

HasAllowTrailingStop returns a boolean if a field has been set.

### GetAllowedSelfTradePreventionModes

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetAllowedSelfTradePreventionModes() []string`

GetAllowedSelfTradePreventionModes returns the AllowedSelfTradePreventionModes field if non-nil, zero value otherwise.

### GetAllowedSelfTradePreventionModesOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetAllowedSelfTradePreventionModesOk() (*[]string, bool)`

GetAllowedSelfTradePreventionModesOk returns a tuple with the AllowedSelfTradePreventionModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSelfTradePreventionModes

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetAllowedSelfTradePreventionModes(v []string)`

SetAllowedSelfTradePreventionModes sets AllowedSelfTradePreventionModes field to given value.

### HasAllowedSelfTradePreventionModes

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasAllowedSelfTradePreventionModes() bool`

HasAllowedSelfTradePreventionModes returns a boolean if a field has been set.

### GetBaseAsset

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetBaseAssetPrecision() int32`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetBaseAssetPrecisionOk() (*int32, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetBaseAssetPrecision(v int32)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetBaseCommissionPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetBaseCommissionPrecision() int32`

GetBaseCommissionPrecision returns the BaseCommissionPrecision field if non-nil, zero value otherwise.

### GetBaseCommissionPrecisionOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetBaseCommissionPrecisionOk() (*int32, bool)`

GetBaseCommissionPrecisionOk returns a tuple with the BaseCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommissionPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetBaseCommissionPrecision(v int32)`

SetBaseCommissionPrecision sets BaseCommissionPrecision field to given value.

### HasBaseCommissionPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasBaseCommissionPrecision() bool`

HasBaseCommissionPrecision returns a boolean if a field has been set.

### GetCancelReplaceAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetCancelReplaceAllowed() bool`

GetCancelReplaceAllowed returns the CancelReplaceAllowed field if non-nil, zero value otherwise.

### GetCancelReplaceAllowedOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetCancelReplaceAllowedOk() (*bool, bool)`

GetCancelReplaceAllowedOk returns a tuple with the CancelReplaceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReplaceAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetCancelReplaceAllowed(v bool)`

SetCancelReplaceAllowed sets CancelReplaceAllowed field to given value.

### HasCancelReplaceAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasCancelReplaceAllowed() bool`

HasCancelReplaceAllowed returns a boolean if a field has been set.

### GetDefaultSelfTradePreventionMode

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetDefaultSelfTradePreventionMode() string`

GetDefaultSelfTradePreventionMode returns the DefaultSelfTradePreventionMode field if non-nil, zero value otherwise.

### GetDefaultSelfTradePreventionModeOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetDefaultSelfTradePreventionModeOk() (*string, bool)`

GetDefaultSelfTradePreventionModeOk returns a tuple with the DefaultSelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSelfTradePreventionMode

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetDefaultSelfTradePreventionMode(v string)`

SetDefaultSelfTradePreventionMode sets DefaultSelfTradePreventionMode field to given value.

### HasDefaultSelfTradePreventionMode

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasDefaultSelfTradePreventionMode() bool`

HasDefaultSelfTradePreventionMode returns a boolean if a field has been set.

### GetFilters

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetFilters() []SpotSymbolFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetFiltersOk() (*[]SpotSymbolFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetFilters(v []SpotSymbolFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIcebergAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetIcebergAllowed() bool`

GetIcebergAllowed returns the IcebergAllowed field if non-nil, zero value otherwise.

### GetIcebergAllowedOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetIcebergAllowedOk() (*bool, bool)`

GetIcebergAllowedOk returns a tuple with the IcebergAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetIcebergAllowed(v bool)`

SetIcebergAllowed sets IcebergAllowed field to given value.

### HasIcebergAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasIcebergAllowed() bool`

HasIcebergAllowed returns a boolean if a field has been set.

### GetIsMarginTradingAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetIsMarginTradingAllowed() bool`

GetIsMarginTradingAllowed returns the IsMarginTradingAllowed field if non-nil, zero value otherwise.

### GetIsMarginTradingAllowedOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetIsMarginTradingAllowedOk() (*bool, bool)`

GetIsMarginTradingAllowedOk returns a tuple with the IsMarginTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarginTradingAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetIsMarginTradingAllowed(v bool)`

SetIsMarginTradingAllowed sets IsMarginTradingAllowed field to given value.

### HasIsMarginTradingAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasIsMarginTradingAllowed() bool`

HasIsMarginTradingAllowed returns a boolean if a field has been set.

### GetIsSpotTradingAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetIsSpotTradingAllowed() bool`

GetIsSpotTradingAllowed returns the IsSpotTradingAllowed field if non-nil, zero value otherwise.

### GetIsSpotTradingAllowedOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetIsSpotTradingAllowedOk() (*bool, bool)`

GetIsSpotTradingAllowedOk returns a tuple with the IsSpotTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpotTradingAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetIsSpotTradingAllowed(v bool)`

SetIsSpotTradingAllowed sets IsSpotTradingAllowed field to given value.

### HasIsSpotTradingAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasIsSpotTradingAllowed() bool`

HasIsSpotTradingAllowed returns a boolean if a field has been set.

### GetOcoAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetOcoAllowed() bool`

GetOcoAllowed returns the OcoAllowed field if non-nil, zero value otherwise.

### GetOcoAllowedOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetOcoAllowedOk() (*bool, bool)`

GetOcoAllowedOk returns a tuple with the OcoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcoAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetOcoAllowed(v bool)`

SetOcoAllowed sets OcoAllowed field to given value.

### HasOcoAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasOcoAllowed() bool`

HasOcoAllowed returns a boolean if a field has been set.

### GetOrderTypes

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetOtoAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetOtoAllowed() bool`

GetOtoAllowed returns the OtoAllowed field if non-nil, zero value otherwise.

### GetOtoAllowedOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetOtoAllowedOk() (*bool, bool)`

GetOtoAllowedOk returns a tuple with the OtoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtoAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetOtoAllowed(v bool)`

SetOtoAllowed sets OtoAllowed field to given value.

### HasOtoAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasOtoAllowed() bool`

HasOtoAllowed returns a boolean if a field has been set.

### GetPermissionSets

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetPermissionSets() [][]string`

GetPermissionSets returns the PermissionSets field if non-nil, zero value otherwise.

### GetPermissionSetsOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetPermissionSetsOk() (*[][]string, bool)`

GetPermissionSetsOk returns a tuple with the PermissionSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSets

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetPermissionSets(v [][]string)`

SetPermissionSets sets PermissionSets field to given value.

### HasPermissionSets

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasPermissionSets() bool`

HasPermissionSets returns a boolean if a field has been set.

### GetPermissions

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetQuoteAssetPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteAssetPrecision() int32`

GetQuoteAssetPrecision returns the QuoteAssetPrecision field if non-nil, zero value otherwise.

### GetQuoteAssetPrecisionOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteAssetPrecisionOk() (*int32, bool)`

GetQuoteAssetPrecisionOk returns a tuple with the QuoteAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAssetPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetQuoteAssetPrecision(v int32)`

SetQuoteAssetPrecision sets QuoteAssetPrecision field to given value.

### HasQuoteAssetPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasQuoteAssetPrecision() bool`

HasQuoteAssetPrecision returns a boolean if a field has been set.

### GetQuoteCommissionPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteCommissionPrecision() int32`

GetQuoteCommissionPrecision returns the QuoteCommissionPrecision field if non-nil, zero value otherwise.

### GetQuoteCommissionPrecisionOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteCommissionPrecisionOk() (*int32, bool)`

GetQuoteCommissionPrecisionOk returns a tuple with the QuoteCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCommissionPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetQuoteCommissionPrecision(v int32)`

SetQuoteCommissionPrecision sets QuoteCommissionPrecision field to given value.

### HasQuoteCommissionPrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasQuoteCommissionPrecision() bool`

HasQuoteCommissionPrecision returns a boolean if a field has been set.

### GetQuoteOrderQtyMarketAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteOrderQtyMarketAllowed() bool`

GetQuoteOrderQtyMarketAllowed returns the QuoteOrderQtyMarketAllowed field if non-nil, zero value otherwise.

### GetQuoteOrderQtyMarketAllowedOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuoteOrderQtyMarketAllowedOk() (*bool, bool)`

GetQuoteOrderQtyMarketAllowedOk returns a tuple with the QuoteOrderQtyMarketAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteOrderQtyMarketAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetQuoteOrderQtyMarketAllowed(v bool)`

SetQuoteOrderQtyMarketAllowed sets QuoteOrderQtyMarketAllowed field to given value.

### HasQuoteOrderQtyMarketAllowed

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasQuoteOrderQtyMarketAllowed() bool`

HasQuoteOrderQtyMarketAllowed returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuotePrecision() int32`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetQuotePrecisionOk() (*int32, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetQuotePrecision(v int32)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetStatus

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotGetExchangeInfoV3RespSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


