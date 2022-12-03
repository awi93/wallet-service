package balance

import "github.com/spf13/cobra"

var HttpCmd = &cobra.Command{
	Use:   "balance-processor",
	Short: "Balance Processor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
