package main

import (
	"fmt"
	"os"
)

func main2() {
	/*打开一个文件，若不存在则创建，
	O_TRUNC:若存在则写入全新的内容，
	O_APPEND:若存在则追加内容
	O_WRONLY新文件可读可写不可执行，
	0666权限
	*/
	file, err := os.OpenFile("b.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()

	content := "shuai\n"
	n, err := file.Write([]byte(content)) //将字符串转换为切片
	if err != nil {
		fmt.Println("写完文件失败", err)
	} else {
		fmt.Println("成功向文件写入", n, "个字节")
	}
}
