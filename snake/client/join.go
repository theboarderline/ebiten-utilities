package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

func (g *GameserverClient) Register(name string) error {
	event := events.Event{
		Type:       events.PLAYER_CONNECT,
		PlayerName: name,
	}

	return g.SendMessage(event)
}

func (g *GameserverClient) Deregister(name string) error {
	event := events.Event{
		Type:       events.PLAYER_DISCONNECT,
		PlayerName: name,
	}

	return g.SendMessage(event)
}
