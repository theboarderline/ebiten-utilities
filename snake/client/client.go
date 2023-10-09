package client

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"github.com/theboarderline/ebiten-utilities/snake/object/snake"
	"github.com/theboarderline/ebiten-utilities/snake/param"
	"net"
)

type GameserverClient struct {
	bufferSize       int
	conn             NetConn
	addr             net.UDPAddr
	incomingMessages chan events.Event
	outgoingMessages chan events.Event
}

func NewGameserverConn(addr net.UDPAddr) *net.UDPConn {
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

func NewGameserverClient(bufferSize int, clientAddr *net.UDPAddr, gameserverConn NetConn, incomingMessages chan events.Event, outgoingMessages chan events.Event) *GameserverClient {
	if bufferSize == 0 {
		bufferSize = param.BufferSize
	}

	if clientAddr == nil {
		clientAddr = &net.UDPAddr{
			IP:   net.ParseIP(param.Localhost),
			Port: param.ClientPort,
		}
		log.Info().Msgf("Using local address: %s", clientAddr.IP)
	} else {
		if clientAddr.IP == nil {
			clientAddr.IP = net.ParseIP(param.Localhost)
			log.Info().Msgf("Using local address: %s", clientAddr.IP)
			return nil
		}

		if clientAddr.Port == 0 {
			clientAddr.Port = param.ClientPort
			log.Info().Msgf("Using default port: %d", clientAddr.Port)
			return nil
		}
	}

	if gameserverConn == nil {
		gameserverConn = NewGameserverConn(net.UDPAddr{
			IP:   net.ParseIP(param.Localhost),
			Port: param.GameserverPort,
		})
	}

	return &GameserverClient{
		bufferSize:       bufferSize,
		addr:             *clientAddr,
		conn:             gameserverConn,
		incomingMessages: incomingMessages,
		outgoingMessages: outgoingMessages,
	}
}

func (g *GameserverClient) GetPlayerCount() {
	event := events.Event{
		Type: events.PLAYER_COUNT,
	}

	g.SendMessage(&event)

	return
}

func (g *GameserverClient) GetPlayers(senderName string) {
	event := events.Event{
		Type:       events.GET_PLAYERS,
		PlayerName: senderName,
	}

	g.SendMessage(&event)
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

	if err := json.Unmarshal([]byte(response), &snakes); err != nil {
		log.Error().Err(err).Msg("Error parsing snake response")
		return nil
	}

	return nil
}
