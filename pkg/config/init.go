package config

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
)

type initConfig struct {
	Debug     bool   `mapstructure:"debug"`
	LogCli    bool   `mapstructure:"log-cli"`
	FleetStr  string `mapstructure:"fleet"`
	Password  string `mapstructure:"password"`
	IName     string `mapstructure:"interface"`
	FleetAddr net.IP
	FleetNet  *net.IPNet
}

var InitConfig initConfig

// config.UnmarshalInitConfig: unmarshal viper to config.InitConfig global instance of config.initConfig
func UnmarshalInitConfig() error {
	err := viper.Unmarshal(&InitConfig)
	if err != nil {
		return fmt.Errorf("Config Error: %w", err)
	}

	return nil
}
