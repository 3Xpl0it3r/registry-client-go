package registry

import (
	"net/http"
)

// repositoriesList list all repositories in registry server
func(registry *registryApi)RepositoriesList()([]string,error){
	req,err := registry.buildRegistryApiRequest("GET", registry.buildRegistryApiUrl("_catalog"), nil, nil)
	if err != nil {
		return nil, err
	}
	response,err := registry.client.Do(req)

	if err != nil {
		return nil,err
	}
	var repoList rdRepositories
	err = registry.simpleParseRegistryApiResponse(response, http.StatusOK, &repoList)
	if err == nil{
		return repoList.Repositories, nil
	} else {
		return nil, err
	}
}
