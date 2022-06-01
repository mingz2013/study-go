package test_301_grpc

import (
	"context"
	"study-go/test-301-grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"testing"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (g greeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("sayhello:", request.GetName())
	return &pb.HelloReply{
		Message: "Hello" + request.GetName(),
	}, nil

}

func serveGreeterServer() {
	l, err := net.Listen("tcp", "localhost")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &greeterServer{})

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

func TestGreeterServer(t *testing.T) {

	serveGreeterServer()

}

func testGreeterClient() {

	conn, err := grpc.Dial("localhost", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "name",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("greeting: %s", r.GetMessage())

}
