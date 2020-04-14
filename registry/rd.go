package registry

import (
	"encoding/json"
)

type rdRegistryCustomResponse struct {
	Code string
	Message string
	Body []byte
}


// rdRegistryStandardError is represent registry server errors
type rdRegistryStandardError struct {
	Errors []struct{
		Code string `json:"code"`
		Message string `json:"message"`
		Detail string `json:"detail"`
	}`json:"errors"`
}

// rdCatalogReturn description response for  request catalog
type rdRepositories struct {
	Repositories []string `json:"repositories"`
}


// rdImageTagsList description response for request list imagestags
type rdImageTagsList struct {
	Name string `json:"name"`
	Tags []string `json:"tags"`
}


// rdHandler trans http.response.body into concrete resource
func rdHandler(in []byte, out interface{})error{

	if err := json.Unmarshal(in, out);err != nil{
		return err
	}
	return nil
}