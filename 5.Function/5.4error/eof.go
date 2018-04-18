package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//
	// waitForServer("http://fanyi.baidusadsad.com/")
	// waitForServer("http://fanyi.baidu.com/")
	waitForServer("https://chrome.google.com/webstore?utm_source=chrome-ntp-icon")
}
func eof() { //EOF测试
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF { //文件读取结束
			break
		}
		if err != nil {
			fmt.Printf("read failed:%v", err)
		}

		fmt.Printf("%q", r)
	}

}
func waitForServer(url string) error {
	//尝试连接url对应的服务器
	//所有尝试失败后返回错误
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout) //在60s内尝试重新连接
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			fmt.Printf("server %s responded successfully after %ds !\n", url, tries)
			return nil
		}
		log.Fatalf("server not responding (%s);retring...\n", err)
		time.Sleep(time.Second << uint(tries)) //指数退避策略

	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)

}
