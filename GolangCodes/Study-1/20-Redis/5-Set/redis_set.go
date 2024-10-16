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

	//1. SAdd 添加集合元素
	err := rdb.SAdd(ctx, "nums", 1, 2, 3, 4, 5, 4, 3).Err()
	if err != nil {
		panic(err)
	}

	//2. SCard 获取集合元素个数
	num, _ := rdb.SCard(ctx, "nums").Result()
	fmt.Println("num = ", num)

	//3. SIsMember 判断元素是否在集合中
	is, _ := rdb.SIsMember(ctx, "nums", 5).Result()
	fmt.Println(is)

	//4. SMembers 获取集合中所有的元素
	all, _ := rdb.SMembers(ctx, "nums").Result()
	fmt.Printf("type of all = %T\n", all)
	fmt.Println("nums:", all)

	//5. SRem 删除集合元素
	rdb.SRem(ctx, "nums", 4, 5)

	//6. SPop,SPopN 随机返回集合中的元素，并且删除返回的元素
	val, _ := rdb.SPop(ctx, "key").Result() // 随机返回集合中的一个元素，并且删除这个元素
	fmt.Println(val)
	vals, _ := rdb.SPopN(ctx, "key", 5).Result() // 随机返回集合中的5个元素，并且删除这些元素
	fmt.Println(vals)

}
