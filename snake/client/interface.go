package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/events"
)

type Client interface {
	HandleIncomingEvents()
	HandleOutgoingEvents()
	IsConnected() bool
	GetPlayers(senderName string)
	Register(name string)
	Deregister(name string)
	SendMessage(event *events.Event)
	GetMessage() *events.Event
	Cleanup() error
}
