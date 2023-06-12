/*
ConductorOne API

The ConductorOne API is a HTTP API for managing ConductorOne resources.

API version: 0.1.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package c1api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PoliciesAPIService PoliciesAPI service
type PoliciesAPIService service

type PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	c1ApiPolicyV1CreatePolicyRequest *C1ApiPolicyV1CreatePolicyRequest
}

func (r PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest) C1ApiPolicyV1CreatePolicyRequest(c1ApiPolicyV1CreatePolicyRequest C1ApiPolicyV1CreatePolicyRequest) PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest {
	r.c1ApiPolicyV1CreatePolicyRequest = &c1ApiPolicyV1CreatePolicyRequest
	return r
}

func (r PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest) Execute() (*C1ApiPolicyV1Policy, *http.Response, error) {
	return r.ApiService.C1ApiPolicyV1PoliciesCreateExecute(r)
}

/*
C1ApiPolicyV1PoliciesCreate Method for C1ApiPolicyV1PoliciesCreate

Invokes the c1.api.policy.v1.Policies.Create method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest
*/
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesCreate(ctx context.Context) PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest {
	return PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return C1ApiPolicyV1Policy
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesCreateExecute(r PoliciesAPIC1ApiPolicyV1PoliciesCreateRequest) (*C1ApiPolicyV1Policy, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *C1ApiPolicyV1Policy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.C1ApiPolicyV1PoliciesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.c1ApiPolicyV1CreatePolicyRequest
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

type PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	id string
	body *map[string]interface{}
}

func (r PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest) Body(body map[string]interface{}) PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest {
	r.body = &body
	return r
}

func (r PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.C1ApiPolicyV1PoliciesDeleteExecute(r)
}

/*
C1ApiPolicyV1PoliciesDelete Method for C1ApiPolicyV1PoliciesDelete

Invokes the c1.api.policy.v1.Policies.Delete method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest
*/
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesDelete(ctx context.Context, id string) PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest {
	return PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesDeleteExecute(r PoliciesAPIC1ApiPolicyV1PoliciesDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.C1ApiPolicyV1PoliciesDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/policies/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.body
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

type PoliciesAPIC1ApiPolicyV1PoliciesGetRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
	id string
}

func (r PoliciesAPIC1ApiPolicyV1PoliciesGetRequest) Execute() (*C1ApiPolicyV1Policy, *http.Response, error) {
	return r.ApiService.C1ApiPolicyV1PoliciesGetExecute(r)
}

/*
C1ApiPolicyV1PoliciesGet Method for C1ApiPolicyV1PoliciesGet

Invokes the c1.api.policy.v1.Policies.Get method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return PoliciesAPIC1ApiPolicyV1PoliciesGetRequest
*/
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesGet(ctx context.Context, id string) PoliciesAPIC1ApiPolicyV1PoliciesGetRequest {
	return PoliciesAPIC1ApiPolicyV1PoliciesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return C1ApiPolicyV1Policy
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesGetExecute(r PoliciesAPIC1ApiPolicyV1PoliciesGetRequest) (*C1ApiPolicyV1Policy, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *C1ApiPolicyV1Policy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.C1ApiPolicyV1PoliciesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/policies/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type PoliciesAPIC1ApiPolicyV1PoliciesListRequest struct {
	ctx context.Context
	ApiService *PoliciesAPIService
}

func (r PoliciesAPIC1ApiPolicyV1PoliciesListRequest) Execute() (*C1ApiPolicyV1ListPolicyResponse, *http.Response, error) {
	return r.ApiService.C1ApiPolicyV1PoliciesListExecute(r)
}

/*
C1ApiPolicyV1PoliciesList Method for C1ApiPolicyV1PoliciesList

Invokes the c1.api.policy.v1.Policies.List method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PoliciesAPIC1ApiPolicyV1PoliciesListRequest
*/
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesList(ctx context.Context) PoliciesAPIC1ApiPolicyV1PoliciesListRequest {
	return PoliciesAPIC1ApiPolicyV1PoliciesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return C1ApiPolicyV1ListPolicyResponse
func (a *PoliciesAPIService) C1ApiPolicyV1PoliciesListExecute(r PoliciesAPIC1ApiPolicyV1PoliciesListRequest) (*C1ApiPolicyV1ListPolicyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *C1ApiPolicyV1ListPolicyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoliciesAPIService.C1ApiPolicyV1PoliciesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/policies"

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
