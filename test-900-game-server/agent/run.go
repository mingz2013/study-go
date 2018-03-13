package agent

import (
	"strconv"
	"fmt"
	"log"
	"io"
	"net"
	"github.com/mingz2013/study.go/test-900-game-server/conf"
)

func HandleConn(conn net.Conn) {

	fmt.Printf(conn.RemoteAddr().String() + "\n")

	io.Copy(conn, conn)
	conn.Close()
}

func Run() {
	c, err := conf.GetAgentAddr()
	if err != nil {
		return
	}

	addr := c.Servers[0].Ip + ":" + strconv.Itoa(c.Servers[0].Port)

	fmt.Println(addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go HandleConn(conn)
	}
}
