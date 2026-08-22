package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getMemCmd)
}

var getMemCmd = &cobra.Command{
	Use:   "ram",
	Short: "Get the memory threshold for VMs",
	Long:  "Get the current maximum memory usage after which new VMs need to be scheduled",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.GetMemThreshold()
	},
}
