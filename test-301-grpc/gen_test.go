package test_301_grpc

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"study-go/test-301-grpc/pb"
	"testing"
)

func TestProto(t *testing.T) {
	info := &pb.Info{}
	info.UID = 10
	info.Power = 20
	info.StartedAt = 30
	info.OutputMoney = 40

	b, err := proto.Marshal(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)

	info2 := &pb.Info{}
	err = proto.Unmarshal(b, info2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info2)

}
