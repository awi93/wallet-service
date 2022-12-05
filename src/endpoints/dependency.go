package endpoints

import (
	"context"
	"log"

	"github.com/awi93/wallet-service/src/models/balance"
	"github.com/awi93/wallet-service/src/models/deposit"
	"github.com/awi93/wallet-service/src/models/threshold"
	depositSvc "github.com/awi93/wallet-service/src/services/deposit"
	walletSvc "github.com/awi93/wallet-service/src/services/wallet"
	"github.com/awi93/wallet-service/src/utils/topicutil"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/storage"
	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var WalletService walletSvc.Service
var DepositService depositSvc.Service
var DepoEmitter *goka.Emitter

func InitDependency(config *viper.Viper) {
	topicutil.EnsureStreamExists(config.GetString("deposit.topic"), config.GetStringSlice("broker.hosts"))
	topicutil.EnsureTableExists(string(balance.Table), config.GetStringSlice("broker.hosts"))
	topicutil.EnsureTableExists(string(threshold.Table), config.GetStringSlice("broker.hosts"))

	opts := &opt.Options{
		BlockCacheCapacity: opt.MiB * 30,
		WriteBuffer:        opt.MiB * 10,
	}
	path := "tmp/goka"
	builder := storage.BuilderWithOptions(path, opts)

	balanceView, err := goka.NewView(config.GetStringSlice("broker.hosts"), balance.Table, new(balance.BalanceCodec), goka.WithViewStorageBuilder(builder))
	if err != nil {
		log.Fatalf(" Error : %v", err)
	}
	go (func() {
		err := balanceView.Run(context.Background())
		log.Fatalf("Error : %v", err)
	})()
	thresholdView, err := goka.NewView(config.GetStringSlice("broker.hosts"), threshold.Table, new(threshold.ThresholdCodec), goka.WithViewStorageBuilder(builder))
	if err != nil {
		log.Fatalf(" Error : %v", err)
	}
	go (func() {
		err := thresholdView.Run(context.Background())
		log.Fatalf(" Error : %v", err)
	})()
	DepoEmitter, err = goka.NewEmitter(config.GetStringSlice("broker.hosts"), goka.Stream(config.GetString("deposit.topic")), new(deposit.DepositCodec))
	if err != nil {
		log.Fatalf(" Error : %v", err)
	}
	WalletService = walletSvc.NewService(&walletSvc.Args{
		BalanceView:   balanceView,
		ThresholdView: thresholdView,
	})
	DepositService = depositSvc.NewService(&depositSvc.Args{
		Emitter: DepoEmitter,
		Config:  config,
	})
}
