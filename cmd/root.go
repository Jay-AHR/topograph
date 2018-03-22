package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// This represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "topograph",
	Short: "A tool to enable interaction with upstream Terraform code as libraries.",
	Long: `A tool to enable interaction with upstream Terraform code as libraries.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// no flags except -h (--help) for this version