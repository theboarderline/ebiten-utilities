package client

import (
	"github.com/theboarderline/ebiten-utilities/snake/events"
	"net"
	"time"
)

var mockEvent = events.Event{
	Type:           events.PLAYER_INPUT,
	PlayerName:     events.ENEMY,
	InputDirection: events.DirectionUp,
}

type MockConn struct {
	ReadBuffer  []byte
	WriteBuffer []byte
}

func (c *MockConn) ReadFrom(b []byte) (int, net.Addr, error) {
	copy(b, c.ReadBuffer)
	return len(c.ReadBuffer), nil, nil
}

func (c *MockConn) Write(b []byte) (int, error) {
	c.WriteBuffer = make([]byte, len(b))
	copy(c.WriteBuffer, b)
	return len(b), nil
}

func (c *MockConn) Close() error {
	return nil
}

func (c *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}
