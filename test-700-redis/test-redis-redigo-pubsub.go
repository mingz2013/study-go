package test_700_redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"log"
	"sync"
)

var (
	RedisPool  *redis.Pool
	REDIS_HOST string
	REDIS_PORT string
	REDIS_DB   int

	wg sync.WaitGroup
)

func init() {
	// 这里可从配置中读取
	REDIS_HOST = "localhost"
	REDIS_PORT = "6379"
	REDIS_DB = 0

	RedisPool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   1024,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", REDIS_HOST+":"+REDIS_PORT)
			if err != nil {
				return nil, err
			}

			// select db

			c.Do("SELECT", REDIS_DB)

			return c, nil
		},
	}
}

func Subscribe() {

	redisChannel := "redChatRoom"
	c := RedisPool.Get()
	psc := redis.PubSubConn{c}
	psc.Subscribe(redisChannel)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {

			psc.Unsubscribe(redisChannel)
			psc.Close()
		}()


		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				log.Println("messages<"+v.Channel+">:", string(v.Data))
			case redis.Subscription:
				log.Println(v.Channel, v.Kind, v.Count)
				continue
			case error:
				log.Println("error", v)
				return

			}
		}

	}()

}

func Pubscribe(s string) {
	//defer wg.Done()
	log.Println("pub msg", s)
	redisChannel := "redChatRoom"
	c := RedisPool.Get()

	defer func() {
		log.Println("close...")
		c.Close()
	}()

	_, err := c.Do("PUBLISH", redisChannel, s)
	if err != nil {
		log.Println("pub err:", err)
		return
	}

	log.Println("pub over...")

}

func TestRedisPubSub() {

	Subscribe()

	//wg.Add(1)
	Pubscribe("hehe")
	Pubscribe("haha")

	wg.Wait()

	ch := make(chan int)
	<-ch

}
