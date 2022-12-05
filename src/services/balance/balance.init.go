package balance

import (
	"github.com/awi93/wallet-service/src/models/balance"
	"github.com/lovoo/goka"
	"github.com/spf13/viper"
)

type Service interface {
	GetBalance(walletId string) (*balance.Balance, error)
	ProcessCallback(ctx goka.Context, msg interface{})
}

type service struct {
	view   *goka.View
	config *viper.Viper
}

type Args struct {
	View   *goka.View
	Config *viper.Viper
}

func NewService(args *Args) Service {
	return &service{
		view:   args.View,
		config: args.Config,
	}
}
