package client

import (
	"github.com/theboarderline/ebiten-utilities/server/events"
)

func (g *GameserverClient) Register(name string) {

	g.SendMessage(&events.Event{
		Type:       events.PLAYER_CONNECT,
		PlayerName: name,
	})

	return
}

func (g *GameserverClient) Deregister(name string) {

	g.SendMessage(&events.Event{
		Type:       events.PLAYER_DISCONNECT,
		PlayerName: name,
	})

	return
}
