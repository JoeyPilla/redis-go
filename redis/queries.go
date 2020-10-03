package redis

import "fmt"

func Get(key string) (interface{}, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error: ", err)
		return val, err
	}
	return val, nil
}
