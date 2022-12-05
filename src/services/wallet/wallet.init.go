package wallet

import (
	"github.com/awi93/wallet-service/src/dtos"
	"github.com/lovoo/goka"
)

type Service interface {
	GetWallet(walletId string) (*dtos.Wallet, error)
}

type Args struct {
	BalanceView   *goka.View
	ThresholdView *goka.View
}

type service struct {
	balanceView   *goka.View
	thresholdView *goka.View
}

func NewService(args *Args) Service {
	return &service{
		balanceView:   args.BalanceView,
		thresholdView: args.ThresholdView,
	}
}
