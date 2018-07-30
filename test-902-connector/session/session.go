package session

import (
	"net"
)

type Session struct {
	conn   net.Conn
	userId int
}

func (s *Session) Send() {

}

func (s *Session) Recv() {

}

func (s *Session) SendMsg() {

}

func (s *Session) RecvMsg() {

}


