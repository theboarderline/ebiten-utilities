package client

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

func (g *GameserverClient) HandleIncomingEvents() {
	buffer := make([]byte, 1024)
	for {
		length, addr, err := g.conn.ReadFrom(buffer)
		if err != nil {
			log.Print(err)
			continue
		}

		if length == 0 {
			continue
		}

		var event events.Event
		if err = json.Unmarshal(buffer[:length], event); err != nil {
			log.Print(err)
			continue
		}

		g.IncomingMessages <- event

		log.Info().Msgf("Received UDP message from %s: %s", addr, string(buffer[:length]))
	}
}

func (g *GameserverClient) HandleOutgoingEvents() {

	for event := range g.OutgoingMessages {
		if _, err := g.conn.Write([]byte(event.String())); err != nil {
			log.Error().Err(err).Msg("Error sending UDP message")
			continue
		}

		log.Info().Msgf("Sent UDP message: %s", event.String())
	}

	return
}
