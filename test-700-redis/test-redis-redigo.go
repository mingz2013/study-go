package test_700_redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"encoding/json"
)

func TestRedigo() {

	host := "localhost:6379"
	rs, err := redis.Dial("tcp", host) // tcp连接redis

	if err != nil {
		log.Println(err)
		return
	}

	defer rs.Close() // 操作完毕后自动关闭

	log.Println("conn success...")

	key := "test-key"
	// 操作redis时调用Do方法，第一个参数传入操作名称（string），然后根据不同操作传入key，value，数字等，返回2个参数，第一个为操作标识，成功则为1，失败则为0；第二个为错误信息
	reply, err := rs.Do("GET", key)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("get key", key, reply)

	// 若value的类型为int，则用redis.Int转换
	// string redis.String
	// json   redis.Byte
	value, err := redis.String(reply, err)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(value)

	// 存入json数据
	key = "test-key-2"
	imap := map[string]string{"key1": "111", "key2": "222"}
	// 将map转换成json数据
	value1, _ := json.Marshal(imap)
	// 存入redis
	n, err := rs.Do("SETNX", key, value1)
	if err != nil {
		log.Println(err)
	}
	if n == int64(1) {
		log.Println("setnx key success", key, value1)
	}

	// 取json数据
	// 先声明imap用来装数据
	var imap1 map[string]string
	// json数据在go中是[]byte类型，所以此处用redis.Bytes转换
	value2, err := redis.Bytes(rs.Do("GET", key))
	if err != nil {
		log.Println(err)
	}

	// 将json解析成map类型
	errShal := json.Unmarshal(value2, &imap1)
	if errShal != nil {
		log.Println(errShal)

	}

	log.Println("get imap1", imap1)

}
