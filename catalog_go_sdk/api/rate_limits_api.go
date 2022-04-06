//Api classes for catalog's golang SDK
package api

import (
    "catalog_go_sdk/client"
	"strings"
	import1 "catalog_go_sdk/models/vmm/v4/images"
	"encoding/json"
	"net/http"
    "net/url"
)

type RateLimitsApi struct {
  ApiClient *client.ApiClient
}

func NewRateLimitsApi() *RateLimitsApi {
	a := &RateLimitsApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    Create an image rate limit.
    Create an image rate limit.

    parameters:-
    -> body (vmm.v4.images.RateLimit) (required) : Image rate limit to create.

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *RateLimitsApi) CreateRateLimit(body *import1.RateLimit) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/rate-limits"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/json"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete multiple image rate limits.
    Delete the image rate limits with the specified extIds.

    parameters:-
    -> body (vmm.v4.images.RateLimitIdList) (required) : ExtIds of the image rate limits to delete.

    returns: (vmm.v4.images.RateLimitTaskListApiResponse, error)
*/
func (api *RateLimitsApi) DeleteMultipleRateLimits(body *import1.RateLimitIdList) (*import1.RateLimitTaskListApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/rate-limits/$actions/delete"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/json"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.RateLimitTaskListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete an image rate limit.
    Delete the image rate limit with the given extId.

    parameters:-
    -> extId (string) (required) : Image rate limit extId.

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *RateLimitsApi) DeleteRateLimitByExtId(extId string) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/rate-limits/{extId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get an image rate limit.
    Get the image rate limit with the given extId.

    parameters:-
    -> extId (string) (required) : Image rate limit extId.

    returns: (vmm.v4.images.RateLimitApiResponse, error)
*/
func (api *RateLimitsApi) GetRateLimitByExtId(extId string) (*import1.RateLimitApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/rate-limits/{extId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.RateLimitApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List image rate limits.
    List image rate limits.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - name - rateLimitKbps 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - name - rateLimitKbps 

    returns: (vmm.v4.images.RateLimitListApiResponse, error)
*/
func (api *RateLimitsApi) GetRateLimitsList(page_ *int, limit_ *int, filter_ *string, orderby_ *string) (*import1.RateLimitListApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/rate-limits"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {
		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.RateLimitListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Resolve effective rate limit for the Prism Elements.
    Resolve the effective image rate limit for the given Prism Elements when multiple image rate limits apply.

    parameters:-
    -> body (vmm.v4.images.ClusterIdList) (required) : List of cluster extIds for resolving effective rate limits.

    returns: (vmm.v4.images.ResolveRateLimitsApiResponse, error)
*/
func (api *RateLimitsApi) ResolveRateLimits(body *import1.ClusterIdList) (*import1.ResolveRateLimitsApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/rate-limits/$actions/resolve"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/json"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.ResolveRateLimitsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update the image rate limit with the given extId.
    Update an image rate limit.

    parameters:-
    -> body (vmm.v4.images.RateLimit) (required) : Image rate limit to update.
    -> extId (string) (required) : Image rate limit extId.

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *RateLimitsApi) UpdateRateLimitByExtId(body *import1.RateLimit, extId string) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/rate-limits/{extId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/json"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

