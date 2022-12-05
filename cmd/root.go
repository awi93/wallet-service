package cmd

import (
	"log"
	"os"

	"github.com/awi93/wallet-service/cmd/balance"
	"github.com/awi93/wallet-service/cmd/http"
	"github.com/awi93/wallet-service/cmd/threshold"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "wallet-service",
	Short: "Wallet Service",
}

func Execute() {
	RootCmd.AddCommand(http.HttpCmd)
	RootCmd.AddCommand(balance.BalanceProcessor)
	RootCmd.AddCommand(threshold.ThresholdProcessor)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Error nih ", err.Error())
		os.Exit(-1)
	}
}
