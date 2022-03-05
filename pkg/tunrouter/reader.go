package tunrouter

import (
	"errors"
	"fmt"
	"net"

	"github.com/prairir/Buoy/pkg/ethrouter"
	"github.com/songgao/water"
)

const (
	maxIPPacketSize = 65535
)

func reader(inf *water.Interface, tun2EthQ chan ethrouter.Packet) error {
	buf := make([]byte, maxIPPacketSize)
	for {
		n, err := inf.Read(buf)
		if err != nil {
			// if connection is closed, exit nicely
			// TODO fix that error, its useless(stupid copy paste)
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			return fmt.Errorf("tunrouter.reader: %w", err)
		}

		go func(buf []byte) {
			cop := make([]byte, n)
			copy(cop, buf[:n])

			//TODO put translation here
			saddr := fmt.Sprintf("%d.%d.%d.%d", buf[16], buf[17], buf[18], buf[19])
			addr, _ := net.ResolveIPAddr("udp", saddr) // TODO think of a way to propegate errors nicely without a wait group
			pack := ethrouter.Packet{
				Addr:    addr, // TODO use translation addr
				Payload: cop,  // TODO use translation payload(encrypted + compressed)
			}
			tun2EthQ <- pack
		}(buf)

	}
}
