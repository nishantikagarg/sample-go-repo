package api

import (
    "github.com/nutanix-core/ntnx-api-go-sdk-internal/storage_go_sdk/v16/client"
	"strings"
	import2 "github.com/nutanix-core/ntnx-api-go-sdk-internal/storage_go_sdk/v16/models/common/v1/config"
	import1 "github.com/nutanix-core/ntnx-api-go-sdk-internal/storage_go_sdk/v16/models/storage/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type VolumeGroupApi struct {
  ApiClient *client.ApiClient
}

func NewVolumeGroupApi() *VolumeGroupApi {
	a := &VolumeGroupApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    AssociateCategory
    Associate category to a Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> body (storage.v4.config.CategoryEntityReferences) (required) : The list of categories to be associated/disassociated with the Volume Group.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.AssociateCategoryApiResponse, error)
*/
func (api *VolumeGroupApi) AssociateCategory(body *import1.CategoryEntityReferences, volumeGroupExtId string, args ...map[string]interface{}) (*import1.AssociateCategoryApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/associate-category"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.AssociateCategoryApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    AttachIscsiClient
    Attach iSCSI initiator to a Volume Group identified by {volumeGroupExtId}

    parameters:-
    -> body (storage.v4.config.IscsiClient) (required) : A model that represents iSCSI client that can be associated with a Volume Group as an external attachment.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.AttachIscsiClientApiResponse, error)
*/
func (api *VolumeGroupApi) AttachIscsiClient(body *import1.IscsiClient, volumeGroupExtId string, args ...map[string]interface{}) (*import1.AttachIscsiClientApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/attach-iscsi-client"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.AttachIscsiClientApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    AttachVm
    Attach VM to a Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> body (storage.v4.config.VmAttachment) (required) : A model that represents a VM reference that can be associated with a Volume Group as a hypervisor attachment.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.AttachVmApiResponse, error)
*/
func (api *VolumeGroupApi) AttachVm(body *import1.VmAttachment, volumeGroupExtId string, args ...map[string]interface{}) (*import1.AttachVmApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/attach-vm"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.AttachVmApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    CreateVolumeDisk
    Create a new Volume Disk.

    parameters:-
    -> body (storage.v4.config.VolumeDisk) (required) : A model that represents Volume Disk which is associated with a Volume Group, and is supported by a backing file on DSF.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.CreateVolumeDiskApiResponse, error)
*/
func (api *VolumeGroupApi) CreateVolumeDisk(body *import1.VolumeDisk, volumeGroupExtId string, args ...map[string]interface{}) (*import1.CreateVolumeDiskApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.CreateVolumeDiskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    CreateVolumeGroup
    Create a new Volume Group.

    parameters:-
    -> body (storage.v4.config.VolumeGroup) (required) : A model that represents Volume Group resources.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.CreateVolumeGroupApiResponse, error)
*/
func (api *VolumeGroupApi) CreateVolumeGroup(body *import1.VolumeGroup, args ...map[string]interface{}) (*import1.CreateVolumeGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups"

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

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.CreateVolumeGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DeleteVolumeDisk
    Delete a Volume Disk identified by {diskExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> diskExtId (string) (required) : The external identifier of the Volume Disk.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.DeleteVolumeDiskApiResponse, error)
*/
func (api *VolumeGroupApi) DeleteVolumeDisk(volumeGroupExtId string, diskExtId string, args ...map[string]interface{}) (*import1.DeleteVolumeDiskApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks/{diskExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"diskExtId"+"}", url.PathEscape(client.ParameterToString(diskExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DeleteVolumeDiskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DeleteVolumeGroupById
    Delete the Volume Group identified by {extId}.

    parameters:-
    -> extId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.DeleteVolumeGroupApiResponse, error)
*/
func (api *VolumeGroupApi) DeleteVolumeGroupById(extId string, args ...map[string]interface{}) (*import1.DeleteVolumeGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{extId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(extId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DeleteVolumeGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DetachIscsiClient
    Detach iSCSI initiator identified by {clientId} from a Volume Group {volumeGroupExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> clientId (string) (required) : The external identifier of the iSCSI client.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.DetachIscsiClientApiResponse, error)
*/
func (api *VolumeGroupApi) DetachIscsiClient(volumeGroupExtId string, clientId string, args ...map[string]interface{}) (*import1.DetachIscsiClientApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/detach-iscsi-client/{clientId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DetachIscsiClientApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DetachVm
    Detach VM identified by {vmExtId} from a Volume Group {volumeGroupExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> vmExtId (string) (required) : The external identifier of the VM.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.DetachVmApiResponse, error)
*/
func (api *VolumeGroupApi) DetachVm(volumeGroupExtId string, vmExtId string, args ...map[string]interface{}) (*import1.DetachVmApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/detach-vm/{vmExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"vmExtId"+"}", url.PathEscape(client.ParameterToString(vmExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DetachVmApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DisassociateCategory
    Disassociate category from a Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> body (storage.v4.config.CategoryEntityReferences) (required) : The list of categories to be associated/disassociated with the Volume Group.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.DisassociateCategoryApiResponse, error)
*/
func (api *VolumeGroupApi) DisassociateCategory(body *import1.CategoryEntityReferences, volumeGroupExtId string, args ...map[string]interface{}) (*import1.DisassociateCategoryApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/disassociate-category"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DisassociateCategoryApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetCategoryAssociations
    Query the category details which are associated with the Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - extId 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on the following fields: - extId 
    -> expand_ (string) (optional) : Odata query for expanding list api response
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetCategoryAssociationsApiResponse, error)
*/
func (api *VolumeGroupApi) GetCategoryAssociations(volumeGroupExtId string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, args ...map[string]interface{}) (*import1.GetCategoryAssociationsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/category-associations"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
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
    unmarshalledResp := new(import1.GetCategoryAssociationsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetExternalAttachments
    Query the list of external attachments for a Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - clusterReference - extId 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on the following fields: - clusterReference - extId 
    -> expand_ (string) (optional) : Odata query for expanding list api response
    -> select_ (string) (optional) : Odata query for selecting projection
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetExternalAttachmentsApiResponse, error)
*/
func (api *VolumeGroupApi) GetExternalAttachments(volumeGroupExtId string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.GetExternalAttachmentsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/iscsi-client-attachments"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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
    unmarshalledResp := new(import1.GetExternalAttachmentsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetVmAttachments
    Query the list of VM Attachments for a Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> expand_ (string) (optional) : Odata query for expanding list api response
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetVmAttachmentsApiResponse, error)
*/
func (api *VolumeGroupApi) GetVmAttachments(volumeGroupExtId string, page_ *int, limit_ *int, expand_ *string, args ...map[string]interface{}) (*import1.GetVmAttachmentsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/vm-attachments"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
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
    unmarshalledResp := new(import1.GetVmAttachmentsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetVolumeDiskById
    Query the Volume Disk identified by {diskExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> diskExtId (string) (required) : The external identifier of the Volume Disk.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetVolumeDiskApiResponse, error)
*/
func (api *VolumeGroupApi) GetVolumeDiskById(volumeGroupExtId string, diskExtId string, args ...map[string]interface{}) (*import1.GetVolumeDiskApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks/{diskExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"diskExtId"+"}", url.PathEscape(client.ParameterToString(diskExtId, "")), -1)
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
    unmarshalledResp := new(import1.GetVolumeDiskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetVolumeDisks
    Query the list of disks corresponding to a Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - diskSizeBytes - index - storageContainerId 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on the following fields: - diskSizeBytes - index - storageContainerId 
    -> expand_ (string) (optional) : Odata query for expanding list api response
    -> select_ (string) (optional) : Odata query for selecting projection
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetVolumeDisksApiResponse, error)
*/
func (api *VolumeGroupApi) GetVolumeDisks(volumeGroupExtId string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.GetVolumeDisksApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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
    unmarshalledResp := new(import1.GetVolumeDisksApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetVolumeGroupById
    Query the Volume Group identified by {extId}.

    parameters:-
    -> extId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetVolumeGroupApiResponse, error)
*/
func (api *VolumeGroupApi) GetVolumeGroupById(extId string, args ...map[string]interface{}) (*import1.GetVolumeGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{extId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(extId, "")), -1)
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
    unmarshalledResp := new(import1.GetVolumeGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Fetches metadata information associated with a Volume Group.
    Query for metadata information which is associated with the Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetVolumeGroupMetadataInfoApiResponse, error)
*/
func (api *VolumeGroupApi) GetVolumeGroupMetadataInfo(volumeGroupExtId string, args ...map[string]interface{}) (*import1.GetVolumeGroupMetadataInfoApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/metadata-info"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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
    unmarshalledResp := new(import1.GetVolumeGroupMetadataInfoApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetVolumeGroups
    Query the list of Volume Groups.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - clusterReference - extId - isHidden - name 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on the following fields: - clusterReference - extId - name 
    -> expand_ (string) (optional) : Odata query for expanding list api response
    -> select_ (string) (optional) : Odata query for selecting projection
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.GetVolumeGroupsApiResponse, error)
*/
func (api *VolumeGroupApi) GetVolumeGroups(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import1.GetVolumeGroupsApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups"


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
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
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
    unmarshalledResp := new(import1.GetVolumeGroupsApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    MigrateVolumeGroup
    Migrate Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> body (storage.v4.config.VolumeGroupMigrationSpec) (required) : Specification for the migrate action on the Volume Group.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.MigrateVolumeGroupApiResponse, error)
*/
func (api *VolumeGroupApi) MigrateVolumeGroup(body *import1.VolumeGroupMigrationSpec, volumeGroupExtId string, args ...map[string]interface{}) (*import1.MigrateVolumeGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/migrate"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.MigrateVolumeGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    PauseVolumeGroupSynchronousReplication
    API to pause replication for a given Volume Group identified by {volumeGroupExtId} protected using protection policy with synchronous replication RPO.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.PauseVolumeGroupSynchronousReplicationApiResponse, error)
*/
func (api *VolumeGroupApi) PauseVolumeGroupSynchronousReplication(volumeGroupExtId string, args ...map[string]interface{}) (*import1.PauseVolumeGroupSynchronousReplicationApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/pause-synchronous-replication"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.PauseVolumeGroupSynchronousReplicationApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    ResumeVolumeGroupSynchronousReplication
    API to resume replication for a given Volume Group identified by {volumeGroupExtId} protected using protection policy with synchronous replication RPO.

    parameters:-
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.ResumeVolumeGroupSynchronousReplicationApiResponse, error)
*/
func (api *VolumeGroupApi) ResumeVolumeGroupSynchronousReplication(volumeGroupExtId string, args ...map[string]interface{}) (*import1.ResumeVolumeGroupSynchronousReplicationApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/resume-synchronous-replication"


    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.ResumeVolumeGroupSynchronousReplicationApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    RevertVolumeGroup
    Revert Volume Group identified by {volumeGroupExtId}. This does an in-place Volume Group restore from a specified Volume Group recovery point.

    parameters:-
    -> body (storage.v4.config.VolumeGroupRevertSpec) (required) : Specify the Volume Group recovery point Id to which the Volume Group would be reverted.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.RevertVolumeGroupApiResponse, error)
*/
func (api *VolumeGroupApi) RevertVolumeGroup(body *import1.VolumeGroupRevertSpec, volumeGroupExtId string, args ...map[string]interface{}) (*import1.RevertVolumeGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/revert"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.RevertVolumeGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateVolumeDisk
    Update a Volume Disk identified by {diskExtId}.

    parameters:-
    -> body (storage.v4.config.VolumeDisk) (required) : A model that represents Volume Disk which is associated with a Volume Group, and is supported by a backing file on DSF.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> diskExtId (string) (required) : The external identifier of the Volume Disk.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.UpdateVolumeDiskApiResponse, error)
*/
func (api *VolumeGroupApi) UpdateVolumeDisk(body *import1.VolumeDisk, volumeGroupExtId string, diskExtId string, args ...map[string]interface{}) (*import1.UpdateVolumeDiskApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/disks/{diskExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"diskExtId"+"}", url.PathEscape(client.ParameterToString(diskExtId, "")), -1)
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
    unmarshalledResp := new(import1.UpdateVolumeDiskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateVolumeGroupById
    Update details of a Volume Group identified by {extId}.

    parameters:-
    -> body (storage.v4.config.VolumeGroup) (required) : A model that represents Volume Group resources.
    -> extId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.UpdateVolumeGroupApiResponse, error)
*/
func (api *VolumeGroupApi) UpdateVolumeGroupById(body *import1.VolumeGroup, extId string, args ...map[string]interface{}) (*import1.UpdateVolumeGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{extId}"

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
    unmarshalledResp := new(import1.UpdateVolumeGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Updates metadata information associated with a Volume Group.
    Update the metadata information associated with the Volume Group identified by {volumeGroupExtId}.

    parameters:-
    -> body (common.v1.config.Metadata) (required) : The list of metadata information associated with the Volume Group to be updated.
    -> volumeGroupExtId (string) (required) : The external identifier of the Volume Group.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (storage.v4.config.UpdateVolumeGroupMetadataInfoApiResponse, error)
*/
func (api *VolumeGroupApi) UpdateVolumeGroupMetadataInfo(body *import2.Metadata, volumeGroupExtId string, args ...map[string]interface{}) (*import1.UpdateVolumeGroupMetadataInfoApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/storage/v4.0.a3/config/volume-groups/{volumeGroupExtId}/$actions/update-metadata-info"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"volumeGroupExtId"+"}", url.PathEscape(client.ParameterToString(volumeGroupExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateVolumeGroupMetadataInfoApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

