package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getCpuCmd)
}

var getCpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Get the CPU threshold for VMs",
	Long:  "Get the current maximum CPU usage after which new VMs need to be scheduled",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.GetCPUThreshold()
	},
}
