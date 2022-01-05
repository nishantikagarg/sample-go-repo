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

type AncConfigApi struct {
  ApiClient *client.ApiClient
}

func NewAncConfigApi() *AncConfigApi {
	a := &AncConfigApi{
		ApiClient: client.NewApiClient(),
	}
	return a
}


/**
    GetAncConfig
    Schema for ANC.

    parameters:-

    returns: (networking.v4.config.AncConfigApiResponse, error)
*/
func (api *AncConfigApi) GetAncConfig() (*import1.AncConfigApiResponse, error) {
    uri := "/api/networking/v4.0.a1/config/anc-config"


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
    unmarshalledResp := new(import1.AncConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

