// attempts to connect client and server, and send a ping
package lighthouse

import (
	"net"
	"testing"
)

func TestPing(t *testing.T) {
	server_conn, err := net.ListenPacket("udp", "127.0.0.1:3333")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	go func() {
		defer server_conn.Close()
		server := Lighthouse{server_conn}
		server.Start()
	}()
	client_conn, err := net.Dial("udp", "127.0.0.1:3333")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer client_conn.Close()
	client_conn.Write([]byte(`{"q": "ping"}`)) // simply test if the ping query will receive a response

	buf := make([]byte, 1000)
	client_conn.Read(buf)
}
