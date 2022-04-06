/*
	lighthouse is the rolodex that allows peer discovery to occur.
	It can run on a central server that is agreed to be the authority and assumed to be a secure server.
*/
package lighthouse

import (
	"encoding/json"
	"fmt"
	"net"
)

type Lighthouse struct {
	conn net.PacketConn
}

func (l *Lighthouse) Start() {
	// listens for connections from new users, and reacts accordingly
	buf := make([]byte, 65535)
	for {
		n, addr, err := l.conn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("packet-received: bytes=%d from=%s\n", n, addr.String())

		query := make(map[string]string)
		err = json.Unmarshal(buf[:n], &query)
		if err != nil {
			panic(err)
		}

		resp := map[string]string{
			"r": "ping",
		}
		respData, _ := json.Marshal(resp)
		n, err = l.conn.WriteTo(respData, addr)
		if err != nil {
			panic(err)
		}

		fmt.Printf("packet-written: bytes=%d to=%s\n", n, addr.String())
	}
}
