package threshold

import (
	thresholdSvc "github.com/awi93/wallet-service/src/services/threshold"
	"github.com/awi93/wallet-service/src/utils/topicutil"
	"github.com/spf13/viper"
)

var ThresholdService thresholdSvc.Service

func InitDependency(config *viper.Viper) {
	topicutil.EnsureStreamExists(config.GetString("deposit.topic"), config.GetStringSlice("broker.hosts"))

	ThresholdService = thresholdSvc.NewService(&thresholdSvc.Args{
		Config: config,
	})
}
