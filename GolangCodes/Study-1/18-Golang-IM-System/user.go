package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 创建一个用户的方法
func NerUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//启动监听当前user channel的消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户的上线业务
func (this *User) Online() {
	//用户上线，将用户加入到OnlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线
	this.server.BroadCast(this, "online....")
}

// 用户的下线业务
func (this *User) Offline() {
	//用户下线，将用户从OnlineMap删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线
	this.server.BroadCast(this, "offline....")
}

// 给当前用户对应的客户端发消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "is online...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式 rename|张三
		newName := strings.Split(msg, "|")[1]

		//判断name是否已存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("This Name :[" + newName + "] had been used!\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("You have successfuly renamed! The new name is: " + this.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息私聊格式 to|张三|消息内容

		//1、获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("Wrong remoteName format! Please use \"to|remoteName|content\" format\n")
			return
		}

		//2、根据用户名 得到对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("The remoteName unexist\n")
			return
		}

		//3、获取消息内容，通过对方的User对象将消息内容发送过去
		if len(strings.Split(msg, "|")) > 2 {
			content := strings.Split(msg, "|")[2]
			if content == "" {
				this.SendMsg("No content!Please send again\n")
				return
			}
			remoteUser.SendMsg("[" + this.Name + "] said to you: " + content + "\n")
		}
	} else {
		this.server.BroadCast(this, msg)
	}

}

// 监听当前User channel的方法，一旦有消息，就直接发送给客服端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
