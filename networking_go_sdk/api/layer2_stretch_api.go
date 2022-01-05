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

type Layer2StretchApi struct {
  ApiClient *client.ApiClient
}

func NewLayer2StretchApi() *Layer2StretchApi {
	a := &Layer2StretchApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CreateLayer2Stretch
    Create a Layer2Stretch configuration.

    parameters:-
    -> body (networking.v4.config.Layer2Stretch) (required) : Request schema to create the Layer2Stretch configuration.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *Layer2StretchApi) CreateLayer2Stretch(body *import1.Layer2Stretch) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/layer2-stretches"

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
    DeleteLayer2Stretch
    Delete the specified Layer2Stretch configuration.

    parameters:-
    -> extId (string) (required) : The UUID of the Layer2Stretch configuration.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *Layer2StretchApi) DeleteLayer2Stretch(extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/layer2-stretches/{extId}"


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
    GetLayer2Stretch
    Get the Layer2Stretch configuration with the specified UUID.

    parameters:-
    -> extId (string) (required) : The UUID of the Layer2Stretch configuration.

    returns: (networking.v4.config.Layer2StretchApiResponse, error)
*/
func (api *Layer2StretchApi) GetLayer2Stretch(extId string) (*import1.Layer2StretchApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/layer2-stretches/{extId}"


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
    unmarshalledResp := new(import1.Layer2StretchApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetLayer2StretchRelatedEntities
    Get the stretch-related entities from the specified PC cluster.

    parameters:-
    -> pcClusterExtId (string) (required) : The UUID of the PC cluster to retrieve stretch-related entities from.

    returns: (networking.v4.config.Layer2StretchRelatedEntitiesApiResponse, error)
*/
func (api *Layer2StretchApi) GetLayer2StretchRelatedEntities(pcClusterExtId string) (*import1.Layer2StretchRelatedEntitiesApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/layer2-stretches/related-entities/{pcClusterExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"pcClusterExtId"+"}", url.PathEscape(client.ParameterToString(pcClusterExtId, "")), -1)
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
    unmarshalledResp := new(import1.Layer2StretchRelatedEntitiesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListLayer2Stretches
    Get the list of existing Layer2Stretch configurations.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 

    returns: (networking.v4.config.Layer2StretchListApiResponse, error)
*/
func (api *Layer2StretchApi) ListLayer2Stretches(page_ *int, limit_ *int) (*import1.Layer2StretchListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/layer2-stretches"


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
    unmarshalledResp := new(import1.Layer2StretchListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateLayer2Stretch
    Update the specified Layer2Stretch configuration.

    parameters:-
    -> body (networking.v4.config.Layer2Stretch) (required) : Request schema to update the specified Layer2Stretch configuration.
    -> extId (string) (required) : The UUID of the Layer2Stretch configuration.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *Layer2StretchApi) UpdateLayer2Stretch(body *import1.Layer2Stretch, extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/layer2-stretches/{extId}"

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

