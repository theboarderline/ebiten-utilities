package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/object/snake"
)

type Client interface {
	HandleIncomingEvents()
	HandleOutgoingEvents()
	IsConnected() bool
	GetPlayers(senderName string) map[string]*snake.Snake
	Register(name string)
	Deregister(name string)
	SendMessage(event *events.Event)
	GetMessage() *events.Event
	Cleanup() error
}
