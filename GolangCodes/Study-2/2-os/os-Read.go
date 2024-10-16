package main

import (
	"fmt"
	"os"
)

func main1() {
	//打开文件
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	//关闭文件
	defer file.Close()

	//读文件到content， 默认为100个0
	content := make([]byte, 100, 100)
	//读取n个字节
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	//转换为字节
	fmt.Println(string(content[:n]))

}
