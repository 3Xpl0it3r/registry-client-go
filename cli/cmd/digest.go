package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func NewDigestCommand()*cobra.Command{
	digetdCmd := &cobra.Command{
		Use: "digest",
		Short: "show digest",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preCheckAndInitRegistryClient(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1{
				return fmt.Errorf("Arguments image:tag must be supplied")
			}
			imageInfo := strings.Split(args[0], ":")
			diget,err := registryClient.Digest(imageInfo[0], imageInfo[1])
			if err != nil{
				return err
			}
			fmt.Fprintln(os.Stdout, diget)
			return nil
		},
	}
	return digetdCmd
}
