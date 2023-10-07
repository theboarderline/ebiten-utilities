package client

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net"
)

func (g *GameserverClient) Connect() error {
	udpServer, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", g.address, g.port))

	if err != nil {
		log.Debug().Err(err).Msg("Error resolving UDP address")
		return nil
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		log.Debug().Err(err).Msg("Error dialing UDP address")
		return nil
	}

	g.conn = conn

	log.Debug().Msgf("Dialed UDP address: %s:%d", g.address, g.port)
	return nil
}

func (g *GameserverClient) Cleanup() error {
	if g.conn == nil {
		return nil
	}

	if err := g.conn.Close(); err != nil {
		log.Debug().Err(err).Msg("Error closing UDP connection")
		return err
	}

	g.conn = nil

	log.Debug().Msg("Closed UDP connection")
	return nil
}
