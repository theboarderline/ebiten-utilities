package client

import (
	"github.com/go-faker/faker/v4"
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

func (g *FakeClient) GetPlayers() map[string]*snake.Snake {
	return g.Players
}

func (g *FakeClient) Connect() error {
	return nil
}

func (g *FakeClient) IsConnected() bool {
	return true
}

func (g *FakeClient) Register(name string) error {
	return nil
}

func (g *FakeClient) Deregister(name string) error {
	return nil
}

func (g *FakeClient) SendMessage(event events.Event) error {
	return nil
}

func (g *FakeClient) GetMessage() (events.Event, error) {

	if rand.Float64() < 0.99 {
		return events.Event{}, nil
	}

	event := events.Event{
		PlayerName: "enemy",
		Type:       events.PLAYER_INPUT,
	}

	if rand.Float64() < 0.25 {
		event.Message = string(snake.DirectionUp)
	} else if rand.Float64() < 0.5 {
		event.Message = string(snake.DirectionDown)
	} else if rand.Float64() < 0.75 {
		event.Message = string(snake.DirectionLeft)
	} else {
		event.Message = string(snake.DirectionRight)
	}

	return event, nil
}

func (g *FakeClient) Cleanup() error {
	return nil
}
