package client

import (
	"github.com/rs/zerolog/log"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/object/snake"
	"net"
	"strings"
)

type GameserverClient struct {
	conn    *net.UDPConn
	address string
	port    int
}

func NewGameserverClient(address string, port int) *GameserverClient {
	return &GameserverClient{
		address: address,
		port:    port,
	}
}

func (g *GameserverClient) IsConnected() bool {
	return g.conn != nil
}

func (g *GameserverClient) GetPlayers() map[string]*snake.Snake {
	event := events.Event{
		Type: events.GET_PLAYERS,
	}
	if err := g.SendMessage(event); err != nil {
		return nil
	}

	response, err := g.GetMessage()
	if err != nil {
		return nil
	}

	return makeSnakes(response.Message)
}

func makeSnakes(response string) map[string]*snake.Snake {
	snakes := make(map[string]*snake.Snake)

	for _, s := range parseSnakeResponse(response) {
		snakes[s.Name] = s
	}

	return snakes
}

func parseSnakeResponse(response string) []*snake.Snake {
	snakes := make([]*snake.Snake, 0)

	for _, item := range strings.Split(response, "\n") {
		if item == "" {
			continue
		}

		s, err := snake.NewSnakeFromResponse(item)
		if err != nil {
			log.Print(err)
		}

		snakes = append(snakes, s)
	}

	return nil
}
