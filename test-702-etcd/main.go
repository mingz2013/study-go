package test_702_etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func createClient() *clientv3.Client {

	//类型中的成员是etcd客户端几何核心功能模块的具体实现，它们分别用于：
	//
	//Cluster：向集群里增加etcd服务端节点之类，属于管理员操作。
	//KV：我们主要使用的功能，即K-V键值库的操作。
	//Lease：租约相关操作，比如申请一个TTL=10秒的租约（应用给key可以实现键值的自动过期）。
	//Watcher：观察订阅，从而监听最新的数据变化。
	//Auth：管理etcd的用户和权限，属于管理员操作。
	//Maintenance：维护etcd，比如主动迁移etcd的leader节点，属于管理员操作。

	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	return cli

}

func put() {
	// crud + watch

	cli := createClient()
	timeout := time.Duration(5)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	cli.Put(ctx, "sample_key", "value")
	cancel()

	//cli := createClient()
	kv := clientv3.NewKV(cli)

	kv.Get(context.TODO(), "/aaa")

	kv.Delete(ctx, "/aaa")

	watch := cli.Watch(ctx, "aaa")
	for ws := range watch {
		fmt.Println("watch ws %d", ws)
	}

	// 租赁时长60s，
	cli.Grant(context.TODO(), 60)

}

func get() {

}
