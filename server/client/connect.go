package client

import (
	"github.com/rs/zerolog/log"
)

func (g *GameserverClient) IsConnected() bool {
	return g.conn != nil
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
