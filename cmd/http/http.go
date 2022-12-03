package http

import "github.com/spf13/cobra"

var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "Http Service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
