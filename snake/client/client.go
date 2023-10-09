package client

import (
	"github.com/rs/zerolog/log"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/object/snake"
	"github.com/theboarderline/ebiten-utilities/snake/param"
	"net"
	"strconv"
	"strings"
)

type GameserverClient struct {
	bufferSize       int
	conn             NetConn
	addr             net.UDPAddr
	incomingMessages chan events.Event
	outgoingMessages chan events.Event
}

func NewUDPConn(addr net.UDPAddr) *net.UDPConn {
	if addr.IP == nil {
		addr.IP = net.ParseIP(param.Localhost)
		log.Info().Msgf("Using local address: %s", addr.IP)
		return nil
	}

	if addr.Port == 0 {
		addr.Port = param.GameserverPort
		log.Info().Msgf("Using default port: %d", addr.Port)
		return nil
	}

	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to UDP")
		return nil
	}

	return conn
}

func NewGameserverClient(bufferSize int, addr *net.UDPAddr, gameserverConn NetConn, incomingMessages chan events.Event, outgoingMessages chan events.Event) *GameserverClient {
	if bufferSize == 0 {
		bufferSize = param.BufferSize
	}

	if addr == nil {
		addr = &net.UDPAddr{
			IP:   net.ParseIP(param.Localhost),
			Port: param.ClientPort,
		}
		log.Info().Msgf("Using local address: %s", addr.IP)
	} else {
		if addr.IP == nil {
			addr.IP = net.ParseIP(param.Localhost)
			log.Info().Msgf("Using local address: %s", addr.IP)
			return nil
		}

		if addr.Port == 0 {
			addr.Port = param.ClientPort
			log.Info().Msgf("Using default port: %d", addr.Port)
			return nil
		}
	}

	if gameserverConn == nil {
		gameserverConn = NewUDPConn(*addr)
	}

	return &GameserverClient{
		addr:             *addr,
		conn:             gameserverConn,
		incomingMessages: incomingMessages,
		outgoingMessages: outgoingMessages,
	}
}

func (g *GameserverClient) GetPlayerCount() int {
	event := events.Event{
		Type: events.PLAYER_COUNT,
	}

	g.SendMessage(&event)

	res := g.GetMessage()
	if res == nil {
		return -1
	}

	count, err := strconv.Atoi(res.Message)
	if err != nil {
		return -1
	}

	return count
}

func (g *GameserverClient) GetPlayers() map[string]*snake.Snake {
	event := events.Event{
		Type: events.GET_PLAYERS,
	}

	g.SendMessage(&event)

	res := g.GetMessage()
	if res == nil {
		return nil
	}

	return makeSnakes(res.Message)
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
			log.Error().Err(err).Msg("Error parsing snake response")
		}

		snakes = append(snakes, s)
	}

	return nil
}
