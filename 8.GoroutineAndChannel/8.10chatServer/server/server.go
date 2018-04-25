package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//这是一个聊天服务器程序
func main() {
	server()
}
func server() { //两个进程,一个主进程,一个广播进程,还会为每个客户端创建一个或多个进程
	// 监听某一地址某一端口某一格式的请求
	listener, err := net.Listen("tcp", "localhost:8080")
	// 日常错误处理
	if err != nil {
		log.Fatal(err)
	}
	// 广播程序goroutine
	go broadcaster()
	// 程序持续运行,不断监听并处理来自各个客户端的请求,当有一个客户端出错,打印错误日志并跳过该客户端
	for {
		con, e := listener.Accept()
		if e != nil {
			log.Fatal(e)
			continue
		}
		go handleCon(con)
	}
}

type client chan string

//三个通道:
var enter = make(chan client) //上线发送该用户发送消息通道
var leave = make(chan client) //下线发送该用户发送消息通道
var messages = make(chan string)

func broadcaster() {
	// 在线用户客户端通道map
	clients := make(map[client]bool)
	// 持续在线监听接收通道消息
	for {
		select {
		// 一台客户端发送消息到服务器,服务器收到后发送到公屏
		case msg := <-messages:
			for client := range clients {
				client <- msg
			}
		case client := <-enter: //存储用户消息通道
			clients[client] = true
		case client := <-leave: //移除用户消息通道
			delete(clients, client)
			close(client)
		}
	}

}
func handleCon(con net.Conn) {
	// 创建单个用户消息通道
	ch := make(chan string)
	go clientWriter(con, ch)
	// 根据链接获取客户端地址
	who := con.RemoteAddr().String()
	ch <- "你是" + who + "!"   //写给自己
	messages <- who + "上线了!" //发给服务器
	enter <- ch              //将消息通道发给服务器
	// 从连接读取内容
	input := bufio.NewScanner(con)
	for input.Scan() {
		messages <- who + ":\t" + input.Text() //把内容发给客户端
	}
	leave <- ch
	messages <- who + "下线了!"
	con.Close()

}
func clientWriter(con net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(con, msg)
	}
}
