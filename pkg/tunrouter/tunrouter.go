package tunrouter

import (
	"github.com/prairir/Buoy/pkg/ethrouter"

	"github.com/songgao/water"
	"golang.org/x/sync/errgroup"
)

func Run(eg *errgroup.Group, inf *water.Interface, tun2EthQ chan ethrouter.Packet, eth2TunQ chan []byte) error {
	eg.Go(func() error {
		return writer(inf, eth2TunQ)
	})

	eg.Go(func() error {
		return reader(inf, tun2EthQ)
	})
}
