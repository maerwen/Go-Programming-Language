package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			log.Fatal(err)
			continue //跳过该链接
		}
		// handleConn(c)//单链接处理
		go handleConn(c) //并发
	}
}

//顺序时钟tcp服务器,每秒向客户端发送时间
func handleConn(c net.Conn) {
	defer c.Close() //无论如何,最终关闭连接
	for {
		_, err := io.WriteString(c, time.Now().Format("\r15:04:05")) //以设定格式将日期写出到连接
		if err != nil {
			return //断开连接
		}
		time.Sleep(1 * time.Second)
	}
}
