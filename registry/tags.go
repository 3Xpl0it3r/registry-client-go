package registry

import (
	"fmt"
	"net/http"
)

func(registry *RegistryApi)TagsList(image string)([]string,error){
	request,err := registry.buildRegistryApiRequest("GET", registry.buildRegistryApiUrl(image, "tags", "list"), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Build Request Failed: %s\n",err.Error())
	}
	response,err := registry.client.Do(request)
	if err != nil {
		return nil,fmt.Errorf("Request Execute Failed: %s\n",err.Error())
	}
	var tagListSpec rdImageTagsList
	err = registry.simpleParseRegistryApiResponse(response, http.StatusOK, &tagListSpec)
	if err != nil{
		return nil, err
	} else {
		return tagListSpec.Tags, nil
	}
}
