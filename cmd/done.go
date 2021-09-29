package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "To mark a task when it is done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println("done called with : ", ids)
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)

}
