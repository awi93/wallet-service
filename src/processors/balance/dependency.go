package balance

import (
	"context"
	"log"

	"github.com/awi93/wallet-service/src/models/balance"
	balanceSvc "github.com/awi93/wallet-service/src/services/balance"
	"github.com/awi93/wallet-service/src/utils/topicutil"
	"github.com/lovoo/goka"
	"github.com/spf13/viper"
)

var BalanceService balanceSvc.Service

func InitDependency(config *viper.Viper) {
	topicutil.EnsureStreamExists(config.GetString("deposit.topic"), config.GetStringSlice("broker.hosts"))
	topicutil.EnsureTableExists(string(balance.Table), config.GetStringSlice("broker.hosts"))

	balanceView, err := goka.NewView(config.GetStringSlice("broker.hosts"), balance.Table, new(balance.BalanceCodec))
	if err != nil {
		log.Fatal(err)
	}
	go balanceView.Run(context.Background())

	BalanceService = balanceSvc.NewService(&balanceSvc.Args{
		View:   balanceView,
		Config: config,
	})
}
