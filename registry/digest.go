package registry

import "fmt"

func(registry *registryApi)Digest(name,reference string)(string,error){
	request,err := registry.buildRegistryApiRequest("HEAD", registry.buildRegistryApiUrl(name, "manifests", reference),
		map[string]string{"Accept":"application/vnd.docker.distribution.manifest.v2+json"}, nil)
	if err != nil{
		return "", err
	}
	response,err := registry.client.Do(request)
	fmt.Printf("%s  %s", request.URL.String(),response)
	if err != nil{
		return "",nil
	}
	digist := response.Header.Get("Docker-Content-Digest")
	if digist == ""{
		return "", fmt.Errorf("Cannot get digest")
	}
	return digist, nil
}
