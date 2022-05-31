package test_703_zookeeper

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"os"
	"sync"
	"time"
)

func Connect() *zk.Conn {
	conn, _, err := zk.Connect([]string{"172.17.0.2", "172.17.0.3", "172.17.0.4"}, time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn
}

func create() {

	conn := Connect()

	// 创建持久节点
	path, err := conn.Create("/hello", []byte("world"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	println("Created", path)

	// 创建临时节点，创建此节点的会话结束后立即清除此节点
	ephemeral, err := conn.Create("/ephemeral", []byte("1"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	println("Ephemeral node created:", ephemeral)

	// 创建持久时序节点
	sequence, err := conn.Create("/sequence", []byte("1"), zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	println("Sequence node created:", sequence)

	// 创建临时时序节点，创建此节点的会话结束后立即清除此节点
	ephemeralSequence, err := conn.Create("/ephemeralSequence", []byte("1"), zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	println("Ephemeral-Sequence node created:", ephemeralSequence)
}

func get() {
	conn := Connect()

	result, state, err := conn.Get("/hello")
	if err != nil {
		panic(err)
	}
	fmt.Println("result: ", string(result))
	fmt.Println("state ->")
	fmt.Printf("cZxid=%d\nctime=%d\nmZxid=%d\nmtime=%d\npZxid=%d\ncversion=%d\ndataVersion=%d\naclVersion=%d\nephemeralOwner=%v\ndataLength=%d\nnumChildren=%d\n", state.Czxid, state.Ctime, state.Mzxid, state.Mtime, state.Pzxid, state.Cversion, state.Version, state.Aversion, state.EphemeralOwner, state.DataLength, state.NumChildren)
}

func set() {
	conn := Connect()

	path := "/hello"
	_, state, _ := conn.Get(path)

	state, err := conn.Set(path, []byte("girl"), state.Version)
	if err != nil {
		panic(err)
	}
	fmt.Println("state ->")
	fmt.Printf("cZxid=%d\nctime=%d\nmZxid=%d\nmtime=%d\npZxid=%d\ncversion=%d\ndataVersion=%d\naclVersion=%d\nephemeralOwner=%v\ndataLength=%d\nnumChildren=%d\n", state.Czxid, state.Ctime, state.Mzxid, state.Mtime, state.Pzxid, state.Cversion, state.Version, state.Aversion, state.EphemeralOwner, state.DataLength, state.NumChildren)

	data, _, err := conn.Get(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nnew value: ", string(data))
}

func delete() {
	conn := Connect()

	path := "/hello"
	exists, state, err := conn.Exists(path)
	fmt.Printf("\npath[%s] exists: %v\n", path, exists)

	err = conn.Delete(path, state.Version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("path[%s] is deleted.", path)

	exists, _, err = conn.Exists(path)
	fmt.Printf("\npath[%s] exists: %v\n", path, exists)
}

func acl() {
	conn := Connect()
	// get acl
	acl, state, err := conn.GetACL("/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("\nget acl:")
	fmt.Println("scheme =", acl[0].Scheme)
	fmt.Println("id =", acl[0].ID)
	fmt.Println("permissions =", acl[0].Perms)

	// set acl
	perms := zk.PermCreate | zk.PermRead | zk.PermWrite | zk.PermAdmin // crwa 权限
	state, err = conn.SetACL("/test", zk.WorldACL(int32(perms)), state.Version)
	if err != nil {
		panic(err)
	}
	fmt.Println("SetAcl successful.")

	// create child node
	_, err = conn.Create("/test/1", []byte("1"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}

	// get child node
	_, state, err = conn.Get("/test/1")
	if err != nil {
		panic(err)
	}

	// delete child node /1
	err = conn.Delete("/test/1", state.Version)
	if err != nil {
		fmt.Println("delete failed: ", err.Error())
		os.Exit(1)
	}
}

func callback(e zk.Event) {
	fmt.Println("++++++++++++++++++++++++")
	fmt.Println("path:", e.Path)
	fmt.Println("type:", e.Type.String())
	fmt.Println("state:", e.State.String())
	fmt.Println("------------------------")
}

func watch() {
	eventCallbackOption := zk.WithEventCallback(callback)

	conn, _, err := zk.Connect([]string{"172.17.0.2", "172.17.0.3", "172.17.0.4"}, time.Second, eventCallbackOption)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 注册一个 watch
	exists, state, _, err := conn.ExistsW("/watch")
	if err != nil {
		panic(err)
	}

	if !exists {
		// 创建 /watch 时，触发监听事件，watch 失效
		_, err = conn.Create("/watch", []byte("watch"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
		if err != nil {
			panic(err)
		}

		// 再注册一个 watch
		_, state, _, err = conn.ExistsW("/watch")
		if err != nil {
			panic(err)
		}
	}

	// 删除 /watch 时，触发监听事件，watch 失效
	err = conn.Delete("/watch", state.Version)
	if err != nil {
		panic(err)
	}
}

func lock() {
	conn, _, err := zk.Connect([]string{"172.17.0.2", "172.17.0.3", "172.17.0.4"}, time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			lock := zk.NewLock(conn, "/root/lock", zk.WorldACL(zk.PermAll))
			err = lock.LockWithData([]byte("it is a lock"))
			if err != nil {
				panic(err)
			}
			fmt.Println("第", n, "个 goroutine 获取到了锁")
			time.Sleep(time.Second) // 1 秒后释放锁

			lock.Unlock()
		}(i)
	}

	wg.Wait()
}
