/*
Binance Spot API

OpenAPI specification for Binance exchange - Spot API

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spot

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// RebateAPIService RebateAPI service
type RebateAPIService service

type ApiGetRebateTaxQueryV1Request struct {
	ctx context.Context
	ApiService *RebateAPIService
	timestamp *int64
	startTime *int64
	endTime *int64
	page *int32
	recvWindow *int64
}

func (r ApiGetRebateTaxQueryV1Request) Timestamp(timestamp int64) ApiGetRebateTaxQueryV1Request {
	r.timestamp = &timestamp
	return r
}

func (r ApiGetRebateTaxQueryV1Request) StartTime(startTime int64) ApiGetRebateTaxQueryV1Request {
	r.startTime = &startTime
	return r
}

func (r ApiGetRebateTaxQueryV1Request) EndTime(endTime int64) ApiGetRebateTaxQueryV1Request {
	r.endTime = &endTime
	return r
}

// Default 1
func (r ApiGetRebateTaxQueryV1Request) Page(page int32) ApiGetRebateTaxQueryV1Request {
	r.page = &page
	return r
}

func (r ApiGetRebateTaxQueryV1Request) RecvWindow(recvWindow int64) ApiGetRebateTaxQueryV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiGetRebateTaxQueryV1Request) Execute() (*GetRebateTaxQueryV1Resp, *http.Response, error) {
	return r.ApiService.GetRebateTaxQueryV1Execute(r)
}

/*
GetRebateTaxQueryV1 Get Spot Rebate History Records (USER_DATA)

Get Spot Rebate History Records

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRebateTaxQueryV1Request
*/
func (a *RebateAPIService) GetRebateTaxQueryV1(ctx context.Context) ApiGetRebateTaxQueryV1Request {
	return ApiGetRebateTaxQueryV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetRebateTaxQueryV1Resp
func (a *RebateAPIService) GetRebateTaxQueryV1Execute(r ApiGetRebateTaxQueryV1Request) (*GetRebateTaxQueryV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetRebateTaxQueryV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RebateAPIService.GetRebateTaxQueryV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/rebate/taxQuery"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
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
