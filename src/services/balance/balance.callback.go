package balance

import (
	"log"

	"github.com/awi93/wallet-service/src/models/balance"
	"github.com/awi93/wallet-service/src/models/deposit"
	"github.com/lovoo/goka"
)

func (s *service) ProcessCallback(ctx goka.Context, msg interface{}) {
	log.Println("Receiving New Message")
	var b *balance.Balance
	var deposit = msg.(*deposit.Deposit)
	if v := ctx.Value(); v == nil {
		b = s.composeBalance(deposit)
	} else {
		var latest = v.(*balance.Balance)
		b = s.calculateBalance(latest, deposit)
	}
	ctx.SetValue(b)
}

func (s *service) composeBalance(deposit *deposit.Deposit) *balance.Balance {
	var balance = &balance.Balance{
		WalletId: deposit.WalletId,
		Amount:   deposit.Amount,
	}
	return balance
}

func (s *service) calculateBalance(balance *balance.Balance, deposit *deposit.Deposit) *balance.Balance {
	balance.Amount += deposit.Amount
	return balance
}
