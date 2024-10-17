package cmd

import (
	"github.com/spf13/cobra"
)

func printTaskList() {
	ensureDBFile()
	printContent()
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the to do items",
	Run: func(cmd *cobra.Command, args []string) {
		printTaskList()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
