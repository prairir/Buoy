package ethrouter

import (
	"errors"
	"fmt"
	"net"
)

const (
	maxUDPPacketSize = 65507
)

func reader(conn *net.UDPConn, eth2TunQ chan []byte) error {

	buf := make([]byte, maxUDPPacketSize)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			// if connection is closed, exit nicely
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			return fmt.Errorf("ethrouter.reader: %w", err)
		}

		go func(buf []byte) {
			//TODO put translation here
			cop := make([]byte, n)
			copy(cop, buf[:n])
			eth2TunQ <- cop
		}(buf)
	}
}
