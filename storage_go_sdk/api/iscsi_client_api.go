//Api classes for storage's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/storage_go_sdk/v16/client"
	"strings"
	import1 "github.com/nishantikagarg/sample-go-repo/storage_go_sdk/v16/models/storage/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type IscsiClientApi struct {
  ApiClient *client.ApiClient
}

func NewIscsiClientApi() *IscsiClientApi {
	a := &IscsiClientApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    GetIscsiClientById
    Query the iSCSI client identified by {clientId}.

    parameters:-
    -> clientId (string) (required) : The external identifier of the iSCSI client.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetIscsiClientApiResponse, error)
*/
func (api *IscsiClientApi) GetIscsiClientById(clientId string, args ...map[string]interface{}) (*import1.GetIscsiClientApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/iscsi-clients/{clientId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"clientId"+"}", url.PathEscape(client.ParameterToString(clientId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.GetIscsiClientApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetIscsiClients
    Query for list of iSCSI clients.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetIscsiClientsApiResponse, error)
*/
func (api *IscsiClientApi) GetIscsiClients(page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.GetIscsiClientsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/iscsi-clients"


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

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.GetIscsiClientsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateIscsiClientById
    Modify an existing iSCSI client configuration identified by {clientId}.

    parameters:-
    -> body (storage.v4.config.IscsiClient) (required) : A model that represents iSCSI client that can be associated with a Volume Group as an external attachment.
    -> clientId (string) (required) : The external identifier of the iSCSI client.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.UpdateIscsiClientApiResponse, error)
*/
func (api *IscsiClientApi) UpdateIscsiClientById(body *import1.IscsiClient, clientId string, args ...map[string]interface{}) (*import1.UpdateIscsiClientApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/iscsi-clients/{clientId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"clientId"+"}", url.PathEscape(client.ParameterToString(clientId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/json"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPatch, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateIscsiClientApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

