//Api classes for tasks's golang SDK
package api

import (
    "github.com/nishantikagarg/sample-go-repo/tasks_go_sdk/v16/client"
	import1 "github.com/nishantikagarg/sample-go-repo/tasks_go_sdk/v16/models/prism/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
    "strings"
)

type TaskApi struct {
  ApiClient *client.ApiClient
}

func NewTaskApi() *TaskApi {
	a := &TaskApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    TaskCancel
    API to cancel an ongoing task.

    parameters:-
    -> taskExtId (string) (required) : Globally unique idntifier of a task.

    returns: (prism.v4.config.TaskCancelResponse, error)
*/
func (api *TaskApi) TaskCancel(taskExtId string) (*import1.TaskCancelResponse, error) {
    uri := "/api/prism/v4.0.a1/config/tasks/{taskExtId}/$actions/cancel"


    // Path Params
    uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(taskExtId, "")), -1)
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
    unmarshalledResp := new(import1.TaskCancelResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    TaskGet
    API to fetch an asyn operation called task.

    parameters:-
    -> taskExtId (string) (required) : Globally unique idntifier of a task.

    returns: (prism.v4.config.TaskGetResponse, error)
*/
func (api *TaskApi) TaskGet(taskExtId string) (*import1.TaskGetResponse, error) {
    uri := "/api/prism/v4.0.a1/config/tasks/{taskExtId}"


    // Path Params
    uri = strings.Replace(uri, "{"+"taskExtId"+"}", url.PathEscape(client.ParameterToString(taskExtId, "")), -1)
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
    unmarshalledResp := new(import1.TaskGetResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

