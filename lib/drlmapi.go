package lib

import (
	"github.com/spf13/viper"
)

type DrlmcoreConfig struct {
	Server string
	Port   string
}

func SetDrlmcoreConfigDefaults() {
	viper.SetDefault("drlmcore.server", "localhost")
	viper.SetDefault("drlmcore.port", "50051")
}
