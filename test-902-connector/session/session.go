package session

import (
	"net"
)

type Session struct {
	conn   net.Conn
	userId int
}
