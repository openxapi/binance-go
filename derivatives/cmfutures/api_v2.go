/*
Binance Cfutures API

OpenAPI specification for Binance cryptocurrency exchange - Cfutures API

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


// V2APIService V2API service
type V2APIService service

type V2APICfuturesGetLeverageBracketV2Request struct {
	ctx context.Context
	ApiService *V2APIService
	timestamp *int64
	symbol *string
	recvWindow *int64
}

func (r V2APICfuturesGetLeverageBracketV2Request) Timestamp(timestamp int64) V2APICfuturesGetLeverageBracketV2Request {
	r.timestamp = &timestamp
	return r
}

func (r V2APICfuturesGetLeverageBracketV2Request) Symbol(symbol string) V2APICfuturesGetLeverageBracketV2Request {
	r.symbol = &symbol
	return r
}

func (r V2APICfuturesGetLeverageBracketV2Request) RecvWindow(recvWindow int64) V2APICfuturesGetLeverageBracketV2Request {
	r.recvWindow = &recvWindow
	return r
}

func (r V2APICfuturesGetLeverageBracketV2Request) Execute() ([]CfuturesGetLeverageBracketV2RespItem, *http.Response, error) {
	return r.ApiService.CfuturesGetLeverageBracketV2Execute(r)
}

/*
CfuturesGetLeverageBracketV2 Notional Bracket for Symbol(USER_DATA)

Get the symbol's notional bracket list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return V2APICfuturesGetLeverageBracketV2Request
*/
func (a *V2APIService) CfuturesGetLeverageBracketV2(ctx context.Context) V2APICfuturesGetLeverageBracketV2Request {
	return V2APICfuturesGetLeverageBracketV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []CfuturesGetLeverageBracketV2RespItem
func (a *V2APIService) CfuturesGetLeverageBracketV2Execute(r V2APICfuturesGetLeverageBracketV2Request) ([]CfuturesGetLeverageBracketV2RespItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CfuturesGetLeverageBracketV2RespItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V2APIService.CfuturesGetLeverageBracketV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dapi/v2/leverageBracket"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	if r.symbol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	} else {
		var defaultValue string = ""
		r.symbol = &defaultValue
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
