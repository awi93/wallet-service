package deposit

import (
	"github.com/awi93/wallet-service/src/dtos"
	"github.com/lovoo/goka"
	"github.com/spf13/viper"
)

type Service interface {
	EmitDeposit(request dtos.DepositRequest) error
}

type service struct {
	emitter *goka.Emitter
	config  *viper.Viper
}

type Args struct {
	Emitter *goka.Emitter
	Config  *viper.Viper
}

func NewService(args *Args) Service {
	return &service{
		emitter: args.Emitter,
		config:  args.Config,
	}
}
