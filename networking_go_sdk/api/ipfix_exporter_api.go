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

type IPFIXExporterApi struct {
  ApiClient *client.ApiClient
}

func NewIPFIXExporterApi() *IPFIXExporterApi {
	a := &IPFIXExporterApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CreateIpfixExporter
    Create an IPFIX Exporter.

    parameters:-
    -> body (networking.v4.config.IPFIXExporter) (required) : Request schema to create the IPFIX Exporter.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *IPFIXExporterApi) CreateIpfixExporter(body *import1.IPFIXExporter) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/ipfix-exporters"

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
    DeleteIpfixExporter
    Delete the specified IPFIX Exporter.

    parameters:-
    -> extId (string) (required) : UUID of IPFIX Exporter.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *IPFIXExporterApi) DeleteIpfixExporter(extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/ipfix-exporters/{extId}"


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
    GetIpfixExporter
    The UUID of the IPFIX Exporter.

    parameters:-
    -> extId (string) (required) : UUID of IPFIX Exporter.

    returns: (networking.v4.config.IPFIXExporterApiResponse, error)
*/
func (api *IPFIXExporterApi) GetIpfixExporter(extId string) (*import1.IPFIXExporterApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/ipfix-exporters/{extId}"


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
    unmarshalledResp := new(import1.IPFIXExporterApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListIpfixExporters
    Get the list of existing IPFIX Exporters.

    parameters:-

    returns: (networking.v4.config.IPFIXExporterListApiResponse, error)
*/
func (api *IPFIXExporterApi) ListIpfixExporters() (*import1.IPFIXExporterListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/ipfix-exporters"


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
    unmarshalledResp := new(import1.IPFIXExporterListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateIpfixExporter
    Update the specified IPFIX Exporter.

    parameters:-
    -> body (networking.v4.config.IPFIXExporter) (required) : Request schema to update the specified IPFIX Exporter.
    -> extId (string) (required) : UUID of IPFIX Exporter.

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *IPFIXExporterApi) UpdateIpfixExporter(body *import1.IPFIXExporter, extId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/ipfix-exporters/{extId}"

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

