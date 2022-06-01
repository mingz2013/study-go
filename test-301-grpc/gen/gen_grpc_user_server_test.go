package gen

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"study-go/test-301-grpc/gen/gen"
	"testing"
	"time"
)

//UserServer  实现User服务的业务对象

type UserServer struct {
	gen.UnimplementedUserServer
}

//UserView 获取详情

func (u *UserServer) UserView(ctx context.Context, in *gen.UserViewRequest) (*gen.UserViewResponse, error) {
	//panic("implement me")
	log.Printf("receive user uid request:uid %d", in.Uid)

	return &gen.UserViewResponse{

		Err: 0,

		Msg: "success",

		Data: &gen.UserEntity{

			Name: "aaaa", Age: 28,
		},
	}, nil
}

//UserIndex 实现了User 服务接口的所有方法

func (u *UserServer) UserIndex(ctx context.Context, in *gen.UserIndexRequest) (*gen.UserIndexResponse, error) {

	log.Printf("receive user index request:page %d page_size %d", in.Page, in.PageSize)

	return &gen.UserIndexResponse{

		Err: 0,

		Msg: "success",

		Data: []*gen.UserEntity{

			{Name: "aaaa", Age: 28},

			{Name: "bbbb", Age: 1},
		},
	}, nil

}

//UserPost 提交数据

func (u *UserServer) UserPost(ctx context.Context, in *gen.UserPostRequest) (*gen.UserPostResponse, error) {

	log.Printf("receive user uid request:name %s password:%s,age:%d", in.Name, in.Password, in.Age)

	return &gen.UserPostResponse{

		Err: 0,

		Msg: "success",
	}, nil

}

//UserDelete 删除数据

func (u *UserServer) UserDelete(ctx context.Context, in *gen.UserDeleteRequest) (*gen.UserDeleteResponse, error) {

	log.Printf("receive user uid request:uid %d", in.Uid)

	return &gen.UserDeleteResponse{

		Err: 0,

		Msg: "success",
	}, nil

}

func testService() {

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {

		log.Fatal("failed to listen", err)

	}

	//创建rpc服务

	grpcServer := grpc.NewServer()

	//为User服务注册业务实现 将User服务绑定到RPC服务器上

	gen.RegisterUserServer(grpcServer, &UserServer{})

	//注册反射服务， 这个服务是CLI使用的， 跟服务本身没有关系

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {

		log.Fatal("faild to server,", err)

	}

}

func testClient() {

	//建立链接

	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())

	if err != nil {

		log.Fatal("did not connect", err)

	}

	defer conn.Close()

	userClient := gen.NewUserClient(conn)

	//设定请求超时时间 3s

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	//UserIndex 请求

	userIndexResponse, err := userClient.UserIndex(ctx, &gen.UserIndexRequest{

		Page: 1,

		PageSize: 12,
	})

	if err != nil {

		log.Printf("user index could not greet: %v", err)

	}

	if 0 == userIndexResponse.Err {

		log.Printf("user index success: %s", userIndexResponse.Msg)

		// 包含 UserEntity 的数组列表

		userEntityList := userIndexResponse.Data

		for _, row := range userEntityList {

			fmt.Println(row.Name, row.Age)

		}

	} else {

		log.Printf("user index error: %d", userIndexResponse.Err)

	}

	// UserView 请求

	userViewResponse, err := userClient.UserView(ctx, &gen.UserViewRequest{Uid: 1})

	if err != nil {

		log.Printf("user view could not greet: %v", err)

	}

	if 0 == userViewResponse.Err {

		log.Printf("user view success: %s", userViewResponse.Msg)

		userEntity := userViewResponse.Data

		fmt.Println(userEntity.Name, userEntity.Age)

	} else {

		log.Printf("user view error: %d", userViewResponse.Err)

	}

	// UserPost 请求

	userPostReponse, err := userClient.UserPost(ctx, &gen.UserPostRequest{Name: "big_cat", Password: "123456", Age: 29})

	if err != nil {

		log.Printf("user post could not greet: %v", err)

	}

	if 0 == userPostReponse.Err {

		log.Printf("user post success: %s", userPostReponse.Msg)

	} else {

		log.Printf("user post error: %d", userPostReponse.Err)

	}

	// UserDelete 请求

	userDeleteReponse, err := userClient.UserDelete(ctx, &gen.UserDeleteRequest{Uid: 1})

	if err != nil {

		log.Printf("user delete could not greet: %v", err)

	}

	if 0 == userDeleteReponse.Err {

		log.Printf("user delete success: %s", userDeleteReponse.Msg)

	} else {

		log.Printf("user delete error: %d", userDeleteReponse.Err)

	}

}

func TestService(t *testing.T) {
	testService()

}
