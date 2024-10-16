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

	// 1. LPush 从列表左边插入数据
	err := rdb.LPush(ctx, "city", "beijing", "shanghai", "tianjing").Err()
	if err != nil {
		panic(err)
	}

	//2. LPushX 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
	rdb.LPushX(ctx, "city", "Hunan")
	rdb.LPushX(ctx, "name", 123)

	//3. RPop 从列表的右边删除第一个数据，并返回删除的数据
	val, _ := rdb.RPop(ctx, "city").Result()
	fmt.Println(val)

	//4. RPush 从列表右边插入数据
	rdb.RPush(ctx, "city", "beijing")

	//5. RPushX 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样

	//6. LPop 从列表左边删除第一个数据，并返回删除的数据

	//7. LLen 返回列表的大小
	size, _ := rdb.LLen(ctx, "city").Result()
	fmt.Println(size)

	//8. LRange 返回列表的一个范围内的数据，也可以返回全部数据
	ss, _ := rdb.LRange(ctx, "city", 0, -1).Result()
	fmt.Println(ss)

	//9. LRem 删除列表中的数据
	rdb.LRem(ctx, "city", 1, "beijing")  // 从列表左边开始，删除100， 如果出现重复元素，仅删除1次，也就是删除第一个
	rdb.LRem(ctx, "city", -2, "beijing") // 如果存在多个100，则从列表右边开始删除2个100
	rdb.LRem(ctx, "city", 0, "beijing")  // 如果存在多个100，第二个参数为0，表示删除所有元素等于100的数据

	//10. LIndex 根据索引坐标，查询列表中的数据
	r1, _ := rdb.LIndex(ctx, "city", 2).Result()
	fmt.Println("city[2]:", r1)

	//11. LInsert 在指定位置插入数据(首个)
	rdb.LInsert(ctx, "city", "before", "shanghai", "taiwan")
	rdb.LInsert(ctx, "city", "after", "shanghai", "taiwan")
	ss, _ = rdb.LRange(ctx, "city", 0, -1).Result()
	fmt.Println(ss)

}
