package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource threshold for VMs",
	Long:  "Get the threshold for a specified resource, which will become the point where new VMs will be scheduled",
}
