package deposit

import (
	"fmt"

	"github.com/awi93/wallet-service/src/dtos"
	"github.com/awi93/wallet-service/src/models/deposit"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *service) EmitDeposit(request dtos.DepositRequest) error {
	var depoTime = timestamppb.Now()
	var deposit = &deposit.Deposit{
		WalletId:    request.WalletId,
		Amount:      request.Amount,
		DepositTime: depoTime,
	}

	err := s.emitter.EmitSync(fmt.Sprintf("%d", request.WalletId), deposit)

	return err
}
