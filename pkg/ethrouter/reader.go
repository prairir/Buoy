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

			return fmt.Errorf("tunrouter.reader: %w", err)
		}

		cop := make([]byte, n)
		copy(cop, buf[:n])

		go func(lBuf []byte) {
			//TODO put translation here
			readq <- lBuf
		}(cop)
	}
}
