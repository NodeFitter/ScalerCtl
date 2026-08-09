package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	setCmd.AddCommand(setRamCmd)
}

var setRamCmd = &cobra.Command{
	Use:   "ram <Mb>",
	Short: "Set the RAM threshold for VMs",
	Long:  "Set the maximum memory usage after which new VMs need to be scheduled",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ram, err := strconv.Atoi(args[0])
		if err != nil || ram <= 0 {
			return fmt.Errorf("RAM threshold must be a positive integer")
		}

		fmt.Printf("Ran the SET RAM command with value %d. TODO: implement", ram)

		return nil
	},
}
