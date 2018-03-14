package client

import (
	"net"
	"fmt"
	"io"
	"os"
)

func HandleConn(conn net.Conn) {
	fmt.Printf("in client game ..." + conn.RemoteAddr().String() + "\n")
	b := []byte(`hello\r\n`)
	conn.Write(b)
	io.Copy(os.Stdout, conn)
	conn.Close()
}

