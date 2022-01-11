package tun

import (
	"fmt"
	"github.com/milosgajdos/tenus"
	"github.com/songgao/water"
	"net"
)

// tun.Init: create, initialize, and bring up tun interface
// Creates interface with name `<iname>` and address `<ip>` with range `<network>`
func Init(iname string, ip net.IP, network *net.IPNet) (*water.Interface, error) {

	// make a new link named iname
	inf, err := tenus.NewLink(iname)
	if err != nil {
		return nil, fmt.Errorf("tun.Init error: %w", err)
	}

	// set the link ip address and range
	// This is the same as running `ip addr add <ip + network in CIDR notation> dev <iname>`
	err = inf.SetLinkIp(ip, network)
	if err != nil {
		return nil, fmt.Errorf("tun.Init error: %w", err)
	}

	// set the link up
	// This is the same as `ip link set dev <iname> up`
	err = inf.SetLinkUp()
	if err != nil {
		return nil, fmt.Errorf("tun.Init error: %w", err)
	}

	// config for tun device
	tunConf := water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: iname,
		},
	}

	// actually create the tun device
	tunInf, err := water.New(tunConf)
	if err != nil {
		return nil, fmt.Errorf("tun.Init error: %w", err)
	}

	return tunInf, nil
}
