package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/object/snake"
	"github.com/theboarderline/ebiten-utilities/snake/param"
)

type FakeClient struct {
	Players map[string]*snake.Snake
}

func NewFakeClient() *FakeClient {
	return &FakeClient{
		Players: map[string]*snake.Snake{
			"enemy": snake.NewSnakeRandDirLoc(param.SnakeLength, param.SnakeSpeedInitial, &param.ColorSnake2),
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

func (g *FakeClient) SendMessage(message []byte) error {
	return nil
}

func (g *FakeClient) GetMessage() (string, error) {
	return "{\"name\":\"Enemy\", \"score\":100}", nil
}

func (g *FakeClient) Cleanup() error {
	return nil
}
