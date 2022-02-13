package tunrouter

import (
	"github.com/prairir/Buoy/pkg/ethrouter"

	"github.com/songgao/water"
	"golang.org/x/sync/errgroup"
)

func Run(eg *errgroup.Group, inf *water.Interface, readq chan ethrouter.Packet, writeq chan []byte) error {
	eg.Go(func() error {
		return writer(inf, writeq)
	})

	eg.Go(func() error {
		return reader(inf, readq)
	})
}
