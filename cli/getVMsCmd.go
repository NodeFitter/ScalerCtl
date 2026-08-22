package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getVMsCmd)
}

var getVMsCmd = &cobra.Command{
	Use:   "vms",
	Short: "Retrieve a list of VMs and their status",
	Long:  "Get the list of active VMs and information about their VM group and the amount of memory and CPU currently available",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.VMs()
	},
}
