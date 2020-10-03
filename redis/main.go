package redis

import (
	"github.com/go-redis/redis"
)

var client redis.Client

func Connect() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}
