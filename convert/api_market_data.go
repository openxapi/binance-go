/*
Binance Convert API

OpenAPI specification for Binance exchange - Convert API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package convert

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MarketDataAPIService MarketDataAPI service
type MarketDataAPIService service

type MarketDataAPIConvertGetConvertAssetInfoV1Request struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	timestamp *int64
	recvWindow *int64
}

func (r MarketDataAPIConvertGetConvertAssetInfoV1Request) Timestamp(timestamp int64) MarketDataAPIConvertGetConvertAssetInfoV1Request {
	r.timestamp = &timestamp
	return r
}

// The value cannot be greater than 60000
func (r MarketDataAPIConvertGetConvertAssetInfoV1Request) RecvWindow(recvWindow int64) MarketDataAPIConvertGetConvertAssetInfoV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r MarketDataAPIConvertGetConvertAssetInfoV1Request) Execute() ([]ConvertGetConvertAssetInfoV1RespItem, *http.Response, error) {
	return r.ApiService.ConvertGetConvertAssetInfoV1Execute(r)
}

/*
ConvertGetConvertAssetInfoV1 Query order quantity precision per asset(USER_DATA)

Query for supported asset’s precision information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MarketDataAPIConvertGetConvertAssetInfoV1Request
*/
func (a *MarketDataAPIService) ConvertGetConvertAssetInfoV1(ctx context.Context) MarketDataAPIConvertGetConvertAssetInfoV1Request {
	return MarketDataAPIConvertGetConvertAssetInfoV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ConvertGetConvertAssetInfoV1RespItem
func (a *MarketDataAPIService) ConvertGetConvertAssetInfoV1Execute(r MarketDataAPIConvertGetConvertAssetInfoV1Request) ([]ConvertGetConvertAssetInfoV1RespItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ConvertGetConvertAssetInfoV1RespItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.ConvertGetConvertAssetInfoV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/convert/assetInfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	if r.recvWindow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "timestamp", r.timestamp, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextBinanceAuth).(Auth); ok {
			localVarHeaderParams["X-MBX-APIKEY"] = auth.APIKey
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MarketDataAPIConvertGetConvertExchangeInfoV1Request struct {
	ctx context.Context
	ApiService *MarketDataAPIService
	fromAsset *string
	toAsset *string
}

// User spends coin
func (r MarketDataAPIConvertGetConvertExchangeInfoV1Request) FromAsset(fromAsset string) MarketDataAPIConvertGetConvertExchangeInfoV1Request {
	r.fromAsset = &fromAsset
	return r
}

// User receives coin
func (r MarketDataAPIConvertGetConvertExchangeInfoV1Request) ToAsset(toAsset string) MarketDataAPIConvertGetConvertExchangeInfoV1Request {
	r.toAsset = &toAsset
	return r
}

func (r MarketDataAPIConvertGetConvertExchangeInfoV1Request) Execute() ([]ConvertGetConvertExchangeInfoV1RespItem, *http.Response, error) {
	return r.ApiService.ConvertGetConvertExchangeInfoV1Execute(r)
}

/*
ConvertGetConvertExchangeInfoV1 List All Convert Pairs

Query for all convertible token pairs and the tokens’ respective upper/lower limits

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MarketDataAPIConvertGetConvertExchangeInfoV1Request
*/
func (a *MarketDataAPIService) ConvertGetConvertExchangeInfoV1(ctx context.Context) MarketDataAPIConvertGetConvertExchangeInfoV1Request {
	return MarketDataAPIConvertGetConvertExchangeInfoV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ConvertGetConvertExchangeInfoV1RespItem
func (a *MarketDataAPIService) ConvertGetConvertExchangeInfoV1Execute(r MarketDataAPIConvertGetConvertExchangeInfoV1Request) ([]ConvertGetConvertExchangeInfoV1RespItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ConvertGetConvertExchangeInfoV1RespItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketDataAPIService.ConvertGetConvertExchangeInfoV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/convert/exchangeInfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromAsset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromAsset", r.fromAsset, "form", "")
	} else {
		var defaultValue string = ""
		r.fromAsset = &defaultValue
	}
	if r.toAsset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toAsset", r.toAsset, "form", "")
	} else {
		var defaultValue string = ""
		r.toAsset = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
