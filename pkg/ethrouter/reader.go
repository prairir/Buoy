package ethrouter

import (
	"errors"
	"fmt"
	"net"
)

const (
	maxUDPPacketSize = 65507
)

func reader(conn *net.UDPConn, readq chan []byte) error {

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

		// TODO possible race condition here. cop can exist in 2 go routines at the same time, the first ones memory gets overwriten when the second one gets dispatched
		cop := make([]byte, n)
		copy(cop, buf[:n])

		go func(lBuf []byte) {
			//TODO put translation here
			readq <- lBuf
		}(cop)
	}
}
