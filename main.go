package main

import (
	"fmt"
	"l0calh0st.cn/docker-registry/registry"
	"log"
)

func main() {
	client := registry.NewInsecureRegistryApi("http://10.23.6.90:5000")
	client.SetBasicAuth("username", "password")
	// list all repo
	if repo,err := client.RepositoriesList();err != nil{
		fmt.Printf("List repo err: %s", err.Error())
	} else {
		fmt.Printf("Repo: %s\n", repo)
	}
	// show tags
	if tags,err := client.TagsList("alpine");err != nil{
		log.Panicln(err)
	} else {
		fmt.Printf("All tags is : %v\n", tags)
	}
	// show digest
	if digest,err := client.Digest("alpine", "v2");err != nil{
		log.Printf("%s", err.Error())
	} else {
		fmt.Printf("%s\n", digest)
	}

	// delete digest
	//if err := client.ImageDelete("alpine", "sha256:cb8a924afdf0229ef7515d9e5b3024e23b3eb03ddbba287f4a19c6ac90b8d221");err !=nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Println("DELETE OK")
	//}


}
