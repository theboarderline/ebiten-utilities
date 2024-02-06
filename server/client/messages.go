package client

import (
	"github.com/theboarderline/ebiten-utilities/server/events"
)

func (g *GameserverClient) SendMessage(event *events.Event) {
	if event == nil {
		return
	}

	g.outgoingMessages <- *event

	return
}

func (g *GameserverClient) GetMessage() *events.Event {
	event := <-g.incomingMessages

	return &event
}
