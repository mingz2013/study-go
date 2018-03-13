package client

import (
	"net"
	"fmt"
	"io"
	"os"
)

func HandleConn(conn net.Conn) {
	fmt.Printf(conn.RemoteAddr().String() + "\n")
	b := []byte(`hello`)
	conn.Write(b)
	io.Copy(os.Stdout, conn)
	conn.Close()
}

