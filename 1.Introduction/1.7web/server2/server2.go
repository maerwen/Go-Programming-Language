package main
import (
	"fmt"
	"log"
	"net/http"
	"sync"
)
// 迷你回声和计数器服务器
var mu sync.Mutex//加锁
var count int
func main(){
	// http://localhost:8080/× 查看地址
	// http://localhost:8080/count 查看访问次数，在增加
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}
// 处理程序回显请求用的url部分
func handler(w http.ResponseWriter,r *http.Request){
	mu.Lock()//确保最多只有一个goroutine在同一时间访问变量
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}
// 回显目前为止的调用次数
func counter(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	fmt.Fprintf(w,"Count %d \n",count)
	mu.Unlock()
}