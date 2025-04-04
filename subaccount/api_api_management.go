/*
Binance Sub Account API

OpenAPI specification for Binance exchange - Subaccount API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package subaccount

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ApiManagementAPIService ApiManagementAPI service
type ApiManagementAPIService service

type ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request struct {
	ctx context.Context
	ApiService *ApiManagementAPIService
	email *string
	status *string
	subAccountApiKey *string
	timestamp *int64
	ipAddress *string
	recvWindow *int64
}

func (r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) Email(email string) ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request {
	r.email = &email
	return r
}

func (r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) Status(status string) ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request {
	r.status = &status
	return r
}

func (r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) SubAccountApiKey(subAccountApiKey string) ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

func (r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) Timestamp(timestamp int64) ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request {
	r.timestamp = &timestamp
	return r
}

func (r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) IpAddress(ipAddress string) ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request {
	r.ipAddress = &ipAddress
	return r
}

func (r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) RecvWindow(recvWindow int64) ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) Execute() (*SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp, *http.Response, error) {
	return r.ApiService.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Execute(r)
}

/*
SubaccountCreateSubAccountSubAccountApiIpRestrictionV2 Add IP Restriction for Sub-Account API key(For Master Account)

Add IP Restriction for Sub-Account API key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request
*/
func (a *ApiManagementAPIService) SubaccountCreateSubAccountSubAccountApiIpRestrictionV2(ctx context.Context) ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request {
	return ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp
func (a *ApiManagementAPIService) SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Execute(r ApiSubaccountCreateSubAccountSubAccountApiIpRestrictionV2Request) (*SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubaccountCreateSubAccountSubAccountApiIpRestrictionV2Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiManagementAPIService.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v2/sub-account/subAccountApi/ipRestriction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}
	if r.status == nil {
		return localVarReturnValue, nil, reportError("status is required and must be specified")
	}
	if r.subAccountApiKey == nil {
		return localVarReturnValue, nil, reportError("subAccountApiKey is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "email", r.email, "", "")
	if r.ipAddress != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ipAddress", r.ipAddress, "", "")
	}
	if r.recvWindow != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "recvWindow", r.recvWindow, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "status", r.status, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "subAccountApiKey", r.subAccountApiKey, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "timestamp", r.timestamp, "", "")
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

type ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request struct {
	ctx context.Context
	ApiService *ApiManagementAPIService
	email *string
	subAccountApiKey *string
	timestamp *int64
	ipAddress *string
	recvWindow *int64
}

// &lt;a href&#x3D;\&quot;/docs/sub_account/api-management/Delete-IP-List-For-a-Sub-account-API-Key#email-address\&quot;&gt;Sub-account email&lt;/a&gt;
func (r ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request) Email(email string) ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request {
	r.email = &email
	return r
}

func (r ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request) SubAccountApiKey(subAccountApiKey string) ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

func (r ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request) Timestamp(timestamp int64) ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request {
	r.timestamp = &timestamp
	return r
}

// Can be added in batches, separated by commas
func (r ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request) IpAddress(ipAddress string) ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request {
	r.ipAddress = &ipAddress
	return r
}

func (r ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request) RecvWindow(recvWindow int64) ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request) Execute() (*SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp, *http.Response, error) {
	return r.ApiService.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Execute(r)
}

/*
SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1 Delete IP List For a Sub-account API Key(For Master Account)

Delete IP List For a Sub-account API Key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request
*/
func (a *ApiManagementAPIService) SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1(ctx context.Context) ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request {
	return ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp
func (a *ApiManagementAPIService) SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Execute(r ApiSubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Request) (*SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiManagementAPIService.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}
	if r.subAccountApiKey == nil {
		return localVarReturnValue, nil, reportError("subAccountApiKey is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "subAccountApiKey", r.subAccountApiKey, "form", "")
	if r.ipAddress != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ipAddress", r.ipAddress, "form", "")
	} else {
		var defaultValue string = ""
		r.ipAddress = &defaultValue
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

type ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request struct {
	ctx context.Context
	ApiService *ApiManagementAPIService
	email *string
	subAccountApiKey *string
	timestamp *int64
	recvWindow *int64
}

// &lt;a href&#x3D;\&quot;/docs/sub_account/api-management#email-address\&quot;&gt;Sub-account email&lt;/a&gt;
func (r ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request) Email(email string) ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request {
	r.email = &email
	return r
}

func (r ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request) SubAccountApiKey(subAccountApiKey string) ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request {
	r.subAccountApiKey = &subAccountApiKey
	return r
}

func (r ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request) Timestamp(timestamp int64) ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request {
	r.timestamp = &timestamp
	return r
}

func (r ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request) RecvWindow(recvWindow int64) ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request) Execute() (*SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp, *http.Response, error) {
	return r.ApiService.SubaccountGetSubAccountSubAccountApiIpRestrictionV1Execute(r)
}

/*
SubaccountGetSubAccountSubAccountApiIpRestrictionV1 Get IP Restriction for a Sub-account API Key(For Master Account)

Get IP Restriction for a Sub-account API Key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request
*/
func (a *ApiManagementAPIService) SubaccountGetSubAccountSubAccountApiIpRestrictionV1(ctx context.Context) ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request {
	return ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp
func (a *ApiManagementAPIService) SubaccountGetSubAccountSubAccountApiIpRestrictionV1Execute(r ApiSubaccountGetSubAccountSubAccountApiIpRestrictionV1Request) (*SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubaccountGetSubAccountSubAccountApiIpRestrictionV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiManagementAPIService.SubaccountGetSubAccountSubAccountApiIpRestrictionV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/sub-account/subAccountApi/ipRestriction"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}
	if r.subAccountApiKey == nil {
		return localVarReturnValue, nil, reportError("subAccountApiKey is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "subAccountApiKey", r.subAccountApiKey, "form", "")
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
