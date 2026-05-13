package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "UWS",
	Short: "Ultimate work space CLI",
	Long:  "A CLI tool to manage your work space and projects",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
