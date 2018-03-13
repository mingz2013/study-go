package client

import (
	"net"
	"strconv"
	"log"
)

func ConnectGate(ip string, port int) (conn net.Conn, err error) {
	addr := ip + ":" + strconv.Itoa(port)
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
