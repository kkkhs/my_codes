package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听.....")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listen.Close()
	//循环等待客户端连接
	for {
		fmt.Println("等待客户端连接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error:", err)
			//return 一个错误不需要return
		} else {
			fmt.Printf("Accept suc conn=%v 客户端IP=%v", conn, conn.RemoteAddr().String())
			break
		}
		//此处启动一个协程为客户端服务

	}
	//fmt.Printf("listen=%v\n", listen)

}
