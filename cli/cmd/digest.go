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
		PreRun: func(cmd *cobra.Command, args []string)  {
			err :=  preCheckAndInitRegistryClient(cmd)
			if err != nil{
				fmt.Fprintf(os.Stderr, "%s",err)
			}
		},
		Run: func(cmd *cobra.Command, args []string)  {
			if len(args) != 1{
				fmt.Fprintf(os.Stderr, "image:reference must be supplied\n")
				return
			}
			imageInfo := strings.Split(args[0], ":")
			if len(imageInfo)!= 2 {
				fmt.Fprintf(os.Stderr, "Argument must be as imagename:reference\n")
				return
			}
			diget,err := registryClient.Digest(imageInfo[0], imageInfo[1])
			if err != nil{
				fmt.Fprintf(os.Stderr, "%s",err.Error())
				return
			}
			fmt.Fprintln(os.Stdout, diget)
			return
		},
	}
	digetdCmd.SetHelpFunc(digetdCmd.HelpFunc())
	return digetdCmd
}
