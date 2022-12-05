package bootstrap

import "github.com/spf13/viper"

var config = viper.New()

func ReadConfig(configFile string) error {
	config.SetConfigFile(configFile)
	config.AddConfigPath(".")
	return config.ReadInConfig()
}

func GetConfig() *viper.Viper {
	return config
}
