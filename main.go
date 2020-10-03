package main

import (
	"fmt"
	"time"

	"github.com/JoeyPilla/redis-go/redis"
)

func main() {
	redis.Connect()
	defer redis.CloseClient()

	redis.Set("name", "joey")
	val, err := redis.Get("name")

	fmt.Println(val, err)
	time.Sleep(time.Second * 1)
	val, err = redis.Get("name")
	fmt.Println(val, err)
	redis.JsonStuff()
	redis.ExampleClient()
}
