package datasource

import (
	"free/conf"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

var (
	client     *redis.Client
	redis_lock sync.Mutex
)

func NewRedisClient() *redis.Client {
	if client != nil {
		return client
	}
	redis_lock.Lock()
	defer redis_lock.Unlock()
	if client != nil {
		return client
	}
	start := time.Now().UnixNano()
	conf := conf.GetRedisConf()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis connect error:%s", err)
	}
	end := time.Now().UnixNano()
	fmt.Println("redis client used")
	fmt.Println(end - start)
	return client
}
