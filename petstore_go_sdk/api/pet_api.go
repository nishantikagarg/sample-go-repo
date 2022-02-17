//Api classes for petstore's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/client"
	"strings"
	import1 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/models/petstore/v1/defaultapi"
	"reflect"
	"encoding/json"
	"net/http"
    "net/url"
)

type PetApi struct {
  ApiClient *client.ApiClient
}

func NewPetApi() *PetApi {
	a := &PetApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    Add a new pet to the store

    parameters:-
    -> body (petstore.v1.defaultapi.Pet) (required) : Create a new pet in the store

    returns: (petstore.v1.defaultapi.PetApiResponse, error)
*/
func (api *PetApi) AddPet(body *import1.Pet) (*import1.PetApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pets"

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
    unmarshalledResp := new(import1.PetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete pet by ID
    Deletes a single pet

    parameters:-
    -> petId (int64) (required) : ID of pet to return

    returns: (petstore.v1.defaultapi.DeletePetApiResponse, error)
*/
func (api *PetApi) DeletePetById(petId int64) (*import1.DeletePetApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pets/{petId}/downloadImage"


    // Path Params
    uri = strings.Replace(uri, "{"+"petId"+"}", url.PathEscape(client.ParameterToString(petId, "")), -1)
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
    unmarshalledResp := new(import1.DeletePetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Download an image
    Downloads a pet image file

    parameters:-
    -> petId (int64) (required) : ID of pet to download the image with

    returns: (string, error)
*/
func (api *PetApi) GetImageById(petId int64) (*string, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pets/{petId}/downloadImage"


    // Path Params
    uri = strings.Replace(uri, "{"+"petId"+"}", url.PathEscape(client.ParameterToString(petId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/octet-stream", "application/pdf", "application/json"} 

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
    Find pet by ID
    Returns a single pet

    parameters:-
    -> petId (int64) (required) : ID of pet to return
    -> select_ (string) (optional) : Odata query for selecting projection
    -> expand_ (string) (optional) : Odata query for expanding list api response

    returns: (petstore.v1.defaultapi.PetApiResponse, error)
*/
func (api *PetApi) GetPetById(petId int64, select_ *string, expand_ *string) (*import1.PetApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pets/{petId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"petId"+"}", url.PathEscape(client.ParameterToString(petId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
	}
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.PetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get pets info from the store

    parameters:-
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - id 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - name 

    returns: (petstore.v1.defaultapi.CategoryListApiResponse, error)
*/
func (api *PetApi) GetPetCategories(filter_ *string, orderby_ *string) (*import1.CategoryListApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pet-categories"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
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
    unmarshalledResp := new(import1.CategoryListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get category from the store

    parameters:-
    -> categoryId (int64) (required)
    -> select_ (string) (optional) : Odata query for selecting projection

    returns: (petstore.v1.defaultapi.CategoryApiResponse, error)
*/
func (api *PetApi) GetPetCategory(categoryId int64, select_ *string) (*import1.CategoryApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pet-categories/{categoryId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"categoryId"+"}", url.PathEscape(client.ParameterToString(categoryId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.CategoryApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get pets info from the store

    parameters:-
    -> petIds ([]int64) (optional)
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - favorite_food - id - last_owner - name - originDetails/ancestors - originDetails/birthPlace/country - originDetails/birthPlace/region - originDetails/birthStatus - photoFiles - size - status - vaccinated 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - favorite_color - favorite_food - id - last_owner - name - originDetails/birthDate - size 
    -> expand_ (string) (optional) : Odata query for expanding list api response

    returns: (petstore.v1.defaultapi.PetListApiResponse, error)
*/
func (api *PetApi) GetPets(petIds *[]int64, filter_ *string, orderby_ *string, expand_ *string) (*import1.PetListApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pets"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if petIds != nil {
		t := *petIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("petIds", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("petIds", client.ParameterToString(t, "multi"))
		}
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

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.PetListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Endpoint to test Retry Mechanism
    Returns a 503 error

    parameters:-

    returns: (petstore.v1.defaultapi.PetApiResponse, error)
*/
func (api *PetApi) GetServerError() (*import1.PetApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pets/error"


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
    unmarshalledResp := new(import1.PetApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get tags info from the store

    parameters:-

    returns: (petstore.v1.defaultapi.TagListApiResponse, error)
*/
func (api *PetApi) GetTags() (*import1.TagListApiResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/tags"


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
    unmarshalledResp := new(import1.TagListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Uploads an image

    parameters:-
    -> body (string) (required) : Uploads an image
    -> petId (int64) (required) : ID of pet to update
    -> additionalMetadata (string) (optional) : Additional Metadata

    returns: (petstore.v1.defaultapi.UrlResponse, error)
*/
func (api *PetApi) UploadFile(body *string, petId int64, additionalMetadata *string) (*import1.UrlResponse, error) {
    uri := "/petstore/v1.0.a1/defaultapi/pets/{petId}/uploadImage"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"petId"+"}", url.PathEscape(client.ParameterToString(petId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{"application/octet-stream", "image/jpeg"}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	if additionalMetadata != nil {
		queryParams.Add("additionalMetadata", client.ParameterToString(*additionalMetadata, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UrlResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

