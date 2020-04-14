package registry

import "net/http"

func(registry *registryApi)TagsList(image string)([]string,error){
	request,err := registry.buildRegistryApiRequest("GET", registry.buildRegistryApiUrl(image, "tags", "list"), nil, nil)
	if err != nil {
		return nil, err
	}
	response,err := registry.client.Do(request)
	if err != nil {
		return nil,err
	}
	var tagListSpec rdImageTagsList
	err = registry.simpleParseRegistryApiResponse(response, http.StatusOK, &tagListSpec)
	if err != nil{
		return nil, err
	} else {
		return tagListSpec.Tags, nil
	}
}
