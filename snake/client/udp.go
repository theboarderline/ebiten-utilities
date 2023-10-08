package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

func (g *GameserverClient) SendMessage(event *events.Event) {
	if event == nil {
		return
	}

	g.OutgoingMessages <- *event

	return
}

func (g *GameserverClient) GetMessage() *events.Event {
	event := <-g.IncomingMessages

	return &event
}
