package mmdb

import (
	"github.com/go-redis/redis"
)



type RedisORM struct {
	*redis.Client
}

var RDB *RedisORM



func ConnectRedis(address, dsn string) {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})
	
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	RDB = &RedisORM{
		Client: client,
	}
}