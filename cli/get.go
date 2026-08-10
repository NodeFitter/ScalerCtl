package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a list of VMs and their status",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.VMs()
	},
}
