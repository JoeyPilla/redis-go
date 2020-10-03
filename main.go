package main

import (
	"fmt"

	"github.com/JoeyPilla/redis-go/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	redis.Connect()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
