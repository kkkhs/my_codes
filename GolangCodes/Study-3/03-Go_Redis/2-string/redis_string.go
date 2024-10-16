package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	//基本使用
	err := rdb.Set(ctx, "1", "khs", 0).Err()
	if err != nil {
		panic(err)
	}
	value, err := rdb.Get(ctx, "1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key : 1, value:", value)
	result, err := rdb.Do(ctx, "get", "1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key : 1, value:", result)

	//GetSet 设置一个key的值，并返回这个key的旧值
	oldvalue, err := rdb.GetSet(ctx, "1", "shuange").Result()
	if err != nil {
		panic(err)
	}
	newvalue, _ := rdb.Do(ctx, "get", "1").Result()
	fmt.Println("key : 1, oldvalue:", oldvalue)
	fmt.Println("key : 1, newvalue:", newvalue)

	//SetNX 如果key不存在,则设置这个Key的值
	rdb.SetNX(ctx, "2", "khx", 0)
	rdb.SetNX(ctx, "1", "khx", 0)
	value, _ = rdb.Get(ctx, "1").Result()
	fmt.Println("key : 1, value :", value) //若key已存在，则无效
	value, _ = rdb.Get(ctx, "2").Result()
	fmt.Println("key : 2, value :", value)

	//MGet 批量查询
	vals, _ := rdb.MGet(ctx, "1", "2").Result()
	fmt.Printf("type of vals =%T\n", vals) //  []interface{}类型
	fmt.Println("vals =", vals)

	//MSet 批量设置
	rdb.MSet(ctx, "key1", "value1", "key2", "value2")
	vals, _ = rdb.MGet(ctx, "key1", "key2").Result()
	fmt.Println("vals =", vals)

	//Incr, IncrBy， IncrByFloat 针对一个key的value进行自增操作
	//Decr, DecrBy	 针对一个key的value进行自减操作
	val, _ := rdb.Incr(ctx, "key").Result()
	fmt.Println(val) //1, 2, 3

	val, _ = rdb.IncrBy(ctx, "key", 2).Result()
	fmt.Println(val) // 4, 6

	res, _ := rdb.DecrBy(ctx, "key", 2).Result()
	fmt.Println(res)

	//Del 删除
	rdb.Del(ctx, "1", "key")
	value, _ = rdb.Get(ctx, "1").Result()
	fmt.Println(value)

	//Expire
	rdb.Expire(ctx, "2", 10*time.Second)
}
