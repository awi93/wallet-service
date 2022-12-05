package threshold

import (
	"context"
	"log"

	"github.com/awi93/wallet-service/src/models/threshold"
	thresholdSvc "github.com/awi93/wallet-service/src/services/threshold"
	"github.com/awi93/wallet-service/src/utils/topicutil"
	"github.com/lovoo/goka"
	"github.com/spf13/viper"
)

var ThresholdService thresholdSvc.Service

func InitDependency(config *viper.Viper) {
	topicutil.EnsureStreamExists(config.GetString("deposit.topic"), config.GetStringSlice("broker.hosts"))
	topicutil.EnsureTableExists(string(threshold.Table), config.GetStringSlice("broker.hosts"))

	thresholdView, err := goka.NewView(config.GetStringSlice("broker.hosts"), threshold.Table, new(threshold.ThresholdCodec))
	if err != nil {
		log.Fatal(err)
	}
	go thresholdView.Run(context.Background())
	ThresholdService = thresholdSvc.NewService(&thresholdSvc.Args{
		View:   thresholdView,
		Config: config,
	})
}
