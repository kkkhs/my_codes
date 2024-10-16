package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	//1. HSet 根据key和field字段设置，field字段的值
	err := rdb.HSet(ctx, "user1", "name", "khs").Err()
	rdb.HSet(ctx, "user1", "age", 18)
	if err != nil {
		panic(err)
	}

	//2. HGet 根据key和field字段，查询field字段的值
	value, _ := rdb.HGet(ctx, "user1", "name").Result()
	fmt.Println("user1-Name:", value)

	//3. HGetAll 根据key查询所有字段和值
	value1, _ := rdb.HGetAll(ctx, "user1").Result()
	fmt.Printf("type of value1 :%T \n ", value1)
	fmt.Println("value1 :", value1)

	//4. HIncrBy 根据key和field字段，累加字段的数值
	rdb.HIncrBy(ctx, "user1", "age", 1)
	value2, _ := rdb.HGetAll(ctx, "user1").Result()
	fmt.Println("value1 :", value2)

	//5. HKeys 根据key返回所有字段名
	keys, _ := rdb.HKeys(ctx, "user1").Result()
	fmt.Println("keys :", keys)

	//6. HLen 根据key，查询hash的字段数量
	num, _ := rdb.HLen(ctx, "user1").Result()
	fmt.Println("Hlen :", num)

	//7. HMGet 根据key和多个字段名，批量查询多个hash字段值
	r1, _ := rdb.HMGet(ctx, "user1", "name", "age").Result()
	fmt.Printf("type of r1 = %T \n", r1)
	fmt.Println("r1 =", r1)

	//8. HMSet 根据key和多个字段名和字段值，批量设置hash字段值
	data := make(map[string]interface{}) // 初始化hash数据的多个字段值
	data["1"] = "hello"
	data["2"] = "bye"

	rdb.HMSet(ctx, "data", data)
	r2, _ := rdb.HGetAll(ctx, "data").Result()
	fmt.Println("r2 :", r2)

	//9. HSetNX 如果field字段不存在，则设置hash字段值
	rdb.HSetNX(ctx, "data", "name", "khs")
	r3, _ := rdb.HGetAll(ctx, "data").Result()
	fmt.Println("r2 :", r3)

	//10. HDel 根据key和字段名，删除hash字段，支持批量删除hash字段

	rdb.HDel(ctx, "data", "name", "1")
	r4, _ := rdb.HGetAll(ctx, "data").Result()
	fmt.Println("r4 :", r4)

	//11. HExists 检测hash字段名是否存在
	is, _ := rdb.HExists(ctx, "user1", "1").Result()
	fmt.Println(is)

}
