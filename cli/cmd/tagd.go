package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewTagCommand()*cobra.Command{
	tagCmd := &cobra.Command{
		Use: "tag",
		Short: "tag",
	}
	tagCmd.AddCommand(newTagListCommand())
	tagCmd.SetHelpFunc(tagCmd.HelpFunc())
	return tagCmd
}



func newTagListCommand()*cobra.Command{
	list := &cobra.Command{
		Use: "list",
		Short: "list a image tag",
		PreRun: func(cmd *cobra.Command, args []string)  {
			err :=  preCheckAndInitRegistryClient(cmd)
			if err != nil{
				fmt.Fprintf(os.Stderr, "[!]Failed: %s",err.Error())
			}
			return
		},
		Run: func(cmd *cobra.Command, args []string)  {
			if len(args) != 1 {
				fmt.Fprintf(os.Stderr, "Image name must be supplied!\n")
				return
			}
			tags,err := registryClient.TagsList(args[0])
			if err != nil{
				fmt.Fprintf(os.Stderr, "Error: list tag error\n%s\n",err.Error())
				return
			}
			if len(tags) ==0{
				fmt.Fprintf(os.Stdout, "No tag in %s\n[*]%s May an empty repo/old image\n", args[0], args[0])
				return
			}
			for _,tag := range tags{
				fmt.Fprintf(os.Stdout,"\t%s\n", tag)
			}
		},
	}
	list.SetHelpFunc(list.HelpFunc())
	return list
}