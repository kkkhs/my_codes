package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("b.txt")
	if err != nil {
		fmt.Println("打开文件失败：", err)
	}
	defer file.Close()

	//带缓冲的读取数据，减少磁盘的损耗，提高程序的性能
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //传入字节参数表示遇到'\n'停一下
		if err != nil {
			if err == io.EOF { //当本行为末行：End Of File
				fmt.Print(line) //break之前打印出最后一行的内容
				break
			} else {
				fmt.Println("读文件发生错误", err)
				return
			}
		} else {
			fmt.Print(line) //line中末尾已有'\n'
		}
	}
}
