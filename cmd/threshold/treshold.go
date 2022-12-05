package threshold

import (
	"context"
	"log"

	config "github.com/awi93/wallet-service/config"
	depositModel "github.com/awi93/wallet-service/src/models/deposit"
	thresholdModel "github.com/awi93/wallet-service/src/models/threshold"
	"github.com/awi93/wallet-service/src/processors/threshold"
	"github.com/lovoo/goka"
	"github.com/spf13/cobra"
)

var ThresholdProcessor = &cobra.Command{
	Use:   "threshold-processor",
	Short: "Deposit Threshold Processor",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Reading Configuration File")
		config.ReadConfig("config.yml")
		log.Println("Configuration File Read")
		log.Println("Initializing Dependency")
		threshold.InitDependency(config.GetConfig())
		log.Println("Dependency Initialized")

		return (func() error {
			config := config.GetConfig()
			var topic = goka.Stream(config.GetString("deposit.topic"))
			g := goka.DefineGroup(
				thresholdModel.Group,
				goka.Input(topic, new(depositModel.DepositCodec), threshold.ThresholdService.ProcessCallback),
				goka.Persist(new(thresholdModel.ThresholdCodec)),
			)
			p, err := goka.NewProcessor(config.GetStringSlice("broker.hosts"), g)
			if err != nil {
				return err
			}
			log.Printf("Start Streaming Kafka Event at Topic : %s\n", topic)
			return p.Run(context.Background())
		})()
	},
}
