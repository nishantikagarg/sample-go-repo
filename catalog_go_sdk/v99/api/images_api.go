//Api classes for Catalog REST API's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/catalog_go_sdk/v99/client"
	import1 "github.com/nishantikagarg/sample-go-repo/catalog_go_sdk/v99/models/vmm/v4/images"
	"encoding/json"
	"net/http"
    "net/url"
    "strings"
)

type ImagesApi struct {
  ApiClient *client.ApiClient
}

func NewImagesApi() *ImagesApi {
	a := &ImagesApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    Create an image.
    Create an image.

    parameters:-
    -> body (vmm.v4.images.Image) (required) : Create image request.

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *ImagesApi) CreateImage(body *import1.Image) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images"

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
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete an image.
    Delete the image with the given extId.

    parameters:-
    -> extId (string) (required)

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *ImagesApi) DeleteImageByExtId(extId string) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/{extId}"


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
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get an image.
    Get image details with the given extId.

    parameters:-
    -> extId (string) (required)

    returns: (vmm.v4.images.ImageApiResponse, error)
*/
func (api *ImagesApi) GetImageByExtId(extId string) (*import1.ImageApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/{extId}"


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
    unmarshalledResp := new(import1.ImageApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get image categories.
    Get image categories for a given extId.

    parameters:-
    -> extId (string) (required)

    returns: (vmm.v4.images.ImageCategoriesApiResponse, error)
*/
func (api *ImagesApi) GetImageCategories(extId string) (*import1.ImageCategoriesApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/{extId}/categories"


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
    unmarshalledResp := new(import1.ImageCategoriesApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get Prism Element locations for an image.
    Get the Prism Element locations where the image is currently available.

    parameters:-
    -> extId (string) (required)

    returns: (vmm.v4.images.ImageClusterLocationListApiResponse, error)
*/
func (api *ImagesApi) GetImageClusterLocations(extId string) (*import1.ImageClusterLocationListApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/{extId}/cluster-locations"


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
    unmarshalledResp := new(import1.ImageClusterLocationListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get placement policies for an image.
    Get the current image placement policies for an image.

    parameters:-
    -> extId (string) (required)

    returns: (vmm.v4.images.ImagePlacementPolicyListApiResponse, error)
*/
func (api *ImagesApi) GetImagePlacementPolicies(extId string) (*import1.ImagePlacementPolicyListApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/{extId}/placement-policies"


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
    unmarshalledResp := new(import1.ImagePlacementPolicyListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List images.
    List images.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - description - name - sizeBytes 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - description - name - sizeBytes 

    returns: (vmm.v4.images.ImageListApiResponse, error)
*/
func (api *ImagesApi) GetImagesList(page_ *int, limit_ *int, filter_ *string, orderby_ *string) (*import1.ImageListApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images"


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

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.ImageListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Import images from Prism Element.
    Import images owned by a registered Prism Element cluster to the current Prism Central.

    parameters:-
    -> body (vmm.v4.images.ClusterImages) (required) : Reference to the Prism Element cluster and its images to import.

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *ImagesApi) ImportImage(body *import1.ClusterImages) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/$actions/import"

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
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Copy images to a remote Prism Central.
    Copy the given images to a remote paired Prism Central.

    parameters:-
    -> body ([]vmm.v4.images.ImageReference) (required) : Ids of the images to copy to the paired Prism Central.

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *ImagesApi) RemoteCopyImage(body *[]import1.ImageReference) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/$actions/remote-copy"

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
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update an image.
    Update the image with the given extId.

    parameters:-
    -> body (vmm.v4.images.Image) (required) : Updated image with the given extId.
    -> extId (string) (required)

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *ImagesApi) UpdateImageByExtId(body *import1.Image, extId string) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/{extId}"

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
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update image categories.
    Replace the attached image categories with the given extId.

    parameters:-
    -> body ([]vmm.v4.images.CategoryReference) (required) : The desired categories for the image.
    -> extId (string) (required)

    returns: (vmm.v4.images.ImagesTaskApiResponse, error)
*/
func (api *ImagesApi) UpdateImageCategories(body *[]import1.CategoryReference, extId string) (*import1.ImagesTaskApiResponse, error) {
    uri := "/api/vmm/v4.0.a1/images/{extId}/categories"

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
    unmarshalledResp := new(import1.ImagesTaskApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
