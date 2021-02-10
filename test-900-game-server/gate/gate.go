package gate

import (
	"fmt"
	"net"
	"strconv"
	"study-go/test-900-game-server/conf"
	"log"
	"io"
)

type Gate struct {
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	fmt.Println("in broadcaster....")
	//clients := make(map[client]bool)

	for {
		fmt.Println("in broadcaster for...")
		select {
		case msg := <-messages:
			//for cli := range clients {
			//	cli <- msg
			//}
			fmt.Println(msg)
			//case cli := <-entering:
			//	clients[cli] = true
			//case cli := <-leaving:
			//	delete(clients, cli)
			//	close(cli)
		}
	}
}

var AgentConn net.Conn

func ReadMsg(conn net.Conn) ([]byte, error) {
	//var bufMsgLen [4]byte
	//bufMsgLen := b[:2]
	bufMsgLen := make([]byte, 2)

	fmt.Println(bufMsgLen)

	// read len
	if _, err := io.ReadFull(conn, bufMsgLen); err != nil {
		return nil, err
	}

	fmt.Println(bufMsgLen)

	if _, err := io.ReadFull(conn, bufMsgLen); err != nil {
		return nil, err
	}

	fmt.Println(bufMsgLen)

	return nil, nil
}

func HandleConn(conn net.Conn) {

	fmt.Printf("in gate..." + conn.RemoteAddr().String() + "\n")

	//who := conn.RemoteAddr().String()

	//input := bufio.NewScanner(conn)
	//
	//for input.Scan() {
	//	messages <- who + ":" + input.Text()
	//	//go fmt.Println(input.Text())
	//}

	ReadMsg(conn)


	io.Copy(conn, conn)
	fmt.Println("conn close in gate...")
	conn.Close()
}

func ConnectAgent() (err error) {
	c, err := conf.GetAgentAddr()

	if err != nil {
		return
	}

	addr := c.Servers[0].Ip + ":" + strconv.Itoa(c.Servers[0].Port)

	//fmt.Println(addr)

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatal(err)
		return

	}

	AgentConn = conn

	return

}
