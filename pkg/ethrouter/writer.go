package ethrouter

import (
	"errors"
	"fmt"
	"net"
)

// ethrouter.Packet: outgoing packet
// contains the address of where its going to be sent do
// and a payload
// TODO: find a better place to put this or a better name
type Packet struct {
	Addr    net.Addr
	Payload []byte
}

// tunrouter.writer: Reads from the writeq channel and writes to a UDP connection
func writer(conn *net.UDPConn, writeq chan Packet) error {
	for {
		pack := <-writeq
		_, err := conn.WriteTo(pack.Payload, pack.Addr)
		if err != nil {
			// if connection is closed, exit nicely
			// TODO verify that is the best way to close the writer
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			return fmt.Errorf("ethrouter.writer: %w", err)
		}
	}
}
