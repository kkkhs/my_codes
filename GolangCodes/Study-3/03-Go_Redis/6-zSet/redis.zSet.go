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

	//1. ZAdd 添加一个或者多个元素到集合，如果元素已经存在则更新分数
	err := rdb.ZAdd(ctx, "key", redis.Z{Score: 2.5, Member: "khs"}, redis.Z{Score: 1, Member: "khx"}, redis.Z{Score: 3, Member: "hsh"}).Err()
	if err != nil {
		panic(err)
	}

	//2. ZCard 返回集合元素个数
	num, err := rdb.ZCard(ctx, "key").Result()
	fmt.Println("num =", num)

	//3. ZCount 统计某个分数范围内的元素个数
	// 返回： 1<分数<=5 的元素个数
	// 说明：默认第二，第三个参数是大于等于和小于等于的关系。
	// 如果加上（ 则表示大于或者小于，相当于去掉了等于关系。
	size, _ := rdb.ZCount(ctx, "key", "(1", "5").Result()
	fmt.Println("size =", size)

	//4. ZIncrBy 增加元素的分数
	rdb.ZIncrBy(ctx, "key", 2, "khx")

	//5. ZRange,ZRevRange 返回集合中某个索引范围的元素，根据分数从小到大排序
	vals, _ := rdb.ZRange(ctx, "key", 0, -1).Result()
	fmt.Println(vals)

	//6. ZRangeByScore,ZRevRangeByScore 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	val1, _ := rdb.ZRangeByScore(ctx, "key", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 0, Count: 5}).Result()
	fmt.Println("val1 =", val1)

	//8. ZRangeByScoreWithScores 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
	val2, _ := rdb.ZRangeByScoreWithScores(ctx, "key", &redis.ZRangeBy{Min: "0", Max: "10", Offset: 0, Count: 5}).Result()
	fmt.Println(val2) //map类型

	//9. ZRem 删除集合元素
	rdb.ZRem(ctx, "key", "khx")

	//10. ZRemRangeByRank 根据索引范围删除元素
	// 集合元素按分数排序，从最低分到高分，删除第0个元素到第5个元素。
	rdb.ZRemRangeByRank(ctx, "key", 0, 5)
	// 这个例子，删除最高分数的两个元素，-1代表最高分数的位置，-2第二高分，以此类推。
	rdb.ZRemRangeByRank(ctx, "key", -1, -2)

	//11.ZRemRangeByScore 根据分数范围删除元素
	// 删除范围： 2<=分数<=5 的元素
	rdb.ZRemRangeByScore(ctx, "key", "2", "5")
	// 删除范围： 2<=分数<5 的元素
	rdb.ZRemRangeByScore(ctx, "key", "2", "(5")

	//12. ZScore 查询元素对应的分数
	rdb.ZAdd(ctx, "key", redis.Z{Score: 2.5, Member: "khs"}, redis.Z{Score: 1, Member: "khx"}, redis.Z{Score: 3, Member: "hsh"})
	score, _ := rdb.ZScore(ctx, "key", "khs").Result()
	fmt.Println("score of khs :", score)

	//13. ZRank 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
	rank, _ := rdb.ZRank(ctx, "key", "khs").Result()
	fmt.Println("rank of khs :", rank)
}
