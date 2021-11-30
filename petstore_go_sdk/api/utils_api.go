//Api classes for Petstore API project's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/client"
	import2 "github.com/nishantikagarg/sample-go-repo/petstore_go_sdk/models/petstore/v1/utils"
	"reflect"
	"encoding/json"
	"net/http"
    "net/url"
    "strings"
)

type UtilsApi struct {
  ApiClient *client.ApiClient
}

func NewUtilsApi() *UtilsApi {
	a := &UtilsApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    Returns a list of all namespaces in the API platform
    Returns a list of all namespaces in the API platform

    parameters:-

    returns: (petstore.v1.utils.MessageApiResponse, error)
*/
func (api *UtilsApi) GetAllNamespaces() (*import2.MessageApiResponse, error) {
    uri := "/petstore/v1.0.a1/utils/namespaces"


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
    unmarshalledResp := new(import2.MessageApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    GetNamespaceById
    Returns details for a single namespace identified by {id}

    parameters:-
    -> id (string) (required) : The uuid of the namespaces
    -> foo (string) (required)
    -> qux ([]string) (required)
    -> bar (string) (optional)

    returns: (petstore.v1.utils.MessageApiResponse, error)
*/
func (api *UtilsApi) GetNamespaceById(id string, foo *string, qux *[]string, bar *string) (*import2.MessageApiResponse, error) {
    uri := "/petstore/v1.0.a1/utils/namespaces/{id}"

    // verify the required parameter 'foo' is set
	if nil == foo {
		return nil, client.ReportError("foo is required and must be specified")
	}
    // verify the required parameter 'qux' is set
	if nil == qux {
		return nil, client.ReportError("qux is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"id"+"}", url.PathEscape(client.ParameterToString(id, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header 
    contentTypes := []string{}

    // to determine the Accept header 
	accepts := []string{"application/json"} 

    // Query Params
	queryParams.Add("foo", client.ParameterToString(*foo, ""))
	if bar != nil {
		queryParams.Add("bar", client.ParameterToString(*bar, ""))
	}
	{
		t := *qux
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				queryParams.Add("qux", client.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			queryParams.Add("qux", client.ParameterToString(t, "multi"))
		}
	}

    authNames := []string{}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import2.MessageApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Verifies complex payload

    parameters:-
    -> body (petstore.v1.utils.Bravo) (required) : Payload to test utils defs

    returns: (petstore.v1.utils.TestUtilsResponse, error)
*/
func (api *UtilsApi) TestUtils(body *import2.Bravo) (*import2.TestUtilsResponse, error) {
    uri := "/petstore/v1.0.a1/utils/testUtils"

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
    unmarshalledResp := new(import2.TestUtilsResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

