package config

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
)

type config struct {
	Debug     bool   `mapstructure:"debug"`
	LogCli    bool   `mapstructure:"log-cli"`
	FleetStr  string `mapstructure:"fleet"`
	Password  string `mapstructure:"password"`
	IName     string `mapstructure:"interface"`
	FleetAddr net.IP
	FleetNet  *net.IPNet
}

var Config config

// config.UnmarshalInitConfig: unmarshal viper to config.InitConfig global instance of config.initConfig
func UnmarshalConfig() error {
	err := viper.Unmarshal(&Config)
	if err != nil {
		return fmt.Errorf("Config Error: %w", err)
	}

	return nil
}
