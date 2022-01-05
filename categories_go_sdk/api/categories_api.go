//Api classes for categories's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/categories_go_sdk/v16/client"
	import1 "github.com/nishantikagarg/sample-go-repo/categories_go_sdk/v16/models/prism/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
    "strings"
)

type CategoriesApi struct {
  ApiClient *client.ApiClient
}

func NewCategoriesApi() *CategoriesApi {
	a := &CategoriesApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    Creates category 
    _Creates a new category using the provided request body._  The users need to specify the `name` of the category to be created and it's `parentExtId` (the extId of the parent category; if present) inside the request body. They can also provide a `description` and `userSpecifiedName` for the category inside the request body.  A sample request body would look like this: ``` { \"name\":\"sample-category\", \"parentExtId\": \"cafc8e9e-b595-46f8-8d43-f62746180a5b\", \"description\": \"This is a sample category\", \"userSpecifiedName\":\"sample-name\" } ```  If the user doesn't specify a `parentExtId` inside the request body, a root level category will be created.  ___Note:___ _Our service currently doesn't support providing `fqName`, `metadata` and `type` inside the request body. Adding them to the request body will not produce any error, but they won't be reflected in the created category_ 

    parameters:-
    -> body (prism.v4.config.Category) (required)

    returns: (prism.v4.config.CategoryApiResponse, error)
*/
func (api *CategoriesApi) CreateCategory(body *import1.Category) (*import1.CategoryApiResponse, error) {
    uri := "/api/prism/v4.0.a1/config/categories"

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
    unmarshalledResp := new(import1.CategoryApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Deletes category 
    _Deletes the category with the given external identifier._  The user has to specify inside the path parameter - a valid external identifier (`extId`) of the category to be deleted.  A sample call would look like this: ``` /prism/v4.0.a1/config/categories/cafc8e9e-b595-46f8-8d43-f62746180a5b ```  A category cannot be deleted in the following cases: - It has some child categories - It has some entities/policies asoociated with it - It is of type `SYSTEM` or `INTERNAL` 

    parameters:-
    -> extId (string) (required) : A globally unique identifier of an instance that is suitable for external consumption. 

    returns: (prism.v4.config.CategoryDeleteApiResponse, error)
*/
func (api *CategoriesApi) DeleteCategoryByExtId(extId string) (*import1.CategoryDeleteApiResponse, error) {
    uri := "/api/prism/v4.0.a1/config/categories/{extId}"


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
    unmarshalledResp := new(import1.CategoryDeleteApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get category list 
    _Get a paginated list of categories._  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported: - name - fqName - parentExtId - type - extId  Users can use 'and', 'or' i.e. logical operators. A sample call would look like this: ``` /prism/v4.0.a1/config/categories?$filter=(fqName eq 'cat1/val1' and extId eq 'bdecc360-b1f8-463f-b963-ecfa7380938f') ``` Following are all the supported logical operators and operations: - eq: Equals - ne: Not Equals - gt: Greater Than - ge: Greater Than or Equal - lt: Less Than - le: Less Than or Equal - and: And - or: Or - not: Not - in: In  The `$orderby` query parameter allows specifying attributes on which to sort the returned list of resource.  The `$expand` query parameter allows expanding the resource with another related resource, and returning this joined entity. This query can have $filter, $orderby and $expand queries nested inside it, to control the \"related resource\" in response.  $expand supported attributes are: - associations [$filter supported on resourceType and count attributes] - childCategories  Some more examples given below: 1. Filter by name (only parent categories): ``` /prism/v4.0.a1/categories?$filter=contains(name, 'something') and parentExtId eq null OR /prism/v4.0.a1/categories?$filter=name eq 'something' and parentExtId eq null ```  2. Filter by value (only child categories) ``` /prism/v4.0.a1/categories?$filter=name eq 'something' and parentExtId ne null ```  3. Filter by entity/policy type ``` /prism/v4.0.a1/categories?$expand=associations($filter=resourceType eq Schema.Enums.ResourceType'VM') ```  4. Filter by multiple entity/policy types ``` /prism/v4.0.a1/categories?$expand=associations($filter=(resourceType eq Schema.Enums.ResourceType'VM' or resourceType eq Schema.Enums.ResourceType'CLUSTER')) ```  5. Filter by name and by entity/policy type ``` /prism/v4.0.a1/categories?$expand=associations($filter=resourceType eq Schema.Enums.ResourceType'VM')&$filter=contains(name, 'something') and parentExtId eq null ```  6. Filter by name, by value and by entity/policy type ``` /prism/v4.0.a1/categories?$expand=associations($filter=resourceType eq Schema.Enums.ResourceType'VM')&$filter=name eq 'something' and parentExtId in ('extId1', 'extId2') ```  7. List categories with associations and child categories ``` /prism/v4.0.a1/categories?$expand=associations,childCategories ```  8. Filter by category type [USER,SYSTEM,INTERNAL] ``` /prism/v4.0.a1/categories?$filter=type eq Schema.Enums.CategoryType'SYSTEM' ```  9. Filter by multiple category types ``` /prism/v4.0.a1/categories?$filter=type eq Schema.Enums.CategoryType'SYSTEM' or type eq Schema.Enums.CategoryType'USER' ```  10. Order by fqName in ascending order ``` /prism/v4.0.a1/categories?$orderby=fqName asc ```  11. Order by name in descending order ``` /prism/v4.0.a1/categories?$orderby=name desc ```  If the user doesn't specify any search query parameters, a list of root level categories is returned. 

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - extId - fqName - name - parentExtId - type 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied on following fields: - fqName - name 
    -> expand_ (string) (optional) : Odata query for expanding list api response

    returns: (prism.v4.config.CategorySummaryListApiResponse, error)
*/
func (api *CategoriesApi) GetAllCategories(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string) (*import1.CategorySummaryListApiResponse, error) {
    uri := "/api/prism/v4.0.a1/config/categories"


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

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.CategorySummaryListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get category associations 
    _Get a list of associations for the given category and given resource type._  The user has to specify inside the path parameter - a valid external identifier (`extId`) of the category, whose associations are to be fetched. User can specify specific resourceType using $filter query param. Refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html) for further information. If not resourceType is specified, all the associated entities are returned grouped with resourceType.  A sample call without filter would look like this: ``` /prism/v4.0.a1/config/categories/cafc8e9e-b595-46f8-8d43-f62746180a5b/associations ```  Sample call with ResourceType filter would look like this: ``` /categories/<extId>/associations?$filter=resourceType in (Schema.Enums.ResourceType'VM', Schema.Enums.ResourceType'CLUSTER') ``` 

    parameters:-
    -> extId (string) (required) : A globally unique identifier of an instance that is suitable for external consumption. 
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on following fields: - categoryId - resourceType 

    returns: (prism.v4.config.AssociationDetailApiResponse, error)
*/
func (api *CategoriesApi) GetAssociationsForCategory(extId string, page_ *int, limit_ *int, filter_ *string) (*import1.AssociationDetailApiResponse, error) {
    uri := "/api/prism/v4.0.a1/config/categories/{extId}/associations"


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
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {
		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.AssociationDetailApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get category by extId 
    _Get a category with the given external identifier._  The user has to specify inside the path parameter - a valid external identifier (`extId`) of the category to be fetched.  A sample call would look like this: ``` /prism/v4.0.a1/config/categories/cafc8e9e-b595-46f8-8d43-f62746180a5b?$view=detail ```  This API supports $expand on associations and childCategories as well. Sample query would look like following: ``` /prism/v4.0.a1/categories/<ext_id>?$expand=association,childCategories ``` 

    parameters:-
    -> extId (string) (required) : A globally unique identifier of an instance that is suitable for external consumption. 
    -> expand_ (string) (optional) : Odata query for expanding list api response

    returns: (prism.v4.config.CategoryGetResponses, error)
*/
func (api *CategoriesApi) GetCategoryByExtId(extId string, expand_ *string) (*import1.CategoryGetResponses, error) {
    uri := "/api/prism/v4.0.a1/config/categories/{extId}"


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
	if expand_ != nil {
		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.CategoryGetResponses)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update category 
    _Updates the category with the given external identifier using the provided request body._  The user has to specify inside the path parameter - a valid external identifier (`extId`) of the category to be updated. They also need to provide a request body for performing the update. They need to specify the `name` of the category to be created and it's `parentExtId` (the extId of the parent category; if present) inside the request body. They can provide an updated value of `description` and `userSpecifiedName` for the category inside the request body.  A sample request body would look like this: ``` { \"name\":\"sample-category\", \"parentExtId\": \"cafc8e9e-b595-46f8-8d43-f62746180a5b\", \"description\": \"This is the updated description\", \"userSpecifiedName\":\"sample-name-updated\" } ```  The fields `name` and `parentExtId` are immutable, hence the user cannot modify them inside this update call. Also, since this is a `PUT` call, the user needs to specify these fields inside the request body (otherwise a 400 error code will be returned). If the category to be updated is a root category, do not specify the `parentExtId` field.  ___Note:___ _Our service currently doesn't support providing `fqName`, `metadata` and `type` inside the request body. Adding them to the request body will not produce any error, but they won't be reflected in the created category._ 

    parameters:-
    -> body (prism.v4.config.Category) (required)
    -> extId (string) (required) : A globally unique identifier of an instance that is suitable for external consumption. 

    returns: (prism.v4.config.CategoryPutResponses, error)
*/
func (api *CategoriesApi) UpdateCategory(body *import1.Category, extId string) (*import1.CategoryPutResponses, error) {
    uri := "/api/prism/v4.0.a1/config/categories/{extId}"

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
    unmarshalledResp := new(import1.CategoryPutResponses)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

