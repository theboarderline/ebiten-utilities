package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/object/snake"
)

type Client interface {
	IsConnected() bool
	GetPlayers() map[string]*snake.Snake
	Register(name string) error
	Deregister(name string) error
	SendMessage(event *events.Event)
	GetMessage() *events.Event
	Cleanup() error
}
