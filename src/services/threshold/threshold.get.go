package threshold

import (
	"github.com/awi93/wallet-service/src/models/threshold"
)

func (s *service) GetThreshold(walletId string) (*threshold.Threshold, error) {
	val, err := s.view.Get(walletId)
	if err != nil {
		return nil, err
	}
	return val.(*threshold.Threshold), nil
}
