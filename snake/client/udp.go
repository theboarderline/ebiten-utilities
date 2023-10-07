package client

import (
	"encoding/json"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"time"

	"github.com/rs/zerolog/log"
)

func (g *GameserverClient) SendMessage(event events.Event) error {

	message, err := json.Marshal(event)
	if err != nil {
		log.Debug().Err(err).Msg("Error marshalling message")
		return err
	}

	if _, err = g.conn.Write(message); err != nil {
		log.Debug().Err(err).Msg("Error writing to UDP address")
		return err
	}

	log.Debug().Msgf("Sent UDP message: %s", message)
	return nil
}

func (g *GameserverClient) GetMessage() (event events.Event, err error) {
	if err = g.conn.SetReadDeadline(time.Now().Add(10 * time.Second)); err != nil {
		log.Debug().Err(err).Msg("Error setting UDP timeout")
		return events.Event{}, err
	}

	received := make([]byte, 1024)
	length, err := g.conn.Read(received)
	if err != nil {
		log.Debug().Err(err).Msg("Error reading from UDP address")
		return events.Event{}, err
	}

	if err = json.Unmarshal(received[:length], &event); err != nil {
		log.Debug().Err(err).Msg("Error unmarshalling message")
		return events.Event{}, err
	}

	log.Debug().Msgf("Received UDP message: %s", string(received))
	return event, nil
}
