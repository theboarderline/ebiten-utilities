package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

func (g *GameserverClient) Register(name string) error {
	event := events.Event{
		Type:    "PLAYER_CONNECT",
		Message: name,
	}

	return g.SendMessage(event)
}

func (g *GameserverClient) Deregister(name string) error {
	event := events.Event{
		Type:    "PLAYER_DISCONNECT",
		Message: name,
	}

	return g.SendMessage(event)
}
