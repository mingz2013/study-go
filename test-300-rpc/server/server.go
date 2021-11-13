package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)

	rpc.RegisterName("HelloService", arith)

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	var r int
	client.Call("HelloService.Multiply", Args{1, 2}, &r)
	log.Println(r)

}
