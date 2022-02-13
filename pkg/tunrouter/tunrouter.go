package tunrouter

import (
	"github.com/prairir/Buoy/pkg/config"
	"github.com/prairir/Buoy/pkg/ethrouter"

	"github.com/songgao/water"
	"golang.org/x/sync/errgroup"
)

func Run(eg *errgroup.Group, inf *water.Interface,  chan ethrouter.Packet, readq chan []byte) error {
	eg.Go(func() error{
		return writer(, writeq chan []byte)
	})
}
