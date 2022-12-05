package balance

import (
	"context"
	"log"

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
		log.Println("Reading Configuration File")
		bootstrap.ReadConfig("config.yml")
		log.Println("Configuration File Read")
		log.Println("Initializing Dependency")
		balance.InitDependency(bootstrap.GetConfig())
		log.Println("Dependency Initialized")

		return (func() error {
			config := bootstrap.GetConfig()
			var topic = goka.Stream(config.GetString("deposit.topic"))
			g := goka.DefineGroup(balanceModel.Group, goka.Input(topic, new(depositModel.DepositCodec), balance.BalanceService.ProcessCallback), goka.Persist(new(balanceModel.BalanceCodec)))
			p, err := goka.NewProcessor(config.GetStringSlice("broker.hosts"), g)
			if err != nil {
				return err
			}
			log.Printf("Start Streaming Kafka Event at Topic : %s\n", topic)
			return p.Run(context.Background())
		})()
	},
}
