package client

import (
	"net"
	"time"
)

type NetConn interface {
	ReadFrom(buf []byte) (n int, addr net.Addr, err error)
	Write(buf []byte) (n int, err error)
	Close() error
	SetReadDeadline(t time.Time) error
}
