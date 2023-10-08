package events

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"strings"
)

type Event struct {
	PlayerName     string     `json:"playerName,omitempty"`
	Type           string     `json:"type,omitempty"`
	Message        string     `json:"message,omitempty"`
	InputDirection DirectionT `json:"direction,omitempty"`
}

func (e Event) String() (event string) {
	event = e.Type

	if e.PlayerName != "" {
		event = fmt.Sprintf("%s %s", event, e.PlayerName)
	}

	if e.Message != "" {
		event = fmt.Sprintf("%s %s", event, e.Message)
	}

	return event
}

func Parse(input string) (event Event) {
	parts := strings.Split(input, " ")

	if len(parts) > 0 {
		event.Type = parts[0]
	}

	if len(parts) > 1 {
		event.PlayerName = parts[1]
	}

	if len(parts) > 2 {
		event.Message = parts[2]
	}

	return event
}

func (e Event) Marshal() (event string) {
	eventBytes, err := json.Marshal(e)
	if err != nil {
		log.Error().Err(err).Msg("Error marshaling event")
		return ""
	}

	if string(eventBytes) == "null" {
		return ""
	}

	event = strings.ReplaceAll(string(eventBytes), "\\", "")

	return event
}
