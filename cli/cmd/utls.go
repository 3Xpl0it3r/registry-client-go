package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"l0calh0st.cn/docker-registry/registry"
	"strings"
)

var (
	registryClient *registry.RegistryApi
)

// newRegistryApiClient create a new registry client according cli flags
func newRegistryApiClient(url string, tls bool, ca, auth string)(*registry.RegistryApi,error){
	var client *registry.RegistryApi
	if tls {
		client = registry.NewTlsRegistryApi(url, ca)
	} else {
		client = registry.NewInsecureRegistryApi(url)
	}
	if auth != ""{
		username,password,err := basicAuth(auth)
		if err != nil{
			return nil, err
		}
		client.SetBasicAuth(username, password)
	}
	return client, nil
}

// basicAuth split auth string to username,password
func basicAuth(auth string)(username,password string,err error){
	if auth == ""{
		return "","", fmt.Errorf("auth must be supplied: <username:password>")
	}
	up := strings.Split(auth, ":")
	if len(up) != 2{
		return "", "", fmt.Errorf("auth invalid format: <username:password>")
	}
	return up[0], up[1], nil
}

// parseCommandFlags parse cobra persistent flags,
// extract arguments from cobra persistent flags,
// type string as url,ca,auth ,and type bool as tls are expected
func parseCommandFlags(cmd *cobra.Command)(url,ca, auth string, tls bool, err error){
	url,err = cmd.Flags().GetString("url")
	if err != nil {
		return "", "", "", false, err
	}
	ca,err = cmd.Flags().GetString("ca")
	if err != nil{
		return "", "", "", false, nil
	}
	auth,err = cmd.Flags().GetString("auth")
	if err != nil{
		return "", "", "", false, err
	}
	tls,err = cmd.Flags().GetBool("tlsclient")
	return

}


func preCheckAndInitRegistryClient(cmd *cobra.Command)error{
	url,ca,auth, tlsclient,err := parseCommandFlags(cmd)
	if err != nil{
		return err
	}
	registryClient,err = newRegistryApiClient(url,tlsclient, ca, auth)
	if err != nil{
		return err
	}
	return nil
}