package main

import (
	"fmt"
	"time"
)

// 时间的格式化
const (
	DATE = "2006-01-02"
	TIME = "2006-01-02 15:04:05"
)

func main() {
	t0 := time.Now()       //time.Time数据
	fmt.Println(t0.Unix()) //时间戳 Unix

	time.Sleep(50 * time.Millisecond)

	t1 := time.Now()
	// Duration = Time - Time
	diff := t1.Sub(t0) //求时间差 Sub
	fmt.Println(diff.Milliseconds())

	//计算从t0到现在的时间 Since
	fmt.Println(time.Since(t0).Milliseconds())

	// Time = Time + Durantion
	d := time.Duration(2 * time.Second)
	t2 := t0.Add(d)
	fmt.Println(t2.Unix())

	//对时间的格式化 Format
	fmt.Println(t0.Format(DATE))
	fmt.Println(t0.Format(TIME))
	s := t0.Format(TIME)
	fmt.Println(s)

	//时间反格式化 Parse
	loc, _ := time.LoadLocation("Asia/Shanghai") //构造时区-东八区
	t3, _ := time.ParseInLocation(TIME, s, loc)
	//t3, _ := time.Parse(TIME, s, loc)
	fmt.Println(t3.Unix())

}
