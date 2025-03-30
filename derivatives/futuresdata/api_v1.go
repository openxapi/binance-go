/*
Binance Futuresdata API

OpenAPI specification for Binance cryptocurrency exchange - Futuresdata API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package futuresdata

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// V1APIService V1API service
type V1APIService service

type V1APIFuturesdataGetFuturesHistDataLinkV1Request struct {
	ctx context.Context
	ApiService *V1APIService
	symbol *string
	dataType *string
	startTime *int64
	endTime *int64
	timestamp *int64
	recvWindow *int64
}

// symbol name, e.g. BTCUSDT or BTCUSD_PERP ｜
func (r V1APIFuturesdataGetFuturesHistDataLinkV1Request) Symbol(symbol string) V1APIFuturesdataGetFuturesHistDataLinkV1Request {
	r.symbol = &symbol
	return r
}

// &#x60;T_DEPTH&#x60; for ticklevel orderbook data, &#x60;S_DEPTH&#x60; for orderbook snapshot data
func (r V1APIFuturesdataGetFuturesHistDataLinkV1Request) DataType(dataType string) V1APIFuturesdataGetFuturesHistDataLinkV1Request {
	r.dataType = &dataType
	return r
}

func (r V1APIFuturesdataGetFuturesHistDataLinkV1Request) StartTime(startTime int64) V1APIFuturesdataGetFuturesHistDataLinkV1Request {
	r.startTime = &startTime
	return r
}

func (r V1APIFuturesdataGetFuturesHistDataLinkV1Request) EndTime(endTime int64) V1APIFuturesdataGetFuturesHistDataLinkV1Request {
	r.endTime = &endTime
	return r
}

func (r V1APIFuturesdataGetFuturesHistDataLinkV1Request) Timestamp(timestamp int64) V1APIFuturesdataGetFuturesHistDataLinkV1Request {
	r.timestamp = &timestamp
	return r
}

func (r V1APIFuturesdataGetFuturesHistDataLinkV1Request) RecvWindow(recvWindow int64) V1APIFuturesdataGetFuturesHistDataLinkV1Request {
	r.recvWindow = &recvWindow
	return r
}

func (r V1APIFuturesdataGetFuturesHistDataLinkV1Request) Execute() (*FuturesdataGetFuturesHistDataLinkV1Resp, *http.Response, error) {
	return r.ApiService.FuturesdataGetFuturesHistDataLinkV1Execute(r)
}

/*
FuturesdataGetFuturesHistDataLinkV1 Get Future TickLevel Orderbook Historical Data Download Link(USER_DATA)

Get Future TickLevel Orderbook Historical Data Download Link

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return V1APIFuturesdataGetFuturesHistDataLinkV1Request
*/
func (a *V1APIService) FuturesdataGetFuturesHistDataLinkV1(ctx context.Context) V1APIFuturesdataGetFuturesHistDataLinkV1Request {
	return V1APIFuturesdataGetFuturesHistDataLinkV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FuturesdataGetFuturesHistDataLinkV1Resp
func (a *V1APIService) FuturesdataGetFuturesHistDataLinkV1Execute(r V1APIFuturesdataGetFuturesHistDataLinkV1Request) (*FuturesdataGetFuturesHistDataLinkV1Resp, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FuturesdataGetFuturesHistDataLinkV1Resp
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1APIService.FuturesdataGetFuturesHistDataLinkV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sapi/v1/futures/histDataLink"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.symbol == nil {
		return localVarReturnValue, nil, reportError("symbol is required and must be specified")
	}
	if r.dataType == nil {
		return localVarReturnValue, nil, reportError("dataType is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, reportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, reportError("endTime is required and must be specified")
	}
	if r.timestamp == nil {
		return localVarReturnValue, nil, reportError("timestamp is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "dataType", r.dataType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
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
