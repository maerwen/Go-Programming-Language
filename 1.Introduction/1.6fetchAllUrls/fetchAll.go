package main
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	/* "os" */
	"time"
)
func main(){//并发获取url并报告他们的时间和大小
	start := time.Now()
	ch := make(chan string)
	go fetch("http://www.baidu.com/index.html",ch)
	fmt.Println(<-ch)
	go fetch("https://aliqin.tmall.com/",ch)
	fmt.Println(<-ch)
	/* 
	for _,url := range os.Args[1:] {
		// 启动一个goroutine
		go fetch(url,ch)
	}
	for range os.Args[1:] {
		// 从通道ch接收
		fmt.Println(<-ch)
	} */
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}
func fetch(url string,ch chan<- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 发送到通道ch
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// 防止资源泄露
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}