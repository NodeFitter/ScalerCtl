package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a list of VMs and their status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Called the GET command. TODO: implement")
	},
}
