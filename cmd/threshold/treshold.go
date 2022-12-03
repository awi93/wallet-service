package threshold

import "github.com/spf13/cobra"

var HttpCmd = &cobra.Command{
	Use:   "threshold-processor",
	Short: "Deposit Threshold Processor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
