package registry

import "net/http"


// ImageDelete Delete a image
func(registry *registryApi)ImageDelete(name,reference string)error{
	request,err := registry.buildRegistryApiRequest("DELETE", registry.buildRegistryApiUrl(name, "manifests", reference),
		map[string]string{"Accept":"application/vnd.docker.distribution.manifest.v2+json"}, nil)
	if err != nil{
		return err
	}
	if response,err := registry.client.Do(request);err != nil{
		return err
	} else {
		return registry.simpleParseRegistryApiResponse(response, http.StatusAccepted, nil)
	}
}
