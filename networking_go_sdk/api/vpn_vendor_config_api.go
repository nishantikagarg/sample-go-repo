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

type VpnVendorConfigApi struct {
  ApiClient *client.ApiClient
}

func NewVpnVendorConfigApi() *VpnVendorConfigApi {
	a := &VpnVendorConfigApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    GetVpnVendorConfig
    Get VPN vendor device configuration steps. If device version is not specified the configuration steps of the latest available device version are returned.

    parameters:-
    -> vpnConnectionId (string) (required) : VPN connection UUID.
    -> vendorName (string) (required) : VPN device vendor name.
    -> deviceVersion (string) (optional) : VPN device version.

    returns: (string, error)
*/
func (api *VpnVendorConfigApi) GetVpnVendorConfig(vpnConnectionId string, vendorName *string, deviceVersion *string) (*string, error) {
    uri := "/api/networking/v4.0.a1/config/vpn-vendor-configs/{vpnConnectionId}"

    // verify the required parameter 'vendorName' is set
	if nil == vendorName {
		return nil, client.ReportError("vendorName is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"vpnConnectionId"+"}", url.PathEscape(client.ParameterToString(vpnConnectionId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"text/plain", "application/json"} 

    // Query Params
	queryParams.Add("vendorName", client.ParameterToString(*vendorName, ""))
	if deviceVersion != nil {
		queryParams.Add("deviceVersion", client.ParameterToString(*deviceVersion, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(string)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ListAvailableVpnVendorConfigs
    List third-party VPN vendors and devices for which configuration steps are available to download.

    parameters:-
    -> vpnConnectionId (string) (required) : VPN connection UUID.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 

    returns: (networking.v4.config.VpnVendorListApiResponse, error)
*/
func (api *VpnVendorConfigApi) ListAvailableVpnVendorConfigs(vpnConnectionId string, page_ *int, limit_ *int) (*import1.VpnVendorListApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/vpn-vendor-configs/{vpnConnectionId}/available"


    // Path Params
    uri = strings.Replace(uri, "{"+"vpnConnectionId"+"}", url.PathEscape(client.ParameterToString(vpnConnectionId, "")), -1)
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
    unmarshalledResp := new(import1.VpnVendorListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

