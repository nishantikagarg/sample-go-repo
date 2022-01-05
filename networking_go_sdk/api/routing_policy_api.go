//Api classes for networking's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/networking_go_sdk/v16/client"
	import1 "github.com/nishantikagarg/sample-go-repo/networking_go_sdk/v16/models/networking/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
    "strings"
)

type RoutingPolicyApi struct {
  ApiClient *client.ApiClient
}

func NewRoutingPolicyApi() *RoutingPolicyApi {
	a := &RoutingPolicyApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CreateRoutingPolicy
    Create a Routing Policy.

    parameters:-
    -> body (networking.v4.config.RoutingPolicy) (required) : Schema to configure a Routing Policy.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *RoutingPolicyApi) CreateRoutingPolicy(body *import1.RoutingPolicy) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/routing-policies"

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
    unmarshalledResp := new(import1.TaskReferenceApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DeleteRoutingPolicy
    Delete the Routing Policy corresponding to the extId.

    parameters:-
    -> extId (string) (required) : extId of the Routing Policy.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *RoutingPolicyApi) DeleteRoutingPolicy(extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/routing-policies/{extId}"


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
    unmarshalledResp := new(import1.TaskReferenceApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetRoutingPolicy
    Get a single Routing Policy corresponding to the extId.

    parameters:-
    -> extId (string) (required) : extId of the Routing Policy.

    returns: (networking.v4.config.RoutingPolicyApiResponse, error)
*/
func (api *RoutingPolicyApi) GetRoutingPolicy(extId string) (*import1.RoutingPolicyApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/routing-policies/{extId}"


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
    unmarshalledResp := new(import1.RoutingPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListRoutingPolicies
    Get a list of Routing Policies.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - name - priority - vpcExtId 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - name - priority 

    returns: (networking.v4.config.RoutingPolicyListApiResponse, error)
*/
func (api *RoutingPolicyApi) ListRoutingPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string) (*import1.RoutingPolicyListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/routing-policies"


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
    unmarshalledResp := new(import1.RoutingPolicyListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    PutRoutingPolicy
    Update the Routing Policy corresponding to the extId.

    parameters:-
    -> body (networking.v4.config.RoutingPolicy) (required) : Schema to configure a Routing Policy.
    -> extId (string) (required) : extId of the Routing Policy.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *RoutingPolicyApi) PutRoutingPolicy(body *import1.RoutingPolicy, extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/routing-policies/{extId}"

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
    unmarshalledResp := new(import1.TaskReferenceApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

