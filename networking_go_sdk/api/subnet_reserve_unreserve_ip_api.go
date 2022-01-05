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

type SubnetReserveUnreserveIpApi struct {
  ApiClient *client.ApiClient
}

func NewSubnetReserveUnreserveIpApi() *SubnetReserveUnreserveIpApi {
	a := &SubnetReserveUnreserveIpApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    FetchSubnetAddressAssignments
    Get the list of assigned and reserved IP addresses on a subnet.

    parameters:-
    -> subnetExtId (string) (required)

    returns: (networking.v4.config.SubnetAddressAssignmentListApiResponse, error)
*/
func (api *SubnetReserveUnreserveIpApi) FetchSubnetAddressAssignments(subnetExtId string) (*import1.SubnetAddressAssignmentListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets/{subnetExtId}/addresses"


    // Path Params
    uri = strings.Replace(uri, "{"+"subnetExtId"+"}", url.PathEscape(client.ParameterToString(subnetExtId, "")), -1)
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
    unmarshalledResp := new(import1.SubnetAddressAssignmentListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ReserveIps
    Reserve IP addresses on a subnet.

    parameters:-
    -> body (networking.v4.config.IpReserveInput) (required) : Request schema to reserve IP addresses on a subnet.
    -> subnetExtId (string) (required)

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetReserveUnreserveIpApi) ReserveIps(body *import1.IpReserveInput, subnetExtId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets/{subnetExtId}/addresses/$actions/reserve"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"subnetExtId"+"}", url.PathEscape(client.ParameterToString(subnetExtId, "")), -1)
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
    UnreserveIps
    Unreserve IP addresses on a subnet.

    parameters:-
    -> body (networking.v4.config.IpUnreserveInput) (required) : Request schema to unreserve IP addresses on a subnet.
    -> subnetExtId (string) (required)

    returns: (networking.v4.config.TaskReferenceApiResponse, error)
*/
func (api *SubnetReserveUnreserveIpApi) UnreserveIps(body *import1.IpUnreserveInput, subnetExtId string) (*import1.TaskReferenceApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/subnets/{subnetExtId}/addresses/$actions/unreserve"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"subnetExtId"+"}", url.PathEscape(client.ParameterToString(subnetExtId, "")), -1)
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

