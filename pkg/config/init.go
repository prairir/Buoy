package config

import (
	"fmt"
	"net"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type initConfig struct {
	Fleet    *net.IPNet `mapstructure:"fleet"`
	Password string     `mapstructure:"password"`
	IName    string     `mapstructure:"interface"`
	Debug    bool       `mapstructure:"debug"`
	LogCli   bool       `mapstructure:"log-cli"`
}

var InitConfig initConfig

// config.UnmarshalInitConfig: unmarshal viper to config.InitConfig global instance of config.initConfig
func UnmarshalInitConfig() error {
	err := viper.Unmarshal(&InitConfig, func(m *mapstructure.DecoderConfig) {
		m.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToIPNetHookFunc(),
		)
	})
	if err != nil {
		return fmt.Errorf("Config Error: %w", err)
	}

	return nil
}
