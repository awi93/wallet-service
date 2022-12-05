package balance

import (
	"context"

	bootstrap "github.com/awi93/wallet-service/config"
	balanceModel "github.com/awi93/wallet-service/src/models/balance"
	depositModel "github.com/awi93/wallet-service/src/models/deposit"
	"github.com/awi93/wallet-service/src/processors/balance"
	"github.com/lovoo/goka"
	"github.com/spf13/cobra"
)

var BalanceProcessor = &cobra.Command{
	Use:   "balance-processor",
	Short: "Balance Processor",
	RunE: func(cmd *cobra.Command, args []string) error {
		bootstrap.ReadConfig("config.yml")
		balance.InitDependency(bootstrap.GetConfig())

		return (func() error {
			config := bootstrap.GetConfig()
			var topic = goka.Stream(config.GetString("deposit.topic"))
			g := goka.DefineGroup(balanceModel.Group, goka.Input(topic, new(depositModel.DepositCodec), balance.BalanceService.ProcessCallback), goka.Persist(new(balanceModel.BalanceCodec)))
			p, err := goka.NewProcessor(config.GetStringSlice("broker.hosts"), g)
			if err != nil {
				return err
			}
			return p.Run(context.Background())
		})()
	},
}
