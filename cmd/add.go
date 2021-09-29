package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the to do list",
	Run: func(cmd *cobra.Command, args []string) {
		new_task := strings.Join(args, " ")
		fmt.Println("added new task: ", new_task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
