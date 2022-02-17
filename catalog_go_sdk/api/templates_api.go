//Api classes for catalog's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/catalog_go_sdk/v16/client"
	"strings"
	import2 "github.com/nishantikagarg/sample-go-repo/catalog_go_sdk/v16/models/vmm/v4/templates"
	"encoding/json"
	"net/http"
    "net/url"
)

type TemplatesApi struct {
  ApiClient *client.ApiClient
}

func NewTemplatesApi() *TemplatesApi {
	a := &TemplatesApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    CancelGuestChanges
    Cancel the guest OS update operation.

    parameters:-
    -> templateExtId (string) (required)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) CancelGuestChanges(templateExtId string) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/$actions/cancel-guest-changes"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    CompleteGuestChanges
    Complete the Guest OS modifications.

    parameters:-
    -> body (vmm.v4.templates.TemplateCompleteGuestChanges) (required)
    -> templateExtId (string) (required)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) CompleteGuestChanges(body *import2.TemplateCompleteGuestChanges, templateExtId string) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/$actions/complete-guest-changes"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
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
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    CreateTemplate
    A Template is an entity that captures configuration and data at rest for a parent entity like a VM. A Template can be used to create an entity similar in configuration and data to the parent entity. A Template is an abstraction to allow grouping Templates as versions inside the template. A template also maintains a notion of an gold template version which can be different from the latest template version. Template version numbers are non-negative monotonically increasing integers which also serve the purpose of relative chronological ordering of template versions within the scope of a template. This API allows creating a template by providing a template and version details description.

    parameters:-
    -> body (vmm.v4.templates.Template) (required)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) CreateTemplate(body *import2.Template) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates"

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
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DeleteTemplateById
    This deletes all versions in the template.

    parameters:-
    -> templateExtId (string) (required) : Template UUID

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) DeleteTemplateById(templateExtId string) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
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
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DeleteTemplateVersionByVerNum
    Delete a specific version of the template. If the version happens to be the last remaining version on the template we will disallow the template delete version and to delete the last remaining version please attempt delete of template.

    parameters:-
    -> templateExtId (string) (required)
    -> versionNumber (int) (required)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) DeleteTemplateVersionByVerNum(templateExtId string, versionNumber int) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/versions/{versionNumber}"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"versionNumber"+"}", url.PathEscape(client.ParameterToString(versionNumber, "")), -1)
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
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    DeployVMFromTemplate
    Deploy a VM from the Template.

    parameters:-
    -> body (vmm.v4.templates.TemplateDeployment) (required)
    -> templateExtId (string) (required) : Template UUID

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) DeployVMFromTemplate(body *import2.TemplateDeployment, templateExtId string) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/$actions/deploy"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
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
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetAllTemplateVersions
    Get all existing template versions.

    parameters:-
    -> templateExtId (string) (required) : Template UUID
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> vmSpec (bool) (optional) : Default value is false. Set spec to true to include the VM spec in the template response.
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - versionName 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - createdAt - versionName 

    returns: (vmm.v4.templates.TemplateVersionListApiResponse, error)
*/
func (api *TemplatesApi) GetAllTemplateVersions(templateExtId string, page_ *int, limit_ *int, vmSpec *bool, filter_ *string, orderby_ *string) (*import2.TemplateVersionListApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/versions"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
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
	if vmSpec != nil {
		queryParams.Add("vmSpec", client.ParameterToString(*vmSpec, ""))
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
    unmarshalledResp := new(import2.TemplateVersionListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetTemplateById
    Get Template details

    parameters:-
    -> templateExtId (string) (required) : Template UUID
    -> vmSpec (bool) (optional) : Default value is false. Set spec to true to include the VM spec in the template response.

    returns: (vmm.v4.templates.TemplateApiResponse, error)
*/
func (api *TemplatesApi) GetTemplateById(templateExtId string, vmSpec *bool) (*import2.TemplateApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if vmSpec != nil {
		queryParams.Add("vmSpec", client.ParameterToString(*vmSpec, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.TemplateApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetTemplateListMetadata
    Get all the Templates Metadata.

    parameters:-
    -> vmSpec (bool) (optional)
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - templateName 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - lastUpdatedAt - templateName 

    returns: (vmm.v4.templates.TemplateListApiResponse, error)
*/
func (api *TemplatesApi) GetTemplateListMetadata(vmSpec *bool, page_ *int, limit_ *int, filter_ *string, orderby_ *string) (*import2.TemplateListApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if vmSpec != nil {
		queryParams.Add("vmSpec", client.ParameterToString(*vmSpec, ""))
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
    unmarshalledResp := new(import2.TemplateListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetTemplateVersionByVerNum
    A description for template version.

    parameters:-
    -> templateExtId (string) (required)
    -> versionNumber (int) (required)
    -> vmSpec (bool) (optional)

    returns: (vmm.v4.templates.TemplateVersionApiResponse, error)
*/
func (api *TemplatesApi) GetTemplateVersionByVerNum(templateExtId string, versionNumber int, vmSpec *bool) (*import2.TemplateVersionApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/versions/{versionNumber}"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"versionNumber"+"}", url.PathEscape(client.ParameterToString(versionNumber, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if vmSpec != nil {
		queryParams.Add("vmSpec", client.ParameterToString(*vmSpec, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.TemplateVersionApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    InitiateGuestChanges
    Modify the Guest OS of the Template.

    parameters:-
    -> templateExtId (string) (required)
    -> body (vmm.v4.templates.TemplateGuestChanges) (optional)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) InitiateGuestChanges(templateExtId string, body *import2.TemplateGuestChanges) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/$actions/initiate-guest-changes"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
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
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    PublishTemplate
    Set the version as Gold Version.

    parameters:-
    -> templateExtId (string) (required)
    -> versionNumber (int) (required)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) PublishTemplate(templateExtId string, versionNumber int) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}/versions/{versionNumber}/$actions/publish"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"versionNumber"+"}", url.PathEscape(client.ParameterToString(versionNumber, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateTemplate
    Modify Template details or create a version by specifying the version reference  details.

    parameters:-
    -> templateExtId (string) (required) : Template UUID
    -> body (vmm.v4.templates.Template) (optional)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) UpdateTemplate(templateExtId string, body *import2.Template) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
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
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    UpdateTemplateMetadata
    Update the Template Metadata (Name, Description etc)

    parameters:-
    -> templateExtId (string) (required) : Template UUID
    -> body (vmm.v4.templates.Template) (optional)

    returns: (vmm.v4.templates.TemplatesTaskApiResponse, error)
*/
func (api *TemplatesApi) UpdateTemplateMetadata(templateExtId string, body *import2.Template) (*import2.TemplatesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/templates/{templateExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"templateExtId"+"}", url.PathEscape(client.ParameterToString(templateExtId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/json"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPatch, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.TemplatesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

