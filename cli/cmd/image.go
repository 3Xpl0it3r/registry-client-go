package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewImageCommand()*cobra.Command{
	rootImg := &cobra.Command{
		Use: "image",
	}
	rootImg.AddCommand(newImageListCommand())
	rootImg.AddCommand(newImageDelCommand())
	rootImg.SetHelpFunc(rootImg.HelpFunc())
	return rootImg
}

// newImageListCommand list images
func newImageListCommand()*cobra.Command{
	listCmd := &cobra.Command{
		Use: "list",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preCheckAndInitRegistryClient(cmd)
		},
		Run: func(cmd *cobra.Command, args []string)  {
			image,err := registryClient.RepositoriesList()
			if err != nil {
				fmt.Fprintf(os.Stdout, "[!]Failed: %s\n",err.Error())
			}
			for _,img := range image{
				fmt.Println(img)
			}
			return
		},
	}
	listCmd.SetHelpFunc(listCmd.HelpFunc())
	return listCmd
}

func newImageDelCommand()*cobra.Command{
	delCmd := &cobra.Command{
		Use: "rm",
		PreRun: func(cmd *cobra.Command, args []string)  {
			err := preCheckAndInitRegistryClient(cmd)
			if err != nil{
				fmt.Fprintf(os.Stderr, "[!]Failed: %s\n",err.Error())
				return
			}
		},
		Run: func(cmd *cobra.Command, args []string)  {
			nameInfo := strings.Split(args[0], ":")
			reference,err := registryClient.Digest(nameInfo[0], nameInfo[1])
			if err != nil{
				return
			}
			err = registryClient.ImageDelete(nameInfo[0], reference)
			if err != nil{
				return
			}else {
				fmt.Fprintf(os.Stdout, "Delete %s:%s Succfully!", nameInfo[0],nameInfo[1])
			}
			return
		},
	}
	delCmd.SetHelpFunc(delCmd.HelpFunc())
	return delCmd
}
