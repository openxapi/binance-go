# CfuturesGetExchangeInfoV1RespSymbolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderType** | Pointer to **[]string** |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**BaseAssetPrecision** | Pointer to **int32** |  | [optional] 
**ContractSize** | Pointer to **int32** |  | [optional] 
**ContractStatus** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**DeliveryDate** | Pointer to **int32** |  | [optional] 
**EqualQtyPrecision** | Pointer to **int32** |  | [optional] 
**Filters** | Pointer to [**[]CfuturesSymbolFilter**](CfuturesSymbolFilter.md) |  | [optional] 
**LiquidationFee** | Pointer to **string** |  | [optional] 
**MaintMarginPercent** | Pointer to **string** |  | [optional] 
**MarginAsset** | Pointer to **string** |  | [optional] 
**MarketTakeBound** | Pointer to **string** |  | [optional] 
**OnboardDate** | Pointer to **int32** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**PricePrecision** | Pointer to **int32** |  | [optional] 
**QuantityPrecision** | Pointer to **int32** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**QuotePrecision** | Pointer to **int32** |  | [optional] 
**RequiredMarginPercent** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **[]string** |  | [optional] 
**TriggerProtect** | Pointer to **string** |  | [optional] 
**UnderlyingSubType** | Pointer to **[]string** |  | [optional] 
**UnderlyingType** | Pointer to **string** |  | [optional] 

## Methods

### NewCfuturesGetExchangeInfoV1RespSymbolsInner

`func NewCfuturesGetExchangeInfoV1RespSymbolsInner() *CfuturesGetExchangeInfoV1RespSymbolsInner`

NewCfuturesGetExchangeInfoV1RespSymbolsInner instantiates a new CfuturesGetExchangeInfoV1RespSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCfuturesGetExchangeInfoV1RespSymbolsInnerWithDefaults

`func NewCfuturesGetExchangeInfoV1RespSymbolsInnerWithDefaults() *CfuturesGetExchangeInfoV1RespSymbolsInner`

NewCfuturesGetExchangeInfoV1RespSymbolsInnerWithDefaults instantiates a new CfuturesGetExchangeInfoV1RespSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetOrderType() []string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetOrderTypeOk() (*[]string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetOrderType(v []string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetBaseAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAssetPrecision() int32`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAssetPrecisionOk() (*int32, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetBaseAssetPrecision(v int32)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetContractSize

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetContractSize() int32`

GetContractSize returns the ContractSize field if non-nil, zero value otherwise.

### GetContractSizeOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetContractSizeOk() (*int32, bool)`

GetContractSizeOk returns a tuple with the ContractSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractSize

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetContractSize(v int32)`

SetContractSize sets ContractSize field to given value.

### HasContractSize

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasContractSize() bool`

HasContractSize returns a boolean if a field has been set.

### GetContractStatus

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.

### HasContractStatus

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasContractStatus() bool`

HasContractStatus returns a boolean if a field has been set.

### GetContractType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetDeliveryDate() int32`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetDeliveryDateOk() (*int32, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetDeliveryDate(v int32)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetEqualQtyPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetEqualQtyPrecision() int32`

GetEqualQtyPrecision returns the EqualQtyPrecision field if non-nil, zero value otherwise.

### GetEqualQtyPrecisionOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetEqualQtyPrecisionOk() (*int32, bool)`

GetEqualQtyPrecisionOk returns a tuple with the EqualQtyPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualQtyPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetEqualQtyPrecision(v int32)`

SetEqualQtyPrecision sets EqualQtyPrecision field to given value.

### HasEqualQtyPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasEqualQtyPrecision() bool`

HasEqualQtyPrecision returns a boolean if a field has been set.

### GetFilters

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetFilters() []CfuturesSymbolFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetFiltersOk() (*[]CfuturesSymbolFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetFilters(v []CfuturesSymbolFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLiquidationFee

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetLiquidationFee() string`

GetLiquidationFee returns the LiquidationFee field if non-nil, zero value otherwise.

### GetLiquidationFeeOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetLiquidationFeeOk() (*string, bool)`

GetLiquidationFeeOk returns a tuple with the LiquidationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationFee

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetLiquidationFee(v string)`

SetLiquidationFee sets LiquidationFee field to given value.

### HasLiquidationFee

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasLiquidationFee() bool`

HasLiquidationFee returns a boolean if a field has been set.

### GetMaintMarginPercent

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetMaintMarginPercent() string`

GetMaintMarginPercent returns the MaintMarginPercent field if non-nil, zero value otherwise.

### GetMaintMarginPercentOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetMaintMarginPercentOk() (*string, bool)`

GetMaintMarginPercentOk returns a tuple with the MaintMarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMarginPercent

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetMaintMarginPercent(v string)`

SetMaintMarginPercent sets MaintMarginPercent field to given value.

### HasMaintMarginPercent

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasMaintMarginPercent() bool`

HasMaintMarginPercent returns a boolean if a field has been set.

### GetMarginAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetMarginAsset() string`

GetMarginAsset returns the MarginAsset field if non-nil, zero value otherwise.

### GetMarginAssetOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetMarginAssetOk() (*string, bool)`

GetMarginAssetOk returns a tuple with the MarginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetMarginAsset(v string)`

SetMarginAsset sets MarginAsset field to given value.

### HasMarginAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasMarginAsset() bool`

HasMarginAsset returns a boolean if a field has been set.

### GetMarketTakeBound

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetMarketTakeBound() string`

GetMarketTakeBound returns the MarketTakeBound field if non-nil, zero value otherwise.

### GetMarketTakeBoundOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetMarketTakeBoundOk() (*string, bool)`

GetMarketTakeBoundOk returns a tuple with the MarketTakeBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTakeBound

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetMarketTakeBound(v string)`

SetMarketTakeBound sets MarketTakeBound field to given value.

### HasMarketTakeBound

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasMarketTakeBound() bool`

HasMarketTakeBound returns a boolean if a field has been set.

### GetOnboardDate

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetOnboardDate() int32`

GetOnboardDate returns the OnboardDate field if non-nil, zero value otherwise.

### GetOnboardDateOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetOnboardDateOk() (*int32, bool)`

GetOnboardDateOk returns a tuple with the OnboardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardDate

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetOnboardDate(v int32)`

SetOnboardDate sets OnboardDate field to given value.

### HasOnboardDate

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasOnboardDate() bool`

HasOnboardDate returns a boolean if a field has been set.

### GetPair

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetPricePrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetPricePrecision() int32`

GetPricePrecision returns the PricePrecision field if non-nil, zero value otherwise.

### GetPricePrecisionOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetPricePrecisionOk() (*int32, bool)`

GetPricePrecisionOk returns a tuple with the PricePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetPricePrecision(v int32)`

SetPricePrecision sets PricePrecision field to given value.

### HasPricePrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasPricePrecision() bool`

HasPricePrecision returns a boolean if a field has been set.

### GetQuantityPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetQuantityPrecision() int32`

GetQuantityPrecision returns the QuantityPrecision field if non-nil, zero value otherwise.

### GetQuantityPrecisionOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetQuantityPrecisionOk() (*int32, bool)`

GetQuantityPrecisionOk returns a tuple with the QuantityPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetQuantityPrecision(v int32)`

SetQuantityPrecision sets QuantityPrecision field to given value.

### HasQuantityPrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasQuantityPrecision() bool`

HasQuantityPrecision returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetQuotePrecision() int32`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetQuotePrecisionOk() (*int32, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetQuotePrecision(v int32)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetRequiredMarginPercent

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetRequiredMarginPercent() string`

GetRequiredMarginPercent returns the RequiredMarginPercent field if non-nil, zero value otherwise.

### GetRequiredMarginPercentOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetRequiredMarginPercentOk() (*string, bool)`

GetRequiredMarginPercentOk returns a tuple with the RequiredMarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMarginPercent

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetRequiredMarginPercent(v string)`

SetRequiredMarginPercent sets RequiredMarginPercent field to given value.

### HasRequiredMarginPercent

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasRequiredMarginPercent() bool`

HasRequiredMarginPercent returns a boolean if a field has been set.

### GetSymbol

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetTimeInForce() []string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetTimeInForceOk() (*[]string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetTimeInForce(v []string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTriggerProtect

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetTriggerProtect() string`

GetTriggerProtect returns the TriggerProtect field if non-nil, zero value otherwise.

### GetTriggerProtectOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetTriggerProtectOk() (*string, bool)`

GetTriggerProtectOk returns a tuple with the TriggerProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerProtect

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetTriggerProtect(v string)`

SetTriggerProtect sets TriggerProtect field to given value.

### HasTriggerProtect

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasTriggerProtect() bool`

HasTriggerProtect returns a boolean if a field has been set.

### GetUnderlyingSubType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingSubType() []string`

GetUnderlyingSubType returns the UnderlyingSubType field if non-nil, zero value otherwise.

### GetUnderlyingSubTypeOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingSubTypeOk() (*[]string, bool)`

GetUnderlyingSubTypeOk returns a tuple with the UnderlyingSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSubType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetUnderlyingSubType(v []string)`

SetUnderlyingSubType sets UnderlyingSubType field to given value.

### HasUnderlyingSubType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasUnderlyingSubType() bool`

HasUnderlyingSubType returns a boolean if a field has been set.

### GetUnderlyingType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingType() string`

GetUnderlyingType returns the UnderlyingType field if non-nil, zero value otherwise.

### GetUnderlyingTypeOk

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingTypeOk() (*string, bool)`

GetUnderlyingTypeOk returns a tuple with the UnderlyingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) SetUnderlyingType(v string)`

SetUnderlyingType sets UnderlyingType field to given value.

### HasUnderlyingType

`func (o *CfuturesGetExchangeInfoV1RespSymbolsInner) HasUnderlyingType() bool`

HasUnderlyingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


