package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx = context.Background()

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Connect() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func CloseClient() {
	client.Close()
}

func JsonStuff() {
	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	Set("id1234", string(json))

	val, err := Get("id1234")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}

func ExampleClient() {
	Set("key", "value")
	val, err := Get("key")
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val)

	val2, err := Get("missing_key")
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val2)
	}

	RPush("queue", "1")
	RPush("queue", "2")
	RPush("queue", "3")
	LPop("queue")
	LPop("queue")
	LPop("queue")
	LPop("queue")
}
