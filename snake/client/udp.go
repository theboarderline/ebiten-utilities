package client

import (
	"time"

	"github.com/rs/zerolog/log"
)

func (g *GameserverClient) SendMessage(message []byte) error {
	if _, err := g.conn.Write(message); err != nil {
		log.Debug().Err(err).Msg("Error writing to UDP address")
		return err
	}

	log.Debug().Msgf("Sent UDP message: %s", message)
	return nil
}

func (g *GameserverClient) GetMessage() (string, error) {
	if err := g.conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Debug().Err(err).Msg("Error setting UDP timeout")
		return "", err
	}

	received := make([]byte, 1024)
	if _, err := g.conn.Read(received); err != nil {
		log.Debug().Err(err).Msg("Error reading from UDP address")
		return "", err
	}

	log.Debug().Msgf("Received UDP message: %s", string(received))
	return string(received), nil
}
