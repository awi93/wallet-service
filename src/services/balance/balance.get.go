package balance

import (
	"github.com/awi93/wallet-service/src/models/balance"
)

func (s *service) GetBalance(walletId string) (*balance.Balance, error) {
	val, err := s.view.Get(walletId)
	if err != nil {
		return nil, err
	}
	return val.(*balance.Balance), err
}
