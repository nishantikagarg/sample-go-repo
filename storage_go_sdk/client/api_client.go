//The api client for storage's golang SDK
package client

import (
	"bytes"
    "context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
    "github.com/google/uuid"
    "github.com/hashicorp/go-retryablehttp"
	"io"
    "io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
    uriCheck  = regexp.MustCompile(`/(?P<namespace>[-\w]+)/v\d+\.\d+(\.[a|b]\d+)?/(?P<suffix>.*)`)
    basic = "basic"
    eTag = "ETag"
    ifMatch = "If-Match"
    retryStatusList = []int{408, 503, 504}
)

/**
    Generic API client for Swagger client library builds.
    Swagger generic API client. This client handles the client-
    server communication, and is invariant across implementations. Specifics of
    the methods and models for each application are generated from the Swagger
    templates.

    Parameters :-
      BasePath (string) : base path or host for all http request made by this client
      DefaultHeaders (map[string]string) : List of default headers as part of each request
      Debug (bool) : flag to enable debug logging (default : empty)
      UserAgent (string) : User-Agent header's value (default: Swagger-Generator/1.0.0/go)
      BackOffPeriod (time.Duration) : Backoff duration between successive retry attempts (default: 3 sec)
      MaxRetryAttempts (int) : Maximum number of retry attempts to be made at a time (default 5) 
*/
type ApiClient struct {
	BasePath      		    string            		`json:"host,omitempty"`
	DefaultHeaders 		    map[string]string 		`json:"defaultHeader,omitempty"`
	Debug         		    bool              		`json:"debug,omitempty"`
	UserAgent     		    string            		`json:"userAgent,omitempty"`
    Cookie					string					`json:"cookie,omitempty"`
    BackOffPeriod 		    time.Duration           `json:"backOffPeriod,omitempty"`
	MaxRetryAttempts 	    int                     `json:"maxRetryAttempts,omitempty"`
	httpClient    		    *http.Client
    authentication		    map[string]interface{}
    refreshCookie           bool
    basicAuth               string
}

/**
    Returns a newly generated ApiClient instance populated with default values
*/
func NewApiClient() *ApiClient {
	a := &ApiClient{
		DefaultHeaders: 	        make(map[string]string),
		UserAgent:     		        "Swagger-Generator/1.0.0/go",
		Debug:         		        false,
		BackOffPeriod: 		        3 * time.Second,
		MaxRetryAttempts: 	        5,
		authentication:		        make(map[string]interface{}),
        refreshCookie:              true,
	}
    a.setupClient()
	return a
}

// Makes the HTTP request with given options and returns response body as byte array.
func (a *ApiClient) CallApi(uri *string, httpMethod string, body interface{},
	queryParams url.Values, headerParams map[string]string, formParams url.Values,
	accepts []string, contentType []string, authNames []string) ([]byte, error) {
	path := a.BasePath + *uri

    if headerParams["Authorization"] != "" {
        a.basicAuth = headerParams["Authorization"]
    }

    if a.DefaultHeaders["Authorization"] != "" {
        a.basicAuth = a.DefaultHeaders["Authorization"]
    }

	// set Content-Type header
	httpContentType := a.selectHeaderContentType(contentType)
	if httpContentType != "" {
		headerParams["Content-Type"] = httpContentType
	}

    // Set Cookie header
    if a.Cookie != "" {
        // Remove basic auth header if cookie is available
        delete(a.DefaultHeaders, "Authorization")
        delete(headerParams, "Authorization")
        headerParams["Cookie"] = a.Cookie
    }

	// set Accept header
	httpHeaderAccept := a.selectHeaderAccept(accepts)
	if httpHeaderAccept != "" {
		headerParams["Accept"] = httpHeaderAccept
	}

    if _, ok := a.DefaultHeaders["NTNX-Request-Id"]; !ok {
		a.DefaultHeaders["NTNX-Request-Id"] = uuid.New().String()
	}

    if body != nil {
		addEtagReferenceToHeader(body, headerParams)
	}

	request, err := a.prepareRequest(path, httpMethod, body, headerParams, queryParams, formParams, authNames)
	if err != nil {
		return nil, err
	}

	if a.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	response, err := a.httpClient.Do(request)

    // Retry one more time without the cookie but with basic auth header
    if response != nil && response.StatusCode == 401 {
        a.refreshCookie = true
        request.Header["Authorization"] = []string{a.basicAuth}
        delete(request.Header, "Cookie")
        response, err = a.httpClient.Do(request)
    }

	if err != nil {
		return nil, err
	}

	if a.Debug {
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	if nil == response {
		return nil, ReportError("response is nil!")
	}

    a.updateCookies(response)

    if response.StatusCode == 204 {
        return nil, nil
    }

    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
    	return nil, err
	}
    response.Body.Close()
    response.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))

    if !(200 <= response.StatusCode && response.StatusCode <= 209) {
		return nil, GenericOpenAPIError {
			Body:   responseBody,
			Status: response.Status,
		}
	} else {
		responseBody := addEtagReferenceToResponse(response.Header, responseBody)
		return responseBody, nil
	}
}

// Get all authentications (key: authentication name, value: authentication)
func (a *ApiClient) GetAuthentications() map[string]interface{} {
    return a.authentication
}

// Get authentication for the given auth name (eg : basic, oauth, bearer, apiKey)
func (a *ApiClient) GetAuthentication(authName string) interface{} {
    return a.authentication[authName]
}

// Helper method to set username for the first HTTP basic authentication.
func (a *ApiClient) SetUserName(username string) error {
	if  auth, ok := a.authentication[basic].(BasicAuth); ok {
		auth.UserName = username
		return nil
	} else {
		return ReportError("no HTTP basic authentication configured!")
	}
}

// Helper method to set password for the first HTTP basic authentication
func (a *ApiClient) SetPassword(password string) error {
	if  auth, ok := a.authentication[basic].(BasicAuth); ok {
		auth.Password = password
		return nil
	} else {
		return ReportError("no HTTP basic authentication configured!")
	}
}

// Helper method to set API key value for the first API key authentication
func (a *ApiClient) SetApiKey(apiKey string) error {
	if authMap, ok := a.authentication["apiKey"].(map[string]interface{}); ok {
		if auth, ok := authMap["apiKey"].(APIKey); ok {
			auth.Key = apiKey
			return nil
		}
	}
	return ReportError("no API key authentication configured!")
}

// Helper method to set API key prefix for the first API key authentication
func (a *ApiClient) SetApiKeyPrefix(apiKeyPrefix string) error {
	if authMap, ok := a.authentication["apiKey"].(map[string]interface{}); ok {
		if auth, ok := authMap["apiKey"].(APIKey); ok {
			auth.Prefix = apiKeyPrefix
			return nil
		}
	}
	return ReportError("no API key authentication configured!")
}

// Helper method to set access token for the first OAuth2 authentication.
func (a *ApiClient) SetAccessToken(accessToken string) error {
	for _, value := range a.authentication {
		if  auth, ok := value.(OAuth); ok {
			auth.AccessToken = accessToken
			return nil
		}
	}
	return ReportError("no OAuth2 authentication configured!")
}

/**
    Helper method to set maximum retry attempts
    !! After the initial instantiation of ApiClient,
    maximum retry attempts must be modified only via this method
*/
func (a *ApiClient) SetMaxRetryAttempts(maxRetryAttempts int) {
	a.MaxRetryAttempts = maxRetryAttempts
	a.setupClient()
}

/**
    Helper method to set retry back off period
    !! After the initial instantiation of ApiClient,
    back off period must be modified only via this method
*/
func (a *ApiClient) SetBackOffPeriodInMilliSeconds(ms int) {
	a.BackOffPeriod = time.Duration(ms) * time.Millisecond
	a.setupClient()
}

func (a *ApiClient) setupClient() {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = a.MaxRetryAttempts
	retryClient.RetryWaitMax = a.BackOffPeriod
	retryClient.CheckRetry = retryPolicy
	if !a.Debug {
		retryClient.Logger = nil
	}
	a.httpClient = retryClient.StandardClient()
}

// SelectHeaderContentType select a content type from the available list.
func (a *ApiClient) selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func (a *ApiClient) selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// prepareRequest build the request
func (a *ApiClient) prepareRequest(
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
    authNames []string) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}
		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", a.UserAgent)

    // Add the cookie to the request.
    if len(a.Cookie) > 0 {
	  localVarRequest.Header.Add("Cookie", a.Cookie)
    }

    // Authentication
    for _, authName := range authNames {
        switch authName {
            case "base":
                // Basic HTTP Authentication
                if auth, ok := a.authentication["basic"].(BasicAuth); ok {
                    localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
                }
            case "apiKey":
                // ApiKey Authentication
                if auth, ok := a.authentication["apiKey"].(map[string]interface{}); ok {
                    var key string
                    if apiKey,ok := auth["apiKey"].(APIKey); ok && apiKey.Prefix != "" {
                        key = apiKey.Prefix + " " + apiKey.Key
                    } else {
                        key = apiKey.Key
                    }
                    if auth["in"] == "header" {
                        headerParams["apiKey"] = key
                    }
                    if auth["in"] == "query" {
                        queryParams.Add("apiKey", key)
                    }
                }
            case "oauth", "bearer":
                // OAuth2/AccessToken authentication
                if auth, ok := a.authentication["bearer"].(OAuth); ok {
                    localVarRequest.Header.Add("Authorization", "Bearer "+auth.AccessToken)
                }
            default:
                return nil, ReportError("unknown authentication type: %s", authName)
        }
    }

	for header, value := range a.DefaultHeaders {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}


// RetryPolicy provides a callback for Client.CheckRetry, specifies retry on
// error codes mentioned in RetryStatusList
func retryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
    if err != nil {
		return false, err
	}
	for _, status := range retryStatusList {
		if resp.StatusCode == status {
			return true, nil
		}
	}
	return false, nil
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// addEtagReferenceToHeader method is used to read ETag and add it to If-Match header
func addEtagReferenceToHeader(body interface{}, headerParams map[string]string) {
	if reflect.ValueOf(body).Elem().Kind() == reflect.Struct {
		if reserved := reflect.ValueOf(body).Elem().FieldByName("Reserved_"); reserved.IsValid() {
			reservedMap := reserved.Interface().(map[string]interface{})
			if etag, etagOk := reservedMap[eTag].(string); etagOk {
				headerParams[ifMatch] = etag
			}
		}
	}
}

// addEtagReferenceToResponse method is used to read ETag and add it to response
func addEtagReferenceToResponse(headers http.Header, body []byte) []byte {
	if etag := headers.Get(eTag); etag != "" {
		responseMap := map[string]interface{}{}
		json.Unmarshal(body, &responseMap)
		if r, ok := responseMap["$reserved"].(map[string]interface{}); ok {
			r[eTag] = etag
			if d, ok := responseMap["data"].(map[string]interface{}); ok {
				if r2, ok := d["$reserved"].(map[string]interface{}); ok {
					r2[eTag] = etag
					m, _ := json.Marshal(responseMap)
					return m
				}
			} else if dList, ok := responseMap["data"].([]interface{}); ok {
				for _, d := range dList {
					if d, ok := d.(map[string]interface{}); ok {
						if r3, ok := d["$reserved"].(map[string]interface{}); ok {
							r3[eTag] = etag
						}
					}
				}
				m, _ := json.Marshal(responseMap)
				return m
			}
		}
	}
	return body
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if nil == bodyBuf {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// Set Cookie information to reuse in subsequent requests for a valid response
func (a *ApiClient) updateCookies(response *http.Response) {
    if a.refreshCookie {
        cookiesList := response.Header["Set-Cookie"]
        if len(cookiesList) > 0 {
            cookie := ""
            for _, value := range cookiesList {
                finalCookie := strings.SplitN(value, ";", 2)[0]
                if strings.Contains(finalCookie, "=") {
                    finalCookie = strings.TrimSpace(finalCookie)
                } else {
                    finalCookie = ""
                }

                if finalCookie != "" {
                    cookie += finalCookie + ";"
                }
            }

            // Remove trailing ";"
            if cookie != "" {
                cookie = strings.TrimSuffix(cookie, ";")
            }

            a.Cookie = cookie
            a.refreshCookie = false
        }
    }
}

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// OAuth provides OAuth authentication
type OAuth struct {
	AccessToken    string
}

// GenericOpenAPIError Provides access to the body (error), status and model on returned errors.
type GenericOpenAPIError struct {
	Body  	[]byte
	Model 	interface{}
    Status  string
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return string(e.Body)
}

// Error returns deserialized response body if compatible with GenericOpenAPIError.Model
func (e GenericOpenAPIError) DeserializedModel() interface{} {
	err := json.Unmarshal(e.Body, e.Model)
	if err != nil {
		return nil
	}
	return e.Model
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func ParameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func ParameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// Prevent trying to import "fmt"
func ReportError(format string, b ...interface{}) error {
	return fmt.Errorf(format, b...)
}