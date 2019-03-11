package lib

import (
	"github.com/spf13/viper"
)

type DrlmcoreConfig struct {
	Server   string
	Port     string
	Tls      bool
	Cert     string
	User     string
	Password string
}

func SetDrlmcoreConfigDefaults() {
	viper.SetDefault("drlmcore.server", "localhost")
	viper.SetDefault("drlmcore.port", "50051")
	viper.SetDefault("drlmcore.tls", false)
	viper.SetDefault("drlmcore.cert", "cert/server.crt")
	viper.SetDefault("drlmcore.user", "drlmadmin")
	viper.SetDefault("drlmcore.password", "drlm3api")
}
