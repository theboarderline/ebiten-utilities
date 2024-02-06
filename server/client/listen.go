package client

import (
	"github.com/rs/zerolog/log"
	"github.com/theboarderline/ebiten-utilities/server/events"
)

func (g *GameserverClient) HandleIncomingEvents() {
	buffer := make([]byte, g.bufferSize)

	for {
		if g.conn == nil {
			return
		}

		length, addr, err := g.conn.ReadFrom(buffer)
		if err != nil {
			log.Print(err)
			continue
		}

		if length == 0 {
			continue
		}

		event := events.Parse(string(buffer[:length]))

		g.incomingMessages <- event

		log.Info().Msgf("Received UDP message from %s: %s", addr, string(buffer[:length]))
	}
}

func (g *GameserverClient) HandleOutgoingEvents() {
	for event := range g.outgoingMessages {
		if g.conn == nil {
			return
		}

		if _, err := g.conn.Write([]byte(event.String())); err != nil {
			log.Error().Err(err).Msg("Error sending UDP message")
			continue
		}

		log.Info().Msgf("Sent UDP message: %s", event.String())
	}

	return
}
