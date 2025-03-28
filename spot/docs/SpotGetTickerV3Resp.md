# SpotGetTickerV3Resp

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

### NewSpotGetTickerV3Resp

`func NewSpotGetTickerV3Resp() *SpotGetTickerV3Resp`

NewSpotGetTickerV3Resp instantiates a new SpotGetTickerV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetTickerV3RespWithDefaults

`func NewSpotGetTickerV3RespWithDefaults() *SpotGetTickerV3Resp`

NewSpotGetTickerV3RespWithDefaults instantiates a new SpotGetTickerV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseTime

`func (o *SpotGetTickerV3Resp) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *SpotGetTickerV3Resp) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *SpotGetTickerV3Resp) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *SpotGetTickerV3Resp) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetCount

`func (o *SpotGetTickerV3Resp) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SpotGetTickerV3Resp) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SpotGetTickerV3Resp) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SpotGetTickerV3Resp) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFirstId

`func (o *SpotGetTickerV3Resp) GetFirstId() int64`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *SpotGetTickerV3Resp) GetFirstIdOk() (*int64, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *SpotGetTickerV3Resp) SetFirstId(v int64)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *SpotGetTickerV3Resp) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### GetHighPrice

`func (o *SpotGetTickerV3Resp) GetHighPrice() string`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *SpotGetTickerV3Resp) GetHighPriceOk() (*string, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *SpotGetTickerV3Resp) SetHighPrice(v string)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *SpotGetTickerV3Resp) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastId

`func (o *SpotGetTickerV3Resp) GetLastId() int64`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *SpotGetTickerV3Resp) GetLastIdOk() (*int64, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *SpotGetTickerV3Resp) SetLastId(v int64)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *SpotGetTickerV3Resp) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetLastPrice

`func (o *SpotGetTickerV3Resp) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *SpotGetTickerV3Resp) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *SpotGetTickerV3Resp) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *SpotGetTickerV3Resp) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *SpotGetTickerV3Resp) GetLowPrice() string`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *SpotGetTickerV3Resp) GetLowPriceOk() (*string, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *SpotGetTickerV3Resp) SetLowPrice(v string)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *SpotGetTickerV3Resp) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetOpenPrice

`func (o *SpotGetTickerV3Resp) GetOpenPrice() string`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *SpotGetTickerV3Resp) GetOpenPriceOk() (*string, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *SpotGetTickerV3Resp) SetOpenPrice(v string)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *SpotGetTickerV3Resp) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetOpenTime

`func (o *SpotGetTickerV3Resp) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *SpotGetTickerV3Resp) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *SpotGetTickerV3Resp) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *SpotGetTickerV3Resp) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPriceChange

`func (o *SpotGetTickerV3Resp) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *SpotGetTickerV3Resp) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *SpotGetTickerV3Resp) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *SpotGetTickerV3Resp) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePercent

`func (o *SpotGetTickerV3Resp) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *SpotGetTickerV3Resp) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *SpotGetTickerV3Resp) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.

### HasPriceChangePercent

`func (o *SpotGetTickerV3Resp) HasPriceChangePercent() bool`

HasPriceChangePercent returns a boolean if a field has been set.

### GetQuoteVolume

`func (o *SpotGetTickerV3Resp) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *SpotGetTickerV3Resp) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *SpotGetTickerV3Resp) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.

### HasQuoteVolume

`func (o *SpotGetTickerV3Resp) HasQuoteVolume() bool`

HasQuoteVolume returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotGetTickerV3Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotGetTickerV3Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotGetTickerV3Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotGetTickerV3Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetVolume

`func (o *SpotGetTickerV3Resp) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SpotGetTickerV3Resp) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SpotGetTickerV3Resp) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SpotGetTickerV3Resp) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWeightedAvgPrice

`func (o *SpotGetTickerV3Resp) GetWeightedAvgPrice() string`

GetWeightedAvgPrice returns the WeightedAvgPrice field if non-nil, zero value otherwise.

### GetWeightedAvgPriceOk

`func (o *SpotGetTickerV3Resp) GetWeightedAvgPriceOk() (*string, bool)`

GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAvgPrice

`func (o *SpotGetTickerV3Resp) SetWeightedAvgPrice(v string)`

SetWeightedAvgPrice sets WeightedAvgPrice field to given value.

### HasWeightedAvgPrice

`func (o *SpotGetTickerV3Resp) HasWeightedAvgPrice() bool`

HasWeightedAvgPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


