/*
Binance Options API

OpenAPI specification for Binance cryptocurrency exchange - Options API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package options

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// AccountAPIService AccountAPI service
type AccountAPIService service

type AccountAPIOptionsGetAccountV1Request struct {
	ctx context.Context
	ApiService *AccountAPIService
	timestamp *int64
	recvWindow *int64
}

func (r AccountAPIOptionsGetAccountV1Request) Timestamp(timestamp int64) AccountAPIOptionsGetAccountV1Request {
	r.timestamp = &timestamp
	return r
}

func (r AccountAPIOptionsGetAccountV1Request) RecvWindow(recvWindow int64) AccountAPIOptionsGetAccountV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r AccountAPIOptionsGetAccountV1Request) Execute() (*OptionsGetAccountV1Resp, *http.Response, error) {
	return r.ApiService.OptionsGetAccountV1Execute(r)
}

/*
OptionsGetAccountV1 Option Account Information(TRADE)

Get current account information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountAPIOptionsGetAccountV1Request
*/
func (a *AccountAPIService) OptionsGetAccountV1(ctx context.Context) AccountAPIOptionsGetAccountV1Request {
	return AccountAPIOptionsGetAccountV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OptionsGetAccountV1Resp
func (a *AccountAPIService) OptionsGetAccountV1Execute(r AccountAPIOptionsGetAccountV1Request) (*OptionsGetAccountV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OptionsGetAccountV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountAPIService.OptionsGetAccountV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eapi/v1/account"

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

type AccountAPIOptionsGetBillV1Request struct {
	ctx context.Context
	ApiService *AccountAPIService
	currency *string
	timestamp *int64
	recordId *int64
	startTime *int64
	endTime *int64
	limit *int32
	recvWindow *int64
}

// Asset type, only support USDT  as of now
func (r AccountAPIOptionsGetBillV1Request) Currency(currency string) AccountAPIOptionsGetBillV1Request {
	r.currency = &currency
	return r
}

func (r AccountAPIOptionsGetBillV1Request) Timestamp(timestamp int64) AccountAPIOptionsGetBillV1Request {
	r.timestamp = &timestamp
	return r
}

// Return the recordId and subsequent data, the latest data is returned by default, e.g 100000
func (r AccountAPIOptionsGetBillV1Request) RecordId(recordId int64) AccountAPIOptionsGetBillV1Request {
	r.recordId = &recordId
	return r
}

// Start Time, e.g 1593511200000
func (r AccountAPIOptionsGetBillV1Request) StartTime(startTime int64) AccountAPIOptionsGetBillV1Request {
	r.startTime = &startTime
	return r
}

// End Time, e.g 1593512200000
func (r AccountAPIOptionsGetBillV1Request) EndTime(endTime int64) AccountAPIOptionsGetBillV1Request {
	r.endTime = &endTime
	return r
}

// Number of result sets returned Default:100 Max:1000
func (r AccountAPIOptionsGetBillV1Request) Limit(limit int32) AccountAPIOptionsGetBillV1Request {
	r.limit = &limit
	return r
}

func (r AccountAPIOptionsGetBillV1Request) RecvWindow(recvWindow int64) AccountAPIOptionsGetBillV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r AccountAPIOptionsGetBillV1Request) Execute() ([]OptionsGetBillV1RespItem, *http.Response, error) {
	return r.ApiService.OptionsGetBillV1Execute(r)
}

/*
OptionsGetBillV1 Account Funding Flow (USER_DATA)

Query account funding flows.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountAPIOptionsGetBillV1Request
*/
func (a *AccountAPIService) OptionsGetBillV1(ctx context.Context) AccountAPIOptionsGetBillV1Request {
	return AccountAPIOptionsGetBillV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []OptionsGetBillV1RespItem
func (a *AccountAPIService) OptionsGetBillV1Execute(r AccountAPIOptionsGetBillV1Request) ([]OptionsGetBillV1RespItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OptionsGetBillV1RespItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountAPIService.OptionsGetBillV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eapi/v1/bill"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.currency == nil {
		return localVarReturnValue, nil, reportError("currency is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "form", "")
	if r.recordId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recordId", r.recordId, "form", "")
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
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

type AccountAPIOptionsGetIncomeAsynIdV1Request struct {
	ctx context.Context
	ApiService *AccountAPIService
	downloadId *string
	timestamp *int64
	recvWindow *int64
}

// get by download id api
func (r AccountAPIOptionsGetIncomeAsynIdV1Request) DownloadId(downloadId string) AccountAPIOptionsGetIncomeAsynIdV1Request {
	r.downloadId = &downloadId
	return r
}

func (r AccountAPIOptionsGetIncomeAsynIdV1Request) Timestamp(timestamp int64) AccountAPIOptionsGetIncomeAsynIdV1Request {
	r.timestamp = &timestamp
	return r
}

func (r AccountAPIOptionsGetIncomeAsynIdV1Request) RecvWindow(recvWindow int64) AccountAPIOptionsGetIncomeAsynIdV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r AccountAPIOptionsGetIncomeAsynIdV1Request) Execute() (*OptionsGetIncomeAsynIdV1Resp, *http.Response, error) {
	return r.ApiService.OptionsGetIncomeAsynIdV1Execute(r)
}

/*
OptionsGetIncomeAsynIdV1 Get Option Transaction History Download Link by Id (USER_DATA)

Get option transaction history download Link by Id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountAPIOptionsGetIncomeAsynIdV1Request
*/
func (a *AccountAPIService) OptionsGetIncomeAsynIdV1(ctx context.Context) AccountAPIOptionsGetIncomeAsynIdV1Request {
	return AccountAPIOptionsGetIncomeAsynIdV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OptionsGetIncomeAsynIdV1Resp
func (a *AccountAPIService) OptionsGetIncomeAsynIdV1Execute(r AccountAPIOptionsGetIncomeAsynIdV1Request) (*OptionsGetIncomeAsynIdV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OptionsGetIncomeAsynIdV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountAPIService.OptionsGetIncomeAsynIdV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eapi/v1/income/asyn/id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.downloadId == nil {
		return localVarReturnValue, nil, reportError("downloadId is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
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

type AccountAPIOptionsGetIncomeAsynV1Request struct {
	ctx context.Context
	ApiService *AccountAPIService
}

func (r AccountAPIOptionsGetIncomeAsynV1Request) Execute() (*OptionsGetIncomeAsynV1Resp, *http.Response, error) {
	return r.ApiService.OptionsGetIncomeAsynV1Execute(r)
}

/*
OptionsGetIncomeAsynV1 Get Download Id For Option Transaction History (USER_DATA)

Get download id for option transaction history

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountAPIOptionsGetIncomeAsynV1Request
*/
func (a *AccountAPIService) OptionsGetIncomeAsynV1(ctx context.Context) AccountAPIOptionsGetIncomeAsynV1Request {
	return AccountAPIOptionsGetIncomeAsynV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OptionsGetIncomeAsynV1Resp
func (a *AccountAPIService) OptionsGetIncomeAsynV1Execute(r AccountAPIOptionsGetIncomeAsynV1Request) (*OptionsGetIncomeAsynV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OptionsGetIncomeAsynV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountAPIService.OptionsGetIncomeAsynV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/eapi/v1/income/asyn"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
