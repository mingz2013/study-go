package network

import "net"

type Msg struct {
	lenMsgLen int
}

func NewMsg() {
	m := Msg{}
	m.lenMsgLen = 2
}

func (m *Msg) Write(conn net.Conn, data []byte) (error) {

	msgLen := len(data)

	conn.Write(data)

}
