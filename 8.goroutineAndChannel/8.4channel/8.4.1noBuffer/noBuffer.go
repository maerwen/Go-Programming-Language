package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 建立tcp拨号连接
	c, err := net.Dial("tcp", "localhost:8080")
	// 错误处理
	if err != nil {
		log.Fatal(err)
	}
	// 创建通道,把连接的流赋值给标准输出并利用通道发送消息
	do := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, c)
		log.Println("do")
		do <- struct{}{}
	}()
	// 把标准输入流内部信息复制给连接
	mustCopy(c, os.Stdin)
	//关闭连接
	c.Close()
	//消息同步
	<-do
}
func mustCopy(dest io.Writer, src io.Reader) { //消息传输
	_, err := io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}
}
