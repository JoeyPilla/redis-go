package redis

import (
	"fmt"
	"time"
)

func Set(key, value string) {
	err := client.Set(ctx, key, value, time.Hour*1).Err()

	if err != nil {
		fmt.Println(err)
	}
}

func SetWithExpiration(key, value string, exp time.Duration) {
	err := client.Set(ctx, key, value, exp).Err()

	if err != nil {
		fmt.Println(err)
	}
}

func Increment(key string) interface{} {
	result, err := client.Incr(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}

	return result
}

func RPush(arr, val string) {
	if err := client.RPush(ctx, arr, val).Err(); err != nil {
		panic(err)
	}
}

func LPop(arr string) {
	result, err := client.BLPop(ctx, 1*time.Second, arr).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(result[0], result[1])
}
