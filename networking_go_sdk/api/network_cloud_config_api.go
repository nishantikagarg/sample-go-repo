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

type NetworkCloudConfigApi struct {
  ApiClient *client.ApiClient
}

func NewNetworkCloudConfigApi() *NetworkCloudConfigApi {
	a := &NetworkCloudConfigApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CreateNetworkCloudConfig
    Create a Network Cloud Config.

    parameters:-
    -> body (networking.v4.config.NetworkCloudConfig) (required) : Request schema to create the Network Cloud Config.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *NetworkCloudConfigApi) CreateNetworkCloudConfig(body *import1.NetworkCloudConfig) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/network-cloud-configs"

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
    DeleteNetworkCloudConfig
    Delete the specified Network Cloud Config.

    parameters:-
    -> extId (string) (required) : The UUID of the Network Cloud Config.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *NetworkCloudConfigApi) DeleteNetworkCloudConfig(extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/network-cloud-configs/{extId}"


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
    GetNetworkCloudConfig
    Get the Network Cloud Config with the specified UUID.

    parameters:-
    -> extId (string) (required) : The UUID of the Network Cloud Config.

    returns: (networking.v4.config.NetworkCloudConfigApiResponse, error)
*/
func (api *NetworkCloudConfigApi) GetNetworkCloudConfig(extId string) (*import1.NetworkCloudConfigApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/network-cloud-configs/{extId}"


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
    unmarshalledResp := new(import1.NetworkCloudConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListNetworkCloudConfigs
    Get the list of existing Network Cloud Configs.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 

    returns: ([]networking.v4.config.NetworkCloudConfigListApiResponse, error)
*/
func (api *NetworkCloudConfigApi) ListNetworkCloudConfigs(page_ *int, limit_ *int) (*[]networking.v4.config.NetworkCloudConfigListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/network-cloud-configs"


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
    unmarshalledResp := new([]networking.v4.config.NetworkCloudConfigListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

