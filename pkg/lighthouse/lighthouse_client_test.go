package lighthouse

import (
	"net"
	"testing"
)

func TestPingClient(t *testing.T) {
	server_conn, err := net.ListenPacket("udp", "127.0.0.1:3334")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	server := Lighthouse{server_conn}
	go func() {
		defer server_conn.Close()
		server.Start()
	}()

	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:3334")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	client_conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		t.FailNow()
	}
	defer client_conn.Close()
	client := Boat{client_conn}

	client.Ping()
}
