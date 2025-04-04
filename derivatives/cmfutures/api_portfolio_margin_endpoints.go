/*
Binance COIN-M Futures API

OpenAPI specification for Binance exchange - Cmfutures API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cmfutures

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// PortfolioMarginEndpointsAPIService PortfolioMarginEndpointsAPI service
type PortfolioMarginEndpointsAPIService service

type ApiCmfuturesGetPmAccountInfoV1Request struct {
	ctx context.Context
	ApiService *PortfolioMarginEndpointsAPIService
	asset *string
	recvWindow *int64
}

func (r ApiCmfuturesGetPmAccountInfoV1Request) Asset(asset string) ApiCmfuturesGetPmAccountInfoV1Request {
	r.asset = &asset
	return r
}

func (r ApiCmfuturesGetPmAccountInfoV1Request) RecvWindow(recvWindow int64) ApiCmfuturesGetPmAccountInfoV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiCmfuturesGetPmAccountInfoV1Request) Execute() (*CmfuturesGetPmAccountInfoV1Resp, *http.Response, error) {
	return r.ApiService.CmfuturesGetPmAccountInfoV1Execute(r)
}

/*
CmfuturesGetPmAccountInfoV1 Classic Portfolio Margin Account Information (USER_DATA)

Get Classic Portfolio Margin current account information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCmfuturesGetPmAccountInfoV1Request
*/
func (a *PortfolioMarginEndpointsAPIService) CmfuturesGetPmAccountInfoV1(ctx context.Context) ApiCmfuturesGetPmAccountInfoV1Request {
	return ApiCmfuturesGetPmAccountInfoV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CmfuturesGetPmAccountInfoV1Resp
func (a *PortfolioMarginEndpointsAPIService) CmfuturesGetPmAccountInfoV1Execute(r ApiCmfuturesGetPmAccountInfoV1Request) (*CmfuturesGetPmAccountInfoV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CmfuturesGetPmAccountInfoV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortfolioMarginEndpointsAPIService.CmfuturesGetPmAccountInfoV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dapi/v1/pmAccountInfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.asset == nil {
		return localVarReturnValue, nil, reportError("asset is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "asset", r.asset, "form", "")
	if r.recvWindow != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recvWindow", r.recvWindow, "form", "")
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
