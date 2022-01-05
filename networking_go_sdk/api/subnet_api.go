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

type SubnetApi struct {
  ApiClient *client.ApiClient
}

func NewSubnetApi() *SubnetApi {
	a := &SubnetApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CreateSubnet
    Create a Subnet.

    parameters:-
    -> body (networking.v4.config.Subnet) (required) : Request schema to create the Subnet.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetApi) CreateSubnet(body *import1.Subnet) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets"

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
    DeleteSubnet
    Delete the specified Subnet.

    parameters:-
    -> extId (string) (required) : The UUID of the Subnet.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetApi) DeleteSubnet(extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets/{extId}"


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
    GetSubnet
    Get the Subnet with the specified UUID.

    parameters:-
    -> extId (string) (required) : The UUID of the Subnet.

    returns: (networking.v4.config.SubnetApiResponse, error)
*/
func (api *SubnetApi) GetSubnet(extId string) (*import1.SubnetApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets/{extId}"


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
    unmarshalledResp := new(import1.SubnetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListSubnets
    Get the list of existing Subnets.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - clusterReference - isExternal - name - subnetType - vpcReference 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - name 

    returns: (networking.v4.config.SubnetListApiResponse, error)
*/
func (api *SubnetApi) ListSubnets(page_ *int, limit_ *int, filter_ *string, orderby_ *string) (*import1.SubnetListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets"


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
    unmarshalledResp := new(import1.SubnetListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateSubnet
    Update the specified Subnet.

    parameters:-
    -> body (networking.v4.config.Subnet) (required) : Request schema to update the specified Subnet.
    -> extId (string) (required) : The UUID of the Subnet.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetApi) UpdateSubnet(body *import1.Subnet, extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets/{extId}"

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

