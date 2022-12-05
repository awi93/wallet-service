package balance

import (
	"github.com/lovoo/goka"
	"github.com/spf13/viper"
)

type Service interface {
	ProcessCallback(ctx goka.Context, msg interface{})
}

type service struct {
	config *viper.Viper
}

type Args struct {
	Config *viper.Viper
}

func NewService(args *Args) Service {
	return &service{
		config: args.Config,
	}
}
