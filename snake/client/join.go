package client

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

func (g *GameserverClient) Register(name string) error {
	message := []byte(fmt.Sprintf("PLAYER_CONNECT %s", name))
	if _, err := g.conn.Write(message); err != nil {
		log.Debug().Err(err).Msg("Error writing to UDP address")
		return err
	}

	log.Debug().Msgf("Sent UDP message: %s", message)
	return nil
}

func (g *GameserverClient) Deregister(name string) error {
	message := []byte(fmt.Sprintf("PLAYER_DISCONNECT %s", name))
	if _, err := g.conn.Write(message); err != nil {
		log.Debug().Err(err).Msg("Error writing to UDP address")
		return err
	}

	log.Debug().Msgf("Sent UDP message: %s", message)
	return nil
}
