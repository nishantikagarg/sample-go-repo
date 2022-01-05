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

type DirectConnectVirtualInterfaceApi struct {
  ApiClient *client.ApiClient
}

func NewDirectConnectVirtualInterfaceApi() *DirectConnectVirtualInterfaceApi {
	a := &DirectConnectVirtualInterfaceApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CreateDirectConnectVirtualInterface
    Create direct connect virtual interface

    parameters:-
    -> body (networking.v4.config.DirectConnectVirtualInterface) (required) : Create direct connect virtual interface request body

    returns: (networking.v4.config.TaskApiResponse, error)
*/
func (api *DirectConnectVirtualInterfaceApi) CreateDirectConnectVirtualInterface(body *import1.DirectConnectVirtualInterface) (*import1.TaskApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/direct-connect-virtual-interfaces"

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
    unmarshalledResp := new(import1.TaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DeleteDirectConnectVirtualInterface
    Delete direct connect virtual interface request body

    parameters:-
    -> extId (string) (required) : Direct connect virtual interface uuid

    returns: (networking.v4.config.TaskApiResponse, error)
*/
func (api *DirectConnectVirtualInterfaceApi) DeleteDirectConnectVirtualInterface(extId string) (*import1.TaskApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/direct-connect-virtual-interfaces/{extId}"


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
    unmarshalledResp := new(import1.TaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetDirectConnectVirtualInterface
    Get direct connect virtual interface

    parameters:-
    -> extId (string) (required) : Direct connect virtual interface uuid

    returns: (networking.v4.config.DirectConnectVirtualInterfaceApiResponse, error)
*/
func (api *DirectConnectVirtualInterfaceApi) GetDirectConnectVirtualInterface(extId string) (*import1.DirectConnectVirtualInterfaceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/direct-connect-virtual-interfaces/{extId}"


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
    unmarshalledResp := new(import1.DirectConnectVirtualInterfaceApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListDirectConnectVirtualInterfaces
    List direct connect virtual interfaces

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 

    returns: (networking.v4.config.DirectConnectVirtualInterfaceListApiResponse, error)
*/
func (api *DirectConnectVirtualInterfaceApi) ListDirectConnectVirtualInterfaces(page_ *int, limit_ *int) (*import1.DirectConnectVirtualInterfaceListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/direct-connect-virtual-interfaces"


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
    unmarshalledResp := new(import1.DirectConnectVirtualInterfaceListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateDirectConnectVirtualInterface
    Update direct connect virtual interface

    parameters:-
    -> body (networking.v4.config.DirectConnectVirtualInterface) (required) : Update direct connect virtual interface request body
    -> extId (string) (required) : Direct connect virtual interface uuid

    returns: (networking.v4.config.TaskApiResponse, error)
*/
func (api *DirectConnectVirtualInterfaceApi) UpdateDirectConnectVirtualInterface(body *import1.DirectConnectVirtualInterface, extId string) (*import1.TaskApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/direct-connect-virtual-interfaces/{extId}"

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
    unmarshalledResp := new(import1.TaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

