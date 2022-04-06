// attempts to connect client and server, and send a ping
package lighthouse

import (
	"net"
	"testing"
)

func TestPing(t *testing.T) {
	server_conn, err := net.ListenPacket("udp", "127.0.0.1:2869")
	if err != nil {
		t.FailNow()
	}
	server := Lighthouse{server_conn}
	go server.Start()

	client_conn, err := net.Dial("udp", "127.0.0.1:2869")
	client_conn.Write([]byte(`{"q": "ping"}`)) // simply test if the ping query will receive a response

	buf := make([]byte, 1000)
	client_conn.Read(buf)
}
