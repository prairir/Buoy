/*
	lighthouse is the rolodex that allows peer discovery to occur.
	It can run on a central server that is agreed to be the authority and assumed to be a secure server.
*/
package lighthouse

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
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
			l.conn.Close()
			panic(err)
		}

		fmt.Printf("packet-received: bytes=%d from=%s\n", n, addr.String())

		fmt.Fprintln(os.Stderr, buf[:n])
		query := make(map[string]string)
		err = json.Unmarshal(buf[:n], &query)
		if err != nil {
			l.conn.Close()
			panic(err)
		}
		if query["q"] == "ping" {
			resp := map[string]string{
				"r": "pong",
			}
			respData, _ := json.Marshal(resp)
			n, err = l.conn.WriteTo(respData, addr)
			fmt.Printf("packet sent: bytes=%d from=%s\n", n, addr.String())
		} else if query["q"] == "locate" {
			networkName := query["network"]
			l.findNetwork(networkName)
		}
	}
}

// given a network name, searches its nodes to find a node that has information regarding the network.
func (l *Lighthouse) findNetwork(networkName string) {
	// get the closest node using the network name
	// this is where the chord code starts.
}
