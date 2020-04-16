package main

import (
	"github.com/spf13/cobra"
	"io"
	"l0calh0st.cn/docker-registry/cli/cmd"

	"os"
)

var (
	rootCmd *cobra.Command
)
var (
	tlsClient bool = false
)
var (
	insecure bool = true
	url string = ""
	ca string = ""
	auth string = ""
)

func main() {
	rootCmd = NewCommand(os.Stdin, os.Stdout, os.Stderr)
	initFlagOptions()
	rootCmd.Execute()

}


func NewCommand(in io.Reader, out io.Writer, err io.Writer)*cobra.Command{
	rootCmd = &cobra.Command{
		Use: "registry",
	}
	// add image sub_command
	rootCmd.AddCommand(cmd.NewImageCommand())
	// show a image digest
	rootCmd.AddCommand(cmd.NewDigestCommand())
	//
	rootCmd.AddCommand(cmd.NewTagCommand())
	rootCmd.SetHelpFunc(rootCmd.HelpFunc())

	return rootCmd
}



func initFlagOptions()error{
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", true, "--insecure")
	rootCmd.PersistentFlags().Set("insecure", "true")

	rootCmd.PersistentFlags().StringVarP(&url, "url", "u","http://localhost:5001", "--url or -u")
	rootCmd.PersistentFlags().Set("url", "https://localh0st:5001")
	rootCmd.MarkPersistentFlagRequired("url")

	rootCmd.PersistentFlags().StringVar(&ca, "ca", "", "--ca")
	rootCmd.PersistentFlags().StringVarP(&auth, "auth", "k", "root:root", "--auth  or -k")

	rootCmd.PersistentFlags().BoolVar(&tlsClient, "tlsclient", false, "")
	rootCmd.PersistentFlags().MarkHidden("tlsclient")


	if insecure || ca == ""{
		rootCmd.PersistentFlags().Set("tlsclient", "false")
	}  else {
		rootCmd.PersistentFlags().Set("tlsclient", "true")
	}

	return nil
}

