package ethrouter

import (
	"fmt"
	"net"

	"github.com/prairir/Buoy/pkg/config"

	"golang.org/x/sync/errgroup"
)

// ethrouter.Run: spawns 2 goroutines for ethrouter.writer and ethrouter.reader
// properly routes each channel to their respective goroutine. writer gets the writeq
// and reader gets the readq. Also creates the socket that they share
// TODO look into more ergonomic way to initialize gorountines and pass around channels
func Run(eg *errgroup.Group, writeq chan Packet, readq chan []byte) error {
	// listen locally with the listen port
	laddr, err := net.ResolveUDPAddr("udp4", ":"+config.Config.ListenPort)
	if err != nil {
		return fmt.Errorf("ethrouter.Run: %w", err)
	}

	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		return fmt.Errorf("ethrouter.Run: %w", err)
	}

	eg.Go(func() error {
		return writer(conn, writeq)
	})

	eg.Go(func() error {
		return reader(conn, readq)
	})

	return nil
}
