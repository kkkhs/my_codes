package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //当前客户端的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	client.conn = conn

	//返回对象
	return client
}

// 处理server回应的消息，直接显示到标准输出即可
func (this *Client) DealResponse() {

	//一旦client有数据，就直接copy到stout标准输出上，永久堵塞监听
	io.Copy(os.Stdout, this.conn)

	/*  等同于 >>>>>>
	for{
		buf := make()
		this.conn.Read(buf)
		fmt.Println(buf)
	}
	*/
}

func (this *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		this.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>>>>> 请输入合法范围内的数字 <<<<<<<<<")
		return false
	}
}

// 查询当前在线的用户
func (this *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}

// 私聊模式
func (this *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	this.SelectUsers()
	fmt.Println(">>>>>>>>>> 请输入聊天对象[用户名], exit退出.")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>>>>>> 请输入消息内容, exit退出")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			//消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := this.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>>>>>>> 请输入聊天内容, exit退出.")
			fmt.Scanln(&chatMsg)
		}

		this.SelectUsers()
		fmt.Println(">>>>>>>>>> 请输入聊天对象[用户名], exit退出.")
		fmt.Scanln(&remoteName)

	}
}

func (this *Client) PublicChat() {
	//提示用户输入消息
	var chatMsg string

	fmt.Println(">>>>>>>>>> 请输入聊天内容, exit退出.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//发给服务器

		//消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := this.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>>>>>> 请输入聊天内容, exit退出.")
		fmt.Scanln(&chatMsg)
	}

}

func (this *Client) UpdateName() bool {
	fmt.Println(">>>>>>>>>> 请输入用户名：")
	fmt.Scanln(&this.Name)

	sendMsg := "rename|" + this.Name + "\n"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return false
	}
	return true
}

func (this *Client) Run() {
	for this.flag != 0 {
		for this.menu() != true {

		}

		//根据不同的模式处理不同的业务
		switch this.flag {
		case 1:
			//公聊模式
			this.PublicChat()
		case 2:
			//私聊模式
			this.PrivateChat()
		case 3:
			//更新用户名
			this.UpdateName()
		}
	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认为127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认为8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>>> 连接服务器失败! <<<<<<<<<")
		return
	}

	//单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>>>>>> 连接服务器成功 <<<<<<<<<")

	//启动客户端业务
	client.Run()
}
