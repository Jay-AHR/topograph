package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/hashicorp/terraform/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Terraform via topograph",
	Long:  `All software has versions. This is Terraform's version via topograph.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The version of Terraform associated with vendoring for this version of topograph is:")
		fmt.Println(version.String())
	},
}