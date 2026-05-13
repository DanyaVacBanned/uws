package cmd

import (
	"fmt"
	"ultimateWorkSpace/internal/scripts"

	"github.com/spf13/cobra"
)

var scriptsRootCmd = &cobra.Command{
	Use:   "scripts",
	Short: "Manage your scripts",
	Long:  "A set of commands to manage your scripts, including creating, listing, and deleting scripts",
}

var scriptsCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new script",
	Long:  "Create a new script with the given name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a name for the script")
			return
		}
		name := args[0]
		fmt.Printf("Creating script: %s\n", name)
		scripts.CreateScript(name)
	},
}

func init() {
	rootCmd.AddCommand(scriptsRootCmd)
	scriptsRootCmd.AddCommand(scriptsCreateCmd)
}
