package deposit

import (
	depositSvc "github.com/awi93/wallet-service/src/services/deposit"
)

type Deposit struct {
	service depositSvc.Service
}

type Args struct {
	Service depositSvc.Service
}

func NewHandler(Args *Args) *Deposit {
	return &Deposit{
		service: Args.Service,
	}
}
