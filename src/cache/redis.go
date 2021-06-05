package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
	"time"
)

// https://www.cnblogs.com/itbsl/p/14198111.html

type RedisOps struct {
	rdb *redis.Client
}

func init() {
}

func NewRedisOps() *RedisOps {
	redisOps := RedisOps{}
	// 通过 redis.NewClient 函数即可创建一个 redis 客户端, 这个方法接收一个 redis.Options 对象参数, 通过这个参数, 我们可以配置 redis 相关的属性, 例如 redis 服务器地址, 数据库名, 数据库密码等。
	redisOps.rdb = redis.NewClient(&redis.Options{
		Addr:     "39.108.118.123:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
		PoolSize: 5,
	})
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := redisOps.rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis init success")
	return &redisOps
}

func (r *RedisOps) Set(key string, value interface{}, duration time.Duration) {
	status := r.rdb.Set(key, value, duration)
	fmt.Println(status)
}
func (r *RedisOps) Get(key string, k reflect.Kind) interface{} {
	status := r.rdb.Get(key)
	fmt.Println(status)
	var result interface{}
	if status.Err() != nil {
		return nil
	}
	if k == reflect.String {
		result = status.Val()
	}
	return result
}
