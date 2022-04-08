package ethrouter

import (
	"errors"
	"fmt"
	"net"

	"github.com/prairir/Buoy/pkg/config"
	"github.com/prairir/Buoy/pkg/translate"
	"github.com/rs/zerolog/log"
)

const (
	maxUDPPacketSize = 65507
)

func reader(conn *net.UDPConn, eth2TunQ chan []byte) error {

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

		go func(buf []byte, n int) {
			pack := buf[:n]

			pack, err = translate.Untranslate(pack, config.Config.Password)
			// TODO: verify this behavior
			if err != nil { // if error, drop packet
				log.Error().Err(err).Msg("Packet dropped")
				return
			}

			eth2TunQ <- pack
		}(buf, n)
	}
}
