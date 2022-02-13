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

func reader(inf *water.Interface, readq chan ethrouter.Packet) error {
	buf := make([]byte, maxIPPacketSize)
	for {
		n, err := inf.Read(buf)
		if err != nil {
			// if connection is closed, exit nicely
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			return fmt.Errorf("tunrouter.reader: %w", err)
		}

		// TODO possible race condition here. cop can exist in 2 go routines at the same time, the first ones memory gets overwriten when the second one gets dispatched
		cop := make([]byte, n)
		copy(cop, buf[:n])

		go func(lBuf []byte) {
			//TODO put translation here
			saddr := fmt.Sprintf("%d.%d.%d.%d", buf[16], buf[17], buf[18], buf[19])
			addr, _ := net.ResolveIPAddr("udp", saddr) // TODO think of a way to propegate errors nicely without a wait group
			pack := ethrouter.Packet{
				Addr:    addr, // TODO use translation addr
				Payload: lBuf, // TODO use translation payload(encrypted + compressed)
			}
			readq <- pack
		}(cop)

	}
}
