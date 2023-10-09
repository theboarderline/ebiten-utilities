package client

import (
	"github.com/go-faker/faker/v4"
	"github.com/rs/zerolog/log"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/object/snake"
	"github.com/theboarderline/ebiten-utilities/snake/param"
	"math/rand"
)

type FakeClient struct {
	Players map[string]*snake.Snake
}

func NewFakeClient(name string) *FakeClient {
	if name == "" {
		name = faker.Name()
	}

	return &FakeClient{
		Players: map[string]*snake.Snake{
			name: snake.NewSnakeRandDirLoc(name, param.SnakeLength, param.SnakeSpeedInitial, &param.ColorSnake2),
		},
	}
}

func (g *FakeClient) HandleIncomingEvents() {
	log.Debug().Msg("HandleIncomingEvents")
}

func (g *FakeClient) HandleOutgoingEvents() {
	log.Debug().Msg("HandleOutgoingEvents")
}

func (g *FakeClient) GetPlayers() map[string]*snake.Snake {
	return g.Players
}

func (g *FakeClient) IsConnected() bool {
	return true
}

func (g *FakeClient) Register(name string) {
	g.Players[name] = snake.NewSnakeRandDirLoc(name, param.SnakeLength, param.SnakeSpeedInitial, &param.ColorSnake2)
	return
}

func (g *FakeClient) Deregister(name string) {
	delete(g.Players, name)
	return
}

func (g *FakeClient) SendMessage(event *events.Event) {
	return
}

func (g *FakeClient) GetMessage() *events.Event {

	if rand.Float64() < 0.99999999999999999999999 {
		return nil
	}

	event := events.Event{
		PlayerName:     events.ENEMY,
		Type:           events.PLAYER_INPUT,
		InputDirection: events.NewRandomDirection(),
	}
	return &event
}

func (g *FakeClient) Cleanup() error {
	return nil
}
