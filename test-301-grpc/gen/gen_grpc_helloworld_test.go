package gen

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"study-go/test-301-grpc/gen/gen"
	"testing"
)

type greeterServer struct {
	gen.UnimplementedGreeterServer
}

func (g greeterServer) SayHello(ctx context.Context, request *gen.HelloRequest) (*gen.HelloReply, error) {
	log.Println("sayhello:", request.GetName())
	return &gen.HelloReply{
		Message: "Hello" + request.GetName(),
	}, nil

}

func TestGreeterServer(t *testing.T) {

	l, err := net.Listen("tcp", "localhost")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	gen.RegisterGreeterServer(s, &greeterServer{})

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}

}

func testGreeterClient() {

	conn, err := grpc.Dial("localhost", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := gen.NewGreeterClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	r, err := c.SayHello(context.Background(), &gen.HelloRequest{
		Name: "name",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("greeting: %s", r.GetMessage())

}
