package game

import (
	"fmt"
	"net"
	"io"
	"strconv"
	"study-go/test-900-game-server/conf"
	"log"
)

type Game struct {
}

var AgentConn net.Conn

func HandleConn(conn net.Conn) {

	fmt.Printf(conn.RemoteAddr().String() + "\n")

	io.Copy(conn, conn)
	conn.Close()
}

func ConnectAgent() (err error) {
	c, err := conf.GetAgentAddr()

	if err != nil {
		return
	}

	addr := c.Servers[0].Ip + ":" + strconv.Itoa(c.Servers[0].Port)

	fmt.Println(addr)

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatal(err)
		return

	}

	AgentConn = conn

	return

}
