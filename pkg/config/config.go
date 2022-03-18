package config

import (
	"fmt"
	"net"

	"github.com/spf13/viper"
)

type config struct {
	Debug       bool   `mapstructure:"debug"`
	LogCli      bool   `mapstructure:"log-cli"`
	FleetStr    string `mapstructure:"fleet"`
	PasswordStr string `mapstructure:"password"`
	IName       string `mapstructure:"interface"`
	ListenPort  string `mapstructure:"listen-port"`
	Password    []byte
	FleetAddr   net.IP
	FleetNet    *net.IPNet
}

var Config config

// config.UnmarshalInitConfig: unmarshal viper to config.InitConfig global instance of config.initConfig
func UnmarshalConfig() error {
	err := viper.Unmarshal(&Config)
	if err != nil {
		return fmt.Errorf("config.UnmarshalConfig: %w", err)
	}

	return nil
}
