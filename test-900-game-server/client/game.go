package client

import (
	"net"
	"fmt"
)

func HandleConn(conn net.Conn) {
	fmt.Printf(conn.RemoteAddr().String() + "\n")

}
