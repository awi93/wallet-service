package threshold

import (
	"context"

	bootstrap "github.com/awi93/wallet-service/config"
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
		bootstrap.ReadConfig("config.yml")
		threshold.InitDependency(bootstrap.GetConfig())

		return (func() error {
			config := bootstrap.GetConfig()
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
			return p.Run(context.Background())
		})()
	},
}
