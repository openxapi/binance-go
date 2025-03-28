# SpotGetTickerTradingDayV3RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseTime** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**FirstId** | Pointer to **int64** |  | [optional] 
**HighPrice** | Pointer to **string** |  | [optional] 
**LastId** | Pointer to **int64** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**LowPrice** | Pointer to **string** |  | [optional] 
**OpenPrice** | Pointer to **string** |  | [optional] 
**OpenTime** | Pointer to **int64** |  | [optional] 
**PriceChange** | Pointer to **string** |  | [optional] 
**PriceChangePercent** | Pointer to **string** |  | [optional] 
**QuoteVolume** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**WeightedAvgPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewSpotGetTickerTradingDayV3RespItem

`func NewSpotGetTickerTradingDayV3RespItem() *SpotGetTickerTradingDayV3RespItem`

NewSpotGetTickerTradingDayV3RespItem instantiates a new SpotGetTickerTradingDayV3RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetTickerTradingDayV3RespItemWithDefaults

`func NewSpotGetTickerTradingDayV3RespItemWithDefaults() *SpotGetTickerTradingDayV3RespItem`

NewSpotGetTickerTradingDayV3RespItemWithDefaults instantiates a new SpotGetTickerTradingDayV3RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseTime

`func (o *SpotGetTickerTradingDayV3RespItem) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *SpotGetTickerTradingDayV3RespItem) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *SpotGetTickerTradingDayV3RespItem) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetCount

`func (o *SpotGetTickerTradingDayV3RespItem) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SpotGetTickerTradingDayV3RespItem) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SpotGetTickerTradingDayV3RespItem) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFirstId

`func (o *SpotGetTickerTradingDayV3RespItem) GetFirstId() int64`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetFirstIdOk() (*int64, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *SpotGetTickerTradingDayV3RespItem) SetFirstId(v int64)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *SpotGetTickerTradingDayV3RespItem) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### GetHighPrice

`func (o *SpotGetTickerTradingDayV3RespItem) GetHighPrice() string`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetHighPriceOk() (*string, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *SpotGetTickerTradingDayV3RespItem) SetHighPrice(v string)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *SpotGetTickerTradingDayV3RespItem) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastId

`func (o *SpotGetTickerTradingDayV3RespItem) GetLastId() int64`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetLastIdOk() (*int64, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *SpotGetTickerTradingDayV3RespItem) SetLastId(v int64)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *SpotGetTickerTradingDayV3RespItem) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetLastPrice

`func (o *SpotGetTickerTradingDayV3RespItem) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *SpotGetTickerTradingDayV3RespItem) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *SpotGetTickerTradingDayV3RespItem) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *SpotGetTickerTradingDayV3RespItem) GetLowPrice() string`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetLowPriceOk() (*string, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *SpotGetTickerTradingDayV3RespItem) SetLowPrice(v string)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *SpotGetTickerTradingDayV3RespItem) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetOpenPrice

`func (o *SpotGetTickerTradingDayV3RespItem) GetOpenPrice() string`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetOpenPriceOk() (*string, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *SpotGetTickerTradingDayV3RespItem) SetOpenPrice(v string)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *SpotGetTickerTradingDayV3RespItem) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetOpenTime

`func (o *SpotGetTickerTradingDayV3RespItem) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *SpotGetTickerTradingDayV3RespItem) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *SpotGetTickerTradingDayV3RespItem) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPriceChange

`func (o *SpotGetTickerTradingDayV3RespItem) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *SpotGetTickerTradingDayV3RespItem) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *SpotGetTickerTradingDayV3RespItem) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePercent

`func (o *SpotGetTickerTradingDayV3RespItem) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *SpotGetTickerTradingDayV3RespItem) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.

### HasPriceChangePercent

`func (o *SpotGetTickerTradingDayV3RespItem) HasPriceChangePercent() bool`

HasPriceChangePercent returns a boolean if a field has been set.

### GetQuoteVolume

`func (o *SpotGetTickerTradingDayV3RespItem) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *SpotGetTickerTradingDayV3RespItem) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.

### HasQuoteVolume

`func (o *SpotGetTickerTradingDayV3RespItem) HasQuoteVolume() bool`

HasQuoteVolume returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotGetTickerTradingDayV3RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotGetTickerTradingDayV3RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotGetTickerTradingDayV3RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetVolume

`func (o *SpotGetTickerTradingDayV3RespItem) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SpotGetTickerTradingDayV3RespItem) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SpotGetTickerTradingDayV3RespItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWeightedAvgPrice

`func (o *SpotGetTickerTradingDayV3RespItem) GetWeightedAvgPrice() string`

GetWeightedAvgPrice returns the WeightedAvgPrice field if non-nil, zero value otherwise.

### GetWeightedAvgPriceOk

`func (o *SpotGetTickerTradingDayV3RespItem) GetWeightedAvgPriceOk() (*string, bool)`

GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAvgPrice

`func (o *SpotGetTickerTradingDayV3RespItem) SetWeightedAvgPrice(v string)`

SetWeightedAvgPrice sets WeightedAvgPrice field to given value.

### HasWeightedAvgPrice

`func (o *SpotGetTickerTradingDayV3RespItem) HasWeightedAvgPrice() bool`

HasWeightedAvgPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


