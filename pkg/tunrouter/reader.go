package tunrouter

import (
	"errors"
	"fmt"
	"net"

	"github.com/prairir/Buoy/pkg/config"
	"github.com/prairir/Buoy/pkg/ethrouter"
	"github.com/prairir/Buoy/pkg/translate"
	"github.com/rs/zerolog/log"
	"github.com/songgao/water"
)

var FleetList map[string]net.UDPAddr = map[string]net.UDPAddr{
	/*
		"192.168.18.2": {
			IP:   net.IP("192.168.2.58"),
			Port: 8080,
		},
	*/
}

const (
	maxIPPacketSize = 65535
)

func reader(inf *water.Interface, tun2EthQ chan ethrouter.Packet) error {
	addr, _ := net.ResolveUDPAddr("udp", "192.168.2.58:8080")
	FleetList["192.168.18.2"] = *addr
	buf := make([]byte, maxIPPacketSize)
	for {
		n, err := inf.Read(buf)
		if err != nil {
			// TODO fix that error, its useless(stupid copy paste). Change it to be the same behavior(interface close)
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			return fmt.Errorf("tunrouter.reader: %w", err)
		}

		go func(buf []byte, n int) {
			cop := make([]byte, n)
			copy(cop, buf[:n])

			// get the version xxxx0000 is the version in the first byte
			version := cop[0] >> 4
			if int(version) != 4 {
				return // exit if not ipv4
			}

			packPayload, err := translate.Translate(cop, config.Config.Password)
			// TODO: verify this behavior
			if err != nil { // if error, drop packet
				log.Error().Err(err).Msg("tunrouter.Reader: couldnt translate package: Packet dropped")
				return
			}

			// Fastest way to convert bytes to IP
			// I benchmarked this, sprintf, and my own way using strconv, int casting, and string concatination.
			// This is fastest(16 ns/op), then mine(48 ns/op), then sprintf(119 ns/op)
			srcAddr := net.IP(cop[16:20])

			destAddr, ok := FleetList[srcAddr.String()]
			// TODO: verify this behavior
			if !ok {
				log.Error().Str("addr", srcAddr.String()).Msg("tunrouter.Reader: couldnt find matching udp addr in fleetlist: Packet dropped")
				return
			}

			pack := ethrouter.Packet{
				Addr:    &destAddr,
				Payload: packPayload,
			}
			tun2EthQ <- pack
		}(buf, n)

	}
}
