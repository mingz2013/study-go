package test_301_grpc

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"log"
	"study-go/test-301-grpc/pb"
	"testing"
	"time"
)

func TestGrpcEtcd(t *testing.T) {
	// 创建etcd客户端
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	if cerr != nil {
		log.Fatalln(cerr)
	}

	go func() {

		// 创建endpoints管理
		em, err := endpoints.NewManager(cli, "foo/bar/my-service")
		if err != nil {
			log.Fatalln(err)
		}
		// 添加节点, 可设置租约
		err = em.AddEndpoint(
			context.TODO(),
			"foo/bar/my-service/e1",
			endpoints.Endpoint{Addr: "localhost"},
		)

		if err != nil {
			log.Fatalln(err)
		}

		defer func() {
			// 删除节点
			err = em.DeleteEndpoint(context.TODO(), "foo/bar/my-service/e1")
			if err != nil {
				log.Fatalln(err)
			}

			////一次修改多个
			//em.Update(context.TODO(), []*endpoints.UpdateWithOpts{
			//
			//})
		}()

		//  启动服务
		serveGreeterServer()

	}()

	// 创建resolover
	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		log.Fatalln(err)
	}

	// 创建grpc连接, 使用etcd resolver，并配置balancer 策略

	grpcDialCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, gerr := grpc.DialContext(
		grpcDialCtx,
		"etcd:///foo/bar/my-service",

		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if gerr != nil {
		log.Fatalln(gerr)
	}

	defer conn.Close()

	// grpc 服务的 greeter 客户端
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
