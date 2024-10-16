package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	//
	ctx := context.Background() //创建一个上下文
	err := rdb.Set(ctx, "goredistest", "goredistestvalue", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "goredistest").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("goredistestkey:", val)

	//2、使用Do
	result, err := rdb.Do(ctx, "get", "goredistest").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Do goredistestkey:", result.(string)) //强制转为string

}
