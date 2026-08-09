package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a resource threshold for VMs",
	Long:  "Set the threshold for a specified resource, which will become the point where new VMs will be scheduled",
}
