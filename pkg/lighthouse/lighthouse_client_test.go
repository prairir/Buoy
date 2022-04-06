package lighthouse

import (
	"net"
	"testing"
)

func TestPingClient(t *testing.T) {
	server_conn, err := net.ListenPacket("udp", "127.0.0.1:2869")
	if err != nil {
		t.FailNow()
	}
	server := Lighthouse{server_conn}
	go server.Start()

	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2869")
	if err != nil {
		t.FailNow()
	}
	client_conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		t.FailNow()
	}
	client := Boat{client_conn}

	client.Ping()

}
