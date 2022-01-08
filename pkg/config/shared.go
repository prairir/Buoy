package config

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

// errors
var ErrInvalidPasswordLength = errors.New("invalid password length")
var ErrInvalidInterfaceName = errors.New("invalid interface name length")
var ErrInvalidFleet = errors.New("invalid fleet")

type sharedConfig struct {
	RWM      sync.RWMutex
	Fleet    *net.IPNet
	Password []byte
	IName    string
}

var SharedConfig sharedConfig

// config.SetSharedConfig: sets config package global instance of sharedConfig struct
//
// global instance of sharedConfig is to be shared between concurrent processes
func SetSharedConfig(fleet *net.IPNet, password string, iname string) error {
	if fleet == nil {
		return fmt.Errorf("Config.SetSharedConfig error: %w, cannot be nil", ErrInvalidFleet)
	}

	// if interface name isnt empty
	if iname == "" {
		return fmt.Errorf("config.SetSharedConfig error: %w, cant be empty string", ErrInvalidPasswordLength)
	}

	// if password length doesnt follow scheme
	if len(password)-16 < 1 {
		return fmt.Errorf("config.SetSharedConfig error: %w, must be at least 16 + 8n where 1 <= n but got %d",
			ErrInvalidPasswordLength,
			len(password))
	}

	// create shared config and set it to package global shared config
	SharedConfig = sharedConfig{
		RWM:      sync.RWMutex{},
		Fleet:    fleet,
		Password: []byte(password),
		IName:    iname,
		// TODO add fleet mates(peer list)
	}

	return nil
}
