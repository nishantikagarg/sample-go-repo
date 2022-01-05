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

type FlowGatewayApi struct {
  ApiClient *client.ApiClient
}

func NewFlowGatewayApi() *FlowGatewayApi {
	a := &FlowGatewayApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CreateAtlasFlowGateway
    Create a Atlas Flow Gateway.

    parameters:-
    -> body (networking.v4.config.FlowGateway) (required) : Request schema to create the Atlas Flow Gateway.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *FlowGatewayApi) CreateAtlasFlowGateway(body *import1.FlowGateway) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/flow-gateways"

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
    DeleteAtlasFlowGateway
    Delete the specified Atlas Flow Gateway.

    parameters:-
    -> extId (string) (required) : The UUID of the Atlas Flow Gateway.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *FlowGatewayApi) DeleteAtlasFlowGateway(extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/flow-gateways/{extId}"


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
    GetAtlasFlowGateway
    Get the Atlas Flow Gateway with the specified UUID.

    parameters:-
    -> extId (string) (required) : The UUID of the Atlas Flow Gateway.

    returns: (networking.v4.config.FlowGatewayApiResponse, error)
*/
func (api *FlowGatewayApi) GetAtlasFlowGateway(extId string) (*import1.FlowGatewayApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/flow-gateways/{extId}"


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
    unmarshalledResp := new(import1.FlowGatewayApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    KeepAliveAtlasFlowGateway
    Send Keep alive for a Flow Gateway.

    parameters:-
    -> body (networking.v4.config.FlowGatewayKeepAliveRequest) (required) : Request schema to send Flow Gateway keep alive heartbeat.

    returns: (networking.v4.config.FlowGatewayKeepAliveResponseApiResponse, error)
*/
func (api *FlowGatewayApi) KeepAliveAtlasFlowGateway(body *import1.FlowGatewayKeepAliveRequest) (*import1.FlowGatewayKeepAliveResponseApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/flow-gateways/keep-alive"

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
    unmarshalledResp := new(import1.FlowGatewayKeepAliveResponseApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListAtlasFlowGateways
    Get the list of existing Atlas Flow Gateways.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 

    returns: ([]networking.v4.config.FlowGatewayListApiResponse, error)
*/
func (api *FlowGatewayApi) ListAtlasFlowGateways(page_ *int, limit_ *int) (*[]networking.v4.config.FlowGatewayListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/flow-gateways"


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

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new([]networking.v4.config.FlowGatewayListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateAtlasFlowGateway
    Update the specified Atlas Flow Gateway.

    parameters:-
    -> body (networking.v4.config.FlowGateway) (required) : Request schema to update the specified Atlas Flow Gateway.
    -> extId (string) (required) : The UUID of the Atlas Flow Gateway.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *FlowGatewayApi) UpdateAtlasFlowGateway(body *import1.FlowGateway, extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/flow-gateways/{extId}"

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

