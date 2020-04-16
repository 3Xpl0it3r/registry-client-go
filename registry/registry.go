package registry

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RegistryApi struct {
	registryUrl string
	client *http.Client
	username, password string
}

func NewInsecureRegistryApi(registryUrl string)*RegistryApi {
	return &RegistryApi{registryUrl: registryUrl, client: newInsecureHttpClient()}
}

func NewTlsRegistryApi(registryUrl ,ca_pem string)*RegistryApi {
	return &RegistryApi{client: newTlsHttpClient(ca_pem), registryUrl: registryUrl}
}

func(registry *RegistryApi)SetBasicAuth(username,password string){
	registry.username = username
	registry.password = password
}

// buildRegistryApiUrl  build a special url by given segments in urls
func(registry *RegistryApi)buildRegistryApiUrl(segments ...string)string{
	url := registry.registryUrl + "/v2"

	if segments!= nil{
		if segments[0] == "v2"{
			segments = segments[1:]
		}
		for _,segment := range segments{
			url +="/" + segment
		}
	}
	return url
}

// buildRegistryApiRequest build a special request by given special method, url, header, and payload
// return a standard http.Request pointer
func(registry *RegistryApi)buildRegistryApiRequest(method,url string, headers map[string]string, body io.Reader)(*http.Request,error){
	request,err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Build Request Faile: %s\n",err.Error())
	}
	if headers != nil{
		for k,v := range headers{
			request.Header.Set(k, v)
		}
	}
	if registry.username != "" && registry.password != ""{
		request.SetBasicAuth(registry.username, registry.password)
	}
	return request,err
}


// disputed
func(registry *RegistryApi)parseRegistryApiResponse_v1(response *http.Response, expectHttpStatus int, object interface{})*rdRegistryCustomResponse {
	defer response.Body.Close()
	raw,_ := ioutil.ReadAll(response.Body)
	if response.StatusCode == expectHttpStatus{
		var err error
		if object != nil{
			err = rdHandler(raw, object)
		}
		if err != nil {
			return &rdRegistryCustomResponse{
				Code:    SUCCESS,
				Message: fmt.Sprintf(SUCCESS_MESSAGE, "Method:%s->Url:%s", response.Request.Method, response.Request.URL.String()),
				Body:    raw,
			}
		} else {
			return &rdRegistryCustomResponse{
				Code:    "",
				Message: "",
				Body:    nil,
			}
		}
	} 
	return &rdRegistryCustomResponse{
		Code:    UNKNOWN,
		Message: fmt.Sprintf(UNKONWN_MESSAGE, "Method: %s->Url:%s->Error:%s",response.Request.Method, response.Request.URL.String(),string(raw)),
		Body:    nil,
	}
	
}

// simpleParseRegistryApiResponse parse response object ,but ignore the registry special error types
// unmarshal response.body to object, return error or nil
func(registry *RegistryApi)simpleParseRegistryApiResponse(response *http.Response, expectHttpStatus int, object interface{})error{
	defer response.Body.Close()
	raw,_ := ioutil.ReadAll(response.Body)
	if response.StatusCode == expectHttpStatus{
		if object != nil{
			return rdHandler(raw, object)
		} else {
			return fmt.Errorf("Meethod:\t%s\tUrl:\t%s\tGot:%s\n", response.Request.Method, response.Request.URL.String(), string(raw))
		}
	} else if response.StatusCode == http.StatusNotFound{
		return fmt.Errorf("Method:\t%s\nUrl:\t%s\nStatusCode:\t%d\n", response.Request.Method, response.Request.URL.String(), response.StatusCode)
	} else {
		return fmt.Errorf("Method:%s\tUrl:%s\tStatusCode:\t%d\tRaw:%s\t", response.Request.Method, response.StatusCode,response.Request.URL.String())
	}
}


