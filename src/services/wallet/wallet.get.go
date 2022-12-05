package wallet

import (
	"github.com/awi93/wallet-service/errors"
	"github.com/awi93/wallet-service/src/dtos"
	"github.com/awi93/wallet-service/src/models/balance"
	"github.com/awi93/wallet-service/src/models/threshold"
)

func (s *service) GetWallet(walletId string) (*dtos.Wallet, error) {
	var wallet *dtos.Wallet

	b, err := s.balanceView.Get(walletId)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, errors.RecordNotFound
	}

	t, err := s.thresholdView.Get(walletId)
	if err != nil {
		return nil, err
	}

	var AboveThreshold = false
	if t != nil {
		AboveThreshold = t.(*threshold.Threshold).IsAboveThreshold
	}

	wallet = &dtos.Wallet{
		WalletId:       int(b.(*balance.Balance).WalletId),
		Balance:        b.(*balance.Balance).Amount,
		AboveThreshold: AboveThreshold,
	}

	return wallet, nil
}
