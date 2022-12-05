package balance

import (
	balanceSvc "github.com/awi93/wallet-service/src/services/balance"
	"github.com/awi93/wallet-service/src/utils/topicutil"
	"github.com/spf13/viper"
)

var BalanceService balanceSvc.Service

func InitDependency(config *viper.Viper) {
	topicutil.EnsureStreamExists(config.GetString("deposit.topic"), config.GetStringSlice("broker.hosts"))

	BalanceService = balanceSvc.NewService(&balanceSvc.Args{
		Config: config,
	})
}
