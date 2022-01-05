//Api classes for flow's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/flow_go_sdk/v16/client"
	"reflect"
	import1 "github.com/nishantikagarg/sample-go-repo/flow_go_sdk/v16/models/microseg/v4/policies"
	"encoding/json"
	"net/http"
    "net/url"
    "strings"
)

type NetworkSecurityPoliciesApi struct {
  ApiClient *client.ApiClient
}

func NewNetworkSecurityPoliciesApi() *NetworkSecurityPoliciesApi {
	a := &NetworkSecurityPoliciesApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    Returns a list of Flow 2.0 policies for the input 1.0 security rules.
    Returns a list of Flow 2.0 policies for the input 1.0 security rules.

    parameters:-
    -> body (microseg.v4.policies.MigrationDryRunInput) (required)

    returns: (microseg.v4.policies.NetworkSecurityPolicyGetListResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) ConfigMigrationDryRun(body *import1.MigrationDryRunInput) (*import1.NetworkSecurityPolicyGetListResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/$actions/migrate-config/dry-run"

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
    unmarshalledResp := new(import1.NetworkSecurityPolicyGetListResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Creates a network security policy.
    Creates a network security policy.

    parameters:-
    -> body (microseg.v4.policies.NetworkSecurityPolicy) (required)

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) CreateNetworkSecurityPolicy(body *import1.NetworkSecurityPolicy) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies"

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
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Deletes many network security policies by UUID.
    Performs a batch delete operation on one or more network security policies. A single task is returned for tracking.

    parameters:-
    -> body (microseg.v4.policies.PolicyExtIdList) (required) : Deletes many network security policies by UUID.

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) DeleteNetworkSecurityPoliciesByExtIds(body *import1.PolicyExtIdList) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/$actions/delete"

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
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Deletes a network security policy by UUID.
    Deletes a network security policy by UUID.

    parameters:-
    -> extId (string) (required) : Network security policy UUID.

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) DeleteNetworkSecurityPolicyByExtId(extId string) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/{extId}"


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
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Gets a network security policy by UUID.
    Gets a network security policy by UUID.

    parameters:-
    -> extId (string) (required) : Network security policy UUID.
    -> includeRefNames (bool) (optional) : A URL query parameter that specifies whether to include names for references. Note that this may increase request completion times.

    returns: (microseg.v4.policies.NetworkSecurityPolicyGetResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) GetNetworkSecurityPolicyByExtId(extId string, includeRefNames *bool) (*import1.NetworkSecurityPolicyGetResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/{extId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if includeRefNames != nil {
		queryParams.Add("includeRefNames", client.ParameterToString(*includeRefNames, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.NetworkSecurityPolicyGetResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Gets counts of policies grouped in requested manner.
    Gets counts of policies grouped in requested manner.

    parameters:-
    -> countFieldName ([]string) (optional) : The field to return counts for. Can be provided multiple times to get counts for multiple fields at once. Supported values are [type, state].
    -> name ([]string) (optional) : Filters returned policies by name. Regular expressions and repetition are supported.
    -> type_ ([]string) (optional) : Filters returned policies by type. Only exact values and repetition are supported.
    -> securedCategoryExtId ([]string) (optional) : Filters returned policies by secured category external id (UUID). Repetition is supported.
    -> status ([]string) (optional) : Filters returned policies by status. Only exact values and repetition are supported.
    -> sourceCategoryExtId ([]string) (optional) : Filters returned policies by rule source category external id (UUID). Repetition is supported.
    -> sourceSubnetFilter ([]string) (optional) : Filters returned policies by rule source subnet. Repetition is supported. 'ALL' is a valid value.
    -> sourceAddressGroupExtId ([]string) (optional) : Filters returned policies by rule source address group external id (UUID). Repetition is supported.
    -> destinationCategoryExtId ([]string) (optional) : Filters returned policies by rule source category external id (UUID). Repetition is supported.
    -> destinationSubnetFilter ([]string) (optional) : Filters returned policies by rule source subnet. Repetition is supported. 'ALL' is a valid value.
    -> destinationAddressGroupExtId ([]string) (optional) : Filters returned policies by rule source address group external id (UUID). Repetition is supported.
    -> serviceExtId ([]string) (optional) : Filters returned policies by rule service external id (UUID). Repetition is supported.
    -> serviceTCPFilter ([]string) (optional) : Filters returned policies by rule TCP port range (for example '1-1000' for TCP ports between 1 to 1000). 'ANY' is a valid value. Repetition is supported.
    -> serviceUDPFilter ([]string) (optional) : Filters returned policies by rule UDP port range (for example '1-1000' for UDP ports between 1 to 1000). 'ANY' is a valid value. Repetition is supported.
    -> serviceICMPFilter ([]string) (optional) : Filters returned policies by rule ICMP type and code (for example '1-100' for type 1 & code 100). 'ANY' is a valid value. Repetition is supported.
    -> lastUpdateTimeRange ([]string) (optional) : Filters returned policies by when they were last updated. This field supports a time range using any of the valid ISO 8601 time interval formats.
    -> creationTimeRange ([]string) (optional) : Filters returned policies by when they were created. This field supports a time range using any of the valid ISO 8601 time interval formats.
    -> vpcExtId ([]string) (optional) : Filters returned policies by VPC id (UUID). Repetition is supported.

    returns: (microseg.v4.policies.NetworkSecurityPolicyGetCountsResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) GetNetworkSecurityPolicyCounts(countFieldName *[]string, name *[]string, type_ *[]string, securedCategoryExtId *[]string, status *[]string, sourceCategoryExtId *[]string, sourceSubnetFilter *[]string, sourceAddressGroupExtId *[]string, destinationCategoryExtId *[]string, destinationSubnetFilter *[]string, destinationAddressGroupExtId *[]string, serviceExtId *[]string, serviceTCPFilter *[]string, serviceUDPFilter *[]string, serviceICMPFilter *[]string, lastUpdateTimeRange *[]string, creationTimeRange *[]string, vpcExtId *[]string) (*import1.NetworkSecurityPolicyGetCountsResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/counts"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if countFieldName != nil {
		t := *countFieldName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("countFieldName", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("countFieldName", client.ParameterToString(t, "multi"))
		}
	}
	if name != nil {
		t := *name
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("name", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("name", client.ParameterToString(t, "multi"))
		}
	}
	if type_ != nil {
		t := *type_
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("type", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("type", client.ParameterToString(t, "multi"))
		}
	}
	if securedCategoryExtId != nil {
		t := *securedCategoryExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("securedCategoryExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("securedCategoryExtId", client.ParameterToString(t, "multi"))
		}
	}
	if status != nil {
		t := *status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("status", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("status", client.ParameterToString(t, "multi"))
		}
	}
	if sourceCategoryExtId != nil {
		t := *sourceCategoryExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("sourceCategoryExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("sourceCategoryExtId", client.ParameterToString(t, "multi"))
		}
	}
	if sourceSubnetFilter != nil {
		t := *sourceSubnetFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("sourceSubnetFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("sourceSubnetFilter", client.ParameterToString(t, "multi"))
		}
	}
	if sourceAddressGroupExtId != nil {
		t := *sourceAddressGroupExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("sourceAddressGroupExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("sourceAddressGroupExtId", client.ParameterToString(t, "multi"))
		}
	}
	if destinationCategoryExtId != nil {
		t := *destinationCategoryExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("destinationCategoryExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("destinationCategoryExtId", client.ParameterToString(t, "multi"))
		}
	}
	if destinationSubnetFilter != nil {
		t := *destinationSubnetFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("destinationSubnetFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("destinationSubnetFilter", client.ParameterToString(t, "multi"))
		}
	}
	if destinationAddressGroupExtId != nil {
		t := *destinationAddressGroupExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("destinationAddressGroupExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("destinationAddressGroupExtId", client.ParameterToString(t, "multi"))
		}
	}
	if serviceExtId != nil {
		t := *serviceExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceExtId", client.ParameterToString(t, "multi"))
		}
	}
	if serviceTCPFilter != nil {
		t := *serviceTCPFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceTCPFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceTCPFilter", client.ParameterToString(t, "multi"))
		}
	}
	if serviceUDPFilter != nil {
		t := *serviceUDPFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceUDPFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceUDPFilter", client.ParameterToString(t, "multi"))
		}
	}
	if serviceICMPFilter != nil {
		t := *serviceICMPFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceICMPFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceICMPFilter", client.ParameterToString(t, "multi"))
		}
	}
	if lastUpdateTimeRange != nil {
		t := *lastUpdateTimeRange
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("lastUpdateTimeRange", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("lastUpdateTimeRange", client.ParameterToString(t, "multi"))
		}
	}
	if creationTimeRange != nil {
		t := *creationTimeRange
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("creationTimeRange", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("creationTimeRange", client.ParameterToString(t, "multi"))
		}
	}
	if vpcExtId != nil {
		t := *vpcExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("vpcExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("vpcExtId", client.ParameterToString(t, "multi"))
		}
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.NetworkSecurityPolicyGetCountsResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Export all network security rules.
    Export all network security policies to save and for subsequent import.

    parameters:-

    returns: (microseg.v4.policies.NetworkSecurityPolicyExportResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) GetNetworkSecurityPolicyExport() (*import1.NetworkSecurityPolicyExportResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/export"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/octet-stream", "application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.NetworkSecurityPolicyExportResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Reports on the impact of importing the policies.
    Returns a summary on the impact of importing the policy data.

    parameters:-
    -> body (string) (required)

    returns: (microseg.v4.policies.NetworkSecurityPolicyImportResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) GetNetworkSecurityPolicyImportDryrun(body *string) (*import1.NetworkSecurityPolicyImportResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/$actions/import/dry_run"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/octet-stream"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.NetworkSecurityPolicyImportResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Gets a list of network security policies.
    Gets a list of network security policies.

    parameters:-
    -> includeRefNames (bool) (optional) : A URL query parameter that specifies whether to include names for references. Note that this may increase request completion times.
    -> name ([]string) (optional) : Filters returned policies by name. Regular expressions and repetition are supported.
    -> type_ ([]string) (optional) : Filters returned policies by type. Only exact values and repetition are supported.
    -> securedCategoryExtId ([]string) (optional) : Filters returned policies by secured category external id (UUID). Repetition is supported.
    -> status ([]string) (optional) : Filters returned policies by status. Only exact values and repetition are supported.
    -> sourceCategoryExtId ([]string) (optional) : Filters returned policies by rule source category external id (UUID). Repetition is supported.
    -> sourceSubnetFilter ([]string) (optional) : Filters returned policies by rule source subnet. Repetition is supported. 'ALL' is a valid value.
    -> sourceAddressGroupExtId ([]string) (optional) : Filters returned policies by rule source address group external id (UUID). Repetition is supported.
    -> destinationCategoryExtId ([]string) (optional) : Filters returned policies by rule source category external id (UUID). Repetition is supported.
    -> destinationSubnetFilter ([]string) (optional) : Filters returned policies by rule source subnet. Repetition is supported. 'ALL' is a valid value.
    -> destinationAddressGroupExtId ([]string) (optional) : Filters returned policies by rule source address group external id (UUID). Repetition is supported.
    -> serviceExtId ([]string) (optional) : Filters returned policies by rule service external id (UUID). Repetition is supported.
    -> serviceTCPFilter ([]string) (optional) : Filters returned policies by rule TCP port range (for example '1-1000' for TCP ports between 1 to 1000). 'ANY' is a valid value. Repetition is supported.
    -> serviceUDPFilter ([]string) (optional) : Filters returned policies by rule UDP port range (for example '1-1000' for UDP ports between 1 to 1000). 'ANY' is a valid value. Repetition is supported.
    -> serviceICMPFilter ([]string) (optional) : Filters returned policies by rule ICMP type and code (for example '1-100' for type 1 & code 100). 'ANY' is a valid value. Repetition is supported.
    -> lastUpdateTimeRange ([]string) (optional) : Filters returned policies by when they were last updated. This field supports a time range using any of the valid ISO 8601 time interval formats.
    -> creationTimeRange ([]string) (optional) : Filters returned policies by when they were created. This field supports a time range using any of the valid ISO 8601 time interval formats.
    -> vpcExtId ([]string) (optional) : Filters returned policies by VPC id (UUID). Repetition is supported.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : Odata query for filtering list api response
    -> orderby_ (string) (optional) : Odata query for sorting list api response

    returns: (microseg.v4.policies.NetworkSecurityPolicyGetListResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) GetNetworkSecurityPolicyList(includeRefNames *bool, name *[]string, type_ *[]string, securedCategoryExtId *[]string, status *[]string, sourceCategoryExtId *[]string, sourceSubnetFilter *[]string, sourceAddressGroupExtId *[]string, destinationCategoryExtId *[]string, destinationSubnetFilter *[]string, destinationAddressGroupExtId *[]string, serviceExtId *[]string, serviceTCPFilter *[]string, serviceUDPFilter *[]string, serviceICMPFilter *[]string, lastUpdateTimeRange *[]string, creationTimeRange *[]string, vpcExtId *[]string, page_ *int, limit_ *int, filter_ *string, orderby_ *string) (*import1.NetworkSecurityPolicyGetListResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if includeRefNames != nil {
		queryParams.Add("includeRefNames", client.ParameterToString(*includeRefNames, ""))
	}
	if name != nil {
		t := *name
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("name", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("name", client.ParameterToString(t, "multi"))
		}
	}
	if type_ != nil {
		t := *type_
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("type", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("type", client.ParameterToString(t, "multi"))
		}
	}
	if securedCategoryExtId != nil {
		t := *securedCategoryExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("securedCategoryExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("securedCategoryExtId", client.ParameterToString(t, "multi"))
		}
	}
	if status != nil {
		t := *status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("status", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("status", client.ParameterToString(t, "multi"))
		}
	}
	if sourceCategoryExtId != nil {
		t := *sourceCategoryExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("sourceCategoryExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("sourceCategoryExtId", client.ParameterToString(t, "multi"))
		}
	}
	if sourceSubnetFilter != nil {
		t := *sourceSubnetFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("sourceSubnetFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("sourceSubnetFilter", client.ParameterToString(t, "multi"))
		}
	}
	if sourceAddressGroupExtId != nil {
		t := *sourceAddressGroupExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("sourceAddressGroupExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("sourceAddressGroupExtId", client.ParameterToString(t, "multi"))
		}
	}
	if destinationCategoryExtId != nil {
		t := *destinationCategoryExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("destinationCategoryExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("destinationCategoryExtId", client.ParameterToString(t, "multi"))
		}
	}
	if destinationSubnetFilter != nil {
		t := *destinationSubnetFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("destinationSubnetFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("destinationSubnetFilter", client.ParameterToString(t, "multi"))
		}
	}
	if destinationAddressGroupExtId != nil {
		t := *destinationAddressGroupExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("destinationAddressGroupExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("destinationAddressGroupExtId", client.ParameterToString(t, "multi"))
		}
	}
	if serviceExtId != nil {
		t := *serviceExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceExtId", client.ParameterToString(t, "multi"))
		}
	}
	if serviceTCPFilter != nil {
		t := *serviceTCPFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceTCPFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceTCPFilter", client.ParameterToString(t, "multi"))
		}
	}
	if serviceUDPFilter != nil {
		t := *serviceUDPFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceUDPFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceUDPFilter", client.ParameterToString(t, "multi"))
		}
	}
	if serviceICMPFilter != nil {
		t := *serviceICMPFilter
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("serviceICMPFilter", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("serviceICMPFilter", client.ParameterToString(t, "multi"))
		}
	}
	if lastUpdateTimeRange != nil {
		t := *lastUpdateTimeRange
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("lastUpdateTimeRange", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("lastUpdateTimeRange", client.ParameterToString(t, "multi"))
		}
	}
	if creationTimeRange != nil {
		t := *creationTimeRange
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("creationTimeRange", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("creationTimeRange", client.ParameterToString(t, "multi"))
		}
	}
	if vpcExtId != nil {
		t := *vpcExtId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("vpcExtId", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("vpcExtId", client.ParameterToString(t, "multi"))
		}
	}
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
    unmarshalledResp := new(import1.NetworkSecurityPolicyGetListResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Sets hitlog state for many security policies by UUID.
    Performs a batch set operation on one or more network security policies that updates the policy hitlog state. A single task is returned for tracking.

    parameters:-
    -> body (microseg.v4.policies.PolicyExtIdListWithBooleanValue) (required) : Sets hitlog state for many security policies by UUID.

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) SetNetworkSecurityPoliciesHitlogByExtIds(body *import1.PolicyExtIdListWithBooleanValue) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/$actions/set-enable-hitlog"

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
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Sets IPv6 allow state for many security policies by UUID.
    Performs a batch set operation on one or more network security policies that updates the policy IPv6 allow state. A single task is returned for tracking.

    parameters:-
    -> body (microseg.v4.policies.PolicyExtIdListWithBooleanValue) (required) : Sets IPv6 allow state for many security policies by UUID.

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) SetNetworkSecurityPoliciesIPv6ByExtIds(body *import1.PolicyExtIdListWithBooleanValue) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/$actions/set-allow-ipv6-traffic"

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
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Sets state for many security policies by UUID.
    Performs a batch set operation on one or more network security policies that updates the policy state. A single task is returned for tracking.

    parameters:-
    -> body (microseg.v4.policies.PolicyExtIdListWithStateValue) (required) : Sets state for many security policies by UUID.

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) SetNetworkSecurityPoliciesStateByExtIds(body *import1.PolicyExtIdListWithStateValue) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/$actions/set-state"

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
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Imports all the network security rules specified by the data.
    Imports previously exported network security rules.

    parameters:-
    -> body (string) (required)

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) SetNetworkSecurityPolicyImportApply(body *string) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/$actions/import/apply"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/octet-stream"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Updates a network security policy by UUID.
    Updates a network security policy by UUID.

    parameters:-
    -> body (microseg.v4.policies.NetworkSecurityPolicy) (required) : Updates a network security policy by UUID.
    -> extId (string) (required) : Network security policy UUID.

    returns: (microseg.v4.policies.NetworkSecurityPolicyTaskResponse, error)
*/
func (api *NetworkSecurityPoliciesApi) UpdateNetworkSecurityPolicyByExtId(body *import1.NetworkSecurityPolicy, extId string) (*import1.NetworkSecurityPolicyTaskResponse, error) {
    uri := "/api/microseg/v4.0.a1/policies/{extId}"

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
    unmarshalledResp := new(import1.NetworkSecurityPolicyTaskResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

