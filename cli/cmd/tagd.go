package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewTagCommand()*cobra.Command{
	tagCmd := &cobra.Command{
		Use: "tag",
		Short: "tags",
	}
	tagCmd.AddCommand(newTagListCommand())
	return tagCmd
}



func newTagListCommand()*cobra.Command{
	list := &cobra.Command{
		Use: "list",
		Short: "list a image tag",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preCheckAndInitRegistryClient(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Argument name must be spcialfied\n")
			}
			tags,err := registryClient.TagsList(args[0])
			if err != nil{
				fmt.Fprintf(os.Stderr, "Error: %s\n",err.Error())
			} else {
				fmt.Fprintf(os.Stdout, "%s", strings.Join(tags, ","))
			}
			return nil
		},
	}
	return list
}