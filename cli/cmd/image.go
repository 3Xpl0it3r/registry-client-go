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
		Short: "image operations",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootImg.AddCommand(newImageListCommand())
	rootImg.AddCommand(newImageDelCommand())
	return rootImg
}

// newImageListCommand list images
func newImageListCommand()*cobra.Command{
	listCmd := &cobra.Command{
		Use: "list",
		Short: "l",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preCheckAndInitRegistryClient(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			image,err := registryClient.RepositoriesList()
			if err != nil {
				return err
			}
			fmt.Fprintln(os.Stdout, strings.Join(image, ","))
			return nil
		},
	}
	return listCmd
}

func newImageDelCommand()*cobra.Command{

	delCmd := &cobra.Command{
		Use: "rm",
		Short: "delete a image",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preCheckAndInitRegistryClient(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			nameInfo := strings.Split(args[0], ":")
			reference,err := registryClient.Digest(nameInfo[0], nameInfo[1])
			if err != nil{
				return err
			}
			err = registryClient.ImageDelete(nameInfo[0], reference)
			if err != nil{
				return err
			}else {
				fmt.Fprintf(os.Stdout, "Delete %s:%s Succfully!", nameInfo[0],nameInfo[1])
			}
			return nil
		},
	}
	return delCmd
}
